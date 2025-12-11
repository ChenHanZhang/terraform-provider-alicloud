---
subcategory: "Tair (Redis OSS-Compatible) And Memcache (KVStore)"
layout: "alicloud"
page_title: "Alicloud: alicloud_kvstore_backup_policy"
description: |-
  Provides a Alicloud Tair (Redis OSS-Compatible) And Memcache (KVStore) Backup resource.
---

# alicloud_kvstore_backup_policy

Provides a Tair (Redis OSS-Compatible) And Memcache (KVStore) Backup resource.

Instance level or database level backup objects.

For information about Tair (Redis OSS-Compatible) And Memcache (KVStore) Backup and how to use it, see [What is Backup](https://www.alibabacloud.com/help/en/redis/developer-reference/api-r-kvstore-2015-01-01-modifybackuppolicy-redis).

-> **NOTE:** Available since v1.15.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

data "alicloud_kvstore_zones" "default" {
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_kvstore_zones.default.zones.0.id
  vswitch_name = var.name
}

resource "alicloud_kvstore_instance" "default" {
  db_instance_name = var.name
  vswitch_id       = alicloud_vswitch.default.id
  zone_id          = data.alicloud_kvstore_zones.default.zones.0.id
  instance_class   = "redis.master.large.default"
  instance_type    = "Redis"
  engine_version   = "5.0"
  security_ips     = ["10.23.12.24"]
  config = {
    appendonly             = "yes"
    lazyfree-lazy-eviction = "yes"
  }
  tags = {
    Created = "TF",
    For     = "example",
  }
}

resource "alicloud_kvstore_backup_policy" "default" {
  instance_id   = alicloud_kvstore_instance.default.id
  backup_period = ["Tuesday", "Wednesday"]
  backup_time   = "10:00Z-11:00Z"
}
```

## Argument Reference

The following arguments are supported:
* `backup_retention_period` - (Optional, Int, Available since v1.266.0) 本次手动备份的过期时长，取值范围为 7~730 天。当您传入-1 时，表示本次手动备份数据不过期（实例生命周期内）；当您不传入任何值（默认情况），表示与当前自动备份策略一致。

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `instance_id` - (Required, ForceNew) InstanceId

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above.The value is formulated as `<instance_id>:<backup_id>`.
* `backup_id` - Backup ID.
* `status` - Backup status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 31 mins) Used when create the Backup.
* `delete` - (Defaults to 5 mins) Used when delete the Backup.

## Import

Tair (Redis OSS-Compatible) And Memcache (KVStore) Backup can be imported using the id, e.g.

```shell
$ terraform import alicloud_kvstore_backup_policy.example <instance_id>:<backup_id>
```