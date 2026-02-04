// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_access.url_access_rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlAccessUrlAccessRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlAccessUrlAccessRuleRead,
		Update: resourceWafUrlAccessUrlAccessRuleUpdate,
		Create: resourceWafUrlAccessUrlAccessRuleCreate,
		Delete: resourceWafUrlAccessUrlAccessRuleDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceWafUrlAccessUrlAccessRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessRule: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlAccessUrlAccessRule(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlAccessUrlAccessRuleRead(d, m)
}

func resourceWafUrlAccessUrlAccessRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafUrlAccessUrlAccessRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessRule resource while getting object: %v", err)
	}

	_, err = c.UpdateWafUrlAccessUrlAccessRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessRule resource: %v", err)
	}

	return resourceWafUrlAccessUrlAccessRuleRead(d, m)
}

func resourceWafUrlAccessUrlAccessRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafUrlAccessUrlAccessRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlAccessUrlAccessRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlAccessUrlAccessRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafUrlAccessUrlAccessRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlAccessUrlAccessRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessRule resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlAccessUrlAccessRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlAccessUrlAccessRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("host_status", flattenWafUrlAccessUrlAccessRule(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}
	if err = d.Set("host", flattenWafUrlAccessUrlAccessRule(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}
	if err = d.Set("action", flattenWafUrlAccessUrlAccessRule(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}
	if err = d.Set("severity", flattenWafUrlAccessUrlAccessRule(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}
	if err = d.Set("trigger", flattenWafUrlAccessUrlAccessRule(o["trigger"], d, "trigger", sv)); err != nil {
		if !fortiAPIPatch(o["trigger"]) {
			return fmt.Errorf("Error reading trigger: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlAccessUrlAccessRule(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlAccessUrlAccessRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlAccessUrlAccessRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafUrlAccessUrlAccessRule(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}
	if v, ok := d.GetOk("host"); ok {
		t, err := expandWafUrlAccessUrlAccessRule(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("action"); ok {
		t, err := expandWafUrlAccessUrlAccessRule(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}
	if v, ok := d.GetOk("severity"); ok {
		t, err := expandWafUrlAccessUrlAccessRule(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	if v, ok := d.GetOk("trigger"); ok {
		t, err := expandWafUrlAccessUrlAccessRule(d, v, "trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
