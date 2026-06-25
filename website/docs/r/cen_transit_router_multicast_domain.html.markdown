---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_multicast_domain"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.
---

# alicloud_cen_transit_router_multicast_domain

Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.



For information about Cloud Enterprise Network (CEN) Transit Router Multicast Domain and how to use it, see [What is Transit Router Multicast Domain](https://www.alibabacloud.com/help/en/cen/developer-reference/api-cbn-2017-09-12-createtransitroutermulticastdomain).

-> **NOTE:** Available since v1.195.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

resource "alicloud_cen_instance" "example" {
  cen_instance_name = var.name
}

resource "alicloud_cen_transit_router" "example" {
  transit_router_name = var.name
  cen_id              = alicloud_cen_instance.example.id
  support_multicast   = true
}

resource "alicloud_cen_transit_router_multicast_domain" "default" {
  transit_router_id                           = alicloud_cen_transit_router.example.transit_router_id
  transit_router_multicast_domain_name        = var.name
  transit_router_multicast_domain_description = var.name
  options {
    igmpv2_support = "disable"
  }
}
```

## Argument Reference

The following arguments are supported:
* `cen_id` - (Optional, ForceNew, Available since v1.283.0) CenId
* `options` - (Optional, Computed, Set, Available since v1.242.0) Options See [`options`](#options) below.
* `tags` - (Optional, Map) The tag of the resource
* `transit_router_id` - (Optional, ForceNew) The ID of the forwarding router instance.
* `transit_router_multicast_domain_description` - (Optional) The description of the multicast domain.
* `transit_router_multicast_domain_name` - (Optional) The name of the multicast domain.

### `options`

The options supports the following:
* `igmpv2_support` - (Optional, Computed) Igmpv2Support

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The ID of the region to which the forwarding router instance belongs.
* `status` - The status of the multicast domain.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Multicast Domain.
* `delete` - (Defaults to 5 mins) Used when delete the Transit Router Multicast Domain.
* `update` - (Defaults to 5 mins) Used when update the Transit Router Multicast Domain.

## Import

Cloud Enterprise Network (CEN) Transit Router Multicast Domain can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_multicast_domain.example <transit_router_multicast_domain_id>
```