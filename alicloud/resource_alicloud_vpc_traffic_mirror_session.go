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

func resourceAliCloudVpcTrafficMirrorSession() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcTrafficMirrorSessionCreate,
		Read:   resourceAlicloudVpcTrafficMirrorSessionRead,
		Update: resourceAlicloudVpcTrafficMirrorSessionUpdate,
		Delete: resourceAlicloudVpcTrafficMirrorSessionDelete,
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
				ForceNew: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"packet_length": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
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
			"traffic_mirror_filter_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"traffic_mirror_session_business_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_mirror_session_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_mirror_session_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_mirror_session_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_mirror_source_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"traffic_mirror_target_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"traffic_mirror_target_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"NetworkInterface", "SLB"}, false),
			},
			"virtual_network_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAlicloudVpcTrafficMirrorSessionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateTrafficMirrorSession"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("traffic_mirror_session_description"); ok {
		request["TrafficMirrorSessionDescription"] = v
	}

	if v, ok := d.GetOk("traffic_mirror_session_name"); ok {
		request["TrafficMirrorSessionName"] = v
	}

	if v, ok := d.GetOk("traffic_mirror_target_id"); ok {
		request["TrafficMirrorTargetId"] = v
	}

	if v, ok := d.GetOk("traffic_mirror_target_type"); ok {
		request["TrafficMirrorTargetType"] = v
	}

	if v, ok := d.GetOk("traffic_mirror_filter_id"); ok {
		request["TrafficMirrorFilterId"] = v
	}

	if v, ok := d.GetOk("virtual_network_id"); ok {
		request["VirtualNetworkId"] = v
	}

	if v, ok := d.GetOk("priority"); ok {
		request["Priority"] = v
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		request["Enabled"] = v
	}

	if v, ok := d.GetOk("packet_length"); ok {
		request["PacketLength"] = v
	}

	if v, ok := d.GetOk("traffic_mirror_source_ids"); ok {
		localData := v
		trafficMirrorSourceIdsMaps := localData.([]interface{})
		request["TrafficMirrorSourceIds"] = trafficMirrorSourceIdsMaps
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_traffic_mirror_session", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["TrafficMirrorSessionId"]))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Created"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcTrafficMirrorSessionStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAlicloudVpcTrafficMirrorSessionUpdate(d, meta)
}

func resourceAlicloudVpcTrafficMirrorSessionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcTrafficMirrorSession(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_traffic_mirror_session .DescribeVpcTrafficMirrorSession Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("enabled", object["enabled"])
	d.Set("packet_length", object["packet_length"])
	d.Set("priority", object["priority"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))
	d.Set("traffic_mirror_filter_id", object["traffic_mirror_filter_id"])
	d.Set("traffic_mirror_session_business_status", object["traffic_mirror_session_business_status"])
	d.Set("traffic_mirror_session_description", object["traffic_mirror_session_description"])
	d.Set("traffic_mirror_session_id", object["traffic_mirror_session_id"])
	d.Set("traffic_mirror_session_name", object["traffic_mirror_session_name"])
	d.Set("traffic_mirror_source_ids", object["traffic_mirror_source_ids"])
	d.Set("traffic_mirror_target_id", object["traffic_mirror_target_id"])
	d.Set("traffic_mirror_target_type", object["traffic_mirror_target_type"])
	d.Set("virtual_network_id", object["virtual_network_id"])

	return nil
}

