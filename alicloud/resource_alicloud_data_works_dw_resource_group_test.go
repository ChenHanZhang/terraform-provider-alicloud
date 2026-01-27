package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test DataWorks DwResourceGroup. >>> Resource test cases, automatically generated.
// Case Dataworks资源组管理_Year_依赖资源自动化创建 8379
func TestAccAliCloudDataWorksDwResourceGroup_basic8379(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_dw_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksDwResourceGroupMap8379)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksDwResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksDwResourceGroupBasicDependence8379)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":          "PostPaid",
					"default_vpc_id":        "${alicloud_vpc.defaulte4zhaL.id}",
					"remark":                "openapi_test",
					"resource_group_name":   "openapi_pop2_test_resg00003",
					"auto_renew":            "false",
					"default_vswitch_id":    "${alicloud_vswitch.default675v38.id}",
					"payment_duration_unit": "Year",
					"specification":         "500",
					"payment_duration":      "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":          "PostPaid",
						"default_vpc_id":        CHECKSET,
						"remark":                "openapi_test",
						"resource_group_name":   "openapi_pop2_test_resg00003",
						"auto_renew":            "false",
						"default_vswitch_id":    CHECKSET,
						"payment_duration_unit": "Year",
						"specification":         "500",
						"payment_duration":      "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"remark":              "openapi_test_update002",
					"resource_group_name": "openapi_pop2_test_resg_update002",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"remark":              "openapi_test_update002",
						"resource_group_name": "openapi_pop2_test_resg_update002",
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
				Config: testAccConfig(map[string]interface{}{}),
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
				ImportStateVerifyIgnore: []string{"auto_renew", "payment_duration", "payment_duration_unit", "project_id", "specification"},
			},
		},
	})
}

var AlicloudDataWorksDwResourceGroupMap8379 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudDataWorksDwResourceGroupBasicDependence8379(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "defaultZImuCO" {
  description      = "default_resg_test"
  project_name     = "default_resg_test"
  pai_task_enabled = false
  display_name     = "default_resg_test"
}

resource "alicloud_vpc" "defaulte4zhaL" {
  description = "default_resgv2"
  vpc_name    = "default_resgv2"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default675v38" {
  description  = "default_resg"
  vpc_id       = alicloud_vpc.defaulte4zhaL.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "default_resg"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case Dataworks资源组管理_Month_依赖资源自动化创建 8020
func TestAccAliCloudDataWorksDwResourceGroup_basic8020(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_dw_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksDwResourceGroupMap8020)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksDwResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksDwResourceGroupBasicDependence8020)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":          "PostPaid",
					"default_vpc_id":        "${alicloud_vpc.defaulte4zhaL.id}",
					"remark":                "openapi_test",
					"resource_group_name":   "openapi_pop2_test_resg00002",
					"auto_renew":            "false",
					"default_vswitch_id":    "${alicloud_vswitch.default675v38.id}",
					"payment_duration_unit": "Month",
					"specification":         "500",
					"payment_duration":      "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":          "PostPaid",
						"default_vpc_id":        CHECKSET,
						"remark":                "openapi_test",
						"resource_group_name":   "openapi_pop2_test_resg00002",
						"auto_renew":            "false",
						"default_vswitch_id":    CHECKSET,
						"payment_duration_unit": "Month",
						"specification":         "500",
						"payment_duration":      "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"remark":              "openapi_test_update",
					"resource_group_name": "openapi_pop2_test_resg_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"remark":              "openapi_test_update",
						"resource_group_name": "openapi_pop2_test_resg_update",
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
				Config: testAccConfig(map[string]interface{}{}),
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
				ImportStateVerifyIgnore: []string{"auto_renew", "payment_duration", "payment_duration_unit", "project_id", "specification"},
			},
		},
	})
}

