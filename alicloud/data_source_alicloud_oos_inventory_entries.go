package alicloud

import (
	"fmt"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

var oosInventoryTypeNames = []string{
	"ACS:InstanceInformation",
	"ACS:Application",
	"ACS:File",
	"ACS:Network",
	"ACS:WindowsRole",
	"ACS:Service",
	"ACS:WindowsRegistry",
	"ACS:WindowsUpdate",
}

var oosInventoryFilterOperators = []string{
	"Equal",
	"NotEqual",
	"BeginWith",
	"LessThan",
	"GreaterThan",
}

func dataSourceAlicloudOosInventoryEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudOosInventoryEntriesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice(oosInventoryTypeNames, false),
			},
			"filter":      oosInventoryFilterSchema(),
			"output_file": oosInventoryOutputFileSchema(),
			"schema_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"capture_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entries_json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAlicloudOosInventoryEntriesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "ListInventoryEntries"
	dataSourceName := "alicloud_oos_inventory_entries"

	request := map[string]interface{}{
		"RegionId":   client.RegionId,
		"MaxResults": PageSizeLarge,
		"InstanceId": d.Get("instance_id"),
		"TypeName":   d.Get("type_name"),
	}
	if v, ok := d.GetOk("filter"); ok {
		expandOosInventoryFilters(request, v.([]interface{}))
	}

	objects, response, err := listOosInventoryObjects(client, dataSourceName, action, request, "$.Entries")
	if err != nil {
		return WrapError(err)
	}

	entries, err := flattenOosInventoryJsonList(objects)
	if err != nil {
		return WrapError(err)
	}
	entriesJson, err := convertInterfaceToJsonString(objects)
	if err != nil {
		return WrapError(err)
	}

	d.SetId(dataResourceIdHash([]string{client.RegionId, fmt.Sprint(d.Get("instance_id")), fmt.Sprint(d.Get("type_name"))}))
	if err := d.Set("schema_version", oosInventoryString(response["SchemaVersion"])); err != nil {
		return WrapError(err)
	}
	if err := d.Set("capture_time", oosInventoryString(response["CaptureTime"])); err != nil {
		return WrapError(err)
	}
	if err := d.Set("entries", entries); err != nil {
		return WrapError(err)
	}
	if err := d.Set("entries_json", entriesJson); err != nil {
		return WrapError(err)
	}
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), objects)
	}

	return nil
}

func dataSourceAlicloudOosInventorySchema() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudOosInventorySchemaRead,
		Schema: map[string]*schema.Schema{
			"type_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(oosInventoryTypeNames, false),
			},
			"aggregator": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"output_file": oosInventoryOutputFileSchema(),
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"schemas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"attributes_json": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"schema_json": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"schemas_json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAlicloudOosInventorySchemaRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "GetInventorySchema"
	dataSourceName := "alicloud_oos_inventory_schema"

	request := map[string]interface{}{
		"RegionId":   client.RegionId,
		"MaxResults": PageSizeLarge,
	}
	if v, ok := d.GetOk("type_name"); ok {
		request["TypeName"] = v
	}
	if v, ok := d.GetOk("aggregator"); ok {
		request["Aggregator"] = v
	}

	objects, _, err := listOosInventoryObjects(client, dataSourceName, action, request, "$.Schemas")
	if err != nil {
		return WrapError(err)
	}

	ids := make([]string, 0, len(objects))
	schemas := make([]map[string]interface{}, 0, len(objects))
	for _, object := range objects {
		schemaObject := object.(map[string]interface{})
		attributesRaw, _ := schemaObject["Attributes"].([]interface{})
		attributes := make([]map[string]interface{}, 0, len(attributesRaw))
		for _, attribute := range attributesRaw {
			attributeObject := attribute.(map[string]interface{})
			attributes = append(attributes, map[string]interface{}{
				"name":      oosInventoryString(attributeObject["Name"]),
				"data_type": oosInventoryString(attributeObject["DataType"]),
			})
		}

		typeName := oosInventoryString(schemaObject["TypeName"])
		attributesJson, err := convertInterfaceToJsonString(attributesRaw)
		if err != nil {
			return WrapError(err)
		}
		schemaJson, err := convertInterfaceToJsonString(schemaObject)
		if err != nil {
			return WrapError(err)
		}

		ids = append(ids, typeName)
		schemas = append(schemas, map[string]interface{}{
			"id":              typeName,
			"type_name":       typeName,
			"version":         oosInventoryString(schemaObject["Version"]),
			"attributes":      attributes,
			"attributes_json": attributesJson,
			"schema_json":     schemaJson,
		})
	}

	schemasJson, err := convertInterfaceToJsonString(objects)
	if err != nil {
		return WrapError(err)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return WrapError(err)
	}
	if err := d.Set("schemas", schemas); err != nil {
		return WrapError(err)
	}
	if err := d.Set("schemas_json", schemasJson); err != nil {
		return WrapError(err)
	}
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), schemas)
	}

	return nil
}

