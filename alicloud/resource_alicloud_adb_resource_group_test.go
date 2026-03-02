package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb ResourceGroup. >>> Resource test cases, automatically generated.
// Case ResourceGroup全流程_Spark_Interactive 10838
func TestAccAliCloudAdbResourceGroup_basic10838(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbResourceGroupMap10838)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbResourceGroupBasicDependence10838)
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
					"group_name":            "TEST328",
					"group_type":            "interactive",
					"db_cluster_id":         "${alicloud_adb_db_cluster_lake_version.defaultibkliK.id}",
					"min_cluster_count":     "1",
					"cluster_size_resource": "16ACU",
					"cluster_mode":          "AutoScale",
					"max_cluster_count":     "1",
					"engine":                "SparkWarehouse",
					"engine_params": map[string]interface{}{
						"\"spark.adb.version\"": "3.2",
					},
					"user_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":            CHECKSET,
						"group_type":            "interactive",
						"db_cluster_id":         CHECKSET,
						"min_cluster_count":     "1",
						"cluster_size_resource": "16ACU",
						"cluster_mode":          "AutoScale",
						"max_cluster_count":     "1",
						"engine":                "SparkWarehouse",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"min_cluster_count":     "2",
					"cluster_size_resource": "24ACU",
					"max_cluster_count":     "4",
					"engine_params": map[string]interface{}{
						"\"spark.adb.version\"": "3.3",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"min_cluster_count":     "2",
						"cluster_size_resource": "24ACU",
						"max_cluster_count":     "4",
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

var AlicloudAdbResourceGroupMap10838 = map[string]string{
	"status":            CHECKSET,
	"create_time":       CHECKSET,
	"port":              CHECKSET,
	"update_time":       CHECKSET,
	"connection_string": CHECKSET,
}

func AlicloudAdbResourceGroupBasicDependence10838(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "defaultQubfrN" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "adb北京vpc-tf-111"
}

resource "alicloud_vswitch" "defaultVAAe0q" {
  vpc_id       = alicloud_vpc.defaultQubfrN.id
  zone_id      = "cn-beijing-k"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "北京k测试tf1111"
}

resource "alicloud_adb_db_cluster_lake_version" "defaultibkliK" {
  storage_resource              = "24ACU"
  product_form                  = "LegacyForm"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  enable_default_resource_group = false
  zone_id                       = "cn-beijing-k"
  vswitch_id                    = alicloud_vswitch.defaultVAAe0q.id
  vpc_id                        = alicloud_vpc.defaultQubfrN.id
  db_cluster_description        = "TF测试专用"
  compute_resource              = "32ACU"
}

resource "alicloud_adb_account" "defaultwxttS6" {
  account_description = "TF测试创建高权限账号"
  db_cluster_id       = alicloud_adb_db_cluster_lake_version.defaultibkliK.id
  account_name        = "admin"
  account_password    = "Aliyun@123"
}

resource "alicloud_adb_account" "defaultabqCUc" {
  db_cluster_id    = alicloud_adb_db_cluster_lake_version.defaultibkliK.id
  account_name     = "user"
  account_password = "Aliyun@123"
}


`, name)
}

// Case ResourceGroup全流程_Xihe_Interactive 10841
func TestAccAliCloudAdbResourceGroup_basic10841(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbResourceGroupMap10841)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbResourceGroupBasicDependence10841)
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
					"group_name":           "TEST505",
					"group_type":           "interactive",
					"db_cluster_id":        "${alicloud_adb_db_cluster_lake_version.defaultgVJVe9.id}",
					"min_compute_resource": "0ACU",
					"max_compute_resource": "16ACU",
					"cluster_mode":         "Disable",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":           CHECKSET,
						"group_type":           "interactive",
						"db_cluster_id":        CHECKSET,
						"min_compute_resource": "0ACU",
						"max_compute_resource": "16ACU",
						"cluster_mode":         "Disable",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"min_compute_resource": "16ACU",
					"max_compute_resource": "32ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"min_compute_resource": "16ACU",
						"max_compute_resource": "32ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cluster_mode":      "AutoScale",
					"min_cluster_count": "1",
					"max_cluster_count": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cluster_mode":      "AutoScale",
						"min_cluster_count": "1",
						"max_cluster_count": "2",
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

var AlicloudAdbResourceGroupMap10841 = map[string]string{
	"status":            CHECKSET,
	"create_time":       CHECKSET,
	"port":              CHECKSET,
	"update_time":       CHECKSET,
	"connection_string": CHECKSET,
}

func AlicloudAdbResourceGroupBasicDependence10841(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "default3pjBMW" {
  cidr_block  = "172.16.0.0/12"
  enable_ipv6 = false
  vpc_name    = "adb北京vpc-tf-111"
}

resource "alicloud_vswitch" "default1vC8QC" {
  vpc_id       = alicloud_vpc.default3pjBMW.id
  zone_id      = "cn-beijing-k"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "北京k测试tf1111"
}

resource "alicloud_adb_db_cluster_lake_version" "defaultgVJVe9" {
  storage_resource              = "24ACU"
  product_form                  = "LegacyForm"
  db_cluster_version            = "5.0"
  payment_type                  = "PayAsYouGo"
  zone_id                       = "cn-beijing-k"
  vpc_id                        = alicloud_vpc.default3pjBMW.id
  vswitch_id                    = alicloud_vswitch.default1vC8QC.id
  db_cluster_description        = "TF测试专用"
  compute_resource              = "32ACU"
  enable_default_resource_group = false
}


`, name)
}

