package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestAccAliCloudNlbListener_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%snlblistener%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependence0)
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
					"listener_protocol":      "TCP",
					"listener_port":          "80",
					"listener_description":   "${var.name}",
					"load_balancer_id":       "${alicloud_nlb_load_balancer.default.id}",
					"server_group_id":        "${alicloud_nlb_server_group.default.id}",
					"idle_timeout":           "900",
					"proxy_protocol_enabled": "true",
					"sec_sensor_enabled":     "true",
					"cps":                    "10000",
					"mss":                    "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_protocol":      "TCP",
						"listener_port":          "80",
						"listener_description":   name,
						"load_balancer_id":       CHECKSET,
						"server_group_id":        CHECKSET,
						"idle_timeout":           "900",
						"proxy_protocol_enabled": "true",
						"sec_sensor_enabled":     "true",
						"cps":                    "10000",
						"mss":                    "0",
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

func TestAccAliCloudNlbListener_TCPSSL(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%snlblistener%d", "cn-hangzhou", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependenceTCPSSL)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, connectivity.NLBSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"listener_protocol":      "TCPSSL",
					"listener_port":          "1883",
					"security_policy_id":     "tls_cipher_policy_1_0",
					"listener_description":   "${var.name}",
					"load_balancer_id":       "${alicloud_nlb_load_balancer.default.id}",
					"server_group_id":        "${alicloud_nlb_server_group.default.id}",
					"idle_timeout":           "900",
					"certificate_ids":        []string{"${local.certificate_id}"},
					"proxy_protocol_enabled": "true",
					"sec_sensor_enabled":     "true",
					"alpn_enabled":           "true",
					"alpn_policy":            "HTTP2Optional",
					"cps":                    "10000",
					"mss":                    "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_protocol":      "TCPSSL",
						"listener_port":          "1883",
						"security_policy_id":     "tls_cipher_policy_1_0",
						"listener_description":   name,
						"load_balancer_id":       CHECKSET,
						"server_group_id":        CHECKSET,
						"idle_timeout":           "900",
						"certificate_ids.#":      "1",
						"alpn_policy":            "HTTP2Optional",
						"proxy_protocol_enabled": "true",
						"sec_sensor_enabled":     "true",
						"alpn_enabled":           "true",
						"cps":                    "10000",
						"mss":                    "0",
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

var AlicloudNlbListenerMap0 = map[string]string{}

func AlicloudNlbListenerBasicDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_nlb_zones" "default" {}
data "alicloud_vpcs" "default" {
    name_regex = "^default-NODELETING$"
}
data "alicloud_resource_manager_resource_groups" "default" {}
data "alicloud_vswitches" "default_1" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_nlb_zones.default.zones.0.id
}
data "alicloud_vswitches" "default_2" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_nlb_zones.default.zones.1.id
}
locals {
  zone_id_1    = data.alicloud_nlb_zones.default.zones.0.id
  vswitch_id_1 = data.alicloud_vswitches.default_1.ids[0]
  zone_id_2    = data.alicloud_nlb_zones.default.zones.1.id
  vswitch_id_2 = data.alicloud_vswitches.default_2.ids[0]
}
resource "alicloud_nlb_load_balancer" "default" {
  load_balancer_name = var.name
  resource_group_id  = data.alicloud_resource_manager_resource_groups.default.ids.0
  load_balancer_type = "Network"
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  tags               = {
    Created = "tfTestAcc0"
    For     = "Tftestacc 0"
  }
  vpc_id = data.alicloud_vpcs.default.ids.0
  zone_mappings {
    vswitch_id = local.vswitch_id_1
    zone_id    = local.zone_id_1
  }
  zone_mappings {
    vswitch_id = local.vswitch_id_2
    zone_id    = local.zone_id_2
  }
}

