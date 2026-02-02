---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_global_database_network"
description: |-
  Provides a Alicloud Polardb Global Database Network resource.
---

# alicloud_polardb_global_database_network

Provides a Polardb Global Database Network resource.



For information about Polardb Global Database Network and how to use it, see [What is Global Database Network](https://www.alibabacloud.com/help/en/polardb/api-polardb-2017-08-01-createglobaldatabasenetwork).

-> **NOTE:** Available since v1.181.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_polardb_node_classes" "default" {
  db_type    = "MySQL"
  db_version = "8.0"
  category   = "Normal"
  pay_type   = "PostPaid"
}

resource "alicloud_vpc" "default" {
  vpc_name   = "terraform-example"
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vpc_id       = alicloud_vpc.default.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_polardb_node_classes.default.classes[0].zone_id
  vswitch_name = "terraform-example"
}

resource "alicloud_polardb_cluster" "default" {
  db_type       = "MySQL"
  db_version    = "8.0"
  db_node_class = data.alicloud_polardb_node_classes.default.classes.0.supported_engines.0.available_resources.0.db_node_class
  pay_type      = "PostPaid"
  vswitch_id    = alicloud_vswitch.default.id
  description   = "terraform-example"
}

resource "alicloud_polardb_global_database_network" "default" {
  db_cluster_id = alicloud_polardb_cluster.default.id
  description   = "terraform-example"
}
```

## Argument Reference

The following arguments are supported:
* `db_cluster_id` - (Required) The cluster ID in the GDN that needs to be switched to the primary cluster.  
You can call the [DescribeGlobalDatabaseNetwork](https://help.aliyun.com/document_detail/264580.html) operation to view the cluster IDs in the GDN.
* `forced` - (Optional, Available since v1.270.0) Specifies whether to forcibly switch the primary and secondary clusters in the GDN. Valid values:  
  - `true`: Forcibly switches the primary and secondary clusters in the GDN.  
  - `false`: Does not forcibly switch the primary and secondary clusters in the GDN.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `gdn_description` - (Optional, Available since v1.270.0) The description of the GDN. Requirements are as follows:
  - It must not start with http:// or https://.
  - It must start with a letter (uppercase or lowercase) or a Chinese character.
  - It can contain letters (uppercase or lowercase), Chinese characters, digits, underscores (_), or hyphens (-).
  - Its length must be between 2 and 126 characters.
* `resource_group_id` - (Optional, Computed, Available since v1.270.0) The ID of the resource group.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the GDN, in the format `YYYY-MM-DDThh:mm:ssZ` (UTC time).
* `status` - The status of the Global Database Network (GDN).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Global Database Network.
* `delete` - (Defaults to 5 mins) Used when delete the Global Database Network.
* `update` - (Defaults to 5 mins) Used when update the Global Database Network.

## Import

Polardb Global Database Network can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_global_database_network.example <gdn_id>
```