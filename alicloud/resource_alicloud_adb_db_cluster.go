// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tidwall/sjson"
)

func resourceAliCloudAdbDbCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbDbClusterCreate,
		Read:   resourceAliCloudAdbDbClusterRead,
		Update: resourceAliCloudAdbDbClusterUpdate,
		Delete: resourceAliCloudAdbDbClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(65 * time.Minute),
			Update: schema.DefaultTimeout(134 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_renew_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"backup_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_resource": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_cluster_category": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_cluster_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_cluster_network_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_cluster_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_node_class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_node_storage": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_cluster_ip_array_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_cluster_ip_array_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disk_encryption": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"disk_performance_level": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"PL1", "PL2", "PL3"}, false),
			},
			"elastic_io_resource": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"elastic_io_resource_size": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"8Core64GB", "12Core96GB", "16Core128GB"}, false),
			},
			"enable_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"executor_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kernel_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kms_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"maintain_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"Reserver", "Flexible", "flexible", "reserver"}, false),
			},
			"modify_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"0", "1", "2"}, false),
			},
			"modify_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Upgrade", "Downgrade"}, false),
			},
			"pay_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"payment_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"PayAsYouGo", "Subscription"}, false),
			},
			"period": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Year", "Month"}, false),
			},
			"period_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Year", "Month"}, false),
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restore_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"restore_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_ips": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_db_instance_name": {
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
			},
			"storage_type": {
				Type:     schema.TypeString,
				Optional: true,
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
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceAliCloudAdbDbClusterCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("db_node_class"); ok {
		request["DBClusterClass"] = v
	}
	if v, ok := d.GetOk("backup_set_id"); ok {
		request["BackupSetID"] = v
	}
	request["DBClusterNetworkType"] = d.Get("db_cluster_network_type")
	request["PayType"] = convertAdbDbClusterPayTypeRequest(d.Get("payment_type").(string))
	if v, ok := d.GetOk("storage_type"); ok {
		request["StorageType"] = v
	}
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
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
	if v, ok := d.GetOk("source_db_instance_name"); ok {
		request["SourceDBInstanceName"] = v
	}
	if v, ok := d.GetOk("db_cluster_name"); ok {
		request["DBClusterDescription"] = v
	}
	request["Mode"] = d.Get("mode")
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("storage_resource"); ok {
		request["StorageResource"] = v
	}
	if v, ok := d.GetOkExists("db_node_storage"); ok {
		request["DBNodeStorage"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("executor_count"); ok {
		request["ExecutorCount"] = v
	}
	request["DBClusterVersion"] = d.Get("db_cluster_version")
	request["DBClusterCategory"] = d.Get("db_cluster_category")
	if v, ok := d.GetOkExists("db_node_count"); ok {
		request["DBNodeGroupCount"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VPCId"] = v
	}
	if v, ok := d.GetOkExists("enable_ssl"); ok {
		request["EnableSSL"] = v
	}
	if v, ok := d.GetOk("elastic_io_resource"); ok {
		request["ElasticIOResource"] = v
	}
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	if v, ok := d.GetOk("restore_time"); ok {
		request["RestoreTime"] = v
	}
	if v, ok := d.GetOk("compute_resource"); ok {
		request["ComputeResource"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_db_cluster", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBClusterId"]))

	adbServiceV2 := AdbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutCreate), 15*time.Minute, adbServiceV2.AdbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{"Deleting"}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudAdbDbClusterUpdate(d, meta)
}

func resourceAliCloudAdbDbClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbDbCluster(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_db_cluster DescribeAdbDbCluster Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("compute_resource", objectRaw["ComputeResource"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("db_cluster_category", objectRaw["Category"])
	d.Set("db_cluster_name", objectRaw["DBClusterDescription"])
	d.Set("db_cluster_network_type", objectRaw["DBClusterNetworkType"])
	d.Set("db_node_class", objectRaw["DBNodeClass"])
	d.Set("db_node_count", objectRaw["DBNodeCount"])
	d.Set("db_node_storage", objectRaw["DBNodeStorage"])
	d.Set("disk_encryption", objectRaw["DiskEncryption"])
	d.Set("disk_performance_level", objectRaw["DiskPerformanceLevel"])
	d.Set("elastic_io_resource", objectRaw["ElasticIOResource"])
	d.Set("elastic_io_resource_size", objectRaw["ElasticIOResourceSize"])
	d.Set("executor_count", objectRaw["ExecutorCount"])
	d.Set("kms_id", objectRaw["KmsId"])
	d.Set("maintain_time", objectRaw["MaintainTime"])
	d.Set("mode", objectRaw["Mode"])
	d.Set("payment_type", convertAdbDbClusterItemsDBClusterPayTypeResponse(objectRaw["PayType"]))
	d.Set("port", objectRaw["Port"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("status", objectRaw["DBClusterStatus"])
	d.Set("storage_resource", objectRaw["StorageResource"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("zone_id", objectRaw["ZoneId"])

	tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	objectRaw, err = adbServiceV2.DescribeDbClusterDescribeDBClusters(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("compute_resource", objectRaw["ComputeResource"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("db_cluster_category", objectRaw["Category"])
	d.Set("db_cluster_name", objectRaw["DBClusterDescription"])
	d.Set("db_cluster_network_type", objectRaw["DBClusterNetworkType"])
	d.Set("db_cluster_version", objectRaw["DBVersion"])
	d.Set("db_node_class", objectRaw["DBNodeClass"])
	d.Set("db_node_count", objectRaw["DBNodeCount"])
	d.Set("db_node_storage", objectRaw["DBNodeStorage"])
	d.Set("elastic_io_resource", objectRaw["ElasticIOResource"])
	d.Set("executor_count", objectRaw["ExecutorCount"])
	d.Set("pay_type", objectRaw["PayType"])
	d.Set("port", formatInt(objectRaw["Port"]))
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("status", objectRaw["DBClusterStatus"])
	d.Set("storage_resource", objectRaw["StorageResource"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("zone_id", objectRaw["ZoneId"])

	tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	objectRaw, err = adbServiceV2.DescribeDbClusterDescribeAutoRenewAttribute(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("region_id", objectRaw["RegionId"])
	d.Set("renewal_status", objectRaw["RenewalStatus"])

	objectRaw, err = adbServiceV2.DescribeDbClusterDescribeDBClusterAccessWhiteList(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("security_ips", objectRaw["SecurityIPList"])

	objectRaw, err = adbServiceV2.DescribeDbClusterDescribeDBClusterSSL(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("enable_ssl", objectRaw["SSLEnabled"])

	objectRaw, err = adbServiceV2.DescribeDbClusterDescribeKernelVersion(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("kernel_version", objectRaw["KernelVersion"])

	return nil
}

func resourceAliCloudAdbDbClusterUpdate(d *schema.ResourceData, meta interface{}) error {
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
	if !d.IsNewResource() && d.HasChange("db_node_class") {
		update = true
		request["DBNodeClass"] = d.Get("db_node_class")
	}

	if !d.IsNewResource() && d.HasChange("db_node_storage") {
		update = true
		request["DBNodeStorage"] = d.Get("db_node_storage")
	}

	if v, ok := d.GetOk("modify_type"); ok {
		request["ModifyType"] = v
	}
	if !d.IsNewResource() && d.HasChange("executor_count") {
		update = true
		request["ExecutorCount"] = d.Get("executor_count")
	}

	if d.HasChange("elastic_io_resource_size") {
		update = true
		request["ElasticIOResourceSize"] = d.Get("elastic_io_resource_size")
	}

	if !d.IsNewResource() && d.HasChange("db_cluster_category") {
		update = true
	}
	request["DBClusterCategory"] = d.Get("db_cluster_category")
	if !d.IsNewResource() && d.HasChange("db_node_count") {
		update = true
		request["DBNodeGroupCount"] = d.Get("db_node_count")
	}

	if !d.IsNewResource() && d.HasChange("mode") {
		update = true
	}
	request["Mode"] = d.Get("mode")
	if !d.IsNewResource() && d.HasChange("elastic_io_resource") {
		update = true
		request["ElasticIOResource"] = d.Get("elastic_io_resource")
	}

	if !d.IsNewResource() && d.HasChange("storage_resource") {
		update = true
		request["StorageResource"] = d.Get("storage_resource")
	}

	if !d.IsNewResource() && d.HasChange("compute_resource") {
		update = true
		request["ComputeResource"] = d.Get("compute_resource")
	}

	if d.HasChange("disk_performance_level") {
		update = true
		request["DiskPerformanceLevel"] = d.Get("disk_performance_level")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"OperationDenied.OrderProcessing"}) || NeedRetry(err) {
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
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 15*time.Minute, adbServiceV2.AdbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
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
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "ModifyDBClusterDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("db_cluster_name") {
		update = true
	}
	request["DBClusterDescription"] = d.Get("db_cluster_name")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "ModifyDBClusterAccessWhiteList"
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
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "ModifyAutoRenewAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("period_unit"); ok {
		request["PeriodUnit"] = v
	}
	if d.HasChange("renewal_status") {
		update = true
		request["RenewalStatus"] = d.Get("renewal_status")
	}

	if v, ok := d.GetOkExists("auto_renew_period"); ok {
		request["Duration"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "ModifyDBClusterPayType"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DbClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("payment_type") {
		update = true
	}
	request["PayType"] = convertAdbDbClusterPayTypeRequest(d.Get("payment_type").(string))
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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

	if v, ok := d.GetOk("connection_string"); ok {
		request["ConnectionString"] = v
	}
	if !d.IsNewResource() && d.HasChange("enable_ssl") {
		update = true
	}
	request["EnableSSL"] = d.Get("enable_ssl")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 30*time.Second, adbServiceV2.AdbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "UpgradeKernelVersion"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("kernel_version") {
		update = true
		request["DBVersion"] = d.Get("kernel_version")
	}

	if v, ok := d.GetOkExists("switch_mode"); ok {
		request["SwitchMode"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "ModifyDBClusterVip"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("connection_string"); ok {
		request["ConnectionString"] = v
	}
	if !d.IsNewResource() && d.HasChange("vpc_id") {
		update = true
	}
	request["VPCId"] = d.Get("vpc_id")
	if !d.IsNewResource() && d.HasChange("vswitch_id") {
		update = true
	}
	request["VSwitchId"] = d.Get("vswitch_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "AttachUserENI"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	action = "DetachUserENI"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
			response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
		if err := adbServiceV2.SetResourceTags(d, "ALIYUN::ADB::CLUSTER"); err != nil {
			return WrapError(err)
		}
	}
	if !d.IsNewResource() && d.HasChange("region_id") {
		oldEntry, newEntry := d.GetChange("region_id")
		oldValue := oldEntry.(string)
		newValue := newEntry.(string)

		if oldValue != "" {
			action := "UntagResources"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["RegionId"] = client.RegionId
			localData := removed.List()
			rightDeltaData, err := jsonpath.Get("$[*].tag_key", localData)
			if err != nil {
				return WrapError(err)
			}
			localDataArray := rightDeltaData
			request["TagKey"] = localDataArray

			request["ResourceType"] = "ALIYUN::ADB::CLUSTER"
			jsonString := convertObjectToJsonString(request)
			jsonString, _ = sjson.Set(jsonString, "ResourceId.0", d.Id())
			_ = json.Unmarshal([]byte(jsonString), &request)

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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

		if newValue != "" {
			action := "TagResources"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["RegionId"] = client.RegionId
			localData := added.List()
			tagMapsArray := make([]interface{}, 0)
			for _, dataLoop := range localData {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["Value"] = dataLoopTmp["tag_value"]
				dataLoopMap["Key"] = dataLoopTmp["tag_key"]
				tagMapsArray = append(tagMapsArray, dataLoopMap)
			}
			request["Tag"] = tagMapsArray

			request["ResourceType"] = "ALIYUN::ADB::CLUSTER"
			jsonString := convertObjectToJsonString(request)
			jsonString, _ = sjson.Set(jsonString, "ResourceId.0", d.Id())
			_ = json.Unmarshal([]byte(jsonString), &request)

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
	}
	d.Partial(false)
	return resourceAliCloudAdbDbClusterRead(d, meta)
}

func resourceAliCloudAdbDbClusterDelete(d *schema.ResourceData, meta interface{}) error {

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
		response, err = client.RpcPost("adb", "2019-03-15", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidDBCluster.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}

func convertAdbDbClusterItemsDBClusterPayTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Prepaid":
		return "Subscription"
	case "Postpaid":
		return "PayAsYouGo"
	}
	return source
}
func convertAdbDbClusterPayTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Subscription":
		return "Prepaid"
	case "PayAsYouGo":
		return "Postpaid"
	}
	return source
}