resource "alicloud_nlb_server_group" "default" {
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  server_group_name = var.name
  server_group_type = "Instance"
  vpc_id            = data.alicloud_vpcs.default.ids.0
  scheduler         = "Wrr"
  protocol          = "TCP"
  health_check {
	health_check_url =           "/test/index.html"
	health_check_domain =       "tf-testAcc.com"
    health_check_enabled         = true
    health_check_type            = "TCP"
    health_check_connect_port    = 0
    healthy_threshold            = 2
    unhealthy_threshold          = 2
    health_check_connect_timeout = 5
    health_check_interval        = 10
    http_check_method            = "GET"
    health_check_http_code       = ["http_2xx", "http_3xx", "http_4xx"]
  }
  connection_drain           = true
  connection_drain_timeout   = 60
  preserve_client_ip_enabled = true
  tags = {
    Created = "TF"
  }
  address_ip_version = "Ipv4"
}
`, name)
}

func AlicloudNlbListenerBasicDependenceTCPSSL(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_nlb_zones" "default" {}
data "alicloud_vpcs" "default" {
    name_regex = "^default-NODELETING$"
}
data "alicloud_resource_manager_resource_groups" "default" {}
data "alicloud_vswitches" "default_1" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_nlb_zones.default.zones.0.id
}
data "alicloud_vswitches" "default_2" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_nlb_zones.default.zones.1.id
}


	resource "alicloud_ssl_certificates_service_certificate" "default" {
  		certificate_name = var.name
  		cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID1zCCAr+gAwIBAgIRAOrWWz1qmkcSg90JDHjuzFwwDQYJKoZIhvcNAQELBQAw
XjELMAkGA1UEBhMCQ04xDjAMBgNVBAoTBU15U1NMMSswKQYDVQQLEyJNeVNTTCBU
ZXN0IFJTQSAtIEZvciB0ZXN0IHVzZSBvbmx5MRIwEAYDVQQDEwlNeVNTTC5jb20w
HhcNMjQxMTI2MDczNjA4WhcNMjkxMTI1MDczNjA4WjAgMQswCQYDVQQGEwJDTjER
MA8GA1UEAxMIdGVzdC5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIB
AQDa7HDGbQ1Km0f4ZaFzYbjVN0q8KkvZ+oQUd4naGOZnlH5k0XFwmjg+TWf88YX3
5IF8c45/rXrTWucPLg7FeqR96Wq9HZEmzEhs6VG031V9Hqa32saRScCOAyhiW7Hj
OWf6BZveuxbZNbgQCR59QzX4CeAIC68xavIDAy3wcTAH9cIkD71BxEPJGGR7BIVH
9DcWXaMAnJqQfrkth0xHBjflZABHAI0wPYPfaw8fd9DRkMYOIkfjwrrcL5IvhI1u
D3wdHJQWA2vR8hjoU4dHiJLbUtQ+xV1UGVkF67CpQ6LDjSQdX7xlZ7WJMc/7dCJ9
a7tr0ZTwq4/3KSgcRvm62oGvAgMBAAGjgc0wgcowDgYDVR0PAQH/BAQDAgWgMB0G
A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBQogSYF0TQa
P8FzD7uTzxUcPwO/fzBjBggrBgEFBQcBAQRXMFUwIQYIKwYBBQUHMAGGFWh0dHA6
Ly9vY3NwLm15c3NsLmNvbTAwBggrBgEFBQcwAoYkaHR0cDovL2NhLm15c3NsLmNv
bS9teXNzbHRlc3Ryc2EuY3J0MBMGA1UdEQQMMAqCCHRlc3QuY29tMA0GCSqGSIb3
DQEBCwUAA4IBAQAxPOlK5WBA9kITzxYyjqe/YvWzfMlsmj0yvpyHrPeZf7HZTTFz
ebYkzrHL8ZLyOHBhag0nL7Poj6ek98NoXTuCYCi8LspdadapOeYQzLce3beu/frk
sqU0A6WLHG9Ol9yUDMCX7xvLoAY/LDrcOM3Z87C/u/ykB4wKfFN2XfR3EZx3PQqw
sV77LOnyQixB4FMHpHlKuDoUkSN9uvxwEPOeGnLZXm96hPsjPwk1bDM8qerNPpVI
CwJ6kNuZ2eLz2Umqu2Gh3l4aADdIwxRY1OOjjZNut8STosABKWVGIwQbbAdRPQze
qHZ05oVTjFy9L1DAzhQ5Zn3oUjLl5KW4tYBA
-----END CERTIFICATE-----
EOF
  		key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA2uxwxm0NSptH+GWhc2G41TdKvCpL2fqEFHeJ2hjmZ5R+ZNFx
cJo4Pk1n/PGF9+SBfHOOf61601rnDy4OxXqkfelqvR2RJsxIbOlRtN9VfR6mt9rG
kUnAjgMoYlux4zln+gWb3rsW2TW4EAkefUM1+AngCAuvMWryAwMt8HEwB/XCJA+9
QcRDyRhkewSFR/Q3Fl2jAJyakH65LYdMRwY35WQARwCNMD2D32sPH3fQ0ZDGDiJH
48K63C+SL4SNbg98HRyUFgNr0fIY6FOHR4iS21LUPsVdVBlZBeuwqUOiw40kHV+8
ZWe1iTHP+3QifWu7a9GU8KuP9ykoHEb5utqBrwIDAQABAoIBAQCErEfIKOymKybZ
pZXLnAxswt563FMtngGPecZEM1TmrvpOVROffwbY0wZTJ3fd/FBwwIM6Y0MNdYiU
DYCMM0AewmeahqGh1qmJv3hx2eswMXQt9driz8RvDADcYt+SagbWYbHNsKovJrwO
k8gzd5jsYeewWIxqsXpLUxDzJ1VJbIqoHgkrirRRPo0onpixPWeA0RbElSwjwIUw
y43cC4WF8N7wot3cTST8yeKM8ujtqpN22ZtKnbkHTd03vnwQTMeUMJeDQmSmY5aJ
yFr7yw/Z66+7Amh6pkWhzZSDHsjI4y/S3CCdpwFlMA7ID590umJB6HFxWsmVacSe
MSs2vIJZAoGBAOiecPH1HVDQqH6PcrN/X9E3pDKSyAj+nHsVDGIZsie9f5g/qA0A
tcJtQLS0CzrpMTLsAnsfdh2T7Lg6pYFz5jnOUyMjOImAEbCtgvqBxqgFea//OhdP
8s/RmxKIAenBsk7Wbwx8/KPhbZLUNe8OnILVHDfS6kLSa49Iu+4UvrpNAoGBAPDt
mky5MMHKdHwbqxPo9jYrz1m3gqqIvv+VihO4t/DE6t2Zg43ctfFm1BVEDSwPjYs/
YV69KfVrVRUnzMZVdtHZ/dBK784YTY0OujemoaIzMKFIL8tbJFldVv2IgB+IelTX
e675hVdHjNUqZhHwccd8X6d/8icohZw62SNHb/HrAoGBAN1HSt1/c6Gau42Y212Q
fw9ARLuvEQYtXaFfxmXTV7uh8axccXndAQmwb+r1kfE6PojYJQwGQ4+jVX1ynFnm
bEz0zfUQ3gk+gJV2mK+/n7/ZZYZb3WCrtqimFUOtiVRZ40pHhV91zcX+/QK9R4je
d1elbbBUvG9QRu0IHW0+4qfJAoGAOmlQvIM1l/ZOsXw/yO71KoMKnXTJYDERJYQK
2ucw6VXEn39FjtJQ5jsI9jLugp0usvDl2YNBNfgUw7FHi1pTGWOhjqtsYmov+x/z
8+QZUerZQnDu7X2mXWgs3AEJFxwOlJ09pllmg5ecRF4oKvdBjpzP0BtMCURgyFTY
Kh56vIsCgYBMbneMvFY6PCESKIAXj16BF4lqYVXFqHVoxyfxIuVlAy3TMNwxvpbS
yDETk05Ux9yNES0WyTb1SWVG1o1wXc0dnDXCwJqLC1tzJUNUSD1AYvktoNIFErcN
gs3ercrzBTX5ezORPj9ErRAPrSq+V3z1Lge5Gl+EqgDvAfnknww75w==
-----END RSA PRIVATE KEY-----
EOF
	}

locals {
  zone_id_1    = data.alicloud_nlb_zones.default.zones.0.id
  vswitch_id_1 = data.alicloud_vswitches.default_1.ids[0]
  zone_id_2    = data.alicloud_nlb_zones.default.zones.1.id
  vswitch_id_2 = data.alicloud_vswitches.default_2.ids[0]
  certificate_id         = join("", [alicloud_ssl_certificates_service_certificate.default.id, "-%s"])
}
resource "alicloud_nlb_load_balancer" "default" {
  load_balancer_name = var.name
  resource_group_id  = data.alicloud_resource_manager_resource_groups.default.ids.0
  load_balancer_type = "Network"
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  tags               = {
    Created = "tfTestAcc0"
    For     = "Tftestacc 0"
  }
  vpc_id = data.alicloud_vpcs.default.ids.0
  zone_mappings {
    vswitch_id = local.vswitch_id_1
    zone_id    = local.zone_id_1
  }
  zone_mappings {
    vswitch_id = local.vswitch_id_2
    zone_id    = local.zone_id_2
  }
}
resource "alicloud_nlb_server_group" "default" {
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  server_group_name = var.name
  server_group_type = "Instance"
  vpc_id            = data.alicloud_vpcs.default.ids.0
  scheduler         = "Wrr"
  protocol          = "TCPSSL"
  health_check {
	health_check_url =           "/test/index.html"
	health_check_domain =       "tf-testAcc.com"
    health_check_enabled         = true
    health_check_type            = "TCP"
    health_check_connect_port    = 0
    healthy_threshold            = 2
    unhealthy_threshold          = 2
    health_check_connect_timeout = 5
    health_check_interval        = 10
    http_check_method            = "GET"
    health_check_http_code       = ["http_2xx", "http_3xx", "http_4xx"]
  }
  tags = {
    Created = "TF"
  }
  address_ip_version = "Ipv4"
}
`, name, "cn-hangzhou")
}

