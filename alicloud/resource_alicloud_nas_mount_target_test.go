package alicloud

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/stretchr/testify/assert"
)

// Test Nas MountTarget. >>> Resource test cases, automatically generated.
// Case resource_MountTarget_test 12383
func TestAccAliCloudNasMountTarget_basic12383(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AlicloudNasMountTargetMap12383)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasMountTargetBasicDependence12383)
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
					"vpc_id":            "${alicloud_vpc.createVpc_7.id}",
					"network_type":      "Vpc",
					"access_group_name": "${alicloud_nas_access_group.resource_AccessGroup_test_6.access_group_name}",
					"vswitch_id":        "${alicloud_vswitch.CreateVswitch_10.id}",
					"file_system_id":    "${alicloud_nas_file_system.resource_FileSystem_test_9.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":            CHECKSET,
						"network_type":      "Vpc",
						"access_group_name": CHECKSET,
						"vswitch_id":        CHECKSET,
						"file_system_id":    CHECKSET,
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
				ImportStateVerifyIgnore: []string{"dual_stack", "security_group_id"},
			},
		},
	})
}

var AlicloudNasMountTargetMap12383 = map[string]string{
	"mount_target_domain": CHECKSET,
}

func AlicloudNasMountTargetBasicDependence12383(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_8" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_kms_key" "resource_Key_test_7" {
}

resource "alicloud_vswitch" "CreateVswitch_9" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_8.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_vpc" "createVpc_12" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_nas_file_system" "resource_FileSystem_test_9" {
  storage_type     = "Performance"
  encrypt_type     = "0"
  kms_key_id       = alicloud_kms_key.resource_Key_test_7.id
  capacity         = "500"
  protocol_type    = "SMB"
  file_system_type = "standard"
  vswitch_id       = alicloud_vswitch.CreateVswitch_9.id
}

resource "alicloud_vpc" "createVpc_7" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_vswitch" "CreateVswitch_10" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_12.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_nas_access_group" "resource_AccessGroup_test_6" {
  access_group_type = "Vpc"
  access_group_name = "testgroup"
}


`, name)
}

// Case resource_MountTarget_test_2 12386
func TestAccAliCloudNasMountTarget_basic12386(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AlicloudNasMountTargetMap12386)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasMountTargetBasicDependence12386)
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
					"vpc_id":            "${alicloud_vpc.createVpc_16.id}",
					"network_type":      "Vpc",
					"access_group_name": "${alicloud_nas_access_group.resource_AccessGroup_test_9.access_group_name}",
					"vswitch_id":        "${alicloud_vswitch.CreateVswitch_14.id}",
					"file_system_id":    "${alicloud_nas_file_system.resource_FileSystem_test_11.id}",
					"dual_stack":        "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":            CHECKSET,
						"network_type":      "Vpc",
						"access_group_name": CHECKSET,
						"vswitch_id":        CHECKSET,
						"file_system_id":    CHECKSET,
						"dual_stack":        "false",
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
				ImportStateVerifyIgnore: []string{"dual_stack", "security_group_id"},
			},
		},
	})
}

var AlicloudNasMountTargetMap12386 = map[string]string{
	"mount_target_domain": CHECKSET,
}

func AlicloudNasMountTargetBasicDependence12386(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_17" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_vswitch" "CreateVswitch_13" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_17.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_kms_key" "resource_Key_test_9" {
}

resource "alicloud_vpc" "createVpc_18" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_nas_access_group" "resource_AccessGroup_test_9" {
  access_group_type = "Vpc"
  access_group_name = "testgroup"
}

resource "alicloud_nas_file_system" "resource_FileSystem_test_11" {
  storage_type     = "Performance"
  encrypt_type     = "0"
  kms_key_id       = alicloud_kms_key.resource_Key_test_9.id
  capacity         = "500"
  protocol_type    = "SMB"
  file_system_type = "standard"
  vswitch_id       = alicloud_vswitch.CreateVswitch_13.id
}

resource "alicloud_vpc" "createVpc_16" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_vswitch" "CreateVswitch_14" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_18.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}


`, name)
}

