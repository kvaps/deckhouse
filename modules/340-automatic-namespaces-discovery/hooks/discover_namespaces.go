/*
Copyright 2021 Flant JSC

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

import (
	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	"github.com/flant/addon-operator/sdk"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	discoverPatch = map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{
				"extended-monitoring.flant.com/enabled": "true",
			},
		},
	}
	undiscoverPatch = map[string]interface{}{
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{
				"extended-monitoring.flant.com/enabled": nil,
			},
		},
	}
)

var _ = sdk.RegisterFunc(&go_hook.HookConfig{
	Queue: "/modules/automatic-namespaces-discovery/namespaces_discovery",
	Kubernetes: []go_hook.KubernetesConfig{
		{
			Name:       "namespaces",
			ApiVersion: "v1",
			Kind:       "Namespace",
			FilterFunc: applyNamespaceFilter,
		},
	},
}, handleNamespace)

func applyNamespaceFilter(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	return obj.GetName(), nil
}

func handleNamespace(input *go_hook.HookInput) error {
	snap := input.Snapshots["namespaces"]
	if len(snap) == 0 {
		input.LogEntry.Warnln("Namespace not found. Skip")
		return nil
	}

	for _, s := range snap {
		ns := s.(corev1.Namespace)
		input.PatchCollector.MergePatch(discoverPatch, "v1", "Namespace", "", ns.Name)
	}

	return nil
}
