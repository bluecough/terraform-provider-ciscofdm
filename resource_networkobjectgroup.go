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
	n.Objects = batchEntries

    for _, vRaw := range entries.List() {
    	val := vRaw.(map[string]interface{})

    	batchEntries = append(batchEntries, &goftd.ReferenceObject{
			ID:      val["id"].(string),
			Version: val["version"].(string),
			Name:    val["netname"].(string),
			Type:    val["type"].(string),
		})
	}
	err := cf.CreateNetworkObjectGroup(n, goftd.DuplicateActionReplace)
   // log.Println(&n.Objects[0].Name)
	d.SetId(n.Name)
	if err != nil{
		fmt.Errorf(err.Error())
	}

	return resourceNetworkObjectGroupRead(d, m)
}

func resourceNetworkObjectGroupRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceNetworkObjectGroupUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceNetworkObjectGroupRead(d, m)
}

func resourceNetworkObjectGroupDelete(d *schema.ResourceData, m interface{}) error {


	return nil
}
