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

func resourceAliCloudComputeNestService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudComputeNestServiceCreate,
		Read:   resourceAliCloudComputeNestServiceRead,
		Update: resourceAliCloudComputeNestServiceUpdate,
		Delete: resourceAliCloudComputeNestServiceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alarm_metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"approval_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deploy_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_metadata": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_private_vpc_connection": {
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"template_configs": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hidden_parameter_keys": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"allowed_regions": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"url": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"predefined_parameters": {
										Type:     schema.TypeList,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"update_info": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"parameters_cause_interruption_if_modified": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"parameters_not_allowed_to_be_modified": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"parameters_conditionally_cause_interruption_if_modified": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"parameters_uncertainly_allowed_to_be_modified": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"parameters_allowed_to_be_modified": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
												"parameters_uncertainly_cause_interruption_if_modified": {
													Type:     schema.TypeList,
													Optional: true,
													Elem:     &schema.Schema{Type: schema.TypeString},
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"supplier_deploy_metadata": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"deploy_timeout": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"deploy_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_support_operated": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"license_metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operation_metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_names": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"service_info": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"locale": {
							Type:     schema.TypeString,
							Required: true,
						},
						"long_description_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"softwares": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"agreements": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"url": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"image": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"short_description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"service_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"share_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trial_duration": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upgrade_metadata": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"version_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudComputeNestServiceCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateService"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	if v, ok := d.GetOk("service_id"); ok {
		request["ServiceId"] = v
	}
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	request["DeployType"] = d.Get("deploy_type")
	if v, ok := d.GetOk("resource_group_id"); ok {
		request["ResourceGroupId"] = v
	}
	request["ServiceType"] = d.Get("service_type")
	if v, ok := d.GetOk("service_info"); ok {
		serviceInfoMapsArray := make([]interface{}, 0)
		for _, dataLoop := range convertToInterfaceArray(v) {
			dataLoopTmp := dataLoop.(map[string]interface{})
			dataLoopMap := make(map[string]interface{})
			dataLoopMap["Locale"] = dataLoopTmp["locale"]
			dataLoopMap["ShortDescription"] = dataLoopTmp["short_description"]
			dataLoopMap["Image"] = dataLoopTmp["image"]
			dataLoopMap["Name"] = dataLoopTmp["name"]
			serviceInfoMapsArray = append(serviceInfoMapsArray, dataLoopMap)
		}
		request["ServiceInfo"] = serviceInfoMapsArray
	}

	if v, ok := d.GetOkExists("is_support_operated"); ok {
		request["IsSupportOperated"] = v
	}
	if v, ok := d.GetOk("policy_names"); ok {
		request["PolicyNames"] = v
	}
	if v, ok := d.GetOkExists("duration"); ok {
		request["Duration"] = v
	}
	if v, ok := d.GetOk("alarm_metadata"); ok {
		request["AlarmMetadata"] = v
	}
	if v, ok := d.GetOk("share_type"); ok {
		request["ShareType"] = v
	}
	if v, ok := d.GetOk("approval_type"); ok {
		request["ApprovalType"] = v
	}
	if v, ok := d.GetOk("version_name"); ok {
		request["VersionName"] = v
	}
	if v, ok := d.GetOk("upgrade_metadata"); ok {
		request["UpgradeMetadata"] = v
	}
	if v, ok := d.GetOkExists("trial_duration"); ok {
		request["TrialDuration"] = v
	}
	if v, ok := d.GetOk("tenant_type"); ok {
		request["TenantType"] = v
	}
	if v, ok := d.GetOk("license_metadata"); ok {
		request["LicenseMetadata"] = v
	}
	if v, ok := d.GetOk("operation_metadata"); ok {
		request["OperationMetadata"] = v
	}
	deployMetadata := make(map[string]interface{})

	if v := d.Get("deploy_metadata"); !IsNil(v) {
		supplierDeployMetadata := make(map[string]interface{})
		deployTimeout1, _ := jsonpath.Get("$[0].supplier_deploy_metadata[0].deploy_timeout", d.Get("deploy_metadata"))
		if deployTimeout1 != nil && deployTimeout1 != "" {
			supplierDeployMetadata["DeployTimeout"] = deployTimeout1
		}

		if len(supplierDeployMetadata) > 0 {
			deployMetadata["SupplierDeployMetadata"] = supplierDeployMetadata
		}
		networkMetadata := make(map[string]interface{})
		enablePrivateVpcConnection1, _ := jsonpath.Get("$[0].network_metadata[0].enable_private_vpc_connection", d.Get("deploy_metadata"))
		if enablePrivateVpcConnection1 != nil && enablePrivateVpcConnection1 != "" {
			networkMetadata["EnablePrivateVpcConnection"] = enablePrivateVpcConnection1
		}

		if len(networkMetadata) > 0 {
			deployMetadata["NetworkMetadata"] = networkMetadata
		}
		localData1, err := jsonpath.Get("$[0].template_configs", v)
		if err != nil {
			localData1 = make([]interface{}, 0)
		}
		localMaps := make([]interface{}, 0)
		for _, dataLoop1 := range convertToInterfaceArray(localData1) {
			dataLoop1Tmp := make(map[string]interface{})
			if dataLoop1 != nil {
				dataLoop1Tmp = dataLoop1.(map[string]interface{})
			}
			dataLoop1Map := make(map[string]interface{})
			dataLoop1Map["Name"] = dataLoop1Tmp["name"]
			dataLoop1Map["Url"] = dataLoop1Tmp["url"]
			dataLoop1Map["AllowedRegions"] = dataLoop1Tmp["allowed_regions"]
			dataLoop1Map["PredefinedParameters"] = dataLoop1Tmp["predefined_parameters"]
			dataLoop1Map["HiddenParameterKeys"] = dataLoop1Tmp["hidden_parameter_keys"]
			localData2 := make(map[string]interface{})
			parametersAllowedToBeModified1, _ := jsonpath.Get("$[0].parameters_allowed_to_be_modified", dataLoop1Tmp["update_info"])
			if parametersAllowedToBeModified1 != nil && parametersAllowedToBeModified1 != "" {
				localData2["ParametersAllowedToBeModified"] = parametersAllowedToBeModified1
			}
			parametersCauseInterruptionIfModified1, _ := jsonpath.Get("$[0].parameters_cause_interruption_if_modified", dataLoop1Tmp["update_info"])
			if parametersCauseInterruptionIfModified1 != nil && parametersCauseInterruptionIfModified1 != "" {
				localData2["ParametersCauseInterruptionIfModified"] = parametersCauseInterruptionIfModified1
			}
			parametersConditionallyCauseInterruptionIfModified1, _ := jsonpath.Get("$[0].parameters_conditionally_cause_interruption_if_modified", dataLoop1Tmp["update_info"])
			if parametersConditionallyCauseInterruptionIfModified1 != nil && parametersConditionallyCauseInterruptionIfModified1 != "" {
				localData2["ParametersConditionallyCauseInterruptionIfModified"] = parametersConditionallyCauseInterruptionIfModified1
			}
			parametersNotAllowedToBeModified1, _ := jsonpath.Get("$[0].parameters_not_allowed_to_be_modified", dataLoop1Tmp["update_info"])
			if parametersNotAllowedToBeModified1 != nil && parametersNotAllowedToBeModified1 != "" {
				localData2["ParametersNotAllowedToBeModified"] = parametersNotAllowedToBeModified1
			}
			parametersUncertainlyAllowedToBeModified1, _ := jsonpath.Get("$[0].parameters_uncertainly_allowed_to_be_modified", dataLoop1Tmp["update_info"])
			if parametersUncertainlyAllowedToBeModified1 != nil && parametersUncertainlyAllowedToBeModified1 != "" {
				localData2["ParametersUncertainlyAllowedToBeModified"] = parametersUncertainlyAllowedToBeModified1
			}
			parametersUncertainlyCauseInterruptionIfModified1, _ := jsonpath.Get("$[0].parameters_uncertainly_cause_interruption_if_modified", dataLoop1Tmp["update_info"])
			if parametersUncertainlyCauseInterruptionIfModified1 != nil && parametersUncertainlyCauseInterruptionIfModified1 != "" {
				localData2["ParametersUncertainlyCauseInterruptionIfModified"] = parametersUncertainlyCauseInterruptionIfModified1
			}
			if len(localData2) > 0 {
				dataLoop1Map["UpdateInfo"] = localData2
			}
			localMaps = append(localMaps, dataLoop1Map)
		}
		deployMetadata["TemplateConfigs"] = localMaps

		request["DeployMetadata"] = deployMetadata
	}

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("ComputeNestSupplier", "2021-05-21", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_compute_nest_service", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["ServiceId"]))

	return resourceAliCloudComputeNestServiceUpdate(d, meta)
}

func resourceAliCloudComputeNestServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	computeNestServiceV2 := ComputeNestServiceV2{client}

	objectRaw, err := computeNestServiceV2.DescribeComputeNestService(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_compute_nest_service DescribeComputeNestService Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("alarm_metadata", objectRaw["AlarmMetadata"])
	d.Set("approval_type", objectRaw["ApprovalType"])
	d.Set("deploy_type", objectRaw["DeployType"])
	d.Set("duration", objectRaw["Duration"])
	d.Set("is_support_operated", objectRaw["IsSupportOperated"])
	d.Set("license_metadata", objectRaw["LicenseMetadata"])
	d.Set("operation_metadata", objectRaw["OperationMetadata"])
	d.Set("policy_names", objectRaw["PolicyNames"])
	d.Set("resource_group_id", objectRaw["ResourceGroupId"])
	d.Set("service_type", objectRaw["ServiceType"])
	d.Set("share_type", objectRaw["ShareType"])
	d.Set("status", objectRaw["Status"])
	d.Set("tenant_type", objectRaw["TenantType"])
	d.Set("trial_duration", objectRaw["TrialDuration"])
	d.Set("upgrade_metadata", objectRaw["UpgradeMetadata"])
	d.Set("version", objectRaw["Version"])
	d.Set("version_name", objectRaw["VersionName"])
	d.Set("service_id", objectRaw["ServiceId"])

	deployMetadataMaps := make([]map[string]interface{}, 0)
	deployMetadataMap := make(map[string]interface{})
	deployMetadataRaw := make(map[string]interface{})
	if objectRaw["DeployMetadata"] != nil {
		deployMetadataRaw = objectRaw["DeployMetadata"].(map[string]interface{})
	}
	if len(deployMetadataRaw) > 0 {

		networkMetadataMaps := make([]map[string]interface{}, 0)
		networkMetadataMap := make(map[string]interface{})
		networkMetadataRaw := make(map[string]interface{})
		if deployMetadataRaw["NetworkMetadata"] != nil {
			networkMetadataRaw = deployMetadataRaw["NetworkMetadata"].(map[string]interface{})
		}
		if len(networkMetadataRaw) > 0 {
			networkMetadataMap["enable_private_vpc_connection"] = formatBool(networkMetadataRaw["EnablePrivateVpcConnection"])

			networkMetadataMaps = append(networkMetadataMaps, networkMetadataMap)
		}
		deployMetadataMap["network_metadata"] = networkMetadataMaps
		supplierDeployMetadataMaps := make([]map[string]interface{}, 0)
		supplierDeployMetadataMap := make(map[string]interface{})
		supplierDeployMetadataRaw := make(map[string]interface{})
		if deployMetadataRaw["SupplierDeployMetadata"] != nil {
			supplierDeployMetadataRaw = deployMetadataRaw["SupplierDeployMetadata"].(map[string]interface{})
		}
		if len(supplierDeployMetadataRaw) > 0 {
			supplierDeployMetadataMap["deploy_timeout"] = formatInt(supplierDeployMetadataRaw["DeployTimeout"])

			supplierDeployMetadataMaps = append(supplierDeployMetadataMaps, supplierDeployMetadataMap)
		}
		deployMetadataMap["supplier_deploy_metadata"] = supplierDeployMetadataMaps
		templateConfigsRaw := deployMetadataRaw["TemplateConfigs"]
		templateConfigsMaps := make([]map[string]interface{}, 0)
		if templateConfigsRaw != nil {
			for _, templateConfigsChildRaw := range convertToInterfaceArray(templateConfigsRaw) {
				templateConfigsMap := make(map[string]interface{})
				templateConfigsChildRaw := templateConfigsChildRaw.(map[string]interface{})
				templateConfigsMap["name"] = templateConfigsChildRaw["Name"]
				templateConfigsMap["url"] = templateConfigsChildRaw["Url"]

				allowedRegionsRaw := make([]interface{}, 0)
				if templateConfigsChildRaw["AllowedRegions"] != nil {
					allowedRegionsRaw = convertToInterfaceArray(templateConfigsChildRaw["AllowedRegions"])
				}

				templateConfigsMap["allowed_regions"] = allowedRegionsRaw
				hiddenParameterKeysRaw := make([]interface{}, 0)
				if templateConfigsChildRaw["HiddenParameterKeys"] != nil {
					hiddenParameterKeysRaw = convertToInterfaceArray(templateConfigsChildRaw["HiddenParameterKeys"])
				}

				templateConfigsMap["hidden_parameter_keys"] = hiddenParameterKeysRaw
				predefinedParametersRaw := make([]interface{}, 0)
				if templateConfigsChildRaw["PredefinedParameters"] != nil {
					predefinedParametersRaw = convertToInterfaceArray(templateConfigsChildRaw["PredefinedParameters"])
				}

				templateConfigsMap["predefined_parameters"] = predefinedParametersRaw
				updateInfoMaps := make([]map[string]interface{}, 0)
				updateInfoMap := make(map[string]interface{})
				updateInfoRaw := make(map[string]interface{})
				if templateConfigsChildRaw["UpdateInfo"] != nil {
					updateInfoRaw = templateConfigsChildRaw["UpdateInfo"].(map[string]interface{})
				}
				if len(updateInfoRaw) > 0 {

					parametersAllowedToBeModifiedRaw := make([]interface{}, 0)
					if updateInfoRaw["ParametersAllowedToBeModified"] != nil {
						parametersAllowedToBeModifiedRaw = convertToInterfaceArray(updateInfoRaw["ParametersAllowedToBeModified"])
					}

					updateInfoMap["parameters_allowed_to_be_modified"] = parametersAllowedToBeModifiedRaw
					parametersCauseInterruptionIfModifiedRaw := make([]interface{}, 0)
					if updateInfoRaw["ParametersCauseInterruptionIfModified"] != nil {
						parametersCauseInterruptionIfModifiedRaw = convertToInterfaceArray(updateInfoRaw["ParametersCauseInterruptionIfModified"])
					}

					updateInfoMap["parameters_cause_interruption_if_modified"] = parametersCauseInterruptionIfModifiedRaw
					parametersConditionallyCauseInterruptionIfModifiedRaw := make([]interface{}, 0)
					if updateInfoRaw["ParametersConditionallyCauseInterruptionIfModified"] != nil {
						parametersConditionallyCauseInterruptionIfModifiedRaw = convertToInterfaceArray(updateInfoRaw["ParametersConditionallyCauseInterruptionIfModified"])
					}

					updateInfoMap["parameters_conditionally_cause_interruption_if_modified"] = parametersConditionallyCauseInterruptionIfModifiedRaw
					parametersNotAllowedToBeModifiedRaw := make([]interface{}, 0)
					if updateInfoRaw["ParametersNotAllowedToBeModified"] != nil {
						parametersNotAllowedToBeModifiedRaw = convertToInterfaceArray(updateInfoRaw["ParametersNotAllowedToBeModified"])
					}

					updateInfoMap["parameters_not_allowed_to_be_modified"] = parametersNotAllowedToBeModifiedRaw
					parametersUncertainlyAllowedToBeModifiedRaw := make([]interface{}, 0)
					if updateInfoRaw["ParametersUncertainlyAllowedToBeModified"] != nil {
						parametersUncertainlyAllowedToBeModifiedRaw = convertToInterfaceArray(updateInfoRaw["ParametersUncertainlyAllowedToBeModified"])
					}

					updateInfoMap["parameters_uncertainly_allowed_to_be_modified"] = parametersUncertainlyAllowedToBeModifiedRaw
					parametersUncertainlyCauseInterruptionIfModifiedRaw := make([]interface{}, 0)
					if updateInfoRaw["ParametersUncertainlyCauseInterruptionIfModified"] != nil {
						parametersUncertainlyCauseInterruptionIfModifiedRaw = convertToInterfaceArray(updateInfoRaw["ParametersUncertainlyCauseInterruptionIfModified"])
					}

					updateInfoMap["parameters_uncertainly_cause_interruption_if_modified"] = parametersUncertainlyCauseInterruptionIfModifiedRaw
					updateInfoMaps = append(updateInfoMaps, updateInfoMap)
				}
				templateConfigsMap["update_info"] = updateInfoMaps
				templateConfigsMaps = append(templateConfigsMaps, templateConfigsMap)
			}
		}
		deployMetadataMap["template_configs"] = templateConfigsMaps
		deployMetadataMaps = append(deployMetadataMaps, deployMetadataMap)
	}
	if err := d.Set("deploy_metadata", deployMetadataMaps); err != nil {
		return err
	}
	serviceInfosRaw := objectRaw["ServiceInfos"]
	serviceInfoMaps := make([]map[string]interface{}, 0)
	if serviceInfosRaw != nil {
		for _, serviceInfosChildRaw := range convertToInterfaceArray(serviceInfosRaw) {
			serviceInfoMap := make(map[string]interface{})
			serviceInfosChildRaw := serviceInfosChildRaw.(map[string]interface{})
			serviceInfoMap["image"] = serviceInfosChildRaw["Image"]
			serviceInfoMap["locale"] = serviceInfosChildRaw["Locale"]
			serviceInfoMap["name"] = serviceInfosChildRaw["Name"]
			serviceInfoMap["short_description"] = serviceInfosChildRaw["ShortDescription"]

			serviceInfoMaps = append(serviceInfoMaps, serviceInfoMap)
		}
	}
	if err := d.Set("service_info", serviceInfoMaps); err != nil {
		return err
	}

	d.Set("service_id", d.Id())

	return nil
}

func resourceAliCloudComputeNestServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateService"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["ServiceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("ComputeNestSupplier", "2021-05-21", action, query, request, true)
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

	return resourceAliCloudComputeNestServiceRead(d, meta)
}

func resourceAliCloudComputeNestServiceDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteService"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["ServiceId"] = d.Id()
	request["RegionId"] = client.RegionId
	request["ClientToken"] = buildClientToken(action)

	request["ServiceVersion"] = d.Get("version")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("ComputeNestSupplier", "2021-05-21", action, query, request, true)
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
