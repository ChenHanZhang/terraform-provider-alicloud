package alicloud

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

// Test CloudFirewall Instance. >>> Resource test cases, automatically generated.
// Case 国内版按量付费2.0 11709
func TestAccAliCloudCloudFirewallInstance_basic11709(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap11709)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence11709)
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
					"product_code":          "cfw",
					"product_type":          "cfw_elasticity_public_cn",
					"payment_type":          "PayAsYouGo",
					"spec":                  "payg_version",
					"cfw_log":               "false",
					"renewal_duration_unit": "M",
					"sdl":                   "false",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "cfw",
						"product_type":          "cfw_elasticity_public_cn",
						"payment_type":          "PayAsYouGo",
						"spec":                  "payg_version",
						"cfw_log":               "false",
						"renewal_duration_unit": "M",
						"sdl":                   "false",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"sdl":         "true",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"sdl":         "true",
						"modify_type": "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap11709 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence11709(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版包年包月2.0_旗舰版 11882
func TestAccAliCloudCloudFirewallInstance_basic11882(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap11882)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence11882)
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
					"product_code":          "cfw",
					"renewal_status":        "ManualRenewal",
					"product_type":          "cfw_sub_public_cn",
					"payment_type":          "Subscription",
					"cfw_log":               "true",
					"spec":                  "ultimate_version",
					"sdl":                   "false",
					"period":                "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "cfw",
						"renewal_status":        "ManualRenewal",
						"product_type":          "cfw_sub_public_cn",
						"payment_type":          "Subscription",
						"cfw_log":               "true",
						"spec":                  "ultimate_version",
						"sdl":                   "false",
						"period":                "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status":        "AutoRenewal",
					"renewal_duration_unit": "Y",
					"modify_type":           "Upgrade",
					"renewal_duration":      "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status":        "AutoRenewal",
						"renewal_duration_unit": "Y",
						"modify_type":           "Upgrade",
						"renewal_duration":      "2",
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap11882 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence11882(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版包年包月2.0_高级版 11859
func TestAccAliCloudCloudFirewallInstance_basic11859(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap11859)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence11859)
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
					"product_code":     "cfw",
					"renewal_status":   "AutoRenewal",
					"product_type":     "cfw_sub_public_cn",
					"payment_type":     "Subscription",
					"cfw_log":          "true",
					"spec":             "premium_version",
					"sdl":              "false",
					"period":           "1",
					"renewal_duration": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":     "cfw",
						"renewal_status":   "AutoRenewal",
						"product_type":     "cfw_sub_public_cn",
						"payment_type":     "Subscription",
						"cfw_log":          "true",
						"spec":             "premium_version",
						"sdl":              "false",
						"period":           "1",
						"renewal_duration": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_type":           "Upgrade",
					"renewal_duration_unit": "Y",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_type":           "Upgrade",
						"renewal_duration_unit": "Y",
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
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap11859 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence11859(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版包年包月2.0_企业版 11879
func TestAccAliCloudCloudFirewallInstance_basic11879(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap11879)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence11879)
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
					"product_code":          "cfw",
					"renewal_status":        "ManualRenewal",
					"product_type":          "cfw_sub_public_cn",
					"payment_type":          "Subscription",
					"cfw_log":               "true",
					"spec":                  "enterprise_version",
					"sdl":                   "false",
					"period":                "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "cfw",
						"renewal_status":        "ManualRenewal",
						"product_type":          "cfw_sub_public_cn",
						"payment_type":          "Subscription",
						"cfw_log":               "true",
						"spec":                  "enterprise_version",
						"sdl":                   "false",
						"period":                "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"modify_type": "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap11879 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence11879(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 创建后付费实例 3387
func TestAccAliCloudCloudFirewallInstance_basic3387(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap3387)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence3387)
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
					"product_code":   "cfw",
					"renewal_status": "AutoRenewal",
					"product_type":   "cfw_elasticity_public_cn",
					"period":         "12",
					"payment_type":   "PayAsYouGo",
					"spec":           "10",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":   "cfw",
						"renewal_status": "AutoRenewal",
						"product_type":   "cfw_elasticity_public_cn",
						"period":         "12",
						"payment_type":   "PayAsYouGo",
						"spec":           "10",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap3387 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence3387(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版按量付费_PayAsYouGo 6852
func TestAccAliCloudCloudFirewallInstance_basic6852(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap6852)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence6852)
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
					"product_code":   "cfw",
					"product_type":   "cfw_elasticity_public_cn",
					"payment_type":   "PayAsYouGo",
					"spec":           "10",
					"cfw_log":        "false",
					"renewal_status": "ManualRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":   "cfw",
						"product_type":   "cfw_elasticity_public_cn",
						"payment_type":   "PayAsYouGo",
						"spec":           "10",
						"cfw_log":        "false",
						"renewal_status": "ManualRenewal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap6852 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence6852(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 线上环境创建后付费实例 3427