func dataSourceAlicloudOosInventorySearch() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudOosInventorySearchRead,
		Schema: map[string]*schema.Schema{
			"filter": oosInventoryFilterSchema(),
			"aggregators": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"output_file": oosInventoryOutputFileSchema(),
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entities_json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAlicloudOosInventorySearchRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "SearchInventory"
	dataSourceName := "alicloud_oos_inventory_search"

	request := map[string]interface{}{
		"RegionId":   client.RegionId,
		"MaxResults": PageSizeLarge,
	}
	if v, ok := d.GetOk("filter"); ok {
		expandOosInventoryFilters(request, v.([]interface{}))
	}
	if v, ok := d.GetOk("aggregators"); ok {
		expandOosInventoryAggregators(request, v.([]interface{}))
	}

	objects, _, err := listOosInventoryObjects(client, dataSourceName, action, request, "$.Entities")
	if err != nil {
		return WrapError(err)
	}

	ids := make([]string, 0, len(objects))
	for i, object := range objects {
		objectMap, ok := object.(map[string]interface{})
		if !ok {
			ids = append(ids, fmt.Sprint(i))
			continue
		}
		id := oosInventoryString(objectMap["Id"])
		if id == "" {
			id = fmt.Sprint(i)
		}
		ids = append(ids, id)
	}

	entities, err := flattenOosInventoryJsonList(objects)
	if err != nil {
		return WrapError(err)
	}
	entitiesJson, err := convertInterfaceToJsonString(objects)
	if err != nil {
		return WrapError(err)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return WrapError(err)
	}
	if err := d.Set("entities", entities); err != nil {
		return WrapError(err)
	}
	if err := d.Set("entities_json", entitiesJson); err != nil {
		return WrapError(err)
	}
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), objects)
	}

	return nil
}

func oosInventoryFilterSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 5,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"operator": {
					Type:         schema.TypeString,
					Optional:     true,
					Default:      "Equal",
					ValidateFunc: validation.StringInSlice(oosInventoryFilterOperators, false),
				},
				"values": {
					Type:     schema.TypeList,
					Required: true,
					MinItems: 1,
					MaxItems: 20,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}

func oosInventoryOutputFileSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Optional: true,
	}
}

func expandOosInventoryFilters(request map[string]interface{}, filters []interface{}) {
	for filterPtr, filter := range filters {
		filterArg := filter.(map[string]interface{})
		request[fmt.Sprintf("Filter.%d.Name", filterPtr+1)] = filterArg["name"]
		request[fmt.Sprintf("Filter.%d.Operator", filterPtr+1)] = filterArg["operator"]
		for valuePtr, value := range filterArg["values"].([]interface{}) {
			request[fmt.Sprintf("Filter.%d.Value.%d", filterPtr+1, valuePtr+1)] = value
		}
	}
}

func expandOosInventoryAggregators(request map[string]interface{}, aggregators []interface{}) {
	for aggregatorPtr, aggregator := range aggregators {
		request[fmt.Sprintf("Aggregator.%d", aggregatorPtr+1)] = aggregator
	}
}

func listOosInventoryObjects(client *connectivity.AliyunClient, dataSourceName, action string, request map[string]interface{}, jsonPath string) ([]interface{}, map[string]interface{}, error) {
	objects := make([]interface{}, 0)
	var response map[string]interface{}

	for {
		resp, err := callOosInventoryApi(client, dataSourceName, action, request)
		if err != nil {
			return nil, nil, err
		}
		response = resp

		resultRaw, err := jsonpath.Get(jsonPath, response)
		if err != nil {
			return nil, nil, WrapErrorf(err, FailedGetAttributeMsg, action, jsonPath, response)
		}
		if result, ok := resultRaw.([]interface{}); ok {
			objects = append(objects, result...)
		}

		if nextToken, ok := response["NextToken"].(string); ok && nextToken != "" {
			request["NextToken"] = nextToken
		} else {
			break
		}
	}

	return objects, response, nil
}

func callOosInventoryApi(client *connectivity.AliyunClient, dataSourceName, action string, request map[string]interface{}) (map[string]interface{}, error) {
	var response map[string]interface{}
	var err error
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("oos", "2019-06-01", action, nil, request, true)
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
		return nil, WrapErrorf(err, DataDefaultErrorMsg, dataSourceName, action, AlibabaCloudSdkGoERROR)
	}

	return response, nil
}

func flattenOosInventoryJsonList(objects []interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0, len(objects))
	for _, object := range objects {
		jsonString, err := convertInterfaceToJsonString(object)
		if err != nil {
			return nil, err
		}
		result = append(result, jsonString)
	}
	return result, nil
}

func oosInventoryString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v)
}
