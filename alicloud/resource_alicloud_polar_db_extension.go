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

func resourceAliCloudPolardbExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbExtensionCreate,
		Read:   resourceAliCloudPolardbExtensionRead,
		Update: resourceAliCloudPolardbExtensionUpdate,
		Delete: resourceAliCloudPolardbExtensionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"account_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"extension_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"installed_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requires": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAliCloudPolardbExtensionCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateExtensions"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("extension_name"); ok {
		request["Extensions"] = v
	}
	if v, ok := d.GetOk("account_name"); ok {
		request["AccountName"] = v
	}
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}
	if v, ok := d.GetOk("db_name"); ok {
		request["DBNames"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationDenied.DBClusterStatus"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_polar_db_extension", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v:%v:%v", request["DBClusterId"], request["AccountName"], request["DBNames"], request["Extensions"]))

	polardbServiceV2 := PolardbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"#CHECKSET"}, d.Timeout(schema.TimeoutCreate), 30*time.Second, polardbServiceV2.PolardbExtensionStateRefreshFunc(d.Id(), "#Name", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudPolardbExtensionUpdate(d, meta)
}

func resourceAliCloudPolardbExtensionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbExtension(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polar_db_extension DescribePolardbExtension Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("default_version", objectRaw["DefaultVersion"])
	d.Set("installed_version", objectRaw["InstalledVersion"])
	d.Set("requires", objectRaw["Requires"])
	d.Set("account_name", objectRaw["Owner"])
	d.Set("extension_name", objectRaw["Name"])

	parts := strings.Split(d.Id(), ":")
	d.Set("db_cluster_id", parts[0])
	d.Set("db_name", parts[2])

	return nil
}

func resourceAliCloudPolardbExtensionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}

	polardbServiceV2 := PolardbServiceV2{client}
	objectRaw, _ := polardbServiceV2.DescribePolardbExtension(d.Id())

	if d.HasChange("installed_version") {
		var err error
		target := d.Get("installed_version").(string)

		currentStatus, err := jsonpath.Get("InstalledVersion", objectRaw)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, d.Id(), "InstalledVersion", objectRaw)
		}
		if fmt.Sprint(currentStatus) != target {
			if target == "$.DefaultVersion" {
				parts := strings.Split(d.Id(), ":")
				action := "UpdateExtensions"
				request = make(map[string]interface{})
				query = make(map[string]interface{})
				request["Extensions"] = parts[3]
				request["DBClusterId"] = parts[0]
				request["DBNames"] = parts[2]
				request["RegionId"] = client.RegionId
				request["ClientToken"] = buildClientToken(action)
				wait := incrementalWait(3*time.Second, 5*time.Second)
				err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
					response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
					if err != nil {
						if IsExpectedErrors(err, []string{"OperationDenied.DBClusterStatus"}) || NeedRetry(err) {
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
				polardbServiceV2 := PolardbServiceV2{client}
				stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("installed_version"))}, d.Timeout(schema.TimeoutUpdate), 10*time.Second, polardbServiceV2.PolardbExtensionStateRefreshFunc(d.Id(), "InstalledVersion", []string{}))
				if _, err := stateConf.WaitForState(); err != nil {
					return WrapErrorf(err, IdMsg, d.Id())
				}

			}
		}
	}

	return resourceAliCloudPolardbExtensionRead(d, meta)
}

func resourceAliCloudPolardbExtensionDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteExtensions"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["Extensions"] = parts[3]
	request["DBClusterId"] = parts[0]
	request["DBNames"] = parts[2]
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("polardb", "2017-08-01", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationDenied.DBClusterStatus"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		if IsExpectedErrors(err, []string{"Extension.NotInstall"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	polardbServiceV2 := PolardbServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 10*time.Second, polardbServiceV2.PolardbExtensionStateRefreshFunc(d.Id(), "Name", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