func TestAccAliCloudCloudFirewallInstance_basic3427(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap3427)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence3427)
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
					"product_code": "cfw",
					"product_type": "cfw_elasticity_public_cn",
					"payment_type": "PayAsYouGo",
					"spec":         "4",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code": "cfw",
						"product_type": "cfw_elasticity_public_cn",
						"payment_type": "PayAsYouGo",
						"spec":         "4",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap3427 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence3427(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版包年包月_Subscription_V4 7235
func TestAccAliCloudCloudFirewallInstance_basic7235(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7235)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7235)
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
					"product_code":   "vipcloudfw",
					"renewal_status": "AutoRenewal",
					"product_type":   "vipcloudfw",
					"period":         "1",
					"payment_type":   "Subscription",
					"cfw_log":        "false",
					"spec":           "4",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":   "vipcloudfw",
						"renewal_status": "AutoRenewal",
						"product_type":   "vipcloudfw",
						"period":         "1",
						"payment_type":   "Subscription",
						"cfw_log":        "false",
						"spec":           "4",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7235 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7235(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 按量付费_PayAsYouGo 7518
func TestAccAliCloudCloudFirewallInstance_basic7518(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7518)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7518)
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
					"product_code":          "cfw",
					"product_type":          "cfw_elasticity_public_cn",
					"payment_type":          "PayAsYouGo",
					"spec":                  "payg_version",
					"cfw_log":               "false",
					"renewal_status":        "ManualRenewal",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "cfw",
						"product_type":          "cfw_elasticity_public_cn",
						"payment_type":          "PayAsYouGo",
						"spec":                  "payg_version",
						"cfw_log":               "false",
						"renewal_status":        "ManualRenewal",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7518 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7518(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V2 7535
func TestAccAliCloudCloudFirewallInstance_basic7535(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7535)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7535)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "ManualRenewal",
					"product_type":          "vipcloudfw",
					"period":                "1",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "premium_version",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "ManualRenewal",
						"product_type":          "vipcloudfw",
						"period":                "1",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "premium_version",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status":   "AutoRenewal",
					"cfw_log":          "true",
					"modify_type":      "Upgrade",
					"renewal_duration": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status":   "AutoRenewal",
						"cfw_log":          "true",
						"modify_type":      "Upgrade",
						"renewal_duration": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7535 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7535(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版包年包月_Subscription_V3 7231
func TestAccAliCloudCloudFirewallInstance_basic7231(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7231)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7231)
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
					"product_code":   "vipcloudfw",
					"renewal_status": "AutoRenewal",
					"product_type":   "vipcloudfw",
					"period":         "1",
					"payment_type":   "Subscription",
					"cfw_log":        "false",
					"spec":           "3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":   "vipcloudfw",
						"renewal_status": "AutoRenewal",
						"product_type":   "vipcloudfw",
						"period":         "1",
						"payment_type":   "Subscription",
						"cfw_log":        "false",
						"spec":           "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7231 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7231(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 国内版包年包月_Subscription_V2 2391
func TestAccAliCloudCloudFirewallInstance_basic2391(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap2391)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence2391)
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
					"product_code":   "vipcloudfw",
					"renewal_status": "ManualRenewal",
					"product_type":   "vipcloudfw",
					"period":         "1",
					"payment_type":   "Subscription",
					"cfw_log":        "false",
					"spec":           "2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":   "vipcloudfw",
						"renewal_status": "ManualRenewal",
						"product_type":   "vipcloudfw",
						"period":         "1",
						"payment_type":   "Subscription",
						"cfw_log":        "false",
						"spec":           "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status": "AutoRenewal",
					"cfw_log":        "true",
					"modify_type":    "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status": "AutoRenewal",
						"cfw_log":        "true",
						"modify_type":    "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap2391 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence2391(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V3 7546
