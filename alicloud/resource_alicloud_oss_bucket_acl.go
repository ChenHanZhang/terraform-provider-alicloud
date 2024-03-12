// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudOssBucketAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOssBucketAclCreate,
		Read:   resourceAliCloudOssBucketAclRead,
		Update: resourceAliCloudOssBucketAclUpdate,
		Delete: resourceAliCloudOssBucketAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"acl": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudOssBucketAclCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/?acl")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	hostMap := make(map[string]*string)
	conn, err := client.NewOssClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	hostMap["bucket"] = StringPointer(d.Get("bucket").(string))

	request["x-oss-acl"] = d.Get("acl")
	body = request
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.Execute(genXmlParam("PutBucketAcl", "PUT", "2019-05-17", action), &openapi.OpenApiRequest{Query: query, Body: body, HostMap: hostMap}, &util.RuntimeOptions{})

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_oss_bucket_acl", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(*hostMap["bucket"]))

	return resourceAliCloudOssBucketAclRead(d, meta)
}

func resourceAliCloudOssBucketAclRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ossServiceV2 := OssServiceV2{client}

	objectRaw, err := ossServiceV2.DescribeOssBucketAcl(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_oss_bucket_acl DescribeOssBucketAcl Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("acl", objectRaw["Grant"])

	d.Set("bucket", d.Id())

	return nil
}

func resourceAliCloudOssBucketAclUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	var hostMap map[string]*string
	update := false
	action := fmt.Sprintf("/?acl")
	conn, err := client.NewOssClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	hostMap = make(map[string]*string)
	hostMap["bucket"] = StringPointer(d.Id())
	if d.HasChange("acl") {
		update = true
	}
	request["x-oss-acl"] = d.Get("acl")
	body = request
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.Execute(genXmlParam("PutBucketAcl", "PUT", "2019-05-17", action), &openapi.OpenApiRequest{Query: query, Body: body, HostMap: hostMap}, &util.RuntimeOptions{})

			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
	}

	return resourceAliCloudOssBucketAclRead(d, meta)
}

func resourceAliCloudOssBucketAclDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Bucket Acl. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
