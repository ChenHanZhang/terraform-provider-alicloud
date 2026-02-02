package alicloud

import (
	"fmt"
	"testing"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Rds DbProxy. >>> Resource test cases, automatically generated.
// Case RDS_DBPROXY_TEST 11914
func TestAccAliCloudRdsDbProxy_basic11914(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_db_proxy.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsDbProxyMap11914)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsDbProxy")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsDbProxyBasicDependence11914)
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
					"db_instance_id":         "rm-2ze87i5thh23z2ym5",
					"db_proxy_instance_type": "exclusive",
					"db_proxy_instance_num":  "4",
					"vpc_id":                 "vpc-2ze7s2hfq6twcdwzs41e4",
					"vswitch_id":             "vsw-2zeq5sk2uyrj7xa8lih3m",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_instance_id":         "rm-2ze87i5thh23z2ym5",
						"db_proxy_instance_type": "exclusive",
						"db_proxy_instance_num":  "4",
						"vpc_id":                 "vpc-2ze7s2hfq6twcdwzs41e4",
						"vswitch_id":             "vsw-2zeq5sk2uyrj7xa8lih3m",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"persistent_connection_status":         "Enabled",
					"causal_consist_read_timeout":          " ",
					"effective_time":                       "MaintainTime",
					"read_only_instance_distribution_type": "Standard",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"persistent_connection_status":         "Enabled",
						"causal_consist_read_timeout":          " ",
						"effective_time":                       "MaintainTime",
						"read_only_instance_distribution_type": "Standard",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_ssl_enabled": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_ssl_enabled": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"causal_consist_read_timeout", "config_db_proxy_features", "db_endpoint_aliases", "db_endpoint_min_slave_count", "db_endpoint_operator", "db_endpoint_read_write_mode", "effective_specific_time", "effective_time", "read_only_instance_distribution_type", "read_only_instance_weight", "vswitch_id", "vpc_id"},
			},
		},
	})
}

var AlicloudRdsDbProxyMap11914 = map[string]string{
	"db_proxy_endpoint_aliases":    CHECKSET,
	"db_proxy_connect_string_port": CHECKSET,
}

func AlicloudRdsDbProxyBasicDependence11914(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-l"
}


