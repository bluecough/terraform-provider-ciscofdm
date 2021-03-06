package main

import (
	"fmt"
	"github.com/bluecough/go-ftd"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
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
    log.Println("GS DEBUG =====NetworkObjectGroupCreate==",total)

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
    //log.Println("GS DEBUG ====== NetworkObjectGroupCreate-E==== ",n.ID)
	d.SetId(n.ID + " " + n.Name)
	if err != nil{
		fmt.Errorf(err.Error())
		return err
	}

	return resourceNetworkObjectGroupRead(d, m)
}

func resourceNetworkObjectGroupRead(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	filter := "name: " + d.Get("name").(string)
	cf.GetNetworkObjectGroupBy(filter)
	return nil
}

func resourceNetworkObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	var err error
	v := []*goftd.NetworkObjectGroup{}
	idsplit := strings.Split(d.Id(), " ")
	n := "name: " + idsplit[1]
	v,err = cf.GetNetworkObjectGroupBy(n)

	if err != nil{
		log.Println("GS DEBUG ====> call for GetNetworkObjectGroupBy ", err)
		return err
	}

	/*
	This section is around displaying the return values of the GetNetworkObjectGroupBy function
	*/

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
	// Create localro as slice pointer and add more slices based upon the number of Objects we had
	// Then assign values that we had from the read
	localnog := new(goftd.NetworkObjectGroup)
	localro := []*goftd.ReferenceObject{}

	for i := 0; i < len(batchEntries); i++ {
		localro = append(localro,new(goftd.ReferenceObject))
		localro[i].Name = batchEntries[i].Name
		localro[i].ID = batchEntries[i].ID
		localro[i].Version = batchEntries[i].Version
		localro[i].Type = batchEntries[i].Type
	}

	localnog.Objects = localro
	localnog.Name = d.Get("name").(string)
	localnog.ID = idsplit[0]
	localnog.Type = v[0].Type
	localnog.Version = v[0].Version

	cf.UpdateNetworkObjectGroup(localnog)

	if err != nil{
		log.Println("GS DEBUG =====NetworkObjectGRPtUpdate-E====== \n", err)
	}
	idsplit[1] = d.Get("name").(string)
	d.SetId(idsplit[0] + " " + idsplit[1])
	return resourceNetworkObjectGroupRead(d, m)
}

func resourceNetworkObjectGroupDelete(d *schema.ResourceData, m interface{}) error {
	cf := m.(*goftd.FTD)
	n := new(goftd.NetworkObjectGroup)
	v := strings.Split(d.Id(), " ")

	n.ID = v[0]
	err := cf.DeleteNetworkObjectGroup(n)


	if err != nil {
		log.Println("Error: %s\n", err)
		return err
	}

	return nil
}
