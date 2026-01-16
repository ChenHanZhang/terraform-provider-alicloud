// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudNasAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNasAccessRuleCreate,
		Read:   resourceAliCloudNasAccessRuleRead,
		Update: resourceAliCloudNasAccessRuleUpdate,
		Delete: resourceAliCloudNasAccessRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"access_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_rule_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_system_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"standard", "extreme"}, false),
			},
			"ipv6_source_cidr_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: IntBetween(0, 100),
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rw_access": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"RDWR", "RDONLY"}, false),
			},
			"source_cidr_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_access": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"no_squash", "root_squash", "all_squash"}, false),
			},
		},
	}
}

func resourceAliCloudNasAccessRuleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAccessRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("access_group_name"); ok {
		request["AccessGroupName"] = v
	}
	if v, ok := d.GetOk("file_system_type"); ok {
		request["FileSystemType"] = v
	}

	if v, ok := d.GetOk("rw_access"); ok {
		request["RWAccessType"] = v
	}
	if v, ok := d.GetOk("user_access"); ok {
		request["UserAccessType"] = v
	}
	if v, ok := d.GetOk("ipv6_source_cidr_ip"); ok {
		request["Ipv6SourceCidrIp"] = v
	}
	if v, ok := d.GetOkExists("priority"); ok && v.(int) > 0 {
		request["Priority"] = v
	}
	if v, ok := d.GetOk("source_cidr_ip"); ok {
		request["SourceCidrIp"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationDenied.InvalidState"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_nas_access_rule", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v:%v", request["AccessGroupName"], request["FileSystemType"], response["AccessRuleId"]))

	return resourceAliCloudNasAccessRuleRead(d, meta)
}

func resourceAliCloudNasAccessRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nasServiceV2 := NasServiceV2{client}

	objectRaw, err := nasServiceV2.DescribeNasAccessRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_nas_access_rule DescribeNasAccessRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("ipv6_source_cidr_ip", objectRaw["Ipv6SourceCidrIp"])
	d.Set("priority", objectRaw["Priority"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("rw_access", objectRaw["RWAccess"])
	d.Set("source_cidr_ip", objectRaw["SourceCidrIp"])
	d.Set("user_access", objectRaw["UserAccess"])
	d.Set("access_group_name", objectRaw["AccessGroupName"])
	d.Set("access_rule_id", objectRaw["AccessRuleId"])
	d.Set("file_system_type", objectRaw["FileSystemType"])

	return nil
}

func resourceAliCloudNasAccessRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyAccessRule"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AccessGroupName"] = parts[0]
	request["AccessRuleId"] = parts[2]
	request["FileSystemType"] = parts[1]

	if d.HasChange("rw_access") {
		update = true
		request["RWAccessType"] = d.Get("rw_access")
	}

	if d.HasChange("user_access") {
		update = true
		request["UserAccessType"] = d.Get("user_access")
	}

	if d.HasChange("ipv6_source_cidr_ip") {
		update = true
		request["Ipv6SourceCidrIp"] = d.Get("ipv6_source_cidr_ip")
	}

	if d.HasChange("priority") {
		update = true
		request["Priority"] = d.Get("priority")
	}

	if d.HasChange("source_cidr_ip") {
		update = true
		request["SourceCidrIp"] = d.Get("source_cidr_ip")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"OperationDenied.InvalidState"}) || NeedRetry(err) {
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

	return resourceAliCloudNasAccessRuleRead(d, meta)
}

func resourceAliCloudNasAccessRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteAccessRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AccessGroupName"] = parts[0]
	request["AccessRuleId"] = parts[2]
	request["FileSystemType"] = parts[1]

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationDenied.InvalidState"}) || NeedRetry(err) {
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
