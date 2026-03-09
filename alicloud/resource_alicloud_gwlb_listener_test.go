package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Gwlb Listener. >>> Resource test cases, automatically generated.
// Case Test Listener 8508
func TestAccAliCloudGwlbListener_basic8508(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_gwlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudGwlbListenerMap8508)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &GwlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeGwlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccgwlb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudGwlbListenerBasicDependence8508)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.EfloSupportRegions)
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"listener_description": "test-tf-lsn",
					"server_group_id":      "${alicloud_gwlb_server_group.defaultoAkLbr.id}",
					"load_balancer_id":     "${alicloud_gwlb_load_balancer.defaultQ5setL.id}",
					"dry_run":              "false",
					"tcp_idle_timeout":     "900",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_description": "test-tf-lsn",
						"server_group_id":      CHECKSET,
						"load_balancer_id":     CHECKSET,
						"dry_run":              "false",
						"tcp_idle_timeout":     "900",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"listener_description": "test-tf-update",
					"server_group_id":      "${alicloud_gwlb_server_group.defaultN4DOzm.id}",
					"tcp_idle_timeout":     "350",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_description": "test-tf-update",
						"server_group_id":      CHECKSET,
						"tcp_idle_timeout":     "350",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tcp_idle_timeout": "360",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tcp_idle_timeout": "360",
					}),
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
				ImportStateVerifyIgnore: []string{"dry_run"},
			},
		},
	})
}

var AlicloudGwlbListenerMap8508 = map[string]string{
	"status":    CHECKSET,
	"region_id": CHECKSET,
}

func AlicloudGwlbListenerBasicDependence8508(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "zone_id2" {
  default = "cn-wulanchabu-c"
}

variable "zone_id1" {
  default = "cn-wulanchabu-b"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaulti9Axhl" {
  cidr_block = "10.0.0.0/8"
  vpc_name   = "tf-gwlb-vpc"
}

resource "alicloud_vswitch" "default9NaKmL" {
  vpc_id       = alicloud_vpc.defaulti9Axhl.id
  zone_id      = var.zone_id1
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "tf-test-vsw1"
}

resource "alicloud_vswitch" "defaultH4pKT4" {
  vpc_id       = alicloud_vpc.defaulti9Axhl.id
  zone_id      = var.zone_id2
  cidr_block   = "10.0.1.0/24"
  vswitch_name = "tf-test-vsw2"
}

resource "alicloud_gwlb_load_balancer" "defaultQ5setL" {
  vpc_id             = alicloud_vpc.defaulti9Axhl.id
  load_balancer_name = "test-tf"
  zone_mappings {
    vswitch_id = alicloud_vswitch.default9NaKmL.id
    zone_id    = var.zone_id1
  }
  address_ip_version = "Ipv4"
  dry_run            = false
}

resource "alicloud_gwlb_server_group" "defaultoAkLbr" {
  scheduler = "5TCH"
  health_check_config {
    health_check_protocol        = "HTTP"
    health_check_connect_port    = "80"
    health_check_connect_timeout = "5"
    health_check_domain          = "www.domain.com"
    health_check_enabled         = true
    health_check_http_code       = ["http_2xx", "http_3xx", "http_4xx"]
    health_check_interval        = "10"
    health_check_path            = "/health-check"
    healthy_threshold            = "2"
    unhealthy_threshold          = "2"
  }
  protocol          = "GENEVE"
  server_group_type = "Ip"
  connection_drain_config {
    connection_drain_enabled = true
    connection_drain_timeout = "1"
  }
  vpc_id  = alicloud_vpc.defaulti9Axhl.id
  dry_run = false
  servers {
    server_id   = "10.0.0.1"
    server_ip   = "10.0.0.1"
    server_type = "Ip"
  }
  servers {
    server_id   = "10.0.0.2"
    server_ip   = "10.0.0.2"
    server_type = "Ip"
  }
  servers {
    server_id   = "10.0.0.3"
    server_ip   = "10.0.0.3"
    server_type = "Ip"
  }
  server_group_name = "tf-test-sgp"
}

resource "alicloud_gwlb_server_group" "defaultN4DOzm" {
  scheduler = "5TCH"
  health_check_config {
    health_check_protocol        = "HTTP"
    health_check_connect_port    = "80"
    health_check_connect_timeout = "5"
    health_check_domain          = "www.domain.com"
    health_check_enabled         = true
    health_check_http_code       = ["http_2xx", "http_3xx", "http_4xx"]
    health_check_interval        = "10"
    health_check_path            = "/health-check"
    healthy_threshold            = "2"
    unhealthy_threshold          = "2"
  }
  protocol          = "GENEVE"
  server_group_type = "Ip"
  connection_drain_config {
    connection_drain_enabled = true
    connection_drain_timeout = "1"
  }
  vpc_id  = alicloud_vpc.defaulti9Axhl.id
  dry_run = false
  servers {
    server_id   = "10.0.0.1"
    server_ip   = "10.0.0.1"
    server_type = "Ip"
  }
  servers {
    server_id   = "10.0.0.2"
    server_ip   = "10.0.0.2"
    server_type = "Ip"
  }
  servers {
    server_id   = "10.0.0.3"
    server_ip   = "10.0.0.3"
    server_type = "Ip"
  }
  server_group_name = "tf-test-sgp-2"
}


`, name)
}

// Test Gwlb Listener. <<< Resource test cases, automatically generated.
