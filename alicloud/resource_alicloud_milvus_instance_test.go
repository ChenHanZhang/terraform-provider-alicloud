// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Milvus Instance. >>> Resource test cases, automatically generated.
// Case instance_terraform_subscription_month_pre 13039
func TestAccAliCloudMilvusInstance_basic13039(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap13039)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence13039)
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
					"ai_function":       "false",
					"zone_id":           "cn-hangzhou-j",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"vswitch_ids": []map[string]interface{}{
						{
							"zone_id": "cn-hangzhou-j",
							"vsw_id":  "vsw-bp1pommb2vygb0kzvf8i6",
						},
						{
							"zone_id": "cn-hangzhou-k",
							"vsw_id":  "vsw-bp1tomony773mb6nlabw9",
						},
					},
					"encrypted":             "false",
					"auto_renew":            "false",
					"payment_duration_unit": "month",
					"auto_pay":              "true",
					"load_replicas":         "2",
					"payment_duration":      "1",
					"db_admin_password":     "@1234Test",
					"instance_name":         name,
					"components": []map[string]interface{}{
						{
							"cu_type": "general",
							"type":    "streaming",
							"cu_num":  "4",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "data",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"cu_type":        "general",
							"type":           "query",
							"cu_num":         "16",
							"disk_size_type": "Normal",
							"replica":        "2",
						},
						{
							"cu_type": "general",
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
						},
					},
					"db_version":          "2.6",
					"vpc_id":              "vpc-bp168d0ay5yft9aira762",
					"is_multi_az_storage": "true",
					"payment_type":        "Subscription",
					"ha":                  "true",
					"multi_zone_mode":     "single",
					"auto_backup":         "true",
					"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_function":           "false",
						"zone_id":               "cn-hangzhou-j",
						"resource_group_id":     CHECKSET,
						"vswitch_ids.#":         "2",
						"encrypted":             "false",
						"auto_renew":            "false",
						"payment_duration_unit": "month",
						"auto_pay":              "true",
						"load_replicas":         "2",
						"payment_duration":      "1",
						"db_admin_password":     "@1234Test",
						"instance_name":         name,
						"components.#":          "5",
						"db_version":            CHECKSET,
						"vpc_id":                "vpc-bp168d0ay5yft9aira762",
						"is_multi_az_storage":   "true",
						"payment_type":          "Subscription",
						"ha":                    "true",
						"multi_zone_mode":       "single",
						"auto_backup":           "true",
						"promotion_no":          "youhuiquan_promotion_option_id_for_blank",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap13039 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence13039(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case instance_terraform_restore_pre 13040
func TestAccAliCloudMilvusInstance_basic13040(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap13040)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence13040)
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
					"ai_function":       "false",
					"zone_id":           "cn-hangzhou-j",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"vswitch_ids": []map[string]interface{}{
						{
							"zone_id": "cn-hangzhou-j",
							"vsw_id":  "vsw-bp1pommb2vygb0kzvf8i6",
						},
						{
							"zone_id": "cn-hangzhou-k",
							"vsw_id":  "vsw-bp1tomony773mb6nlabw9",
						},
					},
					"encrypted":         "false",
					"auto_pay":          "false",
					"load_replicas":     "2",
					"db_admin_password": "@1234Test",
					"instance_name":     name,
					"components": []map[string]interface{}{
						{
							"cu_type": "general",
							"type":    "streaming",
							"cu_num":  "4",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "data",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"cu_type":        "general",
							"type":           "query",
							"cu_num":         "16",
							"disk_size_type": "Normal",
							"replica":        "2",
						},
						{
							"cu_type": "general",
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
						},
					},
					"db_version":          "2.6",
					"vpc_id":              "vpc-bp168d0ay5yft9aira762",
					"is_multi_az_storage": "true",
					"payment_type":        "PayAsYouGo",
					"ha":                  "true",
					"multi_zone_mode":     "single",
					"auto_backup":         "true",
					"backup_restore_info": []map[string]interface{}{
						{
							"backup_id":         "bt-69b8fdaff88db73f",
							"source_cluster_id": "c-2d95c862a7142aec",
							"backup_name":       "auto_backup_2026_08_25_03_00_00_902610000",
						},
					},
					"promotion_no": "youhuiquan_promotion_option_id_for_blank",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_function":         "false",
						"zone_id":             "cn-hangzhou-j",
						"resource_group_id":   CHECKSET,
						"vswitch_ids.#":       "2",
						"encrypted":           "false",
						"auto_pay":            "false",
						"load_replicas":       "2",
						"db_admin_password":   "@1234Test",
						"instance_name":       name,
						"components.#":        "5",
						"db_version":          CHECKSET,
						"vpc_id":              "vpc-bp168d0ay5yft9aira762",
						"is_multi_az_storage": "true",
						"payment_type":        "PayAsYouGo",
						"ha":                  "true",
						"multi_zone_mode":     "single",
						"auto_backup":         "true",
						"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap13040 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence13040(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case instance_terraform_subscription_year_pre 13041
