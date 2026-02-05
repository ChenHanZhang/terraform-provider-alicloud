package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Eflo ExperimentPlan. >>> Resource test cases, automatically generated.
// Case 测试计划_V4_线上 10948
func TestAccAliCloudEfloExperimentPlan_basic10948(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10948)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10948)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-wulanchabu"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${alicloud_eflo_resource.defaultDeWi22.id}",
					"plan_name":         "测试",
					"template_id":       "${alicloud_eflo_experiment_plan_template.defaultpSZN7t.id}",
					"external_params": map[string]interface{}{
						"\"node\"": "test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"plan_name":         "测试",
						"template_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试模版",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试模版",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10948 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10948(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "region_id" {
  default = "cn-wulanchabu"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_eflo_experiment_plan_template" "defaultpSZN7t" {
  template_pipeline {
    workload_id   = "2"
    workload_name = "MatMul"
    setting_params {
    }
    env_params {
      cpu_per_worker     = "90"
      gpu_per_worker     = "8"
      memory_per_worker  = "500"
      share_memory       = "500"
      worker_num         = "1"
      py_torch_version   = "1"
      gpu_driver_version = "1"
      cuda_version       = "1"
      nccl_version       = "1"
    }
    pipeline_order = "1"
    scene          = "baseline"
  }
  privacy_level        = "private"
  template_name        = "1770294188"
  template_description = "测试"
}

resource "alicloud_eflo_resource" "defaultDeWi22" {
  user_access_param {
    access_id    = "dev"
    access_key   = "P89WYVaGQxAHrRHdjVSR"
    workspace_id = "wswjdmh2sbzqnb4s"
    endpoint     = "pai-proxy.c319fe2aa225641b19779ab68533a09a2.cn-wulanchabu.alicontainer.com"
  }
  cluster_id = "20260205202002020808"
  machine_types {
    memory_info  = "32x 64GB DDR4 4800 Memory"
    type         = "Private"
    bond_num     = "5"
    node_count   = "1"
    cpu_info     = "2x Intel Saphhire Rapid 8469C 48C CPU"
    network_info = "1x 200Gbps Dual Port BF3 DPU for VPC\\n4x 200Gbps Dual Port EIC"
    gpu_info     = "8x OAM 810 GPU"
    disk_info    = "2x 480GB SATA SSD\\n4x 3.84TB NVMe SSD"
    network_mode = "net"
    name         = "lingjun"
  }
  cluster_name = "test"
  cluster_desc = "测试集群"
}


`, name)
}

// Case 测试计划_V3_线上 10685
func TestAccAliCloudEfloExperimentPlan_basic10685(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10685)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10685)
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
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${alicloud_eflo_resource.defaultDeWi22.id}",
					"plan_name":         "测试",
					"template_id":       "${alicloud_eflo_experiment_plan_template.defaultpSZN7t.id}",
					"external_params": map[string]interface{}{
						"\"node\"": "test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"plan_name":         "测试",
						"template_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试模版",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试模版",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10685 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10685(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_eflo_experiment_plan_template" "defaultpSZN7t" {
  template_pipeline {
    workload_id   = "2"
    workload_name = "MatMul"
    setting_params {
    }
    env_params {
      cpu_per_worker     = "90"
      gpu_per_worker     = "8"
      memory_per_worker  = "500"
      share_memory       = "500"
      worker_num         = "1"
      py_torch_version   = "1"
      gpu_driver_version = "1"
      cuda_version       = "1"
      nccl_version       = "1"
    }
    pipeline_order = "1"
    scene          = "baseline"
  }
  privacy_level        = "private"
  template_name        = "1770294190"
  template_description = "测试"
}

resource "alicloud_eflo_resource" "defaultDeWi22" {
  user_access_param {
    access_id    = "dev"
    access_key   = "P89WYVaGQxAHrRHdjVSR"
    workspace_id = "wswjdmh2sbzqnb4s"
    endpoint     = "pai-proxy.c319fe2aa225641b19779ab68533a09a2.cn-wulanchabu.alicontainer.com"
  }
  cluster_id = "20260205202002021010"
  machine_types {
    memory_info  = "32x 64GB DDR4 4800 Memory"
    type         = "Private"
    bond_num     = "5"
    node_count   = "1"
    cpu_info     = "2x Intel Saphhire Rapid 8469C 48C CPU"
    network_info = "1x 200Gbps Dual Port BF3 DPU for VPC\\n4x 200Gbps Dual Port EIC"
    gpu_info     = "8x OAM 810 GPU"
    disk_info    = "2x 480GB SATA SSD\\n4x 3.84TB NVMe SSD"
    network_mode = "net"
    name         = "lingjun"
  }
  cluster_name = "test"
  cluster_desc = "测试集群"
}


`, name)
}