// Case resource_MountTarget_test_1 12388
func TestAccAliCloudNasMountTarget_basic12388(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AlicloudNasMountTargetMap12388)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasMountTargetBasicDependence12388)
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
					"vpc_id":            "${alicloud_vpc.createVpc_13.id}",
					"network_type":      "VPC",
					"access_group_name": "${alicloud_nas_access_group.resource_AccessGroup_test_7.access_group_name}",
					"vswitch_id":        "${alicloud_vswitch.CreateVswitch_12.id}",
					"file_system_id":    "${alicloud_nas_file_system.resource_FileSystem_test_10.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":            CHECKSET,
						"network_type":      "VPC",
						"access_group_name": CHECKSET,
						"vswitch_id":        CHECKSET,
						"file_system_id":    CHECKSET,
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
				ImportStateVerifyIgnore: []string{"dual_stack", "security_group_id"},
			},
		},
	})
}

var AlicloudNasMountTargetMap12388 = map[string]string{
	"mount_target_domain": CHECKSET,
}

func AlicloudNasMountTargetBasicDependence12388(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_vpc" "createVpc_14" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_vswitch" "CreateVswitch_11" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_14.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_kms_key" "resource_Key_test_8" {
}

resource "alicloud_vpc" "createVpc_15" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}

resource "alicloud_nas_file_system" "resource_FileSystem_test_10" {
  storage_type     = "Performance"
  encrypt_type     = "0"
  kms_key_id       = alicloud_kms_key.resource_Key_test_8.id
  capacity         = "500"
  protocol_type    = "SMB"
  file_system_type = "standard"
  vswitch_id       = alicloud_vswitch.CreateVswitch_11.id
}

resource "alicloud_vswitch" "CreateVswitch_12" {
  is_default   = false
  vpc_id       = alicloud_vpc.createVpc_15.id
  zone_id      = "cn-hangzhou-b"
  cidr_block   = "10.0.0.0/24"
  vswitch_name = "pl-test-vswpl"
}

resource "alicloud_nas_access_group" "resource_AccessGroup_test_7" {
  access_group_type = "Vpc"
  access_group_name = "testgroup"
}

resource "alicloud_vpc" "createVpc_13" {
  is_default = false
  cidr_block = "10.0.0.0/8"
  vpc_name   = "pl-example-vpc"
}


`, name)
}

// Case 通用型挂载点_副本 7198
func TestAccAliCloudNasMountTarget_basic7198(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AlicloudNasMountTargetMap7198)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccnas%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasMountTargetBasicDependence7198)
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
					"vpc_id":         "${alicloud_vpc.defaultiUcGvP.id}",
					"network_type":   "Vpc",
					"vswitch_id":     "${alicloud_vswitch.defaultulTodx.id}",
					"file_system_id": "${alicloud_nas_file_system.defaultHtxJoZ.id}",
					"status":         "Active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"vpc_id":         CHECKSET,
						"network_type":   "Vpc",
						"vswitch_id":     CHECKSET,
						"file_system_id": CHECKSET,
						"status":         "Active",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Inactive",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Inactive",
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
				ImportStateVerifyIgnore: []string{"dual_stack", "security_group_id"},
			},
		},
	})
}

var AlicloudNasMountTargetMap7198 = map[string]string{
	"mount_target_domain": CHECKSET,
}

func AlicloudNasMountTargetBasicDependence7198(name string) string {
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

resource "alicloud_vpc" "defaultiUcGvP" {
  description = "创建挂载点"
  cidr_block  = "172.16.0.0/12"
}

resource "alicloud_vswitch" "defaultulTodx" {
  description = "挂载点创建测试"
  vpc_id      = alicloud_vpc.defaultiUcGvP.id
  cidr_block  = "172.16.0.0/24"
  zone_id     = var.azone
}

resource "alicloud_nas_access_group" "defaultR4schK" {
  access_group_type = "Vpc"
  description       = "挂载点创建测试"
  access_group_name = "StandardMountTarget"
  file_system_type  = "standard"
}

resource "alicloud_nas_file_system" "defaultHtxJoZ" {
  description      = "挂载点资源测试"
  storage_type     = "Performance"
  encrypt_type     = "0"
  protocol_type    = "NFS"
  file_system_type = "standard"
  zone_id          = var.azone
}

resource "alicloud_nas_access_group" "defaultjWPUuI" {
  access_group_type = "Vpc"
  description       = "挂载点权限组修改"
  access_group_name = "StandardAccessGroupChange"
  file_system_type  = "standard"
}


`, name)
}

