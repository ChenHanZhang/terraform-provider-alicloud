package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test PolarDb Account. >>> Resource test cases, automatically generated.
// Test Polardb Account. >>> Resource test cases, automatically generated.
// Case  Account用例_PG_高权限账号 11819
func TestAccAliCloudPolardbAccount_basic11819(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudPolardbAccountMap11819)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolardbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolardbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolardbAccountBasicDependence11819)
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
					"account_description":         "account_desc",
					"db_cluster_id":               "${alicloud_polardb_cluster.defaultHAnrWv.id}",
					"account_type":                "Super",
					"account_name":                name,
					"account_password":            "test_123",
					"account_lock_state":          "UnLock",
					"account_password_valid_time": "2026-02-02CST1111:0202:434328800",
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
					"account_password_valid_time": "2026-02-02CST1111:0202:434328800",
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

var AlicloudPolardbAccountMap11819 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolardbAccountBasicDependence11819(name string) string {
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
func TestAccAliCloudPolardbAccount_basic11813(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudPolardbAccountMap11813)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolardbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolardbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolardbAccountBasicDependence11813)
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
					"account_description":         "account_desc",
					"db_cluster_id":               "${alicloud_polardb_cluster.defaultHAnrWv.id}",
					"account_type":                "Normal",
					"account_name":                name,
					"account_password":            "test_123",
					"account_lock_state":          "UnLock",
					"account_password_valid_time": "2026-02-02CST1111:0202:434328800",
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
					"account_password_valid_time": "2026-02-02CST1111:0202:434328800",
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

var AlicloudPolardbAccountMap11813 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolardbAccountBasicDependence11813(name string) string {
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
func TestAccAliCloudPolardbAccount_basic9211(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudPolardbAccountMap9211)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolardbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolardbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolardbAccountBasicDependence9211)
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

var AlicloudPolardbAccountMap9211 = map[string]string{
	"status": CHECKSET,
}

func AlicloudPolardbAccountBasicDependence9211(name string) string {
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

// Test Polardb Account. <<< Resource test cases, automatically generated.

// Case  Account用例_PG_高权限账号 11819
func TestAccAliCloudPolarDbAccount_basic11819(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AliCloudPolarDbAccountMap11819)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudPolarDbAccountBasicDependence11819)
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
					"db_cluster_id":    "${alicloud_polardb_cluster.default.id}",
					"account_name":     name,
					"account_password": "YourPassword123!",
					"account_type":     "Super",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id": CHECKSET,
						"account_name":  name,
						"account_type":  "Super",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_password": "YourPassword123!update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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

func TestAccAliCloudPolarDbAccount_basic11819_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AliCloudPolarDbAccountMap11819)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudPolarDbAccountBasicDependence11819)
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
					"db_cluster_id":       "${alicloud_polardb_cluster.default.id}",
					"account_name":        name,
					"account_password":    "YourPassword123!",
					"account_type":        "Super",
					"account_description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":       CHECKSET,
						"account_name":        name,
						"account_type":        "Super",
						"account_description": name,
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

var AliCloudPolarDbAccountMap11819 = map[string]string{
	"account_lock_state":          CHECKSET,
	"account_password_valid_time": CHECKSET,
	"status":                      CHECKSET,
}

func AliCloudPolarDbAccountBasicDependence11819(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_polardb_node_classes" "default" {
  db_type    = "MySQL"
  db_version = "8.0"
  pay_type   = "PostPaid"
  category   = "Normal"
}

data "alicloud_vpcs" "default" {
  name_regex = "default-NODELETING"
}

data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_polardb_node_classes.default.classes.0.zone_id
}

resource "alicloud_polardb_cluster" "default" {
  db_type       = "MySQL"
  db_version    = "8.0"
  db_node_class = data.alicloud_polardb_node_classes.default.classes.0.supported_engines.0.available_resources.0.db_node_class
  pay_type      = "PostPaid"
  vswitch_id    = data.alicloud_vswitches.default.ids.0
}
`, name)
}

// Case  Account用例 9211
func TestAccAliCloudPolarDbAccount_basic9211(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AliCloudPolarDbAccountMap9211)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudPolarDbAccountBasicDependence9211)
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
					"db_cluster_id":          "${alicloud_polardb_cluster.default.id}",
					"account_name":           name,
					"kms_encrypted_password": "${alicloud_kms_ciphertext.default.ciphertext_blob}",
					"kms_encryption_context": map[string]string{
						"name": name,
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id": CHECKSET,
						"account_name":  name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"kms_encrypted_password": "${alicloud_kms_ciphertext.update.ciphertext_blob}",
					"kms_encryption_context": map[string]string{
						"name": name + "update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"kms_encrypted_password":   CHECKSET,
						"kms_encryption_context.%": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_lock_state": "Lock",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_lock_state": "Lock",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_password_valid_time": "2126-09-17T10:00:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_password_valid_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_lock_state": "UnLock",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_lock_state": "UnLock",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"kms_encrypted_password", "kms_encryption_context"},
			},
		},
	})
}

func TestAccAliCloudPolarDbAccount_basic9211_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_account.default"
	ra := resourceAttrInit(resourceId, AliCloudPolarDbAccountMap9211)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolarDbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudPolarDbAccountBasicDependence9211)
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
					"db_cluster_id":          "${alicloud_polardb_cluster.default.id}",
					"account_name":           name,
					"kms_encrypted_password": "${alicloud_kms_ciphertext.default.ciphertext_blob}",
					"kms_encryption_context": map[string]string{
						"name": name,
					},
					"account_type":                "Normal",
					"account_description":         name,
					"account_lock_state":          "Lock",
					"account_password_valid_time": "2126-09-17T10:00:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":               CHECKSET,
						"account_name":                name,
						"account_type":                "Normal",
						"account_description":         name,
						"account_lock_state":          "Lock",
						"account_password_valid_time": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"kms_encrypted_password", "kms_encryption_context"},
			},
		},
	})
}

var AliCloudPolarDbAccountMap9211 = map[string]string{
	"account_lock_state":          CHECKSET,
	"account_password_valid_time": CHECKSET,
	"account_type":                CHECKSET,
	"status":                      CHECKSET}

func AliCloudPolarDbAccountBasicDependence9211(name string) string {
	return fmt.Sprintf(`
variable "name" {
  default = "%s"
}

data "alicloud_polardb_node_classes" "default" {
  db_type  = "PostgreSQL"
  pay_type = "PostPaid"
  category = "Normal"
}

data "alicloud_vpcs" "default" {
  name_regex = "default-NODELETING"
}

data "alicloud_vswitches" "default" {
  vpc_id  = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_polardb_node_classes.default.classes.0.zone_id
}

resource "alicloud_polardb_cluster" "default" {
  db_version    = "14"
  pay_type      = "PostPaid"
  db_node_class = "polar.pg.x4.medium"
  db_type       = "PostgreSQL"
  vswitch_id    = data.alicloud_vswitches.default.ids.0
}

	resource "alicloud_kms_key" "default" {
  		description            = var.name
  		status                 = "Enabled"
  		pending_window_in_days = 7
	}

	resource "alicloud_kms_ciphertext" "default" {
  		key_id    = alicloud_kms_key.default.id
  		plaintext = "YourPassword1234!"
  		encryption_context = {
    		"name" = var.name
  		}
	}

	resource "alicloud_kms_ciphertext" "update" {
  		key_id    = alicloud_kms_key.default.id
  		plaintext = "YourPassword1234!update"
  		encryption_context = {
    		"name" = "${var.name}update"
  		}
	}
`, name)
}

// Test PolarDb Account. <<< Resource test cases, automatically generated.
