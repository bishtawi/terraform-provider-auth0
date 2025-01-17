package auth0

import (
	"github.com/bishtawi/auth0"
	"github.com/bishtawi/auth0/management"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func newCustomDomain() *schema.Resource {
	return &schema.Resource{

		Create: createCustomDomain,
		Read:   readCustomDomain,
		Update: updateCustomDomain,
		Delete: deleteCustomDomain,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"auth0_managed_certs",
					"self_managed_certs",
				}, true),
			},
			"primary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"verification_method": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"txt"}, true),
			},
			"verification": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"methods": {
							Type:     schema.TypeList,
							Elem:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createCustomDomain(d *schema.ResourceData, m interface{}) error {
	c := buildCustomDomain(d)
	api := m.(*management.Management)
	if err := api.CustomDomain.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readCustomDomain(d, m)
}

func readCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.CustomDomain.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	d.Set("domain", c.Domain)
	d.Set("type", c.Type)
	d.Set("primary", c.Primary)
	d.Set("status", c.Status)

	if c.Verification != nil {
		d.Set("verification", []map[string]interface{}{
			{"methods": c.Verification.Methods},
		})
	}

	return nil
}

func updateCustomDomain(d *schema.ResourceData, m interface{}) error {
	c := buildCustomDomain(d)
	c.Verification = nil
	api := m.(*management.Management)
	err := api.CustomDomain.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readCustomDomain(d, m)
}

func deleteCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.CustomDomain.Delete(d.Id())
}

func buildCustomDomain(d *schema.ResourceData) *management.CustomDomain {
	return &management.CustomDomain{
		Domain:             String(d, "domain"),
		Type:               String(d, "type"),
		VerificationMethod: String(d, "verification_method"),
	}
}
