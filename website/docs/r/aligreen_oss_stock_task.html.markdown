---
subcategory: "Aligreen"
layout: "alicloud"
page_title: "Alicloud: alicloud_aligreen_oss_stock_task"
description: |-
  Provides a Alicloud Aligreen Oss Stock Task resource.
---

# alicloud_aligreen_oss_stock_task

Provides a Aligreen Oss Stock Task resource.

Scan task for existing files in OSS.

For information about Aligreen Oss Stock Task and how to use it, see [What is Oss Stock Task](https://next.api.alibabacloud.com/document/Green/2017-08-23/CreateOssStockTask).

-> **NOTE:** Available since v1.228.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "terraform"
}

provider "alicloud" {
  region = "cn-shanghai"
}

resource "random_integer" "default" {
  min = 10000
  max = 99999
}

resource "alicloud_oss_bucket" "defaultPyhXOV" {
  storage_class = "Standard"
  bucket        = "${var.name}-${random_integer.default.result}"
}

resource "alicloud_aligreen_callback" "defaultJnW8Na" {
  callback_url         = "https://www.aliyun.com/"
  crypt_type           = "0"
  callback_name        = "${var.name}${random_integer.default.result}"
  callback_types       = ["machineScan"]
  callback_suggestions = ["block"]
}


resource "alicloud_aligreen_oss_stock_task" "default" {
  image_opened                       = true
  auto_freeze_type                   = "acl"
  audio_max_size                     = "200"
  image_scan_limit                   = "1"
  video_frame_interval               = "1"
  video_scan_limit                   = "1000"
  audio_scan_limit                   = "1000"
  video_max_frames                   = "200"
  video_max_size                     = "500"
  start_date                         = "2024-08-01 00:00:00 +0800"
  end_date                           = "2024-12-31 09:06:42 +0800"
  buckets                            = jsonencode([{ "Bucket" : "${alicloud_oss_bucket.defaultPyhXOV.bucket}", "Selected" : true, "Prefixes" : [] }])
  image_scenes                       = ["porn"]
  audio_antispam_freeze_config       = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  image_live_freeze_config           = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  video_terrorism_freeze_config      = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  image_terrorism_freeze_config      = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  callback_id                        = alicloud_aligreen_callback.defaultJnW8Na.id
  image_ad_freeze_config             = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  biz_type                           = "recommend_massmedia_template_01"
  audio_scenes                       = jsonencode(["antispam"])
  image_porn_freeze_config           = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  video_live_freeze_config           = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  video_porn_freeze_config           = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  video_voice_antispam_freeze_config = jsonencode({ "Type" : "suggestion", "Value" : "block" })
  video_scenes                       = jsonencode(["ad", "terrorism", "live", "porn", "antispam"])
  video_ad_freeze_config             = jsonencode({ "Type" : "suggestion", "Value" : "block" })
}
```

### Deleting `alicloud_aligreen_oss_stock_task` or removing it from your configuration

Terraform cannot destroy resource `alicloud_aligreen_oss_stock_task`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `audio_antispam_freeze_config` - (Optional, ForceNew) The automatic freezing configuration for audio anti-spam scenarios.
* `audio_auto_freeze_opened` - (Optional, ForceNew) Audio files must be automatically frozen.
* `audio_max_size` - (Optional, ForceNew, Int) The maximum size of a single audio file.
* `audio_opened` - (Optional, ForceNew) Specifies whether to scan audio files. Valid values: - `true`: enabled - `false`: disabled.
* `audio_scan_limit` - (Optional, ForceNew, Int) The upper limit for audio scanning.
* `audio_scenes` - (Optional, ForceNew) The list of audio scenarios.
* `auto_freeze_type` - (Optional, ForceNew) The freezing method.
* `biz_type` - (Optional, ForceNew) The business scenario.
* `buckets` - (Optional, ForceNew) Buckets.
* `callback_id` - (Optional, ForceNew, Int) The callback task ID.
* `end_date` - (Optional, ForceNew) The end time. The format is YYYY-MM-DD HH:mm:ss (default time zone: UTC+08:00).
* `image_ad_freeze_config` - (Optional, ForceNew) The scenario for automatically freezing images that contain ads.
* `image_auto_freeze_opened` - (Optional, ForceNew) Images need to be automatically frozen.
* `image_live_freeze_config` - (Optional, ForceNew) The automatic freezing configuration for undesirable image scenarios.
* `image_opened` - (Optional, ForceNew) Specifies whether to scan images. Valid values: - true: enabled - false: disabled.
* `image_porn_freeze_config` - (Optional, ForceNew) The scenario for automatically freezing pornographic images.
* `image_scan_limit` - (Optional, ForceNew, Int) The maximum number of images to scan.
* `image_scenes` - (Optional, ForceNew, List) The list of image scenarios.
* `image_terrorism_freeze_config` - (Optional, ForceNew) The automatic image freezing configuration for terrorism scenarios.
* `scan_image_no_file_type` - (Optional, ForceNew) Scan files without extensions.
* `start_date` - (Optional, ForceNew) The start time. Format: YYYY-MM-DD HH:mm:ss (default time zone: UTC+08:00).
* `video_ad_freeze_config` - (Optional, ForceNew) Automatically freeze videos that contain advertisements.
* `video_auto_freeze_opened` - (Optional, ForceNew) Videos must be automatically frozen.
* `video_frame_interval` - (Optional, ForceNew, Int) The frame capture frequency.
* `video_live_freeze_config` - (Optional, ForceNew) The scenario for automatically freezing videos with inappropriate content.
* `video_max_frames` - (Optional, ForceNew, Int) The maximum number of frames captured per video.
* `video_max_size` - (Optional, ForceNew, Int) The maximum size of a single video file.
* `video_opened` - (Optional, ForceNew) Specifies whether to scan videos. Valid values: `true`: enabled. `false`: disabled.
* `video_porn_freeze_config` - (Optional, ForceNew) Automatically freeze videos that contain pornographic content.
* `video_scan_limit` - (Optional, ForceNew, Int) The upper limit for video scanning.
* `video_scenes` - (Optional, ForceNew) The list of video scenarios.
* `video_terrorism_freeze_config` - (Optional, ForceNew) The automatic video freezing configuration for terrorism scenarios.
* `video_voice_antispam_freeze_config` - (Optional, ForceNew) Automatically freeze videos that contain audio violations.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Oss Stock Task.

## Import

Aligreen Oss Stock Task can be imported using the id, e.g.

```shell
$ terraform import alicloud_aligreen_oss_stock_task.example <oss_stock_task_id>
```