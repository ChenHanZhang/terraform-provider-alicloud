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

func resourceAliCloudDbfsDbfsInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudDbfsDbfsInstanceCreate,
		Read:   resourceAliCloudDbfsDbfsInstanceRead,
		Update: resourceAliCloudDbfsDbfsInstanceUpdate,
		Delete: resourceAliCloudDbfsDbfsInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"advanced_features": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"standard", "enterprise"}, false),
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_snapshot": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_raid": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"fs_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"dbfs.small", "dbfs.medium", "dbfs.large "}, false),
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"performance_level": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"PL0", "PL1", "PL2", "PL3"}, false),
			},
			"raid_stripe_unit_number": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: IntBetween(2, 8),
			},
			"size": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: IntBetween(20, 262144),
			},
			"snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"used_scene": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"MySQL 5.7", "PostgreSQL ", "MongoDB", "DataCube"}, false),
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudDbfsDbfsInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDbfs"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewDbfsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	request["FsName"] = d.Get("fs_name")
	request["Category"] = d.Get("category")
	request["ZoneId"] = d.Get("zone_id")
	request["SizeG"] = d.Get("size")
	if v, ok := d.GetOk("snapshot_id"); ok {
		request["SnapshotId"] = v
	}
	if v, ok := d.GetOkExists("delete_snapshot"); ok {
		request["DeleteSnapshot"] = v
	}
	if v, ok := d.GetOk("performance_level"); ok {
		request["PerformanceLevel"] = v
	}
	if v, ok := d.GetOkExists("enable_raid"); ok {
		request["EnableRaid"] = v
	}
	if v, ok := d.GetOk("raid_stripe_unit_number"); ok {
		request["RaidStripeUnitNumber"] = v
	}
	if v, ok := d.GetOk("kms_key_id"); ok {
		request["KMSKeyId"] = v
	}
	if v, ok := d.GetOkExists("encryption"); ok {
		request["Encryption"] = v
	}
	if v, ok := d.GetOk("used_scene"); ok {
		request["UsedScene"] = v
	}
	if v, ok := d.GetOk("instance_type"); ok {
		request["InstanceType"] = v
	}
	if v, ok := d.GetOk("advanced_features"); ok {
		request["AdvancedFeatures"] = v
	}
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-18"), StringPointer("AK"), query, request, &runtime)
		request["ClientToken"] = buildClientToken(action)

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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_dbfs_dbfs_instance", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["FsId"]))

	dbfsServiceV2 := DbfsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"unattached"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, dbfsServiceV2.DbfsDbfsInstanceStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudDbfsDbfsInstanceRead(d, meta)
}

func resourceAliCloudDbfsDbfsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dbfsServiceV2 := DbfsServiceV2{client}

	objectRaw, err := dbfsServiceV2.DescribeDbfsDbfsInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_dbfs_dbfs_instance DescribeDbfsDbfsInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("advanced_features", objectRaw["AdvancedFeatures"])
	d.Set("category", objectRaw["Category"])
	d.Set("create_time", objectRaw["CreatedTime"])
	d.Set("enable_raid", objectRaw["EnableRaid"])
	d.Set("encryption", objectRaw["Encryption"])
	d.Set("fs_name", objectRaw["FsName"])
	d.Set("instance_type", objectRaw["InstanceType"])
	d.Set("kms_key_id", objectRaw["KMSKeyId"])
	d.Set("performance_level", objectRaw["PerformanceLevel"])
	d.Set("raid_stripe_unit_number", objectRaw["RaidStrip"])
	d.Set("size", objectRaw["SizeG"])
	d.Set("snapshot_id", objectRaw["SnapshotId"])
	d.Set("status", objectRaw["Status"])
	d.Set("used_scene", objectRaw["UsedScene"])
	d.Set("zone_id", objectRaw["ZoneId"])

	return nil
}

func resourceAliCloudDbfsDbfsInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)
	action := "RenameDbfs"
	conn, err := client.NewDbfsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	query["FsId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("fs_name") {
		update = true
	}
	request["FsName"] = d.Get("fs_name")
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-18"), StringPointer("AK"), query, request, &runtime)

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
		d.SetPartial("fs_name")
	}
	update = false
	action = "ResizeDbfs"
	conn, err = client.NewDbfsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	query["FsId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("size") {
		update = true
	}
	request["NewSizeG"] = d.Get("size")
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-18"), StringPointer("AK"), query, request, &runtime)

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
		dbfsServiceV2 := DbfsServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("size"))}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, dbfsServiceV2.DbfsDbfsInstanceStateRefreshFunc(d.Id(), "SizeG", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
		d.SetPartial("size")
	}
	update = false
	action = "UpdateDbfs"
	conn, err = client.NewDbfsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	query["FsId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("used_scene") {
		update = true
		request["UsedScene"] = d.Get("used_scene")
	}

	if d.HasChange("instance_type") {
		update = true
		request["InstanceType"] = d.Get("instance_type")
	}

	if d.HasChange("advanced_features") {
		update = true
		request["AdvancedFeatures"] = d.Get("advanced_features")
	}

	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-18"), StringPointer("AK"), query, request, &runtime)

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
		d.SetPartial("used_scene")
		d.SetPartial("instance_type")
		d.SetPartial("advanced_features")
	}
	update = false
	action = "ModifyPerformanceLevel"
	conn, err = client.NewDbfsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	query["FsId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("performance_level") {
		update = true
		request["PerformanceLevel"] = d.Get("performance_level")
	}

	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-18"), StringPointer("AK"), query, request, &runtime)

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
		d.SetPartial("performance_level")
	}

	d.Partial(false)
	return resourceAliCloudDbfsDbfsInstanceRead(d, meta)
}

func resourceAliCloudDbfsDbfsInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteDbfs"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewDbfsClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query["FsId"] = d.Id()
	request["RegionId"] = client.RegionId

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-04-18"), StringPointer("AK"), query, request, &runtime)

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

	dbfsServiceV2 := DbfsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, dbfsServiceV2.DbfsDbfsInstanceStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}

func convertDbfsDBFSInfoPayTypeResponse(source interface{}) interface{} {
	switch source {
	case "postpaid":
		return "PayAsYouGo"
	}
	return source
}
