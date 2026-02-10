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

func resourceAliCloudOssBucketLogging() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudOssBucketLoggingCreate,
                Read: resourceAliCloudOssBucketLoggingRead,
                Update: resourceAliCloudOssBucketLoggingUpdate,
            Delete: resourceAliCloudOssBucketLoggingDelete,
    Importer: &schema.ResourceImporter{
      State: schema.ImportStatePassthrough,
    },
    Timeouts: &schema.ResourceTimeout{
              Create: schema.DefaultTimeout(5 * time.Minute),
                  Update: schema.DefaultTimeout(5 * time.Minute),
                  Delete: schema.DefaultTimeout(5 * time.Minute),
        },
            Schema: map[string]*schema.Schema {
                        "bucket": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                                    "logging_role": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "target_bucket": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                                                                                                                                                        },
                                    "target_prefix": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
            },
  }
}

func resourceAliCloudOssBucketLoggingCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := fmt.Sprintf("/?logging")
            var request map[string]interface{}
        var response map[string]interface{}
                                                    query := make(map[string]*string)
            body := make(map[string]interface{})
                                            hostMap := make(map[string]*string)
                                                var err error
        request = make(map[string]interface{})
                hostMap["bucket"] = StringPointer(d.Get("bucket").(string))
            
    
                                                                                                                    bucketLoggingStatus := make(map[string]interface{})

                                                                            if v := d.Get("target_prefix");  !IsNil(v)   {
                                                                                    loggingEnabled := make(map[string]interface{})
                                        loggingEnabled["TargetPrefix"] = v                                                            loggingEnabled["LoggingRole"] = v                                                            loggingEnabled["TargetBucket"] = v                    
            if len(loggingEnabled) > 0 {
        bucketLoggingStatus["LoggingEnabled"] = loggingEnabled
        }
                                                        }
                            
            request["BucketLoggingStatus"] = bucketLoggingStatus
    
                                                                    body = request
        wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
            response, err = client.Do("Oss", xmlParam("PUT", "2019-05-17", "PutBucketLogging", action), query, body, nil, hostMap, false)
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
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_oss_bucket_logging", action, AlibabaCloudSdkGoERROR)
    }
        
                        d.SetId(fmt.Sprint(*hostMap["bucket"]))        

                ossServiceV2 := OssServiceV2{client}
        stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("target_bucket"))}, d.Timeout(schema.TimeoutCreate),  5 * time.Second ,ossServiceV2.OssBucketLoggingStateRefreshFunc(d.Id(), "TargetBucket", []string {}))
        if _, err := stateConf.WaitForState(); err != nil {
            return WrapErrorf(err, IdMsg, d.Id())
        }
        
        
        
        return resourceAliCloudOssBucketLoggingRead(d, meta)
    }


    		func resourceAliCloudOssBucketLoggingRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		ossServiceV2 := OssServiceV2{client}

												objectRaw, err := ossServiceV2.DescribeOssBucketLogging(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_oss_bucket_logging DescribeOssBucketLogging Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("logging_role", objectRaw["LoggingRole"])
                                            d.Set("target_bucket", objectRaw["TargetBucket"])
                                            d.Set("target_prefix", objectRaw["TargetPrefix"])
                
                                                                                                                                                            
							
        
        		d.Set("bucket", d.Id())
	
		        	return nil
}


    func resourceAliCloudOssBucketLoggingUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
                    var query map[string]*string
        var body map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                                action := fmt.Sprintf("/?logging")
    request = make(map[string]interface{})
                        query = make(map[string]*string)
                body = make(map[string]interface{})
                                hostMap := make(map[string]*string)
                                        hostMap["bucket"] = StringPointer(d.Id())
            
                                                                                                                                                bucketLoggingStatus := make(map[string]interface{})

                                        if  d.HasChange("target_prefix") {
            update = true
                            }
                                                if v := d.Get("target_prefix");  !IsNil(v)  || d.HasChange("target_prefix")  {
                                                                                    loggingEnabled := make(map[string]interface{})
                                                    if v, ok := d.GetOk("target_prefix"); ok {
                loggingEnabled["TargetPrefix"] = v            }
                                                                                if v, ok := d.GetOk("logging_role"); ok {
                loggingEnabled["LoggingRole"] = v            }
                                                                                if v, ok := d.GetOk("target_bucket"); ok {
                loggingEnabled["TargetBucket"] = v            }
                            
            if len(loggingEnabled) > 0 {
        bucketLoggingStatus["LoggingEnabled"] = loggingEnabled
        }
                                                        }
                            
            request["BucketLoggingStatus"] = bucketLoggingStatus
    
                                                                    body = request
                    if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                    response, err = client.Do("Oss", xmlParam("PUT", "2019-05-17", "PutBucketLogging", action), query, body, nil, hostMap, false)
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
                    ossServiceV2 := OssServiceV2{client}
            stateConf := BuildStateConf([]string{}, []string{fmt.Sprint(d.Get("target_bucket"))}, d.Timeout(schema.TimeoutUpdate),  5 * time.Second ,ossServiceV2.OssBucketLoggingStateRefreshFunc(d.Id(), "TargetBucket", []string {}))
            if _, err := stateConf.WaitForState(); err != nil {
                return WrapErrorf(err, IdMsg, d.Id())
            }
                                }
                
                    return resourceAliCloudOssBucketLoggingRead(d, meta)
}


func resourceAliCloudOssBucketLoggingDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                                        client := meta.(*connectivity.AliyunClient)
                    action := fmt.Sprintf("/?logging")
            var request map[string]interface{}
        var response map[string]interface{}
                                                    query := make(map[string]*string)
                                                    hostMap := make(map[string]*string)
                var err error
        request = make(map[string]interface{})
                hostMap["bucket"] = StringPointer(d.Id())
            
    

                    wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
            response, err = client.Do("Oss", xmlParam("DELETE", "2019-05-17", "DeleteBucketLogging", action), query, nil, nil, hostMap, false)
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




