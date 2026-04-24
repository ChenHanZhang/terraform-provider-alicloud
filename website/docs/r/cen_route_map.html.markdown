---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_route_map"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Route Map resource.
---

# alicloud_cen_route_map

Provides a Cloud Enterprise Network (CEN) Route Map resource.



For information about Cloud Enterprise Network (CEN) Route Map and how to use it, see [What is Route Map](https://next.api.alibabacloud.com/document/Cbn/2017-09-12/CreateRouteMap).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

```terraform
variable "source_region" {
  default = "cn-hangzhou"
}
variable "destination_region" {
  default = "cn-shanghai"
}

provider "alicloud" {
  alias  = "hz"
  region = var.source_region
}
provider "alicloud" {
  alias  = "sh"
  region = var.destination_region
}

resource "alicloud_vpc" "example_hz" {
  provider   = alicloud.hz
  vpc_name   = "tf_example"
  cidr_block = "192.168.0.0/16"
}
resource "alicloud_vpc" "example_sh" {
  provider   = alicloud.sh
  vpc_name   = "tf_example"
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_cen_instance" "example" {
  cen_instance_name = "tf_example"
  description       = "an example for cen"
}
resource "alicloud_cen_instance_attachment" "example_hz" {
  instance_id              = alicloud_cen_instance.example.id
  child_instance_id        = alicloud_vpc.example_hz.id
  child_instance_type      = "VPC"
  child_instance_region_id = var.source_region
}
resource "alicloud_cen_instance_attachment" "example_sh" {
  instance_id              = alicloud_cen_instance.example.id
  child_instance_id        = alicloud_vpc.example_sh.id
  child_instance_type      = "VPC"
  child_instance_region_id = var.destination_region
}

resource "alicloud_cen_route_map" "default" {
  cen_region_id                          = var.source_region
  cen_id                                 = alicloud_cen_instance.example.id
  description                            = "tf_example"
  priority                               = "1"
  transmit_direction                     = "RegionIn"
  map_result                             = "Permit"
  next_priority                          = "1"
  source_region_ids                      = [var.source_region]
  source_instance_ids                    = [alicloud_cen_instance_attachment.example_hz.child_instance_id]
  source_instance_ids_reverse_match      = "false"
  destination_instance_ids               = [alicloud_cen_instance_attachment.example_sh.child_instance_id]
  destination_instance_ids_reverse_match = "false"
  source_route_table_ids                 = [alicloud_vpc.example_hz.route_table_id]
  destination_route_table_ids            = [alicloud_vpc.example_sh.route_table_id]
  source_child_instance_types            = ["VPC"]
  destination_child_instance_types       = ["VPC"]
  destination_cidr_blocks                = [alicloud_vpc.example_sh.cidr_block]
  cidr_match_mode                        = "Include"
  route_types                            = ["System"]
  match_asns                             = ["65501"]
  as_path_match_mode                     = "Include"
  match_community_set                    = ["65501:1"]
  community_match_mode                   = "Include"
  community_operate_mode                 = "Additive"
  operate_community_set                  = ["65501:1"]
  preference                             = "20"
  prepend_as_path                        = ["65501"]
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `direction` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `dry_run` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `name` - (Optional) This property does not have a description in the spec, please add it before generating code.
* `priority` - (Required, Int) This property does not have a description in the spec, please add it before generating code.
* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the resource group
* `tags` - (Optional, Map) Tags

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `route_map_id` - The first ID of the resource.
* `status` - 当前属性没有在镇元上录入属性描述�

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Route Map.
* `delete` - (Defaults to 5 mins) Used when delete the Route Map.
* `update` - (Defaults to 5 mins) Used when update the Route Map.

## Import

Cloud Enterprise Network (CEN) Route Map can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_route_map.example <route_map_id>
```