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

const ModemdKey = "modemd"

func init() {
	allSections[ModemdKey] = section{
		key:         ModemdKey,
		mapToStruct: makeMapToStruct(Modemd{}),
		validate:    noValidateFunc,
	}
}

type Modemd struct {
	TestInterval      time.Duration `mapstructure:"test-interval"`
	InitialOnDuration time.Duration `mapstructure:"initial-on-duration"`
	FindModemTimeout  time.Duration `mapstructure:"find-modem-timeout"`
	ConnectionTimeout time.Duration `mapstructure:"connection-timeout"`
	RequestOnDuration time.Duration `mapstructure:"request-on-duration"`
	Modems            []Modem       `mapstructure:"modems"`
}

type Modem struct {
	Name            string `mapstructure:"name"`
	NetDev          string `mapstructure:"net-dev"`
	VendorProductID string `mapstructure:"vendor-product-id"`
}

func DefaultModemd() Modemd {
	return Modemd{
		TestInterval:      time.Minute * 5,
		InitialOnDuration: time.Hour * 24,
		FindModemTimeout:  time.Minute * 2,
		ConnectionTimeout: time.Minute,
		RequestOnDuration: time.Hour * 24,
		Modems: []Modem{
			Modem{Name: "Huawei 4G modem", NetDev: "eth1", VendorProductID: "12d1:14db"},
			Modem{Name: "Spark 3G modem", NetDev: "usb0", VendorProductID: "19d2:1405"},
		},
	}
}
