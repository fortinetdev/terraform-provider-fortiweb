// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/cors_protection_rule/allowed-headers-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCorsProtectionRuleAllowedHeadersList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCorsProtectionRuleAllowedHeadersListRead,
		Update: resourceWafCorsProtectionRuleAllowedHeadersListUpdate,
		Create: resourceWafCorsProtectionRuleAllowedHeadersListCreate,
		Delete: resourceWafCorsProtectionRuleAllowedHeadersListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

func resourceWafCorsProtectionRuleAllowedHeadersListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedHeadersList: type error")
	}

	obj, err := getObjectWafCorsProtectionRuleAllowedHeadersList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRuleAllowedHeadersList resource while getting object: %v", err)
	}

	_, err = c.CreateWafCorsProtectionRuleAllowedHeadersList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRuleAllowedHeadersList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCorsProtectionRuleAllowedHeadersListRead(d, m)
}

func resourceWafCorsProtectionRuleAllowedHeadersListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedHeadersList: type error")
	}

	obj, err := getObjectWafCorsProtectionRuleAllowedHeadersList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRuleAllowedHeadersList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafCorsProtectionRuleAllowedHeadersList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRuleAllowedHeadersList resource: %v", err)
	}

	return resourceWafCorsProtectionRuleAllowedHeadersListRead(d, m)
}

func resourceWafCorsProtectionRuleAllowedHeadersListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedHeadersList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafCorsProtectionRuleAllowedHeadersList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCorsProtectionRuleAllowedHeadersList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCorsProtectionRuleAllowedHeadersListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedHeadersList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafCorsProtectionRuleAllowedHeadersList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRuleAllowedHeadersList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCorsProtectionRuleAllowedHeadersList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRuleAllowedHeadersList resource from API: %v", err)
	}

	return nil
}

func flattenWafCorsProtectionRuleAllowedHeadersList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCorsProtectionRuleAllowedHeadersList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafCorsProtectionRuleAllowedHeadersList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("header", flattenWafCorsProtectionRuleAllowedHeadersList(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafCorsProtectionRuleAllowedHeadersList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafCorsProtectionRuleAllowedHeadersList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCorsProtectionRuleAllowedHeadersList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafCorsProtectionRuleAllowedHeadersList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("header"); ok {
		t, err := expandWafCorsProtectionRuleAllowedHeadersList(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCorsProtectionRuleAllowedHeadersList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
