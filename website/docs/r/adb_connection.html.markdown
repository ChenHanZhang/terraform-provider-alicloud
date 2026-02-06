---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_connection"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Connection resource.
---

# alicloud_adb_connection

Provides a AnalyticDB for MySQL (ADB) Connection resource.

The access connection of the ADB instance. Users can access the ADB instance through the connection address.

For information about AnalyticDB for MySQL (ADB) Connection and how to use it, see [What is Connection](https://next.api.alibabacloud.com/document/adb/2019-03-15/AllocateClusterPublicConnection).

-> **NOTE:** Available since v1.81.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

data "alicloud_adb_zones" "default" {}

data "alicloud_vpcs" "default" {
  name_regex = "^default-NODELETING$"
}
data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_adb_zones.default.ids.0
}

locals {
  vswitch_id = data.alicloud_vswitches.default.ids.0
}

resource "alicloud_adb_db_cluster" "cluster" {
  db_cluster_category = "MixedStorage"
  mode                = "flexible"
  compute_resource    = "8Core32GB"
  vswitch_id          = local.vswitch_id
  description         = var.name
}

resource "alicloud_adb_connection" "default" {
  db_cluster_id     = alicloud_adb_db_cluster.cluster.id
  connection_prefix = "example"
}
```

## Argument Reference

The following arguments are supported:
* `connection_string` - (Optional, Computed) Connection string of the Data Warehouse Edition cluster to be modified.  

-> **NOTE:**  You can call the [DescribeDBClusterNetInfo](https://help.aliyun.com/document_detail/143384.html) operation to view the cluster's connection strings.


-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `connection_string_prefix` - (Optional, Available since v1.271.0) Prefix of the public endpoint.  
  - Must start with a lowercase letter and consist only of lowercase letters, digits, and hyphens (-).  
  - Must be no longer than 30 characters.  
  - Defaults to the cluster name.
* `db_cluster_id` - (Required, ForceNew) AnalyticDB for MySQL Data Warehouse Edition cluster ID.  

-> **NOTE:**  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to view the cluster IDs of all Data Warehouse Edition clusters in the target region.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `ip_address` - IP address.
* `port` - Port number, fixed as **3306**.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Connection.
* `delete` - (Defaults to 5 mins) Used when delete the Connection.
* `update` - (Defaults to 5 mins) Used when update the Connection.

## Import

AnalyticDB for MySQL (ADB) Connection can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_connection.example <db_cluster_id>
```