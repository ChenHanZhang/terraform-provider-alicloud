package alicloud

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("alicloud_adb_db_instance", &resource.Sweeper{
		Name: "alicloud_adb_db_instance",
		F:    testSweepAdbDbInstances,
	})
}

func testSweepAdbDbInstances(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting AliCloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}

	action := "DescribeDBClusters"
	request := make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1
	var response map[string]interface{}
	for {
		response, err = client.RpcPost("adb", "2019-03-15", action, nil, request, true)
		if err != nil {
			log.Println(WrapErrorf(err, DataDefaultErrorMsg, "alicloud_adb_db_clusters", action, AlibabaCloudSdkGoERROR))
			break
		}

		resp, err := jsonpath.Get("$.Items.DBCluster", response)
		if err != nil {
			log.Println(WrapErrorf(err, FailedGetAttributeMsg, action, "$.Items.DBCluster", response))
			break
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			name := fmt.Sprint(item["DBClusterDescription"])
			id := fmt.Sprint(item["DBClusterId"])
			skip := true
			if !sweepAll() {
				for _, prefix := range prefixes {
					if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
						skip = false
						break
					}
				}
				if skip {
					log.Printf("[INFO] Skipping ADB Instance: %s (%s)", name, id)
					continue
				}
			}
			log.Printf("[INFO] Deleting adb Instance: %s (%s)", name, id)
			action := "DeleteDBCluster"
			request := map[string]interface{}{
				"DBClusterId": id,
			}
			wait := incrementalWait(3*time.Second, 3*time.Second)
			err = resource.Retry(5*time.Minute, func() *resource.RetryError {
				_, err = client.RpcPost("adb", "2019-03-15", action, nil, request, true)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				log.Printf("[ERROR] Deleting ADB cluster failed with error: %#v", err)
				return nil
			})
		}
		if len(result) < PageSizeLarge {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	return nil
}

