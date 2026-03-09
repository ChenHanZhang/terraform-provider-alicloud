// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tidwall/sjson"
)

func resourceAliCloudSlbAccessControlListEntryAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSlbAccessControlListEntryAttachmentCreate,
		Read:   resourceAliCloudSlbAccessControlListEntryAttachmentRead,
		Update: resourceAliCloudSlbAccessControlListEntryAttachmentUpdate,
		Delete: resourceAliCloudSlbAccessControlListEntryAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"acl_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entry": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudSlbAccessControlListEntryAttachmentCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "AddAccessControlListEntry"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("acl_id"); ok {
		request["AclId"] = v
	}
	request["RegionId"] = client.RegionId

	aclEntrysDataList := make(map[string]interface{})

	if v, ok := d.GetOk("entry"); ok {
		aclEntrysDataList["entry"] = v
	}

	if v, ok := d.GetOk("comment"); ok {
		aclEntrysDataList["comment"] = v
	}

	AclEntrysMap := make([]interface{}, 0)
	AclEntrysMap = append(AclEntrysMap, aclEntrysDataList)
	request["AclEntrys"] = AclEntrysMap

	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "AclEntrys.0.entry", d.Get("entry"))
	_ = json.Unmarshal([]byte(jsonString), &request)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Slb", "2014-05-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_slb_access_control_list_entry_attachment", action, AlibabaCloudSdkGoERROR)
	}

	AclEntrysentryVar, _ := jsonpath.Get("AclEntrys[0].entry", request)
	d.SetId(fmt.Sprintf("%v:%v", request["AclId"], AclEntrysentryVar))

	return resourceAliCloudSlbAccessControlListEntryAttachmentRead(d, meta)
}

func resourceAliCloudSlbAccessControlListEntryAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	slbServiceV2 := SlbServiceV2{client}

	objectRaw, err := slbServiceV2.DescribeSlbAccessControlListEntryAttachment(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_slb_access_control_list_entry_attachment DescribeSlbAccessControlListEntryAttachment Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("acl_id", objectRaw["AclId"])

	aclEntryRawArrayObj, _ := jsonpath.Get("$.AclEntrys.AclEntry[*]", objectRaw)
	aclEntryRawArray := make([]interface{}, 0)
	if aclEntryRawArrayObj != nil {
		aclEntryRawArray = convertToInterfaceArray(aclEntryRawArrayObj)
	}
	aclEntryRaw := make(map[string]interface{})
	if len(aclEntryRawArray) > 0 {
		aclEntryRaw = aclEntryRawArray[0].(map[string]interface{})
	}

	d.Set("comment", aclEntryRaw["AclEntryComment"])
	d.Set("entry", aclEntryRaw["AclEntryIP"])

	return nil
}

func resourceAliCloudSlbAccessControlListEntryAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "SetAccessControlListEntryAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AclId"] = parts[0]
	request["RegionId"] = client.RegionId
	aclEntrysDataList := make(map[string]interface{})

	if d.HasChange("entry") {
		update = true
	}
	if v, ok := d.GetOk("entry"); ok {
		aclEntrysDataList["entry"] = v
	}

	if d.HasChange("comment") {
		update = true
	}
	if v, ok := d.GetOk("comment"); ok {
		aclEntrysDataList["comment"] = v
	}

	AclEntrysMap := make([]interface{}, 0)
	AclEntrysMap = append(AclEntrysMap, aclEntrysDataList)
	request["AclEntrys"] = AclEntrysMap

	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "AclEntrys.0.entry", parts[1])
	_ = json.Unmarshal([]byte(jsonString), &request)

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Slb", "2014-05-15", action, query, request, true)
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

	return resourceAliCloudSlbAccessControlListEntryAttachmentRead(d, meta)
}

func resourceAliCloudSlbAccessControlListEntryAttachmentDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "RemoveAccessControlListEntry"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AclId"] = parts[0]
	request["RegionId"] = client.RegionId

	aclEntrysDataList := make(map[string]interface{})

	if v, ok := d.GetOk("entry"); ok {
		aclEntrysDataList["entry"] = v
	}

	if v, ok := d.GetOk("comment"); ok {
		aclEntrysDataList["comment"] = v
	}

	AclEntrysMap := make([]interface{}, 0)
	AclEntrysMap = append(AclEntrysMap, aclEntrysDataList)
	request["AclEntrys"] = AclEntrysMap

	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "AclEntrys.0.entry", parts[1])
	_ = json.Unmarshal([]byte(jsonString), &request)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Slb", "2014-05-15", action, query, request, true)
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
