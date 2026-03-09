package alicloud

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("alicloud_slb_server_certificate", &resource.Sweeper{
		Name: "alicloud_slb_server_certificate",
		F:    testSweepSlbServerCertificate,
	})
}

func testSweepSlbServerCertificate(region string) error {
	rawClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("error getting Alicloud client: %s", err)
	}
	client := rawClient.(*connectivity.AliyunClient)

	prefixes := []string{
		"tf-testAcc",
		"tf_testAcc",
	}

	req := slb.CreateDescribeServerCertificatesRequest()
	req.RegionId = client.RegionId
	raw, err := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
		return slbClient.DescribeServerCertificates(req)
	})
	if err != nil {
		return err
	}
	resp, _ := raw.(*slb.DescribeServerCertificatesResponse)
	for _, serverCertificate := range resp.ServerCertificates.ServerCertificate {
		name := serverCertificate.ServerCertificateName
		id := serverCertificate.ServerCertificateId

		skip := true
		for _, prefix := range prefixes {
			if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
				skip = false
				break
			}
		}
		if skip {
			log.Printf("[INFO] Skipping Slb Server Certificate: %s (%s)", name, id)
			continue
		}
		log.Printf("[INFO] Deleting Slb Server Certificate : %s (%s)", name, id)

		req := slb.CreateDeleteServerCertificateRequest()
		req.ServerCertificateId = id
		_, error := client.WithSlbClient(func(slbClient *slb.Client) (interface{}, error) {
			return slbClient.DeleteServerCertificate(req)
		})
		if error != nil {
			log.Printf("[ERROR] Failed to delete Slb Server Certificate (%s (%s)): %s", name, id, err)
		}
	}
	return nil
}

