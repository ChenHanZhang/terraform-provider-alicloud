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

func resourceAliCloudEfloErAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEfloErAttachmentCreate,
		Read:   resourceAliCloudEfloErAttachmentRead,
		Update: resourceAliCloudEfloErAttachmentUpdate,
		Delete: resourceAliCloudEfloErAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_receive_all_route": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"er_attachment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"er_attachment_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"er_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"VPD", "VCC"}, false),
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_tenant_id": {
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

func resourceAliCloudEfloErAttachmentCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateErAttachment"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("er_id"); ok {
		request["ErId"] = v
	}
	request["RegionId"] = client.RegionId

	request["AutoReceiveAllRoute"] = d.Get("auto_receive_all_route")
	request["InstanceId"] = d.Get("instance_id")
	request["InstanceType"] = d.Get("instance_type")
	request["ResourceTenantId"] = d.Get("resource_tenant_id")
	request["ErAttachmentName"] = d.Get("er_attachment_name")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("eflo", "2022-05-30", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eflo_er_attachment", action, AlibabaCloudSdkGoERROR)
	}

	ContentErAttachmentIdVar, _ := jsonpath.Get("$.Content.ErAttachmentId", response)
	d.SetId(fmt.Sprintf("%v:%v", request["ErId"], ContentErAttachmentIdVar))

	efloServiceV2 := EfloServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, efloServiceV2.EfloErAttachmentStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudEfloErAttachmentRead(d, meta)
}

func resourceAliCloudEfloErAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	efloServiceV2 := EfloServiceV2{client}

	objectRaw, err := efloServiceV2.DescribeEfloErAttachment(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eflo_er_attachment DescribeEfloErAttachment Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("auto_receive_all_route", objectRaw["AutoReceiveAllRoute"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("er_attachment_name", objectRaw["ErAttachmentName"])
	d.Set("instance_id", objectRaw["InstanceId"])
	d.Set("instance_type", objectRaw["InstanceType"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("resource_tenant_id", objectRaw["ResourceTenantId"])
	d.Set("status", objectRaw["Status"])
	d.Set("er_attachment_id", objectRaw["ErAttachmentId"])
	d.Set("er_id", objectRaw["ErId"])

	return nil
}

func resourceAliCloudEfloErAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdateErAttachment"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ErAttachmentId"] = parts[1]
	request["ErId"] = parts[0]
	request["RegionId"] = client.RegionId
	if d.HasChange("er_attachment_name") {
		update = true
		request["ErAttachmentName"] = d.Get("er_attachment_name")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("eflo", "2022-05-30", action, query, request, true)
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

	return resourceAliCloudEfloErAttachmentRead(d, meta)
}

func resourceAliCloudEfloErAttachmentDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteErAttachment"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ErAttachmentId"] = parts[1]
	request["ErId"] = parts[0]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("eflo", "2022-05-30", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"1003"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	efloServiceV2 := EfloServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, efloServiceV2.EfloErAttachmentStateRefreshFunc(d.Id(), "ErAttachmentId", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
