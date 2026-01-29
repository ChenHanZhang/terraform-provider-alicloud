package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestAccAliCloudADBDBClusterLakeVersion_basic0(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterLakeVersionMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%sadbdbclusterlakeversion%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterLakeVersionBasicDependence0)
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
					"db_cluster_version": "5.0",
					"vpc_id":             "${data.alicloud_vpcs.default.ids.0}",
					"vswitch_id":         "${data.alicloud_vswitches.default.ids.0}",
					"zone_id":            "${data.alicloud_adb_zones.default.ids.0}",
					"payment_type":       "PayAsYouGo",
					"compute_resource":   "16ACU",
					"storage_resource":   "0ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_version": "5.0",
						"vpc_id":             CHECKSET,
						"vswitch_id":         CHECKSET,
						"zone_id":            CHECKSET,
						"payment_type":       "PayAsYouGo",
						"compute_resource":   "16ACU",
						"storage_resource":   "0ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource": "32ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource": "32ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource": "24ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource": "24ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource": "16ACU",
					"storage_resource": "0ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource": "16ACU",
						"storage_resource": "0ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_ips": "10.23.1.1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ips": "10.23.1.1",
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
					"db_cluster_description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"enable_default_resource_group", "source_db_cluster_id", "backup_set_id", "restore_type", "restore_to_time"},
			},
		},
	})
}

func TestAccAliCloudADBDBClusterLakeVersion_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterLakeVersionMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%sadbdbclusterlakeversion%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterLakeVersionBasicDependence0Twin)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  nil,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_version":            "5.0",
					"vpc_id":                        "${data.alicloud_vpcs.default.ids.0}",
					"vswitch_id":                    "${data.alicloud_vswitches.default.ids.0}",
					"zone_id":                       "${data.alicloud_adb_zones.default.ids.0}",
					"payment_type":                  "Subscription",
					"secondary_vswitch_id":          "${data.alicloud_vswitches.secondary.ids.0}",
					"secondary_zone_id":             "${data.alicloud_adb_zones.default.ids.1}",
					"compute_resource":              "16ACU",
					"storage_resource":              "0ACU",
					"disk_encryption":               "true",
					"kms_id":                        "${alicloud_kms_key.default.id}",
					"security_ips":                  "10.23.1.0",
					"enable_ssl":                    "true",
					"db_cluster_description":        name,
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
					"period":                        "1",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_version":     "5.0",
						"vpc_id":                 CHECKSET,
						"vswitch_id":             CHECKSET,
						"zone_id":                CHECKSET,
						"payment_type":           "Subscription",
						"secondary_vswitch_id":   CHECKSET,
						"secondary_zone_id":      CHECKSET,
						"compute_resource":       "16ACU",
						"storage_resource":       "0ACU",
						"disk_encryption":        "true",
						"kms_id":                 CHECKSET,
						"security_ips":           "10.23.1.0",
						"enable_ssl":             "true",
						"db_cluster_description": name,
						"resource_group_id":      CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period", "enable_default_resource_group", "source_db_cluster_id", "backup_set_id", "restore_type", "restore_to_time"},
			},
		},
	})
}

func TestAccAliCloudADBDBClusterLakeVersion_basic1(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterLakeVersionMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%sadbdbclusterlakeversion%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterLakeVersionBasicDependence0)
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
					"db_cluster_version":  "5.0",
					"vpc_id":              "${data.alicloud_vpcs.default.ids.0}",
					"vswitch_id":          "${data.alicloud_vswitches.default.ids.0}",
					"zone_id":             "${data.alicloud_adb_zones.default.ids.0}",
					"payment_type":        "PayAsYouGo",
					"product_form":        "IntegrationForm",
					"product_version":     "EnterpriseVersion",
					"reserved_node_size":  "8ACU",
					"reserved_node_count": "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_version":  "5.0",
						"vpc_id":              CHECKSET,
						"vswitch_id":          CHECKSET,
						"zone_id":             CHECKSET,
						"payment_type":        "PayAsYouGo",
						"product_form":        "IntegrationForm",
						"product_version":     "EnterpriseVersion",
						"reserved_node_size":  "8ACU",
						"reserved_node_count": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"reserved_node_size": "12ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"reserved_node_size": "12ACU",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"reserved_node_count": "6",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"reserved_node_count": "6",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"reserved_node_size":  "8ACU",
					"reserved_node_count": "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"reserved_node_size":  "8ACU",
						"reserved_node_count": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"security_ips": "10.23.1.1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"security_ips": "10.23.1.1",
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
					"db_cluster_description": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"enable_default_resource_group", "source_db_cluster_id", "backup_set_id", "restore_type", "restore_to_time"},
			},
		},
	})
}

