// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test SslCertificatesService Instance. >>> Resource test cases, automatically generated.
// Case instance_test 12968
func TestAccAliCloudSslCertificatesServiceInstance_basic12968(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ssl_certificates_service_instance.default"
	ra := resourceAttrInit(resourceId, AlicloudSslCertificatesServiceInstanceMap12968)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SslCertificatesServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSslCertificatesServiceInstance")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsslcertificatesservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSslCertificatesServiceInstanceBasicDependence12968)
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
					"contact_id_list": []string{},
					"certificate_id":  "25747924",
					"product_type":    "cas_dv_public_cn",
					"period":          "3",
					"parameter": []map[string]interface{}{
						{
							"value": "ss.dv.t",
							"code":  "fullSpec",
						},
						{
							"value": "1",
							"code":  "fullDomainCount",
						},
						{
							"value": "testCert_product",
							"code":  "product",
						},
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"certificate_id": "25747924",
						"product_type":   "cas_dv_public_cn",
						"period":         "3",
						"parameter.#":    "3",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"contact_id_list": []string{
						"1354885", "1393155", "1401557"},
					"certificate_id":      "25964769",
					"csr":                 "-----BEGIN CERTIFICATE REQUEST-----\\nMIICrDCCAZQCAQAwZzEaMBgGA1UEAxMRd3d3LmNiYy5jZXJ0cWEuY24xETAPBgNV\\nBAgTCFNoYW5naGFpMREwDwYDVQQHEwhTaGFuZ2hhaTELMAkGA1UECxMCSVQxCTAH\\nBgNVBAoTADELMAkGA1UEBhMCQ04wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\\nAoIBAQC1BNVlWy4xL0nudda/MqiLGcJYYWGrtBjX7rS6niW2Qm+C6/pdEzAWIevt\\nhjI6B1lNLUOtVytEmv7VHka42ESPXbwiOZJHW3A7Y/TTvcrp+INU//w93g55tN5h\\nwUfbyrubDsbmYKvA31JsmsAGzEH/2MGyN6BkXY1cLSe4G0N37dQlkIL+edGMfKvK\\njrHLcFY+UKkIzele1qGdrHs+ROQsK6Wufu76lMLjYjbNsGGQ4iQzFfPx8VvLcE2p\\nlaqR4xJ/64GpaBFVVSJ23v3GDbuYYnpVizacEJ1dyeL5BPmJjyjYX89QfI3w2ng0\\nO8hvUK5cJwT7r579HgoKYlrP/m3HAgMBAAGgADANBgkqhkiG9w0BAQsFAAOCAQEA\\nrpMzwWSAKa88Hlo5sBUu7gXAxrYhONQpavnOXFXSiU0WMTfmSU+1KPzoP+W9LjWd\\niLbdwQMur6r/yEBccDpUzFbDcgVq1dfjy9kwRb5Z2izVqyrOxYBvQ8KG53MRZW7J\\noxZmkO3Vzs8FjEGkRYezW/e8vWwl8izjiP8bWQgj6ZupmbHKKOy/KjLPYnvg6ELJ\\nWNXreNbNzTPgvYQ6ZVzxcAhQooDQWjYuv0TJFd7bPgILRWPK9P3M/8O+IUsbFQWJ\\n+k//9kDjtlCmewFXyWETdUpJjiUZcWTT5QnprDn6ZvoMwXqZ5GUBRLvP/WyKwBs7\\nS7uDNC1wyCRx50fxIBKs2Q==\\n-----END CERTIFICATE REQUEST-----",
					"resource_group_id":   "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"company_id":          "53515",
					"validation_method":   "DNS-TEST",
					"city":                "Shanghai",
					"auto_reissue":        "disable",
					"province":            "Shanghai",
					"key_algorithm":       "RSA2048",
					"generate_csr_method": "test",
					"certificate_name":    "cert-2p4joj",
					"country_code":        "CN_UPDATE",
					"domain":              "www.cbc.certqa.test.cn",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"contact_id_list.#":   "3",
						"certificate_id":      "25964769",
						"csr":                 "-----BEGIN CERTIFICATE REQUEST-----\nMIICrDCCAZQCAQAwZzEaMBgGA1UEAxMRd3d3LmNiYy5jZXJ0cWEuY24xETAPBgNV\nBAgTCFNoYW5naGFpMREwDwYDVQQHEwhTaGFuZ2hhaTELMAkGA1UECxMCSVQxCTAH\nBgNVBAoTADELMAkGA1UEBhMCQ04wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK\nAoIBAQC1BNVlWy4xL0nudda/MqiLGcJYYWGrtBjX7rS6niW2Qm+C6/pdEzAWIevt\nhjI6B1lNLUOtVytEmv7VHka42ESPXbwiOZJHW3A7Y/TTvcrp+INU//w93g55tN5h\nwUfbyrubDsbmYKvA31JsmsAGzEH/2MGyN6BkXY1cLSe4G0N37dQlkIL+edGMfKvK\njrHLcFY+UKkIzele1qGdrHs+ROQsK6Wufu76lMLjYjbNsGGQ4iQzFfPx8VvLcE2p\nlaqR4xJ/64GpaBFVVSJ23v3GDbuYYnpVizacEJ1dyeL5BPmJjyjYX89QfI3w2ng0\nO8hvUK5cJwT7r579HgoKYlrP/m3HAgMBAAGgADANBgkqhkiG9w0BAQsFAAOCAQEA\nrpMzwWSAKa88Hlo5sBUu7gXAxrYhONQpavnOXFXSiU0WMTfmSU+1KPzoP+W9LjWd\niLbdwQMur6r/yEBccDpUzFbDcgVq1dfjy9kwRb5Z2izVqyrOxYBvQ8KG53MRZW7J\noxZmkO3Vzs8FjEGkRYezW/e8vWwl8izjiP8bWQgj6ZupmbHKKOy/KjLPYnvg6ELJ\nWNXreNbNzTPgvYQ6ZVzxcAhQooDQWjYuv0TJFd7bPgILRWPK9P3M/8O+IUsbFQWJ\n+k//9kDjtlCmewFXyWETdUpJjiUZcWTT5QnprDn6ZvoMwXqZ5GUBRLvP/WyKwBs7\nS7uDNC1wyCRx50fxIBKs2Q==\n-----END CERTIFICATE REQUEST-----",
						"resource_group_id":   CHECKSET,
						"company_id":          CHECKSET,
						"validation_method":   "DNS-TEST",
						"city":                "Shanghai",
						"auto_reissue":        "disable",
						"province":            "Shanghai",
						"key_algorithm":       "RSA2048",
						"generate_csr_method": "test",
						"certificate_name":    "cert-2p4joj",
						"country_code":        "CN_UPDATE",
						"domain":              "www.cbc.certqa.test.cn",
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
				ImportStateVerifyIgnore: []string{"auto_reissue", "certificate_id", "certificate_name", "city", "company_id", "contact_id_list", "country_code", "csr", "domain", "generate_csr_method", "key_algorithm", "parameter", "period", "pricing_cycle", "product_type", "province", "resource_group_id", "tags", "validation_method"},
			},
		},
	})
}

var AlicloudSslCertificatesServiceInstanceMap12968 = map[string]string{
	"instance_type":      CHECKSET,
	"status":             CHECKSET,
	"certificate_status": CHECKSET,
	"brand":              CHECKSET,
	"certificate_type":   CHECKSET,
}

func AlicloudSslCertificatesServiceInstanceBasicDependence12968(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test SslCertificatesService Instance. <<< Resource test cases, automatically generated.
