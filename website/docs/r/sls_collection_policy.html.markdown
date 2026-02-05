---
subcategory: "Log Service (SLS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_sls_collection_policy"
description: |-
  Provides a Alicloud Log Service (SLS) Collection Policy resource.
---

# alicloud_sls_collection_policy

Provides a Log Service (SLS) Collection Policy resource.

Orchestration policies for cloud product log collection.

For information about Log Service (SLS) Collection Policy and how to use it, see [What is Collection Policy](https://www.alibabacloud.com/help/zh/sls/developer-reference/api-sls-2020-12-30-upsertcollectionpolicy).

-> **NOTE:** Available since v1.232.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_log_project" "project_create_01" {
  description  = var.name
  project_name = format("%s1%s", var.name, random_integer.default.result)
}

resource "alicloud_log_store" "logstore_create_01" {
  retention_period = "30"
  shard_count      = "2"
  project_name     = alicloud_log_project.project_create_01.project_name
  logstore_name    = format("%s1%s", var.name, random_integer.default.result)
}

resource "alicloud_log_project" "update_01" {
  description  = var.name
  project_name = format("%s2%s", var.name, random_integer.default.result)
}

resource "alicloud_log_store" "logstore002" {
  retention_period = "30"
  shard_count      = "2"
  project_name     = alicloud_log_project.update_01.project_name
  logstore_name    = format("%s2%s", var.name, random_integer.default.result)
}


resource "alicloud_sls_collection_policy" "default" {
  policy_config {
    resource_mode = "all"
    regions       = ["cn-hangzhou"]
  }
  data_code          = "metering_log"
  centralize_enabled = true
  product_code       = "oss"
  policy_name        = "xc-example-oss-01"
  enabled            = true
  data_config {
    data_region = "cn-hangzhou"
  }
  centralize_config {
    dest_ttl      = "3"
    dest_region   = "cn-shanghai"
    dest_project  = alicloud_log_project.project_create_01.project_name
    dest_logstore = alicloud_log_store.logstore_create_01.logstore_name
  }
  resource_directory {
    account_group_type = "custom"
    members            = ["1936728897040477"]
  }
}
```

## Argument Reference

The following arguments are supported:
* `centralize_config` - (Optional, Computed, Set) Centralized transfer configuration. See [`centralize_config`](#centralize_config) below.
* `centralize_enabled` - (Optional) Specifies whether to enable centralized storage. Default value: false.
* `data_code` - (Required, ForceNew) The code of the log type.
* `data_config` - (Optional, ForceNew, Computed, Set) The configuration is supported only when the log type is global. For example, if the productCode is sls, global logs will be collected to the corresponding region during the first configuration. See [`data_config`](#data_config) below.
* `enabled` - (Required) Whether to open.
* `policy_config` - (Required, Set) Collection rule configuration. See [`policy_config`](#policy_config) below.
* `policy_name` - (Required, ForceNew) The name of the rule, with a minimum of 3 characters and a maximum of 63 characters, must start with a letter.
* `product_code` - (Required, ForceNew) The code of the service.
* `resource_directory` - (Optional, Computed, Set) For Resource Directory configuration, the account must have opened the resource directory and be an administrator or a delegated administrator. See [`resource_directory`](#resource_directory) below.

### `centralize_config`

The centralize_config supports the following:
* `dest_logstore` - (Optional) When the central logstore is transferred to the destination logstore, its geographical attribute should be consistent with the destRegion and belong to the destProject.
* `dest_project` - (Optional) The destination project for centralized storage. Make sure that the region of the destination project is consistent with the region specified by destRegion.
* `dest_region` - (Optional) Centralized transfer destination area.
* `dest_ttl` - (Optional, Int) The data retention period for centralized storage. Unit: days. This parameter takes effect only when you use an existing logstore for centralized storage.

### `data_config`

The data_config supports the following:
* `data_region` - (Optional, ForceNew) If and only if the log type is global log type, for example, if productCode is sls, global logs will be collected to the corresponding region during the first configuration.

### `policy_config`

The policy_config supports the following:
* `instance_ids` - (Optional, List) A collection of instance IDs, valid only if resourceMode is instanceMode. Only instances whose instance ID is in the instance ID collection are collected.
* `regions` - (Optional, List) The region collection to which the instance belongs. Valid only when resourceMode is set to attributeMode. Wildcard characters are supported. If the region collection filter item is an empty array, it means that you do not need to filter by region, and all instances meet the filtering condition of the region collection. Otherwise, only instances with region attributes in the region collection are collected. The region collection and resource label of the instance. The instance objects are collected only when all of them are met.
* `resource_mode` - (Required) The resource collection mode. Valid values: all, attributeMode, and instanceMode. The value all specifies that logs of all instances within your account are collected to the default logstore. The value attributeMode specifies that logs are collected based on the regions of instances and resource tags. The value instanceMode specifies that logs are collected based on instance IDs.
* `resource_tags` - (Optional, Map) The resource tags. This parameter takes effect only when resourceMode is set to attributeMode. If you leave this parameter empty, resource tag-based filtering is not performed. The system considers that all instances are matched. If you specify a value for this parameter, logs of instances that use the specified resource tags are collected. Logs are collected from an instance only if the resource tags and region of the instance match the specified conditions.

### `resource_directory`

The resource_directory supports the following:
* `account_group_type` - (Optional) Support all mode all and custom mode custom under this resource directory
* `members` - (Optional, List) When the resource directory is configured in the custom mode, the corresponding member account list

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `data_config` - The configuration is supported only when the log type is global.
  * `data_project` - Valid only when the log type is global.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Collection Policy.
* `delete` - (Defaults to 5 mins) Used when delete the Collection Policy.
* `update` - (Defaults to 5 mins) Used when update the Collection Policy.

## Import

Log Service (SLS) Collection Policy can be imported using the id, e.g.

```shell
$ terraform import alicloud_sls_collection_policy.example <policy_name>
```