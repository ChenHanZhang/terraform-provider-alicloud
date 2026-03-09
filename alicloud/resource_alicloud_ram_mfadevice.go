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

func resourceAliCloudRamMFADevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRamMFADeviceCreate,
		Read:   resourceAliCloudRamMFADeviceRead,
		Update: resourceAliCloudRamMFADeviceUpdate,
		Delete: resourceAliCloudRamMFADeviceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"activate_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_mfa_device_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudRamMFADeviceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateVirtualMFADevice"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["VirtualMFADeviceName"] = d.Get("virtual_mfa_device_name")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Ims", "2019-08-15", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ram_mfadevice", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.VirtualMFADevice.SerialNumber", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudRamMFADeviceRead(d, meta)
}

func resourceAliCloudRamMFADeviceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ramServiceV2 := RamServiceV2{client}

	objectRaw, err := ramServiceV2.DescribeRamMFADevice(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ram_mfadevice DescribeRamMFADevice Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("activate_date", objectRaw["ActivateDate"])

	return nil
}

func resourceAliCloudRamMFADeviceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Cannot update resource Alicloud Resource M F A Device.")
	return nil
}

func resourceAliCloudRamMFADeviceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteVirtualMFADevice"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["SerialNumber"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ims", "2019-08-15", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"EntityNotExist.VirtualMFADevice"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
