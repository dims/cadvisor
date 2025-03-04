// Copyright 2025 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package install

import (
	"time"

	"github.com/google/cadvisor/resctrl"
	"github.com/google/cadvisor/resctrl/intel"

	"k8s.io/klog/v2"
)

type managerplugin struct {
}

func (m *managerplugin) NewManager(interval time.Duration, vendorID string, inHostNamespace bool) (resctrl.ResControlManager, error) {
	return intel.NewManager(interval, intel.Setup, vendorID, inHostNamespace)
}

func init() {
	err := resctrl.RegisterPlugin("intel", &managerplugin{})
	if err != nil {
		klog.Fatalf("Failed to register intel resctrl plugin: %v", err)
	}
}
