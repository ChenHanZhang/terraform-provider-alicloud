package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Arms EnvCustomJob. >>> Resource test cases, automatically generated.
// Case customjob-ecs 4605
func TestAccAliCloudArmsEnvCustomJob_basic4605(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_arms_env_custom_job.default"
	ra := resourceAttrInit(resourceId, AlicloudArmsEnvCustomJobMap4605)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ArmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeArmsEnvCustomJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccarms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudArmsEnvCustomJobBasicDependence4605)
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
					"status":              "run",
					"environment_id":      "${alicloud_arms_environment.env-ecs.id}",
					"env_custom_job_name": name,
					"config_yaml":         `scrape_configs:\n- job_name: job-demo1\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n- job_name: job-demo2\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n  http_sd_configs:\n  - url: 127.0.0.1:9090\n    refresh_interval: 30s`,
					"aliyun_lang":         "zh",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":              "run",
						"environment_id":      CHECKSET,
						"env_custom_job_name": name,
						"config_yaml":         CHECKSET,
						"aliyun_lang":         "zh",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":      "stop",
					"config_yaml": `scrape_configs:\n- job_name: job-demo1\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n- job_name: job-demo2\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n  http_sd_configs:\n  - url: 127.0.0.1:9090\n    refresh_interval: 31s`,
					"aliyun_lang": "en",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":      "stop",
						"config_yaml": CHECKSET,
						"aliyun_lang": "en",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"aliyun_lang"},
			},
		},
	})
}

var AlicloudArmsEnvCustomJobMap4605 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudArmsEnvCustomJobBasicDependence4605(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  description = "api-resource-sub-test-hz-job"
  cidr_block  = "172.16.0.0/12"
  vpc_name    = "api-resource-sub-test-hz-job"
}

resource "alicloud_arms_environment" "env-ecs" {
  environment_type     = "ECS"
  environment_name     = "api-resource-ecs-sub-test-hz-job"
  bind_resource_id     = alicloud_vpc.vpc.id
  environment_sub_type = "ECS"
}


`, name)
}

// Case customjob-cs 4554
func TestAccAliCloudArmsEnvCustomJob_basic4554(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_arms_env_custom_job.default"
	ra := resourceAttrInit(resourceId, AlicloudArmsEnvCustomJobMap4554)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &ArmsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeArmsEnvCustomJob")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccarms%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudArmsEnvCustomJobBasicDependence4554)
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
					"status":              "run",
					"environment_id":      "${alicloud_arms_environment.env-cs.id}",
					"env_custom_job_name": name,
					"config_yaml":         `scrape_configs:\n- job_name: job-demo1\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n- job_name: job-demo2\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n  http_sd_configs:\n  - url: 127.0.0.1:9090\n    refresh_interval: 30s`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":              "run",
						"environment_id":      CHECKSET,
						"env_custom_job_name": name,
						"config_yaml":         CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":      "stop",
					"config_yaml": `scrape_configs:\n- job_name: job-demo1\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n- job_name: job-demo2\n  honor_timestamps: false\n  honor_labels: false\n  scrape_interval: 30s\n  scheme: http\n  metrics_path: /metric\n  static_configs:\n  - targets:\n    - 127.0.0.1:9090\n  http_sd_configs:\n  - url: 127.0.0.1:9090\n    refresh_interval: 31s`,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":      "stop",
						"config_yaml": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"aliyun_lang"},
			},
		},
	})
}

var AlicloudArmsEnvCustomJobMap4554 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudArmsEnvCustomJobBasicDependence4554(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpc" {
  description = "api-resource-sub-test1-hz-job"
  cidr_block  = "172.16.0.0/12"
  vpc_name    = "api-resource-sub-test1-hz-job"
}

resource "alicloud_vswitch" "vsw" {
  description  = "api-resource-test1-hz"
  vpc_id       = alicloud_vpc.vpc.id
  vswitch_name = "api-resource-test1-hz"
  zone_id      = "cn-hangzhou-k"
  cidr_block   = "172.16.0.0/24"
}

resource "alicloud_cs_managed_kubernetes" "ask" {
  deletion_protection = false
  cluster_spec        = "ack.pro.small"
  service_cidr        = "192.168.0.0/24"
  profile             = "Serverless"
}

resource "alicloud_arms_environment" "env-cs" {
  environment_type     = "CS"
  environment_name     = "arms-prom-api-sub-resource1-hz-job"
  bind_resource_id     = alicloud_cs_managed_kubernetes.ask.id
  environment_sub_type = "ACK"
}


`, name)
}

// Test Arms EnvCustomJob. <<< Resource test cases, automatically generated.
