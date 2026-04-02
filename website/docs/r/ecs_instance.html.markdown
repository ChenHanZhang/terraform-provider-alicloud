---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_instance"
description: |-
  Provides a Alicloud Ecs Instance resource.
---

# alicloud_ecs_instance

Provides a Ecs Instance resource.

An ECS instance is equivalent to a virtual machine and contains basic computing components such as CPU, memory, operating system, network, and disk. You can easily customize and change the configuration of an instance. You have full control over the virtual machine.

For information about Ecs Instance and how to use it, see [What is Instance](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/RunInstances).

-> **NOTE:** Available since v1.46.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

variable "instance_type" {
  default = "ecs.n4.large"
}

variable "image_id" {
  default = "ubuntu_18_04_64_20G_alibase_20190624.vhd"
}

# Create a new ECS instance for a VPC
resource "alicloud_security_group" "group" {
  security_group_name = var.name
  description         = "foo"
  vpc_id              = alicloud_vpc.vpc.id
}

resource "alicloud_kms_key" "key" {
  description            = "Hello KMS"
  pending_window_in_days = "7"
  status                 = "Enabled"
}

data "alicloud_zones" "default" {
  available_disk_category     = "cloud_efficiency"
  available_resource_creation = "VSwitch"
  available_instance_type     = var.instance_type
}

# Create a new ECS instance for VPC
resource "alicloud_vpc" "vpc" {
  vpc_name   = var.name
  cidr_block = "172.16.0.0/16"
}

resource "alicloud_vswitch" "vswitch" {
  vpc_id       = alicloud_vpc.vpc.id
  cidr_block   = "172.16.0.0/24"
  zone_id      = data.alicloud_zones.default.zones.0.id
  vswitch_name = var.name
}

resource "alicloud_instance" "instance" {
  # cn-beijing
  availability_zone = data.alicloud_zones.default.zones.0.id
  security_groups   = alicloud_security_group.group.*.id

  # series III
  instance_type              = var.instance_type
  system_disk_category       = "cloud_efficiency"
  system_disk_name           = var.name
  system_disk_description    = "test_foo_system_disk_description"
  image_id                   = var.image_id
  instance_name              = var.name
  vswitch_id                 = alicloud_vswitch.vswitch.id
  internet_max_bandwidth_out = 10
  data_disks {
    name        = "disk2"
    size        = 20
    category    = "cloud_efficiency"
    description = "disk2"
    encrypted   = true
    kms_key_id  = alicloud_kms_key.key.id
  }
}
```

### Deleting `alicloud_ecs_instance` or removing it from your configuration

The `alicloud_ecs_instance` resource allows you to manage  `status = "Running,Stopped"`  instance, but Terraform cannot destroy it.
Deleting the subscription resource or removing it from your configuration will remove it from your state file and management, but will not destroy the Instance.
You can resume managing the subscription instance via the AlibabaCloud Console.

## Argument Reference

The following arguments are supported:
* `action_on_maintenance` - (Optional, ForceNew, Set, Available since v1.274.0) O & M action properties of the instance. See [`action_on_maintenance`](#action_on_maintenance) below.
* `allocate_system_public_ip_address` - (Optional, Available since v1.274.0) Whether to assign the default public IP address to an ECS instance.
* `amount` - (Optional, Int, Available since v1.274.0) Specify the number of ECS instances to be created

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_pay` - (Optional, Available since v1.274.0) Whether to automatically pay when creating an instance

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_reboot_time` - (Optional, Available since v1.274.0) Automatic reboot time after instance configuration change
* `auto_release_time` - (Optional, Available since v1.70.0) Automatic release time for pay-as-you-go instances.
* `auto_renew` - (Optional, Available since v1.274.0) Whether to automatically renew. It takes effect only when the parameter 'InstanceChargeType' takes value 'PrePaid. Value range:
  - true: automatic renewal.
  - false (default): no automatic renewal.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_renew_enabled` - (Optional, Available since v1.274.0) Whether automatic renewal has been set.
