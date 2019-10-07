// go-config - Library for reading cacophony config files.
// Copyright (C) 2018, The Cacophony Project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package config

import "time"

const TestHostsKey = "test-hosts"

func init() {
	allSections[TestHostsKey] = section{
		key:         TestHostsKey,
		mapToStruct: makeMapToStruct(TestHosts{}),
		validate:    noValidateFunc,
	}
}

type TestHosts struct {
	URLs         []string
	PingWaitTime time.Duration `mapstructure:"ping-wait-time"`
	PingRetries  int           `mapstructure:"ping-retries"`
}

func DefaultTestHosts() TestHosts {
	return TestHosts{
		URLs:         []string{"1.1.1.1", "8.8.8.8"},
		PingWaitTime: time.Second * 30,
		PingRetries:  5,
	}
}
