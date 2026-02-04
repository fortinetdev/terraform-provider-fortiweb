// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/custom_protection_group/type-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCustomProtectionGroupTypeList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCustomProtectionGroupTypeListRead,
		Update: resourceWafCustomProtectionGroupTypeListUpdate,
		Create: resourceWafCustomProtectionGroupTypeListCreate,
		Delete: resourceWafCustomProtectionGroupTypeListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_protection_rule": &schema.Schema{
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

func resourceWafCustomProtectionGroupTypeListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionGroupTypeList: type error")
	}

	obj, err := getObjectWafCustomProtectionGroupTypeList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionGroupTypeList resource while getting object: %v", err)
	}

	_, err = c.CreateWafCustomProtectionGroupTypeList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionGroupTypeList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCustomProtectionGroupTypeListRead(d, m)
}

func resourceWafCustomProtectionGroupTypeListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionGroupTypeList: type error")
	}

	obj, err := getObjectWafCustomProtectionGroupTypeList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCustomProtectionGroupTypeList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafCustomProtectionGroupTypeList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCustomProtectionGroupTypeList resource: %v", err)
	}

	return resourceWafCustomProtectionGroupTypeListRead(d, m)
}

func resourceWafCustomProtectionGroupTypeListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionGroupTypeList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafCustomProtectionGroupTypeList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCustomProtectionGroupTypeList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCustomProtectionGroupTypeListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionGroupTypeList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafCustomProtectionGroupTypeList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionGroupTypeList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCustomProtectionGroupTypeList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionGroupTypeList resource from API: %v", err)
	}

	return nil
}

func flattenWafCustomProtectionGroupTypeList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCustomProtectionGroupTypeList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafCustomProtectionGroupTypeList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("custom_protection_rule", flattenWafCustomProtectionGroupTypeList(o["custom-protection-rule"], d, "custom_protection_rule", sv)); err != nil {
		if !fortiAPIPatch(o["custom-protection-rule"]) {
			return fmt.Errorf("Error reading custom_protection_rule: %v", err)
		}
	}

	return nil
}

func expandWafCustomProtectionGroupTypeList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCustomProtectionGroupTypeList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("id"); ok {
		t, err := expandWafCustomProtectionGroupTypeList(d, v, "id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("custom_protection_rule"); ok {
		t, err := expandWafCustomProtectionGroupTypeList(d, v, "custom_protection_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-protection-rule"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCustomProtectionGroupTypeList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