func TestAccAliCloudADBDBClusterLakeVersion_basic1_twin(t *testing.T) {
	var v map[string]interface{}
	checkoutSupportedRegions(t, true, connectivity.TestSalveRegions)
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AliCloudAdbDbClusterLakeVersionMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1000, 9999)
	name := fmt.Sprintf("tf-testacc%sadbdbclusterlakeversion%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudAdbDbClusterLakeVersionBasicDependence0Twin)
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
					"db_cluster_version":            "5.0",
					"vpc_id":                        "${data.alicloud_vpcs.default.ids.0}",
					"vswitch_id":                    "${data.alicloud_vswitches.default.ids.0}",
					"zone_id":                       "${data.alicloud_adb_zones.default.ids.0}",
					"payment_type":                  "PayAsYouGo",
					"product_form":                  "IntegrationForm",
					"product_version":               "EnterpriseVersion",
					"reserved_node_size":            "8ACU",
					"reserved_node_count":           "3",
					"disk_encryption":               "true",
					"kms_id":                        "${alicloud_kms_key.default.id}",
					"security_ips":                  "10.23.1.0",
					"enable_ssl":                    "true",
					"db_cluster_description":        name,
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.groups.1.id}",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_version":     "5.0",
						"vpc_id":                 CHECKSET,
						"vswitch_id":             CHECKSET,
						"zone_id":                CHECKSET,
						"payment_type":           "PayAsYouGo",
						"product_form":           "IntegrationForm",
						"product_version":        "EnterpriseVersion",
						"reserved_node_size":     "8ACU",
						"reserved_node_count":    "3",
						"disk_encryption":        "true",
						"kms_id":                 CHECKSET,
						"security_ips":           "10.23.1.0",
						"enable_ssl":             "true",
						"db_cluster_description": name,
						"resource_group_id":      CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"enable_default_resource_group", "source_db_cluster_id", "backup_set_id", "restore_type", "restore_to_time"},
			},
		},
	})
}

var AliCloudAdbDbClusterLakeVersionMap0 = map[string]string{
	"product_form":           CHECKSET,
	"product_version":        CHECKSET,
	"reserved_node_size":     CHECKSET,
	"security_ips":           CHECKSET,
	"db_cluster_description": CHECKSET,
	"commodity_code":         CHECKSET,
	"connection_string":      CHECKSET,
	"engine":                 CHECKSET,
	"engine_version":         CHECKSET,
	"create_time":            CHECKSET,
	"lock_mode":              CHECKSET,
	"port":                   CHECKSET,
	"resource_group_id":      CHECKSET,
	"status":                 CHECKSET,
}

func AliCloudAdbDbClusterLakeVersionBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_resource_manager_resource_groups" "default" {
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
`, name)
}

func AliCloudAdbDbClusterLakeVersionBasicDependence0Twin(name string) string {
	return fmt.Sprintf(`
	variable "name" {
  		default = "%s"
	}

	data "alicloud_resource_manager_resource_groups" "default" {
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

	data "alicloud_vswitches" "secondary" {
  		vpc_id  = data.alicloud_vpcs.default.ids.0
  		zone_id = data.alicloud_adb_zones.default.ids.1
	}

	resource "alicloud_kms_key" "default" {
  		description            = var.name
  		pending_window_in_days = "7"
  		status                 = "Enabled"
	}
`, name)
}

func TestUnitAliCloudAdbDbClusterLakeVersion(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_adb_db_cluster_lake_version"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_adb_db_cluster_lake_version"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"compute_resource":              "CreateDBClusterValue",
		"db_cluster_version":            "5.0",
		"enable_default_resource_group": true,
		"payment_type":                  "CreateDBClusterValue",
		"storage_resource":              "CreateDBClusterValue",
		"vswitch_id":                    "CreateDBClusterValue",
		"zone_id":                       "CreateDBClusterValue",
	}
	for key, value := range attributes {
		err := dInit.Set(key, value)
		assert.Nil(t, err)
		err = dExisted.Set(key, value)
		assert.Nil(t, err)
		if err != nil {
			log.Printf("[ERROR] the field %s setting error", key)
		}
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		// DescribeDBClusterAttribute
		"Items": map[string]interface{}{
			"DBCluster": []interface{}{
				map[string]interface{}{
					"CommodityCode":    "DefaultValue",
					"ComputeResource":  "CreateDBClusterValue",
					"ConnectionString": "DefaultValue",
					"CreationTime":     "DefaultValue",
					"DBClusterId":      "CreateDBClusterValue",
					"DBVersion":        "5.0",
					"Engine":           "DefaultValue",
					"EngineVersion":    "DefaultValue",
					"ExpireTime":       "DefaultValue",
					"Expired":          "DefaultValue",
					"LockMode":         "DefaultValue",
					"LockReason":       "DefaultValue",
					"PayType":          "CreateDBClusterValue",
					"Port":             "DefaultValue",
					"ResourceGroupId":  "DefaultValue",
					"DBClusterStatus":  "Running",
					"StorageResource":  "CreateDBClusterValue",
					"VPCId":            "CreateDBClusterValue",
					"VSwitchId":        "CreateDBClusterValue",
					"ZoneId":           "CreateDBClusterValue",
				},
			},
		},
	}
	CreateMockResponse := map[string]interface{}{
		// CreateDBCluster
		"DBClusterId":     "CreateDBClusterValue",
		"OrderId":         "MockValue",
		"RequestId":       "MockValue",
		"ResourceGroupId": "MockValue",
	}
	failedResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, &tea.SDKError{
			Code:       String(errorCode),
			Data:       String(errorCode),
			Message:    String(errorCode),
			StatusCode: tea.Int(400),
		}
	}
	notFoundResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_adb_db_cluster_lake_version", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewAdsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudAdbDbClusterLakeVersionCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// DescribeDBClusterAttribute Response
		"Items": map[string]interface{}{
			"DBCluster": []interface{}{
				map[string]interface{}{
					"DBClusterId": "CreateDBClusterValue",
				},
			},
		},
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateDBCluster" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						successResponseMock(ReadMockResponseDiff)
						return CreateMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudAdbDbClusterLakeVersionCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_adb_db_cluster_lake_version"].Schema).Data(dInit.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dInit.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Update
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewAdsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudAdbDbClusterLakeVersionUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// ModifyDBCluster
	attributesDiff := map[string]interface{}{
		"compute_resource": "ModifyDBClusterValue",
		"storage_resource": "ModifyDBClusterValue",
	}
	diff, err := newInstanceDiff("alicloud_adb_db_cluster_lake_version", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_adb_db_cluster_lake_version"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// DescribeDBClusterAttribute Response
		"Items": map[string]interface{}{
			"DBCluster": []interface{}{
				map[string]interface{}{
					"ComputeResource": "ModifyDBClusterValue",
					"StorageResource": "ModifyDBClusterValue",
				},
			},
		},
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyDBCluster" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudAdbDbClusterLakeVersionUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_adb_db_cluster_lake_version"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Read
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "InvalidDBCluster.NotFound", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DescribeDBClusterAttribute" {
				switch errorCode {
				case "{}", "InvalidDBCluster.NotFound":
					return notFoundResponseMock(errorCode)
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudAdbDbClusterLakeVersionRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}", "InvalidDBCluster.NotFound":
			assert.Nil(t, err)
		}
	}

	// Delete
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewAdsClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:    String("loadEndpoint error"),
			Data:    String("loadEndpoint error"),
			Message: String("loadEndpoint error"),
		}
	})
	err = resourceAliCloudAdbDbClusterLakeVersionDelete(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	errorCodes = []string{"NonRetryableError", "Throttling", "nil", "InvalidDBCluster.NotFound"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "DeleteDBCluster" {
				switch errorCode {
				case "NonRetryableError", "InvalidDBCluster.NotFound":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						ReadMockResponse = map[string]interface{}{}
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudAdbDbClusterLakeVersionDelete(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "InvalidDBCluster.NotFound":
			assert.Nil(t, err)
		}
	}
}

// Test Adb DbClusterLakeVersion. >>> Resource test cases, automatically generated.
// Case adb湖仓测试用例 3510
func TestAccAliCloudAdbDbClusterLakeVersion_basic3510(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap3510)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence3510)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-beijing-k",
					"vpc_id":                        "vpc-2ze9l3ppzpg0ueg7rkkmi",
					"vswitch_id":                    "vsw-2zeph51w1r0r9qjzxe933",
					"db_cluster_description":        "tf自动化测试",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "Postpaid",
					"enable_default_resource_group": "false",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"period":                        "Month",
					"used_time":                     "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-beijing-k",
						"vpc_id":                        "vpc-2ze9l3ppzpg0ueg7rkkmi",
						"vswitch_id":                    "vsw-2zeph51w1r0r9qjzxe933",
						"db_cluster_description":        "tf自动化测试",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "Postpaid",
						"enable_default_resource_group": "false",
						"resource_group_id":             CHECKSET,
						"period":                        "Month",
						"used_time":                     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":         "24ACU",
					"db_cluster_description":   "测试修改描述",
					"compute_resource":         "32ACU",
					"security_ips":             "123.23.11.2",
					"db_cluster_ip_array_name": "default",
					"modify_mode":              "Cover",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":         "24ACU",
						"db_cluster_description":   "测试修改描述",
						"compute_resource":         "32ACU",
						"security_ips":             "123.23.11.2",
						"db_cluster_ip_array_name": "default",
						"modify_mode":              "Cover",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource": "48ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource": "48ACU",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap3510 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence3510(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case adb湖仓测试用例_ver2 3562
func TestAccAliCloudAdbDbClusterLakeVersion_basic3562(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap3562)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence3562)
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
					"storage_resource":        "0ACU",
					"zone_id":                 "cn-beijing-k",
					"vpc_id":                  "${alicloud_vpc.VPC ID.id}",
					"vswitch_id":              "${alicloud_vswitch.VSWITCH ID.id}",
					"db_cluster_description":  "tf自动化测试",
					"compute_resource":        "16ACU",
					"db_cluster_network_type": "VPC",
					"db_cluster_version":      "5.0",
					"payment_type":            "Postpaid",
					"period":                  "Month",
					"used_time":               "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":        "0ACU",
						"zone_id":                 "cn-beijing-k",
						"vpc_id":                  CHECKSET,
						"vswitch_id":              CHECKSET,
						"db_cluster_description":  "tf自动化测试",
						"compute_resource":        "16ACU",
						"db_cluster_network_type": "VPC",
						"db_cluster_version":      CHECKSET,
						"payment_type":            "Postpaid",
						"period":                  "Month",
						"used_time":               CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":         "24ACU",
					"db_cluster_description":   "测试修改描述",
					"compute_resource":         "32ACU",
					"security_ips":             "123.23.11.2",
					"db_cluster_ip_array_name": "default",
					"modify_mode":              "Cover",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":         "24ACU",
						"db_cluster_description":   "测试修改描述",
						"compute_resource":         "32ACU",
						"security_ips":             "123.23.11.2",
						"db_cluster_ip_array_name": "default",
						"modify_mode":              "Cover",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource": "48ACU",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource": "48ACU",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap3562 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence3562(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "VPC ID" {
        dry_run = false
        enable_ipv6 = false
        vpc_name = "adb测试北京vpc"
        cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCH ID" {
        vpc_id = "${alicloud_vpc.VPC ID.id}"
        zone_id = "cn-beijing-k"
        vswitch_name = "北京k测试"
        cidr_block = "172.16.0.0/24"
}


`, name)
}

