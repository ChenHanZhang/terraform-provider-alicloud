// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAliKafkaScheduledScalingAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAliKafkaScheduledScalingAttachmentCreate,
		Read:   resourceAliCloudAliKafkaScheduledScalingAttachmentRead,
		Update: resourceAliCloudAliKafkaScheduledScalingAttachmentUpdate,
		Delete: resourceAliCloudAliKafkaScheduledScalingAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"duration_minutes": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"first_scheduled_time": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repeat_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"reserved_pub_flow": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"reserved_sub_flow": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schedule_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"weekly_types": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAliCloudAliKafkaScheduledScalingAttachmentCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateScheduledScalingRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	request["RegionId"] = client.RegionId

	request["RuleName"] = d.Get("rule_name")
	request["ReservedPubFlow"] = d.Get("reserved_pub_flow")
	request["ReservedSubFlow"] = d.Get("reserved_sub_flow")
	request["ScheduleType"] = d.Get("schedule_type")
	request["DurationMinutes"] = d.Get("duration_minutes")
	request["TimeZone"] = d.Get("time_zone")
	request["FirstScheduledTime"] = d.Get("first_scheduled_time")
	if v, ok := d.GetOkExists("enable"); ok {
		request["Enable"] = v
	}
	if v, ok := d.GetOk("repeat_type"); ok {
		request["RepeatType"] = v
	}
	if v, ok := d.GetOk("weekly_types"); ok {
		weeklyTypesMapsArray := convertToInterfaceArray(v)

		weeklyTypesMapsJson, err := json.Marshal(weeklyTypesMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["WeeklyTypes"] = string(weeklyTypesMapsJson)
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ali_kafka_scheduled_scaling_attachment", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["InstanceId"]))

	return resourceAliCloudAliKafkaScheduledScalingAttachmentRead(d, meta)
}

func resourceAliCloudAliKafkaScheduledScalingAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	aliKafkaServiceV2 := AliKafkaServiceV2{client}

	objectRaw, err := aliKafkaServiceV2.DescribeAliKafkaScheduledScalingAttachment(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ali_kafka_scheduled_scaling_attachment DescribeAliKafkaScheduledScalingAttachment Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("duration_minutes", objectRaw["DurationMinutes"])
	d.Set("enable", objectRaw["Enable"])
	d.Set("first_scheduled_time", objectRaw["FirstScheduledTime"])
	d.Set("repeat_type", objectRaw["RepeatType"])
	d.Set("reserved_pub_flow", objectRaw["ReservedPubFlow"])
	d.Set("reserved_sub_flow", objectRaw["ReservedSubFlow"])
	d.Set("rule_name", objectRaw["RuleName"])
	d.Set("schedule_type", objectRaw["ScheduleType"])
	d.Set("time_zone", objectRaw["TimeZone"])

	weeklyTypesRaw, _ := jsonpath.Get("$.WeeklyTypes.WeeklyTypes", objectRaw)
	d.Set("weekly_types", weeklyTypesRaw)

	d.Set("instance_id", d.Id())

	return nil
}

func resourceAliCloudAliKafkaScheduledScalingAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyScheduledScalingRule"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("rule_name") {
		update = true
	}
	request["RuleName"] = d.Get("rule_name")
	if d.HasChange("enable") {
		update = true
	}
	request["Enable"] = d.Get("enable")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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

	return resourceAliCloudAliKafkaScheduledScalingAttachmentRead(d, meta)
}

func resourceAliCloudAliKafkaScheduledScalingAttachmentDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteScheduledScalingRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["RuleName"] = d.Get("rule_name")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
