// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudRocketmqInstanceEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRocketmqInstanceEndpointCreate,
		Read:   resourceAliCloudRocketmqInstanceEndpointRead,
		Update: resourceAliCloudRocketmqInstanceEndpointUpdate,
		Delete: resourceAliCloudRocketmqInstanceEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vswitch_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudRocketmqInstanceEndpointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	instanceId := d.Get("instance_id")
	action := fmt.Sprintf("/instances/%s/endpoint", instanceId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	vpcInfo := make(map[string]interface{})

	if v, ok := d.GetOk("security_group_ids"); ok {
		securityGroupIds1, _ := jsonpath.Get("$", v)
		if securityGroupIds1 != nil && securityGroupIds1 != "" {
			vpcInfo["securityGroupIds"] = securityGroupIds1
		}
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		vpcInfo["vpcId"] = v
	}

	request["vpcInfo"] = vpcInfo

	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("RocketMQ", "2022-08-01", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rocketmq_instance_endpoint", action, AlibabaCloudSdkGoERROR)
	}

	dataidVar, _ := jsonpath.Get("$.data.id", response)
	d.SetId(fmt.Sprintf("%v:%v", instanceId, dataidVar))

	return resourceAliCloudRocketmqInstanceEndpointRead(d, meta)
}

func resourceAliCloudRocketmqInstanceEndpointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rocketmqServiceV2 := RocketmqServiceV2{client}

	objectRaw, err := rocketmqServiceV2.DescribeRocketmqInstanceEndpoint(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rocketmq_instance_endpoint DescribeRocketmqInstanceEndpoint Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("vpc_id", objectRaw["vpcId"])

	securityGroupIdsRaw := make([]interface{}, 0)
	if objectRaw["securityGroupIds"] != nil {
		securityGroupIdsRaw = convertToInterfaceArray(objectRaw["securityGroupIds"])
	}

	d.Set("security_group_ids", securityGroupIdsRaw)
	vSwitchIdsRaw := make([]interface{}, 0)
	if objectRaw["vSwitchIds"] != nil {
		vSwitchIdsRaw = convertToInterfaceArray(objectRaw["vSwitchIds"])
	}

	d.Set("vswitch_ids", vSwitchIdsRaw)

	parts := strings.Split(d.Id(), ":")
	d.Set("instance_id", parts[0])
	d.Set("id", parts[1])

	return nil
}

func resourceAliCloudRocketmqInstanceEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]*string
	var body map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	instanceId := parts[0]
	action := fmt.Sprintf("/instances/%s/endpoint", instanceId)
	request = make(map[string]interface{})
	query = make(map[string]*string)
	body = make(map[string]interface{})
	request["id"] = parts[1]

	if d.HasChange("security_group_ids") {
		update = true
	}
	if v, ok := d.GetOk("security_group_ids"); ok || d.HasChange("security_group_ids") {
		securityGroupIdsMapsArray := convertToInterfaceArray(v)

		request["securityGroupIds"] = securityGroupIdsMapsArray
	}

	body = request
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RoaPatch("RocketMQ", "2022-08-01", action, query, nil, body, true)
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

	return resourceAliCloudRocketmqInstanceEndpointRead(d, meta)
}

func resourceAliCloudRocketmqInstanceEndpointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	instanceId := parts[0]
	action := fmt.Sprintf("/instances/%s/endpoint", instanceId)
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	var err error
	request = make(map[string]interface{})
	query["id"] = StringPointer(parts[1])

	if v, ok := d.GetOk("endpoint_id"); ok {
		query["endpointId"] = StringPointer(v.(string))
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RoaDelete("RocketMQ", "2022-08-01", action, query, nil, nil, true)
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

	return nil
}
