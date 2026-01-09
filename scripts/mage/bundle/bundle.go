// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Package bundle provides utilities for building and locating the OPA bundle.
package bundle

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/sh"
)

// cloudbeatRoot returns the root directory of the cloudbeat module.
func cloudbeatRoot() string {
	_, thisFile, _, _ := runtime.Caller(0)
	// This file is at: cloudbeat/scripts/mage/bundle/bundle.go
	return filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(thisFile))))
}

// Build builds the OPA policy bundle for cloudbeat.
func Build() error {
	root := cloudbeatRoot()
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)
	os.Chdir(root)
	return sh.Run("bin/opa", "build", "-b", "security-policies/bundle", "-e", "security-policies/bundle/compliance")
}

// Path returns the absolute path to the OPA bundle file.
func Path() string {
	return filepath.Join(cloudbeatRoot(), "bundle.tar.gz")
}
