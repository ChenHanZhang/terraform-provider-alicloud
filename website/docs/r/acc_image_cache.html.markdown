---
subcategory: "ACC"
layout: "alicloud"
page_title: "Alicloud: alicloud_acc_image_cache"
description: |-
  Provides a Alicloud ACC Image Cache resource.
---

# alicloud_acc_image_cache

Provides a ACC Image Cache resource.



For information about ACC Image Cache and how to use it, see [What is Image Cache](https://next.api.alibabacloud.com/document/acc/2024-04-02/CreateImageCache).

-> **NOTE:** Available since v1.273.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `acr_registry_infos` - (Optional, List) Information about ACR Enterprise Edition instances. See [`acr_registry_infos`](#acr_registry_infos) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `image_cache_name` - (Required, ForceNew) Image cache name.
* `image_registry_credentials` - (Optional, List) Image registry credential information. See [`image_registry_credentials`](#image_registry_credentials) below.

-> **NOTE:** This parameter is immutable. Changing it after creation has no effect.

* `images` - (Required, ForceNew, List) The container image used to create the image cache. Currently, only one image is supported.
* `network_config` - (Required, ForceNew, Set) An operational attribute field that specifies network configuration. See [`network_config`](#network_config) below.
* `resource_group_id` - (Optional, Computed) The ID of the new resource group to which the resource will be moved.
* `tags` - (Optional, Map) Image cache tags. A maximum of 20 tags are allowed.

### `acr_registry_infos`

The acr_registry_infos supports the following:
* `instance_id` - (Optional) The ID of the ACR registry instance.
* `region_id` - (Optional) The region of the ACR registry.

### `image_registry_credentials`

The image_registry_credentials supports the following:
* `password` - (Optional) An operational attribute field representing the password.
* `server` - (Optional) Represents the operational attribute field of a server.
* `skip_cert_verification` - (Optional) Represents the operational attribute field indicating whether to skip certificate verification.
* `use_plain_http` - (Optional) Represents the operational attribute field indicating whether to use the HTTP protocol.
* `username` - (Optional) Represents the operational attribute field of a username.

### `network_config`

The network_config supports the following:
* `eip_instance` - (Optional, ForceNew, Set) Elastic IP (EIP) instance configuration. See [`eip_instance`](#network_config-eip_instance) below.
* `security_group_id` - (Required, ForceNew) The security group ID.
* `vswitch_ids` - (Required, ForceNew, List) List of vSwitch IDs. Up to 10 IDs can be specified, and one will be selected for pulling the image.

### `network_config-eip_instance`

The network_config-eip_instance supports the following:
* `auto_create` - (Optional, ForceNew) Whether to automatically create an EIP.
* `bandwidth` - (Optional, ForceNew, Int) Bandwidth for the automatically created EIP, in Mbps.
* `instance_id` - (Optional, ForceNew) The Elastic IP address. If this parameter is specified, it takes precedence, and all other parameters become invalid.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the image cache.
* `region_id` - Region ID.
* `status` - The status of the image cache.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Image Cache.
* `delete` - (Defaults to 5 mins) Used when delete the Image Cache.
* `update` - (Defaults to 5 mins) Used when update the Image Cache.

## Import

ACC Image Cache can be imported using the id, e.g.

```shell
$ terraform import alicloud_acc_image_cache.example <image_cache_id>
```