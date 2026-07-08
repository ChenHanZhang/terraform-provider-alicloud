package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Alb LoadBalancerSecurityGroupAttachment. >>> Resource test cases, automatically generated.
// Case LoadBalancerSecurityGroupAttachment_自动化 7132
func TestAccAliCloudAlbLoadBalancerSecurityGroupAttachment_basic7132(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alb_load_balancer_security_group_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudAlbLoadBalancerSecurityGroupAttachmentMap7132)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAlbLoadBalancerSecurityGroupAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAlbLoadBalancerSecurityGroupAttachmentBasicDependence7132)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-chengdu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"security_group_id": "${alicloud_security_group.defaultY97MLV.id}",
					"load_balancer_id":  "${alicloud_alb_load_balancer.defaultHRfk1j.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_group_id": CHECKSET,
						"load_balancer_id":  CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

var AlicloudAlbLoadBalancerSecurityGroupAttachmentMap7132 = map[string]string{}

func AlicloudAlbLoadBalancerSecurityGroupAttachmentBasicDependence7132(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region" {
  default = "cn-chengdu"
}

variable "zone2" {
  default = "cn-chengdu-a"
}

variable "zone1" {
  default = "cn-chengdu-b"
}

resource "alicloud_vpc" "defaultgjB1hS" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "test"
}

resource "alicloud_vswitch" "defaultDB9ylt" {
  vpc_id       = alicloud_vpc.defaultgjB1hS.id
  zone_id      = var.zone1
  cidr_block   = "192.168.1.0/24"
  vswitch_name = "vsw1"
}

resource "alicloud_vswitch" "default2s1hXk" {
  vpc_id       = alicloud_vpc.defaultgjB1hS.id
  zone_id      = var.zone2
  cidr_block   = "192.168.2.0/24"
  vswitch_name = "vsw2"
}

resource "alicloud_security_group" "defaultY97MLV" {
  security_group_name = "test_tf_security_group"
  vpc_id              = alicloud_vpc.defaultgjB1hS.id
}

resource "alicloud_alb_load_balancer" "defaultHRfk1j" {
  load_balancer_name    = "test_tf_security_group"
  load_balancer_edition = "Standard"
  vpc_id                = alicloud_vpc.defaultgjB1hS.id
  load_balancer_billing_config {
    pay_type = "PostPay"
  }
  address_type           = "Intranet"
  address_allocated_mode = "Fixed"
  zone_mappings {
    vswitch_id = alicloud_vswitch.defaultDB9ylt.id
    zone_id    = alicloud_vswitch.defaultDB9ylt.zone_id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.default2s1hXk.id
    zone_id    = alicloud_vswitch.default2s1hXk.zone_id
  }
}


`, name)
}

// Test Alb LoadBalancerSecurityGroupAttachment. <<< Resource test cases, automatically generated.
