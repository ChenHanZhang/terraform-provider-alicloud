---
subcategory: "Container Service for Kubernetes (ACK)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cs_kubernetes_node_pool"
description: |-
  Provides a Alicloud Container Service for Kubernetes (ACK) Node Pool resource.
---

# alicloud_cs_kubernetes_node_pool

Provides a Container Service for Kubernetes (ACK) Node Pool resource.

This resource will help you to manage node pool in Kubernetes Cluster, see [What is kubernetes node pool](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/api-create-node-pools). 

-> **NOTE:** Available since v1.97.0.

-> **NOTE:** From version 1.109.1, support managed node pools, but only for the professional managed clusters.

-> **NOTE:** From version 1.109.1, support remove node pool nodes.

-> **NOTE:** From version 1.111.0, support auto scaling node pool. For more information on how to use auto scaling node pools, see [Use Terraform to create an elastic node pool](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/api-create-node-pools). With auto-scaling is enabled, the nodes in the node pool will be labeled with `k8s.aliyun.com=true` to prevent system pods such as coredns, metrics-servers from being scheduled to elastic nodes, and to prevent node shrinkage from causing business abnormalities.

-> **NOTE:** ACK adds a new RamRole (AliyunCSManagedAutoScalerRole) for the permission control of the node pool with auto-scaling enabled. If you are using a node pool with auto scaling, please click [AliyunCSManagedAutoScalerRole](https://ram.console.aliyun.com/role/authorization?request=%7B%22Services%22%3A%5B%7B%22Service%22%3A%22CS%22%2C%22Roles%22%3A%5B%7B%22RoleName%22%3A%22AliyunCSManagedAutoScalerRole%22%2C%22TemplateId%22%3A%22AliyunCSManagedAutoScalerRole%22%7D%5D%7D%5D%2C%22ReturnUrl%22%3A%22https%3A%2F%2Fcs.console.aliyun.com%2F%22%7D) to complete the authorization. 

-> **NOTE:** ACK adds a new RamRole（AliyunCSManagedNlcRole） for the permission control of the management node pool. If you use the management node pool, please click [AliyunCSManagedNlcRole](https://ram.console.aliyun.com/role/authorization?spm=5176.2020520152.0.0.387f16ddEOZxMv&request=%7B%22Services%22%3A%5B%7B%22Service%22%3A%22CS%22%2C%22Roles%22%3A%5B%7B%22RoleName%22%3A%22AliyunCSManagedNlcRole%22%2C%22TemplateId%22%3A%22AliyunCSManagedNlcRole%22%7D%5D%7D%5D%2C%22ReturnUrl%22%3A%22https%3A%2F%2Fcs.console.aliyun.com%2F%22%7D) to complete the authorization.

-> **NOTE:** From version 1.123.1, supports the creation of a node pool of spot instance.

-> **NOTE:** It is recommended to create a cluster with zero worker nodes, and then use a node pool to manage the cluster nodes. 

-> **NOTE:** From version 1.127.0, support for adding existing nodes to the node pool. In order to distinguish automatically created nodes, it is recommended that existing nodes be placed separately in a node pool for management. 

-> **NOTE:** From version 1.149.0, support for specifying deploymentSet for node pools. 

-> **NOTE:** From version 1.158.0, Support for specifying the desired size of nodes for the node pool, for more information, visit [Modify the expected number of nodes in a node pool](https://www.alibabacloud.com/help/en/doc-detail/160490.html#title-mpp-3jj-oo3)

-> **NOTE:** From version 1.166.0, Support configuring system disk encryption.

-> **NOTE:** From version 1.177.0+, Support `kms_encryption_context`, `rds_instances`, `system_disk_snapshot_policy_id` and `cpu_policy`, add spot strategy `SpotAsPriceGo` and `NoSpot`.

-> **NOTE:** From version 1.180.0+, Support worker nodes customized kubelet parameters by field `kubelet_configuration` and `rollout_policy`.

-> **NOTE:** From version 1.185.0+, Field `rollout_policy` will be deprecated and please use field `rolling_policy` instead.

For information about Container Service for Kubernetes (ACK) Node Pool and how to use it, see [What is Node Pool](https://www.alibabacloud.com/help/en/ack/ack-managed-and-ack-dedicated/developer-reference/api-create-node-pools).

-> **NOTE:** Available since v1.97.0.

## Example Usage

Basic Usage

```terraform
resource "random_integer" "default" {
  max = 99999
  min = 10000
}

variable "name" {
  default = "terraform-example"
}

data "alicloud_enhanced_nat_available_zones" "enhanced" {
}

data "alicloud_instance_types" "cloud_efficiency" {
  availability_zone    = data.alicloud_enhanced_nat_available_zones.enhanced.zones.0.zone_id
  cpu_core_count       = 4
  memory_size          = 8
  kubernetes_node_role = "Worker"
  system_disk_category = "cloud_efficiency"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}
resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_enhanced_nat_available_zones.enhanced.zones.0.zone_id
}

resource "alicloud_cs_managed_kubernetes" "default" {
  name_prefix          = "terraform-example-${random_integer.default.result}"
  cluster_spec         = "ack.pro.small"
  worker_vswitch_ids   = [alicloud_vswitch.default.id]
  new_nat_gateway      = true
  pod_cidr             = cidrsubnet("10.0.0.0/8", 8, 36)
  service_cidr         = cidrsubnet("172.16.0.0/16", 4, 7)
  slb_internet_enabled = true
  enable_rrsa          = true
}

resource "alicloud_key_pair" "default" {
  key_pair_name = "terraform-example-${random_integer.default.result}"
}

resource "alicloud_cs_kubernetes_node_pool" "default" {
  node_pool_name       = var.name
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40
  key_name             = alicloud_key_pair.default.key_pair_name
  // define with multi-labels by defining with labels blocks
  labels {
    key   = "test1"
    value = "nodepool"
  }
  labels {
    key   = "test2"
    value = "nodepool"
  }
  // define with multi-taints by defining with taints blocks
  taints {
    key    = "tf"
    effect = "NoSchedule"
    value  = "example"
  }
  taints {
    key    = "tf2"
    effect = "NoSchedule"
    value  = "example2"
  }
}

#The parameter `node_count` is deprecated from version 1.158.0. Please use the new parameter `desired_size` instead, you can update it as follows.
resource "alicloud_cs_kubernetes_node_pool" "desired_size" {
  node_pool_name       = "desired_size"
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40
  key_name             = alicloud_key_pair.default.key_pair_name
  desired_size         = 0
}

# Create a managed node pool. If you need to enable maintenance window, you need to set the maintenance window in `alicloud_cs_managed_kubernetes`.
resource "alicloud_cs_kubernetes_node_pool" "maintenance" {
  node_pool_name       = "maintenance"
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40

  # only key_name is supported in the management node pool
  key_name = alicloud_key_pair.default.key_pair_name

  # you need to specify the number of nodes in the node pool, which can be zero
  desired_size = 1

  # management node pool configuration.
  management {
    enable      = true
    auto_repair = true
    auto_repair_policy {
      restart_node = true
    }
    auto_upgrade = true
    auto_upgrade_policy {
      auto_upgrade_kubelet = true
    }
    auto_vul_fix = true
    auto_vul_fix_policy {
      vul_level    = "asap"
      restart_node = true
    }
    max_unavailable = 1
  }

  # Enable with automatic scaling node pool configuration.
  # With auto-scaling is enabled, the nodes in the node pool will be labeled with `k8s.aliyun.com=true` to prevent system pods such as coredns, metrics-servers from being scheduled to elastic nodes, and to prevent node shrinkage from causing business abnormalities.
  #  scaling_config {
  #    min_size = 1
  #    max_size = 10
  #    type     = "cpu"
  #  }
}

#Create a node pool with spot instance.
resource "alicloud_cs_kubernetes_node_pool" "spot_instance" {
  node_pool_name       = "spot_instance"
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id, data.alicloud_instance_types.cloud_efficiency.instance_types.1.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40
  key_name             = alicloud_key_pair.default.key_pair_name

  # you need to specify the number of nodes in the node pool, which can be 0
  desired_size = 1

  # spot config
  spot_strategy = "SpotWithPriceLimit"
  spot_price_limit {
    instance_type = data.alicloud_instance_types.cloud_efficiency.instance_types.0.id
    # Different instance types have different price caps
    price_limit = "0.70"
  }
  // define with multi-spot_price_limit by defining with spot_price_limit blocks
  spot_price_limit {
    instance_type = data.alicloud_instance_types.cloud_efficiency.instance_types.1.id
    price_limit   = "0.72"
  }
}


#Use Spot instances to create a node pool with auto-scaling enabled
resource "alicloud_cs_kubernetes_node_pool" "spot_auto_scaling" {
  node_pool_name       = "spot_auto_scaling"
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40
  key_name             = alicloud_key_pair.default.key_pair_name

  # automatic scaling node pool configuration.
  scaling_config {
    min_size = 1
    max_size = 10
    type     = "spot"
  }
  # spot price config
  spot_strategy = "SpotWithPriceLimit"
  spot_price_limit {
    instance_type = data.alicloud_instance_types.cloud_efficiency.instance_types.0.id
    price_limit   = "0.70"
  }
}

#Create a `PrePaid` node pool.
resource "alicloud_cs_kubernetes_node_pool" "prepaid_node" {
  node_pool_name       = "prepaid_node"
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40
  key_name             = alicloud_key_pair.default.key_pair_name
  # use PrePaid
  instance_charge_type = "PrePaid"
  period               = 1
  period_unit          = "Month"
  auto_renew           = true
  auto_renew_period    = 1

  # open cloud monitor
  install_cloud_monitor = true
}

##Create a node pool with customized kubelet parameters
resource "alicloud_cs_kubernetes_node_pool" "customized_kubelet" {
  node_pool_name       = "customized_kubelet"
  cluster_id           = alicloud_cs_managed_kubernetes.default.id
  vswitch_ids          = [alicloud_vswitch.default.id]
  instance_types       = [data.alicloud_instance_types.cloud_efficiency.instance_types.0.id]
  system_disk_category = "cloud_efficiency"
  system_disk_size     = 40
  instance_charge_type = "PostPaid"
  desired_size         = 0

  # kubelet configuration parameters
  kubelet_configuration {
    registry_pull_qps     = 10
    registry_burst        = 5
    event_record_qps      = 10
    event_burst           = 5
    serialize_image_pulls = true
    eviction_hard = {
      "memory.available"  = "1024Mi"
      "nodefs.available"  = "10%"
      "nodefs.inodesFree" = "5%"
      "imagefs.available" = "10%"
    }
    system_reserved = {
      "cpu"               = "1"
      "memory"            = "1Gi"
      "ephemeral-storage" = "10Gi"
    }
    kube_reserved = {
      "cpu"    = "500m"
      "memory" = "1Gi"
    }
    container_log_max_size  = "200Mi"
    container_log_max_files = 3
    max_pods                = 100
    read_only_port          = 0
    allowed_unsafe_sysctls  = ["net.ipv4.route.min_pmtu"]
  }

  # rolling policy: works when updating
  rolling_policy {
    max_parallelism = 1
  }
}
```

## Argument Reference

The following arguments are supported:
* `auto_mode` - (Optional, ForceNew, Computed, Set, Available since v1.266.0) Specifies whether to enable intelligent managed mode. When enabled, the system automatically manages the node pool by using optimized default configurations. Note: When auto_mode is enabled, many parameters are automatically set to their default values and cannot be modified. For more information, see auto_mode.enable below. See [`auto_mode`](#auto_mode) below.
* `auto_renew` - (Optional) Specifies whether auto-renewal is enabled for nodes in the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid`. Valid values:  
  - `true`: Enables auto-renewal.  
  - `false`: Disables auto-renewal.  

Default value: `false`.  
* `auto_renew_period` - (Optional, Int) The auto-renewal period for each renewal. Valid values:
  - If PeriodUnit=Week: 1, 2, or 3.
  - If PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, or 60.

Default value: 1.
* `cis_enabled` - (Optional, ForceNew, Deprecated since v1.223.1) [This field is deprecated.]
Use the `security_hardening_os` parameter instead.
* `cluster_id` - (Required, ForceNew) Cluster ID.
* `compensate_with_on_demand` - (Optional) When `multi_az_policy` is set to `COST_OPTIMIZED`, specifies whether the system is allowed to automatically attempt to create pay-as-you-go instances to meet the required number of ECS instances if sufficient preemptible instances cannot be created due to price or inventory constraints. Valid values:  
  - `true`: Allows automatic creation of pay-as-you-go instances to meet the required number of ECS instances.  
  - `false`: Does not allow automatic creation of pay-as-you-go instances to meet the required number of ECS instances.  
* `cpu_policy` - (Optional, Computed) The CPU management policy for nodes. When the cluster version is 1.12.6 or later, the following two policies are supported:
  - `static`: Enables enhanced CPU affinity and exclusivity for Pods on the node that have specific resource characteristics.
  - `none`: Uses the default CPU affinity scheme.

Default value: `none`.
* `data_disks` - (Optional, List) Data disk configuration for nodes in the node pool. See [`data_disks`](#data_disks) below.
* `deployment_set_id` - (Optional, ForceNew) The deployment set ID. You can use a deployment set to distribute Elastic Compute Service (ECS) instances launched by the node pool across different physical servers, ensuring high availability and underlying disaster recovery capabilities for your workloads. When creating ECS instances within a deployment set, the instances are started across specified zones according to the predefined deployment policy.


-> **NOTE:** After selecting a deployment set, the maximum number of nodes in the node pool becomes limited. By default, a deployment set supports a maximum of 20 × the number of zones (the number of zones is determined by the vSwitches). Choose carefully and ensure sufficient quota within the deployment set to avoid node provisioning failures.>

* `desired_size` - (Optional) The desired number of nodes in the node pool.
This is the total number of nodes the node pool should maintain. We recommend configuring at least two nodes to ensure proper operation of cluster components. You can scale the node pool in or out by adjusting this value.
If you do not want to create nodes immediately, you can set this value to 0 and manually increase it later.
* `eflo_node_group` - (Optional, Set, Available since v1.252.0) Lingjun node pool configuration. See [`eflo_node_group`](#eflo_node_group) below.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `force_delete` - (Optional) Specifies whether to forcibly delete the node pool.
  - true: Forces deletion. If the node pool contains existing nodes, they will be forcibly deleted.
  - false: Does not force deletion. If the node pool contains existing nodes, the node pool cannot be deleted, and the API call returns an error.

-> **NOTE:** This parameter only takes effect when deletion is triggered.

* `image_id` - (Optional, Computed) Custom image ID. You can use `DescribeKubernetesVersionMetadata` to query images supported by the system. By default, the latest system image is used.
* `image_type` - (Optional, Computed, Available since v1.236.0) The operating system distribution type. We recommend that you use this field to specify the node operating system. Valid values:
  - `AliyunLinux`: Alibaba Cloud Linux 2 image.
  - `AliyunLinuxSecurity`: Alibaba Cloud Linux 2 UEFI image.
  - `AliyunLinux3`: Alibaba Cloud Linux 3 image.
  - `AliyunLinux3Arm64`: Alibaba Cloud Linux 3 ARM64 image.
  - `AliyunLinux3Security`: Alibaba Cloud Linux 3 UEFI image.
  - `CentOS`: CentOS image.
  - `Windows`: Windows image.
  - `WindowsCore`: Windows Core image.
  - `ContainerOS`: Container-optimized OS image.
  - `AliyunLinux3ContainerOptimized`: Alibaba Cloud Linux 3 container-optimized image.
* `install_cloud_monitor` - (Optional) Specifies whether to install CloudMonitor on ECS nodes. After installation, you can view monitoring information for the created ECS instances in the CloudMonitor console. We recommend enabling this feature. Valid values:
  - `true`: Install CloudMonitor on ECS nodes.
  - `false`: Do not install CloudMonitor on ECS nodes.

Default value: `false`.
* `instance_charge_type` - (Optional, Computed) Billing method for nodes in the node pool. Valid values:
  - `PrePaid`: Subscription instance.
  - `PostPaid`: Pay-as-you-go instance.

Default value: `PostPaid`.
* `instance_metadata_options` - (Optional, ForceNew, Computed, Set, Available since v1.266.0) ECS instance metadata access configuration. See [`instance_metadata_options`](#instance_metadata_options) below.
* `instance_patterns` - (Optional, List, Available since v1.266.0) Instance attribute configuration. See [`instance_patterns`](#instance_patterns) below.
* `instance_types` - (Optional, List) A list of ECS instance types for nodes in the node pool. When scaling out, the system selects a suitable instance type from this list to create new nodes.

Supported number of instance types: [1, 10].

-> **NOTE:**  To ensure high availability, we recommend selecting multiple instance types.

* `internet_charge_type` - (Optional) The billing method for the node's public IP address.
  - PayByBandwidth: billed based on fixed bandwidth.
  - PayByTraffic: billed based on actual traffic usage.
* `internet_max_bandwidth_out` - (Optional, Int) The maximum outbound bandwidth of the node's public IP address, measured in Mbps (megabits per second). Valid values: 1 to 100.
* `key_name` - (Optional) The name of the key pair for password-free logon. Specify either this parameter or `login_password`.

-> **NOTE:**  If the node pool uses the ContainerOS operating system, only `key_pair` is supported.

* `kubelet_configuration` - (Optional, Set) Supported kubelet configurations. See [`kubelet_configuration`](#kubelet_configuration) below.
* `labels` - (Optional, List) Node labels. See [`labels`](#labels) below.
* `login_as_non_root` - (Optional, ForceNew) Specifies whether the created ECS instances log in as a non-root user.
  - true: Log in as the non-root user (ecs-user).
  - false: Log in as the root user.
* `management` - (Optional, Computed, Set) Managed node pool configuration. See [`management`](#management) below.
* `multi_az_policy` - (Optional, Computed) The ECS instance scaling policy for multi-zone scaling groups. Valid values:
  - `PRIORITY`: Scales instances based on the virtual switches (VSwitchIds.N) you define. If ECS instances cannot be created in the zone of a higher-priority virtual switch, the system automatically uses the next-priority virtual switch to create ECS instances.
  - `COST_OPTIMIZED`: Attempts to create instances in ascending order of vCPU unit price. When multiple instance types are specified in the scaling configuration with the preemptible billing method, preemptible instances are prioritized. You can further use the `CompensateWithOnDemand` parameter to specify whether to automatically attempt creating pay-as-you-go instances if preemptible instances cannot be created due to insufficient inventory or other reasons.

  > `COST_OPTIMIZED` takes effect only when multiple instance types are configured in the scaling configuration or when preemptible instances are selected.
  - `BALANCE`: Distributes ECS instances evenly across the multiple zones specified for the scaling group. If imbalance occurs between zones due to insufficient inventory or other reasons, you can call the [RebalanceInstances](~~71516~~) API to rebalance resources.

Default value: `PRIORITY`.
* `node_name_mode` - (Optional, ForceNew, Computed) Custom node naming. After enabling custom node naming, the node name, ECS instance name, and ECS instance hostname will all be changed accordingly.  

-> **NOTE:**  For Windows instances with custom node naming enabled, the hostname is fixed as the IP address, where dots (.) in the IP address are replaced with hyphens (-), and no prefix or suffix is included.

The node name consists of three parts: a prefix, the node's IP address, and a suffix:
  - The total length must be 2–64 characters. The first and last characters must be lowercase letters or digits.
  - The prefix and suffix may contain uppercase and lowercase letters, digits, hyphens (-), and periods (.). They must start with an uppercase or lowercase letter and cannot start or end with a hyphen (-) or period (.). Consecutive hyphens (-) or periods (.) are not allowed.
  - The prefix is required (due to ECS restrictions), while the suffix is optional.
  - The node IP address is the full private IP address of the node.

Example: If the node IP address is 192.XX.YY.55, the prefix is specified as "aliyun.com", and the suffix is "test":
  - For a Linux node, the node name, ECS instance name, and ECS instance hostname will all be "aliyun.com192.XX.YY.55test".
  - For a Windows node, the ECS instance hostname will be "192-XX-YY-55", while the node name and ECS instance name will both be "aliyun.com192.XX.YY.55test".
* `node_pool_name` - (Optional) The name of the node pool.
Naming rules: The name can contain digits, Chinese characters, English letters, or hyphens (-), must be 1 to 63 characters in length, and cannot start with a hyphen (-).
* `on_demand_base_capacity` - (Optional) The minimum number of pay-as-you-go instances required in the scaling group. Valid values: [0, 1000]. If the current number of pay-as-you-go instances is less than this value, pay-as-you-go instances are preferentially created.
* `on_demand_percentage_above_base_capacity` - (Optional) The percentage of pay-as-you-go instances among the instances that exceed the minimum number of pay-as-you-go instances (`on_demand_base_capacity`) required by the scaling group. Valid values: [0, 100].  
* `password` - (Optional) The SSH login password. You must specify either this parameter or `key_pair`. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
For security reasons, the returned password is encrypted.
* `period` - (Optional, Int) The subscription duration (in months or weeks) for nodes in the node pool. This parameter is required and takes effect only when `instance_charge_type` is set to `PrePaid`.
  - When `period_unit=Week`, valid values for `period` are: {1, 2, 3, 4}.
  - When `period_unit=Month`, valid values for `period` are: {1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, 60}.
* `period_unit` - (Optional) The billing cycle for nodes in the node pool. This parameter takes effect only when `instance_charge_type` is set to `PrePaid` and is required in this case.
  - `Month`: Billed on a monthly basis.
  - `Week`: Billed on a weekly basis.

Default value: `Month`.
* `platform` - (Optional, Computed, Deprecated since v1.145.0) 【This field is deprecated.】Use the `image_type` parameter instead.

Operating system distribution. Valid values:
  - `CentOS`
  - `AliyunLinux`
  - `Windows`
  - `WindowsCore`

Default value: `AliyunLinux`.
* `pre_user_data` - (Optional, Available since v1.232.0) Pre-custom data for the node pool, which refers to scripts executed before node initialization. For more information, see [Generate Instance Custom Data](~~49121~~).
* `private_pool_options` - (Optional, Set) Private node pool configuration. See [`private_pool_options`](#private_pool_options) below.
* `ram_role_name` - (Optional, ForceNew, Computed, Available since v1.242.0) Worker RAM role name.

* If left empty, the default Worker RAM role created for the cluster will be used.
* If specified, the RAM role must be a **general-purpose service role**, and its **trusted service** must be configured as **ECS (Elastic Compute Service)**. For more information, see [Create a general-purpose service role](~~116800~~). When the specified RAM role is not the default Worker RAM role created for the cluster, its name must not start with `KubernetesMasterRole-` or `KubernetesWorkerRole-`.

-> **NOTE:** This parameter is supported only for ACK managed clusters of version 1.22 or later.

* `rds_instances` - (Optional, List) A list of RDS instances.
* `resource_group_id` - (Optional, Computed) The resource group ID of the node pool. Instances scaled out by the node pool belong to this resource group.
Each resource can belong to only one resource group. Depending on your business scenario, you can map resource groups to projects, applications, organizations, or other logical groupings.
* `rolling_policy` - (Optional, Set) Rolling update policy. See [`rolling_policy`](#rolling_policy) below.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `runtime_name` - (Optional, Computed) Container runtime name. ACK supports the following three container runtimes:
  - containerd: Recommended and supported by all cluster versions.
  - Sandboxed-Container.runv: Secure sandboxed containers that provide higher isolation, supported in clusters of version 1.31 and earlier.
  - docker: No longer maintained, supported only in clusters of version 1.22 and earlier.

Default value: containerd.
* `runtime_version` - (Optional, Computed) The container runtime version.
* `scaling_config` - (Optional, Computed, Set) Auto scaling configuration. See [`scaling_config`](#scaling_config) below.
* `scaling_policy` - (Optional, Computed) Scaling mode. Valid values:
  - `release`: Standard mode. Scaling is performed by creating or releasing ECS instances based on resource usage.
  - `recycle`: Rapid mode. Scaling is performed by creating, stopping, and starting instances to accelerate subsequent scaling operations. (When an instance is stopped, you are not charged for computing resources but only for storage, except for instances with local disks.)

Default value: `release`.
* `security_group_id` - (Optional, ForceNew, Computed, Deprecated since v1.145.0) [This field is deprecated.]
Security group ID of the node pool. When multiple security groups are associated with the node pool, this value corresponds to the first entry in `security_group_ids`.
* `security_group_ids` - (Optional, ForceNew, Computed, List) A list of security group IDs. You must specify either `security_group_id` or `security_group_ids`. We recommend that you use `security_group_ids`. If both `security_group_id` and `security_group_ids` are specified, `security_group_ids` takes precedence.
* `security_hardening_os` - (Optional, ForceNew) Alibaba Cloud OS security hardening. Valid values:
  - `true`: Enable Alibaba Cloud OS security hardening.
  - `false`: Disable Alibaba Cloud OS security hardening.

Default value: `false`.
* `soc_enabled` - (Optional, ForceNew) Specifies whether to enable Classified Protection hardening. You can enable Classified Protection hardening for nodes only when Alibaba Cloud Linux 2 or Alibaba Cloud Linux 3 is selected as the system image. Alibaba Cloud provides baseline check standards and scanning tools compliant with the Classified Protection 2.0 Level 3 requirements for Alibaba Cloud Linux 2 and Alibaba Cloud Linux 3 images.
* `spot_instance_pools` - (Optional, Int) Specify the number of available instance types. The auto scaling group creates preemptible instances evenly across the lowest-cost instance types. Valid values: [1, 10].
* `spot_instance_remedy` - (Optional) Specifies whether to enable replenishment of preemptible instances. After this feature is enabled, when the system sends a notification that a preemptible instance will be reclaimed, the scaling group attempts to create a new instance to replace the instance that is about to be reclaimed. Valid values:
  - `true`: Enables replenishment of preemptible instances.
  - `false`: Disables replenishment of preemptible instances.
* `spot_price_limit` - (Optional, List) Current price limit configuration for each spot instance type. See [`spot_price_limit`](#spot_price_limit) below.
* `spot_strategy` - (Optional, Computed) The spot instance strategy. Valid values:
  - `NoSpot`: Regular pay-as-you-go instance.
  - `SpotWithPriceLimit`: Sets a maximum price for spot instances.
  - `SpotAsPriceGo`: The system automatically bids based on the current market price.

For more information, see [Spot Instances](~~165053~~).
* `system_disk_bursting_enabled` - (Optional) Specifies whether to enable burst performance for the node system disk. Valid values:
  - true: Enabled.
  - false: Disabled.

This parameter can be configured only when `system_disk_category` is set to `cloud_auto`. For more information, see [ESSD AutoPL Cloud Disks](~~368372~~).
* `system_disk_categories` - (Optional, Computed, List) Multiple disk types for the system disk. If the system cannot use a higher-priority disk type, it automatically attempts to create the system disk using the next available disk type.
* `system_disk_category` - (Optional, Computed) The system disk type for nodes. Valid values:
  - `cloud_efficiency`: Ultra disk.
  - `cloud_ssd`: SSD cloud disk.
  - `cloud_essd`: ESSD cloud disk.
  - `cloud_auto`: ESSD AutoPL cloud disk.
  - `cloud_essd_entry`: ESSD Entry cloud disk.

Default value: `cloud_efficiency`.
* `system_disk_encrypt_algorithm` - (Optional) The encryption algorithm used for the system disk. Valid values: aes-256.
* `system_disk_encrypted` - (Optional) Specifies whether to encrypt the system disk. Valid values:  
  - true: Encrypts the system disk.  
  - false: Does not encrypt the system disk.
* `system_disk_kms_key` - (Optional) KMS key ID used for the system disk.
* `system_disk_performance_level` - (Optional) Disk performance level for the node's system disk. This setting applies only to ESSD disks. The performance level is related to the disk size. For more information, see [ESSD Cloud Disks](~~122389~~).
  - PL0: Medium concurrent I/O performance with relatively stable read/write latency.
  - PL1: Medium concurrent I/O performance with relatively stable read/write latency.
  - PL2: High concurrent I/O performance with stable read/write latency.
  - PL3: Extremely high concurrent I/O performance with extremely stable read/write latency.
* `system_disk_provisioned_iops` - (Optional, Int) Provisioned read/write IOPS for the node system disk. This parameter applies only when the disk type is cloud_auto.
* `system_disk_size` - (Optional, Int) The size of the system disk for nodes, in GiB.  
Valid values: [20, 2048].  
The value must be greater than or equal to max{20, ImageSize}.  
Default value: max{40, image size corresponding to the ImageId parameter}.  
* `system_disk_snapshot_policy_id` - (Optional) Snapshot policy for the system disk.  
* `tags` - (Optional, Map) Adds tags only to ECS instances.
Tag keys must be unique and can be up to 128 characters in length. Neither tag keys nor tag values can start with "aliyun" or "acs:", or contain "https://" or "http://".
* `taints` - (Optional, List) Node taint information. Taints and tolerations work together to prevent pods from being scheduled onto unsuitable nodes. For more information, see [taint-and-toleration](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/). See [`taints`](#taints) below.
* `tee_config` - (Optional, ForceNew, Computed, Set) Confidential computing cluster configuration. See [`tee_config`](#tee_config) below.
* `type` - (Optional, ForceNew, Computed, Available since v1.252.0) The node pool type. Valid values:
  - `ess`: Standard node pool (supports managed features and auto scaling).
  - `edge`: Edge node pool.
  - `lingjun`: Lingjun node pool.
* `unschedulable` - (Optional) Specifies whether newly scaled-out nodes are unschedulable.
  - true: Unschedulable.
  - false: Schedulable.
* `update_nodes` - (Optional) Synchronize updates to node labels and taints.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `upgrade_policy` - (Optional, Set, Available since v1.269.0) Node pool upgrade policy See [`upgrade_policy`](#upgrade_policy) below.

-> **NOTE:** This parameter only applies during resource update. If modified in isolation without other property changes, Terraform will not trigger any action.

* `user_data` - (Optional) Custom data for the node pool, which refers to scripts executed after node initialization. For more information, see [Generate instance custom data](~~49121~~).  
* `vswitch_ids` - (Optional, List) List of virtual switch IDs. The number of IDs must be in the range [1, 8].

-> **NOTE:**  To ensure high availability, we recommend selecting virtual switches in different zones.


* `kms_encrypted_password` - (Optional, Available since v1.177.0) An KMS encrypts password used to a cs kubernetes. You have to specify one of `password` `key_name` `kms_encrypted_password` fields.
* `kms_encryption_context` - (Optional, Available since v1.177.0) An KMS encryption context used to decrypt `kms_encrypted_password` before creating or updating a cs kubernetes with `kms_encrypted_password`. See [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). It is valid when `kms_encrypted_password` is set.
* `format_disk` - (Optional, Available since v1.127.0) After you select this check box, if data disks have been attached to the specified ECS instances and the file system of the last data disk is uninitialized, the system automatically formats the last data disk to ext4 and mounts the data disk to /var/lib/docker and /var/lib/kubelet. The original data on the disk will be cleared. Make sure that you back up data in advance. If no data disk is mounted on the ECS instance, no new data disk will be purchased. Default is `false`.
* `instances` - (Optional, Available since v1.127.0) The instance list. Add existing nodes under the same cluster VPC to the node pool. 
* `node_count` (Optional, Deprecated) The worker node number of the node pool. From version 1.111.0, `node_count` is not required.
* `keep_instance_name` - (Optional, Available since v1.127.0) Add an existing instance to the node pool, whether to keep the original instance name. It is recommended to set to `true`.
* `rollout_policy` - (Optional, Deprecated since 1.185.0) Rollout policy is used to specify the strategy when the node pool is rolling update. This field works when node pool updating. Please use `rolling_policy` to instead it from provider version 1.185.0. See [`rollout_policy`](#rollout_policy) below.

The following arguments will be discarded. Please use new fields as soon as possible:
* `name` - (Deprecated since v1.219.0). Field 'name' has been deprecated from provider version 1.219.0. New field 'node_pool_name' instead.

### `auto_mode`

The auto_mode supports the following:
* `enabled` - (Optional, ForceNew) Specifies whether to enable intelligent managed mode.

### `data_disks`

The data_disks supports the following:
* `auto_format` - (Optional, Available since v1.229.0) Specifies whether to mount the data disk.
* `auto_snapshot_policy_id` - (Optional) The automatic snapshot policy applied when disk backup is enabled.
* `bursting_enabled` - (Optional) Specifies whether to enable burst performance (Burst) for the data disk. This parameter applies only when the disk type is `cloud_auto`.
* `category` - (Optional) The type of data disk for nodes. Valid values:
  - `cloud`: Basic cloud disk.
  - `cloud_efficiency`: Ultra disk.
  - `cloud_ssd`: SSD cloud disk.
  - `cloud_essd`: ESSD cloud disk.
  - `cloud_auto`: ESSD AutoPL cloud disk.
  - `cloud_essd_entry`: ESSD Entry cloud disk.
  - `elastic_ephemeral_disk_premium`: Premium elastic ephemeral disk.
  - `elastic_ephemeral_disk_standard`: Standard elastic ephemeral disk.

Default value: `cloud_efficiency`.
* `device` - (Optional) If you do not specify this parameter, the system automatically assigns a device name when creating an ECS instance. The assignment starts from /dev/xvdb and ends at /dev/xvdz.
* `encrypted` - (Optional) Specifies whether to encrypt the data disk.
* `file_system` - (Optional, Available since v1.229.0) File system type to mount. This parameter takes effect only when auto_format is true. Valid values: [ext4, xfs].
* `kms_key_id` - (Optional) The KMS key ID associated with the data disk.
* `mount_target` - (Optional, Available since v1.229.0) The mount path. This parameter takes effect only when auto_format is set to true.
* `name` - (Optional, Computed) The name must be 2 to 128 characters in length and can contain letters, Chinese characters, digits, colons (:), underscores (_), or hyphens (-). It must start with a letter or a Chinese character and cannot start with http:// or https://.
* `performance_level` - (Optional) The performance level of the data disk. This parameter applies only to ESSD disks.
* `provisioned_iops` - (Optional, Int) The provisioned read/write IOPS for the data disk, which is configured when the disk type is cloud_auto.
* `size` - (Optional, Int) The size of the data disk. Valid values: 40 to 32767.
* `snapshot_id` - (Optional) When this parameter is specified, the DataDisks.Size parameter is ignored, and the actual size of the created disk matches the size of the specified snapshot. If the snapshot was created on or before July 15, 2013, the request is rejected, and the response includes the error code InvalidSnapshot.TooOld.

### `eflo_node_group`

The eflo_node_group supports the following:
* `cluster_id` - (Optional, Available since v1.252.0) The Lingjun cluster ID that must be associated when creating a Lingjun node pool.
* `group_id` - (Optional) The Lingjun group ID of the Lingjun cluster that must be associated when creating a Lingjun node pool.

### `instance_metadata_options`

The instance_metadata_options supports the following:
* `http_tokens` - (Optional, ForceNew) ECS instance metadata access mode configuration. Valid values:

  - `optional`: Compatible with both standard mode and hardened mode.
  - `required`: Enables hardened mode only (IMDSv2). After this mode is enabled, applications on the node cannot access ECS instance metadata through standard mode. Before enabling this mode, ensure that components in the cluster and the operating system versions meet the minimum version requirements. For more information, see [Access ECS instance metadata in hardened mode only](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/security-and-compliance/secure-access-to-ecs-instance-metadata).

Default value: `optional`.

This parameter is supported only for ACK managed clusters of version 1.28 or later.

### `instance_patterns`

The instance_patterns supports the following:
* `cores` - (Optional, Int, Available since v1.266.0) The number of vCPU cores for the instance type. Example value: 8.
* `cpu_architectures` - (Optional, List, Available since v1.266.0) The CPU architecture of the instance. Valid values:
  - X86
  - ARM.
* `excluded_instance_types` - (Optional, List, Available since v1.266.0) Instance types to exclude. You can use wildcards (*) to exclude a single instance type or an entire instance type family. Examples:
  - ecs.c6.large: excludes the ecs.c6.large instance type.
  - ecs.c6.*: excludes all instance types in the c6 instance family.
* `instance_categories` - (Optional, List, Available since v1.266.0) Instance category. Valid values:
  - General-purpose: general purpose instances.
  - Compute-optimized: compute optimized instances.
  - Memory-optimized: memory optimized instances.
  - Big data: big data instances.
  - Local SSDs: local SSD instances.
  - High Clock Speed: high clock speed instances.
  - Enhanced: enhanced instances.
  - Shared: shared instances.
  - ECS Bare Metal: Elastic Bare Metal instances.
  - High Performance Compute: high performance computing instances.
* `instance_family_level` - (Required, Available since v1.266.0) Instance family level. Valid values:
  - EntryLevel: entry-level, which refers to shared instances. These offer lower costs but do not guarantee consistent compute performance. They are suitable for workloads with low average CPU utilization. For more information, see Shared instances.
  - EnterpriseLevel: enterprise-level instances. These provide stable performance and dedicated resources, making them suitable for business scenarios that require high stability. For more information, see Instance families.
* `instance_type_families` - (Optional, List, Available since v1.266.0) Specifies instance type families. Example value: ["ecs.g8i","ecs.c8i"].
* `max_cpu_cores` - (Optional, Int, Available since v1.266.0) The maximum number of vCPU cores for the instance type. Example value: 8. The value of MaxCpuCores cannot exceed four times the value of MinCpuCores.
* `max_memory_size` - (Optional, Float, Available since v1.266.0) The maximum memory of the instance type. Unit: GiB. Example value: 8. MaxMemorySize cannot exceed four times MinMemorySize.
* `memory` - (Optional, Float, Available since v1.266.0) Memory size of the instance type, in GiB. Example value: 8.
* `min_cpu_cores` - (Optional, Int, Available since v1.266.0) The minimum number of vCPU cores for the instance type. Example value: 4. MaxCpuCores cannot exceed four times MinCpuCores.
* `min_memory_size` - (Optional, Float) The minimum memory of the instance type. Unit: GiB. Example value: 4. MaxMemorySize cannot exceed four times MinMemorySize.

### `kubelet_configuration`

The kubelet_configuration supports the following:
* `allowed_unsafe_sysctls` - (Optional, List) Allowlist of allowed sysctl patterns.
* `cluster_dns` - (Optional, List, Available since v1.242.0) A list of IP addresses of the cluster DNS servers.
* `container_log_max_files` - (Optional) The maximum number of log files allowed per container.  
* `container_log_max_size` - (Optional) The maximum size a log file can reach before it is rotated.
* `container_log_max_workers` - (Optional, Available since v1.242.0) Specifies the maximum number of concurrent workers required to perform log rotation operations.
* `container_log_monitor_interval` - (Optional, Available since v1.242.0) Specifies the duration for monitoring container logs to perform log rotation operations.
* `cpu_cfs_quota` - (Optional, Available since v1.242.0) The switch for CPU CFS quota enforcement.
* `cpu_cfs_quota_period` - (Optional, Available since v1.242.0) The CPU CFS quota period value.  
* `cpu_manager_policy` - (Optional) CPU manager policy.
* `event_burst` - (Optional) The upper limit on the burst count of event records.
* `event_record_qps` - (Optional) The maximum number of events that can be generated per second.  
* `eviction_hard` - (Optional, Map) A set of hard thresholds that trigger pod eviction.
* `eviction_soft` - (Optional, Map) A set of soft eviction thresholds.  
* `eviction_soft_grace_period` - (Optional, Map) Sets a set of eviction grace periods.
* `feature_gates` - (Optional, Map) Feature gates used to enable experimental features.
* `image_gc_high_threshold_percent` - (Optional, Available since v1.242.0) When image usage exceeds this threshold, image garbage collection continues to run.
* `image_gc_low_threshold_percent` - (Optional, Available since v1.242.0) Image garbage collection is not performed when image usage falls below this threshold.
* `kube_api_burst` - (Optional) The maximum number of burst requests per second sent to the API server.
* `kube_api_qps` - (Optional) The number of queries per second for communication with the API Server.
* `kube_reserved` - (Optional, Map) Resource configuration reserved for the Kubernetes system.  
* `max_pods` - (Optional) The maximum number of pods that can run.  
* `memory_manager_policy` - (Optional, Available since v1.242.0) The policy to be used by the memory manager.  
* `pod_pids_limit` - (Optional, Available since v1.242.0) The maximum number of PIDs that can be used by a pod.
* `read_only_port` - (Optional) Read-only port number.
* `registry_burst` - (Optional) Maximum number of burst image pulls.
* `registry_pull_qps` - (Optional) QPS limit for image registry pulls.
* `reserved_memory` - (Optional, List, Available since v1.242.0) Memory reserved for NUMA nodes. See [`reserved_memory`](#kubelet_configuration-reserved_memory) below.
* `serialize_image_pulls` - (Optional) Whether to pull images one at a time.
* `server_tls_bootstrap` - (Optional, Available since v1.266.0) Enables the kubelet server certificate issuance and rotation feature through Certificate Signing Requests (CSRs).
* `system_reserved` - (Optional, Map) Resource configuration reserved for the system.
* `topology_manager_policy` - (Optional, Available since v1.242.0) Name of the topology manager policy to use.
* `tracing` - (Optional, Set) Versioned configuration information for the OpenTelemetry tracing client. See [`tracing`](#kubelet_configuration-tracing) below.

### `kubelet_configuration-reserved_memory`

The kubelet_configuration-reserved_memory supports the following:
* `limits` - (Optional, Map, Available since v1.242.0) Memory resource limits.
* `numa_node` - (Optional, Int) NUMA node.

### `kubelet_configuration-tracing`

The kubelet_configuration-tracing supports the following:
* `endpoint` - (Optional, Available since v1.242.0) The endpoint of the collector.
* `sampling_rate_per_million` - (Optional) The number of samples to collect per million spans.

### `labels`

The labels supports the following:
* `key` - (Required) Key.
* `value` - (Optional) Value.

### `management`

The management supports the following:
* `auto_repair` - (Optional, Computed) Specifies whether to automatically repair nodes. This parameter takes effect only when `enable=true`.
  - `true`: Enables automatic repair.
  - `false`: Disables automatic repair.

When `enable=true`, the default value is `true`. When `enable=false`, the default value is `false`.
* `auto_repair_policy` - (Optional, Computed, Set) Automatic Node Repair Policy   See [`auto_repair_policy`](#management-auto_repair_policy) below.
* `auto_upgrade` - (Optional, Computed) Specifies whether to automatically upgrade nodes. This parameter takes effect only when `enable=true`.
  - `true`: Enables automatic upgrade.
  - `false`: Disables automatic upgrade.

When `enable=true`, the default value is `true`. When `enable=false`, the default value is `false`.
* `auto_upgrade_policy` - (Optional, Computed, Set) Auto upgrade policy. See [`auto_upgrade_policy`](#management-auto_upgrade_policy) below.
* `auto_vul_fix` - (Optional, Computed) Specifies whether to automatically fix CVEs. This parameter takes effect only when `enable=true`.  
  - `true`: Allows automatic CVE fixes.  
  - `false`: Disallows automatic CVE fixes.  
* `auto_vul_fix_policy` - (Optional, Computed, Set) Automatic CVE Fix Policy See [`auto_vul_fix_policy`](#management-auto_vul_fix_policy) below.
* `enable` - (Optional, Computed) Specifies whether to enable the managed node pool. Valid values:  
  - `true`: Enables the managed node pool.  
  - `false`: Disables the managed node pool. Other related configurations take effect only when `enable=true`.  

Default value: `false`.  
* `max_unavailable` - (Optional, Int) The maximum number of unavailable nodes.
Value range: [1, 1000].
Default value: 1.
* `surge` - (Optional, Int, Deprecated since v1.219.0) The number of additional nodes. You must specify either this parameter or `surge_percentage`.

Nodes become unavailable during upgrades. You can create additional nodes to compensate for the cluster workload.

-> **NOTE:**  We recommend that the number of additional nodes does not exceed the current number of nodes.

* `surge_percentage` - (Optional, Int, Deprecated since v1.219.0) The percentage of additional nodes. You must specify either this parameter or `surge`.
Number of additional nodes = Surge percentage × Number of existing nodes. For example, if you set the surge percentage to 50% and there are 6 existing nodes, the number of additional nodes will be 50% × 6, which equals 3 additional nodes.

### `management-auto_repair_policy`

The management-auto_repair_policy supports the following:
* `restart_node` - (Optional, Computed) Specifies whether to allow node restarts. This parameter takes effect only when `auto_repair=true`.  
  - `true`: Allows node restarts.  
  - `false`: Disallows node restarts.  

### `management-auto_upgrade_policy`

The management-auto_upgrade_policy supports the following:
* `auto_upgrade_kubelet` - (Optional, Computed) Specifies whether to allow automatic kubelet upgrades. This parameter takes effect only when `auto_upgrade=true`. Valid values:
  - `true`: Allows automatic kubelet upgrades.
  - `false`: Disallows automatic kubelet upgrades.

When `auto_upgrade=true`, the default value is `true`. When `auto_upgrade=false`, the default value is `false`.

### `management-auto_vul_fix_policy`

The management-auto_vul_fix_policy supports the following:
* `restart_node` - (Optional, Computed) Specifies whether to allow node restarts. This parameter takes effect only when `auto_vul_fix=true`. Valid values:  
  - `true`: Allows node restarts.  
  - `false`: Disallows node restarts.  

When `auto_vul_fix=true`, the default value is `false`. When `auto_vul_fix=false`, the default value is also `false`.  
* `vul_level` - (Optional, Computed) The vulnerability severity levels that are allowed for automatic repair, separated by commas.  
  - `asap`: High  
  - `later`: Medium  
  - `nntf`: Low  

### `private_pool_options`

The private_pool_options supports the following:
* `private_pool_options_id` - (Optional) Private node pool ID. When `match_criteria` is set to `Target`, you must further specify the private pool ID.
* `private_pool_options_match_criteria` - (Optional) Private node pool type, which specifies the private pool capacity option for instance launch. After an elasticity assurance or capacity reservation takes effect, it generates private pool capacity that can be selected when launching instances. Valid values:
  - `Open`: Open mode. Automatically matches open-type private pool capacity. If no matching private pool capacity is available, the instance is launched using public pool resources.
  - `Target`: Targeted mode. Launches the instance using the specified private pool capacity. If the specified private pool capacity is unavailable, the instance launch fails.
  - `None`: Disabled mode. The instance launch does not use any private pool capacity.

### `rolling_policy`

The rolling_policy supports the following:
* `batch_interval` - (Optional, Available since v1.269.0) Time interval between upgrade batches.
* `max_parallelism` - (Optional, Int) Node updates in a node pool are performed in batches. This parameter defines the maximum number of nodes that can be updated in parallel per batch.
Valid range: [1, 10].
Default value: 10.
* `node_names` - (Optional, List, Available since v1.269.0) List of specific nodes to upgrade.
* `pause_policy` - (Optional, Available since v1.269.0) Automatic pause policy during node upgrades.

### `scaling_config`

The scaling_config supports the following:
* `eip_bandwidth` - (Optional, Int) [This field is deprecated.] Use `internet_charge_type` and `internet_max_bandwidth_out` instead.  
Peak bandwidth of the EIP.
Valid range: [1, 100]. Unit: Mbps.
* `eip_internet_charge_type` - (Optional) [This field is deprecated.] Use `internet_charge_type` and `internet_max_bandwidth_out` instead.

Billing method for the EIP. Valid values:
  - `PayByBandwidth`: billed based on fixed bandwidth.
  - `PayByTraffic`: billed based on actual traffic usage.

Default value: `PayByBandwidth`.
* `enable` - (Optional) Specifies whether to enable auto scaling. Valid values:
  - `true`: Enables auto scaling for the node pool. When the cluster capacity is insufficient to schedule application pods, ACK automatically scales the node resources based on the configured minimum and maximum instance counts. For clusters of version 1.24 or later, instant node elasticity is enabled by default. For clusters earlier than version 1.24, auto scaling is enabled by default. For more information, see [Node Auto Scaling](~~2746785~~).
  - `false`: Disables auto scaling. ACK adjusts the number of nodes in the node pool to match the desired node count and maintains this number consistently.

When this parameter is set to `false`, other parameters within `auto_scaling` do not take effect.

Default value: `false`.
* `is_bond_eip` - (Optional) [This field is deprecated.] Use `internet_charge_type` and `internet_max_bandwidth_out` instead.  
Whether to bind an EIP. Valid values:
  - `true`: Bind an EIP.
  - `false`: Do not bind an EIP.

Default value: `false`.
* `max_size` - (Optional, Int) The maximum number of instances that can be added through scaling in the node pool, excluding existing instances. This parameter takes effect only when `enable=true`.
Valid range: [min_instances, 2000]. Default value: 0.
* `min_size` - (Optional, Int) The minimum number of scalable instances in the node pool, excluding your existing instances. This parameter takes effect only when `enable=true`.

Valid range: [0, max_instances]. Default value: 0.

-> **NOTE:**  - If the minimum number of instances is greater than 0, the specified number of ECS instances will be automatically created after the scaling group takes effect.

-> **NOTE:**  - We recommend that the maximum number of instances not be set lower than the current number of nodes in the node pool. Otherwise, enabling auto scaling will immediately trigger scale-in of the node pool.

* `type` - (Optional) Instance type for auto scaling. This parameter takes effect only when `enable=true`. Valid values:
  - `cpu`: General-purpose instance.
  - `gpu`: GPU instance.
  - `gpushare`: Shared GPU instance.
  - `spot`: Spot instance.

Default value: `cpu`.

-> **NOTE:** You cannot modify this field after the node pool is created.>


### `spot_price_limit`

The spot_price_limit supports the following:
* `instance_type` - (Optional) Spot instance type.
* `price_limit` - (Optional) Maximum price per instance.

### `taints`

The taints supports the following:
* `effect` - (Optional) The taint effect policy.
* `key` - (Required) The key.
* `value` - (Optional) The value.

### `tee_config`

The tee_config supports the following:
* `tee_enable` - (Optional, ForceNew) Specifies whether to enable confidential computing.
  - true: Enables confidential computing.
  - false: Disables confidential computing.

### `upgrade_policy`

The upgrade_policy supports the following:
* `image_id` - (Optional, Available since v1.269.0) System image ID of the node
* `kubernetes_version` - (Optional, Available since v1.269.0) Kubernetes version of the node
* `runtime` - (Optional, Available since v1.269.0) Runtime type of the node
* `runtime_version` - (Optional, Available since v1.269.0) Runtime version of the node
* `use_replace` - (Optional, Available since v1.269.0) Specifies whether to use disk replacement for upgrades.

### `rollout_policy`

The rollout_policy mapping supports the following:
* `max_unavailable` - (Optional, Deprecated since 1.185.0) Maximum number of unavailable nodes during rolling upgrade. The value of this field should be greater than `0`, and if it's set to a number less than or equal to `0`, the default setting will be used. Please use `max_parallelism` to instead it from provider version 1.185.0.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<cluster_id>:<node_pool_id>`.
* `node_pool_id` - The node pool ID.
* `scaling_group_id` - The ID of the scaling group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Node Pool.
* `delete` - (Defaults to 5 mins) Used when delete the Node Pool.
* `update` - (Defaults to 15 mins) Used when update the Node Pool.

## Import

Container Service for Kubernetes (ACK) Node Pool can be imported using the id, e.g.

```shell
$ terraform import alicloud_cs_kubernetes_node_pool.example <cluster_id>:<node_pool_id>
```