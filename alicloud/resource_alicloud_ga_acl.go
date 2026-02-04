// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudGaAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudGaAclCreate,
		Read:   resourceAliCloudGaAclRead,
		Update: resourceAliCloudGaAclUpdate,
		Delete: resourceAliCloudGaAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(8 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"acl_entries": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"entry": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"entry_description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
			},
			"acl_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_ip_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
		},
	}
}

func resourceAliCloudGaAclCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAcl"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("acl_entries"); ok {
		aclEntriesMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Entry"] = dataLoopTmp["entry"]
			dataLoopMap["EntryDescription"] = dataLoopTmp["entry_description"]
			aclEntriesMapsArray = append(aclEntriesMapsArray, dataLoopMap)
		}
		request["AclEntries"] = aclEntriesMapsArray
	}

	request["AddressIPVersion"] = d.Get("address_ip_version")
	if v, ok := d.GetOk("acl_name"); ok {
		request["AclName"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ga_acl", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["AclId"]))

	gaServiceV2 := GaServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"active"}, d.Timeout(schema.TimeoutCreate), 50*time.Second, gaServiceV2.GaAclStateRefreshFunc(d.Id(), "AclStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudGaAclRead(d, meta)
}

func resourceAliCloudGaAclRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	gaServiceV2 := GaServiceV2{client}

	objectRaw, err := gaServiceV2.DescribeGaAcl(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ga_acl DescribeGaAcl Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("acl_name", objectRaw["AclName"])
	d.Set("address_ip_version", objectRaw["AddressIPVersion"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("status", objectRaw["AclStatus"])

	aclEntriesRaw := objectRaw["AclEntries"]
	aclEntriesMaps := make([]map[string]interface{}, 0)
	if aclEntriesRaw != nil {
		for _, aclEntriesChildRaw := range convertToInterfaceArray(aclEntriesRaw) {
			aclEntriesMap := make(map[string]interface{})
			aclEntriesChildRaw := aclEntriesChildRaw.(map[string]interface{})
			aclEntriesMap["entry"] = aclEntriesChildRaw["Entry"]
			aclEntriesMap["entry_description"] = aclEntriesChildRaw["EntryDescription"]

			aclEntriesMaps = append(aclEntriesMaps, aclEntriesMap)
		}
	}
	if err := d.Set("acl_entries", aclEntriesMaps); err != nil {
		return err
	}
	tagsMaps := objectRaw["Tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudGaAclUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)

	var err error
	action := "UpdateAclAttribute"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AclId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("acl_name") {
		update = true
	}
	request["AclName"] = d.Get("acl_name")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"StateError.Acl"}) || NeedRetry(err) {
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
		gaServiceV2 := GaServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"active"}, d.Timeout(schema.TimeoutUpdate), 50*time.Second, gaServiceV2.GaAclStateRefreshFunc(d.Id(), "AclStatus", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}
	update = false
	action = "ChangeResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if _, ok := d.GetOk("resource_group_id"); ok && d.HasChange("resource_group_id") {
		update = true
	}
	request["NewResourceGroupId"] = d.Get("resource_group_id")
	request["ResourceType"] = "acl"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
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
		gaServiceV2 := GaServiceV2{client}
		if err := gaServiceV2.SetResourceTags(d, "acl"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAliCloudGaAclRead(d, meta)
}

func resourceAliCloudGaAclDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAcl"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AclId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ga", "2019-11-20", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"StateError.Acl"}) || NeedRetry(err) {
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

	gaServiceV2 := GaServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 50*time.Second, gaServiceV2.GaAclStateRefreshFunc(d.Id(), "$.AclId", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
