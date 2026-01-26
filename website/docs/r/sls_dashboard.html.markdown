---
subcategory: "Log Service (SLS)"
layout: "alicloud"
page_title: "Alicloud: alicloud_sls_dashboard"
description: |-
  Provides a Alicloud Log Service (SLS) Dashboard resource.
---

# alicloud_sls_dashboard

Provides a Log Service (SLS) Dashboard resource.



For information about Log Service (SLS) Dashboard and how to use it, see [What is Dashboard](https://next.api.alibabacloud.com/document/Sls/2020-12-30/CreateDashboard).

-> **NOTE:** Available since v1.269.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform-example"
}

provider "alicloud" {
  region = ""
}

variable "project_name" {
  default = "example-projects213"
}

variable "dashboard_name" {
  default = "exampledashboard"
}

resource "alicloud_log_project" "defaultP71Nui" {
  description = "Description"
  name        = "example-projects213"
}


resource "alicloud_sls_dashboard" "default" {
        description = "description"
                                      charts           {
                    action  {
                                            }
                    display = {
           height = ""5""
           width = ""5""
           xPos = ""0""
           yPos = ""0""
           xAxis = "Array"
           yAxis = "Array"
    }
                title = "eqfqdasdwqd"
                type = "linepro"
                    search = {
           logstore = ""access-log""
           query = ""* | SELECT date_format(__time__ - __time__ % 60, \'%H:%i:%s\') as time, count(1) as count GROUP BY time ORDER BY time""
           start = ""-3600""
           end = ""now""
    }
        }
                      dashboard_name = "${var.dashboard_name}"
          project_name = "${var.project_name}"
          display_name = "asdsadsada"
  }
```

## Argument Reference

The following arguments are supported:
* `attribute` - (Optional, Map) The dashboard attribute value.
* `charts` - (Optional, List) The charts included. See [`charts`](#charts) below.
* `dashboard_name` - (Optional, ForceNew, Computed) The dashboard ID. Within the same project, the dashboard ID is unique and cannot be duplicated. Fuzzy search is supported. For example, entering "da" returns all dashboards whose IDs start with "da".
* `description` - (Optional) The dashboard description.
* `display_name` - (Optional) The display name of the dashboard.
* `project_name` - (Required, ForceNew) The project name.

### `charts`

The charts supports the following:
* `action` - (Optional, Map) The action.
* `display` - (Optional, Map) Display settings for the chart.  
* `search` - (Optional, Map) The query configuration.
* `title` - (Optional) Chart title. It can contain uppercase and lowercase English letters, digits, underscores (_), and hyphens (-). Hyphens and underscores cannot be used at the beginning or end of the name. The length must be between 2 and 64 characters, inclusive.  
* `type` - (Optional) The chart type.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. The value is formulated as `<project_name>:<dashboard_name>`.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Dashboard.
* `delete` - (Defaults to 5 mins) Used when delete the Dashboard.
* `update` - (Defaults to 5 mins) Used when update the Dashboard.

## Import

Log Service (SLS) Dashboard can be imported using the id, e.g.

```shell
$ terraform import alicloud_sls_dashboard.example <project_name>:<dashboard_name>
```