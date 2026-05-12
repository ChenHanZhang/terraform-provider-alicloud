package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Rds Custom. >>> Resource test cases, automatically generated.
// Case resourceCase_20260415_01_clone_1_clone_0 12788
func TestAccAliCloudRdsCustom_basic12788(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap12788)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence12788)
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
					"description":          "ran资源用名称050801",
					"instance_charge_type": "PostPaid",
					"auto_renew":           "false",
					"amount":               "1",
					"vswitch_id":           "${alicloud_vswitch.vswRC_clone_1_clone_0.id}",
					"dry_run":              "false",
					"deletion_protection":  "true",
					"auto_pay":             "true",
					"security_group_ids": []string{
						"${alicloud_security_group.secRC_clone_1_clone_0.id}"},
					"system_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "40",
							"performance_level": "PL1",
						},
					},
					"instance_name": "namecreate",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"performance_level": "PL0",
							"size":              "40",
						},
					},
					"create_mode":        "0",
					"enable_jumbo_frame": "true",
					"instance_type":      "mysql.x2.large.9cm",
					"spot_strategy":      "NoSpot",
					"host_name":          "testhostNameran",
					"period_unit":        "Month",
					"password":           "@MY7yxqc9YXQC",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran资源用名称050801",
						"instance_charge_type": "PostPaid",
						"auto_renew":           "false",
						"amount":               "1",
						"vswitch_id":           CHECKSET,
						"dry_run":              "false",
						"deletion_protection":  "true",
						"auto_pay":             "true",
						"security_group_ids.#": "1",
						"instance_name":        "namecreate",
						"data_disk.#":          "1",
						"create_mode":          CHECKSET,
						"enable_jumbo_frame":   "true",
						"instance_type":        "mysql.x2.large.9cm",
						"spot_strategy":        "NoSpot",
						"host_name":            "testhostNameran",
						"period_unit":          "Month",
						"password":             "@MY7yxqc9YXQC",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "修改描述",
					"force_stop":  "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "修改描述",
						"force_stop":  "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deletion_protection": "false",
					"instance_name":       "update名称",
					"enable_jumbo_frame":  "false",
					"host_name":           "hostnameupdate",
					"password":            "@MY7yxqc9YXQCUPDATE",
					"reboot":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deletion_protection": "false",
						"instance_name":       "update名称",
						"enable_jumbo_frame":  "false",
						"host_name":           "hostnameupdate",
						"password":            "@MY7yxqc9YXQCUPDATE",
						"reboot":              "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap12788 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence12788(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpcRC_clone_1_clone_0" {
  is_default = false
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_vswitch" "vswRC_clone_1_clone_0" {
  is_default   = false
  vpc_id       = alicloud_vpc.vpcRC_clone_1_clone_0.id
  zone_id      = "cn-beijing-i"
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}

resource "alicloud_security_group" "secRC_clone_1_clone_0" {
  description         = "sg"
  security_group_name = "ran_sg"
  vpc_id              = alicloud_vpc.vpcRC_clone_1_clone_0.id
}


