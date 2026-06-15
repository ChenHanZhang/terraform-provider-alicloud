---
subcategory: "Serverless Workflow"
layout: "alicloud"
page_title: "Alicloud: alicloud_fnf_flow"
description: |-
  Provides a Alicloud Serverless Workflow Flow resource.
---

# alicloud_fnf_flow

Provides a Serverless Workflow Flow resource.

The flow defines the business logic description and the general information required for executing the process. For example, an order management process may include order creation, payment processing, inventory reservation, and user notification.

For information about Serverless Workflow Flow and how to use it, see [What is Flow](https://www.alibabacloud.com/help/en/doc-detail/123079.htm).

-> **NOTE:** Available since v1.105.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-shanghai"
}

resource "alicloud_ram_role" "default" {
  name     = "tf-example-fnfflow"
  document = <<EOF
  {
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
          "Service": [
            "fnf.aliyuncs.com"
          ]
        }
      }
    ],
    "Version": "1"
  }
  EOF
}

resource "alicloud_fnf_flow" "example" {
  definition  = <<EOF
  version: v1beta1
  type: flow
  steps:
    - type: pass
      name: helloworld
  EOF
  role_arn    = alicloud_ram_role.default.arn
  description = "Test for terraform fnf_flow."
  name        = "tf-example-flow"
  type        = "FDL"
}
```

## Argument Reference

The following arguments are supported:
* `definition` - (Required) The definition of the flow. It must comply with the Flow Definition Language (FDL) syntax.
* `description` - (Required) The description of the flow.
* `environment` - (Optional, Set, Available since v1.282.0) Environmental information that can be accessed during flow execution See [`environment`](#environment) below.
* `flow_name` - (Required, ForceNew, Available since v1.282.0) The name of the flow. The name must be unique within a region for the same Alibaba Cloud account.
* `resource_group_id` - (Optional, Computed, Available since v1.282.0) The ID of the resource group
* `role_arn` - (Optional, Computed) Tracing ability configuration
* `type` - (Required) The type of the flow. Valid values are FDL or DEFAULT.

### `environment`

The environment supports the following:
* `variables` - (Optional, List, Available since v1.282.0) List of environment variables that can be accessed during flow execution See [`variables`](#environment-variables) below.

### `environment-variables`

The environment-variables supports the following:
* `description` - (Optional, Available since v1.282.0) Environment Variable Description
* `name` - (Optional, Available since v1.282.0) Environment Variable Name
* `value` - (Optional, Available since v1.282.0) Environment variable value

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The time when the flow was created.
* `flow_id` - The unique ID of the flow.
* `last_modified_time` - The time when the flow was last modified.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Flow.
* `delete` - (Defaults to 5 mins) Used when delete the Flow.
* `update` - (Defaults to 5 mins) Used when update the Flow.

## Import

Serverless Workflow Flow can be imported using the id, e.g.

```shell
$ terraform import alicloud_fnf_flow.example <flow_name>
```