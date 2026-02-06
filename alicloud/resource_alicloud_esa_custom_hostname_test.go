// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// Test Esa CustomHostname. >>> Resource test cases, automatically generated.
// Case resource_CustomHostname_test 12471
func TestAccAliCloudEsaCustomHostname_basic12471(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_esa_custom_hostname.default"
	ra := resourceAttrInit(resourceId, AlicloudEsaCustomHostnameMap12471)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EsaServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEsaCustomHostname")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccesa%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEsaCustomHostnameBasicDependence12471)
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
					"site_id":     "${alicloud_esa_site.resource_Site_test_CustomHostname.id}",
					"private_key": "",
					"hostname":    "unittest.ialicdn.com",
					"cert_type":   "free",
					"record_id":   "${alicloud_esa_record.resource_Record_test_CustomHostname.id}",
					"ssl_flag":    "on",
					"cas_id":      "0",
					"cas_region":  "",
					"certificate": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"site_id":     CHECKSET,
						"private_key": "",
						"hostname":    "unittest.ialicdn.com",
						"cert_type":   "free",
						"record_id":   CHECKSET,
						"ssl_flag":    "on",
						"cas_id":      "0",
						"cas_region":  "",
						"certificate": "",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_key": "-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCF0HTDtc4buM4X\\nFpTZELsstgz5fMEAEm3dv4dMfhAlH66/sk9hVoy609Rp3nJS9TVPnrpGsqEgMl9X\\nez7mp1rOCDDIEMGHj4pyJV2Wu5ho++ifm1nmrYilGFq/A2ZkF8hZRIitQ81RBDK7\\ndHHqqqnoQ9huwvQDNf3rGDJZFWiuM2cMObdf5v+eY0lnfkt6yQZXpMynRScOEUSN\\nTOO/Zc3e0qrqY3U2ovb5p7kgbVYtcTmSLee685Qi+zvVj1mUXagGuvSHZ07BzKeV\\nKhHBgPjgKP4/te7pk7hSk9aiWvYUfejr9igJ4w7YXalJ9t39yCb+iSpXPNe0rCPE\\nxC+iEo23AgMBAAECggEAEEstrJb6nc2G6rDJNxUedXdFZSuXpJaZ4iJsKxg+hwaw\\nlI71s3iCck1Q1ANOEGPjNeqx6+HcVLtNeK19H0DJgcTli7bemc8UoImEN9Jn4ICr\\n9qNH8xq2RMQOaKvVT+LFdnkt20siOPc15jVrmZNmVO3N8M60P5/XZ0Tu/IHq2Ssk\\nEihM7wizxQfFoHe6c783yaIpGcDQllHVNqdq7yI4T0FgjiHWnAJmcIQzDhEZuB8l\\n3V2WBuOlmwYLvVNwgFlZ6j+t4bTKSu9lvLCDoFU7mLApWmNxySbz2H9oiIigdgsz\\nxwPP4cmLBBUivLQWAGGYyNilas/zudsrIyi2HiRzYQKBgQC5JbCNsgM0jzTNWO+Z\\nsMvQ3v1Xds6PjlZVEabz5Gjhpgt82HnWU6NzzEZfAvtaSo+wlr+mei/Vj4OK3/RG\\nCcQNVWpOY7vfzahMxVtf2KgKxLnmDvRWABEbz9dfQb8LMy3VTaxM2KI06OQfaqqy\\npD9UuHbHOHdchAkzCx4Es4JZ2wKBgQC5BdWK98AHsw/b56H6LcTkFq7QCw95oMJk\\nn9B9DgXLc+E8M3SyVQr7k7IUXTc4SUDQT+Il58jchHcW3cOyBj3pHAbcPwtq2zGM\\n2qR3n+8A6KBuuaYhjf5or8z8k4PjvJ+0lvzFU5j7YAHaZhJ3bdJwq+mjCpTEJYIO\\n7ynphmmoVQKBgQCPChdE4WqoiJr6quMxke1lCWIg2KDtN8JrJqAdfTGqY2YspfI7\\nK5L0O7WCYgNe1ov8mfqm229pRt8Rw8Qs+A6HXp4qwdK8LV8BMNhPTEtHRoV4v8T1\\nTEfrfL+f0GOQe5HFYaTkGdn3lpXnz7jhGxPymDDr6SwORdP1o3klmKn0UQKBgQCi\\njlCmTrMlUJnvX5v8/LdOmesul3kaKDaCR7LaHDbVIFNUG/U8NOF4VLQRljBMwHc0\\nTern2LOtoqgqq94ii79bDies26sBT+FB2lrDSv2mM7u4bF1cf690dHhJtKUcsf0y\\nobElASSYPiqCokk5KoDIQxH7D/HSPw38Zv3Q0SoFHQKBgDDeasIayiJx0Ae+MXnV\\nFMG2lX61Veh78nHAflEp3Dl0kflWw63lYj5c+8Su9ykx5k2NHzD0bvCOOp/ehJVp\\nQY1y2D0tEJ8J9cd8wAHg1tpKzQcLm0w7jyzUgP/FsLP8P+3WgtXDTiV2dpp0sDwz\\nAKSE4Cu/wacZXyBSk6bWKRmK\\n-----END PRIVATE KEY-----\\n",
					"cert_type":   "upload",
					"certificate": "-----BEGIN CERTIFICATE-----\\nMIIDZzCCAk+gAwIBAgIQAOcTu/gvR354urqG6erS5TANBgkqhkiG9w0BAQsFADAf\\nMR0wGwYDVQQDDBR1bml0dGVzdC5pYWxpY2RuLmNvbTAeFw0yNjAxMDcxMTI1MDda\\nFw0zNjAxMDUxMTI1MDdaMB8xHTAbBgNVBAMMFHVuaXR0ZXN0LmlhbGljZG4uY29t\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhdB0w7XOG7jOFxaU2RC7\\nLLYM+XzBABJt3b+HTH4QJR+uv7JPYVaMutPUad5yUvU1T566RrKhIDJfV3s+5qda\\nzggwyBDBh4+KciVdlruYaPvon5tZ5q2IpRhavwNmZBfIWUSIrUPNUQQyu3Rx6qqp\\n6EPYbsL0AzX96xgyWRVorjNnDDm3X+b/nmNJZ35LeskGV6TMp0UnDhFEjUzjv2XN\\n3tKq6mN1NqL2+ae5IG1WLXE5ki3nuvOUIvs71Y9ZlF2oBrr0h2dOwcynlSoRwYD4\\n4Cj+P7Xu6ZO4UpPWolr2FH3o6/YoCeMO2F2pSfbd/cgm/okqVzzXtKwjxMQvohKN\\ntwIDAQABo4GeMIGbMB0GA1UdDgQWBBSHhDe7PARENpSQFud9jkeEfMcDTzAOBgNV\\nHQ8BAf8EBAMCBLAwDAYDVR0TAQH/BAIwADA7BgNVHSUENDAyBggrBgEFBQcDAgYI\\nKwYBBQUHAwEGCCsGAQUFBwMDBggrBgEFBQcDBAYIKwYBBQUHAwgwHwYDVR0jBBgw\\nFoAUh4Q3uzwERDaUkBbnfY5HhHzHA08wDQYJKoZIhvcNAQELBQADggEBACAq5nic\\noXBRq/1SoPrcYHAv4K8HlLqOltY6Zscuxetk8F6JAy73IVd6S3gVZQ/yZRDbTcrl\\ns64myqFSMbkTz/ACEg8kZIv5pof368ImTua7dTsWVXOFQ2KJ3H6VfbNGZXzNZx+0\\nHyM+ShTZOEO3cwwEfpez6pLYCMDomHRer8DCyJZ8ayMYA80tpD2m5IuXxsrvyp7x\\nA9XsX/p27oeFdqiW2Z6zUxPOJTXXbq00T3wZnFjnZ7fJeAa8BSJrLf+TEt6Y/fHc\\nu0BKm9RB2TxnzCBBF1mhV/5lLT2gP8nASfopPevmdGhnhiRBUqVgnzPPbhpaWfix\\nNXNdOxePSKanJE0=\\n-----END CERTIFICATE-----\\n",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCF0HTDtc4buM4X\nFpTZELsstgz5fMEAEm3dv4dMfhAlH66/sk9hVoy609Rp3nJS9TVPnrpGsqEgMl9X\nez7mp1rOCDDIEMGHj4pyJV2Wu5ho++ifm1nmrYilGFq/A2ZkF8hZRIitQ81RBDK7\ndHHqqqnoQ9huwvQDNf3rGDJZFWiuM2cMObdf5v+eY0lnfkt6yQZXpMynRScOEUSN\nTOO/Zc3e0qrqY3U2ovb5p7kgbVYtcTmSLee685Qi+zvVj1mUXagGuvSHZ07BzKeV\nKhHBgPjgKP4/te7pk7hSk9aiWvYUfejr9igJ4w7YXalJ9t39yCb+iSpXPNe0rCPE\nxC+iEo23AgMBAAECggEAEEstrJb6nc2G6rDJNxUedXdFZSuXpJaZ4iJsKxg+hwaw\nlI71s3iCck1Q1ANOEGPjNeqx6+HcVLtNeK19H0DJgcTli7bemc8UoImEN9Jn4ICr\n9qNH8xq2RMQOaKvVT+LFdnkt20siOPc15jVrmZNmVO3N8M60P5/XZ0Tu/IHq2Ssk\nEihM7wizxQfFoHe6c783yaIpGcDQllHVNqdq7yI4T0FgjiHWnAJmcIQzDhEZuB8l\n3V2WBuOlmwYLvVNwgFlZ6j+t4bTKSu9lvLCDoFU7mLApWmNxySbz2H9oiIigdgsz\nxwPP4cmLBBUivLQWAGGYyNilas/zudsrIyi2HiRzYQKBgQC5JbCNsgM0jzTNWO+Z\nsMvQ3v1Xds6PjlZVEabz5Gjhpgt82HnWU6NzzEZfAvtaSo+wlr+mei/Vj4OK3/RG\nCcQNVWpOY7vfzahMxVtf2KgKxLnmDvRWABEbz9dfQb8LMy3VTaxM2KI06OQfaqqy\npD9UuHbHOHdchAkzCx4Es4JZ2wKBgQC5BdWK98AHsw/b56H6LcTkFq7QCw95oMJk\nn9B9DgXLc+E8M3SyVQr7k7IUXTc4SUDQT+Il58jchHcW3cOyBj3pHAbcPwtq2zGM\n2qR3n+8A6KBuuaYhjf5or8z8k4PjvJ+0lvzFU5j7YAHaZhJ3bdJwq+mjCpTEJYIO\n7ynphmmoVQKBgQCPChdE4WqoiJr6quMxke1lCWIg2KDtN8JrJqAdfTGqY2YspfI7\nK5L0O7WCYgNe1ov8mfqm229pRt8Rw8Qs+A6HXp4qwdK8LV8BMNhPTEtHRoV4v8T1\nTEfrfL+f0GOQe5HFYaTkGdn3lpXnz7jhGxPymDDr6SwORdP1o3klmKn0UQKBgQCi\njlCmTrMlUJnvX5v8/LdOmesul3kaKDaCR7LaHDbVIFNUG/U8NOF4VLQRljBMwHc0\nTern2LOtoqgqq94ii79bDies26sBT+FB2lrDSv2mM7u4bF1cf690dHhJtKUcsf0y\nobElASSYPiqCokk5KoDIQxH7D/HSPw38Zv3Q0SoFHQKBgDDeasIayiJx0Ae+MXnV\nFMG2lX61Veh78nHAflEp3Dl0kflWw63lYj5c+8Su9ykx5k2NHzD0bvCOOp/ehJVp\nQY1y2D0tEJ8J9cd8wAHg1tpKzQcLm0w7jyzUgP/FsLP8P+3WgtXDTiV2dpp0sDwz\nAKSE4Cu/wacZXyBSk6bWKRmK\n-----END PRIVATE KEY-----\n",
						"cert_type":   "upload",
						"certificate": "-----BEGIN CERTIFICATE-----\nMIIDZzCCAk+gAwIBAgIQAOcTu/gvR354urqG6erS5TANBgkqhkiG9w0BAQsFADAf\nMR0wGwYDVQQDDBR1bml0dGVzdC5pYWxpY2RuLmNvbTAeFw0yNjAxMDcxMTI1MDda\nFw0zNjAxMDUxMTI1MDdaMB8xHTAbBgNVBAMMFHVuaXR0ZXN0LmlhbGljZG4uY29t\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhdB0w7XOG7jOFxaU2RC7\nLLYM+XzBABJt3b+HTH4QJR+uv7JPYVaMutPUad5yUvU1T566RrKhIDJfV3s+5qda\nzggwyBDBh4+KciVdlruYaPvon5tZ5q2IpRhavwNmZBfIWUSIrUPNUQQyu3Rx6qqp\n6EPYbsL0AzX96xgyWRVorjNnDDm3X+b/nmNJZ35LeskGV6TMp0UnDhFEjUzjv2XN\n3tKq6mN1NqL2+ae5IG1WLXE5ki3nuvOUIvs71Y9ZlF2oBrr0h2dOwcynlSoRwYD4\n4Cj+P7Xu6ZO4UpPWolr2FH3o6/YoCeMO2F2pSfbd/cgm/okqVzzXtKwjxMQvohKN\ntwIDAQABo4GeMIGbMB0GA1UdDgQWBBSHhDe7PARENpSQFud9jkeEfMcDTzAOBgNV\nHQ8BAf8EBAMCBLAwDAYDVR0TAQH/BAIwADA7BgNVHSUENDAyBggrBgEFBQcDAgYI\nKwYBBQUHAwEGCCsGAQUFBwMDBggrBgEFBQcDBAYIKwYBBQUHAwgwHwYDVR0jBBgw\nFoAUh4Q3uzwERDaUkBbnfY5HhHzHA08wDQYJKoZIhvcNAQELBQADggEBACAq5nic\noXBRq/1SoPrcYHAv4K8HlLqOltY6Zscuxetk8F6JAy73IVd6S3gVZQ/yZRDbTcrl\ns64myqFSMbkTz/ACEg8kZIv5pof368ImTua7dTsWVXOFQ2KJ3H6VfbNGZXzNZx+0\nHyM+ShTZOEO3cwwEfpez6pLYCMDomHRer8DCyJZ8ayMYA80tpD2m5IuXxsrvyp7x\nA9XsX/p27oeFdqiW2Z6zUxPOJTXXbq00T3wZnFjnZ7fJeAa8BSJrLf+TEt6Y/fHc\nu0BKm9RB2TxnzCBBF1mhV/5lLT2gP8nASfopPevmdGhnhiRBUqVgnzPPbhpaWfix\nNXNdOxePSKanJE0=\n-----END CERTIFICATE-----\n",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ssl_flag": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ssl_flag": "off",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"cert_type":  "cas",
					"ssl_flag":   "on",
					"cas_id":     "22636994",
					"cas_region": "cn-hangzhou",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"cert_type":  "cas",
						"ssl_flag":   "on",
						"cas_id":     "22636994",
						"cas_region": "cn-hangzhou",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"private_key": "",
					"cert_type":   "free",
					"cas_id":      "0",
					"cas_region":  "",
					"certificate": "",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"private_key": "",
						"cert_type":   "free",
						"cas_id":      "0",
						"cas_region":  "",
						"certificate": "",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"ssl_flag": "off",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"ssl_flag": "off",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cas_region"},
			},
		},
	})
}