// Case adb湖仓测试用例_ver3 3575
func TestAccAliCloudAdbDbClusterLakeVersion_basic3575(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap3575)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence3575)
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
					"storage_resource":        "0ACU",
					"zone_id":                 "cn-beijing-k",
					"vpc_id":                  "${alicloud_vpc.VPC ID.id}",
					"vswitch_id":              "${alicloud_vswitch.VSWITCH ID.id}",
					"db_cluster_description":  "tf自动化测试",
					"compute_resource":        "16ACU",
					"db_cluster_network_type": "VPC",
					"db_cluster_version":      "5.0",
					"payment_type":            "Postpaid",
					"period":                  "Month",
					"used_time":               "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":        "0ACU",
						"zone_id":                 "cn-beijing-k",
						"vpc_id":                  CHECKSET,
						"vswitch_id":              CHECKSET,
						"db_cluster_description":  "tf自动化测试",
						"compute_resource":        "16ACU",
						"db_cluster_network_type": "VPC",
						"db_cluster_version":      CHECKSET,
						"payment_type":            "Postpaid",
						"period":                  "Month",
						"used_time":               CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":         "24ACU",
					"db_cluster_description":   "测试修改描述",
					"compute_resource":         "32ACU",
					"security_ips":             "123.23.11.2",
					"db_cluster_ip_array_name": "default",
					"modify_mode":              "Cover",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":         "24ACU",
						"db_cluster_description":   "测试修改描述",
						"compute_resource":         "32ACU",
						"security_ips":             "123.23.11.2",
						"db_cluster_ip_array_name": "default",
						"modify_mode":              "Cover",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource":  "48ACU",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource":  "48ACU",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap3575 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence3575(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPC ID" {
        dry_run = false
        enable_ipv6 = false
        vpc_name = "adb测试北京vpc"
        cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCH ID" {
        vpc_id = "${alicloud_vpc.VPC ID.id}"
        zone_id = "cn-beijing-k"
        vswitch_name = "北京k测试"
        cidr_block = "172.16.0.0/24"
}


