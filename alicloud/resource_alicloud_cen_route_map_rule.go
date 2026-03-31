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

func resourceAliCloudCenRouteMapRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudCenRouteMapRuleCreate,
		Read:   resourceAliCloudCenRouteMapRuleRead,
		Update: resourceAliCloudCenRouteMapRuleUpdate,
		Delete: resourceAliCloudCenRouteMapRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dry_run": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"match_conditions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination_instance_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"match_address_type": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"community_set_include": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_instance_ids_reverse_match": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"community_set_match": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_route_table_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"route_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_instance_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"address_prefixes_include": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_instance_ids_reverse_match": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"source_instance_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"as_paths_include": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
						"as_paths_match": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
						"destination_route_table_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination_region_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_instance_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"address_prefixes_match": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"source_region_ids": {
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
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"route_map_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"set_actions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"as_path_replace": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
						"community_add": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"community_replace": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"next_priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"route_action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"as_path_prepend": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeInt},
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAliCloudCenRouteMapRuleCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateRouteMapRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})

	request["ClientToken"] = buildClientToken(action)

	matchConditions := make(map[string]interface{})

	if v := d.Get("match_conditions"); !IsNil(v) {
		communitySetInclude1, _ := jsonpath.Get("$[0].community_set_include", v)
		if communitySetInclude1 != nil && communitySetInclude1 != "" {
			matchConditions["CommunitySetInclude"] = communitySetInclude1
		}
		destinationRouteTableIds1, _ := jsonpath.Get("$[0].destination_route_table_ids", v)
		if destinationRouteTableIds1 != nil && destinationRouteTableIds1 != "" {
			matchConditions["DestinationRouteTableIds"] = destinationRouteTableIds1
		}
		sourceInstanceIdsReverseMatch1, _ := jsonpath.Get("$[0].source_instance_ids_reverse_match", v)
		if sourceInstanceIdsReverseMatch1 != nil && sourceInstanceIdsReverseMatch1 != "" {
			matchConditions["SourceInstanceIdsReverseMatch"] = sourceInstanceIdsReverseMatch1
		}
		destinationInstanceIdsReverseMatch1, _ := jsonpath.Get("$[0].destination_instance_ids_reverse_match", v)
		if destinationInstanceIdsReverseMatch1 != nil && destinationInstanceIdsReverseMatch1 != "" {
			matchConditions["DestinationInstanceIdsReverseMatch"] = destinationInstanceIdsReverseMatch1
		}
		routeTypes1, _ := jsonpath.Get("$[0].route_types", v)
		if routeTypes1 != nil && routeTypes1 != "" {
			matchConditions["RouteTypes"] = routeTypes1
		}
		sourceRegionIds1, _ := jsonpath.Get("$[0].source_region_ids", v)
		if sourceRegionIds1 != nil && sourceRegionIds1 != "" {
			matchConditions["SourceRegionIds"] = sourceRegionIds1
		}
		matchAddressType1, _ := jsonpath.Get("$[0].match_address_type", v)
		if matchAddressType1 != nil && matchAddressType1 != "" {
			matchConditions["MatchAddressType"] = matchAddressType1
		}
		destinationRegionIds1, _ := jsonpath.Get("$[0].destination_region_ids", v)
		if destinationRegionIds1 != nil && destinationRegionIds1 != "" {
			matchConditions["DestinationRegionIds"] = destinationRegionIds1
		}
		asPathsInclude1, _ := jsonpath.Get("$[0].as_paths_include", v)
		if asPathsInclude1 != nil && asPathsInclude1 != "" {
			matchConditions["AsPathsInclude"] = asPathsInclude1
		}
		sourceInstanceTypes1, _ := jsonpath.Get("$[0].source_instance_types", v)
		if sourceInstanceTypes1 != nil && sourceInstanceTypes1 != "" {
			matchConditions["SourceInstanceTypes"] = sourceInstanceTypes1
		}
		sourceInstanceIds1, _ := jsonpath.Get("$[0].source_instance_ids", v)
		if sourceInstanceIds1 != nil && sourceInstanceIds1 != "" {
			matchConditions["SourceInstanceIds"] = sourceInstanceIds1
		}
		sourceRouteTableIds1, _ := jsonpath.Get("$[0].source_route_table_ids", v)
		if sourceRouteTableIds1 != nil && sourceRouteTableIds1 != "" {
			matchConditions["SourceRouteTableIds"] = sourceRouteTableIds1
		}
		asPathsMatch1, _ := jsonpath.Get("$[0].as_paths_match", v)
		if asPathsMatch1 != nil && asPathsMatch1 != "" {
			matchConditions["AsPathsMatch"] = asPathsMatch1
		}
		destinationInstanceIds1, _ := jsonpath.Get("$[0].destination_instance_ids", v)
		if destinationInstanceIds1 != nil && destinationInstanceIds1 != "" {
			matchConditions["DestinationInstanceIds"] = destinationInstanceIds1
		}
		addressPrefixesMatch1, _ := jsonpath.Get("$[0].address_prefixes_match", v)
		if addressPrefixesMatch1 != nil && addressPrefixesMatch1 != "" {
			matchConditions["AddressPrefixesMatch"] = addressPrefixesMatch1
		}
		addressPrefixesInclude1, _ := jsonpath.Get("$[0].address_prefixes_include", v)
		if addressPrefixesInclude1 != nil && addressPrefixesInclude1 != "" {
			matchConditions["AddressPrefixesInclude"] = addressPrefixesInclude1
		}
		destinationInstanceTypes1, _ := jsonpath.Get("$[0].destination_instance_types", v)
		if destinationInstanceTypes1 != nil && destinationInstanceTypes1 != "" {
			matchConditions["DestinationInstanceTypes"] = destinationInstanceTypes1
		}
		communitySetMatch1, _ := jsonpath.Get("$[0].community_set_match", v)
		if communitySetMatch1 != nil && communitySetMatch1 != "" {
			matchConditions["CommunitySetMatch"] = communitySetMatch1
		}

		matchConditionsJson, err := json.Marshal(matchConditions)
		if err != nil {
			return WrapError(err)
		}
		request["MatchConditions"] = string(matchConditionsJson)
	}

	setActions := make(map[string]interface{})

	if v := d.Get("set_actions"); !IsNil(v) {
		communityAdd1, _ := jsonpath.Get("$[0].community_add", v)
		if communityAdd1 != nil && communityAdd1 != "" {
			setActions["CommunityAdd"] = communityAdd1
		}
		routeAction1, _ := jsonpath.Get("$[0].route_action", v)
		if routeAction1 != nil && routeAction1 != "" {
			setActions["RouteAction"] = routeAction1
		}
		nextPriority1, _ := jsonpath.Get("$[0].next_priority", v)
		if nextPriority1 != nil && nextPriority1 != "" {
			setActions["NextPriority"] = nextPriority1
		}
		asPathPrepend1, _ := jsonpath.Get("$[0].as_path_prepend", v)
		if asPathPrepend1 != nil && asPathPrepend1 != "" {
			setActions["AsPathPrepend"] = asPathPrepend1
		}
		communityReplace1, _ := jsonpath.Get("$[0].community_replace", v)
		if communityReplace1 != nil && communityReplace1 != "" {
			setActions["CommunityReplace"] = communityReplace1
		}
		asPathReplace1, _ := jsonpath.Get("$[0].as_path_replace", v)
		if asPathReplace1 != nil && asPathReplace1 != "" {
			setActions["AsPathReplace"] = asPathReplace1
		}

		setActionsJson, err := json.Marshal(setActions)
		if err != nil {
			return WrapError(err)
		}
		request["SetActions"] = string(setActionsJson)
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if v, ok := d.GetOk("name"); ok {
		request["Name"] = v
	}
	request["RouteMapId"] = d.Get("route_map_id")
	request["Priority"] = d.Get("priority")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	wait := incrementalWait(3*time.Second, 0*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("Cbn", "2017-09-12", action, query, request, true)
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_cen_route_map_rule", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(response["RouteMapRuleId"]))

	cenServiceV2 := CenServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, cenServiceV2.CenRouteMapRuleStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return resourceAliCloudCenRouteMapRuleRead(d, meta)
}

func resourceAliCloudCenRouteMapRuleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	cenServiceV2 := CenServiceV2{client}

	objectRaw, err := cenServiceV2.DescribeCenRouteMapRule(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_cen_route_map_rule DescribeCenRouteMapRule Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("description", objectRaw["Description"])
	d.Set("name", objectRaw["Name"])
	d.Set("priority", objectRaw["Priority"])
	d.Set("route_map_id", objectRaw["RouteMapId"])
	d.Set("status", objectRaw["Status"])

	matchConditionsMaps := make([]map[string]interface{}, 0)
	matchConditionsMap := make(map[string]interface{})
	matchConditionsRaw := make(map[string]interface{})
	if objectRaw["MatchConditions"] != nil {
		matchConditionsRaw = objectRaw["MatchConditions"].(map[string]interface{})
	}
	if len(matchConditionsRaw) > 0 {
		matchConditionsMap["destination_instance_ids_reverse_match"] = matchConditionsRaw["DestinationInstanceIdsReverseMatch"]
		matchConditionsMap["match_address_type"] = matchConditionsRaw["MatchAddressType"]
		matchConditionsMap["source_instance_ids_reverse_match"] = matchConditionsRaw["SourceInstanceIdsReverseMatch"]

		addressPrefixesIncludeRaw := make([]interface{}, 0)
		if matchConditionsRaw["AddressPrefixesInclude"] != nil {
			addressPrefixesIncludeRaw = convertToInterfaceArray(matchConditionsRaw["AddressPrefixesInclude"])
		}

		matchConditionsMap["address_prefixes_include"] = addressPrefixesIncludeRaw
		addressPrefixesMatchRaw := make([]interface{}, 0)
		if matchConditionsRaw["AddressPrefixesMatch"] != nil {
			addressPrefixesMatchRaw = convertToInterfaceArray(matchConditionsRaw["AddressPrefixesMatch"])
		}

		matchConditionsMap["address_prefixes_match"] = addressPrefixesMatchRaw
		asPathsIncludeRaw := make([]interface{}, 0)
		if matchConditionsRaw["AsPathsInclude"] != nil {
			asPathsIncludeRaw = convertToInterfaceArray(matchConditionsRaw["AsPathsInclude"])
		}

		matchConditionsMap["as_paths_include"] = asPathsIncludeRaw
		asPathsMatchRaw := make([]interface{}, 0)
		if matchConditionsRaw["AsPathsMatch"] != nil {
			asPathsMatchRaw = convertToInterfaceArray(matchConditionsRaw["AsPathsMatch"])
		}

		matchConditionsMap["as_paths_match"] = asPathsMatchRaw
		communitySetIncludeRaw := make([]interface{}, 0)
		if matchConditionsRaw["CommunitySetInclude"] != nil {
			communitySetIncludeRaw = convertToInterfaceArray(matchConditionsRaw["CommunitySetInclude"])
		}

		matchConditionsMap["community_set_include"] = communitySetIncludeRaw
		communitySetMatchRaw := make([]interface{}, 0)
		if matchConditionsRaw["CommunitySetMatch"] != nil {
			communitySetMatchRaw = convertToInterfaceArray(matchConditionsRaw["CommunitySetMatch"])
		}

		matchConditionsMap["community_set_match"] = communitySetMatchRaw
		destinationInstanceIdsRaw := make([]interface{}, 0)
		if matchConditionsRaw["DestinationInstanceIds"] != nil {
			destinationInstanceIdsRaw = convertToInterfaceArray(matchConditionsRaw["DestinationInstanceIds"])
		}

		matchConditionsMap["destination_instance_ids"] = destinationInstanceIdsRaw
		destinationInstanceTypesRaw := make([]interface{}, 0)
		if matchConditionsRaw["DestinationInstanceTypes"] != nil {
			destinationInstanceTypesRaw = convertToInterfaceArray(matchConditionsRaw["DestinationInstanceTypes"])
		}

		matchConditionsMap["destination_instance_types"] = destinationInstanceTypesRaw
		destinationRegionIdsRaw := make([]interface{}, 0)
		if matchConditionsRaw["DestinationRegionIds"] != nil {
			destinationRegionIdsRaw = convertToInterfaceArray(matchConditionsRaw["DestinationRegionIds"])
		}

		matchConditionsMap["destination_region_ids"] = destinationRegionIdsRaw
		destinationRouteTableIdsRaw := make([]interface{}, 0)
		if matchConditionsRaw["DestinationRouteTableIds"] != nil {
			destinationRouteTableIdsRaw = convertToInterfaceArray(matchConditionsRaw["DestinationRouteTableIds"])
		}

		matchConditionsMap["destination_route_table_ids"] = destinationRouteTableIdsRaw
		routeTypesRaw := make([]interface{}, 0)
		if matchConditionsRaw["RouteTypes"] != nil {
			routeTypesRaw = convertToInterfaceArray(matchConditionsRaw["RouteTypes"])
		}

		matchConditionsMap["route_types"] = routeTypesRaw
		sourceInstanceIdsRaw := make([]interface{}, 0)
		if matchConditionsRaw["SourceInstanceIds"] != nil {
			sourceInstanceIdsRaw = convertToInterfaceArray(matchConditionsRaw["SourceInstanceIds"])
		}

		matchConditionsMap["source_instance_ids"] = sourceInstanceIdsRaw
		sourceInstanceTypesRaw := make([]interface{}, 0)
		if matchConditionsRaw["SourceInstanceTypes"] != nil {
			sourceInstanceTypesRaw = convertToInterfaceArray(matchConditionsRaw["SourceInstanceTypes"])
		}

		matchConditionsMap["source_instance_types"] = sourceInstanceTypesRaw
		sourceRegionIdsRaw := make([]interface{}, 0)
		if matchConditionsRaw["SourceRegionIds"] != nil {
			sourceRegionIdsRaw = convertToInterfaceArray(matchConditionsRaw["SourceRegionIds"])
		}

		matchConditionsMap["source_region_ids"] = sourceRegionIdsRaw
		sourceRouteTableIdsRaw := make([]interface{}, 0)
		if matchConditionsRaw["SourceRouteTableIds"] != nil {
			sourceRouteTableIdsRaw = convertToInterfaceArray(matchConditionsRaw["SourceRouteTableIds"])
		}

		matchConditionsMap["source_route_table_ids"] = sourceRouteTableIdsRaw
		matchConditionsMaps = append(matchConditionsMaps, matchConditionsMap)
	}
	if err := d.Set("match_conditions", matchConditionsMaps); err != nil {
		return err
	}
	setActionsMaps := make([]map[string]interface{}, 0)
	setActionsMap := make(map[string]interface{})
	setActionsRaw := make(map[string]interface{})
	if objectRaw["SetActions"] != nil {
		setActionsRaw = objectRaw["SetActions"].(map[string]interface{})
	}
	if len(setActionsRaw) > 0 {
		setActionsMap["next_priority"] = setActionsRaw["NextPriority"]
		setActionsMap["route_action"] = setActionsRaw["RouteAction"]

		asPathPrependRaw := make([]interface{}, 0)
		if setActionsRaw["AsPathPrepend"] != nil {
			asPathPrependRaw = convertToInterfaceArray(setActionsRaw["AsPathPrepend"])
		}

		setActionsMap["as_path_prepend"] = asPathPrependRaw
		asPathReplaceRaw := make([]interface{}, 0)
		if setActionsRaw["AsPathReplace"] != nil {
			asPathReplaceRaw = convertToInterfaceArray(setActionsRaw["AsPathReplace"])
		}

		setActionsMap["as_path_replace"] = asPathReplaceRaw
		communityAddRaw := make([]interface{}, 0)
		if setActionsRaw["CommunityAdd"] != nil {
			communityAddRaw = convertToInterfaceArray(setActionsRaw["CommunityAdd"])
		}

		setActionsMap["community_add"] = communityAddRaw
		communityReplaceRaw := make([]interface{}, 0)
		if setActionsRaw["CommunityReplace"] != nil {
			communityReplaceRaw = convertToInterfaceArray(setActionsRaw["CommunityReplace"])
		}

		setActionsMap["community_replace"] = communityReplaceRaw
		setActionsMaps = append(setActionsMaps, setActionsMap)
	}
	if err := d.Set("set_actions", setActionsMaps); err != nil {
		return err
	}

	return nil
}

func resourceAliCloudCenRouteMapRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false

	var err error
	action := "UpdateRouteMapRule"
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["RouteMapRuleId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)
	if d.HasChange("match_conditions") {
		update = true
		matchConditions := make(map[string]interface{})

		if v := d.Get("match_conditions"); v != nil {
			sourceRouteTableIds1, _ := jsonpath.Get("$[0].source_route_table_ids", v)
			if sourceRouteTableIds1 != nil && sourceRouteTableIds1 != "" {
				matchConditions["SourceRouteTableIds"] = sourceRouteTableIds1
			}
			communitySetInclude1, _ := jsonpath.Get("$[0].community_set_include", v)
			if communitySetInclude1 != nil && communitySetInclude1 != "" {
				matchConditions["CommunitySetInclude"] = communitySetInclude1
			}
			destinationRouteTableIds1, _ := jsonpath.Get("$[0].destination_route_table_ids", v)
			if destinationRouteTableIds1 != nil && destinationRouteTableIds1 != "" {
				matchConditions["DestinationRouteTableIds"] = destinationRouteTableIds1
			}
			sourceInstanceIdsReverseMatch1, _ := jsonpath.Get("$[0].source_instance_ids_reverse_match", v)
			if sourceInstanceIdsReverseMatch1 != nil && sourceInstanceIdsReverseMatch1 != "" {
				matchConditions["SourceInstanceIdsReverseMatch"] = sourceInstanceIdsReverseMatch1
			}
			asPathsMatch1, _ := jsonpath.Get("$[0].as_paths_match", v)
			if asPathsMatch1 != nil && asPathsMatch1 != "" {
				matchConditions["AsPathsMatch"] = asPathsMatch1
			}
			destinationInstanceIdsReverseMatch1, _ := jsonpath.Get("$[0].destination_instance_ids_reverse_match", v)
			if destinationInstanceIdsReverseMatch1 != nil && destinationInstanceIdsReverseMatch1 != "" {
				matchConditions["DestinationInstanceIdsReverseMatch"] = destinationInstanceIdsReverseMatch1
			}
			routeTypes1, _ := jsonpath.Get("$[0].route_types", v)
			if routeTypes1 != nil && routeTypes1 != "" {
				matchConditions["RouteTypes"] = routeTypes1
			}
			destinationInstanceIds1, _ := jsonpath.Get("$[0].destination_instance_ids", v)
			if destinationInstanceIds1 != nil && destinationInstanceIds1 != "" {
				matchConditions["DestinationInstanceIds"] = destinationInstanceIds1
			}
			sourceRegionIds1, _ := jsonpath.Get("$[0].source_region_ids", v)
			if sourceRegionIds1 != nil && sourceRegionIds1 != "" {
				matchConditions["SourceRegionIds"] = sourceRegionIds1
			}
			matchAddressType1, _ := jsonpath.Get("$[0].match_address_type", v)
			if matchAddressType1 != nil && matchAddressType1 != "" {
				matchConditions["MatchAddressType"] = matchAddressType1
			}
			addressPrefixesMatch1, _ := jsonpath.Get("$[0].address_prefixes_match", v)
			if addressPrefixesMatch1 != nil && addressPrefixesMatch1 != "" {
				matchConditions["AddressPrefixesMatch"] = addressPrefixesMatch1
			}
			addressPrefixesInclude1, _ := jsonpath.Get("$[0].address_prefixes_include", v)
			if addressPrefixesInclude1 != nil && addressPrefixesInclude1 != "" {
				matchConditions["AddressPrefixesInclude"] = addressPrefixesInclude1
			}
			destinationInstanceTypes1, _ := jsonpath.Get("$[0].destination_instance_types", v)
			if destinationInstanceTypes1 != nil && destinationInstanceTypes1 != "" {
				matchConditions["DestinationInstanceTypes"] = destinationInstanceTypes1
			}
			destinationRegionIds1, _ := jsonpath.Get("$[0].destination_region_ids", v)
			if destinationRegionIds1 != nil && destinationRegionIds1 != "" {
				matchConditions["DestinationRegionIds"] = destinationRegionIds1
			}
			asPathsInclude1, _ := jsonpath.Get("$[0].as_paths_include", v)
			if asPathsInclude1 != nil && asPathsInclude1 != "" {
				matchConditions["AsPathsInclude"] = asPathsInclude1
			}
			sourceInstanceTypes1, _ := jsonpath.Get("$[0].source_instance_types", v)
			if sourceInstanceTypes1 != nil && sourceInstanceTypes1 != "" {
				matchConditions["SourceInstanceTypes"] = sourceInstanceTypes1
			}
			communitySetMatch1, _ := jsonpath.Get("$[0].community_set_match", v)
			if communitySetMatch1 != nil && communitySetMatch1 != "" {
				matchConditions["CommunitySetMatch"] = communitySetMatch1
			}
			sourceInstanceIds1, _ := jsonpath.Get("$[0].source_instance_ids", v)
			if sourceInstanceIds1 != nil && sourceInstanceIds1 != "" {
				matchConditions["SourceInstanceIds"] = sourceInstanceIds1
			}

			matchConditionsJson, err := json.Marshal(matchConditions)
			if err != nil {
				return WrapError(err)
			}
			request["MatchConditions"] = string(matchConditionsJson)
		}
	}

	if d.HasChange("priority") {
		update = true
	}
	request["Priority"] = d.Get("priority")
	if d.HasChange("set_actions") {
		update = true
		setActions := make(map[string]interface{})

		if v := d.Get("set_actions"); v != nil {
			communityAdd1, _ := jsonpath.Get("$[0].community_add", v)
			if communityAdd1 != nil && communityAdd1 != "" {
				setActions["CommunityAdd"] = communityAdd1
			}
			routeAction1, _ := jsonpath.Get("$[0].route_action", v)
			if routeAction1 != nil && routeAction1 != "" {
				setActions["RouteAction"] = routeAction1
			}
			nextPriority1, _ := jsonpath.Get("$[0].next_priority", v)
			if nextPriority1 != nil && nextPriority1 != "" {
				setActions["NextPriority"] = nextPriority1
			}
			asPathReplace1, _ := jsonpath.Get("$[0].as_path_replace", v)
			if asPathReplace1 != nil && asPathReplace1 != "" {
				setActions["AsPathReplace"] = asPathReplace1
			}
			asPathPrepend1, _ := jsonpath.Get("$[0].as_path_prepend", v)
			if asPathPrepend1 != nil && asPathPrepend1 != "" {
				setActions["AsPathPrepend"] = asPathPrepend1
			}
			communityReplace1, _ := jsonpath.Get("$[0].community_replace", v)
			if communityReplace1 != nil && communityReplace1 != "" {
				setActions["CommunityReplace"] = communityReplace1
			}

			setActionsJson, err := json.Marshal(setActions)
			if err != nil {
				return WrapError(err)
			}
			request["SetActions"] = string(setActionsJson)
		}
	}

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	if d.HasChange("description") {
		update = true
		request["Description"] = d.Get("description")
	}

	if d.HasChange("name") {
		update = true
		request["Name"] = d.Get("name")
	}

	if update {
		wait := incrementalWait(3*time.Second, 0*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("Cbn", "2017-09-12", action, query, request, true)
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
		cenServiceV2 := CenServiceV2{client}
		stateConf := BuildStateConf([]string{}, []string{"Active"}, d.Timeout(schema.TimeoutUpdate), 5*time.Second, cenServiceV2.CenRouteMapRuleStateRefreshFunc(d.Id(), "Status", []string{}))
		if _, err := stateConf.WaitForState(); err != nil {
			return WrapErrorf(err, IdMsg, d.Id())
		}
	}

	return resourceAliCloudCenRouteMapRuleRead(d, meta)
}

func resourceAliCloudCenRouteMapRuleDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteRouteMapRule"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	var err error
	request = make(map[string]interface{})
	request["RouteMapRuleId"] = d.Id()

	request["ClientToken"] = buildClientToken(action)

	if v, ok := d.GetOkExists("dry_run"); ok {
		request["DryRun"] = v
	}
	wait := incrementalWait(3*time.Second, 0*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("Cbn", "2017-09-12", action, query, request, true)
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

	cenServiceV2 := CenServiceV2{client}
	stateConf := BuildStateConf([]string{}, []string{""}, d.Timeout(schema.TimeoutDelete), 5*time.Second, cenServiceV2.CenRouteMapRuleStateRefreshFunc(d.Id(), "Status", []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	return nil
}