func TestAccAliCloudNlbListener_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap1)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%snlblistener%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependence1)
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
					"listener_protocol":      "TCP",
					"listener_port":          "80",
					"load_balancer_id":       "${alicloud_nlb_load_balancer.default.id}",
					"server_group_id":        "${alicloud_nlb_server_group.default.0.id}",
					"proxy_protocol_enabled": "true",
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_eps_id_enabled": "true",
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_protocol":      "TCP",
						"listener_port":          "80",
						"load_balancer_id":       CHECKSET,
						"server_group_id":        CHECKSET,
						"proxy_protocol_enabled": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"listener_description":   "${var.name}_update",
					"proxy_protocol_enabled": "false",
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_eps_id_enabled": "true",
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"listener_description":   name + "_update",
						"proxy_protocol_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cps":                    "0",
					"proxy_protocol_enabled": "true",
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_eps_id_enabled": "true",
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cps":                    "0",
						"proxy_protocol_enabled": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"mss": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"mss": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_id": "${alicloud_nlb_server_group.default.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"idle_timeout": "800",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"idle_timeout": "800",
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
					"proxy_protocol_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_protocol_enabled": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"sec_sensor_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sec_sensor_enabled": "true",
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

var AlicloudNlbListenerMap1 = map[string]string{}

func AlicloudNlbListenerBasicDependence1(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_nlb_zones" "default" {}
data "alicloud_vpcs" "default" {
    name_regex = "^default-NODELETING$"
}
data "alicloud_resource_manager_resource_groups" "default" {}
data "alicloud_vswitches" "default_1" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_nlb_zones.default.zones.0.id
}
data "alicloud_vswitches" "default_2" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_nlb_zones.default.zones.1.id
}
locals {
  zone_id_1    = data.alicloud_nlb_zones.default.zones.0.id
  vswitch_id_1 = data.alicloud_vswitches.default_1.ids[0]
  zone_id_2    = data.alicloud_nlb_zones.default.zones.1.id
  vswitch_id_2 = data.alicloud_vswitches.default_2.ids[0]
}
resource "alicloud_nlb_load_balancer" "default" {
  load_balancer_name = var.name
  resource_group_id  = data.alicloud_resource_manager_resource_groups.default.ids.0
  load_balancer_type = "Network"
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  tags               = {
    Created = "tfTestAcc0"
    For     = "Tftestacc 0"
  }
  vpc_id = data.alicloud_vpcs.default.ids.0
  zone_mappings {
    vswitch_id = local.vswitch_id_1
    zone_id    = local.zone_id_1
  }
  zone_mappings {
    vswitch_id = local.vswitch_id_2
    zone_id    = local.zone_id_2
  }
}
resource "alicloud_nlb_server_group" "default" {
  count = 2
  resource_group_id = data.alicloud_resource_manager_resource_groups.default.ids.0
  server_group_name = var.name
  server_group_type = "Instance"
  vpc_id            = data.alicloud_vpcs.default.ids.0
  scheduler         = "Wrr"
  protocol          = "TCP"
  health_check {
	health_check_url =           "/test/index.html"
	health_check_domain =       "tf-testAcc.com"
    health_check_enabled         = true
    health_check_type            = "TCP"
    health_check_connect_port    = 0
    healthy_threshold            = 2
    unhealthy_threshold          = 2
    health_check_connect_timeout = 5
    health_check_interval        = 10
    http_check_method            = "GET"
    health_check_http_code       = ["http_2xx", "http_3xx", "http_4xx"]
  }
  connection_drain           = true
  connection_drain_timeout   = 60
  preserve_client_ip_enabled = true
  tags = {
    Created = "TF"
  }
  address_ip_version = "Ipv4"
}
`, name)
}

// lintignore: R001
func TestUnitAlicloudNlbListener(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"listener_description":   "CreateListenerValue",
		"listener_port":          10,
		"listener_protocol":      "CreateListenerValue",
		"load_balancer_id":       "CreateListenerValue",
		"server_group_id":        "CreateListenerValue",
		"idle_timeout":           10,
		"cps":                    10,
		"proxy_protocol_enabled": true,
		"mss":                    10,
		"sec_sensor_enabled":     true,
		"ca_enabled":             true,
		"end_port":               20,
		"start_port":             10,
		"alpn_policy":            "CreateListenerValue",
		"alpn_enabled":           true,
		"ca_certificate_ids":     []string{"CreateListenerValue"},
		"certificate_ids":        []string{"CreateListenerValue"},
	}
	for key, value := range attributes {
		err := dInit.Set(key, value)
		assert.Nil(t, err)
		err = dExisted.Set(key, value)
		assert.Nil(t, err)
		if err != nil {
			log.Printf("[ERROR] the field %s setting error", key)
		}
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		// GetListenerAttribute
		"CaCertificateIds":     []interface{}{"CreateListenerValue"},
		"CertificateIds":       []interface{}{"CreateListenerValue"},
		"EndPort":              "20",
		"ListenerDescription":  "CreateListenerValue",
		"ListenerId":           "CreateListenerValue",
		"ListenerPort":         10,
		"ListenerProtocol":     "CreateListenerValue",
		"LoadBalancerId":       "CreateListenerValue",
		"ServerGroupId":        "CreateListenerValue",
		"StartPort":            "10",
		"ListenerStatus":       "Running",
		"Cps":                  10,
		"IdleTimeout":          10,
		"Mss":                  10,
		"ProxyProtocolEnabled": true,
		"SecSensorEnabled":     true,
		"CaEnabled":            true,
		"AlpnPolicy":           "CreateListenerValue",
		"AlpnEnabled":          true,
	}
	CreateMockResponse := map[string]interface{}{
		// CreateListener
		"ListenerId": "CreateListenerValue",
	}
	failedResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, &tea.SDKError{
			Code:       String(errorCode),
			Data:       String(errorCode),
			Message:    String(errorCode),
			StatusCode: tea.Int(400),
		}
	}
	notFoundResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_nlb_listener", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNlbClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudNlbListenerCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// GetListenerAttribute Response
		"ListenerId": "CreateListenerValue",
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateListener" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						successResponseMock(ReadMockResponseDiff)
						return CreateMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudNlbListenerCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dInit.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dInit.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Update
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNlbClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudNlbListenerUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// UpdateListenerAttribute
	attributesDiff := map[string]interface{}{
		"alpn_enabled":           false,
		"alpn_policy":            "UpdateListenerAttributeValue",
		"ca_certificate_ids":     []interface{}{"UpdateListenerAttributeValue3"},
		"ca_enabled":             false,
		"certificate_ids":        []interface{}{"UpdateListenerAttributeValue3"},
		"cps":                    15,
		"idle_timeout":           15,
		"listener_description":   "UpdateListenerAttributeValue",
		"mss":                    15,
		"proxy_protocol_enabled": false,
		"sec_sensor_enabled":     false,
		"security_policy_id":     "UpdateListenerAttributeValue",
		"server_group_id":        "UpdateListenerAttributeValue",
	}
	diff, err := newInstanceDiff("alicloud_nlb_listener", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// GetListenerAttribute Response
		"AlpnEnabled": false,
		"AlpnPolicy":  "UpdateListenerAttributeValue",
		"CaCertificateIds": []interface{}{
			"UpdateListenerAttributeValue3",
		},
		"CaEnabled": false,
		"CertificateIds": []interface{}{
			"UpdateListenerAttributeValue3",
		},
		"Cps":                  15,
		"IdleTimeout":          15,
		"ListenerDescription":  "UpdateListenerAttributeValue",
		"Mss":                  15,
		"ProxyProtocolEnabled": false,
		"SecSensorEnabled":     false,
		"SecurityPolicyId":     "UpdateListenerAttributeValue",
		"ServerGroupId":        "UpdateListenerAttributeValue",
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "UpdateListenerAttribute" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudNlbListenerUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// StopListener
	attributesDiff = map[string]interface{}{
		"status": "Stopped",
	}
	diff, err = newInstanceDiff("alicloud_nlb_listener", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// GetListenerAttribute Response
		"ListenerStatus": "Stopped",
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "StopListener" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudNlbListenerUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// StartListener
	attributesDiff = map[string]interface{}{
		"status": "Running",
	}
	diff, err = newInstanceDiff("alicloud_nlb_listener", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// GetListenerAttribute Response
		"ListenerStatus": "Running",
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "StartListener" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudNlbListenerUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_nlb_listener"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Read
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "ResourceNotFound.listener", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "GetListenerAttribute" {
				switch errorCode {
				case "{}", "ResourceNotFound.listener":
					return notFoundResponseMock(errorCode)
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudNlbListenerRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}", "ResourceNotFound.listener":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNlbClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudNlbListenerDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteListener" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						ReadMockResponse = map[string]interface{}{}
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudNlbListenerDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		}
	}
}

// Test Nlb Listener. >>> Resource test cases, automatically generated.
// Case terraform覆盖tcpssl 4683
func TestAccAliCloudNlbListener_basic4683(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap4683)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnlb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependence4683)
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
					"ca_enabled":        "true",
					"load_balancer_id":  "${alicloud_nlb_load_balancer.nlb.id}",
					"server_group_id":   "${alicloud_nlb_server_group.sg1.id}",
					"listener_protocol": "TCPSSL",
					"certificate_ids": []string{
						"${alicloud_ssl_certificates_service_certificate.ssl0.id}"},
					"listener_description":   "test",
					"status":                 "Running",
					"start_port":             "1",
					"cps":                    "1",
					"idle_timeout":           "10",
					"mss":                    "100",
					"end_port":               "65535",
					"proxy_protocol_enabled": "true",
					"sec_sensor_enabled":     "false",
					"listener_port":          "0",
					"security_policy_id":     "tls_cipher_policy_1_0",
					"alpn_enabled":           "true",
					"alpn_policy":            "HTTP1Only",
					"ca_certificate_ids": []string{
						"${alicloud_ssl_certificates_service_pca_certificate.pca.id}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ca_enabled":             "true",
						"load_balancer_id":       CHECKSET,
						"server_group_id":        CHECKSET,
						"listener_protocol":      "TCPSSL",
						"certificate_ids.#":      "1",
						"listener_description":   "test",
						"status":                 "Running",
						"start_port":             CHECKSET,
						"cps":                    "1",
						"idle_timeout":           "10",
						"mss":                    "100",
						"end_port":               CHECKSET,
						"proxy_protocol_enabled": "true",
						"sec_sensor_enabled":     "false",
						"listener_port":          "0",
						"security_policy_id":     "tls_cipher_policy_1_0",
						"alpn_enabled":           "true",
						"alpn_policy":            "HTTP1Only",
						"ca_certificate_ids.#":   "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_id": "${alicloud_nlb_server_group.sg2.id}",
					"certificate_ids": []string{
						"${alicloud_ssl_certificates_service_certificate.ssl1.id}"},
					"listener_description":   "testupdate",
					"cps":                    "0",
					"idle_timeout":           "900",
					"mss":                    "1500",
					"proxy_protocol_enabled": "false",
					"security_policy_id":     "tls_cipher_policy_1_2",
					"alpn_policy":            "HTTP2Preferred",
					"ca_certificate_ids":     []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_id":        CHECKSET,
						"certificate_ids.#":      "1",
						"listener_description":   "testupdate",
						"cps":                    "0",
						"idle_timeout":           "900",
						"mss":                    "1500",
						"proxy_protocol_enabled": "false",
						"security_policy_id":     "tls_cipher_policy_1_2",
						"alpn_policy":            "HTTP2Preferred",
						"ca_certificate_ids.#":   "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ca_enabled":         "false",
					"alpn_enabled":       "false",
					"ca_certificate_ids": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ca_enabled":           "false",
						"alpn_enabled":         "false",
						"ca_certificate_ids.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"certificate_ids":    []string{},
					"ca_certificate_ids": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"certificate_ids.#":    "0",
						"ca_certificate_ids.#": "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"certificate_ids": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"certificate_ids.#": "0",
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
					"sec_sensor_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"sec_sensor_enabled": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ca_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ca_enabled": "true",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNlbListenerMap4683 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudNlbListenerBasicDependence4683(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "nlbvpc"
}

resource "alicloud_vswitch" "vswtich" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1a"
  vswitch_name = "nlbvs"
  cidr_block   = "192.168.1.0/24"
}

resource "alicloud_vswitch" "vswtich2" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1b"
  vswitch_name = "nlbvs2"
  cidr_block   = "192.168.2.0/24"
}

resource "alicloud_nlb_load_balancer" "nlb" {
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich.id
    zone_id    = alicloud_vswitch.vswtich.zone_id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich2.id
    zone_id    = alicloud_vswitch.vswtich2.zone_id
  }
  load_balancer_type = "Network"
  vpc_id             = alicloud_vpc.vpc.id
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  load_balancer_name = "tf-test-nlb"
}

resource "alicloud_nlb_server_group" "sg1" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCPSSL"
  server_group_name = "tf01"
}

resource "alicloud_nlb_server_group" "sg2" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCPSSL"
  server_group_name = "tf02"
}

resource "alicloud_ssl_certificates_service_certificate" "ssl0" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID3zCCAsegAwIBAgIUI2j6Zirw9BU72IE/YMcI7JfzYnAwDQYJKoZIhvcNAQEL
BQAwfzELMAkGA1UEBhMCQWwxEDAOBgNVBAgMB0FsaWJhYmExEDAOBgNVBAcMB0Fs
aWJhYmExEDAOBgNVBAoMB0FsaWJhYmExEDAOBgNVBAsMB0FsaWJhYmExEDAOBgNV
BAMMB0FsaWJhYmExFjAUBgkqhkiG9w0BCQEWB0FsaWJhYmEwHhcNMjUxMjMwMDcy
NjEyWhcNMzUxMjI4MDcyNjEyWjB/MQswCQYDVQQGEwJBbDEQMA4GA1UECAwHQWxp
YmFiYTEQMA4GA1UEBwwHQWxpYmFiYTEQMA4GA1UECgwHQWxpYmFiYTEQMA4GA1UE
CwwHQWxpYmFiYTEQMA4GA1UEAwwHQWxpYmFiYTEWMBQGCSqGSIb3DQEJARYHQWxp
YmFiYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKcODnQT7wrYdEqH
cTcNu6WNI6ij+lJd6LtORQ6UgYP4HwiBukR/GgkhXYvrK5zeCVy9Gu8IaKOG9dyU
k55a7UCC3fXLAbjKx8CeKlZY5YMJvxEhGNwo06iTN6AC8QsUT9QsdiEbzimi9ubZ
a8Jvsu1/c8WT/frZrlkodJQAJ/Cd4WWcBwj3RewyXYn0/LzQNmahaVK3VunjKdBw
7DBzyhOeECxkWaD9wmg7zqZ+cNyghXqip5UAm9Pji1V2xHhDc1P4f37LheWsSKG6
DfS9BwFXl566e5VvzTu0RjjNlQbHfYivahrpvO2hoiTxkRyqVUdf3q31McptKttf
gQgICwECAwEAAaNTMFEwHQYDVR0OBBYEFOTEyExA4rMHLPnTGh6Cyuh0fMhfMB8G
A1UdIwQYMBaAFOTEyExA4rMHLPnTGh6Cyuh0fMhfMA8GA1UdEwEB/wQFMAMBAf8w
DQYJKoZIhvcNAQELBQADggEBAKJDQDedLZYXQxbI7yBV1KKxa6IrzV7OPTZX8e6L
IdHiMCZ7sHYkmanBm1VEDMY4jSoPtkwsK3tng02EStzX6is3Sy/hbsqHwU8H8W0a
gZN7+QEJW2/Gze0wvov2xbvGU1MjJdw7ooJE7LkPPkXJdRCHHbYKEPgjZMcYAuod
nTOQX1c7S1fh0seeKZVjEVDW7ePBvCPHzZM6xevH6kjHPUhpszmvbXIPSl+NEJyL
acTa54urnT7oZT1Pda3PkJSEWy8VB36hQ/ygp4uNEQeF43baL4Vp5GWXr1ninIQr
/mpqDUjv8prPjKmseGU7Be49Ivh9vea+BOnXpFvZwuGf/8Y=
-----END CERTIFICATE-----
EOF
  certificate_name = "tf-test-552"
  key              = <<EOF
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCnDg50E+8K2HRK
h3E3DbuljSOoo/pSXei7TkUOlIGD+B8IgbpEfxoJIV2L6yuc3glcvRrvCGijhvXc
lJOeWu1Agt31ywG4ysfAnipWWOWDCb8RIRjcKNOokzegAvELFE/ULHYhG84povbm
2WvCb7Ltf3PFk/362a5ZKHSUACfwneFlnAcI90XsMl2J9Py80DZmoWlSt1bp4ynQ
cOwwc8oTnhAsZFmg/cJoO86mfnDcoIV6oqeVAJvT44tVdsR4Q3NT+H9+y4XlrEih
ug30vQcBV5eeunuVb807tEY4zZUGx32Ir2oa6bztoaIk8ZEcqlVHX96t9THKbSrb
X4EICAsBAgMBAAECggEAEaMQJ1Tt7z4STAaP70UrxyMXLVAvptfLkQ/mxk7rPBq7
cwdYOjBVmDptsQ6hdoCyKhh8ghlIC+C9Lx5AFg9JPl76rl8yFbeR/o3sUBy3UX37
TkSJZoAIdaMiU6pP1tC8i1zs1VrFSbvO96kjgX4PvNeQil7xKPVnvPnCkhgQn2xt
83OPvILoQM8xefU4o3npjWpIGaSo0z5E5dBqRb3IWb8/o0ruMoTTPWh3xq9/ob6Y
XieQIX4fgOMToTGnfJBdTNTj7p/IUUO9ZoTod7UJ0E6GcHLwiArIoGgwQ70OJgX0
V3pC1lvCGIpEE0BevVaEfnbR2/QS66vyIBkxl2sUeQKBgQDfD/9APGhQiRId3je6
r1wz19feRT/qlbP/qPompKnkeUfZ6nTl5UL4eLhloZnbza8ReL/cnL8Aq/uOqPY/
PXkAAErLFuKiTQPkgm8uwzr/JXFASEm4YnpSrjbNlaq0dD4Z9WUb+lm55bKPGjk2
Tw6kaEUwcljXSXngLRW4ZyreiQKBgQC/uOmQwT2e3KVt9vp+BA22K7EeY74xqCUm
uCxu3HCJWYLSr/expyEyYlTrPrPoYCIsAxwD1JNIAn4jHWvh2lqlFUg3WozABSep
RKnvtrjlCO3Do5ZxJTrayyyTeAG+iP8UBV9fp9dCcNEm4cOuaudynxy9lid2/Qf2
jrwfLITquQKBgQC6aLP3CoRiSSpKq5oG+OUkrgxIQ7bWY6S40o26HYGc3X2eLsDi
HmWJH9l5lULu3USgZThxNMyp0w+2eJzQ3J5x0cyvS8C5wYKvRBGGEsEK2E98WOzC
mgl/bvd9dsXhaAM9zkVgzCaPo9VEobWOHSMvYbPweJ6ly1F+di7gm9bHgQKBgFKy
A6y4bUfUjFZ+LVWlkfR9rAwbobHxgPTOg/vdgXz86vzNgd1S5XLCDzbY+OjGbnbl
cHQGgPCSgq3KxvnIIUkLgPa/S+6XSqAsSdBn1CCWVVgJe56aHGS7NiL1pGr21G9i
Ud0EnEjAOLa4sprM6b/6+X6dAbDFO2YR8vrPTRWhAoGAao1jyluslKpL4LcjEImd
nNzHGosZ9liKecWZHGE84r52BNOX+O130HOStDYlKYPUCUJuO5dDpZSjBIVrR2CO
5WHaQin/p+S0akWYSowIjR12igUIxEv4GJq9E9prTZhOWN/BB1t1Q89rWir5XlmD
D9RiDhntjbRzI7zYd5QRWGk=
-----END PRIVATE KEY-----
EOF
}

resource "alicloud_ssl_certificates_service_certificate" "ssl1" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIID3zCCAsegAwIBAgIUI2j6Zirw9BU72IE/YMcI7JfzYnAwDQYJKoZIhvcNAQEL
BQAwfzELMAkGA1UEBhMCQWwxEDAOBgNVBAgMB0FsaWJhYmExEDAOBgNVBAcMB0Fs
aWJhYmExEDAOBgNVBAoMB0FsaWJhYmExEDAOBgNVBAsMB0FsaWJhYmExEDAOBgNV
BAMMB0FsaWJhYmExFjAUBgkqhkiG9w0BCQEWB0FsaWJhYmEwHhcNMjUxMjMwMDcy
NjEyWhcNMzUxMjI4MDcyNjEyWjB/MQswCQYDVQQGEwJBbDEQMA4GA1UECAwHQWxp
YmFiYTEQMA4GA1UEBwwHQWxpYmFiYTEQMA4GA1UECgwHQWxpYmFiYTEQMA4GA1UE
CwwHQWxpYmFiYTEQMA4GA1UEAwwHQWxpYmFiYTEWMBQGCSqGSIb3DQEJARYHQWxp
YmFiYTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKcODnQT7wrYdEqH
cTcNu6WNI6ij+lJd6LtORQ6UgYP4HwiBukR/GgkhXYvrK5zeCVy9Gu8IaKOG9dyU
k55a7UCC3fXLAbjKx8CeKlZY5YMJvxEhGNwo06iTN6AC8QsUT9QsdiEbzimi9ubZ
a8Jvsu1/c8WT/frZrlkodJQAJ/Cd4WWcBwj3RewyXYn0/LzQNmahaVK3VunjKdBw
7DBzyhOeECxkWaD9wmg7zqZ+cNyghXqip5UAm9Pji1V2xHhDc1P4f37LheWsSKG6
DfS9BwFXl566e5VvzTu0RjjNlQbHfYivahrpvO2hoiTxkRyqVUdf3q31McptKttf
gQgICwECAwEAAaNTMFEwHQYDVR0OBBYEFOTEyExA4rMHLPnTGh6Cyuh0fMhfMB8G
A1UdIwQYMBaAFOTEyExA4rMHLPnTGh6Cyuh0fMhfMA8GA1UdEwEB/wQFMAMBAf8w
DQYJKoZIhvcNAQELBQADggEBAKJDQDedLZYXQxbI7yBV1KKxa6IrzV7OPTZX8e6L
IdHiMCZ7sHYkmanBm1VEDMY4jSoPtkwsK3tng02EStzX6is3Sy/hbsqHwU8H8W0a
gZN7+QEJW2/Gze0wvov2xbvGU1MjJdw7ooJE7LkPPkXJdRCHHbYKEPgjZMcYAuod
nTOQX1c7S1fh0seeKZVjEVDW7ePBvCPHzZM6xevH6kjHPUhpszmvbXIPSl+NEJyL
acTa54urnT7oZT1Pda3PkJSEWy8VB36hQ/ygp4uNEQeF43baL4Vp5GWXr1ninIQr
/mpqDUjv8prPjKmseGU7Be49Ivh9vea+BOnXpFvZwuGf/8Y=
-----END CERTIFICATE-----
EOF
  certificate_name = "tf-test-50"
  key              = <<EOF
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCnDg50E+8K2HRK
h3E3DbuljSOoo/pSXei7TkUOlIGD+B8IgbpEfxoJIV2L6yuc3glcvRrvCGijhvXc
lJOeWu1Agt31ywG4ysfAnipWWOWDCb8RIRjcKNOokzegAvELFE/ULHYhG84povbm
2WvCb7Ltf3PFk/362a5ZKHSUACfwneFlnAcI90XsMl2J9Py80DZmoWlSt1bp4ynQ
cOwwc8oTnhAsZFmg/cJoO86mfnDcoIV6oqeVAJvT44tVdsR4Q3NT+H9+y4XlrEih
ug30vQcBV5eeunuVb807tEY4zZUGx32Ir2oa6bztoaIk8ZEcqlVHX96t9THKbSrb
X4EICAsBAgMBAAECggEAEaMQJ1Tt7z4STAaP70UrxyMXLVAvptfLkQ/mxk7rPBq7
cwdYOjBVmDptsQ6hdoCyKhh8ghlIC+C9Lx5AFg9JPl76rl8yFbeR/o3sUBy3UX37
TkSJZoAIdaMiU6pP1tC8i1zs1VrFSbvO96kjgX4PvNeQil7xKPVnvPnCkhgQn2xt
83OPvILoQM8xefU4o3npjWpIGaSo0z5E5dBqRb3IWb8/o0ruMoTTPWh3xq9/ob6Y
XieQIX4fgOMToTGnfJBdTNTj7p/IUUO9ZoTod7UJ0E6GcHLwiArIoGgwQ70OJgX0
V3pC1lvCGIpEE0BevVaEfnbR2/QS66vyIBkxl2sUeQKBgQDfD/9APGhQiRId3je6
r1wz19feRT/qlbP/qPompKnkeUfZ6nTl5UL4eLhloZnbza8ReL/cnL8Aq/uOqPY/
PXkAAErLFuKiTQPkgm8uwzr/JXFASEm4YnpSrjbNlaq0dD4Z9WUb+lm55bKPGjk2
Tw6kaEUwcljXSXngLRW4ZyreiQKBgQC/uOmQwT2e3KVt9vp+BA22K7EeY74xqCUm
uCxu3HCJWYLSr/expyEyYlTrPrPoYCIsAxwD1JNIAn4jHWvh2lqlFUg3WozABSep
RKnvtrjlCO3Do5ZxJTrayyyTeAG+iP8UBV9fp9dCcNEm4cOuaudynxy9lid2/Qf2
jrwfLITquQKBgQC6aLP3CoRiSSpKq5oG+OUkrgxIQ7bWY6S40o26HYGc3X2eLsDi
HmWJH9l5lULu3USgZThxNMyp0w+2eJzQ3J5x0cyvS8C5wYKvRBGGEsEK2E98WOzC
mgl/bvd9dsXhaAM9zkVgzCaPo9VEobWOHSMvYbPweJ6ly1F+di7gm9bHgQKBgFKy
A6y4bUfUjFZ+LVWlkfR9rAwbobHxgPTOg/vdgXz86vzNgd1S5XLCDzbY+OjGbnbl
cHQGgPCSgq3KxvnIIUkLgPa/S+6XSqAsSdBn1CCWVVgJe56aHGS7NiL1pGr21G9i
Ud0EnEjAOLa4sprM6b/6+X6dAbDFO2YR8vrPTRWhAoGAao1jyluslKpL4LcjEImd
nNzHGosZ9liKecWZHGE84r52BNOX+O130HOStDYlKYPUCUJuO5dDpZSjBIVrR2CO
5WHaQin/p+S0akWYSowIjR12igUIxEv4GJq9E9prTZhOWN/BB1t1Q89rWir5XlmD
D9RiDhntjbRzI7zYd5QRWGk=
-----END PRIVATE KEY-----
EOF
}

resource "alicloud_ssl_certificates_service_pca_certificate" "pca" {
  organization      = "Alibaba"
  years             = "10"
  locality          = "Alibaba"
  organization_unit = "Alibaba"
  state             = "Alibaba"
  country_code      = "Alibaba"
  common_name       = "Alibaba"
  algorithm         = "RSA_1024"
}

resource "alicloud_nlb_hd_monitor_region_config" "HDMonitor" {
  metric_store = "eu-central-1-secsensor"
  log_project  = "eu-central-1-secsensor"
}


`, name)
}

