// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test AliKafka Instance. >>> Resource test cases, automatically generated.
// Case 后付费覆盖-停用-启用-网段-172.16.16.0 10761
func TestAccAliCloudAliKafkaInstance_basic10761(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap10761)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence10761)
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
					"deploy_type":   "5",
					"spec_type":     "normal",
					"disk_type":     "0",
					"disk_size":     "500",
					"io_max_spec":   "alikafka.hw.2xlarge",
					"paid_type":     "1",
					"deploy_module": "vpc",
					"vswitch_ids": []string{
						"${alicloud_vswitch.default7stf9N.id}"},
					"vswitch_id":      "${alicloud_vswitch.default7stf9N.id}",
					"vpc_id":          "${alicloud_vpc.defaultRAZ2YU.id}",
					"service_version": "2.2.0",
					"status":          "15",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":     CHECKSET,
						"spec_type":       "normal",
						"disk_type":       "0",
						"disk_size":       "500",
						"io_max_spec":     "alikafka.hw.2xlarge",
						"paid_type":       "1",
						"deploy_module":   "vpc",
						"vswitch_ids.#":   "1",
						"vswitch_id":      CHECKSET,
						"vpc_id":          CHECKSET,
						"service_version": "2.2.0",
						"status":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "5",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap10761 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence10761(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-后付费自动创建TopicGroup-预发133"
}

variable "region_id" {
  default = "cn-beijing"
}

variable "tag_value" {
  default = "刘江"
}

variable "tagkey" {
  default = "owner"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "172.16.16.0/24"
  zone_id     = "cn-beijing-a"
}


`, name)
}

// Case 创建confluent实例并部署-预付费方式进行升配 9769
func TestAccAliCloudAliKafkaInstance_basic9769(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap9769)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence9769)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-shanghai"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"deploy_type":    "5",
					"spec_type":      "professional",
					"paid_type":      "4",
					"security_group": "${alicloud_security_group.defaultS9n3Ql.resource_group_id}",
					"deploy_module":  "vpc",
					"vpc_id":         "${alicloud_vpc.defaultRAZ2YU.id}",
					"confluent_config": []map[string]interface{}{
						{
							"kafka_storage":            "800",
							"kafka_replica":            "3",
							"zoo_keeper_storage":       "100",
							"zoo_keeper_replica":       "3",
							"control_center_storage":   "300",
							"control_center_replica":   "1",
							"schema_registry_replica":  "2",
							"connect_replica":          "2",
							"ksql_storage":             "100",
							"ksql_replica":             "2",
							"kafka_rest_proxy_replica": "2",
							"kafka_cu":                 "4",
							"control_center_cu":        "4",
							"zoo_keeper_cu":            "2",
							"connect_cu":               "4",
							"kafka_rest_proxy_cu":      "4",
							"schema_registry_cu":       "2",
							"ksql_cu":                  "4",
						},
					},
					"duration":   "1",
					"password":   "LTAI5tMDeoDB86TfQTRsr6m9",
					"vswitch_id": "${alicloud_vswitch.default7stf9N.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":    CHECKSET,
						"spec_type":      "professional",
						"paid_type":      "4",
						"security_group": CHECKSET,
						"deploy_module":  "vpc",
						"vpc_id":         CHECKSET,
						"duration":       "1",
						"password":       "LTAI5tMDeoDB86TfQTRsr6m9",
						"vswitch_id":     CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"confluent_config": []map[string]interface{}{
						{
							"kafka_storage":            "2000",
							"kafka_replica":            "5",
							"zoo_keeper_storage":       "300",
							"control_center_storage":   "400",
							"schema_registry_replica":  "3",
							"connect_replica":          "4",
							"ksql_storage":             "200",
							"ksql_replica":             "4",
							"kafka_rest_proxy_replica": "4",
							"kafka_cu":                 "8",
							"control_center_cu":        "5",
							"zoo_keeper_cu":            "5",
							"connect_cu":               "5",
							"kafka_rest_proxy_cu":      "5",
							"schema_registry_cu":       "5",
							"ksql_cu":                  "5",
							"zoo_keeper_replica":       "3",
							"control_center_replica":   "1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap9769 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence9769(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-创建confluent实例覆盖属性-预发399"
}

variable "region_id" {
  default = "cn-shanghai"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "172.16.118.0/24"
  zone_id     = "cn-shanghai-a"
}

resource "alicloud_security_group" "defaultS9n3Ql" {
  description         = "sg"
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.defaultRAZ2YU.id
  security_group_type = "normal"
}


`, name)
}

