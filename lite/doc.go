// Copyright 2024 Google Inc. All Rights Reserved.
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

// Package lite provides a minimal subset of cadvisor for Kubernetes.
//
// This module is designed for Kubernetes and includes only the packages
// that Kubernetes needs, with minimal dependencies. It excludes:
//
//   - Docker container runtime support
//   - Podman container runtime support
//   - ZFS filesystem support (unless built with -tags zfs)
//   - AWS cloud provider dependencies
//   - GCE cloud provider dependencies
//
// Usage:
//
// Instead of importing the full cadvisor module:
//
//	import "github.com/google/cadvisor/lite/manager"
//
// Import via the lite module to get reduced dependencies:
//
//	import "github.com/google/cadvisor/lite/manager" // with lite go.mod
//
// To register container handlers for Kubernetes, import:
//
//	import _ "github.com/google/cadvisor/lite/container/install"
//
// This will register containerd, crio, and systemd handlers.
package lite