// Test Nas MountTarget. <<< Resource test cases, automatically generated.

func TestAccAliCloudNASMountTarget_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AlicloudNasMountTarget0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sAlicloudNasMountTarget%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasMountTargetBasicDependence0)
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
					"access_group_name": "${alicloud_nas_access_group.example.access_group_name}",
					"file_system_id":    "${alicloud_nas_file_system.example.id}",
					"vswitch_id":        "${alicloud_vswitch.main.id}",
					"security_group_id": "${alicloud_security_group.example.id}",
					"status":            "Active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_group_name": name,
						"file_system_id":    CHECKSET,
						"vswitch_id":        CHECKSET,
						"security_group_id": CHECKSET,
						"status":            "Active",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"security_group_id"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_group_name": "${alicloud_nas_access_group.example1.access_group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_group_name": name + "change",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Inactive",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"status": "Inactive",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_group_name": "${alicloud_nas_access_group.example.access_group_name}",
					"file_system_id":    "${alicloud_nas_file_system.example.id}",
					"vswitch_id":        "${alicloud_vswitch.main.id}",
					"security_group_id": "${alicloud_security_group.example.id}",
					"status":            "Active",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_group_name": name,
						"file_system_id":    CHECKSET,
						"vswitch_id":        CHECKSET,
						"security_group_id": CHECKSET,
						"status":            "Active",
					}),
				),
			},
		},
	})
}

func TestAccAliCloudNasExtremeMountTarget_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AlicloudNasMountTarget1)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sAlicloudNasMountTarget%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudNasMountTargetBasicDependence1)
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
					"access_group_name": "${alicloud_nas_access_group.example.access_group_name}",
					"file_system_id":    "${alicloud_nas_file_system.example.id}",
					"vswitch_id":        "${alicloud_vswitch.main.id}",
					"vpc_id":            "${alicloud_vpc.main.id}",
					"network_type":      "${alicloud_nas_access_group.example.access_group_type}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_group_name":   name,
						"file_system_id":      CHECKSET,
						"vswitch_id":          CHECKSET,
						"vpc_id":              CHECKSET,
						"mount_target_domain": CHECKSET,
						"status":              CHECKSET,
						"network_type":        "Vpc",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_group_name": "${alicloud_nas_access_group.example1.access_group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"access_group_name": name + "change",
					}),
				),
			},
		},
	})
}

var AlicloudNasMountTarget0 = map[string]string{}

func AlicloudNasMountTargetBasicDependence0(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}

variable "name1" {
	default = "%schange"
}

data "alicloud_nas_protocols" "example" {
	type = "Performance"
}

data "alicloud_nas_zones" "default" {
	file_system_type = "standard"
}
	
locals {
	count_size = length(data.alicloud_nas_zones.default.zones)
	zone_id    = data.alicloud_nas_zones.default.zones[local.count_size - 1].zone_id
}


resource "alicloud_vpc" "main" {
	vpc_name   = "terraform-example"
	cidr_block = "172.17.3.0/24"
}

resource "alicloud_vswitch" "main" {
  	vswitch_name = alicloud_vpc.main.vpc_name
  	cidr_block   = alicloud_vpc.main.cidr_block
  	vpc_id       = alicloud_vpc.main.id
  	zone_id      = local.zone_id
}

resource "alicloud_security_group" "example" {
	name = var.name
	vpc_id = "${alicloud_vpc.main.id}"
}

