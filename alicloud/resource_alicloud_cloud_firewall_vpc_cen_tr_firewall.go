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

func resourceAliCloudCloudFirewallVpcCenTrFirewall() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudFirewallVpcCenTrFirewallCreate,
		Read:   resourceAliCloudCloudFirewallVpcCenTrFirewallRead,
		Update: resourceAliCloudCloudFirewallVpcCenTrFirewallUpdate,
		Delete: resourceAliCloudCloudFirewallVpcCenTrFirewallDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(41 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(46 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"firewall_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"firewall_eni_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firewall_eni_vpc_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firewall_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"firewall_subnet_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"firewall_vpc_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region_no": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"route_mode": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"managed"}, false),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tr_attachment_master_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tr_attachment_master_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tr_attachment_slave_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tr_attachment_slave_zone": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"transit_router_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudCloudFirewallVpcCenTrFirewallCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateTrFirewallV2"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["CenId"] = d.Get("cen_id")
	request["TrAttachmentMasterCidr"] = d.Get("tr_attachment_master_cidr")
	if v, ok := d.GetOk("firewall_description"); ok {
		request["FirewallDescription"] = v
	}
	if v, ok := d.GetOk("tr_attachment_slave_zone"); ok {
		request["TrAttachmentSlaveZone"] = v
	}
	request["FirewallSubnetCidr"] = d.Get("firewall_subnet_cidr")
	request["RouteMode"] = d.Get("route_mode")
	request["RegionNo"] = d.Get("region_no")
	request["TransitRouterId"] = d.Get("transit_router_id")
	request["FirewallName"] = d.Get("firewall_name")
	request["TrAttachmentSlaveCidr"] = d.Get("tr_attachment_slave_cidr")
	if v, ok := d.GetOk("tr_attachment_master_zone"); ok {
		request["TrAttachmentMasterZone"] = v
	}
	request["FirewallVpcCidr"] = d.Get("firewall_vpc_cidr")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"ErrorTrResourceNotReady"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cloud_firewall_vpc_cen_tr_firewall", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["FirewallId"]))

	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Ready"}, d.Timeout(schema.TimeoutCreate), 30*time.Second, cloudFirewallServiceV2.CloudFirewallVpcCenTrFirewallStateRefreshFunc(d.Id(), "FirewallStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudCloudFirewallVpcCenTrFirewallRead(d, meta)
}

func resourceAliCloudCloudFirewallVpcCenTrFirewallRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}

	objectRaw, err := cloudFirewallServiceV2.DescribeCloudFirewallVpcCenTrFirewall(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cloud_firewall_vpc_cen_tr_firewall DescribeCloudFirewallVpcCenTrFirewall Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("cen_id", objectRaw["CenId"])
	d.Set("firewall_description", objectRaw["FirewallDescription"])
	d.Set("firewall_eni_id", objectRaw["FirewallEniId"])
	d.Set("firewall_eni_vpc_id", objectRaw["FirewallEniVpcId"])
	d.Set("firewall_name", objectRaw["FirewallName"])
	d.Set("firewall_subnet_cidr", objectRaw["FirewallSubnetCidr"])
	d.Set("firewall_vpc_cidr", objectRaw["FirewallVpcCidr"])
	d.Set("region_no", objectRaw["RegionNo"])
	d.Set("route_mode", objectRaw["RouteMode"])
	d.Set("status", objectRaw["FirewallStatus"])
	d.Set("tr_attachment_master_cidr", objectRaw["TrAttachmentMasterCidr"])
	d.Set("tr_attachment_slave_cidr", objectRaw["TrAttachmentSlaveCidr"])
	d.Set("transit_router_id", objectRaw["TransitRouterId"])

	return nil
}

func resourceAliCloudCloudFirewallVpcCenTrFirewallUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyTrFirewallV2Configuration"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["FirewallId"] = d.Id()

	if d.HasChange("firewall_name") {
		update = true
	}
	request["FirewallName"] = d.Get("firewall_name")
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

	return resourceAliCloudCloudFirewallVpcCenTrFirewallRead(d, meta)
}

func resourceAliCloudCloudFirewallVpcCenTrFirewallDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteTrFirewallV2"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["FirewallId"] = d.Id()

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
		if IsExpectedErrors(err, []string{"ErrorTrFirewallNotExist"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 30*time.Second, cloudFirewallServiceV2.CloudFirewallVpcCenTrFirewallStateRefreshFunc(d.Id(), "FirewallStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
