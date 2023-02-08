package k8sutil

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	k8types "github.com/Mirantis/mcc-api/v2/pkg/apis/util/ipam/k8sutil/types"
	"github.com/google/uuid"
	funk "github.com/thoas/go-funk"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
	"regexp"
	"strings"
	"time"
)

// IsMACaddr is MAC address checker
// +gocode:public-api=true
var IsMACaddr = regexp.MustCompile(`[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}`)

//-----------------------------------------------------------------------------
// +gocode:public-api=true
func setSomething(m map[string]string, k, v string, shouldBeChanged []bool) (map[string]string, bool) {
	var changed bool
	if m == nil {
		m = map[string]string{}
	}
	if len(shouldBeChanged) > 0 {
		changed = shouldBeChanged[0]
	}
	if tmp, ok := m[k]; ok && tmp == v {

		return m, changed
	}
	m[k] = v
	return m, true
}

// +gocode:public-api=true
func clearSomething(m map[string]string, k string, shouldBeChanged []bool) (map[string]string, bool) {
	var changed bool
	if len(shouldBeChanged) > 0 {
		changed = shouldBeChanged[0]
	}
	if len(m) < 1 {
		return m, changed
	}
	if _, ok := m[k]; ok {
		delete(m, k)
		changed = true
	}
	return m, changed
}

// SetLabel returns `true` if value really was changed,
// or `true` if optional parameter shouldBeChanged == true
// +gocode:public-api=true
func SetLabel(o k8types.K8sObject, label, value string, shouldBeChanged ...bool) bool {
	tmp, changed := setSomething(o.GetLabels(), label, value, shouldBeChanged)
	if changed {
		o.SetLabels(tmp)
	}
	return changed
}

// RemoveLabel returns `true` if value really was changed
// or `true` if optional parameter shouldBeChanged == true
// +gocode:public-api=true
func RemoveLabel(o k8types.K8sObject, label string, shouldBeChanged ...bool) bool {
	tmp, changed := clearSomething(o.GetLabels(), label, shouldBeChanged)
	if changed {
		o.SetLabels(tmp)
	}
	return changed
}

// SetAnnotation returns `true` if value really was changed,
// or `true` if optional parameter shouldBeChanged == true
// +gocode:public-api=true
func SetAnnotation(o k8types.K8sObject, annotation, value string, shouldBeChanged ...bool) bool {
	tmp, changed := setSomething(o.GetAnnotations(), annotation, value, shouldBeChanged)
	if changed {
		o.SetAnnotations(tmp)
	}
	return changed
}

// RemoveAnnotation returns `true` if value really was changed
// or `true` if optional parameter shouldBeChanged == true
// +gocode:public-api=true
func RemoveAnnotation(o k8types.K8sObject, annotation string, shouldBeChanged ...bool) bool {
	tmp, changed := clearSomething(o.GetAnnotations(), annotation, shouldBeChanged)
	if changed {
		o.SetAnnotations(tmp)
	}
	return changed
}

// CheckFinalizer returns `true` if provided finalizer is set
// +gocode:public-api=true
func CheckFinalizer(o k8types.K8sObject, finalizer string) bool {
	return funk.ContainsString(o.GetFinalizers(), finalizer)
}

// AddFinalizer add provided finalizer to resource.
// returns `true` if finalizer was really set
// or `true` if optional parameter shouldBeChanged == true
// +gocode:public-api=true
func AddFinalizer(o k8types.K8sObject, finalizer string, shouldBeChanged ...bool) (rv bool) {
	if len(shouldBeChanged) > 0 {
		rv = shouldBeChanged[0]
	}
	if !CheckFinalizer(o, finalizer) {
		o.SetFinalizers(append(o.GetFinalizers(), finalizer))
		return true
	}
	return rv
}

