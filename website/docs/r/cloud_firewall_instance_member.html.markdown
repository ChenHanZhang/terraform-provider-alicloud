---
subcategory: "Cloud Firewall"
layout: "alicloud"
page_title: "Alicloud: alicloud_cloud_firewall_instance_member"
description: |-
  Provides a Alicloud Cloud Firewall Instance Member resource.
---

# alicloud_cloud_firewall_instance_member

Provides a Cloud Firewall Instance Member resource.

Cloud Firewall Member Account.

For information about Cloud Firewall Instance Member and how to use it, see [What is Instance Member](https://www.alibabacloud.com/help/en/cloud-firewall/cloudfirewall/developer-reference/api-cloudfw-2017-12-07-addinstancemembers).

-> **NOTE:** Available since v1.194.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "AliyunTerraform"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_resource_manager_account" "default" {
  display_name = "${var.name}-${random_integer.default.result}"
  timeouts {
    delete = "5m"
  }
}

resource "alicloud_cloud_firewall_instance_member" "default" {
  member_desc = "${var.name}-${random_integer.default.result}"
  member_uid  = alicloud_resource_manager_account.default.id
}
```

## Argument Reference

The following arguments are supported:
* `member_desc` - (Optional) The description of the Cloud Firewall member account.
* `member_uid` - (Required, ForceNew, Int) The UID of the Cloud Firewall member account. You can add up to 20 member accounts.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the Cloud Firewall member account was added.
* `member_display_name` - The name of the Cloud Firewall member account.
* `modify_time` - The last modification time of the Cloud Firewall member account.
* `status` - The status of the Cloud Firewall member account.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Instance Member.
* `delete` - (Defaults to 5 mins) Used when delete the Instance Member.
* `update` - (Defaults to 5 mins) Used when update the Instance Member.

## Import

Cloud Firewall Instance Member can be imported using the id, e.g.

```shell
$ terraform import alicloud_cloud_firewall_instance_member.example <member_uid>
```