---
subcategory: "Elastic High Performance Computing (Ehpc)"
layout: "alicloud"
page_title: "Alicloud: alicloud_ehpc_cluster_v2"
description: |-
  Provides a Alicloud Ehpc Cluster V2 resource.
---

# alicloud_ehpc_cluster_v2

Provides a Ehpc Cluster V2 resource.

E-HPC cluster resource  .

For information about Ehpc Cluster V2 and how to use it, see [What is Cluster V2](https://next.api.alibabacloud.com/document/EHPC/2024-07-30/CreateCluster).

-> **NOTE:** Available since v1.266.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_vpc" "example" {
  is_default = false
  cidr_block = "10.0.0.0/24"
  vpc_name   = "example-cluster-vpc"
}

resource "alicloud_nas_access_group" "example" {
  access_group_type = "Vpc"
  description       = var.name
  access_group_name = "StandardMountTarget"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "example" {
  description  = "example-cluster-nas"
  storage_type = "Capacity"
  nfs_acl {
    enabled = false
  }
  zone_id          = "cn-hangzhou-k"
  encrypt_type     = "0"
  protocol_type    = "NFS"
  file_system_type = "standard"
}

resource "alicloud_vswitch" "example" {
  is_default   = false
  vpc_id       = alicloud_vpc.example.id
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "example-cluster-vsw"
}

resource "alicloud_nas_access_rule" "example" {
  priority          = "1"
  access_group_name = alicloud_nas_access_group.example.access_group_name
  file_system_type  = alicloud_nas_file_system.example.file_system_type
  source_cidr_ip    = "10.0.0.0/24"
}

resource "alicloud_ecs_key_pair" "example" {
  key_pair_name = var.name
}

resource "alicloud_nas_mount_target" "example" {
  vpc_id            = alicloud_vpc.example.id
  network_type      = "Vpc"
  access_group_name = alicloud_nas_access_group.example.access_group_name
  vswitch_id        = alicloud_vswitch.example.id
  file_system_id    = alicloud_nas_file_system.example.id
}

resource "alicloud_security_group" "example" {
  vpc_id              = alicloud_vpc.example.id
  security_group_type = "normal"
}

resource "alicloud_ehpc_cluster_v2" "default" {
  cluster_credentials {
    key_pair_name = alicloud_ecs_key_pair.example.id
  }

  cluster_mode        = "Integrated"
  cluster_vpc_id      = alicloud_vpc.example.id
  deletion_protection = "true"
  shared_storages {
    mount_directory     = "/home"
    nas_directory       = "/"
    mount_target_domain = alicloud_nas_mount_target.example.mount_target_domain
    protocol_type       = "NFS"
    file_system_id      = alicloud_nas_file_system.example.id
    mount_options       = "-t nfs -o vers=3,nolock,proto=tcp,noresvport"
  }
  shared_storages {
    nas_directory       = "/"
    mount_target_domain = alicloud_nas_mount_target.example.mount_target_domain
    protocol_type       = "NFS"
    file_system_id      = alicloud_nas_file_system.example.id
    mount_options       = "-t nfs -o vers=3,nolock,proto=tcp,noresvport"
    mount_directory     = "/opt"
  }
  shared_storages {
    mount_options       = "-t nfs -o vers=3,nolock,proto=tcp,noresvport"
    mount_directory     = "/ehpcdata"
    nas_directory       = "/"
    mount_target_domain = alicloud_nas_mount_target.example.mount_target_domain
    protocol_type       = "NFS"
    file_system_id      = alicloud_nas_file_system.example.id
  }

  cluster_vswitch_id = alicloud_vswitch.example.id
  cluster_category   = "Standard"
  security_group_id  = alicloud_security_group.example.id
  cluster_name       = var.name
  manager {
    manager_node {
      spot_strategy = "NoSpot"
      system_disk {
        category = "cloud_essd"
        size     = "40"
        level    = "PL0"
      }
      enable_ht            = "true"
      instance_charge_type = "PostPaid"
      image_id             = "centos_7_6_x64_20G_alibase_20211130.vhd"
      instance_type        = "ecs.c6.xlarge"
    }
    scheduler {
      type    = "SLURM"
      version = "22.05.8"
    }
    dns {
      type    = "nis"
      version = "1.0"
    }
    directory_service {
      type    = "nis"
      version = "1.0"
    }
  }
}
```

## Argument Reference

The following arguments are supported:
* `addons` - (Optional, ForceNew, List) Custom add-on configurations for the cluster. Only one add-on is supported.   See [`addons`](#addons) below.
* `client_version` - (Optional, Computed) The client version of the cluster. By default, the latest version is used.  
* `cluster_category` - (Optional, ForceNew) The cluster edition. Valid values:  
  - Standard: Standard Edition  
  - Serverless: Serverless Edition  
* `cluster_credentials` - (Required, ForceNew, Set) Security credentials for the cluster.   See [`cluster_credentials`](#cluster_credentials) below.
* `cluster_mode` - (Optional, ForceNew) The cluster deployment type. Valid values:  
  - Integrated: Public Cloud Cluster  
  - Hybrid: Hybrid Cloud Cluster  
  - Custom: Custom Cluster
* `cluster_name` - (Optional) Cluster name. It must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, hyphens (-), and underscores (_).
* `cluster_vswitch_id` - (Optional, ForceNew) The ID of the vSwitch used by the cluster. The vSwitch ID must belong to the VPC specified by `ClusterVpcId`.
You can call [DescribeVpcs](https://help.aliyun.com/document_detail/448581.html) to query information about existing VPCs and vSwitches.  
* `cluster_vpc_id` - (Optional, ForceNew) The VPC ID used by the cluster.  
* `deletion_protection` - (Optional) Cluster deletion protection attribute that specifies whether the cluster can be deleted via the console or the [DeleteCluster](https://help.aliyun.com/document_detail/424406.html) API.
  - true: Enables cluster deletion protection  
  - false: Disables cluster deletion protection (default)  
* `grow_interval` - (Optional, Computed, Int, Available since v1.283.0) The time interval for automatic cluster scaling.  
* `idle_interval` - (Optional, Computed, Int, Available since v1.283.0) The idle time of the cluster compute nodes.  
* `manager` - (Optional, ForceNew, Set) Configuration of the cluster management node. See [`manager`](#manager) below.
* `max_core_count` - (Optional, Computed, Int, Available since v1.283.0) The number of compute nodes that the cluster can manage.  
* `max_node_count` - (Optional, Computed, Int, Available since v1.283.0) The total number of CPU cores across all compute nodes that the cluster can manage.  
* `resource_group_id` - (Optional, ForceNew, Computed) The ID of the resource group.  
You can call [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) to query the resource group ID.  
* `security_group_id` - (Optional, ForceNew) The ID of the security group to which the newly created cluster belongs.  
You can call [DescribeSecurityGroups](https://help.aliyun.com/document_detail/25556.html) to query the available security groups in the current region.  
* `shared_storages` - (Required, ForceNew, List) Shared storage configuration for the cluster.   See [`shared_storages`](#shared_storages) below.

### `addons`

The addons supports the following:
* `name` - (Required, ForceNew) Detailed configuration information of the custom add-on.
* `resources_spec` - (Optional, ForceNew) Resource configuration of the custom add-on.  
* `services_spec` - (Optional, ForceNew) Service configuration of the custom add-on.  
* `version` - (Required, ForceNew) The version of the custom add-on.  

### `cluster_credentials`

The cluster_credentials supports the following:
* `key_pair_name` - (Optional, ForceNew, Available since v1.270.0) The root key pair for cluster nodes.  
* `password` - (Optional, ForceNew) The root password for logging into cluster nodes. It must be 8 to 20 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Valid special characters include:  
`() ~ ! @ # $ % ^ & * - = + { } [ ] : ; ‘ < > , . ? /`

-> **NOTE:**  We recommend that you call the API over HTTPS to prevent password leakage.  


### `manager`

The manager supports the following:
* `directory_service` - (Optional, ForceNew, Set) Directory service configuration information.   See [`directory_service`](#manager-directory_service) below.
* `dns` - (Optional, ForceNew, Set) DNS service configuration information.   See [`dns`](#manager-dns) below.
* `manager_node` - (Optional, ForceNew, Set) Hardware configuration information of the management node. See [`manager_node`](#manager-manager_node) below.
* `scheduler` - (Optional, ForceNew, Set) Scheduler service configuration information. See [`scheduler`](#manager-scheduler) below.

### `manager-directory_service`

The manager-directory_service supports the following:
* `type` - (Optional, ForceNew) Directory service type.  
* `version` - (Optional, ForceNew) Directory service version.  

### `manager-dns`

The manager-dns supports the following:
* `type` - (Optional, ForceNew) DNS type.  
* `version` - (Optional, ForceNew) DNS version.  

### `manager-manager_node`

The manager-manager_node supports the following:
* `auto_renew` - (Optional, ForceNew) Specifies whether to enable auto-renewal. This parameter takes effect only when InstanceChargeType is set to PrePaid. Valid values:
  - true: Enable auto-renewal  
  - false: Disable auto-renewal (default)  
* `auto_renew_period` - (Optional, ForceNew, Int) The auto-renewal period for a single renewal. Valid values:
  - When PeriodUnit=Week: 1 (default), 2, 3
  - When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, 60
* `duration` - (Optional, ForceNew, Computed, Int) The retention period for a preemptible instance, measured in hours. Valid values:
  - 1: After creation, Alibaba Cloud guarantees that the instance will run for at least 1 hour without being automatically released. After 1 hour, the system continuously compares your bid price with the current market price and checks resource availability to determine whether to retain or reclaim the instance (default)  
  - 0: After creation, Alibaba Cloud does not guarantee any minimum runtime. The system continuously compares your bid price with the current market price and checks resource availability to determine whether to retain or reclaim the instance  
* `enable_ht` - (Optional, ForceNew) EnableHT.  
* `image_id` - (Optional, ForceNew) ImageId.  
* `instance_charge_type` - (Optional, ForceNew) The billing method for the manager node instance. Valid values:
  - PostPaid: Pay-as-you-go  
  - PrePaid: Subscription  
* `instance_type` - (Optional, ForceNew) Instance type of the management node.
* `period` - (Optional, ForceNew, Int) The subscription duration of the resource, measured in the unit specified by PeriodUnit. This parameter is required and takes effect only when InstanceChargeType is set to PrePaid. If DedicatedHostId is specified, the value cannot exceed the subscription duration of the dedicated host. Valid values:
  - When PeriodUnit=Week, valid Period values are: 1, 2, 3, 4  
  - When PeriodUnit=Month, valid Period values are: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60  
* `period_unit` - (Optional, ForceNew) The time unit for subscription billing. Valid values:
  - Week
  - Month (default)
* `spot_price_limit` - (Optional, ForceNew, Float) Set the maximum hourly price for the instance. Supports up to three decimal places. This parameter takes effect only when the SpotStrategy parameter is set to SpotWithPriceLimit.
* `spot_strategy` - (Optional, ForceNew) The spot strategy for pay-as-you-go instances. This parameter takes effect only when the InstanceChargeType parameter is set to PostPaid. Valid values:
  - NoSpot: standard pay-as-you-go instance (default)
  - SpotWithPriceLimit: spot instance with a specified maximum price
  - SpotAsPriceGo: system automatically bids based on the current market price
* `system_disk` - (Optional, ForceNew, Set) The system disk configuration for the manager node.   See [`system_disk`](#manager-manager_node-system_disk) below.

### `manager-scheduler`

The manager-scheduler supports the following:
* `type` - (Optional, ForceNew) Scheduler type. Valid values:
  - SLURM
  - PBS
  - OPENGRIDSCHEDULER
  - LSF_PLUGIN
  - PBS_PLUGIN
* `version` - (Optional, ForceNew) Scheduler version.

### `manager-manager_node-system_disk`

The manager-manager_node-system_disk supports the following:
* `category` - (Optional, ForceNew) System disk configuration for the manager node. Valid values:  
  - cloud_efficiency: Ultra disk  
  - cloud_ssd: SSD cloud disk  
  - cloud_essd: ESSD cloud disk  
  - cloud: Basic cloud disk  
* `level` - (Optional, ForceNew) Performance level of the ESSD cloud disk when used as a system disk. Valid values:  
  - PL0: Maximum random read/write IOPS per disk is 10,000  
  - PL1: Maximum random read/write IOPS per disk is 50,000 (default)  
  - PL2: Maximum random read/write IOPS per disk is 100,000  
  - PL3: Maximum random read/write IOPS per disk is 1,000,000  
* `size` - (Optional, ForceNew, Int) System disk size for the manager node, in GiB. Valid ranges:  
  - Basic cloud disk: 20 to 500  
  - ESSD cloud disk:  
  - PL0: 1 to 2048  
  - PL1: 20 to 2048  
  - PL2: 461 to 2048  
  - PL3: 1261 to 2048  
  - Other disk types: 20 to 2048  

### `shared_storages`

The shared_storages supports the following:
* `file_system_id` - (Optional, ForceNew) The ID of the mounted file system.
* `mount_directory` - (Optional, ForceNew) The local mount directory for the mounted file system.  
* `mount_options` - (Optional, ForceNew) The mount options for the mounted file system.
* `mount_target_domain` - (Optional, ForceNew) The mount target domain of the mounted file system.
* `nas_directory` - (Optional, ForceNew) The remote directory of the file system to be mounted.  
* `protocol_type` - (Optional, ForceNew) The protocol type of the mounted file system. Valid values:
  - NFS
  - SMB

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `cluster_status` - Cluster status.
* `create_time` - The creation time of the cluster.
* `ehpc_version` - E-HPC product version number.
* `manager` - Configuration of the cluster management node.
  * `manager_node` - Hardware configuration information of the management node.
    * `expired_time` - Expiration time of the manager node.
    * `instance_id` - The instance ID of the management node.
* `modify_time` - The time when the cluster was last modified.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 8 mins) Used when create the Cluster V2.
* `delete` - (Defaults to 5 mins) Used when delete the Cluster V2.
* `update` - (Defaults to 5 mins) Used when update the Cluster V2.

## Import

Ehpc Cluster V2 can be imported using the id, e.g.

```shell
$ terraform import alicloud_ehpc_cluster_v2.example <cluster_id>
```