// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
  "encoding/json"
  "fmt"
  "log"
  "time"
  "strings"

  util "github.com/alibabacloud-go/tea-utils/service"

    "github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudFnFFlow() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudFnFFlowCreate,
                Read: resourceAliCloudFnFFlowRead,
                Update: resourceAliCloudFnFFlowUpdate,
            Delete: resourceAliCloudFnFFlowDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Timeouts: &schema.ResourceTimeout{
              Create: schema.DefaultTimeout(5 * time.Minute),
                  Update: schema.DefaultTimeout(5 * time.Minute),
                  Delete: schema.DefaultTimeout(5 * time.Minute),
        },
            Schema: map[string]*schema.Schema {
                        "create_time": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "definition": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                                                                                                                        },
                                    "description": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                                                                                                                        },
                                    "environment": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "variables": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                                                                                                                Elem: &schema.Resource{
                                        Schema: map[string]*schema.Schema {
                        "description": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "value": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "name": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                    },
                                        },
            },
                },
                        },
                                    "flow_id": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "flow_name": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                                        "last_modified_time": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                            "resource_group_id": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                    "role_arn": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                    "tracing_configuration": {
            Type:schema.TypeList,
                                                                                                            Computed: true,
                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "enabled": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                },
                        },
                                    "type": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                                                                ValidateFunc: StringInSlice([]string{"FDL","DEFAULT"}, false),
                                                                                    },
            },
  }
}

func resourceAliCloudFnFFlowCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := "CreateFlow"
            var request map[string]interface{}
        var response map[string]interface{}
                            query := make(map[string]interface{})
                    var err error
        request = make(map[string]interface{})
                if v, ok := d.GetOk("flow_name"); ok {
request["Name"] = v
}
            
    
                                                                if v, ok := d.GetOk("role_arn");  ok    {
                                                            request["RoleArn"] = v                                        }
                                                                                if v, ok := d.GetOk("resource_group_id");  ok    {
                                                            request["ResourceGroupId"] = v                                        }
                                                                                                request["Description"] = d.Get("description")
                                                                                                                            environment := make(map[string]interface{})
    
