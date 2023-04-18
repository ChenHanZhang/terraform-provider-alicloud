// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudVpcFlowLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcFlowLogCreate,
		Read:   resourceAlicloudVpcFlowLogRead,
		Update: resourceAlicloudVpcFlowLogUpdate,
		Delete: resourceAlicloudVpcFlowLogDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"aggregation_interval": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"business_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_log_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_log_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_store_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": tagsSchema(),
			"traffic_path": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"traffic_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAlicloudVpcFlowLogCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateFlowLog"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("project_name"); ok {
		request["ProjectName"] = v
	}

	if v, ok := d.GetOk("resource_id"); ok {
		request["ResourceId"] = v
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}

	if v, ok := d.GetOk("log_store_name"); ok {
		request["LogStoreName"] = v
	}

	if v, ok := d.GetOk("traffic_type"); ok {
		request["TrafficType"] = v
	}

	if v, ok := d.GetOk("flow_log_name"); ok {
		request["FlowLogName"] = v
	}

	if v, ok := d.GetOk("aggregation_interval"); ok {
		request["AggregationInterval"] = v
	}

	if v, ok := d.GetOk("traffic_path"); ok {
		localData := v
		trafficPathMaps := localData.([]interface{})
		request["TrafficPath"] = trafficPathMaps
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("resource_type"); ok {
		request["ResourceType"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"TaskConflict", "OperationFailed.LastTokenProcessing", "LastTokenProcessing", "IncorrectStatus.%s", "InvalidHdMonitorStatus", "OperationConflict", "ServiceUnavailable", "SystemBusy"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_flow_log", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["FlowLogId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcFlowLogStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcFlowLogUpdate(d, meta)
}

func resourceAlicloudVpcFlowLogRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcFlowLog(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_flow_log .DescribeVpcFlowLog Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("aggregation_interval", object["aggregation_interval"])
	d.Set("business_status", object["business_status"])
	d.Set("create_time", object["create_time"])
	d.Set("description", object["description"])
	d.Set("flow_log_id", object["flow_log_id"])
	d.Set("flow_log_name", object["flow_log_name"])
	d.Set("log_store_name", object["log_store_name"])
	d.Set("project_name", object["project_name"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("resource_id", object["resource_id"])
	d.Set("resource_type", object["resource_type"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))
	d.Set("traffic_path", object["traffic_path"])
	d.Set("traffic_type", object["traffic_type"])

	return nil
}

func resourceAlicloudVpcFlowLogUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyFlowLogAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["FlowLogId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("flow_log_name") {
		update = true
		if v, ok := d.GetOk("flow_log_name"); ok {
			request["FlowLogName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("aggregation_interval") {
		update = true
		if v, ok := d.GetOk("aggregation_interval"); ok {
			request["AggregationInterval"] = v
		}
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("description")
		d.SetPartial("flow_log_name")
		d.SetPartial("aggregation_interval")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["NewResourceGroupId"] = v
		}
	}
	request["ResourceType"] = "FLOWLOG"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("resource_group_id")
	}

	if d.HasChange("status") {
		client := meta.(*connectivity.AliyunClient)
		vpcServiceV2 := VpcServiceV2{client}
		object, err := vpcServiceV2.DescribeVpcFlowLog(d.Id())
		if err != nil {
			return WrapError(err)
		}

		target := d.Get("status").(string)
		if object["status"].(string) != target {
			if target == "Active" {
				action = "ActiveFlowLog"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["FlowLogId"] = d.Id()
				request["RegionId"] = client.RegionId

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{"TaskConflict", "LastTokenProcessing", "IncorrectStatus.%s", "IncorrectStatus.flowlog", "InvalidStatus", "OperationConflict", "ServiceUnavailable", "SystemBusy"}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
				{
					vpcServiceV2 := VpcServiceV2{client}
					stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcFlowLogStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
			if target == "Inactive" {
				action = "DeactiveFlowLog"
				conn, err = client.NewVpcClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["FlowLogId"] = d.Id()
				request["RegionId"] = client.RegionId

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{"TaskConflict", "LastTokenProcessing", "IncorrectStatus.%s", "IncorrectStatus.flowlog", "InvalidStatus", "OperationConflict", "ServiceUnavailable", "SystemBusy"}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
				{
					vpcServiceV2 := VpcServiceV2{client}
					stateConf := BuildStateConf([]string{}, []string{"Inactive"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcFlowLogStateRefreshFunc(d.Id(), []string{}))
					if _, err := stateConf.WaitForState(); err != nil {
						return WrapErrorf(err, IdMsg, d.Id())
					}
				}

			}
		}
	}

	update = false
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "FLOWLOG"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcFlowLogRead(d, meta)
}

func resourceAlicloudVpcFlowLogDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteFlowLog"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["FlowLogId"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationFailed.LastTokenProcessing", "LastTokenProcessing", "InvalidHdMonitorStatus", "IncorrectStatus.%s", "IncorrectStatus.flowlog", "InvalidStatus", "OperationConflict", "SystemBusy", "ServiceUnavailable"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"Instance.IsNotAvailable", "Instance.IsNotPostPay"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcFlowLogStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
