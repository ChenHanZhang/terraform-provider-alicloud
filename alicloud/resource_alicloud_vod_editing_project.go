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

func resourceAliCloudVodEditingProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudVodEditingProjectCreate,
		Read:   resourceAliCloudVodEditingProjectRead,
		Update: resourceAliCloudVodEditingProjectUpdate,
		Delete: resourceAliCloudVodEditingProjectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cover_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"editing_project_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timeline": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudVodEditingProjectCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "AddEditingProject"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["Title"] = d.Get("title")
	if v, ok := d.GetOk("timeline"); ok {
		request["Timeline"] = v
	}
	if v, ok := d.GetOk("editing_project_name"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("cover_url"); ok {
		request["CoverURL"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("vod", "2017-03-21", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vod_editing_project", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Project.ProjectId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudVodEditingProjectRead(d, meta)
}

func resourceAliCloudVodEditingProjectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vodServiceV2 := VodServiceV2{client}

	objectRaw, err := vodServiceV2.DescribeVodEditingProject(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vod_editing_project DescribeVodEditingProject Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cover_url", objectRaw["CoverURL"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("editing_project_name", objectRaw["Description"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("status", objectRaw["Status"])
	d.Set("timeline", objectRaw["Timeline"])
	d.Set("title", objectRaw["Title"])

	return nil
}

func resourceAliCloudVodEditingProjectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateEditingProject"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ProjectId"] = d.Id()

	if d.HasChange("title") {
		update = true
	}
	request["Title"] = d.Get("title")
	if d.HasChange("timeline") {
		update = true
		request["Timeline"] = d.Get("timeline")
	}

	if d.HasChange("editing_project_name") {
		update = true
		request["Description"] = d.Get("editing_project_name")
	}

	if d.HasChange("cover_url") {
		update = true
		request["CoverURL"] = d.Get("cover_url")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("vod", "2017-03-21", action, query, request, true)
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

	return resourceAliCloudVodEditingProjectRead(d, meta)
}

func resourceAliCloudVodEditingProjectDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteEditingProject"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ProjectIds"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("vod", "2017-03-21", action, query, request, true)
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
