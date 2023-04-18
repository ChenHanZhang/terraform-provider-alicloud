package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Vpc VswitchCidrReservation. >>> Resource test cases, automatically generated.
// Case 1984
func TestAccAlicloudVpcVswitchCidrReservation_basic1984(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_vpc_vswitch_cidr_reservation.default"
	ra := resourceAttrInit(resourceId, AlicloudVpcVswitchCidrReservationMap1984)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &VpcServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeVpcVswitchCidrReservation")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sVpcVswitchCidrReservation%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudVpcVswitchCidrReservationBasicDependence1984)
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
					"resource_group_id":             "rg-aek2xl5qajpkquq",
					"ip_version":                    "IPv4",
					"vswitch_id":                    "vsw-bp19icd33jy4ji9rrnhoc",
					"vswitch_cidr_reservation_name": "rdk-test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id":             "rg-aek2xl5qajpkquq",
						"ip_version":                    "IPv4",
						"vswitch_id":                    "vsw-bp19icd33jy4ji9rrnhoc",
						"vswitch_cidr_reservation_name": "rdk-test",
					}),
				),
			}, {
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cidr_reservation_mask", "resource_type"},
			},
		},
	})
}

var AlicloudVpcVswitchCidrReservationMap1984 = map[string]string{}

func AlicloudVpcVswitchCidrReservationBasicDependence1984(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Vpc VswitchCidrReservation. <<< Resource test cases, automatically generated.
