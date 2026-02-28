---
subcategory: "Function Compute Service V3 (FCV3)"
layout: "alicloud"
page_title: "Alicloud: alicloud_fcv3_service"
description: |-
  Provides a Function Compute Service V3 (FCV3) Service resource.
---

# alicloud_fcv3_service

Provides a Function Compute Service V3 (FCV3) Service resource.

Function Compute Service V3 (FCV3) is the latest version of Alibaba Cloud's serverless computing service, offering enhanced performance, improved networking capabilities, and advanced security features compared to the previous FC service. FCV3 provides better scalability, faster cold start times, and more flexible configuration options.

Function Compute Service (FCV3) is a fully managed service that allows you to run code without managing servers. It scales automatically and executes code only when needed.

For information about Function Compute Service V3 (FCV3) Service and how to use it, see [What is Service](https://www.alibabacloud.com/help/en/functioncompute/developer-reference/api-fc-2023-03-30-createservice).

-> **NOTE:** Available since v1.228.0.

-> **NOTE:** **FCV3 is the recommended version for all new projects. It offers improved performance, enhanced security, and additional features compared to the legacy FC service.**

## Example Usage

Basic Usage

<div style="display: block;margin-bottom: 40px;"><div class="oics-button" style="float: right;position: absolute;margin-bottom: 10px;">
  <a href="https://api.aliyun.com/terraform?resource=alicloud_fcv3_service&exampleId=example-id&activeTab=example&spm=docs.r.fcv3_service.0.example&intl_lang=EN_US" target="_blank">
    <img alt="Open in AliCloud" src="https://img.alicdn.com/imgextra/i1/O1CN01hjjqXv1uYUlY56FyX_!!6000000006049-55-tps-254-36.svg" style="max-height: 44px; max-width: 100%;">
  </a>
</div></div>

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "random_uuid" "default" {
}

resource "alicloud_log_project" "default" {
  project_name = "${var.name}-${random_uuid.default.result}"
}

resource "alicloud_log_store" "default" {
  project_name  = alicloud_log_project.default.project_name
  logstore_name = "example-logstore"
}

resource "alicloud_ram_role" "default" {
  name = "${var.name}-role-${random_uuid.default.result}"

  assume_role_policy_document = <<EOF
{
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "fc.aliyuncs.com"
        ]
      }
    }
  ],
  "Version": "1"
}
EOF

  description = "Role for FCV3 service"
}

resource "alicloud_ram_role_policy_attachment" "default" {
  role_name   = alicloud_ram_role.default.name
  policy_name = "AliyunLogFullAccess"
  policy_type = "System"
}

resource "alicloud_fcv3_service" "default" {
  service_name    = "${var.name}-service-${random_uuid.default.result}"
  description     = "Example FCV3 service"
  role            = alicloud_ram_role.default.arn
  internet_access = true

  log_config {
    project                 = alicloud_log_project.default.project_name
    logstore                = alicloud_log_store.default.logstore_name
    enable_request_metrics  = true
    enable_instance_metrics = true
  }

  tags = {
    "Environment" = "test"
    "ManagedBy"   = "Terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `service_name` - (Required, ForceNew) The name of the FCV3 service. It must be globally unique and consist of lowercase letters, digits, and hyphens. It must start with a letter and end with a letter or digit. Length is 1-128 characters.
* `description` - (Optional) The description of the FCV3 service.
* `role` - (Optional) The ARN of the RAM role that Function Compute assumes when executing functions in this service.
* `internet_access` - (Optional) Whether functions in this service can access the internet. Default to `true`.
* `vpc_config` - (Optional) VPC configuration for the service. See [`vpc_config`](#vpc_config) below.
* `nas_config` - (Optional) NAS configuration for the service. See [`nas_config`](#nas_config) below.
* `log_config` - (Optional) Log configuration for the service. See [`log_config`](#log_config) below.
* `tags` - (Optional) A mapping of tags to assign to the resource.

### `vpc_config`

The vpc_config supports the following:

* `vpc_id` - (Optional) The ID of the VPC.
* `vswitch_ids` - (Optional) A list of vSwitch IDs.
* `security_group_id` - (Optional) The ID of the security group.

### `nas_config`

The nas_config supports the following:

* `user_id` - (Optional) The user ID for NAS access.
* `group_id` - (Optional) The group ID for NAS access.
* `mount_points` - (Optional) A list of mount points. See [`mount_points`](#nas_config-mount_points) below.

### `nas_config-mount_points`

The nas_config-mount_points supports the following:

* `server_addr` - (Optional) The server address of the NAS.
* `mount_dir` - (Optional) The mount directory.

### `log_config`

The log_config supports the following:

* `project` - (Optional) The name of the SLS project.
* `logstore` - (Optional) The name of the SLS logstore.
* `log_begin_rule` - (Optional) The log begin rule.
* `enable_request_metrics` - (Optional) Whether to enable request metrics.
* `enable_instance_metrics` - (Optional) Whether to enable instance metrics.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the resource (same as service_name).
* `created_time` - The creation time of the service.
* `last_modified_time` - The last modification time of the service.
* `service_id` - The system-generated ID of the service.
* `service_arn` - The ARN of the service.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 5 mins) Used when creating the FCV3 Service.
* `delete` - (Defaults to 5 mins) Used when deleting the FCV3 Service.
* `update` - (Defaults to 5 mins) Used when updating the FCV3 Service.

## Import

FCV3 Service can be imported using the service name, e.g.

```shell
$ terraform import alicloud_fcv3_service.example my-service-name
```