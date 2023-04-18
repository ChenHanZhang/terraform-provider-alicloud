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

func resourceAliCloudEipSegmentAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudEipSegmentAddressCreate,
		Read:   resourceAlicloudEipSegmentAddressRead,
		Delete: resourceAlicloudEipSegmentAddressDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bandwidth": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"descritpion": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"eip_mask": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"internet_charge_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ip_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BGP", "BGP_PRO", "ChinaTelecom", "ChinaUnicom", "ChinaMobile", "ChinaTelecom_L2", "ChinaUnicom_L2", "ChinaMobile_L2"}, false),
			},
			"netmode": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_group_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("(.*)"), "The ID of the resource group."),
			},
			"segment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segment_address_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segment_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAlicloudEipSegmentAddressCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "AllocateEipSegmentAddress"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("isp"); ok {
		request["Isp"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("internet_charge_type"); ok {
		request["InternetChargeType"] = v
	}

	if v, ok := d.GetOk("eip_mask"); ok {
		request["EipMask"] = v
	}

	if v, ok := d.GetOk("bandwidth"); ok {
		request["Bandwidth"] = v
	}

	if v, ok := d.GetOk("netmode"); ok {
		request["Netmode"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationFailed.LastTokenProcessing", "IncorrectStatus.%s", "SystemBusy", "OperationConflict", "ServiceUnavailable"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_eip_segment_address", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["EipSegmentInstanceId"]))

	return resourceAlicloudEipSegmentAddressRead(d, meta)
}

func resourceAlicloudEipSegmentAddressRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	eipServiceV2 := EipServiceV2{client}

	object, err := eipServiceV2.DescribeEipSegmentAddress(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_eip_segment_address .DescribeEipSegmentAddress Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", object["create_time"])
	d.Set("descritpion", object["descritpion"])
	d.Set("ip_count", object["ip_count"])
	d.Set("segment", object["segment"])
	d.Set("segment_address_name", object["segment_address_name"])
	d.Set("segment_instance_id", object["segment_instance_id"])
	d.Set("status", object["status"])
	d.Set("zone", object["zone"])

	return nil
}

func resourceAlicloudEipSegmentAddressDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ReleaseEipSegmentAddress"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewEipClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["SegmentInstanceId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationFailed.LastTokenProcessing", "LastTokenProcessing", "IncorrectStatus.%s", "IncorrectEipStatus", "OperationFailed.EipStatusInvalid", "OperationConflict", "TaskConflict.AssociateGlobalAccelerationInstance", "SystemBusy", "ServiceUnavailable"}) || NeedRetry(err) {
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
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, eipServiceV2.EipSegmentAddressStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
