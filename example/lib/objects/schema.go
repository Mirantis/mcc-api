package objects

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	scheme "github.com/Mirantis/mcc-api/pkg/apis/public"
)

var Codecs = serializer.NewCodecFactory(scheme.Scheme)

func Decode(data []byte) (runtime.Object, error) {
	obj, _, err := Codecs.UniversalDeserializer().Decode(data, nil, nil)
	return obj, err
}