func TestAccAliCloudMilvusInstance_basic13041(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap13041)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence13041)
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
					"ai_function":       "false",
					"zone_id":           "cn-hangzhou-j",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"vswitch_ids": []map[string]interface{}{
						{
							"zone_id": "cn-hangzhou-j",
							"vsw_id":  "vsw-bp1pommb2vygb0kzvf8i6",
						},
						{
							"zone_id": "cn-hangzhou-k",
							"vsw_id":  "vsw-bp1tomony773mb6nlabw9",
						},
					},
					"encrypted":             "false",
					"auto_renew":            "true",
					"payment_duration_unit": "year",
					"auto_pay":              "true",
					"load_replicas":         "2",
					"payment_duration":      "1",
					"db_admin_password":     "@1234Test",
					"instance_name":         name,
					"components": []map[string]interface{}{
						{
							"cu_type": "general",
							"type":    "streaming",
							"cu_num":  "4",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "data",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"cu_type":        "general",
							"type":           "query",
							"cu_num":         "16",
							"disk_size_type": "Large",
							"replica":        "2",
						},
						{
							"cu_type": "general",
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
						},
					},
					"db_version":          "2.6",
					"vpc_id":              "vpc-bp168d0ay5yft9aira762",
					"is_multi_az_storage": "true",
					"payment_type":        "Subscription",
					"ha":                  "true",
					"multi_zone_mode":     "single",
					"auto_backup":         "true",
					"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_function":           "false",
						"zone_id":               "cn-hangzhou-j",
						"resource_group_id":     CHECKSET,
						"vswitch_ids.#":         "2",
						"encrypted":             "false",
						"auto_renew":            "true",
						"payment_duration_unit": "year",
						"auto_pay":              "true",
						"load_replicas":         "2",
						"payment_duration":      "1",
						"db_admin_password":     "@1234Test",
						"instance_name":         name,
						"components.#":          "5",
						"db_version":            CHECKSET,
						"vpc_id":                "vpc-bp168d0ay5yft9aira762",
						"is_multi_az_storage":   "true",
						"payment_type":          "Subscription",
						"ha":                    "true",
						"multi_zone_mode":       "single",
						"auto_backup":           "true",
						"promotion_no":          "youhuiquan_promotion_option_id_for_blank",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap13041 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence13041(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case instance_terraform_parity_pre 13042
func TestAccAliCloudMilvusInstance_basic13042(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap13042)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence13042)
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
					"ai_function":       "false",
					"zone_id":           "cn-hangzhou-j",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"configuration":     "{rootCoord: {maxDatabaseNum: 64}}",
					"vswitch_ids": []map[string]interface{}{
						{
							"zone_id": "cn-hangzhou-j",
							"vsw_id":  "vsw-bp1pommb2vygb0kzvf8i6",
						},
						{
							"zone_id": "cn-hangzhou-k",
							"vsw_id":  "vsw-bp1tomony773mb6nlabw9",
						},
					},
					"encrypted":         "false",
					"auto_pay":          "false",
					"load_replicas":     "2",
					"db_admin_password": "@1234Test",
					"instance_name":     name,
					"components": []map[string]interface{}{
						{
							"cu_type": "general",
							"type":    "query",
							"cu_num":  "16",
							"data_disk": []map[string]interface{}{
								{
									"storage_class":     "alicloud-disk-essd-pl1",
									"size":              "100",
									"performance_level": "PL1",
									"enabled":           "true",
								},
							},
							"disk_size_type": "Normal",
							"replica":        "2",
						},
						{
							"cu_type": "general",
							"type":    "data",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"cu_type": "general",
							"type":    "streaming",
							"cu_num":  "4",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
						},
					},
					"db_version":          "2.6",
					"vpc_id":              "vpc-bp168d0ay5yft9aira762",
					"is_multi_az_storage": "true",
					"payment_type":        "PayAsYouGo",
					"ha":                  "true",
					"multi_zone_mode":     "single",
					"auto_backup":         "true",
					"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_function":         "false",
						"zone_id":             "cn-hangzhou-j",
						"resource_group_id":   CHECKSET,
						"configuration":       "{rootCoord: {maxDatabaseNum: 64}}",
						"vswitch_ids.#":       "2",
						"encrypted":           "false",
						"auto_pay":            "false",
						"load_replicas":       "2",
						"db_admin_password":   "@1234Test",
						"instance_name":       name,
						"components.#":        "5",
						"db_version":          CHECKSET,
						"vpc_id":              "vpc-bp168d0ay5yft9aira762",
						"is_multi_az_storage": "true",
						"payment_type":        "PayAsYouGo",
						"ha":                  "true",
						"multi_zone_mode":     "single",
						"auto_backup":         "true",
						"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"configuration":     "{rootCoord: {maxDatabaseNum: 96}}",
					"instance_name":     name + "_update",
					"components": []map[string]interface{}{
						{
							"cu_type": "general",
							"type":    "query",
							"cu_num":  "16",
							"data_disk": []map[string]interface{}{
								{
									"storage_class":     "alicloud-disk-essd-pl1",
									"size":              "120",
									"performance_level": "PL1",
									"enabled":           "true",
								},
							},
							"replica": "2",
						},
					},
					"auto_backup": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"configuration":     "{rootCoord: {maxDatabaseNum: 96}}",
						"instance_name":     name + "_update",
						"components.#":      "1",
						"auto_backup":       "false",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap13042 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence13042(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case instance_terraform_encrypted_pre 13043
func TestAccAliCloudMilvusInstance_basic13043(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap13043)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence13043)
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
					"ai_function":       "false",
					"zone_id":           "cn-hangzhou-j",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"kms_key_id":        "key-hzz68ca89cbjfuikcksoq",
					"vswitch_ids": []map[string]interface{}{
						{
							"zone_id": "cn-hangzhou-j",
							"vsw_id":  "vsw-bp1pommb2vygb0kzvf8i6",
						},
						{
							"zone_id": "cn-hangzhou-k",
							"vsw_id":  "vsw-bp1tomony773mb6nlabw9",
						},
					},
					"encrypted":         "true",
					"auto_pay":          "false",
					"load_replicas":     "2",
					"db_admin_password": "@1234Test",
					"instance_name":     name,
					"components": []map[string]interface{}{
						{
							"cu_type": "general",
							"type":    "streaming",
							"cu_num":  "4",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "data",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"cu_type":        "general",
							"type":           "query",
							"cu_num":         "16",
							"disk_size_type": "Large",
							"replica":        "2",
						},
						{
							"cu_type": "general",
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
						},
						{
							"cu_type": "general",
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
						},
					},
					"db_version":          "2.6",
					"vpc_id":              "vpc-bp168d0ay5yft9aira762",
					"is_multi_az_storage": "true",
					"payment_type":        "PayAsYouGo",
					"ha":                  "true",
					"multi_zone_mode":     "single",
					"auto_backup":         "true",
					"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ai_function":         "false",
						"zone_id":             "cn-hangzhou-j",
						"resource_group_id":   CHECKSET,
						"kms_key_id":          "key-hzz68ca89cbjfuikcksoq",
						"vswitch_ids.#":       "2",
						"encrypted":           "true",
						"auto_pay":            "false",
						"load_replicas":       "2",
						"db_admin_password":   "@1234Test",
						"instance_name":       name,
						"components.#":        "5",
						"db_version":          CHECKSET,
						"vpc_id":              "vpc-bp168d0ay5yft9aira762",
						"is_multi_az_storage": "true",
						"payment_type":        "PayAsYouGo",
						"ha":                  "true",
						"multi_zone_mode":     "single",
						"auto_backup":         "true",
						"promotion_no":        "youhuiquan_promotion_option_id_for_blank",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap13043 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence13043(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case instance-按量更新_张家口_资源组 12774
func TestAccAliCloudMilvusInstance_basic12774(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap12774)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence12774)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-zhangjiakou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":           "data",
							"cu_num":         "2",
							"replica":        "1",
							"cu_type":        "general",
							"disk_size_type": "Normal",
						},
						{
							"type":           "index",
							"cu_num":         "4",
							"replica":        "2",
							"cu_type":        "general",
							"disk_size_type": "Normal",
						},
						{
							"type":           "query",
							"cu_num":         "8",
							"replica":        "2",
							"cu_type":        "general",
							"disk_size_type": "Large",
						},
						{
							"type":           "proxy",
							"cu_num":         "2",
							"replica":        "2",
							"cu_type":        "general",
							"disk_size_type": "Normal",
						},
						{
							"type":           "mix_coordinator",
							"cu_num":         "4",
							"replica":        "2",
							"cu_type":        "general",
							"disk_size_type": "Normal",
						},
					},
					"instance_name":     name,
					"db_version":        "2.4",
					"vpc_id":            "${alicloud_vpc.defaultILXuit.id}",
					"ha":                "false",
					"payment_type":      "PayAsYouGo",
					"multi_zone_mode":   "Single",
					"kms_key_id":        "key-test-milvus",
					"encrypted":         "false",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"auto_backup":       "false",
					"configuration":     "rootCoord:\\n    maxDatabaseNum: 64 # Maximum number of database",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":           CHECKSET,
						"vswitch_ids.#":     "1",
						"db_admin_password": "Test123321@",
						"components.#":      "5",
						"instance_name":     name,
						"db_version":        CHECKSET,
						"vpc_id":            CHECKSET,
						"ha":                "false",
						"payment_type":      "PayAsYouGo",
						"multi_zone_mode":   "Single",
						"kms_key_id":        "key-test-milvus",
						"encrypted":         "false",
						"resource_group_id": CHECKSET,
						"auto_backup":       "false",
						"configuration":     "rootCoord:\n    maxDatabaseNum: 64 # Maximum number of database",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name":     name + "_update",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name":     name + "_update",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap12774 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence12774(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-zhangjiakou"
}

