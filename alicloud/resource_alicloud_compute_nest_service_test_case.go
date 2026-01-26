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

func resourceAliCloudComputeNestServiceTestCase() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudComputeNestServiceTestCaseCreate,
		Read:   resourceAliCloudComputeNestServiceTestCaseRead,
		Update: resourceAliCloudComputeNestServiceTestCaseUpdate,
		Delete: resourceAliCloudComputeNestServiceTestCaseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"test_case_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"test_config": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudComputeNestServiceTestCaseCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateServiceTestCase"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	request["TemplateName"] = d.Get("template_name")
	request["TestCaseName"] = d.Get("test_case_name")
	request["TestConfig"] = d.Get("test_config")
	request["ServiceId"] = d.Get("service_id")
	request["ServiceVersion"] = d.Get("service_version")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ComputeNestSupplier", "2021-05-21", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_compute_nest_service_test_case", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["TestCaseId"]))

	return resourceAliCloudComputeNestServiceTestCaseRead(d, meta)
}

func resourceAliCloudComputeNestServiceTestCaseRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	computeNestServiceV2 := ComputeNestServiceV2{client}

	objectRaw, err := computeNestServiceV2.DescribeComputeNestServiceTestCase(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_compute_nest_service_test_case DescribeComputeNestServiceTestCase Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("template_name", objectRaw["TemplateName"])
	d.Set("test_case_name", objectRaw["TestCaseName"])
	d.Set("test_config", objectRaw["TestConfig"])

	return nil
}

func resourceAliCloudComputeNestServiceTestCaseUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateServiceTestCase"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["TestCaseId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("test_case_name") {
		update = true
	}
	request["TestCaseName"] = d.Get("test_case_name")
	if d.HasChange("test_config") {
		update = true
	}
	request["TestConfig"] = d.Get("test_config")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("ComputeNestSupplier", "2021-05-21", action, query, request, true)
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

	return resourceAliCloudComputeNestServiceTestCaseRead(d, meta)
}

func resourceAliCloudComputeNestServiceTestCaseDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteServiceTestCase"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["TestCaseId"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("ComputeNestSupplier", "2021-05-21", action, query, request, true)
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
