---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_endpoint_connector"
description: |-
  Provides a Alicloud Cms Endpoint Connector resource.
---

# alicloud_cms_endpoint_connector

Provides a Cms Endpoint Connector resource.



For information about Cms Endpoint Connector and how to use it, see [What is Endpoint Connector](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateEndpointConnector).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `alias` - (Optional) Configuration alias (used for display).
* `credential` - (Required, Map) Credential configuration (cohesive object)
* `description` - (Optional) Configuration description.
* `endpoint` - (Required) Service endpoint address.
* `headers` - (Optional, List) Custom HTTP request headers. See [`headers`](#headers) below.
* `name` - (Required) Configuration name (unique within workspace and type).
* `properties` - (Optional, Map) Type-specific configuration (required keys validated based on type).
* `type` - (Required, ForceNew) Configuration type: model_service or agent_app.
* `workspace` - (Optional) Workspace.

### `headers`

The headers supports the following:
* `key` - (Optional) HTTP header key
* `value` - (Optional) HTTP header value

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `region_id` - Resource property field representing the region ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Endpoint Connector.
* `delete` - (Defaults to 5 mins) Used when delete the Endpoint Connector.
* `update` - (Defaults to 5 mins) Used when update the Endpoint Connector.

## Import

Cms Endpoint Connector can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_endpoint_connector.example <connector_id>
```