// Case terraform-multi-rsp 12231
func TestAccAliCloudNlbListener_basic12231(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap12231)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnlb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependence12231)
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
					"ca_enabled":             "false",
					"load_balancer_id":       "${alicloud_nlb_load_balancer.nlb.id}",
					"listener_protocol":      "TCP",
					"listener_description":   "test",
					"status":                 "Running",
					"start_port":             "1",
					"cps":                    "1",
					"idle_timeout":           "10",
					"mss":                    "100",
					"end_port":               "65535",
					"proxy_protocol_enabled": "true",
					"sec_sensor_enabled":     "false",
					"listener_port":          "0",
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "true",
							"proxy_protocol_config_private_link_eps_id_enabled": "true",
						},
					},
					"server_group_tuples": []map[string]interface{}{
						{
							"server_group_id": "${alicloud_nlb_server_group.sg1.id}",
							"weight":          "100",
						},
						{
							"server_group_id": "${alicloud_nlb_server_group.sg2.id}",
							"weight":          "1000",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ca_enabled":             "false",
						"load_balancer_id":       CHECKSET,
						"listener_protocol":      "TCP",
						"listener_description":   "test",
						"status":                 "Running",
						"start_port":             CHECKSET,
						"cps":                    "1",
						"idle_timeout":           "10",
						"mss":                    "100",
						"end_port":               CHECKSET,
						"proxy_protocol_enabled": "true",
						"sec_sensor_enabled":     "false",
						"listener_port":          "0",
						"server_group_tuples.#":  "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_tuples": []map[string]interface{}{
						{
							"server_group_id": "${alicloud_nlb_server_group.sg1.id}",
							"weight":          "1000",
						},
						{
							"server_group_id": "${alicloud_nlb_server_group.sg2.id}",
							"weight":          "100",
						},
						{
							"server_group_id": "${alicloud_nlb_server_group.sg3.id}",
							"weight":          "1000",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_tuples.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_tuples": []map[string]interface{}{
						{
							"server_group_id": "${alicloud_nlb_server_group.sg2.id}",
							"weight":          "100",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_tuples.#": "1",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNlbListenerMap12231 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudNlbListenerBasicDependence12231(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "nlbvpc"
}

resource "alicloud_vswitch" "vswtich" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1a"
  vswitch_name = "nlbvs"
  cidr_block   = "192.168.1.0/24"
}

resource "alicloud_vswitch" "vswtich2" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1b"
  vswitch_name = "nlbvs2"
  cidr_block   = "192.168.2.0/24"
}

resource "alicloud_nlb_load_balancer" "nlb" {
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich.id
    zone_id    = alicloud_vswitch.vswtich.zone_id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich2.id
    zone_id    = alicloud_vswitch.vswtich2.zone_id
  }
  load_balancer_type = "Network"
  vpc_id             = alicloud_vpc.vpc.id
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  load_balancer_name = "tf-test-nlb"
}

resource "alicloud_nlb_server_group" "sg1" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCP"
  server_group_name = "tf01"
}

