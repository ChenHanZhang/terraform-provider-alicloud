---
subcategory: "Express Connect Router"
layout: "alicloud"
page_title: "Alicloud: alicloud_express_connect_router_route_map_association"
description: |-
  Provides a Alicloud Express Connect Router Route Map Association resource.
---

# alicloud_express_connect_router_route_map_association

Provides a Express Connect Router Route Map Association resource.

Route map association.

For information about Express Connect Router Route Map Association and how to use it, see [What is Route Map Association](https://next.api.alibabacloud.com/document/ExpressConnectRouter/2023-09-01/CreateRouteMapAssociation).

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

variable "ecr_id" {
  default = "ecr-jl48tkimaa8g51o7nd"
}

variable "region_id_list" {
  default = "cn-wulanchabu-example-5,cn-wulanchabu-example-6"
}

variable "route_map_id" {
  default = "rm-yfs0jpvp487aj2dzrg"
}


resource "alicloud_express_connect_router_route_map_association" "default" {
  ecr_id         = var.ecr_id
  region_id_list = ["${var.region_id_list}"]
  route_map_id   = var.route_map_id
  dry_run        = false
}
```

## Argument Reference

The following arguments are supported:
* `dry_run` - (Optional) DryRun

-> **NOTE:** This parameter is only evaluated during resource creation, update and deletion. Modifying it in isolation will not trigger any action.

* `ecr_id` - (Required, ForceNew) Express Connect Router instance ID.
* `region_id_list` - (Optional, List) List of regions where the association takes effect.
* `route_map_id` - (Required, ForceNew) Route map instance ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<ecr_id>:<association_id>`.
* `association_id` - Route map association instance ID.
* `status` - The deployment status of the associated resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Route Map Association.
* `delete` - (Defaults to 5 mins) Used when delete the Route Map Association.
* `update` - (Defaults to 5 mins) Used when update the Route Map Association.

## Import

Express Connect Router Route Map Association can be imported using the id, e.g.

```shell
$ terraform import alicloud_express_connect_router_route_map_association.example <ecr_id>:<association_id>
```