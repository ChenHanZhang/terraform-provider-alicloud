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

func resourceAliCloudActionTrailAnalysisTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudActionTrailAnalysisTaskCreate,
		Read:   resourceAliCloudActionTrailAnalysisTaskRead,
		Delete: resourceAliCloudActionTrailAnalysisTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"query_conditions": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"query_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudActionTrailAnalysisTaskCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAnalysisTask"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["QueryType"] = d.Get("query_type")
	request["StartTime"] = d.Get("start_time")
	request["QueryConditions"] = d.Get("query_conditions")
	request["EndTime"] = d.Get("end_time")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Actiontrail", "2020-07-06", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_action_trail_analysis_task", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["AnalysisTaskId"]))

	return resourceAliCloudActionTrailAnalysisTaskRead(d, meta)
}

func resourceAliCloudActionTrailAnalysisTaskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	actionTrailServiceV2 := ActionTrailServiceV2{client}

	objectRaw, err := actionTrailServiceV2.DescribeActionTrailAnalysisTask(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_action_trail_analysis_task DescribeActionTrailAnalysisTask Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("end_time", objectRaw["EndTime"])
	d.Set("query_conditions", objectRaw["QueryConditions"])
	d.Set("query_type", objectRaw["QueryType"])
	d.Set("start_time", objectRaw["StartTime"])
	d.Set("status", objectRaw["Status"])

	return nil
}

func resourceAliCloudActionTrailAnalysisTaskDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAnalysisTask"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AnalysisTaskId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Actiontrail", "2020-07-06", action, query, request, true)
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