// Test Slb ServerCertificate. >>> Resource test cases, automatically generated.
// Case 全生命周期 6748
func TestAccAliCloudSlbServerCertificate_basic6748(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_slb_server_certificate.default"
	ra := resourceAttrInit(resourceId, AlicloudSlbServerCertificateMap6748)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSlbServerCertificate")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccslb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSlbServerCertificateBasicDependence6748)
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
					"server_certificate":      "-----BEGIN CERTIFICATE-----\\nMIIDUjCCAjoCCQDYCjnX6QH2NTANBgkqhkiG9w0BAQUFADBrMQswCQYDVQQGEwJD\\nTjEQMA4GA1UECAwHQmVpSmluZzEQMA4GA1UEBwwHSGFpRGlhbjESMBAGA1UECgwJ\\nTWVuVG91R291MQswCQYDVQQLDAJJVDEXMBUGA1UEAwwOKi5kengudGVzdC5jb20w\\nHhcNMjMwODMwMDgxMTQ4WhcNMjQwODI5MDgxMTQ4WjBrMQswCQYDVQQGEwJDTjEQ\\nMA4GA1UECAwHQmVpSmluZzEQMA4GA1UEBwwHSGFpRGlhbjESMBAGA1UECgwJTWVu\\nVG91R291MQswCQYDVQQLDAJJVDEXMBUGA1UEAwwOKi5kengudGVzdC5jb20wggEi\\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCawIxT1owaUNlR6jmbUtFDtiqM\\nxGjk29DTjsXjHqY1FtZJRjlmNRC3VlcG7iK28IdRjkBz+da3bUVDE42Fg0daU8n2\\n+ls6ZGDdrHZpnBVw//rBuSiQZKrvbcswOLXfpQK5U0HiHgZA/HMJcISbp3AZ8vqa\\nnzQj1ly7Bpp4kwtMjO4JY0KXE4IjkDQmW2M+A8VbrvmP8xJ87Uy15uRqntKUdEV1\\ni6W1SD2TkAzHW8KfflHEYxRUCIFD5uoo/YxHDzFHeI5/TNtWJCXB+0p4BeWvzun2\\nwxbXNRbIUXNTRUXlsQmEHMGaks6g97ieKEV6xyZLQqPlOkc/w7bzvFEaRfi9AgMB\\nAAEwDQYJKoZIhvcNAQEFBQADggEBAJP8B1ms1Z8J7oBeG5QzqxmxjFeH6+/j+cHp\\nZm9xLlUrI0pzwdZ3c+yQxiq0u+fa5uWOD6UDyTMWu6XpRNOHy++gOHlHdNf774fk\\nBj3Cfnc2WIrtHF4AgPXIv7l6jZnceIFpDVuvwTLek0aDwsxH2WDSrLcgg3wlzxTK\\nfPxEMcvjLbAJXprsS+SlKLf9EtLx3LSGBUjwP0qfFkRJ1Kyze5H3HtUcIVhPqbOq\\n9iHUL8vQ0UYWJH5bEthBjAaTZVo/eQTisFlo5WCI9OwS/eQhjMKuYIN+IP1SX92L\\ndcZgGmUhQdSd4ALpDM2wxx8Hpo38DJ+08B6hb9ESsI938YGRUtk=\\n-----END CERTIFICATE-----",
					"private_key":             "-----BEGIN RSA PRIVATE KEY-----\\nMIIEogIBAAKCAQEAmsCMU9aMGlDZUeo5m1LRQ7YqjMRo5NvQ047F4x6mNRbWSUY5\\nZjUQt1ZXBu4itvCHUY5Ac/nWt21FQxONhYNHWlPJ9vpbOmRg3ax2aZwVcP/6wbko\\nkGSq723LMDi136UCuVNB4h4GQPxzCXCEm6dwGfL6mp80I9ZcuwaaeJMLTIzuCWNC\\nlxOCI5A0JltjPgPFW675j/MSfO1Mtebkap7SlHRFdYultUg9k5AMx1vCn35RxGMU\\nVAiBQ+bqKP2MRw8xR3iOf0zbViQlwftKeAXlr87p9sMW1zUWyFFzU0VF5bEJhBzB\\nmpLOoPe4nihFescmS0Kj5TpHP8O287xRGkX4vQIDAQABAoIBAGXAvYRSygRzrLFj\\n5UTwC5EOMqQkcSQqNQEmG4/fE0JNJdFAT3WY5sjmCIsSAdlwBknl0xNu73PkcWpN\\nbPyg+UW3WlD/BQU8A0U+pQ2EB8UpS2Qdr0JiLtMHylaGs2++PDBHQka/nzUTyUAn\\nr8n5Koilb4pDrDD0Pjzrdp5ZcVB/RDQBDh40CBi8SUaTFlJJRdat5obPKDATKNuR\\nV+ON2tTcDWkpEihJj6Xu/c6zqqu0sKohsg3P9eE1J4aGvkjc4kgtQ78u3XYF6c2s\\nKu2JA3/J1Ev6seZywe4oR+Io1U1YblEr6pE4MoN8Mte3iXnGvPLTK2Ll0dYAdCTP\\nWPE1gzECgYEAyEJn/JPB+wH6XgV9YFnDkbPMMQf30N3MqGzDD4n/vn/643xoGBUb\\nfwrtsEUhF74tRB/wbH+61SVzc80M+HQs8903v49K6SH54VLKciF4944+LY5Ujd19\\nPROsgrpMUwizzAFHdNw+pZgGZNLy/aKTaigzSccy4GL8tSh3WA2l7QcCgYEAxdN/\\nSeod5pVDCKfQ11lz7u7kP92Ex4gCh0E2ufjlne/7C2WgYoj70QArtrQTTsEY/nDk\\nRBMfAOAB3szG/LdrJO7IYJpNY+GXjg4PK08Kra5IWl/V5LfRs9Jh8IMC1lu7OnW5\\n+UfabGe9y4lew5I+AhYVqO1g79WEGeHfBMj0/xsCgYBGqlgzYphy0JHel6VUWJiQ\\nU5fcPWmPRJUxYZ+7XgDm4hZQK9g3l0sqm1qgduRkknE6tuKKBtCQ4jRJSrQnACfU\\nrd08NX5Dof+hKSvS9kjPYlxCixT/Moc9BDI9tyuzqUB43oKulAFvQaQP8/hW1AWc\\n0i62/BnR4Fw4ON1ULwy8BQKBgGDhTcIF/HSy/mwbuyPrc+I8bd10/5Sz4AEbB9EV\\numWZZAlV0LDQwvm8qqvEDnyQBkx9PtyzvIgyK9hP0tdqf/dLVSgyCNp0XUM9UWJ1\\nSBZ8doSD0H09JF57Fmmxz07pB4z+oIAbFzXlrEYkVRtT4DgnVp4u+j0aBKKkQhLW\\n01ynAoGAC5kV59AUWZAaeMMI6GtxRvgBE6aEB5/bWd99mISEW1igcP42kKmxPwsE\\ncepQFFWpEf6+FXDrzxi3wSsKe34SqnEbhI3f2MDF7mRVSZU1VWrKZ1ZWci0mVdba\\nbtbnVPYR2Ya6Y0o8/++CAYc1f4tlVffiwRgiEQq7JQBURvtqHaE=\\n-----END RSA PRIVATE KEY-----",
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"server_certificate_name": name,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"server_certificate":      "-----BEGIN CERTIFICATE-----\nMIIDUjCCAjoCCQDYCjnX6QH2NTANBgkqhkiG9w0BAQUFADBrMQswCQYDVQQGEwJD\nTjEQMA4GA1UECAwHQmVpSmluZzEQMA4GA1UEBwwHSGFpRGlhbjESMBAGA1UECgwJ\nTWVuVG91R291MQswCQYDVQQLDAJJVDEXMBUGA1UEAwwOKi5kengudGVzdC5jb20w\nHhcNMjMwODMwMDgxMTQ4WhcNMjQwODI5MDgxMTQ4WjBrMQswCQYDVQQGEwJDTjEQ\nMA4GA1UECAwHQmVpSmluZzEQMA4GA1UEBwwHSGFpRGlhbjESMBAGA1UECgwJTWVu\nVG91R291MQswCQYDVQQLDAJJVDEXMBUGA1UEAwwOKi5kengudGVzdC5jb20wggEi\nMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCawIxT1owaUNlR6jmbUtFDtiqM\nxGjk29DTjsXjHqY1FtZJRjlmNRC3VlcG7iK28IdRjkBz+da3bUVDE42Fg0daU8n2\n+ls6ZGDdrHZpnBVw//rBuSiQZKrvbcswOLXfpQK5U0HiHgZA/HMJcISbp3AZ8vqa\nnzQj1ly7Bpp4kwtMjO4JY0KXE4IjkDQmW2M+A8VbrvmP8xJ87Uy15uRqntKUdEV1\ni6W1SD2TkAzHW8KfflHEYxRUCIFD5uoo/YxHDzFHeI5/TNtWJCXB+0p4BeWvzun2\nwxbXNRbIUXNTRUXlsQmEHMGaks6g97ieKEV6xyZLQqPlOkc/w7bzvFEaRfi9AgMB\nAAEwDQYJKoZIhvcNAQEFBQADggEBAJP8B1ms1Z8J7oBeG5QzqxmxjFeH6+/j+cHp\nZm9xLlUrI0pzwdZ3c+yQxiq0u+fa5uWOD6UDyTMWu6XpRNOHy++gOHlHdNf774fk\nBj3Cfnc2WIrtHF4AgPXIv7l6jZnceIFpDVuvwTLek0aDwsxH2WDSrLcgg3wlzxTK\nfPxEMcvjLbAJXprsS+SlKLf9EtLx3LSGBUjwP0qfFkRJ1Kyze5H3HtUcIVhPqbOq\n9iHUL8vQ0UYWJH5bEthBjAaTZVo/eQTisFlo5WCI9OwS/eQhjMKuYIN+IP1SX92L\ndcZgGmUhQdSd4ALpDM2wxx8Hpo38DJ+08B6hb9ESsI938YGRUtk=\n-----END CERTIFICATE-----",
						"private_key":             "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAmsCMU9aMGlDZUeo5m1LRQ7YqjMRo5NvQ047F4x6mNRbWSUY5\nZjUQt1ZXBu4itvCHUY5Ac/nWt21FQxONhYNHWlPJ9vpbOmRg3ax2aZwVcP/6wbko\nkGSq723LMDi136UCuVNB4h4GQPxzCXCEm6dwGfL6mp80I9ZcuwaaeJMLTIzuCWNC\nlxOCI5A0JltjPgPFW675j/MSfO1Mtebkap7SlHRFdYultUg9k5AMx1vCn35RxGMU\nVAiBQ+bqKP2MRw8xR3iOf0zbViQlwftKeAXlr87p9sMW1zUWyFFzU0VF5bEJhBzB\nmpLOoPe4nihFescmS0Kj5TpHP8O287xRGkX4vQIDAQABAoIBAGXAvYRSygRzrLFj\n5UTwC5EOMqQkcSQqNQEmG4/fE0JNJdFAT3WY5sjmCIsSAdlwBknl0xNu73PkcWpN\nbPyg+UW3WlD/BQU8A0U+pQ2EB8UpS2Qdr0JiLtMHylaGs2++PDBHQka/nzUTyUAn\nr8n5Koilb4pDrDD0Pjzrdp5ZcVB/RDQBDh40CBi8SUaTFlJJRdat5obPKDATKNuR\nV+ON2tTcDWkpEihJj6Xu/c6zqqu0sKohsg3P9eE1J4aGvkjc4kgtQ78u3XYF6c2s\nKu2JA3/J1Ev6seZywe4oR+Io1U1YblEr6pE4MoN8Mte3iXnGvPLTK2Ll0dYAdCTP\nWPE1gzECgYEAyEJn/JPB+wH6XgV9YFnDkbPMMQf30N3MqGzDD4n/vn/643xoGBUb\nfwrtsEUhF74tRB/wbH+61SVzc80M+HQs8903v49K6SH54VLKciF4944+LY5Ujd19\nPROsgrpMUwizzAFHdNw+pZgGZNLy/aKTaigzSccy4GL8tSh3WA2l7QcCgYEAxdN/\nSeod5pVDCKfQ11lz7u7kP92Ex4gCh0E2ufjlne/7C2WgYoj70QArtrQTTsEY/nDk\nRBMfAOAB3szG/LdrJO7IYJpNY+GXjg4PK08Kra5IWl/V5LfRs9Jh8IMC1lu7OnW5\n+UfabGe9y4lew5I+AhYVqO1g79WEGeHfBMj0/xsCgYBGqlgzYphy0JHel6VUWJiQ\nU5fcPWmPRJUxYZ+7XgDm4hZQK9g3l0sqm1qgduRkknE6tuKKBtCQ4jRJSrQnACfU\nrd08NX5Dof+hKSvS9kjPYlxCixT/Moc9BDI9tyuzqUB43oKulAFvQaQP8/hW1AWc\n0i62/BnR4Fw4ON1ULwy8BQKBgGDhTcIF/HSy/mwbuyPrc+I8bd10/5Sz4AEbB9EV\numWZZAlV0LDQwvm8qqvEDnyQBkx9PtyzvIgyK9hP0tdqf/dLVSgyCNp0XUM9UWJ1\nSBZ8doSD0H09JF57Fmmxz07pB4z+oIAbFzXlrEYkVRtT4DgnVp4u+j0aBKKkQhLW\n01ynAoGAC5kV59AUWZAaeMMI6GtxRvgBE6aEB5/bWd99mISEW1igcP42kKmxPwsE\ncepQFFWpEf6+FXDrzxi3wSsKe34SqnEbhI3f2MDF7mRVSZU1VWrKZ1ZWci0mVdba\nbtbnVPYR2Ya6Y0o8/++CAYc1f4tlVffiwRgiEQq7JQBURvtqHaE=\n-----END RSA PRIVATE KEY-----",
						"resource_group_id":       CHECKSET,
						"server_certificate_name": name,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"server_certificate_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id":       CHECKSET,
						"server_certificate_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"private_key", "server_certificate"},
			},
		},
	})
}

