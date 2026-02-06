package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Esa RatePlanInstance. >>> Resource test cases, automatically generated.
// Case resource_RatePlanInstance_test1 12540
func TestAccAliCloudEsaRatePlanInstance_basic12540(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap12540)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence12540)
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
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"payment_type": "Subscription",
					"coverage":     "overseas",
					"auto_pay":     "true",
					"plan_name":    "basic",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"payment_type": "Subscription",
						"coverage":     "overseas",
						"auto_pay":     "true",
						"plan_name":    "basic",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "medium",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "medium",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap12540 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence12540(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case resource_RatePlanInstance_test0 12541
func TestAccAliCloudEsaRatePlanInstance_basic12541(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap12541)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence12541)
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
					"coverage":  "domestic,overseas,global",
					"plan_name": "Youheng_all2412",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"coverage":  "domestic,overseas,global",
						"plan_name": "Youheng_all2412",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_burst_gbps": "30",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_burst_gbps": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_burst_gbps": "100",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_burst_gbps": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"max_burst_gbps": "300",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"max_burst_gbps": CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap12541 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence12541(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 套餐1729154520851 8245
func TestAccAliCloudEsaRatePlanInstance_basic8245(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap8245)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence8245)
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
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"payment_type": "Subscription",
					"coverage":     "overseas",
					"plan_name":    "basic",
					"auto_pay":     "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"payment_type": "Subscription",
						"coverage":     "overseas",
						"plan_name":    "basic",
						"auto_pay":     "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "medium",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "medium",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap8245 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence8245(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 套餐_副本 8110
func TestAccAliCloudEsaRatePlanInstance_basic8110(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap8110)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence8110)
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
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"coverage":     "overseas",
					"plan_name":    "basic",
					"auto_pay":     "true",
					"payment_type": "Subscription",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"coverage":     "overseas",
						"plan_name":    "basic",
						"auto_pay":     "true",
						"payment_type": "Subscription",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "medium",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "medium",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap8110 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence8110(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 套餐 8052
func TestAccAliCloudEsaRatePlanInstance_basic8052(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap8052)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence8052)
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
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"payment_type": "Subscription",
					"coverage":     "overseas",
					"plan_name":    "basic",
					"auto_pay":     "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"payment_type": "Subscription",
						"coverage":     "overseas",
						"plan_name":    "basic",
						"auto_pay":     "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "medium",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "medium",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap8052 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence8052(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 套餐_2.0 8489
func TestAccAliCloudEsaRatePlanInstance_basic8489(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap8489)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence8489)
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
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"payment_type": "Subscription",
					"coverage":     "overseas",
					"plan_name":    "basic",
					"auto_pay":     "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"payment_type": "Subscription",
						"coverage":     "overseas",
						"plan_name":    "basic",
						"auto_pay":     "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "medium",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "medium",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap8489 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence8489(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Case 套餐_3.0 8926
func TestAccAliCloudEsaRatePlanInstance_basic8926(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap8926)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence8926)
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
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"payment_type": "Subscription",
					"coverage":     "overseas",
					"plan_name":    "basic",
					"auto_pay":     "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"payment_type": "Subscription",
						"coverage":     "overseas",
						"plan_name":    "basic",
						"auto_pay":     "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "medium",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "medium",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}

var AlicloudEsaRatePlanInstanceMap8926 = map[string]string{
	"status":          CHECKSET,
	"create_time":     CHECKSET,
	"instance_status": CHECKSET,
}

func AlicloudEsaRatePlanInstanceBasicDependence8926(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test Esa RatePlanInstance. <<< Resource test cases, automatically generated.

func TestAccAliCloudEsaRatePlanInstance_intl(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_rate_plan_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaRatePlanInstanceMap8489)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaRatePlanInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sesarateplaninstance%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaRatePlanInstanceBasicDependence8489)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithAccountSiteType(t, IntlSite)
			testAccPreCheckWithRegions(t, true, []connectivity.Region{"cn-hangzhou"})
			testAccPreCheck(t)
		},
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"type":         "NS",
					"auto_renew":   "true",
					"period":       "1",
					"payment_type": "Subscription",
					"coverage":     "overseas",
					"plan_name":    "entranceplan_intl",
					"auto_pay":     "true",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"type":         "NS",
						"auto_renew":   "true",
						"period":       "1",
						"payment_type": "Subscription",
						"coverage":     "overseas",
						"plan_name":    "entranceplan_intl",
						"auto_pay":     "true",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"plan_name": "basicplan_intl",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"plan_name": "basicplan_intl",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auto_pay", "auto_renew", "coverage", "period", "type"},
			},
		},
	})
}
