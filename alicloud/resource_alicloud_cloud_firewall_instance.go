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

func resourceAliCloudCloudFirewallInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudFirewallInstanceCreate,
		Read:   resourceAliCloudCloudFirewallInstanceRead,
		Update: resourceAliCloudCloudFirewallInstanceUpdate,
		Delete: resourceAliCloudCloudFirewallInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_asset_protection": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cfw_log": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
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
				ValidateFunc: StringInSlice([]string{"PayAsYouGo", "Subscription"}, false),
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"product_code": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"cfw"}, false),
			},
			"product_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"cfw_elasticity_public_cn", "cfw_elasticity_public_intl", "cfw_sub_public_cn", "cfw_sub_public_intl"}, false),
			},
			"release_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"renewal_duration": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"renewal_duration_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"renewal_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"spec": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"payg_version", "premium_version", "enterprise_version", "ultimate_version"}, false),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudCloudFirewallInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	parameterMapList := make([]map[string]interface{}, 0)
	if v, ok := d.GetOk("sdl"); ok {
		parameterMapList = append(parameterMapList, map[string]interface{}{
			"Code":  "cfw_ndlp_enable",
			"Value": fmt.Sprint(v),
		})
	}
	if v, ok := d.GetOk("cfw_log"); ok {
		parameterMapList = append(parameterMapList, map[string]interface{}{
			"Code":  "CfwLog",
			"Value": fmt.Sprint(v),
		})
	}
	if v, ok := d.GetOk("spec"); ok {
		parameterMapList = append(parameterMapList, map[string]interface{}{
			"Code":  "cfw_spec",
			"Value": v,
		})
	}
	if v, ok := d.GetOk("auto_asset_protection"); ok {
		parameterMapList = append(parameterMapList, map[string]interface{}{
			"Code":  "AutoAssetProtection",
			"Value": v,
		})
	}
	request["Parameter"] = parameterMapList

	request["ProductCode"] = d.Get("product_code")
	request["SubscriptionType"] = d.Get("payment_type")
	request["ProductType"] = d.Get("product_type")
	if v, ok := d.GetOk("renewal_status"); ok {
		request["RenewalStatus"] = v
	}
	if v, ok := d.GetOkExists("renewal_duration"); ok {
		request["RenewPeriod"] = v
	}
	if v, ok := d.GetOkExists("period"); ok {
		request["Period"] = v
	}
	var endpoint string
	request["ProductCode"] = ""
	request["ProductType"] = ""
	if client.IsInternationalAccount() {
		request["ProductType"] = ""
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cloud_firewall_instance", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Data.InstanceId", response)
	d.SetId(fmt.Sprint(id))

	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Normal"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, cloudFirewallServiceV2.CloudFirewallInstanceStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudCloudFirewallInstanceUpdate(d, meta)
}

func resourceAliCloudCloudFirewallInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}

	objectRaw, err := cloudFirewallServiceV2.DescribeCloudFirewallInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cloud_firewall_instance DescribeCloudFirewallInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("end_time", objectRaw["EndTime"])
	d.Set("payment_type", objectRaw["SubscriptionType"])
	d.Set("product_code", objectRaw["ProductCode"])
	d.Set("product_type", objectRaw["ProductType"])
	d.Set("release_time", objectRaw["ReleaseTime"])
	d.Set("renewal_duration", objectRaw["RenewalDuration"])
	d.Set("renewal_duration_unit", objectRaw["RenewalDurationUnit"])
	d.Set("renewal_status", objectRaw["RenewStatus"])
	d.Set("status", objectRaw["Status"])

	objectRaw, err = cloudFirewallServiceV2.DescribeInstanceDescribeUserBuyVersion(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("cfw_log", objectRaw["LogStatus"])
	d.Set("sdl", objectRaw["Sdl"])
	d.Set("spec", convertCloudFirewallInstanceVersionResponse(objectRaw["Version"]))
	d.Set("status", objectRaw["InstanceStatus"])
	d.Set("user_status", objectRaw["UserStatus"])

	objectRaw, err = cloudFirewallServiceV2.DescribeInstanceDescribeAssetStatistic(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("auto_asset_protection", objectRaw["AutoResourceEnable"])

	return nil
}

func resourceAliCloudCloudFirewallInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "ModifyCfwInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	if d.HasChange("sdl") {
		update = true
	}
	updateList := make(map[string]interface{})

	if v := d.Get("sdl"); !IsNil(v) || d.HasChange("sdl") {
		1 := make(map[string]interface{})
		if v, ok := d.GetOkExists("sdl"); ok {
			1["Value"] = v
		}

		if len(1) > 0 {
			updateList["1"] = 1
		}

		request["UpdateList"] = updateList
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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
	action = "ModifyInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("product_code") {
		update = true
	}
	request["ProductCode"] = d.Get("product_code")
	if !d.IsNewResource() && d.HasChange("payment_type") {
		update = true
	}
	request["SubscriptionType"] = d.Get("payment_type")
	if !d.IsNewResource() && d.HasChange("product_type") {
		update = true
	}
	request["ProductType"] = d.Get("product_type")
	request["ModifyType"] = d.Get("modify_type")
	if d.HasChange("cfw_log") {
		update = true
	}
	parameter := make(map[string]interface{})

	if v := d.Get("cfw_log"); !IsNil(v) || d.HasChange("cfw_log") {
		1 := make(map[string]interface{})
		if v, ok := d.GetOkExists("cfw_log"); ok {
			1["Value"] = v
		}

		if len(1) > 0 {
			parameter["1"] = 1
		}

		request["Parameter"] = parameter
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
		cloudFirewallServiceV2 := CloudFirewallServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("cfw_log"))}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, cloudFirewallServiceV2.CloudFirewallInstanceStateRefreshFuncWithApi(d.Id(), "LogStatus", []string{}, cloudFirewallServiceV2.DescribeInstanceDescribeUserBuyVersion))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "SetRenewal"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceIDs"] = d.Id()

	if !d.IsNewResource() && d.HasChange("product_code") {
		update = true
	}
	request["ProductCode"] = d.Get("product_code")
	if !d.IsNewResource() && d.HasChange("payment_type") {
		update = true
	}
	request["SubscriptionType"] = d.Get("payment_type")
	if !d.IsNewResource() && d.HasChange("product_type") {
		update = true
	}
	request["ProductType"] = d.Get("product_type")
	if d.HasChange("renewal_duration_unit") {
		update = true
	}
	if v, ok := d.GetOk("renewal_duration_unit"); ok || d.HasChange("renewal_duration_unit") {
		request["RenewalPeriodUnit"] = v
	}
	if !d.IsNewResource() && d.HasChange("renewal_status") {
		update = true
	}
	request["RenewalStatus"] = d.Get("renewal_status")
	if !d.IsNewResource() && d.HasChange("renewal_duration") {
		update = true
	}
	if v, ok := d.GetOkExists("renewal_duration"); ok || d.HasChange("renewal_duration") {
		request["RenewalPeriod"] = v
	}
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
	action = "SetAutoProtectNewAssets"
	request = make(map[string]interface{})
	query = make(map[string]interface{})

	if !d.IsNewResource() && d.HasChange("auto_asset_protection") {
		update = true
	}
	request["AutoProtect"] = d.Get("auto_asset_protection")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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
	return resourceAliCloudCloudFirewallInstanceRead(d, meta)
}

func resourceAliCloudCloudFirewallInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	enableDelete := false
	if v, ok := d.GetOkExists("payment_type"); ok {
		if InArray(fmt.Sprint(v), []string{"PayAsYouGo"}) {
			enableDelete = true
		}
	}
	if enableDelete {
		action := "ReleasePostInstance"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		request["InstanceId"] = d.Id()

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
			response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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

		cloudFirewallServiceV2 := CloudFirewallServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"#CHECKSET"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, cloudFirewallServiceV2.CloudFirewallInstanceStateRefreshFunc(d.Id(), "#Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}

	}
	return nil
}

func convertCloudFirewallInstanceVersionResponse(source interface{}) interface{} {
	source = fmt.Sprint(source)
	switch source {
	case "2":
		return "premium_version"
	case "3":
		return "enterprise_version"
	case "4":
		return "ultimate_version"
	case "10":
		return "payg_version"
	}
	return source
}
