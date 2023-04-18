---
subcategory: "VPC"
layout: "alicloud"
page_title: "Alicloud: alicloud_havip"
sidebar_current: "docs-alicloud-resource-havip"
description: |-
  Provides a Alicloud Vpc Havip resource.
---

# alicloud_havip

Provides a Vpc Havip resource.

For information about Vpc Havip and how to use it, see [What is Havip](https://www.alibabacloud.com/help/en/).

-> **NOTE:** Available in v1.204.0+.

## Example Usage

Basic Usage

```terraform
resource "alicloud_havip" "default" {
  description = "test"
  vswitch_id  = alicloud_vswitch.defaultVswitch.vswitch_id
  ha_vip_name = "tf-testacc-chenyi"
  ip_address  = "192.168.1.101"
  region_id   = "cn-shanghai"
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Dependence of a HaVip instance.
* `ha_vip_name` - (Optional) The name of the HaVip instance
* `ip_address` - (ForceNew,Computed,Optional) IP address of private network
* `resource_group_id` - (Computed,Optional) The ID of the resource group to which the VPC belongs.
* `tags` - (Optional) The tags of PrefixList.See the following `Block Tags`.
* `vswitch_id` - (Required,ForceNew) The switch ID to which the HaVip instance belongs

The following arguments will be discarded. Please use new fields as soon as possible:
* `havip_name` - Field 'havip_name' has been deprecated from provider version 1.204.0. New field 'ha_vip_name' instead.

#### Block Tags

The Tags supports the following:
* `tag_key` - (Optional) The key of tag.
* `tag_value` - (Optional) The value of tag.



## Attributes Reference

The following attributes are exported:
* `id` - The `key` of the resource supplied above.
* `associated_eip_addresses` - EIP bound to HaVip
* `associated_instance_type` - The type of the instance that is bound to the VIIP. Value:-**EcsInstance**: ECS instance.-**NetworkInterface**: ENI instance.
* `associated_instances` - An ECS instance that is bound to HaVip
* `create_time` - The creation time of the  resource
* `ha_vip_id` - The  ID of the resource
* `ip_address` - IP address of private network
* `master_instance_id` - The primary instance ID bound to HaVip
* `resource_group_id` - The ID of the resource group to which the VPC belongs.
* `status` - The status
* `vpc_id` - The VPC ID to which the HaVip instance belongs

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration-0-11/resources.html#timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Havip.
* `delete` - (Defaults to 5 mins) Used when delete the Havip.
* `update` - (Defaults to 5 mins) Used when update the Havip.

## Import

Vpc Havip can be imported using the id, e.g.

```shell
$ terraform import alicloud_vpc_havip.example 
```