// Case 创建预付费实例-部署0.10.2-升级版本-升配预付费实例-网段-10.5.93.0 9765
func TestAccAliCloudAliKafkaInstance_basic9765(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap9765)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence9765)
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
					"deploy_type":    "5",
					"spec_type":      "normal",
					"disk_type":      "0",
					"disk_size":      "500",
					"io_max_spec":    "alikafka.hw.2xlarge",
					"paid_type":      "0",
					"zone_id":        "cn-beijing-a",
					"security_group": "${alicloud_security_group.defaultS9n3Ql.resource_group_id}",
					"deploy_module":  "vpc",
					"vswitch_ids": []string{
						"${alicloud_vswitch.default7stf9N.id}"},
					"vswitch_id":      "${alicloud_vswitch.default7stf9N.id}",
					"vpc_id":          "${alicloud_vpc.defaultRAZ2YU.id}",
					"service_version": "0.10.2",
					"duration":        "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":     CHECKSET,
						"spec_type":       "normal",
						"disk_type":       "0",
						"disk_size":       "500",
						"io_max_spec":     "alikafka.hw.2xlarge",
						"paid_type":       "0",
						"zone_id":         "cn-beijing-a",
						"security_group":  CHECKSET,
						"deploy_module":   "vpc",
						"vswitch_ids.#":   "1",
						"vswitch_id":      CHECKSET,
						"vpc_id":          CHECKSET,
						"service_version": "0.10.2",
						"duration":        "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec_type":       "professional",
					"disk_size":       "1000",
					"io_max_spec":     "alikafka.hw.3xlarge",
					"service_version": "2.2.0",
					"eip_model":       "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec_type":       "professional",
						"disk_size":       "1000",
						"io_max_spec":     "alikafka.hw.3xlarge",
						"service_version": "2.2.0",
						"eip_model":       "false",
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap9765 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence9765(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-beijing"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "10.5.93.0/24"
  zone_id     = "cn-beijing-a"
}

resource "alicloud_security_group" "defaultS9n3Ql" {
  description         = "sg"
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.defaultRAZ2YU.id
  security_group_type = "normal"
}


