package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb Account. >>> Resource test cases, automatically generated.
// Case 数仓account测试用例 3881
func TestAccAliCloudAdbAccount_basic3881(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_account.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbAccountMap3881)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbAccountBasicDependence3881)
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
					"account_description": "testtag",
					"db_cluster_id":       "${alicloud_adb_db_cluster.createADBCluster.id}",
					"account_type":        "Super",
					"account_name":        name,
					"account_password":    "Aliyun@123",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description": "testtag",
						"db_cluster_id":       CHECKSET,
						"account_type":        "Super",
						"account_name":        name,
						"account_password":    "Aliyun@123",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_description": "testmodifydesc",
					"account_password":    "Aliyun@1234",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description": "testmodifydesc",
						"account_password":    "Aliyun@1234",
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
				ImportStateVerifyIgnore: []string{"account_password"},
			},
		},
	})
}

var AlicloudAdbAccountMap3881 = map[string]string{
	"status": CHECKSET,
}

func AlicloudAdbAccountBasicDependence3881(name string) string {
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
  db_cluster_category    = "Cluster"
  zone_id                = "cn-beijing-k"
  mode                   = "reserver"
  db_node_class          = "C8"
}


`, name)
}

// Test Adb Account. <<< Resource test cases, automatically generated.
