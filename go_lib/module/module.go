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

package module

import (
	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
)

func GetHTTPSMode(moduleName string, input *go_hook.HookInput) string {
	var (
		modulePath = moduleName + ".https.mode"
		globalPath = "global.modules.https.mode"
	)

	v, ok := input.Values.GetOk(modulePath)
	if ok {
		return v.String()
	}

	v, ok = input.Values.GetOk(globalPath)
	if ok {
		return v.String()
	}

	panic("https mode is not defined")
}

// IsEnabled check module on enable. moduleName should be in `kebab-case` without order prefix
func IsEnabled(moduleName string, input *go_hook.HookInput) bool {
	modules := input.Values.Get("global.enabledModules").Array()

	for _, m := range modules {
		if m.String() == moduleName {
			return true
		}
	}

	return false
}
