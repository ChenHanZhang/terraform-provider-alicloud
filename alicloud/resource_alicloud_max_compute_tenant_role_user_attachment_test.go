package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test MaxCompute TenantRoleUserAttachment. >>> Resource test cases, automatically generated.
// Case 国际站RamRole_prod 12483
func TestAccAliCloudMaxComputeTenantRoleUserAttachment_basic12483(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_max_compute_tenant_role_user_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudMaxComputeTenantRoleUserAttachmentMap12483)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MaxComputeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMaxComputeTenantRoleUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmaxcompute%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12483)
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
					"account_id":  "v4_301631103718367601",
					"tenant_role": "super_administrator",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_id":  "v4_301631103718367601",
						"tenant_role": "super_administrator",
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

var AlicloudMaxComputeTenantRoleUserAttachmentMap12483 = map[string]string{}

func AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12483(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 中国站RamRole_prod 12480
func TestAccAliCloudMaxComputeTenantRoleUserAttachment_basic12480(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_max_compute_tenant_role_user_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudMaxComputeTenantRoleUserAttachmentMap12480)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MaxComputeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMaxComputeTenantRoleUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmaxcompute%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12480)
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
					"account_id":  "v4_300703454561238226",
					"tenant_role": "super_administrator",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_id":  "v4_300703454561238226",
						"tenant_role": "super_administrator",
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

var AlicloudMaxComputeTenantRoleUserAttachmentMap12480 = map[string]string{}

func AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12480(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 中国站RamUser_prod 12477
func TestAccAliCloudMaxComputeTenantRoleUserAttachment_basic12477(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_max_compute_tenant_role_user_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudMaxComputeTenantRoleUserAttachmentMap12477)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MaxComputeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMaxComputeTenantRoleUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmaxcompute%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12477)
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
					"account_id":  "p4_202991450150769897",
					"tenant_role": "super_administrator",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_id":  "p4_202991450150769897",
						"tenant_role": "super_administrator",
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

var AlicloudMaxComputeTenantRoleUserAttachmentMap12477 = map[string]string{}

func AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12477(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test MaxCompute TenantRoleUserAttachment. <<< Resource test cases, automatically generated.

func TestAccAliCloudMaxComputeTenantRoleUserAttachment_basic12482(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_max_compute_tenant_role_user_attachment.default"
	ra := resourceAttrInit(resourceId, AlicloudMaxComputeTenantRoleUserAttachmentMap12482)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &MaxComputeServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeMaxComputeTenantRoleUserAttachment")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccmaxcompute%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12482)
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
					"account_id":  "p4_200053869413670560",
					"tenant_role": "super_administrator",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"account_id":  "p4_200053869413670560",
						"tenant_role": "super_administrator",
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

var AlicloudMaxComputeTenantRoleUserAttachmentMap12482 = map[string]string{}

func AlicloudMaxComputeTenantRoleUserAttachmentBasicDependence12482(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_max_compute_tenant_role_user_attachment" "default0" {
  account_id = "p4_200053869413670560"
  tenant_role = "admin"
} 

`, name)
}
