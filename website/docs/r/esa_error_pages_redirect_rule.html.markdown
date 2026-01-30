---
subcategory: "ESA"
layout: "alicloud"
page_title: "Alicloud: alicloud_esa_error_pages_redirect_rule"
description: |-
  Provides a Alicloud ESA Error Pages Redirect Rule resource.
---

# alicloud_esa_error_pages_redirect_rule

Provides a ESA Error Pages Redirect Rule resource.



For information about ESA Error Pages Redirect Rule and how to use it, see [What is Error Pages Redirect Rule](https://next.api.alibabacloud.com/document/ESA/2024-09-10/CreateErrorPagesRedirectRule).

-> **NOTE:** Available since v1.270.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_example_ErrorPagesRedirectRule" {
  type         = "NS"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "overseas"
  auto_pay     = true
  plan_name    = "high"
}

resource "alicloud_esa_site" "resource_Site_example_ErrorPagesRedirectRule" {
  site_name   = "gositecdn.cn"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_example_ErrorPagesRedirectRule.id
  coverage    = "overseas"
  access_type = "NS"
}


resource "alicloud_esa_error_pages_redirect_rule" "default" {
  site_id      = alicloud_esa_site.resource_Site_example_ErrorPagesRedirectRule.id
  rule_enable  = "off"
  rule         = "(http.host eq \"video.example.com\")"
  sequence     = "1"
  site_version = "0"
  rule_name    = "rule_example"
  error_pages_redirect {
    target_url  = "https://example.com/foo/bar"
    status_code = "500"
  }
  error_pages_redirect {
    target_url  = "https://example.com/foo"
    status_code = "400"
  }
  error_pages_redirect {
    target_url  = "https://example.com/example"
    status_code = "503"
  }
}
```

## Argument Reference

The following arguments are supported:
* `error_pages_redirect` - (Required, List) The configurations of error pages redirect.  See [`error_pages_redirect`](#error_pages_redirect) below.
* `rule` - (Optional) Rule content, using conditional expressions to match user requests. When adding global configuration, this parameter does not need to be set. There are two usage scenarios:
  - Match all incoming requests: value set to true
  - Match specified request: Set the value to a custom expression, for example: (http.host eq "video.example.com")
* `rule_enable` - (Optional) Rule switch. When adding global configuration, this parameter does not need to be set. Value range:
  - `on`: open.
  - `off`: close.
* `rule_name` - (Optional) Rule name. When adding global configuration, this parameter does not need to be set.
* `sequence` - (Optional, ForceNew, Int) Order of rule execution. The smaller the value, the higher the priority for execution.
* `site_id` - (Required, ForceNew, Int) The website ID, which can be obtained by calling the [ListSites](https://www.alibabacloud.com/help/en/doc-detail/2850189.html) operation.
* `site_version` - (Optional, ForceNew, Int) The version number of the site configuration. For sites that have enabled configuration version management, this parameter can be used to specify the effective version of the configuration site, which defaults to version 0.

### `error_pages_redirect`

The error_pages_redirect supports the following:
* `status_code` - (Required) The response code that you want to use to indicate URL redirection. Valid values:
  - 400
  - 403
  - 404
  - 405
  - 414
  - 416
  - 500
  - 501
  - 502
  - 503
  - 504
* `target_url` - (Required) The destination URL after the redirect.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<site_id>:<config_id>`.
* `config_id` - The configuration ID.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Error Pages Redirect Rule.
* `delete` - (Defaults to 5 mins) Used when delete the Error Pages Redirect Rule.
* `update` - (Defaults to 5 mins) Used when update the Error Pages Redirect Rule.

## Import

ESA Error Pages Redirect Rule can be imported using the id, e.g.

```shell
$ terraform import alicloud_esa_error_pages_redirect_rule.example <site_id>:<config_id>
```