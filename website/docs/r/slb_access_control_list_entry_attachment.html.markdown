---
subcategory: "Classic Load Balancer (SLB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_slb_access_control_list_entry_attachment"
description: |-
  Provides a Alicloud SLB Access Control List Entry Attachment resource.
---

# alicloud_slb_access_control_list_entry_attachment

Provides a SLB Access Control List Entry Attachment resource.

IP entry in the access control policy group.

For information about SLB Access Control List Entry Attachment and how to use it, see [What is Access Control List Entry Attachment](https://next.api.alibabacloud.com/document/Slb/2014-05-15/AddAccessControlListEntry).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-beijing"
}

resource "alicloud_slb_acl" "创建acl" {
}


resource "alicloud_slb_access_control_list_entry_attachment" "default" {
  comment = "example-111"
  entry   = "10.0.17.0/24"
  acl_id  = alicloud_slb_acl.创建acl.id
}
```

## Argument Reference

The following arguments are supported:
* `acl_id` - (Required, ForceNew) The ID of the access control policy group.
* `comment` - (Optional) Comments on access control entries.
* `entry` - (Required, ForceNew) The IP address of the access control entry.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<acl_id>:<entry>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Access Control List Entry Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Access Control List Entry Attachment.
* `update` - (Defaults to 5 mins) Used when update the Access Control List Entry Attachment.

## Import

SLB Access Control List Entry Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_slb_access_control_list_entry_attachment.example <acl_id>:<entry>
```