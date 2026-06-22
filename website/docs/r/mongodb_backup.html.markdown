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

-> **NOTE:** Available since v1.283.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

variable "zone_id" {
  default = "cn-shanghai-b"
}

variable "region_id" {
  default = "cn-shanghai"
}

variable "ipv4网段-b" {
  default = "10.0.0.0/24"
}

resource "alicloud_vpc" "defaulttAyVQI" {
  cidr_block = "10.0.0.0/8"
  vpc_name   = "bgg-vpc-shanghai-b"
}

resource "alicloud_vswitch" "defaultGQF7Gl" {
  vpc_id     = alicloud_vpc.defaulttAyVQI.id
  zone_id    = var.zone_id
  cidr_block = var.ipv4网段-b
}

resource "alicloud_mongodb_instance" "defaultpUQD63" {
  engine_version      = "5.0"
  storage_type        = "cloud_essd1"
  vswitch_id          = alicloud_vswitch.defaultGQF7Gl.id
  db_instance_storage = "20"
  vpc_id              = alicloud_vpc.defaulttAyVQI.id
  db_instance_class   = "mdb.shard.4x.large.d"
  storage_engine      = "WiredTiger"
  network_type        = "VPC"
  zone_id             = var.zone_id
  replication_factor  = "3"
  readonly_replicas   = "0"
}


resource "alicloud_mongodb_backup" "default" {
  backup_method           = "Snapshot"
  db_instance_id          = alicloud_mongodb_instance.defaultpUQD63.id
  backup_retention_period = "7"
}
```

## Argument Reference

The following arguments are supported:
* `backup_method` - (Optional, ForceNew) Backup Method
* `backup_retention_period` - (Optional, Int) Backup retention days.

-> **NOTE:**  No transmission means consistent with the default backup policy. When passing, can pass

-> **NOTE:** - 7-730 days

-> **NOTE:** - - 1 (Long-term retention)


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `db_instance_id` - (Required, ForceNew) DB Instance Id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_instance_id>:<backup_id>`.
* `backup_db_names` - Backup DB Names.
* `backup_download_url` - Backup Download URL.
* `backup_id` - Backup Id.
* `backup_intranet_download_url` - Backup Intranet DownloadURL.
* `backup_job_id` - The backup task ID.
* `backup_mode` - Backup Mode.
* `backup_size` - Backup Size.
* `backup_type` - Backup Type.
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