if v := d.Get("environment");  !IsNil(v)   {
                                                        localData, err := jsonpath.Get("$[0].variables", v)
        if err != nil {
            localData = make([]interface{}, 0)
        }
    localMaps := make([]interface{}, 0)
for _, dataLoop := range convertToInterfaceArray(localData) {
    dataLoopTmp := make(map[string]interface{})
    if dataLoop != nil {
        dataLoopTmp = dataLoop.(map[string]interface{})
    }
    dataLoopMap := make(map[string]interface{})
    	                        dataLoopMap["Description"] = dataLoopTmp["description"]
                                                            dataLoopMap["Name"] = dataLoopTmp["name"]
                                                            dataLoopMap["Value"] = dataLoopTmp["value"]
                                    		localMaps = append(localMaps, dataLoopMap)
}
    environment["Variables"] = localMaps

    
    environmentJson, err := json.Marshal(environment)
    if err != nil {
    return WrapError(err)
    }
    request["Environment"] = string(environmentJson)
}

                                                                                            request["Definition"] = d.Get("definition")
                                                                                                request["Type"] = d.Get("type")
                                                    wait := incrementalWait(1*time.Second, 1*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
                        response, err = client.RpcPost("fnf", "2019-03-15", action, query, request, true)
                    if err != nil {
                        if  IsExpectedErrors(err, []string{"InternalServerError"}) ||  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
        }
        return nil
    })
    addDebug(action, response, request)

    if err != nil {
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_fnf_flow", action, AlibabaCloudSdkGoERROR)
    }
        
                        d.SetId(fmt.Sprint(response["Name"]))        

            
        
        
        return resourceAliCloudFnFFlowRead(d, meta)
    }


    		func resourceAliCloudFnFFlowRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		fnFServiceV2 := FnFServiceV2{client}

												objectRaw, err := fnFServiceV2.DescribeFnFFlow(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_fnf_flow DescribeFnFFlow Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("create_time", objectRaw["CreatedTime"])
                                            d.Set("definition", objectRaw["Definition"])
                                            d.Set("description", objectRaw["Description"])
                                            d.Set("flow_id", objectRaw["Id"])
                                            d.Set("last_modified_time", objectRaw["LastModifiedTime"])
                                            d.Set("resource_group_id", objectRaw["ResourceGroupId"])
                                            d.Set("role_arn", objectRaw["RoleArn"])
                                            d.Set("type", objectRaw["Type"])
                                            d.Set("flow_name", objectRaw["Name"])
                
                                                                                                                                                                                                                                environmentMaps := make([]map[string]interface{}, 0)
        environmentMap := make(map[string]interface{})
                environmentRaw := make(map[string]interface{})
if objectRaw["Environment"] != nil {
environmentRaw = objectRaw["Environment"].(map[string]interface{})}
if len(environmentRaw) > 0 {
            
                                                                                                                                                variablesRaw := environmentRaw["Variables"]
                                    variablesMaps := make([]map[string]interface{}, 0)
                                                if variablesRaw != nil {
for _, variablesChildRaw := range convertToInterfaceArray(variablesRaw) {
                                variablesMap := make(map[string]interface{})
                                                        variablesChildRaw := variablesChildRaw.(map[string]interface{})
                                        variablesMap["description"] = variablesChildRaw["Description"]
                                            variablesMap["name"] = variablesChildRaw["Name"]
                                            variablesMap["value"] = variablesChildRaw["Value"]
                
                                                                                                                                                                                            variablesMaps = append(variablesMaps, variablesMap)
            }
                        }
                                                environmentMap["variables"] = variablesMaps
            environmentMaps = append(environmentMaps, environmentMap)
}
if err := d.Set("environment", environmentMaps); err != nil {
return err
}
                                                                                                                                                                                                                                                                                                                                                                                                                        tracingConfigurationMaps := make([]map[string]interface{}, 0)
        tracingConfigurationMap := make(map[string]interface{})
                
                                        tracingConfigurationMap["enabled"] = tracingConfigurationRaw["Enabled"]
                
                                                    tracingConfigurationMaps = append(tracingConfigurationMaps, tracingConfigurationMap)
}
if err := d.Set("tracing_configuration", tracingConfigurationMaps); err != nil {
return err
}
                            
							
        
        		d.Set("flow_name", d.Id())
	
		        	return nil
}


    func resourceAliCloudFnFFlowUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
            var query map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                action := "UpdateFlow"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["Name"] = d.Id()
            
                                                                                        if d.HasChange("role_arn") {
        update = true
                    request["RoleArn"] = d.Get("role_arn")}

                                                                                                                                        if d.HasChange("description") {
        update = true
        }
                    request["Description"] = d.Get("description")
                                                                                                                                                                            if  d.HasChange("environment") {
    update = true
                environment := make(map[string]interface{})
    
if v := d.Get("environment");  v != nil  {
                                                        localData, err := jsonpath.Get("$[0].variables", v)
        if err != nil {
            localData = make([]interface{}, 0)
        }
    localMaps := make([]interface{}, 0)
for _, dataLoop := range convertToInterfaceArray(localData) {
    dataLoopTmp := make(map[string]interface{})
    if dataLoop != nil {
        dataLoopTmp = dataLoop.(map[string]interface{})
    }
    dataLoopMap := make(map[string]interface{})
    	                        dataLoopMap["Description"] = dataLoopTmp["description"]
                                                            dataLoopMap["Name"] = dataLoopTmp["name"]
                                                            dataLoopMap["Value"] = dataLoopTmp["value"]
                                    		localMaps = append(localMaps, dataLoopMap)
}
    environment["Variables"] = localMaps

    
    environmentJson, err := json.Marshal(environment)
    if err != nil {
    return WrapError(err)
    }
    request["Environment"] = string(environmentJson)
}
    }

                                                                                                        if d.HasChange("definition") {
        update = true
        }
                    request["Definition"] = d.Get("definition")
                                                                                                                                        if d.HasChange("type") {
        update = true
        }
                    request["Type"] = d.Get("type")
                                                                                            if update  {
            wait := incrementalWait(1*time.Second, 1*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("fnf", "2019-03-15", action, query, request, true)
                                        if err != nil {
                if  IsExpectedErrors(err, []string{"ConcurrentUpdateError","InternalServerError"}) ||  NeedRetry(err) {
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
                
                    return resourceAliCloudFnFFlowRead(d, meta)
}


func resourceAliCloudFnFFlowDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                        client := meta.(*connectivity.AliyunClient)
                    action := "DeleteFlow"
            var request map[string]interface{}
        var response map[string]interface{}
                                        query := make(map[string]interface{})
                        var err error
        request = make(map[string]interface{})
                request["Name"] = d.Id()
            
    

                    wait := incrementalWait(1*time.Second, 1*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
                        response, err = client.RpcPost("fnf", "2019-03-15", action, query, request, true)
                    if err != nil {
            if  IsExpectedErrors(err, []string{"InternalServerError"}) ||  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
       }
       return nil
    })
    addDebug(action, response, request)

    if err != nil {
        if IsExpectedErrors(err, []string{"FlowNotExists"}) ||  NotFoundError(err) {
            return nil
        }
        return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
    }

        
                        return nil
}




