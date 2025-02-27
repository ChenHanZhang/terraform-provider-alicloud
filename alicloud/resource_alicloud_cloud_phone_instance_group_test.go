package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test CloudPhone InstanceGroup. >>> Resource test cases, automatically generated.
// Case chuyuan_CreateInstanceGroup_prod_all_vsw 10335
func TestAccAliCloudCloudPhoneInstanceGroup_basic10335(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_phone_instance_group.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudPhoneInstanceGroupMap10335)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudPhoneServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudPhoneInstanceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudphone%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudPhoneInstanceGroupBasicDependence10335)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_group_spec": "acp.basic.small",
					"instance_group_name": name,
					"period":              "1",
					"number_of_instances": "1",
					"charge_type":         "PostPaid",
					"image_id":            "imgc-075cllfeuazh03tg9",
					"period_unit":         "Hour",
					"auto_renew":          "false",
					"amount":              "1",
					"auto_pay":            "false",
					"gpu_acceleration":    "false",
					"policy_group_id":     "${alicloud_cloud_phone_policy.defaultjZ1gi0.id}",
					"office_site_id":      "${alicloud_ecd_simple_office_site.defaultq0h9R9.id}",
					"vswitch_id":          "${alicloud_vswitch.default4uQKQR.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_group_spec": "acp.basic.small",
						"instance_group_name": name,
						"period":              "1",
						"number_of_instances": "1",
						"charge_type":         "PostPaid",
						"image_id":            "imgc-075cllfeuazh03tg9",
						"period_unit":         "Hour",
						"auto_renew":          "false",
						"amount":              "1",
						"auto_pay":            "false",
						"gpu_acceleration":    "false",
						"policy_group_id":     CHECKSET,
						"office_site_id":      CHECKSET,
						"vswitch_id":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_group_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_group_name": name + "_update",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "gpu_acceleration", "office_site_id", "period", "period_unit", "policy_group_id", "vswitch_id"},
			},
		},
	})
}

var AlicloudCloudPhoneInstanceGroupMap10335 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCloudPhoneInstanceGroupBasicDependence10335(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "vsw_cidr_block" {
    default = "192.168.0.0/24"
}

variable "vpc_cidr_block" {
    default = "192.168.0.0/16"
}

variable "region_id" {
    default = "cn-shanghai"
}

resource "alicloud_cloud_phone_policy" "defaultjZ1gi0" {
	lock_resolution = "off"
	resolution_width = "720"
	camera_redirect = "on"
	policy_group_name = "defaultPolicyGroup"
	resolution_height = "1280"
	clipboard = "readwrite"
	net_redirect_policy  {
		net_redirect = "off"
		custom_proxy = "off"
	}
}

resource "alicloud_vpc" "defaultBhvqvn" {
	is_default = false
	cidr_block = var.vpc_cidr_block
	enable_ipv6 = false
	vpc_name = "tf-test-01"
}

resource "alicloud_vswitch" "default4uQKQR" {
	vpc_id = alicloud_vpc.defaultBhvqvn.id
	zone_id = "cn-shanghai-l"
	cidr_block = var.vsw_cidr_block
}

resource "alicloud_ecd_simple_office_site" "defaultq0h9R9" {
    # It is necessary to specify the VSwitchId; otherwise, the instance startup will fail.
	cidr_block = "${var.vpc_cidr_block}"
	office_site_name = "Init-Terraform-SimpleOfficeSite"
}


`, name)
}

// Case chuyuan_CreateInstanceGroup_prod_all_1 10200
func TestAccAliCloudCloudPhoneInstanceGroup_basic10200(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_phone_instance_group.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudPhoneInstanceGroupMap10200)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudPhoneServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudPhoneInstanceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudphone%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudPhoneInstanceGroupBasicDependence10200)
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
					"instance_group_spec": "acp.basic.small",
					"instance_group_name": name,
					"period":              "1",
					"number_of_instances": "1",
					"charge_type":         "PostPaid",
					"image_id":            "imgc-075cllfeuazh03tg9",
					"period_unit":         "Hour",
					"auto_renew":          "false",
					"amount":              "1",
					"auto_pay":            "false",
					"gpu_acceleration":    "false",
					"policy_group_id":     "${alicloud_cloud_phone_policy.defaultjZ1gi0.id}",
					"office_site_id":      "${alicloud_ecd_simple_office_site.defaultH2a5KS.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_group_spec": "acp.basic.small",
						"instance_group_name": name,
						"period":              "1",
						"number_of_instances": "1",
						"charge_type":         "PostPaid",
						"image_id":            "imgc-075cllfeuazh03tg9",
						"period_unit":         "Hour",
						"auto_renew":          "false",
						"amount":              "1",
						"auto_pay":            "false",
						"gpu_acceleration":    "false",
						"policy_group_id":     CHECKSET,
						"office_site_id":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_group_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_group_name": name + "_update",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "gpu_acceleration", "office_site_id", "period", "period_unit", "policy_group_id", "vswitch_id"},
			},
		},
	})
}

var AlicloudCloudPhoneInstanceGroupMap10200 = map[string]string{
	"status": CHECKSET,
}

func AlicloudCloudPhoneInstanceGroupBasicDependence10200(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

resource "alicloud_cloud_phone_policy" "defaultjZ1gi0" {
  lock_resolution   = "off"
  resolution_width  = "720"
  camera_redirect   = "on"
  policy_group_name = "defaultPolicyGroup"
  resolution_height = "1280"
  clipboard         = "readwrite"
  net_redirect_policy {
    net_redirect = "off"
    custom_proxy = "off"
  }
}

resource "alicloud_ecd_simple_office_site" "defaultH2a5KS" {
  office_site_name = "InitOfficeSite"
  cidr_block       = "172.16.0.0/12"
}


`, name)
}

// Test CloudPhone InstanceGroup. <<< Resource test cases, automatically generated.