`, name)
}

// Test Rds DbProxy. <<< Resource test cases, automatically generated.

func TestAccAliCloudRdsDBProxy_MySQL(t *testing.T) {
	var connection map[string]interface{}
	var primary map[string]interface{}
	var readonly map[string]interface{}

	resourceId := "alicloud_rds_db_proxy.default"
	var DBProxyMap = map[string]string{
		"instance_id":           CHECKSET,
		"instance_network_type": "VPC",
		"db_proxy_instance_num": "2",
		"vswitch_id":            CHECKSET,
		"vpc_id":                CHECKSET,
	}
	ra := resourceAttrInit(resourceId, DBProxyMap)

	rc_connection := resourceCheckInitWithDescribeMethod(resourceId, &connection, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsProxyEndpoint")
	rc_primary := resourceCheckInitWithDescribeMethod("alicloud_db_instance.default", &primary, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDBInstance")
	rc_readonly := resourceCheckInitWithDescribeMethod("alicloud_db_readonly_instance.default", &readonly, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDBReadonlyInstance")
	rand := acctest.RandString(5)
	effectiveSpecificTime := time.Now().UTC().Add(45 * time.Minute).Format("2006-01-02T15:04:05Z")
	switchTime := time.Now().UTC().Add(45 * time.Minute).Format("2006-01-02T15:04:05Z")
	rac := resourceAttrCheckInit(rc_connection, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAccDBProxy%s", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceRdsDBProxyDependence_MySQL)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: resourceId,

		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_id":            "${alicloud_db_instance.default.id}",
					"instance_network_type":  "VPC",
					"db_proxy_instance_num":  "2",
					"db_proxy_instance_type": "common",
					"vswitch_id":             "${alicloud_db_instance.default.vswitch_id}",
					"vpc_id":                 "${alicloud_db_instance.default.vpc_id}",
					"depends_on":             []string{"alicloud_db_readonly_instance.default"},
					"resource_group_id":      "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					rc_primary.checkResourceExists(),
					rc_readonly.checkResourceExists(),
					testAccCheck(map[string]string{
						"instance_id":            CHECKSET,
						"db_proxy_instance_type": CHECKSET,
						"instance_network_type":  "VPC",
						"db_proxy_instance_num":  "2",
						"vswitch_id":             CHECKSET,
						"vpc_id":                 CHECKSET,
						"resource_group_id":      CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"upgrade_time": "Immediate",
					"switch_time":  switchTime,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"upgrade_time": "Immediate",
						"switch_time":  CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_instance_num":   "2",
					"db_proxy_instance_type":  "exclusive",
					"effective_time":          "SpecificTime",
					"effective_specific_time": effectiveSpecificTime,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_instance_num":   "2",
						"db_proxy_instance_type":  "exclusive",
						"effective_time":          "SpecificTime",
						"effective_specific_time": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_endpoint_read_write_mode": "ReadWrite",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_endpoint_read_write_mode": "ReadWrite",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"read_only_instance_distribution_type": "Standard",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"read_only_instance_distribution_type": "Standard",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"read_only_instance_max_delay_time": "90",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"read_only_instance_max_delay_time": "90",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_features": "TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_features": "TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"read_only_instance_distribution_type": "Custom",
					"db_proxy_features":                    "TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:1",
					"read_only_instance_weight": []map[string]interface{}{
						{
							"instance_id": "${alicloud_db_instance.default.id}",
							"weight":      "100",
						},
						{
							"instance_id": "${alicloud_db_readonly_instance.default.id}",
							"weight":      "500",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"read_only_instance_distribution_type": "Custom",
						"db_proxy_features":                    "TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:1",
						"read_only_instance_weight.#":          "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_instance_num":  "3",
					"effective_time":         "Immediate",
					"db_proxy_instance_type": "exclusive",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_instance_num":  "3",
						"effective_time":         "Immediate",
						"db_proxy_instance_type": "exclusive",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_connect_string_port": "3306",
					"db_proxy_connection_prefix":   fmt.Sprintf("tf-testacc%s", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_connect_string_port": "3306",
						"db_proxy_connection_prefix":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_proxy_ssl_enabled": "Open",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_proxy_ssl_enabled": "Open",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       false,
				ImportStateVerifyIgnore: []string{"switch_time", "effective_specific_time"},
			},
		},
	})
}

func resourceRdsDBProxyDependence_MySQL(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}
data "alicloud_db_zones" "default"{
	engine = "MySQL"
	engine_version = "8.0"
	instance_charge_type = "PostPaid"
	category = "HighAvailability"
 	db_instance_storage_type = "cloud_essd"
}

data "alicloud_db_instance_classes" "default" {
    zone_id = data.alicloud_db_zones.default.zones.0.id
	engine = "MySQL"
	engine_version = "8.0"
    category = "HighAvailability"
 	db_instance_storage_type = "cloud_essd"
	instance_charge_type = "PostPaid"
}

data "alicloud_vpcs" "default" {
    name_regex = "^default-NODELETING$"
}
data "alicloud_vswitches" "default" {
  vpc_id = data.alicloud_vpcs.default.ids.0
  zone_id = data.alicloud_db_zones.default.zones.0.id
}

resource "alicloud_vswitch" "this" {
 count = length(data.alicloud_vswitches.default.ids) > 0 ? 0 : 1
 vswitch_name = var.name
 vpc_id = data.alicloud_vpcs.default.ids.0
 zone_id = data.alicloud_db_zones.default.ids.0
 cidr_block = cidrsubnet(data.alicloud_vpcs.default.vpcs.0.cidr_block, 8, 4)
}
locals {
  vswitch_id = length(data.alicloud_vswitches.default.ids) > 0 ? data.alicloud_vswitches.default.ids.0 : concat(alicloud_vswitch.this.*.id, [""])[0]
  zone_id = data.alicloud_db_zones.default.ids[length(data.alicloud_db_zones.default.ids)-1]
}

data "alicloud_resource_manager_resource_groups" "default" {
	status = "OK"
}

resource "alicloud_db_instance" "default" {
    engine = "MySQL"
	engine_version = "8.0"
 	db_instance_storage_type = "cloud_essd"
	instance_type = "mysql.x2.large.2c"
	instance_storage = data.alicloud_db_instance_classes.default.instance_classes.0.storage_range.min
	vswitch_id = local.vswitch_id
	instance_name = var.name
}

resource "alicloud_db_readonly_instance" "default" {
	master_db_instance_id = alicloud_db_instance.default.id
	zone_id = alicloud_db_instance.default.zone_id
	engine_version = alicloud_db_instance.default.engine_version
	instance_type = "mysqlro.n4.medium.1c"
	instance_storage = alicloud_db_instance.default.instance_storage
	instance_name = "${var.name}_ro"
	vswitch_id = alicloud_db_instance.default.vswitch_id
}

`, name)
}
