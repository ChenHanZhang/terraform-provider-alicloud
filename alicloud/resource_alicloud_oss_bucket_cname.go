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

func resourceAliCloudOssBucketCname() *schema.Resource {
  return &schema.Resource{
            Create: resourceAliCloudOssBucketCnameCreate,
                Read: resourceAliCloudOssBucketCnameRead,
                Update: resourceAliCloudOssBucketCnameUpdate,
            Delete: resourceAliCloudOssBucketCnameDelete,
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
                                    "certificate": {
            Type:schema.TypeList,
                                                                                            Optional: true,
                                                                                                                                                                                                            MaxItems: 1,
                Elem: &schema.Resource {
                              Schema: map[string]*schema.Schema {
                        "status": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "creation_date": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "valid_end_date": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "type": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "fingerprint": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "valid_start_date": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
                                    "private_key": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                Sensitive: true,
                                                                                                                                    },
                                    "cert_id": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                    Computed: true,
                                                                                                                                                                                        },
                                    "certificate": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                Sensitive: true,
                                                                                                                                    },
            },
                },
                        },
                                    "delete_certificate": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "domain": {
            Type:schema.TypeString,
                                                            Required: true,
                                                                                                        ForceNew: true,
                                                                                                                                                            },
                                    "force": {
            Type:schema.TypeBool,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                        "previous_cert_id": {
            Type:schema.TypeString,
                                                                                            Optional: true,
                                                                                                                                                                                                        },
                                    "status": {
            Type:schema.TypeString,
                                                                                                            Computed: true,
                                                                                                                                                                                        },
            },
  }
}

