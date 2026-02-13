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

func resourceAliCloudOssBucketUserDefinedLogFields() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOssBucketUserDefinedLogFieldsCreate,
		Read:   resourceAliCloudOssBucketUserDefinedLogFieldsRead,
		Update: resourceAliCloudOssBucketUserDefinedLogFieldsUpdate,
		Delete: resourceAliCloudOssBucketUserDefinedLogFieldsDelete,
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
			"header_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"param_set": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAliCloudOssBucketUserDefinedLogFieldsCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/?userDefinedLogFieldsConfig")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["bucket"] = StringPointer(d.Get("bucket").(string))

	userDefinedLogFieldsConfiguration := make(map[string]interface{})

	if v := d.Get("header_set"); !IsNil(v) {
		headerSet := make(map[string]interface{})
		headerSet1, _ := jsonpath.Get("$", d.Get("header_set"))
		if headerSet1 != nil && headerSet1 != "" {
			headerSet["header"] = convertToInterfaceArray(headerSet1)
		}

		if len(headerSet) > 0 {
			userDefinedLogFieldsConfiguration["HeaderSet"] = headerSet
		}
	}

	if v := d.Get("param_set"); !IsNil(v) {
		paramSet := make(map[string]interface{})
		paramSet1, _ := jsonpath.Get("$", d.Get("param_set"))
		if paramSet1 != nil && paramSet1 != "" {
			paramSet["parameter"] = convertToInterfaceArray(paramSet1)
		}

		if len(paramSet) > 0 {
			userDefinedLogFieldsConfiguration["ParamSet"] = paramSet
		}
	}

	request["UserDefinedLogFieldsConfiguration"] = userDefinedLogFieldsConfiguration

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.Do("Oss", xmlParam("PUT", "2019-05-17", "PutUserDefinedLogFieldsConfig", action), query, body, nil, hostMap, false)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_oss_bucket_user_defined_log_fields", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(*hostMap["bucket"]))

	return resourceAliCloudOssBucketUserDefinedLogFieldsRead(d, meta)
}

func resourceAliCloudOssBucketUserDefinedLogFieldsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ossServiceV2 := OssServiceV2{client}

	objectRaw, err := ossServiceV2.DescribeOssBucketUserDefinedLogFields(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_oss_bucket_user_defined_log_fields DescribeOssBucketUserDefinedLogFields Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("header_set", headerRaw)

	d.Set("param_set", parameterRaw)

	d.Set("bucket", d.Id())

	return nil
}

func resourceAliCloudOssBucketUserDefinedLogFieldsUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	action := fmt.Sprintf("/?userDefinedLogFieldsConfig")
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap := make(map[string]*string)
	hostMap["bucket"] = StringPointer(d.Id())

	userDefinedLogFieldsConfiguration := make(map[string]interface{})

	if d.HasChange("header_set") {
		update = true
	}
	if v := d.Get("header_set"); !IsNil(v) || d.HasChange("header_set") {
		headerSet := make(map[string]interface{})
		headerSet1, _ := jsonpath.Get("$", d.Get("header_set"))
		if headerSet1 != nil && headerSet1 != "" {
			headerSet["header"] = convertToInterfaceArray(headerSet1)
		}

		if len(headerSet) > 0 {
			userDefinedLogFieldsConfiguration["HeaderSet"] = headerSet
		}
	}

	if d.HasChange("param_set") {
		update = true
	}
	if v := d.Get("param_set"); !IsNil(v) || d.HasChange("param_set") {
		paramSet := make(map[string]interface{})
		paramSet1, _ := jsonpath.Get("$", d.Get("param_set"))
		if paramSet1 != nil && paramSet1 != "" {
			paramSet["parameter"] = convertToInterfaceArray(paramSet1)
		}

		if len(paramSet) > 0 {
			userDefinedLogFieldsConfiguration["ParamSet"] = paramSet
		}
	}

	request["UserDefinedLogFieldsConfiguration"] = userDefinedLogFieldsConfiguration

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.Do("Oss", xmlParam("PUT", "2019-05-17", "PutUserDefinedLogFieldsConfig", action), query, body, nil, hostMap, false)
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

	return resourceAliCloudOssBucketUserDefinedLogFieldsRead(d, meta)
}

func resourceAliCloudOssBucketUserDefinedLogFieldsDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := fmt.Sprintf("/?userDefinedLogFieldsConfig")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	hostMap := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	hostMap["bucket"] = StringPointer(d.Id())

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.Do("Oss", xmlParam("DELETE", "2019-05-17", "DeleteUserDefinedLogFieldsConfig", action), query, nil, nil, hostMap, false)
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
		if IsExpectedErrors(err, []string{"NoSuchBucket", "NoSuchUserDefinedLogFieldsConfig"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
