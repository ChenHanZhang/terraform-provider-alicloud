// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionWebLockConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionWebLockConfigCreate,
		Read:   resourceAliCloudThreatDetectionWebLockConfigRead,
		Update: resourceAliCloudThreatDetectionWebLockConfigUpdate,
		Delete: resourceAliCloudThreatDetectionWebLockConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"config_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"defence_mode": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"exclusive_dir": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclusive_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclusive_file_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inclusive_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inclusive_file_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_backup_dir": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudThreatDetectionWebLockConfigCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ModifyWebLockCreateConfig"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("uuid"); ok {
		request["Uuid"] = v
	}

	if v, ok := d.GetOk("exclusive_file_type"); ok {
		request["ExclusiveFileType"] = v
	}
	if v, ok := d.GetOk("exclusive_file"); ok {
		request["ExclusiveFile"] = v
	}
	request["Dir"] = d.Get("dir")
	if v, ok := d.GetOk("inclusive_file"); ok {
		request["InclusiveFile"] = v
	}
	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	request["LocalBackupDir"] = d.Get("local_backup_dir")
	request["DefenceMode"] = d.Get("defence_mode")
	if v, ok := d.GetOk("mode"); ok {
		request["Mode"] = v
	}
	if v, ok := d.GetOk("inclusive_file_type"); ok {
		request["InclusiveFileType"] = v
	}
	if v, ok := d.GetOk("exclusive_dir"); ok {
		request["ExclusiveDir"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Sas", "2018-12-03", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_threat_detection_web_lock_config", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["Uuid"], response["ConfigId"]))

	return resourceAliCloudThreatDetectionWebLockConfigRead(d, meta)
}

func resourceAliCloudThreatDetectionWebLockConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionWebLockConfig(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_web_lock_config DescribeThreatDetectionWebLockConfig Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("defence_mode", objectRaw["DefenceMode"])
	d.Set("dir", objectRaw["Dir"])
	d.Set("exclusive_dir", objectRaw["ExclusiveDir"])
	d.Set("exclusive_file", objectRaw["ExclusiveFile"])
	d.Set("exclusive_file_type", objectRaw["ExclusiveFileType"])
	d.Set("inclusive_file", objectRaw["InclusiveFile"])
	d.Set("inclusive_file_type", objectRaw["InclusiveFileType"])
	d.Set("local_backup_dir", objectRaw["LocalBackupDir"])
	d.Set("mode", objectRaw["Mode"])
	d.Set("config_id", formatInt(objectRaw["Id"]))
	d.Set("uuid", objectRaw["Uuid"])

	return nil
}

func resourceAliCloudThreatDetectionWebLockConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyWebLockUpdateConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["Id"] = parts[1]
	request["Uuid"] = parts[0]

	if d.HasChange("exclusive_file_type") {
		update = true
		request["ExclusiveFileType"] = d.Get("exclusive_file_type")
	}

	if d.HasChange("exclusive_file") {
		update = true
		request["ExclusiveFile"] = d.Get("exclusive_file")
	}

	if d.HasChange("dir") {
		update = true
	}
	request["Dir"] = d.Get("dir")
	if d.HasChange("inclusive_file") {
		update = true
		request["InclusiveFile"] = d.Get("inclusive_file")
	}

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if d.HasChange("local_backup_dir") {
		update = true
	}
	request["LocalBackupDir"] = d.Get("local_backup_dir")
	if d.HasChange("defence_mode") {
		update = true
	}
	request["DefenceMode"] = d.Get("defence_mode")
	if d.HasChange("mode") {
		update = true
		request["Mode"] = d.Get("mode")
	}

	if d.HasChange("inclusive_file_type") {
		update = true
		request["InclusiveFileType"] = d.Get("inclusive_file_type")
	}

	if d.HasChange("exclusive_dir") {
		update = true
		request["ExclusiveDir"] = d.Get("exclusive_dir")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Sas", "2018-12-03", action, query, request, true)
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

	return resourceAliCloudThreatDetectionWebLockConfigRead(d, meta)
}

func resourceAliCloudThreatDetectionWebLockConfigDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "ModifyWebLockDeleteConfig"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["Id"] = parts[1]
	request["Uuid"] = parts[0]

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Sas", "2018-12-03", action, query, request, true)
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
