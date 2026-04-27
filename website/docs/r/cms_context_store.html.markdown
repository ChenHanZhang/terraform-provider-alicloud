---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_context_store"
description: |-
  Provides a Alicloud Cms Context Store resource.
---

# alicloud_cms_context_store

Provides a Cms Context Store resource.

Context Store  .

For information about Cms Context Store and how to use it, see [What is Context Store](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateContextStore).

-> **NOTE:** Available since v1.277.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `config` - (Optional, Set) Configuration. See [`config`](#config) below.
* `context_store_name` - (Required, ForceNew) A resource property field representing the resource name.  
* `context_type` - (Required) Context type.
* `description` - (Optional) Description.

### `config`

The config supports the following:
* `metadata_field` - (Optional, ForceNew, Map) Metadata information
* `source` - (Optional, ForceNew, Set) Data source See [`source`](#config-source) below.

### `config-source`

The config-source supports the following:
* `logstore` - (Optional, ForceNew) SLS LogStore name
* `project` - (Optional, ForceNew) SLS Project name
* `start_time` - (Optional, ForceNew) Start time

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<workspace>:<context_store_name>`.
* `create_time` - Creation time.
* `region_id` - Region ID.
* `status` - Status.
* `workspace` - Workspace.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Context Store.
* `delete` - (Defaults to 5 mins) Used when delete the Context Store.
* `update` - (Defaults to 5 mins) Used when update the Context Store.

## Import

Cms Context Store can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_context_store.example <workspace>:<context_store_name>
```