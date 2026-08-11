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

func resourceAliCloudEnsBucketLifecycle() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEnsBucketLifecycleCreate,
		Read:   resourceAliCloudEnsBucketLifecycleRead,
		Update: resourceAliCloudEnsBucketLifecycleUpdate,
		Delete: resourceAliCloudEnsBucketLifecycleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allow_same_action_overlap": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"created_before_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"expiration_days": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"Enabled", "Disabled"}, false),
			},
		},
	}
}

func resourceAliCloudEnsBucketLifecycleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "PutBucketLifecycle"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("bucket_name"); ok {
		request["BucketName"] = v
	}
	if v, ok := d.GetOk("rule_id"); ok {
		request["RuleId"] = v
	}

	if v, ok := d.GetOk("prefix"); ok {
		request["Prefix"] = v
	}
	if v, ok := d.GetOk("created_before_date"); ok {
		request["CreatedBeforeDate"] = v
	}
	if v, ok := d.GetOkExists("expiration_days"); ok {
		request["ExpirationDays"] = v
	}
	if v, ok := d.GetOkExists("allow_same_action_overlap"); ok {
		request["AllowSameActionOverlap"] = v
	}
	request["Status"] = d.Get("status")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ens_bucket_lifecycle", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["BucketName"], response["RuleId"]))

	return resourceAliCloudEnsBucketLifecycleRead(d, meta)
}

func resourceAliCloudEnsBucketLifecycleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ensServiceV2 := EnsServiceV2{client}

	objectRaw, err := ensServiceV2.DescribeEnsBucketLifecycle(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ens_bucket_lifecycle DescribeEnsBucketLifecycle Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("prefix", objectRaw["Prefix"])
	d.Set("status", objectRaw["Status"])
	d.Set("rule_id", objectRaw["ID"])

	expirationRawObj, _ := jsonpath.Get("$.Expiration", objectRaw)
	expirationRaw := make(map[string]interface{})
	if expirationRawObj != nil {
		expirationRaw = expirationRawObj.(map[string]interface{})
	}
	d.Set("created_before_date", expirationRaw["CreatedBeforeDate"])
	d.Set("expiration_days", formatInt(expirationRaw["Days"]))

	parts := strings.Split(d.Id(), ":")
	d.Set("bucket_name", parts[0])

	return nil
}

func resourceAliCloudEnsBucketLifecycleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "PutBucketLifecycle"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["BucketName"] = parts[0]
	request["RuleId"] = parts[1]

	if d.HasChange("prefix") {
		update = true
		request["Prefix"] = d.Get("prefix")
	}

	if d.HasChange("created_before_date") {
		update = true
		request["CreatedBeforeDate"] = d.Get("created_before_date")
	}

	if d.HasChange("expiration_days") {
		update = true
		request["ExpirationDays"] = d.Get("expiration_days")
	}

	if v, ok := d.GetOkExists("allow_same_action_overlap"); ok {
		request["AllowSameActionOverlap"] = v
	}
	if d.HasChange("status") {
		update = true
	}
	request["Status"] = d.Get("status")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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

	return resourceAliCloudEnsBucketLifecycleRead(d, meta)
}

func resourceAliCloudEnsBucketLifecycleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteBucketLifecycle"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["BucketName"] = parts[0]
	request["RuleId"] = parts[1]

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ens", "2017-11-10", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"NoSuchBucket"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
