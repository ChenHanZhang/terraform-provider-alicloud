package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudHBRNasRestoreJob_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_hbr_nas_restore_job.default"
	ra := resourceAttrInit(resourceId, AlicloudHBRNasRestoreJobMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &HbrService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeHbrNasRestoreJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%shbrnasrestorejob%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudHBRNasRestoreJobBasicDependence0)
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
					"snapshot_hash":         "e449cf3911ff19aa32311c8d99d456d945fe536d706c2be28e9f765a8d69b69b",
					"vault_id":              "v-0008va0n86482iwkqswb",
					"source_type":           "NAS",
					"restore_type":          "NAS",
					"snapshot_id":           "s-0006ulw1dfq2muo5g9ml",
					"target_file_system_id": "02ddf4ad63",
					"target_create_time":    "1628856766",
					"target_path":           "/",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"snapshot_hash":         "e449cf3911ff19aa32311c8d99d456d945fe536d706c2be28e9f765a8d69b69b",
						"vault_id":              "v-0008va0n86482iwkqswb",
						"source_type":           "NAS",
						"restore_type":          "NAS",
						"snapshot_id":           "s-0006ulw1dfq2muo5g9ml",
						"target_file_system_id": "02ddf4ad63",
						"target_create_time":    "1628856766",
						"target_path":           "/",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"target_container", "exclude", "udm_region_id", "target_bucket", "target_container_cluster_id", "target_prefix", "target_instance_id", "target_path", "include"},
			},
		},
	})
}

var AlicloudHBRNasRestoreJobMap0 = map[string]string{
	"target_data_source_id":       "",
	"include":                     NOSET,
	"status":                      CHECKSET,
	"target_container":            NOSET,
	"exclude":                     NOSET,
	"options":                     "",
	"udm_region_id":               NOSET,
	"cluster_id":                  "",
	"target_bucket":               NOSET,
	"target_client_id":            "",
	"target_container_cluster_id": NOSET,
	"target_prefix":               NOSET,
	"target_instance_id":          NOSET,
	"target_path":                 NOSET,
	"target_create_time":          "1628856766",
	"snapshot_hash":               "e449cf3911ff19aa32311c8d99d456d945fe536d706c2be28e9f765a8d69b69b",
	"vault_id":                    "v-0008va0n86482iwkqswb",
	"source_type":                 "NAS",
	"restore_type":                "NAS",
	"snapshot_id":                 "s-0006ulw1dfq2muo5g9ml",
	"target_file_system_id":       "02ddf4ad63",
}

func AlicloudHBRNasRestoreJobBasicDependence0(name string) string {
	return fmt.Sprintf(` 
variable "name" {
  default = "%s"
}
`, name)
}
