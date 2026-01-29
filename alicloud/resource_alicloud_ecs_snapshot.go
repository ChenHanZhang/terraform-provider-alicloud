// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudEcsSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEcsSnapshotCreate,
		Read:   resourceAliCloudEcsSnapshotRead,
		Update: resourceAliCloudEcsSnapshotUpdate,
		Delete: resourceAliCloudEcsSnapshotDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"category": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"standard", "flash"}, false),
			},
			"cool_off_period": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntInSlice([]int{0, 1}),
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"disk_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"force": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"instant_access": {
				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "Field 'instant_access' has been deprecated from provider version 1.270.0. Field `instant_access` has been deprecated from provider version 1.231.0.",
			},
			"instant_access_retention_days": {
				Type:       schema.TypeInt,
				Optional:   true,
				ForceNew:   true,
				Deprecated: "Field 'instant_access_retention_days' has been deprecated from provider version 1.270.0. Field `instant_access_retention_days` has been deprecated from provider version 1.231.0.",
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"lock_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntInSlice([]int{0, 1}),
			},
			"lock_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"compliance"}, false),
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retention_days": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snapshot_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_region_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"cn-hangzhou"}, false),
			},
			"source_snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudEcsSnapshotCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateSnapshot"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	request["DiskId"] = d.Get("disk_id")
	if v, ok := d.GetOkExists("instant_access"); ok {
		request["InstantAccess"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("category"); ok {
		request["Category"] = v
	}
	request["Description"] = d.Get("description")
	request["SnapshotName"] = d.Get("snapshot_name")
	if v, ok := d.GetOkExists("instant_access_retention_days"); ok {
		request["InstantAccessRetentionDays"] = v
	}
	if v, ok := d.GetOkExists("retention_days"); ok {
		request["RetentionDays"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_snapshot", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["SnapshotId"]))

	ecsServiceV2 := EcsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"accomplished"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, ecsServiceV2.EcsSnapshotStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	action = "CopySnapshot"
	request = make(map[string]interface{})
	request["DestinationRegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["DestinationSnapshotDescription"] = d.Get("description")
	if v, ok := d.GetOkExists("encrypted"); ok {
		request["Encrypted"] = v
	}
	request["DestinationSnapshotName"] = d.Get("snapshot_name")
	request["SnapshotId"] = d.Get("source_snapshot_id")
	if v, ok := d.GetOk("kms_key_id"); ok {
		request["KMSKeyId"] = v
	}
	if v, ok := d.GetOkExists("retention_days"); ok {
		request["RetentionDays"] = v
	}
	request["RegionId"] = d.Get("source_region_id")
	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_snapshot", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["SnapshotId"]))

	return resourceAliCloudEcsSnapshotUpdate(d, meta)
}

func resourceAliCloudEcsSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ecsServiceV2 := EcsServiceV2{client}

	objectRaw, err := ecsServiceV2.DescribeEcsSnapshot(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ecs_snapshot DescribeEcsSnapshot Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("category", objectRaw["Category"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("description", objectRaw["Description"])
	d.Set("disk_id", objectRaw["SourceDiskId"])
	d.Set("encrypted", objectRaw["Encrypted"])
	d.Set("instant_access", objectRaw["InstantAccess"])
	d.Set("instant_access_retention_days", objectRaw["InstantAccessRetentionDays"])
	d.Set("kms_key_id", objectRaw["KMSKeyId"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("retention_days", objectRaw["RetentionDays"])
	d.Set("snapshot_name", objectRaw["SnapshotName"])
	d.Set("status", objectRaw["Status"])
	d.Set("source_region_id", objectRaw["SourceRegionId"])

	tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	objectRaw, err = ecsServiceV2.DescribeSnapshotDescribeLockedSnapshots(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("cool_off_period", objectRaw["CoolOffPeriod"])
	d.Set("lock_duration", objectRaw["LockDuration"])

	return nil
}

func resourceAliCloudEcsSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	ecsServiceV2 := EcsServiceV2{client}
	objectRaw, _ := ecsServiceV2.DescribeEcsSnapshot(d.Id())

	if d.HasChange("lock_status") {
		var err error
		target := d.Get("lock_status").(string)
		if target == "unlocked" {
			action := "UnlockSnapshot"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["SnapshotId"] = d.Id()
			request["RegionId"] = client.RegionId
			request["ClientToken"] = buildClientToken(action)
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
			ecsServiceV2 := EcsServiceV2{client}
			stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ecsServiceV2.EcsSnapshotStateRefreshFuncWithApi(d.Id(), "LockDuration", []string{}, ecsServiceV2.DescribeSnapshotDescribeLockedSnapshots))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}

		}
	}

	var err error
	action := "ModifySnapshotAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["SnapshotId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("instant_access") {
		update = true
		request["DisableInstantAccess"] = d.Get("instant_access")
	}

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
	}
	request["Description"] = d.Get("description")
	if !d.IsNewResource() && d.HasChange("snapshot_name") {
		update = true
	}
	request["SnapshotName"] = d.Get("snapshot_name")
	if !d.IsNewResource() && d.HasChange("retention_days") {
		update = true
		request["RetentionDays"] = d.Get("retention_days")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
		ecsServiceV2 := EcsServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"accomplished"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, ecsServiceV2.EcsSnapshotStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "JoinResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	request["ResourceType"] = "snapshot"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
	action = "ModifySnapshotCategory"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["SnapshotId"] = d.Id()
	request["SourceRegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("category") {
		update = true
		request["Category"] = d.Get("category")
	}

	if !d.IsNewResource() && d.HasChange("retention_days") {
		update = true
		request["RetentionDays"] = d.Get("retention_days")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
	action = "LockSnapshot"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["SnapshotId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("cool_off_period") {
		update = true
	}
	request["CoolOffPeriod"] = d.Get("cool_off_period")
	if d.HasChange("lock_duration") {
		update = true
	}
	request["LockDuration"] = d.Get("lock_duration")
	request["LockMode"] = d.Get("lock_mode")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
		ecsServiceV2 := EcsServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"#CHECKSET"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, ecsServiceV2.EcsSnapshotStateRefreshFuncWithApi(d.Id(), "#LockDuration", []string{}, ecsServiceV2.DescribeSnapshotDescribeLockedSnapshots))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	if d.HasChange("tags") {
		ecsServiceV2 := EcsServiceV2{client}
		if err := ecsServiceV2.SetResourceTags(d, "snapshot"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudEcsSnapshotRead(d, meta)
}

func resourceAliCloudEcsSnapshotDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteSnapshot"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["SnapshotId"] = d.Id()

	if v, ok := d.GetOkExists("force"); ok {
		request["Force"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidResource.NotFound", "InvalidParameter"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}

func convertEcsSnapshotLockedSnapshotsInfoLockStatusResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "unlocked":
		return ""
	}
	return source
}
