package alicloud

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAlicloudDtsSubscriptionJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudDtsSubscriptionJobCreate,
		Read:   resourceAlicloudDtsSubscriptionJobRead,
		Update: resourceAlicloudDtsSubscriptionJobUpdate,
		Delete: resourceAlicloudDtsSubscriptionJobDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"auto_pay": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_start": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"checkpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compute_unit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"database_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"db_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay_notice": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"delay_phone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay_rule_time": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_endpoint_engine_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"MySQL","PolarDB","polardb_o","polardb_pg","Redis","DRDS","PostgreSQL","odps","oracle","mongodb","tidb","ADS","Greenplum","MSSQL","kafka","DataHub","clickhouse","DB2","as400","Tablestore"}, false),
			},
			"destination_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dts_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dts_job_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"error_notice": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"error_phone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_class": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"large", "medium", "micro", "small", "xlarge", "xxlarge"}, false),
			},
			"job_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"period": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Month", "Year"}, false),
			},
			"quantity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"reserve": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_database_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_engine_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"MySQL", "PolarDB", "polardb_o", "polardb_pg", "Redis", "DRDS", "PostgreSQL", "odps", "oracle", "mongodb", "tidb", "ADS", "Greenplum", "MSSQL", "kafka", "DataHub", "clickhouse", "DB2", "as400", "Tablestore"}, false),
			},
			"source_endpoint_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"CEN", "DRDS", "ECS", "Express", "LocalInstance", "PolarDB", "RDS", "dg"}, false),
			},
			"source_endpoint_oracle_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_owner_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_endpoint_role": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_endpoint_user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringInSlice([]string{"Abnormal", "Downgrade", "Locked", "Normal", "NotStarted", "NotStarted", "PreCheckPass", "PrecheckFailed", "Prechecking", "Retrying", "Starting", "Upgrade"}, false),
			},
			"subscription_data_type_ddl": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"subscription_data_type_dml": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"subscription_instance_network_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"classic", "vpc"}, false),
			},
			"subscription_instance_vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subscription_instance_vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sync_architecture": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"bidirectional", "oneway"}, false),
			},
			"synchronization_direction": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": tagsSchema(),
			"used_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAlicloudDtsSubscriptionJobCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "CreateDtsInstance"
	request := make(map[string]interface{})
	conn, err := client.NewDtsClient()
	if err != nil {
		return WrapError(err)
	}
	if v, ok := d.GetOkExists("auto_pay"); ok {
		request["AutoPay"] = v
	}
	if v, ok := d.GetOkExists("auto_start"); ok {
		request["AutoStart"] = v
	}
	if v, ok := d.GetOk("compute_unit"); ok {
		request["ComputeUnit"] = v
	}
	if v, ok := d.GetOk("database_count"); ok {
		request["DatabaseCount"] = v
	}
	if v, ok := d.GetOk("destination_endpoint_engine_name"); ok {
		request["DestinationEndpointEngineName"] = v
	}
	if v, ok := d.GetOk("destination_region"); ok {
		request["DestinationRegion"] = v
	}
	if v, ok := d.GetOk("instance_class"); ok {
		request["InstanceClass"] = v
	}
	if v, ok := d.GetOk("job_id"); ok {
		request["JobId"] = v
	}
	if v, ok := d.GetOk("payment_type"); ok {
		request["PayType"] = v
	}
	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("quantity"); ok {
		request["Quantity"] = v
	}
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("source_endpoint_engine_name"); ok {
		request["SourceEndpointEngineName"] = v
	}
	if v, ok := d.GetOk("source_endpoint_region"); ok {
		request["SourceRegion"] = v
	}
	if v, ok := d.GetOk("sync_architecture"); ok {
		request["SyncArchitecture"] = v
	}
	request["Type"] = "SUBSCRIBE"
	if v, ok := d.GetOk("used_time"); ok {
		request["UsedTime"] = v
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_dts_subscription_job", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["JobId"]))

	return resourceAlicloudDtsSubscriptionJobUpdate(d, meta)
}
func resourceAlicloudDtsSubscriptionJobRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dtsService := DtsService{client}
	object, err := dtsService.DescribeDtsSubscriptionJob(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_dts_subscription_job dtsService.DescribeDtsSubscriptionJob Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	d.Set("checkpoint", fmt.Sprint(formatInt(object["Checkpoint"])))
	d.Set("db_list", object["DbObject"])
	d.Set("dts_instance_id", object["DtsInstanceID"])
	d.Set("dts_job_name", object["DtsJobName"])
	d.Set("payment_type", object["PayType"])
	d.Set("source_endpoint_database_name", object["SourceEndpoint"].(map[string]interface{})["DatabaseName"])
	d.Set("source_endpoint_engine_name", object["SourceEndpoint"].(map[string]interface{})["EngineName"])
	d.Set("source_endpoint_ip", object["SourceEndpoint"].(map[string]interface{})["Ip"])
	d.Set("source_endpoint_instance_id", object["SourceEndpoint"].(map[string]interface{})["InstanceID"])
	d.Set("source_endpoint_instance_type", object["SourceEndpoint"].(map[string]interface{})["InstanceType"])
	d.Set("source_endpoint_oracle_sid", object["SourceEndpoint"].(map[string]interface{})["OracleSID"])
	d.Set("source_endpoint_owner_id", object["SourceEndpoint"].(map[string]interface{})["AliyunUid"])
	d.Set("source_endpoint_port", object["SourceEndpoint"].(map[string]interface{})["Port"])
	d.Set("source_endpoint_region", object["SourceEndpoint"].(map[string]interface{})["Region"])
	d.Set("source_endpoint_role", object["SourceEndpoint"].(map[string]interface{})["RoleName"])
	d.Set("source_endpoint_user_name", object["SourceEndpoint"].(map[string]interface{})["UserName"])
	d.Set("status", object["Status"])
	d.Set("subscription_data_type_ddl", object["SubscriptionDataType"].(map[string]interface{})["Ddl"])
	d.Set("subscription_data_type_dml", object["SubscriptionDataType"].(map[string]interface{})["Dml"])
	d.Set("subscription_instance_network_type", object["Reserved"].(map[string]interface{})["srcNetType"])
	d.Set("subscription_instance_vpc_id", object["Reserved"].(map[string]interface{})["vpcId"])
	d.Set("subscription_instance_vswitch_id", object["Reserved"].(map[string]interface{})["vswitchId"])
	listTagResourcesObject, err := dtsService.ListTagResources(d.Id(), "ALIYUN::DTS::INSTANCE")
	if err != nil {
		return WrapError(err)
	}
	d.Set("tags", tagsToMap(listTagResourcesObject))
	return nil
}
func resourceAlicloudDtsSubscriptionJobUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	dtsService := DtsService{client}
	var response map[string]interface{}
	d.Partial(true)

	if d.HasChange("tags") {
		if err := dtsService.SetResourceTags(d, "ALIYUN::DTS::INSTANCE"); err != nil {
			return WrapError(err)
		}
		d.SetPartial("tags")
	}
	update := false
	request := map[string]interface{}{
		"DtsJobId": d.Id(),
	}
	if d.HasChange("dts_job_name") {
		update = true
		if v, ok := d.GetOk("dts_job_name"); ok {
			request["DtsJobName"] = v
		}
	}
	request["RegionId"] = client.RegionId
	if update {
		action := "ModifyDtsJobName"
		conn, err := client.NewDtsClient()
		if err != nil {
			return WrapError(err)
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
		if fmt.Sprint(response["Success"]) == "false" {
			return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
		}
		d.SetPartial("dts_job_name")
	}
	update = false
	modifyDtsJobPasswordReq := map[string]interface{}{
		"DtsJobId": d.Id(),
	}
	modifyDtsJobPasswordReq["Endpoint"] = "src"
	modifyDtsJobPasswordReq["RegionId"] = client.RegionId
	if d.HasChange("source_endpoint_password") {
		update = true
		if v, ok := d.GetOk("source_endpoint_password"); ok {
			modifyDtsJobPasswordReq["Password"] = v
		}
	}
	if d.HasChange("source_endpoint_user_name") {
		update = true
		if v, ok := d.GetOk("source_endpoint_user_name"); ok {
			modifyDtsJobPasswordReq["UserName"] = v
		}
	}
	if update {
		action := "ModifyDtsJobPassword"
		conn, err := client.NewDtsClient()
		if err != nil {
			return WrapError(err)
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, modifyDtsJobPasswordReq, &util.RuntimeOptions{})
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, modifyDtsJobPasswordReq)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		if fmt.Sprint(response["Success"]) == "false" {
			return WrapError(fmt.Errorf("%s failed, response: %v", action, response))
		}
		d.SetPartial("source_endpoint_password")
		d.SetPartial("source_endpoint_user_name")
	}
	update = false
	configureSubscriptionReq := map[string]interface{}{
		"DtsJobId": d.Id(),
	}
	if d.HasChange("db_list") {
		update = true
	}
	if v, ok := d.GetOk("db_list"); ok {
		configureSubscriptionReq["DbList"] = v
	}

	if v, ok := d.GetOk("dts_job_name"); ok {
		configureSubscriptionReq["DtsJobName"] = v
	}
	configureSubscriptionReq["RegionId"] = client.RegionId
	if d.HasChange("subscription_instance_network_type") {
		update = true
	}
	if v, ok := d.GetOk("subscription_instance_network_type"); ok {
		configureSubscriptionReq["SubscriptionInstanceNetworkType"] = v
	}
	if d.HasChange("checkpoint") {
		update = true
		if v, ok := d.GetOk("checkpoint"); ok {
			configureSubscriptionReq["Checkpoint"] = v
		}
	}
	if d.HasChange("source_endpoint_database_name") {
		update = true
		if v, ok := d.GetOk("source_endpoint_database_name"); ok {
			configureSubscriptionReq["SourceEndpointDatabaseName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("source_endpoint_engine_name") {
		update = true
		if v, ok := d.GetOk("source_endpoint_engine_name"); ok {
			configureSubscriptionReq["SourceEndpointEngineName"] = v
		}
	}
	if d.HasChange("source_endpoint_ip") {
		update = true
		if v, ok := d.GetOk("source_endpoint_ip"); ok {
			configureSubscriptionReq["SourceEndpointIP"] = v
		}
	}
	if d.HasChange("source_endpoint_instance_id") {
		update = true
		if v, ok := d.GetOk("source_endpoint_instance_id"); ok {
			configureSubscriptionReq["SourceEndpointInstanceID"] = v
		}
	}
	if d.HasChange("source_endpoint_instance_type") {
		update = true
		if v, ok := d.GetOk("source_endpoint_instance_type"); ok {
			configureSubscriptionReq["SourceEndpointInstanceType"] = v
		}
	}
	if d.HasChange("source_endpoint_oracle_sid") {
		update = true
		if v, ok := d.GetOk("source_endpoint_oracle_sid"); ok {
			configureSubscriptionReq["SourceEndpointOracleSID"] = v
		}
	}
	if d.HasChange("source_endpoint_owner_id") {
		update = true
		if v, ok := d.GetOk("source_endpoint_owner_id"); ok {
			configureSubscriptionReq["SourceEndpointOwnerID"] = v
		}
	}

	if v, ok := d.GetOk("source_endpoint_password"); ok {
		configureSubscriptionReq["SourceEndpointPassword"] = v
	}
	if d.HasChange("source_endpoint_port") {
		update = true
		if v, ok := d.GetOk("source_endpoint_port"); ok {
			configureSubscriptionReq["SourceEndpointPort"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("source_endpoint_region") {
		update = true
		if v, ok := d.GetOk("source_endpoint_region"); ok {
			configureSubscriptionReq["SourceEndpointRegion"] = v
		}
	}
	if d.HasChange("source_endpoint_role") {
		update = true
		if v, ok := d.GetOk("source_endpoint_role"); ok {
			configureSubscriptionReq["SourceEndpointRole"] = v
		}
	}

	if v, ok := d.GetOk("source_endpoint_user_name"); ok {
		configureSubscriptionReq["SourceEndpointUserName"] = v
	}
	if d.HasChange("subscription_data_type_ddl") || d.IsNewResource() {
		update = true
		if v, ok := d.GetOkExists("subscription_data_type_ddl"); ok {
			configureSubscriptionReq["SubscriptionDataTypeDDL"] = v
		}
	}
	if d.HasChange("subscription_data_type_dml") || d.IsNewResource() {
		update = true
		if v, ok := d.GetOkExists("subscription_data_type_dml"); ok {
			configureSubscriptionReq["SubscriptionDataTypeDML"] = v
		}
	}
	if d.HasChange("subscription_instance_vpc_id") {
		update = true
		if v, ok := d.GetOk("subscription_instance_vpc_id"); ok {
			configureSubscriptionReq["SubscriptionInstanceVPCId"] = v
		}
	}
	if d.HasChange("subscription_instance_vswitch_id") {
		update = true
		if v, ok := d.GetOk("subscription_instance_vswitch_id"); ok {
			configureSubscriptionReq["SubscriptionInstanceVSwitchId"] = v
		}
	}
	if update {
		if v, ok := d.GetOkExists("delay_notice"); ok {
			configureSubscriptionReq["DelayNotice"] = v
		}
		if v, ok := d.GetOk("delay_phone"); ok {
			configureSubscriptionReq["DelayPhone"] = v
		}
		if v, ok := d.GetOk("delay_rule_time"); ok {
			configureSubscriptionReq["DelayRuleTime"] = v
		}
		if v, ok := d.GetOkExists("error_notice"); ok {
			configureSubscriptionReq["ErrorNotice"] = v
		}
		if v, ok := d.GetOk("error_phone"); ok {
			configureSubscriptionReq["ErrorPhone"] = v
		}
		if v, ok := d.GetOk("reserve"); ok {
			configureSubscriptionReq["Reserve"] = v
		}
		action := "ConfigureSubscription"
		conn, err := client.NewDtsClient()
		if err != nil {
			return WrapError(err)
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, configureSubscriptionReq, &util.RuntimeOptions{})
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, configureSubscriptionReq)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("db_list")
		d.SetPartial("dts_job_name")
		d.SetPartial("subscription_instance_network_type")
		d.SetPartial("checkpoint")
		d.SetPartial("source_endpoint_database_name")
		d.SetPartial("source_endpoint_engine_name")
		d.SetPartial("source_endpoint_ip")
		d.SetPartial("source_endpoint_instance_id")
		d.SetPartial("source_endpoint_instance_type")
		d.SetPartial("source_endpoint_oracle_sid")
		d.SetPartial("source_endpoint_owner_id")
		d.SetPartial("source_endpoint_password")
		d.SetPartial("source_endpoint_port")
		d.SetPartial("source_endpoint_region")
		d.SetPartial("source_endpoint_role")
		d.SetPartial("source_endpoint_user_name")
		d.SetPartial("subscription_data_type_ddl")
		d.SetPartial("subscription_data_type_dml")
		d.SetPartial("subscription_instance_vpc_id")
		d.SetPartial("subscription_instance_vswitch_id")
	}
	if d.HasChange("status") {
		object, err := dtsService.DescribeDtsSubscriptionJob(d.Id())
		if err != nil {
			return WrapError(err)
		}
		target := d.Get("status").(string)
		if object["Status"].(string) != target {
			if target == "NotStarted" {
				request := map[string]interface{}{
					"DtsJobId": d.Id(),
				}
				request["RegionId"] = client.RegionId
				if v, ok := d.GetOk("synchronization_direction"); ok {
					request["SynchronizationDirection"] = v
				}
				action := "StopDtsJob"
				conn, err := client.NewDtsClient()
				if err != nil {
					return WrapError(err)
				}
				wait := incrementalWait(3*time.Second, 3*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
			if target == "Starting" {
				request := map[string]interface{}{
					"DtsJobId": d.Id(),
				}
				request["RegionId"] = client.RegionId
				if v, ok := d.GetOk("synchronization_direction"); ok {
					request["SynchronizationDirection"] = v
				}
				action := "StartDtsJob"
				conn, err := client.NewDtsClient()
				if err != nil {
					return WrapError(err)
				}
				wait := incrementalWait(3*time.Second, 3*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
			d.SetPartial("status")
		}
	}
	d.Partial(false)
	return resourceAlicloudDtsSubscriptionJobRead(d, meta)
}
func resourceAlicloudDtsSubscriptionJobDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "DeleteDtsJob"
	var response map[string]interface{}
	conn, err := client.NewDtsClient()
	if err != nil {
		return WrapError(err)
	}
	request := map[string]interface{}{
		"DtsJobId": d.Id(),
	}

	if v, ok := d.GetOk("dts_instance_id"); ok {
		request["DtsInstanceId"] = v
	}
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("synchronization_direction"); ok {
		request["SynchronizationDirection"] = v
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2020-01-01"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
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
	return nil
}