resource "alicloud_nlb_server_group" "sg2" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCP"
  server_group_name = "tf02"
}

resource "alicloud_nlb_server_group" "sg3" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCP"
  server_group_name = "tf03"
}


`, name)
}

// Case terraform覆盖udp 4673
func TestAccAliCloudNlbListener_basic4673(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap4673)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnlb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependence4673)
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
					"ca_enabled":           "false",
					"load_balancer_id":     "${alicloud_nlb_load_balancer.nlb.id}",
					"server_group_id":      "${alicloud_nlb_server_group.sg1.id}",
					"listener_protocol":    "UDP",
					"listener_description": "test",
					"listener_port":        "100",
					"status":               "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ca_enabled":           "false",
						"load_balancer_id":     CHECKSET,
						"server_group_id":      CHECKSET,
						"listener_protocol":    "UDP",
						"listener_description": "test",
						"listener_port":        "100",
						"status":               "Running",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_id":      "${alicloud_nlb_server_group.sg2.id}",
					"listener_description": "testupdate",
					"sec_sensor_enabled":   "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_id":      CHECKSET,
						"listener_description": "testupdate",
						"sec_sensor_enabled":   "false",
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
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_ep_id_enabled":  "false",
							"proxy_protocol_config_vpc_id_enabled":              "false",
							"proxy_protocol_config_private_link_eps_id_enabled": "false",
						},
					},
					"proxy_protocol_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_protocol_enabled": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "true",
							"proxy_protocol_config_private_link_eps_id_enabled": "true",
						},
					},
					"proxy_protocol_enabled": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"proxy_protocol_enabled": "true",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNlbListenerMap4673 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudNlbListenerBasicDependence4673(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "nlbvpc"
}

resource "alicloud_vswitch" "vswtich" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1a"
  vswitch_name = "nlbvs"
  cidr_block   = "192.168.1.0/24"
}

resource "alicloud_vswitch" "vswtich2" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1b"
  vswitch_name = "nlbvs2"
  cidr_block   = "192.168.2.0/24"
}

resource "alicloud_nlb_load_balancer" "nlb" {
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich.id
    zone_id    = alicloud_vswitch.vswtich.zone_id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich2.id
    zone_id    = alicloud_vswitch.vswtich2.zone_id
  }
  load_balancer_type = "Network"
  vpc_id             = alicloud_vpc.vpc.id
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  load_balancer_name = "tf-test-nlb"
}

resource "alicloud_nlb_server_group" "sg1" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = false
  protocol          = "UDP"
  server_group_name = "tf01"
}

resource "alicloud_nlb_server_group" "sg2" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = false
  protocol          = "UDP"
  server_group_name = "tf02"
}


`, name)
}

