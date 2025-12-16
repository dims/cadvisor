// Copyright 2014 Google Inc. All Rights Reserved.
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

//go:build linux && !zfs

package fs

import "errors"

// getZfstats returns an error when ZFS support is not compiled in.
func getZfstats(poolName string) (uint64, uint64, uint64, error) {
	return 0, 0, 0, errors.New("ZFS support not compiled in (build with -tags zfs)")
}
