package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/bluecough/go-ftd"
	"log"
	"strings"
)

func resourcePortObjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourcePortObjectGroupCreate,
		Read:   resourcePortObjectGroupRead,
		Update: resourcePortObjectGroupUpdate,
		Delete: resourcePortObjectGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"objects": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
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
							Required: true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourcePortObjectGroupCreate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)

	total := d.Get("objects.#")
	log.Println("GS DEBUG =====PortkObjectGroupCreate==",total)

	po := new(goftd.PortObjectGroup)
	po.Name = d.Get("name").(string)

	entries := d.Get("objects").(*schema.Set)
	var batchEntries = []*goftd.ReferenceObject{}
	for _, vRaw := range entries.List() {
		val := vRaw.(map[string]interface{})

		batchEntries = append(batchEntries, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["name"].(string),
			Type:    val["type"].(string),
		})
	}

	po.Objects = batchEntries
	err := cf.CreatePortObjectGroup(po, goftd.DuplicateActionReplace)

	//log.Println("GS DEBUG ====== NetworkObjectGroupCreate-E==== ",n.ID)
	d.SetId(po.ID + " " + po.Name)
	if err != nil{
		log.Println("==== > Error %s",err)
		return err
	}
	return resourceServerRead(d, m)
}

func resourcePortObjectGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePortObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceServerRead(d, m)
}

func resourcePortObjectGroupDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	n := new(goftd.PortObjectGroup)
	v := strings.Split(d.Id(), " ")

	n.ID = v[0]
	err := cf.DeletePortObjectGroup(n)


	if err != nil {
		log.Println("Error: %s\n", err)
		return err
	}


	return nil
}