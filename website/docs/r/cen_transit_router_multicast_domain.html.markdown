---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_multicast_domain"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.
---

# alicloud_cen_transit_router_multicast_domain

Provides a Cloud Enterprise Network (CEN) Transit Router Multicast Domain resource.

Transit Router Multicast Domain.

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
* `cen_id` - (Optional, ForceNew, Available since v1.284.0) The ID of the Cloud Enterprise Network (CEN) instance.
* `options` - (Optional, Computed, Set, Available since v1.242.0) Feature options. See [`options`](#options) below.
* `tags` - (Optional, Map) The tag information.
You can specify up to 20 tags at a time.
* `transit_router_id` - (Optional, ForceNew) The ID of the transit router instance.
* `transit_router_multicast_domain_description` - (Optional) The description of the multicast domain.
The description can be empty or contain 1 to 256 characters. It cannot start with http:// or https://.
* `transit_router_multicast_domain_name` - (Optional) The name of the multicast domain.
The name can be empty or contain 1 to 128 characters, and must not start with http:// or https://.

### `options`

The options supports the following:
* `igmpv2_support` - (Optional, Computed) Indicates whether IGMP is enabled for the multicast domain. After this feature is enabled, hosts can dynamically join or leave multicast groups by using the Internet Group Management Protocol (IGMP). Valid values:
  - `enable`: enables IGMP.
  - `disable` (default): disables IGMP.

-> **NOTE:**  - The IGMP feature is in public preview. To use it, contact your account manager to apply for access.

-> **NOTE:**  - After IGMP is enabled, it cannot be disabled.

* `strict_source_control` - (Optional, Available since v1.284.0) Indicates whether strict multicast source control is enabled. If disabled, all ECS instances in the associated VSwitch can act as multicast sources. If enabled, only ENIs that are statically configured or have sent an IGMP Join message can act as multicast sources. Valid values:
  - ``enable``: enables strict multicast source control.
  - ``disable``: disables strict multicast source control.

-> **NOTE:**  - Strict multicast source control takes effect only for multicast domains with IGMP enabled.

-> **NOTE:**  - Currently, you can create only one multicast domain with strict multicast source control disabled under a single transit router.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - The region ID of the transit router instance.
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