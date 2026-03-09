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

func resourceAliCloudPaiAbTestCrowd() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPaiAbTestCrowdCreate,
		Read:   resourceAliCloudPaiAbTestCrowdRead,
		Update: resourceAliCloudPaiAbTestCrowdUpdate,
		Delete: resourceAliCloudPaiAbTestCrowdDelete,
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
			"crowd_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPaiAbTestCrowdCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/api/v1/crowds")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	query["Regionid"] = StringPointer(client.RegionId)

	request["WorkspaceId"] = d.Get("workspace_id")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("label"); ok {
		request["Label"] = v
	}
	request["Users"] = d.Get("users")
	request["Name"] = d.Get("crowd_name")
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("PAIABTest", "2024-01-19", action, query, nil, body, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"503"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_pai_ab_test_crowd", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["CrowdId"]))

	return resourceAliCloudPaiAbTestCrowdRead(d, meta)
}

func resourceAliCloudPaiAbTestCrowdRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	paiAbTestServiceV2 := PaiAbTestServiceV2{client}

	objectRaw, err := paiAbTestServiceV2.DescribePaiAbTestCrowd(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_pai_ab_test_crowd DescribePaiAbTestCrowd Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["GmtCreateTime"])
	d.Set("crowd_name", objectRaw["Name"])
	d.Set("description", objectRaw["Description"])
	d.Set("label", objectRaw["Label"])
	d.Set("users", objectRaw["Users"])
	d.Set("workspace_id", objectRaw["WorkspaceId"])

	return nil
}

func resourceAliCloudPaiAbTestCrowdUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	CrowdId := d.Id()
	action := fmt.Sprintf("/api/v1/crowds/%s", CrowdId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)
	if d.HasChange("description") {
		update = true
	}
	if v, ok := d.GetOk("description"); ok || d.HasChange("description") {
		request["Description"] = v
	}
	if d.HasChange("label") {
		update = true
	}
	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		request["Label"] = v
	}
	if d.HasChange("users") {
		update = true
	}
	request["Users"] = d.Get("users")
	if d.HasChange("crowd_name") {
		update = true
	}
	request["Name"] = d.Get("crowd_name")
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("PAIABTest", "2024-01-19", action, query, nil, body, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"503"}) || NeedRetry(err) {
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

	return resourceAliCloudPaiAbTestCrowdRead(d, meta)
}

func resourceAliCloudPaiAbTestCrowdDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	CrowdId := d.Id()
	action := fmt.Sprintf("/api/v1/crowds/%s", CrowdId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	query["RegionId"] = StringPointer(client.RegionId)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("PAIABTest", "2024-01-19", action, query, nil, nil, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"503"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if IsExpectedErrors(err, []string{"404"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
