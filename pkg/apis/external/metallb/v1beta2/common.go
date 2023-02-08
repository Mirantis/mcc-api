// SPDX-License-Identifier:Apache-2.0

package v1beta2

import (
	"github.com/go-kit/log"

	"github.com/Mirantis/mcc-api/v2/pkg/apis/external/metallb/validate"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// log is for logging addresspool-webhook.
var (
	Logger           log.Logger
	WebhookClient    client.Reader
	Validator        validate.ClusterObjects
	MetalLBNamespace string
)
