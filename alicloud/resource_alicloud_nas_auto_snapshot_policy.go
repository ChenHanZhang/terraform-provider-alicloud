// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/blues/jsonata-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudNasAutoSnapshotPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNasAutoSnapshotPolicyCreate,
		Read:   resourceAliCloudNasAutoSnapshotPolicyRead,
		Update: resourceAliCloudNasAutoSnapshotPolicyUpdate,
		Delete: resourceAliCloudNasAutoSnapshotPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_snapshot_policy_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_system_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"extreme"}, false),
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repeat_weekdays": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"retention_days": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.Any(IntBetween(0, 65536), IntBetween(-1, -1)),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_points": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAliCloudNasAutoSnapshotPolicyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAutoSnapshotPolicy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("auto_snapshot_policy_name"); ok {
		request["AutoSnapshotPolicyName"] = v
	}
	repeatWeekdaysJsonPath, err := jsonpath.Get("$", d.Get("repeat_weekdays"))
	if err == nil {
		request["RepeatWeekdays"] = convertListToCommaSeparate(convertToInterfaceArray(repeatWeekdaysJsonPath))
	}

	request["FileSystemType"] = d.Get("file_system_type")
	timePointsJsonPath, err := jsonpath.Get("$", d.Get("time_points"))
	if err == nil {
		request["TimePoints"] = convertListToCommaSeparate(convertToInterfaceArray(timePointsJsonPath))
	}

	if v, ok := d.GetOkExists("retention_days"); ok && v.(int) > 0 {
		request["RetentionDays"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_nas_auto_snapshot_policy", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["AutoSnapshotPolicyId"]))

	nasServiceV2 := NasServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, nasServiceV2.NasAutoSnapshotPolicyStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudNasAutoSnapshotPolicyRead(d, meta)
}

func resourceAliCloudNasAutoSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nasServiceV2 := NasServiceV2{client}

	objectRaw, err := nasServiceV2.DescribeNasAutoSnapshotPolicy(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_nas_auto_snapshot_policy DescribeNasAutoSnapshotPolicy Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("auto_snapshot_policy_name", objectRaw["AutoSnapshotPolicyName"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("file_system_type", objectRaw["FileSystemType"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("retention_days", objectRaw["RetentionDays"])
	d.Set("status", objectRaw["Status"])

	e := jsonata.MustCompile("$split($.RepeatWeekdays, \",\")")
	evaluation, _ := e.Eval(objectRaw)
	d.Set("repeat_weekdays", evaluation)
	e = jsonata.MustCompile("$split($.TimePoints, \",\")")
	evaluation, _ = e.Eval(objectRaw)
	d.Set("time_points", evaluation)

	return nil
}

func resourceAliCloudNasAutoSnapshotPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyAutoSnapshotPolicy"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AutoSnapshotPolicyId"] = d.Id()

	if d.HasChange("auto_snapshot_policy_name") {
		update = true
		request["AutoSnapshotPolicyName"] = d.Get("auto_snapshot_policy_name")
	}

	if d.HasChange("repeat_weekdays") {
		update = true
	}
	repeatWeekdaysJsonPath, err := jsonpath.Get("$", d.Get("repeat_weekdays"))
	if err == nil {
		request["RepeatWeekdays"] = convertListToCommaSeparate(convertToInterfaceArray(repeatWeekdaysJsonPath))
	}

	if d.HasChange("time_points") {
		update = true
	}
	timePointsJsonPath, err := jsonpath.Get("$", d.Get("time_points"))
	if err == nil {
		request["TimePoints"] = convertListToCommaSeparate(convertToInterfaceArray(timePointsJsonPath))
	}

	if d.HasChange("retention_days") {
		update = true
		request["RetentionDays"] = d.Get("retention_days")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
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
		nasServiceV2 := NasServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, nasServiceV2.NasAutoSnapshotPolicyStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudNasAutoSnapshotPolicyRead(d, meta)
}

func resourceAliCloudNasAutoSnapshotPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAutoSnapshotPolicy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AutoSnapshotPolicyId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"	InvalidLifecyclePolicy.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	nasServiceV2 := NasServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, nasServiceV2.NasAutoSnapshotPolicyStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
