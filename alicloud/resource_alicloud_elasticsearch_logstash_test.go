// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Elasticsearch Logstash. >>> Resource test cases, automatically generated.
// Case Logstash横向接入Terraform_副本1677130071168 2582
func TestAccAliCloudElasticsearchLogstash_basic2582(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_elasticsearch_logstash.default"
	ra := resourceAttrInit(resourceId, AlicloudElasticsearchLogstashMap2582)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ElasticsearchServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeElasticsearchLogstash")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccelasticsearch%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudElasticsearchLogstashBasicDependence2582)
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
					"description":       "tf-acc-create-test-1",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"version":           "7.4_with_X-Pack",
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_efficiency",
							"spec":      "elasticsearch.sn1ne.large",
							"disk":      "20",
						},
					},
					"network_config": []map[string]interface{}{
						{
							"type":       "vpc",
							"vpc_id":     "vpc-bp1jy348ibzulk6hn65xf",
							"vswitch_id": "vsw-bp13kz5zhn6flmqmh9fyn",
							"vs_area":    "cn-hangzhou-i",
						},
					},
					"payment_type": "Subscription",
					"node_amount":  "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "tf-acc-create-test-1",
						"resource_group_id": CHECKSET,
						"version":           "7.4_with_X-Pack",
						"payment_type":      "Subscription",
						"node_amount":       "1",
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
				ImportStateVerifyIgnore: []string{"payment_info"},
			},
		},
	})
}

var AlicloudElasticsearchLogstashMap2582 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"updated_at":  CHECKSET,
}

func AlicloudElasticsearchLogstashBasicDependence2582(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case Logstash横向接入Terraform 2543
func TestAccAliCloudElasticsearchLogstash_basic2543(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_elasticsearch_logstash.default"
	ra := resourceAttrInit(resourceId, AlicloudElasticsearchLogstashMap2543)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ElasticsearchServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeElasticsearchLogstash")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccelasticsearch%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudElasticsearchLogstashBasicDependence2543)
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
					"description":       "tf-acc-create-test-1",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"version":           "7.4_with_X-Pack",
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_efficiency",
							"spec":      "elasticsearch.sn1ne.large",
							"disk":      "20",
						},
					},
					"network_config": []map[string]interface{}{
						{
							"type":       "vpc",
							"vpc_id":     "vpc-bp1jy348ibzulk6hn65xf",
							"vswitch_id": "vsw-bp13kz5zhn6flmqmh9fyn",
							"vs_area":    "cn-hangzhou-i",
						},
					},
					"payment_type": "Subscription",
					"node_amount":  "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "tf-acc-create-test-1",
						"resource_group_id": CHECKSET,
						"version":           "7.4_with_X-Pack",
						"payment_type":      "Subscription",
						"node_amount":       "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_efficiency",
							"spec":      "elasticsearch.sn1ne.large",
							"disk":      "30",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
					"description": "tf-test-updatedescription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "tf-test-updatedescription",
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
				ImportStateVerifyIgnore: []string{"payment_info"},
			},
		},
	})
}

var AlicloudElasticsearchLogstashMap2543 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"updated_at":  CHECKSET,
}

func AlicloudElasticsearchLogstashBasicDependence2543(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case Logstash横向接入04 2718
func TestAccAliCloudElasticsearchLogstash_basic2718(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_elasticsearch_logstash.default"
	ra := resourceAttrInit(resourceId, AlicloudElasticsearchLogstashMap2718)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ElasticsearchServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeElasticsearchLogstash")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccelasticsearch%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudElasticsearchLogstashBasicDependence2718)
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
					"description":       "tf-acc-create-test-1",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"version":           "7.4_with_X-Pack",
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_efficiency",
							"spec":      "elasticsearch.sn1ne.large",
							"disk":      "20",
						},
					},
					"network_config": []map[string]interface{}{
						{
							"type":       "vpc",
							"vpc_id":     "vpc-bp1jy348ibzulk6hn65xf",
							"vswitch_id": "vsw-bp13kz5zhn6flmqmh9fyn",
							"vs_area":    "cn-hangzhou-i",
						},
					},
					"payment_type": "Subscription",
					"node_amount":  "1",
					"payment_info": []map[string]interface{}{
						{
							"duration":            "1",
							"pricing_cycle":       "Month",
							"auto_renew":          "true",
							"auto_renew_duration": "1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "tf-acc-create-test-1",
						"resource_group_id": CHECKSET,
						"version":           "7.4_with_X-Pack",
						"payment_type":      "Subscription",
						"node_amount":       "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_ssd",
							"spec":      "elasticsearch.sn2ne.large",
							"disk":      "30",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
					"description": "tf-test-updatedescription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "tf-test-updatedescription",
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
				ImportStateVerifyIgnore: []string{"payment_info"},
			},
		},
	})
}

