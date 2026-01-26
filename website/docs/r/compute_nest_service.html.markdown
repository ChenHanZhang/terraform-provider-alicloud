---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_service"
description: |-
  Provides a Alicloud Compute Nest Service resource.
---

# alicloud_compute_nest_service

Provides a Compute Nest Service resource.

Compute nest service.

For information about Compute Nest Service and how to use it, see [What is Service](https://next.api.alibabacloud.com/document/ComputeNestSupplier/2021-05-21/CreateService).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_compute_nest_service" "default" {
  deploy_type  = "operation"
  duration     = "2592000"
  policy_names = "AliyunComputeNestPolicyForReadOnly"
  service_info {
    locale            = "zh-CN"
    short_description = "自动化exampleservice创建服务"
    image             = "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png"
    name              = "自动化exampleservice创建服务"
  }
  deploy_metadata {
    supplier_deploy_metadata {
      deploy_timeout = "7200"
    }
  }
  service_type       = "operation"
  approval_type      = "Manual"
  version_name       = "自动化exampleservice创建服务"
  operation_metadata = "{\"PrometheusConfigMap\":{}}"
  share_type         = "Public"
}
```

## Argument Reference

The following arguments are supported:
* `alarm_metadata` - (Optional) AlarmMetadata
* `approval_type` - (Optional) Service usage requisition approval type.

Value:
  - Manual: Approval received.
  - AutoPass: automatically passed.
* `deploy_metadata` - (Optional, Set) DeployMetadata See [`deploy_metadata`](#deploy_metadata) below.
* `deploy_type` - (Required, ForceNew) DeployType
* `duration` - (Optional, Int) Generation operation and maintenance time. Unit: seconds.
* `is_support_operated` - (Optional) Whether to operate on behalf.

Value:
  - true: On.
  - false: closed.
* `license_metadata` - (Optional) LicenseMetadata元数据
* `operation_metadata` - (Optional) OperationMetadata
 
* `policy_names` - (Optional) The policy name.
The length of a policy name cannot exceed 128 characters. Separate multiple names with commas (,). For the time being, only parameter policies related to Agent operation and maintenance are supported.
* `resource_group_id` - (Optional, Computed) ResourceGroupId
* `service_id` - (Optional, ForceNew, Computed) The service ID.
You can call [ListServices-query service information](~~ 2264368 ~~) to obtain the service ID.
* `service_info` - (Optional, ForceNew, List) Service information. See [`service_info`](#service_info) below.
* `service_type` - (Required, ForceNew) ServiceType
* `share_type` - (Optional) The sharing type.

Value:
  - Public: Public, formal deployment and trial deployment are not restricted.
  - Restricted: Restricted, both formal deployment and trial deployment are Restricted.
  - OnlyFormalRestricted: limited deployment only.
  - OnlyTrailRestricted: limited deployment for trial only.
  - Hidden: Hidden, invisible, and cannot apply for deployment permissions.
* `tenant_type` - (Optional) TenantType
* `trial_duration` - (Optional, Int) Trial duration. Unit: days.
* `upgrade_metadata` - (Optional) UpgradeMetadata
* `version` - (Optional, ForceNew) service version
* `version_name` - (Optional) The version name.

### `deploy_metadata`

The deploy_metadata supports the following:
* `network_metadata` - (Optional, Set) NetworkMetadata See [`network_metadata`](#deploy_metadata-network_metadata) below.
* `supplier_deploy_metadata` - (Optional, Set) 商家部署元数据 See [`supplier_deploy_metadata`](#deploy_metadata-supplier_deploy_metadata) below.
* `template_configs` - (Optional, List) TemplateConfigs See [`template_configs`](#deploy_metadata-template_configs) below.

### `deploy_metadata-network_metadata`

The deploy_metadata-network_metadata supports the following:
* `enable_private_vpc_connection` - (Optional) EnablePrivateVpcConnection

### `deploy_metadata-supplier_deploy_metadata`

The deploy_metadata-supplier_deploy_metadata supports the following:
* `deploy_timeout` - (Optional, Int) 部署超时时间

### `deploy_metadata-template_configs`

The deploy_metadata-template_configs supports the following:
* `allowed_regions` - (Optional, List) AllowedRegions
* `hidden_parameter_keys` - (Optional, List) HiddenParameterKeys
* `name` - (Optional) Template Name
* `predefined_parameters` - (Optional, List) PredefinedParameters
* `update_info` - (Optional, Set) UpdateInfo See [`update_info`](#deploy_metadata-template_configs-update_info) below.
* `url` - (Optional) TemplateUrl

### `deploy_metadata-template_configs-update_info`

The deploy_metadata-template_configs-update_info supports the following:
* `parameters_allowed_to_be_modified` - (Optional, List) ParametersAllowedToBeModified
* `parameters_cause_interruption_if_modified` - (Optional, List) parameterCauseInterruptionIfModified
* `parameters_conditionally_cause_interruption_if_modified` - (Optional, List) ParametersConditionallyCauseInterruptionIfModified
* `parameters_not_allowed_to_be_modified` - (Optional, List) ParametersNotAllowedToBeModified
* `parameters_uncertainly_allowed_to_be_modified` - (Optional, List) parametersUncertainlyAllowedToBeModified
* `parameters_uncertainly_cause_interruption_if_modified` - (Optional, List) ParametersUncertainlyCauseInterruptionIfModified

### `service_info`

The service_info supports the following:
* `agreements` - (Optional, ForceNew, List) The protocol document information for the service. See [`agreements`](#service_info-agreements) below.
* `image` - (Optional) The service icon address.
* `locale` - (Required) The service configuration language.

Value:
  - zh-CN: Chinese.
  - en-US: English.
* `long_description_url` - (Optional) The address of the detailed description of the service.
* `name` - (Required) The service name.
* `short_description` - (Optional) Overview of services.
* `softwares` - (Optional, ForceNew, List) Software information used in the service. See [`softwares`](#service_info-softwares) below.

### `service_info-agreements`

The service_info-agreements supports the following:
* `name` - (Optional) The agreement name.
* `url` - (Optional) Agreement link.

### `service_info-softwares`

The service_info-softwares supports the following:
* `name` - (Optional) The software name.
* `version` - (Optional) Software version.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Service.
* `delete` - (Defaults to 5 mins) Used when delete the Service.
* `update` - (Defaults to 5 mins) Used when update the Service.

## Import

Compute Nest Service can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_service.example <service_id>
```