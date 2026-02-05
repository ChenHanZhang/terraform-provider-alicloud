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

func resourceAliCloudPolardbDbCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbDbClusterCreate,
		Read:   resourceAliCloudPolardbDbClusterRead,
		Update: resourceAliCloudPolardbDbClusterUpdate,
		Delete: resourceAliCloudPolardbDbClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(17 * time.Minute),
			Update: schema.DefaultTimeout(29 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allow_shut_down": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"architecture": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"X86", "ARM"}, false),
			},
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"backup_retention_policy_on_cluster_deletion": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"ALL", "LATEST", "NONE"}, false),
			},
			"category": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"Normal", "Basic", "ArchiveNormal", "NormalMultimaster", "SENormal"}, false),
			},
			"clone_data_point": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"LATEST", "BackupID", "Timestamp"}, false),
			},
			"cluster_network_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"collector_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_option": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Normal", "CloneFromPolarDB", "CloneFromRDS", "CreateGdnStandby", "MigrationFromRDS"}, false),
			},
			"db_node_num": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_cluster_ip_array_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_cluster_ip_array_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_minor_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_node_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_node_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_version": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"8.0", "5.7", "5.6", "11", "14"}, false),
			},
			"default_time_zone": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encrypt_new_tables": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"ON", "OFF"}, false),
			},
			"encryption_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"existed_endpoint_switch_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"from_time_service": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"gdn_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hot_standby_cluster": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"ON", "OFF"}, false),
			},
			"is_switch_over_for_disaster": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"true", "false"}, false),
			},
			"loose_polar_log_bin": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"ON", "OFF"}, false),
			},
			"loose_xengine": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"ON", "OFF"}, false),
			},
			"loose_xengine_use_memory_pct": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lower_case_table_names": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"1", "0"}, false),
			},
			"maintain_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modify_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Cover", "Append", "Delete"}, false),
			},
			"modify_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Downgrade", "Upgrade"}, false),
			},
			"parameter_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_modifiable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"checking_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"force_restart": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"parameter_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"payment_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"PrePaid", "PostPaid"}, false),
			},
			"period": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Year", "Month"}, false),
			},
			"period_unit": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"planned_end_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"planned_flashing_off_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"planned_start_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"provisioned_iops": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"proxy_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role_arn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scale_max": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scale_min": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scale_ro_num_max": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scale_ro_num_min": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"seconds_until_auto_pause": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group_ids": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_ips": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverless_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"standby_az": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
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
				Computed: true,
			},
			"storage_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"PSL5", "PSL4", "ESSDPL0", "ESSDPL1", "ESSDPL2", "ESSDPL3", "ESSDAUTOPL"}, false),
			},
			"storage_upper_bound": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"strict_consistency": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"ON", "OFF"}, false),
			},
			"sub_category": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"normal_exclusive", "normal_general"}, false),
			},
			"tags": tagsSchema(),
			"tde_status": {
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
				Computed: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"white_list_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"IP", "SecurityGroup"}, false),
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudPolardbDbClusterCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDBCluster"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("scale_min"); ok {
		request["ScaleMin"] = v
	}
	request["PayType"] = convertPolardbDbClusterPayTypeRequest(d.Get("payment_type").(string))
	if v, ok := d.GetOk("storage_type"); ok {
		request["StorageType"] = v
	}
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	if v, ok := d.GetOk("strict_consistency"); ok {
		request["StrictConsistency"] = v
	}
	if v, ok := d.GetOk("tde_status"); ok {
		request["TDEStatus"] = convertPolardbDbClusterTDEStatusRequest(v.(string))
	}
	request["DBNodeClass"] = d.Get("db_node_class")
	if v, ok := d.GetOkExists("provisioned_iops"); ok {
		request["ProvisionedIops"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["DBType"] = d.Get("db_type")
	if v, ok := d.GetOk("storage_auto_scale"); ok {
		request["StorageAutoScale"] = v
	}
	if v, ok := d.GetOkExists("auto_renew"); ok {
		request["AutoRenew"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("scale_ro_num_max"); ok {
		request["ScaleRoNumMax"] = v
	}
	if v, ok := d.GetOk("storage_pay_type"); ok {
		request["StoragePayType"] = v
	}
	if v, ok := d.GetOk("cluster_network_type"); ok {
		request["ClusterNetworkType"] = v
	}
	if v, ok := d.GetOk("clone_data_point"); ok {
		request["CloneDataPoint"] = v
	}
	if v, ok := d.GetOkExists("db_node_num"); ok {
		request["DBNodeNum"] = v
	}
	if v, ok := d.GetOkExists("storage_space"); ok {
		request["StorageSpace"] = v
	}
	if v, ok := d.GetOk("loose_xengine"); ok {
		request["LooseXEngine"] = v
	}
	if v, ok := d.GetOk("hot_standby_cluster"); ok {
		request["HotStandbyCluster"] = v
	}
	if v, ok := d.GetOk("proxy_type"); ok {
		request["ProxyType"] = v
	}
	if v, ok := d.GetOk("scale_ro_num_min"); ok {
		request["ScaleRoNumMin"] = v
	}
	if v, ok := d.GetOk("gdn_id"); ok {
		request["GDNId"] = v
	}
	if v, ok := d.GetOk("scale_max"); ok {
		request["ScaleMax"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VPCId"] = v
	}
	if v, ok := d.GetOk("allow_shut_down"); ok {
		request["AllowShutDown"] = v
	}
	if v, ok := d.GetOk("loose_polar_log_bin"); ok {
		request["LoosePolarLogBin"] = v
	}
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("proxy_class"); ok {
		request["ProxyClass"] = v
	}
	if v, ok := d.GetOk("security_ips"); ok {
		request["SecurityIPList"] = v
	}
	if v, ok := d.GetOk("loose_xengine_use_memory_pct"); ok {
		request["LooseXEngineUseMemoryPct"] = v
	}
	if v, ok := d.GetOk("db_minor_version"); ok {
		request["DBMinorVersion"] = v
	}
	if v, ok := d.GetOkExists("storage_upper_bound"); ok {
		request["StorageUpperBound"] = v
	}
	if v, ok := d.GetOk("category"); ok {
		request["CreationCategory"] = v
	}
	if v, ok := d.GetOk("source_resource_id"); ok {
		request["SourceResourceId"] = v
	}
	if v, ok := d.GetOk("backup_retention_policy_on_cluster_deletion"); ok {
		request["BackupRetentionPolicyOnClusterDeletion"] = v
	}
	if v, ok := d.GetOk("default_time_zone"); ok {
		request["DefaultTimeZone"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["DBClusterDescription"] = v
	}
	if v, ok := d.GetOk("standby_az"); ok {
		request["StandbyAZ"] = v
	}
	if v, ok := d.GetOk("lower_case_table_names"); ok {
		request["LowerCaseTableNames"] = v
	}
	if v, ok := d.GetOk("architecture"); ok {
		request["Architecture"] = v
	}
	if v, ok := d.GetOk("serverless_type"); ok {
		request["ServerlessType"] = v
	}
	if v, ok := d.GetOk("creation_option"); ok {
		request["CreationOption"] = v
	}
	if v, ok := d.GetOk("parameter_group_id"); ok {
		request["ParameterGroupId"] = v
	}
	request["DBVersion"] = d.Get("db_version")
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"undefined"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polardb_cluster", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["DBClusterId"]))

	polardbServiceV2 := PolardbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutCreate), 2*time.Minute, polardbServiceV2.PolardbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudPolardbDbClusterUpdate(d, meta)
}

func resourceAliCloudPolardbDbClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbDbCluster(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_cluster DescribePolardbDbCluster Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("architecture", objectRaw["Architecture"])
	d.Set("category", objectRaw["Category"])
	d.Set("cluster_network_type", objectRaw["DBClusterNetworkType"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("db_type", objectRaw["DBType"])
	d.Set("db_version", objectRaw["DBVersion"])
	d.Set("description", objectRaw["DBClusterDescription"])
	d.Set("hot_standby_cluster", convertPolardbDbClusterHotStandbyClusterResponse(objectRaw["HotStandbyCluster"]))
	d.Set("maintain_time", objectRaw["MaintainTime"])
	d.Set("payment_type", convertPolardbDbClusterPayTypeResponse(objectRaw["PayType"]))
	d.Set("provisioned_iops", formatInt(objectRaw["ProvisionedIops"]))
	d.Set("proxy_type", objectRaw["ProxyType"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("status", objectRaw["DBClusterStatus"])
	d.Set("storage_pay_type", objectRaw["StoragePayType"])
	d.Set("storage_space", objectRaw["StorageSpace"])
	d.Set("storage_type", convertPolardbDbClusterStorageTypeResponse(objectRaw["StorageType"]))
	d.Set("strict_consistency", objectRaw["StrictConsistency"])
	d.Set("sub_category", objectRaw["SubCategory"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VPCId"])
	d.Set("zone_id", objectRaw["ZoneIds"])

	tagsMaps := objectRaw["Tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	objectRaw, err = polardbServiceV2.DescribeDbClusterDescribeDBClusterParameters(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("db_type", objectRaw["DBType"])
	d.Set("db_version", objectRaw["DBVersion"])

	parameterRaw, _ := jsonpath.Get("$.RunningParameters.Parameter", objectRaw)
	parametersMaps := make([]map[string]interface{}, 0)
	if parameterRaw != nil {
		for _, parameterChildRaw := range convertToInterfaceArray(parameterRaw) {
			parametersMap := make(map[string]interface{})
			parameterChildRaw := parameterChildRaw.(map[string]interface{})
			parametersMap["checking_code"] = parameterChildRaw["CheckingCode"]
			parametersMap["data_type"] = parameterChildRaw["DataType"]
			parametersMap["force_restart"] = parameterChildRaw["ForceRestart"]
			parametersMap["is_modifiable"] = parameterChildRaw["IsModifiable"]
			parametersMap["name"] = parameterChildRaw["ParameterName"]
			parametersMap["parameter_description"] = parameterChildRaw["ParameterDescription"]
			parametersMap["parameter_status"] = parameterChildRaw["ParameterStatus"]
			parametersMap["value"] = parameterChildRaw["ParameterValue"]

			parametersMaps = append(parametersMaps, parametersMap)
		}
	}
	if err := d.Set("parameters", parametersMaps); err != nil {
		return err
	}

	objectRaw, err = polardbServiceV2.DescribeDbClusterDescribeDBClusterVersion(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("db_minor_version", objectRaw["DBMinorVersion"])
	d.Set("db_version", objectRaw["DBVersion"])

	checkValue00 := d.Get("serverless_type")
	if checkValue00 == "" {
		objectRaw, err = polardbServiceV2.DescribeDbClusterDescribeDBClusterServerlessConf(d.Id())
		if err != nil && !NotFoundError(err) {
			return WrapError(err)
		}

		d.Set("allow_shut_down", objectRaw["AllowShutDown"])
		d.Set("scale_max", objectRaw["ScaleMax"])
		d.Set("scale_min", objectRaw["ScaleMin"])
		d.Set("scale_ro_num_max", objectRaw["ScaleRoNumMax"])
		d.Set("scale_ro_num_min", objectRaw["ScaleRoNumMin"])
		d.Set("seconds_until_auto_pause", objectRaw["SecondsUntilAutoPause"])

	}
	objectRaw, err = polardbServiceV2.DescribeDbClusterDescribeDBClusterAccessWhitelist(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("security_ips", objectRaw["SecurityIps"])
	d.Set("db_cluster_ip_array_attribute", objectRaw["DBClusterIPArrayAttribute"])
	d.Set("db_cluster_ip_array_name", objectRaw["DBClusterIPArrayName"])

	objectRaw, err = polardbServiceV2.DescribeDbClusterDescribeDBClusterTDE(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("encrypt_new_tables", objectRaw["EncryptNewTables"])
	d.Set("encryption_key", objectRaw["EncryptionKey"])
	d.Set("tde_status", objectRaw["TDEStatus"])

	objectRaw, err = polardbServiceV2.DescribeDbClusterDescribeDBClusterAuditLogCollector(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("collector_status", objectRaw["CollectorStatus"])

	return nil
}

func resourceAliCloudPolardbDbClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyDBClusterDescription"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
	}
	request["DBClusterDescription"] = d.Get("description")
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
	action = "ModifyDBClusterAccessWhitelist"
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
	if !d.IsNewResource() && d.HasChange("security_ips") {
		update = true
		request["SecurityIps"] = d.Get("security_ips")
	}

	if v, ok := d.GetOk("white_list_type"); ok {
		request["WhiteListType"] = v
	}
	if v, ok := d.GetOk("security_group_ids"); ok {
		request["SecurityGroupIds"] = v
	}
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
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, polardbServiceV2.PolardbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
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
	action = "ModifyDBClusterAuditLogCollector"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if d.HasChange("collector_status") {
		update = true
	}
	request["CollectorStatus"] = d.Get("collector_status")
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
	action = "ModifyAutoRenewAttribute"
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

	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	if d.HasChange("renewal_status") {
		update = true
		request["RenewalStatus"] = d.Get("renewal_status")
	}

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
	action = "TransformDBClusterPayType"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if _, ok := d.GetOk("resource_group_id"); ok && !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	if !d.IsNewResource() && d.HasChange("payment_type") {
		update = true
	}
	request["PayType"] = convertPolardbDbClusterPayTypeRequest(d.Get("payment_type").(string))
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
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
	action = "ModifyDBClusterTDE"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if !d.IsNewResource() && d.HasChange("tde_status") {
		update = true
	}
	request["TDEStatus"] = d.Get("tde_status")
	if v, ok := d.GetOk("role_arn"); ok {
		request["RoleArn"] = v
	}
	if d.HasChange("encrypt_new_tables") {
		update = true
		request["EncryptNewTables"] = d.Get("encrypt_new_tables")
	}

	if d.HasChange("encryption_key") {
		update = true
		request["EncryptionKey"] = d.Get("encryption_key")
	}

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
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, polardbServiceV2.PolardbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ModifyDBNodeClass"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	request["DBNodeTargetClass"] = d.Get("db_node_class")
	if v, ok := d.GetOk("planned_end_time"); ok {
		request["PlannedEndTime"] = v
	}
	if v, ok := d.GetOk("planned_flashing_off_time"); ok {
		request["PlannedFlashingOffTime"] = v
	}
	if v, ok := d.GetOk("db_node_type"); ok {
		request["DBNodeType"] = v
	}
	request["ModifyType"] = d.Get("modify_type")
	if v, ok := d.GetOk("planned_start_time"); ok {
		request["PlannedStartTime"] = v
	}
	if d.HasChange("sub_category") {
		update = true
		request["SubCategory"] = d.Get("sub_category")
	}

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
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, polardbServiceV2.PolardbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
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
	action = "ModifyDBClusterPrimaryZone"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	if v, ok := d.GetOk("is_switch_over_for_disaster"); ok {
		request["IsSwitchOverForDisaster"] = v
	}
	if v, ok := d.GetOk("planned_end_time"); ok {
		request["PlannedEndTime"] = v
	}
	if !d.IsNewResource() && d.HasChange("vpc_id") {
		update = true
		request["VPCId"] = d.Get("vpc_id")
	}

	if !d.IsNewResource() && d.HasChange("vswitch_id") {
		update = true
		request["VSwitchId"] = d.Get("vswitch_id")
	}

	if v, ok := d.GetOk("planned_start_time"); ok {
		request["PlannedStartTime"] = v
	}
	request["ZoneType"] = d.Get("zone_type")
	if v, ok := d.GetOkExists("from_time_service"); ok {
		request["FromTimeService"] = v
	}
	if !d.IsNewResource() && d.HasChange("zone_id") {
		update = true
	}
	request["ZoneId"] = d.Get("zone_id")
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
	action = "ModifyDBClusterStorageSpace"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("planned_end_time"); ok {
		request["PlannedEndTime"] = v
	}
	if v, ok := d.GetOk("planned_start_time"); ok {
		request["PlannedStartTime"] = v
	}
	if d.HasChange("sub_category") {
		update = true
		request["SubCategory"] = d.Get("sub_category")
	}

	if !d.IsNewResource() && d.HasChange("storage_space") {
		update = true
	}
	request["StorageSpace"] = d.Get("storage_space")
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
	action = "ModifyDBClusterVpc"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBClusterId"] = d.Id()

	request["ExistedEndpointSwitchType"] = d.Get("existed_endpoint_switch_type")
	if !d.IsNewResource() && d.HasChange("vpc_id") {
		update = true
		request["VPCId"] = d.Get("vpc_id")
	}

	if !d.IsNewResource() && d.HasChange("vswitch_id") {
		update = true
	}
	request["VSwitchId"] = d.Get("vswitch_id")
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

	if d.HasChange("tags") {
		polardbServiceV2 := PolardbServiceV2{client}
		if err := polardbServiceV2.SetResourceTags(d, "cluster"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudPolardbDbClusterRead(d, meta)
}

func resourceAliCloudPolardbDbClusterDelete(d *schema.ResourceData, meta interface{}) error {

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
			if IsExpectedErrors(err, []string{"OperationDenied.PolarDBClusterStatus", "OperationDenied.DBClusterStatus", "OperationDenied.ReadPolarDBClusterStatus"}) || NeedRetry(err) {
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

	polardbServiceV2 := PolardbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 10*time.Second, polardbServiceV2.PolardbDbClusterStateRefreshFunc(d.Id(), "DBClusterStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}

func convertPolardbDbClusterHotStandbyClusterResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "StandbyClusterON":
		return "ON"
	case "StandbyClusterOFF":
		return "OFF"
	case "equal":
		return "EQUAL"
	}
	return source
}
func convertPolardbDbClusterPayTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Postpaid":
		return "PostPaid"
	case "Prepaid":
		return "PrePaid"
	}
	return source
}
func convertPolardbDbClusterStorageTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "HighPerformance":
		return "PSL5"
	}
	return source
}
func convertPolardbDbClusterPayTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "PostPaid":
		return "Postpaid"
	case "PrePaid":
		return "Prepaid"
	}
	return source
}
func convertPolardbDbClusterTDEStatusRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Enabled":
		return "true"
	case "Disabled":
		return "false"
	}
	return source
}
