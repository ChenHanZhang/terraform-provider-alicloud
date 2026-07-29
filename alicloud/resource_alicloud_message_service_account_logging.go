// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudMessageServiceAccountLogging() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudMessageServiceAccountLoggingCreate,
		Read:   resourceAliCloudMessageServiceAccountLoggingRead,
		Delete: resourceAliCloudMessageServiceAccountLoggingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"log_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"log_store_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"message_trace_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudMessageServiceAccountLoggingCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "SetAccountAttributes"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("project_name"); ok {
		request["ProjectName"] = v
	}
	if v, ok := d.GetOk("log_store_name"); ok {
		request["LogStoreName"] = v
	}
	if v, ok := d.GetOkExists("log_enabled"); ok {
		request["LogEnabled"] = v
	}
	if v, ok := d.GetOkExists("message_trace_enabled"); ok {
		request["MessageTraceEnabled"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Mns-open", "2022-01-19", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_message_service_account_logging", action, AlibabaCloudSdkGoERROR)
	}

	accountId, err := client.AccountId()
	d.SetId(accountId)

	return resourceAliCloudMessageServiceAccountLoggingRead(d, meta)
}

func resourceAliCloudMessageServiceAccountLoggingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	messageServiceServiceV2 := MessageServiceServiceV2{client}

	objectRaw, err := messageServiceServiceV2.DescribeMessageServiceAccountLogging(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_message_service_account_logging DescribeMessageServiceAccountLogging Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("log_enabled", objectRaw["LogEnabled"])
	d.Set("log_store_name", objectRaw["LogStoreName"])
	d.Set("message_trace_enabled", objectRaw["MessageTraceEnabled"])
	d.Set("project_name", objectRaw["ProjectName"])

	return nil
}

func resourceAliCloudMessageServiceAccountLoggingDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Account Logging. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
