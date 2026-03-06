// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ens LoadBalancerTCPListener. >>> Resource test cases, automatically generated.
// Case TCPListener_20240514 6762
func TestAccAliCloudEnsLoadBalancerTCPListener_basic6762(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_load_balancer_t_c_p_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsLoadBalancerTCPListenerMap6762)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsLoadBalancerTCPListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsLoadBalancerTCPListenerBasicDependence6762)
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
					"listener_port":                "5000",
					"description":                  "test",
					"health_check_interval":        "10",
					"unhealthy_threshold":          "10",
					"scheduler":                    "wrr",
					"health_check_uri":             "/test.html",
					"health_check_connect_timeout": "20",
					"load_balancer_id":             "${alicloud_ens_load_balancer.defaultrlO87U.id}",
					"backend_server_port":          "5000",
					"persistence_timeout":          "0",
					"health_check_connect_port":    "5000",
					"healthy_threshold":            "10",
					"health_check_domain":          "test.com",
					"eip_transmit":                 "off",
					"health_check_http_code":       "http_2xx",
					"health_check_type":            "http",
					"established_timeout":          "50",
					"status":                       "Stopped",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_port":                "5000",
						"description":                  "test",
						"health_check_interval":        "10",
						"unhealthy_threshold":          "10",
						"scheduler":                    "wrr",
						"health_check_uri":             "/test.html",
						"health_check_connect_timeout": "20",
						"load_balancer_id":             CHECKSET,
						"backend_server_port":          "5000",
						"persistence_timeout":          "0",
						"health_check_connect_port":    "5000",
						"healthy_threshold":            "10",
						"health_check_domain":          "test.com",
						"eip_transmit":                 "off",
						"health_check_http_code":       "http_2xx",
						"health_check_type":            "http",
						"established_timeout":          "50",
						"status":                       "Stopped",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"health_check_interval":        "5",
					"unhealthy_threshold":          "5",
					"health_check_connect_timeout": "5",
					"persistence_timeout":          "5",
					"health_check_connect_port":    "6000",
					"healthy_threshold":            "5",
					"health_check_type":            "tcp",
					"established_timeout":          "10",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"health_check_interval":        "5",
						"unhealthy_threshold":          "5",
						"health_check_connect_timeout": "5",
						"persistence_timeout":          "5",
						"health_check_connect_port":    "6000",
						"healthy_threshold":            "5",
						"health_check_type":            "tcp",
						"established_timeout":          "10",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description":                  "test2",
					"health_check_interval":        "10",
					"unhealthy_threshold":          "10",
					"scheduler":                    "wlc",
					"health_check_connect_timeout": "10",
					"persistence_timeout":          "10",
					"health_check_connect_port":    "4000",
					"healthy_threshold":            "10",
					"health_check_domain":          "test2.com",
					"eip_transmit":                 "on",
					"health_check_type":            "http",
					"established_timeout":          "20",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":                  "test2",
						"health_check_interval":        "10",
						"unhealthy_threshold":          "10",
						"scheduler":                    "wlc",
						"health_check_connect_timeout": "10",
						"persistence_timeout":          "10",
						"health_check_connect_port":    "4000",
						"healthy_threshold":            "10",
						"health_check_domain":          "test2.com",
						"eip_transmit":                 "on",
						"health_check_type":            "http",
						"established_timeout":          "20",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"health_check_uri":       "/test3.html",
					"health_check_domain":    "test3.com",
					"health_check_http_code": "http_3xx",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"health_check_uri":       "/test3.html",
						"health_check_domain":    "test3.com",
						"health_check_http_code": "http_3xx",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Running",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Stopped",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Stopped",
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

var AlicloudEnsLoadBalancerTCPListenerMap6762 = map[string]string{}

func AlicloudEnsLoadBalancerTCPListenerBasicDependence6762(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "ens_region_id" {
  default = "cn-hangzhou-44"
}

resource "alicloud_ens_network" "defaultzsNbyZ" {
  network_name  = "测试用例-tcp监听示例"
  cidr_block    = "10.0.0.0/8"
  ens_region_id = var.ens_region_id
}

resource "alicloud_ens_vswitch" "defaultYueZqF" {
  cidr_block    = "10.0.8.0/24"
  ens_region_id = alicloud_ens_network.defaultzsNbyZ.ens_region_id
  network_id    = alicloud_ens_network.defaultzsNbyZ.id
}

resource "alicloud_ens_load_balancer" "defaultrlO87U" {
  load_balancer_name = "测试用例-TCP监听"
  vswitch_id         = alicloud_ens_vswitch.defaultYueZqF.id
  payment_type       = "PayAsYouGo"
  ens_region_id      = alicloud_ens_vswitch.defaultYueZqF.ens_region_id
  network_id         = alicloud_ens_vswitch.defaultYueZqF.network_id
  load_balancer_spec = "elb.s1.small"
}


`, name)
}

// Test Ens LoadBalancerTCPListener. <<< Resource test cases, automatically generated.
