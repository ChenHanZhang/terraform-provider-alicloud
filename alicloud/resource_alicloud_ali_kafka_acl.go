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

func resourceAliCloudAliKafkaAcl() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAliKafkaAclCreate,
		Read:   resourceAliCloudAliKafkaAclRead,
		Update: resourceAliCloudAliKafkaAclUpdate,
		Delete: resourceAliCloudAliKafkaAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"acl_operation_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acl_operation_types": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"acl_permission_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"acl_resource_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acl_resource_pattern_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"acl_resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudAliKafkaAclCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAcl"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	if v, ok := d.GetOk("acl_resource_type"); ok {
		request["AclResourceType"] = v
	}
	if v, ok := d.GetOk("acl_resource_pattern_type"); ok {
		request["AclResourcePatternType"] = v
	}
	request["RegionId"] = client.RegionId

	request["AclResourceName"] = d.Get("acl_resource_name")
	request["Username"] = d.Get("username")
	request["AclOperationType"] = d.Get("acl_operation_type")
	if v, ok := d.GetOk("host"); ok {
		request["Host"] = v
	}
	if v, ok := d.GetOk("acl_permission_type"); ok {
		request["AclPermissionType"] = v
	}
	if v, ok := d.GetOk("acl_operation_types"); ok {
		request["AclOperationTypes"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ali_kafka_acl", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v:%v", request["InstanceId"], request["AclResourceType"], request["AclResourcePatternType"]))

	return resourceAliCloudAliKafkaAclRead(d, meta)
}

func resourceAliCloudAliKafkaAclRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	aliKafkaServiceV2 := AliKafkaServiceV2{client}

	objectRaw, err := aliKafkaServiceV2.DescribeAliKafkaAcl(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ali_kafka_acl DescribeAliKafkaAcl Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("acl_operation_type", objectRaw["AclOperationType"])
	d.Set("acl_permission_type", objectRaw["AclPermissionType"])
	d.Set("acl_resource_name", objectRaw["AclResourceName"])
	d.Set("host", objectRaw["Host"])
	d.Set("username", objectRaw["Username"])
	d.Set("acl_resource_pattern_type", objectRaw["AclResourcePatternType"])
	d.Set("acl_resource_type", objectRaw["AclResourceType"])

	parts := strings.Split(d.Id(), ":")
	d.Set("instance_id", parts[0])

	return nil
}

func resourceAliCloudAliKafkaAclUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Cannot update resource Alicloud Resource Acl.")
	return nil
}

func resourceAliCloudAliKafkaAclDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteAcl"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["AclResourceType"] = parts[1]
	request["AclResourcePatternType"] = parts[2]
	request["RegionId"] = client.RegionId

	request["AclResourceName"] = d.Get("acl_resource_name")
	request["Username"] = d.Get("username")
	request["AclOperationType"] = d.Get("acl_operation_type")
	if v, ok := d.GetOk("host"); ok {
		request["Host"] = v
	}
	if v, ok := d.GetOk("acl_permission_type"); ok {
		request["AclPermissionType"] = v
	}
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

	return nil
}
