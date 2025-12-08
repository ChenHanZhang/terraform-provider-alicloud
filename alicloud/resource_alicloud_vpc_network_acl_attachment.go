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

func resourceAliCloudVpcNetworkAclAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudVpcNetworkAclAttachmentCreate,
		Read:   resourceAliCloudVpcNetworkAclAttachmentRead,
		Delete: resourceAliCloudVpcNetworkAclAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"network_acl_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"VSwitch"}, false),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudVpcNetworkAclAttachmentCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "AssociateNetworkAcl"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("network_acl_id"); ok {
		request["NetworkAclId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	resourceDataList := make(map[string]interface{})

	if v, ok := d.GetOk("resource_id"); ok {
		resourceDataList["ResourceId"] = v
	}

	if v, ok := d.GetOk("resource_type"); ok {
		resourceDataList["ResourceType"] = v
	}

	ResourceMap := make([]interface{}, 0)
	ResourceMap = append(ResourceMap, resourceDataList)
	request["Resource"] = ResourceMap

	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "Resource.0.ResourceId", d.Get("resource_id"))
	_ = json.Unmarshal([]byte(jsonString), &request)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"LastTokenProcessing", "NetworkStatus.Modifying", "OperationConflict", "SystemBusy", "ResourceStatus.Error", "ServiceUnavailable", "OperationDenied.NetworkAclAttachmentInMiddleStatus", "IncorrectStatus"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_network_acl_attachment", action, AlibabaCloudSdkGoERROR)
	}

	ResourceResourceIdVar, _ := jsonpath.Get("Resource[0].ResourceId", request)
	d.SetId(fmt.Sprintf("%v:%v", request["NetworkAclId"], ResourceResourceIdVar))

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"BINDED"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcServiceV2.VpcNetworkAclAttachmentStateRefreshFunc(d.Id(), "$.Resources.Resource[0].Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudVpcNetworkAclAttachmentRead(d, meta)
}

func resourceAliCloudVpcNetworkAclAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	objectRaw, err := vpcServiceV2.DescribeVpcNetworkAclAttachment(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_network_acl_attachment DescribeVpcNetworkAclAttachment Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("network_acl_id", objectRaw["NetworkAclId"])

	resourceRawObj, _ := jsonpath.Get("$.Resources.Resource[*]", objectRaw)
	resourceRaw := make([]interface{}, 0)
	if resourceRawObj != nil {
		resourceRaw = convertToInterfaceArray(resourceRawObj)
	}

	d.Set("resource_type", resourceRaw["ResourceType"])
	d.Set("status", resourceRaw["Status"])
	d.Set("resource_id", resourceRaw["ResourceId"])

	return nil
}

func resourceAliCloudVpcNetworkAclAttachmentDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "UnassociateNetworkAcl"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["NetworkAclId"] = parts[0]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	resourceDataList := make(map[string]interface{})

	if v, ok := d.GetOk("resource_id"); ok {
		resourceDataList["ResourceId"] = v
	}

	if v, ok := d.GetOk("resource_type"); ok {
		resourceDataList["ResourceType"] = v
	}

	ResourceMap := make([]interface{}, 0)
	ResourceMap = append(ResourceMap, resourceDataList)
	request["Resource"] = ResourceMap

	jsonString := convertObjectToJsonString(request)
	jsonString, _ = sjson.Set(jsonString, "Resource.0.ResourceId", parts[1])
	_ = json.Unmarshal([]byte(jsonString), &request)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Vpc", "2016-04-28", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"LastTokenProcessing", "NetworkStatus.Modifying", "OperationConflict", "SystemBusy", "ResourceStatus.Error", "OperationDenied.NetworkAclAttachmentInMiddleStatus", "IncorrectStatus"}) || NeedRetry(err) {
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

	vpcServiceV2 := VpcServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcServiceV2.VpcNetworkAclAttachmentStateRefreshFunc(d.Id(), "$.Resources.Resource[0].Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
