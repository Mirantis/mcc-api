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

package types

import (
	"errors"
	"fmt"
)

var (
	Error                  = errors.New("")
	ErrorUndefined         = fmt.Errorf("not defined%w", Error)
	ErrorClusterUndefined  = fmt.Errorf("cluster %w", ErrorUndefined)
	ErrorRegionUndefined   = fmt.Errorf("region %w", ErrorUndefined)
	ErrorProviderUndefined = fmt.Errorf("provider %w", ErrorUndefined)

	ErrorObjLookup          = fmt.Errorf("%w", Error)
	ErrorNotFound           = fmt.Errorf("not found%w", ErrorObjLookup)
	ErrorMoreThanOne        = fmt.Errorf("more than one%w", ErrorObjLookup)
	ErrorInoperable         = fmt.Errorf("inoperable%w", Error)
	ErrorWrong              = fmt.Errorf("wrong%w", Error)
	ErrorWrongFormat        = fmt.Errorf("%w format", ErrorWrong)
	ErrorWrongObject        = fmt.Errorf("%w object", ErrorWrong)
	ErrorWrongName          = fmt.Errorf("%w name", ErrorWrong)
	ErrorWrongParametr      = fmt.Errorf("%w parameter", ErrorWrong)
	ErrorWrongRequest       = fmt.Errorf("%w request", ErrorWrong)
	ErrorTimeoutOrCancel    = fmt.Errorf("timeout or cancel%w", Error)
	ErrorUnableToAllocate   = fmt.Errorf("unable to allocate%w", Error)
	ErrorDoNothing          = fmt.Errorf("do nothing%w", Error)
	ErrorSomethingWentWrong = fmt.Errorf("something went wrong%w", Error)
)