resource "alicloud_nas_file_system" "example" {
	protocol_type = "${data.alicloud_nas_protocols.example.protocols.0}"
	storage_type = "Performance"
}

resource "alicloud_nas_access_group" "example" {
	access_group_name = "${var.name}"
	access_group_type = "Vpc"
}

resource "alicloud_nas_access_group" "example1" {
	access_group_name = "${var.name1}"
	access_group_type = "Vpc"
}

resource "alicloud_nas_mount_target" "example" {
	file_system_id    = "${alicloud_nas_file_system.example.id}"
	access_group_name = "${alicloud_nas_access_group.example.access_group_name}"
	vswitch_id        = "${alicloud_vswitch.main.id}"
	security_group_id = "${alicloud_security_group.example.id}"
}
`, name, name)
}

var AlicloudNasMountTarget1 = map[string]string{}

func AlicloudNasMountTargetBasicDependence1(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}

variable "name1" {
	default = "%schange"
}

data "alicloud_nas_zones" "default" {
	file_system_type = "extreme"
}

locals {
	count_size = length(data.alicloud_nas_zones.default.zones)
	zone_id    = data.alicloud_nas_zones.default.zones[local.count_size - 1].zone_id
}


resource "alicloud_vpc" "main" {
  vpc_name   = "terraform-example"
  cidr_block = "172.17.3.0/24"
}

resource "alicloud_vswitch" "main" {
  vswitch_name = alicloud_vpc.main.vpc_name
  cidr_block   = alicloud_vpc.main.cidr_block
  vpc_id       = alicloud_vpc.main.id
  zone_id      = local.zone_id
}

resource "alicloud_nas_file_system" "example" {
  file_system_type = "extreme"
  protocol_type    = "NFS"
  zone_id          = data.alicloud_nas_zones.default.zones[local.count_size - 1].zone_id
  storage_type     = "advance"
  capacity         = 100
}

resource "alicloud_nas_access_group" "example" {
	access_group_name = "${var.name}"
	access_group_type = "Vpc"
	file_system_type  = "extreme"
}

resource "alicloud_nas_access_group" "example1" {
	access_group_name = "${var.name1}"
	access_group_type = "Vpc"
	file_system_type  = "extreme"
}
`, name, name)
}

