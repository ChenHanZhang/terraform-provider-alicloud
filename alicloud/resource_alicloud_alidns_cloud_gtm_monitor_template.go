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

func resourceAliCloudAlidnsCloudGtmMonitorTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAlidnsCloudGtmMonitorTemplateCreate,
		Read:   resourceAliCloudAlidnsCloudGtmMonitorTemplateRead,
		Update: resourceAliCloudAlidnsCloudGtmMonitorTemplateUpdate,
		Delete: resourceAliCloudAlidnsCloudGtmMonitorTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"evaluation_count": {
				Type:     schema.TypeString,
				Required: true,
			},
			"extend_info": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"failure_rate": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"interval": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"remark": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudAlidnsCloudGtmMonitorTemplateCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateCloudGtmMonitorTemplate"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	request["Interval"] = d.Get("interval")
	request["Protocol"] = d.Get("protocol")
	request["IpVersion"] = d.Get("ip_version")
	request["EvaluationCount"] = d.Get("evaluation_count")
	request["Timeout"] = d.Get("timeout")
	if v, ok := d.GetOk("extend_info"); ok {
		request["ExtendInfo"] = v
	}
	request["FailureRate"] = d.Get("failure_rate")
	request["Name"] = d.Get("name")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Alidns", "2015-01-09", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_alidns_cloud_gtm_monitor_template", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["TemplateId"]))

	return resourceAliCloudAlidnsCloudGtmMonitorTemplateUpdate(d, meta)
}

func resourceAliCloudAlidnsCloudGtmMonitorTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	alidnsServiceV2 := AlidnsServiceV2{client}

	objectRaw, err := alidnsServiceV2.DescribeAlidnsCloudGtmMonitorTemplate(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_alidns_cloud_gtm_monitor_template DescribeAlidnsCloudGtmMonitorTemplate Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("evaluation_count", objectRaw["EvaluationCount"])
	d.Set("extend_info", objectRaw["ExtendInfo"])
	d.Set("failure_rate", objectRaw["FailureRate"])
	d.Set("interval", objectRaw["Interval"])
	d.Set("ip_version", objectRaw["IpVersion"])
	d.Set("name", objectRaw["Name"])
	d.Set("protocol", objectRaw["Protocol"])
	d.Set("remark", objectRaw["Remark"])
	d.Set("timeout", objectRaw["Timeout"])

	return nil
}

func resourceAliCloudAlidnsCloudGtmMonitorTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UpdateCloudGtmMonitorTemplate"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["TemplateId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("interval") {
		update = true
	}
	request["Interval"] = d.Get("interval")
	if !d.IsNewResource() && d.HasChange("evaluation_count") {
		update = true
	}
	request["EvaluationCount"] = d.Get("evaluation_count")
	if !d.IsNewResource() && d.HasChange("timeout") {
		update = true
	}
	request["Timeout"] = d.Get("timeout")
	if !d.IsNewResource() && d.HasChange("extend_info") {
		update = true
		request["ExtendInfo"] = d.Get("extend_info")
	}

	if !d.IsNewResource() && d.HasChange("failure_rate") {
		update = true
	}
	request["FailureRate"] = d.Get("failure_rate")
	if !d.IsNewResource() && d.HasChange("name") {
		update = true
	}
	request["Name"] = d.Get("name")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Alidns", "2015-01-09", action, query, request, true)
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
	action = "UpdateCloudGtmMonitorTemplateRemark"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["TemplateId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("remark") {
		update = true
		request["Remark"] = d.Get("remark")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Alidns", "2015-01-09", action, query, request, true)
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
	return resourceAliCloudAlidnsCloudGtmMonitorTemplateRead(d, meta)
}

func resourceAliCloudAlidnsCloudGtmMonitorTemplateDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteCloudGtmMonitorTemplate"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["TemplateId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Alidns", "2015-01-09", action, query, request, true)
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