`, name)
}

// Case resourceCase_20260415_01_clone_1_clone_0_clone_0 12789
func TestAccAliCloudRdsCustom_basic12789(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap12789)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence12789)
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
					"description":          "ran资源用名称050801",
					"instance_charge_type": "PostPaid",
					"auto_renew":           "false",
					"amount":               "1",
					"vswitch_id":           "${alicloud_vswitch.vswRC_clone_1_clone_0_clone_0.id}",
					"dry_run":              "false",
					"deletion_protection":  "true",
					"auto_pay":             "true",
					"security_group_ids": []string{
						"${alicloud_security_group.secRC_clone_1_clone_0_clone_0.id}"},
					"system_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "40",
							"performance_level": "PL1",
						},
					},
					"instance_name": "namecreate",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"performance_level": "PL0",
							"size":              "40",
						},
					},
					"create_mode":        "0",
					"enable_jumbo_frame": "true",
					"instance_type":      "mysql.x2.large.9cm",
					"spot_strategy":      "NoSpot",
					"host_name":          "testhostNameran",
					"period_unit":        "Month",
					"password":           "@MY7yxqc9YXQC",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran资源用名称050801",
						"instance_charge_type": "PostPaid",
						"auto_renew":           "false",
						"amount":               "1",
						"vswitch_id":           CHECKSET,
						"dry_run":              "false",
						"deletion_protection":  "true",
						"auto_pay":             "true",
						"security_group_ids.#": "1",
						"instance_name":        "namecreate",
						"data_disk.#":          "1",
						"create_mode":          CHECKSET,
						"enable_jumbo_frame":   "true",
						"instance_type":        "mysql.x2.large.9cm",
						"spot_strategy":        "NoSpot",
						"host_name":            "testhostNameran",
						"period_unit":          "Month",
						"password":             "@MY7yxqc9YXQC",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "修改描述",
					"force_stop":  "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "修改描述",
						"force_stop":  "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deletion_protection": "false",
					"instance_name":       "update名称",
					"enable_jumbo_frame":  "false",
					"host_name":           "hostnameupdate",
					"password":            "@MY7yxqc9YXQCUPDATE",
					"reboot":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deletion_protection": "false",
						"instance_name":       "update名称",
						"enable_jumbo_frame":  "false",
						"host_name":           "hostnameupdate",
						"password":            "@MY7yxqc9YXQCUPDATE",
						"reboot":              "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap12789 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence12789(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpcRC_clone_1_clone_0_clone_0" {
  is_default = false
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_security_group" "secRC_clone_1_clone_0_clone_0" {
  description         = "sg"
  security_group_name = "ran_sg"
  vpc_id              = alicloud_vpc.vpcRC_clone_1_clone_0_clone_0.id
}

resource "alicloud_vswitch" "vswRC_clone_1_clone_0_clone_0" {
  is_default   = false
  vpc_id       = alicloud_vpc.vpcRC_clone_1_clone_0_clone_0.id
  zone_id      = "cn-beijing-i"
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}


`, name)
}

// Case resourceCase_20260415_01_clone_1 12771
func TestAccAliCloudRdsCustom_basic12771(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap12771)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence12771)
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
					"description":          "ran资源用名称042301",
					"private_ip_address":   "172.16.0.114",
					"instance_charge_type": "PostPaid",
					"auto_renew":           "false",
					"amount":               "1",
					"vswitch_id":           "${alicloud_vswitch.vswRC_clone_1.id}",
					"dry_run":              "false",
					"deletion_protection":  "true",
					"auto_pay":             "true",
					"security_group_ids": []string{
						"${alicloud_security_group.secRC_clone_1.id}"},
					"system_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "40",
							"performance_level": "PL1",
						},
					},
					"instance_name": "namecreate",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"performance_level": "PL0",
							"size":              "40",
						},
					},
					"create_mode":        "0",
					"enable_jumbo_frame": "true",
					"instance_type":      "mysql.x2.large.9cm",
					"spot_strategy":      "NoSpot",
					"host_name":          "testhostNameran",
					"period_unit":        "Month",
					"password":           "@MY7yxqc9YXQC",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran资源用名称042301",
						"private_ip_address":   "172.16.0.114",
						"instance_charge_type": "PostPaid",
						"auto_renew":           "false",
						"amount":               "1",
						"vswitch_id":           CHECKSET,
						"dry_run":              "false",
						"deletion_protection":  "true",
						"auto_pay":             "true",
						"security_group_ids.#": "1",
						"instance_name":        "namecreate",
						"data_disk.#":          "1",
						"create_mode":          CHECKSET,
						"enable_jumbo_frame":   "true",
						"instance_type":        "mysql.x2.large.9cm",
						"spot_strategy":        "NoSpot",
						"host_name":            "testhostNameran",
						"period_unit":          "Month",
						"password":             "@MY7yxqc9YXQC",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "修改描述",
					"force_stop":  "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "修改描述",
						"force_stop":  "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"deletion_protection": "false",
					"instance_name":       "update名称",
					"enable_jumbo_frame":  "false",
					"host_name":           "hostnameupdate",
					"password":            "@MY7yxqc9YXQCUPDATE",
					"reboot":              "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"deletion_protection": "false",
						"instance_name":       "update名称",
						"enable_jumbo_frame":  "false",
						"host_name":           "hostnameupdate",
						"password":            "@MY7yxqc9YXQCUPDATE",
						"reboot":              "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap12771 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence12771(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpcRC_clone_1" {
  is_default = false
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_security_group" "secRC_clone_1" {
  description         = "sg"
  security_group_name = "ran_sg"
  vpc_id              = alicloud_vpc.vpcRC_clone_1.id
}

