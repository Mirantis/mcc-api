/*
Copyright © 2020 Mirantis

Inspired by https://github.com/inwinstack/ipam/, https://github.com/inwinstack/blended/

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8sutil

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thoas/go-funk"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"

	"github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

var IsMACaddr = regexp.MustCompile(`[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}`)

//-----------------------------------------------------------------------------
func setSomething(m map[string]string, k, v string, shouldBeChanged []bool) (map[string]string, bool) {
	var changed bool
	if m == nil {
		m = map[string]string{}
	}
	if len(shouldBeChanged) > 0 {
		changed = shouldBeChanged[0]
	}
	if tmp, ok := m[k]; ok && tmp == v {
		// present, equal
		return m, changed
	}
	m[k] = v
	return m, true
}

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

//-----------------------------------------------------------------------------

// SetLabel returns `true` if value really was changed,
// or `true` if optional parameter shouldBeChanged == true
func SetLabel(o k8types.K8sObject, label, value string, shouldBeChanged ...bool) bool {
	tmp, changed := setSomething(o.GetLabels(), label, value, shouldBeChanged)
	if changed {
		o.SetLabels(tmp)
	}
	return changed
}

// RemoveLabel returns `true` if value really was changed
// or `true` if optional parameter shouldBeChanged == true
func RemoveLabel(o k8types.K8sObject, label string, shouldBeChanged ...bool) bool {
	tmp, changed := clearSomething(o.GetLabels(), label, shouldBeChanged)
	if changed {
		o.SetLabels(tmp)
	}
	return changed
}

//-----------------------------------------------------------------------------

// SetAnnotation returns `true` if value really was changed,
// or `true` if optional parameter shouldBeChanged == true
func SetAnnotation(o k8types.K8sObject, annotation, value string, shouldBeChanged ...bool) bool {
	tmp, changed := setSomething(o.GetAnnotations(), annotation, value, shouldBeChanged)
	if changed {
		o.SetAnnotations(tmp)
	}
	return changed
}

// RemoveAnnotation returns `true` if value really was changed
// or `true` if optional parameter shouldBeChanged == true
func RemoveAnnotation(o k8types.K8sObject, annotation string, shouldBeChanged ...bool) bool {
	tmp, changed := clearSomething(o.GetAnnotations(), annotation, shouldBeChanged)
	if changed {
		o.SetAnnotations(tmp)
	}
	return changed
}

//-----------------------------------------------------------------------------

func MakeNeedToUpdate(o k8types.K8sObject, oldVal, newVal interface{}) {
	if !reflect.DeepEqual(oldVal, newVal) {
		SetNeedToUpdate(o)
	}
}

// SetNeedToUpdate returns `true` if value really was changed,
// or `true` if optional parameter shouldBeChanged == true
func SetNeedToUpdate(o k8types.K8sObject, shouldBeChanged ...bool) bool {
	tmp, changed := setSomething(o.GetAnnotations(), config.NeedUpdateKey, "true", shouldBeChanged)
	if changed {
		o.SetAnnotations(tmp)
	}
	return changed
}

// ClearNeedToUpdate returns `true` if value really was changed
// or `true` if optional parameter shouldBeChanged == true
func ClearNeedToUpdate(o k8types.K8sObject, shouldBeChanged ...bool) bool {
	tmp, changed := clearSomething(o.GetAnnotations(), config.NeedUpdateKey, shouldBeChanged)
	if changed {
		o.SetAnnotations(tmp)
	}
	return changed
}

// IsNeedToUpdate returns `true` if "NeedToUpdate" flag is set
func IsNeedToUpdate(o k8types.K8sObject) bool {
	tmp := o.GetAnnotations()
	if tmp == nil {
		return false
	}
	return tmp[config.NeedUpdateKey] != ""
}

//-----------------------------------------------------------------------------

// CheckFinalizer returns `true` if provided finalizer is set
func CheckFinalizer(o k8types.K8sObject, finalizer string) bool {
	return funk.ContainsString(o.GetFinalizers(), finalizer)
}

// AddFinalizer add provided finalizer to resource.
// returns `true` if finalizer was really set
// or `true` if optional parameter shouldBeChanged == true
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

//-----------------------------------------------------------------------------

// KeyToNamespacedName convert key (namespace/name) to types.NamespacedName
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

func Clone(in, out interface{}) (err error) {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(in); err != nil {
		return err
	}
	err = dec.Decode(out)
	return err
}

func GetLogKeyFromContext(ctx context.Context) (string, error) {
	ctxValue := ctx.Value("LogKey")
	clusterName, ok := ctxValue.(string)
	if !ok {
		return "", fmt.Errorf("%w of LogKey in the context, should be string, given: %T", k8types.ErrorWrongFormat, ctxValue)
	}
	return clusterName, nil
}

func GetNamespaceFromContext(ctx context.Context) (string, error) {
	ctxValue := ctx.Value(config.NamespaceLabel)
	clusterName, ok := ctxValue.(string)
	if !ok {
		return "", fmt.Errorf("%w of Namespace in the context, should be string, given: %T", k8types.ErrorWrongFormat, ctxValue)
	}
	return clusterName, nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func GetServiceLabels(obj k8types.K8sObject) map[string]string {
	rv := map[string]string{}
	labels := obj.GetLabels()
	for k := range labels {
		if strings.HasPrefix(k, config.PerServiceLabelPrefix) {
			rv[k] = labels[k]
		}
	}
	return rv
}

// ContextWithReducedDeadline -- returns new ctx/cancel pair with reduced remaining time (by divider) of given context
// if given context does not contains deadline fallbackDeadline will be used if not zero
func ContextWithReducedDeadline(ctx context.Context, divider uint, fallbackDeadline time.Duration) (newCtx context.Context, cancel context.CancelFunc) {
	if deadline, ok := ctx.Deadline(); ok {
		duration := time.Duration(time.Until(deadline).Nanoseconds() / int64(divider))
		newCtx, cancel = context.WithTimeout(ctx, duration)
	} else {
		// context has no deadline
		if fallbackDeadline != 0 {
			newCtx, cancel = context.WithTimeout(ctx, fallbackDeadline)
		} else {
			newCtx, cancel = context.WithCancel(ctx)
		}
	}
	return newCtx, cancel
}

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
