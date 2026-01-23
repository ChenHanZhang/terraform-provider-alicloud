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

func resourceAliCloudAmqpInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAmqpInstanceCreate,
		Read:   resourceAliCloudAmqpInstanceRead,
		Update: resourceAliCloudAmqpInstanceUpdate,
		Delete: resourceAliCloudAmqpInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(200 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_renew": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"edition": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"professional", "enterprise", "vip", "serverless"}, false),
			},
			"max_connections": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_eip_tps": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_tps": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modify_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Upgrade", "Downgrade"}, false),
			},
			"payment_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"Subscription", "PayAsYouGo"}, false),
			},
			"period": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntInSlice([]int{0, 1, 2, 3, 6, 12, 24}),
			},
			"period_cycle": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"Month", "Year"}, false),
			},
			"provisioned_capacity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"queue_capacity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"renewal_duration": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: IntInSlice([]int{0, 1, 2, 3, 6, 12}),
			},
			"renewal_duration_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"Month", "Year"}, false),
			},
			"renewal_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"AutoRenewal", "ManualRenewal", "NotRenewal"}, false),
			},
			"serverless_charge_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"support_eip": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"support_tracing": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tracing_storage_time": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: IntInSlice([]int{0, 3, 7, 15}),
			},
		},
	}
}

func resourceAliCloudAmqpInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("auto_renew"); ok {
		request["AutoRenew"] = v
	}
	if v, ok := d.GetOk("max_tps"); ok {
		request["MaxPrivateTps"] = v
	}
	request["PaymentType"] = d.Get("payment_type")
	if v, ok := d.GetOk("period_cycle"); ok {
		request["PeriodCycle"] = v
	}
	if v, ok := d.GetOkExists("tracing_storage_time"); ok {
		request["TracingStorageTime"] = v
	}
	if v, ok := d.GetOk("renewal_duration_unit"); ok {
		request["RenewalDurationUnit"] = v
	}
	if v, ok := d.GetOk("instance_name"); ok {
		request["InstanceName"] = v
	}
	if v, ok := d.GetOkExists("renewal_duration"); ok {
		request["AutoRenewPeriod"] = v
	}
	if v, ok := d.GetOkExists("support_tracing"); ok {
		request["SupportTracing"] = v
	}
	if v, ok := d.GetOkExists("support_eip"); ok {
		request["SupportEip"] = v
	}
	if v, ok := d.GetOk("max_eip_tps"); ok {
		request["MaxEipTps"] = v
	}
	if v, ok := d.GetOk("serverless_charge_type"); ok {
		request["ServerlessChargeType"] = v
	}
	if v, ok := d.GetOk("edition"); ok {
		request["Edition"] = v
	}
	if v, ok := d.GetOk("queue_capacity"); ok {
		request["QueueCapacity"] = v
	}
	if v, ok := d.GetOk("instance_type"); ok {
		request["InstanceType"] = v
	}
	if v, ok := d.GetOk("renewal_status"); ok {
		request["RenewStatus"] = v
	}
	if v, ok := d.GetOkExists("provisioned_capacity"); ok {
		request["ProvisionedCapacity"] = v
	}
	if v, ok := d.GetOkExists("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOk("storage_size"); ok {
		request["StorageSize"] = v
	}
	if v, ok := d.GetOkExists("max_connections"); ok {
		request["MaxConnections"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("amqp-open", "2019-12-12", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_amqp_instance", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["Data"]))

	amqpServiceV2 := AmqpServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"SERVING"}, d.Timeout(schema.TimeoutCreate), 60*time.Second, amqpServiceV2.AmqpInstanceStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudAmqpInstanceRead(d, meta)
}

func resourceAliCloudAmqpInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	amqpServiceV2 := AmqpServiceV2{client}

	objectRaw, err := amqpServiceV2.DescribeAmqpInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_amqp_instance DescribeAmqpInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["OrderCreateTime"])
	d.Set("edition", objectRaw["Edition"])
	d.Set("instance_name", objectRaw["InstanceName"])
	d.Set("instance_type", convertAmqpInstanceDataInstanceTypeResponse(objectRaw["InstanceType"]))
	d.Set("max_connections", objectRaw["MaxConnections"])
	d.Set("max_eip_tps", objectRaw["MaxEipTps"])
	d.Set("max_tps", objectRaw["MaxTps"])
	d.Set("provisioned_capacity", objectRaw["ProvisionedCapacity"])
	d.Set("queue_capacity", objectRaw["MaxQueue"])
	d.Set("status", objectRaw["Status"])
	d.Set("storage_size", objectRaw["StorageSize"])
	d.Set("support_eip", objectRaw["SupportEIP"])
	d.Set("support_tracing", objectRaw["SupportTracing"])
	d.Set("tracing_storage_time", objectRaw["TracingStorageTime"])

	objectRaw, err = amqpServiceV2.DescribeInstanceQueryAvailableInstances(d)
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("create_time", formatInt(objectRaw["CreateTime"]))
	d.Set("payment_type", objectRaw["SubscriptionType"])
	d.Set("renewal_duration", objectRaw["RenewalDuration"])
	d.Set("renewal_duration_unit", convertAmqpInstanceDataInstanceListRenewalDurationUnitResponse(objectRaw["RenewalDurationUnit"]))
	d.Set("renewal_status", objectRaw["RenewStatus"])

	return nil
}

func resourceAliCloudAmqpInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "SetRenewal"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceIDs"] = d.Id()

	request["ProductCode"] = "ons"
	if d.HasChange("payment_type") {
		update = true
	}
	request["SubscriptionType"] = d.Get("payment_type")
	if d.HasChange("renewal_duration_unit") {
		update = true
		request["RenewalPeriodUnit"] = convertAmqpInstanceRenewalPeriodUnitRequest(d.Get("renewal_duration_unit").(string))
	}

	if d.HasChange("renewal_status") {
		update = true
	}
	request["RenewalStatus"] = d.Get("renewal_status")
	if d.HasChange("renewal_duration") {
		update = true
		request["RenewalPeriod"] = d.Get("renewal_duration")
	}

	var endpoint string
	request["ProductCode"] = ""
	request["ProductType"] = ""
	if client.IsInternationalAccount() {
		request["ProductType"] = ""
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPostWithEndpoint("BssOpenApi", "2017-12-14", action, query, request, true, endpoint)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				if !client.IsInternationalAccount() && IsExpectedErrors(err, []string{""}) {
					request["ProductCode"] = ""
					request["ProductType"] = ""
					endpoint = connectivity.BssOpenAPIEndpointInternational
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
	action = "UpdateInstanceName"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if d.HasChange("instance_name") {
		update = true
	}
	request["InstanceName"] = d.Get("instance_name")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("amqp-open", "2019-12-12", action, query, request, true)
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
	action = "UpdateInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("max_tps") {
		update = true
		request["MaxPrivateTps"] = d.Get("max_tps")
	}

	if d.HasChange("tracing_storage_time") {
		update = true
		request["TracingStorageTime"] = d.Get("tracing_storage_time")
	}

	request["ModifyType"] = convertAmqpInstanceModifyTypeRequest(d.Get("modify_type").(string))
	if d.HasChange("support_tracing") {
		update = true
		request["SupportTracing"] = d.Get("support_tracing")
	}

	if d.HasChange("support_eip") {
		update = true
		request["SupportEip"] = d.Get("support_eip")
	}

	if d.HasChange("max_eip_tps") {
		update = true
		request["MaxEipTps"] = d.Get("max_eip_tps")
	}

	if v, ok := d.GetOk("serverless_charge_type"); ok {
		request["ServerlessChargeType"] = v
	}
	if d.HasChange("edition") {
		update = true
		request["Edition"] = d.Get("edition")
	}

	if d.HasChange("queue_capacity") {
		update = true
		request["QueueCapacity"] = d.Get("queue_capacity")
	}

	if d.HasChange("instance_type") {
		update = true
		request["InstanceType"] = d.Get("instance_type")
	}

	if d.HasChange("provisioned_capacity") {
		update = true
		request["ProvisionedCapacity"] = d.Get("provisioned_capacity")
	}

	if d.HasChange("storage_size") {
		update = true
		request["StorageSize"] = d.Get("storage_size")
	}

	if d.HasChange("max_connections") {
		update = true
		request["MaxConnections"] = d.Get("max_connections")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("amqp-open", "2019-12-12", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"InstanceUpgradeOrDownGradeTopicMigrating"}) || NeedRetry(err) {
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
		amqpServiceV2 := AmqpServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"SERVING"}, d.Timeout(schema.TimeoutUpdate), 60*time.Second, amqpServiceV2.AmqpInstanceStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	d.Partial(false)
	return resourceAliCloudAmqpInstanceRead(d, meta)
}

func resourceAliCloudAmqpInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	enableDelete := false
	if v, ok := d.GetOkExists("payment_type"); ok {
		if InArray(fmt.Sprint(v), []string{"Subscription"}) {
			enableDelete = true
		}
	}
	if enableDelete {
		action := "RefundInstance"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["InstanceId"] = d.Id()

		request["ClientToken"] = buildClientToken(action)

		request["ProductCode"] = "ons"
		request["ProductType"] = "ons_onsproxy_pre"
		request["ImmediatelyRelease"] = "1"
		var endpoint string
		request["ProductCode"] = ""
		request["ProductType"] = ""
		if client.IsInternationalAccount() {
			request["ProductType"] = ""
		}
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
			response, err = client.RpcPostWithEndpoint("BssOpenApi", "2017-12-14", action, query, request, true, endpoint)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				if !client.IsInternationalAccount() && IsExpectedErrors(err, []string{""}) {
					request["ProductCode"] = ""
					request["ProductType"] = ""
					endpoint = connectivity.BssOpenAPIEndpointInternational
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

		amqpServiceV2 := AmqpServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"RELEASED"}, d.Timeout(schema.TimeoutDelete), 3*time.Minute, amqpServiceV2.AmqpInstanceStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}

	}
	return nil
}

func convertAmqpInstanceSupportEipRequest(source interface{}) interface{} {
	switch source {
	case false:
		return "eip_false"
	case true:
		return "eip_true"
	}
	return ""
}
func convertAmqpInstanceInstanceTypeResponse(source interface{}) interface{} {
	switch source {
	case "PROFESSIONAL":
		return "professional"
	case "ENTERPRISE":
		return "enterprise"
	case "VIP":
		return "vip"
	}
	return source
}
func convertAmqpInstanceRenewalDurationUnitResponse(source interface{}) interface{} {
	switch source {
	case "M":
		return "Month"
	case "Y":
		return "Year"
	}
	return source
}
func convertAmqpInstanceRenewalDurationUnitRequest(source interface{}) interface{} {
	switch source {
	case "Month":
		return "M"
	case "Year":
		return "Y"
	}
	return source
}

func convertAmqpInstanceDataInstanceTypeResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "PROFESSIONAL":
		return "professional"
	case "VIP":
		return "vip"
	case "ENTERPRISE":
		return "enterprise"
	}
	return source
}
func convertAmqpInstanceDataInstanceListRenewalDurationUnitResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "M":
		return "Month"
	case "Y":
		return "Year"
	}
	return source
}
func convertAmqpInstanceRenewalPeriodUnitRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Month":
		return "M"
	case "Year":
		return "Y"
	}
	return source
}
func convertAmqpInstanceModifyTypeRequest(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "Downgrade":
		return "DOWNGRADE"
	case "Upgrade":
		return "UPGRADE"
	}
	return source
}