// Test Adb DbCluster. >>> Resource test cases, automatically generated.
// Case C8变E32测试用例正式 11900
func TestAccAliCloudAdbDbCluster_basic11900(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterMap11900)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterBasicDependence11900)
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
					"db_cluster_category":     "Cluster",
					"zone_id":                 "cn-beijing-k",
					"mode":                    "reserver",
					"db_cluster_version":      "3.0",
					"payment_type":            "Subscription",
					"disk_performance_level":  "PL1",
					"db_cluster_network_type": "vpc",
					"enable_ssl":              "false",
					"period":                  "Month",
					"used_time":               "1",
					"vswitch_id":              "${alicloud_vswitch.VSWITCHID.id}",
					"vpc_id":                  "${alicloud_vpc.VPCID.id}",
					"db_node_count":           "1",
					"db_node_storage":         "200",
					"renewal_status":          "Normal",
					"db_node_class":           "C8",
					"db_cluster_name":         name,
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"disk_encryption":         "false",
					"maintain_time":           "18:00Z-19:00Z",
					"kernel_version":          "3.2.4",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category":     "Cluster",
						"zone_id":                 "cn-beijing-k",
						"mode":                    "reserver",
						"db_cluster_version":      CHECKSET,
						"payment_type":            "Subscription",
						"disk_performance_level":  "PL1",
						"db_cluster_network_type": "vpc",
						"enable_ssl":              "false",
						"period":                  "Month",
						"used_time":               CHECKSET,
						"vswitch_id":              CHECKSET,
						"vpc_id":                  CHECKSET,
						"db_node_count":           "1",
						"db_node_storage":         "200",
						"renewal_status":          "Normal",
						"db_node_class":           "C8",
						"db_cluster_name":         name,
						"resource_group_id":       CHECKSET,
						"disk_encryption":         "false",
						"maintain_time":           "18:00Z-19:00Z",
						"kernel_version":          "3.2.4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"used_time":      "2",
					"renewal_status": "NotRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"used_time":      CHECKSET,
						"renewal_status": "NotRenewal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_category":      "MIXED_STORAGE",
					"mode":                     "flexible",
					"payment_type":             "PayAsYouGo",
					"db_node_class":            "E32",
					"db_cluster_name":          name + "_update",
					"resource_group_id":        "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"maintain_time":            "22:00Z-23:00Z",
					"switch_mode":              "0",
					"elastic_io_resource":      "1",
					"executor_count":           "3",
					"auto_renew_period":        "2",
					"elastic_io_resource_size": "8Core64GB",
					"compute_resource":         "48Core192GBNEW",
					"period_unit":              "Month",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category":      "MIXED_STORAGE",
						"mode":                     "flexible",
						"payment_type":             "PayAsYouGo",
						"db_node_class":            "E32",
						"db_cluster_name":          name + "_update",
						"resource_group_id":        CHECKSET,
						"maintain_time":            "22:00Z-23:00Z",
						"switch_mode":              "0",
						"elastic_io_resource":      CHECKSET,
						"executor_count":           CHECKSET,
						"auto_renew_period":        "2",
						"elastic_io_resource_size": "8Core64GB",
						"compute_resource":         "48Core192GBNEW",
						"period_unit":              "Month",
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
				ImportStateVerifyIgnore: []string{"auto_renew_period", "backup_set_id", "connection_string", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "modify_mode", "modify_type", "period", "period_unit", "restore_time", "restore_type", "source_db_instance_name", "storage_type", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterMap11900 = map[string]string{
	"port":        CHECKSET,
	"status":      CHECKSET,
	"pay_type":    CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudAdbDbClusterBasicDependence11900(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "adb北京vpc-tf-test"
}

resource "alicloud_vswitch" "VSWITCHID" {
  description  = "北京测试-tf"
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-beijing-k"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "北京k测试tf"
}


`, name)
}

// Case C8磁盘变配相关测试用例 11858
func TestAccAliCloudAdbDbCluster_basic11858(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterMap11858)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterBasicDependence11858)
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
					"db_cluster_category":     "Cluster",
					"zone_id":                 "cn-beijing-k",
					"mode":                    "reserver",
					"db_cluster_version":      "3.0",
					"payment_type":            "Subscription",
					"disk_performance_level":  "PL1",
					"db_cluster_network_type": "vpc",
					"enable_ssl":              "false",
					"period":                  "Month",
					"used_time":               "1",
					"vswitch_id":              "${alicloud_vswitch.VSWITCHID.id}",
					"vpc_id":                  "${alicloud_vpc.VPCID.id}",
					"db_node_count":           "1",
					"db_node_storage":         "200",
					"renewal_status":          "Normal",
					"db_node_class":           "C8",
					"db_cluster_name":         name,
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"disk_encryption":         "false",
					"maintain_time":           "18:00Z-19:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category":     "Cluster",
						"zone_id":                 "cn-beijing-k",
						"mode":                    "reserver",
						"db_cluster_version":      CHECKSET,
						"payment_type":            "Subscription",
						"disk_performance_level":  "PL1",
						"db_cluster_network_type": "vpc",
						"enable_ssl":              "false",
						"period":                  "Month",
						"used_time":               CHECKSET,
						"vswitch_id":              CHECKSET,
						"vpc_id":                  CHECKSET,
						"db_node_count":           "1",
						"db_node_storage":         "200",
						"renewal_status":          "Normal",
						"db_node_class":           "C8",
						"db_cluster_name":         name,
						"resource_group_id":       CHECKSET,
						"disk_encryption":         "false",
						"maintain_time":           "18:00Z-19:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_storage": "500",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_storage": "500",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_count": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_count": "2",
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
				ImportStateVerifyIgnore: []string{"auto_renew_period", "backup_set_id", "connection_string", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "modify_mode", "modify_type", "period", "period_unit", "restore_time", "restore_type", "source_db_instance_name", "storage_type", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterMap11858 = map[string]string{
	"port":        CHECKSET,
	"status":      CHECKSET,
	"pay_type":    CHECKSET,
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudAdbDbClusterBasicDependence11858(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "adb北京vpc-tf-test"
}

resource "alicloud_vswitch" "VSWITCHID" {
  description  = "北京测试-tf"
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-beijing-k"
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "北京k测试tf"
}


`, name)
}

// Test Adb DbCluster. <<< Resource test cases, automatically generated.

