---
subcategory: "Ai Content"
layout: "alicloud"
page_title: "Alicloud: alicloud_ai_content_personalized_text_to_image_image_asset"
description: |-
  Provides a Alicloud Ai Content Personalized Text To Image Image Asset resource.
---

# alicloud_ai_content_personalized_text_to_image_image_asset

Provides a Ai Content Personalized Text To Image Image Asset resource.

Personalized Wensheng Figure v2 Version Picture Resources.

For information about Ai Content Personalized Text To Image Image Asset and how to use it, see [What is Personalized Text To Image Image Asset](https://next.api.alibabacloud.com/document/AiContent/20240611/PersonalizedTextToImageAddInferenceJob).

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


resource "alicloud_ai_content_personalized_text_to_image_image_asset" "default" {
  train_steps  = "10"
  seed         = "1"
  image_url    = ["https://repo.singsound.com/repository/singsound_artifacts/700-1280.jpg"]
  prompt       = "girl"
  image_number = "1"
  strength     = 0.3
}
```

### Deleting `alicloud_ai_content_personalized_text_to_image_image_asset` or removing it from your configuration

Terraform cannot destroy resource `alicloud_ai_content_personalized_text_to_image_image_asset`. Terraform will remove this resource from the state file, however resources may remain.

## Argument Reference

The following arguments are supported:
* `image_number` - (Optional, Int) The number of pictures to generate. Note: Due to the resource limitation of the test environment, a maximum of 10 pictures are generated at a time. If the value is greater than 10, it will be processed by 10.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `image_url` - (Required, List) A list containing links to images, which can contain links to single or multiple images, such as ["url_1", "url_2", ...]

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `prompt` - (Required) English prompt describes the generated image, and the character to be generated is replaced by , for example, "a man in the snow" should be changed to "a  in the snow", "a photo of a girl" should be changed to "a photo of a "

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `seed` - (Optional, Int) Random number seed, when the seed is fixed, the same generated image can be obtained to ensure the reproducibility of the result. Note: the input value must be in [-1, 2147483647], and a suitable seed will be automatically generated when the input value is outside the interval or not entered.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `strength` - (Optional, Float) Represents the strength of the input reference image's influence on the generation process,
The value range is [0.3, 0.4, 0.5, 0.6, 0.7, 0.8],
A smaller value indicates that the influence intensity of the reference image on the generation is lower, and correspondingly, the influence intensity of the text on the generation process is higher.
The default is 0.5, generally do not need to set, keep the default.

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.

* `train_steps` - (Optional, Int) Model Training Task Training Steps

-> **NOTE:** The parameter is immutable after resource creation. It only applies during resource creation and has no effect when modified post-creation.


## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `status` - The status of the resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Personalized Text To Image Image Asset.

## Import

Ai Content Personalized Text To Image Image Asset can be imported using the id, e.g.

```shell
$ terraform import alicloud_ai_content_personalized_text_to_image_image_asset.example <inference_job_id>
```