// lintignore: R001
func TestUnitAlicloudNASMountTarget(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	d, _ := schema.InternalMap(p["alicloud_nas_mount_target"].Schema).Data(nil, nil)
	dCreate, _ := schema.InternalMap(p["alicloud_nas_mount_target"].Schema).Data(nil, nil)
	dCreate.MarkNewResource()
	for key, value := range map[string]interface{}{
		"access_group_name": "access_group_name",
		"file_system_id":    "file_system_id",
		"vswitch_id":        "vswitch_id",
		"vpc_id":            "vpc_id",
		"network_type":      "network_type",
	} {
		err := dCreate.Set(key, value)
		assert.Nil(t, err)
		err = d.Set(key, value)
		assert.Nil(t, err)
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		"MountTargets": map[string]interface{}{
			"MountTarget": []interface{}{
				map[string]interface{}{
					"AccessGroup":       "access_group_name",
					"FileSystemId":      "file_system_id",
					"Status":            "Active",
					"VpcId":             "vpc_id",
					"VswId":             "vswitch_id",
					"MountTargetDomain": "MockMountTargetDomain",
					"NetworkType":       "network_type",
				},
			},
		},
	}

	responseMock := map[string]func(errorCode string) (map[string]interface{}, error){
		"RetryError": func(errorCode string) (map[string]interface{}, error) {
			return nil, &tea.SDKError{
				Code:       String(errorCode),
				Data:       String(errorCode),
				Message:    String(errorCode),
				StatusCode: tea.Int(400),
			}
		},
		"NotFoundError": func(errorCode string) (map[string]interface{}, error) {
			return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_nas_mount_target", "MockMountTargetDomain"))
		},
		"NoRetryError": func(errorCode string) (map[string]interface{}, error) {
			return nil, &tea.SDKError{
				Code:       String(errorCode),
				Data:       String(errorCode),
				Message:    String(errorCode),
				StatusCode: tea.Int(400),
			}
		},
		"CreateNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			result["MountTargetDomain"] = "MockMountTargetDomain"
			return result, nil
		},
		"UpdateNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"DeleteNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"ReadNormal": func(errorCode string) (map[string]interface{}, error) {
			result := ReadMockResponse
			return result, nil
		},
		"VpcNormal": func(errorCode string) (map[string]interface{}, error) {
			result := map[string]interface{}{
				"VSwitchId": "VSwitchId",
				"VpcId":     "VpcId",
			}
			return result, nil
		},
	}
	// Create
	t.Run("CreateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNasClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudNasMountTargetCreate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("CreateAbnormal", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		patcheDescribeVSwitchWithTeadsl := gomonkey.ApplyMethod(reflect.TypeOf(&VpcService{}), "DescribeVSwitchWithTeadsl", func(*VpcService, string) (map[string]interface{}, error) {
			return responseMock["VpcNormal"]("")
		})
		err := resourceAliCloudNasMountTargetCreate(d, rawClient)
		patches.Reset()
		patcheDescribeVSwitchWithTeadsl.Reset()
		assert.NotNil(t, err)
	})
	t.Run("CreateNormal", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["CreateNormal"]("")
		})
		patcheDescribeVSwitchWithTeadsl := gomonkey.ApplyMethod(reflect.TypeOf(&VpcService{}), "DescribeVSwitchWithTeadsl", func(*VpcService, string) (map[string]interface{}, error) {
			return responseMock["VpcNormal"]("")
		})
		err := resourceAliCloudNasMountTargetCreate(dCreate, rawClient)
		patches.Reset()
		patcheDescribeVSwitchWithTeadsl.Reset()
		assert.Nil(t, err)
	})

	// Set ID for Update and Delete Method
	d.SetId(fmt.Sprint("file_system_id", ":", "MockMountTargetDomain"))
	// Update
	t.Run("UpdateClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNasClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})

		err := resourceAliCloudNasMountTargetUpdate(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyMountTargetAbnormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"access_group_name", "status"} {
			switch p["alicloud_nas_mount_target"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeInt:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.Itoa(d.Get(key).(int)), New: strconv.Itoa(3)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_nas_mount_target"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudNasMountTargetUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})

	t.Run("UpdateModifyMountTargetNormal", func(t *testing.T) {
		diff := terraform.NewInstanceDiff()
		for _, key := range []string{"access_group_name", "status"} {
			switch p["alicloud_nas_mount_target"].Schema[key].Type {
			case schema.TypeString:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: d.Get(key).(string), New: d.Get(key).(string) + "_update"})
			case schema.TypeBool:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.FormatBool(d.Get(key).(bool)), New: strconv.FormatBool(true)})
			case schema.TypeInt:
				diff.SetAttribute(key, &terraform.ResourceAttrDiff{Old: strconv.Itoa(d.Get(key).(int)), New: strconv.Itoa(3)})
			case schema.TypeMap:
				diff.SetAttribute("tags.%", &terraform.ResourceAttrDiff{Old: "0", New: "2"})
				diff.SetAttribute("tags.For", &terraform.ResourceAttrDiff{Old: "", New: "Test"})
				diff.SetAttribute("tags.Created", &terraform.ResourceAttrDiff{Old: "", New: "TF"})
			}
		}
		resourceData1, _ := schema.InternalMap(p["alicloud_nas_mount_target"].Schema).Data(nil, diff)
		resourceData1.SetId(d.Id())
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudNasMountTargetUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("UpdateMockParseResourceId", func(t *testing.T) {
		resourceData1, _ := schema.InternalMap(p["alicloud_nas_mount_target"].Schema).Data(nil, nil)
		resourceData1.SetId("MockId")
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("RetryError")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["UpdateNormal"]("")
		})
		err := resourceAliCloudNasMountTargetUpdate(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	// Delete
	t.Run("DeleteClientAbnormal", func(t *testing.T) {
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewNasClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
			return nil, &tea.SDKError{
				Code:       String("loadEndpoint error"),
				Data:       String("loadEndpoint error"),
				Message:    String("loadEndpoint error"),
				StatusCode: tea.Int(400),
			}
		})
		err := resourceAliCloudNasMountTargetDelete(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("DeleteMockAbnormal", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := true
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		err := resourceAliCloudNasMountTargetDelete(d, rawClient)
		patches.Reset()
		assert.NotNil(t, err)
	})
	t.Run("DeleteMockNormal", func(t *testing.T) {
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		patcheDescribeNasMountTarget := gomonkey.ApplyMethod(reflect.TypeOf(&NasService{}), "DescribeNasMountTarget", func(*NasService, string) (map[string]interface{}, error) {
			return responseMock["NotFoundError"]("ResourceNotfound")
		})
		err := resourceAliCloudNasMountTargetDelete(d, rawClient)
		patches.Reset()
		patcheDescribeNasMountTarget.Reset()
		assert.Nil(t, err)
	})

	t.Run("DeleteIsExpectedErrors", func(t *testing.T) {
		retryFlag := true
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				return responseMock["RetryError"]("Forbidden.NasNotFound")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		err := resourceAliCloudNasMountTargetDelete(d, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	t.Run("DeleteMockParseResourceId", func(t *testing.T) {
		resourceData1, _ := schema.InternalMap(p["alicloud_nas_mount_target"].Schema).Data(nil, nil)
		resourceData1.SetId("MockId")
		retryFlag := false
		noRetryFlag := false
		patches := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if retryFlag {
				retryFlag = false
				return responseMock["RetryError"]("RetryError")
			} else if noRetryFlag {
				noRetryFlag = false
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["DeleteNormal"]("")
		})
		err := resourceAliCloudNasMountTargetDelete(resourceData1, rawClient)
		patches.Reset()
		assert.Nil(t, err)
	})

	//Read
	t.Run("ReadDescribeNasMountTargetNotFound", func(t *testing.T) {
		patcheDorequest := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			NotFoundFlag := true
			noRetryFlag := false
			if NotFoundFlag {
				return responseMock["NotFoundError"]("ResourceNotfound")
			} else if noRetryFlag {
				return responseMock["NoRetryError"]("NoRetryError")
			}
			return responseMock["ReadNormal"]("")
		})
		err := resourceAliCloudNasMountTargetRead(d, rawClient)
		patcheDorequest.Reset()
		assert.Nil(t, err)
	})

	t.Run("ReadDescribeNasMountTargetAbnormal", func(t *testing.T) {
		patcheDorequest := gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, _ *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			retryFlag := false
			noRetryFlag := true
			if retryFlag {
				return responseMock["RetryError"]("Throttling")
			} else if noRetryFlag {
				return responseMock["NoRetryError"]("NonRetryableError")
			}
			return responseMock["ReadNormal"]("")
		})
		err := resourceAliCloudNasMountTargetRead(d, rawClient)
		patcheDorequest.Reset()
		assert.NotNil(t, err)
	})
}

