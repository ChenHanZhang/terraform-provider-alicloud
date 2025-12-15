package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test ESA WafRule. >>> Resource test cases, automatically generated.
// Case resource_wafrule_custom_test
func TestAccAliCloudESAWafRuleresource_wafrule_custom_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_custom_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_custom_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_custom",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"action":     "deny",
							"expression": "(http.host in {\\\"123.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "111",
						},
					},
					"site_version": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_custom",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"action":     "deny",
							"expression": "(http.host in {\\\"123.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "10",
											"code": "401",
										},
									},
								},
							},
							"name": "111",
						},
					},
					"site_version": "0",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"action":     "monitor",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"name":       "2222",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"action":     "js",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"name":       "2222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status":     "off",
							"action":     "captcha",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"name":       "2222",
						},
					},
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
				ImportStateVerifyIgnore: []string{"site_version"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_custom_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_custom_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_custom"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_anti_scan_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_anti_scan_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_anti_scan_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_anti_scan_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_anti_scan",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "http.request.session",
												},
											},
											"logic": "and",
										},
									},
									"ttl": "1800",
									"threshold": []map[string]interface{}{
										{
											"response_status": []map[string]interface{}{
												{
													"ratio": "70",
													"count": "50",
													"code":  "404",
												},
											},
											"request": "50",
										},
									},
									"interval": "10",
								},
							},
							"type": "http_directory_traversal",
						},
					},
					"shared": []map[string]interface{}{
						{
							"action":     "deny",
							"expression": "(ip.src eq 1.1.1.1)",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "tttt",
							"match": []map[string]interface{}{
								{
									"criteria": []map[string]interface{}{

										{
											"match_type": "ip.src",
										},

										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "ip.src",
												},

												{
													"criteria": []map[string]interface{}{

														{
															"match_type": "ip.src",
														},

														{
															"match_type": "ip.src",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"logic": "and",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_anti_scan",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "ip.src",
												},
											},
											"logic": "and",
										},
									},
									"ttl": "1801",
									"threshold": []map[string]interface{}{
										{
											"response_status": []map[string]interface{}{
												{
													"ratio": "71",
													"count": "51",
													"code":  "404",
												},
											},
											"request": "51",
										},
									},
									"interval": "11",
								},
							},
							"type": "http_directory_traversal",
						},
					},
					"shared": []map[string]interface{}{
						{
							"action":     "deny",
							"expression": "(ip.src eq 1.1.1.2)",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "tttt1",
							"match": []map[string]interface{}{
								{
									"criteria": []map[string]interface{}{

										{
											"match_type": "ip.src",
										},

										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "ip.src",
												},

												{
													"criteria": []map[string]interface{}{

														{
															"match_type": "ip.src",
														},

														{
															"match_type": "ip.src",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"logic": "and",
								},
							},
						},
					},
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
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_anti_scan_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_anti_scan_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_anti_scan"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_ratelimit_test
func TestAccAliCloudESAWafRuleresource_wafrule_ratelimit_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_ratelimit_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_ratelimit_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_ratelimit",
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"criteria": []map[string]interface{}{

														{
															"criteria": []map[string]interface{}{

																{
																	"match_type": "ip.src",
																},

																{
																	"match_type": "http.host",
																},

																{
																	"match_type": "ip.src",
																},
															},
															"logic": "and",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "true",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"criteria": []map[string]interface{}{

														{
															"criteria": []map[string]interface{}{

																{
																	"match_type": "http.host",
																},
															},
															"logic": "or",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "60",
									"threshold": []map[string]interface{}{
										{
											"traffic": "200",
											"request": "20",
										},
									},
									"interval": "60",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"126.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "2222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"criteria": []map[string]interface{}{

														{
															"criteria": []map[string]interface{}{

																{
																	"match_type": "http.host",
																},

																{
																	"match_type": "http.host",
																},

																{
																	"match_type": "http.host",
																},
															},
															"logic": "or",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "60",
									"threshold": []map[string]interface{}{
										{
											"traffic": "200",
											"request": "20",
										},
									},
									"interval": "60",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"126.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "2222",
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_ratelimit_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_ratelimit_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}



data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_ratelimit"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_ratelimit_test2
func TestAccAliCloudESAWafRuleresource_wafrule_ratelimit_test2(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_ratelimit_test2Map)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_ratelimit_test2BasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_ratelimit",
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"criteria": []map[string]interface{}{

														{
															"match_type": "ip.src",
														},

														{
															"match_type": "http.host",
														},

														{
															"match_type": "ip.src",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "true",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"criteria": []map[string]interface{}{

														{
															"match_type": "http.host",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"129.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_ratelimit_test2Map = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_ratelimit_test2BasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_ratelimit"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_ratelimit_test3
func TestAccAliCloudESAWafRuleresource_wafrule_ratelimit_test3(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_ratelimit_test3Map)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_ratelimit_test3BasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_ratelimit",
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "ip.src",
												},

												{
													"match_type": "ip.src",
												},

												{
													"match_type": "http.host",
												},
											},
											"logic": "or",
										},
									},
									"on_hit": "true",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "http.host",
												},

												{
													"match_type": "ip.src",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"126.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_ratelimit_test3Map = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_ratelimit_test3BasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_ratelimit"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_security_level_rule_test
func TestAccAliCloudESAWafRuleresource_wafrule_security_level_rule_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_security_level_rule_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_security_level_rule_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_security_level_rule",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"action":     "captcha",
							"expression": "(ip.src eq 1.1.1.2)",
							"security_level": []map[string]interface{}{
								{
									"value": "medium",
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_security_level_rule",
					"config": []map[string]interface{}{
						{
							"status":     "off",
							"action":     "captcha",
							"expression": "(ip.src eq 1.1.1.3)",
							"security_level": []map[string]interface{}{
								{
									"value": "high",
								},
							},
							"name": "12222",
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_security_level_rule_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_security_level_rule_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_security_level_rule"
  site_version = "0"
}


`, name)
}

// Case resource_wafrule_http_managed_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_managed_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_managed_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_managed_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_managed",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{

								{
									"protection_level": "2",
									"action":           "deny",
									"attack_type":      "11",
								},

								{
									"protection_level": "3",
									"action":           "deny",
									"attack_type":      "34",
								},

								{
									"protection_level": "-1",
									"action":           "monitor",
									"managed_rules": []map[string]interface{}{

										{
											"status": "on",
											"action": "deny",
											"id":     "900896",
										},

										{
											"status": "on",
											"action": "deny",
											"id":     "900897",
										},

										{
											"status": "on",
											"action": "deny",
											"id":     "900895",
										},
									},
									"attack_type": "35",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{

								{
									"protection_level": "1",
									"action":           "monitor",
									"attack_type":      "17",
								},

								{
									"protection_level": "2",
									"action":           "monitor",
									"attack_type":      "15",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{

								{
									"protection_level": "-1",
									"action":           "monitor",
									"managed_rules": []map[string]interface{}{

										{
											"status": "on",
											"action": "monitor",
											"id":     "900896",
										},

										{
											"status": "on",
											"action": "deny",
											"id":     "900897",
										},

										{
											"status": "off",
											"action": "deny",
											"id":     "900895",
										},
									},
									"attack_type": "17",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{

								{
									"protection_level": "-1",
									"action":           "monitor",
									"managed_rules": []map[string]interface{}{

										{
											"status": "on",
											"action": "deny",
											"id":     "900896",
										},

										{
											"status": "on",
											"action": "deny",
											"id":     "900897",
										},
									},
									"attack_type": "17",
								},
							},
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_http_managed_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_managed_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_managed"
  site_version = "0"
}


`, name)
}

// Case resource_wafrule_http_intelligence_crawler_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_intelligence_crawler_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_intelligence_crawler_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_intelligence_crawler_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_security_level_rule",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"action":     "captcha",
							"expression": "true",
							"security_level": []map[string]interface{}{
								{
									"value": "high",
								},
							},
							"name": "aaa",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"action":     "captcha",
							"expression": "true",
							"security_level": []map[string]interface{}{
								{
									"value": "low",
								},
							},
							"name": "aaa1",
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_http_intelligence_crawler_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_intelligence_crawler_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_security_level_rule"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_sigchl_weekly_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_sigchl_weekly_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_sigchl_weekly_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_sigchl_weekly_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "8",
									"weekly_periods": []map[string]interface{}{

										{
											"days": "3",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},

												{
													"start": "09:00:00",
													"end":   "10:00:00",
												},
											},
										},

										{
											"days": "4",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},

												{
													"start": "09:00:00",
													"end":   "10:00:00",
												},
											},
										},

										{
											"days": "5",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},

												{
													"start": "09:00:00",
													"end":   "10:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig",
								"driver",
								"replay",
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "web",
							"expression": "(ip.src eq 1.1.1.1)",
							"mode":       "automatic",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "9",
									"weekly_periods": []map[string]interface{}{

										{
											"days": "3",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},

										{
											"days": "4",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig",
								"replay",
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "web",
							"expression": "(ip.src eq 1.1.1.2)",
							"mode":       "automatic",
							"name":       "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "8",
									"weekly_periods": []map[string]interface{}{

										{
											"days": "3,4,5",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig",
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "web",
							"expression": "(ip.src eq 1.1.1.2)",
							"mode":       "automatic",
							"name":       "tttt1",
						},
					},
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
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_sigchl_weekly_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_sigchl_weekly_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_bot"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_whitelist_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_whitelist_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_whitelist_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_whitelist_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_whitelist",
					"config": []map[string]interface{}{
						{
							"actions": []map[string]interface{}{
								{
									"bypass": []map[string]interface{}{
										{
											"skip": "part",
											"custom_rules": []string{
												"11111",
												"2222",
												"34444",
											},
											"regular_rules": []string{
												"111111",
												"222222",
												"333333",
											},
											"regular_types": []string{
												"webshell",
												"rfi",
												"lfi",
											},
											"tags": []string{
												"http_bot",
												"http_custom",
												"http_security_level",
												"http_intelligent_cc",
												"http_managed",
											},
										},
									},
								},
							},
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "2222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"actions": []map[string]interface{}{
								{
									"bypass": []map[string]interface{}{
										{
											"skip": "all",
											"custom_rules": []string{
												"1111",
												"2222",
											},
											"regular_rules": []string{
												"111111",
												"222222",
											},
											"regular_types": []string{
												"webshell",
												"rfi",
											},
											"tags": []string{
												"http_security_level",
												"http_custom",
											},
										},
									},
								},
							},
							"expression": "(ip.src eq 1.1.1.2)",
							"name":       "4444",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "off",
							"actions": []map[string]interface{}{
								{
									"bypass": []map[string]interface{}{
										{
											"skip": "part",
											"custom_rules": []string{
												"1111",
												"2222",
											},
											"regular_rules": []string{
												"111111",
												"222222",
											},
											"regular_types": []string{
												"webshell",
												"rfi",
											},
											"tags": []string{
												"http_security_level",
												"http_custom",
											},
										},
									},
								},
							},
							"expression": "(ip.src eq 1.1.1.2)",
							"name":       "4444",
						},
					},
				}),
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

var AliCloudESAWafRuleresource_wafrule_http_whitelist_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_whitelist_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_whitelist"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_app_sdk_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_app_sdk_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_app_sdk_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_app_sdk_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"app_sdk": []map[string]interface{}{
								{
									"custom_sign": []map[string]interface{}{
										{
											"value": "222222",
											"key":   "cookie",
										},
									},
									"custom_sign_status": "on",
									"feature_abnormal": []string{
										"signInvalid",
										"debugged",
										"root",
										"simulator",
									},
								},
							},
							"type":   "http_app_sdk",
							"action": "monitor",
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"app_sdk": []map[string]interface{}{
								{
									"custom_sign_status": "off",
									"feature_abnormal": []string{
										"signInvalid",
									},
								},
							},
							"type":   "http_app_sdk",
							"action": "monitor",
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.2)",
							"name":       "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"app_sdk": []map[string]interface{}{
								{
									"custom_sign": []map[string]interface{}{
										{
											"value": "44444",
											"key":   "header",
										},
									},
									"custom_sign_status": "on",
									"feature_abnormal": []string{
										"signInvalid",
									},
								},
							},
							"type":   "http_app_sdk",
							"action": "monitor",
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.2)",
							"name":       "tttt1",
						},
					},
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
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_app_sdk_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_app_sdk_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_bot"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_scan_tools_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_scan_tools_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_scan_tools_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_scan_tools_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_anti_scan",
					"config": []map[string]interface{}{
						{
							"status":       "on",
							"type":         "http_scan_tools",
							"managed_list": "intelligence_scan_tools",
						},
					},
					"shared": []map[string]interface{}{
						{
							"action":     "deny",
							"expression": "(ip.src eq 1.1.1.1)",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_anti_scan",
					"config": []map[string]interface{}{
						{
							"status":       "off",
							"type":         "http_scan_tools",
							"managed_list": "intelligence_scan_tools",
						},
					},
					"shared": []map[string]interface{}{
						{
							"action":     "monitor",
							"expression": "(ip.src eq 2.2.2.2)",
							"name":       "vvv",
						},
					},
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
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_scan_tools_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_scan_tools_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_anti_scan"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_ip_access_rule_test
func TestAccAliCloudESAWafRuleresource_wafrule_ip_access_rule_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_ip_access_rule_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_ip_access_rule_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "ip_access_rule",
					"config": []map[string]interface{}{
						{
							"type":   "ip",
							"action": "deny",
							"value":  "1.1.1.1",
							"notes":  "test",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"type":   "ip",
							"action": "deny",
							"value":  "2.2.2.2",
							"notes":  "test1",
						},
					},
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

var AliCloudESAWafRuleresource_wafrule_ip_access_rule_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_ip_access_rule_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "ip_access_rule"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_high_frequency_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_high_frequency_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_high_frequency_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_high_frequency_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_anti_scan",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "ip.src",
												},
											},
											"logic": "and",
										},
									},
									"ttl": "1800",
									"threshold": []map[string]interface{}{
										{
											"managed_rules_blocked":  "20",
											"distinct_managed_rules": "2",
										},
									},
									"interval": "10",
								},
							},
							"type": "http_high_frequency",
						},
					},
					"shared": []map[string]interface{}{
						{
							"action":     "deny",
							"expression": "(ip.src eq 1.1.1.1)",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "tttt",
							"match": []map[string]interface{}{
								{
									"match_type": "ip.src",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_anti_scan",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{

												{
													"match_type": "ip.src",
												},
											},
											"logic": "and",
										},
									},
									"ttl": "1800",
									"threshold": []map[string]interface{}{
										{
											"managed_rules_blocked":  "30",
											"distinct_managed_rules": "3",
										},
									},
									"interval": "10",
								},
							},
							"type": "http_high_frequency",
						},
					},
					"shared": []map[string]interface{}{
						{
							"action":     "deny",
							"expression": "(ip.src eq 1.1.1.1)",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "tttt",
							"match": []map[string]interface{}{
								{
									"match_type": "ip.src",
								},
							},
						},
					},
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
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_high_frequency_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_high_frequency_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_anti_scan"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_sigchl_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_sigchl_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_sigchl_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_sigchl_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"periods": []map[string]interface{}{

										{
											"start": "2025-06-22T16:00:00Z",
											"end":   "2025-06-23T16:00:00Z",
										},
									},
									"scopes": "periods",
								},
							},
							"sigchl": []string{
								"sig",
								"driver",
								"replay",
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "web",
							"expression": "(ip.src eq 1.1.1.1)",
							"mode":       "automatic",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"periods": []map[string]interface{}{

										{
											"start": "2025-06-23T16:00:00Z",
											"end":   "2025-06-24T16:00:00Z",
										},

										{
											"start": "2025-06-25T16:00:00Z",
											"end":   "2025-06-26T16:00:00Z",
										},

										{
											"start": "2025-06-27T16:00:00Z",
											"end":   "2025-06-30T16:00:00Z",
										},
									},
									"scopes": "periods",
								},
							},
							"sigchl": []string{
								"sig",
								"replay",
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "web",
							"expression": "(ip.src eq 1.1.1.2)",
							"mode":       "automatic",
							"name":       "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"periods": []map[string]interface{}{

										{
											"start": "2025-07-23T16:00:00Z",
											"end":   "2025-07-29T16:00:00Z",
										},
									},
									"scopes": "periods",
								},
							},
							"sigchl": []string{
								"sig",
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "web",
							"expression": "(ip.src eq 1.1.1.2)",
							"mode":       "automatic",
							"name":       "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "9",
									"weekly_periods": []map[string]interface{}{

										{
											"days": "3",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},

										{
											"days": "4",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig",
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "9",
									"weekly_periods": []map[string]interface{}{

										{
											"days": "3",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},

										{
											"days": "4",
											"daily_periods": []map[string]interface{}{

												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},

												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig",
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_sigchl_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_sigchl_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_bot"
  site_version = "0"
}

`, name)
}

// Case resource_wafrule_http_app_package_test
func TestAccAliCloudESAWafRuleresource_wafrule_http_app_package_test(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AliCloudESAWafRuleresource_wafrule_http_app_package_testMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testacc%sESAWafRule%d", defaultRegionToTest, rand)

	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AliCloudESAWafRuleresource_wafrule_http_app_package_testBasicDependence)
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
					"site_id":    "${data.alicloud_esa_sites.default.sites.0.site_id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.default.ruleset_id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_app_package",
							"action": "monitor",
							"app_package": []map[string]interface{}{
								{
									"package_signs": []map[string]interface{}{

										{
											"sign": "222大三大四",
											"name": "asdsada",
										},

										{
											"sign": "22222",
											"name": "asdasdasd",
										},

										{
											"sign": "333333",
											"name": "awdadasdsadas",
										},
									},
								},
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"phase": "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_app_package",
							"action": "monitor",
							"app_package": []map[string]interface{}{
								{
									"package_signs": []map[string]interface{}{

										{
											"sign": "233123",
											"name": "asdsada2",
										},

										{
											"sign": "dasdasas",
											"name": "asdasdasd2",
										},
									},
								},
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"phase": "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_app_package",
							"action": "monitor",
							"app_package": []map[string]interface{}{
								{
									"package_signs": []map[string]interface{}{

										{
											"sign": "222",
											"name": "asdsada3",
										},
									},
								},
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "tttt",
						},
					},
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
				ImportStateVerifyIgnore: []string{"shared"},
			},
		},
	})
}

var AliCloudESAWafRuleresource_wafrule_http_app_package_testMap = map[string]string{
	"id": CHECKSET,
}

func AliCloudESAWafRuleresource_wafrule_http_app_package_testBasicDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


data "alicloud_esa_sites" "default" {
  plan_subscribe_type = "enterpriseplan"
}

resource "alicloud_esa_waf_ruleset" "default" {
  site_id      = data.alicloud_esa_sites.default.sites.0.site_id
  phase        = "http_bot"
  site_version = "0"
}

`, name)
}

// Test ESA WafRule. <<< Resource test cases, automatically generated.
// Test Esa WafRule. >>> Resource test cases, automatically generated.
// Case resource_wafrule_ip_access_rule_test 11998
func TestAccAliCloudEsaWafRule_basic11998(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap11998)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence11998)
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
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test7.id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test7.ruleset_id}",
					"phase":      "ip_access_rule",
					"config": []map[string]interface{}{
						{
							"type":   "ip",
							"action": "deny",
							"value":  "1.1.1.1",
							"notes":  "test",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"site_id":    CHECKSET,
						"ruleset_id": CHECKSET,
						"phase":      "ip_access_rule",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"type":   "ip",
							"action": "deny",
							"value":  "2.2.2.2",
							"notes":  "test1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap11998 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence11998(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test7" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test7" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test7.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test7" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test7.id
  phase   = "ip_access_rule"
}


`, name)
}

// Case resource_wafrule_ratelimit_test 11999
func TestAccAliCloudEsaWafRule_basic11999(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap11999)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence11999)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test8.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test8.id}",
					"phase":      "http_ratelimit",
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"criteria": []map[string]interface{}{
														{
															"criteria": []map[string]interface{}{
																{
																	"match_type": "ip.src",
																},
																{
																	"match_type": "http.host",
																},
																{
																	"match_type": "ip.src",
																},
															},
															"logic": "and",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "true",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_ratelimit",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"criteria": []map[string]interface{}{
														{
															"criteria": []map[string]interface{}{
																{
																	"match_type": "http.host",
																},
															},
															"logic": "or",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "60",
									"threshold": []map[string]interface{}{
										{
											"traffic": "200",
											"request": "20",
										},
									},
									"interval": "60",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"126.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "2222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"criteria": []map[string]interface{}{
														{
															"criteria": []map[string]interface{}{
																{
																	"match_type": "http.host",
																},
																{
																	"match_type": "http.host",
																},
																{
																	"match_type": "http.host",
																},
															},
															"logic": "or",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "60",
									"threshold": []map[string]interface{}{
										{
											"traffic": "200",
											"request": "20",
										},
									},
									"interval": "60",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"126.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "2222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap11999 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence11999(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test8" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test8" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test8.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test8" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test8.id
  phase   = "http_ratelimit"
}


`, name)
}

