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

func resourceAliCloudAliKafkaTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAliKafkaTopicCreate,
		Read:   resourceAliCloudAliKafkaTopicRead,
		Update: resourceAliCloudAliKafkaTopicUpdate,
		Delete: resourceAliCloudAliKafkaTopicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(16 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"compact_topic": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"config": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"local_topic": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"min_insync_replicas": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"partition_num": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remark": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_factor": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"topic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudAliKafkaTopicCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateTopic"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("topic"); ok {
		request["Topic"] = v
	}
	request["RegionId"] = client.RegionId

	request["Remark"] = d.Get("remark")
	if v, ok := d.GetOkExists("compact_topic"); ok {
		request["CompactTopic"] = v
	}
	if v, ok := d.GetOkExists("partition_num"); ok {
		request["PartitionNum"] = v
	}
	if v, ok := d.GetOkExists("local_topic"); ok {
		request["LocalTopic"] = v
	}
	if v, ok := d.GetOkExists("replication_factor"); ok {
		request["ReplicationFactor"] = v
	}
	if v, ok := d.GetOkExists("min_insync_replicas"); ok {
		request["MinInsyncReplicas"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if v, ok := d.GetOk("config"); ok {
		request["Config"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_alikafka_topic", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], request["Topic"]))

	aliKafkaServiceV2 := AliKafkaServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"0"}, d.Timeout(schema.TimeoutCreate), 10*time.Second, aliKafkaServiceV2.AliKafkaTopicStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudAliKafkaTopicUpdate(d, meta)
}

func resourceAliCloudAliKafkaTopicRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	aliKafkaServiceV2 := AliKafkaServiceV2{client}

	objectRaw, err := aliKafkaServiceV2.DescribeAliKafkaTopic(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_alikafka_topic DescribeAliKafkaTopic Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("compact_topic", objectRaw["CompactTopic"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("local_topic", objectRaw["LocalTopic"])
	d.Set("partition_num", objectRaw["PartitionNum"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("remark", objectRaw["Remark"])
	d.Set("status", objectRaw["Status"])
	d.Set("instance_id", objectRaw["InstanceId"])
	d.Set("topic", objectRaw["Topic"])

	tagsMaps, _ := jsonpath.Get("$.Tags.TagVO", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	checkValue00 := d.Get("status")
	if checkValue00 == "0" {
		objectRaw, err = aliKafkaServiceV2.DescribeTopicGetTopicDetail(d.Id())
		if err != nil && !NotFoundError(err) {
			return WrapError(err)
		}

		d.Set("config", objectRaw["Config"])
		d.Set("value", objectRaw["Value"])
		d.Set("instance_id", objectRaw["InstanceId"])
		d.Set("topic", objectRaw["Topic"])

	}

	return nil
}

func resourceAliCloudAliKafkaTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyTopicRemark"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["Topic"] = parts[1]
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("remark") {
		update = true
	}
	request["Remark"] = d.Get("remark")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "UpdateTopicConfig"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["Topic"] = parts[1]
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("config") {
		update = true
	}
	request["Config"] = d.Get("config")
	if d.HasChange("value") {
		update = true
	}
	request["Value"] = d.Get("value")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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
	update = false
	parts = strings.Split(d.Id(), ":")
	action = "ModifyPartitionNum"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["Topic"] = parts[1]
	request["RegionId"] = client.RegionId
	if !d.IsNewResource() && d.HasChange("partition_num") {
		update = true
	}
	request["AddPartitionNum"] = d.Get("partition_num")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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

	if d.HasChange("tags") {
		aliKafkaServiceV2 := AliKafkaServiceV2{client}
		if err := aliKafkaServiceV2.SetResourceTags(d, "TOPIC"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudAliKafkaTopicRead(d, meta)
}

func resourceAliCloudAliKafkaTopicDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteTopic"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["Topic"] = parts[1]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
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

	aliKafkaServiceV2 := AliKafkaServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 30*time.Second, aliKafkaServiceV2.AliKafkaTopicStateRefreshFunc(d.Id(), "Topic", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
