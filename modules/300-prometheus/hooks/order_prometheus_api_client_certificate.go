/*
Copyright 2022 Flant JSC

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

package hooks

import "github.com/deckhouse/deckhouse/go_lib/hooks/tls_certificate"

var _ = tls_certificate.RegisterOrderCertificateHook(
	[]tls_certificate.OrderCertificateRequest{
		{
			Namespace:  "d8-monitoring",
			SecretName: "prometheus-api-client-tls",
			CommonName: "d8-monitoring:prometheus-api-client",
			ValueName:  "internal.prometheusAPIClientTLS",
			Groups:     []string{"prometheus:auth"},
			ModuleName: "prometheus",
		},
	},
)
