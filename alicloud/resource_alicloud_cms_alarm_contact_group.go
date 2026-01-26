// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCloudMonitorServiceAlarmContactGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudMonitorServiceAlarmContactGroupCreate,
		Read:   resourceAliCloudCloudMonitorServiceAlarmContactGroupRead,
		Update: resourceAliCloudCloudMonitorServiceAlarmContactGroupUpdate,
		Delete: resourceAliCloudCloudMonitorServiceAlarmContactGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alarm_contact_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"contact_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"describe": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_subscribed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudCloudMonitorServiceAlarmContactGroupCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "PutContactGroup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("alarm_contact_group_name"); ok {
		request["ContactGroupName"] = v
	}

	if v, ok := d.GetOk("describe"); ok {
		request["Describe"] = v
	}
	if v, ok := d.GetOkExists("enable_subscribed"); ok {
		request["EnableSubscribed"] = v
	}
	if v, ok := d.GetOk("contact_names"); ok {
		contactNamesMapsArray := convertToInterfaceArray(v)

		request["ContactNames"] = contactNamesMapsArray
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cms_alarm_contact_group", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["ContactGroupName"]))

	return resourceAliCloudCloudMonitorServiceAlarmContactGroupRead(d, meta)
}

func resourceAliCloudCloudMonitorServiceAlarmContactGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudMonitorServiceServiceV2 := CloudMonitorServiceServiceV2{client}

	objectRaw, err := cloudMonitorServiceServiceV2.DescribeCloudMonitorServiceAlarmContactGroup(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cms_alarm_contact_group DescribeCloudMonitorServiceAlarmContactGroup Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("describe", objectRaw["Describe"])
	d.Set("enable_subscribed", objectRaw["EnableSubscribed"])
	d.Set("alarm_contact_group_name", objectRaw["Name"])

	contactRaw, _ := jsonpath.Get("$.Contacts.Contact", objectRaw)
	d.Set("contact_names", contactRaw)

	d.Set("alarm_contact_group_name", d.Id())

	return nil
}

func resourceAliCloudCloudMonitorServiceAlarmContactGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "PutContactGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ContactGroupName"] = d.Id()

	if d.HasChange("describe") {
		update = true
		request["Describe"] = d.Get("describe")
	}

	if d.HasChange("enable_subscribed") {
		update = true
		request["EnableSubscribed"] = d.Get("enable_subscribed")
	}

	if d.HasChange("contact_names") {
		update = true
		if v, ok := d.GetOk("contact_names"); ok || d.HasChange("contact_names") {
			contactNamesMapsArray := convertToInterfaceArray(v)

			request["ContactNames"] = contactNamesMapsArray
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
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
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	return resourceAliCloudCloudMonitorServiceAlarmContactGroupRead(d, meta)
}

func resourceAliCloudCloudMonitorServiceAlarmContactGroupDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteContactGroup"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ContactGroupName"] = d.Id()

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
		if IsExpectedErrors(err, []string{"400", "403", "404"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
