// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionLogMeta() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionLogMetaCreate,
		Read:   resourceAliCloudThreatDetectionLogMetaRead,
		Update: resourceAliCloudThreatDetectionLogMetaUpdate,
		Delete: resourceAliCloudThreatDetectionLogMetaDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"log_meta_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"aegis-log-crack", "aegis-log-dns-query", "aegis-log-login", "aegis-log-network", "aegis-log-process", "aegis-snapshot-host", "aegis-snapshot-port", "aegis-snapshot-process", "local-dns", "sas-cspm-log", "sas-hc-log", "sas-log-dns", "sas-log-http", "sas-log-session", "sas-security-log", "sas-vul-log"}, false),
			},
			"status": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"disabled", "enabled"}, false),
			},
		},
	}
}

func resourceAliCloudThreatDetectionLogMetaCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ModifyLogMetaStatus"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewThreatdetectionClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["LogStore"] = d.Get("log_meta_id")

	request["Status"] = d.Get("status")
	request["From"] = "sas"
	request["Project"] = "sas"
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-12-03"), StringPointer("AK"), nil, request, &runtime)

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_threat_detection_log_meta", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["LogStore"]))

	return resourceAliCloudThreatDetectionLogMetaRead(d, meta)
}

func resourceAliCloudThreatDetectionLogMetaRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionLogMeta(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_log_meta DescribeThreatDetectionLogMeta Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("status", objectRaw["Status"])
	d.Set("log_meta_id", objectRaw["LogStore"])

	return nil
}

func resourceAliCloudThreatDetectionLogMetaUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	action := "ModifyLogMetaStatus"
	conn, err := client.NewThreatdetectionClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["LogStore"] = d.Id()
	if d.HasChange("status") {
		update = true
	}
	request["Status"] = d.Get("status")
	request["From"] = "sas"
	request["Project"] = "sas"
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-12-03"), StringPointer("AK"), nil, request, &runtime)

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

	return resourceAliCloudThreatDetectionLogMetaRead(d, meta)
}

func resourceAliCloudThreatDetectionLogMetaDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Log Meta. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
