// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudNATGatewayFullNat() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNATGatewayFullNatCreate,
		Read:   resourceAliCloudNATGatewayFullNatRead,
		Update: resourceAliCloudNATGatewayFullNatUpdate,
		Delete: resourceAliCloudNATGatewayFullNatDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"access_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_port": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[0-9]*$"), "The backend port used for port mapping in the FULLNAT entry. Valid values: **1** to **65535**."),
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_nat_entry_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"full_nat_entry_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_nat_entry_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"full_nat_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_protocol": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"TCP", "UDP"}, false),
			},
			"nat_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"nat_ip_port": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[0-9]*$"), "The frontend port used for port mapping in the FULLNAT entry. Valid values: **1** to **65535**."),
			},
			"network_interface_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudNATGatewayFullNatCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateFullNatEntry"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("full_nat_table_id"); ok {
		request["FullNatTableId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("full_nat_entry_name"); ok {
		request["FullNatEntryName"] = v
	}
	request["IpProtocol"] = d.Get("ip_protocol")
	if v, ok := d.GetOk("nat_ip_port"); ok {
		request["NatIpPort"] = v
	}
	request["NatIp"] = d.Get("nat_ip")
	request["NetworkInterfaceId"] = d.Get("network_interface_id")
	if v, ok := d.GetOk("full_nat_entry_description"); ok {
		request["FullNatEntryDescription"] = v
	}
	request["AccessPort"] = d.Get("access_port")
	request["AccessIp"] = d.Get("access_ip")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_n_a_t_gateway_full_nat", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["FullNatTableId"], response["FullNatEntryId"]))

	return resourceAliCloudNATGatewayFullNatRead(d, meta)
}

func resourceAliCloudNATGatewayFullNatRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nATGatewayServiceV2 := NATGatewayServiceV2{client}

	objectRaw, err := nATGatewayServiceV2.DescribeNATGatewayFullNat(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_n_a_t_gateway_full_nat DescribeNATGatewayFullNat Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("access_ip", objectRaw["AccessIp"])
	d.Set("access_port", objectRaw["AccessPort"])
	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("full_nat_entry_description", objectRaw["FullNatEntryDescription"])
	d.Set("full_nat_entry_name", objectRaw["FullNatEntryName"])
	d.Set("ip_protocol", objectRaw["IpProtocol"])
	d.Set("nat_ip", objectRaw["NatIp"])
	d.Set("nat_ip_port", objectRaw["NatIpPort"])
	d.Set("network_interface_id", objectRaw["NetworkInterfaceId"])
	d.Set("status", objectRaw["FullNatEntryStatus"])
	d.Set("full_nat_entry_id", objectRaw["FullNatEntryId"])
	d.Set("full_nat_table_id", objectRaw["FullNatTableId"])

	return nil
}

func resourceAliCloudNATGatewayFullNatUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyFullNatEntryAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["FullNatEntryId"] = parts[1]
	request["FullNatTableId"] = parts[0]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("full_nat_entry_name") {
		update = true
		request["FullNatEntryName"] = d.Get("full_nat_entry_name")
	}

	if d.HasChange("ip_protocol") {
		update = true
	}
	request["IpProtocol"] = d.Get("ip_protocol")
	if d.HasChange("nat_ip_port") {
		update = true
		request["NatIpPort"] = d.Get("nat_ip_port")
	}

	if d.HasChange("nat_ip") {
		update = true
	}
	request["NatIp"] = d.Get("nat_ip")
	if d.HasChange("network_interface_id") {
		update = true
	}
	request["NetworkInterfaceId"] = d.Get("network_interface_id")
	if d.HasChange("full_nat_entry_description") {
		update = true
		request["FullNatEntryDescription"] = d.Get("full_nat_entry_description")
	}

	if d.HasChange("access_port") {
		update = true
	}
	request["AccessPort"] = d.Get("access_port")
	if d.HasChange("access_ip") {
		update = true
	}
	request["AccessIp"] = d.Get("access_ip")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
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

	return resourceAliCloudNATGatewayFullNatRead(d, meta)
}

func resourceAliCloudNATGatewayFullNatDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteFullNatEntry"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["FullNatEntryId"] = parts[1]
	request["FullNatTableId"] = parts[0]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict"}) || NeedRetry(err) {
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

	return nil
}
