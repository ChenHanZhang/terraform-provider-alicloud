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

func resourceAliCloudNasAccessPoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNasAccessPointCreate,
		Read:   resourceAliCloudNasAccessPointRead,
		Update: resourceAliCloudNasAccessPointUpdate,
		Delete: resourceAliCloudNasAccessPointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"access_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"access_point_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_point_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled_ram": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"file_system_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"posix_user": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"posix_group_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"posix_user_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"posix_secondary_group_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"root_path_permission": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"owner_user_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"permission": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"owner_group_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"vswitch_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudNasAccessPointCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateAccessPoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("file_system_id"); ok {
		request["FileSystemId"] = v
	}

	if v, ok := d.GetOkExists("enabled_ram"); ok {
		request["EnabledRam"] = v
	}
	if v, ok := d.GetOkExists("posix_user"); ok {
		posixUserPosixGroupIdJsonPath, err := jsonpath.Get("$[0].posix_group_id", v)
		if err == nil && posixUserPosixGroupIdJsonPath != "" {
			request["PosixGroupId"] = posixUserPosixGroupIdJsonPath
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	request["VswId"] = d.Get("vswitch_id")
	if v, ok := d.GetOk("access_point_name"); ok {
		request["AccessPointName"] = v
	}
	if v, ok := d.GetOkExists("root_path_permission"); ok {
		rootPathPermissionOwnerGroupIdJsonPath, err := jsonpath.Get("$[0].owner_group_id", v)
		if err == nil && rootPathPermissionOwnerGroupIdJsonPath != "" {
			request["OwnerGroupId"] = rootPathPermissionOwnerGroupIdJsonPath
		}
	}
	if v, ok := d.GetOkExists("root_path_permission"); ok {
		rootPathPermissionOwnerUserIdJsonPath, err := jsonpath.Get("$[0].owner_user_id", v)
		if err == nil && rootPathPermissionOwnerUserIdJsonPath != "" {
			request["OwnerUserId"] = rootPathPermissionOwnerUserIdJsonPath
		}
	}
	request["VpcId"] = d.Get("vpc_id")
	if v, ok := d.GetOk("root_path_permission"); ok {
		rootPathPermissionPermissionJsonPath, err := jsonpath.Get("$[0].permission", v)
		if err == nil && rootPathPermissionPermissionJsonPath != "" {
			request["Permission"] = rootPathPermissionPermissionJsonPath
		}
	}
	request["AccessGroup"] = d.Get("access_group")
	if v, ok := d.GetOk("root_path"); ok {
		request["RootDirectory"] = v
	}
	if v, ok := d.GetOkExists("posix_user"); ok {
		posixUserPosixUserIdJsonPath, err := jsonpath.Get("$[0].posix_user_id", v)
		if err == nil && posixUserPosixUserIdJsonPath != "" {
			request["PosixUserId"] = posixUserPosixUserIdJsonPath
		}
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_nas_access_point", action, AlibabaCloudSdkGoERROR)
	}

	AccessPointAccessPointIdVar, _ := jsonpath.Get("$.AccessPoint.AccessPointId", response)
	d.SetId(fmt.Sprintf("%v:%v", request["FileSystemId"], AccessPointAccessPointIdVar))

	nasServiceV2 := NasServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"active"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, nasServiceV2.NasAccessPointStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudNasAccessPointRead(d, meta)
}

func resourceAliCloudNasAccessPointRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nasServiceV2 := NasServiceV2{client}

	objectRaw, err := nasServiceV2.DescribeNasAccessPoint(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_nas_access_point DescribeNasAccessPoint Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("access_group", objectRaw["AccessGroup"])
	d.Set("access_point_name", objectRaw["AccessPointName"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("enabled_ram", objectRaw["EnabledRam"])
	d.Set("region_id", objectRaw["RegionId"])
	d.Set("root_path", objectRaw["RootPath"])
	d.Set("status", objectRaw["Status"])
	d.Set("vswitch_id", objectRaw["VSwitchId"])
	d.Set("vpc_id", objectRaw["VpcId"])
	d.Set("access_point_id", objectRaw["AccessPointId"])
	d.Set("file_system_id", objectRaw["FileSystemId"])

	posixUserMaps := make([]map[string]interface{}, 0)
	posixUserMap := make(map[string]interface{})
	posixUserRaw := make(map[string]interface{})
	if objectRaw["PosixUser"] != nil {
		posixUserRaw = objectRaw["PosixUser"].(map[string]interface{})
	}
	if len(posixUserRaw) > 0 {
		posixUserMap["posix_group_id"] = posixUserRaw["PosixGroupId"]
		posixUserMap["posix_user_id"] = posixUserRaw["PosixUserId"]

		posixSecondaryGroupIdsRaw := make([]interface{}, 0)
		if posixUserRaw["PosixSecondaryGroupIds"] != nil {
			posixSecondaryGroupIdsRaw = convertToInterfaceArray(posixUserRaw["PosixSecondaryGroupIds"])
		}

		posixUserMap["posix_secondary_group_ids"] = posixSecondaryGroupIdsRaw
		posixUserMaps = append(posixUserMaps, posixUserMap)
	}
	if err := d.Set("posix_user", posixUserMaps); err != nil {
		return err
	}
	rootPathPermissionMaps := make([]map[string]interface{}, 0)
	rootPathPermissionMap := make(map[string]interface{})
	rootPathPermissionRaw := make(map[string]interface{})
	if objectRaw["RootPathPermission"] != nil {
		rootPathPermissionRaw = objectRaw["RootPathPermission"].(map[string]interface{})
	}
	if len(rootPathPermissionRaw) > 0 {
		rootPathPermissionMap["owner_group_id"] = rootPathPermissionRaw["OwnerGroupId"]
		rootPathPermissionMap["owner_user_id"] = rootPathPermissionRaw["OwnerUserId"]
		rootPathPermissionMap["permission"] = rootPathPermissionRaw["Permission"]

		rootPathPermissionMaps = append(rootPathPermissionMaps, rootPathPermissionMap)
	}
	if err := d.Set("root_path_permission", rootPathPermissionMaps); err != nil {
		return err
	}
	tagsMaps := objectRaw["Tags"]
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudNasAccessPointUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "ModifyAccessPoint"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["AccessPointId"] = parts[1]
	request["FileSystemId"] = parts[0]

	if d.HasChange("enabled_ram") {
		update = true
		request["EnabledRam"] = d.Get("enabled_ram")
	}

	if d.HasChange("access_group") {
		update = true
	}
	request["AccessGroup"] = d.Get("access_group")
	if d.HasChange("access_point_name") {
		update = true
		request["AccessPointName"] = d.Get("access_point_name")
	}

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
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
		nasServiceV2 := NasServiceV2{client}
		if err := nasServiceV2.SetResourceTags(d, "ACCESSPOINT"); err != nil {
			return WrapError(err)
		}
	}
	return resourceAliCloudNasAccessPointRead(d, meta)
}

func resourceAliCloudNasAccessPointDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	parts := strings.Split(d.Id(), ":")
	action := "DeleteAccessPoint"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["AccessPointId"] = parts[1]
	request["FileSystemId"] = parts[0]

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("NAS", "2017-06-26", action, query, request, true)
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

	nasServiceV2 := NasServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, nasServiceV2.NasAccessPointStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
