// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tempdir

import (
	"fmt"
	"os"
)

// CreateTemporaryPath creates a temporary path.
func CreateTemporaryPath(reposTempPath, prefix string) (string, error) {
	if reposTempPath != "" {
		if err := os.MkdirAll(reposTempPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create directory %s: %w", reposTempPath, err)
		}
	}
	basePath, err := os.MkdirTemp(reposTempPath, prefix+".git")
	if err != nil {
		return "", fmt.Errorf("failed to create dir %s-*.git: %w", prefix, err)
	}
	return basePath, nil
}

// RemoveTemporaryPath removes the temporary path.
func RemoveTemporaryPath(basePath string) error {
	if _, err := os.Stat(basePath); !os.IsNotExist(err) {
		return os.RemoveAll(basePath)
	}
	return nil
}
