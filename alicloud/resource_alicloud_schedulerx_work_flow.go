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

func resourceAliCloudSchedulerxWorkFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSchedulerxWorkFlowCreate,
		Read:   resourceAliCloudSchedulerxWorkFlowRead,
		Update: resourceAliCloudSchedulerxWorkFlowUpdate,
		Delete: resourceAliCloudSchedulerxWorkFlowDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"max_concurrency": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"namespace_source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Enable", "Disable"}, false),
			},
			"time_expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"time_type": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"work_flow_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"workflow_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudSchedulerxWorkFlowCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateWorkflow"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("namespace"); ok {
		request["Namespace"] = v
	}
	if v, ok := d.GetOk("group_id"); ok {
		request["GroupId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("time_expression"); ok {
		request["TimeExpression"] = v
	}
	request["Name"] = d.Get("workflow_name")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	request["TimeType"] = d.Get("time_type")
	if v, ok := d.GetOk("timezone"); ok {
		request["Timezone"] = v
	}
	if v, ok := d.GetOk("namespace_source"); ok {
		request["NamespaceSource"] = v
	}
	if v, ok := d.GetOkExists("max_concurrency"); ok {
		request["MaxConcurrency"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("schedulerx2", "2019-04-30", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_schedulerx_work_flow", action, AlibabaCloudSdkGoERROR)
	}

	DataWorkflowIdVar, _ := jsonpath.Get("$.Data.WorkflowId", response)
	d.SetId(fmt.Sprintf("%v:%v:%v", request["Namespace"], request["GroupId"], DataWorkflowIdVar))

	return resourceAliCloudSchedulerxWorkFlowUpdate(d, meta)
}

func resourceAliCloudSchedulerxWorkFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	schedulerxServiceV2 := SchedulerxServiceV2{client}

	objectRaw, err := schedulerxServiceV2.DescribeSchedulerxWorkFlow(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_schedulerx_work_flow DescribeSchedulerxWorkFlow Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["Description"])
	d.Set("max_concurrency", formatInt(objectRaw["MaxConcurrency"]))
	d.Set("status", convertSchedulerxWorkFlowDataWorkFlowInfoStatusResponse(objectRaw["Status"]))
	d.Set("time_expression", objectRaw["TimeExpression"])
	d.Set("time_type", formatInt(convertSchedulerxWorkFlowDataWorkFlowInfoTimeTypeResponse(objectRaw["TimeType"])))
	d.Set("workflow_name", objectRaw["Name"])
	d.Set("group_id", objectRaw["GroupId"])
	d.Set("namespace", objectRaw["Namespace"])
	d.Set("work_flow_id", objectRaw["WorkflowId"])

	return nil
}

func resourceAliCloudSchedulerxWorkFlowUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	schedulerxServiceV2 := SchedulerxServiceV2{client}
	objectRaw, _ := schedulerxServiceV2.DescribeSchedulerxWorkFlow(d.Id())

	if d.HasChange("status") {
		var err error
		target := d.Get("status").(string)

		currentStatus, err := jsonpath.Get("Status", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "Status", objectRaw)
		}
		if fmt.Sprint(convertSchedulerxWorkFlowDataWorkFlowInfoStatusResponse(currentStatus)) != target {
			if target == "Disable" {
				parts := strings.Split(d.Id(), ":")
				action := "DisableWorkflow"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				query["Namespace"] = parts[0]
				query["GroupId"] = parts[1]
				query["WorkflowId"] = parts[2]
				query["RegionId"] = client.RegionId
				if v, ok := d.GetOk("namespace_source"); ok {
					query["NamespaceSource"] = v.(string)
				}

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcGet("schedulerx2", "2019-04-30", action, query, request)
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
			if target == "Enable" {
				parts := strings.Split(d.Id(), ":")
				action := "EnableWorkflow"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				query["Namespace"] = parts[0]
				query["GroupId"] = parts[1]
				query["WorkflowId"] = parts[2]
				query["RegionId"] = client.RegionId
				if v, ok := d.GetOk("namespace_source"); ok {
					query["NamespaceSource"] = v.(string)
				}

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcGet("schedulerx2", "2019-04-30", action, query, request)
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

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdateWorkflow"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["Namespace"] = parts[0]
	request["GroupId"] = parts[1]
	request["WorkflowId"] = parts[2]
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("time_expression") {
		update = true
		request["TimeExpression"] = d.Get("time_expression")
	}

	if !d.IsNewResource() && d.HasChange("workflow_name") {
		update = true
	}
	request["Name"] = d.Get("workflow_name")
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if !d.IsNewResource() && d.HasChange("time_type") {
		update = true
	}
	request["TimeType"] = d.Get("time_type")
	if v, ok := d.GetOk("namespace_source"); ok {
		request["NamespaceSource"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("schedulerx2", "2019-04-30", action, query, request, true)
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

	return resourceAliCloudSchedulerxWorkFlowRead(d, meta)
}

func resourceAliCloudSchedulerxWorkFlowDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteWorkflow"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	query["Namespace"] = parts[0]
	query["GroupId"] = parts[1]
	query["WorkflowId"] = parts[2]
	query["RegionId"] = client.RegionId

	if v, ok := d.GetOk("namespace_source"); ok {
		query["NamespaceSource"] = v.(string)
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcGet("schedulerx2", "2019-04-30", action, query, request)
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

func convertSchedulerxWorkFlowDataWorkFlowInfoStatusResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "enable":
		return "Enable"
	case "disable":
		return "Disable"
	}
	return source
}

func convertSchedulerxWorkFlowDataWorkFlowInfoTimeTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "cron":
		return "1"
	case "api":
		return "100"
	}
	return source
}
