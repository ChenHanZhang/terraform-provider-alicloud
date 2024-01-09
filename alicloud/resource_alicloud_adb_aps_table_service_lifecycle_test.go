package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb ApsTableServiceLifecycle. >>> Resource test cases, automatically generated.
// Case 5712
func TestAccAliCloudAdbApsTableServiceLifecycle_basic5712(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_aps_table_service_lifecycle.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbApsTableServiceLifecycleMap5712)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbApsTableServiceLifecycle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbapstableservicelifecycle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbApsTableServiceLifecycleBasicDependence5712)
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
					"strategy_type":  "KEEP_BY_TIME",
					"strategy_name":  "Lifecycle-202312270000001",
					"strategy_value": "100",
					"db_cluster_id":  "${var.db_cluster_id}",
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_type":      "KEEP_BY_TIME",
						"strategy_name":      "Lifecycle-202312270000001",
						"strategy_value":     "100",
						"db_cluster_id":      CHECKSET,
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_desc": "100天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_desc": "100天生命周期管理",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_value": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_value": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "off",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_value": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_value": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_desc": "1000天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_desc": "1000天生命周期管理",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"process_all": "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":         "on",
					"strategy_type":  "KEEP_BY_TIME",
					"strategy_name":  "Lifecycle-202312270000001",
					"strategy_value": "100",
					"db_cluster_id":  "${var.db_cluster_id}",
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
						},
					},
					"strategy_desc": "100天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":             "on",
						"strategy_type":      "KEEP_BY_TIME",
						"strategy_name":      "Lifecycle-202312270000001",
						"strategy_value":     "100",
						"db_cluster_id":      CHECKSET,
						"operation_tables.#": "1",
						"strategy_desc":      "100天生命周期管理",
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

var AlicloudAdbApsTableServiceLifecycleMap5712 = map[string]string{
	"status":      "on",
	"create_time": CHECKSET,
	"aps_job_id":  CHECKSET,
}

func AlicloudAdbApsTableServiceLifecycleBasicDependence5712(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "db_cluster_id" {
  default = "amv-bp1u30028ta370f7"
}


`, name)
}

// Case 5524
func TestAccAliCloudAdbApsTableServiceLifecycle_basic5524(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_aps_table_service_lifecycle.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbApsTableServiceLifecycleMap5524)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbApsTableServiceLifecycle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbapstableservicelifecycle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbApsTableServiceLifecycleBasicDependence5524)
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
					"strategy_type":  "KEEP_BY_TIME",
					"strategy_name":  "Lifecycle-202312270000001",
					"strategy_value": "100",
					"db_cluster_id":  "${var.db_cluster_id}",
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
							"process_all":   "true",
						},
					},
					"status": "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_type":      "KEEP_BY_TIME",
						"strategy_name":      "Lifecycle-202312270000001",
						"strategy_value":     "100",
						"db_cluster_id":      CHECKSET,
						"operation_tables.#": "1",
						"status":             "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_desc": "100天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_desc": "100天生命周期管理",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_value": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_value": "100",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
							"process_all":   "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "off",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_value": "1000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_value": "1000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"strategy_desc": "1000天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"strategy_desc": "1000天生命周期管理",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"operation_tables": []map[string]interface{}{
						{
							"process_all": "true",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"operation_tables.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":         "on",
					"strategy_type":  "KEEP_BY_TIME",
					"strategy_name":  "Lifecycle-202312270000001",
					"strategy_value": "100",
					"db_cluster_id":  "${var.db_cluster_id}",
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
							"process_all":   "true",
						},
					},
					"strategy_desc": "100天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":             "on",
						"strategy_type":      "KEEP_BY_TIME",
						"strategy_name":      "Lifecycle-202312270000001",
						"strategy_value":     "100",
						"db_cluster_id":      CHECKSET,
						"operation_tables.#": "1",
						"strategy_desc":      "100天生命周期管理",
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

var AlicloudAdbApsTableServiceLifecycleMap5524 = map[string]string{
	"status":      "on",
	"create_time": CHECKSET,
	"aps_job_id":  CHECKSET,
}

func AlicloudAdbApsTableServiceLifecycleBasicDependence5524(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "db_cluster_id" {
  default = "amv-bp1u30028ta370f7"
}


`, name)
}

// Case 5712  twin
func TestAccAliCloudAdbApsTableServiceLifecycle_basic5712_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_aps_table_service_lifecycle.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbApsTableServiceLifecycleMap5712)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbApsTableServiceLifecycle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbapstableservicelifecycle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbApsTableServiceLifecycleBasicDependence5712)
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
					"status":         "on",
					"strategy_type":  "KEEP_BY_TIME",
					"strategy_name":  "Lifecycle-202312270000001",
					"strategy_value": "1000",
					"db_cluster_id":  "${var.db_cluster_id}",
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
							"process_all":   "true",
						},
						{
							"table_names": []string{
								"schema002_tb"},
							"database_name": "schema002_db",
							"process_all":   "true",
						},
						{
							"table_names": []string{
								"schema003_tb"},
							"database_name": "schema003_db",
						},
					},
					"strategy_desc": "1000天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":             "on",
						"strategy_type":      "KEEP_BY_TIME",
						"strategy_name":      "Lifecycle-202312270000001",
						"strategy_value":     "1000",
						"db_cluster_id":      CHECKSET,
						"operation_tables.#": "3",
						"strategy_desc":      "1000天生命周期管理",
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

// Case 5524  twin
func TestAccAliCloudAdbApsTableServiceLifecycle_basic5524_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_aps_table_service_lifecycle.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbApsTableServiceLifecycleMap5524)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbApsTableServiceLifecycle")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbapstableservicelifecycle%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbApsTableServiceLifecycleBasicDependence5524)
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
					"status":         "on",
					"strategy_type":  "KEEP_BY_TIME",
					"strategy_name":  "Lifecycle-202312270000001",
					"strategy_value": "1000",
					"db_cluster_id":  "${var.db_cluster_id}",
					"operation_tables": []map[string]interface{}{
						{
							"table_names": []string{
								"schema001_tb"},
							"database_name": "schema001_db",
							"process_all":   "true",
						},
						{
							"table_names": []string{
								"schema002_tb"},
							"database_name": "schema002_db",
							"process_all":   "true",
						},
						{
							"table_names": []string{
								"schema003_tb"},
							"database_name": "schema003_db",
							"process_all":   "true",
						},
					},
					"strategy_desc": "1000天生命周期管理",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":             "on",
						"strategy_type":      "KEEP_BY_TIME",
						"strategy_name":      "Lifecycle-202312270000001",
						"strategy_value":     "1000",
						"db_cluster_id":      CHECKSET,
						"operation_tables.#": "3",
						"strategy_desc":      "1000天生命周期管理",
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

// Test Adb ApsTableServiceLifecycle. <<< Resource test cases, automatically generated.
