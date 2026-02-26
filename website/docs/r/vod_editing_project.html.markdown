---
subcategory: "ApsaraVideo VoD (VOD)"
layout: "alicloud"
page_title: "Alicloud: alicloud_vod_editing_project"
description: |-
  Provides a Alicloud VOD Editing Project resource.
---

# alicloud_vod_editing_project

Provides a VOD Editing Project resource.



For information about VOD Editing Project and how to use it, see [What is Editing Project](https://www.alibabacloud.com/help/en/apsaravideo-for-vod/latest/addeditingproject#doc-api-vod-AddEditingProject).

-> **NOTE:** Available since v1.187.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tfexample"
}
data "alicloud_regions" "default" {
  current = true
}
resource "alicloud_vod_editing_project" "example" {
  editing_project_name = var.name
  title                = var.name
  timeline             = <<EOF
  {
    "VideoTracks":[
      {
        "VideoTrackClips":[
          {
          "MediaId":"0c60e6f02dae71edbfaa472190a90102",
          "In":2811
          }
        ]
      }
    ]
  }
  EOF
  cover_url            = "https://demo.aliyundoc.com/6AB4D0E1E1C74468883516C2349D1FC2-6-2.png"
  division             = data.alicloud_regions.default.regions.0.id
}
```

## Argument Reference

The following arguments are supported:
* `cover_url` - (Optional) Cover url
* `editing_project_name` - (Optional, Computed) The name of the resource
* `timeline` - (Optional, Computed) Timeline
* `title` - (Required) Title

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.
* `status` - Status.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Editing Project.
* `delete` - (Defaults to 5 mins) Used when delete the Editing Project.
* `update` - (Defaults to 5 mins) Used when update the Editing Project.

## Import

VOD Editing Project can be imported using the id, e.g.

```shell
$ terraform import alicloud_vod_editing_project.example <editing_project_id>
```