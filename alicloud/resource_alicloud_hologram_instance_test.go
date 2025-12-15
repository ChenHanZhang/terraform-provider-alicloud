package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Hologram Instance. >>> Resource test cases, automatically generated.
// Case warehouse_test 3920
func TestAccAliCloudHologramInstance_basic3920(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hologram_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudHologramInstanceMap3920)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HologramServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHologramInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHologramInstanceBasicDependence3920)
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
					"zone_id":       "cn-beijing-l",
					"pricing_cycle": "Hour",
					"cpu":           "32",
					"duration":      "1",
					"auto_pay":      "true",
					"instance_name": name,
					"gateway_count": "2",
					"payment_type":  "PayAsYouGo",
					"instance_type": "Warehouse",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":            "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":           "cn-beijing-l",
						"pricing_cycle":     "Hour",
						"cpu":               "32",
						"duration":          "1",
						"auto_pay":          "true",
						"instance_name":     name,
						"gateway_count":     "2",
						"payment_type":      "PayAsYouGo",
						"instance_type":     "Warehouse",
						"endpoints.#":       "2",
						"resource_group_id": CHECKSET,
						"status":            "Running",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"gateway_count":     "4",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scale_type":        "UPGRADE",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"gateway_count":     "4",
						"resource_group_id": CHECKSET,
						"scale_type":        "UPGRADE",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":            "Suspended",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"status":            "Suspended",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"status":            "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"status":            "Running",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "duration", "initial_databases", "pricing_cycle", "scale_type"},
			},
		},
	})
}

var AlicloudHologramInstanceMap3920 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudHologramInstanceBasicDependence3920(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVpc"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  zone_id      = "cn-beijing-l"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch"
}


`, name)
}

// Case 按量付费测试_副本1764148617865_副本1765436193872 12058
func TestAccAliCloudHologramInstance_basic12058(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hologram_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudHologramInstanceMap12058)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HologramServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHologramInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHologramInstanceBasicDependence12058)
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
					"cold_storage_size": "0",
					"zone_id":           "cn-shenzhen-f",
					"pricing_cycle":     "Hour",
					"cpu":               "32",
					"storage_size":      "0",
					"duration":          "1",
					"auto_pay":          "true",
					"instance_name":     name,
					"payment_type":      "PayAsYouGo",
					"instance_type":     "Standard",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
					"status":            "Running",
					"initial_databases": "abcd, 123, _underline_db",
					"enable_ssl":        "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cold_storage_size": "0",
						"zone_id":           "cn-shenzhen-f",
						"pricing_cycle":     "Hour",
						"cpu":               "32",
						"storage_size":      "0",
						"duration":          "1",
						"auto_pay":          "true",
						"instance_name":     name,
						"payment_type":      "PayAsYouGo",
						"instance_type":     "Standard",
						"resource_group_id": CHECKSET,
						"endpoints.#":       "2",
						"status":            "Running",
						"initial_databases": "abcd, 123, _underline_db",
						"enable_ssl":        "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cpu":               "64",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scale_type":        "UPGRADE",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cpu":               "64",
						"resource_group_id": CHECKSET,
						"scale_type":        "UPGRADE",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch2.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch2.vpc_id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"endpoints.#":       "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"endpoints.#":       "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_ssl": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_ssl": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_ssl": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_ssl": "false",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "duration", "initial_databases", "pricing_cycle", "scale_type"},
			},
		},
	})
}

var AlicloudHologramInstanceMap12058 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudHologramInstanceBasicDependence12058(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVpc"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  zone_id      = "cn-shenzhen-f"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch"
}

resource "alicloud_vpc" "defaultVPC2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVPC2"
}

resource "alicloud_vswitch" "defaultVSwitch2" {
  vpc_id       = alicloud_vpc.defaultVPC2.id
  zone_id      = "cn-shenzhen-f"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch2"
}


`, name)
}

// Case 包年包月测试 3916
func TestAccAliCloudHologramInstance_basic3916(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hologram_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudHologramInstanceMap3916)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HologramServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHologramInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHologramInstanceBasicDependence3916)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"zone_id":           "cn-shenzhen-f",
					"pricing_cycle":     "Month",
					"cpu":               "8",
					"storage_size":      "100",
					"duration":          "1",
					"auto_pay":          "true",
					"instance_name":     name,
					"payment_type":      "Subscription",
					"instance_type":     "Standard",
					"cold_storage_size": "100",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":           "cn-shenzhen-f",
						"pricing_cycle":     "Month",
						"cpu":               "8",
						"storage_size":      "100",
						"duration":          "1",
						"auto_pay":          "true",
						"instance_name":     name,
						"payment_type":      "Subscription",
						"instance_type":     "Standard",
						"cold_storage_size": "100",
						"endpoints.#":       "2",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_size":      "200",
					"instance_name":     name + "_update",
					"cold_storage_size": "200",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scale_type":        "UPGRADE",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_size":      "200",
						"instance_name":     name + "_update",
						"cold_storage_size": "200",
						"resource_group_id": CHECKSET,
						"scale_type":        "UPGRADE",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoints.#":       "1",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"status":            "Suspended",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"status":            "Suspended",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_size":      "100",
					"instance_name":     name + "_update",
					"cold_storage_size": "100",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_size":      "100",
						"instance_name":     name + "_update",
						"cold_storage_size": "100",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_size":      "200",
					"instance_name":     name + "_update",
					"cold_storage_size": "200",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_size":      "200",
						"instance_name":     name + "_update",
						"cold_storage_size": "200",
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"auto_pay", "duration", "initial_databases", "pricing_cycle", "scale_type"},
			},
		},
	})
}

