// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test SslCertificatesService Contact. >>> Resource test cases, automatically generated.
// Case contact_test 12965
func TestAccAliCloudSslCertificatesServiceContact_basic12965(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ssl_certificates_service_contact.default"
	ra := resourceAttrInit(resourceId, AlicloudSslCertificatesServiceContactMap12965)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SslCertificatesServiceServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSslCertificatesServiceContact")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccsslcertificatesservice%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSslCertificatesServiceContactBasicDependence12965)
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
					"email":  "testcontact@example.com",
					"idcard": "110101199001011234",
					"mobile": "13800138001",
					"name":   name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"email":  "testcontact@example.com",
						"idcard": CHECKSET,
						"mobile": CHECKSET,
						"name":   name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"email":  "testcontact-updated@example.com",
					"idcard": "110101199001015678",
					"mobile": "13800138002",
					"name":   name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"email":  "testcontact-updated@example.com",
						"idcard": CHECKSET,
						"mobile": CHECKSET,
						"name":   name + "_update",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"email", "mobile", "webhooks"},
			},
		},
	})
}

var AlicloudSslCertificatesServiceContactMap12965 = map[string]string{}

func AlicloudSslCertificatesServiceContactBasicDependence12965(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}


`, name)
}

// Test SslCertificatesService Contact. <<< Resource test cases, automatically generated.