`, name)
}

// Case adb湖仓测试用例_ver6 3713
func TestAccAliCloudAdbDbClusterLakeVersion_basic3713(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap3713)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence3713)
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
					"storage_resource":        "0ACU",
					"zone_id":                 "cn-beijing-k",
					"vpc_id":                  "${alicloud_vpc.VPC ID.id}",
					"vswitch_id":              "${alicloud_vswitch.VSWITCH ID.id}",
					"db_cluster_description":  "tf自动化测试",
					"compute_resource":        "16ACU",
					"db_cluster_network_type": "VPC",
					"db_cluster_version":      "5.0",
					"payment_type":            "Postpaid",
					"period":                  "Month",
					"used_time":               "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":        "0ACU",
						"zone_id":                 "cn-beijing-k",
						"vpc_id":                  CHECKSET,
						"vswitch_id":              CHECKSET,
						"db_cluster_description":  "tf自动化测试",
						"compute_resource":        "16ACU",
						"db_cluster_network_type": "VPC",
						"db_cluster_version":      CHECKSET,
						"payment_type":            "Postpaid",
						"period":                  "Month",
						"used_time":               CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":         "24ACU",
					"db_cluster_description":   "测试修改描述",
					"compute_resource":         "32ACU",
					"security_ips":             "123.23.11.2",
					"db_cluster_ip_array_name": "default",
					"modify_mode":              "Cover",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":         "24ACU",
						"db_cluster_description":   "测试修改描述",
						"compute_resource":         "32ACU",
						"security_ips":             "123.23.11.2",
						"db_cluster_ip_array_name": "default",
						"modify_mode":              "Cover",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"compute_resource":  "48ACU",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"compute_resource":  "48ACU",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_default_resource_group": "false",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap3713 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence3713(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPC ID" {
        dry_run = false
        enable_ipv6 = false
        vpc_name = "adb测试北京vpc"
        cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCH ID" {
        vpc_id = "${alicloud_vpc.VPC ID.id}"
        zone_id = "cn-beijing-k"
        vswitch_name = "北京k测试"
        cidr_block = "172.16.0.0/24"
}


`, name)
}

