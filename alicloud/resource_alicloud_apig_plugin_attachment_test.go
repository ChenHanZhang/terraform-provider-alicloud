// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Apig PluginAttachment. >>> Resource test cases, automatically generated.
// Case TestAttachGateway 8724
func TestAccAliCloudApigPluginAttachment_basic8724(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginAttachmentMap8724)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPluginAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginAttachmentBasicDependence8724)
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
					"attach_resource_ids": []string{
						"${alicloud_apig_gateway.defaultgateway.id}"},
					"enable": "false",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_id":     "${alicloud_apig_plugin.defaultplugin.id}",
							"gateway_id":    "${alicloud_apig_gateway.defaultgateway.id}",
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
					"attach_resource_type": "Gateway",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "1",
						"enable":                "false",
						"attach_resource_type":  "Gateway",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"attach_resource_ids": []string{
						"${alicloud_apig_plugin.defaultplugin.gateway_id}"},
					"enable": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "1",
						"enable":                "true",
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

var AlicloudApigPluginAttachmentMap8724 = map[string]string{}

func AlicloudApigPluginAttachmentBasicDependence8724(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-test"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.26.0/24"
  vswitch_name = "zhenyuan-test"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = alicloud_vpc.defaultvpc.vpc_name
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "test"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin" "defaultplugin" {
  plugin_class_id = "pls-cqebrgh4ckt6ppatmprc"
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
}


`, name)
}

// Case TestAttachApi 8779
func TestAccAliCloudApigPluginAttachment_basic8779(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginAttachmentMap8779)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPluginAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginAttachmentBasicDependence8779)
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
					"attach_resource_ids": []string{
						"${alicloud_apig_http_api.defaultapi.id}"},
					"enable": "false",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_id":     "${alicloud_apig_plugin.defaultplugin.id}",
							"gateway_id":    "${alicloud_apig_gateway.defaultgateway.id}",
							"plugin_config": "YWxsb3dfb3JpZ2luX3BhdHRlcm5zOgogIC0gJyonCmFsbG93X21ldGhvZHM6CiAgLSAnKicgIAphbGxvd19oZWFkZXJzOgogIC0gJyonCmV4cG9zZV9oZWFkZXJzOgogIC0gJyonCmFsbG93X2NyZWRlbnRpYWxzOiB0cnVlCm1heF9hZ2U6IDcyMDA=",
						},
					},
					"attach_resource_type": "HttpApi",
					"environment_id":       "${alicloud_apig_environment.defaultenvironment.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "1",
						"enable":                "false",
						"attach_resource_type":  "HttpApi",
						"environment_id":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable": "true",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable": "true",
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

var AlicloudApigPluginAttachmentMap8779 = map[string]string{}

func AlicloudApigPluginAttachmentBasicDependence8779(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-test"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.15.0/24"
  vswitch_name = "zhenyuan-test"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = alicloud_vpc.defaultvpc.vpc_name
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "test"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_environment" "defaultenvironment" {
  description      = "测试环境"
  environment_name = "test-env-2"
  gateway_id       = alicloud_apig_gateway.defaultgateway.id
}

resource "alicloud_apig_plugin" "defaultplugin" {
  plugin_class_id = "pls-cqebrgh4ckt6ppatmprc"
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
}

resource "alicloud_apig_http_api" "defaultapi" {
  http_api_name = "zhenyuan-test"
  protocols     = ["HTTP"]
  type          = "Rest"
  description   = "zhenyuan test api"
  base_path     = "/v1"
}


`, name)
}

// Case TestAttachGatewayDomain 8783
func TestAccAliCloudApigPluginAttachment_basic8783(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginAttachmentMap8783)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPluginAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginAttachmentBasicDependence8783)
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
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain1.id}", "${alicloud_apig_domain.defaultdomain2.id}", "${alicloud_apig_domain.defaultdomain3.id}"},
					"enable": "false",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_id":     "${alicloud_apig_plugin.defaultplugin.id}",
							"gateway_id":    "${alicloud_apig_gateway.defaultgateway.id}",
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQ=",
						},
					},
					"attach_resource_type": "GatewayDomain",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "3",
						"enable":                "false",
						"attach_resource_type":  "GatewayDomain",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain1.id}", "${alicloud_apig_domain.defaultdomain2.id}"},
					"enable": "true",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "2",
						"enable":                "true",
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

var AlicloudApigPluginAttachmentMap8783 = map[string]string{}

func AlicloudApigPluginAttachmentBasicDependence8783(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-test"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.15.0/24"
  vswitch_name = "zhenyuan-test"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = alicloud_vpc.defaultvpc.vpc_name
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "test"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin" "defaultplugin" {
  plugin_class_id = "pls-cqebrgh4ckt6ppatmprc"
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
}

