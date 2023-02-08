package types

import (
	"errors"
	"fmt"
)

var (
	// +gocode:public-api=true
	Error = errors.New("")
	// +gocode:public-api=true
	ErrorUndefined = fmt.Errorf("not defined%w", Error)
	// +gocode:public-api=true
	ErrorClusterUndefined = fmt.Errorf("cluster %w", ErrorUndefined)
	// +gocode:public-api=true
	ErrorRegionUndefined = fmt.Errorf("region %w", ErrorUndefined)
	// +gocode:public-api=true
	ErrorProviderUndefined = fmt.Errorf("provider %w", ErrorUndefined)

	// +gocode:public-api=true
	ErrorObjLookup = fmt.Errorf("%w", Error)
	// +gocode:public-api=true
	ErrorNotFound = fmt.Errorf("not found%w", ErrorObjLookup)
	// +gocode:public-api=true
	ErrorMoreThanOne = fmt.Errorf("more than one%w", ErrorObjLookup)
	// +gocode:public-api=true
	ErrorInoperable = fmt.Errorf("inoperable%w", Error)
	// +gocode:public-api=true
	ErrorWrong = fmt.Errorf("wrong%w", Error)
	// +gocode:public-api=true
	ErrorWrongFormat = fmt.Errorf("%w format", ErrorWrong)
	// +gocode:public-api=true
	ErrorWrongObject = fmt.Errorf("%w object", ErrorWrong)
	// +gocode:public-api=true
	ErrorWrongName = fmt.Errorf("%w name", ErrorWrong)
	// +gocode:public-api=true
	ErrorWrongParametr = fmt.Errorf("%w parameter", ErrorWrong)
	// +gocode:public-api=true
	ErrorWrongRequest = fmt.Errorf("%w request", ErrorWrong)
	// +gocode:public-api=true
	ErrorTimeoutOrCancel = fmt.Errorf("timeout or cancel%w", Error)
	// +gocode:public-api=true
	ErrorUnableToAllocate = fmt.Errorf("unable to allocate%w", Error)
	// +gocode:public-api=true
	ErrorDoNothing = fmt.Errorf("do nothing%w", Error)
	// +gocode:public-api=true
	ErrorSomethingWentWrong = fmt.Errorf("something went wrong%w", Error)
)