// Case ResourceGroup全流程_数仓_副本1770775849704_副本1770791063321 12553
func TestAccAliCloudAdbResourceGroup_basic12553(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbResourceGroupMap12553)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbResourceGroupBasicDependence12553)
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
					"group_name":    "TEST697",
					"group_type":    "interactive",
					"db_cluster_id": "${alicloud_adb_db_cluster.defaulteOVp71.id}",
					"user_list": []string{
						"${alicloud_adb_account.default0GWsrd.account_name}", "${alicloud_adb_account.default5N8gc9.account_name}", "${alicloud_adb_account.defaultUlPzQj.account_name}"},
					"node_num": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":    CHECKSET,
						"group_type":    "interactive",
						"db_cluster_id": CHECKSET,
						"user_list.#":   "3",
						"node_num":      "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_list": []string{
						"${alicloud_adb_account.default5N8gc9.account_name}"},
					"node_num": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_list.#": "1",
						"node_num":    "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_type": "batch",
					"user_list": []string{
						"${alicloud_adb_account.default0GWsrd.account_name}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_type":  "batch",
						"user_list.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_list": []string{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_list.#": "0",
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

var AlicloudAdbResourceGroupMap12553 = map[string]string{
	"status":            CHECKSET,
	"create_time":       CHECKSET,
	"port":              CHECKSET,
	"update_time":       CHECKSET,
	"connection_string": CHECKSET,
}

func AlicloudAdbResourceGroupBasicDependence12553(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_adb_db_cluster" "defaulteOVp71" {
  disk_performance_level   = "PL1"
  elastic_io_resource      = "1"
  db_cluster_version       = "3.0"
  payment_type             = "PayAsYouGo"
  db_cluster_category      = "MIXED_STORAGE"
  elastic_io_resource_size = "8Core64GB"
  zone_id                  = "cn-beijing-k"
  mode                     = "flexible"
  compute_resource         = "32Core128GBNEW"
}

resource "alicloud_adb_account" "default2DHJXc" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "admin"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建高权限账号"
}

resource "alicloud_adb_account" "default0GWsrd" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}

resource "alicloud_adb_account" "default5N8gc9" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user1"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}

resource "alicloud_adb_account" "defaultUlPzQj" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user2"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}


