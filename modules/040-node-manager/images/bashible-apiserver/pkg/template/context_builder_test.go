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

package template

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2"
	"sigs.k8s.io/yaml"
)

var golden = flag.Bool("golden", false, "Generate golden files")

func TestGenerateContext(t *testing.T) {
	b := NewContextBuilder(context.Background(), "../../../../candi/")

	reg := registryInputData{
		Address:      "registry.deckhouse.io",
		Path:         "/sys/deckhouse-oss",
		Scheme:       "https",
		CA:           "",
		DockerConfig: []byte(`{"auths": {"registry.deckhouse.io": {"auth": "dGVzdDp0ZXN0Cg=="}}}`),
	}
	b.SetRegistryData(reg.toRegistry())

	vdata, err := ioutil.ReadFile("../../../../candi/version_map.yml")
	require.NoError(t, err)

	var versionMap map[string]interface{}

	err = yaml.Unmarshal(vdata, &versionMap)
	require.NoError(t, err)
	b.SetVersionMapData(versionMap)

	dirs, err := ioutil.ReadDir("test_data")
	require.NoError(t, err)

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		if dir.Name() == "openstack" {
			continue
		}

		stepsWriter := func(bundle, ngName string, m map[string]string) {
			dir := path.Join("test_data", dir.Name(), "golden")

			for fileName, content := range m {
				if *golden {
					_ = os.MkdirAll(path.Join(dir, bundle, ngName), 0777)
					_ = ioutil.WriteFile(path.Join(dir, bundle, ngName, fileName), []byte(content), 0666)
				} else {
					existed, err := ioutil.ReadFile(path.Join(dir, bundle, ngName, fileName))
					require.NoError(t, err)
					assert.Equal(t, string(existed), content, fmt.Sprintf("File %q is nod valid", fileName))
				}
			}
		}
		b.setStepsOutput(stepsWriter)

		inputYaml, err := ioutil.ReadFile(path.Join("test_data", dir.Name(), "/input.yaml"))
		require.NoError(t, err)

		var input inputData
		err = yaml.Unmarshal(inputYaml, &input)
		require.NoError(t, err)
		b.SetInputData(input)

		imagesData, err := ioutil.ReadFile(path.Join("test_data", dir.Name(), "images_tags.json"))
		require.NoError(t, err)

		var images map[string]map[string]string
		err = json.Unmarshal(imagesData, &images)
		require.NoError(t, err)

		b.SetImagesData(images)

		data, checksum, errs := b.Build()
		assert.Len(t, errs, 0)

		checkStr := strings.Builder{}
		checkStr.WriteString("Next checksums will be written to checksum-secret:\n")
		for k, v := range checksum {
			checkStr.WriteString(fmt.Sprintf("\t%s: %s\n", k, v))
		}
		fmt.Print(checkStr.String())

		rawData, err := yaml.Marshal(data.Map())
		if err != nil {
			klog.Errorf("Failed to marshal data", err)
			return
		}

		goldenContextFile := path.Join("test_data", dir.Name(), "golden", "context.yml")

		if *golden {
			err = ioutil.WriteFile(goldenContextFile, rawData, 0666)
			require.NoError(t, err)
		} else {
			existed, err := ioutil.ReadFile(goldenContextFile)
			require.NoError(t, err)
			assert.YAMLEq(t, string(existed), string(rawData))
		}
	}
}
