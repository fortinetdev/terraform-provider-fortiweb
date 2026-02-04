// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/cors_protection_rule/exposed-headers-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCorsProtectionRuleExposedHeadersList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCorsProtectionRuleExposedHeadersListRead,
		Update: resourceWafCorsProtectionRuleExposedHeadersListUpdate,
		Create: resourceWafCorsProtectionRuleExposedHeadersListCreate,
		Delete: resourceWafCorsProtectionRuleExposedHeadersListDelete,

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

func resourceWafCorsProtectionRuleExposedHeadersListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleExposedHeadersList: type error")
	}

	obj, err := getObjectWafCorsProtectionRuleExposedHeadersList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRuleExposedHeadersList resource while getting object: %v", err)
	}

	_, err = c.CreateWafCorsProtectionRuleExposedHeadersList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRuleExposedHeadersList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCorsProtectionRuleExposedHeadersListRead(d, m)
}

func resourceWafCorsProtectionRuleExposedHeadersListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleExposedHeadersList: type error")
	}

	obj, err := getObjectWafCorsProtectionRuleExposedHeadersList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRuleExposedHeadersList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafCorsProtectionRuleExposedHeadersList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRuleExposedHeadersList resource: %v", err)
	}

	return resourceWafCorsProtectionRuleExposedHeadersListRead(d, m)
}

func resourceWafCorsProtectionRuleExposedHeadersListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleExposedHeadersList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafCorsProtectionRuleExposedHeadersList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCorsProtectionRuleExposedHeadersList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCorsProtectionRuleExposedHeadersListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleExposedHeadersList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafCorsProtectionRuleExposedHeadersList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRuleExposedHeadersList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCorsProtectionRuleExposedHeadersList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRuleExposedHeadersList resource from API: %v", err)
	}

	return nil
}

func flattenWafCorsProtectionRuleExposedHeadersList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCorsProtectionRuleExposedHeadersList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafCorsProtectionRuleExposedHeadersList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("header", flattenWafCorsProtectionRuleExposedHeadersList(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafCorsProtectionRuleExposedHeadersList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafCorsProtectionRuleExposedHeadersList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCorsProtectionRuleExposedHeadersList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafCorsProtectionRuleExposedHeadersList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("header"); ok {
		t, err := expandWafCorsProtectionRuleExposedHeadersList(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCorsProtectionRuleExposedHeadersList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