var AlicloudSlbServerCertificateMap6748 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudSlbServerCertificateBasicDependence6748(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}


`, name)
}

// Case 全生命周期_alicert 6751
func TestAccAliCloudSlbServerCertificate_basic6751(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_slb_server_certificate.default"
	ra := resourceAttrInit(resourceId, AlicloudSlbServerCertificateMap6751)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &SlbServiceV2{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeSlbServerCertificate")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tfaccslb%d", rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudSlbServerCertificateBasicDependence6751)
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
					"resource_group_id":          "${data.alicloud_resource_manager_resource_groups.default.ids.0}",
					"server_certificate_name":    name,
					"ali_cloud_certificate_name": "${alicloud_ssl_certificates_service_certificate.cert.certificate_name}",
					"ali_cloud_certificate_id":   "${alicloud_ssl_certificates_service_certificate.cert.id}",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id":          CHECKSET,
						"server_certificate_name":    name,
						"ali_cloud_certificate_name": CHECKSET,
						"ali_cloud_certificate_id":   CHECKSET,
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"resource_group_id":       "${data.alicloud_resource_manager_resource_groups.default.ids.1}",
					"server_certificate_name": name + "_update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"resource_group_id":       CHECKSET,
						"server_certificate_name": name + "_update",
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
				ImportStateVerifyIgnore: []string{"private_key", "server_certificate"},
			},
		},
	})
}

var AlicloudSlbServerCertificateMap6751 = map[string]string{
	"create_time": CHECKSET,
	"region_id":   CHECKSET,
}

func AlicloudSlbServerCertificateBasicDependence6751(name string) string {
	return fmt.Sprintf(`
