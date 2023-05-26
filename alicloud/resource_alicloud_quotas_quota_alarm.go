// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/blues/jsonata-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudQuotasQuotaAlarm() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudQuotasQuotaAlarmCreate,
		Read:   resourceAlicloudQuotasQuotaAlarmRead,
		Update: resourceAlicloudQuotasQuotaAlarmUpdate,
		Delete: resourceAlicloudQuotasQuotaAlarmDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alarm_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"quota_action_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"quota_alarm_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"quota_dimensions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"threshold": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"threshold_percent": {
				Type:         schema.TypeFloat,
				ExactlyOneOf: []string{"threshold_percent", "threshold"},
				Optional:     true,
			},
			"threshold_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_hook": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudQuotasQuotaAlarmCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateQuotaAlarm"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewQuotasClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	if v, ok := d.GetOk("product_code"); ok {
		request["ProductCode"] = v
	}
	if v, ok := d.GetOk("quota_action_code"); ok {
		request["QuotaActionCode"] = v
	}
	if v, ok := d.GetOk("threshold"); ok {
		request["Threshold"] = v
	}
	if v, ok := d.GetOk("threshold_percent"); ok {
		request["ThresholdPercent"] = v
	}
	if v, ok := d.GetOk("web_hook"); ok {
		request["WebHook"] = v
	}
	if v, ok := d.GetOk("quota_dimensions"); ok {
		quotaDimensionsMaps := make([]map[string]interface{}, 0)
		for _, dataLoop := range v.([]interface{}) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Key"] = dataLoopTmp["key"]
			dataLoopMap["Value"] = dataLoopTmp["value"]
			quotaDimensionsMaps = append(quotaDimensionsMaps, dataLoopMap)
		}
		request["QuotaDimensions"] = quotaDimensionsMaps
	}

	if v, ok := d.GetOk("quota_alarm_name"); ok {
		request["AlarmName"] = v
	}
	if v, ok := d.GetOk("threshold_type"); ok {
		request["ThresholdType"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-05-10"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_quotas_quota_alarm", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["AlarmId"]))

	return resourceAlicloudQuotasQuotaAlarmUpdate(d, meta)
}

func resourceAlicloudQuotasQuotaAlarmRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	quotasServiceV2 := QuotasServiceV2{client}

	objectRaw, err := quotasServiceV2.DescribeQuotasQuotaAlarm(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_quotas_quota_alarm DescribeQuotasQuotaAlarm Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("product_code", objectRaw["ProductCode"])
	d.Set("quota_action_code", objectRaw["QuotaActionCode"])
	d.Set("quota_alarm_name", objectRaw["AlarmName"])
	d.Set("threshold", objectRaw["Threshold"])
	d.Set("threshold_percent", objectRaw["ThresholdPercent"])
	d.Set("threshold_type", objectRaw["ThresholdType"])
	d.Set("web_hook", objectRaw["Webhook"])
	d.Set("alarm_id", objectRaw["AlarmId"])

	e := jsonata.MustCompile("$each($.QuotaDimension, function($v, $k) {{\"value\":$v, \"key\": $k}})[]")
	evaluation, _ := e.Eval(objectRaw)
	d.Set("quota_dimensions", evaluation)

	return nil
}

func resourceAlicloudQuotasQuotaAlarmUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	action := "UpdateQuotaAlarm"
	conn, err := client.NewQuotasClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["AlarmId"] = d.Id()
	if !d.IsNewResource() && d.HasChange("threshold") {
		update = true
		if v, ok := d.GetOk("threshold"); ok {
			request["Threshold"] = v
		}
	}

	if !d.IsNewResource() && d.HasChange("threshold_percent") {
		update = true
		if v, ok := d.GetOk("threshold_percent"); ok {
			request["ThresholdPercent"] = v
		}
	}

	if !d.IsNewResource() && d.HasChange("web_hook") {
		update = true
		if v, ok := d.GetOk("web_hook"); ok {
			request["WebHook"] = v
		}
	}

	if !d.IsNewResource() && d.HasChange("quota_alarm_name") {
		update = true
	}
	request["AlarmName"] = d.Get("quota_alarm_name")
	if !d.IsNewResource() && d.HasChange("threshold_type") {
		update = true
		if v, ok := d.GetOk("threshold_type"); ok {
			request["ThresholdType"] = v
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-05-10"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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
		d.SetPartial("threshold")
		d.SetPartial("threshold_percent")
		d.SetPartial("web_hook")
		d.SetPartial("quota_alarm_name")
		d.SetPartial("threshold_type")
	}

	return resourceAlicloudQuotasQuotaAlarmRead(d, meta)
}

func resourceAlicloudQuotasQuotaAlarmDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteQuotaAlarm"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewQuotasClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["AlarmId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-05-10"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})

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

	return nil
}
