---
subcategory: "Threat Detection"
layout: "alicloud"
page_title: "Alicloud: alicloud_threat_detection_web_lock_config"
description: |-
  Provides a Alicloud Threat Detection Web Lock Config resource.
---

# alicloud_threat_detection_web_lock_config

Provides a Threat Detection Web Lock Config resource.

List information of servers for Web tamper protection.

For information about Threat Detection Web Lock Config and how to use it, see [What is Web Lock Config](https://www.alibabacloud.com/help/en/security-center/developer-reference/api-sas-2018-12-03-modifyweblockstart).

-> **NOTE:** Available since v1.195.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_threat_detection_assets" "default" {
  machine_types = "ecs"
}
resource "alicloud_threat_detection_web_lock_config" "default" {
  inclusive_file_type = "php;jsp;asp;aspx;js;cgi;html;htm;xml;shtml;shtm;jpg"
  uuid                = data.alicloud_threat_detection_assets.default.ids.0
  mode                = "whitelist"
  local_backup_dir    = "/usr/local/aegis/bak"
  dir                 = "/tmp/"
  defence_mode        = "audit"
}
```

## Argument Reference

The following arguments are supported:
* `defence_mode` - (Required) The prevention mode. Valid values:

  - `block`: Interception Mode
  - `audit`: Alert Mode
* `dir` - (Required) The directory for which you want to enable web tamper proofing.
* `exclusive_dir` - (Optional) The directory for which you want to disable web tamper proofing.

-> **NOTE:**  If you set `Mode` to `blacklist`, you must specify this parameter.

* `exclusive_file` - (Optional) The file for which you want to disable web tamper proofing.

-> **NOTE:**  If you set `Mode` to `blacklist`, you must specify this parameter.

* `exclusive_file_type` - (Optional) The type of file for which you want to disable web tamper proofing. Separate multiple types with semicolons (;). Valid values:

*   php
*   jsp
*   asp
*   aspx
*   js
*   cgi
*   html
*   htm
*   xml
*   shtml
*   shtm
*   jpg
*   gif
*   png

-> **NOTE:**  If you set `Mode` to `blacklist`, you must specify this parameter.

* `inclusive_file` - (Optional, Available since v1.272.0) The file that has web tamper proofing enabled.

-> **NOTE:**  If the value of `Mode` is `whitelist`, this parameter is returned.

* `inclusive_file_type` - (Optional) The type of file for which you want to enable web tamper proofing. Separate multiple types with semicolons (;). Valid values:

*   php
*   jsp
*   asp
*   aspx
*   js
*   cgi
*   html
*   htm
*   xml
*   shtml
*   shtm
*   jpg
*   gif
*   png

-> **NOTE:**  If you set `Mode` to `whitelist`, you must specify this parameter.

* `lang` - (Optional, Available since v1.272.0) The language of the content within the request and response. Valid values:

  - `zh`: Chinese
  - `en`: English

-> **NOTE:** This parameter is only evaluated during resource creation and update. Modifying it in isolation will not trigger any action.

* `local_backup_dir` - (Required) The local path to the backup files of the protected directory.\
The directory format of a Linux server is different from that of a Windows server. You must enter the directory in the required format based on your operating system. Examples:

  - Linux server: /usr/local/aegis/bak
  - Windows server: C:\Program Files (x86)\Alibaba\Aegis\bak
* `mode` - (Optional) The protection mode of web tamper proofing. Valid values:

  - `whitelist`: In this mode, web tamper proofing is enabled for the specified directories and file types.
  - `blacklist`: In this mode, web tamper proofing is enabled for the unspecified sub-directories, file types, and files in the protected directories.
* `uuid` - (Required, ForceNew) The UUID of the server for which you want to add a directory to protect.

-> **NOTE:**  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<uuid>:<config_id>`.
* `config_id` - The configuration ID of the protected directory.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Web Lock Config.
* `delete` - (Defaults to 5 mins) Used when delete the Web Lock Config.
* `update` - (Defaults to 5 mins) Used when update the Web Lock Config.

## Import

Threat Detection Web Lock Config can be imported using the id, e.g.

```shell
$ terraform import alicloud_threat_detection_web_lock_config.example <uuid>:<config_id>
```