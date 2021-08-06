package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudHBRNasBackupPlan_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hbr_nas_backup_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudHBRNasBackupPlanMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HbrService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHbrNasBackupPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%shbrnasbackupplan%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHBRNasBackupPlanBasicDependence0)
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
					"backup_type":          "COMPLETE",
					"vault_id":             "v-0003gxoksflhu46w185s",
					"file_system_id":       "031cf4964f",
					"create_time":          "1603163444",
					"schedule":             "I|1602673264|PT2H",
					"nas_backup_plan_name": "tf-testAcc测试",
					"retention":            "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"backup_type":          "COMPLETE",
						"vault_id":             "v-0003gxoksflhu46w185s",
						"file_system_id":       "031cf4964f",
						"create_time":          "1603163444",
						"schedule":             "I|1602673264|PT2H",
						"nas_backup_plan_name": "tf-testAcc测试",
						"retention":            "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"nas_backup_plan_name": "tf-testAcc测试2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"nas_backup_plan_name": "tf-testAcc测试2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"schedule": "I|1602673264|P1D",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"schedule": "I|1602673264|P1D",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"path": []string{"/home/test"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"path.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"nas_backup_plan_name": "tf-testAcc测试3",
					"schedule":             "I|1602673264|PT2H",
					"retention":            "4",
					"path":                 []string{"/home"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"nas_backup_plan_name": "tf-testAcc测试3",
						"schedule":             "I|1602673264|PT2H",
						"retention":            "4",
						"path.#":               "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"update_paths", "disk_id", "file_system_id", "resource", "bucket", "prefix", "rule", "udm_region_id"},
			},
		},
	})
}

var AlicloudHBRNasBackupPlanMap0 = map[string]string{
	"path.#":               NOSET,
	"retention":            "",
	"disk_id":              NOSET,
	"options":              "",
	"exclude":              NOSET,
	"resource":             NOSET,
	"rule":                 NOSET,
	"file_system_id":       "031cf4964f",
	"udm_region_id":        NOSET,
	"speed_limit":          NOSET,
	"include":              NOSET,
	"detail":               NOSET,
	"prefix":               NOSET,
	"update_paths":         NOSET,
	"bucket":               NOSET,
	"instance_id":          NOSET,
	"create_time":          "1603163444",
	"schedule":             "I|1602673264|PT2H",
	"nas_backup_plan_name": "tf-testAcc测试",
	"backup_type":          "COMPLETE",
	"vault_id":             "v-0003gxoksflhu46w185s",
}

func AlicloudHBRNasBackupPlanBasicDependence0(name string) string {
	return fmt.Sprintf(` 
variable "name" {
  default = "%s"
}
`, name)
}