var AlicloudEsaCustomHostnameMap12471 = map[string]string{
	"status":      CHECKSET,
	"create_time": CHECKSET,
}

func AlicloudEsaCustomHostnameBasicDependence12471(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

resource "alicloud_esa_rate_plan_instance" "resource_RatePlanInstance_test_CustomHostname" {
  type         = "NS"
  auto_renew   = false
  period       = "1"
  payment_type = "Subscription"
  coverage     = "overseas"
  auto_pay     = true
  plan_name    = "high"
}

resource "alicloud_esa_site" "resource_Site_test_CustomHostname" {
  site_name   = "gositecdn.cn"
  instance_id = alicloud_esa_rate_plan_instance.resource_RatePlanInstance_test_CustomHostname.id
  coverage    = "overseas"
  access_type = "NS"
}

resource "alicloud_esa_record" "resource_Record_test_CustomHostname" {
  record_name = "www.gositecdn.cn"
  comment     = "This is a remark"
  proxied     = true
  site_id     = alicloud_esa_site.resource_Site_test_CustomHostname.id
  record_type = "CNAME"
  source_type = "S3"
  data {
    value = "www.idltestr.com"
  }
  biz_name    = "api"
  host_policy = "follow_hostname"
  ttl         = "100"
  auth_conf {
    secret_key = "hijklmnhijklmnhijklmnhijklmn"
    version    = "v4"
    region     = "us-east-1"
    auth_type  = "private"
    access_key = "abcdefgabcdefgabcdefgabcdefg"
  }
}


`, name)
}

// Test Esa CustomHostname. <<< Resource test cases, automatically generated.
