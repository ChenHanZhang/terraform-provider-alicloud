// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test NATGateway FullNat. >>> Resource test cases, automatically generated.
// Case 全生命周期_FullNat_1.0.0 12370
func TestAccAliCloudNATGatewayFullNat_basic12370(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_n_a_t_gateway_full_nat.default"
	ra := resourceAttrInit(resourceId, AlicloudNATGatewayFullNatMap12370)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NATGatewayServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNATGatewayFullNat")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnatgateway%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNATGatewayFullNatBasicDependence12370)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNATGatewayFullNatMap12370 = map[string]string{
	"status":            CHECKSET,
	"create_time":       CHECKSET,
	"full_nat_entry_id": CHECKSET,
}

func AlicloudNATGatewayFullNatBasicDependence12370(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "VPC" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "tf-test-fullnat-vpc"
}

resource "alicloud_vswitch" "VSwitch" {
  vpc_id       = alicloud_vpc.VPC.id
  zone_id      = "eu-central-1b"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "tf-test-fullnat-vsw"
}


`, name)
}

// Test NATGateway FullNat. <<< Resource test cases, automatically generated.
