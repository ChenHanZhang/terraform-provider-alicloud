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

func resourceAliCloudAdbDbClusterLakeVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbDbClusterLakeVersionCreate,
		Read:   resourceAliCloudAdbDbClusterLakeVersionRead,
		Update: resourceAliCloudAdbDbClusterLakeVersionUpdate,
		Delete: resourceAliCloudAdbDbClusterLakeVersionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(9 * time.Minute),
			Update: schema.DefaultTimeout(47 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"audit_log_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_renewal_period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"auto_renewal_period_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_renewal_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"capacity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"commodity_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_resource": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_cluster_description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_cluster_network_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_cluster_version": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"5.0"}, false),
			},
			"db_cluster_ip_array_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_cluster_ip_array_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"default"}, false),
			},
			"db_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"enable_compaction_service": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_default_resource_group": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_essd_cache": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"enable_lake_cache": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"engine": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"XIHE", "SPARK"}, false),
			},
			"engine_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"essd_cache_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"expire_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expired": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kms_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"lock_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lock_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintain_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"modify_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"PayAsYouGo", "Subscription"}, false),
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"product_form": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reserved_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reserved_node_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restore_to_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"restore_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"secondary_zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"security_ips": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_db_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_resource": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_mode": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntInSlice([]int{0, 1}),
			},
			"tags": tagsSchema(),
			"used_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudAdbDbClusterLakeVersionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("backup_set_id"); ok {
		request["BackupSetId"] = v
	}
	request["PayType"] = convertAdbDbClusterLakeVersionPayTypeRequest(d.Get("payment_type").(string))
	request["DBClusterNetworkType"] = d.Get("db_cluster_network_type")
	if v, ok := d.GetOk("restore_to_time"); ok {
		request["RestoreToTime"] = v
	}
	request["VSwitchId"] = d.Get("vswitch_id")
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("source_db_cluster_id"); ok {
		request["SourceDbClusterId"] = v
	}
	if v, ok := d.GetOk("secondary_vswitch_id"); ok {
		request["SecondaryVSwitchId"] = v
	}
	if v, ok := d.GetOkExists("disk_encryption"); ok {
		request["DiskEncryption"] = v
	}
	if v, ok := d.GetOk("kms_id"); ok {
		request["KmsId"] = v
	}
	if v, ok := d.GetOk("restore_type"); ok {
		request["RestoreType"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("storage_resource"); ok {
		request["StorageResource"] = v
	}
	if v, ok := d.GetOk("secondary_zone_id"); ok {
		request["SecondaryZoneId"] = v
	}
	if v, ok := d.GetOk("product_version"); ok {
		request["ProductVersion"] = v
	}
	if v, ok := d.GetOkExists("enable_default_resource_group"); ok {
		request["EnableDefaultResourcePool"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	request["DBClusterVersion"] = d.Get("db_cluster_version")
	if v, ok := d.GetOkExists("enable_ssl"); ok {
		request["EnableSSL"] = v
	}
	if v, ok := d.GetOk("db_cluster_description"); ok {
		request["DBClusterDescription"] = v
	}
	if v, ok := d.GetOkExists("reserved_node_count"); ok {
		request["ReservedNodeCount"] = v
	}
	if v, ok := d.GetOk("reserved_node_size"); ok {
		request["ReservedNodeSize"] = v
	}
	request["VPCId"] = d.Get("vpc_id")
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("product_form"); ok {
		request["ProductForm"] = v
	}
	request["ZoneId"] = d.Get("zone_id")
	if v, ok := d.GetOk("compute_resource"); ok {
		request["ComputeResource"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_db_cluster_lake_version", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBClusterId"]))

	adbServiceV2 := AdbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutCreate), 9*time.Minute, adbServiceV2.AdbDbClusterLakeVersionStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudAdbDbClusterLakeVersionUpdate(d, meta)
}

func resourceAliCloudAdbDbClusterLakeVersionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbDbClusterLakeVersion(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_db_cluster_lake_version DescribeAdbDbClusterLakeVersion Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("commodity_code", objectRaw["CommodityCode"])
	d.Set("compute_resource", objectRaw["ComputeResource"])
	d.Set("connection_string", objectRaw["ConnectionString"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("db_cluster_description", objectRaw["DBClusterDescription"])
	d.Set("db_cluster_network_type", objectRaw["DBClusterNetworkType"])
	d.Set("db_cluster_version", objectRaw["DBVersion"])
	d.Set("disk_encryption", objectRaw["DiskEncryption"])
	d.Set("engine", objectRaw["Engine"])
	d.Set("engine_version", objectRaw["EngineVersion"])
	d.Set("expire_time", objectRaw["ExpireTime"])
	d.Set("expired", objectRaw["Expired"])
	d.Set("kms_id", objectRaw["KmsId"])
	d.Set("lock_mode", objectRaw["LockMode"])
	d.Set("lock_reason", objectRaw["LockReason"])
	d.Set("maintain_time", objectRaw["MaintainTime"])
	d.Set("payment_type", convertAdbDbClusterLakeVersionItemsDBClusterPayTypeResponse(objectRaw["PayType"]))
	d.Set("port", objectRaw["Port"])
	d.Set("product_form", objectRaw["ProductForm"])
	d.Set("product_version", objectRaw["ProductVersion"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("reserved_node_count", objectRaw["ReservedNodeCount"])
	d.Set("reserved_node_size", objectRaw["ReservedNodeSize"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("secondary_vswitch_id", objectRaw["SecondaryVSwitchId"])
	d.Set("secondary_zone_id", objectRaw["SecondaryZoneId"])
	d.Set("status", objectRaw["DBClusterStatus"])
	d.Set("storage_resource", objectRaw["StorageResource"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("zone_id", objectRaw["ZoneId"])

	d.Set("tags", tagsToMap(tagsMaps))

	objectRaw, err = adbServiceV2.DescribeDbClusterLakeVersionDescribeAuditLogConfig(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("audit_log_status", objectRaw["AuditLogStatus"])

	objectRaw, err = adbServiceV2.DescribeDbClusterLakeVersionDescribeEssdCacheConfig(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("enable_essd_cache", objectRaw["EnableEssdCache"])
	d.Set("essd_cache_size", objectRaw["EssdCacheSize"])

	objectRaw, err = adbServiceV2.DescribeDbClusterLakeVersionDescribeLakeCacheSize(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("capacity", objectRaw["Capacity"])
	d.Set("enable_lake_cache", objectRaw["EnableLakeCache"])

	objectRaw, err = adbServiceV2.DescribeDbClusterLakeVersionDescribeDBClusterSSL(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("enable_ssl", objectRaw["SSLEnabled"])

	objectRaw, err = adbServiceV2.DescribeDbClusterLakeVersionDescribeClusterAccessWhiteList(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("security_ips", objectRaw["SecurityIPList"])

	objectRaw, err = adbServiceV2.DescribeDbClusterLakeVersionDescribeCompactionServiceSwitch(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("enable_compaction_service", objectRaw["EnableCompactionService"])

	return nil
}

func resourceAliCloudAdbDbClusterLakeVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyDBCluster"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOkExists("enable_default_resource_group"); ok {
		request["EnableDefaultResourcePool"] = v
	}
	if !d.IsNewResource() && d.HasChange("reserved_node_count") {
		update = true
		request["ReservedNodeCount"] = d.Get("reserved_node_count")
	}

	if !d.IsNewResource() && d.HasChange("reserved_node_size") {
		update = true
		request["ReservedNodeSize"] = d.Get("reserved_node_size")
	}

	if !d.IsNewResource() && d.HasChange("product_form") {
		update = true
		request["ProductForm"] = d.Get("product_form")
	}

	if !d.IsNewResource() && d.HasChange("compute_resource") {
		update = true
		request["ComputeResource"] = d.Get("compute_resource")
	}

	if !d.IsNewResource() && d.HasChange("storage_resource") {
		update = true
		request["StorageResource"] = d.Get("storage_resource")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		adbServiceV2 := AdbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, adbServiceV2.AdbDbClusterLakeVersionStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ModifyDBClusterDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("db_cluster_description") {
		update = true
	}
	request["DBClusterDescription"] = d.Get("db_cluster_description")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyClusterAccessWhiteList"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if v, ok := d.GetOk("db_cluster_ip_array_attribute"); ok {
		request["DBClusterIPArrayAttribute"] = v
	}
	if v, ok := d.GetOk("modify_mode"); ok {
		request["ModifyMode"] = v
	}
	if v, ok := d.GetOk("db_cluster_ip_array_name"); ok {
		request["DBClusterIPArrayName"] = v
	}
	if d.HasChange("security_ips") {
		update = true
	}
	request["SecurityIps"] = d.Get("security_ips")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyDBClusterResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
	}
	request["NewResourceGroupId"] = d.Get("resource_group_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyDBClusterMaintainTime"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("maintain_time") {
		update = true
	}
	request["MaintainTime"] = d.Get("maintain_time")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyAuditLogConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("audit_log_status") {
		update = true
	}
	request["AuditLogStatus"] = d.Get("audit_log_status")
	if v, ok := d.GetOk("engine_type"); ok {
		request["EngineType"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyEssdCacheConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("essd_cache_size") {
		update = true
	}
	request["EssdCacheSize"] = d.Get("essd_cache_size")
	if d.HasChange("enable_essd_cache") {
		update = true
	}
	request["EnableEssdCache"] = d.Get("enable_essd_cache")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyLakeCacheSize"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("capacity") {
		update = true
		request["Capacity"] = d.Get("capacity")
	}

	if d.HasChange("enable_lake_cache") {
		update = true
	}
	request["EnableLakeCache"] = d.Get("enable_lake_cache")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyCompactionServiceSwitch"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("enable_compaction_service") {
		update = true
	}
	request["EnableCompactionService"] = d.Get("enable_compaction_service")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyDBClusterSSL"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("connection_string") {
		update = true
		request["ConnectionString"] = d.Get("connection_string")
	}

	if !d.IsNewResource() && d.HasChange("enable_ssl") {
		update = true
	}
	request["EnableSSL"] = d.Get("enable_ssl")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		adbServiceV2 := AdbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"true"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, adbServiceV2.AdbDbClusterLakeVersionStateRefreshFuncWithApi(d.Id(), "SSLEnabled", []string{}, adbServiceV2.DescribeDbClusterLakeVersionDescribeDBClusterSSL))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "AttachUserENI"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		adbServiceV2 := AdbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, adbServiceV2.AdbDbClusterLakeVersionStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "DetachUserENI"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		adbServiceV2 := AdbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, adbServiceV2.AdbDbClusterLakeVersionStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "DisableAdviceService"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "EnableAdviceService"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "UpgradeKernelVersion"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if v, ok := d.GetOk("db_version"); ok {
		request["DBVersion"] = v
	}
	if v, ok := d.GetOkExists("switch_mode"); ok {
		request["SwitchMode"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
	action = "ModifyAutoRenewalAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("auto_renewal_status"); ok {
		request["AutoRenewalStatus"] = v
	}
	if v, ok := d.GetOkExists("auto_renewal_period"); ok {
		request["AutoRenewalPeriod"] = v
	}
	if v, ok := d.GetOk("auto_renewal_period_unit"); ok {
		request["AutoRenewalPeriodUnit"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		adbServiceV2 := AdbServiceV2{client}
		if err := adbServiceV2.SetResourceTags(d, "dbclusterlakeversion"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudAdbDbClusterLakeVersionRead(d, meta)
}

func resourceAliCloudAdbDbClusterLakeVersionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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

	adbServiceV2 := AdbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 30*time.Second, adbServiceV2.AdbDbClusterLakeVersionStateRefreshFunc(d.Id(), "$.DBClusterId", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}

func convertAdbDbClusterLakeVersionItemsDBClusterPayTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Prepaid":
		return "Subscription"
	case "Postpaid":
		return "PayAsYouGo"
	}
	return source
}
func convertAdbDbClusterLakeVersionPayTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "PayAsYouGo":
		return "Postpaid"
	case "Subscription":
		return "Prepaid"
	}
	return source
}
