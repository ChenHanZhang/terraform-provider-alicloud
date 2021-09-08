package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudDTSSubscriptionJob_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_dts_subscription_job.default"
	ra := resourceAttrInit(resourceId, AlicloudDTSSubscriptionJobMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &DtsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDtsSubscriptionJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sdtssubscriptionjob%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudDTSSubscriptionJobBasicDependence0)
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
					"payment_type":                "PostPaid",
					"source_endpoint_engine_name": "MySQL",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type":                "PostPaid",
						"source_endpoint_engine_name": "MySQL",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"dts_job_name": "tf-testAccCase",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"dts_job_name": "tf-testAccCase",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "tfTestAcc2",
						"For":     "Tftestacc 2",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "tfTestAcc2",
						"tags.For":     "Tftestacc 2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"subscription_data_type_dml": "true",
					"subscription_data_type_ddl": "true",
					"dts_job_name":               "tf-testAccCase2",
					"tags": map[string]string{
						"Created": "tfTestAcc3",
						"For":     "Tftestacc 3",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"subscription_data_type_dml": "true",
						"subscription_data_type_ddl": "true",
						"dts_job_name":               "tf-testAccCase2",
						"tags.%":                     "2",
						"tags.Created":               "tfTestAcc3",
						"tags.For":                   "Tftestacc 3",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true, ImportStateVerifyIgnore: []string{"error_notice", "delay_notice", "error_phone", "job_id", "database_count", "used_time", "auto_pay", "destination_endpoint_engine_name", "instance_class", "delay_rule_time", "sync_architecture", "period", "auto_start", "reserve", "synchronization_direction", "delay_phone", "destination_region", "quantity", "compute_unit"},
			},
		},
	})
}

var AlicloudDTSSubscriptionJobMap0 = map[string]string{
	"auto_pay":                         NOSET,
	"destination_endpoint_engine_name": NOSET,
	"instance_class":                   NOSET,
	"used_time":                        NOSET,
	"delay_rule_time":                  NOSET,
	"sync_architecture":                NOSET,
	"period":                           NOSET,
	"status":                           CHECKSET,
	"auto_start":                       NOSET,
	"reserve":                          NOSET,
	"subscription_data_type_dml":       CHECKSET,
	"synchronization_direction":        NOSET,
	"delay_phone":                      NOSET,
	"destination_region":               NOSET,
	"quantity":                         NOSET,
	"compute_unit":                     NOSET,
	"subscription_data_type_ddl":       CHECKSET,
	"tags.%":                           CHECKSET,
	"error_notice":                     NOSET,
	"delay_notice":                     NOSET,
	"error_phone":                      NOSET,
	"job_id":                           NOSET,
	"database_count":                   NOSET,
}

func AlicloudDTSSubscriptionJobBasicDependence0(name string) string {
	return fmt.Sprintf(` 
variable "name" {
  default = "%s"
}
`, name)
}
