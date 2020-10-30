package meraki

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type Client struct {
	APIKey string
}

type GetDevicesInput struct {
	OrgID string
}

func (v *Client) GetOrganizations(ctx context.Context) ([]map[string]interface{}, diag.Diagnostics) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.meraki.com/api/v1/organizations", nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	req.Header.Set("X-Cisco-Meraki-API-Key", v.APIKey)
	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	defer r.Body.Close()
	items := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return items, nil
}

func (v *Client) GetDevices(ctx context.Context, input *GetDevicesInput) ([]map[string]interface{}, diag.Diagnostics) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.meraki.com/api/v1/%s/devices", input.OrgID), nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	req.Header.Set("X-Cisco-Meraki-API-Key", v.APIKey)
	c := http.DefaultClient
	r, err := c.Do(req)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	defer r.Body.Close()
	items := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return items, nil
}
