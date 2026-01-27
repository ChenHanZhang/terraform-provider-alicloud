package alicloud

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
}

var v map[string]interface{}

func TestAccAliCloudPolarDBZonalCluster_complete(t *testing.T) {
	name := "tf-testAccPolarDBClusterDBNodeConfigMYSQL"
	resourceId := "alicloud_polardb_zonal_db_cluster.default"
	var basicMap = map[string]string{
		"description":   CHECKSET,
		"db_node_class": CHECKSET,
		"vswitch_id":    CHECKSET,
		"db_type":       CHECKSET,
		"db_version":    CHECKSET,
	}
	ra := resourceAttrInit(resourceId, basicMap)
	serviceFunc := func() interface{} {
		return &PolarDbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, serviceFunc, "DescribePolarDbZonalCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourcePolarDBClusterDBNodeConfigConfigDependence)

	preparedNodes, preparedNodesEscape := prepareNodeConfigs()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		// module name
		IDRefreshName: resourceId,

		Providers:    testAccProviders,
		CheckDestroy: rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_class":        "polar.mysql.x4.medium.c",
					"ens_region_id":        "tr-Istanbul-1",
					"db_type":              "MySQL",
					"storage_type":         "ESSDPL0",
					"db_version":           "8.0",
					"db_minor_version":     "8.0.2",
					"creation_category":    "SENormal",
					"used_time":            "1",
					"pay_type":             "PrePaid",
					"cluster_version":      "8.0.2.2.28",
					"vpc_id":               "${alicloud_ens_network.network.id}",
					"vswitch_id":           "${alicloud_ens_vswitch.switch.id}",
					"target_minor_version": "innovate_x86#20250311",
					"db_cluster_nodes_configs": map[string]string{
						"db_node_1": preparedNodesEscape["db_node_1"],
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"id":                                 CHECKSET,
						"db_cluster_nodes_configs.%":         "1",
						"db_cluster_nodes_attributes.%":      "1",
						"db_cluster_nodes_configs.db_node_1": preparedNodes["db_node_1"],
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"db_node_class", "target_minor_version", "used_time"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_class": "polar.mysql.x8.medium.c",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_class": "polar.mysql.x8.medium.c",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"target_minor_version": "innovate_x86#20250312",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"target_minor_version": "innovate_x86#20250312",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "tf-testAccPolarDBClusterDescriptionEdit",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "tf-testAccPolarDBClusterDescriptionEdit",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_space": "30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_space": "30",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status":    "AutoRenewal",
					"auto_renew_period": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status":    "AutoRenewal",
						"auto_renew_period": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_nodes_configs": map[string]string{
						"db_node_1": preparedNodesEscape["db_node_1"],
						"node1":     preparedNodesEscape["node1"],
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_nodes_configs.%":     "2",
						"db_cluster_nodes_attributes.%":  "2",
						"db_cluster_nodes_configs.node1": preparedNodes["node1"],
					}),
				),
				PreventDiskCleanup: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_nodes_configs": map[string]string{
						"db_node_1": preparedNodesEscape["nodeWriterDemote"],
						"node1":     preparedNodesEscape["node1Promote"],
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_nodes_configs.%":         "2",
						"db_cluster_nodes_attributes.%":      "2",
						"db_cluster_nodes_configs.db_node_1": preparedNodes["nodeWriterDemote"],
						"db_cluster_nodes_configs.node1":     preparedNodes["node1Promote"],
					}),
				),
				PreventDiskCleanup: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_nodes_configs": map[string]string{
						"db_node_1": preparedNodesEscape["nodeWriterDemote"],
						"node1":     preparedNodesEscape["node1Upgrade"],
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_nodes_configs.%":     "2",
						"db_cluster_nodes_attributes.%":  "2",
						"db_cluster_nodes_configs.node1": preparedNodes["node1Upgrade"],
					}),
				),
				PreventDiskCleanup: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_cluster_nodes_configs": map[string]string{
						"node1": preparedNodesEscape["node1Upgrade"],
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_nodes_configs.%":         "1",
						"db_cluster_nodes_attributes.%":      "1",
						"db_cluster_nodes_configs.db_node_1": REMOVEKEY,
						"db_cluster_nodes_configs.node1":     preparedNodes["node1Upgrade"],
					}),
				),
				PreventDiskCleanup: true,
			},
		},
	})
}

func resourcePolarDBClusterDBNodeConfigConfigDependence(name string) string {
	return fmt.Sprintf(`
	variable "name" {
		default = "%s"
	}

	resource "alicloud_ens_network" "network" {
	  network_name = var.name
	
	  description   = "LoadBalancerNetworkDescription_autotest"
	  cidr_block    = "192.168.2.0/24"
	  ens_region_id = "tr-Istanbul-1"
	}
	
	resource "alicloud_ens_vswitch" "switch" {
	  description  = "LoadBalancerVSwitchDescription_autotest"
	  cidr_block   = "192.168.2.0/24"
	  vswitch_name = var.name
	
	  ens_region_id = "tr-Istanbul-1"
	  network_id    = alicloud_ens_network.network.id
	}
	
	locals {
		vpc_id = alicloud_ens_network.network.id
		vswitch_id = alicloud_ens_vswitch.switch.id
	}

`, name)
}

