package cidr32

import (
	"encoding/binary"
	"fmt"
	"net"
)

// ----------------------------------------------------------------------------

type Net32 struct {
	Addr    uint32
	Masklen uint8
}

type Net32List []Net32

func (r Net32) String() string {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, r.Addr)
	return fmt.Sprintf("%d.%d.%d.%d/%d", bs[0], bs[1], bs[2], bs[3], r.Masklen)
}

func (r *Net32) Empty() bool {
	return r == nil || r.Addr == 0
}

func (r *Net32) Equal(n *Net32) bool {
	return r == n || (r != nil && r.Addr == n.Addr && r.Masklen == n.Masklen)
}

func (r *Net32) IPMask() net.IPMask {
	if r == nil {
		return net.IPMask{}
	}
	return MasklenToIPMask(r.Masklen)
}

func (r *Net32) IPNet() (rv *net.IPNet) {
	return Net32toIPnet(r)
}

func (r *Net32) First32() uint32 {
	if r == nil {
		return 0
	}
	return r.Addr
}

func (r *Net32) Last32() uint32 {
	if r == nil {
		return 0
	}
	return r.First32() | ^MasklenToMask(r.Masklen)
}

func (r *Net32) IPRange() *IPRange {
	return &IPRange{i32: [2]uint32{r.First32(), r.Last32()}}
}

// Capacity -- capacity in IP addrsses (include Net-addr and Broadcast)
func (r *Net32) Capacity() int {
	if r == nil {
		return 0
	}
	return r.IPRange().Capacity()
}

// CapacityFor -- capacity in blocks (with corresponded size)
func (r *Net32) CapacityFor(masklen uint8) int {
	if r == nil {
		return 0
	}
	if masklen < r.Masklen {
		return 0
	}
	netCapacityBits := 32 - r.Masklen
	reqCapacityBits := 32 - masklen
	bitsToBlockCount := netCapacityBits - reqCapacityBits
	return 1 + (0xffffffff >> (32 - bitsToBlockCount))
}

// GetBlocks -- returns net block manipulator for corresponded blockSize
func (r *Net32) GetBlocks(blockSize uint8) *Net32Blocks {
	if r == nil || blockSize < r.Masklen || blockSize > 32 {
		// unable to allocate
		return nil
	}
	return &Net32Blocks{
		Net:       r,
		blockSize: blockSize,
	}
}

// GetFreeBlock -- returns first free net block, corresponded to blockSize. BusyList should be given
func (r *Net32) GetFreeBlock(blockSize uint8, busy Net32List) (rv *Net32) {
	if r == nil || blockSize < r.Masklen {
		// unable to allocate
		return nil
	}
	blocks := r.GetBlocks(blockSize)
	if blocks == nil {
		return nil
	}

exLoop:
	for i := 0; i < blocks.Amount(); i++ {
		block := blocks.At(i)
		blockRange := block.IPRange()
		for j := range busy {
			busyBlockRange := busy[j].IPRange()
			if blockRange.IsIntersect(busyBlockRange) {
				continue exLoop
			}
		}
		rv = block
		break
	}
	return rv
}

// ----------------------------------------------------------------------------

func MasklenToMask(ml uint8) uint32 {
	if ml > 32 {
		return 0
	}
	return ^(uint32(0xffffffff) >> ml)
}

func MasklenToIPMask(ml uint8) net.IPMask {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, MasklenToMask(ml))
	return net.IPMask(bs)
}

// IPnetToNet32 -- convert net.IPNet to Net32
func IPnetToNet32(n *net.IPNet) *Net32 {
	if n == nil {
		return nil
	}
	a, _ := n.Mask.Size()
	return &Net32{
		Addr:    IPtoUint32(n.IP),
		Masklen: uint8(a),
	}
}

// Net32toIPnet -- convert Net32 to net.IPnet
func Net32toIPnet(n *Net32) *net.IPNet {
	if n == nil {
		return nil
	}
	return &net.IPNet{
		IP:   Uint32toIP(n.Addr),
		Mask: n.IPMask(),
	}
}

// CidrToNet32 -- parse CIDR-notated string and convert one to Net32
func CidrToNet32(s string) *Net32 {
	_, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil
	}
	return IPnetToNet32(ipNet)
}

// AmountInBlock -- amount of IP addresses in block with given masklen
func AmountInBlock(masklen uint8) int {
	return int(^MasklenToMask(masklen) + 1)
}

//-----------------------------------------------------------------------------

type Net32Blocks struct {
	Net       *Net32
	blockSize uint8
}

func (r *Net32Blocks) Amount() int {
	if r == nil {
		return 0
	}
	return r.Net.CapacityFor(r.blockSize)
}

func (r *Net32Blocks) First() *Net32 {
	if r == nil {
		return nil
	}
	return &Net32{
		Addr:    r.Net.Addr,
		Masklen: r.blockSize,
	}
}

func (r *Net32Blocks) Last() *Net32 {
	if r == nil {
		return nil
	}
	lastIP := r.Net.Last32()
	startBlock := lastIP ^ uint32(0xffffffff>>r.blockSize)
	return &Net32{
		Addr:    startBlock,
		Masklen: r.blockSize,
	}
}

func (r *Net32Blocks) At(i int) *Net32 {
	if r == nil || i < 0 {
		return nil
	}
	a := AmountInBlock(r.blockSize)
	firstIP := r.Net.First32()
	startIP := firstIP + uint32(i*a)
	if startIP > r.Net.Last32() || startIP < r.Net.First32() {
		return nil
	}
	return &Net32{
		Addr:    startIP,
		Masklen: r.blockSize,
	}
}