resource "alicloud_vswitch" "vswRC_clone_1" {
  is_default   = false
  vpc_id       = alicloud_vpc.vpcRC_clone_1.id
  zone_id      = "cn-beijing-i"
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}


`, name)
}

// Case resourceCase_20260415_01 12772
func TestAccAliCloudRdsCustom_basic12772(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap12772)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence12772)
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
					"description":          "ran资源用例描述042001",
					"private_ip_address":   "172.16.0.113",
					"instance_charge_type": "PostPaid",
					"auto_renew":           "false",
					"amount":               "1",
					"vswitch_id":           "${alicloud_vswitch.vswRC.id}",
					"dry_run":              "false",
					"deletion_protection":  "false",
					"auto_pay":             "true",
					"security_group_ids": []string{
						"${alicloud_security_group.secRC.id}"},
					"system_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "40",
							"performance_level": "PL1",
						},
					},
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"performance_level": "PL0",
							"size":              "40",
						},
					},
					"create_mode":   "0",
					"instance_type": "mysql.x2.large.9cm",
					"spot_strategy": "NoSpot",
					"host_name":     "testhostNameran",
					"period_unit":   "Month",
					"password":      "@MY7yxqc9YXQC",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":          "ran资源用例描述042001",
						"private_ip_address":   "172.16.0.113",
						"instance_charge_type": "PostPaid",
						"auto_renew":           "false",
						"amount":               "1",
						"vswitch_id":           CHECKSET,
						"dry_run":              "false",
						"deletion_protection":  "false",
						"auto_pay":             "true",
						"security_group_ids.#": "1",
						"data_disk.#":          "1",
						"create_mode":          CHECKSET,
						"instance_type":        "mysql.x2.large.9cm",
						"spot_strategy":        "NoSpot",
						"host_name":            "testhostNameran",
						"period_unit":          "Month",
						"password":             "@MY7yxqc9YXQC",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "修改描述",
					"force_stop":  "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "修改描述",
						"force_stop":  "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"password":      "@MY7yxqc9YXQCUPDATE",
					"instance_name": "update名称",
					"reboot":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"password":      "@MY7yxqc9YXQCUPDATE",
						"instance_name": "update名称",
						"reboot":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap12772 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence12772(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "vpcRC" {
  is_default = false
  cidr_block = "172.16.0.0/16"
  vpc_name   = "ran_vpc1"
}

resource "alicloud_security_group" "secRC" {
  description         = "sg"
  security_group_name = "ran_sg"
  vpc_id              = alicloud_vpc.vpcRC.id
}

resource "alicloud_vswitch" "vswRC" {
  is_default   = false
  vpc_id       = alicloud_vpc.vpcRC.id
  zone_id      = "cn-beijing-i"
  cidr_block   = "172.16.0.0/23"
  vswitch_name = "rante_vsw1"
}


`, name)
}

// Case rdscustom_tmk_11_26_2 11916
func TestAccAliCloudRdsCustom_basic11916(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap11916)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence11916)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
							"encrypted":         "false",
							"device":            "/dev/xvdb",
							"snapshot_id":       "${alicloud_rds_custom_snapshot.RCSnapshot.id}",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_custom_tocreateDisk",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":              "40",
							"category":          "cloud_essd",
							"performance_level": "PL1",
						},
					},
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":           "1778575459",
					"create_mode":         "0",
					"spot_strategy":       "NoSpot",
					"support_case":        "eni",
					"user_data":           "IyEvYmluL3NoCmVjaG8gIkhlbGxvIFdvcmxkLiBUaGUgdGltZSBpcyBub3cgJChkYXRlIC1SKSEiIHwgdGVlIC9yb290L3VzZXJkYXRhX3Rlc3QudHh0",
					"dry_run":             "false",
					"user_data_in_base64": "true",
					"deletion_protection": "false",
					"promotion_code":      "1",
					"auto_use_coupon":     "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_custom_tocreateDisk",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"support_case":                  "eni",
						"user_data":                     "IyEvYmluL3NoCmVjaG8gIkhlbGxvIFdvcmxkLiBUaGUgdGltZSBpcyBub3cgJChkYXRlIC1SKSEiIHwgdGVlIC9yb290L3VzZXJkYXRhX3Rlc3QudHh0",
						"dry_run":                       "false",
						"user_data_in_base64":           "true",
						"deletion_protection":           "false",
						"promotion_code":                CHECKSET,
						"auto_use_coupon":               "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":        "mysql.x2.xlarge.7cm",
					"status":               "Stopped",
					"resource_group_id":    "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"auto_use_coupon":      "true",
					"force_stop":           "false",
					"direction":            "Up",
					"reboot_when_finished": "true",
					"reboot_time":          "2026-05-12CST1616:050528800",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":        "mysql.x2.xlarge.7cm",
						"status":               "Stopped",
						"resource_group_id":    CHECKSET,
						"auto_use_coupon":      "true",
						"force_stop":           "false",
						"direction":            "Up",
						"reboot_when_finished": "true",
						"reboot_time":          CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap11916 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence11916(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  description         = "custom用例"
  security_group_name = "test_r_sg"
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_type = "normal"
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}

