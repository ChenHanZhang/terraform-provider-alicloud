package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudFcv3Service_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_fcv3_service.default"
	ra := resourceAttrInit(resourceId, AlicloudFcv3ServiceMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &Fcv3ServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeFcv3Service")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"service_name":    "tf-testacc-fcv3service",
					"description":     "test description",
					"role":            "${alicloud_ram_role.default.arn}",
					"internet_access": true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(nil),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

var AlicloudFcv3ServiceMap = map[string]string{
	"service_name":       CHECKSET,
	"created_time":       CHECKSET,
	"last_modified_time": CHECKSET,
	"service_id":         CHECKSET,
	"service_arn":        CHECKSET,
}