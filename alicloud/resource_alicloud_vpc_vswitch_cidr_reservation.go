// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudVpcVswitchCidrReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcVswitchCidrReservationCreate,
		Read:   resourceAlicloudVpcVswitchCidrReservationRead,
		Update: resourceAlicloudVpcVswitchCidrReservationUpdate,
		Delete: resourceAlicloudVpcVswitchCidrReservationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cdir_reservation_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cidr_reservation_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"cidr_reservation_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cidr_reservation_mask": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vswitch_cidr_reservation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vswitch_cidr_reservation_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAlicloudVpcVswitchCidrReservationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateVSwitchCidrReservation"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["VSwitchId"] = d.Get("vswitch_id")
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("ip_version"); ok {
		request["IpVersion"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("cdir_reservation_type"); ok {
		request["VSwitchCidrReservationType"] = v
	}

	if v, ok := d.GetOk("cidr_reservation_description"); ok {
		request["VSwitchCidrReservationDescription"] = v
	}

	if v, ok := d.GetOk("cidr_reservation_cidr"); ok {
		request["VSwitchCidrReservationCidr"] = v
	}

	if v, ok := d.GetOk("vswitch_cidr_reservation_name"); ok {
		request["VSwitchCidrReservationName"] = v
	}

	if v, ok := d.GetOk("cidr_reservation_mask"); ok {
		request["VSwitchCidrReservationMask"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.Vpc", "OperationConflict", "IncorrectStatus.%s", "ServiceUnavailable", "SystemBusy", "LastTokenProcessing"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_vswitch_cidr_reservation", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["VSwitchId"], response["VSwitchCidrReservationId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Assigned"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcVswitchCidrReservationStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcVswitchCidrReservationUpdate(d, meta)
}

func resourceAlicloudVpcVswitchCidrReservationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcVswitchCidrReservation(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_vswitch_cidr_reservation .DescribeVpcVswitchCidrReservation Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cdir_reservation_type", object["cdir_reservation_type"])
	d.Set("cidr_reservation_cidr", object["cidr_reservation_cidr"])
	d.Set("cidr_reservation_description", object["cidr_reservation_description"])
	d.Set("create_time", object["create_time"])
	d.Set("ip_version", object["ip_version"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("status", object["status"])
	d.Set("vswitch_cidr_reservation_id", object["vswitch_cidr_reservation_id"])
	d.Set("vswitch_cidr_reservation_name", object["vswitch_cidr_reservation_name"])
	d.Set("vswitch_id", object["vswitch_id"])
	d.Set("vpc_instance_id", object["vpc_instance_id"])

	return nil
}

func resourceAlicloudVpcVswitchCidrReservationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyVSwitchCidrReservationAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", d.Id(), 2, len(parts)))
	}

	request["VSwitchCidrReservationId"] = parts[1]
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("cidr_reservation_description") {
		update = true
		if v, ok := d.GetOk("cidr_reservation_description"); ok {
			request["VSwitchCidrReservationDescription"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("vswitch_cidr_reservation_name") {
		update = true
		if v, ok := d.GetOk("vswitch_cidr_reservation_name"); ok {
			request["VSwitchCidrReservationName"] = v
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
		d.SetPartial("cidr_reservation_description")
		d.SetPartial("vswitch_cidr_reservation_name")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	parts = strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", d.Id(), 2, len(parts)))
	}

	request["ResourceId"] = parts[1]
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["NewResourceGroupId"] = v
		}
	}
	if d.HasChange("resource_type") {
		update = true
		if v, ok := d.GetOk("resource_type"); ok {
			request["ResourceType"] = v
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
		d.SetPartial("resource_group_id")
	}

	d.Partial(false)
	return resourceAlicloudVpcVswitchCidrReservationRead(d, meta)
}

func resourceAlicloudVpcVswitchCidrReservationDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteVSwitchCidrReservation"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		err = WrapError(fmt.Errorf("invalid Resource Id %s. Expected parts' length %d, got %d", d.Id(), 2, len(parts)))
	}

	request["VSwitchCidrReservationId"] = parts[1]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.Vpc", "OperationConflict", "IncorrectStatus.%s", "ServiceUnavailable", "SystemBusy", "LastTokenProcessing"}) || NeedRetry(err) {
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

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcVswitchCidrReservationStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
