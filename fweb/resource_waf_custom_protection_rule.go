// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/custom_protection_rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCustomProtectionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCustomProtectionRuleRead,
		Update: resourceWafCustomProtectionRuleUpdate,
		Create: resourceWafCustomProtectionRuleCreate,
		Delete: resourceWafCustomProtectionRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			//High, Meduim, Low, Info
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//critical, severe, substantial, moderate, low, informational
			"threat_weight": &schema.Schema{
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

func resourceWafCustomProtectionRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionRule: type error")
	}

	obj, err := getObjectWafCustomProtectionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafCustomProtectionRule(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCustomProtectionRuleRead(d, m)
}

func resourceWafCustomProtectionRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafCustomProtectionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCustomProtectionRule resource while getting object: %v", err)
	}

	_, err = c.UpdateWafCustomProtectionRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCustomProtectionRule resource: %v", err)
	}

	return resourceWafCustomProtectionRuleRead(d, m)
}

func resourceWafCustomProtectionRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafCustomProtectionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCustomProtectionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCustomProtectionRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafCustomProtectionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCustomProtectionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionRule resource from API: %v", err)
	}

	return nil
}

func flattenWafCustomProtectionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCustomProtectionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenWafCustomProtectionRule(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("action", flattenWafCustomProtectionRule(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}
	if err = d.Set("block_period", flattenWafCustomProtectionRule(o["block-period"], d, "block_period", sv)); err != nil {
		if !fortiAPIPatch(o["block-period"]) {
			return fmt.Errorf("Error reading block_period: %v", err)
		}
	}
	if err = d.Set("severity", flattenWafCustomProtectionRule(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}
	if err = d.Set("trigger", flattenWafCustomProtectionRule(o["trigger"], d, "trigger", sv)); err != nil {
		if !fortiAPIPatch(o["trigger"]) {
			return fmt.Errorf("Error reading trigger: %v", err)
		}
	}
	if err = d.Set("threat_weight", flattenWafCustomProtectionRule(o["threat-weight"], d, "threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["threat-weight"]) {
			return fmt.Errorf("Error reading threat_weight: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafCustomProtectionRule(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafCustomProtectionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCustomProtectionRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("action"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}
	if v, ok := d.GetOk("block_period"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-period"] = t
		}
	}
	if v, ok := d.GetOk("severity"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	if v, ok := d.GetOk("trigger"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger"] = t
		}
	}
	if v, ok := d.GetOk("threat_weight"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCustomProtectionRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
