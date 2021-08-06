package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudHbrNasBackupPlansDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000000, 9999999)

	nasBackupIdsconf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand, map[string]string{
			"ids": `["${alicloud_hbr_nas_backup_plan.default.id}"]`,
		}),
		fakeConfig: testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand, map[string]string{
			"ids": `["${alicloud_hbr_nas_backup_plan.default.id}_fake"]`,
		}),
	}

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand, map[string]string{
			"name_regex": `"${alicloud_hbr_nas_backup_plan.default.nas_backup_plan_name}"`,
		}),
		fakeConfig: testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand, map[string]string{
			"name_regex": `"${alicloud_hbr_nas_backup_plan.default.nas_backup_plan_name}_fake"`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand, map[string]string{
			"ids":        `["${alicloud_hbr_nas_backup_plan.default.id}"]`,
			"name_regex": `"${alicloud_hbr_nas_backup_plan.default.nas_backup_plan_name}"`,
		}),
		fakeConfig: testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand, map[string]string{
			"ids":        `["${alicloud_hbr_nas_backup_plan.default.id}_fake"]`,
			"name_regex": `"${alicloud_hbr_nas_backup_plan.default.nas_backup_plan_name}_fake"`,
		}),
	}

	HbrNasBackupPlanCheckInfo.dataSourceTestCheck(t, rand, nasBackupIdsconf, nameRegexConf, allConf)
}

func testAccCheckAlicloudHbrNasBackupPlanSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
variable "name" {
	default = "tf-test%d"
}

resource "alicloud_hbr_nas_backup_plan" "default" {
  nas_backup_plan_name = var.name
  file_system_id =      "031cf4964f"
  schedule =            "I|1602673264|PT2H"
  backup_type =         "COMPLETE"
  vault_id =            "v-0003gxoksflhu46w185s"
  create_time =         "1603163444"
  retention =			"2"
}

data "alicloud_hbr_nas_backup_plans" "default" {
%s
}
`, rand, strings.Join(pairs, "\n   "))
	return config
}

var existHbrNasBackupPlanMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"plans.#":                      "1",
		"plans.0.id":                   CHECKSET,
		"plans.0.nas_backup_plan_name": fmt.Sprintf("tf-test%d", rand),
		"plans.0.file_system_id":       "031cf4964f",
		"plans.0.schedule":             "I|1602673264|PT2H",
		"plans.0.backup_type":          "COMPLETE",
		"plans.0.vault_id":             "v-0003gxoksflhu46w185s",
		"plans.0.create_time":          "1603163444",
	}
}

var fakeHbrNasBackupPlanMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"plans.#": "0",
	}
}

var HbrNasBackupPlanCheckInfo = dataSourceAttr{
	resourceId:   "data.alicloud_hbr_nas_backup_plans.default",
	existMapFunc: existHbrNasBackupPlanMapFunc,
	fakeMapFunc:  fakeHbrNasBackupPlanMapFunc,
}
