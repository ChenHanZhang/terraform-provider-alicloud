---
subcategory: "Tair (Redis OSS-Compatible) And Memcache (KVStore)"
layout: "alicloud"
page_title: "Alicloud: alicloud_kvstore_connection"
description: |-
  Provides a Alicloud Tair (Redis OSS-Compatible) And Memcache (KVStore) Connection resource.
---

# alicloud_kvstore_connection

Provides a Tair (Redis OSS-Compatible) And Memcache (KVStore) Connection resource.



For information about Tair (Redis OSS-Compatible) And Memcache (KVStore) Connection and how to use it, see [What is Connection](https://next.api.alibabacloud.com/document/R-kvstore/2015-01-01/AllocateInstancePublicConnection).

-> **NOTE:** Available since v1.101.0.

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

resource "alicloud_kvstore_connection" "default" {
  connection_string_prefix = "exampleconnection"
  instance_id              = alicloud_kvstore_instance.default.id
  port                     = "6370"
}
```

## Argument Reference

The following arguments are supported:
* `connection_string` - (Optional, Computed) Connection String
* `connection_string_prefix` - (Required) Connection String Prefix
* `instance_id` - (Required, ForceNew) Instance Id
* `port` - (Required) Port

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Create Time.
* `status` - The status of the instance.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Connection.
* `delete` - (Defaults to 5 mins) Used when delete the Connection.
* `update` - (Defaults to 5 mins) Used when update the Connection.

## Import

Tair (Redis OSS-Compatible) And Memcache (KVStore) Connection can be imported using the id, e.g.

```shell
$ terraform import alicloud_kvstore_connection.example <instance_id>
```