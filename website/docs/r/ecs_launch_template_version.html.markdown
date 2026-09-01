---
subcategory: "ECS"
layout: "alicloud"
page_title: "Alicloud: alicloud_ecs_launch_template_version"
description: |-
  Provides a Alicloud ECS Launch Template Version resource.
---

# alicloud_ecs_launch_template_version

Provides a ECS Launch Template Version resource.

Launch template version.

For information about ECS Launch Template Version and how to use it, see [What is Launch Template Version](https://next.api.alibabacloud.com/document/Ecs/2014-05-26/CreateLaunchTemplateVersion).

-> **NOTE:** Available since v1.287.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_ecs_launch_template" "defaulttiqFgu" {
  image_owner_alias    = "system"
  instance_charge_type = "PrePaid"
  system_disk {
    category                = "cloud_essd"
    size                    = "40"
    performance_level       = "PL0"
    delete_with_instance    = true
    encrypted               = "true"
    auto_snapshot_policy_id = "sp-bp1chlvcgt3d30ec1a1u"
    disk_name               = "system-disk"
    description             = "example-system-disk"
    bursting_enabled        = true
    provisioned_iops        = "200"
  }
  network_type                  = "vpc"
  image_id                      = "aliyun_3_x64_20G_alibase_20230424.vhd"
  instance_type                 = "ecs.g8i.large"
  launch_template_name          = "launch_template_example"
  io_optimized                  = "optimized"
  zone_id                       = "random"
  vswitch_id                    = "vsw-bp13h5fqotbnqy7s7fdsm"
  internet_charge_type          = "PayByBandwidth"
  internet_max_bandwidth_out    = "1"
  vpc_id                        = "vpc-bp1xk1tto1jev2zzo1cuh"
  version_description           = "version_1"
  password_inherit              = false
  description                   = "example-description"
  private_ip_address            = "0.0.0.0"
  user_data                     = "example-userdata"
  ram_role_name                 = "lt-ram"
  spot_price_limit              = 3
  default_version_number        = "1"
  host_name                     = "example-host-name"
  spot_strategy                 = "SpotWithPriceLimit"
  key_pair_name                 = "example-key-pair-name"
  period                        = "1"
  deployment_set_id             = "example-deployment-set-id"
  instance_name                 = "example-instance-name"
  internet_max_bandwidth_in     = "10"
  security_enhancement_strategy = "Active"
  auto_release_time             = "2018-01-01T12:05:00Z"
  security_group_ids            = ["sg-bp1i7foe6v0pzol3i511"]
  template_resource_group_id    = "521"
}