func TestAccAliCloudADBDbCluster_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, false, connectivity.AdbReserverUnSupportRegions)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_category": "Cluster",
					"db_node_class":       "C8",
					"description":         name,
					"db_node_count":       "1",
					"db_node_storage":     "100",
					"mode":                "reserver",
					"vswitch_id":          "${local.vswitch_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category": "Cluster",
						"db_node_class":       "C8",
						"description":         name,
						"db_node_count":       "1",
						"db_node_storage":     "100",
						"mode":                "reserver",
						"vswitch_id":          CHECKSET,
						"kernel_version":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_class": "C32",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_class": "C32",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_count": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_count": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_storage": "200",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_storage": "200",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"maintain_time": "23:00Z-00:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"maintain_time": "23:00Z-00:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_ips": []string{"10.168.1.12"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ips.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type":      "Subscription",
					"period":            "1",
					"renewal_status":    "AutoRenewal",
					"auto_renew_period": "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":      "Subscription",
						"renewal_status":    "AutoRenewal",
						"auto_renew_period": "2",
						"pay_type":          "PrePaid",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type": "PayAsYouGo",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type": "PayAsYouGo",
						"pay_type":     "PostPaid",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_class":   "C8",
					"db_node_count":   "1",
					"db_node_storage": "100",
					"description":     name,
					"maintain_time":   "01:00Z-02:00Z",
					"security_ips":    []string{"10.168.1.13"},
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_class":   "C8",
						"db_node_count":   "1",
						"db_node_storage": "100",
						"description":     name,
						"maintain_time":   "01:00Z-02:00Z",
						"security_ips.#":  "1",
						"tags.%":          "2",
						"tags.Created":    "TF-update",
						"tags.For":        "test-update",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_renew_period", "modify_type", "period", "renewal_status"},
			},
		},
	})
}

func TestAccAliCloudADBDbCluster_flexible8C(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterMap1)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterBasicDependence0)
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
					"db_cluster_category": "MixedStorage",
					"description":         name,
					"mode":                "flexible",
					"compute_resource":    "8Core32GB",
					"vswitch_id":          "${local.vswitch_id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category": "MixedStorage",
						"description":         name,
						"mode":                "flexible",
						"compute_resource":    "8Core32GB",
						"vswitch_id":          CHECKSET,
						"kernel_version":      CHECKSET,
					}),
				),
			},
			// API does not support to updating the compute_resource
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"compute_resource": "16Core64GB",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"compute_resource": "16Core64GB",
			//		}),
			//	),
			//},
			// API does not support updating elastic_io_resource when compute_resource is 8Core32GB or 16Core64GB
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"elastic_io_resource": "1",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"elastic_io_resource": "1",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"maintain_time": "23:00Z-00:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"maintain_time": "23:00Z-00:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"kernel_version": "3.2.1",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"kernel_version": "3.2.1",
			//		}),
			//	),
			//},
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
					"security_ips": []string{"10.168.1.12"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ips.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource": "8Core32GB",
					// API does not support updating elastic_io_resource when compute_resource is 8Core32GB or 16Core64GB
					//"elastic_io_resource": "1",
					"description":   name,
					"maintain_time": "01:00Z-02:00Z",
					"security_ips":  []string{"10.168.1.13"},
					"tags": map[string]string{
						"Created": "TF-update",
						"For":     "test-update",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource": "8Core32GB",
						//"elastic_io_resource": "1",
						"description":    name,
						"maintain_time":  "01:00Z-02:00Z",
						"security_ips.#": "1",
						"tags.%":         "2",
						"tags.Created":   "TF-update",
						"tags.For":       "test-update",
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

func TestAccAliCloudADBDbCluster_flexible32C(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_db_cluster.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterMap1)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterBasicDependence0)
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
					"db_cluster_category": "MixedStorage",
					"description":         name,
					"mode":                "flexible",
					"compute_resource":    "32Core128GBNEW",
					"elastic_io_resource": "1",
					"vswitch_id":          "${local.vswitch_id}",
					"enable_ssl":          "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category": "MixedStorage",
						"description":         name,
						"mode":                "flexible",
						"compute_resource":    "32Core128GBNEW",
						"elastic_io_resource": "1",
						"vswitch_id":          CHECKSET,
						"enable_ssl":          "true",
						"db_node_class":       "E32",
					}),
				),
			},
			// API does not support updating elastic_io_resource when compute_resource is 32Core128GBNEW
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"compute_resource": "48Core192GBNEW",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"compute_resource": "48Core192GBNEW",
			//			"db_node_count":    CHECKSET,
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"elastic_io_resource": "2",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"elastic_io_resource": "2",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": name + "update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": name + "update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"maintain_time": "23:00Z-00:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"maintain_time": "23:00Z-00:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_ips": []string{"10.168.1.12"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ips.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test",
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

