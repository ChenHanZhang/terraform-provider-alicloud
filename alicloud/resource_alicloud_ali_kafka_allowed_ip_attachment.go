// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/blues/jsonata-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAliKafkaAllowedIpAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAliKafkaAllowedIpAttachmentCreate,
		Read:   resourceAliCloudAliKafkaAllowedIpAttachmentRead,
		Update: resourceAliCloudAliKafkaAllowedIpAttachmentUpdate,
		Delete: resourceAliCloudAliKafkaAllowedIpAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allowed_list_ip": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allowed_list_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port_range": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudAliKafkaAllowedIpAttachmentCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "UpdateAllowedIp"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("port_range"); ok {
		request["PortRange"] = v
	}
	if v, ok := d.GetOk("allowed_list_type"); ok {
		request["AllowedListType"] = v
	}
	if v, ok := d.GetOk("allowed_list_ip"); ok {
		request["AllowedListIp"] = v
	}
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	request["UpdateType"] = "add"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"Instance.StatusNotRunning", "INSTANCE_STATUS_NOT_RUNNING"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ali_kafka_allowed_ip_attachment", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v:%v:%v", request["InstanceId"], request["PortRange"], request["AllowedListType"], request["AllowedListIp"]))

	return resourceAliCloudAliKafkaAllowedIpAttachmentRead(d, meta)
}

func resourceAliCloudAliKafkaAllowedIpAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	aliKafkaServiceV2 := AliKafkaServiceV2{client}

	checkValue00 := d.Get("port_range")
	if checkValue00 == "9093/9093" {
		objectRaw, err := aliKafkaServiceV2.DescribeAliKafkaAllowedIpAttachment(d.Id())
		if err != nil {
			if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_ali_kafka_allowed_ip_attachment DescribeAliKafkaAllowedIpAttachment Failed!!! %s", err)
				d.SetId("")
				return nil
			}
			return WrapError(err)
		}

		internetListRawObj, _ := jsonpath.Get("$.InternetList[*]", objectRaw)
		internetListRaw := make([]interface{}, 0)
		if internetListRawObj != nil {
			internetListRaw = convertToInterfaceArray(internetListRawObj)
		}

		d.Set("port_range", internetListRaw["PortRange"])

		e := jsonata.MustCompile("$reduce($split($reduce($.AllowedList.VpcList.AllowedIpList, function($i, $j){$i & ',' & $j}), ',', 1), function($i, $j){$i & '' & $j})")
		evaluation, _ := e.Eval(objectRaw)
		d.Set("allowed_list_ip", evaluation)
	}

	parts := strings.Split(d.Id(), ":")
	d.Set("instance_id", parts[0])
	d.Set("allowed_list_type", parts[2])

	return nil
}

func resourceAliCloudAliKafkaAllowedIpAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Cannot update resource Alicloud Resource Allowed Ip Attachment.")
	return nil
}

func resourceAliCloudAliKafkaAllowedIpAttachmentDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "UpdateAllowedIp"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["PortRange"] = parts[1]
	request["AllowedListIp"] = parts[3]
	request["AllowedListType"] = parts[2]
	request["RegionId"] = client.RegionId

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	request["UpdateType"] = "delete"
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("alikafka", "2019-09-16", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"Instance.StatusNotRunning", "INSTANCE_STATUS_NOT_RUNNING"}) || NeedRetry(err) {
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