// Case 测试用例1008_hangzhou_00-160 4752
func TestAccAliCloudAdbDbClusterLakeVersion_basic4752(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap4752)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence4752)
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
					"storage_resource":        "0ACU",
					"zone_id":                 "cn-hangzhou-k",
					"vpc_id":                  "${alicloud_vpc.VPCID.id}",
					"vswitch_id":              "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":  "tf自动化测试-beijing",
					"compute_resource":        "16ACU",
					"db_cluster_network_type": "VPC",
					"db_cluster_version":      "5.0",
					"payment_type":            "Postpaid",
					"period":                  "Month",
					"used_time":               "1",
					"backup_set_id":           "1699271771",
					"source_db_cluster_id":    "am-2ze38892366s6hcn3",
					"restore_type":            "backup",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":        "0ACU",
						"zone_id":                 "cn-hangzhou-k",
						"vpc_id":                  CHECKSET,
						"vswitch_id":              CHECKSET,
						"db_cluster_description":  "tf自动化测试-beijing",
						"compute_resource":        "16ACU",
						"db_cluster_network_type": "VPC",
						"db_cluster_version":      CHECKSET,
						"payment_type":            "Postpaid",
						"period":                  "Month",
						"used_time":               CHECKSET,
						"backup_set_id":           CHECKSET,
						"source_db_cluster_id":    "am-2ze38892366s6hcn3",
						"restore_type":            "backup",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description":   "测试修改描述",
					"compute_resource":         "32ACU",
					"security_ips":             "123.23.11.2",
					"db_cluster_ip_array_name": "default",
					"modify_mode":              "Cover",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description":   "测试修改描述",
						"compute_resource":         "32ACU",
						"security_ips":             "123.23.11.2",
						"db_cluster_ip_array_name": "default",
						"modify_mode":              "Cover",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap4752 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence4752(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "adb杭州测试vpc"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "adb测试杭州"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case 测试用例1008_hangzhou_00-160_2 4789
func TestAccAliCloudAdbDbClusterLakeVersion_basic4789(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap4789)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence4789)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "Postpaid",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "Postpaid",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description":        "测试修改描述",
					"compute_resource":              "32ACU",
					"security_ips":                  "123.23.11.2",
					"db_cluster_ip_array_name":      "default",
					"modify_mode":                   "Cover",
					"db_cluster_ip_array_attribute": "show",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description":        "测试修改描述",
						"compute_resource":              "32ACU",
						"security_ips":                  "123.23.11.2",
						"db_cluster_ip_array_name":      "default",
						"modify_mode":                   "Cover",
						"db_cluster_ip_array_attribute": "show",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "24ACU",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "24ACU",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "false",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap4789 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence4789(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "adb杭州测试vpc"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "adb测试杭州"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case 测试用例1008_hangzhou_00-160_3 4922
func TestAccAliCloudAdbDbClusterLakeVersion_basic4922(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap4922)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence4922)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "Postpaid",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "Postpaid",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description":        "测试修改描述",
					"compute_resource":              "32ACU",
					"security_ips":                  "123.23.11.2",
					"db_cluster_ip_array_name":      "default",
					"modify_mode":                   "Cover",
					"db_cluster_ip_array_attribute": "show",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description":        "测试修改描述",
						"compute_resource":              "32ACU",
						"security_ips":                  "123.23.11.2",
						"db_cluster_ip_array_name":      "default",
						"modify_mode":                   "Cover",
						"db_cluster_ip_array_attribute": "show",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "24ACU",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "24ACU",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "false",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap4922 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence4922(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "adb杭州测试vpc"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "adb测试杭州"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case 测试用例hangzhou_pre_20231129 5257
func TestAccAliCloudAdbDbClusterLakeVersion_basic5257(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap5257)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence5257)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "Postpaid",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "Postpaid",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description":        "测试修改描述",
					"compute_resource":              "32ACU",
					"security_ips":                  "123.23.11.2",
					"maintain_time":                 "05:00Z-06:00Z",
					"db_cluster_ip_array_name":      "default",
					"modify_mode":                   "Cover",
					"db_cluster_ip_array_attribute": "show",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description":        "测试修改描述",
						"compute_resource":              "32ACU",
						"security_ips":                  "123.23.11.2",
						"maintain_time":                 "05:00Z-06:00Z",
						"db_cluster_ip_array_name":      "default",
						"modify_mode":                   "Cover",
						"db_cluster_ip_array_attribute": "show",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "24ACU",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "24ACU",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "false",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap5257 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence5257(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "adb杭州测试vpc"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "adb测试杭州"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case 测试用例hangzhou_pre_20231208_perth相关_包年包月 5442
func TestAccAliCloudAdbDbClusterLakeVersion_basic5442(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap5442)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence5442)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "Subscription",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "Subscription",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description":        "测试修改描述",
					"compute_resource":              "32ACU",
					"security_ips":                  "123.23.11.2",
					"maintain_time":                 "05:00Z-06:00Z",
					"db_cluster_ip_array_name":      "default",
					"modify_mode":                   "Cover",
					"db_cluster_ip_array_attribute": "show",
					"audit_log_status":              "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description":        "测试修改描述",
						"compute_resource":              "32ACU",
						"security_ips":                  "123.23.11.2",
						"maintain_time":                 "05:00Z-06:00Z",
						"db_cluster_ip_array_name":      "default",
						"modify_mode":                   "Cover",
						"db_cluster_ip_array_attribute": "show",
						"audit_log_status":              "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "24ACU",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "24ACU",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "0ACU",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"security_ips":                  "127.0.0.2",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"security_ips":                  "127.0.0.2",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description": "测试修改描述",
					"compute_resource":       "32ACU",
					"security_ips":           "123.23.11.2",
					"maintain_time":          "05:00Z-06:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description": "测试修改描述",
						"compute_resource":       "32ACU",
						"security_ips":           "123.23.11.2",
						"maintain_time":          "05:00Z-06:00Z",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap5442 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence5442(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "tf816"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "tf77"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case 测试用例hangzhou_pre_20231208_perth相关_按量付费仅创建 5444
func TestAccAliCloudAdbDbClusterLakeVersion_basic5444(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap5444)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence5444)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "PayAsYouGo",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "PayAsYouGo",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap5444 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence5444(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "tf453"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "tf901"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case hangzhou_pre__perth_aps_spark_按量付费创建_副本1703606693176 5610
func TestAccAliCloudAdbDbClusterLakeVersion_basic5610(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap5610)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence5610)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "PayAsYouGo",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
					"audit_log_status":              "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "PayAsYouGo",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
						"audit_log_status":              "off",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap5610 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence5610(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "tf650"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "tf526"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Case 勿动勿运行_测试用例hangzhou_pre_20231222_perth相关_包年包月2 5544
func TestAccAliCloudAdbDbClusterLakeVersion_basic5544(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_db_cluster_lake_version.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbDbClusterLakeVersionMap5544)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbDbClusterLakeVersion")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbDbClusterLakeVersionBasicDependence5544)
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
					"storage_resource":              "0ACU",
					"zone_id":                       "cn-hangzhou-k",
					"vpc_id":                        "${alicloud_vpc.VPCID.id}",
					"vswitch_id":                    "${alicloud_vswitch.VSWITCHID.id}",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"db_cluster_network_type":       "VPC",
					"db_cluster_version":            "5.0",
					"payment_type":                  "Subscription",
					"period":                        "Month",
					"used_time":                     "1",
					"backup_set_id":                 "1699271771",
					"source_db_cluster_id":          "am-2ze38892366s6hcn3",
					"restore_type":                  "backup",
					"security_ips":                  "127.0.0.2",
					"restore_to_time":               "2023-09-20T03:13:56Z",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"zone_id":                       "cn-hangzhou-k",
						"vpc_id":                        CHECKSET,
						"vswitch_id":                    CHECKSET,
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"db_cluster_network_type":       "VPC",
						"db_cluster_version":            CHECKSET,
						"payment_type":                  "Subscription",
						"period":                        "Month",
						"used_time":                     CHECKSET,
						"backup_set_id":                 CHECKSET,
						"source_db_cluster_id":          "am-2ze38892366s6hcn3",
						"restore_type":                  "backup",
						"security_ips":                  "127.0.0.2",
						"restore_to_time":               "2023-09-20T03:13:56Z",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description":        "测试修改描述",
					"compute_resource":              "32ACU",
					"security_ips":                  "123.23.11.2",
					"maintain_time":                 "05:00Z-06:00Z",
					"db_cluster_ip_array_name":      "default",
					"modify_mode":                   "Cover",
					"db_cluster_ip_array_attribute": "show",
					"audit_log_status":              "on",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description":        "测试修改描述",
						"compute_resource":              "32ACU",
						"security_ips":                  "123.23.11.2",
						"maintain_time":                 "05:00Z-06:00Z",
						"db_cluster_ip_array_name":      "default",
						"modify_mode":                   "Cover",
						"db_cluster_ip_array_attribute": "show",
						"audit_log_status":              "on",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "24ACU",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"enable_default_resource_group": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "24ACU",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_resource":              "0ACU",
					"db_cluster_description":        "tf自动化测试-beijing",
					"compute_resource":              "16ACU",
					"security_ips":                  "127.0.0.2",
					"resource_group_id":             "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"enable_default_resource_group": "true",
					"maintain_time":                 "04:00Z-05:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_resource":              "0ACU",
						"db_cluster_description":        "tf自动化测试-beijing",
						"compute_resource":              "16ACU",
						"security_ips":                  "127.0.0.2",
						"resource_group_id":             CHECKSET,
						"enable_default_resource_group": "true",
						"maintain_time":                 "04:00Z-05:00Z",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_description": "测试修改描述",
					"compute_resource":       "32ACU",
					"security_ips":           "123.23.11.2",
					"maintain_time":          "05:00Z-06:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_description": "测试修改描述",
						"compute_resource":       "32ACU",
						"security_ips":           "123.23.11.2",
						"maintain_time":          "05:00Z-06:00Z",
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
				ImportStateVerifyIgnore: []string{"auto_renewal_period", "auto_renewal_period_unit", "auto_renewal_status", "backup_set_id", "db_cluster_ip_array_attribute", "db_cluster_ip_array_name", "db_version", "enable_default_resource_group", "engine_type", "modify_mode", "period", "restore_to_time", "restore_type", "source_db_cluster_id", "switch_mode", "used_time"},
			},
		},
	})
}

var AlicloudAdbDbClusterLakeVersionMap5544 = map[string]string{
	"port":              CHECKSET,
	"lock_reason":       CHECKSET,
	"engine":            CHECKSET,
	"status":            CHECKSET,
	"engine_version":    CHECKSET,
	"expired":           CHECKSET,
	"lock_mode":         CHECKSET,
	"expire_time":       CHECKSET,
	"connection_string": CHECKSET,
	"commodity_code":    CHECKSET,
	"create_time":       CHECKSET,
	"region_id":         CHECKSET,
}

func AlicloudAdbDbClusterLakeVersionBasicDependence5544(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "tf339"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "tf108"
  cidr_block   = "172.16.0.0/24"
}


`, name)
}

// Test Adb DbClusterLakeVersion. <<< Resource test cases, automatically generated.
