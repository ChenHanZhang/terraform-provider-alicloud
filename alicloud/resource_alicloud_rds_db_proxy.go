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

func resourceAliCloudRdsDbProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRdsDbProxyCreate,
		Read:   resourceAliCloudRdsDbProxyRead,
		Update: resourceAliCloudRdsDbProxyUpdate,
		Delete: resourceAliCloudRdsDbProxyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"causal_consist_read_timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_db_proxy_features": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_proxy_connect_string": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_proxy_connect_string_net_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_proxy_connect_string_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_proxy_endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"db_proxy_instance_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"db_proxy_new_connect_string": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_endpoint_aliases": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_endpoint_min_slave_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_endpoint_operator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_endpoint_read_write_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_proxy_endpoint_aliases": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_proxy_instance_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_proxy_new_connect_string_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_proxy_ssl_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"effective_specific_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"effective_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistent_connection_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"read_only_instance_distribution_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"read_only_instance_max_delay_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"read_only_instance_weight": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudRdsDbProxyCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ModifyDBProxy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_instance_id"); ok {
		request["DBInstanceId"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("db_proxy_instance_type"); ok {
		request["DBProxyInstanceType"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	request["ConfigDBProxyService"] = "Startup"
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VPCId"] = v
	}
	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	if v, ok := d.GetOkExists("db_proxy_instance_num"); ok {
		request["DBProxyInstanceNum"] = v
	}
	request["InstanceNetworkType"] = "VPC"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rds_db_proxy", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["DBInstanceId"]))

	rdsServiceV2 := RdsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, rdsServiceV2.RdsDbProxyStateRefreshFunc(d.Id(), "$.DBProxyInstanceStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudRdsDbProxyUpdate(d, meta)
}

func resourceAliCloudRdsDbProxyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rdsServiceV2 := RdsServiceV2{client}

	objectRaw, err := rdsServiceV2.DescribeRdsDbProxy(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rds_db_proxy DescribeRdsDbProxy Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("db_proxy_instance_num", objectRaw["DBProxyInstanceNum"])
	d.Set("db_proxy_instance_type", convertRdsDbProxyDBProxyInstanceTypeResponse(objectRaw["DBProxyInstanceType"]))
	d.Set("persistent_connection_status", objectRaw["DBProxyPersistentConnectionStatus"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])

	dBProxyConnectStringItemsRawArrayObj, _ := jsonpath.Get("$.DBProxyConnectStringItems.DBProxyConnectStringItems[*]", objectRaw)
	dBProxyConnectStringItemsRawArray := make([]interface{}, 0)
	if dBProxyConnectStringItemsRawArrayObj != nil {
		dBProxyConnectStringItemsRawArray = convertToInterfaceArray(dBProxyConnectStringItemsRawArrayObj)
	}
	dBProxyConnectStringItemsRaw := make(map[string]interface{})
	if len(dBProxyConnectStringItemsRawArray) > 0 {
		dBProxyConnectStringItemsRaw = dBProxyConnectStringItemsRawArray[0].(map[string]interface{})
	}

	d.Set("db_proxy_connect_string", dBProxyConnectStringItemsRaw["DBProxyConnectString"])
	d.Set("db_proxy_connect_string_net_type", dBProxyConnectStringItemsRaw["DBProxyConnectStringNetType"])
	d.Set("db_proxy_connect_string_port", dBProxyConnectStringItemsRaw["DBProxyConnectStringPort"])
	d.Set("db_proxy_endpoint_id", dBProxyConnectStringItemsRaw["DBProxyEndpointId"])

	dbProxyEndpointItemsRawArrayObj, _ := jsonpath.Get("$.DbProxyEndpointItems.DbProxyEndpointItems[*]", objectRaw)
	dbProxyEndpointItemsRawArray := make([]interface{}, 0)
	if dbProxyEndpointItemsRawArrayObj != nil {
		dbProxyEndpointItemsRawArray = convertToInterfaceArray(dbProxyEndpointItemsRawArrayObj)
	}
	dbProxyEndpointItemsRaw := make(map[string]interface{})
	if len(dbProxyEndpointItemsRawArray) > 0 {
		dbProxyEndpointItemsRaw = dbProxyEndpointItemsRawArray[0].(map[string]interface{})
	}

	d.Set("db_proxy_endpoint_aliases", dbProxyEndpointItemsRaw["DbProxyEndpointAliases"])

	objectRaw, err = rdsServiceV2.DescribeDbProxyDescribeDBProxyEndpoint(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("db_proxy_connect_string", objectRaw["DBProxyConnectString"])
	d.Set("db_proxy_connect_string_net_type", objectRaw["DBProxyConnectStringNetType"])
	d.Set("db_proxy_connect_string_port", objectRaw["DBProxyConnectStringPort"])
	d.Set("db_proxy_endpoint_id", objectRaw["DBProxyEndpointId"])
	d.Set("db_proxy_endpoint_aliases", objectRaw["DbProxyEndpointAliases"])
	d.Set("read_only_instance_max_delay_time", objectRaw["ReadOnlyInstanceMaxDelayTime"])
	d.Set("causal_consist_read_timeout", objectRaw["CausalConsistReadTimeout"])
	d.Set("read_only_instance_distribution_type", objectRaw["ReadOnlyInstanceDistributionType"])
	d.Set("read_only_instance_weight", objectRaw["ReadOnlyInstanceWeight"])

	d.Set("db_instance_id", d.Id())

	return nil
}

func resourceAliCloudRdsDbProxyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyDBProxy"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ConfigDBProxyService"] = "Modify"
	if d.HasChange("persistent_connection_status") {
		update = true
		request["PersistentConnectionStatus"] = d.Get("persistent_connection_status")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
	action = "ModifyDBProxyInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("db_proxy_instance_type") {
		update = true
	}
	request["DBProxyInstanceType"] = d.Get("db_proxy_instance_type")
	if !d.IsNewResource() && d.HasChange("db_proxy_instance_num") {
		update = true
	}
	request["DBProxyInstanceNum"] = d.Get("db_proxy_instance_num")
	if v, ok := d.GetOk("effective_time"); ok {
		request["EffectiveTime"] = v
	}
	if v, ok := d.GetOk("effective_specific_time"); ok {
		request["EffectiveSpecificTime"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
		rdsServiceV2 := RdsServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Running"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, rdsServiceV2.RdsDbProxyStateRefreshFunc(d.Id(), "$.DBProxyInstanceStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ModifyDBProxyEndpoint"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("read_only_instance_distribution_type"); ok {
		request["ReadOnlyInstanceDistributionType"] = v
	}
	if v, ok := d.GetOk("causal_consist_read_timeout"); ok {
		request["CausalConsistReadTimeout"] = v
	}
	if d.HasChange("read_only_instance_max_delay_time") {
		update = true
		request["ReadOnlyInstanceMaxDelayTime"] = d.Get("read_only_instance_max_delay_time")
	}

	if v, ok := d.GetOk("vswitch_id"); ok {
		request["VSwitchId"] = v
	}
	if d.HasChange("db_proxy_endpoint_id") {
		update = true
		request["DBProxyEndpointId"] = d.Get("db_proxy_endpoint_id")
	}

	if v, ok := d.GetOk("db_endpoint_aliases"); ok {
		request["DbEndpointAliases"] = v
	}
	if v, ok := d.GetOk("db_endpoint_min_slave_count"); ok {
		request["DbEndpointMinSlaveCount"] = v
	}
	if v, ok := d.GetOk("effective_time"); ok {
		request["EffectiveTime"] = v
	}
	if v, ok := d.GetOk("effective_specific_time"); ok {
		request["EffectiveSpecificTime"] = v
	}
	if v, ok := d.GetOk("db_endpoint_read_write_mode"); ok {
		request["DbEndpointReadWriteMode"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VpcId"] = v
	}
	if v, ok := d.GetOk("config_db_proxy_features"); ok {
		request["ConfigDBProxyFeatures"] = v
	}
	if v, ok := d.GetOk("db_endpoint_operator"); ok {
		request["DbEndpointOperator"] = v
	}
	if v, ok := d.GetOk("read_only_instance_weight"); ok {
		request["ReadOnlyInstanceWeight"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
	action = "ModifyDbProxyInstanceSsl"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DbInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("db_proxy_ssl_enabled") {
		update = true
	}
	request["DbProxySslEnabled"] = d.Get("db_proxy_ssl_enabled")
	if d.HasChange("db_proxy_connect_string") {
		update = true
	}
	request["DbProxyConnectString"] = d.Get("db_proxy_connect_string")
	if d.HasChange("db_proxy_endpoint_id") {
		update = true
	}
	request["DbProxyEndpointId"] = d.Get("db_proxy_endpoint_id")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
	action = "ModifyDBProxyEndpointAddress"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["DBInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("db_proxy_new_connect_string") {
		update = true
		request["DBProxyNewConnectString"] = d.Get("db_proxy_new_connect_string")
	}

	if d.HasChange("db_proxy_endpoint_id") {
		update = true
	}
	request["DBProxyEndpointId"] = d.Get("db_proxy_endpoint_id")
	if d.HasChange("db_proxy_connect_string_net_type") {
		update = true
		request["DBProxyConnectStringNetType"] = d.Get("db_proxy_connect_string_net_type")
	}

	if d.HasChange("db_proxy_new_connect_string_port") {
		update = true
		request["DBProxyNewConnectStringPort"] = d.Get("db_proxy_new_connect_string_port")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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
	return resourceAliCloudRdsDbProxyRead(d, meta)
}

func resourceAliCloudRdsDbProxyDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "ModifyDBProxy"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["ConfigDBProxyService"] = "Shutdown"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Rds", "2014-08-15", action, query, request, true)
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

	rdsServiceV2 := RdsServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, rdsServiceV2.RdsDbProxyStateRefreshFunc(d.Id(), "$.DBProxyInstanceStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}

func convertRdsDbProxyDBProxyInstanceTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "3":
		return "common"
	case "2":
		return "exclusive"
	}
	return source
}
