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

func resourceAliCloudSslCertificatesServiceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudSslCertificatesServiceInstanceCreate,
		Read:   resourceAliCloudSslCertificatesServiceInstanceRead,
		Update: resourceAliCloudSslCertificatesServiceInstanceUpdate,
		Delete: resourceAliCloudSslCertificatesServiceInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_reissue": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"enable", "disable"}, false),
			},
			"brand": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"certificate_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certificate_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"company_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_id_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"country_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"csr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"generate_csr_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"online", "upload"}, false),
			},
			"instance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"code": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pricing_cycle": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"product_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"cas_dv_public_cn", "cas_intl"}, false),
			},
			"province": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": tagsSchema(),
			"validation_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: StringInSlice([]string{"DNS", "HTTP"}, false),
			},
		},
	}
}

func resourceAliCloudSslCertificatesServiceInstanceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	request["ProductCode"] = "cas"
	request["SubscriptionType"] = "Subscription"
	if v, ok := d.GetOk("product_type"); ok {
		request["ProductType"] = v
	}
	parameterDataList := make(map[string]interface{})
	if v, ok := d.GetOk("parameter"); ok {
		value1, _ := jsonpath.Get("$.value", v)
		if value1 != nil && value1 != "" {
			parameterDataList["Value"] = value1
		}
	}
	if v, ok := d.GetOk("parameter"); ok {
		code1, _ := jsonpath.Get("$.code", v)
		if code1 != nil && code1 != "" {
			parameterDataList["Code"] = code1
		}
	}

	ParameterMap := make([]interface{}, 0)
	ParameterMap = append(ParameterMap, parameterDataList)
	request["Parameter"] = ParameterMap

	if v, ok := d.GetOkExists("period"); ok {
		request["Period"] = v
	}
	if v, ok := d.GetOkExists("pricing_cycle"); ok {
		request["PricingCycle"] = v
	}
	var endpoint string
	request["ProductCode"] = ""
	request["ProductType"] = ""
	if client.IsInternationalAccount() {
		request["ProductType"] = ""
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPostWithEndpoint("BssOpenApi", "2017-12-14", action, query, request, true, endpoint)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			if !client.IsInternationalAccount() && IsExpectedErrors(err, []string{""}) {
				request["ProductCode"] = ""
				request["ProductType"] = ""
				endpoint = connectivity.BssOpenAPIEndpointInternational
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ssl_certificates_service_instance", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.Data.InstanceId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudSslCertificatesServiceInstanceRead(d, meta)
}

func resourceAliCloudSslCertificatesServiceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	sslCertificatesServiceServiceV2 := SslCertificatesServiceServiceV2{client}

	objectRaw, err := sslCertificatesServiceServiceV2.DescribeSslCertificatesServiceInstance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ssl_certificates_service_instance DescribeSslCertificatesServiceInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("brand", objectRaw["Brand"])
	d.Set("certificate_status", objectRaw["CertificateStatus"])
	d.Set("certificate_type", objectRaw["CertificateType"])
	d.Set("instance_type", objectRaw["InstanceType"])
	d.Set("status", objectRaw["Status"])

	return nil
}

func resourceAliCloudSslCertificatesServiceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOk("key_algorithm"); ok {
		request["KeyAlgorithm"] = v
	}
	if v, ok := d.GetOk("country_code"); ok {
		request["CountryCode"] = v
	}
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	if v, ok := d.GetOk("city"); ok {
		request["City"] = v
	}
	if v, ok := d.GetOk("province"); ok {
		request["Province"] = v
	}
	if v, ok := d.GetOk("generate_csr_method"); ok {
		request["GenerateCsrMethod"] = v
	}
	if d.HasChange("contact_id_list") {
		update = true
		if v, ok := d.GetOk("contact_id_list"); ok || d.HasChange("contact_id_list") {
			contactIdListMapsArray := convertToInterfaceArray(v)

			request["ContactIdList"] = contactIdListMapsArray
		}
	}

	tagsDataList := make(map[string]interface{})
	tagsMap := ConvertTags(v.(map[string]interface{}))
	tagsDataList["TagValue"] = tagsMap
	tagsMap := ConvertTags(v.(map[string]interface{}))
	tagsDataList["TagKey"] = tagsMap

	TagsMap := make([]interface{}, 0)
	TagsMap = append(TagsMap, tagsDataList)
	request["Tags"] = TagsMap

	if v, ok := d.GetOk("csr"); ok {
		request["Csr"] = v
	}
	if v, ok := d.GetOk("certificate_name"); ok {
		request["CertificateName"] = v
	}
	if v, ok := d.GetOk("domain"); ok {
		request["Domain"] = v
	}
	if v, ok := d.GetOk("auto_reissue"); ok {
		request["AutoReissue"] = v
	}
	if v, ok := d.GetOk("company_id"); ok {
		request["CompanyId"] = v
	}
	if v, ok := d.GetOk("validation_method"); ok {
		request["ValidationMethod"] = v
	}
	if update {
		wait := incrementalWait(5*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cas", "2020-04-07", action, query, request, true)
			if err != nil {
				if IsExpectedErrors(err, []string{"InvalidStatus.UpdateProtection"}) || NeedRetry(err) {
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
	update = false
	action = "ApplyCertificate"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cas", "2020-04-07", action, query, request, true)
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
	update = false
	action = "RevokeCertificate"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if v, ok := d.GetOkExists("certificate_id"); ok {
		request["CertificateId"] = v
	}
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cas", "2020-04-07", action, query, request, true)
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
	update = false
	action = "RefundInstance"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cas", "2020-04-07", action, query, request, true)
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
	update = false
	action = "CancelPendingCertificate"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("cas", "2020-04-07", action, query, request, true)
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

	return resourceAliCloudSslCertificatesServiceInstanceRead(d, meta)
}

func resourceAliCloudSslCertificatesServiceInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteInstance"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["InstanceId"] = d.Id()

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("cas", "2020-04-07", action, query, request, true)
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
