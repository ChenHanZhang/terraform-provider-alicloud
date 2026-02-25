---
subcategory: "PolarDB"
layout: "alicloud"
page_title: "Alicloud: alicloud_polardb_application"
description: |-
  Provides a Alicloud Polardb Application resource.
---

# alicloud_polardb_application

Provides a Polardb Application resource.

PolarDB AI Application.

For information about Polardb Application and how to use it, see [What is Application](https://next.api.alibabacloud.com/document/polardb/2017-08-01/CreateApplication).

-> **NOTE:** Available since v1.272.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

### Deleting `alicloud_polardb_application` or removing it from your configuration

The `alicloud_polardb_application` resource allows you to manage  `payment_type = "PayAsYouGo"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `application_type` - (Required, ForceNew) Application type. Valid values:
  - supabase: Specify this value when creating a managed Supabase application.
  - raycluster: Specify this value when creating a managed Ray Cluster application.
* `architecture` - (Required, ForceNew) The CPU architecture. Valid values:  
  - x86  
* `auto_create_polar_fs` - (Optional) Specifies whether to automatically create a cold-storage PolarFS instance. Valid values:
  - false (default): Do not automatically create.
  - true: Automatically create.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `auto_renew` - (Optional) Automatic renewal.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `auto_use_coupon` - (Optional) Specifies whether to automatically apply a coupon. Valid values:
  - true (default): Apply a coupon.
  - false: Do not apply a coupon.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `components` - (Optional, ForceNew, List) A list of user-defined application subcomponents.   See [`components`](#components) below.
* `db_cluster_id` - (Required, ForceNew) The PolarDB cluster ID. If specified, only application information related to this PolarDB cluster is returned.
* `description` - (Optional) The description or notes for the application.
* `endpoints` - (Optional, Computed, List) List of user-defined service endpoints. By default, a VPC endpoint is created. See [`endpoints`](#endpoints) below.
* `modify_mode` - (Optional) The modification mode for the IP whitelist. Valid values:  
  - `Cover`: Overwrites the existing IP whitelist (default).  
  - `Append`: Adds new IP addresses.  
  - `Delete`: Removes IP addresses.  

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `payment_type` - (Required, ForceNew) Billing type.
* `period` - (Optional) Subscription period type (yearly or monthly).

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `polar_fs_instance_id` - (Optional, ForceNew, Computed) The instance ID of PolarFS Cold Storage Edition or High-Performance Edition. This field is empty by default. If specified, the corresponding storage will be mounted to the application.

Currently, only the following applications are supported:  
  - supabase  
  - raycluster.  
* `promotion_code` - (Optional) Coupon code. If left blank, the default coupon is used.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `security_groups` - (Optional, List) A list of security groups at the application level. See [`security_groups`](#security_groups) below.
* `security_ip_arrays` - (Optional, Computed, List) The application-level whitelist. See [`security_ip_arrays`](#security_ip_arrays) below.
* `tags` - (Optional, Map) List of tags.
* `used_time` - (Optional) Duration.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `vswitch_id` - (Optional, ForceNew, Computed) The vSwitch ID. The default value is the current vSwitch in the primary zone of the instance.
* `vpc_id` - (Optional, ForceNew) VPC ID.
* `zone_id` - (Optional, ForceNew, Computed) The zone ID. The default value is the primary zone of the instance.  

### `components`

The components supports the following:
* `component_class` - (Optional, ForceNew) Specification of the application subcomponent.
* `component_max_replica` - (Optional, ForceNew, Int) The maximum number of replicas for the application subcomponent.  
* `component_replica` - (Optional, ForceNew, Int) The number of replicas for the application subcomponent.  
* `component_type` - (Optional, ForceNew) The type of the application subcomponent.  
* `scale_max` - (Optional) The upper scaling limit per node. Valid values: 0 PCU to 16 PCU.  
* `scale_min` - (Optional) The lower scaling limit per node. Valid values: 0 PCU to 16 PCU.  
* `security_groups` - (Optional, ForceNew, List) A list of security groups at the subcomponent level.   See [`security_groups`](#components-security_groups) below.
* `security_ip_arrays` - (Optional, ForceNew, List) A list of allowlist IP addresses at the subcomponent level. See [`security_ip_arrays`](#components-security_ip_arrays) below.

### `components-security_groups`

The components-security_groups supports the following:

### `components-security_ip_arrays`

The components-security_ip_arrays supports the following:
* `security_ip_array_name` - (Optional, ForceNew) Name of the IP address group. Default value: default.  
* `security_ip_list` - (Optional, ForceNew) Whitelist IP addresses, separated by commas.  
* `security_ip_net_type` - (Optional, ForceNew) Network type of the whitelist IP addresses. Default value: mix.

### `endpoints`

The endpoints supports the following:
* `net_type` - (Optional, Computed) Type of connection address  
  - Private: VPC address  
  - Public: Public network address

### `security_groups`

The security_groups supports the following:
* `security_group_id` - (Optional) The ID of the security group.

### `security_ip_arrays`

The security_ip_arrays supports the following:
* `security_ip_array_name` - (Optional, Computed) The name of the IP address group. Default value: default.
* `security_ip_list` - (Optional, Computed) Whitelist IP addresses, separated by commas.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `components` - A list of user-defined application subcomponents.
  * `component_class_description` - Specification description of the application subcomponent.
  * `component_id` - Subcomponent ID.
  * `component_replica_group_name` - Group name of the replicas of the application subcomponent.
  * `security_groups` - A list of security groups at the subcomponent level.
    * `net_type` - The network type.
    * `region_id` - The region ID.
    * `security_group_id` - The ID of the security group.
    * `security_group_name` - The name of the security group.
  * `security_ip_arrays` - A list of allowlist IP addresses at the subcomponent level.
    * `security_ip_array_tag` - Tag of the IP address group.
    * `security_ip_type` - Type of the IP address.
  * `status` - Component status.
  * `topology` - The topology information of the application subcomponent.
    * `children` - A list of topology child node IDs of the current application subcomponent, or a list of subcomponent types of the child nodes.
    * `layer` - The topology layer of the topology information for the current application subcomponent.
    * `parents` - A list of topology parent node IDs or parent subcomponent types for the current application subcomponent.
* `endpoints` - List of user-defined service endpoints.
  * `description` - Description of the endpoint for the connection address.
  * `endpoint_id` - Endpoint ID of the connection address.
  * `ip` - IP address.
  * `port` - Port.
  * `port_description` - Port description.
* `region_id` - The region ID.
* `security_groups` - A list of security groups at the application level.
  * `net_type` - The network type.
  * `region_id` - The region ID.
  * `security_group_name` - The name of the security group.
* `security_ip_arrays` - The application-level whitelist.
  * `security_ip_array_tag` - The tag of the IP address group.
  * `security_ip_net_type` - The network type of the whitelist IP addresses.
  * `security_ip_type` - The type of the IP address.
* `status` - The application status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Application.
* `delete` - (Defaults to 5 mins) Used when delete the Application.
* `update` - (Defaults to 5 mins) Used when update the Application.

## Import

Polardb Application can be imported using the id, e.g.

```shell
$ terraform import alicloud_polardb_application.example <application_id>
```