variable "name" {
    default = "%s"
}

data "alicloud_resource_manager_resource_groups" "default" {}

resource "alicloud_ssl_certificates_service_certificate" "cert" {
  cert             = <<EOF
-----BEGIN CERTIFICATE-----
MIIDPjCCAiYCCQDh5cM3cK8OXDANBgkqhkiG9w0BAQsFADBhMQswCQYDVQQGEwJh
YTELMAkGA1UECAwCYWExCzAJBgNVBAcMAmFhMQswCQYDVQQKDAJhYTELMAkGA1UE
CwwCYWExCzAJBgNVBAMMAmFhMREwDwYJKoZIhvcNAQkBFgJhYTAeFw0yNDEyMDUw
ODA3MzhaFw0yNTEyMDUwODA3MzhaMGExCzAJBgNVBAYTAmFhMQswCQYDVQQIDAJh
YTELMAkGA1UEBwwCYWExCzAJBgNVBAoMAmFhMQswCQYDVQQLDAJhYTELMAkGA1UE
AwwCYWExETAPBgkqhkiG9w0BCQEWAmFhMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEAuIxwdPXqG8GVf0f2NIJueHDRQkezLSZCFs5dbNfdRBVHQACf3N1/
2UhrlG+PWbSFcRjMUThUxOGx3p/WqvFAPxoHaQE2F3TNmy4ckSpKpnWmvTrZLDjC
TBa/62u8yrMKsDp7RDUPSODK5LCUaWa8r7488+DuyI759Fv7rQ+N2Lo/2NUYrmZ4
bM8ToJA19RberK0lPwkvkBzvV5OOiDa/7+Fil9+ofG6/FBvYLoUaI+MkN2SdTtlB
yiibJZc5PjV4z19nER0d7m7yK9rF8Vjaju6yrFq3rLvfNc/MI/bYnmoYXbyfaPP9
Lc9t9wowAz/7OSNuhkBFwdA+4oDRQwqK/QIDAQABMA0GCSqGSIb3DQEBCwUAA4IB
AQBj+k3MUTadQjtT0lcMT9O6A1vqz6IJU1Ry+yJ+Y88SzzXpxoPQ/7E1QrGPfZFs
fEqRQbyb8cnn0NkOTlOGL2/N5TiNtZtSib4q3ZbAvzMV6xUCbzmpkFg8hzvzX9NG
ll0U005Owc/5KO/KtaHWKtLRolxWqJoSzHJsAomNqp4PRQTwSUTP7ZN4z7ucSjzX
M4WYeY45/LLF0XqM1N7SUAFVbphf1xyrVtpAYWPOOvSV1Qg7ZSol4ljRSGdgf+cl
orAobJk78+tSEgbRSykffUEnE1NWWV43qxnk/pUUPNTvmZxmdl3WOJvGNqzbiko2
PXWOE2t5SISK0vIWpYw7kr2w
-----END CERTIFICATE-----
EOF
  certificate_name = "test-cert-name-374"
  key              = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAuIxwdPXqG8GVf0f2NIJueHDRQkezLSZCFs5dbNfdRBVHQACf
3N1/2UhrlG+PWbSFcRjMUThUxOGx3p/WqvFAPxoHaQE2F3TNmy4ckSpKpnWmvTrZ
LDjCTBa/62u8yrMKsDp7RDUPSODK5LCUaWa8r7488+DuyI759Fv7rQ+N2Lo/2NUY
rmZ4bM8ToJA19RberK0lPwkvkBzvV5OOiDa/7+Fil9+ofG6/FBvYLoUaI+MkN2Sd
TtlByiibJZc5PjV4z19nER0d7m7yK9rF8Vjaju6yrFq3rLvfNc/MI/bYnmoYXbyf
aPP9Lc9t9wowAz/7OSNuhkBFwdA+4oDRQwqK/QIDAQABAoIBAQCVhWOPv7ZfqqZf
K4S04oByoKVKmCD5kae2JpjXGMYy5TKHDnp3ThbJir1u1DxGp9X93eOcYpF7uoiM
IOCiOzicZ5BdaRfURVRYSDHpA2TcHTJs2oMeKxZw43W3XeDmMc2VHqzEhGDP4SFB
zDdSdmSl1vI3faS5Ze3qJ6RvgMyiLFH/HWlHTJru4gX1bkTknauF5iZIVkXyrc7W
hzznQycnJTdHvU5JAVATi0lxTFkE0Ghuzo+S40GU9cX+sN46F3grmNerlezdmXcB
jO1gqbjn5w9LhfHDhQW9M+vNMqHZYF9D2OXLuPFkyB/XQRYDvzKRFvwJzHUga1+4
/s+bz6uZAoGBAPQKxgFc1Jyn7iezIgx2o29RtYrrAPzn+5YQolVKekdCPCWuawKX
X/gbVUzfWHXRsP6nwC0oN4K9QuXcSkK6fAwqYH2qg/QPUi59y6zMOJfgDzJ0Jy49
TAgoRHH/Qp/R+djkUyNAhyCyM9j79o86tWuelj/5iI9JaWi4exPR18ebAoGBAMGX
Y3ibx4SkNvU+qu2IRiI1F8T+gnRRYxgNpVlgmnhDIdFOXyoQS/4YFoY8BnmjNyA7
iQXB+CGxTkvXw97T7D++YTOldcN6x7Hk5/bbrGXQDaACGcQJICrRDut66AHHNg1U
+1ITKOyT3G9KgexMvo/LkRTlj8nFPaIfHTd5qv1HAoGABF8lCOU20c+YTpHA7GQ7
rUTTcbef/ufQ8/E8VzbMVfZzyWEEfblkeLdUodnmZcXp28X5hVWRGQB28FMv6V4Z
1sXzkIy7bGXR3Q9X1LKb2OykmQzHnuZuND80aq5d7Sr/xMjiTGFdilJm514sHoi4
s0LWAjNgJOD1yHB1k/dQXD0CgYB5O2XWdK+IWBH2mM0PvUQ2TctZ2Qw27XTB7mG6
y+1QHLJD0fF0dQox0EHixBpF+DkgXUgcq/DIoLFoe3E0zL/o1z8lyZsq1GvT8oZ1
J/IZPCYcKtSZ2CElzw+p1akO3AL0bGCKcGczaBEmPKGKKjE3+YIcdPoLoBJNwXBZ
ej7jJwKBgQC45BUsy2P1Hq8YZqXmvJMwfURSDItz72GXZiCl+BErqX1FcC5vvOdK
8kUafqlVX6DkIhIuMjJ939icNop2wWrI318fYYZVuua5tGZN2yEKE+B7SJWKbq1Q
Fz8DK5vCe8boiD00V2rUd+6O5ElGLUeRACdQNeMudLQfJCRn1+9Jkw==
-----END RSA PRIVATE KEY-----
EOF
}