// Case resource_wafrule_http_managed_test 12000
func TestAccAliCloudEsaWafRule_basic12000(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12000)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12000)
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
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test3.id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test3.ruleset_id}",
					"phase":      "http_managed",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{
								{
									"protection_level": "2",
									"action":           "deny",
									"attack_type":      "11",
								},
								{
									"protection_level": "3",
									"action":           "deny",
									"attack_type":      "34",
								},
								{
									"protection_level": "-1",
									"action":           "monitor",
									"managed_rules": []map[string]interface{}{
										{
											"status": "on",
											"action": "deny",
											"id":     "900896",
										},
										{
											"status": "on",
											"action": "deny",
											"id":     "900897",
										},
										{
											"status": "on",
											"action": "deny",
											"id":     "900895",
										},
									},
									"attack_type": "35",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"site_id":    CHECKSET,
						"ruleset_id": CHECKSET,
						"phase":      "http_managed",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{
								{
									"protection_level": "1",
									"action":           "monitor",
									"attack_type":      "17",
								},
								{
									"protection_level": "2",
									"action":           "monitor",
									"attack_type":      "15",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{
								{
									"protection_level": "-1",
									"action":           "monitor",
									"managed_rules": []map[string]interface{}{
										{
											"status": "on",
											"action": "deny",
											"id":     "900896",
										},
										{
											"status": "on",
											"action": "deny",
											"id":     "900897",
										},
										{
											"status": "on",
											"action": "deny",
											"id":     "900895",
										},
									},
									"attack_type": "17",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"type":       "http_managed",
							"action":     "deny",
							"expression": "true",
							"managed_rulesets": []map[string]interface{}{
								{
									"protection_level": "-1",
									"action":           "monitor",
									"managed_rules": []map[string]interface{}{
										{
											"status": "on",
											"action": "deny",
											"id":     "900896",
										},
										{
											"status": "on",
											"action": "deny",
											"id":     "900897",
										},
									},
									"attack_type": "17",
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12000 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12000(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test3" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test3" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test3.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test3" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test3.id
  phase   = "http_managed"
}


`, name)
}

// Case resource_wafrule_http_app_sdk_test 12004
func TestAccAliCloudEsaWafRule_basic12004(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12004)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12004)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test11.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test11.id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"app_sdk": []map[string]interface{}{
								{
									"custom_sign": []map[string]interface{}{
										{
											"value": "222222",
											"key":   "cookie",
										},
									},
									"custom_sign_status": "on",
									"feature_abnormal": []string{
										"signInvalid", "debugged", "root", "simulator"},
								},
							},
							"type":   "http_app_sdk",
							"action": "monitor",
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_bot",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"app_sdk": []map[string]interface{}{
								{
									"custom_sign": []map[string]interface{}{
										{
											"value": "44444",
											"key":   "header",
										},
									},
									"custom_sign_status": "on",
									"feature_abnormal": []string{
										"signInvalid"},
								},
							},
							"type":   "http_app_sdk",
							"action": "monitor",
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.2)",
							"name":       "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12004 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12004(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test11" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test11" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test11.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test11" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test11.id
  phase   = "http_bot"
}


`, name)
}

// Case resource_wafrule_http_sigchl_weekly_test 12005
func TestAccAliCloudEsaWafRule_basic12005(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12005)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12005)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test13.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test13.id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "8",
									"weekly_periods": []map[string]interface{}{
										{
											"days": "3",
											"daily_periods": []map[string]interface{}{
												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
												{
													"start": "09:00:00",
													"end":   "10:00:00",
												},
											},
										},
										{
											"days": "4",
											"daily_periods": []map[string]interface{}{
												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
												{
													"start": "09:00:00",
													"end":   "10:00:00",
												},
											},
										},
										{
											"days": "5",
											"daily_periods": []map[string]interface{}{
												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
												{
													"start": "09:00:00",
													"end":   "10:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig", "driver", "replay"},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":        "web",
							"expression":    "(ip.src eq 1.1.1.1)",
							"mode":          "automatic",
							"cross_site_id": "763558725360816",
							"name":          "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_bot",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "9",
									"weekly_periods": []map[string]interface{}{
										{
											"days": "3",
											"daily_periods": []map[string]interface{}{
												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},
										{
											"days": "4",
											"daily_periods": []map[string]interface{}{
												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
												{
													"start": "06:00:00",
													"end":   "07:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig", "replay"},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":        "web",
							"expression":    "(ip.src eq 1.1.1.2)",
							"mode":          "automatic",
							"cross_site_id": "763558725360816",
							"name":          "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"scopes": "weekly",
									"zone":   "8",
									"weekly_periods": []map[string]interface{}{
										{
											"days": "3,4,5",
											"daily_periods": []map[string]interface{}{
												{
													"start": "02:00:00",
													"end":   "05:00:00",
												},
											},
										},
									},
								},
							},
							"sigchl": []string{
								"sig"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12005 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12005(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test13" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test13" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test13.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test13" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test13.id
  phase   = "http_bot"
}


`, name)
}

// Case resource_wafrule_http_sigchl_test 12007
func TestAccAliCloudEsaWafRule_basic12007(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12007)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12007)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test10.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test10.id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"periods": []map[string]interface{}{
										{
											"start": "2025-06-22T16:00:00Z",
											"end":   "2025-06-23T16:00:00Z",
										},
									},
									"scopes": "periods",
								},
							},
							"sigchl": []string{
								"sig", "driver", "replay"},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":        "web",
							"expression":    "(ip.src eq 1.1.1.1)",
							"mode":          "automatic",
							"cross_site_id": "763558725360816",
							"name":          "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_bot",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"periods": []map[string]interface{}{
										{
											"start": "2025-06-23T16:00:00Z",
											"end":   "2025-06-24T16:00:00Z",
										},
										{
											"start": "2025-06-25T16:00:00Z",
											"end":   "2025-06-26T16:00:00Z",
										},
										{
											"start": "2025-06-27T16:00:00Z",
											"end":   "2025-06-30T16:00:00Z",
										},
									},
									"scopes": "periods",
								},
							},
							"sigchl": []string{
								"sig", "replay"},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":        "web",
							"expression":    "(ip.src eq 1.1.1.2)",
							"mode":          "automatic",
							"cross_site_id": "763558725360816",
							"name":          "tttt1",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_sigchl",
							"action": "sigchl",
							"timer": []map[string]interface{}{
								{
									"periods": []map[string]interface{}{
										{
											"start": "2025-07-23T16:00:00Z",
											"end":   "2025-07-29T16:00:00Z",
										},
									},
									"scopes": "periods",
								},
							},
							"sigchl": []string{
								"sig"},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12007 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12007(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test10" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test10" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test10.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test10" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test10.id
  phase   = "http_bot"
}


`, name)
}

// Case resource_wafrule_ratelimit_test2 12008
func TestAccAliCloudEsaWafRule_basic12008(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12008)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12008)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test82.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test82.id}",
					"phase":      "http_ratelimit",
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"criteria": []map[string]interface{}{
														{
															"match_type": "ip.src",
														},
														{
															"match_type": "http.host",
														},
														{
															"match_type": "ip.src",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "true",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_ratelimit",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"criteria": []map[string]interface{}{
														{
															"match_type": "http.host",
														},
													},
													"logic": "and",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"129.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12008 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12008(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test82" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test82" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test82.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test82" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test82.id
  phase   = "http_ratelimit"
}


`, name)
}

// Case resource_wafrule_http_whitelist_test 12009
func TestAccAliCloudEsaWafRule_basic12009(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12009)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12009)
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
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test6.id}",
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test6.ruleset_id}",
					"phase":      "http_whitelist",
					"config": []map[string]interface{}{
						{
							"actions": []map[string]interface{}{
								{
									"bypass": []map[string]interface{}{
										{
											"skip": "part",
											"custom_rules": []string{
												"11111", "2222", "34444"},
											"regular_rules": []string{
												"111111", "222222", "333333"},
											"regular_types": []string{
												"webshell", "rfi", "lfi"},
											"tags": []string{
												"http_bot", "http_custom", "http_security_level", "http_intelligent_cc", "http_managed"},
										},
									},
								},
							},
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "2222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"site_id":    CHECKSET,
						"ruleset_id": CHECKSET,
						"phase":      "http_whitelist",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "off",
							"actions": []map[string]interface{}{
								{
									"bypass": []map[string]interface{}{
										{
											"skip": "part",
											"custom_rules": []string{
												"1111", "2222"},
											"regular_rules": []string{
												"111111", "222222"},
											"regular_types": []string{
												"webshell", "rfi"},
											"tags": []string{
												"http_security_level", "http_custom"},
										},
									},
								},
							},
							"expression": "(ip.src eq 1.1.1.2)",
							"name":       "4444",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12009 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12009(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test6" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test6" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test6.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test6" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test6.id
  phase   = "http_whitelist"
  name    = "2222"
}


`, name)
}

// Case resource_wafrule_ratelimit_test3 12010
func TestAccAliCloudEsaWafRule_basic12010(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12010)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12010)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test83.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test83.id}",
					"phase":      "http_ratelimit",
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"match_type": "ip.src",
												},
												{
													"match_type": "ip.src",
												},
												{
													"match_type": "http.host",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "true",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"125.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_ratelimit",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"rate_limit": []map[string]interface{}{
								{
									"characteristics": []map[string]interface{}{
										{
											"criteria": []map[string]interface{}{
												{
													"match_type": "http.host",
												},
												{
													"match_type": "ip.src",
												},
											},
											"logic": "and",
										},
									},
									"on_hit": "false",
									"ttl":    "300",
									"threshold": []map[string]interface{}{
										{
											"traffic": "100",
											"request": "10",
										},
									},
									"interval": "10",
								},
							},
							"action":     "deny",
							"expression": "(http.host in {\\\"126.maqi0415.top\\\"})",
							"actions": []map[string]interface{}{
								{
									"response": []map[string]interface{}{
										{
											"id":   "0",
											"code": "403",
										},
									},
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12010 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12010(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test83" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test83" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test83.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test83" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test83.id
  phase   = "http_ratelimit"
}


`, name)
}

// Case resource_wafrule_security_level_rule_test 12013
func TestAccAliCloudEsaWafRule_basic12013(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12013)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12013)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test14.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test14.id}",
					"phase":      "http_security_level_rule",
					"config": []map[string]interface{}{
						{
							"status":     "on",
							"action":     "captcha",
							"expression": "(ip.src eq 1.1.1.2)",
							"security_level": []map[string]interface{}{
								{
									"value": "medium",
								},
							},
							"name": "11111",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_security_level_rule",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status":     "off",
							"action":     "captcha",
							"expression": "(ip.src eq 1.1.1.3)",
							"security_level": []map[string]interface{}{
								{
									"value": "high",
								},
							},
							"name": "12222",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12013 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12013(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test14" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test14" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test14.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test14" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test14.id
  phase   = "http_security_level_rule"
}


`, name)
}

// Case resource_wafrule_http_app_package_test 12014
func TestAccAliCloudEsaWafRule_basic12014(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_waf_rule.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaWafRuleMap12014)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaWafRule")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaWafRuleBasicDependence12014)
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
					"ruleset_id": "${alicloud_esa_waf_ruleset.resource_WafRuleset_test12.ruleset_id}",
					"site_id":    "${alicloud_esa_site.resource_Site_waf_rule_test12.id}",
					"phase":      "http_bot",
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_app_package",
							"action": "monitor",
							"app_package": []map[string]interface{}{
								{
									"package_signs": []map[string]interface{}{
										{
											"sign": "222大三大四",
											"name": "asdsada",
										},
										{
											"sign": "22222",
											"name": "asdasdasd",
										},
										{
											"sign": "333333",
											"name": "awdadasdsadas",
										},
									},
								},
							},
						},
					},
					"shared": []map[string]interface{}{
						{
							"target":     "app",
							"expression": "(ip.src eq 1.1.1.1)",
							"name":       "tttt",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ruleset_id": CHECKSET,
						"site_id":    CHECKSET,
						"phase":      "http_bot",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_app_package",
							"action": "monitor",
							"app_package": []map[string]interface{}{
								{
									"package_signs": []map[string]interface{}{
										{
											"sign": "233123",
											"name": "asdsada2",
										},
										{
											"sign": "dasdasas",
											"name": "asdasdasd2",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"config": []map[string]interface{}{
						{
							"status": "on",
							"type":   "http_app_package",
							"action": "monitor",
							"app_package": []map[string]interface{}{
								{
									"package_signs": []map[string]interface{}{
										{
											"sign": "222",
											"name": "asdsada3",
										},
									},
								},
							},
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"shared", "site_version"},
			},
		},
	})
}

var AlicloudEsaWafRuleMap12014 = map[string]string{
	"waf_rule_id": CHECKSET,
}

func AlicloudEsaWafRuleBasicDependence12014(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_waf_rule_test12" {
  type         = "CNAME"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "domestic"
  auto_pay     = true
  plan_name    = "weizhaotest"
}

resource "alicloud_esa_site" "resource_Site_waf_rule_test12" {
  site_name   = "cname.chengkun.com"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_waf_rule_test12.id
  coverage    = "domestic"
  access_type = "CNAME"
}

resource "alicloud_esa_waf_ruleset" "resource_WafRuleset_test12" {
  site_id = alicloud_esa_site.resource_Site_waf_rule_test12.id
  phase   = "http_bot"
}


`, name)
}

// Test Esa WafRule. <<< Resource test cases, automatically generated.
