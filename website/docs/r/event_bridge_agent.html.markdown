---
subcategory: "Event Bridge"
layout: "alicloud"
page_title: "Alicloud: alicloud_event_bridge_agent"
description: |-
  Provides a Alicloud Event Bridge Agent resource.
---

# alicloud_event_bridge_agent

Provides a Event Bridge Agent resource.

User-defined Agent  .

For information about Event Bridge Agent and how to use it, see [What is Agent](https://next.api.alibabacloud.com/document/eventbridge/2020-04-01/CreateAgent).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

variable "agent_name" {
  default = "AgentForPopTest"
}


resource "alicloud_event_bridge_agent" "default" {
  description = "You are a helpful assistant"
  prompt      = "You are a helpful assistant"
  name        = var.agent_name
}
```

## Argument Reference

The following arguments are supported:
* `description` - (Optional) Agent description  
* `metadata` - (Optional, ForceNew, Set) Metadata such as associated resources   See [`metadata`](#metadata) below.
* `name` - (Optional, ForceNew, Computed) Agent name  
* `prompt` - (Optional) Prompt  

### `metadata`

The metadata supports the following:
* `attachments` - (Optional, List) List of associated resources   See [`attachments`](#metadata-attachments) below.

### `metadata-attachments`

The metadata-attachments supports the following:
* `arn` - (Optional) Resource ARN  
* `mime_type` - (Optional) Resource type  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Agent.
* `delete` - (Defaults to 5 mins) Used when delete the Agent.
* `update` - (Defaults to 5 mins) Used when update the Agent.

## Import

Event Bridge Agent can be imported using the id, e.g.

```shell
$ terraform import alicloud_event_bridge_agent.example <name>
```