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

func resourceAliCloudEventBridgeTable() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudEventBridgeTableCreate,
                Read: resourceAliCloudEventBridgeTableRead,
                Update: resourceAliCloudEventBridgeTableUpdate,
            Delete: resourceAliCloudEventBridgeTableDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Timeouts: &schema.ResourceTimeout{
              Create: schema.DefaultTimeout(5 * time.Minute),
                  Update: schema.DefaultTimeout(5 * time.Minute),
                  Delete: schema.DefaultTimeout(5 * time.Minute),
        },
            Schema: map[string]*schema.Schema {
                        "catalog": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                                                                            },
                                    "columns": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                                MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "comment": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "type": {
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
                                    "comment": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                        "name": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                                    "namespace": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                        ForceNew: true,
                                                                                                                                                            },
                                        "retention_policy": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                        ForceNew: true,
                                                                                                                                                                MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "hot_ttl": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "cold_ttl": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                },
                        },
                },
  }
}

func resourceAliCloudEventBridgeTableCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := "CreateTable"
            var request map[string]interface{}
        var response map[string]interface{}
                            query := make(map[string]interface{})
                    var err error
        request = make(map[string]interface{})
                if v, ok := d.GetOk("namespace"); ok {
request["Namespace"] = v
}
            if v, ok := d.GetOk("catalog"); ok {
request["Catalog"] = v
}
            if v, ok := d.GetOk("name"); ok {
request["Name"] = v
}
            
                        request["ClientToken"] = buildClientToken(action)
            
                                                                if v, ok := d.GetOk("comment");  ok    {
                                                            request["Comment"] = v                                        }
                                                                                                                            retentionPolicy := make(map[string]interface{})
    
if v := d.Get("retention_policy");  !IsNil(v)   {
                                coldTTL1, _ := jsonpath.Get("$[0].cold_ttl", v)
        if coldTTL1 != nil                && coldTTL1 != "" {
                retentionPolicy["ColdTTL"] = coldTTL1        }
                                            hotTTL1, _ := jsonpath.Get("$[0].hot_ttl", v)
        if hotTTL1 != nil                && hotTTL1 != "" {
                retentionPolicy["HotTTL"] = hotTTL1        }
            
    retentionPolicyJson, err := json.Marshal(retentionPolicy)
    if err != nil {
    return WrapError(err)
    }
    request["RetentionPolicy"] = string(retentionPolicyJson)
}

                                                                                                                                columnsDataList := make(map[string]interface{})

                                                                                                if v, ok := d.GetOk("columns"); ok {
                                                                                                    type1, _ := jsonpath.Get("$[0].type", v)
        if type1 != nil                && type1 != "" {
                columnsDataList["Type"] = type1        }
                                                                }
                            
                                                                                                if v, ok := d.GetOk("columns"); ok {
                                                                                                    name1, _ := jsonpath.Get("$[0].name", v)
        if name1 != nil                && name1 != "" {
                columnsDataList["Name"] = name1        }
                                                                }
                            
                                                                                                if v, ok := d.GetOk("columns"); ok {
                                                                                                    comment1, _ := jsonpath.Get("$[0].comment", v)
        if comment1 != nil                && comment1 != "" {
                columnsDataList["Comment"] = comment1        }
                                                                }
                            
    ColumnsMap := make([]interface{}, 0)
    ColumnsMap = append(ColumnsMap, columnsDataList)
            columnsDataListJson, err := json.Marshal(ColumnsMap)
        if err != nil {
            return WrapError(err)
        }
        request["Columns"] = string(columnsDataListJson)
    
                                                                wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
                        response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
                    if err != nil {
                        if  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
        }
        return nil
    })
    addDebug(action, response, request)

    if err != nil {
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_event_bridge_table", action, AlibabaCloudSdkGoERROR)
    }
        
    
                                                                            DataTableARNVar, _ := jsonpath.Get("$.Data.TableARN", response)
                            d.SetId(fmt.Sprintf("%v:%v:%v",request["Catalog"],request["Namespace"],DataTableARNVar));
    

            
        
        
        return resourceAliCloudEventBridgeTableRead(d, meta)
    }


    		func resourceAliCloudEventBridgeTableRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		eventBridgeServiceV2 := EventBridgeServiceV2{client}

												objectRaw, err := eventBridgeServiceV2.DescribeEventBridgeTable(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_event_bridge_table DescribeEventBridgeTable Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("comment", objectRaw["Comment"])
                                            d.Set("catalog", objectRaw["Catalog"])
                                            d.Set("name", objectRaw["Name"])
                                            d.Set("namespace", objectRaw["Namespace"])
                
                                                                    columnsMaps := make([]map[string]interface{}, 0)
        columnsMap := make(map[string]interface{})
                columnsChildRaw := columnsChildRaw.(map[string]interface{})
                                        columnsMap["comment"] = columnsChildRaw["Comment"]
                                            columnsMap["name"] = columnsChildRaw["Name"]
                                            columnsMap["type"] = columnsChildRaw["Type"]
                
                                                                                                                                                            columnsMaps = append(columnsMaps, columnsMap)
}
if err := d.Set("columns", columnsMaps); err != nil {
return err
}
                                                                                                                                                    retentionPolicyMaps := make([]map[string]interface{}, 0)
        retentionPolicyMap := make(map[string]interface{})
                retentionPolicyRaw := make(map[string]interface{})
