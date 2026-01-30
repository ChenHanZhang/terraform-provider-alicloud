// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionIncidentInvestigation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionIncidentInvestigationCreate,
		Read:   resourceAliCloudThreatDetectionIncidentInvestigationRead,
		Delete: resourceAliCloudThreatDetectionIncidentInvestigationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"incident_investigation_display_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"incident_investigation_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"incident_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudThreatDetectionIncidentInvestigationCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudThreatDetectionIncidentInvestigationRead(d, meta)
}

func resourceAliCloudThreatDetectionIncidentInvestigationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionIncidentInvestigation(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_incident_investigation DescribeThreatDetectionIncidentInvestigation Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("incident_investigation_display_id", objectRaw["IncidentInvestigationDisplayId"])
	d.Set("incident_uuid", objectRaw["IncidentUuid"])
	d.Set("incident_investigation_id", objectRaw["IncidentInvestigationId"])

	return nil
}

func resourceAliCloudThreatDetectionIncidentInvestigationDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Incident Investigation. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
