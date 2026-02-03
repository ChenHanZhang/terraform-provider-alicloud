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

func resourceAliCloudPaiStudioAlgorithm() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPaiStudioAlgorithmCreate,
		Read:   resourceAliCloudPaiStudioAlgorithmRead,
		Update: resourceAliCloudPaiStudioAlgorithmUpdate,
		Delete: resourceAliCloudPaiStudioAlgorithmDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"algorithm_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"algorithm_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPaiStudioAlgorithmCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/api/v1/algorithms")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("workspace_id"); ok {
		request["WorkspaceId"] = v
	}
	if v, ok := d.GetOk("algorithm_description"); ok {
		request["AlgorithmDescription"] = v
	}
	if v, ok := d.GetOk("algorithm_name"); ok {
		request["AlgorithmName"] = v
	}
	if v, ok := d.GetOk("display_name"); ok {
		request["DisplayName"] = v
	}
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("PaiStudio", "2022-01-12", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_pai_studio_algorithm", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["AlgorithmId"]))

	return resourceAliCloudPaiStudioAlgorithmRead(d, meta)
}

func resourceAliCloudPaiStudioAlgorithmRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	paiStudioServiceV2 := PaiStudioServiceV2{client}

	objectRaw, err := paiStudioServiceV2.DescribePaiStudioAlgorithm(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_pai_studio_algorithm DescribePaiStudioAlgorithm Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("algorithm_description", objectRaw["AlgorithmDescription"])
	d.Set("algorithm_name", objectRaw["AlgorithmName"])
	d.Set("create_time", objectRaw["GmtCreateTime"])
	d.Set("display_name", objectRaw["DisplayName"])
	d.Set("workspace_id", objectRaw["WorkspaceId"])

	return nil
}

func resourceAliCloudPaiStudioAlgorithmUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	AlgorithmId := d.Id()
	action := fmt.Sprintf("/api/v1/algorithms/%s", AlgorithmId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("algorithm_description") {
		update = true
	}
	if v, ok := d.GetOk("algorithm_description"); ok || d.HasChange("algorithm_description") {
		request["AlgorithmDescription"] = v
	}
	if d.HasChange("display_name") {
		update = true
	}
	if v, ok := d.GetOk("display_name"); ok || d.HasChange("display_name") {
		request["DisplayName"] = v
	}
	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("PaiStudio", "2022-01-12", action, query, header, body, true)
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

	return resourceAliCloudPaiStudioAlgorithmRead(d, meta)
}

func resourceAliCloudPaiStudioAlgorithmDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	AlgorithmId := d.Id()
	action := fmt.Sprintf("/api/v1/algorithms/%s", AlgorithmId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("PaiStudio", "2022-01-12", action, query, nil, nil, true)
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