`, name)
}

// Test Slb ServerCertificate. <<< Resource test cases, automatically generated.

func TestAccAlicloudSLBServerCertificate_basic(t *testing.T) {
	var v *slb.ServerCertificate
	resourceId := "alicloud_slb_server_certificate.default"
	ra := resourceAttrInit(resourceId, serverCertificateMap)
	rc := resourceCheckInit(resourceId, &v, func() interface{} {
		return &SlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	})
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := fmt.Sprintf("tf-testAccSlbServerCertificate")
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceServerCertificateBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"name":               name,
					"server_certificate": `-----BEGIN CERTIFICATE-----\nMIICWDCCAcGgAwIBAgIJAP7vOtjPtQIjMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkNOMRMwEQYDVQQIDApjbi1iZWlqaW5nMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMjAxMDIwMDYxOTUxWhcNMjAxMTE5MDYxOTUxWjBF\nMQswCQYDVQQGEwJDTjETMBEGA1UECAwKY24tYmVpamluZzEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB\ngQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9BVuFIBoU8nrP\nY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2CNIzxr9DjCzN5\ntWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQABo1AwTjAdBgNV\nHQ4EFgQUYDwuuqC2a2UPrfm1v31vE7+GRM4wHwYDVR0jBBgwFoAUYDwuuqC2a2UP\nrfm1v31vE7+GRM4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQAovSB0\n5JRKrg7lYR/KlTuKHmozfyL9UER0/dpTSoqsCyt8yc1BbtAKUJWh09BujBE1H22f\nlKvCAjhPmnNdfd/l9GrmAWNDWEDPLdUTkGSkKAScMpdS+mLmOBuYWgdnOtq3eQGf\nt07tlBL+dtzrrohHpfLeuNyYb40g8VQdp3RRRQ==\n-----END CERTIFICATE-----`,
					"private_key":        `-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9\nBVuFIBoU8nrPY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2C\nNIzxr9DjCzN5tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQAB\nAoGARe2oaCo5lTDK+c4Zx3392hoqQ94r0DmWHPBvNmwAooYd+YxLPrLMe5sMjY4t\ndmohnLNevCK1Uzw5eIX6BNSo5CORBcIDRmiAgwiYiS3WOv2+qi9g5uIdMiDr+EED\nK8wZJjB5E2WyfxL507vtW4T5L36yfr8SkmqH3GvzpI2jCqECQQDsy0AmBzyfK0tG\nNw1+iF9SReJWgb1f5iHvz+6Dt5ueVQngrl/5++Gp5bNoaQMkLEDsy0iHIj9j43ji\n0DON05uDAkEA1GXgGn8MXXKyuzYuoyYXCBH7aF579d7KEGET/jjnXx9DHcfRJZBY\nB9ghMnnonSOGboF04Zsdd3xwYF/3OHYssQJAekd/SeQEzyE5TvoQ8t2Tc9X4yrlW\nxNX/gmp6/fPr3biGUEtb7qi+4NBodCt+XsingmB7hKUP3RJTk7T2WnAC5wJAMqHi\njY5x3SkFkHl3Hq9q2CKpQxUbCd7FXqg1wum/xj5GmqfSpNjHE3+jUkwbdrJMTrWP\nrmRy3tQMWf0mixAo0QJBAN4IcZChanq8cZyNqqoNbxGm4hkxUmE0W4hxHmLC2CYZ\nV4JpNm8dpi4CiMWLasF6TYlVMgX+aPxYRUWc/qqf1/Q=\n-----END RSA PRIVATE KEY-----`,
					"resource_group_id":  "${data.alicloud_resource_manager_resource_groups.default.groups.0.id}",
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test123",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					// the alicloud_certificate_id/alicloud_certificate_name depend on anothor alibaba cloud certificate product.
					// but now it is not suppot on alibaba cloud international site.
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test123",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"server_certificate", "private_key"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF1",
						"For":     "acceptance test1231",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF1",
						"tags.For":     "acceptance test1231",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name": "tf-testAccSlbServerCertificateUpdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					// the alicloud_certificate_id/alicloud_certificate_name depend on anothor alibaba cloud certificate product.
					// but now it is not suppot on alibaba cloud international site.
					testAccCheck(map[string]string{
						"name": "tf-testAccSlbServerCertificateUpdate",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name": name,
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test123",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					// the alicloud_certificate_id/alicloud_certificate_name depend on anothor alibaba cloud certificate product.
					// but now it is not suppot on alibaba cloud international site.
					testAccCheck(map[string]string{
						"name":         name,
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test123",
					}),
				),
			},
		},
	})
}

// There is an unknown error for the testcase
func SkipTestAccAlicloudSLBServerCertificate_aliCloud_server_certificate(t *testing.T) {
	var v *slb.ServerCertificate
	resourceId := "alicloud_slb_server_certificate.default"
	ra := resourceAttrInit(resourceId, serverAliCloudCertificateMap)
	rc := resourceCheckInit(resourceId, &v, func() interface{} {
		return &SlbService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	})
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	randInt := acctest.RandInt()
	alicloudCertificateName := fmt.Sprintf("tf_testAcc_%d", randInt)
	testAccConfig := resourceTestAccConfigFunc(resourceId, alicloudCertificateName, resourceServerCertificateAliCloudServerCertificateBasicDependence)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckWithRegions(t, true, connectivity.CasClassicSupportedRegions)
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"name":                      "tf-testAccSlbServerCertificate",
					"alicloud_certificate_id":   "${alicloud_cas_certificate.default.id}",
					"alicloud_certificate_name": "${alicloud_cas_certificate.default.name}",
					//"alicloud_certificate_region_id": os.Getenv("ALICLOUD_REGION"),
					"resource_group_id": "${data.alicloud_resource_manager_resource_groups.default.groups.0.id}",
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test123",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					// the alicloud_certificate_id/alicloud_certificate_name depend on anothor alibaba cloud certificate product.
					// but now it is not suppot on alibaba cloud international site.
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test123",
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"server_certificate", "private_key", "alicloud_certificate_region_id"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"tags": map[string]string{
						"Created": "TF1",
						"For":     "acceptance test1231",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"tags.%":       "2",
						"tags.Created": "TF1",
						"tags.For":     "acceptance test1231",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name": "tf-testAccSlbServerCertificateUpdate",
				}),
				Check: resource.ComposeTestCheckFunc(
					// the alicloud_certificate_id/alicloud_certificate_name depend on anothor alibaba cloud certificate product.
					// but now it is not suppot on alibaba cloud international site.
					testAccCheck(map[string]string{
						"name": "tf-testAccSlbServerCertificateUpdate",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name": "tf-testAccSlbServerCertificate",
					"tags": map[string]string{
						"Created": "TF",
						"For":     "acceptance test123",
					},
				}),
				Check: resource.ComposeTestCheckFunc(
					// the alicloud_certificate_id/alicloud_certificate_name depend on anothor alibaba cloud certificate product.
					// but now it is not suppot on alibaba cloud international site.
					testAccCheck(map[string]string{
						"name":         "tf-testAccSlbServerCertificate",
						"tags.%":       "2",
						"tags.Created": "TF",
						"tags.For":     "acceptance test123",
					}),
				),
			},
		},
	})
}

func resourceServerCertificateBasicDependence(name string) string {
	return `
