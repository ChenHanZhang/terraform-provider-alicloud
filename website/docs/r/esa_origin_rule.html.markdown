---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_origin_rule"
description: |-
  Provides a Alicloud ESA Origin Rule resource.
---

# alicloud_esa_origin_rule

Provides a ESA Origin Rule resource.

Site origin rules.

For information about ESA Origin Rule and how to use it, see [What is Origin Rule](https://next.api.alibabacloud.com/document/ESA/2024-09-10/CreateOriginRule).

-> **NOTE:** Available since v1.244.0.

## Example Usage

Basic Usage

```terraform
data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_origin_rule" "default" {
  origin_sni        = "origin.example.com"
  site_id           = data.alicloud_esa_sites.default.sites.0.id
  origin_host       = "origin.example.com"
  dns_record        = "tf.example.com"
  site_version      = "0"
  rule_name         = "tf"
  origin_https_port = "443"
  origin_scheme     = "http"
  range             = "on"
  origin_http_port  = "8080"
  rule              = "(http.host eq \"video.example.com\")"
  rule_enable       = "on"
}
```

## Argument Reference

The following arguments are supported:
* `dns_record` - (Optional) Rewrite the DNS resolution record of the origin request.
* `follow302_enable` - (Optional, Available since v1.262.1) 302 redirect following toggle for origin requests. Valid values:
  - `on`: Enable.
  - `off`: Disable.
* `follow302_max_tries` - (Optional, Available since v1.262.1) Maximum number of 302 redirects to follow. Valid values: [1–5].
* `follow302_retain_args` - (Optional, Available since v1.262.1) Switch for retaining original request parameters. Valid values:
  - `on`: Enable.
  - `off`: Disable.
* `follow302_retain_header` - (Optional, Available since v1.262.1) Switch for retaining original request headers. Valid values:
  - `on`: Enable.
  - `off`: Disable.
* `follow302_target_host` - (Optional, Available since v1.262.1) Modify the origin host after a 302 redirect.
* `origin_host` - (Optional) The HOST header carried in the origin request.
* `origin_http_port` - (Optional) Origin server port used when fetching content via the HTTP protocol.
* `origin_https_port` - (Optional) The origin server port accessed when using the HTTPS protocol for origin fetching.
* `origin_mtls` - (Optional, Available since v1.262.1) mTLS toggle. Valid values:
  - `on`: Enable.
  - `off`: Disable.
* `origin_read_timeout` - (Optional, Available since v1.262.1) Origin read timeout (seconds).
* `origin_scheme` - (Optional) The protocol used for origin requests. Valid values:
  - `http`: Use the HTTP protocol for origin fetching.
  - `https`: Use the HTTPS protocol for origin fetching.
  - `follow`: Follow the client's protocol for origin fetching.
* `origin_sni` - (Optional) SNI included in the origin request.
* `origin_verify` - (Optional, Available since v1.262.1) Origin certificate verification switch. Valid values:  
  - `on`: enabled.  
  - `off`: disabled.
* `range` - (Optional) Use the range-based chunked method to download files from the origin. Valid values:
  - `on`: Enable.
  - `off`: Disable.
  - `force`: Force enable.
* `range_chunk_size` - (Optional, Available since v1.262.1) Range chunk size.
* `rule` - (Optional) Rule content, which uses conditional expressions to match user requests. This parameter is not required when adding global configurations. There are two usage scenarios:  
  - Match all incoming requests: set the value to true.  
  - Match specific requests: set the value to a custom expression, for example: (http.host eq "video.example.com")
* `rule_enable` - (Optional) The rule switch. This parameter is not required when adding a global configuration. Valid values:
  - `on`: Enable.
  - `off`: Disable.
* `rule_name` - (Optional) The rule name. This parameter is not required when adding a global configuration.
* `sequence` - (Optional, Computed, Int, Available since v1.262.1) The execution order of the rule. Rules with smaller values take higher priority. This parameter is used only when setting or modifying the order of a single rule configuration.
* `site_id` - (Required, ForceNew) Site ID.
* `site_version` - (Optional, ForceNew, Int) Site configuration version number. For sites with configuration version management enabled, you can use this parameter to specify the site version to which the configuration applies. The default value is version 0.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<site_id>:<config_id>`.
* `config_id` - Origin rule configuration ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Origin Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Origin Rule.
* `update` - (Defaults to 5 mins) Used when update the Origin Rule.

## Import

ESA Origin Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_origin_rule.example <site_id>:<config_id>
```