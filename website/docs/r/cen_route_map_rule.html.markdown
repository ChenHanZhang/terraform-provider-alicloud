---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_route_map_rule"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Route Map Rule resource.
---

# alicloud_cen_route_map_rule

Provides a Cloud Enterprise Network (CEN) Route Map Rule resource.

Route map rule.

For information about Cloud Enterprise Network (CEN) Route Map Rule and how to use it, see [What is Route Map Rule](https://next.api.alibabacloud.com/document/Cbn/2017-09-12/CreateRouteMapRule).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_cen_route_map" "defaultxtaOyM" {
  priority = "94"
}


resource "alicloud_cen_route_map_rule" "default" {
  description = "example"
  match_conditions {
    address_prefixes_include               = ["192.168.1.0/24", "192.168.2.0/24", "192.168.3.0/24"]
    as_paths_include                       = ["65500", "65501", "65502"]
    community_set_include                  = ["400:1", "400:2", "400:3"]
    destination_instance_ids               = ["vpc-1", "vpc-2", "vpc-3"]
    destination_instance_ids_reverse_match = false
    destination_instance_types             = ["VPC", "VPN", "VBR"]
    destination_region_ids                 = ["cn-shanghai", "cn-beijing", "cn-qingdao"]
    route_types                            = ["BGP"]
    source_instance_ids                    = ["vpc-1", "vpc-2", "vpc-3"]
    source_instance_ids_reverse_match      = false
    source_instance_types                  = ["VPC", "VPN", "VBR"]
    source_region_ids                      = ["cn-beijing", "cn-qingdao", "cn-shanghai"]
    destination_route_table_ids            = ["vtb-1", "vtb-2", "vtb-3"]
    source_route_table_ids                 = ["vtb-1", "vtb-2", "vtb-3"]
    match_address_type                     = "IPv4"
  }
  priority = "1"
  dry_run  = false
  set_actions {
    route_action    = "Permit"
    as_path_prepend = ["65500", "65501", "65502"]
    community_add   = ["400:1", "400:2", "400:3"]
  }
  route_map_id = alicloud_cen_route_map.defaultxtaOyM.id
  name         = "example"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) The description can be empty or contain 1 to 256 characters, and must not start with http:// or https://.
* `dry_run` - (Optional) Specifies whether to perform a dry run for this request, including checks on permissions and instance status.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `match_conditions` - (Optional, ForceNew, Set) Match conditions. See [`match_conditions`](#match_conditions) below.
* `name` - (Optional) The name can be empty or contain 1 to 128 characters, and must not start with http:// or https://.
* `priority` - (Required, Int) Value range: 1 to 100. A smaller number indicates a higher priority.
* `route_map_id` - (Required, ForceNew) ID of the route policy to which the rule belongs.
* `set_actions` - (Optional, ForceNew, Set) Modification actions. See [`set_actions`](#set_actions) below.

### `match_conditions`

The match_conditions supports the following:
* `address_prefixes_include` - (Optional, List) A match is successful if a route prefix in the match condition contains the route prefix of the route being matched.  
For example, a policy defined with 10.10.0.0/16 can fuzzy-match a route with 10.10.1.0/24.
* `address_prefixes_match` - (Optional, List) For a match to succeed, the route prefix in the matching condition must exactly match the route prefix of the target route.  
For example, a policy defined for 10.10.0.0/16 matches only the route 10.10.0.0/16.
* `as_paths_include` - (Optional, List) A match is successful if the AS Path in the match condition overlaps with the AS Path of the route being matched.
* `as_paths_match` - (Optional, List) A match is successful only if the AS Path in the match condition exactly matches the AS Path of the route being matched.
* `community_set_include` - (Optional, List) A match is considered successful if there is any overlap between the Community in the match condition and the Community of the route being matched.
* `community_set_match` - (Optional, List) The Community in the match condition must exactly match the Community of the route being matched for the match to succeed.
* `destination_instance_ids` - (Optional, List) You can enter instance IDs of the following types:
  - Virtual Private Cloud (VPC) instance ID  
  - Virtual Border Router (VBR) instance ID  
  - Cloud Connect Network (CCN) instance ID  
  - Smart Access Gateway instance ID  
  - IPsec connection ID  

You can specify up to 64 instance IDs.
* `destination_instance_ids_reverse_match` - (Optional) Valid values:
  - false (default): No. The match succeeds if the destination instance ID of the route is included in DestinationInstanceIds.N.  
  - true: Yes. The match succeeds if the destination instance ID of the route is not included in DestinationInstanceIds.N.
* `destination_instance_types` - (Optional, List) The following instance types are supported:
VPC: Virtual Private Cloud instance.
VBR: Virtual Border Router instance.
CCN: Cloud Connect Network instance.
VPN: IPsec connection.
* `destination_region_ids` - (Optional, List) You can specify up to 64 region IDs.
* `destination_route_table_ids` - (Optional, List) You can specify up to 64 route table IDs.
* `match_address_type` - (Optional) Valid values:
IPv4: Matches only IPv4 routes.
IPv6: Matches only IPv6 routes.
This parameter can be left empty, which indicates matching all route types.
* `route_types` - (Optional, List) Supported route types:
System: System routes automatically generated by the system.
Custom: Custom routes manually added by users.
BGP: BGP routes propagated through the BGP routing protocol.
* `source_instance_ids` - (Optional, List) You can enter instance IDs of the following types:
  - Virtual Private Cloud (VPC) instance ID
  - Virtual Border Router (VBR) instance ID
  - Cloud Connect Network (CCN) instance ID
  - Smart Access Gateway instance ID
  - IPsec connection ID

You can specify up to 64 instance IDs.
* `source_instance_ids_reverse_match` - (Optional) Valid values:
false (default): No. If the source instance ID of the route is included in SourceInstanceIds.N, the match succeeds.
true: Yes. If the source instance ID of the route is not included in SourceInstanceIds.N, the match succeeds.
* `source_instance_types` - (Optional, List) The following instance types are supported:
VPC: Virtual Private Cloud instance.
VBR: Virtual Border Router instance.
CCN: Cloud Connect Network instance.
VPN: VPN gateway instance or IPsec connection.
* `source_region_ids` - (Optional, List) You can specify up to 64 region IDs.
* `source_route_table_ids` - (Optional, List) You can specify up to 64 route table IDs.

### `set_actions`

The set_actions supports the following:
* `as_path_prepend` - (Optional, List) This parameter specifies the action to perform after a route matches the conditions. You can enter up to 64 AS numbers.
* `as_path_replace` - (Optional, List) This parameter specifies the action to be performed after a route matches the conditions. You can enter up to 64 AS numbers.
* `community_add` - (Optional, List) This parameter specifies the action to be performed after a route matches the conditions. You can enter up to 64 Communities.
* `community_replace` - (Optional, List) This parameter specifies the action to be performed after a route matches the conditions. You can enter up to 64 Communities.
* `next_priority` - (Optional, Int) You can set the priority of the associated route policy only when the SetAction value is Permit. Only routes that are permitted will continue to match the next associated route policy rule. The priority of the next associated route policy rule must be lower than that of the current route policy rule.
* `route_action` - (Optional) Valid values:
Permit: Allows matching routes to pass.
Deny: Denies matching routes from passing.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The status of the route map rule.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Route Map Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Route Map Rule.
* `update` - (Defaults to 5 mins) Used when update the Route Map Rule.

## Import

Cloud Enterprise Network (CEN) Route Map Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_route_map_rule.example <route_map_rule_id>
```