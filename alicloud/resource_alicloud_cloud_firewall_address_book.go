// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudCloudFirewallAddressBook() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCloudFirewallAddressBookCreate,
		Read:   resourceAliCloudCloudFirewallAddressBookRead,
		Update: resourceAliCloudCloudFirewallAddressBookUpdate,
		Delete: resourceAliCloudCloudFirewallAddressBookDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"address_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"address_list_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"asset_member_uids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"asset_region_resource_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asset_region_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_type": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv6": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ai_gateway_e_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"api_gateway_e_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
												},
												"nlb_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"eni_e_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"slb_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"ecs_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"alb_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"ga_e_ipv6": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
									"ipv4": {
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"slb_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"ai_gateway_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"ga_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"nat_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"bastion_host_ingress_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"eip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"nat_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"slb_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"eni_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"nlb_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"ecs_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"hav_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"api_gateway_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"ecs_public_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"bastion_host_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"bastion_host_egress_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
												"alb_e_ip": {
													Type:     schema.TypeBool,
													Optional: true,
													ForceNew: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"auto_add_tag_ecs": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: IntInSlice([]int{0, 1}),
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"ip", "domain", "port", "tag", "ipv6", "asset", "assetIpv6"}, false),
			},
			"lang": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reference_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tag_relation": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"and", "or"}, false),
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAliCloudCloudFirewallAddressBookCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "AddAddressBook"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	assetRegionResourceTypesDataList := make(map[string]interface{})
	if v, ok := d.GetOk("asset_region_resource_types"); ok {
		assetRegionId1, _ := jsonpath.Get("$.asset_region_id", v)
		if assetRegionId1 != nil && assetRegionId1 != "" {
			assetRegionResourceTypesDataList["AssetRegionId"] = assetRegionId1
		}
	}
	if v := d.Get("asset_region_resource_types"); !IsNil(v) {
		resourceType := make(map[string]interface{})
		ipv6 := make(map[string]interface{})
		albIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].alb_ipv6", d.Get("asset_region_resource_types"))
		if albIPv61 != nil && albIPv61 != "" {
			ipv6["AlbIPv6"] = albIPv61
		}
		slbIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].slb_ipv6", d.Get("asset_region_resource_types"))
		if slbIPv61 != nil && slbIPv61 != "" {
			ipv6["SlbIPv6"] = slbIPv61
		}
		nlbIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].nlb_ipv6", d.Get("asset_region_resource_types"))
		if nlbIPv61 != nil && nlbIPv61 != "" {
			ipv6["NlbIPv6"] = nlbIPv61
		}
		apiGatewayEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].api_gateway_e_ipv6", d.Get("asset_region_resource_types"))
		if apiGatewayEIPv61 != nil && apiGatewayEIPv61 != "" {
			ipv6["ApiGatewayEIPv6"] = apiGatewayEIPv61
		}
		ecsIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].ecs_ipv6", d.Get("asset_region_resource_types"))
		if ecsIPv61 != nil && ecsIPv61 != "" {
			ipv6["EcsIPv6"] = ecsIPv61
		}
		gaEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].ga_e_ipv6", d.Get("asset_region_resource_types"))
		if gaEIPv61 != nil && gaEIPv61 != "" {
			ipv6["GaEIPv6"] = gaEIPv61
		}
		eniEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].eni_e_ipv6", d.Get("asset_region_resource_types"))
		if eniEIPv61 != nil && eniEIPv61 != "" {
			ipv6["EniEIPv6"] = eniEIPv61
		}
		aiGatewayEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].ai_gateway_e_ipv6", d.Get("asset_region_resource_types"))
		if aiGatewayEIPv61 != nil && aiGatewayEIPv61 != "" {
			ipv6["AiGatewayEIPv6"] = aiGatewayEIPv61
		}

		if len(ipv6) > 0 {
			resourceType["Ipv6"] = ipv6
		}
		ipv4 := make(map[string]interface{})
		ecsEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ecs_e_ip", d.Get("asset_region_resource_types"))
		if ecsEIP1 != nil && ecsEIP1 != "" {
			ipv4["EcsEIP"] = ecsEIP1
		}
		nlbEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].nlb_e_ip", d.Get("asset_region_resource_types"))
		if nlbEIP1 != nil && nlbEIP1 != "" {
			ipv4["NlbEIP"] = nlbEIP1
		}
		bastionHostIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].bastion_host_ip", d.Get("asset_region_resource_types"))
		if bastionHostIP1 != nil && bastionHostIP1 != "" {
			ipv4["BastionHostIP"] = bastionHostIP1
		}
		ecsPublicIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ecs_public_ip", d.Get("asset_region_resource_types"))
		if ecsPublicIP1 != nil && ecsPublicIP1 != "" {
			ipv4["EcsPublicIP"] = ecsPublicIP1
		}
		aiGatewayEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ai_gateway_e_ip", d.Get("asset_region_resource_types"))
		if aiGatewayEIP1 != nil && aiGatewayEIP1 != "" {
			ipv4["AiGatewayEIP"] = aiGatewayEIP1
		}
		bastionHostIngressIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].bastion_host_ingress_ip", d.Get("asset_region_resource_types"))
		if bastionHostIngressIP1 != nil && bastionHostIngressIP1 != "" {
			ipv4["BastionHostIngressIP"] = bastionHostIngressIP1
		}
		natPublicIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].nat_public_ip", d.Get("asset_region_resource_types"))
		if natPublicIP1 != nil && natPublicIP1 != "" {
			ipv4["NatPublicIP"] = natPublicIP1
		}
		slbPublicIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].slb_public_ip", d.Get("asset_region_resource_types"))
		if slbPublicIP1 != nil && slbPublicIP1 != "" {
			ipv4["SlbPublicIP"] = slbPublicIP1
		}
		bastionHostEgressIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].bastion_host_egress_ip", d.Get("asset_region_resource_types"))
		if bastionHostEgressIP1 != nil && bastionHostEgressIP1 != "" {
			ipv4["BastionHostEgressIP"] = bastionHostEgressIP1
		}
		eIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].eip", d.Get("asset_region_resource_types"))
		if eIP1 != nil && eIP1 != "" {
			ipv4["EIP"] = eIP1
		}
		slbEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].slb_e_ip", d.Get("asset_region_resource_types"))
		if slbEIP1 != nil && slbEIP1 != "" {
			ipv4["SlbEIP"] = slbEIP1
		}
		albEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].alb_e_ip", d.Get("asset_region_resource_types"))
		if albEIP1 != nil && albEIP1 != "" {
			ipv4["AlbEIP"] = albEIP1
		}
		gaEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ga_e_ip", d.Get("asset_region_resource_types"))
		if gaEIP1 != nil && gaEIP1 != "" {
			ipv4["GaEIP"] = gaEIP1
		}
		eniEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].eni_e_ip", d.Get("asset_region_resource_types"))
		if eniEIP1 != nil && eniEIP1 != "" {
			ipv4["EniEIP"] = eniEIP1
		}
		apiGatewayEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].api_gateway_e_ip", d.Get("asset_region_resource_types"))
		if apiGatewayEIP1 != nil && apiGatewayEIP1 != "" {
			ipv4["ApiGatewayEIP"] = apiGatewayEIP1
		}
		hAVIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].hav_ip", d.Get("asset_region_resource_types"))
		if hAVIP1 != nil && hAVIP1 != "" {
			ipv4["HAVIP"] = hAVIP1
		}
		natEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].nat_e_ip", d.Get("asset_region_resource_types"))
		if natEIP1 != nil && natEIP1 != "" {
			ipv4["NatEIP"] = natEIP1
		}

		if len(ipv4) > 0 {
			resourceType["Ipv4"] = ipv4
		}

		if len(resourceType) > 0 {
			assetRegionResourceTypesDataList["ResourceType"] = resourceType
		}
	}

	AssetRegionResourceTypesMap := make([]interface{}, 0)
	AssetRegionResourceTypesMap = append(AssetRegionResourceTypesMap, assetRegionResourceTypesDataList)
	assetRegionResourceTypesDataListJson, err := json.Marshal(AssetRegionResourceTypesMap)
	if err != nil {
		return WrapError(err)
	}
	request["AssetRegionResourceTypes"] = string(assetRegionResourceTypesDataListJson)

	if v, ok := d.GetOk("tag_relation"); ok {
		request["TagRelation"] = v
	}
	if v, ok := d.GetOkExists("auto_add_tag_ecs"); ok {
		request["AutoAddTagEcs"] = v
	}
	request["GroupType"] = d.Get("group_type")
	if v, ok := d.GetOk("asset_member_uids"); ok {
		assetMemberUidsMapsArray := convertToInterfaceArray(v)

		assetMemberUidsMapsJson, err := json.Marshal(assetMemberUidsMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["AssetMemberUids"] = string(assetMemberUidsMapsJson)
	}

	tagListDataList := make(map[string]interface{})
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagListDataList["TagKey"] = tagsMap
	}
	if v, ok := d.GetOk("tags"); ok {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		tagListDataList["TagValue"] = tagsMap
	}

	TagListMap := make([]interface{}, 0)
	TagListMap = append(TagListMap, tagListDataList)
	request["TagList"] = TagListMap

	if v, ok := d.GetOk("address_list"); ok {
		addressListJsonPath, err := jsonpath.Get("$", v)
		if err == nil && addressListJsonPath != "" {
			request["AddressList"] = convertListToCommaSeparate(addressListJsonPath.([]interface{}))
		}
	}
	request["Description"] = d.Get("description")
	request["GroupName"] = d.Get("group_name")
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cloud_firewall_address_book", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["GroupUuid"]))

	return resourceAliCloudCloudFirewallAddressBookRead(d, meta)
}

func resourceAliCloudCloudFirewallAddressBookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cloudFirewallServiceV2 := CloudFirewallServiceV2{client}

	objectRaw, err := cloudFirewallServiceV2.DescribeCloudFirewallAddressBook(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cloud_firewall_address_book DescribeCloudFirewallAddressBook Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("address_list_count", objectRaw["AddressListCount"])
	d.Set("auto_add_tag_ecs", objectRaw["AutoAddTagEcs"])
	d.Set("description", objectRaw["Description"])
	d.Set("group_name", objectRaw["GroupName"])
	d.Set("group_type", objectRaw["GroupType"])
	d.Set("reference_count", objectRaw["ReferenceCount"])
	d.Set("tag_relation", objectRaw["TagRelation"])

	addressListRaw := make([]interface{}, 0)
	if objectRaw["AddressList"] != nil {
		addressListRaw = convertToInterfaceArray(objectRaw["AddressList"])
	}

	d.Set("address_list", addressListRaw)
	assetMemberUidsRaw := make([]interface{}, 0)
	if objectRaw["AssetMemberUids"] != nil {
		assetMemberUidsRaw = convertToInterfaceArray(objectRaw["AssetMemberUids"])
	}

	d.Set("asset_member_uids", assetMemberUidsRaw)
	assetRegionResourceTypesRaw := objectRaw["AssetRegionResourceTypes"]
	assetRegionResourceTypesMaps := make([]map[string]interface{}, 0)
	if assetRegionResourceTypesRaw != nil {
		for _, assetRegionResourceTypesChildRaw := range convertToInterfaceArray(assetRegionResourceTypesRaw) {
			assetRegionResourceTypesMap := make(map[string]interface{})
			assetRegionResourceTypesChildRaw := assetRegionResourceTypesChildRaw.(map[string]interface{})
			assetRegionResourceTypesMap["asset_region_id"] = assetRegionResourceTypesChildRaw["AssetRegionId"]

			resourceTypeMaps := make([]map[string]interface{}, 0)
			resourceTypeMap := make(map[string]interface{})
			resourceTypeRaw := make(map[string]interface{})
			if assetRegionResourceTypesChildRaw["ResourceType"] != nil {
				resourceTypeRaw = assetRegionResourceTypesChildRaw["ResourceType"].(map[string]interface{})
			}
			if len(resourceTypeRaw) > 0 {

				ipv4Maps := make([]map[string]interface{}, 0)
				ipv4Map := make(map[string]interface{})
				ipv4Raw := make(map[string]interface{})
				if resourceTypeRaw["Ipv4"] != nil {
					ipv4Raw = resourceTypeRaw["Ipv4"].(map[string]interface{})
				}
				if len(ipv4Raw) > 0 {
					ipv4Map["ai_gateway_e_ip"] = ipv4Raw["AiGatewayEIP"]
					ipv4Map["alb_e_ip"] = ipv4Raw["AlbEIP"]
					ipv4Map["api_gateway_e_ip"] = ipv4Raw["ApiGatewayEIP"]
					ipv4Map["bastion_host_egress_ip"] = ipv4Raw["BastionHostEgressIP"]
					ipv4Map["bastion_host_ip"] = ipv4Raw["BastionHostIP"]
					ipv4Map["bastion_host_ingress_ip"] = ipv4Raw["BastionHostIngressIP"]
					ipv4Map["eip"] = ipv4Raw["EIP"]
					ipv4Map["ecs_e_ip"] = ipv4Raw["EcsEIP"]
					ipv4Map["ecs_public_ip"] = ipv4Raw["EcsPublicIP"]
					ipv4Map["eni_e_ip"] = ipv4Raw["EniEIP"]
					ipv4Map["ga_e_ip"] = ipv4Raw["GaEIP"]
					ipv4Map["hav_ip"] = ipv4Raw["HAVIP"]
					ipv4Map["nat_e_ip"] = ipv4Raw["NatEIP"]
					ipv4Map["nat_public_ip"] = ipv4Raw["NatPublicIP"]
					ipv4Map["nlb_e_ip"] = ipv4Raw["NlbEIP"]
					ipv4Map["slb_e_ip"] = ipv4Raw["SlbEIP"]
					ipv4Map["slb_public_ip"] = ipv4Raw["SlbPublicIP"]

					ipv4Maps = append(ipv4Maps, ipv4Map)
				}
				resourceTypeMap["ipv4"] = ipv4Maps
				ipv6Maps := make([]map[string]interface{}, 0)
				ipv6Map := make(map[string]interface{})
				ipv6Raw := make(map[string]interface{})
				if resourceTypeRaw["Ipv6"] != nil {
					ipv6Raw = resourceTypeRaw["Ipv6"].(map[string]interface{})
				}
				if len(ipv6Raw) > 0 {
					ipv6Map["ai_gateway_e_ipv6"] = ipv6Raw["AiGatewayEIPv6"]
					ipv6Map["alb_ipv6"] = ipv6Raw["AlbIPv6"]
					ipv6Map["api_gateway_e_ipv6"] = ipv6Raw["ApiGatewayEIPv6"]
					ipv6Map["ecs_ipv6"] = ipv6Raw["EcsIPv6"]
					ipv6Map["eni_e_ipv6"] = ipv6Raw["EniEIPv6"]
					ipv6Map["ga_e_ipv6"] = ipv6Raw["GaEIPv6"]
					ipv6Map["nlb_ipv6"] = ipv6Raw["NlbIPv6"]
					ipv6Map["slb_ipv6"] = ipv6Raw["SlbIPv6"]

					ipv6Maps = append(ipv6Maps, ipv6Map)
				}
				resourceTypeMap["ipv6"] = ipv6Maps
				resourceTypeMaps = append(resourceTypeMaps, resourceTypeMap)
			}
			assetRegionResourceTypesMap["resource_type"] = resourceTypeMaps
			assetRegionResourceTypesMaps = append(assetRegionResourceTypesMaps, assetRegionResourceTypesMap)
		}
	}
	if err := d.Set("asset_region_resource_types", assetRegionResourceTypesMaps); err != nil {
		return err
	}
	tagsMaps := objectRaw["TagList"]
	d.Set("tags", tagsToMap(tagsMaps))

	return nil
}

func resourceAliCloudCloudFirewallAddressBookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "ModifyAddressBook"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["GroupUuid"] = d.Id()

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	if d.HasChange("tag_relation") {
		update = true
	}
	if v, ok := d.GetOk("tag_relation"); ok || d.HasChange("tag_relation") {
		request["TagRelation"] = v
	}
	assetRegionResourceTypesDataList := make(map[string]interface{})
	if d.HasChange("asset_region_resource_types") {
		update = true
	}
	assetRegionId1, _ := jsonpath.Get("$.asset_region_id", d.Get("asset_region_resource_types"))
	if assetRegionId1 != nil && assetRegionId1 != "" {
		assetRegionResourceTypesDataList["AssetRegionId"] = assetRegionId1
	}
	if d.HasChange("asset_region_resource_types") {
		update = true
	}
	if v := d.Get("asset_region_resource_types"); !IsNil(v) || d.HasChange("asset_region_resource_types") {
		resourceType := make(map[string]interface{})
		ipv6 := make(map[string]interface{})
		albIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].alb_ipv6", d.Get("asset_region_resource_types"))
		if albIPv61 != nil && albIPv61 != "" {
			ipv6["AlbIPv6"] = albIPv61
		}
		slbIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].slb_ipv6", d.Get("asset_region_resource_types"))
		if slbIPv61 != nil && slbIPv61 != "" {
			ipv6["SlbIPv6"] = slbIPv61
		}
		nlbIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].nlb_ipv6", d.Get("asset_region_resource_types"))
		if nlbIPv61 != nil && nlbIPv61 != "" {
			ipv6["NlbIPv6"] = nlbIPv61
		}
		apiGatewayEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].api_gateway_e_ipv6", d.Get("asset_region_resource_types"))
		if apiGatewayEIPv61 != nil && apiGatewayEIPv61 != "" {
			ipv6["ApiGatewayEIPv6"] = apiGatewayEIPv61
		}
		ecsIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].ecs_ipv6", d.Get("asset_region_resource_types"))
		if ecsIPv61 != nil && ecsIPv61 != "" {
			ipv6["EcsIPv6"] = ecsIPv61
		}
		gaEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].ga_e_ipv6", d.Get("asset_region_resource_types"))
		if gaEIPv61 != nil && gaEIPv61 != "" {
			ipv6["GaEIPv6"] = gaEIPv61
		}
		eniEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].eni_e_ipv6", d.Get("asset_region_resource_types"))
		if eniEIPv61 != nil && eniEIPv61 != "" {
			ipv6["EniEIPv6"] = eniEIPv61
		}
		aiGatewayEIPv61, _ := jsonpath.Get("$.resource_type[0].ipv6[0].ai_gateway_e_ipv6", d.Get("asset_region_resource_types"))
		if aiGatewayEIPv61 != nil && aiGatewayEIPv61 != "" {
			ipv6["AiGatewayEIPv6"] = aiGatewayEIPv61
		}

		if len(ipv6) > 0 {
			resourceType["Ipv6"] = ipv6
		}
		ipv4 := make(map[string]interface{})
		ecsEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ecs_e_ip", d.Get("asset_region_resource_types"))
		if ecsEIP1 != nil && ecsEIP1 != "" {
			ipv4["EcsEIP"] = ecsEIP1
		}
		nlbEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].nlb_e_ip", d.Get("asset_region_resource_types"))
		if nlbEIP1 != nil && nlbEIP1 != "" {
			ipv4["NlbEIP"] = nlbEIP1
		}
		bastionHostIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].bastion_host_ip", d.Get("asset_region_resource_types"))
		if bastionHostIP1 != nil && bastionHostIP1 != "" {
			ipv4["BastionHostIP"] = bastionHostIP1
		}
		ecsPublicIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ecs_public_ip", d.Get("asset_region_resource_types"))
		if ecsPublicIP1 != nil && ecsPublicIP1 != "" {
			ipv4["EcsPublicIP"] = ecsPublicIP1
		}
		aiGatewayEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ai_gateway_e_ip", d.Get("asset_region_resource_types"))
		if aiGatewayEIP1 != nil && aiGatewayEIP1 != "" {
			ipv4["AiGatewayEIP"] = aiGatewayEIP1
		}
		bastionHostIngressIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].bastion_host_ingress_ip", d.Get("asset_region_resource_types"))
		if bastionHostIngressIP1 != nil && bastionHostIngressIP1 != "" {
			ipv4["BastionHostIngressIP"] = bastionHostIngressIP1
		}
		natPublicIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].nat_public_ip", d.Get("asset_region_resource_types"))
		if natPublicIP1 != nil && natPublicIP1 != "" {
			ipv4["NatPublicIP"] = natPublicIP1
		}
		slbPublicIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].slb_public_ip", d.Get("asset_region_resource_types"))
		if slbPublicIP1 != nil && slbPublicIP1 != "" {
			ipv4["SlbPublicIP"] = slbPublicIP1
		}
		bastionHostEgressIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].bastion_host_egress_ip", d.Get("asset_region_resource_types"))
		if bastionHostEgressIP1 != nil && bastionHostEgressIP1 != "" {
			ipv4["BastionHostEgressIP"] = bastionHostEgressIP1
		}
		eIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].eip", d.Get("asset_region_resource_types"))
		if eIP1 != nil && eIP1 != "" {
			ipv4["EIP"] = eIP1
		}
		slbEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].slb_e_ip", d.Get("asset_region_resource_types"))
		if slbEIP1 != nil && slbEIP1 != "" {
			ipv4["SlbEIP"] = slbEIP1
		}
		albEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].alb_e_ip", d.Get("asset_region_resource_types"))
		if albEIP1 != nil && albEIP1 != "" {
			ipv4["AlbEIP"] = albEIP1
		}
		gaEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].ga_e_ip", d.Get("asset_region_resource_types"))
		if gaEIP1 != nil && gaEIP1 != "" {
			ipv4["GaEIP"] = gaEIP1
		}
		eniEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].eni_e_ip", d.Get("asset_region_resource_types"))
		if eniEIP1 != nil && eniEIP1 != "" {
			ipv4["EniEIP"] = eniEIP1
		}
		apiGatewayEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].api_gateway_e_ip", d.Get("asset_region_resource_types"))
		if apiGatewayEIP1 != nil && apiGatewayEIP1 != "" {
			ipv4["ApiGatewayEIP"] = apiGatewayEIP1
		}
		hAVIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].hav_ip", d.Get("asset_region_resource_types"))
		if hAVIP1 != nil && hAVIP1 != "" {
			ipv4["HAVIP"] = hAVIP1
		}
		natEIP1, _ := jsonpath.Get("$.resource_type[0].ipv4[0].nat_e_ip", d.Get("asset_region_resource_types"))
		if natEIP1 != nil && natEIP1 != "" {
			ipv4["NatEIP"] = natEIP1
		}

		if len(ipv4) > 0 {
			resourceType["Ipv4"] = ipv4
		}

		if len(resourceType) > 0 {
			assetRegionResourceTypesDataList["ResourceType"] = resourceType
		}
	}

	AssetRegionResourceTypesMap := make([]interface{}, 0)
	AssetRegionResourceTypesMap = append(AssetRegionResourceTypesMap, assetRegionResourceTypesDataList)
	assetRegionResourceTypesDataListJson, err := json.Marshal(AssetRegionResourceTypesMap)
	if err != nil {
		return WrapError(err)
	}
	request["AssetRegionResourceTypes"] = string(assetRegionResourceTypesDataListJson)

	if d.HasChange("auto_add_tag_ecs") {
		update = true
	}
	if v, ok := d.GetOkExists("auto_add_tag_ecs"); ok || d.HasChange("auto_add_tag_ecs") {
		request["AutoAddTagEcs"] = v
	}
	if d.HasChange("asset_member_uids") {
		update = true
	}
	if v, ok := d.GetOk("asset_member_uids"); ok || d.HasChange("asset_member_uids") {
		assetMemberUidsMapsArray := convertToInterfaceArray(v)

		assetMemberUidsMapsJson, err := json.Marshal(assetMemberUidsMapsArray)
		if err != nil {
			return WrapError(err)
		}
		request["AssetMemberUids"] = string(assetMemberUidsMapsJson)
	}

	if d.HasChange("tags") {
		update = true
	}
	if v, ok := d.GetOk("tags"); ok || d.HasChange("tags") {
		tagsMap := ConvertTags(v.(map[string]interface{}))
		request = expandTagsToMap(request, tagsMap)
	}

	if d.HasChange("address_list") {
		update = true
	}
	if v, ok := d.GetOk("address_list"); ok || d.HasChange("address_list") {
		addressListJsonPath, err := jsonpath.Get("$", v)
		if err == nil && addressListJsonPath != "" {
			request["AddressList"] = convertListToCommaSeparate(addressListJsonPath.([]interface{}))
		}
	}
	if d.HasChange("description") {
		update = true
	}
	request["Description"] = d.Get("description")
	if d.HasChange("group_name") {
		update = true
	}
	request["GroupName"] = d.Get("group_name")
	if update {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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

	return resourceAliCloudCloudFirewallAddressBookRead(d, meta)
}

func resourceAliCloudCloudFirewallAddressBookDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteAddressBook"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["GroupUuid"] = d.Id()

	if v, ok := d.GetOk("lang"); ok {
		request["Lang"] = v
	}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Cloudfw", "2017-12-07", action, query, request, true)
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