// Case 测试计划_V3_预发 10683
func TestAccAliCloudEfloExperimentPlan_basic10683(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10683)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10683)
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
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${alicloud_eflo_resource.defaultDeWi22.id}",
					"plan_name":         "测试",
					"template_id":       "${var.template_id}",
					"external_params": map[string]interface{}{
						"\"node\"": "test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"plan_name":         "测试",
						"template_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试模版",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试模版",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10683 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10683(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_eflo_experiment_plan_template" "defaultpSZN7t" {
  template_pipeline {
    workload_id   = "2"
    workload_name = "MatMul"
    setting_params {
    }
    env_params {
      cpu_per_worker     = "90"
      gpu_per_worker     = "8"
      memory_per_worker  = "500"
      share_memory       = "500"
      worker_num         = "1"
      py_torch_version   = "1"
      gpu_driver_version = "1"
      cuda_version       = "1"
      nccl_version       = "1"
    }
    pipeline_order = "1"
    scene          = "baseline"
  }
  privacy_level        = "private"
  template_name        = "1770294192"
  template_description = "测试"
}

resource "alicloud_eflo_resource" "defaultDeWi22" {
  user_access_param {
    access_id    = "dev"
    access_key   = "P89WYVaGQxAHrRHdjVSR"
    workspace_id = "wswjdmh2sbzqnb4s"
    endpoint     = "pai-proxy.c319fe2aa225641b19779ab68533a09a2.cn-wulanchabu.alicontainer.com"
  }
  cluster_id = "20260205202002021212"
  machine_types {
    memory_info  = "32x 64GB DDR4 4800 Memory"
    type         = "Private"
    bond_num     = "5"
    node_count   = "1"
    cpu_info     = "2x Intel Saphhire Rapid 8469C 48C CPU"
    network_info = "1x 200Gbps Dual Port BF3 DPU for VPC\\n4x 200Gbps Dual Port EIC"
    gpu_info     = "8x OAM 810 GPU"
    disk_info    = "2x 480GB SATA SSD\\n4x 3.84TB NVMe SSD"
    network_mode = "net"
    name         = "lingjun"
  }
  cluster_name = "test"
  cluster_desc = "测试集群"
}


`, name)
}

// Case 测试计划_V2_预发_测试 10572
func TestAccAliCloudEfloExperimentPlan_basic10572(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10572)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10572)
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
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${alicloud_eflo_resource.defaultDeWi22.id}",
					"plan_name":         "测试",
					"template_id":       "${var.template_id}",
					"external_params": map[string]interface{}{
						"\"node\"": "test",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"plan_name":         "测试",
						"template_id":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试模版",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试模版",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10572 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10572(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_eflo_experiment_plan_template" "defaultpSZN7t" {
  template_pipeline {
    workload_id   = "2"
    workload_name = "MatMul"
    setting_params {
    }
    env_params {
      cpu_per_worker     = "90"
      gpu_per_worker     = "8"
      memory_per_worker  = "500"
      share_memory       = "500"
      worker_num         = "1"
      py_torch_version   = "1"
      gpu_driver_version = "1"
      cuda_version       = "1"
      nccl_version       = "1"
    }
    pipeline_order = "1"
    scene          = "baseline"
  }
  privacy_level        = "private"
  template_name        = "1770294193"
  template_description = "测试"
}

resource "alicloud_eflo_resource" "defaultDeWi22" {
  user_access_param {
    access_id    = "cnp-test"
    access_key   = "cnp-test"
    workspace_id = "wswjdmh2sbzqnb4s"
    endpoint     = "pai-proxy.test.cn-wulanchabu.alicontainer.com"
  }
  cluster_id = "20260205202002021313"
  machine_types {
    memory_info  = "32x 64GB DDR4 4800 Memory"
    type         = "Private"
    bond_num     = "5"
    node_count   = "1"
    cpu_info     = "2x Intel Saphhire Rapid 8469C 48C CPU"
    network_info = "1x 200Gbps Dual Port BF3 DPU for VPC\\n4x 200Gbps Dual Port EIC"
    gpu_info     = "8x OAM 810 GPU"
    disk_info    = "2x 480GB SATA SSD\\n4x 3.84TB NVMe SSD"
    network_mode = "net"
    name         = "lingjun"
  }
  cluster_name = "test"
  cluster_desc = "测试集群"
}


`, name)
}