func resourceAlicloudVpcTrafficMirrorSessionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "UpdateTrafficMirrorSessionAttribute"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["TrafficMirrorSessionId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("traffic_mirror_session_description") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_session_description"); ok {
			request["TrafficMirrorSessionDescription"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("traffic_mirror_session_name") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_session_name"); ok {
			request["TrafficMirrorSessionName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("traffic_mirror_target_id") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_target_id"); ok {
			request["TrafficMirrorTargetId"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("traffic_mirror_target_type") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_target_type"); ok {
			request["TrafficMirrorTargetType"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("traffic_mirror_filter_id") {
		update = true
		if v, ok := d.GetOk("traffic_mirror_filter_id"); ok {
			request["TrafficMirrorFilterId"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("virtual_network_id") {
		update = true
		if v, ok := d.GetOk("virtual_network_id"); ok {
			request["VirtualNetworkId"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("priority") {
		update = true
		if v, ok := d.GetOk("priority"); ok {
			request["Priority"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("enabled") {
		update = true
		if v, ok := d.GetOkExists("enabled"); ok {
			request["Enabled"] = v
		}
	}
	if d.HasChange("traffic_mirror_target_id") || d.HasChange("traffic_mirror_target_type") {
		update = true
		request["TrafficMirrorTargetId"] = d.Get("traffic_mirror_target_id")
		request["TrafficMirrorTargetType"] = d.Get("traffic_mirror_target_type")
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectStatus.TrafficMirrorSession", "OperationConflict", "IncorrectStatus.%s", "SystemBusy", "LastTokenProcessing", "OperationFailed.LastTokenProcessing", "ServiceUnavailable"}) || NeedRetry(err) {
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
		d.SetPartial("traffic_mirror_session_description")
		d.SetPartial("traffic_mirror_session_name")
		d.SetPartial("traffic_mirror_target_id")
		d.SetPartial("traffic_mirror_target_type")
		d.SetPartial("traffic_mirror_filter_id")
		d.SetPartial("virtual_network_id")
		d.SetPartial("priority")
		d.SetPartial("enabled")
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
	request["ResourceType"] = "TrafficMirrorSession"

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
	if !d.IsNewResource() && d.HasChange("traffic_mirror_source_ids") {
		update = true
		oldEntry, newEntry := d.GetChange("traffic_mirror_source_ids")
		removed := oldEntry
		added := newEntry

		if len(removed.([]interface{})) > 0 {
			action = "RemoveSourcesFromTrafficMirrorSession"
			conn, err = client.NewVpcClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["TrafficMirrorSessionId"] = d.Id()
			request["RegionId"] = client.RegionId
			request["ClientToken"] = buildClientToken(action)
			localData := removed

			trafficMirrorSourceIdsMaps := localData.([]interface{})
			request["TrafficMirrorSourceIds"] = trafficMirrorSourceIdsMaps

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
				if err != nil {
					if IsExpectedErrors(err, []string{"IncorrectStatus.TrafficMirrorSession", "OperationConflict", "IncorrectStatus.%s", "SystemBusy", "LastTokenProcessing", "OperationFailed.LastTokenProcessing", "ServiceUnavailable"}) || NeedRetry(err) {
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
				vpcServiceV2 := VpcServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Created"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcTrafficMirrorSessionStateRefreshFunc(d.Id(), []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}
			}

		}

		if len(added.([]interface{})) > 0 {
			action = "AddSourcesToTrafficMirrorSession"
			conn, err = client.NewVpcClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["TrafficMirrorSessionId"] = d.Id()
			request["RegionId"] = client.RegionId
			request["ClientToken"] = buildClientToken(action)
			localData := added

			trafficMirrorSourceIdsMaps := localData.([]interface{})
			request["TrafficMirrorSourceIds"] = trafficMirrorSourceIdsMaps

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
				if err != nil {
					if IsExpectedErrors(err, []string{"IncorrectStatus.TrafficMirrorSession", "OperationConflict", "IncorrectStatus.%s", "SystemBusy", "LastTokenProcessing", "OperationFailed.LastTokenProcessing", "ServiceUnavailable"}) || NeedRetry(err) {
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
				vpcServiceV2 := VpcServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{"Created"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcTrafficMirrorSessionStateRefreshFunc(d.Id(), []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}
			}

		}
	}
	update = false
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "TrafficMirrorSession"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcTrafficMirrorSessionRead(d, meta)
}

func resourceAlicloudVpcTrafficMirrorSessionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteTrafficMirrorSession"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["TrafficMirrorSessionId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectStatus.TrafficMirrorSession", "OperationConflict", "IncorrectStatus.%s", "SystemBusy", "LastTokenProcessing", "OperationFailed.LastTokenProcessing", "ServiceUnavailable", "IncorrectStatus.TrafficMirrorFilter"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"ResourceNotFound.TrafficMirrorSession"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcTrafficMirrorSessionStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
