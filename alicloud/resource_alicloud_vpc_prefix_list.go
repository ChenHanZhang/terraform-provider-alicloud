// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudVpcPrefixList() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVpcPrefixListCreate,
		Read:   resourceAlicloudVpcPrefixListRead,
		Update: resourceAlicloudVpcPrefixListUpdate,
		Delete: resourceAlicloudVpcPrefixListDelete,
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
			"entries": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cidr": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"max_entries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"owner_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_list_association": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix_list_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_uid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reason": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"prefix_list_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefix_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix_list_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"share_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"entrys": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.All(validation.StringLenBetween(2, 256), validation.StringDoesNotMatch(regexp.MustCompile(`(^http://.*)|(^https://.*)`), "It cannot begin with \"http://\", \"https://\".")),
						},
					},
				},
			},
		},
	}
}

func resourceAlicloudVpcPrefixListCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "CreateVpcPrefixList"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("ip_version"); ok {
		request["IpVersion"] = v
	}

	if v, ok := d.GetOk("max_entries"); ok {
		request["MaxEntries"] = v
	}

	if v, ok := d.GetOk("prefix_list_description"); ok {
		request["PrefixListDescription"] = v
	}

	if v, ok := d.GetOk("prefix_list_name"); ok {
		request["PrefixListName"] = v
	}

	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}

	if v, ok := d.GetOk("entries"); ok {
		localData := v
		prefixListEntriesMaps := make([]map[string]interface{}, 0)
		for _, dataLoop := range localData.(*schema.Set).List() {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Cidr"] = dataLoopTmp["cidr"]
			dataLoopMap["Description"] = dataLoopTmp["description"]
			prefixListEntriesMaps = append(prefixListEntriesMaps, dataLoopMap)
		}
		request["PrefixListEntries"] = prefixListEntriesMaps
	}

	if v, ok := d.GetOk("entrys"); ok {
		for entrysPtr, entrys := range v.(*schema.Set).List() {
			entrysArg := entrys.(map[string]interface{})
			request["PrefixListEntries."+fmt.Sprint(entrysPtr+1)+".Cidr"] = entrysArg["cidr"]
			request["PrefixListEntries."+fmt.Sprint(entrysPtr+1)+".Description"] = entrysArg["description"]
		}
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "SystemBusy", "IncorrectStatus.%s"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vpc_prefix_list", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["PrefixListId"]))

	return resourceAlicloudVpcPrefixListUpdate(d, meta)
}

func resourceAlicloudVpcPrefixListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcServiceV2 := VpcServiceV2{client}

	object, err := vpcServiceV2.DescribeVpcPrefixList(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vpc_prefix_list .DescribeVpcPrefixList Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", object["create_time"])
	d.Set("entries", object["entries"])
	d.Set("ip_version", object["ip_version"])
	d.Set("max_entries", object["max_entries"])
	d.Set("owner_id", object["owner_id"])
	d.Set("prefix_list_association", object["prefix_list_association"])
	d.Set("prefix_list_description", object["prefix_list_description"])
	d.Set("prefix_list_id", object["prefix_list_id"])
	d.Set("prefix_list_name", object["prefix_list_name"])
	d.Set("resource_group_id", object["resource_group_id"])
	d.Set("share_type", object["share_type"])
	d.Set("status", object["status"])
	d.Set("tags", tagsToMap(object["tags"]))

	d.Set("entrys", d.Get("entries"))
	return nil
}

func resourceAlicloudVpcPrefixListUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	update := false
	d.Partial(true)
	update = false
	action := "ModifyVpcPrefixList"
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["PrefixListId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if !d.IsNewResource() && d.HasChange("prefix_list_name") {
		update = true
		if v, ok := d.GetOk("prefix_list_name"); ok {
			request["PrefixListName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("prefix_list_description") {
		update = true
		if v, ok := d.GetOk("prefix_list_description"); ok {
			request["PrefixListDescription"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("max_entries") {
		update = true
		if v, ok := d.GetOk("max_entries"); ok {
			request["MaxEntries"] = v
		}
	}
	if d.HasChange("entrys") {
		update = true
		oldEntry, newEntry := d.GetChange("entrys")
		oldEntrySet := oldEntry.(*schema.Set)
		newEntrySet := newEntry.(*schema.Set)
		removed := oldEntrySet.Difference(newEntrySet)
		added := newEntrySet.Difference(oldEntrySet)

		for entrysPtr, entrys := range removed.List() {
			entrysArg := entrys.(map[string]interface{})
			request["RemovePrefixListEntry."+fmt.Sprint(entrysPtr+1)+".Cidr"] = entrysArg["cidr"]
			request["RemovePrefixListEntry."+fmt.Sprint(entrysPtr+1)+".Description"] = entrysArg["description"]
		}

		for entrysPtr, entrys := range added.List() {
			entrysArg := entrys.(map[string]interface{})
			request["AddPrefixListEntry."+fmt.Sprint(entrysPtr+1)+".Cidr"] = entrysArg["cidr"]
			request["AddPrefixListEntry."+fmt.Sprint(entrysPtr+1)+".Description"] = entrysArg["description"]
		}
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("prefix_list_name")
		d.SetPartial("prefix_list_description")
		d.SetPartial("max_entries")
	}
	update = false
	action = "MoveResourceGroup"
	conn, err = client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId

	if !d.IsNewResource() && d.HasChange("resource_group_id") {
		update = true
		if v, ok := d.GetOk("resource_group_id"); ok {
			request["NewResourceGroupId"] = v
		}
	}
	request["ResourceType"] = "PrefixList"

	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("resource_group_id")
	}

	update = false
	if !d.IsNewResource() && d.HasChange("entries") {
		update = true
		oldEntry, newEntry := d.GetChange("entries")
		oldEntrySet := oldEntry.(*schema.Set)
		newEntrySet := newEntry.(*schema.Set)
		removed := oldEntrySet.Difference(newEntrySet)
		added := newEntrySet.Difference(oldEntrySet)

		if removed.Len() > 0 {
			action = "ModifyVpcPrefixList"
			conn, err = client.NewVpcClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["PrefixListId"] = d.Id()
			request["RegionId"] = client.RegionId
			request["ClientToken"] = buildClientToken(action)
			localData := removed.List()

			removePrefixListEntryMaps := make([]map[string]interface{}, 0)
			for _, dataLoop := range localData {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["Cidr"] = dataLoopTmp["cidr"]
				dataLoopMap["Description"] = dataLoopTmp["description"]
				removePrefixListEntryMaps = append(removePrefixListEntryMaps, dataLoopMap)
			}
			request["RemovePrefixListEntry"] = removePrefixListEntryMaps

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
				if err != nil {
					if IsExpectedErrors(err, []string{"IncorrectStatus.PrefixList", "IncorrectStatus.%s", "SystemBusy", "LastTokenProcessing"}) || NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}

		if added.Len() > 0 {
			action = "ModifyVpcPrefixList"
			conn, err = client.NewVpcClient()
			if err != nil {
				return WrapError(err)
			}
			request = make(map[string]interface{})

			request["PrefixListId"] = d.Id()
			request["RegionId"] = client.RegionId
			request["ClientToken"] = buildClientToken(action)
			localData := added.List()

			addPrefixListEntryMaps := make([]map[string]interface{}, 0)
			for _, dataLoop := range localData {
				dataLoopTmp := dataLoop.(map[string]interface{})
				dataLoopMap := make(map[string]interface{})
				dataLoopMap["Cidr"] = dataLoopTmp["cidr"]
				dataLoopMap["Description"] = dataLoopTmp["description"]
				addPrefixListEntryMaps = append(addPrefixListEntryMaps, dataLoopMap)
			}
			request["AddPrefixListEntry"] = addPrefixListEntryMaps

			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
				response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
				if err != nil {
					if IsExpectedErrors(err, []string{"IncorrectStatus.PrefixList", "IncorrectStatus.%s", "SystemBusy", "LastTokenProcessing"}) || NeedRetry(err) {
						wait()
						return resource.RetryableError(err)
					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}

		}

	}
	update = false
	if d.HasChange("tags") {
		update = true
		vpcServiceV2 := VpcServiceV2{client}
		if err := vpcServiceV2.SetResourceTags(d, "PrefixList"); err != nil {
			return WrapError(err)
		}
	}
	d.Partial(false)
	return resourceAlicloudVpcPrefixListRead(d, meta)
}

func resourceAlicloudVpcPrefixListDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "DeleteVpcPrefixList"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	request["PrefixListId"] = d.Id()
	request["RegionId"] = client.RegionId

	request["ClientToken"] = buildClientToken(action)

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"OperationConflict", "SystemBusy", "DependencyViolation.ShareResource", "IncorrectStatus.PrefixList", "IncorrectStatus.SystemPrefixList", "IncorrectStatus.%s", "OperationFailed.LastTokenProcessing", "LastTokenProcessing"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
