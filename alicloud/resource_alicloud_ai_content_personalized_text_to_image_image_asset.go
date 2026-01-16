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

func resourceAliCloudAiContentPersonalizedTextToImageImageAsset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudAiContentPersonalizedTextToImageImageAssetCreate,
		Read:   resourceAliCloudAiContentPersonalizedTextToImageImageAssetRead,
		Update: resourceAliCloudAiContentPersonalizedTextToImageImageAssetUpdate,
		Delete: resourceAliCloudAiContentPersonalizedTextToImageImageAssetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"image_number": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"image_url": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"prompt": {
				Type:     schema.TypeString,
				Required: true,
			},
			"seed": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"strength": {
				Type:         schema.TypeFloat,
				Optional:     true,
				ValidateFunc: FloatBetween(0.1, 0.9),
			},
			"train_steps": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudAiContentPersonalizedTextToImageImageAssetCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := fmt.Sprintf("/api/v1/personalizedtxt2img/addPreModelInferenceJob")
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]*string)
	body := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("image_url"); ok {
		imageUrlMapsArray := convertToInterfaceArray(v)

		request["imageUrl"] = imageUrlMapsArray
	}

	if v, ok := d.GetOkExists("train_steps"); ok {
		request["trainSteps"] = v
	}
	request["prompt"] = d.Get("prompt")
	if v, ok := d.GetOkExists("image_number"); ok {
		request["imageNumber"] = v
	}
	if v, ok := d.GetOkExists("seed"); ok {
		request["seed"] = v
	}
	if v, ok := d.GetOk("strength"); ok && v.(float64) > 0 {
		request["strength"] = v
	}
	body = request
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RoaPost("AiContent", "20240611", action, query, nil, body, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ai_content_personalized_text_to_image_image_asset", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.data.id", response)
	d.SetId(fmt.Sprint(id))

	aiContentServiceV2 := AiContentServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"FINISHED"}, d.Timeout(schema.TimeoutCreate), 60*time.Second, aiContentServiceV2.AiContentPersonalizedTextToImageImageAssetStateRefreshFunc(d.Id(), "jobStatus", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudAiContentPersonalizedTextToImageImageAssetRead(d, meta)
}

func resourceAliCloudAiContentPersonalizedTextToImageImageAssetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	aiContentServiceV2 := AiContentServiceV2{client}

	objectRaw, err := aiContentServiceV2.DescribeAiContentPersonalizedTextToImageImageAsset(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ai_content_personalized_text_to_image_image_asset DescribeAiContentPersonalizedTextToImageImageAsset Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("status", objectRaw["jobStatus"])

	return nil
}

func resourceAliCloudAiContentPersonalizedTextToImageImageAssetUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Cannot update resource Alicloud Resource Personalized Text To Image Image Asset.")
	return nil
}

func resourceAliCloudAiContentPersonalizedTextToImageImageAssetDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Personalized Text To Image Image Asset. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
