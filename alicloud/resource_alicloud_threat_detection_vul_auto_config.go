// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionVulAutoConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionVulAutoConfigCreate,
		Read:   resourceAliCloudThreatDetectionVulAutoConfigRead,
		Update: resourceAliCloudThreatDetectionVulAutoConfigUpdate,
		Delete: resourceAliCloudThreatDetectionVulAutoConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"all_uuid": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"enable": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"necessity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"need_snapshot": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"period_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshot_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snapshot_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"start_time": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"target_end_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"target_start_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudThreatDetectionVulAutoConfigCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "AddOrUpdateAutoFixConfig"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("config_id"); ok {
		request["ConfigId"] = v
	}

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("period_unit"); ok {
		request["PeriodUnit"] = v
	}
	request["AllUuid"] = d.Get("all_uuid")
	request["NeedSnapshot"] = d.Get("need_snapshot")
	request["StartTime"] = d.Get("start_time")
	if v, ok := d.GetOk("snapshot_name"); ok {
		request["SnapshotName"] = v
	}
	if v, ok := d.GetOkExists("target_start_time"); ok {
		request["TargetStartTime"] = v
	}
	request["Enable"] = d.Get("enable")
	if v, ok := d.GetOk("necessity"); ok {
		request["Necessity"] = v
	}
	if v, ok := d.GetOk("rules"); ok {
		request["Rules"] = v
	}
	if v, ok := d.GetOkExists("target_end_time"); ok {
		request["TargetEndTime"] = v
	}
	request["Type"] = d.Get("type")
	if v, ok := d.GetOkExists("snapshot_time"); ok {
		request["SnapshotTime"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_threat_detection_vul_auto_config", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ConfigId"]))

	return resourceAliCloudThreatDetectionVulAutoConfigRead(d, meta)
}

func resourceAliCloudThreatDetectionVulAutoConfigRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionVulAutoConfig(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_vul_auto_config DescribeThreatDetectionVulAutoConfig Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("all_uuid", objectRaw["AllUuid"])
	d.Set("enable", objectRaw["Enable"])
	d.Set("necessity", objectRaw["Necessity"])
	d.Set("need_snapshot", objectRaw["NeedSnapshot"])
	d.Set("period_unit", objectRaw["PeriodUnit"])
	d.Set("snapshot_name", objectRaw["SnapshotName"])
	d.Set("snapshot_time", objectRaw["SnapshotTime"])
	d.Set("start_time", objectRaw["StartTime"])
	d.Set("target_end_time", objectRaw["TargetEndTime"])
	d.Set("target_start_time", objectRaw["TargetStartTime"])
	d.Set("type", objectRaw["Type"])
	d.Set("config_id", objectRaw["ConfigId"])

	d.Set("config_id", d.Id())

	return nil
}

func resourceAliCloudThreatDetectionVulAutoConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UpdateAutoFixConfigStatus"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ConfigId"] = d.Id()

	if d.HasChange("enable") {
		update = true
	}
	request["Enable"] = d.Get("enable")
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
	update = false
	action = "AddOrUpdateAutoFixConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ConfigId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("period_unit") {
		update = true
		request["PeriodUnit"] = d.Get("period_unit")
	}

	if d.HasChange("all_uuid") {
		update = true
	}
	request["AllUuid"] = d.Get("all_uuid")
	if d.HasChange("need_snapshot") {
		update = true
	}
	request["NeedSnapshot"] = d.Get("need_snapshot")
	if d.HasChange("start_time") {
		update = true
	}
	request["StartTime"] = d.Get("start_time")
	if d.HasChange("snapshot_name") {
		update = true
		request["SnapshotName"] = d.Get("snapshot_name")
	}

	if d.HasChange("target_start_time") {
		update = true
		request["TargetStartTime"] = d.Get("target_start_time")
	}

	if d.HasChange("enable") {
		update = true
	}
	request["Enable"] = d.Get("enable")
	if d.HasChange("necessity") {
		update = true
		request["Necessity"] = d.Get("necessity")
	}

	if v, ok := d.GetOk("rules"); ok {
		request["Rules"] = v
	}
	if d.HasChange("target_end_time") {
		update = true
		request["TargetEndTime"] = d.Get("target_end_time")
	}

	if d.HasChange("type") {
		update = true
	}
	request["Type"] = d.Get("type")
	if d.HasChange("snapshot_time") {
		update = true
		request["SnapshotTime"] = d.Get("snapshot_time")
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

	d.Partial(false)
	return resourceAliCloudThreatDetectionVulAutoConfigRead(d, meta)
}

func resourceAliCloudThreatDetectionVulAutoConfigDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Vul Auto Config. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