`, name)
}

// Case 创建后付费实例全生命周期-网段10.97.97.0-10.96.96.0 9751
func TestAccAliCloudAliKafkaInstance_basic9751(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap9751)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence9751)
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
					"deploy_type":          "4",
					"spec_type":            "professional",
					"disk_type":            "0",
					"disk_size":            "500",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"io_max_spec":          "alikafka.hw.2xlarge",
					"paid_type":            "1",
					"eip_max":              "3",
					"zone_id":              "cn-beijing-a",
					"security_group":       "${alicloud_security_group.defaultS9n3Ql.resource_group_id}",
					"partition_num_of_buy": "12",
					"config":               "{\\\"enable.vpc_sasl_ssl\\\":\\\"false\\\",\\\"kafka.log.retention.hours\\\":\\\"33\\\",\\\"kafka.offsets.retention.minutes\\\":\\\"10080\\\",\\\"enable.tiered\\\":\\\"false\\\",\\\"cloud.maxTieredStoreSpace\\\":\\\"0\\\",\\\"enable.acl\\\":\\\"false\\\",\\\"kafka.ssl.bit\\\":\\\"4096\\\",\\\"message.timestamp.type\\\":\\\"CreateTime\\\",\\\"enable.compact\\\":\\\"true\\\",\\\"kafka.message.max.bytes\\\":\\\"1048576\\\",\\\"message.timestamp.difference.max.ms\\\":\\\"9223372036854775807\\\"}",
					"deploy_module":        "eip",
					"vswitch_ids": []string{
						"${alicloud_vswitch.default7stf9N.id}", "${alicloud_vswitch.defaultU3njGI.id}"},
					"vswitch_id":               "${alicloud_vswitch.default7stf9N.id}",
					"instance_name":            name,
					"vpc_id":                   "${alicloud_vpc.defaultRAZ2YU.id}",
					"service_version":          "2.2.0",
					"is_set_user_and_password": "true",
					"cross_zone":               "true",
					"is_force_selected_zones":  "false",
					"password":                 "NmyzhIl4cMpgrKUoZT0yQcRAFRdnM4ta",
					"is_eip_inner":             "true",
					"username":                 "AdministratorAccessAdmini1",
					"notifier":                 "刘江",
					"selected_zones":           "[[\\\"zonea\\\"],[]]",
					"user_phone_num":           "18049581490",
					"kms_key_id":               "${alicloud_kms_key.defaultbxem2C.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":              CHECKSET,
						"spec_type":                "professional",
						"disk_type":                "0",
						"disk_size":                "500",
						"resource_group_id":        CHECKSET,
						"io_max_spec":              "alikafka.hw.2xlarge",
						"paid_type":                "1",
						"eip_max":                  "3",
						"zone_id":                  "cn-beijing-a",
						"security_group":           CHECKSET,
						"partition_num_of_buy":     "12",
						"config":                   CHECKSET,
						"deploy_module":            "eip",
						"vswitch_ids.#":            "2",
						"vswitch_id":               CHECKSET,
						"instance_name":            name,
						"vpc_id":                   CHECKSET,
						"service_version":          "2.2.0",
						"is_set_user_and_password": "true",
						"cross_zone":               "true",
						"is_force_selected_zones":  "false",
						"password":                 "NmyzhIl4cMpgrKUoZT0yQcRAFRdnM4ta",
						"is_eip_inner":             "true",
						"username":                 "AdministratorAccessAdmini1",
						"notifier":                 "刘江",
						"selected_zones":           CHECKSET,
						"user_phone_num":           CHECKSET,
						"kms_key_id":               CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"config":            "{\\\"enable.vpc_sasl_ssl\\\":\\\"false\\\",\\\"kafka.log.retention.hours\\\":\\\"33\\\",\\\"kafka.offsets.retention.minutes\\\":\\\"10080\\\",\\\"enable.tiered\\\":\\\"false\\\",\\\"cloud.maxTieredStoreSpace\\\":\\\"0\\\",\\\"enable.acl\\\":\\\"false\\\",\\\"kafka.ssl.bit\\\":\\\"4096\\\",\\\"message.timestamp.type\\\":\\\"CreateTime\\\",\\\"enable.compact\\\":\\\"true\\\",\\\"kafka.message.max.bytes\\\":\\\"2048576\\\",\\\"message.timestamp.difference.max.ms\\\":\\\"9223372036854775807\\\"}",
					"instance_name":     name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"config":            CHECKSET,
						"instance_name":     name + "_update",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap9751 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence9751(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-创建后付费实例覆盖属性-预发61"
}

variable "region_id" {
  default = "cn-beijing"
}

variable "tag_value" {
  default = "刘江"
}

variable "tagkey" {
  default = "owner"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "10.97.97.0/24"
  zone_id     = "cn-beijing-a"
}

resource "alicloud_vswitch" "defaultU3njGI" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "10.96.96.0/24"
  zone_id     = "cn-beijing-c"
}

resource "alicloud_security_group" "defaultS9n3Ql" {
  description         = "sg"
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.defaultRAZ2YU.id
  security_group_type = "normal"
}

resource "alicloud_kms_instance" "defaultBiZU50" {
  vpc_num         = "1"
  key_num         = "1000"
  renew_period    = "1"
  secret_num      = "1"
  product_version = "3"
  renew_status    = "AutoRenewal"
  vpc_id          = alicloud_vpc.defaultRAZ2YU.id
  vswitch_ids     = ["${alicloud_vswitch.defaultU3njGI.id}"]
  zone_ids        = ["cn-beijing-c"]
  spec            = "1000"
}

resource "alicloud_kms_key" "defaultbxem2C" {
  protection_level = "SOFTWARE"
  key_spec         = "Aliyun_AES_256"
  key_usage        = "ENCRYPT/DECRYPT"
  dkms_instance_id = alicloud_kms_instance.defaultBiZU50.id
  origin           = "Aliyun_KMS"
  status           = "Enabled"
}


`, name)
}