resource "alicloud_rds_custom" "RCInstance" {
  description          = "快照依赖实例"
  instance_charge_type = "Prepaid"
  auto_renew           = false
  system_disk {
    category = "cloud_essd"
    size     = "20"
  }
  image_id           = "aliyun_2_1903_x64_20G_alibase_20240628.vhd"
  instance_type      = "mysql.x2.xlarge.6cm"
  host_name          = "1778575411"
  spot_strategy      = "NoSpot"
  password           = "jingyiTEST@123"
  status             = "Running"
  key_pair_name      = alicloud_ecs_key_pair.KeyPairName.id
  io_optimized       = "optimized"
  zone_id            = var.test_zone_id
  amount             = "1"
  vswitch_id         = alicloud_vswitch.vSwitchId.id
  period             = "1"
  auto_pay           = true
  security_group_ids = ["${alicloud_security_group.securityGroupId.id}"]
  data_disk {
    category          = "cloud_essd"
    performance_level = "PL1"
    size              = "50"
  }
  internet_max_bandwidth_out    = "0"
  create_mode                   = "0"
  security_enhancement_strategy = "Active"
  period_unit                   = "Month"
}

resource "alicloud_rds_custom_disk" "RCDisk" {
  description          = "包年disk快照用"
  zone_id              = var.test_zone_id
  size                 = "40"
  instance_charge_type = "Prepaid"
  auto_renew           = false
  disk_category        = "cloud_essd"
  period               = "1"
  auto_pay             = true
  disk_name            = "custom_disk_tosnapshot"
  period_unit          = "Month"
}

resource "alicloud_rds_custom_snapshot" "RCSnapshot" {
  description                   = "创建实例使用快照"
  zone_id                       = var.test_zone_id
  retention_days                = "3"
  instant_access_retention_days = "1"
  disk_id                       = alicloud_rds_custom_disk.RCDisk.id
}