func TestAccAliCloudCloudFirewallInstance_basic7546(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7546)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7546)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "AutoRenewal",
					"product_type":          "vipcloudfw",
					"period":                "1",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "enterprise_version",
					"renewal_duration":      "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "AutoRenewal",
						"product_type":          "vipcloudfw",
						"period":                "1",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "enterprise_version",
						"renewal_duration":      "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7546 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7546(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V4 7547
func TestAccAliCloudCloudFirewallInstance_basic7547(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7547)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7547)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "AutoRenewal",
					"product_type":          "vipcloudfw",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "ultimate_version",
					"period":                "1",
					"renewal_duration":      "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "AutoRenewal",
						"product_type":          "vipcloudfw",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "ultimate_version",
						"period":                "1",
						"renewal_duration":      "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":               "true",
					"renewal_duration":      "2",
					"renewal_duration_unit": "Y",
					"modify_type":           "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":               "true",
						"renewal_duration":      "2",
						"renewal_duration_unit": "Y",
						"modify_type":           "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7547 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7547(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 按量付费_PayAsYouGo_1 7992
func TestAccAliCloudCloudFirewallInstance_basic7992(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7992)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7992)
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
					"product_code":          "cfw",
					"product_type":          "cfw_elasticity_public_cn",
					"payment_type":          "PayAsYouGo",
					"spec":                  "payg_version",
					"cfw_log":               "false",
					"renewal_status":        "ManualRenewal",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "cfw",
						"product_type":          "cfw_elasticity_public_cn",
						"payment_type":          "PayAsYouGo",
						"spec":                  "payg_version",
						"cfw_log":               "false",
						"renewal_status":        "ManualRenewal",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Downgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Downgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7992 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7992(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V4_1 7997
func TestAccAliCloudCloudFirewallInstance_basic7997(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7997)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7997)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "AutoRenewal",
					"product_type":          "vipcloudfw",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "ultimate_version",
					"period":                "1",
					"renewal_duration":      "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "AutoRenewal",
						"product_type":          "vipcloudfw",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "ultimate_version",
						"period":                "1",
						"renewal_duration":      "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":               "true",
					"renewal_duration":      "2",
					"renewal_duration_unit": "Y",
					"modify_type":           "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":               "true",
						"renewal_duration":      "2",
						"renewal_duration_unit": "Y",
						"modify_type":           "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7997 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7997(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V2_1 7993
func TestAccAliCloudCloudFirewallInstance_basic7993(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7993)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7993)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "ManualRenewal",
					"product_type":          "vipcloudfw",
					"period":                "1",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "premium_version",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "ManualRenewal",
						"product_type":          "vipcloudfw",
						"period":                "1",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "premium_version",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status":   "AutoRenewal",
					"cfw_log":          "true",
					"modify_type":      "Upgrade",
					"renewal_duration": "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status":   "AutoRenewal",
						"cfw_log":          "true",
						"modify_type":      "Upgrade",
						"renewal_duration": "1",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7993 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7993(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V3_1 7995
func TestAccAliCloudCloudFirewallInstance_basic7995(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap7995)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence7995)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "AutoRenewal",
					"product_type":          "vipcloudfw",
					"period":                "1",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "enterprise_version",
					"renewal_duration":      "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "AutoRenewal",
						"product_type":          "vipcloudfw",
						"period":                "1",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "enterprise_version",
						"renewal_duration":      "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":     "true",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":     "true",
						"modify_type": "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap7995 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence7995(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 包年包月_Subscription_V4_2 8004
