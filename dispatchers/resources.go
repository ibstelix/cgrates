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

package dispatchers

import (
	"github.com/cgrates/cgrates/engine"
	"github.com/cgrates/cgrates/utils"
)

func (dS *DispatcherService) ResourceSv1Ping(args *CGREvWithApiKey, rpl *string) (err error) {
	if dS.attrS != nil {
		if err = dS.authorize(utils.ResourceSv1Ping,
			args.CGREvent.Tenant,
			args.APIKey, args.CGREvent.Time); err != nil {
			return
		}
	}
	return dS.Dispatch(&args.CGREvent, utils.MetaResources, args.RouteID,
		utils.ResourceSv1Ping, args.CGREvent, rpl)
}

func (dS *DispatcherService) ResourceSv1GetResourcesForEvent(args *ArgsV1ResUsageWithApiKey,
	reply *engine.Resources) (err error) {
	if dS.attrS != nil {
		if err = dS.authorize(utils.ResourceSv1GetResourcesForEvent,
			args.ArgRSv1ResourceUsage.CGREvent.Tenant,
			args.APIKey, args.ArgRSv1ResourceUsage.CGREvent.Time); err != nil {
			return
		}

	}
	return dS.Dispatch(&args.CGREvent, utils.MetaResources, args.RouteID,
		utils.ResourceSv1GetResourcesForEvent, args.ArgRSv1ResourceUsage, reply)
}

func (dS *DispatcherService) ResourceSv1AuthorizeResources(args *ArgsV1ResUsageWithApiKey,
	reply *string) (err error) {
	if dS.attrS != nil {
		if err = dS.authorize(utils.ResourceSv1AuthorizeResources,
			args.ArgRSv1ResourceUsage.CGREvent.Tenant,
			args.APIKey, args.ArgRSv1ResourceUsage.CGREvent.Time); err != nil {
			return
		}

	}
	return dS.Dispatch(&args.CGREvent, utils.MetaResources, args.RouteID,
		utils.ResourceSv1AuthorizeResources, args.ArgRSv1ResourceUsage, reply)
}

func (dS *DispatcherService) ResourceSv1AllocateResources(args *ArgsV1ResUsageWithApiKey,
	reply *string) (err error) {
	if dS.attrS != nil {
		if err = dS.authorize(utils.ResourceSv1AllocateResources,
			args.ArgRSv1ResourceUsage.CGREvent.Tenant,
			args.APIKey, args.ArgRSv1ResourceUsage.CGREvent.Time); err != nil {
			return
		}

	}
	return dS.Dispatch(&args.CGREvent, utils.MetaResources, args.RouteID,
		utils.ResourceSv1AllocateResources, args.ArgRSv1ResourceUsage, reply)
}

func (dS *DispatcherService) ResourceSv1ReleaseResources(args *ArgsV1ResUsageWithApiKey,
	reply *string) (err error) {
	if dS.attrS != nil {
		if err = dS.authorize(utils.ResourceSv1ReleaseResources,
			args.ArgRSv1ResourceUsage.CGREvent.Tenant,
			args.APIKey, args.ArgRSv1ResourceUsage.CGREvent.Time); err != nil {
			return
		}

	}
	return dS.Dispatch(&args.CGREvent, utils.MetaResources, args.RouteID,
		utils.ResourceSv1ReleaseResources, args.ArgRSv1ResourceUsage, reply)
}
