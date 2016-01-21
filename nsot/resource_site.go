package nsot

import (
	"fmt"
	"log"
	"strings"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sarguru/go-nsot-api"
)

func resourceNsotSite() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: SiteCreate,
		Read:   SiteRead,
		Update: SiteUpdate,
		Delete: SiteDelete,
	}
}



func SiteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsot.Client)

	// Create the new Site
	newSite := &nsot.SiteOpts{
		Name: d.Get("name").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		newSite.Desc = desc.(string)
	}

	log.Printf("[DEBUG] NSOT site create configuration: %#v", newSite)

	site, err := client.CreateSite(newSite)

	if err != nil {
		return fmt.Errorf("Failed to create NSOT site: %s", err)
	}

	d.SetId(strconv.Itoa(site.Id))
	log.Printf("[INFO] NSOT Site ID: %s", d.Id())

	return SiteRead(d, meta)
}

func SiteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsot.Client)

	site, err := client.RetrieveSitebyName(d.Get("name").(string))

	log.Printf("[DEBUG] NSOT site create retrieved: %#v", site)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Couldn't find NSOT site by name (%s): %s",
			d.Get("domain").(string), err)
	}

	d.Set("name", site.Name)
	d.SetId(strconv.Itoa(site.Id))
	if site.Desc != "" {
		d.Set("description", site.Desc)
	}
	return nil
}

func SiteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsot.Client)

	updateSite := &nsot.SiteOpts{
		Name: d.Get("name").(string),
	}

	if desc, ok := d.GetOk("description"); ok {
		updateSite.Desc = desc.(string)
	}

	log.Printf("[DEBUG] NSOT site update configuration: %#v", updateSite)

	id,err := strconv.Atoi(d.Id())

	if err != nil {
		return fmt.Errorf("Error fetching site ID: %s", err)
	}

	_, err = client.UpdateSitebyID(id, updateSite)
	if err != nil {
		return fmt.Errorf("Failed to update NSOT site: %s", err)
	}

	return SiteRead(d, meta)
}

func SiteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsot.Client)

	log.Printf("[INFO] Deleting NSOT Site: %s, %s", d.Get("name").(string), d.Id())

	id,err := strconv.Atoi(d.Id())

	if err != nil {
		return fmt.Errorf("Error fetching site ID: %s", err)
	}

	err = client.DestroySitebyID(id)

	if err != nil {
		return fmt.Errorf("Error deleting NSOT site: %s", err)
	}

	return nil
}