variable "zone_id" {
  default = "cn-zhangjiakou-b"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance-按量更新_tag 11665
func TestAccAliCloudMilvusInstance_basic11665(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11665)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11665)
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
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "data",
							"cu_num":  "2",
							"replica": "1",
							"cu_type": "general",
						},
						{
							"type":    "index",
							"cu_num":  "4",
							"replica": "2",
							"cu_type": "general",
						},
						{
							"type":    "query",
							"cu_num":  "4",
							"replica": "2",
							"cu_type": "general",
						},
						{
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
							"cu_type": "general",
						},
						{
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
							"cu_type": "general",
						},
					},
					"instance_name":   name,
					"db_version":      "2.4",
					"vpc_id":          "${alicloud_vpc.defaultILXuit.id}",
					"ha":              "false",
					"payment_type":    "PayAsYouGo",
					"multi_zone_mode": "Single",
					"kms_key_id":      "k-test",
					"encrypted":       "false",
					"auto_pay":        "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":           CHECKSET,
						"vswitch_ids.#":     "1",
						"db_admin_password": "Test123321@",
						"components.#":      "5",
						"instance_name":     name,
						"db_version":        CHECKSET,
						"vpc_id":            CHECKSET,
						"ha":                "false",
						"payment_type":      "PayAsYouGo",
						"multi_zone_mode":   "Single",
						"kms_key_id":        "k-test",
						"encrypted":         "false",
						"auto_pay":          "false",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11665 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11665(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance包年包月-月_张家口 11772
func TestAccAliCloudMilvusInstance_basic11772(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11772)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11772)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-zhangjiakou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "standalone_pro",
							"cu_num":  "8",
							"replica": "1",
							"cu_type": "general",
						},
					},
					"instance_name":         name,
					"db_version":            "2.4",
					"vpc_id":                "${alicloud_vpc.defaultILXuit.id}",
					"ha":                    "false",
					"payment_type":          "Subscription",
					"multi_zone_mode":       "Single",
					"payment_duration_unit": "month",
					"payment_duration":      "1",
					"auto_pay":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":               CHECKSET,
						"vswitch_ids.#":         "1",
						"db_admin_password":     "Test123321@",
						"components.#":          "1",
						"instance_name":         name,
						"db_version":            CHECKSET,
						"vpc_id":                CHECKSET,
						"ha":                    "false",
						"payment_type":          "Subscription",
						"multi_zone_mode":       "Single",
						"payment_duration_unit": "month",
						"payment_duration":      "1",
						"auto_pay":              "true",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11772 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11772(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-zhangjiakou"
}

