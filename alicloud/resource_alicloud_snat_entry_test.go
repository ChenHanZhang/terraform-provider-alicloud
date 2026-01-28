package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test NATGateway SnatEntry. >>> Resource test cases, automatically generated.
// Case 私网NAT_SnatEntry_1.0.1 8016
func TestAccAliCloudNATGatewaySnatEntry_basic8016(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_snat_entry.default"
	ra := resourceAttrInit(resourceId, AlicloudNATGatewaySnatEntryMap8016)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NATGatewayServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNATGatewaySnatEntry")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnatgateway%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNATGatewaySnatEntryBasicDependence8016)
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
					"snat_ip":           "${alicloud_vpc_nat_ip.NatIp1.nat_ip}",
					"snat_table_id":     "${alicloud_nat_gateway.NATGateway.snat_table_ids[0]}",
					"source_vswitch_id": "${alicloud_vswitch.VSwitch.id}",
					"snat_entry_name":   name,
					"eip_affinity":      "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"snat_ip":           CHECKSET,
						"snat_table_id":     CHECKSET,
						"source_vswitch_id": CHECKSET,
						"snat_entry_name":   name,
						"eip_affinity":      "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"snat_ip":         "${alicloud_vpc_nat_ip.NatIp1.nat_ip}",
					"snat_entry_name": name + "_update",
					"eip_affinity":    "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"snat_ip":         CHECKSET,
						"snat_entry_name": name + "_update",
						"eip_affinity":    "1",
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

var AlicloudNATGatewaySnatEntryMap8016 = map[string]string{
	"status":        CHECKSET,
	"snat_entry_id": CHECKSET,
}

func AlicloudNATGatewaySnatEntryBasicDependence8016(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "tf-test-nat-snat-vpc"
}

resource "alicloud_vswitch" "VSwitch" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1b"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "tf-test-nat-snat-vsw"
}

resource "alicloud_nat_gateway" "NATGateway" {
  nat_gateway_private_info {
    iz_no      = "eu-central-1b"
    vswitch_id = alicloud_vswitch.VSwitch.id
  }
  vpc_id                      = alicloud_vpc.vpc.id
  network_type                = "intranet"
  icmp_reply_enabled          = false
  nat_gateway_name            = "tf-test-snat-private-nat-1"
  eip_bind_mode               = "NAT"
  security_protection_enabled = false
  auto_pay                    = false
  nat_type                    = "Enhanced"
  payment_type                = "PayAsYouGo"
}

resource "alicloud_vpc_nat_ip" "NatIp1" {
  nat_ip         = "172.16.0.66"
  nat_ip_name    = "tf-test-snat-natip1"
  nat_gateway_id = alicloud_nat_gateway.NATGateway.id
  nat_ip_cidr    = alicloud_vswitch.VSwitch.cidr_block
}

resource "alicloud_vpc_nat_ip" "NatIp2" {
  nat_ip         = "172.16.0.88"
  nat_ip_cidr    = alicloud_vswitch.VSwitch.cidr_block
  nat_ip_name    = "tf-test-snat-natip2"
  nat_gateway_id = alicloud_nat_gateway.NATGateway.id
}


`, name)
}

// Case 私网NAT_SnatEntry_VPC粒度 12474
func TestAccAliCloudNATGatewaySnatEntry_basic12474(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_snat_entry.default"
	ra := resourceAttrInit(resourceId, AlicloudNATGatewaySnatEntryMap12474)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NATGatewayServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNATGatewaySnatEntry")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnatgateway%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNATGatewaySnatEntryBasicDependence12474)
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
					"snat_ip":         "${alicloud_vpc_nat_ip.NatIp1.nat_ip}",
					"snat_table_id":   "${alicloud_nat_gateway.NATGateway.snat_table_ids[0]}",
					"snat_entry_name": name,
					"eip_affinity":    "0",
					"source_cidr":     "0.0.0.0/0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"snat_ip":         CHECKSET,
						"snat_table_id":   CHECKSET,
						"snat_entry_name": name,
						"eip_affinity":    "0",
						"source_cidr":     "0.0.0.0/0",
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

var AlicloudNATGatewaySnatEntryMap12474 = map[string]string{
	"status":        CHECKSET,
	"snat_entry_id": CHECKSET,
}

func AlicloudNATGatewaySnatEntryBasicDependence12474(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "172.16.0.0/16"
  vpc_name   = "tf-test-nat-snat-vpc"
}

resource "alicloud_vswitch" "VSwitch" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1b"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "tf-test-nat-snat-vsw"
}

resource "alicloud_nat_gateway" "NATGateway" {
  nat_gateway_private_info {
    iz_no      = "eu-central-1b"
    vswitch_id = alicloud_vswitch.VSwitch.id
  }
  vpc_id                      = alicloud_vpc.vpc.id
  network_type                = "intranet"
  icmp_reply_enabled          = false
  nat_gateway_name            = "tf-test-snat-private-nat-2"
  eip_bind_mode               = "NAT"
  security_protection_enabled = false
  auto_pay                    = false
  nat_type                    = "Enhanced"
  payment_type                = "PayAsYouGo"
}

resource "alicloud_vpc_nat_ip" "NatIp1" {
  nat_ip         = "172.16.0.68"
  nat_ip_name    = "tf-test-snat-natip1"
  nat_gateway_id = alicloud_nat_gateway.NATGateway.id
  nat_ip_cidr    = alicloud_vswitch.VSwitch.cidr_block
}


`, name)
}

// Test NATGateway SnatEntry. <<< Resource test cases, automatically generated.
