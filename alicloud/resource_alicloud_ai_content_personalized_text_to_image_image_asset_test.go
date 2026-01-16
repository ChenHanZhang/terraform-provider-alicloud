// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test AiContent PersonalizedTextToImageImageAsset. >>> Resource test cases, automatically generated.
// Case PersonalizedTextToImageImageAssetTest 7645
func TestAccAliCloudAiContentPersonalizedTextToImageImageAsset_basic7645(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ai_content_personalized_text_to_image_image_asset.default"
	ra := resourceAttrInit(resourceId, AlicloudAiContentPersonalizedTextToImageImageAssetMap7645)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AiContentServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAiContentPersonalizedTextToImageImageAsset")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccaicontent%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAiContentPersonalizedTextToImageImageAssetBasicDependence7645)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"train_steps": "10",
					"seed":        "1",
					"image_url": []string{
						"https://repo.singsound.com/repository/singsound_artifacts/700-1280.jpg"},
					"prompt":       "girl",
					"image_number": "1",
					"strength":     "0.3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"train_steps":  "10",
						"seed":         "1",
						"image_url.#":  "1",
						"prompt":       "girl",
						"image_number": "1",
						"strength":     CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"image_number", "image_url", "prompt", "seed", "strength", "train_steps"},
			},
		},
	})
}

var AlicloudAiContentPersonalizedTextToImageImageAssetMap7645 = map[string]string{
	"status": CHECKSET,
}

func AlicloudAiContentPersonalizedTextToImageImageAssetBasicDependence7645(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test AiContent PersonalizedTextToImageImageAsset. <<< Resource test cases, automatically generated.
