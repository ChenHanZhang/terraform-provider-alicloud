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

func resourceAliCloudNATGatewayNatIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNATGatewayNatIpCreate,
		Read:   resourceAliCloudNATGatewayNatIpRead,
		Update: resourceAliCloudNATGatewayNatIpUpdate,
		Delete: resourceAliCloudNATGatewayNatIpDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ipv4_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"nat_gateway_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nat_ip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nat_ip_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nat_ip_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat_ip_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat_ip_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudNATGatewayNatIpCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateNatIp"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("nat_gateway_id"); ok {
		request["NatGatewayId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("nat_ip_name"); ok {
		request["NatIpName"] = v
	}
	request["NatIpCidr"] = d.Get("nat_ip_cidr")
	if v, ok := d.GetOk("nat_ip"); ok {
		request["NatIp"] = v
	}
	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if v, ok := d.GetOk("ipv4_prefix"); ok {
		request["Ipv4Prefix"] = v
	}
	if v, ok := d.GetOk("nat_ip_description"); ok {
		request["NatIpDescription"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "IncorrectStatus.NatGateway"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_nat_ip", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["NatGatewayId"], response["NatIpId"]))

	nATGatewayServiceV2 := NATGatewayServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, nATGatewayServiceV2.NATGatewayNatIpStateRefreshFunc(d.Id(), "NatIpStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudNATGatewayNatIpRead(d, meta)
}

func resourceAliCloudNATGatewayNatIpRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nATGatewayServiceV2 := NATGatewayServiceV2{client}

	objectRaw, err := nATGatewayServiceV2.DescribeNATGatewayNatIp(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_nat_ip DescribeNATGatewayNatIp Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("ipv4_prefix", objectRaw["Ipv4Prefix"])
	d.Set("nat_ip", objectRaw["NatIp"])
	d.Set("nat_ip_cidr", objectRaw["NatIpCidr"])
	d.Set("nat_ip_description", objectRaw["NatIpDescription"])
	d.Set("nat_ip_name", objectRaw["NatIpName"])
	d.Set("status", objectRaw["NatIpStatus"])
	d.Set("nat_gateway_id", objectRaw["NatGatewayId"])
	d.Set("nat_ip_id", objectRaw["NatIpId"])

	return nil
}

func resourceAliCloudNATGatewayNatIpUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyNatIpAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["NatIpId"] = parts[1]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("nat_ip_name") {
		update = true
		request["NatIpName"] = d.Get("nat_ip_name")
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if d.HasChange("nat_ip_description") {
		update = true
		request["NatIpDescription"] = d.Get("nat_ip_description")
	}

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

	return resourceAliCloudNATGatewayNatIpRead(d, meta)
}

func resourceAliCloudNATGatewayNatIpDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteNatIp"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["NatGatewayId"] = parts[0]
	request["NatIpId"] = parts[1]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if v, ok := d.GetOk("ipv4_prefix"); ok {
		request["Ipv4Prefix"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "IncorrectStatus.NatGateway"}) || NeedRetry(err) {
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

	nATGatewayServiceV2 := NATGatewayServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, nATGatewayServiceV2.NATGatewayNatIpStateRefreshFunc(d.Id(), "NatIpStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
