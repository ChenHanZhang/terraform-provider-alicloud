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

func resourceAliCloudEcsDisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEcsDiskCreate,
		Read:   resourceAliCloudEcsDiskRead,
		Update: resourceAliCloudEcsDiskUpdate,
		Delete: resourceAliCloudEcsDiskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(25 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"advanced_features": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_snapshot_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bursting_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"category": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"cloud", "cloud_efficiency", "cloud_ssd", "cloud_essd", "cloud_auto", "cloud_essd_entry", "elastic_ephemeral_disk_standard", "elastic_ephemeral_disk_premium"}, false),
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_auto_snapshot": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"delete_with_instance": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_auto_snapshot": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"encrypt_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"image_id": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"iops": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"iops_write": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"multi_attach": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payment_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Subscription", "PayAsYouGo"}, false),
			},
			"performance_level": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"PL0", "PL1", "PL2", "PL3"}, false),
			},
			"provisioned_iops": {
				Type:     schema.TypeInt,
				Optional: true,
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
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"snapshot_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source_disk_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"storage_set_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"storage_set_partition_number": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"tags": tagsSchema(),
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"offline", "online"}, false),
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudEcsDiskCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDisk"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("performance_level"); ok {
		request["PerformanceLevel"] = v
	}
	if v, ok := d.GetOkExists("bursting_enabled"); ok {
		request["BurstingEnabled"] = v
	}
	if v, ok := d.GetOk("advanced_features"); ok {
		request["AdvancedFeatures"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOkExists("encrypted"); ok {
		request["Encrypted"] = v
	}
	if v, ok := d.GetOk("snapshot_id"); ok {
		request["SnapshotId"] = v
	}
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	request["DiskCategory"] = d.Get("category")
	request["Size"] = d.Get("size")
	if v, ok := d.GetOk("disk_name"); ok {
		request["DiskName"] = v
	}
	request["MultiAttach"] = d.Get("multi_attach")
	if v, ok := d.GetOkExists("storage_set_partition_number"); ok {
		request["StorageSetPartitionNumber"] = v
	}
	if v, ok := d.GetOkExists("provisioned_iops"); ok {
		request["ProvisionedIops"] = v
	}
	if v, ok := d.GetOk("storage_set_id"); ok {
		request["StorageSetId"] = v
	}
	if v, ok := d.GetOk("encrypt_algorithm"); ok {
		request["EncryptAlgorithm"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	if v, ok := d.GetOk("kms_key_id"); ok {
		request["KMSKeyId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"LastTokenProcessing", "UnknownError", "ServiceUnavailable"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_disk", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DiskId"]))

	ecsServiceV2 := EcsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available", "In_use"}, d.Timeout(schema.TimeoutCreate), 10*time.Second, ecsServiceV2.EcsDiskStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	action = "CreateDisks"
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("bursting_enabled"); ok {
		request["BurstingEnabled"] = v
	}
	if v, ok := d.GetOk("performance_level"); ok {
		request["PerformanceLevel"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("storage_cluster_id"); ok {
		request["StorageClusterId"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOkExists("encrypted"); ok {
		request["Encrypted"] = v
	}
	if v, ok := d.GetOk("snapshot_id"); ok {
		request["SnapshotId"] = v
	}
	request["Category"] = d.Get("category")
	request["Size"] = d.Get("size")
	if v, ok := d.GetOk("disk_name"); ok {
		request["DiskName"] = v
	}
	request["MultiAttach"] = d.Get("multi_attach")
	if v, ok := d.GetOkExists("storage_set_partition_number"); ok {
		request["StorageSetPartitionNumber"] = v
	}
	if v, ok := d.GetOkExists("provisioned_iops"); ok {
		request["ProvisionedIops"] = v
	}
	if v, ok := d.GetOk("storage_set_id"); ok {
		request["StorageSetId"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	if v, ok := d.GetOk("kms_key_id"); ok {
		request["KMSKeyId"] = v
	}
	if v, ok := d.GetOk("auto_snapshot_policy_id"); ok {
		request["AutoSnapshotPolicyId"] = v
	}
	wait = incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2016-03-14", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_disk", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.DiskId.DiskIds[0]", response)
	d.SetId(fmt.Sprint(id))

	action = "CloneDisks"
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("bursting_enabled"); ok {
		request["BurstingEnabled"] = v
	}
	if v, ok := d.GetOk("performance_level"); ok {
		request["PerformanceLevel"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	request["SourceDiskId"] = d.Get("source_disk_id")
	if v, ok := d.GetOkExists("encrypted"); ok {
		request["Encrypted"] = v
	}
	tagDataList := make(map[string]interface{})

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagDataList["Value"] = tagsMap
	}

	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagDataList["Key"] = tagsMap
	}

	TagMap := make([]interface{}, 0)
	TagMap = append(TagMap, tagDataList)
	request["Tag"] = TagMap

	request["DiskCategory"] = d.Get("category")
	request["Size"] = d.Get("size")
	if v, ok := d.GetOk("disk_name"); ok {
		request["DiskName"] = v
	}
	request["MultiAttach"] = d.Get("multi_attach")
	if v, ok := d.GetOkExists("provisioned_iops"); ok {
		request["ProvisionedIops"] = v
	}
	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if v, ok := d.GetOk("kms_key_id"); ok {
		request["KmsKeyId"] = v
	}
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_disk", action, AlibabaCloudSdkGoERROR)
	}

	return resourceAliCloudEcsDiskUpdate(d, meta)
}

func resourceAliCloudEcsDiskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ecsServiceV2 := EcsServiceV2{client}

	objectRaw, err := ecsServiceV2.DescribeEcsDisk(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ecs_disk DescribeEcsDisk Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("auto_snapshot_policy_id", objectRaw["AutoSnapshotPolicyId"])
	d.Set("bursting_enabled", objectRaw["BurstingEnabled"])
	d.Set("category", objectRaw["Category"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("delete_auto_snapshot", objectRaw["DeleteAutoSnapshot"])
	d.Set("delete_with_instance", objectRaw["DeleteWithInstance"])
	d.Set("description", objectRaw["Description"])
	d.Set("disk_name", objectRaw["DiskName"])
	d.Set("enable_auto_snapshot", objectRaw["EnableAutoSnapshot"])
	d.Set("encrypted", objectRaw["Encrypted"])
	d.Set("image_id", objectRaw["ImageId"])
	d.Set("instance_id", objectRaw["InstanceId"])
	d.Set("iops", objectRaw["IOPS"])
	d.Set("iops_write", objectRaw["IOPSWrite"])
	d.Set("kms_key_id", objectRaw["KMSKeyId"])
	d.Set("multi_attach", objectRaw["MultiAttach"])
	d.Set("payment_type", convertEcsDiskDisksDiskDiskChargeTypeResponse(objectRaw["DiskChargeType"]))
	d.Set("performance_level", objectRaw["PerformanceLevel"])
	d.Set("provisioned_iops", objectRaw["ProvisionedIops"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("size", objectRaw["Size"])
	d.Set("snapshot_id", objectRaw["SourceSnapshotId"])
	d.Set("source_disk_id", objectRaw["SourceDiskId"])
	d.Set("status", objectRaw["Status"])
	d.Set("storage_cluster_id", objectRaw["StorageClusterId"])
	d.Set("storage_set_id", objectRaw["StorageSetId"])
	d.Set("storage_set_partition_number", objectRaw["StorageSetPartitionNumber"])
	d.Set("zone_id", objectRaw["ZoneId"])

	tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudEcsDiskUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	ecsServiceV2 := EcsServiceV2{client}
	objectRaw, _ := ecsServiceV2.DescribeEcsDisk(d.Id())

	var err error
	action := "ModifyDiskAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DiskId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("delete_auto_snapshot") {
		update = true
		request["DeleteAutoSnapshot"] = d.Get("delete_auto_snapshot")
	}

	if d.HasChange("delete_with_instance") {
		update = true
		request["DeleteWithInstance"] = d.Get("delete_with_instance")
	}

	if !d.IsNewResource() && d.HasChange("bursting_enabled") {
		update = true
		request["BurstingEnabled"] = d.Get("bursting_enabled")
	}

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if !d.IsNewResource() && d.HasChange("disk_name") {
		update = true
		request["DiskName"] = d.Get("disk_name")
	}

	if d.HasChange("enable_auto_snapshot") {
		update = true
		request["EnableAutoSnapshot"] = d.Get("enable_auto_snapshot")
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
	action = "ResizeDisk"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DiskId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("size") {
		update = true
	}
	request["NewSize"] = d.Get("size")
	if v, ok := d.GetOk("type"); ok {
		request["Type"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"InternalError"}) || NeedRetry(err) {
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
	action = "JoinResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	request["ResourceType"] = "disk"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"InternalError"}) || NeedRetry(err) {
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
	action = "ModifyDiskSpec"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DiskId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("performance_level") {
		update = true
		request["PerformanceLevel"] = d.Get("performance_level")
	}

	if !d.IsNewResource() && d.HasChange("category") {
		update = true
	}
	request["DiskCategory"] = d.Get("category")
	if d.HasChange("iops_write") {
		update = true
		request["PerformanceControlOptions.Throughput"] = d.Get("iops_write")
	}

	if !d.IsNewResource() && d.HasChange("provisioned_iops") {
		update = true
		request["ProvisionedIops"] = d.Get("provisioned_iops")
	}

	if d.HasChange("iops") {
		update = true
		request["PerformanceControlOptions.IOPS"] = d.Get("iops")
	}

	if !d.IsNewResource() && d.HasChange("zone_id") {
		update = true
		request["DestinationZoneId"] = d.Get("zone_id")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"Throttling.ConcurrentLimitExceeded", "ServiceUnavailable"}) || NeedRetry(err) {
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
		stateConf := BuildStateConf([]string{}, []string{"Available", "In_use"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, ecsServiceV2.EcsDiskStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	objectRaw, _ = ecsServiceV2.DescribeEcsDisk(d.Id())
	enableModifyDiskChargeType1 := false
	checkValue00 := objectRaw["Status"]
	if checkValue00 == "In_use" {
		enableModifyDiskChargeType1 = true
	}
	action = "ModifyDiskChargeType"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DiskIds"] = expandSingletonToList(d.Id())
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("instance_id") {
		update = true
	}
	request["InstanceId"] = d.Get("instance_id")
	if d.HasChange("payment_type") {
		update = true
		request["DiskChargeType"] = convertEcsDiskDiskChargeTypeRequest(d.Get("payment_type").(string))
	}

	request["AutoPay"] = "true"
	if update && enableModifyDiskChargeType1 {
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
	action = "ReplaceSystemDisk"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DiskId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("instance_id") {
		update = true
	}
	request["InstanceId"] = d.Get("instance_id")
	if !d.IsNewResource() && d.HasChange("size") {
		update = true
	}
	request["SystemDisk\\.Size"] = d.Get("size")
	if !d.IsNewResource() && d.HasChange("encrypted") {
		update = true
		request["Encrypted"] = d.Get("encrypted")
	}

	if d.HasChange("image_id") {
		update = true
		request["ImageId"] = d.Get("image_id")
	}

	if !d.IsNewResource() && d.HasChange("kms_key_id") {
		update = true
		request["KMSKeyId"] = d.Get("kms_key_id")
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
	action = "ModifyDiskSpec"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DiskId"] = d.Id()
	request["SourceRegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("performance_level") {
		update = true
		request["PerformanceLevel"] = d.Get("performance_level")
	}

	if !d.IsNewResource() && d.HasChange("category") {
		update = true
	}
	request["DiskCategory"] = d.Get("category")
	request["AutoPay"] = "true"
	if !d.IsNewResource() && d.HasChange("provisioned_iops") {
		update = true
		request["ProvisionedIops"] = d.Get("provisioned_iops")
	}

	if !d.IsNewResource() && d.HasChange("zone_id") {
		update = true
		request["DestinationZoneId"] = d.Get("zone_id")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2016-03-14", action, query, request, true)
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

	if d.HasChange("tags") {
		ecsServiceV2 := EcsServiceV2{client}
		if err := ecsServiceV2.SetResourceTags(d, "disk"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudEcsDiskRead(d, meta)
}

func resourceAliCloudEcsDiskDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	enableDelete := false
	if v, ok := d.GetOkExists("status"); ok {
		if InArray(fmt.Sprint(v), []string{"Available"}) {
			enableDelete = true
		}
	}
	if enableDelete {
		action := "DeleteDisk"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["DiskId"] = d.Id()

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectDiskStatus.Initializing"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)

		if err != nil {
			if IsExpectedErrors(err, []string{"InvalidDiskId.NotFound"}) || NotFoundError(err) {
				return nil
			}
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}

	}
	return nil
}

func convertEcsDiskDisksDiskDiskChargeTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "PrePaid":
		return "Subscription"
	case "PostPaid":
		return "PayAsYouGo"
	}
	return source
}
func convertEcsDiskDiskChargeTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Subscription":
		return "PrePaid"
	case "PayAsYouGo":
		return "PostPaid"
	}
	return source
}
