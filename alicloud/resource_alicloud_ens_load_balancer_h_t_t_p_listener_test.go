// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Ens LoadBalancerHTTPListener. >>> Resource test cases, automatically generated.
// Case HTTPListener_20240508 6685
func TestAccAliCloudEnsLoadBalancerHTTPListener_basic6685(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ens_load_balancer_h_t_t_p_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudEnsLoadBalancerHTTPListenerMap6685)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EnsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEnsLoadBalancerHTTPListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccens%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEnsLoadBalancerHTTPListenerBasicDependence6685)
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
					"listener_port":             "8082",
					"request_timeout":           "60",
					"health_check_interval":     "5",
					"description":               "testD5esc1",
					"unhealthy_threshold":       "5",
					"scheduler":                 "wrr",
					"health_check_uri":          "/test1.html",
					"health_check":              "on",
					"idle_timeout":              "5",
					"load_balancer_id":          "${alicloud_ens_load_balancer.defaultmaG48i.id}",
					"health_check_timeout":      "5",
					"health_check_connect_port": "8082",
					"health_check_method":       "head",
					"healthy_threshold":         "5",
					"x_forwarded_for":           "off",
					"health_check_domain":       "test1.com",
					"health_check_http_code":    "http_2xx",
					"status":                    "Stopped",
					"backend_server_port":       "8082",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_port":             "8082",
						"request_timeout":           "60",
						"health_check_interval":     "5",
						"description":               "testD5esc1",
						"unhealthy_threshold":       "5",
						"scheduler":                 "wrr",
						"health_check_uri":          "/test1.html",
						"health_check":              "on",
						"idle_timeout":              "5",
						"load_balancer_id":          CHECKSET,
						"health_check_timeout":      "5",
						"health_check_connect_port": "8082",
						"health_check_method":       "head",
						"healthy_threshold":         "5",
						"x_forwarded_for":           "off",
						"health_check_domain":       "test1.com",
						"health_check_http_code":    "http_2xx",
						"status":                    "Stopped",
						"backend_server_port":       "8082",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"request_timeout":           "10",
					"health_check_interval":     "10",
					"description":               "testDesc2",
					"unhealthy_threshold":       "10",
					"scheduler":                 "rr",
					"health_check_uri":          "/test2.html",
					"idle_timeout":              "10",
					"health_check_timeout":      "10",
					"health_check_connect_port": "8085",
					"health_check_method":       "get",
					"healthy_threshold":         "10",
					"x_forwarded_for":           "on",
					"health_check_domain":       "test2.com",
					"health_check_http_code":    "http_3xx",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"request_timeout":           "10",
						"health_check_interval":     "10",
						"description":               "testDesc2",
						"unhealthy_threshold":       "10",
						"scheduler":                 "rr",
						"health_check_uri":          "/test2.html",
						"idle_timeout":              "10",
						"health_check_timeout":      "10",
						"health_check_connect_port": "8085",
						"health_check_method":       "get",
						"healthy_threshold":         "10",
						"x_forwarded_for":           "on",
						"health_check_domain":       "test2.com",
						"health_check_http_code":    "http_3xx",
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
				Config: testAccConfig(map[string]interface{}{
					"health_check": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"health_check": "off",
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

var AlicloudEnsLoadBalancerHTTPListenerMap6685 = map[string]string{}

func AlicloudEnsLoadBalancerHTTPListenerBasicDependence6685(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "ens_region_id" {
  default = "cn-hangzhou-47"
}

resource "alicloud_ens_network" "defaultgkJXuV" {
  network_name  = "测试用例-http监听示例"
  cidr_block    = "10.0.0.0/8"
  ens_region_id = var.ens_region_id
}

resource "alicloud_ens_vswitch" "defaultQy3IhQ" {
  cidr_block    = "10.0.8.0/24"
  vswitch_name  = "测试用例-负载均衡-http监听"
  ens_region_id = alicloud_ens_network.defaultgkJXuV.ens_region_id
  network_id    = alicloud_ens_network.defaultgkJXuV.id
}

resource "alicloud_ens_load_balancer" "defaultmaG48i" {
  load_balancer_name = "测试用例-http监听"
  vswitch_id         = alicloud_ens_vswitch.defaultQy3IhQ.id
  payment_type       = "PayAsYouGo"
  ens_region_id      = alicloud_ens_vswitch.defaultQy3IhQ.ens_region_id
  network_id         = alicloud_ens_vswitch.defaultQy3IhQ.network_id
  load_balancer_spec = "elb.s1.small"
}


`, name)
}

// Test Ens LoadBalancerHTTPListener. <<< Resource test cases, automatically generated.
