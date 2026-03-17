package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Nas AccessPoint. >>> Resource test cases, automatically generated.
// Case resource_AccessPoint_test_2 12389
func TestAccAliCloudNasAccessPoint_basic12389(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_access_point.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAccessPointMap12389)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAccessPoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(1, 999)
	name := fmt.Sprintf("tfacc%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAccessPointBasicDependence12389)
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
					"access_point_name": name,
					"vpc_id":            "${alicloud_vpc.createVpc_11.id}",
					"access_group":      "${alicloud_nas_access_group.resource_AccessGroup_test_8.access_group_name}",
					"root_path_permission": []map[string]interface{}{
						{
							"owner_user_id":  "1",
							"permission":     "0755",
							"owner_group_id": "1",
						},
					},
					"vswitch_id":     "${alicloud_vswitch.CreateVswitch_7.id}",
					"file_system_id": "${alicloud_nas_file_system.resource_FileSystem_test_5.id}",
					"enabled_ram":    "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_point_name": name,
						"vpc_id":            CHECKSET,
						"access_group":      CHECKSET,
						"vswitch_id":        CHECKSET,
						"file_system_id":    CHECKSET,
						"enabled_ram":       "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_point_name": name + "_update",
					"enabled_ram":       "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_point_name": name + "_update",
						"enabled_ram":       "true",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNasAccessPointMap12389 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"access_point_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudNasAccessPointBasicDependence12389(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_10" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_vpc" "createVpc_9" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_vswitch" "CreateVswitch_8" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_10.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_kms_key" "resource_Key_test_4" {
}

resource "alicloud_vswitch" "CreateVswitch_7" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_9.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_nas_access_group" "resource_AccessGroup_test_8" {
  access_group_type = "Vpc"
  access_group_name = "testgroup"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "resource_FileSystem_test_5" {
  storage_type     = "Performance"
  encrypt_type     = "2"
  kms_key_id       = alicloud_kms_key.resource_Key_test_4.id
  capacity         = "200"
  protocol_type    = "NFS"
  file_system_type = "standard"
  vswitch_id       = alicloud_vswitch.CreateVswitch_8.id
}

resource "alicloud_vpc" "createVpc_11" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}


`, name)
}

// Case 通用型接入点_Terraform发布 10938
func TestAccAliCloudNasAccessPoint_basic10938(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_access_point.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAccessPointMap10938)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAccessPoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAccessPointBasicDependence10938)
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
					"access_point_name": name,
					"root_path":         "/",
					"vpc_id":            "${alicloud_vpc.defaultkyVC70.id}",
					"access_group":      "${alicloud_nas_access_group.defaultBbc7ev.access_group_name}",
					"root_path_permission": []map[string]interface{}{
						{
							"owner_group_id": "1",
							"owner_user_id":  "1",
							"permission":     "0755",
						},
					},
					"vswitch_id":     "${alicloud_vswitch.defaultoZAPmO.id}",
					"file_system_id": "${alicloud_nas_file_system.defaultVtUpDh.id}",
					"posix_user": []map[string]interface{}{
						{
							"posix_group_id":            "123",
							"posix_user_id":             "123",
							"posix_secondary_group_ids": []string{},
						},
					},
					"enabled_ram": "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_point_name": name,
						"root_path":         "/",
						"vpc_id":            CHECKSET,
						"access_group":      CHECKSET,
						"vswitch_id":        CHECKSET,
						"file_system_id":    CHECKSET,
						"enabled_ram":       "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_point_name": name + "_update",
					"access_group":      "${alicloud_nas_access_group.default6mnIjY.access_group_name}",
					"enabled_ram":       "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_point_name": name + "_update",
						"access_group":      CHECKSET,
						"enabled_ram":       "true",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNasAccessPointMap10938 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"access_point_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudNasAccessPointBasicDependence10938(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "azone" {
  default = "cn-hangzhou-g"
}

resource "alicloud_vpc" "defaultkyVC70" {
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultoZAPmO" {
  vpc_id     = alicloud_vpc.defaultkyVC70.id
  zone_id    = var.azone
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_nas_access_group" "defaultBbc7ev" {
  access_group_type = "Vpc"
  access_group_name = "AccessPointTest"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "defaultVtUpDh" {
  storage_type     = "Performance"
  zone_id          = var.azone
  encrypt_type     = "0"
  protocol_type    = "NFS"
  file_system_type = "standard"
}

resource "alicloud_nas_access_group" "default6mnIjY" {
  access_group_type = "Vpc"
  access_group_name = "AccessPointChangeTest"
  file_system_type  = "standard"
}


`, name)
}