`, name)
}

// Case ResourceGroup全流程_数仓_副本1770775849704 12551
func TestAccAliCloudAdbResourceGroup_basic12551(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbResourceGroupMap12551)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbResourceGroupBasicDependence12551)
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
					"group_name":    "TEST121",
					"group_type":    "interactive",
					"db_cluster_id": "${alicloud_adb_db_cluster.defaulteOVp71.id}",
					"user_list": []string{
						"${alicloud_adb_account.default0GWsrd.account_name}", "${alicloud_adb_account.default5N8gc9.account_name}", "${alicloud_adb_account.defaultUlPzQj.account_name}"},
					"node_num": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":    CHECKSET,
						"group_type":    "interactive",
						"db_cluster_id": CHECKSET,
						"user_list.#":   "3",
						"node_num":      "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_list": []string{
						"${alicloud_adb_account.default5N8gc9.account_name}"},
					"node_num": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_list.#": "1",
						"node_num":    "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_type": "batch",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_type": "batch",
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

var AlicloudAdbResourceGroupMap12551 = map[string]string{
	"status":            CHECKSET,
	"create_time":       CHECKSET,
	"port":              CHECKSET,
	"update_time":       CHECKSET,
	"connection_string": CHECKSET,
}

func AlicloudAdbResourceGroupBasicDependence12551(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_adb_db_cluster" "defaulteOVp71" {
  disk_performance_level   = "PL1"
  elastic_io_resource      = "1"
  db_cluster_version       = "3.0"
  payment_type             = "PayAsYouGo"
  db_cluster_category      = "MIXED_STORAGE"
  elastic_io_resource_size = "8Core64GB"
  zone_id                  = "cn-beijing-k"
  mode                     = "flexible"
  compute_resource         = "32Core128GBNEW"
}

resource "alicloud_adb_account" "default2DHJXc" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "admin"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建高权限账号"
}

resource "alicloud_adb_account" "default0GWsrd" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}

resource "alicloud_adb_account" "default5N8gc9" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user1"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}

resource "alicloud_adb_account" "defaultUlPzQj" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user2"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}


`, name)
}

// Case ResourceGroup全流程_数仓 7215
func TestAccAliCloudAdbResourceGroup_basic7215(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbResourceGroupMap7215)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbResourceGroupBasicDependence7215)
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
					"group_name":    "TEST44",
					"group_type":    "interactive",
					"db_cluster_id": "${alicloud_adb_db_cluster.defaulteOVp71.id}",
					"user_list": []string{
						"${alicloud_adb_account.default0GWsrd.account_name}"},
					"node_num": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_name":    CHECKSET,
						"group_type":    "interactive",
						"db_cluster_id": CHECKSET,
						"user_list.#":   "1",
						"node_num":      "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"user_list": []string{
						"${alicloud_adb_account.default5N8gc9.account_name}"},
					"node_num": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"user_list.#": "1",
						"node_num":    "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_type": "batch",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_type": "batch",
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

var AlicloudAdbResourceGroupMap7215 = map[string]string{
	"status":            CHECKSET,
	"create_time":       CHECKSET,
	"port":              CHECKSET,
	"update_time":       CHECKSET,
	"connection_string": CHECKSET,
}

func AlicloudAdbResourceGroupBasicDependence7215(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_adb_db_cluster" "defaulteOVp71" {
  disk_performance_level   = "PL1"
  elastic_io_resource      = "1"
  db_cluster_version       = "3.0"
  payment_type             = "PayAsYouGo"
  db_cluster_category      = "MIXED_STORAGE"
  elastic_io_resource_size = "8Core64GB"
  zone_id                  = "cn-beijing-k"
  mode                     = "flexible"
  compute_resource         = "32Core128GBNEW"
}

resource "alicloud_adb_account" "default2DHJXc" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "admin"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建高权限账号"
}

resource "alicloud_adb_account" "default0GWsrd" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}

resource "alicloud_adb_account" "default5N8gc9" {
  db_cluster_id       = alicloud_adb_db_cluster.defaulteOVp71.id
  account_name        = "user1"
  account_password    = "Aliyun@123"
  account_description = "TF测试创建普通账号"
}


