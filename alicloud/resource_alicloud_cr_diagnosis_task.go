// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCrDiagnosisTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCrDiagnosisTaskCreate,
		Read:   resourceAliCloudCrDiagnosisTaskRead,
		Update: resourceAliCloudCrDiagnosisTaskUpdate,
		Delete: resourceAliCloudCrDiagnosisTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"diagnosis_task_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"diagnosis_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"diagnosis_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"BUILD_DIAGNOSIS", "SYNC_DIAGNOSIS", "INSTANCE_DIAGNOSIS", "TAG_LIFECYCLE_ANALYSIS", "PULL_PUSH_DIAGNOSIS"}, false),
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"end_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"extra": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"start_time": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"related_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceAliCloudCrDiagnosisTaskCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDiagnosisTask"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}
	request["RegionId"] = client.RegionId

	diagnosisTarget := make(map[string]interface{})

	if v := d.Get("targets"); !IsNil(v) {
		namespace1, _ := jsonpath.Get("$[0].namespace", v)
		if namespace1 != nil && namespace1 != "" {
			diagnosisTarget["Namespace"] = namespace1
		}
		extra1, _ := jsonpath.Get("$[0].extra", v)
		if extra1 != nil && extra1 != "" {
			diagnosisTarget["Extra"] = extra1
		}
		repository1, _ := jsonpath.Get("$[0].repository", v)
		if repository1 != nil && repository1 != "" {
			diagnosisTarget["Repository"] = repository1
		}
		endTime1, _ := jsonpath.Get("$[0].end_time", v)
		if endTime1 != nil && endTime1 != "" {
			diagnosisTarget["EndTime"] = endTime1
		}
		relatedId1, _ := jsonpath.Get("$[0].related_id", v)
		if relatedId1 != nil && relatedId1 != "" {
			diagnosisTarget["RelatedId"] = relatedId1
		}
		startTime1, _ := jsonpath.Get("$[0].start_time", v)
		if startTime1 != nil && startTime1 != "" {
			diagnosisTarget["StartTime"] = startTime1
		}
		tag1, _ := jsonpath.Get("$[0].tag", v)
		if tag1 != nil && tag1 != "" {
			diagnosisTarget["Tag"] = tag1
		}

		diagnosisTargetJson, err := json.Marshal(diagnosisTarget)
		if err != nil {
			return WrapError(err)
		}
		request["DiagnosisTarget"] = string(diagnosisTargetJson)
	}

	request["DiagnosisType"] = d.Get("diagnosis_type")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("cr", "2018-12-01", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cr_diagnosis_task", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprintf("%v:%v", request["InstanceId"], response["DiagnosisTaskId"]))

	return resourceAliCloudCrDiagnosisTaskRead(d, meta)
}

func resourceAliCloudCrDiagnosisTaskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	crServiceV2 := CrServiceV2{client}

	objectRaw, err := crServiceV2.DescribeCrDiagnosisTask(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cr_diagnosis_task DescribeCrDiagnosisTask Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("diagnosis_time", objectRaw["DiagnosisTime"])
	d.Set("diagnosis_type", objectRaw["DiagnosisType"])
	d.Set("status", objectRaw["TaskStatus"])
	d.Set("diagnosis_task_id", objectRaw["DiagnosisTaskId"])
	d.Set("instance_id", objectRaw["InstanceId"])

	return nil
}

func resourceAliCloudCrDiagnosisTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	parts := strings.Split(d.Id(), ":")
	action := "CancelDiagnosisTask"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = parts[0]
	request["DiagnosisTaskId"] = parts[1]
	request["RegionId"] = client.RegionId
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cr", "2018-12-01", action, query, request, true)
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

	return resourceAliCloudCrDiagnosisTaskRead(d, meta)
}

func resourceAliCloudCrDiagnosisTaskDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Diagnosis Task. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
