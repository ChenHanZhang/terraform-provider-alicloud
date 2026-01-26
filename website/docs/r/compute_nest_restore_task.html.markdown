---
subcategory: "Compute Nest"
layout: "alicloud"
page_title: "Alicloud: alicloud_compute_nest_restore_task"
description: |-
  Provides a Alicloud Compute Nest Restore Task resource.
---

# alicloud_compute_nest_restore_task

Provides a Compute Nest Restore Task resource.

Backup and restore tasks for computing nest instances.

For information about Compute Nest Restore Task and how to use it, see [What is Restore Task](https://next.api.alibabacloud.com/document/ComputeNest/2021-06-01/CreateRestoreTask).

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

resource "alicloud_compute_nest_service" "defaultpV4cW6" {
  deploy_type = "ros"
  service_info {
    locale            = "zh-CN"
    short_description = "自动化example备份service创建服务"
    image             = "https://service-info-public.oss-cn-hangzhou.aliyuncs.com/1563457855438522/service-image/05498c54-80ea-445f-9a44-1e2ba25b7462.png"
    name              = "自动化example备份service创建服务-229"
  }
  policy_names = "AliyunComputeNestPolicyForFullAccess"
  deploy_metadata {
  }
  service_type       = "private"
  approval_type      = "Manual"
  version_name       = "自动化exampleservice创建服务"
  operation_metadata = "{\"PrometheusConfigMap\":{\"模板1\":{\"EnablePrometheus\":false}},\"ModifyParametersConfig\":[], \"SupportBackup\":true,\"StatusOperationConfigs\":[{\"TemplateName\":\"模板1\",\"SupportOperations\":[\"Start\",\"Stop\"]}]}"
}

resource "alicloud_compute_nest_service_instance" "defaultslP8EP" {
  parameters = "{\"RegionId\":\"cn-hangzhou\",\"PayType\":\"PostPaid\",\"ZoneId\":\"cn-hangzhou-i\",\"VpcCidrBlock\":\"192.168.0.0/16\",\"VSwitchCidrBlock\":\"192.168.11.0/24\",\"EcsInstanceType\":\"ecs.g6.large\",\"SystemDiskCategory\":\"cloud_essd\",\"SystemDiskSize\":40,\"InstancePassword\":\"liuzheng121@\",\"InstanceCount\":1,\"UserEnablePrometheus\":false}"
  service {
    service_id = alicloud_compute_nest_service.defaultpV4cW6.id
  }
  template_name = "模板1"
}

resource "alicloud_compute_nest_backup" "defaulthJ0HNs" {
  service_instance_id = alicloud_compute_nest_service_instance.defaultslP8EP.id
  description         = "fdsfdsfsd_326"
}


resource "alicloud_compute_nest_restore_task" "default" {
  service_instance_id = alicloud_compute_nest_service_instance.defaultslP8EP.id
  backup_id           = alicloud_compute_nest_backup.defaulthJ0HNs.id
}
```

### Deleting `alicloud_compute_nest_restore_task` or removing it from your configuration

Terraform cannot destroy resource `alicloud_compute_nest_restore_task`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `backup_id` - (Required, ForceNew) The backup ID of the data source to be restored.
* `service_instance_id` - (Required, ForceNew) ID of the backup instance.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `status` - The status of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Restore Task.

## Import

Compute Nest Restore Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_compute_nest_restore_task.example <restore_task_id>
```