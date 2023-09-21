// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudRocketMQTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRocketMQTopicCreate,
		Read:   resourceAliCloudRocketMQTopicRead,
		Update: resourceAliCloudRocketMQTopicUpdate,
		Delete: resourceAliCloudRocketMQTopicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"message_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"NORMAL", "FIFO", "DELAY", "TRANSACTION"}, false),
			},
			"remark": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[\u4E00-\u9FA5A-Za-z0-9_]+$"), "Custom remarks"),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"topic_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringMatch(regexp.MustCompile("^[%a-zA-Z0-9_-]+$"), "Topic name and identification"),
			},
		},
	}
}

func resourceAliCloudRocketMQTopicCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateTopic"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewRocketmqClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["instanceId"] = d.Get("instance_id")
	request["topicName"] = d.Get("topic_name")

	objectDataLocalMap := make(map[string]interface{})
	if v, ok := d.GetOk("message_type"); ok {
		objectDataLocalMap["messageType"] = v
	}
	if v, ok := d.GetOk("remark"); ok {
		objectDataLocalMap["remark"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_rocket_m_q_topic", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["instanceId"], request["topicName"]))

	rocketMQServiceV2 := RocketMQServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"RUNNING"}, d.Timeout(schema.TimeoutCreate), 50*time.Second, rocketMQServiceV2.RocketMQTopicStateRefreshFunc(d.Id(), "status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudRocketMQTopicUpdate(d, meta)
}

func resourceAliCloudRocketMQTopicRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rocketMQServiceV2 := RocketMQServiceV2{client}

	objectRaw, err := rocketMQServiceV2.DescribeRocketMQTopic(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rocket_m_q_topic DescribeRocketMQTopic Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["createTime"])
	d.Set("message_type", objectRaw["messageType"])
	d.Set("remark", objectRaw["remark"])
	d.Set("status", objectRaw["status"])
	d.Set("instance_id", objectRaw["instanceId"])
	d.Set("topic_name", objectRaw["topicName"])

	return nil
}

func resourceAliCloudRocketMQTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	parts := strings.Split(d.Id(), ":")
	action := "UpdateTopic"
	conn, err := client.NewRocketmqClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["instanceId"] = parts[0]
	request["topicName"] = parts[1]
	if !d.IsNewResource() && d.HasChange("remark") {
		update = true
		objectDataLocalMap := make(map[string]interface{})
		if v := d.Get("remark"); !IsNil(v) {
			objectDataLocalMap["remark"] = v
		}
		objectDataLocalMapJson, err := json.Marshal(objectDataLocalMap)
		if err != nil {
			return WrapError(err)
		}
		request["body"] = string(objectDataLocalMapJson)
	}

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
		stateConf := BuildStateConf([]string{}, []string{"RUNNING"}, d.Timeout(schema.TimeoutUpdate), 30*time.Second, rocketMQServiceV2.RocketMQTopicStateRefreshFunc(d.Id(), "status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudRocketMQTopicRead(d, meta)
}

func resourceAliCloudRocketMQTopicDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteTopic"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewRocketmqClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["instanceId"] = parts[0]
	request["topicName"] = parts[1]

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
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 10*time.Second, rocketMQServiceV2.RocketMQTopicStateRefreshFunc(d.Id(), "topicName", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
