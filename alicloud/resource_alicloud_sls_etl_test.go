package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Sls Etl. >>> Resource test cases, automatically generated.
// Case test1 10468
func TestAccAliCloudSlsEtl_basic10468(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_sls_etl.default"
	ra := resourceAttrInit(resourceId, AlicloudSlsEtlMap10468)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SlsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSlsEtl")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsls%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSlsEtlBasicDependence10468)
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
					"project":     "${alicloud_log_project.defaulthhAPo6.id}",
					"description": "etl-1740472705-185721",
					"configuration": []map[string]interface{}{
						{
							"script":   "* | extend a=1",
							"lang":     "SPL",
							"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
							"sink": []map[string]interface{}{
								{
									"name":     "11111",
									"endpoint": "cn-hangzhou-intranet.log.aliyuncs.com",
									"project":  "gy-hangzhou-huolang-1",
									"logstore": "gy-rm2",
									"datasets": []string{
										"__UNNAMED__"},
									"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
								},
							},
							"logstore":  "${alicloud_log_store.defaultzWKLkp.name}",
							"from_time": "1706771697",
							"to_time":   "1738394097",
						},
					},
					"job_name":     "etl-1740472705-185721",
					"display_name": "etl-1740472705-185721",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"project":      CHECKSET,
						"description":  "etl-1740472705-185721",
						"job_name":     "etl-1740472705-185721",
						"display_name": "etl-1740472705-185721",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test",
					"configuration": []map[string]interface{}{
						{
							"script":   "* | extend a=2",
							"lang":     "SPL",
							"logstore": "${alicloud_log_store.defaultzWKLkp.name}",
							"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
							"sink": []map[string]interface{}{
								{
									"name":     "11111",
									"endpoint": "https://cn-qingdao.log.aliyuncs.com",
									"project":  "hclcn-qingdao",
									"logstore": "test",
									"datasets": []string{
										"__UNNAMED__"},
									"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
								},
							},
							"from_time": "1706771697",
							"to_time":   "1738394097",
						},
					},
					"display_name": "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "test",
						"display_name": "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "test22",
					"configuration": []map[string]interface{}{
						{
							"script":   "* | extend a=3",
							"lang":     "SPL",
							"logstore": "${alicloud_log_store.defaultzWKLkp.name}",
							"role_arn": "acs:ram::1395894005860:role/aliyunlogetlrole",
							"sink": []map[string]interface{}{
								{
									"name":     "11111",
									"endpoint": "https://cn-shanghai.log.aliyuncs.com",
									"project":  "hclcn-shanghai-b",
									"logstore": "test",
									"datasets": []string{
										"__set__"},
									"role_arn": "acs:ram::1395894005860:role/aliyunlogetlrole",
								},
								{
									"name":     "22222",
									"endpoint": "https://cn-beijing.log.aliyuncs.com",
									"project":  "hclcn-beijing",
									"logstore": "test",
									"datasets": []string{
										"__UNNAMED__"},
									"role_arn": "acs:ram::13958940050:role/aliyunlogetlrole",
								},
							},
							"parameters": map[string]interface{}{
								"\"AK\"": "11111",
								"\"SK\"": "22222",
							},
							"from_time": "1706771697",
							"to_time":   "1738394097",
						},
					},
					"display_name": "55555",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "test22",
						"display_name": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "88888",
					"configuration": []map[string]interface{}{
						{
							"script":   "* | extend c=2",
							"logstore": "${alicloud_log_store.defaultzWKLkp.name}",
							"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
							"sink": []map[string]interface{}{
								{
									"name":     "11111",
									"endpoint": "https://cn-nanjing.log.aliyuncs.com",
									"project":  "hcl-cn-nanjing",
									"logstore": "test",
									"datasets": []string{
										"__UNNAMED__"},
									"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
								},
								{
									"name":     "22222",
									"endpoint": "https://cn-shenzhen.log.aliyuncs.com",
									"project":  "hclcn-shenzhen",
									"logstore": "test",
									"datasets": []string{
										"__UU__"},
									"role_arn": "acs:ram::1395894005860:role/aliyunlogetlrole",
								},
								{
									"name":     "33333",
									"endpoint": "https://cn-wulanchabu.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test",
									"datasets": []string{
										"qqq"},
									"role_arn": "acs:ram::1395894005860:role/aliyunlogetlrole",
								},
								{
									"name":     "44444",
									"endpoint": "https://cn-huhehaote.log.aliyuncs.com",
									"project":  "test1",
									"logstore": "test1",
									"datasets": []string{
										"aaa"},
									"role_arn": "acs:ram::13958940022220:role/aliyunlogetlrole",
								},
								{
									"name":     "55555",
									"endpoint": "https://cn-zhangjiakou.log.aliyuncs.com",
									"project":  "test2",
									"logstore": "test2",
									"datasets": []string{
										"ggg"},
									"role_arn": "acs:ram::13958940022220:role/aliyunlogetlrole",
								},
							},
							"lang":      "SPL",
							"from_time": "1706771697",
							"to_time":   "1738394097",
							"parameters": map[string]interface{}{
								"\"AK\"": "333",
								"\"SK\"": "444",
							},
						},
					},
					"display_name": "3456776",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  CHECKSET,
						"display_name": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "2222222",
					"configuration": []map[string]interface{}{
						{
							"script":    "* | extend d=1",
							"lang":      "SPL",
							"logstore":  "${alicloud_log_store.defaultzWKLkp.name}",
							"role_arn":  "acs:ram::1395894005868720:role/aliyunlogetlrole",
							"from_time": "1706771697",
							"to_time":   "1738394097",
							"sink": []map[string]interface{}{
								{
									"name":     "33333",
									"endpoint": "https://cn-huhehaote.log.aliyuncs.com",
									"project":  "test3",
									"logstore": "test3",
									"datasets": []string{
										"cc"},
									"role_arn": "acs:ram::1395894005820:role/aliyunlogetlrole",
								},
								{
									"name":     "44444",
									"endpoint": "https://cn-zhangjiakou.log.aliyuncs.com",
									"project":  "test4",
									"logstore": "test4",
									"datasets": []string{
										"lll"},
									"role_arn": "acs:ram::1395894005820:role/aliyunlogetlrole",
								},
								{
									"name":     "55555",
									"endpoint": "https://cn-shenzhen.log.aliyuncs.com",
									"project":  "test5",
									"logstore": "test5",
									"datasets": []string{
										"kkk"},
									"role_arn": "acs:ram::1395894005820:role/aliyunlogetlrole",
								},
							},
							"parameters": map[string]interface{}{
								"\"system.join.enable_cache\"": "true",
							},
						},
					},
					"display_name": "6666666",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  CHECKSET,
						"display_name": CHECKSET,
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

var AlicloudSlsEtlMap10468 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudSlsEtlBasicDependence10468(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_log_project" "defaulthhAPo6" {
  description = "terraform-etl-test-188"
  name        = "terraform-etl-test-96"
}

resource "alicloud_log_store" "defaultzWKLkp" {
  hot_ttl          = "8"
  retention_period = "30"
  shard_count      = "2"
  project          = alicloud_log_project.defaulthhAPo6.id
  name             = "test"
}


`, name)
}

// Case test2 10469
func TestAccAliCloudSlsEtl_basic10469(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_sls_etl.default"
	ra := resourceAttrInit(resourceId, AlicloudSlsEtlMap10469)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SlsServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSlsEtl")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsls%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSlsEtlBasicDependence10469)
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
					"project":     "${alicloud_log_project.defaultpYdUWa.id}",
					"description": "etl-1740562885-856622",
					"configuration": []map[string]interface{}{
						{
							"script":    "* | extend d=1",
							"logstore":  "${alicloud_log_store.defaultDtiWb2.name}",
							"role_arn":  "acs:ram::1395894005868720:role/aliyunlogetlrole",
							"from_time": "1",
							"to_time":   "1738394097",
							"sink": []map[string]interface{}{
								{
									"name":     "11111",
									"endpoint": "https://cn-guangzhou.log.aliyuncs.com",
									"project":  "test-cn-guangzhou-ha",
									"logstore": "test",
									"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
									"datasets": []string{
										"__SET__"},
								},
								{
									"name":     "22222",
									"endpoint": "https://cn-chengdu.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test",
									"datasets": []string{
										"__eeee__"},
									"role_arn": "acs:ram::1395894008720:role/aliyunlogetlrole",
								},
								{
									"name":     "33333",
									"endpoint": "https://cn-heyuan.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test",
									"datasets": []string{
										"ddd"},
									"role_arn": "acs:ram::1395894008720:role/aliyunlogetlrole",
								},
								{
									"name":     "44444",
									"endpoint": "https://cn-wulanchabu.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test",
									"datasets": []string{
										"fff"},
									"role_arn": "acs:ram::1395894008720:role/aliyunlogetlrole",
								},
							},
							"parameters": map[string]interface{}{
								"\"AK\"": "888",
							},
							"lang": "SPL",
						},
					},
					"job_name":     "etl-1740562885-856622",
					"display_name": "etl-1740562885-856622",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"project":      CHECKSET,
						"description":  "etl-1740562885-856622",
						"job_name":     "etl-1740562885-856622",
						"display_name": "etl-1740562885-856622",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "DESSSSSS",
					"configuration": []map[string]interface{}{
						{
							"script":    "* ",
							"logstore":  "${alicloud_log_store.defaultDtiWb2.name}",
							"role_arn":  "acs:ram::1395894005720:role/aliyunlogetlrole",
							"from_time": "1",
							"to_time":   "1738394097",
							"sink": []map[string]interface{}{
								{
									"name":     "11111",
									"endpoint": "https://cn-wulanchabu.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test-logstore",
									"datasets": []string{
										"__UNNAMED__"},
									"role_arn": "acs:ram::1395894005720:role/aliyunlogetlrole",
								},
								{
									"name":     "22222",
									"endpoint": "https://cn-heyuan.log.aliyuncs.com",
									"project":  "tet",
									"logstore": "test",
									"datasets": []string{
										"rrr"},
									"role_arn": "acs:ram::13958940009000:role/aliyunlogetlrole",
								},
								{
									"name":     "33333",
									"endpoint": "https://cn-chengdu.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test",
									"datasets": []string{
										"ttt"},
									"role_arn": "acs:ram::13958940059998720:role/aliyunlogetlrole",
								},
								{
									"name":     "44444",
									"endpoint": "https://cn-guangzhou.log.aliyuncs.com",
									"project":  "test",
									"logstore": "test",
									"datasets": []string{
										"DDDD"},
									"role_arn": "acs:ram::1395894005868720:role/aliyunlogetlrole",
								},
							},
							"parameters": map[string]interface{}{
								"\"SK\"": "888",
							},
							"lang": "SPL",
						},
					},
					"display_name": "ESSSSSS",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description":  "DESSSSSS",
						"display_name": "ESSSSSS",
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

var AlicloudSlsEtlMap10469 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudSlsEtlBasicDependence10469(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_log_project" "defaultpYdUWa" {
  description = "terraform-etl-test-235"
  name        = "terraform-etl-test-695"
}

resource "alicloud_log_store" "defaultDtiWb2" {
  hot_ttl          = "8"
  retention_period = "30"
  shard_count      = "2"
  project          = alicloud_log_project.defaultpYdUWa.id
  name             = "test1"
}


`, name)
}

// Test Sls Etl. <<< Resource test cases, automatically generated.