resource "alicloud_ecs_launch_template_version" "default" {
  launch_template_data {
    vpc_id                              = "vpc-bp1xk1tto1jev2zzo1cuh"
    network_type                        = "vpc"
    description                         = "example_host"
    instance_name                       = "example_vm"
    system_disk_disk_name               = "example_system_disk"
    system_disk_size                    = "40"
    image_id                            = "aliyun_3_x64_20G_alibase_20230424.vhd"
    system_disk_delete_with_instance    = true
    system_disk_category                = "cloud_essd"
    system_disk_description             = "example_system_disk"
    image_owner_alias                   = "system"
    host_name                           = "example_host"
    system_disk_auto_snapshot_policy_id = "sp-bp1chlvcgt3d30ec1a1u"
    internet_max_bandwidth_out          = "1"
    instance_type                       = "ecs.g8i.large"
    instance_charge_type                = "PrePaid"
    io_optimized                        = "optimized"
    vswitch_id                          = "vsw-bp13h5fqotbnqy7s7fdsm"
    internet_charge_type                = "PayByBandwidth"
    zone_id                             = "random"
    data_disks {
      description             = "example_datadisk"
      size                    = "40"
      disk_name               = "example_datadisk_name"
      category                = "cloud_auto"
      delete_with_instance    = true
      encrypted               = "true"
      provisioned_iops        = "200"
      bursting_enabled        = true
      auto_snapshot_policy_id = "sp-bp1chlvcgt3d30ec1a1u"
      performance_level       = "PL0"
      snapshot_id             = "s-bp17441ohwka0yuh0000"
      device                  = "null"
    }
    system_disk_bursting_enabled  = false
    system_disk_encrypted         = "true"
    system_disk_performance_level = "PL1"
    deployment_set_id             = "example-deployment-set-id"
    key_pair_name                 = "example-key-pair-name"
    network_interfaces {
      network_interface_name         = "exampleNetworkInterfaceName"
      vswitch_id                     = "vsw-bp1s5fnvk4gn2tws00000"
      description                    = "exampleNetworkInterfaceDescription"
      primary_ip_address             = "0.0.0.0"
      instance_type                  = "Secondary"
      network_interface_traffic_mode = "Standard"
      security_group_ids             = []
      security_group_id              = "sg-example"
    }
    tags {
      key   = "TestKey"
      value = "TestValue"
    }
    spot_strategy                 = "SpotWithPriceLimit"
    spot_duration                 = "0"
    security_enhancement_strategy = "Active"
    user_data                     = "example-user-data"
    spot_price_limit              = 3
    private_ip_address            = "0.0.0.0"
    auto_release_time             = "2018-01-01T12:05:00Z"
    system_disk_iops              = "200"
    internet_max_bandwidth_in     = "10"
    period                        = "1"
    ram_role_name                 = "exampleRamRoleName"
    ipv6_address_count            = "1"
    system_disk_provisioned_iops  = "200"
    security_group_ids            = ["sg-bp15ed6xe1yxeycg0000"]
    enable_vm_os_config           = false
    password_inherit              = false
    deletion_protection           = false
    security_group_id             = "sg-example"
    credit_specification          = "Standard"
  }
  version_description  = "example_version"
  launch_template_id   = alicloud_ecs_launch_template.defaulttiqFgu.id
  launch_template_name = alicloud_ecs_launch_template.defaulttiqFgu.launch_template_name
}
```

## Argument Reference

The following arguments are supported:
* `launch_template_data` - (Optional, ForceNew, Set) The specific configuration of the template. See [`launch_template_data`](#launch_template_data) below.
* `launch_template_id` - (Optional, ForceNew, Computed) Template ID.
* `launch_template_name` - (Optional, ForceNew) Template name.
* `resource_group_id` - (Optional, ForceNew, Computed) The resource attribute field that represents the resource group.
* `version_description` - (Optional, ForceNew) Template version description.

### `launch_template_data`

The launch_template_data supports the following:
* `auto_release_time` - (Optional, ForceNew) Automatic release time.
* `credit_specification` - (Optional, ForceNew) Set the running mode of the burst performance instance.
* `data_disks` - (Optional, ForceNew, List) A collection of data disks. See [`data_disks`](#launch_template_data-data_disks) below.
* `deletion_protection` - (Optional, ForceNew) Instance deletion protection attribute.
* `deployment_set_id` - (Optional, ForceNew) The ID of the deployment set.
* `description` - (Optional, ForceNew) Description of the instance.
* `enable_vm_os_config` - (Optional, ForceNew) Whether to enable instance operating system configuration.
* `host_name` - (Optional, ForceNew) Instance hostname.
* `image_id` - (Optional, ForceNew) The ID of the image used by the instance.
* `image_owner_alias` - (Optional, ForceNew) Mirror source.
* `instance_charge_type` - (Optional, ForceNew) Instance billing type.
* `instance_name` - (Optional, ForceNew) The name of the instance.
* `instance_type` - (Optional, ForceNew) Instance type.
* `internet_charge_type` - (Optional, ForceNew) Public network bandwidth billing method.
* `internet_max_bandwidth_in` - (Optional, ForceNew, Int) The maximum public network inbound bandwidth.
* `internet_max_bandwidth_out` - (Optional, ForceNew, Int) The maximum outbound bandwidth of the public network.
* `io_optimized` - (Optional, ForceNew) Whether to optimize the instance for I/O.
* `ipv6_address_count` - (Optional, ForceNew, Int) The number of IPV6 addresses.
* `key_pair_name` - (Optional, ForceNew) Key pair name.
* `network_interfaces` - (Optional, ForceNew, List) A collection of attributes for a secondary ENI. See [`network_interfaces`](#launch_template_data-network_interfaces) below.
* `network_type` - (Optional, ForceNew) Network type.
* `password_inherit` - (Optional, ForceNew) Whether to inherit the username and password set in the original image.
* `period` - (Optional, ForceNew, Int) The duration of the purchase of resources.
* `private_ip_address` - (Optional, ForceNew) The private IP address of the instance.
* `ram_role_name` - (Optional, ForceNew) The instance RAM role name.
* `security_enhancement_strategy` - (Optional, ForceNew) Whether to turn on safety reinforcement.
* `security_group_id` - (Optional, ForceNew) The ID of the security group of the instance.

-> **NOTE:**  'SecurityGroupId' and 'SecurityGroupIds' do not return values at the same time.

* `security_group_ids` - (Optional, ForceNew, List) One or more security groups to which the instance is added.
* `spot_duration` - (Optional, ForceNew, Int) The retention period of a preemptible instance, in hours. Possible value: 0~6
  - The reserved duration is 2~6 and is under invitation test. Please submit the work order if you need to open it.
  - If the value is 0, it is the unprotected period mode.
* `spot_price_limit` - (Optional, ForceNew, Float) Set the maximum price per hour for the instance.
* `spot_strategy` - (Optional, ForceNew) The bidding strategy for pay-as-you-go instances.
* `system_disk_auto_snapshot_policy_id` - (Optional, ForceNew) The ID of the system disk automatic snapshot policy.
* `system_disk_bursting_enabled` - (Optional, ForceNew) Whether Burst is enabled on the system disk.
* `system_disk_category` - (Optional, ForceNew) System disk type.
* `system_disk_delete_with_instance` - (Optional, ForceNew) Whether the system disk is released with the instance.
* `system_disk_description` - (Optional, ForceNew) Description of the system disk.
* `system_disk_disk_name` - (Optional, ForceNew) Disk name.
* `system_disk_encrypted` - (Optional, ForceNew) Whether the system disk is encrypted.
* `system_disk_iops` - (Optional, ForceNew, Int) System disk I/O times per second.

-> **NOTE:**  This parameter is about to stop using. To improve code compatibility, please try to use other parameters.

* `system_disk_performance_level` - (Optional, ForceNew) When creating an ESSD cloud disk for use as a system disk, set the performance level of the cloud disk. This parameter has a return value when 'SystemDisk.Category = cloud_essd. Possible values:
  - PL0: The highest random read-write IOPS 10,000 per disk.
  - PL1: The highest random read-write IOPS 50,000 per disk.
  - PL2: The highest random read-write IOPS 100,000 per disk.
  - PL3: The highest random read-write IOPS 1 million per disk.
* `system_disk_provisioned_iops` - (Optional, ForceNew, Int) The pre-configured read and write IOPS of the system disk ESSD AutoPL cloud disk.
* `system_disk_size` - (Optional, ForceNew, Int) The size of the system disk, in GiB.
* `tags` - (Optional, ForceNew, List) The label of the instance. See [`tags`](#launch_template_data-tags) below.
* `user_data` - (Optional, ForceNew) Instance custom data, encoded in Base64.
* `vswitch_id` - (Optional, ForceNew) The ID of the virtual switch to which the instance belongs.
* `vpc_id` - (Optional, ForceNew) VPC ID.
* `zone_id` - (Optional, ForceNew) The zone ID.

### `launch_template_data-data_disks`

The launch_template_data-data_disks supports the following:
* `auto_snapshot_policy_id` - (Optional, ForceNew) The ID of the automatic snapshot policy used by the data disk.
* `bursting_enabled` - (Optional, ForceNew) Whether Burst is enabled for the data disk.
* `category` - (Optional, ForceNew) The type of cloud disk of the data disk.
* `delete_with_instance` - (Optional, ForceNew) Whether the data disk is released with the release of the instance.
* `description` - (Optional, ForceNew) Description of the data disk.
* `device` - (Optional, ForceNew) The device name of the data disk.

-> **NOTE:**  This parameter is about to stop using. In order to improve code compatibility, it is recommended that you try not to use this parameter.

* `disk_name` - (Optional, ForceNew) The name of the data disk.
* `encrypted` - (Optional, ForceNew) Whether the data disk is encrypted.
* `performance_level` - (Optional, ForceNew) When creating an ESSD cloud disk for use as a data disk, set the performance level of the cloud disk. This parameter has a return value when 'Category = cloud_essd. Possible values:
  - PL0: The highest random read-write IOPS 10,000 per disk.
  - PL1: The highest random read-write IOPS 50,000 per disk.
  - PL2: The highest random read-write IOPS 100,000 per disk.
  - PL3: The highest random read-write IOPS 1 million per disk.
* `provisioned_iops` - (Optional, ForceNew, Int) The pre-configured read and write IOPS of the data disk ESSD AutoPL cloud disk.
* `size` - (Optional, ForceNew, Int) The size of the data disk.
* `snapshot_id` - (Optional, ForceNew) The ID of the snapshot used by the data disk.

### `launch_template_data-network_interfaces`

The launch_template_data-network_interfaces supports the following:
* `description` - (Optional, ForceNew) Description information of secondary ENI.
* `instance_type` - (Optional, ForceNew) The type of the ENI.
* `network_interface_name` - (Optional, ForceNew) The name of the secondary ENI.
* `network_interface_traffic_mode` - (Optional, ForceNew) The traffic mode of the main network card.
* `primary_ip_address` - (Optional, ForceNew) The primary private IP address of the secondary Eni.
* `security_group_id` - (Optional, ForceNew) The ID of the security group to which the secondary Eni belongs. Must be a security group under the same VPC.

-> **NOTE:**  SecurityGroupId and SecurityGroupIds do not return values at the same time.

* `security_group_ids` - (Optional, ForceNew, List) One or more security groups to which the secondary ENI is attached.
* `vswitch_id` - (Optional, ForceNew) The ID of the virtual switch to which the Eni belongs.

### `launch_template_data-tags`

The launch_template_data-tags supports the following:
* `key` - (Optional, ForceNew) The tag key of the instance.
* `value` - (Optional, ForceNew) The tag value of the instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<launch_template_id>:<version_number>`.
* `create_time` - When the template was created.
* `created_by` - The creator of the template.
* `default_version` - The default version of the template.
* `modified_time` - Template modification time.
* `version_number` - Template version number.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Launch Template Version.
* `delete` - (Defaults to 5 mins) Used when delete the Launch Template Version.

## Import

ECS Launch Template Version can be imported using the id, e.g.

```shell
$ terraform import alicloud_ecs_launch_template_version.example <launch_template_id>:<version_number>
```