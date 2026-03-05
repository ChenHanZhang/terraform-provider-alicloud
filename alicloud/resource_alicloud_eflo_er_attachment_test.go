// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Eflo ErAttachment. >>> Resource test cases, automatically generated.
// Case er_at_tf_vcc 11361
func TestAccAliCloudEfloErAttachment_basic11361(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_er_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloErAttachmentMap11361)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloErAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloErAttachmentBasicDependence11361)
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
					"instance_id":            "${alicloud_eflo_vcc.VCC.id}",
					"resource_tenant_id":     "1511928242963727",
					"er_id":                  "${alicloud_eflo_er.ER.id}",
					"er_attachment_name":     name,
					"instance_type":          "VCC",
					"auto_receive_all_route": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":            CHECKSET,
						"resource_tenant_id":     CHECKSET,
						"er_id":                  CHECKSET,
						"er_attachment_name":     name,
						"instance_type":          "VCC",
						"auto_receive_all_route": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"er_attachment_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"er_attachment_name": name + "_update",
					}),
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

var AlicloudEfloErAttachmentMap11361 = map[string]string{
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"er_attachment_id": CHECKSET,
	"region_id":        CHECKSET,
}

func AlicloudEfloErAttachmentBasicDependence11361(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

variable "region" {
  default = "cn-wulanchabu"
}

resource "alicloud_eflo_er" "ER" {
  er_name        = "er_at_test"
  master_zone_id = var.zone_id
}

resource "alicloud_vpc" "VPC" {
  is_default  = false
  dry_run     = false
  cidr_block  = "192.168.0.0/16"
  enable_ipv6 = false
}

resource "alicloud_vswitch" "VSW" {
  vpc_id     = alicloud_vpc.VPC.id
  zone_id    = var.zone_id
  cidr_block = "192.168.0.0/24"
}

resource "alicloud_eflo_vcc" "VCC" {
  connection_type = "VPC"
  zone_id         = var.zone_id
  vswitch_id      = alicloud_vswitch.VSW.id
  vpc_id          = alicloud_vpc.VPC.id
  bandwidth       = "1000"
  vcc_name        = "ER_AT_TEST"
}


`, name)
}

// Case er_at_tf_vpd 12464
func TestAccAliCloudEfloErAttachment_basic12464(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_er_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloErAttachmentMap12464)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloErAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloErAttachmentBasicDependence12464)
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
					"resource_tenant_id":     "1511928242963727",
					"instance_id":            "${alicloud_eflo_vpd.VPD.id}",
					"auto_receive_all_route": "true",
					"er_id":                  "${alicloud_eflo_er.ER.id}",
					"er_attachment_name":     name,
					"instance_type":          "VPD",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_tenant_id":     CHECKSET,
						"instance_id":            CHECKSET,
						"auto_receive_all_route": "true",
						"er_id":                  CHECKSET,
						"er_attachment_name":     name,
						"instance_type":          "VPD",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"er_attachment_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"er_attachment_name": name + "_update",
					}),
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

var AlicloudEfloErAttachmentMap12464 = map[string]string{
	"status":           CHECKSET,
	"create_time":      CHECKSET,
	"er_attachment_id": CHECKSET,
	"region_id":        CHECKSET,
}

func AlicloudEfloErAttachmentBasicDependence12464(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id" {
  default = "cn-wulanchabu-b"
}

resource "alicloud_eflo_vpd" "VPD" {
  cidr                  = "10.0.0.0/8"
  secondary_cidr_blocks = []
  vpd_name              = "er_at_tf_test"
}

resource "alicloud_eflo_er" "ER" {
  er_name        = "TEST"
  master_zone_id = var.zone_id
}


`, name)
}

// Test Eflo ErAttachment. <<< Resource test cases, automatically generated.