// Case terraform覆盖tcp 4675
func TestAccAliCloudNlbListener_basic4675(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nlb_listener.default"
	ra := resourceAttrInit(resourceId, AlicloudNlbListenerMap4675)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNlbListener")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnlb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNlbListenerBasicDependence4675)
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
					"ca_enabled":             "false",
					"load_balancer_id":       "${alicloud_nlb_load_balancer.nlb.id}",
					"server_group_id":        "${alicloud_nlb_server_group.sg1.id}",
					"listener_protocol":      "TCP",
					"listener_description":   "test",
					"status":                 "Running",
					"start_port":             "1",
					"cps":                    "1",
					"idle_timeout":           "10",
					"mss":                    "100",
					"end_port":               "65535",
					"proxy_protocol_enabled": "true",
					"sec_sensor_enabled":     "false",
					"listener_port":          "0",
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "true",
							"proxy_protocol_config_private_link_eps_id_enabled": "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ca_enabled":             "false",
						"load_balancer_id":       CHECKSET,
						"server_group_id":        CHECKSET,
						"listener_protocol":      "TCP",
						"listener_description":   "test",
						"status":                 "Running",
						"start_port":             CHECKSET,
						"cps":                    "1",
						"idle_timeout":           "10",
						"mss":                    "100",
						"end_port":               CHECKSET,
						"proxy_protocol_enabled": "true",
						"sec_sensor_enabled":     "false",
						"listener_port":          "0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"server_group_id":      "${alicloud_nlb_server_group.sg2.id}",
					"listener_description": "testupdate",
					"cps":                  "0",
					"idle_timeout":         "900",
					"mss":                  "1500",
					"proxy_protocol_config": []map[string]interface{}{
						{
							"proxy_protocol_config_private_link_ep_id_enabled":  "true",
							"proxy_protocol_config_vpc_id_enabled":              "false",
							"proxy_protocol_config_private_link_eps_id_enabled": "false",
						},
					},
					"alpn_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_group_id":      CHECKSET,
						"listener_description": "testupdate",
						"cps":                  "0",
						"idle_timeout":         "900",
						"mss":                  "1500",
						"alpn_enabled":         "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":                 "Stopped",
					"proxy_protocol_enabled": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":                 "Stopped",
						"proxy_protocol_enabled": "false",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNlbListenerMap4675 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudNlbListenerBasicDependence4675(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "nlbvpc"
}

resource "alicloud_vswitch" "vswtich" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1a"
  vswitch_name = "nlbvs"
  cidr_block   = "192.168.1.0/24"
}

