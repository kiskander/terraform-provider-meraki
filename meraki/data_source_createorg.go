package meraki

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCreateOrgs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationsCreate,
		Schema: map[string]*schema.Schema{
			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: false,
				Required: true,
			},
			"organizations": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOrganizationsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	OrgName := d.Get("org_name").(string)
	if OrgName == "" {
		return diag.FromErr(errors.New("missing org_name"))
	}

	client := m.(*Client)
	organizations, diags := client.CreateOrganization(ctx, &CreateOrganizationInput{OrgName: OrgName})
	if diags != nil {
		return diags
	}
	items := make([]map[string]interface{}, 0)
	for _, organizations := range organizations {
		items = append(items, map[string]interface{}{
			"id":   organizations.Id,
			"name": organizations.Name,
			"url":  organizations.Url,
		})
	}

	if err := d.Set("organizations", items); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
