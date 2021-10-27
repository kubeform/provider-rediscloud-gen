/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"log"
	"os"

	"kubeform.dev/generator-v2/util"

	"github.com/RedisLabs/terraform-provider-rediscloud/rediscloud/provider"
)

func main() {
	apisPath := flag.String("apis-path", "<empty>", "path to generate the apis. Pass empty string to use default path <$GOPATH/kubeform.dev/provider-rediscloud-api")
	controllerPath := flag.String("controller-path", "<empty>", "path to generate the controller. Pass empty string to use default path <$GOPATH/kubeform.dev/provider-rediscloud-controller")
	flag.Parse()

	if *apisPath == "<empty>" {
		apisPath = nil
	}
	if *controllerPath == "<empty>" {
		controllerPath = nil
	}

	opts := &util.GeneratorOptions{
		ProviderName:             "rediscloud",
		ProviderNameOriginal:     "provider",
		ProviderData:             provider.New(os.Getenv("PROVIDER_VERSION"))(),
		ProviderGetterMethodBody: `provider.New("")()`,
		ProviderImportPath:       "github.com/RedisLabs/terraform-provider-rediscloud/rediscloud/provider",
		ProviderConfig:           "RedisLabs/rediscloud",
		Version:                  "v1alpha1",
		APIsPath:                 apisPath,
		ControllerPath:           controllerPath,
	}
	err := util.Generate(opts)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