func TestAccAliCloudCloudFirewallInstance_basic8004(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudCloudFirewallInstanceMap8004)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &CloudFirewallServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeCloudFirewallInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfacccloudfirewall%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudCloudFirewallInstanceBasicDependence8004)
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
					"product_code":          "vipcloudfw",
					"renewal_status":        "AutoRenewal",
					"product_type":          "vipcloudfw",
					"payment_type":          "Subscription",
					"cfw_log":               "false",
					"spec":                  "ultimate_version",
					"period":                "1",
					"renewal_duration":      "1",
					"renewal_duration_unit": "M",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"product_code":          "vipcloudfw",
						"renewal_status":        "AutoRenewal",
						"product_type":          "vipcloudfw",
						"payment_type":          "Subscription",
						"cfw_log":               "false",
						"spec":                  "ultimate_version",
						"period":                "1",
						"renewal_duration":      "1",
						"renewal_duration_unit": "M",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":               "true",
					"renewal_duration":      "2",
					"renewal_duration_unit": "Y",
					"modify_type":           "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":               "true",
						"renewal_duration":      "2",
						"renewal_duration_unit": "Y",
						"modify_type":           "Upgrade",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"modify_type", "period"},
			},
		},
	})
}

var AlicloudCloudFirewallInstanceMap8004 = map[string]string{
	"end_time":     CHECKSET,
	"user_status":  CHECKSET,
	"status":       CHECKSET,
	"create_time":  CHECKSET,
	"release_time": CHECKSET,
}

func AlicloudCloudFirewallInstanceBasicDependence8004(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test CloudFirewall Instance. <<< Resource test cases, automatically generated.

func TestAccAliCloudCloudFirewallInstance_basic0(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallInstanceMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &BssOpenApiService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "QueryAvailableInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%ssddpinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallInstanceBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckForCleanUpInstances(t, "", "vipcloudfw", "vipcloudfw", "cfw", "cfw_pre_intl")
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type": "PayAsYouGo",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type": "PayAsYouGo",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log": "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log": "true",
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

func TestAccAliCloudCloudFirewallInstance_basic0_twin(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallInstanceMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &BssOpenApiService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "QueryAvailableInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%ssddpinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallInstanceBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckForCleanUpInstances(t, "", "vipcloudfw", "vipcloudfw", "cfw", "cfw_pre_intl")
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type": "PayAsYouGo",
					"spec":         "payg_version",
					"cfw_log":      "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type": "PayAsYouGo",
						"spec":         "payg_version",
						"cfw_log":      "true",
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

func TestAccAliCloudCloudFirewallInstance_basic1(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_cloud_firewall_instance.default"
	ra := resourceAttrInit(resourceId, AliCloudCloudFirewallInstanceMap0)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &BssOpenApiService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "QueryAvailableInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%ssddpinstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudCloudFirewallInstanceBasicDependence0)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckForCleanUpInstances(t, "", "vipcloudfw", "vipcloudfw", "cfw", "cfw_pre_intl")
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  nil,
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"payment_type": "Subscription",
					"spec":         "enterprise_version",
					"ip_number":    "50",
					"band_width":   "50",
					"cfw_log":      "false",
					"period":       "1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"payment_type": "Subscription",
						"spec":         "enterprise_version",
						"ip_number":    "50",
						"cfw_log":      "false",
						"period":       "1",
					}),
				),
			},
			// premium_version does not support fw_vpc_number
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"fw_vpc_number": "3",
			//		"modify_type":   "Upgrade",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"fw_vpc_number": "3",
			//			"modify_type":   "Upgrade",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"band_width":  "55",
					"modify_type": "Upgrade",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fw_vpc_number": "2",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log":         "true",
					"cfw_log_storage": "3000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log":         "true",
						"cfw_log_storage": "3000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_log_storage": "5000",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_log_storage": "5000",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ip_number": "55",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ip_number": "55",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"fw_vpc_number": "5",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"fw_vpc_number": "5",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_duration":      "1",
					"renewal_duration_unit": "Month",
					"renewal_status":        "AutoRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_duration":      "1",
						"renew_period":          "1",
						"renewal_duration_unit": "Month",
						"renewal_status":        "AutoRenewal",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"renewal_status": "ManualRenewal",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"renewal_status":        "ManualRenewal",
						"renewal_duration":      REMOVEKEY,
						"renew_period":          REMOVEKEY,
						"renewal_duration_unit": REMOVEKEY,
					}),
				),
			},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"renewal_duration": REMOVEKEY,
			//		"renew_period":     "2",
			//		"renewal_status":   "AutoRenewal",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"renewal_duration":      "2",
			//			"renew_period":          "2",
			//			"renewal_duration_unit": "Month",
			//			"renewal_status":        "AutoRenewal",
			//		}),
			//	),
			//},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"renewal_status": "NotRenewal",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"renewal_status":        "NotRenewal",
			//			"renewal_duration":      REMOVEKEY,
			//			"renew_period":          REMOVEKEY,
			//			"renewal_duration_unit": REMOVEKEY,
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"cfw_account":    "true",
					"account_number": "10",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cfw_account":    "true",
						"account_number": "10",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"band_width", "period", "modify_type", "cfw_account", "account_number"},
			},
		},
	})
}