resource "alicloud_vswitch" "vswtich2" {
  vpc_id       = alicloud_vpc.vpc.id
  zone_id      = "eu-central-1b"
  vswitch_name = "nlbvs2"
  cidr_block   = "192.168.2.0/24"
}

resource "alicloud_nlb_load_balancer" "nlb" {
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich.id
    zone_id    = alicloud_vswitch.vswtich.zone_id
  }
  zone_mappings {
    vswitch_id = alicloud_vswitch.vswtich2.id
    zone_id    = alicloud_vswitch.vswtich2.zone_id
  }
  load_balancer_type = "Network"
  vpc_id             = alicloud_vpc.vpc.id
  address_type       = "Internet"
  address_ip_version = "Ipv4"
  load_balancer_name = "tf-test-nlb"
}

resource "alicloud_nlb_server_group" "sg1" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCP"
  server_group_name = "tf01"
}

resource "alicloud_nlb_server_group" "sg2" {
  scheduler = "Wrr"
  health_check {
    health_check_enabled = false
  }
  server_group_type = "Instance"
  vpc_id            = alicloud_vpc.vpc.id
  any_port_enabled  = true
  protocol          = "TCP"
  server_group_name = "tf02"
}


`, name)
}

// Test Nlb Listener. <<< Resource test cases, automatically generated.
