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

func resourceAliCloudCloudSsoAccessConfigurationProvisioning() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudSsoAccessConfigurationProvisioningCreate,
		Read:   resourceAliCloudCloudSsoAccessConfigurationProvisioningRead,
		Delete: resourceAliCloudCloudSsoAccessConfigurationProvisioningDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"access_configuration_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"directory_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"RD-Account"}, false),
			},
		},
	}
}

func resourceAliCloudCloudSsoAccessConfigurationProvisioningCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "ProvisionAccessConfiguration"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("target_type"); ok {
		request["TargetType"] = v
	}
	if v, ok := d.GetOk("target_id"); ok {
		request["TargetId"] = v
	}
	if v, ok := d.GetOk("directory_id"); ok {
		request["DirectoryId"] = v
	}
	if v, ok := d.GetOk("access_configuration_id"); ok {
		request["AccessConfigurationId"] = v
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("cloudsso", "2021-05-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cloud_sso_access_configuration_provisioning", action, AlibabaCloudSdkGoERROR)
	}

	TasksAccessConfigurationIdVar, _ := jsonpath.Get("$.Tasks[0].AccessConfigurationId", response)
	TasksTargetTypeVar, _ := jsonpath.Get("$.Tasks[0].TargetType", response)
	TasksTargetIdVar, _ := jsonpath.Get("$.Tasks[0].TargetId", response)
	d.SetId(fmt.Sprintf("%v:%v:%v:%v", request["DirectoryId"], TasksAccessConfigurationIdVar, TasksTargetTypeVar, TasksTargetIdVar))

	cloudSSOServiceV2 := CloudSSOServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Success"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, cloudSSOServiceV2.DescribeAsyncCloudSsoAccessConfigurationProvisioningStateRefreshFunc(d, response, "TaskStatus.Status", []string{}))
	if jobDetail, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id(), jobDetail)
	}

	return resourceAliCloudCloudSsoAccessConfigurationProvisioningRead(d, meta)
}

func resourceAliCloudCloudSsoAccessConfigurationProvisioningRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudSSOServiceV2 := CloudSSOServiceV2{client}

	objectRaw, err := cloudSSOServiceV2.DescribeCloudSsoAccessConfigurationProvisioning(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cloud_sso_access_configuration_provisioning DescribeCloudSsoAccessConfigurationProvisioning Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("status", objectRaw["Status"])
	d.Set("access_configuration_id", objectRaw["AccessConfigurationId"])
	d.Set("target_id", objectRaw["TargetId"])
	d.Set("target_type", objectRaw["TargetType"])

	parts := strings.Split(d.Id(), ":")
	d.Set("directory_id", parts[0])

	return nil
}

func resourceAliCloudCloudSsoAccessConfigurationProvisioningDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeprovisionAccessConfiguration"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AccessConfigurationId"] = parts[1]
	request["TargetType"] = parts[2]
	request["TargetId"] = parts[3]
	request["DirectoryId"] = parts[0]

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("cloudsso", "2021-05-15", action, query, request, true)
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

	cloudSSOServiceV2 := CloudSSOServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Success"}, d.Timeout(schema.TimeoutDelete), 5*time.Second, cloudSSOServiceV2.DescribeAsyncCloudSsoAccessConfigurationProvisioningStateRefreshFunc(d, response, "TaskStatus.Status", []string{}))
	if jobDetail, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id(), jobDetail)
	}

	return nil
}
