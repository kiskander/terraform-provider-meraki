package meraki

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("api_key").(string)
	var diags diag.Diagnostics

	if apiKey != "" {
		return &Client{APIKey: apiKey}, diags

	}
	return nil, diag.FromErr(errors.New("Missing API Key or Org ID"))
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("MERAKI_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"meraki_organizations": dataSourceOrganizations(),
			"meraki_devices":       dataSourceDevice(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