variable "zone_id" {
  default = "cn-zhangjiakou-b"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance_包年包月-年_张家口 11774
func TestAccAliCloudMilvusInstance_basic11774(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11774)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11774)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-zhangjiakou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "standalone_pro",
							"cu_num":  "8",
							"replica": "1",
							"cu_type": "general",
						},
					},
					"instance_name":         name,
					"db_version":            "2.4",
					"vpc_id":                "${alicloud_vpc.defaultILXuit.id}",
					"ha":                    "false",
					"payment_type":          "Subscription",
					"multi_zone_mode":       "Single",
					"payment_duration_unit": "year",
					"payment_duration":      "1",
					"auto_pay":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":               CHECKSET,
						"vswitch_ids.#":         "1",
						"db_admin_password":     "Test123321@",
						"components.#":          "1",
						"instance_name":         name,
						"db_version":            CHECKSET,
						"vpc_id":                CHECKSET,
						"ha":                    "false",
						"payment_type":          "Subscription",
						"multi_zone_mode":       "Single",
						"payment_duration_unit": "year",
						"payment_duration":      "1",
						"auto_pay":              "true",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11774 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11774(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-zhangjiakou"
}

variable "zone_id" {
  default = "cn-zhangjiakou-b"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance-按量更新_tag_张家口 11771
func TestAccAliCloudMilvusInstance_basic11771(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11771)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11771)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-zhangjiakou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "data",
							"cu_num":  "2",
							"replica": "1",
							"cu_type": "general",
						},
						{
							"type":    "index",
							"cu_num":  "4",
							"replica": "2",
							"cu_type": "general",
						},
						{
							"type":    "query",
							"cu_num":  "4",
							"replica": "2",
							"cu_type": "general",
						},
						{
							"type":    "proxy",
							"cu_num":  "2",
							"replica": "2",
							"cu_type": "general",
						},
						{
							"type":    "mix_coordinator",
							"cu_num":  "4",
							"replica": "2",
							"cu_type": "general",
						},
					},
					"instance_name":   name,
					"db_version":      "2.4",
					"vpc_id":          "${alicloud_vpc.defaultILXuit.id}",
					"ha":              "false",
					"payment_type":    "PayAsYouGo",
					"multi_zone_mode": "Single",
					"kms_key_id":      "k-test",
					"encrypted":       "false",
					"auto_pay":        "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":           CHECKSET,
						"vswitch_ids.#":     "1",
						"db_admin_password": "Test123321@",
						"components.#":      "5",
						"instance_name":     name,
						"db_version":        CHECKSET,
						"vpc_id":            CHECKSET,
						"ha":                "false",
						"payment_type":      "PayAsYouGo",
						"multi_zone_mode":   "Single",
						"kms_key_id":        "k-test",
						"encrypted":         "false",
						"auto_pay":          "false",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11771 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11771(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-zhangjiakou"
}

