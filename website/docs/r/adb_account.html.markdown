---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_account"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Account resource.
---

# alicloud_adb_account

Provides a AnalyticDB for MySQL (ADB) Account resource.

ADB Account  .

For information about AnalyticDB for MySQL (ADB) Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/en/analyticdb-for-mysql/latest/api-doc-adb-2019-03-15-api-doc-createaccount).

-> **NOTE:** Available since v1.71.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform_example"
}

data "alicloud_adb_zones" "default" {
}

data "alicloud_vpcs" "default" {
  name_regex = "^default-NODELETING$"
}

data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_adb_zones.default.ids.0
}

resource "alicloud_adb_db_cluster" "cluster" {
  db_cluster_category = "MixedStorage"
  mode                = "flexible"
  compute_resource    = "8Core32GB"
  vswitch_id          = data.alicloud_vswitches.default.ids.0
  description         = var.name
}

resource "alicloud_adb_account" "default" {
  db_cluster_id       = alicloud_adb_db_cluster.cluster.id
  account_name        = var.name
  account_password    = "tf_example123"
  account_description = var.name
}
```

## Argument Reference

The following arguments are supported:
* `account_description` - (Optional) Modify the account description:
  - Must start with a Chinese character or an English letter.
  - Can contain Chinese characters, English letters, digits, underscores (_), and hyphens (-).
  - Cannot start with `http://` or `https://`.
  - Must be 2 to 256 characters in length.
* `account_name` - (Required, ForceNew) The database account name must meet the following requirements:  
  - It must start with a lowercase letter and end with a lowercase letter or digit.  
  - It can contain only lowercase letters, digits, or underscores (_).  
  - Its length must be between 2 and 16 characters.  
  - Reserved account names such as root, admin, and opsadmin cannot be used.  
* `account_password` - (Required) The password for the database account must meet the following requirements:  
  - It must consist of uppercase letters, lowercase letters, digits, and special characters.  
  - The allowed special characters are: (!), (@), (#), ($), (%), (^), (&), (*), (()), (_), (+), (-), (=).  
  - Its length must be between 8 and 32 characters.  

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `account_type` - (Optional, ForceNew, Computed, Available since v1.273.0) The account type. Valid values:
  - `Normal`: A standard account. You can create up to 256 standard accounts for a cluster.
  - `Super` (default): A privileged account. You can create only one privileged account for a cluster.

-> **NOTE:**  - If no account exists in the cluster, you can create either a privileged account or a standard account by calling the API. However, if a privileged account already exists in the cluster, you must specify Normal to successfully create a new account.

-> **NOTE:**  - After creation, the privileged account has permissions on all databases in the cluster. By default, a standard account has no permissions. You must manually grant permissions on specific databases to a standard account by using the privileged account. For more information, see [Grant Permissions to Users](https://help.aliyun.com/document_detail/123662.html).

* `db_cluster_id` - (Required, ForceNew) The cluster ID of the data warehouse edition.

-> **NOTE:**  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to view the cluster IDs of all data warehouse edition clusters in the target region.

* `tags` - (Optional, ForceNew, Map) The tag of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<db_cluster_id>:<account_name>`.
* `status` - A resource attribute field that indicates the resource status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Account.
* `delete` - (Defaults to 5 mins) Used when delete the Account.
* `update` - (Defaults to 5 mins) Used when update the Account.

## Import

AnalyticDB for MySQL (ADB) Account can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_account.example <db_cluster_id>:<account_name>
```