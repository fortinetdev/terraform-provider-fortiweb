// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_encryption.url_encryption_rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlEncryptionUrlEncryptionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlEncryptionUrlEncryptionRuleRead,
		Update: resourceWafUrlEncryptionUrlEncryptionRuleUpdate,
		Create: resourceWafUrlEncryptionUrlEncryptionRuleCreate,
		Delete: resourceWafUrlEncryptionUrlEncryptionRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_unencrypted": &schema.Schema{
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

func resourceWafUrlEncryptionUrlEncryptionRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRule: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlEncryptionUrlEncryptionRule(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlEncryptionUrlEncryptionRuleRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionRule resource while getting object: %v", err)
	}

	_, err = c.UpdateWafUrlEncryptionUrlEncryptionRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionRule resource: %v", err)
	}

	return resourceWafUrlEncryptionUrlEncryptionRuleRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafUrlEncryptionUrlEncryptionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlEncryptionUrlEncryptionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlEncryptionUrlEncryptionRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafUrlEncryptionUrlEncryptionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlEncryptionUrlEncryptionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionRule resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlEncryptionUrlEncryptionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlEncryptionUrlEncryptionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("host_status", flattenWafUrlEncryptionUrlEncryptionRule(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}
	if err = d.Set("host", flattenWafUrlEncryptionUrlEncryptionRule(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}
	if err = d.Set("allow_unencrypted", flattenWafUrlEncryptionUrlEncryptionRule(o["allow-unencrypted"], d, "allow_unencrypted", sv)); err != nil {
		if !fortiAPIPatch(o["allow-unencrypted"]) {
			return fmt.Errorf("Error reading allow_unencrypted: %v", err)
		}
	}
	if err = d.Set("action", flattenWafUrlEncryptionUrlEncryptionRule(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}
	if err = d.Set("block_period", flattenWafUrlEncryptionUrlEncryptionRule(o["block-period"], d, "block_period", sv)); err != nil {
		if !fortiAPIPatch(o["block-period"]) {
			return fmt.Errorf("Error reading block_period: %v", err)
		}
	}
	if err = d.Set("severity", flattenWafUrlEncryptionUrlEncryptionRule(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}
	if err = d.Set("trigger", flattenWafUrlEncryptionUrlEncryptionRule(o["trigger"], d, "trigger", sv)); err != nil {
		if !fortiAPIPatch(o["trigger"]) {
			return fmt.Errorf("Error reading trigger: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlEncryptionUrlEncryptionRule(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlEncryptionUrlEncryptionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlEncryptionUrlEncryptionRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}
	if v, ok := d.GetOk("host"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("allow_unencrypted"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "allow_unencrypted", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-unencrypted"] = t
		}
	}
	if v, ok := d.GetOk("action"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}
	if v, ok := d.GetOk("block_period"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-period"] = t
		}
	}
	if v, ok := d.GetOk("severity"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	if v, ok := d.GetOk("trigger"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
