// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/ldap_user/extracted-attributes.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserLdapUserExtractedAttributes() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserLdapUserExtractedAttributesRead,
		Update: resourceUserLdapUserExtractedAttributesUpdate,
		Create: resourceUserLdapUserExtractedAttributesCreate,
		Delete: resourceUserLdapUserExtractedAttributesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceUserLdapUserExtractedAttributesCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing UserLdapUserExtractedAttributes: type error")
	}

	obj, err := getObjectUserLdapUserExtractedAttributes(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserLdapUserExtractedAttributes resource while getting object: %v", err)
	}

	_, err = c.CreateUserLdapUserExtractedAttributes(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserLdapUserExtractedAttributes resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserLdapUserExtractedAttributesRead(d, m)
}

func resourceUserLdapUserExtractedAttributesUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserLdapUserExtractedAttributes(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserLdapUserExtractedAttributes resource while getting object: %v", err)
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing UserLdapUserExtractedAttributes: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey

	_, err = c.UpdateUserLdapUserExtractedAttributes(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserLdapUserExtractedAttributes resource: %v", err)
	}

	return resourceUserLdapUserExtractedAttributesRead(d, m)
}

func resourceUserLdapUserExtractedAttributesDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing UserLdapUserExtractedAttributes: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey

	err := c.DeleteUserLdapUserExtractedAttributes(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserLdapUserExtractedAttributes resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapUserExtractedAttributesRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing UserLdapUserExtractedAttributes: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey

	o, err := c.ReadUserLdapUserExtractedAttributes(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserLdapUserExtractedAttributes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdapUserExtractedAttributes(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserLdapUserExtractedAttributes resource from API: %v", err)
	}

	return nil
}

func flattenUserLdapUserExtractedAttributes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserLdapUserExtractedAttributes(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenUserLdapUserExtractedAttributes(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}
	if err = d.Set("name", flattenUserLdapUserExtractedAttributes(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("attribute", flattenUserLdapUserExtractedAttributes(o["attribute"], d, "attribute", sv)); err != nil {
		if !fortiAPIPatch(o["attribute"]) {
			return fmt.Errorf("Error reading attribute: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserLdapUserExtractedAttributes(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserLdapUserExtractedAttributes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserLdapUserExtractedAttributes(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandUserLdapUserExtractedAttributes(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserLdapUserExtractedAttributes(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("attribute"); ok {
		t, err := expandUserLdapUserExtractedAttributes(d, v, "attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserLdapUserExtractedAttributes(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
