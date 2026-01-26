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

func resourceAliCloudCloudSsoGroup() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudCloudSsoGroupCreate,
                Read: resourceAliCloudCloudSsoGroupRead,
                Update: resourceAliCloudCloudSsoGroupUpdate,
            Delete: resourceAliCloudCloudSsoGroupDelete,
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
                                    "description": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                Sensitive: true,
                                                                            ValidateFunc: StringMatch(regexp.MustCompile("^[\s\S]{0,1024}$"), "Description"),
                                                                                    },
                                    "directory_id": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                                    "group_id": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "group_name": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                Sensitive: true,
                                                                            ValidateFunc: StringMatch(regexp.MustCompile("^[0-9a-zA-Z.\\\-_]{1,128}$"), "GroupName"),
                                                                                    },
                                    },
  }
}

func resourceAliCloudCloudSsoGroupCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := "CreateGroup"
            var request map[string]interface{}
        var response map[string]interface{}
                            query := make(map[string]interface{})
                    var err error
        request = make(map[string]interface{})
                if v, ok := d.GetOk("directory_id"); ok {
request["DirectoryId"] = v
}
            
    
                                                                if v, ok := d.GetOk("description");  ok    {
                                                            request["Description"] = v                                        }
                                                                                                request["GroupName"] = d.Get("group_name")
                                                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
                        response, err = client.RpcPost("cloudsso", "2021-05-15", action, query, request, true)
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
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_cloud_sso_group", action, AlibabaCloudSdkGoERROR)
    }
        
    
                                                        GroupGroupIdVar, _ := jsonpath.Get("$.Group.GroupId", response)
                            d.SetId(fmt.Sprintf("%v:%v",request["DirectoryId"],GroupGroupIdVar));
    

            
        
        
        return resourceAliCloudCloudSsoGroupRead(d, meta)
    }


    		func resourceAliCloudCloudSsoGroupRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		cloudSSOServiceV2 := CloudSSOServiceV2{client}

												objectRaw, err := cloudSSOServiceV2.DescribeCloudSsoGroup(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_cloud_sso_group DescribeCloudSsoGroup Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("create_time", objectRaw["CreateTime"])
                                            d.Set("description", objectRaw["Description"])
                                            d.Set("group_name", objectRaw["GroupName"])
                                            d.Set("group_id", objectRaw["GroupId"])
                
                                                                                                                                                                                                                
							
        		parts := strings.Split(d.Id(), ":")
					d.Set("directory_id", parts[0])
			
        
		        	return nil
}


    func resourceAliCloudCloudSsoGroupUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
            var query map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                parts := strings.Split(d.Id(), ":")
            action := "UpdateGroup"
    request = make(map[string]interface{})
                            query = make(map[string]interface{})
                            request["GroupId"] = parts[1]
            request["DirectoryId"] = parts[0]
            
                                                                                        if d.HasChange("description") {
        update = true
                    request["NewDescription"] = d.Get("description")}

                                                                                                                                        if d.HasChange("group_name") {
        update = true
        }
                    request["NewGroupName"] = d.Get("group_name")
                                                                                            if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                                    response, err = client.RpcPost("cloudsso", "2021-05-15", action, query, request, true)
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
                
                    return resourceAliCloudCloudSsoGroupRead(d, meta)
}


func resourceAliCloudCloudSsoGroupDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                        client := meta.(*connectivity.AliyunClient)
                    parts := strings.Split(d.Id(), ":")
            action := "DeleteGroup"
            var request map[string]interface{}
        var response map[string]interface{}
                                        query := make(map[string]interface{})
                        var err error
        request = make(map[string]interface{})
                request["GroupId"] = parts[1]
            request["DirectoryId"] = parts[0]
            
    

                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
                        response, err = client.RpcPost("cloudsso", "2021-05-15", action, query, request, true)
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
        if IsExpectedErrors(err, []string{"EntityNotExists.Group"}) ||  NotFoundError(err) {
            return nil
        }
        return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
    }

        
                        return nil
}