// RemoveFinalizer remove provided finalizer from resource
// returns `true` if finalizer was really removed
// or `true` if optional parameter shouldBeChanged == true
// +gocode:public-api=true
func RemoveFinalizer(o k8types.K8sObject, finalizer string, shouldBeChanged ...bool) (rv bool) {
	if len(shouldBeChanged) > 0 {
		rv = shouldBeChanged[0]
	}
	if CheckFinalizer(o, finalizer) {
		o.SetFinalizers(funk.FilterString(o.GetFinalizers(), func(s string) bool {
			return s != finalizer
		}))
		return true
	}
	return rv
}

// KeyToNamespacedName convert key (namespace/name) to types.NamespacedName
// +gocode:public-api=true
func KeyToNamespacedName(ref string) (rv types.NamespacedName) {
	namespace, name, err := cache.SplitMetaNamespaceKey(ref)
	if err != nil {
		return types.NamespacedName{}
	}
	return types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
}

// +gocode:public-api=true
func GetRuntimeObjectNamespacedName(rtobj runtime.Object) (types.NamespacedName, error) {
	obj, ok := rtobj.(k8types.K8sObject)
	if !ok {
		return types.NamespacedName{}, fmt.Errorf("GetRuntimeObjectNamespacedName: %w: given object is not a k8s object", k8types.ErrorWrongParametr)
	}
	return types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}, nil
}

// +gocode:public-api=true
func GetRuntimeObjectKey(rtobj runtime.Object) (rv string, err error) {
	nn, err := GetRuntimeObjectNamespacedName(rtobj)
	if err != nil {
		return "", err
	}
	if nn.Namespace == "" {
		rv = nn.Name
	} else {
		rv = nn.String()
	}
	return rv, err
}

// +gocode:public-api=true
func Clone(in, out interface{}) (err error) {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if encErr := enc.Encode(in); encErr != nil {
		return encErr
	}
	err = dec.Decode(out)
	return err
}

// +gocode:public-api=true
func GetLogKeyFromContext(ctx context.Context) (string, error) {
	ctxValue := ctx.Value("LogKey")
	clusterName, ok := ctxValue.(string)
	if !ok {
		return "", fmt.Errorf("%w of LogKey in the context, should be string, given: %T", k8types.ErrorWrongFormat, ctxValue)
	}
	return clusterName, nil
}

// +gocode:public-api=true
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

// ContextWithReducedDeadline -- returns new ctx/cancel pair with reduced remaining time (by divider) of given context
// if given context does not contains deadline fallbackDeadline will be used if not zero
// +gocode:public-api=true
func ContextWithReducedDeadline(ctx context.Context, divider uint, fallbackDeadline time.Duration) (newCtx context.Context, cancel context.CancelFunc) {
	if deadline, ok := ctx.Deadline(); ok {
		duration := time.Duration(time.Until(deadline).Nanoseconds() / int64(divider))
		newCtx, cancel = context.WithTimeout(ctx, duration)
	} else {

		if fallbackDeadline != 0 {
			newCtx, cancel = context.WithTimeout(ctx, fallbackDeadline)
		} else {
			newCtx, cancel = context.WithCancel(ctx)
		}
	}
	return newCtx, cancel
}

// +gocode:public-api=true
func EnsureContextDeadline(ctx context.Context, fallbackTimeout time.Duration) (newCtx context.Context, cancel context.CancelFunc, isFallback bool) {
	var (
		deadline time.Time
		ok       bool
	)
	if deadline, ok = ctx.Deadline(); ok {
		newCtx, cancel = context.WithDeadline(ctx, deadline)
	} else {
		isFallback = true
		newCtx, cancel = context.WithTimeout(ctx, fallbackTimeout)
	}
	return newCtx, cancel, isFallback
}

// +gocode:public-api=true
func EnsureNamespaceInRef(ref, fallbackNamespace string) (rv string, fallback bool) {
	if fallbackNamespace == "" {
		return ref, false
	}
	refParts := strings.Split(ref, "/")
	if len(refParts) < 2 {
		rv = fmt.Sprintf("%s/%s", fallbackNamespace, refParts[0])
		fallback = true
	} else {
		rv = fmt.Sprintf("%s/%s", refParts[0], refParts[1])
	}
	return rv, fallback
}
