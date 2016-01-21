package nsot

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSOT_EMAIL", nil),
				Description: "Your nsot email.",
			},

			"secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSOT_SECRET", nil),
				Description: "The secret key for API operations.",
			},

			"url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NSOT_URL", nil),
				Description: "The URL to your nsot instance.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"nsot_site": resourceNsotSite(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Email:  d.Get("email").(string),
		Secret: d.Get("secret").(string),
		Url:    d.Get("url").(string),
	}

	return config.Client()
}
