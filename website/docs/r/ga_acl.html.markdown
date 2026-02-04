---
subcategory: "Global Accelerator (GA)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ga_acl"
description: |-
  Provides a Alicloud Global Accelerator (GA) Acl resource.
---

# alicloud_ga_acl

Provides a Global Accelerator (GA) Acl resource.



For information about Global Accelerator (GA) Acl and how to use it, see [What is Acl](https://www.alibabacloud.com/help/en/global-accelerator/latest/api-ga-2019-11-20-createacl).

-> **NOTE:** Available since v1.150.0.

## Example Usage

Basic Usage

```terraform
resource "alicloud_ga_acl" "default" {
  acl_name           = "terraform-example"
  address_ip_version = "IPv4"
}

resource "alicloud_ga_acl_entry_attachment" "default" {
  acl_id            = alicloud_ga_acl.default.id
  entry             = "192.168.1.1/32"
  entry_description = "terraform-example"
}
```

## Argument Reference

The following arguments are supported:
* `acl_entries` - (Optional, ForceNew, Computed, List) The entries of the Acl See [`acl_entries`](#acl_entries) below.
* `acl_name` - (Optional) The name of the acl
* `address_ip_version` - (Required, ForceNew) ipVersion
* `resource_group_id` - (Optional, Computed, Available since v1.226.0) The ID of the resource group
* `tags` - (Optional, Map, Available since v1.207.1) The tag of the resource

### `acl_entries`

The acl_entries supports the following:
* `entry` - (Optional, ForceNew, Computed) acl entry
* `entry_description` - (Optional, ForceNew, Computed) The description of the entry

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Acl.
* `delete` - (Defaults to 5 mins) Used when delete the Acl.
* `update` - (Defaults to 8 mins) Used when update the Acl.

## Import

Global Accelerator (GA) Acl can be imported using the id, e.g.

```shell
$ terraform import alicloud_ga_acl.example <acl_id>
```