var AlicloudElasticsearchLogstashMap2718 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"updated_at":  CHECKSET,
}

func AlicloudElasticsearchLogstashBasicDependence2718(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case Logstash横向接入05_副本1678086683429 2744
func TestAccAliCloudElasticsearchLogstash_basic2744(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_elasticsearch_logstash.default"
	ra := resourceAttrInit(resourceId, AlicloudElasticsearchLogstashMap2744)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ElasticsearchServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeElasticsearchLogstash")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccelasticsearch%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudElasticsearchLogstashBasicDependence2744)
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
					"description":       "tf-acc-create-test-1",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"version":           "7.4_with_X-Pack",
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_efficiency",
							"spec":      "elasticsearch.sn1ne.large",
							"disk":      "20",
						},
					},
					"network_config": []map[string]interface{}{
						{
							"type":       "vpc",
							"vpc_id":     "vpc-bp1jy348ibzulk6hn65xf",
							"vswitch_id": "vsw-bp13kz5zhn6flmqmh9fyn",
							"vs_area":    "cn-hangzhou-i",
						},
					},
					"payment_type": "PayAsYouGo",
					"node_amount":  "1",
					"payment_info": []map[string]interface{}{
						{
							"auto_renew": "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "tf-acc-create-test-1",
						"resource_group_id": CHECKSET,
						"version":           "7.4_with_X-Pack",
						"payment_type":      "PayAsYouGo",
						"node_amount":       "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_ssd",
							"spec":      "elasticsearch.sn2ne.large",
							"disk":      "30",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
					"description": "tf-test-updatedescription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "tf-test-updatedescription",
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
				ImportStateVerifyIgnore: []string{"payment_info"},
			},
		},
	})
}

var AlicloudElasticsearchLogstashMap2744 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"updated_at":  CHECKSET,
}

func AlicloudElasticsearchLogstashBasicDependence2744(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case Logstash_换账号重入 3472
func TestAccAliCloudElasticsearchLogstash_basic3472(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_elasticsearch_logstash.default"
	ra := resourceAttrInit(resourceId, AlicloudElasticsearchLogstashMap3472)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ElasticsearchServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeElasticsearchLogstash")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccelasticsearch%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudElasticsearchLogstashBasicDependence3472)
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
					"description":       "tf-acc-create-test-1",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"version":           "7.4_with_X-Pack",
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_efficiency",
							"spec":      "elasticsearch.sn1ne.large",
							"disk":      "20",
						},
					},
					"network_config": []map[string]interface{}{
						{
							"type":       "vpc",
							"vpc_id":     "${alicloud_vpc.defaultZFtcRh.id}",
							"vswitch_id": "${alicloud_vswitch.defaultMiMSn6.id}",
							"vs_area":    "cn-hangzhou-i",
						},
					},
					"payment_type": "PayAsYouGo",
					"node_amount":  "1",
					"payment_info": []map[string]interface{}{
						{
							"auto_renew": "false",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":       "tf-acc-create-test-1",
						"resource_group_id": CHECKSET,
						"version":           "7.4_with_X-Pack",
						"payment_type":      "PayAsYouGo",
						"node_amount":       "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"node_spec": []map[string]interface{}{
						{
							"disk_type": "cloud_ssd",
							"spec":      "elasticsearch.sn2ne.large",
							"disk":      "30",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
					"description": "tf-test-updatedescription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "tf-test-updatedescription",
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
				ImportStateVerifyIgnore: []string{"payment_info"},
			},
		},
	})
}

var AlicloudElasticsearchLogstashMap3472 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
	"updated_at":  CHECKSET,
}

func AlicloudElasticsearchLogstashBasicDependence3472(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultZFtcRh" {
  classic_link_enabled = false
  dry_run              = false
  enable_ipv6          = false
  vpc_name             = "rdktest01"
  cidr_block           = "10.0.0.0/8"
}

resource "alicloud_vswitch" "defaultMiMSn6" {
  vpc_id       = alicloud_vpc.defaultZFtcRh.id
  cidr_block   = "10.0.10.0/24"
  vswitch_name = "rdktest01"
  zone_id      = "cn-hangzhou-i"
}


`, name)
}

// Test Elasticsearch Logstash. <<< Resource test cases, automatically generated.
