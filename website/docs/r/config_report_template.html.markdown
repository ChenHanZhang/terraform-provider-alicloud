---
subcategory: "Cloud Config (Config)"
layout: "alicloud"
page_title: "Alicloud: alicloud_config_report_template"
description: |-
  Provides a Alicloud Cloud Config (Config) Report Template resource.
---

# alicloud_config_report_template

Provides a Cloud Config (Config) Report Template resource.

Config Compliance Report Tempalte.

For information about Cloud Config (Config) Report Template and how to use it, see [What is Report Template](https://next.api.alibabacloud.com/document/Config/2020-09-07/CreateReportTemplate).

-> **NOTE:** Available since v1.266.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = "cn-hangzhou"
}


resource "alicloud_config_report_template" "default" {
  report_granularity = "AllInOne"
  report_scope {
    key        = "RuleId"
    value      = "cr-xxx"
    match_type = "In"
  }
  report_file_formats         = "excel"
  report_template_name        = "example-name"
  report_template_description = "example-desc"
  subscription_frequency      = " "
  report_language             = "en-US"
}
```

## Argument Reference

The following arguments are supported:
* `report_file_formats` - (Optional) Report file format. Currently, only Excel is supported.
* `report_granularity` - (Optional) Report aggregation granularity.

Supported from the management account perspective:
  - AllInOne (aggregates all accounts within the template scope into a single report)
  - GroupByAggregator (aggregates reports by account group, generating a compressed file)
  - GroupByAccount (generates a separate report for each account (default), packaged into a compressed file)

Member accounts support only GroupByAccount.
* `report_language` - (Optional) Report language. Supported values are zh-CN and en-US. If left empty, it defaults to en-US.
* `report_scope` - (Optional, ForceNew, List) An array of report scopes used to specify the set of rules included in the audit report. The relationship between each ReportScope entry in the array is logical OR (additive).

-> **NOTE:**  For example, if the array contains two entries, where the first ReportScope specifies rule In cr-1 and the second specifies rule In cr-2, then the resulting report includes rules cr-1 and cr-2.
 See [`report_scope`](#report_scope) below.
* `report_template_description` - (Optional) Report template description.
* `report_template_name` - (Required) Report template name.
* `subscription_frequency` - (Optional) Report subscription frequency. If this field is not empty, it must be a Quartz-formatted Cron expression that triggers subscription notifications.

The format is: second minute hour day-of-month month day-of-week. Below are common Cron expression examples:
  - Execute at 00:00 every day: 0 0 0 * * ?
  - Execute at 15:30 every Monday: 0 30 15 ? * MON
  - Execute at 02:00 on the 1st of every month: 0 0 2 1 * ?

Where:
  - "*" indicates any value.
  - "?" is used in the day-of-month and day-of-week fields to indicate no specific value.
  - "MON" stands for Monday.

-> **NOTE:**  The trigger time uses UTC+8. You can adjust the Cron expression according to your time zone.

-> **NOTE:**  Execution will occur as closely as possible to the scheduled time defined by the Cron expression, but may be delayed due to report generation status. The Cron expression ensures that each template triggers at most one notification per day.

-> **NOTE:**  In addition to using "MON" for Monday, you can also use numbers. Note that in the Quartz framework:
1 represents Sunday; 2 represents Monday; 3 represents Tuesday; 4 represents Wednesday; 5 represents Thursday; 6 represents Friday; 7 represents Saturday.

### `report_scope`

The report_scope supports the following:
* `key` - (Optional) The key of the report scope. Currently supported keys include:
  - AggregatorId
  - CompliancePackId
  - RuleId.
* `match_type` - (Optional) Matching logic. Currently, only "In" is supported.
* `value` - (Optional) The value of the report scope. Multiple rule IDs can be specified, separated by English commas (,).

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Report Template.
* `delete` - (Defaults to 5 mins) Used when delete the Report Template.
* `update` - (Defaults to 5 mins) Used when update the Report Template.

## Import

Cloud Config (Config) Report Template can be imported using the id, e.g.

```shell
$ terraform import alicloud_config_report_template.example <report_template_id>
```