func resourceAliCloudOssBucketCnameCreate(d *schema.ResourceData, meta interface{}) error {

        
        
                                                                                                    client := meta.(*connectivity.AliyunClient)
             
        action := fmt.Sprintf("/?cname&comp=add")
            var request map[string]interface{}
        var response map[string]interface{}
                                                    query := make(map[string]*string)
            body := make(map[string]interface{})
                                            hostMap := make(map[string]*string)
                                                var err error
        request = make(map[string]interface{})
                hostMap["bucket"] = StringPointer(d.Get("bucket").(string))
            
    
                                                                                                            bucketCnameConfiguration := make(map[string]interface{})
    
if v := d.Get("domain");  !IsNil(v)   {
                    cname := make(map[string]interface{})
                                        cname["Domain"] = v                                            certificateConfiguration := make(map[string]interface{})
                                    certId1, _ := jsonpath.Get("$[0].cert_id", d.Get("certificate"))
        if certId1 != nil                && certId1 != "" {
                certificateConfiguration["CertId"] = certId1        }
                                                    certificate1, _ := jsonpath.Get("$[0].certificate", d.Get("certificate"))
        if certificate1 != nil                && certificate1 != "" {
                certificateConfiguration["Certificate"] = certificate1        }
                                                    privateKey1, _ := jsonpath.Get("$[0].private_key", d.Get("certificate"))
        if privateKey1 != nil                && privateKey1 != "" {
                certificateConfiguration["PrivateKey"] = privateKey1        }
                
            if len(certificateConfiguration) > 0 {
        cname["CertificateConfiguration"] = certificateConfiguration
        }
                
            if len(cname) > 0 {
        bucketCnameConfiguration["Cname"] = cname
        }
            
                request["BucketCnameConfiguration"] = bucketCnameConfiguration
    }

                                    jsonString := convertObjectToJsonString(hostMap)
                    jsonString, _ = sjson.Set(jsonString, "BucketCnameConfiguration.Cname.Domain", d.Get("domain"))
                _ = json.Unmarshal([]byte(jsonString), &hostMap)

                            body = request
        wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
            response, err = client.Do("Oss", xmlParam("POST", "2019-05-17", "PutCname", action), query, body, nil, hostMap, false)
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
        return WrapErrorf(err, DefaultErrorMsg, "alicloud_oss_bucket_cname", action, AlibabaCloudSdkGoERROR)
    }
        
    
                                                        BucketCnameConfigurationCnameDomainVar, _ := jsonpath.Get("BucketCnameConfiguration.Cname.Domain", request)
                            d.SetId(fmt.Sprintf("%v:%v",*hostMap["bucket"],BucketCnameConfigurationCnameDomainVar));
    

            
        
        
        return resourceAliCloudOssBucketCnameRead(d, meta)
    }


    		func resourceAliCloudOssBucketCnameRead(d *schema.ResourceData, meta interface{}) error {
	    		client := meta.(*connectivity.AliyunClient)
		ossServiceV2 := OssServiceV2{client}

												objectRaw, err := ossServiceV2.DescribeOssBucketCname(d.Id())
				if err != nil  {
					if !d.IsNewResource() && NotFoundError(err) {
				log.Printf("[DEBUG] Resource alicloud_oss_bucket_cname DescribeOssBucketCname Failed!!! %s", err)
				d.SetId("")
				return nil
			}
					return WrapError(err)
		}

				

                            
                                        d.Set("bucket", objectRaw["Bucket"])
                
            cnameChildRawObj, _ := jsonpath.Get("$.Cname[*]", objectRaw)
cnameChildRaw := make([]interface{}, 0)
if cnameChildRawObj != nil {
cnameChildRaw = convertToInterfaceArray(cnameChildRawObj)}

                                        d.Set("status", cnameChildRaw["Status"])
                                            d.Set("domain", cnameChildRaw["Domain"])
                
                                                                certificateMaps := make([]map[string]interface{}, 0)
        certificateMap := make(map[string]interface{})
                certificateRawObj, _ := jsonpath.Get("$.Cname[*].Certificate", objectRaw)
certificateRaw := make(map[string]interface{})
if certificateRawObj != nil {
certificateRaw = certificateRawObj.(map[string]interface{})
}
if len(certificateRaw) > 0 {
                                        certificateMap["cert_id"] = certificateRaw["CertId"]
                                            certificateMap["creation_date"] = certificateRaw["CreationDate"]
                                            certificateMap["fingerprint"] = certificateRaw["Fingerprint"]
                                            certificateMap["status"] = certificateRaw["Status"]
                                            certificateMap["type"] = certificateRaw["Type"]
                                            certificateMap["valid_end_date"] = certificateRaw["ValidEndDate"]
                                            certificateMap["valid_start_date"] = certificateRaw["ValidStartDate"]
                
                                                                                                                                                                                                                                                                                        certificateMaps = append(certificateMaps, certificateMap)
}
if err := d.Set("certificate", certificateMaps); err != nil {
return err
}
                                                                                                                                            
							
        
        
		        	return nil
}


    func resourceAliCloudOssBucketCnameUpdate(d *schema.ResourceData, meta interface{}) error {
                client := meta.(*connectivity.AliyunClient)
        var request map[string]interface{}
    var response map[string]interface{}
            var header map[string]*string
        var query map[string]*string
        var body map[string]interface{}
                        update := false
        
        
        
                                                                var err error
                                                                                                parts := strings.Split(d.Id(), ":")
            action := fmt.Sprintf("/?cname&comp=add")
    request = make(map[string]interface{})
                        query = make(map[string]*string)
                body = make(map[string]interface{})
                                hostMap := make(map[string]*string)
                                        hostMap["bucket"] = StringPointer(parts[0])
            
                                                                                                                                        bucketCnameConfiguration := make(map[string]interface{})
    
if v := d.Get("domain");  !IsNil(v)   {
                    cname := make(map[string]interface{})
                                                    if v, ok := d.GetOk("domain"); ok {
                cname["Domain"] = v            }
                                                    certificateConfiguration := make(map[string]interface{})
                                                    if v, ok := d.GetOk("previous_cert_id"); ok {
                certificateConfiguration["PreviousCertId"] = v            }
                                                                                if v, ok := d.GetOkExists("force"); ok {
                certificateConfiguration["Force"] = v            }
                                                                                if v, ok := d.GetOkExists("delete_certificate"); ok {
                certificateConfiguration["DeleteCertificate"] = v            }
                                                                certId1, _ := jsonpath.Get("$[0].cert_id", d.Get("certificate"))
        if certId1 != nil&&  certId1 != ""  {
                            certificateConfiguration["CertId"] = certId1                }
                                                    certificate1, _ := jsonpath.Get("$[0].certificate", d.Get("certificate"))
        if certificate1 != nil&&  certificate1 != ""  {
                            certificateConfiguration["Certificate"] = certificate1                }
                                                    privateKey1, _ := jsonpath.Get("$[0].private_key", d.Get("certificate"))
        if privateKey1 != nil&&  privateKey1 != ""  {
                            certificateConfiguration["PrivateKey"] = privateKey1                }
                
            if len(certificateConfiguration) > 0 {
        cname["CertificateConfiguration"] = certificateConfiguration
        }
                
            if len(cname) > 0 {
        bucketCnameConfiguration["Cname"] = cname
        }
            
                request["BucketCnameConfiguration"] = bucketCnameConfiguration
    }

                                    jsonString := convertObjectToJsonString(hostMap)
                    jsonString, _ = sjson.Set(jsonString, "BucketCnameConfiguration.Cname.Domain", parts[1])
                _ = json.Unmarshal([]byte(jsonString), &hostMap)

                            body = request
                    if update  {
            wait := incrementalWait(3*time.Second, 5*time.Second)
        err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
                    response, err = client.Do("Oss", xmlParam("POST", "2019-05-17", "PutCname", action), query, body, nil, hostMap, false)
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
                
                    return resourceAliCloudOssBucketCnameRead(d, meta)
}


func resourceAliCloudOssBucketCnameDelete(d *schema.ResourceData, meta interface{}) error {
                        
                                                                                        client := meta.(*connectivity.AliyunClient)
                    parts := strings.Split(d.Id(), ":")
            action := fmt.Sprintf("/?cname&comp=delete")
            var request map[string]interface{}
        var response map[string]interface{}
                                                    query := make(map[string]*string)
                            body := make(map[string]interface{})
                                                    hostMap := make(map[string]*string)
                var err error
        request = make(map[string]interface{})
                hostMap["bucket"] = StringPointer(parts[0])
            
    

                jsonString := convertObjectToJsonString(hostMap)
                    jsonString, _ = sjson.Set(jsonString, "BucketCnameConfiguration.Cname.Domain", parts[1])
                _ = json.Unmarshal([]byte(jsonString), &hostMap)

                    body = request
        wait := incrementalWait(3*time.Second, 5*time.Second)
    err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
            response, err = client.Do("Oss", xmlParam("POST", "2019-05-17", "DeleteCname", action), query, body, nil, hostMap, false)
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
        if IsExpectedErrors(err, []string{"NoSuchBucket","NoSuchCname"}) ||  NotFoundError(err) {
            return nil
        }
        return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
    }

        
                        return nil
}




