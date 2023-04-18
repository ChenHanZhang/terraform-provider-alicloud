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
)

func resourceAliCloudVpcTrafficMirrorFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcTrafficMirrorFilterCreate,
		Read:   resourceAlicloudVpcTrafficMirrorFilterRead,
		Update: resourceAlicloudVpcTrafficMirrorFilterUpdate,
		Delete: resourceAlicloudVpcTrafficMirrorFilterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"egress_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_port_range": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"action": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_port_range": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"source_cidr_block": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"traffic_mirror_filter_rule_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_cidr_block": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"ingress_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_port_range": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"action": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_port_range": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"source_cidr_block": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"traffic_mirror_filter_rule_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_cidr_block": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"traffic_mirror_filter_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_mirror_filter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_mirror_filter_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dry_run": {
				Type:       schema.TypeBool,
				Optional:   true,
				Deprecated: "This Field was deprecated.",
			},
		},
	}
}

func resourceAlicloudVpcTrafficMirrorFilterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateTrafficMirrorFilter"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("traffic_mirror_filter_description"); ok {
		request["TrafficMirrorFilterDescription"] = v
	}

	if v, ok := d.GetOk("traffic_mirror_filter_name"); ok {
		request["TrafficMirrorFilterName"] = v
	}

	if v, ok := d.GetOk("ingress_rules"); ok {
		localData := v
		ingressRulesMaps := make([]map[string]interface{}, 0)
		for _, dataLoop := range localData.([]interface{}) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Action"] = dataLoopTmp["action"]
			dataLoopMap["SourceCidrBlock"] = dataLoopTmp["source_cidr_block"]
			dataLoopMap["Protocol"] = dataLoopTmp["protocol"]
			dataLoopMap["DestinationPortRange"] = dataLoopTmp["destination_port_range"]
			dataLoopMap["Priority"] = dataLoopTmp["priority"]
			dataLoopMap["DestinationCidrBlock"] = dataLoopTmp["destination_cidr_block"]
			dataLoopMap["SourcePortRange"] = dataLoopTmp["source_port_range"]
			ingressRulesMaps = append(ingressRulesMaps, dataLoopMap)
		}
		request["IngressRules"] = ingressRulesMaps
	}

	if v, ok := d.GetOk("egress_rules"); ok {
		localData1 := v
		egressRulesMaps := make([]map[string]interface{}, 0)
		for _, dataLoop1 := range localData1.([]interface{}) {
			dataLoop1Tmp := dataLoop1.(map[string]interface{})
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["Action"] = dataLoop1Tmp["action"]
			dataLoop1Map["SourceCidrBlock"] = dataLoop1Tmp["source_cidr_block"]
			dataLoop1Map["Protocol"] = dataLoop1Tmp["protocol"]
			dataLoop1Map["DestinationPortRange"] = dataLoop1Tmp["destination_port_range"]
			dataLoop1Map["Priority"] = dataLoop1Tmp["priority"]
			dataLoop1Map["DestinationCidrBlock"] = dataLoop1Tmp["destination_cidr_block"]
			dataLoop1Map["SourcePortRange"] = dataLoop1Tmp["source_port_range"]
			egressRulesMaps = append(egressRulesMaps, dataLoop1Map)
		}
		request["EgressRules"] = egressRulesMaps
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "IncorrectStatus.%s", "ServiceUnavailable", "SystemBusy"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_traffic_mirror_filter", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["TrafficMirrorFilterId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Created"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcTrafficMirrorFilterStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcTrafficMirrorFilterUpdate(d, meta)
}

func resourceAlicloudVpcTrafficMirrorFilterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcTrafficMirrorFilter(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_traffic_mirror_filter .DescribeVpcTrafficMirrorFilter Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("egress_rules", object["egress_rules"])
	d.Set("ingress_rules", object["ingress_rules"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))
	d.Set("traffic_mirror_filter_description", object["traffic_mirror_filter_description"])
	d.Set("traffic_mirror_filter_id", object["traffic_mirror_filter_id"])
	d.Set("traffic_mirror_filter_name", object["traffic_mirror_filter_name"])

	return nil
}

func resourceAlicloudVpcTrafficMirrorFilterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "UpdateTrafficMirrorFilterAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["TrafficMirrorFilterId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("traffic_mirror_filter_description") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_filter_description"); ok {
			request["TrafficMirrorFilterDescription"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("traffic_mirror_filter_name") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_filter_name"); ok {
			request["TrafficMirrorFilterName"] = v
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
		d.SetPartial("traffic_mirror_filter_description")
		d.SetPartial("traffic_mirror_filter_name")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewVpcClient()
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
	request["ResourceType"] = "TRAFFICMIRRORFILTER"
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
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "TRAFFICMIRRORFILTER"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcTrafficMirrorFilterRead(d, meta)
}

func resourceAlicloudVpcTrafficMirrorFilterDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteTrafficMirrorFilter"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["TrafficMirrorFilterId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.TrafficMirrorFilter", "IncorrectStatus.TrafficMirrorRule", "OperationConflict", "IncorrectStatus.%s", "ServiceUnavailable", "SystemBusy"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"ResourceNotFound.TrafficMirrorFilter"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcTrafficMirrorFilterStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
