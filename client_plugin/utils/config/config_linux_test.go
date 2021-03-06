// Copyright 2016-2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config_test

// Test Loading JSON config files

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmware/vsphere-storage-for-docker/client_plugin/utils/config"
)

func TestLoad(t *testing.T) {
	conf, err := config.Load("../../default-config.json")
	assert.Nil(t, err)
	assert.Equal(t, conf.MaxLogSizeMb, 10)
	assert.Equal(t, conf.MaxLogAgeDays, 28)
	assert.Equal(t, conf.MaxLogFiles, 10)
	assert.Equal(t, conf.LogPath, "/var/log/vsphere-storage-for-docker.log")
}