// Case 测试计划_接入TAG_副本1742033263565 10542
func TestAccAliCloudEfloExperimentPlan_basic10542(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10542)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10542)
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
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${var.resource_id}",
					"template_id":       "${var.template_id}",
					"plan_name":         "测试",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"template_id":       CHECKSET,
						"plan_name":         "测试",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试模版",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试模版",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10542 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10542(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "resource_id" {
  default = <<EOF
36
EOF
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "external_params" {
  default = ""
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case 测试计划_接入TAG 10178
func TestAccAliCloudEfloExperimentPlan_basic10178(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10178)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10178)
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
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${var.resource_id}",
					"template_id":       "${var.template_id}",
					"plan_name":         "测试",
					"external_params":   map[string]interface{}{},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"template_id":       CHECKSET,
						"plan_name":         "测试",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试模版",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试模版",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10178 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10178(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "resource_id" {
  default = <<EOF
36
EOF
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "external_params" {
  default = ""
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case 实验计划测试用例 6045
func TestAccAliCloudEfloExperimentPlan_basic6045(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap6045)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence6045)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_id": "30",
					"template_id": "277",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_id": CHECKSET,
						"template_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap6045 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence6045(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "resource_id" {
  default = <<EOF
30
EOF
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "external_params" {
  default = <<EOF
{}
EOF
}

variable "template_id" {
  default = <<EOF
242
EOF
}


`, name)
}

// Case 实验计划测试用例_资源组测试 6145
func TestAccAliCloudEfloExperimentPlan_basic6145(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap6145)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence6145)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_id":       "30",
					"template_id":       "277",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_id":       CHECKSET,
						"template_id":       CHECKSET,
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap6145 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence6145(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "resource_id" {
  default = <<EOF
30
EOF
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "external_params" {
  default = <<EOF
{}
EOF
}

variable "template_id" {
  default = <<EOF
242
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case 实验计划测试用例_资源组测试_覆盖率 6305
func TestAccAliCloudEfloExperimentPlan_basic6305(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap6305)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence6305)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_id":       "${var.resource_id}",
					"template_id":       "${var.template_id}",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_id":       CHECKSET,
						"template_id":       CHECKSET,
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap6305 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence6305(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "resource_id" {
  default = <<EOF
36
EOF
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "external_params" {
  default = <<EOF
{}
EOF
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case 测试计划_副本1736992623067 10041
func TestAccAliCloudEfloExperimentPlan_basic10041(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_eflo_experiment_plan.default"
	ra := resourceAttrInit(resourceId, AlicloudEfloExperimentPlanMap10041)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EfloServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEfloExperimentPlan")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacceflo%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEfloExperimentPlanBasicDependence10041)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-heyuan"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"resource_id":       "${var.resource_id}",
					"external_params":   map[string]interface{}{},
					"template_id":       "${var.template_id}",
					"plan_name":         "测试1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"resource_id":       CHECKSET,
						"external_params":   CHECKSET,
						"template_id":       CHECKSET,
						"plan_name":         "测试1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"plan_name":         "测试用例",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"plan_name":         "测试用例",
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
				ImportStateVerifyIgnore: []string{"external_params"},
			},
		},
	})
}

var AlicloudEfloExperimentPlanMap10041 = map[string]string{
	"create_time": CHECKSET,
}

func AlicloudEfloExperimentPlanBasicDependence10041(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "resource_id" {
  default = <<EOF
36
EOF
}

variable "plan_id" {
  default = <<EOF
354
EOF
}

variable "experiment_plan_id" {
  default = <<EOF
354
EOF
}

variable "external_params" {
  default = ""
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "template_id" {
  default = <<EOF
54
EOF
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test Eflo ExperimentPlan. <<< Resource test cases, automatically generated.
