// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudEipAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudEipAddressCreate,
		Read:   resourceAlicloudEipAddressRead,
		Update: resourceAlicloudEipAddressUpdate,
		Delete: resourceAlicloudEipAddressDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"activity_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"address_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
			},
			"allocation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_pay": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"bandwidth": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_package_bandwidth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bandwidth_package_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"bandwidth_package_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"business_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deletion_protection": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"eip_bandwidth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expired_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"has_reservation_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hd_monitor_log_project": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"log_project"},
			},
			"hd_monitor_log_store": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"log_store"},
			},
			"hd_monitor_status": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"high_definition_monitor_log_status"},
				ForceNew:      true,
				ValidateFunc:  validation.StringInSlice([]string{"false", "true"}, false),
			},
			"instance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_region_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PayByBandwidth", "PayByTraffic", "PayByDominantTraffic"}, false),
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BGP", "BGP_PRO", "ChinaTelecom", "ChinaUnicom", "ChinaMobile", "ChinaTelecom_L2", "ChinaUnicom_L2", "ChinaMobile_L2"}, false),
			},
			"netmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"operation_locks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lock_reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"payment_type": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"instance_charge_type"},
				ForceNew:      true,
				ValidateFunc:  validation.StringInSlice([]string{"Subscription", "PayAsYouGo"}, false),
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"pricing_cycle": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"public_ip_address_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"reservation_active_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_bandwidth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_internet_charge_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_order_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("(.*)"), "The ID of the resource group."),
			},
			"second_limited": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"security_protection_types": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"segment_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_managed": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"vpc_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"address_name"},
				Deprecated:    "Field 'name' has been deprecated from provider version 1.126.0. New field 'address_name' instead.",
			},
			"instance_charge_type": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"payment_type"},
				Deprecated:    "Field 'instance_charge_type' has been deprecated from provider version 1.126.0. New field 'payment_type' instead.",
			},
			"high_definition_monitor_log_status": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"hd_monitor_status"},
				Deprecated:    "Field 'high_definition_monitor_log_status' has been deprecated from provider version 1.204.0. New field 'hd_monitor_status' instead.",
			},
			"log_project": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"hd_monitor_log_project"},
				Deprecated:    "Field 'log_project' has been deprecated from provider version 1.204.0. New field 'hd_monitor_log_project' instead.",
			},
			"log_store": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"hd_monitor_log_store"},
				Deprecated:    "Field 'log_store' has been deprecated from provider version 1.204.0. New field 'hd_monitor_log_store' instead.",
			},
		},
	}
}

func resourceAlicloudEipAddressCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "AllocateEipAddress"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("bandwidth"); ok {
		request["Bandwidth"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}

	if v, ok := d.GetOk("address_name"); ok {
		request["Name"] = v
	}
	if v, ok := d.GetOk("name"); ok {
		request["Name"] = v
	}

	if v, ok := d.GetOk("netmode"); ok {
		request["Netmode"] = v
	}

	if v, ok := d.GetOk("security_protection_types"); ok {
		localData := v
		securityProtectionTypesMaps := localData.([]interface{})
		request["SecurityProtectionTypes"] = securityProtectionTypesMaps
	}

	if v, ok := d.GetOk("isp"); ok {
		request["ISP"] = v
	}

	if v, ok := d.GetOk("period"); ok {
		request["Period"] = v
	}

	if v, ok := d.GetOk("activity_id"); ok {
		request["ActivityId"] = v
	}

	if v, ok := d.GetOkExists("auto_pay"); ok {
		request["AutoPay"] = v
	}

	if v, ok := d.GetOk("pricing_cycle"); ok {
		request["PricingCycle"] = v
	}

	if v, ok := d.GetOk("public_ip_address_pool_id"); ok {
		request["PublicIpAddressPoolId"] = v
	}

	if v, ok := d.GetOk("payment_type"); ok {
		request["InstanceChargeType"] = convertEipAddressPaymentTypeRequest(v.(string))
	}
	if v, ok := d.GetOk("instance_charge_type"); ok {
		request["InstanceChargeType"] = convertEipAddressPaymentTypeRequest(v.(string))
	}

	if v, ok := d.GetOk("internet_charge_type"); ok {
		request["InternetChargeType"] = v
	}

	if v, ok := d.GetOk("instance_charge_type"); ok {
		request["InstanceChargeType"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable", "FrequentPurchase.EIP"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eip_address", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["AllocationId"]))

	eipServiceV2 := EipServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, eipServiceV2.EipAddressStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudEipAddressUpdate(d, meta)
}

func resourceAlicloudEipAddressRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eipServiceV2 := EipServiceV2{client}

	object, err := eipServiceV2.DescribeEipAddress(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eip_address .DescribeEipAddress Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("address_name", object["address_name"])
	d.Set("allocation_id", object["allocation_id"])
	d.Set("bandwidth", object["bandwidth"])
	d.Set("bandwidth_package_bandwidth", object["bandwidth_package_bandwidth"])
	d.Set("bandwidth_package_id", object["bandwidth_package_id"])
	d.Set("bandwidth_package_type", object["bandwidth_package_type"])
	d.Set("business_status", object["business_status"])
	d.Set("create_time", object["create_time"])
	d.Set("deletion_protection", object["deletion_protection"])
	d.Set("description", object["description"])
	d.Set("eip_bandwidth", object["eip_bandwidth"])
	d.Set("expired_time", object["expired_time"])
	d.Set("has_reservation_data", object["has_reservation_data"])
	d.Set("hd_monitor_log_project", object["hd_monitor_log_project"])
	d.Set("hd_monitor_log_store", object["hd_monitor_log_store"])
	d.Set("hd_monitor_status", object["hd_monitor_status"])
	d.Set("instance_id", object["instance_id"])
	d.Set("instance_region_id", object["instance_region_id"])
	d.Set("internet_charge_type", object["internet_charge_type"])
	d.Set("ip_address", object["ip_address"])
	d.Set("isp", object["isp"])
	d.Set("netmode", object["netmode"])
	d.Set("operation_locks", object["operation_locks"])
	d.Set("payment_type", convertEipAddressPaymentTypeResponse(object["payment_type"].(string)))
	d.Set("public_ip_address_pool_id", object["public_ip_address_pool_id"])
	d.Set("reservation_active_time", object["reservation_active_time"])
	d.Set("reservation_bandwidth", object["reservation_bandwidth"])
	d.Set("reservation_internet_charge_type", object["reservation_internet_charge_type"])
	d.Set("reservation_order_type", object["reservation_order_type"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("second_limited", object["second_limited"])
	d.Set("security_protection_types", object["security_protection_types"])
	d.Set("segment_instance_id", object["segment_instance_id"])
	d.Set("service_managed", object["service_managed"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))
	d.Set("vpc_id", object["vpc_id"])
	d.Set("zone", object["zone"])

	d.Set("name", d.Get("address_name"))
	d.Set("instance_charge_type", d.Get("payment_type"))
	d.Set("high_definition_monitor_log_status", d.Get("hd_monitor_status"))
	d.Set("log_project", d.Get("hd_monitor_log_project"))
	d.Set("log_store", d.Get("hd_monitor_log_store"))
	d.Set("instance_charge_type", object["payment_type"])
	return nil
}

func resourceAlicloudEipAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyEipAddressAttribute"
	conn, err := client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["AllocationId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("bandwidth") {
		update = true
		if v, ok := d.GetOk("bandwidth"); ok {
			request["Bandwidth"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if !d.IsNewResource() && (d.HasChange("address_name") || d.HasChange("name")) {
		update = true
		if d.HasChange("address_name") {
			if v, ok := d.GetOk("address_name"); ok {
				request["Name"] = v
			}
		}
		if d.HasChange("name") {
			if v, ok := d.GetOk("name"); ok {
				request["Name"] = v
			}
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable", "IncorrectEipStatus", "IncorrectStatus.ResourceStatus"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("bandwidth")
		d.SetPartial("description")
		d.SetPartial("address_name")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["NewResourceGroupId"] = v
		}
	}
	request["ResourceType"] = "EIP"

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("resource_group_id")
	}
	update = false
	action = "SetHighDefinitionMonitorLogStatus"
	conn, err = client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId

	if d.HasChange("hd_monitor_log_project") || d.HasChange("log_project") {
		update = true
		if d.HasChange("hd_monitor_log_project") {
			if v, ok := d.GetOk("hd_monitor_log_project"); ok {
				request["LogProject"] = v
			}
		}
		if d.HasChange("log_project") {
			if v, ok := d.GetOk("log_project"); ok {
				request["LogProject"] = v
			}
		}
	}
	if d.HasChange("hd_monitor_log_store") || d.HasChange("log_store") {
		update = true
		if d.HasChange("hd_monitor_log_store") {
			if v, ok := d.GetOk("hd_monitor_log_store"); ok {
				request["LogStore"] = v
			}
		}
		if d.HasChange("log_store") {
			if v, ok := d.GetOk("log_store"); ok {
				request["LogStore"] = v
			}
		}
	}
	if d.HasChange("hd_monitor_status") || d.HasChange("high_definition_monitor_log_status") {
		update = true
		if d.HasChange("hd_monitor_status") {
			if v, ok := d.GetOk("hd_monitor_status"); ok {
				request["Status"] = v
			}
		}
		if d.HasChange("high_definition_monitor_log_status") {
			if v, ok := d.GetOk("high_definition_monitor_log_status"); ok {
				request["Status"] = v
			}
		}
	}
	if d.HasChange("instance_type") {
		update = true
		if v, ok := d.GetOk("instance_type"); ok {
			request["InstanceType"] = v
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable", "IncorrectEipStatus", "IncorrectInstanceStatus", "InvalidBindingStatus", "IncorrectStatus.NatGateway", "InvalidStatus.EcsStatusNotSupport", "InvalidStatus.InstanceHasBandWidth", "InvalidStatus.EniStatusNotSupport", "TaskConflict"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("hd_monitor_log_project")
		d.SetPartial("hd_monitor_log_store")
		d.SetPartial("instance_type")
	}
	update = false
	action = "DeletionProtection"
	conn, err = client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["InstanceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("deletion_protection") {
		update = true
		if v, ok := d.GetOkExists("deletion_protection"); ok {
			request["ProtectionEnable"] = v
		}
	}
	request["Type"] = "EIP"

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("deletion_protection")
	}

	if d.HasChange("second_limited") {
		client := meta.(*connectivity.AliyunClient)
		eipServiceV2 := EipServiceV2{client}
		object, err := eipServiceV2.DescribeEipAddress(d.Id())
		if err != nil {
			return WrapError(err)
		}

		target := d.Get("second_limited").(bool)
		if object["second_limited"].(bool) != target {
			if target == true {
				action = "CancelCommonBandwidthPackageIpBandwidth"
				conn, err = client.NewEipClient()
				if err != nil {
					return WrapError(err)
				}
				request = make(map[string]interface{})

				request["EipId"] = d.Id()
				request["RegionId"] = client.RegionId

				if v, ok := d.GetOk("bandwidth_package_id"); ok {
					request["BandwidthPackageId"] = v
				}

				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
					if err != nil {
						if IsExpectedErrors(err, []string{"OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable"}) || NeedRetry(err) {
							wait()
							return resource.RetryableError(err)
						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}

			}
		}
	}

	update = false
	if d.HasChange("tags") {
		update = true
		eipServiceV2 := EipServiceV2{client}
		if err := eipServiceV2.SetResourceTags(d, "EIP"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudEipAddressRead(d, meta)
}

func resourceAlicloudEipAddressDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ReleaseEipAddress"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["AllocationId"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable", "IncorrectEipStatus", "TaskConflict.AssociateGlobalAccelerationInstance"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	eipServiceV2 := EipServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, eipServiceV2.EipAddressStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}

func convertEipAddressPaymentTypeRequest(source interface{}) interface{} {
	switch source {
	case "PayAsYouGo":
		return "PostPaid"
	case "Subscription":
		return "PrePaid"
	}
	return source
}

func convertEipAddressPaymentTypeResponse(source interface{}) interface{} {
	switch source {
	case "PostPaid":
		return "PayAsYouGo"
	case "PrePaid":
		return "Subscription"
	}
	return source
}
