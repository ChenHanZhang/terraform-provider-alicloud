package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Eip SegmentAddress. >>> Resource test cases, automatically generated.
// Case 2198
func TestAccAlicloudEipSegmentAddress_basic2198(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eip_segment_address.default"
	ra := resourceAttrInit(resourceId, AlicloudEipSegmentAddressMap2198)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EipServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEipSegmentAddress")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sEipSegmentAddress%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEipSegmentAddressBasicDependence2198)
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
					"isp":      "BGP",
					"eip_mask": "28",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"isp":      "BGP",
						"eip_mask": "28",
					}),
				),
			}, {
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bandwidth", "eip_mask", "internet_charge_type", "isp", "netmode", "resource_group_id"},
			},
		},
	})
}

var AlicloudEipSegmentAddressMap2198 = map[string]string{}

func AlicloudEipSegmentAddressBasicDependence2198(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 2203
func TestAccAlicloudEipSegmentAddress_basic2203(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eip_segment_address.default"
	ra := resourceAttrInit(resourceId, AlicloudEipSegmentAddressMap2203)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EipServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEipSegmentAddress")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sEipSegmentAddress%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEipSegmentAddressBasicDependence2203)
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
					"eip_mask":          "28",
					"resource_group_id": "${var.ResourceGroupId}",
					"isp":               "${var.ISP}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"eip_mask":          "28",
						"resource_group_id": CHECKSET,
						"isp":               CHECKSET,
					}),
				),
			}, {
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bandwidth", "eip_mask", "internet_charge_type", "isp", "netmode", "resource_group_id"},
			},
		},
	})
}

var AlicloudEipSegmentAddressMap2203 = map[string]string{}

func AlicloudEipSegmentAddressBasicDependence2203(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "ResourceGroupId" {
  default = "rg-acfmv7pftjmydwa"
}

variable "ISP" {
  default = "BGP"
}

variable "PaymentType" {
  default = "PayByTraffic"
}

variable "RegionId" {
  default = "cn-shanghai"
}


`, name)
}

// Test Eip SegmentAddress. <<< Resource test cases, automatically generated.