// Test NAS MountTarget. >>> Resource test cases, automatically generated.
// Case test_extreme_MountTarget
func TestAccAliCloudNASMountTargettest_extreme_MountTarget(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AliCloudNASMountTargettest_extreme_MountTargetMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sNASMountTarget%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudNASMountTargettest_extreme_MountTargetBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"vpc_id":            "${alicloud_vpc.createEVpc.id}",
					"network_type":      "Vpc",
					"vswitch_id":        "${alicloud_vswitch.CreateEVswitch1.id}",
					"security_group_id": "${alicloud_security_group.CreateSecurityGroup.id}",
					"access_group_name": "${alicloud_nas_access_group.create_eaccess_group.access_group_name}",
					"file_system_id":    "${alicloud_nas_file_system.create_extreme_file_system.id}",
					"dual_stack":        true,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dual_stack", "security_group_id"},
			},
		},
	})
}

var AliCloudNASMountTargettest_extreme_MountTargetMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudNASMountTargettest_extreme_MountTargetBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

variable "var_zone_id1" {
  type = object({
    zone_id1 = string
  })
  default = {
    zone_id1 = "cn-hangzhou-g"
  }
}

resource "alicloud_vpc" "createEVpc" {
  cidr_block  = "192.168.0.0/16"
  vpc_name    = "nas-teste0308-vpc"
  enable_ipv6 = true
}

