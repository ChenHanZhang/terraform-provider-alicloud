---
subcategory: "unreleased"
layout: "alicloud"
page_title: "Alicloud: alicloud_event_bridge_api_destination"
description: |-
  Provides a Alicloud Event Bridge Api Destination resource.
---

# alicloud_event_bridge_api_destination

Provides a Event Bridge Api Destination resource.



For information about Event Bridge Api Destination and how to use it, see [What is Api Destination](https://next.api.alibabacloud.com/document/eventbridge/2020-04-01/CreateApiDestination).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = var.region
}

variable "region" {
  default = "cn-chengdu"
}

variable "name" {
  default = "terraform-example"
}

resource "alicloud_event_bridge_connection" "default" {
  connection_name = var.name
  network_parameters {
    network_type = "PublicNetwork"
  }
}

resource "alicloud_event_bridge_api_destination" "default" {
  connection_name      = alicloud_event_bridge_connection.default.connection_name
  api_destination_name = var.name
  description          = "test-api-destination-connection"
  http_api_parameters {
    endpoint = "http://127.0.0.1:8001"
    method   = "POST"
  }
}
```

## Argument Reference

The following arguments are supported:
* `api_destination_name` - (Required, ForceNew) The name of the api. The maximum length is 127 characters. Minimum length 2 characters.
* `connection_name` - (Required) The connection configuration name. The maximum length is 127 characters. Minimum length 2 characters.
  -

Tip:
You must first call the Connection interface to create a Connection configuration. Enter the name of the existing Connection.
* `description` - (Optional) The description of the api. It must be no more than 255 characters.
* `http_api_parameters` - (Required, Set) Configuration information for API endpoints See [`http_api_parameters`](#http_api_parameters) below.

### `http_api_parameters`

The http_api_parameters supports the following:
* `endpoint` - (Required) The access point address of the API endpoint. The maximum length is 127 characters.
* `method` - (Required) HTTP request method.
  - GET
  - POST
  - HEAD
  - DELETE
  - PUT
  - PATCH

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - Creation time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Api Destination.
* `delete` - (Defaults to 5 mins) Used when delete the Api Destination.
* `update` - (Defaults to 5 mins) Used when update the Api Destination.

## Import

Event Bridge Api Destination can be imported using the id, e.g.

```shell
$ terraform import alicloud_event_bridge_api_destination.example <api_destination_name>
```