`, name)
}

// Test Adb ResourceGroup. <<< Resource test cases, automatically generated.

func TestAccAliCloudAdbResourceGroup_basic0(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbResourceGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tftestaccAdbResourceGroup%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbResourceGroupBasicDependence0)
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
					"db_cluster_id": "${alicloud_adb_db_cluster.default.id}",
					"group_name":    name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id": CHECKSET,
						"group_name":    CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"group_type": "batch",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"group_type": "batch",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"node_num": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"node_num": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"users": []string{"${alicloud_adb_account.default.account_name}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"users.#": "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudAdbResourceGroup_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbResourceGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc-AdbResourceGroup%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbResourceGroupBasicDependence0)
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
					"db_cluster_id": "${alicloud_adb_db_cluster.default.id}",
					"group_name":    name,
					"group_type":    "batch",
					"node_num":      "1",
					"users":         []string{"${alicloud_adb_account.default.account_name}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id": CHECKSET,
						"group_name":    CHECKSET,
						"group_type":    "batch",
						"node_num":      "1",
						"users.#":       "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudAdbResourceGroup_basic1(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbResourceGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("TFADBRG%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbResourceGroupBasicDependence1)
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
					"db_cluster_id":        "${alicloud_adb_lake_account.default.db_cluster_id}",
					"group_name":           name,
					"min_compute_resource": "16ACU",
					"max_compute_resource": "16ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":        CHECKSET,
						"group_name":           CHECKSET,
						"min_compute_resource": "16ACU",
						"max_compute_resource": "16ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_compute_resource": "128ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_compute_resource": "128ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"min_compute_resource": "128ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"min_compute_resource": "128ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cluster_mode":          "AutoScale",
					"cluster_size_resource": "16ACU",
					"max_cluster_count":     "2",
					"min_cluster_count":     "1",
					"min_compute_resource":  REMOVEKEY,
					"max_compute_resource":  REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cluster_mode":          "AutoScale",
						"cluster_size_resource": "16ACU",
						"max_cluster_count":     "2",
						"min_cluster_count":     "1",
						"min_compute_resource":  CHECKSET,
						"max_compute_resource":  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cluster_size_resource": "32ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cluster_size_resource": "32ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_cluster_count": "6",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_cluster_count": "6",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"min_cluster_count": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"min_cluster_count": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"users": []string{"${alicloud_adb_lake_account.default.account_name}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"users.#": "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudAdbResourceGroup_basic1_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbResourceGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("TFADBRG%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbResourceGroupBasicDependence1)
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
					"db_cluster_id":        "${alicloud_adb_lake_account.default.db_cluster_id}",
					"group_name":           name,
					"group_type":           "job",
					"cluster_mode":         "Disable",
					"min_compute_resource": "16ACU",
					"max_compute_resource": "16ACU",
					"users":                []string{"${alicloud_adb_lake_account.default.account_name}"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":        CHECKSET,
						"group_name":           CHECKSET,
						"group_type":           "job",
						"cluster_mode":         "Disable",
						"min_compute_resource": "16ACU",
						"max_compute_resource": "16ACU",
						"users.#":              "1",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudAdbResourceGroup_basic2(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbResourceGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("TFADBRG%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbResourceGroupBasicDependence1)
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
					"db_cluster_id":         "${alicloud_adb_lake_account.default.db_cluster_id}",
					"group_name":            name,
					"cluster_mode":          "AutoScale",
					"cluster_size_resource": "16ACU",
					"max_cluster_count":     "2",
					"min_cluster_count":     "1",
					"engine":                "SparkWarehouse",
					"engine_params": map[string]interface{}{
						"\"spark.adb.version\"":      "3.5",
						"\"spark.app.log.rootPath\"": "oss://" + "${data.alicloud_oss_buckets.default.buckets.0.name}" + "/",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":         CHECKSET,
						"group_name":            CHECKSET,
						"cluster_mode":          "AutoScale",
						"cluster_size_resource": "16ACU",
						"max_cluster_count":     "2",
						"min_cluster_count":     "1",
						"engine":                "SparkWarehouse",
						"engine_params.%":       "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cluster_size_resource": "36ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cluster_size_resource": "36ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_cluster_count": "6",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_cluster_count": "6",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"min_cluster_count": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"min_cluster_count": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"engine_params": map[string]interface{}{
						"\"spark.adb.version\"":                 "3.5",
						"\"spark.app.log.rootPath\"":            "oss://" + "${data.alicloud_oss_buckets.default.buckets.0.name}" + "/",
						"\"spark.driver.memoryOverheadFactor\"": "0.5",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"engine_params.%": "3",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccAliCloudAdbResourceGroup_basic2_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_resource_group.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbResourceGroupMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbResourceGroup")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("TFADBRG%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbResourceGroupBasicDependence1)
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
					"db_cluster_id":         "${alicloud_adb_lake_account.default.db_cluster_id}",
					"group_name":            name,
					"group_type":            "interactive",
					"cluster_mode":          "AutoScale",
					"cluster_size_resource": "16ACU",
					"max_cluster_count":     "2",
					"min_cluster_count":     "1",
					"engine":                "SparkWarehouse",
					"engine_params": map[string]interface{}{
						"\"spark.adb.version\"":      "3.5",
						"\"spark.app.log.rootPath\"": "oss://" + "${data.alicloud_oss_buckets.default.buckets.0.name}" + "/",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":         CHECKSET,
						"group_name":            name,
						"group_type":            "interactive",
						"cluster_mode":          "AutoScale",
						"cluster_size_resource": "16ACU",
						"max_cluster_count":     "2",
						"min_cluster_count":     "1",
						"engine":                "SparkWarehouse",
						"engine_params.%":       "2",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

var AliCloudAdbResourceGroupMap0 = map[string]string{
	"group_type": CHECKSET,
	"status":     CHECKSET,
}

func AliCloudAdbResourceGroupBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_adb_zones" "default" {
	}
	
	data "alicloud_vpcs" "default" {
  		name_regex = "^default-NODELETING$"
	}
	
	data "alicloud_vswitches" "default" {
  		vpc_id  = data.alicloud_vpcs.default.ids.0
  		zone_id = data.alicloud_adb_zones.default.ids.0
	}
	
	resource "alicloud_adb_db_cluster" "default" {
  		compute_resource    = "32Core128GBNEW"
  		db_cluster_category = "MixedStorage"
  		description         = var.name
  		elastic_io_resource = 1
  		mode                = "flexible"
  		payment_type        = "PayAsYouGo"
  		vpc_id              = data.alicloud_vpcs.default.ids.0
  		vswitch_id          = data.alicloud_vswitches.default.ids.0
  		zone_id             = data.alicloud_adb_zones.default.zones.0.id
	}
	
	resource "alicloud_adb_account" "default" {
  		db_cluster_id    = alicloud_adb_db_cluster.default.id
  		account_name     = "tf_account_name"
  		account_password = "YourPassword123!"
	}
`, name)
}

func AliCloudAdbResourceGroupBasicDependence1(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_oss_buckets" "default" {
	}

	data "alicloud_adb_zones" "default" {
	}
	
	data "alicloud_vpcs" "default" {
  		name_regex = "^default-NODELETING$"
	}
	
	data "alicloud_vswitches" "default" {
  		vpc_id  = data.alicloud_vpcs.default.ids.0
  		zone_id = data.alicloud_adb_zones.default.ids.0
	}
	
	resource "alicloud_adb_db_cluster_lake_version" "default" {
  		db_cluster_version            = "5.0"
  		vpc_id                        = data.alicloud_vpcs.default.ids.0
  		vswitch_id                    = data.alicloud_vswitches.default.ids.0
  		zone_id                       = data.alicloud_adb_zones.default.ids.0
  		compute_resource              = "128ACU"
  		storage_resource              = "0ACU"
  		payment_type                  = "PayAsYouGo"
  		enable_default_resource_group = false
	}
	
	resource "alicloud_adb_lake_account" "default" {
  		db_cluster_id    = alicloud_adb_db_cluster_lake_version.default.id
		account_type     = "Super"
  		account_name     = "tf_account_name"
  		account_password = "YourPassword123!"
	}
`, name)
}
