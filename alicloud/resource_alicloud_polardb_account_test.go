package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PolarDb Account. >>> Resource test cases, automatically generated.
// Case  Account用例_PG_高权限账号 11819
func TestAccAliCloudPolarDbAccount_basic11819(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudPolarDbAccountMap11819)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolarDbAccountBasicDependence11819)
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
					"account_description":         "account_desc",
					"db_cluster_id":               "${alicloud_polardb_cluster.defaultHAnrWv.id}",
					"account_type":                "Super",
					"account_name":                name,
					"account_password":            "test_123",
					"account_lock_state":          "UnLock",
					"account_password_valid_time": "2025-12-09CST2323:1212:181828800",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description":         "account_desc",
						"db_cluster_id":               CHECKSET,
						"account_type":                "Super",
						"account_name":                name,
						"account_password":            "test_123",
						"account_lock_state":          "UnLock",
						"account_password_valid_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_description":         "account_desc_modified",
					"account_password":            "test_1234",
					"account_lock_state":          "Lock",
					"account_password_valid_time": "2025-12-09CST2323:1212:181828800",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description":         "account_desc_modified",
						"account_password":            "test_1234",
						"account_lock_state":          "Lock",
						"account_password_valid_time": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_password"},
			},
		},
	})
}

var AlicloudPolarDbAccountMap11819 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolarDbAccountBasicDependence11819(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_polardb_cluster" "defaultHAnrWv" {
  default_time_zone = "SYSTEM"
  creation_category = "Normal"
  db_version        = "14"
  pay_type          = "PayAsYouGo"
  db_node_class     = "polar.pg.x4.medium"
  zone_id           = "cn-beijing-k"
  db_type           = "PostgreSQL"
  creation_option   = "Normal"
}


`, name)
}

// Case  Account用例_PG 11813
func TestAccAliCloudPolarDbAccount_basic11813(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudPolarDbAccountMap11813)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolarDbAccountBasicDependence11813)
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
					"account_description":         "account_desc",
					"db_cluster_id":               "${alicloud_polardb_cluster.defaultHAnrWv.id}",
					"account_type":                "Normal",
					"account_name":                name,
					"account_password":            "test_123",
					"account_lock_state":          "UnLock",
					"account_password_valid_time": "2025-12-09CST2323:1212:191928800",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description":         "account_desc",
						"db_cluster_id":               CHECKSET,
						"account_type":                "Normal",
						"account_name":                name,
						"account_password":            "test_123",
						"account_lock_state":          "UnLock",
						"account_password_valid_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_description":         "account_desc_modified",
					"account_password":            "test_1234",
					"account_lock_state":          "Lock",
					"account_password_valid_time": "2025-12-09CST2323:1212:191928800",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description":         "account_desc_modified",
						"account_password":            "test_1234",
						"account_lock_state":          "Lock",
						"account_password_valid_time": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_password"},
			},
		},
	})
}

var AlicloudPolarDbAccountMap11813 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolarDbAccountBasicDependence11813(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_polardb_cluster" "defaultHAnrWv" {
  default_time_zone = "SYSTEM"
  creation_category = "Normal"
  db_version        = "14"
  pay_type          = "PayAsYouGo"
  db_node_class     = "polar.pg.x4.medium"
  zone_id           = "cn-beijing-k"
  db_type           = "PostgreSQL"
  creation_option   = "Normal"
}


`, name)
}

// Case  Account用例 9211
func TestAccAliCloudPolarDbAccount_basic9211(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudPolarDbAccountMap9211)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolarDbAccountBasicDependence9211)
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
					"account_description": "account_desc",
					"db_cluster_id":       "${alicloud_polardb_cluster.defaultHgpFHo.id}",
					"account_type":        "Normal",
					"account_name":        name,
					"account_password":    "test_123",
					"account_lock_state":  "UnLock",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description": "account_desc",
						"db_cluster_id":       CHECKSET,
						"account_type":        "Normal",
						"account_name":        name,
						"account_password":    "test_123",
						"account_lock_state":  "UnLock",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_password": "test_1234",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_password": "test_1234",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_password"},
			},
		},
	})
}

var AlicloudPolarDbAccountMap9211 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolarDbAccountBasicDependence9211(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_polardb_cluster" "defaultHgpFHo" {
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

// Test PolarDb Account. <<< Resource test cases, automatically generated.