resource "alicloud_apig_domain" "defaultdomain1" {
  domain_name = "test-1.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_domain" "defaultdomain2" {
  domain_name = "test-2.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_domain" "defaultdomain3" {
  domain_name = "test-3.com"
  protocol    = "HTTP"
}


`, name)
}

// Case TestAttachGatewayDomain_测试 8961
func TestAccAliCloudApigPluginAttachment_basic8961(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginAttachmentMap8961)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPluginAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginAttachmentBasicDependence8961)
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
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain1.id}", "${alicloud_apig_domain.defaultdomain2.id}"},
					"enable": "false",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_id":     "${alicloud_apig_plugin.defaultplugin.id}",
							"gateway_id":    "${alicloud_apig_gateway.defaultgateway.id}",
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQ=",
						},
					},
					"attach_resource_type": "GatewayDomain",
					"environment_id":       "${alicloud_apig_environment.defaultB0iva8.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "2",
						"enable":                "false",
						"attach_resource_type":  "GatewayDomain",
						"environment_id":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain1.id}", "${alicloud_apig_domain.defaultdomain2.id}", "${alicloud_apig_domain.defaultdomain3.id}"},
					"enable": "true",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "3",
						"enable":                "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain3.id}", "${alicloud_apig_domain.defaultdomain1.id}"},
					"plugin_info": []map[string]interface{}{
						{
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "2",
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

var AlicloudApigPluginAttachmentMap8961 = map[string]string{}

func AlicloudApigPluginAttachmentBasicDependence8961(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-test"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.15.0/24"
  vswitch_name = "zhenyuan-test"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = alicloud_vpc.defaultvpc.vpc_name
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "test"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin" "defaultplugin" {
  plugin_class_id = "pls-cqebrgh4ckt6ppatmprc"
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
}

resource "alicloud_apig_domain" "defaultdomain1" {
  domain_name = "test-1.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_domain" "defaultdomain2" {
  domain_name = "test-2.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_domain" "defaultdomain3" {
  domain_name = "test-3.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_environment" "defaultB0iva8" {
  description      = "测试环境"
  environment_name = "test-env-2"
  gateway_id       = alicloud_apig_gateway.defaultgateway.id
}


`, name)
}

// Case TestAttachGatewayDomain 9275
func TestAccAliCloudApigPluginAttachment_basic9275(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_apig_plugin_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudApigPluginAttachmentMap9275)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ApigServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeApigPluginAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccapig%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudApigPluginAttachmentBasicDependence9275)
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
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain1.id}", "${alicloud_apig_domain.defaultdomain2.id}"},
					"enable": "false",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_id":     "${alicloud_apig_plugin.defaultplugin.id}",
							"gateway_id":    "${alicloud_apig_gateway.defaultgateway.id}",
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQ=",
						},
					},
					"attach_resource_type": "GatewayDomain",
					"environment_id":       "${alicloud_apig_environment.defaultB0iva8.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "2",
						"enable":                "false",
						"attach_resource_type":  "GatewayDomain",
						"environment_id":        CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain1.id}", "${alicloud_apig_domain.defaultdomain2.id}", "${alicloud_apig_domain.defaultdomain3.id}"},
					"enable": "true",
					"plugin_info": []map[string]interface{}{
						{
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "3",
						"enable":                "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"attach_resource_ids": []string{
						"${alicloud_apig_domain.defaultdomain3.id}", "${alicloud_apig_domain.defaultdomain1.id}"},
					"plugin_info": []map[string]interface{}{
						{
							"plugin_config": "c3RhdHVzX2NvZGU6IDIwMApoZWFkZXJzOgotIENvbnRlbnQtVHlwZT1hcHBsaWNhdGlvbi9qc29uCi0gSGVsbG89V29ybGQKYm9keTogIntcImhlbGxvXCI6XCJ3b3JsZFwifSI=",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"attach_resource_ids.#": "2",
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

var AlicloudApigPluginAttachmentMap9275 = map[string]string{}

func AlicloudApigPluginAttachmentBasicDependence9275(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultvpc" {
  cidr_block = "192.168.0.0/16"
  vpc_name   = "zhenyuan-test"
}

resource "alicloud_vswitch" "defaultvswitch" {
  vpc_id       = alicloud_vpc.defaultvpc.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "192.168.15.0/24"
  vswitch_name = "zhenyuan-test"
}

resource "alicloud_apig_gateway" "defaultgateway" {
  network_access_config {
    type = "Intranet"
  }
  vswitch {
    name       = alicloud_vswitch.defaultvswitch.vswitch_name
    vswitch_id = alicloud_vswitch.defaultvswitch.id
  }
  zone_config {
    select_option = "Auto"
  }
  vpc {
    name   = alicloud_vpc.defaultvpc.vpc_name
    vpc_id = alicloud_vpc.defaultvpc.id
  }
  payment_type = "PayAsYouGo"
  gateway_name = "test"
  spec         = "apigw.small.x1"
}

resource "alicloud_apig_plugin" "defaultplugin" {
  plugin_class_id = "pls-cqebrgh4ckt6ppatmprc"
  gateway_id      = alicloud_apig_gateway.defaultgateway.id
}

resource "alicloud_apig_domain" "defaultdomain1" {
  domain_name = "test-1.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_domain" "defaultdomain2" {
  domain_name = "test-2.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_domain" "defaultdomain3" {
  domain_name = "test-3.com"
  protocol    = "HTTP"
}

resource "alicloud_apig_environment" "defaultB0iva8" {
  description      = "测试环境"
  environment_name = "test-env-2"
  gateway_id       = alicloud_apig_gateway.defaultgateway.id
}


`, name)
}

// Test Apig PluginAttachment. <<< Resource test cases, automatically generated.