variable "zone_id" {
  default = "cn-zhangjiakou-b"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance_包年包月-年 11635
func TestAccAliCloudMilvusInstance_basic11635(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11635)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11635)
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
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "standalone_pro",
							"cu_num":  "8",
							"replica": "1",
							"cu_type": "general",
						},
					},
					"instance_name":         name,
					"db_version":            "2.4",
					"vpc_id":                "${alicloud_vpc.defaultILXuit.id}",
					"ha":                    "false",
					"payment_type":          "Subscription",
					"multi_zone_mode":       "Single",
					"payment_duration_unit": "year",
					"payment_duration":      "1",
					"auto_pay":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":               CHECKSET,
						"vswitch_ids.#":         "1",
						"db_admin_password":     "Test123321@",
						"components.#":          "1",
						"instance_name":         name,
						"db_version":            CHECKSET,
						"vpc_id":                CHECKSET,
						"ha":                    "false",
						"payment_type":          "Subscription",
						"multi_zone_mode":       "Single",
						"payment_duration_unit": "year",
						"payment_duration":      "1",
						"auto_pay":              "true",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11635 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11635(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance包年包月-月-标准版2.5升配预发 11718
func TestAccAliCloudMilvusInstance_basic11718(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11718)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11718)
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
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "data",
							"cu_num":  "2",
							"replica": "1",
							"cu_type": "general",
						},
						{
							"type":    "index",
							"cu_type": "general",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"type":    "query",
							"cu_type": "general",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"type":    "mix_coordinator",
							"cu_type": "general",
							"cu_num":  "4",
							"replica": "1",
						},
						{
							"type":    "proxy",
							"cu_type": "general",
							"cu_num":  "2",
							"replica": "1",
						},
					},
					"instance_name":         name,
					"db_version":            "2.5",
					"vpc_id":                "${alicloud_vpc.defaultILXuit.id}",
					"ha":                    "false",
					"payment_type":          "Subscription",
					"multi_zone_mode":       "Single",
					"payment_duration_unit": "month",
					"payment_duration":      "1",
					"auto_pay":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":               CHECKSET,
						"vswitch_ids.#":         "1",
						"db_admin_password":     "Test123321@",
						"components.#":          "5",
						"instance_name":         name,
						"db_version":            CHECKSET,
						"vpc_id":                CHECKSET,
						"ha":                    "false",
						"payment_type":          "Subscription",
						"multi_zone_mode":       "Single",
						"payment_duration_unit": "month",
						"payment_duration":      "1",
						"auto_pay":              "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"components": []map[string]interface{}{
						{
							"type":    "proxy",
							"cu_type": "general",
							"cu_num":  "4",
							"replica": "2",
						},
						{
							"type":    "data",
							"cu_type": "general",
							"cu_num":  "4",
							"replica": "2",
						},
					},
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"components.#":  "2",
						"instance_name": name + "_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"components": []map[string]interface{}{
						{
							"type":    "data",
							"cu_type": "general",
							"cu_num":  "2",
							"replica": "1",
						},
					},
					"instance_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"components.#":  "1",
						"instance_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11718 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11718(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Case instance包年包月-月 11679
