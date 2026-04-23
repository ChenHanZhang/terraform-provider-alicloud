---
subcategory: "APIG"
layout: "alicloud"
page_title: "Alicloud: alicloud_apig_http_api"
description: |-
  Provides a Alicloud APIG Http Api resource.
---

# alicloud_apig_http_api

Provides a APIG Http Api resource.

HTTP API endpoint  .

For information about APIG Http Api and how to use it, see [What is Http Api](https://next.api.aliyun.com/api/APIG/2024-03-27/CreateHttpApi).

-> **NOTE:** Available since v1.240.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}

variable "protocol" {
  default = "HTTP"
}

variable "protocol_https" {
  default = "HTTPS"
}

data "alicloud_resource_manager_resource_groups" "default" {}


resource "alicloud_apig_http_api" "default" {
  http_api_name = var.name
  protocols     = ["${var.protocol}"]
  base_path     = "/v1"
  description   = "zhiwei_pop_examplecase"
  type          = "Rest"
}
```

## Argument Reference

The following arguments are supported:
* `ai_protocols` - (Optional, List, Available since v1.277.0) AI protocols  
* `base_path` - (Optional) API base path, which must start with a forward slash (/).
* `description` - (Optional) API description.
* `environments` - (Optional, ForceNew, List, Available since v1.277.0) Published environment information. See [`environments`](#environments) below.
* `gateway_id` - (Optional, ForceNew, Available since v1.277.0) Cloud-native API gateway ID.
* `gateway_type` - (Optional, ForceNew, Available since v1.277.0) Gateway type filter. Currently, `AI` and `API` gateway types are supported.
* `http_api_name` - (Required, ForceNew) Perform an exact search by name.
* `model_category` - (Optional, Available since v1.277.0) AI model category  
* `protocols` - (Required, List) List of API access protocols.
* `resource_group_id` - (Optional, Computed) Target resource group ID.
* `type` - (Optional, ForceNew) The type of the HTTP API. Multiple types are supported and must be separated by commas (",").  
  - Http  
  - Rest  
  - LLM  
  - WebSocket  
  - HttpIngress  

### `environments`

The environments supports the following:

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `environments` - Published environment information.
  * `environment_id` - The environment ID.
  * `name` - The environment name.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Http Api.
* `delete` - (Defaults to 5 mins) Used when delete the Http Api.
* `update` - (Defaults to 5 mins) Used when update the Http Api.

## Import

APIG Http Api can be imported using the id, e.g.

```shell
$ terraform import alicloud_apig_http_api.example <http_api_id>
```