// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system dns.

package fortiweb

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Create: resourceSystemDnsCreate,
		Delete: resourceSystemDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type: schema.TypeString,
				//Computed: true,
				Optional: true,
			},
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceSystemDnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""
	if v, ok := d.GetOk("mkey"); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	if mkey == "" {
		mkey = "system_dns"
	}
	d.SetId(mkey)

	obj, err := getObjectSystemDns(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := d.Id()
	if mkey == "" {
		mkey = "system_dns"
	}
	d.SetId(mkey)

	obj, err := getObjectSystemDns(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := d.Id()
	if mkey == "" {
		mkey = "system_dns"
	}

	//special case: reset default value
	obj := make(map[string]interface{})
	obj["domain"] = ""
	obj["primary"] = "8.8.8.8"
	obj["secondary"] = "0.0.0.0"
	_, err := c.UpdateSystemDns(&obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := d.Id()
	if mkey == "" {
		mkey = "system_dns"
	}
	d.SetId(mkey)

	o, err := c.ReadSystemDns(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] DNS resource (%s) returned nil, but keeping in state", d.Id())
		o = make(map[string]interface{})
	}

	err = refreshObjectSystemDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}

	c.Logout()

	return nil
}
func flattenSystemDns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("domain", flattenSystemDns(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("primary", flattenSystemDns(o["primary"], d, "primary", sv)); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemDns(o["secondary"], d, "secondary", sv)); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}
	return nil
}

func expandSystemDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDns(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if !ok {
		//default value
		obj["domain"] = ""
	}

	if v, ok := d.GetOk("primary"); ok {
		t, err := expandSystemDns(d, v, "primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	} else if !ok {
		//default value
		obj["primary"] = "8.8.8.8"
	}

	if v, ok := d.GetOk("secondary"); ok {
		t, err := expandSystemDns(d, v, "secondary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	} else if !ok {
		//default value
		obj["secondary"] = "0.0.0.0"
	}

	return &obj, nil
}
