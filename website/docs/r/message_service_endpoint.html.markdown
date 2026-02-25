---
subcategory: "Message Service"
layout: "alicloud"
page_title: "Alicloud: alicloud_message_service_endpoint"
description: |-
  Provides a Alicloud Message Service Endpoint resource.
---

# alicloud_message_service_endpoint

Provides a Message Service Endpoint resource.

-> **NOTE**: Destruction of this resource will disable the endpont, and this will cause all requests originating from the corresponding endpoint to be intercepted and return an error.

For information about Message Service Endpoint and how to use it, see [What is Endpoint](https://www.alibabacloud.com/help/en/mns/developer-reference/api-mns-open-2022-01-19-enableendpoint).

-> **NOTE:** Available since v1.243.0.

## Example Usage

Basic Usage

```terraform
provider "alicloud" {
  region = "cn-hangzhou"
}

resource "alicloud_message_service_endpoint" "default" {
  endpoint_enabled = true
  endpoint_type    = "public"
}
```

## Argument Reference

The following arguments are supported:
* `cidr_list` - (Optional, List, Available since v1.272.0) The CIDR blocks. See [`cidr_list`](#cidr_list) below.
* `endpoint_type` - (Required, ForceNew) Access point type. Value:
  - public: indicates a public access point. (Currently only public is supported)

### `cidr_list`

The cidr_list supports the following:
* `acl_strategy` - (Required, Available since v1.272.0) The ACL policy. Valid value:
  - allow: indicates that the current endpoint allows access from the corresponding CIDR block. (Only allow is supported)
* `cidr` - (Required, Available since v1.272.0) The CIDR block.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Endpoint.
* `delete` - (Defaults to 5 mins) Used when delete the Endpoint.
* `update` - (Defaults to 5 mins) Used when update the Endpoint.

## Import

Message Service Endpoint can be imported using the id, e.g.

```shell
$ terraform import alicloud_message_service_endpoint.example <endpoint_type>
```