// Case 通用型接入点 4833
func TestAccAliCloudNasAccessPoint_basic4833(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_access_point.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAccessPointMap4833)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAccessPoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAccessPointBasicDependence4833)
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
					"vpc_id":            "${alicloud_vpc.defaultkyVC70.id}",
					"access_group":      "${alicloud_nas_access_group.defaultBbc7ev.access_group_name}",
					"vswitch_id":        "${alicloud_vswitch.defaultoZAPmO.id}",
					"enabled_ram":       "false",
					"file_system_id":    "${alicloud_nas_file_system.defaultVtUpDh.id}",
					"access_point_name": name,
					"root_path":         "/",
					"posix_user": []map[string]interface{}{
						{
							"posix_group_id": "123",
							"posix_user_id":  "123",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":            CHECKSET,
						"access_group":      CHECKSET,
						"vswitch_id":        CHECKSET,
						"enabled_ram":       "false",
						"file_system_id":    CHECKSET,
						"access_point_name": name,
						"root_path":         "/",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_group":      "${alicloud_nas_access_group.default6mnIjY.access_group_name}",
					"enabled_ram":       "true",
					"access_point_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_group":      CHECKSET,
						"enabled_ram":       "true",
						"access_point_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNasAccessPointMap4833 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"access_point_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudNasAccessPointBasicDependence4833(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "azone" {
  default = "cn-hangzhou-g"
}

resource "alicloud_vpc" "defaultkyVC70" {
  cidr_block  = "172.16.0.0/12"
  description = "接入点测试"
}

resource "alicloud_vswitch" "defaultoZAPmO" {
  vpc_id     = alicloud_vpc.defaultkyVC70.id
  zone_id    = var.azone
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_nas_access_group" "defaultBbc7ev" {
  access_group_type = "Vpc"
  access_group_name = "StandardAccessPointTest"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "defaultVtUpDh" {
  storage_type     = "Performance"
  zone_id          = var.azone
  encrypt_type     = "0"
  protocol_type    = "NFS"
  file_system_type = "standard"
}

resource "alicloud_nas_access_group" "default6mnIjY" {
  access_group_type = "Vpc"
  access_group_name = "StandardAccessPointTestChange"
  file_system_type  = "standard"
}


`, name)
}

// Case 通用型接入点不传入 RootDirectory 6611
func TestAccAliCloudNasAccessPoint_basic6611(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_access_point.default"
	ra := resourceAttrInit(resourceId, AlicloudNasAccessPointMap6611)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasAccessPoint")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasAccessPointBasicDependence6611)
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
					"vpc_id":            "${alicloud_vpc.defaultkyVC70.id}",
					"access_group":      "${alicloud_nas_access_group.defaultBbc7ev.access_group_name}",
					"vswitch_id":        "${alicloud_vswitch.defaultoZAPmO.id}",
					"enabled_ram":       "false",
					"file_system_id":    "${alicloud_nas_file_system.defaultVtUpDh.id}",
					"access_point_name": name,
					"posix_user": []map[string]interface{}{
						{
							"posix_group_id": "123",
							"posix_user_id":  "123",
						},
					},
					"root_path_permission": []map[string]interface{}{
						{
							"owner_group_id": "1",
							"owner_user_id":  "1",
							"permission":     "0777",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":            CHECKSET,
						"access_group":      CHECKSET,
						"vswitch_id":        CHECKSET,
						"enabled_ram":       "false",
						"file_system_id":    CHECKSET,
						"access_point_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"enabled_ram":       "true",
					"access_point_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"enabled_ram":       "true",
						"access_point_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

var AlicloudNasAccessPointMap6611 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"access_point_id": CHECKSET,
	"region_id":       CHECKSET,
}

func AlicloudNasAccessPointBasicDependence6611(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "region_id" {
  default = "cn-hangzhou"
}

variable "azone" {
  default = "cn-hangzhou-g"
}

resource "alicloud_vpc" "defaultkyVC70" {
  cidr_block  = "172.16.0.0/12"
  description = "接入点测试noRootDirectory"
}

resource "alicloud_vswitch" "defaultoZAPmO" {
  vpc_id     = alicloud_vpc.defaultkyVC70.id
  zone_id    = var.azone
  cidr_block = "172.16.0.0/24"
}

resource "alicloud_nas_access_group" "defaultBbc7ev" {
  access_group_type = "Vpc"
  access_group_name = "StandardAccessPointTestnoRootDirectory"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "defaultVtUpDh" {
  storage_type     = "Performance"
  zone_id          = var.azone
  encrypt_type     = "0"
  protocol_type    = "NFS"
  file_system_type = "standard"
  description      = "AccessPointnoRootDirectory"
}


`, name)
}

// Test Nas AccessPoint. <<< Resource test cases, automatically generated.
