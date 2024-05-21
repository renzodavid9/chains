/*
Copyright 2024 The Tekton Authors
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

package metadata

import (
	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v1"
	"github.com/tektoncd/chains/pkg/chains/objects"
)

// GetBuildMetadata returns SLSA metadata.
func GetBuildMetadata(obj objects.TektonObject) slsa.BuildMetadata {
	return slsa.BuildMetadata{
		InvocationID: string(obj.GetUID()),
		StartedOn:    obj.GetStartTime(),
		FinishedOn:   obj.GetCompletitionTime(),
	}
}
