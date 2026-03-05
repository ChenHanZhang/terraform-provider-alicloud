---
subcategory: "Tair (Redis OSS-Compatible) And Memcache (KVStore)"
layout: "alicloud"
page_title: "Alicloud: alicloud_kvstore_audit_log_config"
description: |-
  Provides a Alicloud Tair (Redis OSS-Compatible) And Memcache (KVStore) Audit Log Config resource.
---

# alicloud_kvstore_audit_log_config

Provides a Tair (Redis OSS-Compatible) And Memcache (KVStore) Audit Log Config resource.

Set the switch and retention duration of the database instance audit log.

For information about Tair (Redis OSS-Compatible) And Memcache (KVStore) Audit Log Config and how to use it, see [What is Audit Log Config](https://next.api.alibabacloud.com/document/R-kvstore/2015-01-01/ModifyAuditLogConfig).

-> **NOTE:** Available since v1.130.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf-example"
}
data "alicloud_kvstore_zones" "default" {

}
data "alicloud_resource_manager_resource_groups" "default" {
  status = "OK"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}
resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_kvstore_zones.default.zones.0.id
}

resource "alicloud_kvstore_instance" "default" {
  db_instance_name  = var.name
  vswitch_id        = alicloud_vswitch.default.id
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  zone_id           = data.alicloud_kvstore_zones.default.zones.0.id
  instance_class    = "redis.master.large.default"
  instance_type     = "Redis"
  engine_version    = "5.0"
  security_ips      = ["10.23.12.24"]
  config = {
    appendonly             = "yes"
    lazyfree-lazy-eviction = "yes"
  }
  tags = {
    Created = "TF",
    For     = "example",
  }
}

resource "alicloud_kvstore_audit_log_config" "example" {
  instance_id = alicloud_kvstore_instance.default.id
  db_audit    = true
  retention   = 1
}
```

### Deleting `alicloud_kvstore_audit_log_config` or removing it from your configuration

Terraform cannot destroy resource `alicloud_kvstore_audit_log_config`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `db_audit` - (Optional) Indicates Whether to Enable the Audit Log Value: True: Default Value, Open. False: Closed. Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn't Support Separate Open.
* `instance_id` - (Required, ForceNew) Instance ID, Call the Describeinstances Get.
* `retention` - (Optional, Int) Audit Log Retention Period Value: 1~365. Note: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Instance Creation Time.
* `region_id` - The region ID of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 11 mins) Used when create the Audit Log Config.
* `update` - (Defaults to 5 mins) Used when update the Audit Log Config.

## Import

Tair (Redis OSS-Compatible) And Memcache (KVStore) Audit Log Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_kvstore_audit_log_config.example <instance_id>
```