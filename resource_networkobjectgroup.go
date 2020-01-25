package main

import (
	"fmt"
	"github.com/bluecough/go-ftd"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceNetworkObjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkObjectGroupCreate,
		Read:   resourceNetworkObjectGroupRead,
		Update: resourceNetworkObjectGroupUpdate,
		Delete: resourceNetworkObjectGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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
						"netname": &schema.Schema{
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNetworkObjectGroupCreate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
    total := d.Get("objects.#")
    log.Println("GS DEBUG ===",total)

    n := new(goftd.NetworkObjectGroup)
    n.Name = d.Get("name").(string)

    entries := d.Get("objects").(*schema.Set)
    var batchEntries = []*goftd.ReferenceObject{}
    for _, vRaw := range entries.List() {
    	val := vRaw.(map[string]interface{})

    	batchEntries = append(batchEntries, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["netname"].(string),
			Type:    val["type"].(string),
		})
	}

	n.Objects = batchEntries

	err := cf.CreateNetworkObjectGroup(n, goftd.DuplicateActionReplace)
   // log.Println(&n.Objects[0].Name)
	d.SetId(n.ID)

	if err != nil{
		fmt.Errorf(err.Error())
	}

	return resourceNetworkObjectGroupRead(d, m)
}

func resourceNetworkObjectGroupRead(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	cf.GetNetworkObjectGroupBy(d.Get("name").(string))
    //log.Println("GS DEBUG === ", reflect.TypeOf(d.Get("name")))
	return nil
}

func resourceNetworkObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	total := d.Get("objects.#")
	log.Println("GS DEBUG ======NetworkObjectUpdate-B====== \n",total)

	n := new(goftd.NetworkObjectGroup)
	n.Name = d.Get("name").(string)

	existing, err := cf.GetNetworkObjectGroupBy(d.Get("name").(string))

	entries := d.Get("objects").(*schema.Set)
	var batchEntries = []*goftd.ReferenceObject{}
	for _, vRaw := range entries.List() {
		val := vRaw.(map[string]interface{})

		batchEntries = append(batchEntries, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["netname"].(string),
			Type:    val["type"].(string),
		})
	}

	n.Objects = batchEntries
	n.ID = d.Id()
	err = cf.UpdateNetworkObjectGroup(n)
	log.Println(n.Objects[0].Name)
	log.Println(n.Name)
	log.Println(n.ID)
	log.Println(n.Version)
	//log.Println(d.Get("version").(string))
	log.Println(len(existing))

	if err != nil{
		log.Println("GS DEBUG =====NetworkObjectUpdate-E====== \n", err)
	}

	return resourceNetworkObjectGroupRead(d, m)
}

func resourceNetworkObjectGroupDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	n := new(goftd.NetworkObjectGroup)
	n.ID = d.Id()
	err := cf.DeleteNetworkObjectGroup(n)
	if err != nil {
		log.Println("Error: %s\n", err)
		return err
	}

	return nil
}
