---
subcategory: "Serverless App Engine (SAE)"
layout: "alicloud"
page_title: "Alicloud: alicloud_sae_application"
description: |-
  Provides a Alicloud Serverless App Engine (SAE) Application resource.
---

# alicloud_sae_application

Provides a Serverless App Engine (SAE) Application resource.



For information about Serverless App Engine (SAE) Application and how to use it, see [What is Application](https://www.alibabacloud.com/help/en/sae/latest/createapplication).

-> **NOTE:** Available since v1.161.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = var.region
}

variable "region" {
  default = "cn-hangzhou"
}

variable "name" {
  default = "tf-example"
}

resource "random_integer" "default" {
  max = 99999
  min = 10000
}

data "alicloud_regions" "default" {
  current = true
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "default" {
  vpc_name   = var.name
  cidr_block = "10.4.0.0/16"
}

resource "alicloud_vswitch" "default" {
  vswitch_name = var.name
  cidr_block   = "10.4.0.0/24"
  vpc_id       = alicloud_vpc.default.id
  zone_id      = data.alicloud_zones.default.zones.0.id
}

resource "alicloud_security_group" "default" {
  vpc_id = alicloud_vpc.default.id
}

resource "alicloud_sae_namespace" "default" {
  namespace_id              = "${data.alicloud_regions.default.regions.0.id}:example${random_integer.default.result}"
  namespace_name            = var.name
  namespace_description     = var.name
  enable_micro_registration = false
}

resource "alicloud_sae_application" "default" {
  app_description   = var.name
  app_name          = "${var.name}-${random_integer.default.result}"
  namespace_id      = alicloud_sae_namespace.default.id
  image_url         = "registry-vpc.${data.alicloud_regions.default.regions.0.id}.aliyuncs.com/sae-demo-image/consumer:1.0"
  package_type      = "Image"
  security_group_id = alicloud_security_group.default.id
  vpc_id            = alicloud_vpc.default.id
  vswitch_id        = alicloud_vswitch.default.id
  timezone          = "Asia/Beijing"
  replicas          = "5"
  cpu               = "500"
  memory            = "2048"
}
```

## Argument Reference

The following arguments are supported:
* `acr_assume_role_arn` - (Optional, Computed) The ARN of the RAM role required when pulling images across accounts. For more information, see [implement cross-account authorization through RAM roles](~~ 223585 ~~).
* `acr_instance_id` - (Optional) The ID of the container Image Service ACR Enterprise Edition instance. This parameter is required when `ImageUrl` is the container Image Service Enterprise Edition.
* `alb_ingress_readiness_gate` - (Optional, Available since v1.280.0) ALB gateway ReadinessGate configuration
* `app_description` - (Optional) Application description information. No more than 1024 characters.
* `app_name` - (Required, ForceNew) Application Name. Combinations of numbers, letters, and dashes (-) are allowed. It must start with a letter and the maximum length is 36 characters.
* `app_source` - (Optional, ForceNew, Available since v1.280.0) Select micro_service, which is the microservice application.
* `auto_config` - (Optional) Automatically configure network settings

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `auto_enable_application_scaling_rule` - (Optional) Whether to automatically enable the auto scaling policy. The values are described as follows:
  - `true`: on.
  - `false`: closed.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `batch_wait_time` - (Optional, Computed, Int) The waiting time between batches when the batch is released, in minutes.
* `change_order_desc` - (Optional) Release sheet description information.

-> **NOTE:** This parameter only takes effect when other resource properties are also modified. Changing this parameter alone will not trigger a resource update.

* `command` - (Optional) Mirror start command. The command must be an executable object in the container. For example: sleep. Setting this command will cause the original startup command of the mirror to become invalid.
* `command_args` - (Optional, Computed, List) Mirror startup command parameters. The parameters required for the above start command. For example: 1d
* `config_map_mount_desc` - (Optional, Computed, List) ConfigMap information. See [`config_map_mount_desc`](#config_map_mount_desc) below.
* `cpu` - (Optional, Int) The CPU required for each instance, in millicores, cannot be 0.
* `custom_host_alias` - (Optional, Computed, List) Custom host mapping in the container. For example: [{"hostName":"samplehost","ip":"127.0.0.1"}] See [`custom_host_alias`](#custom_host_alias) below.
* `custom_image_network_type` - (Optional, Available since v1.280.0) Custom Image Type

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `deploy` - (Optional) This parameter is only valid for applications in the stopped state. If the running application calls the `DeployApplication` interface, it will be redeployed immediately.
  - `true`: The default value. Indicates that the new deployment configuration takes effect immediately, and the instance is pulled up immediately.
  - `false`: indicates that only the new deployment configuration takes effect and the application instance is not pulled up.

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `dotnet` - (Optional, Available since v1.280.0) Version number of the. NET Framework

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `edas_container_version` - (Optional) Pandora Application for Use in the Operation of Environment.
* `enable_ahas` - (Optional, Computed) Whether to access AHAS. The values are described as follows:
  - `true`: access AHAS.
  - `false`: AHAS is not connected.
* `enable_cpu_burst` - (Optional, Available since v1.280.0) Whether the CPU Burst function is enabled

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `enable_grey_tag_route` - (Optional, Computed, Deprecated since v1.280.0) Whether the traffic grayscale rule is enabled. This rule applies only to Spring Cloud and Dubbo applications. The value description is as follows:
  - `true`: enables grayscale rules.
  - `false`: Disables grayscale rules.
* `enable_new_arms` - (Optional, Available since v1.280.0) Whether to enable the new ARMS feature

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `enable_prometheus` - (Optional, Available since v1.280.0) Enable Prometheus custom indicator collection
* `envs` - (Optional, Computed, List) The virtual switch where the elastic network card of the application instance is located. The switch must be located in the aforementioned VPC. The switch also has a binding relationship with the SAE namespace. If it is left blank, the default is the vSwitch ID bound to the namespace. See [`envs`](#envs) below.
* `image_pull_secrets` - (Optional) The ID of the Security Dictionary.
* `image_url` - (Optional) Mirror address. Only Image type applications can configure the mirror address.
* `jar_start_args` - (Optional) The JAR package starts application parameters. Application default startup command: $JAVA_HOME/bin/java $JarStartOptions - jar $CATALINA_OPTS "$package_path" $JarStartArgs
* `jar_start_options` - (Optional) The JAR package starts the application option. Application default startup command: $JAVA_HOME/bin/java $JarStartOptions - jar $CATALINA_OPTS "$package_path" $JarStartArgs
* `jdk` - (Optional) The JDK version that the deployment package depends on. Image type applications are not supported.
* `kafka_configs` - (Optional, Set) The configuration summary of Kafka is collected from logs. The values are described as follows:
  - `kafkaEndpoint`: the access address of the Kafka API.
  - `kafkaInstanceId`: the ID of the Kafka instance.
  - `kafkaConfigs`: the configuration summary of single or multiple logs. For more information, see the request parameter `kafkalogfilerconfig` in this article * *. See [`kafka_configs`](#kafka_configs) below.
* `liveness` - (Optional, Computed, Set) Container health check. Containers that fail the health check will be shut down and restored. Currently, only the method of issuing commands in the container is supported. See [`liveness`](#liveness) below.
* `max_surge_instance_ratio` - (Optional, Int, Available since v1.280.0) The proportion of maximum peak instances. The value description is as follows:
If the minimum number of surviving instances is 100 percent, the maximum peak cannot be set to 0. If it is set to **-1**, the system recommended value of 30% will be used, which is 30% of the number of existing instances. If there are currently 10 instances, then 10 x 30%= 3.
* `max_surge_instances` - (Optional, Int, Available since v1.280.0) Maximum number of peak instances. The value description is as follows:
If the minimum number of surviving instances is 100 percent, the maximum peak cannot be set to 0.
If **-1** is set, the maximum peak number of instances uses the system recommended value of 30%, which is 30% of the existing number of instances. If there are currently 10 instances, then 10 x 30%= 3.
* `memory` - (Optional, Int) The memory required for each instance, in MB, cannot be 0. One-to-one correspondence with CPU
* `micro_registration` - (Optional, Computed) Select the Nacos registry. The values are as follows:
  - `0`:SAE built-in Nacos.
  - `1`: User-created Nacos.
  - `2`:MSE commercial Nacos.
* `min_ready_instance_ratio` - (Optional, Computed, Int) The minimum number of surviving instances. The values are described as follows:
  - **-1**: Initialization value, indicating that the percentage is not used.
  - **0~100**: The unit is a percentage, rounded up. For example, set it to `50`%. If there are currently 5 instances, the minimum number of surviving instances is 3.

-> **NOTE:**  If `MinReadyInstance` and `MinReadyInstanceRatio` are passed at the same time, and the value of `MinReadyInstanceRatio` is not **-1**, the `MinReadyInstanceRatio` parameter shall prevail. Assume that `MinReadyInstances` is set to `5` and `MinReadyInstanceRatio` is set to `50`, then `50` is used to calculate the minimum number of surviving instances.

* `min_ready_instances` - (Optional, Computed, Int) The Minimum Available Instance. On the Change Had Promised during the Available Number of Instances to Be.
* `namespace_id` - (Optional, ForceNew) SAE namespace ID. Only namespaces whose names are lowercase letters and dashes (-) are supported, and must start with a letter. The namespace can be obtained by calling the DescribeNamespaceList interface.
* `nas_configs` - (Optional, List) Mount the NAS configuration. See [`nas_configs`](#nas_configs) below.
* `oidc_role_name` - (Optional, Available since v1.280.0) Configure the identity authentication service RAM role

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `oss_ak_id` - (Optional) OSS AccessKey ID
* `oss_ak_secret` - (Optional) OSS  AccessKey Secret
* `oss_mount_descs` - (Optional, Computed, List) OSS Mount description information. See [`oss_mount_descs`](#oss_mount_descs) below.
* `package_type` - (Required, ForceNew) Application package type. The values are described as follows:
  - When you choose to deploy in Java, `FatJar`, `War`, and `Image` are supported * *.
  - When you choose to deploy in PHP, the support types are as follows:
  - `PhpZip`
  - **IMAGE\_PHP\_5\_4**
  - **IMAGE\_PHP\_5\_4\_ALPINE**
  - **IMAGE\_PHP\_5\_5**
  - **IMAGE\_PHP\_5\_5\_ALPINE**
  - **IMAGE\_PHP\_5\_6**
  - **IMAGE\_PHP\_5\_6\_ALPINE**
  - **IMAGE\_PHP\_7\_0**
  - **IMAGE\_PHP\_7\_0\_ALPINE**
  - **IMAGE\_PHP\_7\_1**
  - **IMAGE\_PHP\_7\_1\_ALPINE**
  - **IMAGE\_PHP\_7\_2**
  - **IMAGE\_PHP\_7\_2\_ALPINE**
  - **IMAGE\_PHP\_7\_3**
  - **IMAGE\_PHP\_7\_3\_ALPINE**
  - When you choose to deploy in Python, you can support `PythonZip` and **Image * *.
* `package_url` - (Optional) Deployment package address. Only FatJar or War type applications can configure the deployment package address.
* `package_version` - (Optional, Computed) The version number of the deployment package. Required when the Package Type is War and FatJar.
* `php` - (Optional) The PHP version on which the PHP deployment package depends. Mirror not supported
* `php_arms_config_location` - (Optional, Computed) The PHP application monitors the Mount path, and you must ensure that the PHP server loads the configuration file of this path.
You don't need to pay attention to the configuration content, SAE will automatically render the correct configuration file.This field is obsolete, it is recommended to use the ebpf field.
* `php_config` - (Optional) PHP configuration file content.
* `php_config_location` - (Optional) Configure the Mount path for the PHP application startup. You must ensure that the PHP server is started using this configuration file.
* `post_start` - (Optional, Computed, Set) Execute the script after startup, the format is like: {"exec":{"command":["cat","/etc/group"]}} See [`post_start`](#post_start) below.
* `pre_stop` - (Optional, Computed, Set) Execute the script before stopping, the format is like: {"exec":{"command":["cat","/etc/group"]}} See [`pre_stop`](#pre_stop) below.
* `programming_language` - (Optional, ForceNew, Computed) Create the technology stack language of the application. The values are described as follows:
  - `java`:Java language.
  - `php`:PHP language.
  - `other`: Multilingual, such as Python, C ++, Go,. NET, and Node.js.
* `pvtz_discovery_svc` - (Optional, Set) Enable the K8s Service Service registration discovery. See [`pvtz_discovery_svc`](#pvtz_discovery_svc) below.
* `readiness` - (Optional, Computed, Set) Application startup status checks, containers that fail multiple health checks will be shut down and restarted. Containers that do not pass the health check will not receive SLB traffic. For example: {"exec":{"command":["sh","-c","cat /home/admin/start.sh"]},"initialDelaySeconds":30,"periodSeconds":30,"timeoutSeconds ":2} See [`readiness`](#readiness) below.
* `replicas` - (Required, Int) Initial number of instances
* `resource_type` - (Optional, Available since v1.280.0) Resource Type

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `sae_version` - (Optional, Available since v1.280.0) SAE version, supported versions are as follows: v1, v2

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `secret_mount_desc` - (Optional, Available since v1.280.0) Mount description

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `security_group_id` - (Optional, Computed) Security group ID
* `sls_configs` - (Optional, List) SLS  configuration. See [`sls_configs`](#sls_configs) below.
* `termination_grace_period_seconds` - (Optional, Computed, Int) Graceful offline timeout, the default is 30, the unit is seconds. The value range is 1~60.
* `timezone` - (Optional, Computed) Time zone, the default value is Asia/Shanghai.
* `tomcat_config` - (Optional, Computed, Set) Tomcat file configuration, set to "" or "{}" means to delete the configuration:  useDefaultConfig: Whether to use a custom configuration, if it is true, it means that the custom configuration is not used; if it is false, it means that the custom configuration is used. If you do not use custom configuration, the following parameter configuration will not take effect.  contextInputType: Select the access path of the application.  war: No need to fill in the custom path, the access path of the application is the WAR package name. root: No need to fill in the custom path, the access path of the application is /. custom: You need to fill in the custom path in the custom path below. contextPath: custom path, this parameter only needs to be configured when the contextInputType type is custom.  httpPort: The port range is 1024~65535. Ports less than 1024 need Root permission to operate. Because the container is configured with Admin permissions, please fill in a port greater than 1024. If not configured, the default is 8080. maxThreads: Configure the number of connections in the connection pool, the default size is 400. uriEncoding: Tomcat encoding format, including UTF-8, ISO-8859-1, GBK and GB2312. If not set, the default is ISO-8859-1. useBodyEncoding: Whether to use BodyEncoding for URL. See [`tomcat_config`](#tomcat_config) below.
* `update_strategy` - (Optional, Computed, Set) Deployment strategy. When the minimum number of surviving instances is equal to 1, the value of the `updategrategy` field is "". When the minimum number of surviving instances is greater than 1, an example is as follows:
  - 1 set of gray scale + subsequent 2 batches + automatic batch + batch interval 1 minute: '{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchWaitTime":1},"grayUpdate":{"gray":1}}'
  - Grayscale 1 set, followed by 2 batches, manual batch: '{"type":"GrayBatchUpdate","batchUpdate":{"batch":2,"reliasetype":"manual"},"grayUpdate":{"gray":1}}'
  - 2 batches + automatic batch + batch interval 0 minutes: '{"type":"BatchUpdate","batchUpdate":{"batch":2,"releaseType":"auto","batchwaitime":0}}'

The parameters are described as follows:
  - `type`: The type of the release policy. You can select **grayscale release** or batch release **BatchUpdate * *.
  - `batchUpdate`: Batch release policy.
  - `batch`: Release batch.
  - `Reliasetype`: The batch processing method. You can select `auto` or manual **manual * *.
  - `batchWaitTime`: The time interval between deployment in batches. Unit: seconds.
  - `grayUpdate`: the remaining batch after Grayscale. This parameter is required when `type` is **GrayBatchUpdate. See [`update_strategy`](#update_strategy) below.
* `vswitch_id` - (Optional) vSwitch ID
* `vpc_id` - (Optional, ForceNew) The VPC corresponding to the SAE namespace. In SAE, a namespace can only correspond to one VPC and cannot be modified. Creating a SAE application in the namespace for the first time will form a binding relationship. Multiple namespaces can correspond to a VPC. If you leave it blank, it will default to the VPC ID bound to the namespace.
* `war_start_options` - (Optional) WAR package launch application option. Application default startup command: java $JAVA_OPTS $CATALINA_OPTS [-Options] org.apache.catalina.startup.Bootstrap "$@" start
* `web_container` - (Optional) The version of tomcat that the deployment package depends on. Image type applications are not supported.

### `config_map_mount_desc`

The config_map_mount_desc supports the following:
* `config_map_id` - (Optional, Int, Available since v1.280.0) ConfigMap ID.
* `key` - (Optional, Available since v1.280.0) ConfigMap key-value pair.
* `mount_path` - (Optional, Available since v1.280.0) The container Mount path.

### `custom_host_alias`

The custom_host_alias supports the following:
* `host_name` - (Optional, Available since v1.280.0) hostname config.
* `ip` - (Optional, Available since v1.280.0) ip config.

### `envs`

The envs supports the following:
* `name` - (Optional, Available since v1.280.0) env config name.
* `value` - (Optional, Available since v1.280.0) env value.
* `value_from` - (Optional, Set, Available since v1.280.0) env ref config. See [`value_from`](#envs-value_from) below.

### `envs-value_from`

The envs-value_from supports the following:
* `config_map_ref` - (Optional, Set, Available since v1.280.0) env ref config item. See [`config_map_ref`](#envs-value_from-config_map_ref) below.

### `envs-value_from-config_map_ref`

The envs-value_from-config_map_ref supports the following:
* `config_map_id` - (Optional, Int, Available since v1.280.0) configmap id.
* `key` - (Optional, Available since v1.280.0) configmap key.if ref all key not need set this field.

### `kafka_configs`

The kafka_configs supports the following:
* `kafka_configs` - (Optional, List) kafka configs. See [`kafka_configs`](#kafka_configs-kafka_configs) below.
* `kafka_endpoint` - (Optional) kafka endpoints.
* `kafka_instance_id` - (Optional) kafka instance id.

### `kafka_configs-kafka_configs`

The kafka_configs-kafka_configs supports the following:
* `kafka_topic` - (Optional) kafka topic config.
* `log_dir` - (Optional) log dir.
* `log_type` - (Optional) log type.support stdout/file_log.

### `liveness`

The liveness supports the following:
* `exec` - (Optional, Set, Available since v1.280.0) exec config. See [`exec`](#liveness-exec) below.
* `failure_threshold` - (Optional, Int, Available since v1.280.0) Failur Threshold
* `http_get` - (Optional, Set, Available since v1.280.0) http check config. See [`http_get`](#liveness-http_get) below.
* `initial_delay_seconds` - (Optional, Int, Available since v1.280.0) liveness check init delay.
* `period_seconds` - (Optional, Int, Available since v1.280.0) liveness check period.
* `tcp_socket` - (Optional, Set, Available since v1.280.0) tcp check config. See [`tcp_socket`](#liveness-tcp_socket) below.
* `timeout_seconds` - (Optional, Int, Available since v1.280.0) liveness check timeout.

### `liveness-exec`

The liveness-exec supports the following:
* `command` - (Optional, List, Available since v1.280.0) commands.

### `liveness-http_get`

The liveness-http_get supports the following:
* `is_contain_key_word` - (Optional, Available since v1.280.0) http check contain keyword.
* `key_word` - (Optional, Available since v1.280.0) http check keyword.
* `path` - (Optional, Available since v1.280.0) http check path config.
* `port` - (Optional, Int, Available since v1.280.0) http check port config.
* `scheme` - (Optional, Available since v1.280.0) http check scheme config.support HTTP/HTTPS

### `liveness-tcp_socket`

The liveness-tcp_socket supports the following:
* `port` - (Optional, Int, Available since v1.280.0) tcp check port.

### `nas_configs`

The nas_configs supports the following:
* `mount_domain` - (Optional) nas domain.
* `mount_path` - (Optional) container mount path.
* `nas_id` - (Optional) nas instance id.
* `nas_path` - (Optional) nas path.
* `read_only` - (Optional) readonly.

### `oss_mount_descs`

The oss_mount_descs supports the following:
* `bucket_name` - (Optional, Available since v1.280.0) The name of the Bucket.
* `bucket_path` - (Optional, Available since v1.280.0) If the OSS Mount directory does not exist, an exception is triggered.
* `mount_path` - (Optional, Available since v1.280.0) Your container path in SAE. If the path already exists, it is an overlay relationship; If the path does not exist, a new one will be created.
* `read_only` - (Optional, Available since v1.280.0) Whether the container path has the read permission to mount Directory resources, the values are as follows:
  - `true`: Read-only permission.
  - `false`: read and write permissions.

### `post_start`

The post_start supports the following:
* `exec` - (Optional, Set, Available since v1.280.0) exec command config. See [`exec`](#post_start-exec) below.

### `post_start-exec`

The post_start-exec supports the following:
* `command` - (Optional, List, Available since v1.280.0) exec command array.

### `pre_stop`

The pre_stop supports the following:
* `exec` - (Optional, Set, Available since v1.280.0) exec comamnd config. See [`exec`](#pre_stop-exec) below.

### `pre_stop-exec`

The pre_stop-exec supports the following:
* `command` - (Optional, List, Available since v1.280.0) exec command array.

### `pvtz_discovery_svc`

The pvtz_discovery_svc supports the following:
* `enable` - (Optional) isenable.
* `namespace_id` - (Optional, ForceNew) namespaceid.
* `port_protocols` - (Optional, List) port and protocol config list. See [`port_protocols`](#pvtz_discovery_svc-port_protocols) below.
* `service_name` - (Optional, ForceNew) servicename.

### `pvtz_discovery_svc-port_protocols`

The pvtz_discovery_svc-port_protocols supports the following:
* `port` - (Optional, Int) port.
* `protocol` - (Optional) protocol.support TCP/UDP.

### `readiness`

The readiness supports the following:
* `exec` - (Optional, Set, Available since v1.280.0) exec command config. See [`exec`](#readiness-exec) below.
* `http_get` - (Optional, Set, Available since v1.280.0) http check config. See [`http_get`](#readiness-http_get) below.
* `initial_delay_seconds` - (Optional, Int, Available since v1.280.0) readiness check init delay config.
* `period_seconds` - (Optional, Int, Available since v1.280.0) readiness check period config.
* `tcp_socket` - (Optional, Set, Available since v1.280.0) tcp check config. See [`tcp_socket`](#readiness-tcp_socket) below.
* `timeout_seconds` - (Optional, Int, Available since v1.280.0) readiness timeout config.

### `readiness-exec`

The readiness-exec supports the following:
* `command` - (Optional, List, Available since v1.280.0) commands.

### `readiness-http_get`

The readiness-http_get supports the following:
* `is_contain_key_word` - (Optional, Available since v1.280.0) http check contain keyword.
* `key_word` - (Optional, Available since v1.280.0) http check keyword.
* `path` - (Optional, Available since v1.280.0) http check path.
* `port` - (Optional, Int, Available since v1.280.0) http check port.
* `scheme` - (Optional, Available since v1.280.0) http check schema.support HTTP/HTTPS.

### `readiness-tcp_socket`

The readiness-tcp_socket supports the following:
* `port` - (Optional, Int, Available since v1.280.0) tcp check port.

### `sls_configs`

The sls_configs supports the following:
* `log_dir` - (Optional, Available since v1.280.0) log dir.
* `log_type` - (Optional, Available since v1.280.0) sls logtype.support stdout/empty string.
* `logstore_name` - (Optional, Available since v1.280.0) sls storename config.
* `logtail_name` - (Optional, Available since v1.280.0) sls logtail name.
* `project_name` - (Optional, Available since v1.280.0) sls projectname.

### `tomcat_config`

The tomcat_config supports the following:
* `context_path` - (Optional, Available since v1.280.0) access url,default /.
* `max_threads` - (Optional, Int, Available since v1.280.0) maxthreads,default 400.
* `port` - (Optional, Int, Available since v1.280.0) port config,1024~65535.
* `uri_encoding` - (Optional, Available since v1.280.0) tomcat encoding config，includ UTF-8、ISO-8859-1、GBK、GB2312.default ISO-8859-1.
* `use_body_encoding_for_uri` - (Optional, Available since v1.280.0) BodyEncoding for URL default true.

### `update_strategy`

The update_strategy supports the following:
* `batch_update` - (Optional, Set, Available since v1.280.0) batch update config. See [`batch_update`](#update_strategy-batch_update) below.
* `type` - (Optional, Available since v1.280.0) deloy strategy type.

### `update_strategy-batch_update`

The update_strategy-batch_update supports the following:
* `batch` - (Optional, Int, Available since v1.280.0) total batch.
* `batch_wait_time` - (Optional, Int, Available since v1.280.0) instance ready interval time within a batch.
* `release_type` - (Optional, Available since v1.280.0) batch-to-batch processing.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Indicates That the Application of the Creation Time.
* `region_id` - The region ID of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 8 mins) Used when create the Application.
* `delete` - (Defaults to 8 mins) Used when delete the Application.
* `update` - (Defaults to 8 mins) Used when update the Application.

## Import

Serverless App Engine (SAE) Application can be imported using the id, e.g.

```shell
$ terraform import alicloud_sae_application.example <application_id>
```