data "alicloud_resource_manager_resource_groups" "default" {
  name_regex = "default"
}
`
}

func resourceServerCertificateAliCloudServerCertificateBasicDependence(name string) string {
	return fmt.Sprintf(`
data "alicloud_resource_manager_resource_groups" "default" {
  name_regex = "default"
}
resource "alicloud_cas_certificate" "default" {
  name = "%s"
  cert = "-----BEGIN CERTIFICATE-----\nMIICWDCCAcGgAwIBAgIJAP7vOtjPtQIjMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkNOMRMwEQYDVQQIDApjbi1iZWlqaW5nMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMjAxMDIwMDYxOTUxWhcNMjAxMTE5MDYxOTUxWjBF\nMQswCQYDVQQGEwJDTjETMBEGA1UECAwKY24tYmVpamluZzEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB\ngQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9BVuFIBoU8nrP\nY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2CNIzxr9DjCzN5\ntWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQABo1AwTjAdBgNV\nHQ4EFgQUYDwuuqC2a2UPrfm1v31vE7+GRM4wHwYDVR0jBBgwFoAUYDwuuqC2a2UP\nrfm1v31vE7+GRM4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQAovSB0\n5JRKrg7lYR/KlTuKHmozfyL9UER0/dpTSoqsCyt8yc1BbtAKUJWh09BujBE1H22f\nlKvCAjhPmnNdfd/l9GrmAWNDWEDPLdUTkGSkKAScMpdS+mLmOBuYWgdnOtq3eQGf\nt07tlBL+dtzrrohHpfLeuNyYb40g8VQdp3RRRQ==\n-----END CERTIFICATE-----"
  key = "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9\nBVuFIBoU8nrPY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2C\nNIzxr9DjCzN5tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQAB\nAoGARe2oaCo5lTDK+c4Zx3392hoqQ94r0DmWHPBvNmwAooYd+YxLPrLMe5sMjY4t\ndmohnLNevCK1Uzw5eIX6BNSo5CORBcIDRmiAgwiYiS3WOv2+qi9g5uIdMiDr+EED\nK8wZJjB5E2WyfxL507vtW4T5L36yfr8SkmqH3GvzpI2jCqECQQDsy0AmBzyfK0tG\nNw1+iF9SReJWgb1f5iHvz+6Dt5ueVQngrl/5++Gp5bNoaQMkLEDsy0iHIj9j43ji\n0DON05uDAkEA1GXgGn8MXXKyuzYuoyYXCBH7aF579d7KEGET/jjnXx9DHcfRJZBY\nB9ghMnnonSOGboF04Zsdd3xwYF/3OHYssQJAekd/SeQEzyE5TvoQ8t2Tc9X4yrlW\nxNX/gmp6/fPr3biGUEtb7qi+4NBodCt+XsingmB7hKUP3RJTk7T2WnAC5wJAMqHi\njY5x3SkFkHl3Hq9q2CKpQxUbCd7FXqg1wum/xj5GmqfSpNjHE3+jUkwbdrJMTrWP\nrmRy3tQMWf0mixAo0QJBAN4IcZChanq8cZyNqqoNbxGm4hkxUmE0W4hxHmLC2CYZ\nV4JpNm8dpi4CiMWLasF6TYlVMgX+aPxYRUWc/qqf1/Q=\n-----END RSA PRIVATE KEY-----"
}
`, name)
}

var serverCertificateMap = map[string]string{
	"name":                      "tf-testAccSlbServerCertificate",
	"server_certificate":        "-----BEGIN CERTIFICATE-----\nMIICWDCCAcGgAwIBAgIJAP7vOtjPtQIjMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkNOMRMwEQYDVQQIDApjbi1iZWlqaW5nMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMjAxMDIwMDYxOTUxWhcNMjAxMTE5MDYxOTUxWjBF\nMQswCQYDVQQGEwJDTjETMBEGA1UECAwKY24tYmVpamluZzEhMB8GA1UECgwYSW50\nZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB\ngQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9BVuFIBoU8nrP\nY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2CNIzxr9DjCzN5\ntWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQABo1AwTjAdBgNV\nHQ4EFgQUYDwuuqC2a2UPrfm1v31vE7+GRM4wHwYDVR0jBBgwFoAUYDwuuqC2a2UP\nrfm1v31vE7+GRM4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQAovSB0\n5JRKrg7lYR/KlTuKHmozfyL9UER0/dpTSoqsCyt8yc1BbtAKUJWh09BujBE1H22f\nlKvCAjhPmnNdfd/l9GrmAWNDWEDPLdUTkGSkKAScMpdS+mLmOBuYWgdnOtq3eQGf\nt07tlBL+dtzrrohHpfLeuNyYb40g8VQdp3RRRQ==\n-----END CERTIFICATE-----",
	"private_key":               "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9\nBVuFIBoU8nrPY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2C\nNIzxr9DjCzN5tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQAB\nAoGARe2oaCo5lTDK+c4Zx3392hoqQ94r0DmWHPBvNmwAooYd+YxLPrLMe5sMjY4t\ndmohnLNevCK1Uzw5eIX6BNSo5CORBcIDRmiAgwiYiS3WOv2+qi9g5uIdMiDr+EED\nK8wZJjB5E2WyfxL507vtW4T5L36yfr8SkmqH3GvzpI2jCqECQQDsy0AmBzyfK0tG\nNw1+iF9SReJWgb1f5iHvz+6Dt5ueVQngrl/5++Gp5bNoaQMkLEDsy0iHIj9j43ji\n0DON05uDAkEA1GXgGn8MXXKyuzYuoyYXCBH7aF579d7KEGET/jjnXx9DHcfRJZBY\nB9ghMnnonSOGboF04Zsdd3xwYF/3OHYssQJAekd/SeQEzyE5TvoQ8t2Tc9X4yrlW\nxNX/gmp6/fPr3biGUEtb7qi+4NBodCt+XsingmB7hKUP3RJTk7T2WnAC5wJAMqHi\njY5x3SkFkHl3Hq9q2CKpQxUbCd7FXqg1wum/xj5GmqfSpNjHE3+jUkwbdrJMTrWP\nrmRy3tQMWf0mixAo0QJBAN4IcZChanq8cZyNqqoNbxGm4hkxUmE0W4hxHmLC2CYZ\nV4JpNm8dpi4CiMWLasF6TYlVMgX+aPxYRUWc/qqf1/Q=\n-----END RSA PRIVATE KEY-----",
	"alicloud_certificate_id":   NOSET,
	"alicloud_certificate_name": NOSET,
	"resource_group_id":         CHECKSET,
}

var serverAliCloudCertificateMap = map[string]string{
	"name":                           "tf-testAccSlbServerCertificate",
	"server_certificate":             NOSET,
	"private_key":                    NOSET,
	"alicloud_certificate_region_id": NOSET,
	"alicloud_certificate_id":        CHECKSET,
	"alicloud_certificate_name":      CHECKSET,
	"resource_group_id":              CHECKSET,
}
