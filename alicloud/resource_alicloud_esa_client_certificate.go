// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEsaClientCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEsaClientCertificateCreate,
		Read:   resourceAliCloudEsaClientCertificateRead,
		Update: resourceAliCloudEsaClientCertificateUpdate,
		Delete: resourceAliCloudEsaClientCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"csr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pkey_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validity_days": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudEsaClientCertificateCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateClientCertificate"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("site_id"); ok {
		request["SiteId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("csr"); ok {
		request["CSR"] = v
	}
	request["ValidityDays"] = d.Get("validity_days")
	if v, ok := d.GetOk("pkey_type"); ok {
		request["PkeyType"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ESA", "2024-09-10", action, query, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_esa_client_certificate", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["SiteId"], response["Id"]))

	return resourceAliCloudEsaClientCertificateUpdate(d, meta)
}

func resourceAliCloudEsaClientCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	esaServiceV2 := EsaServiceV2{client}

	objectRaw, err := esaServiceV2.DescribeEsaClientCertificate(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_esa_client_certificate DescribeEsaClientCertificate Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("site_id", formatInt(objectRaw["SiteId"]))

	resultRawObj, _ := jsonpath.Get("$.Result", objectRaw)
	resultRaw := make(map[string]interface{})
	if resultRawObj != nil {
		resultRaw = resultRawObj.(map[string]interface{})
	}
	d.Set("create_time", resultRaw["CreateTime"])
	d.Set("status", resultRaw["Status"])
	d.Set("client_cert_id", resultRaw["Id"])

	return nil
}

func resourceAliCloudEsaClientCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}

	if d.HasChange("status") {
		esaServiceV2 := EsaServiceV2{client}
		object, err := esaServiceV2.DescribeEsaClientCertificate(d.Id())
		if err != nil {
			return WrapError(err)
		}

		target := d.Get("status").(string)
		status, err := jsonpath.Get("$.Result.Status", object)
		if fmt.Sprint(status) != target {
			if target == "revoked" {
				parts := strings.Split(d.Id(), ":")
				action := "RevokeClientCertificate"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				query["SiteId"] = parts[0]
				query["Id"] = parts[1]
				query["RegionId"] = client.RegionId
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcGet("ESA", "2024-09-10", action, query, request)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					return nil
				})
				addDebug(action, response, request)
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
			if target == "active" {
				parts := strings.Split(d.Id(), ":")
				action := "ActivateClientCertificate"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				query["SiteId"] = parts[0]
				query["Id"] = parts[1]
				query["RegionId"] = client.RegionId
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcGet("ESA", "2024-09-10", action, query, request)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					return nil
				})
				addDebug(action, response, request)
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
		}
	}

	return resourceAliCloudEsaClientCertificateRead(d, meta)
}

func resourceAliCloudEsaClientCertificateDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteClientCertificate"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	query["SiteId"] = parts[0]
	query["Id"] = parts[1]
	query["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcGet("ESA", "2024-09-10", action, query, request)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
