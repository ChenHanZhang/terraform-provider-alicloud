---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_address_book"
description: |-
  Provides a Alicloud Cloud Firewall Address Book resource.
---

# alicloud_cloud_firewall_address_book

Provides a Cloud Firewall Address Book resource.

Security access control address books, including IPv4 address books, ECS tag address books, IPv6 address books, port address books, and domain name address books.

For information about Cloud Firewall Address Book and how to use it, see [What is Address Book](https://www.alibabacloud.com/help/en/cloud-firewall/developer-reference/api-cloudfw-2017-12-07-addaddressbook).

-> **NOTE:** Available since v1.178.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_cloud_firewall_address_book" "example" {
  description      = "example_value"
  group_name       = "example_value"
  group_type       = "tag"
  tag_relation     = "and"
  auto_add_tag_ecs = 0
  ecs_tags {
    tag_key   = "created"
    tag_value = "tfTestAcc0"
  }
}
```

## Argument Reference

The following arguments are supported:
* `address_list` - (Optional, List) Returns the address book table.
* `asset_member_uids` - (Optional, List, Available since v1.285.0) Asset address book, member account list.
* `asset_region_resource_types` - (Optional, List, Available since v1.285.0) Asset address book, region and resource type list. See [`asset_region_resource_types`](#asset_region_resource_types) below.
* `auto_add_tag_ecs` - (Optional, Int) Whether you want to automatically add new matching tags of the ECS IP address to the address book.
  - `1`: the automatically added
  - `0`: indicates that does not automatically add
* `description` - (Required) After the description of.
* `group_name` - (Required) Address book name.
* `group_type` - (Required, ForceNew) Address book type, optional values include:
  - `ip`:IPv4 address book
  - `domain`: domain Address Book
  - `port`: port Address Book
  - `tag`:ECS tag address book
  - `ipv6`:IPv6 address book
* `lang` - (Optional) The language type of the received message.
  - `zh` (default): Chinese.
  - `en`: English.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `tag_relation` - (Optional, Computed) One or more tags for the relationship between.
  - `and`: the one or more tags for the Inter-as "and" relationship, that is to say, in the meantime matching the plurality of tags of the ECS IP address will be added to the address book.
  - `or`: a plurality of inter-tag "or" relationship, that is, as long as the matching one of the tags of the ECS IP address will be added to the address book.
* `tags` - (Optional, Map, Available since v1.285.0) ECS tags

### `asset_region_resource_types`

The asset_region_resource_types supports the following:
* `asset_region_id` - (Optional, ForceNew, Available since v1.285.0) Asset region ID.
* `resource_type` - (Optional, ForceNew, Set, Available since v1.285.0) Resource type. See [`resource_type`](#asset_region_resource_types-resource_type) below.

### `asset_region_resource_types-resource_type`

The asset_region_resource_types-resource_type supports the following:
* `ipv4` - (Optional, ForceNew, Set, Available since v1.285.0) IPv4 resource type. See [`ipv4`](#asset_region_resource_types-resource_type-ipv4) below.
* `ipv6` - (Optional, ForceNew, Set, Available since v1.285.0) IPv6 resource type. See [`ipv6`](#asset_region_resource_types-resource_type-ipv6) below.

### `asset_region_resource_types-resource_type-ipv4`

The asset_region_resource_types-resource_type-ipv4 supports the following:
* `ai_gateway_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `alb_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `api_gateway_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `bastion_host_egress_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `bastion_host_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `bastion_host_ingress_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `eip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `ecs_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `ecs_public_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `eni_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `ga_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `hav_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `nat_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `nat_public_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `nlb_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `slb_e_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `slb_public_ip` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.

### `asset_region_resource_types-resource_type-ipv6`

The asset_region_resource_types-resource_type-ipv6 supports the following:
* `ai_gateway_e_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `alb_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `api_gateway_e_ipv6` - (Optional, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `ecs_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `eni_e_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `ga_e_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `nlb_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.
* `slb_ipv6` - (Optional, ForceNew, Available since v1.285.0) This property does not have a description in the spec, please add it before generating code.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `address_list_count` - The returned address book list address number.
* `reference_count` - The address book is reference.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Address Book.
* `delete` - (Defaults to 5 mins) Used when delete the Address Book.
* `update` - (Defaults to 5 mins) Used when update the Address Book.

## Import

Cloud Firewall Address Book can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_address_book.example <group_uuid>
```