resource "alicloud_nas_file_system" "create_extreme_file_system" {
  description      = "挂载点E0307测试资源"
  storage_type     = "standard"
  zone_id          = var.var_zone_id1.zone_id1
  encrypt_type     = 0
  capacity         = 100
  protocol_type    = "NFS"
  file_system_type = "extreme"
}

resource "alicloud_security_group" "CreateSecurityGroup" {
  security_group_name = "testMountTarget"
  vpc_id              = alicloud_vpc.createEVpc.id
}

resource "alicloud_nas_access_group" "create_eaccess_group" {
  access_group_type = "Vpc"
  description       = "挂载点创建测试"
  access_group_name = "ExtremeMountTarget"
  file_system_type  = "extreme"
}

resource "alicloud_vswitch" "CreateEVswitch1" {
  ipv6_cidr_block_mask = 38
  vpc_id               = alicloud_vpc.createEVpc.id
  zone_id              = var.var_zone_id1.zone_id1
  cidr_block           = "192.168.0.0/24"
  vswitch_name         = "nas-teste0307-vsw1sdw"
  enable_ipv6          = true
}

`, name)
}

// Case test_MountTarget
func TestAccAliCloudNASMountTargettest_MountTarget(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_nas_mount_target.default"
	ra := resourceAttrInit(resourceId, AliCloudNASMountTargettest_MountTargetMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &NasServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeNasMountTarget")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sNASMountTarget%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudNASMountTargettest_MountTargetBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"status":            "Active",
					"vpc_id":            "${alicloud_vpc.createVpc.id}",
					"network_type":      "Vpc",
					"access_group_name": "${alicloud_nas_access_group.create_access_group.access_group_name}",
					"vswitch_id":        "${alicloud_vswitch.CreateVswitch1.id}",
					"file_system_id":    "${alicloud_nas_file_system.create_file_system.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"access_group_name": "${var.access_group_default.access_group_name}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"status": "Inactive",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
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

var AliCloudNASMountTargettest_MountTargetMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudNASMountTargettest_MountTargetBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


variable "var_region_id" {
  type = object({
    region_id = string
  })
  default = {
    region_id = "cn-hangzhou"
  }
}

variable "var_zone_id1" {
  type = object({
    zone_id1 = string
  })
  default = {
    zone_id1 = "cn-hangzhou-g"
  }
}

variable "access_group_default" {
  type = object({
    access_group_name = string
  })
  default = {
    access_group_name = "DEFAULT_VPC_GROUP_NAME"
  }
}

resource "alicloud_vpc" "createVpc" {
  cidr_block = "172.16.0.0/12"
  vpc_name   = "nas-test0307-vpc"
}

resource "alicloud_nas_file_system" "create_file_system" {
  description      = "挂载点0307测试资源"
  storage_type     = "Performance"
  zone_id          = var.var_zone_id1.zone_id1
  encrypt_type     = 0
  protocol_type    = "NFS"
  file_system_type = "standard"
}

resource "alicloud_vswitch" "CreateVswitch1" {
  vpc_id       = alicloud_vpc.createVpc.id
  zone_id      = var.var_zone_id1.zone_id1
  cidr_block   = "172.16.0.0/24"
  vswitch_name = "nas-test0307-vsw1sdw"
}

resource "alicloud_nas_access_group" "create_access_group" {
  access_group_type = "Vpc"
  description       = "挂载点创建测试"
  access_group_name = "StandardMountTarget"
  file_system_type  = "standard"
}

`, name)
}

// Test NAS MountTarget. <<< Resource test cases, automatically generated.
