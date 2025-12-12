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

func resourceAliCloudOssBucketWorm() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOssBucketWormCreate,
		Read:   resourceAliCloudOssBucketWormRead,
		Update: resourceAliCloudOssBucketWormUpdate,
		Delete: resourceAliCloudOssBucketWormDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"retention_period_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"worm_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudOssBucketWormCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/?worm")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["bucket"] = StringPointer(d.Get("bucket").(string))

	initiateWormConfiguration := make(map[string]interface{})

	if v := d.Get("retention_period_in_days"); !IsNil(v) {
		initiateWormConfiguration["RetentionPeriodInDays"] = v
		request["InitiateWormConfiguration"] = initiateWormConfiguration
	}

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Oss", xmlParam("POST", "2019-05-17", "InitiateBucketWorm", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_oss_bucket_worm", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", *hostMap["bucket"], response["x-oss-worm-id"]))

	return resourceAliCloudOssBucketWormUpdate(d, meta)
}

func resourceAliCloudOssBucketWormRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ossServiceV2 := OssServiceV2{client}

	objectRaw, err := ossServiceV2.DescribeOssBucketWorm(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_oss_bucket_worm DescribeOssBucketWorm Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreationDate"])
	d.Set("retention_period_in_days", objectRaw["RetentionPeriodInDays"])
	d.Set("status", objectRaw["State"])
	d.Set("worm_id", objectRaw["WormId"])

	parts := strings.Split(d.Id(), ":")
	d.Set("bucket", parts[0])

	return nil
}

func resourceAliCloudOssBucketWormUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var header map[string]*string
	var query map[string]*string
	var body map[string]interface{}
	update := false

	ossServiceV2 := OssServiceV2{client}
	objectRaw, _ := ossServiceV2.DescribeOssBucketWorm(d.Id())

	if d.HasChange("status") {
		var err error
		target := d.Get("status").(string)

		currentStatus, err := jsonpath.Get("State", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "State", objectRaw)
		}
		if fmt.Sprint(currentStatus) != target {
			if target == "Locked" {
				parts := strings.Split(d.Id(), ":")
				action := fmt.Sprintf("/")
				request = make(map[string]interface{})
				query = make(map[string]*string)
				body = make(map[string]interface{})
				hostMap := make(map[string]*string)
				hostMap["bucket"] = StringPointer(parts[0])
				query["wormId"] = StringPointer(parts[1])

				body = request
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.Do("Oss", xmlParam("POST", "2019-05-17", "CompleteBucketWorm", action), query, body, nil, hostMap, false)
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
		}
	}

	var err error
	parts := strings.Split(d.Id(), ":")
	action := fmt.Sprintf("/?wormExtend")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	hostMap["bucket"] = StringPointer(parts[0])
	query["wormId"] = StringPointer(parts[1])

	if d.HasChange("retention_period_in_days") {
		update = true
	}
	extendWormConfiguration := make(map[string]interface{})

	if v := d.Get("retention_period_in_days"); !IsNil(v) || d.HasChange("retention_period_in_days") {
		if v, ok := d.GetOkExists("retention_period_in_days"); ok {
			extendWormConfiguration["RetentionPeriodInDays"] = v
		}

		request["ExtendWormConfiguration"] = extendWormConfiguration
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Oss", xmlParam("POST", "2019-05-17", "ExtendBucketWorm", action), query, body, nil, hostMap, false)
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
		ossServiceV2 := OssServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("retention_period_in_days"))}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ossServiceV2.OssBucketWormStateRefreshFunc(d.Id(), "RetentionPeriodInDays", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudOssBucketWormRead(d, meta)
}

func resourceAliCloudOssBucketWormDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	enableDelete := true
	if v, ok := d.GetOkExists("status"); ok {
		if InArray(fmt.Sprint(v), []string{"Locked"}) {
			enableDelete = false
			log.Printf("[WARN] Cannot destroy resource alicloud_oss_bucket_worm which status valued Locked. Terraform will remove this resource from the state file, however resources may remain.")
		}
	}
	if enableDelete {
		action := fmt.Sprintf("/?worm")
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]*string)
		hostMap := make(map[string]*string)
		var err error
		request = make(map[string]interface{})
		hostMap["bucket"] = StringPointer(parts[0])

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
			response, err = client.Do("Oss", xmlParam("DELETE", "2019-05-17", "AbortBucketWorm", action), query, nil, nil, hostMap, false)
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
			if IsExpectedErrors(err, []string{"NoSuchBucket", "NoSuchWORMConfiguration"}) || NotFoundError(err) {
				return nil
			}
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}

	}
	return nil
}
