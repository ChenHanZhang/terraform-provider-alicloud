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

func resourceAliCloudCenTransitRouterPolicyBasedRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCenTransitRouterPolicyBasedRouteCreate,
		Read:   resourceAliCloudCenTransitRouterPolicyBasedRouteRead,
		Update: resourceAliCloudCenTransitRouterPolicyBasedRouteUpdate,
		Delete: resourceAliCloudCenTransitRouterPolicyBasedRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"address_family": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dest_port_range": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"destination_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dscp": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_port_range": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_route_table_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudCenTransitRouterPolicyBasedRouteCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateTransitRouterPolicyBasedRoute"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	request["TargetRouteTableId"] = d.Get("target_route_table_id")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("protocol"); ok {
		request["Protocol"] = v
	}
	request["SourceCidr"] = d.Get("source_cidr")
	if v, ok := d.GetOk("source_port_range"); ok {
		request["SourcePortRange"] = v
	}
	request["Priority"] = d.Get("priority")
	request["DestinationCidr"] = d.Get("destination_cidr")
	if v, ok := d.GetOk("name"); ok {
		request["Name"] = v
	}
	if v, ok := d.GetOk("dscp"); ok {
		request["Dscp"] = v
	}
	if v, ok := d.GetOk("address_family"); ok {
		request["AddressFamily"] = v
	}
	if v, ok := d.GetOk("dest_port_range"); ok {
		request["DestPortRange"] = v
	}
	request["PolicyTableId"] = d.Get("policy_table_id")
	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	wait := incrementalWait(3*time.Second, 0*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Cbn", "2017-09-12", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cen_transit_router_policy_based_route", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["PolicyBasedRouteId"]))

	cenServiceV2 := CenServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, cenServiceV2.CenTransitRouterPolicyBasedRouteStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudCenTransitRouterPolicyBasedRouteRead(d, meta)
}

func resourceAliCloudCenTransitRouterPolicyBasedRouteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cenServiceV2 := CenServiceV2{client}

	objectRaw, err := cenServiceV2.DescribeCenTransitRouterPolicyBasedRoute(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cen_transit_router_policy_based_route DescribeCenTransitRouterPolicyBasedRoute Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("address_family", objectRaw["AddressFamily"])
	d.Set("description", objectRaw["Description"])
	d.Set("dest_port_range", objectRaw["DestPortRange"])
	d.Set("destination_cidr", objectRaw["DestinationCidr"])
	d.Set("dscp", objectRaw["Dscp"])
	d.Set("name", objectRaw["Name"])
	d.Set("policy_table_id", objectRaw["PolicyTableId"])
	d.Set("priority", objectRaw["Priority"])
	d.Set("protocol", objectRaw["Protocol"])
	d.Set("source_cidr", objectRaw["SourceCidr"])
	d.Set("source_port_range", objectRaw["SourcePortRange"])
	d.Set("status", objectRaw["Status"])
	d.Set("target_route_table_id", objectRaw["TargetRouteTableId"])

	return nil
}

func resourceAliCloudCenTransitRouterPolicyBasedRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateTransitRouterPolicyBasedRoute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["PolicyBasedRouteId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if d.HasChange("name") {
		update = true
		request["Name"] = d.Get("name")
	}

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Cbn", "2017-09-12", action, query, request, true)
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
		cenServiceV2 := CenServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, cenServiceV2.CenTransitRouterPolicyBasedRouteStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudCenTransitRouterPolicyBasedRouteRead(d, meta)
}

func resourceAliCloudCenTransitRouterPolicyBasedRouteDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteTransitRouterPolicyBasedRoute"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["PolicyBasedRouteId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	wait := incrementalWait(3*time.Second, 0*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Cbn", "2017-09-12", action, query, request, true)
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

	cenServiceV2 := CenServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, cenServiceV2.CenTransitRouterPolicyBasedRouteStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
