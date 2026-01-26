// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCloudMonitorServiceMonitoringAgentProcess() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudMonitorServiceMonitoringAgentProcessCreate,
		Read:   resourceAliCloudCloudMonitorServiceMonitoringAgentProcessRead,
		Delete: resourceAliCloudCloudMonitorServiceMonitoringAgentProcessDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"process_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"process_user": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudCloudMonitorServiceMonitoringAgentProcessCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateMonitorAgentProcess"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}

	request["ProcessName"] = d.Get("process_name")
	if v, ok := d.GetOk("process_user"); ok {
		request["ProcessUser"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Cms", "2019-01-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cloud_monitor_service_monitoring_agent_process", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], response["Id"]))

	return resourceAliCloudCloudMonitorServiceMonitoringAgentProcessRead(d, meta)
}

func resourceAliCloudCloudMonitorServiceMonitoringAgentProcessRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudMonitorServiceServiceV2 := CloudMonitorServiceServiceV2{client}

	objectRaw, err := cloudMonitorServiceServiceV2.DescribeCloudMonitorServiceMonitoringAgentProcess(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cloud_monitor_service_monitoring_agent_process DescribeCloudMonitorServiceMonitoringAgentProcess Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("process_name", objectRaw["ProcessName"])
	d.Set("process_user", objectRaw["ProcessUser"])
	d.Set("instance_id", objectRaw["InstanceId"])
	d.Set("process_id", objectRaw["ProcessId"])

	return nil
}

func resourceAliCloudCloudMonitorServiceMonitoringAgentProcessDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteMonitoringAgentProcess"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["ProcessId"] = parts[1]

	request["ProcessName"] = d.Get("process_name")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Cms", "2019-01-01", action, query, request, true)
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
