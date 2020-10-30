package meraki

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type Client struct {
	APIKey         string
	OrganizationID string
}

type GetOrganizationDevicesInput struct {
	OrganizationID string
}

type GetOrganizationDevicesOutput struct {
	Name                   string   `json:"name"`
	Serial                 string   `json:"serial"`
	Mac                    string   `json:"mac"`
	NetworkID              string   `json:"networkId"`
	Model                  string   `json:"model"`
	Address                string   `json:"address"`
	Lat                    float64  `json:"lat"`
	Lng                    float64  `json:"lng"`
	Notes                  string   `json:"notes"`
	Tags                   []string `json:"tags"`
	LanIP                  string   `json:"lanIp"`
	ConfigurationUpdatedAt string   `json:"configurationUpdatedAt"`
	Firmware               string   `json:"firmware"`
	URL                    string   `json:"url"`
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

func (v *Client) GetOrganizationDevices(ctx context.Context, input *GetOrganizationDevicesInput) ([]GetOrganizationDevicesOutput, diag.Diagnostics) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.meraki.com/api/v1/organizations/%s/devices", input.OrganizationID), nil)
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
	items := make([]GetOrganizationDevicesOutput, 0)
	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return items, nil
}
