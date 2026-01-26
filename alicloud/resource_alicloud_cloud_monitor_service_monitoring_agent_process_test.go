package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAliCloudCloudMonitorServiceMonitoringAgentProcess_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_monitor_service_monitoring_agent_process.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceMonitoringAgentProcessMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceMonitoringAgentProcess")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudmonitorservicemonitoringagentprocess%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence0)
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
					"instance_id":  "${alicloud_instance.default.id}",
					"process_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":  CHECKSET,
						"process_name": name,
					}),
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

func TestAccAliCloudCloudMonitorServiceMonitoringAgentProcess_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_monitor_service_monitoring_agent_process.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceMonitoringAgentProcessMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceMonitoringAgentProcess")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%scloudmonitorservicemonitoringagentprocess%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence0)
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
					"instance_id":  "${alicloud_instance.default.id}",
					"process_name": name,
					"process_user": "root",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_id":  CHECKSET,
						"process_name": name,
						"process_user": "root",
					}),
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

var AlicloudCloudMonitorServiceMonitoringAgentProcessMap0 = map[string]string{
	"process_id": CHECKSET,
}

func AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence0(name string) string {
	return fmt.Sprintf(`
	variable "name" {
    	default = "%s"
	}

	data "alicloud_zones" "default" {
  		available_disk_category     = "cloud_efficiency"
  		available_resource_creation = "VSwitch"
	}

	data "alicloud_instance_types" "default" {
  		availability_zone    = data.alicloud_zones.default.zones.0.id
  		instance_type_family = "ecs.sn1ne"
	}

	data "alicloud_images" "default" {
  		name_regex  = "^ubuntu_[0-9]+_[0-9]+_x64*"
  		most_recent = true
  		owners      = "system"
	}

	resource "alicloud_vpc" "default" {
  		vpc_name   = var.name
  		cidr_block = "172.16.0.0/16"
	}

	resource "alicloud_vswitch" "default" {
  		vpc_id       = alicloud_vpc.default.id
  		cidr_block   = "172.16.0.0/24"
  		zone_id      = data.alicloud_zones.default.zones.0.id
  		vswitch_name = var.name
	}

	resource "alicloud_security_group" "default" {
  		name   = var.name
  		vpc_id = alicloud_vswitch.default.vpc_id
	}

	resource "alicloud_instance" "default" {
  		image_id                   = data.alicloud_images.default.images.0.id
  		instance_type              = data.alicloud_instance_types.default.instance_types.0.id
  		instance_name              = var.name
  		security_groups            = alicloud_security_group.default.*.id
  		internet_charge_type       = "PayByTraffic"
  		internet_max_bandwidth_out = "10"
  		availability_zone          = data.alicloud_zones.default.zones.0.id
  		instance_charge_type       = "PostPaid"
  		system_disk_category       = "cloud_efficiency"
  		vswitch_id                 = alicloud_vswitch.default.id
	}
`, name)
}

// Test CloudMonitorService MonitoringAgentProcess. >>> Resource test cases, automatically generated.
// Case MonitoringAgentProcess资源用例_副本1687758154615_副本1697079153930 4792
func TestAccAliCloudCloudMonitorServiceMonitoringAgentProcess_basic4792(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_monitor_service_monitoring_agent_process.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceMonitoringAgentProcessMap4792)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceMonitoringAgentProcess")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudmonitorservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence4792)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"process_name": "saposcol",
					"instance_id":  "i-bp11lgn8z7syr0kktcus",
					"process_user": "root",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"process_name": "saposcol",
						"instance_id":  "i-bp11lgn8z7syr0kktcus",
						"process_user": "root",
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

var AlicloudCloudMonitorServiceMonitoringAgentProcessMap4792 = map[string]string{
	"process_id": CHECKSET,
}

func AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence4792(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case MonitoringAgentProcess资源用例_副本1687758154615 3482
func TestAccAliCloudCloudMonitorServiceMonitoringAgentProcess_basic3482(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_monitor_service_monitoring_agent_process.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceMonitoringAgentProcessMap3482)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceMonitoringAgentProcess")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudmonitorservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence3482)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-beijing"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"process_name": "saposcol",
					"instance_id":  "i-bp16nxdu7n2u9t1qqmga",
					"process_user": "root",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"process_name": "saposcol",
						"instance_id":  "i-bp16nxdu7n2u9t1qqmga",
						"process_user": "root",
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

var AlicloudCloudMonitorServiceMonitoringAgentProcessMap3482 = map[string]string{
	"process_id": CHECKSET,
}

func AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence3482(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case MonitoringAgentProcess资源用例 2756
func TestAccAliCloudCloudMonitorServiceMonitoringAgentProcess_basic2756(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_monitor_service_monitoring_agent_process.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudMonitorServiceMonitoringAgentProcessMap2756)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudMonitorServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudMonitorServiceMonitoringAgentProcess")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudmonitorservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence2756)
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
					"process_name": "saposcol",
					"instance_id":  "i-bp16nxdu7n2u9t1qqmga",
					"process_user": "root",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"process_name": "saposcol",
						"instance_id":  "i-bp16nxdu7n2u9t1qqmga",
						"process_user": "root",
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

var AlicloudCloudMonitorServiceMonitoringAgentProcessMap2756 = map[string]string{
	"process_id": CHECKSET,
}

func AlicloudCloudMonitorServiceMonitoringAgentProcessBasicDependence2756(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test CloudMonitorService MonitoringAgentProcess. <<< Resource test cases, automatically generated.
