/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package engine

import (
	"sort"

	"github.com/cgrates/cgrates/config"
	"github.com/cgrates/cgrates/utils"
)

type Attribute struct {
	FilterIDs  []string
	FieldName  string
	Substitute config.RSRParsers
}

type AttributeProfile struct {
	Tenant             string
	ID                 string
	Contexts           []string // bind this AttributeProfile to multiple contexts
	FilterIDs          []string
	ActivationInterval *utils.ActivationInterval // Activation interval
	Attributes         []*Attribute
	Blocker            bool // blocker flag to stop processing on multiple runs
	Weight             float64
}

func (ap *AttributeProfile) compileSubstitutes() (err error) {
	for _, attr := range ap.Attributes {
		if err = attr.Substitute.Compile(); err != nil {
			return
		}
	}
	return
}

// Compile is a wrapper for convenience setting up the AttributeProfile
func (ap *AttributeProfile) Compile() error {
	return ap.compileSubstitutes()
}

func (als *AttributeProfile) TenantID() string {
	return utils.ConcatenatedKey(als.Tenant, als.ID)
}

// AttributeProfiles is a sortable list of Attribute profiles
type AttributeProfiles []*AttributeProfile

// Sort is part of sort interface, sort based on Weight
func (aps AttributeProfiles) Sort() {
	sort.Slice(aps, func(i, j int) bool { return aps[i].Weight > aps[j].Weight })
}
