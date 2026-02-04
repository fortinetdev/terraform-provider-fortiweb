// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/custom_protection_group.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCustomProtectionGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCustomProtectionGroupRead,
		Update: resourceWafCustomProtectionGroupUpdate,
		Create: resourceWafCustomProtectionGroupCreate,
		Delete: resourceWafCustomProtectionGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_alert_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

func resourceWafCustomProtectionGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionGroup: type error")
	}

	obj, err := getObjectWafCustomProtectionGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionGroup resource while getting object: %v", err)
	}

	_, err = c.CreateWafCustomProtectionGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCustomProtectionGroupRead(d, m)
}

func resourceWafCustomProtectionGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceWafCustomProtectionGroupRead(d, m)
}

func resourceWafCustomProtectionGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafCustomProtectionGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCustomProtectionGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCustomProtectionGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafCustomProtectionGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCustomProtectionGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionGroup resource from API: %v", err)
	}

	return nil
}

func flattenWafCustomProtectionGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCustomProtectionGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_alert_interval", flattenWafCustomProtectionGroup(o["max-alert-interval"], d, "max_alert_interval", sv)); err != nil {
		if !fortiAPIPatch(o["max-alert-interval"]) {
			return fmt.Errorf("Error reading max_alert_interval: %v", err)
		}
	}
	if err = d.Set("custom_protection_rule", flattenWafCustomProtectionGroup(o["custom-protection-rule"], d, "custom_protection_rule", sv)); err != nil {
		if !fortiAPIPatch(o["custom-protection-rule"]) {
			return fmt.Errorf("Error reading custom_protection_rule: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafCustomProtectionGroup(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafCustomProtectionGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCustomProtectionGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_alert_interval"); ok {
		t, err := expandWafCustomProtectionGroup(d, v, "max_alert_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-alert-interval"] = t
		}
	}
	if v, ok := d.GetOk("custom_protection_rule"); ok {
		t, err := expandWafCustomProtectionGroup(d, v, "custom_protection_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-protection-rule"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCustomProtectionGroup(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
