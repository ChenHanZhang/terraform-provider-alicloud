// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudRocketMQConsumerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRocketMQConsumerGroupCreate,
		Read:   resourceAliCloudRocketMQConsumerGroupRead,
		Update: resourceAliCloudRocketMQConsumerGroupUpdate,
		Delete: resourceAliCloudRocketMQConsumerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"consume_retry_policy": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retry_policy": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_retry_times": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"consumer_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delivery_order_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"remark": {
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

func resourceAliCloudRocketMQConsumerGroupCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateConsumerGroup"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewRocketmqClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["instanceId"] = d.Get("instance_id")
	request["consumerGroupId"] = d.Get("consumer_group_id")

	objectDataLocalMap := make(map[string]interface{})
	if v, ok := d.GetOk("remark"); ok {
		objectDataLocalMap["remark"] = v
	}
	if v, ok := d.GetOk("delivery_order_type"); ok {
		objectDataLocalMap["deliveryOrderType"] = v
	}
	if v := d.Get("consume_retry_policy"); !IsNil(v) {
		consumeRetryPolicy := make(map[string]interface{})
		nodeNative2, _ := jsonpath.Get("$[0].max_retry_times", v)
		if nodeNative2 != "" {
			consumeRetryPolicy["maxRetryTimes"] = nodeNative2
		}
		nodeNative3, _ := jsonpath.Get("$[0].retry_policy", v)
		if nodeNative3 != "" {
			consumeRetryPolicy["retryPolicy"] = nodeNative3
		}
		objectDataLocalMap["consumeRetryPolicy"] = consumeRetryPolicy
	}
	request["body"] = objectDataLocalMap
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2022-08-01"), StringPointer("AK"), nil, request, &runtime)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rocket_m_q_consumer_group", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["instanceId"], request["consumerGroupId"]))

	rocketMQServiceV2 := RocketMQServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"RUNNING"}, d.Timeout(schema.TimeoutCreate), 30*time.Second, rocketMQServiceV2.RocketMQConsumerGroupStateRefreshFunc(d.Id(), "status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudRocketMQConsumerGroupUpdate(d, meta)
}

func resourceAliCloudRocketMQConsumerGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rocketMQServiceV2 := RocketMQServiceV2{client}

	objectRaw, err := rocketMQServiceV2.DescribeRocketMQConsumerGroup(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rocket_m_q_consumer_group DescribeRocketMQConsumerGroup Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["createTime"])
	d.Set("delivery_order_type", objectRaw["deliveryOrderType"])
	d.Set("remark", objectRaw["remark"])
	d.Set("status", objectRaw["status"])
	d.Set("consumer_group_id", objectRaw["consumerGroupId"])
	d.Set("instance_id", objectRaw["instanceId"])
	consumeRetryPolicyMaps := make([]map[string]interface{}, 0)
	consumeRetryPolicyMap := make(map[string]interface{})
	consumeRetryPolicy1Raw := make(map[string]interface{})
	if objectRaw["consumeRetryPolicy"] != nil {
		consumeRetryPolicy1Raw = objectRaw["consumeRetryPolicy"].(map[string]interface{})
	}
	if len(consumeRetryPolicy1Raw) > 0 {
		consumeRetryPolicyMap["max_retry_times"] = consumeRetryPolicy1Raw["maxRetryTimes"]
		consumeRetryPolicyMap["retry_policy"] = consumeRetryPolicy1Raw["retryPolicy"]
		consumeRetryPolicyMaps = append(consumeRetryPolicyMaps, consumeRetryPolicyMap)
	}
	d.Set("consume_retry_policy", consumeRetryPolicyMaps)

	return nil
}

func resourceAliCloudRocketMQConsumerGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	parts := strings.Split(d.Id(), ":")
	action := "UpdateConsumerGroup"
	conn, err := client.NewRocketmqClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["instanceId"] = parts[0]
	request["consumerGroupId"] = parts[1]
	objectDataLocalMap := make(map[string]interface{})
	if d.HasChange("remark") {
		update = true
		if v, ok := d.GetOk("remark"); ok {
			objectDataLocalMap["remark"] = v
		}
	}
	if d.HasChange("delivery_order_type") {
		update = true
		if v, ok := d.GetOk("delivery_order_type"); ok {
			objectDataLocalMap["deliveryOrderType"] = v
		}
	}
	if d.HasChange("consume_retry_policy") {
		update = true
		if v := d.Get("consume_retry_policy"); !IsNil(v) {
			consumeRetryPolicy := make(map[string]interface{})
			nodeNative2, _ := jsonpath.Get("$[0].max_retry_times", v)
			if nodeNative2 != "" {
				consumeRetryPolicy["maxRetryTimes"] = nodeNative2
			}
			nodeNative3, _ := jsonpath.Get("$[0].retry_policy", v)
			if nodeNative3 != "" {
				consumeRetryPolicy["retryPolicy"] = nodeNative3
			}
			objectDataLocalMap["consumeRetryPolicy"] = consumeRetryPolicy
		}
	}
	request["body"] = objectDataLocalMap
	if update {
		runtime := util.RuntimeOptions{}
		runtime.SetAutoretry(true)
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("PATCH"), StringPointer("2022-08-01"), StringPointer("AK"), nil, request, &runtime)

			if err != nil {
				if NeedRetry(err) {
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
		rocketMQServiceV2 := RocketMQServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"RUNNING"}, d.Timeout(schema.TimeoutUpdate), 30*time.Second, rocketMQServiceV2.RocketMQConsumerGroupStateRefreshFunc(d.Id(), "status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudRocketMQConsumerGroupRead(d, meta)
}

func resourceAliCloudRocketMQConsumerGroupDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteConsumerGroup"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewRocketmqClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["instanceId"] = parts[0]
	request["consumerGroupId"] = parts[1]

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("DELETE"), StringPointer("2022-08-01"), StringPointer("AK"), nil, request, &runtime)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	rocketMQServiceV2 := RocketMQServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 30*time.Second, rocketMQServiceV2.RocketMQConsumerGroupStateRefreshFunc(d.Id(), "consumerGroupId", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
