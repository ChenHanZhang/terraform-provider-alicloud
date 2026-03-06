// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudOpenSearchAppGroupCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOpenSearchAppGroupCredentialCreate,
		Read:   resourceAliCloudOpenSearchAppGroupCredentialRead,
		Update: resourceAliCloudOpenSearchAppGroupCredentialUpdate,
		Delete: resourceAliCloudOpenSearchAppGroupCredentialDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"app_group_credential_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudOpenSearchAppGroupCredentialCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	appGroupIdentity := d.Get("app_group_id")
	action := fmt.Sprintf("/v4/openapi/app-groups/%s/credentials", appGroupIdentity)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["type"] = d.Get("type")
	if v, ok := d.GetOkExists("dry_run"); ok {
		query["dryRun"] = StringPointer(strconv.FormatBool(v.(bool)))
	}

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("OpenSearch", "2017-12-25", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_open_search_app_group_credential", action, AlibabaCloudSdkGoERROR)
	}

	resultappGroupIdVar, _ := jsonpath.Get("$.result.appGroupId", response)
	resulttokenVar, _ := jsonpath.Get("$.result.token", response)
	d.SetId(fmt.Sprintf("%v:%v", resultappGroupIdVar, resulttokenVar))

	return resourceAliCloudOpenSearchAppGroupCredentialUpdate(d, meta)
}

func resourceAliCloudOpenSearchAppGroupCredentialRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	openSearchServiceV2 := OpenSearchServiceV2{client}

	objectRaw, err := openSearchServiceV2.DescribeOpenSearchAppGroupCredential(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_open_search_app_group_credential DescribeOpenSearchAppGroupCredential Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("enabled", objectRaw["enabled"])
	d.Set("type", objectRaw["type"])
	d.Set("app_group_credential_id", objectRaw["token"])
	d.Set("app_group_id", objectRaw["appGroupId"])

	return nil
}

func resourceAliCloudOpenSearchAppGroupCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	token := parts[1]
	appGroupIdentity := parts[0]
	action := fmt.Sprintf("/v4/openapi/app-groups/%s/credentials/%s", appGroupIdentity, token)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})

	if d.HasChange("enabled") {
		update = true
	}
	if v, ok := d.GetOkExists("enabled"); ok || d.HasChange("enabled") {
		request["enabled"] = v
	}
	if v, ok := d.GetOkExists("dry_run"); ok {
		query["dryRun"] = StringPointer(strconv.FormatBool(v.(bool)))
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPut("OpenSearch", "2017-12-25", action, query, nil, body, true)
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

	return resourceAliCloudOpenSearchAppGroupCredentialRead(d, meta)
}

func resourceAliCloudOpenSearchAppGroupCredentialDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	token := parts[1]
	appGroupIdentity := parts[0]
	action := fmt.Sprintf("/v4/openapi/app-groups/%s/credentials/%s", appGroupIdentity, token)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("OpenSearch", "2017-12-25", action, query, nil, nil, true)
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
