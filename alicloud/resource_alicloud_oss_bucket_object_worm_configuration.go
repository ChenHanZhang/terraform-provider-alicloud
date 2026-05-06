// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
  "encoding/json"
  "fmt"
  "log"
  "time"
  "strings"

  util "github.com/alibabacloud-go/tea-utils/v2/service"
  sls "github.com/aliyun/aliyun-log-go-sdk"

    "github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceAliCloudOssBucketObjectWormConfiguration() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudOssBucketObjectWormConfigurationCreate,
                Read: resourceAliCloudOssBucketObjectWormConfigurationRead,
                Update: resourceAliCloudOssBucketObjectWormConfigurationUpdate,
            Delete: resourceAliCloudOssBucketObjectWormConfigurationDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Timeouts: &schema.ResourceTimeout{
              Create: schema.DefaultTimeout(5 * time.Minute),
                  Update: schema.DefaultTimeout(5 * time.Minute),
                  Delete: schema.DefaultTimeout(5 * time.Minute),
        },
            Schema: map[string]*schema.Schema {
                        "bucket_name": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                                    "object_worm_enabled": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                                                                ValidateFunc: StringInSlice([]string{"Enabled"}, false),
                                                                                    },
                                    "rule": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "default_retention": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                                                                                                                Elem: &schema.Resource{
                                        Schema: map[string]*schema.Schema {
                        "years": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "mode": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                ValidateFunc: StringInSlice([]string{"COMPLIANCE"}, false),
                                                                                    },
                                    "days": {
            Type:schema.TypeInt,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
                    },
                                        },
            },
                },
                        },
            },
  }
}

func resourceAliCloudOssBucketObjectWormConfigurationCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := fmt.Sprintf("/?objectWorm")
            var request map[string]interface{}
        var response map[string]interface{}
                                                    query := make(map[string]*string)
            body := make(map[string]interface{})
                                            hostMap := make(map[string]*string)
                                                var err error
        request = make(map[string]interface{})
                hostMap["bucket"] = StringPointer(d.Get("bucket_name").(string))
            
    
                                                                                                            objectWormConfiguration := make(map[string]interface{})
    
if v := d.Get("object_worm_enabled");  !IsNil(v)   {
                                    objectWormConfiguration["ObjectWormEnabled"] = v                                    rule := make(map[string]interface{})

            if len(rule) > 0 {
        objectWormConfiguration["Rule"] = rule
        }
            
                request["ObjectWormConfiguration"] = objectWormConfiguration
    }

                                                    body = request
        wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
            response, err = client.Do("Oss", xmlParam("PUT", "2019-05-17", "PutBucketObjectWormConfiguration", action), query, body, nil, hostMap, false)
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
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_oss_bucket_object_worm_configuration", action, AlibabaCloudSdkGoERROR)
    }
        
                        d.SetId(fmt.Sprint(*hostMap["bucket"]))        

            
        
        
        return resourceAliCloudOssBucketObjectWormConfigurationRead(d, meta)
    }


    		func resourceAliCloudOssBucketObjectWormConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		ossServiceV2 := OssServiceV2{client}

												objectRaw, err := ossServiceV2.DescribeOssBucketObjectWormConfiguration(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_oss_bucket_object_worm_configuration DescribeOssBucketObjectWormConfiguration Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("object_worm_enabled", objectRaw["ObjectWormEnabled"])
                
                                                                                                                        ruleMaps := make([]map[string]interface{}, 0)
        ruleMap := make(map[string]interface{})
                defaultRetentionChildRaw := make(map[string]interface{})
if defaultRetentionRaw["DefaultRetentionChild"] != nil {
defaultRetentionChildRaw = defaultRetentionRaw["DefaultRetentionChild"].(map[string]interface{})}
if len(defaultRetentionChildRaw) > 0 {
            
                                                                                                                                                defaultRetentionChildRaw := defaultRetentionChildRaw.(map[string]interface{})
                                    defaultRetentionMaps := make([]map[string]interface{}, 0)
                                                if defaultRetentionChildRaw != nil {
for _, daysRaw := range convertToInterfaceArray(defaultRetentionChildRaw) {
                                defaultRetentionMap := make(map[string]interface{})
                                                        defaultRetentionChildRaw := defaultRetentionChildRaw.(map[string]interface{})
                                        defaultRetentionMap["days"] = defaultRetentionChildRaw["Days"]
                                            defaultRetentionMap["mode"] = defaultRetentionChildRaw["Mode"]
                                            defaultRetentionMap["years"] = defaultRetentionChildRaw["Years"]
                
                                                                                                                                                                                            defaultRetentionMaps = append(defaultRetentionMaps, defaultRetentionMap)
            }
                        }
                                                ruleMap["default_retention"] = defaultRetentionMaps
            ruleMaps = append(ruleMaps, ruleMap)
}
if err := d.Set("rule", ruleMaps); err != nil {
return err
}
                            
							
        
        		d.Set("bucket_name", d.Id())
	
		        	return nil
}


    func resourceAliCloudOssBucketObjectWormConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
                    var query map[string]*string
        var body map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                                action := fmt.Sprintf("/?objectWorm")
    request = make(map[string]interface{})
                        query = make(map[string]*string)
                body = make(map[string]interface{})
                                hostMap := make(map[string]*string)
                                        hostMap["bucket"] = StringPointer(d.Id())
            
                                                                                                                            if  d.HasChange("object_worm_enabled") {
    update = true
    }
            objectWormConfiguration := make(map[string]interface{})
    
if v := d.Get("object_worm_enabled");  !IsNil(v)  || d.HasChange("object_worm_enabled")  {
                                                if v, ok := d.GetOk("object_worm_enabled"); ok {
                objectWormConfiguration["ObjectWormEnabled"] = v            }
                                            rule := make(map[string]interface{})

            if len(rule) > 0 {
        objectWormConfiguration["Rule"] = rule
        }
            
                request["ObjectWormConfiguration"] = objectWormConfiguration
    }

                                                    body = request
                    if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                    response, err = client.Do("Oss", xmlParam("PUT", "2019-05-17", "PutBucketObjectWormConfiguration", action), query, body, nil, hostMap, false)
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
                
                    return resourceAliCloudOssBucketObjectWormConfigurationRead(d, meta)
}


func resourceAliCloudOssBucketObjectWormConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
    log.Printf("[WARN] Cannot destroy resource AliCloud Resource Bucket Object Worm Configuration. Terraform will remove this resource from the state file, however resources may remain.")
    return nil
}