if objectRaw["RetentionPolicy"] != nil {
retentionPolicyRaw = objectRaw["RetentionPolicy"].(map[string]interface{})}
if len(retentionPolicyRaw) > 0 {
                                        retentionPolicyMap["cold_ttl"] = retentionPolicyRaw["ColdTTL"]
                                            retentionPolicyMap["hot_ttl"] = retentionPolicyRaw["HotTTL"]
                
                                                                                                        retentionPolicyMaps = append(retentionPolicyMaps, retentionPolicyMap)
}
if err := d.Set("retention_policy", retentionPolicyMaps); err != nil {
return err
}
                                                                                                                                                                                        
							
        
        
		        	return nil
}


    func resourceAliCloudEventBridgeTableUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
            var query map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                parts := strings.Split(d.Id(), ":")
            action := "UpdateTable"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["Namespace"] = parts[1]
            request["Catalog"] = parts[0]
            request["Name"] = parts[2]
            
                request["ClientToken"] = buildClientToken(action)
                                                                                                if d.HasChange("comment") {
        update = true
                    request["UpdateComment"] = d.Get("comment")}

                                                                                                                                                                            if  d.HasChange("retention_policy") {
    update = true
                updateRetentionPolicy := make(map[string]interface{})
    
if v := d.Get("retention_policy");  v != nil  {
                                coldTTL1, _ := jsonpath.Get("$[0].cold_ttl", v)
        if coldTTL1 != nil&&  coldTTL1 != ""  {
                            updateRetentionPolicy["ColdTTL"] = coldTTL1                }
                                            hotTTL1, _ := jsonpath.Get("$[0].hot_ttl", v)
        if hotTTL1 != nil&&  hotTTL1 != ""  {
                            updateRetentionPolicy["HotTTL"] = hotTTL1                }
            
    updateRetentionPolicyJson, err := json.Marshal(updateRetentionPolicy)
    if err != nil {
    return WrapError(err)
    }
    request["UpdateRetentionPolicy"] = string(updateRetentionPolicyJson)
}
    }

                                                                                                                                            if  d.HasChange("columns") {
    update = true
                updateColumnType := make(map[string]interface{})
    
if v := d.Get("columns");  v != nil  {
                                type1, _ := jsonpath.Get("$[0].type", v)
        if type1 != nil&&  type1 != ""  {
                            updateColumnType["Type"] = type1                }
                                            name1, _ := jsonpath.Get("$[0].name", v)
        if name1 != nil&&  name1 != ""  {
                            updateColumnType["Name"] = name1                }
            
    updateColumnTypeJson, err := json.Marshal(updateColumnType)
    if err != nil {
    return WrapError(err)
    }
    request["UpdateColumnType"] = string(updateColumnTypeJson)
}
    }

                                                                                                                                            if  d.HasChange("columns") {
    update = true
                updateColumnComment := make(map[string]interface{})
    
if v := d.Get("columns");  v != nil  {
                                comment1, _ := jsonpath.Get("$[0].comment", v)
        if comment1 != nil&&  comment1 != ""  {
                            updateColumnComment["Comment"] = comment1                }
            
    updateColumnCommentJson, err := json.Marshal(updateColumnComment)
    if err != nil {
    return WrapError(err)
    }
    request["UpdateColumnComment"] = string(updateColumnCommentJson)
}
    }

                                                            if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
                                        if err != nil {
                if  NeedRetry(err) {
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
                
                    return resourceAliCloudEventBridgeTableRead(d, meta)
}


func resourceAliCloudEventBridgeTableDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                        client := meta.(*connectivity.AliyunClient)
                    parts := strings.Split(d.Id(), ":")
            action := "DeleteTable"
            var request map[string]interface{}
        var response map[string]interface{}
                                        query := make(map[string]interface{})
                        var err error
        request = make(map[string]interface{})
                request["Namespace"] = parts[1]
            request["Catalog"] = parts[0]
            request["Name"] = parts[2]
            
                        request["ClientToken"] = buildClientToken(action)
            

                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
                        response, err = client.RpcPost("eventbridge", "2020-04-01", action, query, request, true)
                    if err != nil {
            if  NeedRetry(err) {
                wait()
                return resource.RetryableError(err)
            }
                        return resource.NonRetryableError(err)
       }
       return nil
    })
    addDebug(action, response, request)

    if err != nil {
        if  NotFoundError(err) {
            return nil
        }
        return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
    }

        
                        return nil
}




