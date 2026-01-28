package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test SslCertificatesService PcaCertificate. >>> Resource test cases, automatically generated.
// Case PcaCertificate_sub 12329
func TestAccAliCloudSslCertificatesServicePcaCertificate_basic12329(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ssl_certificates_service_pca_certificate.default"
	ra := resourceAttrInit(resourceId, AlicloudSslCertificatesServicePcaCertificateMap12329)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SslCertificatesServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSslCertificatesServicePcaCertificate")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsslcertificatesservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSslCertificatesServicePcaCertificateBasicDependence12329)
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
					"organization":        "a",
					"years":               "1",
					"locality":            "a",
					"organization_unit":   "a",
					"state":               "a",
					"country_code":        "cn",
					"common_name":         "cbc.certqa.cn",
					"algorithm":           "RSA_2048",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"parent_identifier":   "${alicloud_ssl_certificates_service_pca_certificate.创建根CA.id}",
					"enable_crl":          "true",
					"crl_day":             "1",
					"path_len_constraint": "0",
					"extended_key_usages": []string{
						"serverAuth"},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"organization":          "a",
						"years":                 "1",
						"locality":              "a",
						"organization_unit":     "a",
						"state":                 "a",
						"country_code":          "cn",
						"common_name":           "cbc.certqa.cn",
						"algorithm":             "RSA_2048",
						"resource_group_id":     CHECKSET,
						"parent_identifier":     CHECKSET,
						"enable_crl":            "true",
						"crl_day":               "1",
						"path_len_constraint":   "0",
						"extended_key_usages.#": "1",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"alias_name":        "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"alias_name":        "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"alias_name":        "testupdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"alias_name":        "testupdate",
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
				ImportStateVerifyIgnore: []string{"alias_name", "enable_crl", "extended_key_usages", "path_len_constraint"},
			},
		},
	})
}

var AlicloudSslCertificatesServicePcaCertificateMap12329 = map[string]string{
	"status":           CHECKSET,
	"certificate_type": CHECKSET,
}

func AlicloudSslCertificatesServicePcaCertificateBasicDependence12329(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_ssl_certificates_service_pca_certificate" "创建根CA" {
  organization      = "a"
  years             = "1"
  locality          = "a"
  organization_unit = "a"
  state             = "a"
  country_code      = "cn"
  common_name       = "cbc.certqa.cn"
  algorithm         = "RSA_2048"
  certificate_type  = "ROOT"
}


`, name)
}

// Case PcaCertificate_root 12327
func TestAccAliCloudSslCertificatesServicePcaCertificate_basic12327(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ssl_certificates_service_pca_certificate.default"
	ra := resourceAttrInit(resourceId, AlicloudSslCertificatesServicePcaCertificateMap12327)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SslCertificatesServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSslCertificatesServicePcaCertificate")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsslcertificatesservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSslCertificatesServicePcaCertificateBasicDependence12327)
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
					"organization":      "a",
					"years":             "1",
					"locality":          "a",
					"organization_unit": "a",
					"state":             "a",
					"country_code":      "cn",
					"common_name":       "cbc.certqa.cn",
					"algorithm":         "RSA_2048",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"organization":      "a",
						"years":             "1",
						"locality":          "a",
						"organization_unit": "a",
						"state":             "a",
						"country_code":      "cn",
						"common_name":       "cbc.certqa.cn",
						"algorithm":         "RSA_2048",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"alias_name":        "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"alias_name":        "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"alias_name":        "testupdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"alias_name":        "testupdate",
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
				ImportStateVerifyIgnore: []string{"alias_name", "enable_crl", "extended_key_usages", "path_len_constraint"},
			},
		},
	})
}

var AlicloudSslCertificatesServicePcaCertificateMap12327 = map[string]string{
	"status":           CHECKSET,
	"certificate_type": CHECKSET,
}

func AlicloudSslCertificatesServicePcaCertificateBasicDependence12327(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case PcaCertificate 11010
func TestAccAliCloudSslCertificatesServicePcaCertificate_basic11010(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ssl_certificates_service_pca_certificate.default"
	ra := resourceAttrInit(resourceId, AlicloudSslCertificatesServicePcaCertificateMap11010)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SslCertificatesServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSslCertificatesServicePcaCertificate")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsslcertificatesservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSslCertificatesServicePcaCertificateBasicDependence11010)
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
					"organization":      "a",
					"years":             "1",
					"locality":          "a",
					"organization_unit": "a",
					"state":             "a",
					"country_code":      "cn",
					"common_name":       "cbc.certqa.cn",
					"algorithm":         "RSA_2048",
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"organization":      "a",
						"years":             "1",
						"locality":          "a",
						"organization_unit": "a",
						"state":             "a",
						"country_code":      "cn",
						"common_name":       "cbc.certqa.cn",
						"algorithm":         "RSA_2048",
						"resource_group_id": CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"alias_name":        "test",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"alias_name":        "test",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"alias_name":        "testupdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id": CHECKSET,
						"alias_name":        "testupdate",
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
				ImportStateVerifyIgnore: []string{"alias_name", "enable_crl", "extended_key_usages", "path_len_constraint"},
			},
		},
	})
}

var AlicloudSslCertificatesServicePcaCertificateMap11010 = map[string]string{
	"status":           CHECKSET,
	"certificate_type": CHECKSET,
}

func AlicloudSslCertificatesServicePcaCertificateBasicDependence11010(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Test SslCertificatesService PcaCertificate. <<< Resource test cases, automatically generated.
