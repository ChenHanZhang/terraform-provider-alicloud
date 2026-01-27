// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudNATGatewaySnatEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNATGatewaySnatEntryCreate,
		Read:   resourceAliCloudNATGatewaySnatEntryRead,
		Update: resourceAliCloudNATGatewaySnatEntryUpdate,
		Delete: resourceAliCloudNATGatewaySnatEntryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"eip_affinity": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"network_interface_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"snat_entry_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snat_entry_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snat_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snat_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"source_vswitch_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudNATGatewaySnatEntryCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateSnatEntry"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("snat_table_id"); ok {
		request["SnatTableId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("source_vswitch_id"); ok {
		request["SourceVSwitchId"] = v
	}
	if v, ok := d.GetOkExists("eip_affinity"); ok {
		request["EipAffinity"] = v
	}
	if v, ok := d.GetOk("snat_entry_name"); ok {
		request["SnatEntryName"] = v
	}
	if v, ok := d.GetOk("source_cidr"); ok {
		request["SourceCIDR"] = v
	}
	if v, ok := d.GetOk("snat_ip"); ok {
		request["SnatIp"] = v
	}
	if v, ok := d.GetOk("network_interface_id"); ok {
		request["NetworkInterfaceId"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.NATGW", "OperationConflict", "EIP_NOT_IN_GATEWAY", "OperationUnsupported.EipInBinding", "OperationUnsupported.EipNatBWPCheck", "InternalError"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_snat_entry", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["SnatTableId"], response["SnatEntryId"]))

	nATGatewayServiceV2 := NATGatewayServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, nATGatewayServiceV2.NATGatewaySnatEntryStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudNATGatewaySnatEntryRead(d, meta)
}

func resourceAliCloudNATGatewaySnatEntryRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nATGatewayServiceV2 := NATGatewayServiceV2{client}

	objectRaw, err := nATGatewayServiceV2.DescribeNATGatewaySnatEntry(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_snat_entry DescribeNATGatewaySnatEntry Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("eip_affinity", formatInt(objectRaw["EipAffinity"]))
	d.Set("network_interface_id", objectRaw["NetworkInterfaceId"])
	d.Set("snat_entry_name", objectRaw["SnatEntryName"])
	d.Set("snat_ip", objectRaw["SnatIp"])
	d.Set("source_cidr", objectRaw["SourceCIDR"])
	d.Set("source_vswitch_id", objectRaw["SourceVSwitchId"])
	d.Set("status", objectRaw["Status"])
	d.Set("snat_entry_id", objectRaw["SnatEntryId"])
	d.Set("snat_table_id", objectRaw["SnatTableId"])

	return nil
}

func resourceAliCloudNATGatewaySnatEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifySnatEntry"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["SnatEntryId"] = parts[1]
	request["SnatTableId"] = parts[0]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("eip_affinity") {
		update = true
		request["EipAffinity"] = d.Get("eip_affinity")
	}

	if d.HasChange("snat_entry_name") {
		update = true
		request["SnatEntryName"] = d.Get("snat_entry_name")
	}

	if d.HasChange("snat_ip") {
		update = true
		request["SnatIp"] = d.Get("snat_ip")
	}

	if d.HasChange("network_interface_id") {
		update = true
		request["NetworkInterfaceId"] = d.Get("network_interface_id")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectStatus.NATGW"}) || NeedRetry(err) {
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
		nATGatewayServiceV2 := NATGatewayServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, nATGatewayServiceV2.NATGatewaySnatEntryStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudNATGatewaySnatEntryRead(d, meta)
}

func resourceAliCloudNATGatewaySnatEntryDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteSnatEntry"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["SnatEntryId"] = parts[1]
	request["SnatTableId"] = parts[0]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.NATGW", "IncorretSnatEntryStatus", "OperationConflict"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidSnatTableId.NotFound", "InvalidSnatEntryId.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	nATGatewayServiceV2 := NATGatewayServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, nATGatewayServiceV2.NATGatewaySnatEntryStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
