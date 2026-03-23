---
subcategory: "Cms"
layout: "alicloud"
page_title: "Alicloud: alicloud_cms_alert_webhook"
description: |-
  Provides a Alicloud Cms Alert Webhook resource.
---

# alicloud_cms_alert_webhook

Provides a Cms Alert Webhook resource.

Notification target: Webhook  .

For information about Cms Alert Webhook and how to use it, see [What is Alert Webhook](https://next.api.alibabacloud.com/document/Cms/2024-03-30/CreateAlertWebhook).

-> **NOTE:** Available since v1.274.0.

## Example Usage

Basic Usage

没有资源测试用例，请先通过资源测试用例后再生成示例代码。

## Argument Reference

The following arguments are supported:
* `alert_webhook_id` - (Optional, ForceNew, Computed) The unique identifier of the webhook.  
* `alert_webhook_name` - (Required) The name of the webhook.  
* `content_type` - (Optional) The data content type. Supported types include:  
  - JSON (default)  
  - FORM.  
* `headers` - (Optional, Map) Headers.  
* `lang` - (Optional) Language. Supported languages include:  
  - zh_CN  
  - en_US.  
* `method` - (Optional) The HTTP request method. Supported methods include:  
  - GET  
  - POST.  
* `url` - (Required) The URL address for alert callbacks.  

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Alert Webhook.
* `delete` - (Defaults to 5 mins) Used when delete the Alert Webhook.
* `update` - (Defaults to 5 mins) Used when update the Alert Webhook.

## Import

Cms Alert Webhook can be imported using the id, e.g.

```shell
$ terraform import alicloud_cms_alert_webhook.example <alert_webhook_id>
```