// Case 后付费覆盖自动创建Topic和Group-网段-10.1.1.0 9834
func TestAccAliCloudAliKafkaInstance_basic9834(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap9834)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence9834)
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
					"deploy_type":    "5",
					"spec_type":      "normal",
					"disk_type":      "0",
					"disk_size":      "500",
					"io_max_spec":    "alikafka.hw.2xlarge",
					"paid_type":      "1",
					"security_group": "${alicloud_security_group.defaultS9n3Ql.resource_group_id}",
					"deploy_module":  "vpc",
					"vswitch_ids": []string{
						"${alicloud_vswitch.default7stf9N.id}"},
					"vswitch_id":                  "${alicloud_vswitch.default7stf9N.id}",
					"instance_name":               name,
					"vpc_id":                      "${alicloud_vpc.defaultRAZ2YU.id}",
					"service_version":             "2.2.0",
					"enable_auto_group":           "true",
					"enable_auto_topic":           "enable",
					"default_topic_partition_num": "12",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":                 CHECKSET,
						"spec_type":                   "normal",
						"disk_type":                   "0",
						"disk_size":                   "500",
						"io_max_spec":                 "alikafka.hw.2xlarge",
						"paid_type":                   "1",
						"security_group":              CHECKSET,
						"deploy_module":               "vpc",
						"vswitch_ids.#":               "1",
						"vswitch_id":                  CHECKSET,
						"instance_name":               name,
						"vpc_id":                      CHECKSET,
						"service_version":             "2.2.0",
						"enable_auto_group":           "true",
						"enable_auto_topic":           "enable",
						"default_topic_partition_num": "12",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_auto_group":                  "false",
					"default_topic_partition_num":        "24",
					"update_default_topic_partition_num": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_auto_group":                  "false",
						"default_topic_partition_num":        "24",
						"update_default_topic_partition_num": "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enable_auto_topic": "disable",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enable_auto_topic": "disable",
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap9834 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence9834(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-后付费自动创建TopicGroup-预发77"
}

variable "region_id" {
  default = "cn-beijing"
}

variable "tag_value" {
  default = "刘江"
}

variable "tagkey" {
  default = "owner"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "10.1.1.0/24"
  zone_id     = "cn-beijing-a"
}

resource "alicloud_security_group" "defaultS9n3Ql" {
  description         = "sg"
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.defaultRAZ2YU.id
  security_group_type = "normal"
}


`, name)
}

// Case 后付费覆盖-停用-启用-网段-172.16.0.0 10755
func TestAccAliCloudAliKafkaInstance_basic10755(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap10755)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence10755)
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
					"deploy_type":   "5",
					"spec_type":     "normal",
					"disk_type":     "0",
					"disk_size":     "500",
					"io_max_spec":   "alikafka.hw.2xlarge",
					"paid_type":     "1",
					"deploy_module": "vpc",
					"vswitch_ids": []string{
						"${alicloud_vswitch.default7stf9N.id}"},
					"vswitch_id":      "${alicloud_vswitch.default7stf9N.id}",
					"vpc_id":          "${alicloud_vpc.defaultRAZ2YU.id}",
					"service_version": "2.2.0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":     CHECKSET,
						"spec_type":       "normal",
						"disk_type":       "0",
						"disk_size":       "500",
						"io_max_spec":     "alikafka.hw.2xlarge",
						"paid_type":       "1",
						"deploy_module":   "vpc",
						"vswitch_ids.#":   "1",
						"vswitch_id":      CHECKSET,
						"vpc_id":          CHECKSET,
						"service_version": "2.2.0",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "15",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "5",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": CHECKSET,
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap10755 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence10755(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-后付费自动创建TopicGroup-预发39"
}

variable "region_id" {
  default = "cn-beijing"
}

variable "tag_value" {
  default = "刘江"
}

variable "tagkey" {
  default = "owner"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "172.16.0.0/24"
  zone_id     = "cn-beijing-a"
}


`, name)
}

// Case 后付费转预付费覆盖-公网升配-网段-10.4.1.0 9800
func TestAccAliCloudAliKafkaInstance_basic9800(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap9800)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence9800)
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
					"deploy_type":          "5",
					"spec_type":            "normal",
					"disk_type":            "0",
					"disk_size":            "500",
					"io_max_spec":          "alikafka.hw.2xlarge",
					"paid_type":            "1",
					"zone_id":              "cn-beijing-a",
					"security_group":       "${alicloud_security_group.defaultS9n3Ql.resource_group_id}",
					"partition_num_of_buy": "12",
					"deploy_module":        "vpc",
					"vswitch_ids": []string{
						"${alicloud_vswitch.default7stf9N.id}"},
					"vswitch_id":      "${alicloud_vswitch.default7stf9N.id}",
					"instance_name":   name,
					"vpc_id":          "${alicloud_vpc.defaultRAZ2YU.id}",
					"service_version": "2.2.0",
					"user_phone_num":  "18049581490",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":          CHECKSET,
						"spec_type":            "normal",
						"disk_type":            "0",
						"disk_size":            "500",
						"io_max_spec":          "alikafka.hw.2xlarge",
						"paid_type":            "1",
						"zone_id":              "cn-beijing-a",
						"security_group":       CHECKSET,
						"partition_num_of_buy": "12",
						"deploy_module":        "vpc",
						"vswitch_ids.#":        "1",
						"vswitch_id":           CHECKSET,
						"instance_name":        name,
						"vpc_id":               CHECKSET,
						"service_version":      "2.2.0",
						"user_phone_num":       CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"spec_type":            "professional",
					"disk_size":            "1000",
					"io_max_spec":          "alikafka.hw.3xlarge",
					"partition_num_of_buy": "100",
					"eip_max":              "6",
					"duration":             "2",
					"eip_model":            "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"spec_type":            "professional",
						"disk_size":            "1000",
						"io_max_spec":          "alikafka.hw.3xlarge",
						"partition_num_of_buy": "100",
						"eip_max":              "6",
						"duration":             "2",
						"eip_model":            "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"paid_type": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"paid_type": "0",
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap9800 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence9800(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-创建后付费实例覆盖属性-预发178"
}

