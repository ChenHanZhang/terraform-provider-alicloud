package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/adb"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb Connection. >>> Resource test cases, automatically generated.
// Case Connection测试用例V2 12536
func TestAccAliCloudAdbConnection_basic12536(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_connection.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbConnectionMap12536)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbConnectionBasicDependence12536)
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
					"db_cluster_id":            "${alicloud_adb_db_cluster.createADBCluster.id}",
					"connection_string_prefix": "test-123",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"db_cluster_id":            CHECKSET,
						"connection_string_prefix": "test-123",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"connection_string_prefix": "test-1234",
					"connection_string":        "test-123.analyticdb.pre.rds.aliyuncs.com",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"connection_string_prefix": "test-1234",
						"connection_string":        "test-123.analyticdb.pre.rds.aliyuncs.com",
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

var AlicloudAdbConnectionMap12536 = map[string]string{
	"ip_address": CHECKSET,
}

func AlicloudAdbConnectionBasicDependence12536(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_adb_db_cluster" "createADBCluster" {
  disk_performance_level = "PL1"
  db_cluster_version     = "3.0"
  db_node_count          = "1"
  payment_type           = "PayAsYouGo"
  db_node_storage        = "100"
  zone_id                = "cn-beijing-k"
  mode                   = "reserver"
  db_node_class          = "C8"
  db_cluster_category    = "Cluster"
}


`, name)
}

// Test Adb Connection. <<< Resource test cases, automatically generated.

func TestAccAlicloudADBConnectionConfig(t *testing.T) {
	var v *adb.Address
	rand := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("tf-testAccAdbConnection%s", rand)
	var basicMap = map[string]string{
		"db_cluster_id":     CHECKSET,
		"connection_string": CHECKSET,
		"ip_address":        CHECKSET,
		"port":              CHECKSET,
	}
	resourceId := "alicloud_adb_connection.default"
	ra := resourceAttrInit(resourceId, basicMap)
	serviceFunc := func() interface{} {
		return &AdbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, serviceFunc, "DescribeAdbConnection")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceAdbConnectionConfigDependence)
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
					"db_cluster_id":     "${alicloud_adb_db_cluster.cluster.id}",
					"connection_prefix": fmt.Sprintf("tf-testacc%s", rand),
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(nil),
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

func resourceAdbConnectionConfigDependence(name string) string {
	return fmt.Sprintf(`
	%s
	variable "creation" {
		default = "ADB"
	}

	variable "name" {
		default = "%s"
	}

	resource "alicloud_adb_db_cluster" "cluster" {
	db_cluster_category = "MixedStorage"
	mode = "flexible"
	compute_resource = "8Core32GB"
	vswitch_id              = local.vswitch_id
	description             = "${var.name}"
	}`, AdbCommonTestCase, name)
}