func TestAccAliCloudMilvusInstance_basic11679(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_milvus_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudMilvusInstanceMap11679)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MilvusServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMilvusInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmilvus%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMilvusInstanceBasicDependence11679)
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
					"zone_id": "${var.zone_id}",
					"vswitch_ids": []map[string]interface{}{
						{
							"vsw_id":  "${alicloud_vswitch.defaultN80M7S.id}",
							"zone_id": "${alicloud_vswitch.defaultN80M7S.zone_id}",
						},
					},
					"db_admin_password": "Test123321@",
					"components": []map[string]interface{}{
						{
							"type":    "standalone_pro",
							"cu_num":  "8",
							"replica": "1",
							"cu_type": "general",
						},
					},
					"instance_name":         name,
					"db_version":            "2.4",
					"vpc_id":                "${alicloud_vpc.defaultILXuit.id}",
					"ha":                    "false",
					"payment_type":          "Subscription",
					"multi_zone_mode":       "Single",
					"payment_duration_unit": "month",
					"payment_duration":      "1",
					"auto_pay":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"zone_id":               CHECKSET,
						"vswitch_ids.#":         "1",
						"db_admin_password":     "Test123321@",
						"components.#":          "1",
						"instance_name":         name,
						"db_version":            CHECKSET,
						"vpc_id":                CHECKSET,
						"ha":                    "false",
						"payment_type":          "Subscription",
						"multi_zone_mode":       "Single",
						"payment_duration_unit": "month",
						"payment_duration":      "1",
						"auto_pay":              "true",
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
				ImportStateVerifyIgnore: []string{"ai_function", "auto_pay", "auto_renew", "backup_restore_info", "db_admin_password", "is_multi_az_storage", "load_replicas", "payment_duration", "payment_duration_unit", "promotion_no"},
			},
		},
	})
}

var AlicloudMilvusInstanceMap11679 = map[string]string{
	"status":               CHECKSET,
	"create_time":          CHECKSET,
	"order_id":             CHECKSET,
	"security_group_ids.#": CHECKSET,
	"expire_time":          CHECKSET,
	"running_time":         CHECKSET,
}

func AlicloudMilvusInstanceBasicDependence11679(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "zone_id" {
  default = "cn-hangzhou-j"
}

resource "alicloud_vpc" "defaultILXuit" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultN80M7S" {
  vpc_id       = alicloud_vpc.defaultILXuit.id
  zone_id      = var.zone_id
  cidr_block   = "172.16.1.0/24"
  vswitch_name = "milvus-test"
}


`, name)
}

// Test Milvus Instance. <<< Resource test cases, automatically generated.