variable "region_id" {
  default = "cn-beijing"
}

variable "tag_value" {
  default = "刘江"
}

variable "tagkey" {
  default = "owner"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "10.4.1.0/24"
  zone_id     = "cn-beijing-a"
}

resource "alicloud_security_group" "defaultS9n3Ql" {
  description         = "sg"
  security_group_name = "sg_name"
  vpc_id              = alicloud_vpc.defaultRAZ2YU.id
  security_group_type = "normal"
}


`, name)
}

// Case 创建v3实例并部署-后付费方式进行升配到公网-网段-10.31.5.0 9764
func TestAccAliCloudAliKafkaInstance_basic9764(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_alikafka_instance_v2.default"
	ra := resourceAttrInit(resourceId, AlicloudAliKafkaInstanceMap9764)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &AliKafkaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeAliKafkaInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccalikafka%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudAliKafkaInstanceBasicDependence9764)
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
					"deploy_type":   "4",
					"spec_type":     "normal",
					"paid_type":     "3",
					"deploy_module": "vpc",
					"vswitch_id":    "${alicloud_vswitch.default7stf9N.id}",
					"vpc_id":        "${alicloud_vpc.defaultRAZ2YU.id}",
					"serverless_config": []map[string]interface{}{
						{
							"reserved_publish_capacity":   "60",
							"reserved_subscribe_capacity": "60",
						},
					},
					"zone_id": "cn-beijing-a",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deploy_type":   CHECKSET,
						"spec_type":     "normal",
						"paid_type":     "3",
						"deploy_module": "vpc",
						"vswitch_id":    CHECKSET,
						"vpc_id":        CHECKSET,
						"zone_id":       "cn-beijing-a",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"serverless_config": []map[string]interface{}{
						{
							"reserved_publish_capacity":   "100",
							"reserved_subscribe_capacity": "100",
						},
					},
					"eip_model": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"eip_model": "true",
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
				ImportStateVerifyIgnore: []string{"cross_zone", "deploy_module", "duration", "eip_model", "is_eip_inner", "is_force_selected_zones", "is_set_user_and_password", "notifier", "order_id", "password", "selected_zones", "update_default_topic_partition_num", "user_phone_num", "username"},
			},
		},
	})
}

var AlicloudAliKafkaInstanceMap9764 = map[string]string{
	"is_partition_buy":     CHECKSET,
	"group_left":           CHECKSET,
	"topic_left":           CHECKSET,
	"partition_used":       CHECKSET,
	"group_used":           CHECKSET,
	"domain_endpoint":      CHECKSET,
	"end_point":            CHECKSET,
	"sasl_domain_endpoint": CHECKSET,
	"ssl_domain_endpoint":  CHECKSET,
	"create_time":          CHECKSET,
	"topic_used":           CHECKSET,
	"partition_left":       CHECKSET,
	"region_id":            CHECKSET,
	"topic_num_of_buy":     CHECKSET,
}

func AlicloudAliKafkaInstanceBasicDependence9764(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "实例名" {
  default = "刘江-创建v3实例覆盖属性-预发573"
}

variable "region_id" {
  default = "cn-beijing"
}

variable "tag_value" {
  default = "刘江"
}

variable "tagkey" {
  default = "owner"
}

resource "alicloud_vpc" "defaultRAZ2YU" {
  cidr_block = "10.0.0.0/8"
}

resource "alicloud_vswitch" "default7stf9N" {
  description = "资源用例创建vsw"
  vpc_id      = alicloud_vpc.defaultRAZ2YU.id
  cidr_block  = "10.31.5.0/24"
  zone_id     = "cn-beijing-a"
}


`, name)
}

// Test AliKafka Instance. <<< Resource test cases, automatically generated.
