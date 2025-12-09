package alicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PolarDb Database. >>> Resource test cases, automatically generated.
// Case Database用例_pg 11853
func TestAccAliCloudPolarDbDatabase_basic11853(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_database.default"
	ra := resourceAttrInit(resourceId, AlicloudPolarDbDatabaseMap11853)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbDatabase")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpolardb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolarDbDatabaseBasicDependence11853)
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
					"character_set_name": "UTF8",
					"db_cluster_id":      "${alicloud_polardb_cluster.defaultmRyliI.id}",
					"db_name":            "testdb001",
					"collate":            "C",
					"ctype":              "C",
					"db_description":     "testDesc",
					"account_name":       "${alicloud_polardb_account.defaultIcSx3U.account_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"character_set_name": "UTF8",
						"db_cluster_id":      CHECKSET,
						"db_name":            "testdb001",
						"collate":            "C",
						"ctype":              "C",
						"db_description":     "testDesc",
						"account_name":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_name": "${alicloud_polardb_account.defaultvHpplR.account_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_name": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"collate", "ctype"},
			},
		},
	})
}

var AlicloudPolarDbDatabaseMap11853 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolarDbDatabaseBasicDependence11853(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_polardb_cluster" "defaultmRyliI" {
  default_time_zone   = "SYSTEM"
  db_node_class       = "polar.pg.x4.medium"
  creation_category   = "Normal"
  hot_standby_cluster = "OFF"
  db_version          = "14"
  pay_type            = "PostPaid"
  zone_id             = "cn-beijing-k"
  creation_option     = "Normal"
  db_type             = "PostgreSQL"
}

resource "alicloud_polardb_account" "defaultIcSx3U" {
  db_cluster_id    = alicloud_polardb_cluster.defaultmRyliI.id
  account_type     = "Normal"
  account_name     = "testacc001"
  account_password = "Test1234"
}

resource "alicloud_polardb_account" "defaultvHpplR" {
  db_cluster_id    = alicloud_polardb_cluster.defaultmRyliI.id
  account_type     = "Normal"
  account_name     = "testacc002"
  account_password = "Test1234"
}


`, name)
}

// Case Database用例_mysql 11854
func TestAccAliCloudPolarDbDatabase_basic11854(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_database.default"
	ra := resourceAttrInit(resourceId, AlicloudPolarDbDatabaseMap11854)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbDatabase")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpolardb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolarDbDatabaseBasicDependence11854)
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
					"character_set_name": "utf8mb4",
					"db_cluster_id":      "${alicloud_polardb_cluster.defaultmRyliI.id}",
					"db_name":            "testdb001",
					"db_description":     "testDesc",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"character_set_name": "utf8mb4",
						"db_cluster_id":      CHECKSET,
						"db_name":            "testdb001",
						"db_description":     "testDesc",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_description": "testDescModified",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_description": "testDescModified",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"collate", "ctype"},
			},
		},
	})
}

var AlicloudPolarDbDatabaseMap11854 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolarDbDatabaseBasicDependence11854(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_polardb_cluster" "defaultmRyliI" {
  default_time_zone = "SYSTEM"
  creation_category = "Normal"
  db_version        = "8.0"
  pay_type          = "PayAsYouGo"
  db_node_class     = "polar.mysql.x4.medium"
  zone_id           = "cn-beijing-k"
  db_type           = "MySQL"
  creation_option   = "Normal"
}


`, name)
}

// Case  Database用例1 9535
func TestAccAliCloudPolarDbDatabase_basic9535(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_database.default"
	ra := resourceAttrInit(resourceId, AlicloudPolarDbDatabaseMap9535)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbDatabase")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpolardb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolarDbDatabaseBasicDependence9535)
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
					"character_set_name": "utf8",
					"db_cluster_id":      "${alicloud_polardb_cluster.database依赖dbCluster.id}",
					"db_name":            "test_db",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"character_set_name": "utf8",
						"db_cluster_id":      CHECKSET,
						"db_name":            "test_db",
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"collate", "ctype"},
			},
		},
	})
}

var AlicloudPolarDbDatabaseMap9535 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolarDbDatabaseBasicDependence9535(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_polardb_cluster" "database依赖dbCluster" {
  db_node_class     = "polar.mysql.x4.large"
  creation_category = "Normal"
  db_version        = "8.0"
  pay_type          = "Postpaid"
  creation_option   = "Normal"
  db_type           = "MySQL"
  zone_id           = "cn-beijing-k"
}


`, name)
}

// Test PolarDb Database. <<< Resource test cases, automatically generated.