func TestAccAliCloudADBDbCluster_basic1(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_db_cluster.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterMap2)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sadbCluster%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterBasicDependence1)
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
					"db_cluster_category":      "MixedStorage",
					"description":              name,
					"mode":                     "flexible",
					"compute_resource":         "32Core128GBNEW",
					"vswitch_id":               "${local.vswitch_id}",
					"vpc_id":                   "${data.alicloud_vpcs.default.ids.0}",
					"elastic_io_resource":      "1",
					"disk_encryption":          "true",
					"elastic_io_resource_size": "8Core64GB",
					"disk_performance_level":   "PL1",
					"kms_id":                   "${alicloud_kms_key.default.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_category":      "MixedStorage",
						"description":              name,
						"mode":                     "flexible",
						"compute_resource":         "32Core128GBNEW",
						"elastic_io_resource_size": "8Core64GB",
						"disk_performance_level":   "PL1",
						"vswitch_id":               CHECKSET,
						"vpc_id":                   CHECKSET,
						"db_node_class":            "E32",
						"elastic_io_resource":      "1",
						"disk_encryption":          "true",
						"kms_id":                   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"elastic_io_resource_size": "12Core96GB",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"elastic_io_resource_size": "12Core96GB",
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

var AliCloudAdbDbClusterMap0 = map[string]string{
	"auto_renew_period":   NOSET,
	"compute_resource":    "",
	"connection_string":   CHECKSET,
	"port":                CHECKSET,
	"db_cluster_version":  "3.0",
	"db_node_storage":     "0",
	"elastic_io_resource": "0",
	"maintain_time":       CHECKSET,
	"modify_type":         NOSET,
	"payment_type":        "PayAsYouGo",
	"pay_type":            "PostPaid",
	"renewal_status":      NOSET,
	"resource_group_id":   CHECKSET,
	"security_ips.#":      "1",
	"status":              "Running",
	"tags.%":              "0",
	"zone_id":             CHECKSET,
}

var AliCloudAdbDbClusterMap1 = map[string]string{
	"auto_renew_period": NOSET,
	//"compute_resource": "8c16g",
	"connection_string":   CHECKSET,
	"port":                CHECKSET,
	"db_cluster_version":  "3.0",
	"db_node_class":       "E8",
	"db_node_count":       "1",
	"db_node_storage":     "100",
	"elastic_io_resource": "0",
	"maintain_time":       CHECKSET,
	"modify_type":         NOSET,
	"payment_type":        "PayAsYouGo",
	"pay_type":            "PostPaid",
	"renewal_status":      NOSET,
	"resource_group_id":   CHECKSET,
	"security_ips.#":      "1",
	"status":              "Running",
	"tags.%":              "0",
	"zone_id":             CHECKSET,
}

var AliCloudAdbDbClusterMap2 = map[string]string{}

func AliCloudAdbDbClusterBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "%s"
	}

	data "alicloud_resource_manager_resource_groups" "default" {
	}
	%s
`, name, AdbCommonTestCase)
}

func AliCloudAdbDbClusterBasicDependence1(name string) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "%s"
	}

	data "alicloud_resource_manager_resource_groups" "default" {
	}

	resource "alicloud_kms_key" "default" {
  		description            = var.name
  		pending_window_in_days = "7"
  		status                 = "Enabled"
	}
	%s
`, name, AdbCommonTestCase)
}
