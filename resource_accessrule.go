package main

import (
	"github.com/bluecough/go-ftd"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessRuleCreate,
		Read:   resourceAccessRuleRead,
		Update: resourceAccessRuleUpdate,
		Delete: resourceAccessRuleDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ruleid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"sourcezones": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"destinationzones": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sourcenetworks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"destinationnetworks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sourceports": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"destinationports": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ruleaction": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"eventlogaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlantags": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"intrusionpolicy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"filepolicy": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"logfiles": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"syslogserver": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource {
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parent": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	// Define all the Structures and associated variables
	pAR := new(goftd.AccessRule)

	// Name
	pAR.Name = d.Get("name").(string)

	// Source Zones
	psourcezones := d.Get("sourcezones").(*schema.Set)
	var localSourceZoneObject = []*goftd.ReferenceObject{}
	for _, vRaw := range psourcezones.List() {
		val := vRaw.(map[string]interface{})

		localSourceZoneObject = append(localSourceZoneObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.SourceZones = localSourceZoneObject

	// Destination Zones
	pdestinationzones := d.Get("destinationzones").(*schema.Set)
	var localDestinationZoneObject = []*goftd.ReferenceObject{}
	for _, vRaw := range pdestinationzones.List() {
		val := vRaw.(map[string]interface{})

		localDestinationZoneObject = append(localDestinationZoneObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.DestinationZones = localDestinationZoneObject

	// Source Networks
	psourcenetworks := d.Get("sourcenetworks").(*schema.Set)
	var localSourceNetworkObject = []*goftd.ReferenceObject{}
	for _, vRaw := range psourcenetworks.List() {
		val := vRaw.(map[string]interface{})

		localSourceNetworkObject = append(localSourceNetworkObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.SourceNetworks = localSourceNetworkObject

	// Destination Networks
	pdestinationnetworks := d.Get("destinationnetworks").(*schema.Set)
	var localDestinationNetworkObject = []*goftd.ReferenceObject{}
	for _, vRaw := range pdestinationnetworks.List() {
		val := vRaw.(map[string]interface{})

		localDestinationNetworkObject = append(localDestinationNetworkObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.DestinationNetworks = localDestinationNetworkObject

	// Source Ports
	psourceports := d.Get("sourceports").(*schema.Set)
	var localSourcePortObject = []*goftd.ReferenceObject{}
	for _, vRaw := range psourceports.List() {
		val := vRaw.(map[string]interface{})

		localSourcePortObject = append(localSourcePortObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.SourcePorts = localSourcePortObject

	// Destination Ports
	pdestinationports := d.Get("destinationports").(*schema.Set)
	var localDestinationPortObject = []*goftd.ReferenceObject{}
	for _, vRaw := range pdestinationports.List() {
		val := vRaw.(map[string]interface{})

		localDestinationPortObject = append(localDestinationPortObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.DestinationPorts = localDestinationPortObject

	// Rule Action
	pAR.RuleAction = d.Get("ruleaction").(string)

	// Event Log Action
	pAR.EventLogAction = d.Get("eventlogaction").(string)

	// Vlan Tags
	pvlantags := d.Get("vlantags").(*schema.Set)
	var localVlanTagObject = []*goftd.ReferenceObject{}
	for _, vRaw := range pvlantags.List() {
		val := vRaw.(map[string]interface{})

		localVlanTagObject = append(localVlanTagObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.VLANTags = localVlanTagObject

	// Users
	pusers := d.Get("users").(*schema.Set)
	var localUsersObject = []*goftd.ReferenceObject{}
	for _, vRaw := range pusers.List() {
		val := vRaw.(map[string]interface{})

		localUsersObject = append(localUsersObject, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}
	pAR.Users = localUsersObject


	// Intrusion Policy
	tf := d.Get("intrusionpolicy").(map[string]interface{})


	if tf["name"] != nil {
		var localIntrusionPolicyObject = new(goftd.ReferenceObject)

		
		localIntrusionPolicyObject.Name = tf["name"].(string)
		localIntrusionPolicyObject.Type = tf["type"].(string)

		pAR.IntrusionPolicy = localIntrusionPolicyObject
	}
	// File Policy
	//var localFilePolicyObject = new(goftd.ReferenceObject)
   	tg := d.Get("filepolicy").(map[string]interface{})


    if tg["name"] != nil {
 		var localFilePolicyObject = new(goftd.ReferenceObject)


		localFilePolicyObject.Name = tg["name"].(string)
		localFilePolicyObject.Type = tg["type"].(string)

		pAR.FilePolicy = localFilePolicyObject
	}


	// Log Files
	pAR.LogFiles = d.Get("logfiles").(bool)

	// Syslog Server
	th := d.Get("syslogserver").(map[string]interface{})


	if th["name"] != nil {
		var localSyslogServerObject = new(goftd.ReferenceObject)


		localSyslogServerObject.Name = th["name"].(string)
		localSyslogServerObject.Type = th["type"].(string)

		pAR.SyslogServer = localSyslogServerObject
	}

	// Parent

	// Call CreateAccessRule
	err := cf.CreateAccessRule(pAR, "default" )
	if err != nil{
		log.Println("==== > Error %s",err)
		return err
	}

	d.SetId(pAR.ID)
	return resourceServerRead(d, m)
}

func resourceAccessRuleRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceServerRead(d, m)
}

func resourceAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	pAP, err := cf.GetAccessPolicies( 10)
	if err != nil{
		log.Println("==== > Error %s",err)
		return err
	}

	n := new(goftd.AccessRule)
	n.ID = d.Id()
	n.Parent = pAP[0].ID
	cf.DeleteAccessRule(n)

	return nil
}
