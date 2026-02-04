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

func resourceAliCloudEcsKeyPair() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudEcsKeyPairCreate,
		Read:   resourceAliCloudEcsKeyPairRead,
		Update: resourceAliCloudEcsKeyPairUpdate,
		Delete: resourceAliCloudEcsKeyPairDelete,
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
			"finger_print": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_pair_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudEcsKeyPairCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	if v, ok := d.GetOk("is_import"); ok && InArray(fmt.Sprint(v), []string{"false"}) {
		action := "CreateKeyPair"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		if v, ok := d.GetOk("key_pair_name"); ok {
			request["KeyPairName"] = v
		}
		request["RegionId"] = client.RegionId

		if v, ok := d.GetOk("resource_group_id"); ok {
			request["ResourceGroupId"] = v
		}
		if v, ok := d.GetOk("tags"); ok {
			tagsMap := ConvertTags(v.(map[string]interface{}))
			request = expandTagsToMap(request, tagsMap)
		}

		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
			return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_key_pair", action, AlibabaCloudSdkGoERROR)
		}

		d.SetId(fmt.Sprint(response["KeyPairName"]))

	}

	if v, ok := d.GetOk("is_import"); ok && InArray(fmt.Sprint(v), []string{"true"}) {
		action := "ImportKeyPair"
		var request map[string]interface{}
		var response map[string]interface{}
		query := make(map[string]interface{})
		var err error
		request = make(map[string]interface{})
		if v, ok := d.GetOk("key_pair_name"); ok {
			request["KeyPairName"] = v
		}
		request["RegionId"] = client.RegionId

		if v, ok := d.GetOk("resource_group_id"); ok {
			request["ResourceGroupId"] = v
		}
		if v, ok := d.GetOk("tags"); ok {
			tagsMap := ConvertTags(v.(map[string]interface{}))
			request = expandTagsToMap(request, tagsMap)
		}

		request["PublicKeyBody"] = d.Get("public_key")
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
			return WrapErrorf(err, DefaultErrorMsg, "alicloud_ecs_key_pair", action, AlibabaCloudSdkGoERROR)
		}

		d.SetId(fmt.Sprint(response["KeyPairName"]))

	}

	return resourceAliCloudEcsKeyPairRead(d, meta)
}

func resourceAliCloudEcsKeyPairRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ecsServiceV2 := EcsServiceV2{client}

	objectRaw, err := ecsServiceV2.DescribeEcsKeyPair(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ecs_key_pair DescribeEcsKeyPair Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreationTime"])
	d.Set("finger_print", objectRaw["KeyPairFingerPrint"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("key_pair_name", objectRaw["KeyPairName"])

	tagsMaps, _ := jsonpath.Get("$.Tags.Tag", objectRaw)
	d.Set("tags", tagsToMap(tagsMaps))

	d.Set("key_pair_name", d.Id())

	return nil
}

func resourceAliCloudEcsKeyPairUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "JoinResourceGroup"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ResourceId"] = d.Id()
	request["RegionId"] = client.RegionId
	if _, ok := d.GetOk("resource_group_id"); ok && d.HasChange("resource_group_id") {
		update = true
		request["ResourceGroupId"] = d.Get("resource_group_id")
	}

	request["ResourceType"] = "keypair"
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
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
		ecsServiceV2 := EcsServiceV2{client}
		if err := ecsServiceV2.SetResourceTags(d, "keypair"); err != nil {
			return WrapError(err)
		}
	}
	return resourceAliCloudEcsKeyPairRead(d, meta)
}

func resourceAliCloudEcsKeyPairDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteKeyPairs"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["KeyPairNames"] = expandSingletonToList(d.Id())
	request["RegionId"] = client.RegionId

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)
		if err != nil {
			if IsExpectedErrors(err, []string{"ServiceUnavailable", "InvalidParameter.KeypairAlreadyAttachedInstance"}) || NeedRetry(err) {
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
