---
subcategory: "AnalyticDB for MySQL (ADB)"
layout: "alicloud"
page_title: "Alicloud: alicloud_adb_db_cluster"
description: |-
  Provides a Alicloud AnalyticDB for MySQL (ADB) Db Cluster resource.
---

# alicloud_adb_db_cluster

Provides a AnalyticDB for MySQL (ADB) Db Cluster resource.

ADB Cluster Instance.

For information about AnalyticDB for MySQL (ADB) Db Cluster and how to use it, see [What is Db Cluster](https://www.alibabacloud.com/help/en/analyticdb/analyticdb-for-mysql/product-overview/what-is-analyticdb-for-mysql).

-> **NOTE:** Available since v1.121.0.

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
```

## Argument Reference

The following arguments are supported:
* `auto_renew_period` - (Optional, Computed, Int) AutoRenewPeriod

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `backup_set_id` - (Optional, Available since v1.271.0) Instance Backup set ID

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `compute_resource` - (Optional) ComputeResource
* `connection_string` - (Optional, Computed) ConnectionString

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_cluster_category` - (Required) Category
* `db_cluster_name` - (Optional, Available since v1.271.0) DBClusterName
* `db_cluster_network_type` - (Required, ForceNew, Available since v1.271.0) DBClusterNetworkType
* `db_cluster_version` - (Required, ForceNew) DBClusterVersion
* `db_node_class` - (Optional, Computed) DBNodeClass
* `db_node_count` - (Optional, Computed, Int) DBNodeCount
* `db_node_storage` - (Optional, Computed, Int) DBNodeStorage
* `db_cluster_ip_array_attribute` - (Optional, Available since v1.271.0) DbClusterIpArrayAttribute

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `db_cluster_ip_array_name` - (Optional, Available since v1.271.0) DbClusterIpArrayName

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `disk_encryption` - (Optional, ForceNew, Computed, Available since v1.219.0) Whether to Enable Cloud Disk encryption
* `disk_performance_level` - (Optional, Computed, Available since v1.207.0) Disk performance level
* `elastic_io_resource` - (Optional, Computed) ElasticIOResource
* `elastic_io_resource_size` - (Optional, Computed, Available since v1.207.0) EIU Specification
* `enable_ssl` - (Optional, Available since v1.230.0) Whether to enable SSL encryption
* `executor_count` - (Optional, Available since v1.271.0) ExecutorCount
* `kernel_version` - (Optional, Computed, Available since v1.240.0) Instance kernel minor version number
* `kms_id` - (Optional, ForceNew, Available since v1.219.0) Key service id
* `maintain_time` - (Optional, Computed) MaintainTime
* `mode` - (Required) Mode. Value Description:
Reserver: the reserved mode.
Flexible: elastic mode
* `modify_mode` - (Optional, Available since v1.271.0) ModifyMode

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `modify_type` - (Optional) Change type, value description:
Upgrade: Upgrade.
Downgrade: Downgrade.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `payment_type` - (Required) The paymen type of the resource
* `period` - (Optional) Specifies that the subscription cluster is of the annual or monthly subscription type. Value Description:
Year: type of package Year.
Month: the type of the Month.

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `period_unit` - (Optional, Available since v1.271.0) Unit of renewal duration. The default value is Month.
Year: Year.
Month: Month.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `renewal_status` - (Optional, Computed) RenewalStatus
* `resource_group_id` - (Optional, Computed) Alibaba Cloud Resource Group ID
* `restore_time` - (Optional, Available since v1.271.0) Instance backup and recovery time

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `restore_type` - (Optional, Available since v1.271.0) Reserved parameters, do not involve

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `security_ips` - (Optional, Computed) SecurityIps
* `source_db_instance_name` - (Optional, Available since v1.271.0) The id of the source instance of the data synchronization link.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `storage_resource` - (Optional, Available since v1.271.0) StorageResource
* `storage_type` - (Optional, Available since v1.271.0) Reserved parameters, not involved.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `switch_mode` - (Optional, Int, Available since v1.240.0) Execution time, value:
0: Execute immediately. Default value.
1: Execute in maintainable time period

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `tags` - (Optional, Map) The tag of the resource
* `used_time` - (Optional, Available since v1.271.0) The purchase duration of the prepaid cluster. Value Description:
When Period is set to Year, the value range of UsedTime is: 1, 2, 3, and 5 (integer).
When Period is set to Month, the value range of UsedTime is: 1~11 (integer)

-> **NOTE:** This parameter only applies during resource creation, update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `vswitch_id` - (Optional) VSwitchId
* `vpc_id` - (Optional, Computed, Available since v1.178.0) VPCId
* `zone_id` - (Optional, ForceNew, Computed) The zone ID  of the resource

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `pay_type` - Value Description:.
* `port` - Port.
* `region_id` - The region ID of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 65 mins) Used when create the Db Cluster.
* `delete` - (Defaults to 5 mins) Used when delete the Db Cluster.
* `update` - (Defaults to 134 mins) Used when update the Db Cluster.

## Import

AnalyticDB for MySQL (ADB) Db Cluster can be imported using the id, e.g.

```shell
$ terraform import alicloud_adb_db_cluster.example <db_cluster_id>
```