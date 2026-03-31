---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_policy_based_route"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Policy Based Route resource.
---

# alicloud_cen_transit_router_policy_based_route

Provides a Cloud Enterprise Network (CEN) Transit Router Policy Based Route resource.

Transit Router Policy Based Route.

For information about Cloud Enterprise Network (CEN) Transit Router Policy Based Route and how to use it, see [What is Transit Router Policy Based Route](https://next.api.alibabacloud.com/document/Cbn/2017-09-12/CreateTransitRouterPolicyBasedRoute).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_cen_instance" "defaultO0XU5Y" {
}

resource "alicloud_cen_transit_router" "defaultF4twp9" {
  cen_id = alicloud_cen_instance.defaultO0XU5Y.id
}

resource "alicloud_cen_transit_router_policy_table" "defaultEN72S0" {
  transit_router_id = alicloud_cen_transit_router.defaultF4twp9.transit_router_id
}

resource "alicloud_cen_transit_router_route_table" "default9xsUKi" {
  transit_router_id = alicloud_cen_transit_router.defaultF4twp9.transit_router_id
}


resource "alicloud_cen_transit_router_policy_based_route" "default" {
  description           = "路由描述"
  priority              = "50"
  name                  = "路由名称"
  policy_table_id       = alicloud_cen_transit_router_policy_table.defaultEN72S0.id
  dry_run               = false
  source_port_range     = "1/10"
  source_cidr           = "192.168.1.0/24"
  target_route_table_id = alicloud_cen_transit_router_route_table.default9xsUKi.transit_router_route_table_id
  dest_port_range       = "200/2000"
  address_family        = "IPv4"
  protocol              = "ALL"
  dscp                  = "23/24"
  destination_cidr      = "192.168.1.0/24"
}
```

## Argument Reference

The following arguments are supported:
* `address_family` - (Optional, ForceNew) The creation time of the resource
* `description` - (Optional) The ID of the resource group
* `dest_port_range` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `destination_cidr` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `dry_run` - (Optional) This property does not have a description in the spec, please add it before generating code.

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `dscp` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `name` - (Optional) The name of the resource
* `policy_table_id` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `priority` - (Required, ForceNew, Int) This property does not have a description in the spec, please add it before generating code.
* `protocol` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `source_cidr` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `source_port_range` - (Optional, ForceNew) This property does not have a description in the spec, please add it before generating code.
* `target_route_table_id` - (Required, ForceNew) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - 当前属性没有在镇元上录入属性描述�

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Policy Based Route.
* `delete` - (Defaults to 5 mins) Used when delete the Transit Router Policy Based Route.
* `update` - (Defaults to 5 mins) Used when update the Transit Router Policy Based Route.

## Import

Cloud Enterprise Network (CEN) Transit Router Policy Based Route can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_policy_based_route.example <policy_based_route_id>
```