var AlicloudHologramInstanceMap3916 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudHologramInstanceBasicDependence3916(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVpc"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  zone_id      = "cn-shenzhen-f"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch"
}


`, name)
}

// Case tag测试final 4132
func TestAccAliCloudHologramInstance_basic4132(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hologram_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudHologramInstanceMap4132)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HologramServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHologramInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHologramInstanceBasicDependence4132)
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
					"duration":      "1",
					"instance_type": "Standard",
					"zone_id":       "cn-hangzhou-j",
					"pricing_cycle": "Hour",
					"cpu":           "8",
					"instance_name": name,
					"payment_type":  "PayAsYouGo",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"duration":          "1",
						"instance_type":     "Standard",
						"zone_id":           "cn-hangzhou-j",
						"pricing_cycle":     "Hour",
						"cpu":               "8",
						"instance_name":     name,
						"payment_type":      "PayAsYouGo",
						"endpoints.#":       "2",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
						{
							"type": "Internet",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scale_type":        "UPGRADE",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoints.#":       "3",
						"resource_group_id": CHECKSET,
						"scale_type":        "UPGRADE",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"status":            "Running",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoints.#":       "1",
						"resource_group_id": CHECKSET,
						"status":            "Running",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cpu":               "32",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cpu":               "32",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"endpoints": []map[string]interface{}{
						{
							"endpoint": "Intranet",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"endpoints.#":       "1",
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"auto_pay", "duration", "initial_databases", "pricing_cycle", "scale_type"},
			},
		},
	})
}

var AlicloudHologramInstanceMap4132 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudHologramInstanceBasicDependence4132(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaulVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVPc"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaulVpc.id
  zone_id      = "cn-hangzhou-j"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch"
}


`, name)
}

// Case 按量付费测试 4858
func TestAccAliCloudHologramInstance_basic4858(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hologram_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudHologramInstanceMap4858)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HologramServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHologramInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHologramInstanceBasicDependence4858)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shenzhen"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"cold_storage_size": "0",
					"zone_id":           "cn-shenzhen-f",
					"pricing_cycle":     "Hour",
					"cpu":               "8",
					"storage_size":      "0",
					"duration":          "1",
					"auto_pay":          "true",
					"instance_name":     name,
					"payment_type":      "PayAsYouGo",
					"instance_type":     "Standard",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
					"status":            "Running",
					"initial_databases": "abcd, 123, _underline_db",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cold_storage_size": "0",
						"zone_id":           "cn-shenzhen-f",
						"pricing_cycle":     "Hour",
						"cpu":               "8",
						"storage_size":      "0",
						"duration":          "1",
						"auto_pay":          "true",
						"instance_name":     name,
						"payment_type":      "PayAsYouGo",
						"instance_type":     "Standard",
						"resource_group_id": CHECKSET,
						"endpoints.#":       "2",
						"status":            "Running",
						"initial_databases": "abcd, 123, _underline_db",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cpu":               "32",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"scale_type":        "UPGRADE",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cpu":               "32",
						"resource_group_id": CHECKSET,
						"scale_type":        "UPGRADE",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch2.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch2.vpc_id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"endpoints.#":       "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"endpoints": []map[string]interface{}{
						{
							"type": "Intranet",
						},
						{
							"type":       "VPCSingleTunnel",
							"vswitch_id": "${alicloud_vswitch.defaultVSwitch.id}",
							"vpc_id":     "${alicloud_vswitch.defaultVSwitch.vpc_id}",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"endpoints.#":       "2",
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
				ImportStateVerifyIgnore: []string{"auto_pay", "duration", "initial_databases", "pricing_cycle", "scale_type"},
			},
		},
	})
}

var AlicloudHologramInstanceMap4858 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudHologramInstanceBasicDependence4858(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVpc"
}

resource "alicloud_vswitch" "defaultVSwitch" {
  vpc_id       = alicloud_vpc.defaultVpc.id
  zone_id      = "cn-shenzhen-f"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch"
}

resource "alicloud_vpc" "defaultVPC2" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "defaultVPC2"
}

resource "alicloud_vswitch" "defaultVSwitch2" {
  vpc_id       = alicloud_vpc.defaultVPC2.id
  zone_id      = "cn-shenzhen-f"
  cidr_block   = "172.16.53.0/24"
  vswitch_name = "defaultVSwitch2"
}


`, name)
}

// Test Hologram Instance. <<< Resource test cases, automatically generated.