var AlicloudDataWorksDwResourceGroupMap8020 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudDataWorksDwResourceGroupBasicDependence8020(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "defaultZImuCO" {
  description      = "default_resg_test03"
  project_name     = "default_resg_test03"
  pai_task_enabled = false
  display_name     = "default_resg_test03"
}

resource "alicloud_vpc" "defaulte4zhaL" {
  description = "default_resgv2"
  vpc_name    = "default_resgv2"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default675v38" {
  description  = "default_resg"
  vpc_id       = alicloud_vpc.defaulte4zhaL.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "default_resg"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case Dataworks资源组接入阿里云资源组与TAG测试用例_YEAR_无依赖资源创建_正式 9319
func TestAccAliCloudDataWorksDwResourceGroup_basic9319(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_dw_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksDwResourceGroupMap9319)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksDwResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksDwResourceGroupBasicDependence9319)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":          "PostPaid",
					"default_vpc_id":        "${alicloud_vpc.defaulte4zhaL.id}",
					"remark":                "openapi接入TAG和资源组测试用资源组",
					"resource_group_name":   "tag_aliyunresg1769480648",
					"auto_renew":            "false",
					"default_vswitch_id":    "${alicloud_vswitch.default675v38.id}",
					"payment_duration_unit": "Year",
					"specification":         "500",
					"payment_duration":      "1",
					"resource_group_id":     "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":          "PostPaid",
						"default_vpc_id":        CHECKSET,
						"remark":                "openapi接入TAG和资源组测试用资源组",
						"resource_group_name":   CHECKSET,
						"auto_renew":            "false",
						"default_vswitch_id":    CHECKSET,
						"payment_duration_unit": "Year",
						"specification":         "500",
						"payment_duration":      "1",
						"resource_group_id":     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"remark":              "openapi接入TAG和资源组测试用资源组_update",
					"resource_group_name": "tag1769480649",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"remark":              "openapi接入TAG和资源组测试用资源组_update",
						"resource_group_name": CHECKSET,
						"resource_group_id":   CHECKSET,
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
				Config: testAccConfig(map[string]interface{}{}),
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
				Config: testAccConfig(map[string]interface{}{}),
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
				ImportStateVerifyIgnore: []string{"auto_renew", "payment_duration", "payment_duration_unit", "project_id", "specification"},
			},
		},
	})
}

var AlicloudDataWorksDwResourceGroupMap9319 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudDataWorksDwResourceGroupBasicDependence9319(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "defaultZImuCO" {
  description      = "default_resg_test0666"
  project_name     = "default_resg_test0666"
  pai_task_enabled = false
  display_name     = "default_resg_test0666"
}

resource "alicloud_vpc" "defaulte4zhaL" {
  description = "default_resgv2"
  vpc_name    = "default_resgv2"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default675v38" {
  description  = "default_resg"
  vpc_id       = alicloud_vpc.defaulte4zhaL.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "default_resg"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case Dataworks资源组接入阿里云资源组与TAG测试用例_YEAR_无依赖资源创建_正式ForRMC 10190
func TestAccAliCloudDataWorksDwResourceGroup_basic10190(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_data_works_dw_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudDataWorksDwResourceGroupMap10190)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DataWorksServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDataWorksDwResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccdataworks%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDataWorksDwResourceGroupBasicDependence10190)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":          "PostPaid",
					"default_vpc_id":        "${alicloud_vpc.defaulte4zhaL.id}",
					"remark":                "openapi接入TAG和资源组测试用资源组",
					"resource_group_name":   "tag_aliyunresg1769480649",
					"auto_renew":            "false",
					"default_vswitch_id":    "${alicloud_vswitch.default675v38.id}",
					"payment_duration_unit": "Year",
					"specification":         "500",
					"payment_duration":      "1",
					"resource_group_id":     "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":          "PostPaid",
						"default_vpc_id":        CHECKSET,
						"remark":                "openapi接入TAG和资源组测试用资源组",
						"resource_group_name":   CHECKSET,
						"auto_renew":            "false",
						"default_vswitch_id":    CHECKSET,
						"payment_duration_unit": "Year",
						"specification":         "500",
						"payment_duration":      "1",
						"resource_group_id":     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"remark":              "openapi接入TAG和资源组测试用资源组_update",
					"resource_group_name": "tag1769480649",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"remark":              "openapi接入TAG和资源组测试用资源组_update",
						"resource_group_name": CHECKSET,
						"resource_group_id":   CHECKSET,
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
				Config: testAccConfig(map[string]interface{}{}),
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
				Config: testAccConfig(map[string]interface{}{}),
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
				ImportStateVerifyIgnore: []string{"auto_renew", "payment_duration", "payment_duration_unit", "project_id", "specification"},
			},
		},
	})
}

var AlicloudDataWorksDwResourceGroupMap10190 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudDataWorksDwResourceGroupBasicDependence10190(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_data_works_project" "defaultZImuCO" {
  description      = "default_resg_test_rmc001"
  project_name     = "default_resg_test_rmc001"
  pai_task_enabled = false
  display_name     = "default_resg_test_rmc001"
}

resource "alicloud_vpc" "defaulte4zhaL" {
  description = "default_resgv2"
  vpc_name    = "default_resgv2"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default675v38" {
  description  = "default_resg"
  vpc_id       = alicloud_vpc.defaulte4zhaL.id
  zone_id      = "cn-beijing-g"
  vswitch_name = "default_resg"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Test DataWorks DwResourceGroup. <<< Resource test cases, automatically generated.
