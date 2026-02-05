---
subcategory: "MongoDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_mongodb_backup"
description: |-
  Provides a Alicloud Mongodb Backup resource.
---

# alicloud_mongodb_backup

Provides a Mongodb Backup resource.

Instance-level or database-level backup objects.

For information about Mongodb Backup and how to use it, see [What is Backup](https://next.api.alibabacloud.com/document/Dds/2015-12-01/CreateBackup).

-> **NOTE:** Available since v1.271.0.

## Example Usage

Basic Usage

```terraform
variable "zone_id" {
  default = "cn-shanghai-n"
}

variable "region_id" {
  default = "cn-shanghai"
}

variable "other_zone_id" {
  default = "cn-shanghai-g"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "10.0.0.0/8"
  vpc_name   = "bgg-vpc-shanghai-n"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id     = alicloud_vpc.defaultVpc.id
  zone_id    = var.zone_id
  cidr_block = "10.0.0.0/24"
}

resource "alicloud_vswitch" "defaulteotherVswitch" {
  vpc_id     = alicloud_vpc.defaultVpc.id
  zone_id    = var.other_zone_id
  cidr_block = "10.1.0.0/24"
}

resource "alicloud_security_group" "defaultSecurityGroup" {
  vpc_id = alicloud_vpc.defaultVpc.id
}


```

## Argument Reference

The following arguments are supported:
* `backup_method` - (Optional, ForceNew) Backup Method
* `backup_retention_period` - (Optional, Int) Backup retention days.

-> **NOTE:**  No transmission means consistent with the default backup policy. When passing, can pass

-> **NOTE:** - 7-730 days

-> **NOTE:** - - 1 (Long-term retention)


-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `db_instance_id` - (Required, ForceNew) DB Instance Id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_instance_id>:<backup_id>`.
* `backup_id` - Backup Id.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 23 mins) Used when create the Backup.
* `delete` - (Defaults to 5 mins) Used when delete the Backup.

## Import

Mongodb Backup can be imported using the id, e.g.

```shell
$ terraform import alicloud_mongodb_backup.example <db_instance_id>:<backup_id>
```