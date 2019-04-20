/*
Copyright 2016 The Kubernetes Authors.

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

package version

import (
	"fmt"
	"runtime"
)

// VERSION is the app-global version string, which should be substituted with a
// real value during build.
var VERSION = "UNKNOWN"

// BUILD_DATE is the app-global version string, which should be substituted with a
// real value during build.
var BUILD_DATE = "UNKNOWN"

// COMMIT_SHA1 is the app-global version string, which should be substituted with a
// real value during build.
var COMMIT_SHA1 = "UNKNOWN"

func GetVersion() string {
	return fmt.Sprintf("Version: %s, Build date: %s, Commit SHA: %s, Go version: %s", VERSION, BUILD_DATE, COMMIT_SHA1, runtime.Version())
}
