// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionGlobalConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionGlobalConfigCreate,
		Read:   resourceAliCloudThreatDetectionGlobalConfigRead,
		Update: resourceAliCloudThreatDetectionGlobalConfigUpdate,
		Delete: resourceAliCloudThreatDetectionGlobalConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"global_config_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_config_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role_for": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudThreatDetectionGlobalConfigCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudThreatDetectionGlobalConfigUpdate(d, meta)
}

func resourceAliCloudThreatDetectionGlobalConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionGlobalConfig(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_global_config DescribeThreatDetectionGlobalConfig Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("global_config_name", objectRaw["GlobalConfigName"])
	d.Set("global_config_value", objectRaw["GlobalConfigValue"])

	return nil
}

func resourceAliCloudThreatDetectionGlobalConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateGlobalConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	if d.HasChange("global_config_name") {
		update = true
	}
	request["GlobalConfigName"] = d.Get("global_config_name")
	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if d.HasChange("global_config_value") {
		update = true
		request["GlobalConfigValue"] = d.Get("global_config_value")
	}

	if v, ok := d.GetOkExists("role_for"); ok {
		request["RoleFor"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cloud-siem", "2024-12-12", action, query, request, true)
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

	return resourceAliCloudThreatDetectionGlobalConfigRead(d, meta)
}

func resourceAliCloudThreatDetectionGlobalConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Global Config. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
