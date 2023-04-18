// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
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

func resourceAliCloudCbwpCommonBandwidthPackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudCbwpCommonBandwidthPackageCreate,
		Read:   resourceAlicloudCbwpCommonBandwidthPackageRead,
		Update: resourceAlicloudCbwpCommonBandwidthPackageUpdate,
		Delete: resourceAlicloudCbwpCommonBandwidthPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bandwidth": {
				Type:     schema.TypeString,
				Required: true,
			},
			"business_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"common_bandwidth_package_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"common_bandwidth_package_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"bandwidth_package_name", "name"},
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
			"expired_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"force_delete": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"has_reservation_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_charge_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PayBy95", "PayByBandwidth", "PayByTraffic", "PayByDominantTraffic"}, false),
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BGP", "BGP_PRO", "ChinaTelecom", "ChinaUnicom", "ChinaMobile", "ChinaTelecom_L2", "ChinaUnicom_L2", "ChinaMobile_L2"}, false),
			},
			"payment_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allocation_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bandwidth_package_ip_relation_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ratio": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_protection_types": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"un_tag_all_tags": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bandwidth_package_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"common_bandwidth_package_name"},
				Deprecated:    "Field 'bandwidth_package_name' has been deprecated from provider version 1.201.3. New field 'common_bandwidth_package_name' instead.",
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"common_bandwidth_package_name"},
				Deprecated:    "Field 'name' has been deprecated from provider version 1.120.0. New field 'common_bandwidth_package_name' instead.",
			},
		},
	}
}

func resourceAlicloudCbwpCommonBandwidthPackageCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateCommonBandwidthPackage"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewCbwpClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("bandwidth"); ok {
		request["Bandwidth"] = v
	}

	if v, ok := d.GetOk("ratio"); ok {
		request["Ratio"] = v
	}

	if v, ok := d.GetOk("security_protection_types"); ok {
		localData := v
		securityProtectionTypesMaps := localData.([]interface{})
		request["SecurityProtectionTypes"] = securityProtectionTypesMaps
	}

	if v, ok := d.GetOk("isp"); ok {
		request["ISP"] = v
	}

	if v, ok := d.GetOk("internet_charge_type"); ok {
		request["InternetChargeType"] = v
	}

	if v, ok := d.GetOk("common_bandwidth_package_name"); ok {
		request["Name"] = v
	}
	if v, ok := d.GetOk("bandwidth_package_name"); ok {
		request["Name"] = v
	}
	if v, ok := d.GetOk("name"); ok {
		request["Name"] = v
	}

	if v, ok := d.GetOk("zone"); ok {
		request["Zone"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"BandwidthPackageOperation.conflict", "OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_common_bandwidth_package", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["BandwidthPackageId"]))

	cbwpServiceV2 := CbwpServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, cbwpServiceV2.CbwpCommonBandwidthPackageStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudCbwpCommonBandwidthPackageUpdate(d, meta)
}