func prepareNodeConfigs() (map[string]string, map[string]string) {
	nodes := map[string]map[string]string{
		"db_node_1":        {"db_node_class": "polar.mysql.x4.medium.c", "db_node_role": "Writer"},
		"node1":            {"db_node_class": "polar.mysql.x4.medium.c", "db_node_role": "Reader"},
		"node1Upgrade":     {"db_node_class": "polar.mysql.x8.medium.c", "db_node_role": "Writer"},
		"node1Promote":     {"db_node_class": "polar.mysql.x4.medium.c", "db_node_role": "Writer"},
		"nodeWriterDemote": {"db_node_class": "polar.mysql.x4.medium.c", "db_node_role": "Reader"},
		"node2":            {"db_node_class": "polar.mysql.x4.medium.c", "db_node_role": "Reader"},
	}

	preparedNodes := make(map[string]string)
	preparedNodesEscape := make(map[string]string)

	for key, value := range nodes {
		nodeJSON, _ := json.Marshal(value)
		preparedNodes[key] = string(nodeJSON)
		preparedNodesEscape[key] = strings.ReplaceAll(string(nodeJSON), `"`, `\"`)
	}
	return preparedNodes, preparedNodesEscape
}

// Test Polardb ZonalDbCluster. >>> Resource test cases, automatically generated.
// Case 基础用例 11408
func TestAccAliCloudPolardbZonalDbCluster_basic11408(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_polardb_zonal_db_cluster.default"
	ra := resourceAttrInit(resourceId, AlicloudPolardbZonalDbClusterMap11408)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &PolardbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribePolardbZonalDbCluster")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccpolardb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudPolardbZonalDbClusterBasicDependence11408)
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
					"db_node_class":          "polar.mysql.x4.medium.c",
					"storage_type":           "ESSDPL0",
					"db_version":             "8.0",
					"db_type":                "MySQL",
					"storage_space":          "20",
					"vswitch_id":             "${alicloud_ens_vswitch.defaultvJRlrU.id}",
					"period":                 "Month",
					"vpc_id":                 "${alicloud_ens_network.defaultE5FDFD.id}",
					"ens_region_id":          "${alicloud_ens_network.defaultE5FDFD.ens_region_id}",
					"used_time":              "1",
					"db_cluster_description": "ens-dbcluster-basic-test-case",
					"zone_id":                "${alicloud_ens_network.defaultE5FDFD.ens_region_id}",
					"storage_pay_type":       "Prepaid",
					"db_node_num":            "1",
					"hot_standby_cluster":    "OFF",
					"storage_auto_scale":     "Disable",
					"cluster_network_type":   "VPC",
					"architecture":           "X86",
					"pay_type":               "Subscription",
					"target_minor_version":   "innovate_x86#20250311",
					"db_minor_version":       "8.0.2",
					"creation_category":      "SENormal",
					"db_node_target_class":   "polar.mysql.x4.medium.c",
					"duration":               "1",
					"auto_renew_status":      "AutoRenewal",
					"period_unit":            "Month",
					"modify_type":            "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_class":          "polar.mysql.x4.medium.c",
						"storage_type":           "ESSDPL0",
						"db_version":             CHECKSET,
						"db_type":                "MySQL",
						"storage_space":          "20",
						"vswitch_id":             CHECKSET,
						"period":                 "Month",
						"vpc_id":                 CHECKSET,
						"ens_region_id":          CHECKSET,
						"used_time":              CHECKSET,
						"db_cluster_description": "ens-dbcluster-basic-test-case",
						"zone_id":                CHECKSET,
						"storage_pay_type":       "Prepaid",
						"db_node_num":            "1",
						"hot_standby_cluster":    "OFF",
						"storage_auto_scale":     "Disable",
						"cluster_network_type":   "VPC",
						"architecture":           "X86",
						"pay_type":               "Subscription",
						"target_minor_version":   "innovate_x86#20250311",
						"db_minor_version":       "8.0.2",
						"creation_category":      "SENormal",
						"db_node_target_class":   "polar.mysql.x4.medium.c",
						"duration":               CHECKSET,
						"auto_renew_status":      "AutoRenewal",
						"period_unit":            "Month",
						"modify_type":            "Upgrade",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"storage_space": "30",
					"modify_type":   "UPGRADE",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"storage_space": "30",
						"modify_type":   "UPGRADE",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_renew_status": "Normal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_renew_status": "Normal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"db_node_target_class": "polar.mysql.x8.medium.c",
					"modify_type":          "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_node_target_class": "polar.mysql.x8.medium.c",
						"modify_type":          "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"architecture", "cluster_network_type", "db_node", "db_node_class", "db_node_num", "ens_region_id", "hot_standby_cluster", "period", "storage_auto_scale", "target_minor_version", "used_time"},
			},
		},
	})
}

var AlicloudPolardbZonalDbClusterMap11408 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudPolardbZonalDbClusterBasicDependence11408(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "ens_region_id" {
  default = "sg-singapore-9"
}

resource "alicloud_ens_network" "defaultE5FDFD" {
  ens_region_id = "sg-singapore-9"
}

resource "alicloud_ens_vswitch" "defaultvJRlrU" {
  ens_region_id = alicloud_ens_network.defaultE5FDFD.ens_region_id
  network_id    = alicloud_ens_network.defaultE5FDFD.id
}


`, name)
}

// Test Polardb ZonalDbCluster. <<< Resource test cases, automatically generated.