var AliCloudCloudFirewallInstanceMap0 = map[string]string{
	"user_status": CHECKSET,
	"status":      CHECKSET,
}

func AliCloudCloudFirewallInstanceBasicDependence0(name string) string {
	return fmt.Sprintf(` 
	variable "name" {
  		default = "%s"
	}
`, name)
}

// lintignore: R001
func TestUnitAliCloudCloudFirewallInstance(t *testing.T) {
	p := Provider().(*schema.Provider).ResourcesMap
	dInit, _ := schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(nil, nil)
	dExisted, _ := schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(nil, nil)
	dInit.MarkNewResource()
	attributes := map[string]interface{}{
		"payment_type":    "CreateInstanceValue",
		"spec":            "CreateInstanceValue",
		"renewal_status":  "CreateInstanceValue",
		"ip_number":       20,
		"band_width":      10,
		"cfw_log":         false,
		"cfw_log_storage": 1000,
		"cfw_service":     false,
		"period":          6,
		"fw_vpc_number":   10,
		"instance_count":  10,
		"logistics":       "CreateInstanceValue",
	}
	for key, value := range attributes {
		err := dInit.Set(key, value)
		assert.Nil(t, err)
		err = dExisted.Set(key, value)
		assert.Nil(t, err)
		if err != nil {
			log.Printf("[ERROR] the field %s setting error", key)
		}
	}
	region := os.Getenv("ALICLOUD_REGION")
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		t.Skipf("Skipping the test case with err: %s", err)
		t.Skipped()
	}
	rawClient = rawClient.(*connectivity.AliyunClient)
	ReadMockResponse := map[string]interface{}{
		// QueryAvailableInstances
		"Data": map[string]interface{}{
			"InstanceList": []interface{}{
				map[string]interface{}{
					"InstanceID":          "CreateInstanceValue",
					"CreateTime":          "CreateInstanceValue",
					"RenewStatus":         "CreateInstanceValue",
					"RenewalDurationUnit": "M",
					"Status":              "CreateInstanceValue",
					"SubscriptionType":    "CreateInstanceValue",
					"EndTime":             "CreateInstanceValue",
				},
			},
			"InstanceId": "CreateInstanceValue",
		},
		"Code": "Success",
	}
	CreateMockResponse := map[string]interface{}{
		// CreateInstance
		"Data": map[string]interface{}{
			"InstanceId": "CreateInstanceValue",
		},
		"Code": "Success",
	}
	failedResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, &tea.SDKError{
			Code:       String(errorCode),
			Data:       String(errorCode),
			Message:    String(errorCode),
			StatusCode: tea.Int(400),
		}
	}
	notFoundResponseMock := func(errorCode string) (map[string]interface{}, error) {
		return nil, GetNotFoundErrorFromString(GetNotFoundMessage("alicloud_cloud_firewall_instance", errorCode))
	}
	successResponseMock := func(operationMockResponse map[string]interface{}) (map[string]interface{}, error) {
		if len(operationMockResponse) > 0 {
			mapMerge(ReadMockResponse, operationMockResponse)
		}
		return ReadMockResponse, nil
	}

	// Create
	patches := gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewBssopenapiClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudCloudFirewallInstanceCreate(dInit, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	ReadMockResponseDiff := map[string]interface{}{
		// QueryAvailableInstances Response
		"Data": map[string]interface{}{
			"InstanceId": "CreateInstanceValue",
		},
		"Code": "Success",
	}
	errorCodes := []string{"NonRetryableError", "Throttling", "NotApplicable", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1 // a counter used to cover retry scenario; the same below
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "CreateInstance" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						successResponseMock(ReadMockResponseDiff)
						return CreateMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudCloudFirewallInstanceCreate(dInit, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(dInit.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dInit.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Update
	patches = gomonkey.ApplyMethod(reflect.TypeOf(&connectivity.AliyunClient{}), "NewBssopenapiClient", func(_ *connectivity.AliyunClient) (*client.Client, error) {
		return nil, &tea.SDKError{
			Code:       String("loadEndpoint error"),
			Data:       String("loadEndpoint error"),
			Message:    String("loadEndpoint error"),
			StatusCode: tea.Int(400),
		}
	})
	err = resourceAliCloudCloudFirewallInstanceUpdate(dExisted, rawClient)
	patches.Reset()
	assert.NotNil(t, err)
	// RenewInstance
	attributesDiff := map[string]interface{}{
		"renew_period": 1,
	}
	diff, err := newInstanceDiff("alicloud_cloud_firewall_instance", attributes, attributesDiff, dInit.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(dInit.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// QueryAvailableInstances Response
		"Data": map[string]interface{}{
			"InstanceList": []interface{}{
				map[string]interface{}{
					"RenewPeriod": 1,
				},
			},
		},
		"Code": "Success",
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "NotApplicable", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "RenewInstance" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudCloudFirewallInstanceUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// ModifyInstance
	attributesDiff = map[string]interface{}{
		"cfw_service":     true,
		"fw_vpc_number":   20,
		"ip_number":       30,
		"cfw_log_storage": 2000,
		"cfw_log":         true,
		"band_width":      20,
		"spec":            "enterprise_version",
		"instance_count":  20,
		"modify_type":     "ModifyInstanceValue",
	}
	diff, err = newInstanceDiff("alicloud_cloud_firewall_instance", attributes, attributesDiff, dExisted.State())
	if err != nil {
		t.Error(err)
	}
	dExisted, _ = schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(dExisted.State(), diff)
	ReadMockResponseDiff = map[string]interface{}{
		// QueryAvailableInstances Response
		"Data": map[string]interface{}{
			"InstanceList": []interface{}{
				map[string]interface{}{
					"CfwService":    true,
					"FwVpcNumber":   20,
					"IpNumber":      30,
					"CfwLogStorage": 2000,
					"CfwLog":        true,
					"BandWidth":     20,
					"Spec":          "enterprise_version",
					"InstanceCount": 20,
					"ModifyType":    "ModifyInstanceValue",
				},
			},
		},
		"Code": "Success",
	}
	errorCodes = []string{"NonRetryableError", "Throttling", "NotApplicable", "nil"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "ModifyInstance" {
				switch errorCode {
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if retryIndex >= len(errorCodes)-1 {
						return successResponseMock(ReadMockResponseDiff)
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudCloudFirewallInstanceUpdate(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		default:
			assert.Nil(t, err)
			dCompare, _ := schema.InternalMap(p["alicloud_cloud_firewall_instance"].Schema).Data(dExisted.State(), nil)
			for key, value := range attributes {
				_ = dCompare.Set(key, value)
			}
			assert.Equal(t, dCompare.State().Attributes, dExisted.State().Attributes)
		}
		if retryIndex >= len(errorCodes)-1 {
			break
		}
	}

	// Read
	errorCodes = []string{"NonRetryableError", "Throttling", "NotApplicable", "nil", "{}"}
	for index, errorCode := range errorCodes {
		retryIndex := index - 1
		patches = gomonkey.ApplyMethod(reflect.TypeOf(&client.Client{}), "DoRequest", func(_ *client.Client, action *string, _ *string, _ *string, _ *string, _ *string, _ map[string]interface{}, _ map[string]interface{}, _ *util.RuntimeOptions) (map[string]interface{}, error) {
			if *action == "QueryAvailableInstances" {
				switch errorCode {
				case "{}":
					return notFoundResponseMock(errorCode)
				case "NonRetryableError":
					return failedResponseMock(errorCode)
				default:
					retryIndex++
					if errorCodes[retryIndex] == "nil" {
						return ReadMockResponse, nil
					}
					return failedResponseMock(errorCodes[retryIndex])
				}
			}
			return ReadMockResponse, nil
		})
		err := resourceAliCloudCloudFirewallInstanceRead(dExisted, rawClient)
		patches.Reset()
		switch errorCode {
		case "NonRetryableError":
			assert.NotNil(t, err)
		case "{}":
			assert.Nil(t, err)
		}
	}

	// Delete
	err = resourceAliCloudCloudFirewallInstanceDelete(dExisted, rawClient)
	assert.Nil(t, err)

}
