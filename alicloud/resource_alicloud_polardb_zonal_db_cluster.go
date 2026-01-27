// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/blues/jsonata-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPolardbZonalDbCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbZonalDbClusterCreate,
		Read:   resourceAliCloudPolardbZonalDbClusterRead,
		Update: resourceAliCloudPolardbZonalDbClusterUpdate,
		Delete: resourceAliCloudPolardbZonalDbClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(38 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"architecture": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_renew_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_network_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_category": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"SENormal"}, false),
			},
			"db_cluster_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_minor_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_node": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_class": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"db_node_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"db_node_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_node_num": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_node_target_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ens_region_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hot_standby_cluster": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"modify_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pay_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"period_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_auto_scale": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_pay_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"storage_space": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"target_minor_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"used_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPolardbZonalDbClusterCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("storage_type"); ok {
		request["StorageType"] = v
	}
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("db_minor_version"); ok {
		request["DBMinorVersion"] = v
	}
	request["PayType"] = convertPolardbZonalDbClusterPayTypeRequest(d.Get("pay_type").(string))
	request["DBNodeClass"] = d.Get("db_node_class")
	if v, ok := d.GetOk("db_cluster_description"); ok {
		request["DBClusterDescription"] = v
	}
	if v, ok := d.GetOk("creation_category"); ok {
		request["CreationCategory"] = v
	}
	request["DBType"] = d.Get("db_type")
	if v, ok := d.GetOk("storage_auto_scale"); ok {
		request["StorageAutoScale"] = v
	}
	if v, ok := d.GetOk("target_minor_version"); ok {
		request["TargetMinorVersion"] = v
	}
	if v, ok := d.GetOkExists("db_node_num"); ok {
		request["DBNodeNum"] = v
	}
	if v, ok := d.GetOk("storage_pay_type"); ok {
		request["StoragePayType"] = v
	}
	if v, ok := d.GetOk("cluster_network_type"); ok {
		request["ClusterNetworkType"] = v
	}
	if v, ok := d.GetOkExists("storage_space"); ok {
		request["StorageSpace"] = v
	}
	if v, ok := d.GetOk("architecture"); ok {
		request["Architecture"] = v
	}
	if v, ok := d.GetOk("hot_standby_cluster"); ok {
		request["HotStandbyCluster"] = v
	}
	if v, ok := d.GetOk("ens_region_id"); ok {
		request["EnsRegionId"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VPCId"] = v
	}
	request["DBVersion"] = d.Get("db_version")
	request["CloudProvider"] = "ENS"
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_zonal_db_cluster", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBClusterId"]))

	polardbServiceV2 := PolardbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutCreate), 10*time.Second, polardbServiceV2.PolardbZonalDbClusterStateRefreshFunc(d.Id(), "$.DBClusterStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudPolardbZonalDbClusterUpdate(d, meta)
}

func resourceAliCloudPolardbZonalDbClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbZonalDbCluster(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_zonal_db_cluster DescribePolardbZonalDbCluster Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("creation_category", objectRaw["Category"])
	d.Set("db_cluster_description", objectRaw["DBClusterDescription"])
	d.Set("db_node_target_class", objectRaw["DBClusterClass"])
	d.Set("db_type", objectRaw["DBType"])
	d.Set("db_version", objectRaw["DBVersion"])
	d.Set("pay_type", convertPolardbZonalDbClusterPayTypeResponse(objectRaw["PayType"]))
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("storage_pay_type", objectRaw["StoragePayType"])
	d.Set("storage_type", objectRaw["StorageType"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("architecture", objectRaw["Architecture"])
	d.Set("hot_standby_cluster", objectRaw["HotStandbyCluster"])

	dBNodesRaw := objectRaw["DBNodes"]
	dbNodeMaps := make([]map[string]interface{}, 0)
	if dBNodesRaw != nil {
		for _, dBNodesChildRaw := range convertToInterfaceArray(dBNodesRaw) {
			dbNodeMap := make(map[string]interface{})
			dBNodesChildRaw := dBNodesChildRaw.(map[string]interface{})
			dbNodeMap["db_node_id"] = dBNodesChildRaw["DBNodeId"]
			dbNodeMap["target_class"] = dBNodesChildRaw["DBNodeClass"]

			dbNodeMaps = append(dbNodeMaps, dbNodeMap)
		}
	}
	if err := d.Set("db_node", dbNodeMaps); err != nil {
		return err
	}

	e := jsonata.MustCompile("$.StorageSpace / 1024 / 1024")
	evaluation, _ := e.Eval(objectRaw)
	d.Set("storage_space", evaluation)
	objectRaw, err = polardbServiceV2.DescribeZonalDbClusterDescribeDBClusterVersionZonal(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("db_minor_version", objectRaw["DBMinorVersion"])

	objectRaw, err = polardbServiceV2.DescribeZonalDbClusterDescribeAutoRenewAttribute(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("auto_renew_status", objectRaw["RenewalStatus"])
	d.Set("duration", objectRaw["Duration"])
	d.Set("period_unit", objectRaw["PeriodUnit"])

	return nil
}

func resourceAliCloudPolardbZonalDbClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyAutoRenewAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterIds"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("duration") {
		update = true
		request["Duration"] = d.Get("duration")
	}

	if d.HasChange("period_unit") {
		update = true
		request["PeriodUnit"] = d.Get("period_unit")
	}

	if d.HasChange("auto_renew_status") {
		update = true
		request["RenewalStatus"] = d.Get("auto_renew_status")
	}

	request["CloudProvider"] = "ENS"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	action = "ModifyDBNodeClass"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("modify_type") {
		update = true
	}
	request["ModifyType"] = d.Get("modify_type")
	request["CloudProvider"] = "ENS"
	if d.HasChange("db_node_target_class") {
		update = true
	}
	request["DBNodeTargetClass"] = d.Get("db_node_target_class")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		polardbServiceV2 := PolardbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 3*time.Minute, polardbServiceV2.PolardbZonalDbClusterStateRefreshFunc(d.Id(), "$.DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ModifyDBClusterStorageSpace"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	request["CloudProvider"] = "ENS"
	if !d.IsNewResource() && d.HasChange("storage_space") {
		update = true
	}
	request["StorageSpace"] = d.Get("storage_space")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		polardbServiceV2 := PolardbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, polardbServiceV2.PolardbZonalDbClusterStateRefreshFunc(d.Id(), "$.DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ModifyDBNodesClass"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("modify_type") {
		update = true
	}
	request["ModifyType"] = d.Get("modify_type")
	if d.HasChange("db_node") {
		update = true
		if v, ok := d.GetOk("db_node"); ok || d.HasChange("db_node") {
			dBNodeMapsArray := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(v) {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["TargetClass"] = dataLoopTmp["target_class"]
				dataLoopMap["DBNodeId"] = dataLoopTmp["db_node_id"]
				dBNodeMapsArray = append(dBNodeMapsArray, dataLoopMap)
			}
			request["DBNode"] = dBNodeMapsArray
		}
	}

	request["CloudProvider"] = "ENS"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		polardbServiceV2 := PolardbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 4*time.Minute, polardbServiceV2.PolardbZonalDbClusterStateRefreshFunc(d.Id(), "$.DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "UpgradeDBClusterVersionZonal"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		polardbServiceV2 := PolardbServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutUpdate), 5*time.Minute, polardbServiceV2.PolardbZonalDbClusterStateRefreshFunc(d.Id(), "", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "FailoverDBClusterZonal"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	action = "ModifyDBClusterDescriptionZonal"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("db_cluster_description") {
		update = true
	}
	request["DBClusterDescription"] = d.Get("db_cluster_description")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
	return resourceAliCloudPolardbZonalDbClusterRead(d, meta)
}

func resourceAliCloudPolardbZonalDbClusterDelete(d *schema.ResourceData, meta interface{}) error {

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
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidDBClusterId.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}

func convertPolardbZonalDbClusterPayTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Postpaid":
		return "PayAsYouGo"
	case "Prepaid":
		return "Subscription"
	}
	return source
}
func convertPolardbZonalDbClusterPayTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "PayAsYouGo":
		return "Postpaid"
	case "Subscription":
		return "Prepaid"
	}
	return source
}
