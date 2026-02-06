// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudAdbLakeStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAdbLakeStorageCreate,
		Read:   resourceAliCloudAdbLakeStorageRead,
		Update: resourceAliCloudAdbLakeStorageUpdate,
		Delete: resourceAliCloudAdbLakeStorageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"lake_storage_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permissions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"read": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"write": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"account": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudAdbLakeStorageCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateLakeStorage"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("db_cluster_id"); ok {
		request["DBClusterId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOk("permissions"); ok {
		permissionsMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Write"] = dataLoopTmp["write"]
			dataLoopMap["Type"] = dataLoopTmp["type"]
			dataLoopMap["Account"] = dataLoopTmp["account"]
			dataLoopMap["Read"] = dataLoopTmp["read"]
			permissionsMapsArray = append(permissionsMapsArray, dataLoopMap)
		}
		permissionsMapsJson, err := json.Marshal(permissionsMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["Permissions"] = string(permissionsMapsJson)
	}

	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_adb_lake_storage", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["DBClusterId"], response["Data"]))

	return resourceAliCloudAdbLakeStorageRead(d, meta)
}

func resourceAliCloudAdbLakeStorageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	adbServiceV2 := AdbServiceV2{client}

	objectRaw, err := adbServiceV2.DescribeAdbLakeStorage(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_adb_lake_storage DescribeAdbLakeStorage Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("description", objectRaw["Description"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])
	d.Set("lake_storage_id", objectRaw["LakeStorageId"])

	permissionsRaw := objectRaw["Permissions"]
	permissionsMaps := make([]map[string]interface{}, 0)
	if permissionsRaw != nil {
		for _, permissionsChildRaw := range convertToInterfaceArray(permissionsRaw) {
			permissionsMap := make(map[string]interface{})
			permissionsChildRaw := permissionsChildRaw.(map[string]interface{})
			permissionsMap["account"] = permissionsChildRaw["Account"]
			permissionsMap["read"] = permissionsChildRaw["Read"]
			permissionsMap["type"] = permissionsChildRaw["Type"]
			permissionsMap["write"] = permissionsChildRaw["Write"]

			permissionsMaps = append(permissionsMaps, permissionsMap)
		}
	}
	if err := d.Set("permissions", permissionsMaps); err != nil {
		return err
	}

	return nil
}

func resourceAliCloudAdbLakeStorageUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "UpdateLakeStorage"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["LakeStorageId"] = parts[1]
	request["DBClusterId"] = parts[0]
	request["RegionId"] = client.RegionId
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if d.HasChange("permissions") {
		update = true
		if v, ok := d.GetOk("permissions"); ok || d.HasChange("permissions") {
			permissionsMapsArray := make([]interface{}, 0)
			for _, dataLoop := range convertToInterfaceArray(v) {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["Write"] = dataLoopTmp["write"]
				dataLoopMap["Type"] = dataLoopTmp["type"]
				dataLoopMap["Account"] = dataLoopTmp["account"]
				dataLoopMap["Read"] = dataLoopTmp["read"]
				permissionsMapsArray = append(permissionsMapsArray, dataLoopMap)
			}
			permissionsMapsJson, err := json.Marshal(permissionsMapsArray)
			if err != nil {
				return WrapError(err)
			}
			request["Permissions"] = string(permissionsMapsJson)
		}
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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

	return resourceAliCloudAdbLakeStorageRead(d, meta)
}

func resourceAliCloudAdbLakeStorageDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteLakeStorage"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["DBClusterId"] = parts[0]
	request["LakeStorageId"] = parts[1]
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("adb", "2021-12-01", action, query, request, true)
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
		if IsExpectedErrors(err, []string{"400"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