func resourceAlicloudCbwpCommonBandwidthPackageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cbwpServiceV2 := CbwpServiceV2{client}

	object, err := cbwpServiceV2.DescribeCbwpCommonBandwidthPackage(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_common_bandwidth_package .DescribeCbwpCommonBandwidthPackage Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("bandwidth", object["bandwidth"])
	d.Set("business_status", object["business_status"])
	d.Set("common_bandwidth_package_id", object["common_bandwidth_package_id"])
	d.Set("common_bandwidth_package_name", object["common_bandwidth_package_name"])
	d.Set("create_time", object["create_time"])
	d.Set("deletion_protection", object["deletion_protection"])
	d.Set("description", object["description"])
	d.Set("expired_time", object["expired_time"])
	d.Set("has_reservation_data", object["has_reservation_data"])
	d.Set("internet_charge_type", object["internet_charge_type"])
	d.Set("isp", object["isp"])
	d.Set("payment_type", convertCbwpCommonBandwidthPackagePaymentTypeResponse(object["payment_type"].(string)))
	d.Set("public_ip_addresses", object["public_ip_addresses"])
	d.Set("ratio", object["ratio"])
	d.Set("reservation_active_time", object["reservation_active_time"])
	d.Set("reservation_bandwidth", object["reservation_bandwidth"])
	d.Set("reservation_internet_charge_type", object["reservation_internet_charge_type"])
	d.Set("reservation_order_type", object["reservation_order_type"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("security_protection_types", object["security_protection_types"])
	d.Set("service_managed", object["service_managed"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))

	d.Set("bandwidth_package_name", d.Get("common_bandwidth_package_name"))
	d.Set("name", d.Get("common_bandwidth_package_name"))
	return nil
}

func resourceAlicloudCbwpCommonBandwidthPackageUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyCommonBandwidthPackageAttribute"
	conn, err := client.NewCbwpClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["BandwidthPackageId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if !d.IsNewResource() && (d.HasChange("common_bandwidth_package_name") || d.HasChange("bandwidth_package_name") || d.HasChange("name")) {
		update = true
		if d.HasChange("common_bandwidth_package_name") {
			if v, ok := d.GetOk("common_bandwidth_package_name"); ok {
				request["Name"] = v
			}
		}
		if d.HasChange("bandwidth_package_name") {
			if v, ok := d.GetOk("bandwidth_package_name"); ok {
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
		d.SetPartial("description")
		d.SetPartial("common_bandwidth_package_name")
	}
	update = false
	action = "ModifyCommonBandwidthPackageSpec"
	conn, err = client.NewCbwpClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["BandwidthPackageId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("bandwidth") {
		update = true
		if v, ok := d.GetOk("bandwidth"); ok {
			request["Bandwidth"] = v
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"BandwidthPackageOperation.conflict", "OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable"}) || NeedRetry(err) {
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
		{
			cbwpServiceV2 := CbwpServiceV2{client}
			stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, cbwpServiceV2.CbwpCommonBandwidthPackageStateRefreshFunc(d.Id(), []string{}))
			if _, err := stateConf.WaitForState(); err != nil {
				return WrapErrorf(err, IdMsg, d.Id())
			}
		}
		d.SetPartial("bandwidth")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewCbwpClient()
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
	request["ResourceType"] = "bandwidthpackage"

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
	action = "DeletionProtection"
	conn, err = client.NewCbwpClient()
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
	request["Type"] = "CBWP"
	if d.HasChange("deletion_protection") {
		update = true
		request["ProtectionEnable"] = d.Get("deletion_protection")
	}
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

	update = false
	if d.HasChange("tags") {
		update = true
		cbwpServiceV2 := CbwpServiceV2{client}
		if err := cbwpServiceV2.SetResourceTags(d, "bandwidthpackage"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudCbwpCommonBandwidthPackageRead(d, meta)
}

func resourceAlicloudCbwpCommonBandwidthPackageDelete(d *schema.ResourceData, meta interface{}) error {
	// Pre paid instance can not be release.
	if d.Get("internet_charge_type").(string) == string(PayBy95) {
		log.Printf("[WARN] Cannot destroy CommonBandwidthPackage. Because internet_charge_type = PayBy95. Terraform will remove this resource from the state file, however resources may remain.")
		return nil
	}

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteCommonBandwidthPackage"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewCbwpClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["BandwidthPackageId"] = d.Id()
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("force_delete"); ok {
		request["Force"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"BandwidthPackageOperation.conflict", "OperationConflict", "LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "ServiceUnavailable"}) || NeedRetry(err) {
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

	cbwpServiceV2 := CbwpServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, cbwpServiceV2.CbwpCommonBandwidthPackageStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}

func convertCbwpCommonBandwidthPackagePaymentTypeRequest(source interface{}) interface{} {
	switch source {
	case "PayAsYouGo":
		return "PostPaid"
	case "Subscription":
		return "PrePaid"
	}
	return source
}

func convertCbwpCommonBandwidthPackagePaymentTypeResponse(source interface{}) interface{} {
	switch source {
	case "PostPaid":
		return "PayAsYouGo"
	case "PrePaid":
		return "Subscription"
	}
	return source
}
