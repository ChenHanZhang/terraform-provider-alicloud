---
subcategory: "Function Compute Service V3 (FCV3)"
layout: "alicloud"
page_title: "Alicloud: alicloud_fcv3_function"
description: |-
  Provides a Alicloud Function Compute Service V3 (FCV3) Function resource.
---

# alicloud_fcv3_function

Provides a Function Compute Service V3 (FCV3) Function resource.

Function Compute schedules and runs resources on a per-function basis. A Function Compute function consists of function code and function configuration.

For information about Function Compute Service V3 (FCV3) Function and how to use it, see [What is Function](https://www.alibabacloud.com/help/en/functioncompute/developer-reference/api-fc-2023-03-30-getfunction).

-> **NOTE:** Available since v1.228.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "random_uuid" "default" {
}

resource "alicloud_oss_bucket" "default" {
  bucket = "${var.name}-${random_uuid.default.result}"
}

resource "alicloud_oss_bucket_object" "default" {
  bucket  = alicloud_oss_bucket.default.bucket
  key     = "FCV3Py39.zip"
  content = "print('hello')"
}

resource "alicloud_fcv3_function" "default" {
  description = "Create"
  memory_size = "512"
  layers = [
    "acs:fc:cn-shanghai:official:layers/Python39-Aliyun-SDK/versions/3"
  ]
  timeout   = "3"
  runtime   = "custom.debian10"
  handler   = "index.handler"
  disk_size = "512"
  custom_runtime_config {
    command = [
      "python",
      "-c",
      "example"
    ]
    args = [
      "app.py",
      "xx",
      "x"
    ]
    port = "9000"
    health_check_config {
      http_get_url          = "/ready"
      initial_delay_seconds = "1"
      period_seconds        = "10"
      success_threshold     = "1"
      timeout_seconds       = "1"
      failure_threshold     = "3"
    }

  }

  log_config {
    log_begin_rule = "None"
  }

  code {
    oss_bucket_name = alicloud_oss_bucket.default.bucket
    oss_object_name = alicloud_oss_bucket_object.default.key
    checksum        = "4270285996107335518"
  }

  instance_lifecycle_config {
    initializer {
      timeout = "1"
      handler = "index.init"
    }

    pre_stop {
      timeout = "1"
      handler = "index.stop"
    }

  }

  cpu                  = "0.5"
  instance_concurrency = "2"
  function_name        = "${var.name}-${random_uuid.default.result}"
  environment_variables = {
    "EnvKey" = "EnvVal"
  }
  internet_access = "true"
}
```

## Argument Reference

The following arguments are supported:
* `code` - (Optional, Set) The function code ZIP package. You must specify either code or customContainerConfig. See [`code`](#code) below.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `cpu` - (Optional, Computed, Float) CPU specification for the function, measured in vCPUs, and must be a multiple of 0.05 vCPU.
* `custom_container_config` - (Optional, Set) The configuration for a custom container runtime. After successful configuration, the function can execute using a custom container image. Specify either code or customContainerConfig, but not both. See [`custom_container_config`](#custom_container_config) below.
* `custom_dns` - (Optional, Set) Custom DNS configuration for the function. See [`custom_dns`](#custom_dns) below.
* `custom_runtime_config` - (Optional, Set) Custom runtime configuration. See [`custom_runtime_config`](#custom_runtime_config) below.
* `description` - (Optional) The description of the function. Function Compute does not use this property value, but we recommend that you provide a concise and clear description for the function.
* `disk_size` - (Optional, Computed, Int) The disk size of the function, in MB. Valid values are 512 MB or 10240 MB.
* `environment_variables` - (Optional, Map) Environment variables configured for the function. You can retrieve the values of these environment variables within the function.
* `function_name` - (Optional, ForceNew, Computed) Name of the function that is prohibited from being invoked.
* `gpu_config` - (Optional, Set) GPU configuration for the function. See [`gpu_config`](#gpu_config) below.
* `handler` - (Required) The handler, which serves as the entry point invoked by the Function Compute system to execute your function.  
* `idle_timeout` - (Optional, ForceNew, Int, Available since v1.266.0) If an instance remains idle for longer than this duration, it is destroyed. A value of - 1 removes this threshold and uses the system default behavior.
* `instance_concurrency` - (Optional, Computed, Int) Maximum instance concurrency.
* `instance_isolation_mode` - (Optional, Computed, Available since v1.256.0) Instance isolation mode.
* `instance_lifecycle_config` - (Optional, Set) Instance lifecycle callback method configuration.   See [`instance_lifecycle_config`](#instance_lifecycle_config) below.
* `internet_access` - (Optional, Computed) Specifies whether the function is allowed to access the Internet.
* `invocation_restriction` - (Optional, Set) Details about invocation restrictions. See [`invocation_restriction`](#invocation_restriction) below.
* `layers` - (Optional, List) A list of layers.
* `log_config` - (Optional, Set) Logs generated by the function are written to the configured Logstore. See [`log_config`](#log_config) below.
* `memory_size` - (Optional, Computed, Int) The memory size of the function, specified in MB. The memory size must be a multiple of 64 MB, with a minimum of 128 MB and a maximum of 32 GB. Additionally, the ratio of CPU to memory size (in GB) must be between 1:1 and 1:4.  
* `nas_config` - (Optional, Computed, Set) NAS configuration. After this parameter is configured, the function can access the specified NAS resources. See [`nas_config`](#nas_config) below.
* `oss_mount_config` - (Optional, Computed, Set) OSS mount configuration. See [`oss_mount_config`](#oss_mount_config) below.
* `qualifier` - (Optional, Available since v1.287.0) The version or alias of the function.

-> **NOTE:** This parameter configures deletion behavior and is only evaluated when Terraform attempts to destroy the resource. Changes to this parameter during updates are stored but have no immediate effect.

* `resource_group_id` - (Optional, Computed, Available since v1.260.0) The resource group ID.
* `role` - (Optional) The RAM role that the user grants to Function Compute. After this role is configured, Function Compute assumes the role to generate temporary security credentials. The function can use these temporary credentials to access specified Alibaba Cloud services, such as OSS and Tablestore (OTS).
* `runtime` - (Required) The runtime type of the function.
* `session_affinity` - (Optional, Computed, Available since v1.256.0) Session affinity policy for Function Compute invocation requests. To enable request affinity for the MCP SSE protocol, set this parameter to MCP_SSE. For cookie-based affinity, set it to GENERATED_COOKIE. For header-based affinity, set it to HEADER_FIELD. If this parameter is not set or is set to NONE, no session affinity is applied, and requests are routed according to the default scheduling policy of Function Compute.
* `session_affinity_config` - (Optional, Available since v1.256.0) When session affinity is configured, the corresponding affinity configuration must be specified. For MCP_SSE affinity, configure MCPSSESessionAffinityConfig. For Cookie affinity, configure CookieSessionAffinityConfig. For Header Field affinity, configure HeaderFieldSessionAffinityConfig.
* `tags` - (Optional, Map, Available since v1.242.0) A resource property field representing resource tags.
* `timeout` - (Optional, Computed, Int) The maximum execution duration of the function, in seconds.  
* `vpc_config` - (Optional, Computed, Set) VPC configuration. After this parameter is configured, the function can access the specified VPC resources. See [`vpc_config`](#vpc_config) below.

### `code`

The code supports the following:
* `checksum` - (Optional) The CRC-64 checksum of the function code package.
* `oss_bucket_name` - (Optional) The name of the OSS bucket that stores the function code ZIP package.
* `oss_object_name` - (Optional) The name of the OSS object that stores the function code ZIP package.
* `zip_file` - (Optional) The Base64-encoded content of the function code ZIP package.

### `custom_container_config`

The custom_container_config supports the following:
* `acceleration_type` - (Optional, Deprecated since v1.287.0) Specifies whether to enable image acceleration. Default: The default value, which enables image acceleration. None: Disables image acceleration. (Deprecated.)
* `acr_instance_id` - (Optional, Deprecated since v1.287.0) The ID of the ACR Enterprise Edition image repository. This parameter must be provided when using an ACR Enterprise Edition image. (Deprecated.)
* `command` - (Optional, List) Container startup arguments.
* `entrypoint` - (Optional, List) The container entrypoint command.
* `health_check_config` - (Optional, Computed, Set) Custom health check configuration for the function. See [`health_check_config`](#custom_container_config-health_check_config) below.
* `image` - (Optional) The container image address.
* `port` - (Optional, Int) The listening port of the HTTP server in the custom container runtime.
* `registry_config` - (Optional, Set, Available since v1.287.0) The configuration information for the image registry. See [`registry_config`](#custom_container_config-registry_config) below.

### `custom_container_config-health_check_config`

The custom_container_config-health_check_config supports the following:
* `failure_threshold` - (Optional, Computed, Int) The threshold for the number of consecutive failed health checks. After this number is reached, the system considers the check to have failed. Valid values are 1 to 120. Default value is 3.
* `http_get_url` - (Optional, Computed) The custom health check URL for the container.
* `initial_delay_seconds` - (Optional, Computed, Int) The delay, in seconds, between the container startup and the first health check. Valid values: 0 to 120. Default value: 0.
* `period_seconds` - (Optional, Computed, Int) The health check interval, in seconds. Valid values: 1 to 120. Default value: 3.
* `success_threshold` - (Optional, Computed, Int) The threshold for the number of consecutive successful health checks. After this number is reached, the system considers the check to have succeeded. Valid values are 1 to 120. Default value is 1.
* `timeout_seconds` - (Optional, Computed, Int) The timeout duration, in seconds, for health checks. Valid values: 1 to 3. Default value: 1.

### `custom_container_config-registry_config`

The custom_container_config-registry_config supports the following:
* `auth_config` - (Optional, Set, Available since v1.287.0) Authentication configuration for the custom image repository. See [`auth_config`](#custom_container_config-registry_config-auth_config) below.
* `cert_config` - (Optional, Set, Available since v1.287.0) Certificate configuration for the custom image registry. See [`cert_config`](#custom_container_config-registry_config-cert_config) below.
* `network_config` - (Optional, Set, Available since v1.287.0) Network configuration for the custom image registry. See [`network_config`](#custom_container_config-registry_config-network_config) below.

### `custom_container_config-registry_config-auth_config`

The custom_container_config-registry_config-auth_config supports the following:
* `password` - (Optional, Available since v1.287.0) Password for the image repository.
* `user_name` - (Optional, Available since v1.287.0) Username for the image repository.

### `custom_container_config-registry_config-cert_config`

The custom_container_config-registry_config-cert_config supports the following:
* `insecure` - (Optional, Available since v1.287.0) Whether to skip certificate verification.
* `root_ca_cert_base64` - (Optional, Available since v1.287.0) CA certificate of the image registry.

### `custom_container_config-registry_config-network_config`

The custom_container_config-registry_config-network_config supports the following:
* `security_group_id` - (Optional, Available since v1.287.0) The ID of the security group that can connect to the image registry.
* `v_switch_id` - (Optional, Available since v1.287.0) The VSwitch ID that can connect to the image repository.
* `vpc_id` - (Optional, Available since v1.287.0) The VPC ID that can connect to the image repository.

### `custom_dns`

The custom_dns supports the following:
* `dns_options` - (Optional, List) A list of configuration options from the resolv.conf file. Each item corresponds to a key-value pair in the format key:value, where the key is required. See [`dns_options`](#custom_dns-dns_options) below.
* `name_servers` - (Optional, List) A list of IP addresses of DNS servers.
* `searches` - (Optional, List) A list of DNS search domains.

### `custom_dns-dns_options`

The custom_dns-dns_options supports the following:
* `name` - (Optional) The name of the configuration option.
* `value` - (Optional) The value of the configuration option.

### `custom_runtime_config`

The custom_runtime_config supports the following:
* `args` - (Optional, List) Instance startup arguments.
* `command` - (Optional, List) Instance startup command.
* `health_check_config` - (Optional, Computed, Set) Custom health check configuration for the function. See [`health_check_config`](#custom_runtime_config-health_check_config) below.
* `port` - (Optional, Computed, Int) The listening port of the HTTP server.

### `custom_runtime_config-health_check_config`

The custom_runtime_config-health_check_config supports the following:
* `failure_threshold` - (Optional, Computed, Int) Failure threshold for health checks. After this number of consecutive failures, the system considers the check failed. Valid values: 1 to 120. Default value: 3.
* `http_get_url` - (Optional, Computed) The custom health check URL for the container. The length must not exceed 2,048 characters.
* `initial_delay_seconds` - (Optional, Computed, Int) The delay from container startup to the initiation of the health check. Valid values: 0 to 120. Default value: 0.
* `period_seconds` - (Optional, Computed, Int) Health check period. Valid values: 1 to 120. Default value: 3.
* `success_threshold` - (Optional, Computed, Int) Success threshold for health checks. After this number of consecutive successes, the system considers the check successful. Valid values: 1 to 120. Default value: 1.
* `timeout_seconds` - (Optional, Computed, Int) The health check timeout period. Valid values: 1 to 3. Default value: 1.

### `gpu_config`

The gpu_config supports the following:
* `gpu_memory_size` - (Optional, Int) GPU memory size, in MB. The value must be a multiple of 1024 MB.
* `gpu_type` - (Optional) GPU card architecture.
  - fc.gpu.tesla.1 indicates a GPU instance with the Tesla architecture series (equivalent to the NVIDIA T4).
  - fc.gpu.ampere.1 indicates a GPU instance with the Ampere architecture series (equivalent to the NVIDIA A10).
  - fc.gpu.ada.1 indicates a GPU instance with the Ada Lovelace architecture series.

### `instance_lifecycle_config`

The instance_lifecycle_config supports the following:
* `initializer` - (Optional, Set) Initializer callback method configuration.   See [`initializer`](#instance_lifecycle_config-initializer) below.
* `pre_stop` - (Optional, Set) PreStop callback method configuration.   See [`pre_stop`](#instance_lifecycle_config-pre_stop) below.

### `instance_lifecycle_config-initializer`

The instance_lifecycle_config-initializer supports the following:
* `command` - (Optional, List, Available since v1.260.0) Callback command executed during the function initialization phase. The handler and command for lifecycle callbacks cannot be configured simultaneously; only one can take effect. Configuring both will result in an error message.  
* `handler` - (Optional) Entry point for executing the callback method, similar in meaning to a request handler.
* `timeout` - (Optional, Int) Timeout duration for the callback method, in seconds.

### `instance_lifecycle_config-pre_stop`

The instance_lifecycle_config-pre_stop supports the following:
* `handler` - (Optional) Execution entry point of the callback method, similar in meaning to a request handler.  
* `timeout` - (Optional, Int) Timeout duration for the callback method, in seconds.  

### `invocation_restriction`

The invocation_restriction supports the following:
* `disable` - (Optional, Available since v1.255.0) Specifies whether function invocation is disabled.
* `reason` - (Optional) The reason why function invocation is disabled.

### `log_config`

The log_config supports the following:
* `enable_instance_metrics` - (Optional, Computed) When this feature is enabled, you can view core metrics at the instance level, such as CPU usage, memory usage, network status, and the number of requests processed by the instance. false: The default value, which disables instance-level metrics. true: Enables instance-level metrics.
* `enable_request_metrics` - (Optional, Computed) When this feature is enabled, you can view the duration and memory consumption of each invocation for all functions in the service. false: Disables request-level metrics. true: The default value, which enables request-level metrics.
* `log_begin_rule` - (Optional, Computed) The rule for matching the beginning of log lines.
* `logstore` - (Optional) The name of the Log Service Logstore.
* `project` - (Optional) The name of the Log Service project.

### `nas_config`

The nas_config supports the following:
* `group_id` - (Optional, Computed, Int) Group ID.
* `mount_points` - (Optional, List) A list of mount points. See [`mount_points`](#nas_config-mount_points) below.
* `user_id` - (Optional, Computed, Int) Account ID.

### `nas_config-mount_points`

The nas_config-mount_points supports the following:
* `enable_tls` - (Optional) Mount using transport encryption. Note: Only General-purpose NAS supports transport encryption.
* `mount_dir` - (Optional) Local mount directory.
* `server_addr` - (Optional) NAS server address.

### `oss_mount_config`

The oss_mount_config supports the following:
* `mount_points` - (Optional, List) List of OSS mount points. See [`mount_points`](#oss_mount_config-mount_points) below.

### `oss_mount_config-mount_points`

The oss_mount_config-mount_points supports the following:
* `bucket_name` - (Optional) The OSS bucket to mount.
* `bucket_path` - (Optional) The path within the mounted OSS bucket.
* `endpoint` - (Optional) The OSS endpoint.
* `mount_dir` - (Optional) The mount directory.
* `read_only` - (Optional) Whether the mount point is read-only.

### `vpc_config`

The vpc_config supports the following:
* `security_group_id` - (Optional, Computed) The security group ID.
* `vswitch_ids` - (Optional, List) The list of vSwitches.
* `vpc_id` - (Optional, Computed) The VPC network ID.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `code_size` - The size of the function code package returned by the system, in bytes.
* `create_time` - The creation time of the function.
* `custom_container_config` - The configuration for a custom container runtime.
  * `acceleration_info` - Image acceleration information.
    * `status` - Image acceleration status.
  * `resolved_image_uri` - The actual digest version of the deployed image.
* `function_arn` - A list of resource identifiers.
* `function_id` - A unique ID generated by the system for each function.
* `invocation_restriction` - Details about invocation restrictions.
  * `last_modified_time` - The last modification time of the invocation restriction.
* `last_modified_time` - The time when the function was last updated.
* `last_update_status` - The status of the most recent function update operation.
* `last_update_status_reason` - The reason why the status of the most recent function update operation is currently in its present state.
* `last_update_status_reason_code` - The status code indicating why the most recent function update operation resulted in its current status.
* `state` - Function status.
* `state_reason` - The reason why the function is in its current state.
* `state_reason_code` - The status code indicating why the function is in its current state.
* `tracing_config` - Tracing configuration.
  * `params` - Distributed tracing parameters.
  * `type` - The tracing protocol type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Function.
* `delete` - (Defaults to 5 mins) Used when delete the Function.
* `update` - (Defaults to 5 mins) Used when update the Function.

## Import

Function Compute Service V3 (FCV3) Function can be imported using the id, e.g.

```shell
$ terraform import alicloud_fcv3_function.example <function_name>
```