`, name)
}

// Case rdscustom_0925 11570
func TestAccAliCloudRdsCustom_basic11570(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap11570)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence11570)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_custom_tocreateDisk",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575460",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
					"support_case":      "eni",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_custom_tocreateDisk",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"support_case":                  "eni",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x2.xlarge.7cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x2.xlarge.7cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
						"dry_run":           "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap11570 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence11570(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  description         = "custom用例"
  security_group_name = "test_r_sg"
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_type = "normal"
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_run_custom_20250905 11416
func TestAccAliCloudRdsCustom_basic11416(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap11416)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence11416)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_custom_tocreateDisk",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575461",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
					"support_case":      "eni",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_custom_tocreateDisk",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"support_case":                  "eni",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x2.xlarge.7cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x2.xlarge.7cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
						"dry_run":           "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap11416 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence11416(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  description         = "custom用例"
  security_group_name = "test_r_sg"
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_type = "normal"
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_run_custom_testB-ran 11412
func TestAccAliCloudRdsCustom_basic11412(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap11412)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence11412)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_custom_tocreateDisk",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575462",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
					"support_case":      "eni",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_custom_tocreateDisk",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_3_x64_20G_alibase_20250629.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"support_case":                  "eni",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x2.xlarge.7cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x2.xlarge.7cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
						"dry_run":           "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap11412 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence11412(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  description         = "custom用例"
  security_group_name = "test_r_sg"
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_type = "normal"
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_run_custom_pro_four不带创建者标签 10885
func TestAccAliCloudRdsCustom_basic10885(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap10885)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence10885)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.xa2.xlarge.7cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_custom_tocreateDisk",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_3_x64_20G_alibase_20250117.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575463",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
					"support_case":      "eni",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.xa2.xlarge.7cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_custom_tocreateDisk",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_3_x64_20G_alibase_20250117.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"support_case":                  "eni",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
						"dry_run":           "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap10885 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence10885(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-chengdu"
}

variable "test_zone_id" {
  default = "cn-chengdu-b"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  description         = "custom用例"
  security_group_name = "test_r_sg"
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_type = "normal"
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_run_ins_extra_param 10836
func TestAccAliCloudRdsCustom_basic10836(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap10836)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence10836)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_test_ram_code",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575464",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_test_ram_code",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x4.xlarge.6cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x4.xlarge.6cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap10836 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence10836(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_ran_ram_code_one-pre 9224
func TestAccAliCloudRdsCustom_basic9224(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap9224)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence9224)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_test_ram_code",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_charge_type":          "PayByTraffic",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575465",
					"create_mode":       "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_test_ram_code",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_charge_type":          "PayByTraffic",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x4.xlarge.6cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x4.xlarge.6cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap9224 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence9224(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_ran_ram_code_three-pre 9664
func TestAccAliCloudRdsCustom_basic9664(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap9664)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence9664)
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
					"amount":        "1",
					"vswitch_id":    "${alicloud_vswitch.vSwitchId.id}",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "ran_test_ram_code",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575466",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"vswitch_id":                    CHECKSET,
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   "ran_test_ram_code",
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x4.xlarge.6cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x4.xlarge.6cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap9664 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence9664(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  cidr_block   = "172.16.5.0/24"
  zone_id      = var.test_zone_id
  vswitch_name = "test_vswitch"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id = alicloud_vpc.vpcId.id
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_ran_rc_createnodepoolapi_three 9893
func TestAccAliCloudRdsCustom_basic9893(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap9893)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence9893)
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
					"amount":        "1",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "${var.description}",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575467",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
					"vswitch_id":        "${alicloud_vswitch.vSwitchId.id}",
					"support_case":      "eni",
					"deployment_set_id": "${var.deploymentsetid}",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   CHECKSET,
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"vswitch_id":                    CHECKSET,
						"support_case":                  "eni",
						"deployment_set_id":             CHECKSET,
						"dry_run":                       "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x4.xlarge.6cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x4.xlarge.6cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_pay":          "false",
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_pay":          "false",
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap9893 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence9893(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cluster_id" {
  default = "c18c40b2b336840e2b2bbf8ab291758e2"
}

variable "deploymentsetid" {
  default = "ds-2ze78ef5kyj9eveue92m"
}

variable "vswtich-id" {
  default = "test_vswitch"
}

variable "vpc_name" {
  default = "beijing111"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "description" {
  default = "ran_1-08_rccreatenodepool_api"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

variable "securitygroup_name" {
  default = "rds_custom_init_sg_cn_beijing"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  vpc_name = var.vpc_name
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  zone_id      = var.test_zone_id
  vswitch_name = var.vswtich-id
  cidr_block   = "172.16.5.0/24"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_name = var.securitygroup_name
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Case rdscustom_ran_rc_createnodepoolapi_one 9895
func TestAccAliCloudRdsCustom_basic9895(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_rds_custom.default"
	ra := resourceAttrInit(resourceId, AlicloudRdsCustomMap9895)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &RdsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeRdsCustom")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccrds%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudRdsCustomBasicDependence9895)
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
					"amount":        "1",
					"auto_renew":    "false",
					"period":        "1",
					"auto_pay":      "true",
					"instance_type": "mysql.x2.xlarge.6cm",
					"data_disk": []map[string]interface{}{
						{
							"category":          "cloud_essd",
							"size":              "50",
							"performance_level": "PL1",
						},
					},
					"status": "Running",
					"security_group_ids": []string{
						"${alicloud_security_group.securityGroupId.id}"},
					"io_optimized":                  "optimized",
					"description":                   "${var.description}",
					"key_pair_name":                 "${alicloud_ecs_key_pair.KeyPairName.id}",
					"zone_id":                       "${var.test_zone_id}",
					"instance_charge_type":          "Prepaid",
					"internet_max_bandwidth_out":    "0",
					"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
					"security_enhancement_strategy": "Active",
					"period_unit":                   "Month",
					"password":                      "jingyiTEST@123",
					"system_disk": []map[string]interface{}{
						{
							"size":     "40",
							"category": "cloud_essd",
						},
					},
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"host_name":         "1778575468",
					"create_mode":       "0",
					"spot_strategy":     "NoSpot",
					"vswitch_id":        "${alicloud_vswitch.vSwitchId.id}",
					"support_case":      "eni",
					"deployment_set_id": "${var.deploymentsetid}",
					"dry_run":           "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"amount":                        "1",
						"auto_renew":                    "false",
						"period":                        "1",
						"auto_pay":                      "true",
						"instance_type":                 "mysql.x2.xlarge.6cm",
						"data_disk.#":                   "1",
						"status":                        "Running",
						"security_group_ids.#":          "1",
						"io_optimized":                  "optimized",
						"description":                   CHECKSET,
						"key_pair_name":                 CHECKSET,
						"zone_id":                       CHECKSET,
						"instance_charge_type":          "Prepaid",
						"internet_max_bandwidth_out":    "0",
						"image_id":                      "aliyun_2_1903_x64_20G_alibase_20240628.vhd",
						"security_enhancement_strategy": "Active",
						"period_unit":                   "Month",
						"password":                      "jingyiTEST@123",
						"resource_group_id":             CHECKSET,
						"host_name":                     CHECKSET,
						"create_mode":                   CHECKSET,
						"spot_strategy":                 "NoSpot",
						"vswitch_id":                    CHECKSET,
						"support_case":                  "eni",
						"deployment_set_id":             CHECKSET,
						"dry_run":                       "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_type":     "mysql.x4.xlarge.6cm",
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "false",
					"direction":         "Up",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_type":     "mysql.x4.xlarge.6cm",
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "false",
						"direction":         "Up",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"auto_pay":          "false",
					"status":            "Running",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"auto_pay":          "false",
						"status":            "Running",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Stopped",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"force_stop":        "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status":            "Stopped",
						"resource_group_id": CHECKSET,
						"force_stop":        "true",
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
				ImportStateVerifyIgnore: []string{"amount", "auto_pay", "auto_renew", "auto_use_coupon", "create_mode", "direction", "dry_run", "force_stop", "host_name", "image_id", "instance_charge_type", "internet_charge_type", "internet_max_bandwidth_out", "io_optimized", "key_pair_name", "password", "period", "period_unit", "promotion_code", "reboot", "reboot_time", "reboot_when_finished", "security_enhancement_strategy", "spot_strategy", "support_case", "user_data_in_base64"},
			},
		},
	})
}

var AlicloudRdsCustomMap9895 = map[string]string{
	"region_id": CHECKSET,
}

func AlicloudRdsCustomBasicDependence9895(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "cluster_id" {
  default = "c18c40b2b336840e2b2bbf8ab291758e2"
}

variable "deploymentsetid" {
  default = "ds-2ze78ef5kyj9eveue92m"
}

variable "vswtich-id" {
  default = "test_vswitch"
}

variable "vpc_name" {
  default = "beijing111"
}

variable "test_region_id" {
  default = "cn-beijing"
}

variable "description" {
  default = "ran_1-08_rccreatenodepool_api"
}

variable "test_zone_id" {
  default = "cn-beijing-h"
}

variable "securitygroup_name" {
  default = "rds_custom_init_sg_cn_beijing"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_vpc" "vpcId" {
  vpc_name = var.vpc_name
}

resource "alicloud_vswitch" "vSwitchId" {
  vpc_id       = alicloud_vpc.vpcId.id
  zone_id      = var.test_zone_id
  vswitch_name = var.vswtich-id
  cidr_block   = "172.16.5.0/24"
}

resource "alicloud_security_group" "securityGroupId" {
  vpc_id              = alicloud_vpc.vpcId.id
  security_group_name = var.securitygroup_name
}

resource "alicloud_ecs_deployment_set" "deploymentSet" {
}

resource "alicloud_ecs_key_pair" "KeyPairName" {
  key_pair_name = alicloud_vswitch.vSwitchId.id
}


`, name)
}

// Test Rds Custom. <<< Resource test cases, automatically generated.
