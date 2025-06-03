// Copyright Meshery Authors
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

package config

import (
	"github.com/layer5io/meshkit/errors"
)

const (
	ErrEmptyConfigCode           = "1000"
	ErrInstallBinaryCode         = "1001"
	ErrGetLatestReleasesCode     = "1002"
	ErrGetLatestReleaseNamesCode = "1003"
	ErrStatusCheckCode           = "1004"
	ErrUnmarshalCode             = "1005"
)

var (
	// ErrEmptyConfig error is the error when config is invalid
	ErrEmptyConfig = errors.New(ErrEmptyConfigCode, errors.Alert, []string{"Config is empty"}, []string{}, []string{}, []string{})
)

// ErrGetLatestReleases is the error for fetching nsm-mesh releases
func ErrGetLatestReleases(err error) error {
	return errors.New(ErrGetLatestReleasesCode, errors.Alert, []string{"Unable to fetch release info"}, []string{err.Error()}, []string{}, []string{})
}

// ErrGetLatestReleaseNames is the error for fetching nsm-mesh releases
func ErrGetLatestReleaseNames(err error) error {
	return errors.New(ErrGetLatestReleaseNamesCode, errors.Alert, []string{"Failed to extract release names"}, []string{err.Error()}, []string{}, []string{})
}

// ErrInstallBinary captures failure to update filesystem permissions
func ErrInstallBinary(err error) error {
	return errors.New(ErrInstallBinaryCode, errors.Alert, []string{"Failed to change permission of the binary"}, []string{err.Error()}, []string{}, []string{})
}

func ErrStatusCheck(status string) error {
	return errors.New(ErrStatusCheckCode, errors.Alert, []string{"Error Bad Status", status}, []string{}, []string{}, []string{})
}

func ErrUnmarshal(err error, obj string) error {
	return errors.New(ErrUnmarshalCode, errors.Alert, []string{"Unable to unmarshal the : ", obj}, []string{err.Error()}, []string{"Object is not a valid json object"}, []string{"Make sure if the object passed is a valid json"})
}