* `auto_renew_period` - (Optional, Int) The duration of each auto-renewal. This parameter is required when the value of AutoRenew is True.
 When PeriodUnit is Week, AutoRenewPeriod values: 1, 2, 3. 
When PeriodUnit is Month, the value of AutoRenewPeriod is 1, 2, 3, 6, and 12.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `clock_options` - (Optional, Set, Available since v1.274.0) Instance clock-related property parameters See [`clock_options`](#clock_options) below.
* `cpu` - (Optional, ForceNew, Computed, Int) VCPU auditing.
* `cpu_options` - (Optional, Computed, Set) CPU configuration details. See [`cpu_options`](#cpu_options) below.
* `credit_specification` - (Optional, Computed, Available since v1.57.1) Modify the running mode of the burst performance instance
* `data_disk` - (Optional, List, Available since v1.274.0) List of data disks. See [`data_disk`](#data_disk) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `dedicated_host_attribute` - (Optional, ForceNew, Set, Available since v1.274.0) The dedicated host attribute See [`dedicated_host_attribute`](#dedicated_host_attribute) below.
* `dedicated_instance_attribute` - (Optional, ForceNew, Set, Available since v1.274.0) Host instance parameters See [`dedicated_instance_attribute`](#dedicated_instance_attribute) below.
* `deletion_protection` - (Optional) Instance release protection
* `deployment_set_group_no` - (Optional, ForceNew, Computed, Int) ECS Instance Binding The grouped location of the instances in the deployment set when the deployment set is decentralized
* `deployment_set_id` - (Optional, ForceNew, Available since v1.176.0) Deployment set ID.
* `description` - (Optional, Computed) The description
* `dry_run` - (Optional) Whether to PreCheck only this request

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `duration` - (Optional, Int, Available since v1.274.0) The duration of automatic renewal.
* `force_delete` - (Optional, Available since v1.18.0) Whether to forcibly release The `running`('Running') instance.
  - true: forcibly release the running **('Running') instance. Forced release is equivalent to power failure, and temporary data in the instance memory and storage will be erased and cannot be recovered.
  - false: The instance is released normally. The instance must be in the `Stopped`('Stopped') state.

The default value is false.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `force_stop` - (Optional, Available since v1.274.0) Force shutdown policy when stopping an instance

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `host_name` - (Optional, Computed) Instance host name.
* `host_names` - (Optional, List, Available since v1.274.0) Specify a different host name for each instance when creating multiple instances

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `hpc_cluster_id` - (Optional, ForceNew, Available since v1.144.0) ID of the HPC cluster to which the instance belongs.
* `image_family` - (Optional, Available since v1.274.0) The name of the Image family. Set this parameter to obtain the latest custom image available in the current image family to create an instance.
  - If the'imageid' is set, this parameter cannot be set.
  - If the'imageid' is not set, you can set this parameter.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `image_id` - (Optional, ForceNew, Computed) The image ID where the instance is running.
* `image_options` - (Optional, Computed, Set, Available since v1.237.0) Image-related property information See [`image_options`](#image_options) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `instance_name` - (Optional, Computed) Instance name
* `instance_type` - (Optional, Computed) Instance specifications.
* `internet_charge_type` - (Optional, ForceNew, Computed) Network billing types
* `internet_max_bandwidth_in` - (Optional, Computed, Int) Maximum public access bandwidth.
* `internet_max_bandwidth_out` - (Optional, Computed, Int) Maximum public network output bandwidth.
* `io_optimized` - (Optional, ForceNew, Computed) Whether it is an I/O optimized instance.
* `ipv6_address` - (Optional, List, Available since v1.274.0) Specify one or more IPv6 addresses for the primary NIC

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `ipv6_address_count` - (Optional, Computed, Int, Available since v1.193.0) Specify the number of randomly generated IPv6 addresses for the primary network adapter

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `key_pair_name` - (Optional, Available since v1.274.0) The name of the key pair.
* `launch_template_id` - (Optional, Available since v1.213.1) Launch template ID

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `launch_template_name` - (Optional, Available since v1.213.1) Launch Template Name

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `launch_template_version` - (Optional, Available since v1.213.1) Launch Template Version

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `maintenance_windows` - (Optional, ForceNew, List, Available since v1.274.0) The list of instances in the O & M window. See [`maintenance_windows`](#maintenance_windows) below.
* `metadata_options` - (Optional, ForceNew, Set, Available since v1.274.0) Collection of metadata options. See [`metadata_options`](#metadata_options) below.
* `min_amount` - (Optional, Int, Available since v1.274.0) Specify the minimum purchase quantity of ECS instances

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `network_interface` - (Optional, List, Available since v1.274.0) Elastic Network Card Information See [`network_interface`](#network_interface) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `network_interface_queue_number` - (Optional, Int, Available since v1.274.0) Number of primary NIC queues

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `network_interfaces` - (Optional, ForceNew, Computed, List, Available since v1.212.0) Instance contains an elastic set of network cards See [`network_interfaces`](#network_interfaces) below.
* `notify_on_maintenance` - (Optional, Available since v1.274.0) Whether to send an event notification before the instance goes down.
* `password` - (Optional) The password of the instance. It is 8 to 30 characters in length and must contain both uppercase and lowercase letters, numbers, and special symbols. Special symbols can be:

'''
()'~! @#$% & *-_+ =|{}[]:;',.?/
'''

You need to pay attention:
  - If you pass in the Password parameter, we recommend that you use the HTTPS protocol to send requests to avoid Password leakage.
  - Windows instances cannot use forward slash (/) as the first character of the password.
  - Some instances of operating systems do not support password configuration, only key pairs are supported. For example, Others Linux and Fedora CoreOS.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `password_inherit` - (Optional, Available since v1.232.0) Whether to use the password preset by the image. When using this parameter, the Password parameter must be empty, and you must ensure that the Password has been set for the image used.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `payment_type` - (Optional, Computed, Available since v1.274.0) This property does not have a description in the spec, please add it before generating code.
* `period` - (Optional, Computed, Int) The duration of the purchase of the resource, the unit specified by 'PeriodUnit. The parameter 'InstanceChargeType' takes effect only when the value of the parameter 'PrePaid' and is required. Once 'DedicatedHostId' is specified, the value range cannot exceed the subscription duration of the dedicated host. Value range:

  - Period unit = Week, Period values: 1, 2, 3, 4.
When-PeriodUnit = Month, Period values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60.



When-PeriodUnit = Month, Period values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60.



When-PeriodUnit = Month, Period values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60.


-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `period_unit` - (Optional, ForceNew) Length of purchase of resources. Value range:

  - Week.
  - Month.


  - Month.


  - Month.


The default value is Month.
* `private_dns_name_options` - (Optional, Set, Available since v1.274.0) The private domain name configuration information of the instance. See [`private_dns_name_options`](#private_dns_name_options) below.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `private_pool_options_id` - (Optional, Computed, Available since v1.253.0) Private pool ID. When the return value is' open', the private pool ID is the private pool ID assigned when the system automatically matches.
* `private_pool_options_match_criteria` - (Optional, Computed, Available since v1.253.0) The private pool matching pattern of the instance. Possible values:
  - Open: Open mode. Instances automatically match open-type private pools.
  - Target: specifies the mode. The instance matches the specified private pool.
  - None: Do not use mode. Instances do not use private pools.
* `public_ip_address` - (Optional, Computed, List, Available since v1.274.0) IP address of instance public network.
* `ram_role_name` - (Optional, Available since v1.274.0) 实例 RAM 角色名称

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `recyclable` - (Optional, Available since v1.274.0) Whether the instance can be recycled.
* `renewal_status` - (Optional) The automatic renewal status of the instance. Possible values:
  - AutoRenewal: set to automatic renewal.
  - Normal: cancel automatic renewal.
  - NotRenewal: No renewal fee, the system will no longer send expiration reminder, only send non-renewal reminder on the third day before expiration. ECS instances that are no longer renewed can be changed to a fee ('Normal') by [ModifyInstanceAutoRenewAttribute](~~ 52843 ~~), and then self-renewal or set to automatic renewal.
* `resource_group_id` - (Optional, Computed, Available since v1.57.0) ID of the enterprise resource group to which the instance belongs.
* `security_enhancement_strategy` - (Optional, Computed) Whether security reinforcement is enabled

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_group_id` - (Optional, Available since v1.274.0) Specifies the ID of the security group to which the newly created instance belongs. Instances in the same security group can access each other.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_group_ids` - (Optional, ForceNew, List, Available since v1.274.0) Security group ID
* `security_options` - (Optional, Set, Available since v1.274.0) Security Options Collection See [`security_options`](#security_options) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `spot_duration` - (Optional, ForceNew, Computed, Int, Available since v1.188.0) The retention length of a preemptive instance
* `spot_interruption_behavior` - (Optional, ForceNew, Available since v1.274.0) Preemptive instance Interrupt mode
* `spot_price_limit` - (Optional, ForceNew, Computed, Float) Maximum hourly price for an instance
* `spot_strategy` - (Optional, ForceNew, Computed) Preemptive policies for preemptive instances
* `start_time` - (Optional, Computed) The start time of the bidding mode for the instance.
* `status` - (Optional, Computed) The status of the resource
* `stopped_mode` - (Optional, Computed, Available since v1.170.0) Whether to continue charging after the shutdown of the instance
* `storage_set_id` - (Optional, Available since v1.274.0) Save set ID

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `storage_set_partition_number` - (Optional, Int, Available since v1.274.0) Maximum number of partitions in a storage set

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `system_disk` - (Optional, Set, Available since v1.274.0) Parameters related to the system disk. Currently, the dedicated block storage cluster ID('storageclusterid') needs to use this parameter to set the parameter value. See [`system_disk`](#system_disk) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `tags` - (Optional, Map) The tags
* `unique_suffix` - (Optional, Available since v1.274.0) Whether to automatically add an ordered suffix to HostName and InstanceName when creating multiple instances

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `user_data` - (Optional) The custom data of the instance.
* `vnc_password` - (Optional, Available since v1.274.0) The new ECS instance management terminal connection password.
* `vpc_attributes` - (Optional, ForceNew, Set, Available since v1.274.0) Instance VPC information See [`vpc_attributes`](#vpc_attributes) below.
* `zone_id` - (Optional, ForceNew, Available since v1.274.0) Instance belongs to the available zone.

### `action_on_maintenance`

The action_on_maintenance supports the following:
* `value` - (Optional, ForceNew, Available since v1.274.0) Maintenance action, currently effective value. Possible values:
  - Stop: Stop state (I. E. Downtime).
  - AutoRecover: automatic recovery.
  - Autorededeploy: down migration, data disk damage.

### `clock_options`

The clock_options supports the following:
* `ptp_status` - (Optional, Available since v1.274.0) PTP state value

### `cpu_options`

The cpu_options supports the following:
* `core_count` - (Optional, Computed, Int) Number of physical CPU cores.
* `numa` - (Optional, ForceNew, Available since v1.274.0) Number of threads allocated
* `threads_per_core` - (Optional, Computed, Int) CPU threads
* `topology_type` - (Optional, Computed) The Cpu topology type of the instance. Value range:
Continuousoretohtmapping: when continuousoretohtmapping is selected, in the Cpu topology of the instance, the HT of the same Core of the instance is continuous.
DiscreteCoreToHTMapping: When DiscreteCoreToHTMapping is selected, the HT of the same Core of the instance is discrete.
Default value: None.

### `data_disk`

The data_disk supports the following:
* `auto_snapshot_policy_id` - (Optional, Available since v1.274.0) The ID of the automatic snapshot policy used by the system disk.
* `bursting_enabled` - (Optional, Available since v1.274.0) Whether to enable Burst (performance Burst)
* `category` - (Optional, Available since v1.274.0) The type of cloud disk of data disk N. Value range:
  - cloud_efficiency: efficient cloud disk.
  - cloud_ssd:SSD cloud disk.
  - cloud_essd:ESSD cloud disk.
  - cloud: ordinary cloud disk.

The default value of I/O optimized instances is cloud_efficiency, and the default value of non-I/O optimized instances is cloud.
* `delete_with_instance` - (Optional, Available since v1.274.0) Whether the data disk is released with the instance.
The default value is true.
* `description` - (Optional, Available since v1.274.0) Data disk description. 2 to 256 English or Chinese characters in length and cannot start with' http:// 'or' https.
* `device` - (Optional, Available since v1.274.0) Mount point.

-> **NOTE:**  This parameter is about to be stopped. In order to improve code compatibility, it is recommended that you try not to use this parameter.

* `disk_name` - (Optional, Available since v1.274.0) The name of the data disk. The length is 2~128 English or Chinese characters. It must start with a large or small letter or Chinese, and cannot start with' http:// 'or' https. It can contain numbers, colons (:), underscores (_), or dashes (-).
* `encrypt_algorithm` - (Optional, Available since v1.274.0) Encryption algorithm.
* `encrypted` - (Optional, Available since v1.274.0) Whether data disk N is encrypted.
The default value is false.
* `kms_key_id` - (Optional, Available since v1.274.0) The ID of the KMS key used by the cloud disk.
* `performance_level` - (Optional, Available since v1.274.0) When creating an ESSD cloud disk for use as a data disk, set the performance level of the cloud disk. The value of N must be consistent with N in 'DataDisk.N.Category = cloud_essd. Value range:
  - PL0: The highest random read-write IOPS 10,000 per disk.
  - PL1 (default): The highest random read-write IOPS 50,000 per disk.
  - PL2: The highest random read-write IOPS 100,000 per disk.
  - PL3: The highest random read-write IOPS 1 million per disk.

For how to select an ESSD performance level, see [ESSD cloud disk](~~ 122389 ~~).
* `provisioned_iops` - (Optional, Int, Available since v1.274.0) ESSD AutoPL cloud disk pre-configured read/write IOPS
* `size` - (Optional, Int, Available since v1.274.0) The capacity of the nth data disk. The value of n ranges from 1 to 16. The memory unit is GiB. Value range:
  - cloud_efficiency:20~32768.
  - cloud_ssd:20~32768.
  - cloud_essd: The specific value range is related to the value of' datadisk. N. Performanclevel.
  - PL0:40~32768.
  - PL1:20~32768.
  - PL2:461~32768.
  - PL3:1261~32768.
  - cloud:5~2000.

The value of this parameter must be greater than or equal to the size of the snapshot specified by the parameter' snapshotid.
* `snapshot_id` - (Optional, Available since v1.274.0) Create a snapshot used by data disk n. The value range of N is 1 to 16.
  - After the parameter 'DataDisk.N.SnapshotId' is specified, the parameter 'DataDisk.N.Size' is ignored, and the actual size of the created cloud disk is the size of the specified snapshot.
  - Snapshots created earlier than July 15, 2013 (inclusive) cannot be used. The request will report an error and be rejected.
* `storage_cluster_id` - (Optional, Available since v1.274.0) The ID of the dedicated block storage cluster. If you need to use the cloud disk resources in the dedicated block storage cluster as data disks when creating an ECS instance, set this parameter.

### `dedicated_host_attribute`

The dedicated_host_attribute supports the following:
* `dedicated_host_cluster_id` - (Optional, ForceNew, Available since v1.274.0) Proprietary host cluster ID.
* `dedicated_host_id` - (Optional, ForceNew, Available since v1.274.0) Private host ID

### `dedicated_instance_attribute`

The dedicated_instance_attribute supports the following:
* `affinity` - (Optional, ForceNew, Available since v1.274.0) Whether a proprietary host instance is connected to a proprietary host mechanism
* `tenancy` - (Optional, ForceNew, Available since v1.274.0) Whether the instance's host type is a proprietary host

### `image_options`

The image_options supports the following:
* `login_as_non_root` - (Optional) Whether the instance using the image supports using ecsuser user logon

### `maintenance_windows`

The maintenance_windows supports the following:
* `end_time` - (Optional, Available since v1.274.0) Maintenance time window end time.
* `start_time` - (Optional, Available since v1.274.0) Maintenance time window start time.

### `metadata_options`

The metadata_options supports the following:
* `http_endpoint` - (Optional, Available since v1.274.0) Whether to enable the access channel of instance metadata. Possible values:
  - enabled: enabled
  - disabled: disabled
* `http_put_response_hop_limit` - (Optional, Int, Available since v1.274.0) This parameter is not available.
* `http_tokens` - (Optional, Available since v1.274.0) Whether to force reinforcement mode (IMDSv2) when accessing instance metadata. Possible values:
  - optional: not mandatory
  - required: mandatory use

### `network_interface`

The network_interface supports the following:
* `bandwidth_weighting` - (Optional, Available since v1.274.0) Bandwidth weight value of the example
* `delete_on_release` - (Optional, Available since v1.274.0) Whether to retain the NIC when releasing the instance
* `description` - (Optional, Available since v1.274.0) Description of the flexible network card
* `enable_instance_id_dns_aaaa_record` - (Optional, Available since v1.274.0) Whether the domain name of the instance ID type is resolved to IPv6
* `enable_jumbo_frame` - (Optional, Available since v1.274.0) 实例是否开启 Jumbo frame 特性
* `instance_type` - (Optional, Available since v1.274.0) Elastic Network Card Type
* `ipv6_address` - (Optional, List, Available since v1.274.0) Specify one or more IPv6 addresses for the primary NIC
* `ipv6_address_count` - (Optional, Int, Available since v1.274.0) 
Specify the number of randomly generated IPv6 addresses for the primary network adapter
* `network_card_index` - (Optional, Available since v1.274.0) Index of the physical NIC specified by the NIC
* `network_interface_id` - (Optional, Available since v1.274.0) The ID of the Eni attached to the instance.
* `network_interface_name` - (Optional, Available since v1.274.0) ENI Name
* `network_interface_traffic_mode` - (Optional, Available since v1.274.0) Communication mode of network card
* `primary_ip_address` - (Optional, Available since v1.274.0) Add an Eni and set the primary IP address
* `queue_number` - (Optional, Int, Available since v1.274.0) Number of ENI queues
* `queue_pair_number` - (Optional, Int, Available since v1.274.0) Number of RDMA network card queues
* `rx_queue_size` - (Optional, Int, Available since v1.274.0) ENI inbound queue depth
* `security_group_id` - (Optional, Available since v1.274.0) ID of the security group to which the Eni belongs
* `security_group_ids` - (Optional, List, Available since v1.274.0) 弹性网卡所属的一个或多个安全组 ID
* `source_dest_check` - (Optional, Available since v1.274.0) Whether the source or destination check function is on
* `tx_queue_size` - (Optional, Int, Available since v1.274.0) ENI outbound queue depth
* `vswitch_id` - (Optional, Available since v1.274.0) ID of the virtual switch to which the Eni belongs

### `network_interfaces`

The network_interfaces supports the following:
* `private_ip_sets` - (Optional, ForceNew, List, Available since v1.274.0) A collection of privateIPsets See [`private_ip_sets`](#network_interfaces-private_ip_sets) below.

### `network_interfaces-private_ip_sets`

The network_interfaces-private_ip_sets supports the following:
* `private_ip_address` - (Optional, Available since v1.274.0) The private network IP address of the instance

### `private_dns_name_options`

The private_dns_name_options supports the following:
* `enable_instance_id_dns_record` - (Optional, Available since v1.274.0) Enable/disable the resolution of domain names of the instance ID type to IPv4.
* `enable_ip_dns_ptr_record` - (Optional, Available since v1.274.0) Enable/disable domain name resolution from IPv4 to IP.
* `enable_ip_dns_record` - (Optional, Available since v1.274.0) Enable/disable the resolution of IP domain names to IPv4.
* `hostname_type` - (Optional, Available since v1.274.0) Host name type

### `security_options`

The security_options supports the following:
* `confidential_computing_mode` - (Optional, Available since v1.274.0) Confidential computing mode
* `trusted_system_mode` - (Optional, Available since v1.274.0) Trusted System Mode

### `system_disk`

The system_disk supports the following:
* `auto_snapshot_policy_id` - (Optional, Available since v1.274.0) The ID of the automatic snapshot policy used by the system disk.
* `bursting_enabled` - (Optional, Available since v1.274.0) 是否开启性能突发
* `category` - (Optional, Available since v1.274.0) The cloud disk type of the system disk. Value range:
cloud_essd:ESSD cloud disk. You can set the performance level of the cloud disk by using the parameter SystemDisk. Performanclevel.
cloud_efficiency: efficient cloud disk.
cloud_ssd:SSD cloud disk.
cloud: ordinary cloud disk.
The default value of the instance type that has been discontinued and is not I/O optimized is cloud. Otherwise, the default value is cloud_efficiency.
* `description` - (Optional, Available since v1.274.0) Description of the system disk. 2 to 256 English or Chinese characters in length and cannot start with http:// or https.
* `disk_name` - (Optional, Available since v1.274.0) The name of the system disk. The length is 2~128 English or Chinese characters. It must start with a letter or Chinese, and cannot start with http:// or https. It can contain numbers, colons (:), underscores (_), or dashes (-).
* `encrypted` - (Optional, Available since v1.274.0) Whether data disk N is encrypted.
* `kms_key_id` - (Optional, Available since v1.274.0) KMS key ID of the system disk
* `performance_level` - (Optional, Available since v1.274.0) When creating an ESSD cloud disk for use as a system disk, set the performance level of the cloud disk. Value range:
PL0: The highest random read and write IOPS 10,000 per disk.
PL1 (default): The highest random read-write IOPS 50,000 per disk.
PL2: The highest random read and write IOPS 100,000 per disk.
PL3: The highest random read and write IOPS 1 million per disk.
* `provisioned_iops` - (Optional, Int, Available since v1.274.0) ESSD AutoPL cloud disk pre-configured read/write IOPS
* `size` - (Optional, Available since v1.274.0) The size of the system disk. Unit: GiB. The value range is 20~500.
The value of this parameter must be greater than or equal to max{20, ImageSize}.
* `storage_cluster_id` - (Optional, Available since v1.274.0) Dedicated block storage cluster ID

### `vpc_attributes`

The vpc_attributes supports the following:
* `vswitch_id` - (Optional, ForceNew, Available since v1.274.0) Switch ID
* `vpc_id` - (Optional, Available since v1.274.0) Virtual private network IDs

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `action_on_maintenance` - O & M action properties of the instance.
  * `default_value` - Maintenance action, default value.
  * `supported_values` - An array format consisting of maintenance actions, which returns the supported O & M action values.
* `create_time` - The create time.
* `dedicated_host_attribute` - The dedicated host attribute.
  * `dedicated_host_name` - The dedicated host name.
* `expired_time` - Expiration time.
* `memory` - Memory size.
* `network_interfaces` - Instance contains an elastic set of network cards.
  * `ipv_sets` - IPv6 address information.
    * `ipv_address` - The IPv6 address specified for the elastic network card.
  * `mac_address` - The MAC address of the elastic network card.
  * `network_interface_id` - Elastic NIC ID.
  * `primary_ip_address` - Elastic network card master private IP address.
  * `private_ip_sets` - A collection of privateIPsets.
    * `primary` - Whether it is the IP address of the primary and private network.
  * `type` - Elastic network card type.
* `os_name` - The operating system name of the instance.
* `os_type` - The operating system type of the instance.
* `region_id` - The region.
* `vpc_attributes` - Instance VPC information.
  * `nat_ip_address` - IP of cloud products is used for network interworking among cloud products.
  * `private_ip_address` - Private IP address.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Instance.
* `delete` - (Defaults to 5 mins) Used when delete the Instance.
* `update` - (Defaults to 5 mins) Used when update the Instance.

## Import

Ecs Instance can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_instance.example <instance_id>
```