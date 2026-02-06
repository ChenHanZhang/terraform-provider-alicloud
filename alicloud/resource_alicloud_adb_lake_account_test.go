package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Adb LakeAccount. >>> Resource test cases, automatically generated.
// Case 湖仓账号测试用例 5218
func TestAccAliCloudAdbLakeAccount_basic5218(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_adb_lake_account.default"
	ra := resourceAttrInit(resourceId, AlicloudAdbLakeAccountMap5218)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AdbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAdbLakeAccount")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccadb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAdbLakeAccountBasicDependence5218)
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
					"account_description": "test_tf_des",
					"db_cluster_id":       "${alicloud_adb_db_cluster_lake_version.CreateInstance.id}",
					"account_type":        "Super",
					"account_name":        "tfnormal",
					"account_password":    "normal@2022",
					"account_privileges": []map[string]interface{}{
						{
							"privilege_type": "Column",
							"privilege_object": []map[string]interface{}{
								{
									"database": "MYSQL",
									"table":    "COLUMNS_PRIV",
									"column":   "DB",
								},
							},
							"privileges": []string{
								"create", "select", "update"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description":  "test_tf_des",
						"db_cluster_id":        CHECKSET,
						"account_type":         "Super",
						"account_name":         "tfnormal",
						"account_password":     "normal@2022",
						"account_privileges.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_description": "test_tf_des2",
					"account_password":    "normal@2023",
					"account_privileges": []map[string]interface{}{
						{
							"privilege_type": "Database",
							"privilege_object": []map[string]interface{}{
								{
									"database": "MYSQL",
								},
							},
							"privileges": []string{
								"select"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_description":  "test_tf_des2",
						"account_password":     "normal@2023",
						"account_privileges.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_privileges": []map[string]interface{}{
						{
							"privilege_type": "Column",
							"privilege_object": []map[string]interface{}{
								{
									"database": "MYSQL",
									"table":    "COLUMNS_PRIV",
									"column":   "HOST",
								},
							},
							"privileges": []string{
								"select"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_privileges.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_privileges": []map[string]interface{}{
						{
							"privilege_type": "Database",
							"privilege_object": []map[string]interface{}{
								{
									"database": "MYSQL",
								},
							},
							"privileges": []string{
								"create", "select", "update"},
						},
						{
							"privilege_type": "Table",
							"privilege_object": []map[string]interface{}{
								{
									"database": "INFORMATION_SCHEMA",
									"table":    "ENGINES",
								},
							},
							"privileges": []string{
								"update"},
						},
						{
							"privilege_type": "Column",
							"privilege_object": []map[string]interface{}{
								{
									"database": "INFORMATION_SCHEMA",
									"table":    "COLUMNS",
									"column":   "PRIVILEGES",
								},
							},
							"privileges": []string{
								"update"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_privileges.#": "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"account_privileges": REMOVEKEY,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_privileges.#": "0",
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

var AlicloudAdbLakeAccountMap5218 = map[string]string{
	"status": CHECKSET,
}

func AlicloudAdbLakeAccountBasicDependence5218(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "VPCID" {
  dry_run     = false
  enable_ipv6 = false
  vpc_name    = "湖仓资源组tf882"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "VSWITCHID" {
  vpc_id       = alicloud_vpc.VPCID.id
  zone_id      = "cn-hangzhou-k"
  vswitch_name = "湖仓资源组tf262"
  cidr_block   = "172.16.0.0/24"
}

resource "alicloud_adb_db_cluster_lake_version" "CreateInstance" {
  storage_resource              = "0ACU"
  zone_id                       = "cn-hangzhou-k"
  vpc_id                        = alicloud_vpc.VPCID.id
  vswitch_id                    = alicloud_vswitch.VSWITCHID.id
  db_cluster_description        = "tf自动化测试-杭州-资源组"
  compute_resource              = "16ACU"
  db_cluster_version            = "5.0"
  payment_type                  = "Postpaid"
  period                        = "Month"
  enable_default_resource_group = false
  security_ips                  = "127.0.0.1"
}


`, name)
}

// Test Adb LakeAccount. <<< Resource test cases, automatically generated.
