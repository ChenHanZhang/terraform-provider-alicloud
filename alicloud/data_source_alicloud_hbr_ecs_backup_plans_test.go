package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
)

func TestAccAlicloudHbrEcsBackupPlansDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000000, 9999999)

	ecsBackupIdsconf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand, map[string]string{
			"ids": `["${alicloud_hbr_ecs_backup_plan.default.id}"]`,
		}),
		fakeConfig: testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand, map[string]string{
			"ids": `["${alicloud_hbr_ecs_backup_plan.default.id}_fake"]`,
		}),
	}

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand, map[string]string{
			"name_regex": `"${alicloud_hbr_ecs_backup_plan.default.ecs_backup_plan_name}"`,
		}),
		fakeConfig: testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand, map[string]string{
			"name_regex": `"${alicloud_hbr_ecs_backup_plan.default.ecs_backup_plan_name}_fake"`,
		}),
	}

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand, map[string]string{
			"ids":        `["${alicloud_hbr_ecs_backup_plan.default.id}"]`,
			"name_regex": `"${alicloud_hbr_ecs_backup_plan.default.ecs_backup_plan_name}"`,
		}),
		fakeConfig: testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand, map[string]string{
			"ids":        `["${alicloud_hbr_ecs_backup_plan.default.id}_fake"]`,
			"name_regex": `"${alicloud_hbr_ecs_backup_plan.default.ecs_backup_plan_name}_fake"`,
		}),
	}

	HbrEcsBackupPlanCheckInfo.dataSourceTestCheck(t, rand, ecsBackupIdsconf, nameRegexConf, allConf)
}

func testAccCheckAlicloudHbrEcsBackupPlanSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
variable "name" {
	default = "tf-test%d"
}

resource "alicloud_hbr_ecs_backup_plan" "default" {
  ecs_backup_plan_name = var.name
  instance_id =          "i-bp1567rc0o5rz8bxnltz"
  schedule =            "I|1602673264|PT2H"
  backup_type =         "COMPLETE"
  vault_id =            "v-0003gxoksflhu46w185s"
  retention =			"2"
}

data "alicloud_hbr_ecs_backup_plans" "default" {
%s
}
`, rand, strings.Join(pairs, "\n   "))
	return config
}

var existHbrEcsBackupPlanMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"plans.#":                      "1",
		"plans.0.id":                   CHECKSET,
		"plans.0.ecs_backup_plan_name": fmt.Sprintf("tf-test%d", rand),
		"plans.0.instance_id":          "i-bp1567rc0o5rz8bxnltz",
		"plans.0.schedule":             "I|1602673264|PT2H",
		"plans.0.backup_type":          "COMPLETE",
		"plans.0.vault_id":             "v-0003gxoksflhu46w185s",
	}
}

var fakeHbrEcsBackupPlanMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"plans.#": "0",
	}
}

var HbrEcsBackupPlanCheckInfo = dataSourceAttr{
	resourceId:   "data.alicloud_hbr_ecs_backup_plans.default",
	existMapFunc: existHbrEcsBackupPlanMapFunc,
	fakeMapFunc:  fakeHbrEcsBackupPlanMapFunc,
}
