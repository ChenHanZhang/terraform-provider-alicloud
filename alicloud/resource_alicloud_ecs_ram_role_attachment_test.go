// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ecs RamRoleAttachment. >>> Resource test cases, automatically generated.
// Case RamRoleAttachment 10656
func TestAccAliCloudEcsRamRoleAttachment_basic10656(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_ram_role_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsRamRoleAttachmentMap10656)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsRamRoleAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccecs%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsRamRoleAttachmentBasicDependence10656)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"policy":        "{\\\"Statement\\\": [{\\\"Action\\\": [\\\"*\\\"],\\\"Effect\\\": \\\"Allow\\\",\\\"Resource\\\": [\\\"*\\\"]}],\\\"Version\\\":\\\"1\\\"}",
					"ram_role_name": "${alicloud_ram_role.ram.id}",
					"instance_id":   "${alicloud_ecs_instance.instance.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"policy":        CHECKSET,
						"ram_role_name": CHECKSET,
						"instance_id":   CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"policy"},
			},
		},
	})
}

var AlicloudEcsRamRoleAttachmentMap10656 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudEcsRamRoleAttachmentBasicDependence10656(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "172.16.0.0/18"
  vpc_name   = "tf-test"
}

resource "alicloud_vswitch" "vsw" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "cn-hangzhou-i"
  cidr_block   = "172.16.0.0/18"
  vswitch_name = "tf-test"
}

resource "alicloud_security_group" "sg" {
  description         = "sg"
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.vpc.id
  security_group_type = "normal"
}

resource "alicloud_ecs_instance" "instance" {
  system_disk {
    size     = "20"
    category = "cloud_essd"
  }
  status       = "Running"
  image_family = "acs:alibaba_cloud_linux_3_2104_lts_x64"
  vpc_attributes {
    vpc_id     = alicloud_vpc.vpc.id
    vswitch_id = alicloud_vswitch.vsw.id
  }
  instance_name     = "tf-image-ecs"
  password          = "Ali@qa1234"
  security_group_id = alicloud_security_group.sg.id
  instance_type     = "ecs.g6e.large"
}

resource "alicloud_ram_role" "ram" {
  role_name                   = "tfTest1770205954"
  assume_role_policy_document = "{\"Statement\": [{\"Action\": \"sts:AssumeRole\", \"Effect\": \"Allow\", \"Principal\": {\"Service\": [\"ecs.aliyuncs.com\"]}}], \"Version\": \"1\"}"
}


`, name)
}

// Test Ecs RamRoleAttachment. <<< Resource test cases, automatically generated.
