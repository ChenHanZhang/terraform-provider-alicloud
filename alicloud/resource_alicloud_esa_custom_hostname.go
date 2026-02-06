// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEsaCustomHostname() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEsaCustomHostnameCreate,
		Read:   resourceAliCloudEsaCustomHostnameRead,
		Update: resourceAliCloudEsaCustomHostnameUpdate,
		Delete: resourceAliCloudEsaCustomHostnameDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cas_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cas_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cert_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"record_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ssl_flag": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudEsaCustomHostnameCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateCustomHostname"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["RecordId"] = d.Get("record_id")
	if v, ok := d.GetOk("private_key"); ok {
		request["PrivateKey"] = v
	}
	request["SiteId"] = d.Get("site_id")
	request["SslFlag"] = d.Get("ssl_flag")
	request["Hostname"] = d.Get("hostname")
	if v, ok := d.GetOkExists("cas_id"); ok {
		request["CasId"] = v
	}
	if v, ok := d.GetOk("certificate"); ok {
		request["Certificate"] = v
	}
	request["CertType"] = d.Get("cert_type")
	if v, ok := d.GetOk("cas_region"); ok {
		request["CasRegion"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ESA", "2024-09-10", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"TooManyRequests"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_esa_custom_hostname", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["HostnameId"]))

	return resourceAliCloudEsaCustomHostnameRead(d, meta)
}

func resourceAliCloudEsaCustomHostnameRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	esaServiceV2 := EsaServiceV2{client}

	objectRaw, err := esaServiceV2.DescribeEsaCustomHostname(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_esa_custom_hostname DescribeEsaCustomHostname Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cas_id", objectRaw["CasId"])
	d.Set("cert_type", objectRaw["CertType"])
	d.Set("certificate", objectRaw["Certificate"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("hostname", objectRaw["Hostname"])
	d.Set("private_key", objectRaw["PrivateKey"])
	d.Set("record_id", objectRaw["RecordId"])
	d.Set("site_id", objectRaw["SiteId"])
	d.Set("ssl_flag", objectRaw["SslFlag"])
	d.Set("status", objectRaw["Status"])

	return nil
}

func resourceAliCloudEsaCustomHostnameUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UpdateCustomHostname"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["HostnameId"] = d.Id()

	if d.HasChange("record_id") {
		update = true
	}
	request["RecordId"] = d.Get("record_id")
	if d.HasChange("private_key") {
		update = true
		request["PrivateKey"] = d.Get("private_key")
	}

	if d.HasChange("ssl_flag") {
		update = true
	}
	request["SslFlag"] = d.Get("ssl_flag")
	if d.HasChange("cas_id") {
		update = true
		request["CasId"] = d.Get("cas_id")
	}

	if d.HasChange("certificate") {
		update = true
		request["Certificate"] = d.Get("certificate")
	}

	if d.HasChange("cert_type") {
		update = true
	}
	request["CertType"] = d.Get("cert_type")
	if v, ok := d.GetOk("cas_region"); ok {
		request["CasRegion"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	action = "ApplyCustomHostnameCertificate"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["HostnameId"] = d.Id()

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}
	update = false
	action = "VerifyCustomHostname"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["HostnameId"] = d.Id()

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	d.Partial(false)
	return resourceAliCloudEsaCustomHostnameRead(d, meta)
}

func resourceAliCloudEsaCustomHostnameDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteCustomHostname"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["HostnameId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
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
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
