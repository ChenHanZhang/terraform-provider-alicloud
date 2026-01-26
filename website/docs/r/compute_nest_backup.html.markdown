---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_backup"
description: |-
  Provides a Alicloud Compute Nest Backup resource.
---

# alicloud_compute_nest_backup

Provides a Compute Nest Backup resource.

Service Instance Backup.

For information about Compute Nest Backup and how to use it, see [What is Backup](https://next.api.alibabacloud.com/document/ComputeNest/2021-06-01/CreateBackup).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_compute_nest_service" "defaultelCW8F" {
  deploy_type = "ros"
  service_info {
    locale            = "zh-CN"
    short_description = "自动化example备份service创建服务"
    image             = "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png"
    name              = "自动化example备份service创建服务-573"
  }
  policy_names = "AliyunComputeNestPolicyForFullAccess"
  deploy_metadata {
  }
  service_type       = "private"
  approval_type      = "Manual"
  version_name       = "自动化exampleservice创建服务"
  operation_metadata = "{\"PrometheusConfigMap\":{\"模板1\":{\"EnablePrometheus\":false}},\"ModifyParametersConfig\":[], \"SupportBackup\":true,\"StatusOperationConfigs\":[{\"TemplateName\":\"模板1\",\"SupportOperations\":[\"Start\",\"Stop\"]}]}"
}

resource "alicloud_compute_nest_service_instance" "defaultXRlJaY" {
  parameters = "{\"RegionId\":\"cn-hangzhou\",\"PayType\":\"PostPaid\",\"ZoneId\":\"cn-hangzhou-i\",\"VpcCidrBlock\":\"192.168.0.0/16\",\"VSwitchCidrBlock\":\"192.168.11.0/24\",\"EcsInstanceType\":\"ecs.g6.large\",\"SystemDiskCategory\":\"cloud_essd\",\"SystemDiskSize\":40,\"InstancePassword\":\"liuzheng121@\",\"InstanceCount\":1,\"UserEnablePrometheus\":false}"
  service {
    service_id = alicloud_compute_nest_service.defaultelCW8F.id
  }
  template_name = "模板1"
}


resource "alicloud_compute_nest_backup" "default" {
  description         = "example"
  service_instance_id = alicloud_compute_nest_service_instance.defaultXRlJaY.id
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional, ForceNew) Description
* `service_instance_id` - (Required, ForceNew) Service instance id

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `status` - Backup Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 10 mins) Used when create the Backup.
* `delete` - (Defaults to 8 mins) Used when delete the Backup.

## Import

Compute Nest Backup can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_backup.example <backup_id>
```