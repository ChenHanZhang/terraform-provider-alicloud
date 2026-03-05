// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Eflo Vcc. >>> Resource test cases, automatically generated.
// Case vcc-tf-cross-account_副本1770022786191 12518
func TestAccAliCloudEfloVcc_basic12518(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_vcc.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloVccMap12518)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloVcc")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloVccBasicDependence12518)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-wulanchabu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"description":          "test-tf-cross-account",
					"access_could_service": "false",
					"connection_type":      "CENTR",
					"zone_id":              "${var.zone_id}",
					"cen_owner_id":         "${var.cen_owner_id}",
					"bandwidth":            "1000",
					"vcc_name":             name,
					"vswitch_id":           "${var.vsw_id}",
					"vpc_id":               "${var.vpc_id}",
					"cen_id":               "${var.cen_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "test-tf-cross-account",
						"access_could_service": "false",
						"connection_type":      "CENTR",
						"zone_id":              CHECKSET,
						"cen_owner_id":         CHECKSET,
						"bandwidth":            "1000",
						"vcc_name":             name,
						"vswitch_id":           CHECKSET,
						"vpc_id":               CHECKSET,
						"cen_id":               CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "Test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "Test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "Test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF-update",
						"tags.For":     "Test-update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "0",
						"tags.Created": REMOVEKEY,
						"tags.For":     REMOVEKEY,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"access_could_service", "bgp_asn", "description", "zone_id"},
			},
		},
	})
}

var AlicloudEfloVccMap12518 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudEfloVccBasicDependence12518(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cen_owner_id" {
  default = <<EOF
1013666993027780
EOF
}

variable "vsw_id" {
  default = "vsw-0jlpbevcjvsbe1b29cwkb"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "cen_id" {
  default = "cen-czo1beum1uku1vq4gp"
}

variable "vpc_id" {
  default = "vpc-0jll0blr5o98680qxqujn"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}


`, name)
}

// Test Eflo Vcc. <<< Resource test cases, automatically generated.
