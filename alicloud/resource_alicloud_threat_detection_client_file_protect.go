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

func resourceAliCloudThreatDetectionClientFileProtect() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionClientFileProtectCreate,
		Read:   resourceAliCloudThreatDetectionClientFileProtectRead,
		Update: resourceAliCloudThreatDetectionClientFileProtectUpdate,
		Delete: resourceAliCloudThreatDetectionClientFileProtectDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alert_level": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"file_ops": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"file_paths": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"proc_paths": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"rule_action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rule_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"switch_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudThreatDetectionClientFileProtectCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateFileProtectRule"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewThreatdetectionClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["RuleName"] = d.Get("rule_name")
	if v, ok := d.GetOk("proc_paths"); ok {
		procPathsMaps := v.([]interface{})
		request["ProcPaths"] = procPathsMaps
	}

	if v, ok := d.GetOk("file_paths"); ok {
		filePathsMaps := v.([]interface{})
		request["FilePaths"] = filePathsMaps
	}

	if v, ok := d.GetOk("file_ops"); ok {
		fileOpsMaps := v.([]interface{})
		request["FileOps"] = fileOpsMaps
	}

	request["RuleAction"] = d.Get("rule_action")
	if v, ok := d.GetOk("alert_level"); ok {
		request["AlertLevel"] = v
	}
	if v, ok := d.GetOk("switch_id"); ok {
		request["SwitchId"] = v
	}
	if v, ok := d.GetOk("status"); ok {
		request["Status"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-12-03"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_threat_detection_client_file_protect", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["RuleId"]))

	return resourceAliCloudThreatDetectionClientFileProtectRead(d, meta)
}

func resourceAliCloudThreatDetectionClientFileProtectRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionClientFileProtect(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_client_file_protect DescribeThreatDetectionClientFileProtect Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("alert_level", objectRaw["AlertLevel"])
	d.Set("rule_action", objectRaw["Action"])
	d.Set("rule_name", objectRaw["RuleName"])
	d.Set("status", objectRaw["Status"])
	d.Set("switch_id", objectRaw["SwitchId"])
	fileOps1Raw := make([]interface{}, 0)
	if objectRaw["FileOps"] != nil {
		fileOps1Raw = objectRaw["FileOps"].([]interface{})
	}

	d.Set("file_ops", fileOps1Raw)
	filePaths1Raw := make([]interface{}, 0)
	if objectRaw["FilePaths"] != nil {
		filePaths1Raw = objectRaw["FilePaths"].([]interface{})
	}

	d.Set("file_paths", filePaths1Raw)
	procPaths1Raw := make([]interface{}, 0)
	if objectRaw["ProcPaths"] != nil {
		procPaths1Raw = objectRaw["ProcPaths"].([]interface{})
	}

	d.Set("proc_paths", procPaths1Raw)

	return nil
}

func resourceAliCloudThreatDetectionClientFileProtectUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	action := "UpdateFileProtectRule"
	conn, err := client.NewThreatdetectionClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["Id"] = d.Id()
	if d.HasChange("rule_name") {
		update = true
	}
	request["RuleName"] = d.Get("rule_name")
	if d.HasChange("proc_paths") {
		update = true
		if v, ok := d.GetOk("proc_paths"); ok {
			procPathsMaps := v.([]interface{})
			request["ProcPaths"] = procPathsMaps
		}
	}

	if d.HasChange("file_paths") {
		update = true
		if v, ok := d.GetOk("file_paths"); ok {
			filePathsMaps := v.([]interface{})
			request["FilePaths"] = filePathsMaps
		}
	}

	if d.HasChange("file_ops") {
		update = true
		if v, ok := d.GetOk("file_ops"); ok {
			fileOpsMaps := v.([]interface{})
			request["FileOps"] = fileOpsMaps
		}
	}

	if d.HasChange("rule_action") {
		update = true
	}
	request["RuleAction"] = d.Get("rule_action")
	if d.HasChange("alert_level") {
		update = true
		request["AlertLevel"] = d.Get("alert_level")
	}

	if d.HasChange("status") {
		update = true
		request["Status"] = d.Get("status")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-12-03"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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

	return resourceAliCloudThreatDetectionClientFileProtectRead(d, meta)
}

func resourceAliCloudThreatDetectionClientFileProtectDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteFileProtectRule"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewThreatdetectionClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["Id.1"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2018-12-03"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
