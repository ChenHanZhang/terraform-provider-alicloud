// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudMessageServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudMessageServiceEndpointCreate,
		Read:   resourceAliCloudMessageServiceEndpointRead,
		Update: resourceAliCloudMessageServiceEndpointUpdate,
		Delete: resourceAliCloudMessageServiceEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cidr_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:     schema.TypeString,
							Required: true,
						},
						"acl_strategy": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: StringInSlice([]string{"allow"}, false),
						},
					},
				},
			},
			"endpoint_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"public"}, false),
			},
		},
	}
}

func resourceAliCloudMessageServiceEndpointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "EnableEndpoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("endpoint_type"); ok {
		request["EndpointType"] = v
	}
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Mns-open", "2022-01-19", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_message_service_endpoint", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["EndpointType"]))

	return resourceAliCloudMessageServiceEndpointUpdate(d, meta)
}

func resourceAliCloudMessageServiceEndpointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	messageServiceServiceV2 := MessageServiceServiceV2{client}

	objectRaw, err := messageServiceServiceV2.DescribeMessageServiceEndpoint(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_message_service_endpoint DescribeMessageServiceEndpoint Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	cidrListRaw := objectRaw["CidrList"]
	cidrListMaps := make([]map[string]interface{}, 0)
	if cidrListRaw != nil {
		for _, cidrListChildRaw := range convertToInterfaceArray(cidrListRaw) {
			cidrListMap := make(map[string]interface{})
			cidrListChildRaw := cidrListChildRaw.(map[string]interface{})
			cidrListMap["acl_strategy"] = cidrListChildRaw["AclStrategy"]
			cidrListMap["cidr"] = cidrListChildRaw["Cidr"]

			cidrListMaps = append(cidrListMaps, cidrListMap)
		}
	}
	if err := d.Set("cidr_list", cidrListMaps); err != nil {
		return err
	}

	d.Set("endpoint_type", d.Id())

	return nil
}

func resourceAliCloudMessageServiceEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}

	if d.HasChange("cidr_list") {
		var err error
		oldEntry, newEntry := d.GetChange("cidr_list")
		removed := oldEntry
		added := newEntry

		if len(added.([]interface{})) > 0 {
			action := "AuthorizeEndpointAcl"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["EndpointType"] = d.Id()
			request["RegionId"] = client.RegionId
			cidrListAclStrategyJsonPath, err := jsonpath.Get("$.acl_strategy", d.Get("cidr_list"))
			if err == nil {
				request["AclStrategy"] = cidrListAclStrategyJsonPath
			}

			cidrListCidrJsonPath, err := jsonpath.Get("$.cidr", d.Get("cidr_list"))
			if err == nil {
				request["CidrList"] = cidrListCidrJsonPath
			}

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Mns-open", "2022-01-19", action, query, request, true)
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

		if len(removed.([]interface{})) > 0 {
			action := "RevokeEndpointAcl"
			request = make(map[string]interface{})
			query = make(map[string]interface{})
			request["EndpointType"] = d.Id()
			request["RegionId"] = client.RegionId
			cidrListAclStrategyJsonPath, err := jsonpath.Get("$.acl_strategy", d.Get("cidr_list"))
			if err == nil {
				request["AclStrategy"] = cidrListAclStrategyJsonPath
			}

			cidrListCidrJsonPath, err := jsonpath.Get("$.cidr", d.Get("cidr_list"))
			if err == nil {
				request["CidrList"] = cidrListCidrJsonPath
			}

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = client.RpcPost("Mns-open", "2022-01-19", action, query, request, true)
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

	}
	return resourceAliCloudMessageServiceEndpointRead(d, meta)
}

func resourceAliCloudMessageServiceEndpointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DisableEndpoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["EndpointType"] = d.Id()
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Mns-open", "2022-01-19", action, query, request, true)
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
