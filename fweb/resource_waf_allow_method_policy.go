// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/allow_method_policy.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafAllowMethodPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafAllowMethodPolicyRead,
		Update: resourceWafAllowMethodPolicyUpdate,
		Create: resourceWafAllowMethodPolicyCreate,
		Delete: resourceWafAllowMethodPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"allow_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"override_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"override_parameter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"triggered_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_method_exception": &schema.Schema{
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

func resourceWafAllowMethodPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowMethodPolicy: type error")
	}

	obj, err := getObjectWafAllowMethodPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowMethodPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateWafAllowMethodPolicy(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowMethodPolicy resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafAllowMethodPolicyRead(d, m)
}

func resourceWafAllowMethodPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafAllowMethodPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowMethodPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateWafAllowMethodPolicy(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowMethodPolicy resource: %v", err)
	}

	return resourceWafAllowMethodPolicyRead(d, m)
}

func resourceWafAllowMethodPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafAllowMethodPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafAllowMethodPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafAllowMethodPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafAllowMethodPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowMethodPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafAllowMethodPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowMethodPolicy resource from API: %v", err)
	}

	return nil
}

func flattenWafAllowMethodPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafAllowMethodPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("allow_method", flattenWafAllowMethodPolicy(o["allow-method"], d, "allow_method", sv)); err != nil {
		if !fortiAPIPatch(o["allow-method"]) {
			return fmt.Errorf("Error reading allow_method: %v", err)
		}
	}
	if err = d.Set("override_header", flattenWafAllowMethodPolicy(o["override-header"], d, "override_header", sv)); err != nil {
		if !fortiAPIPatch(o["override-header"]) {
			return fmt.Errorf("Error reading override_header: %v", err)
		}
	}
	if err = d.Set("override_parameter", flattenWafAllowMethodPolicy(o["override-parameter"], d, "override_parameter", sv)); err != nil {
		if !fortiAPIPatch(o["override-parameter"]) {
			return fmt.Errorf("Error reading override_parameter: %v", err)
		}
	}
	if err = d.Set("severity", flattenWafAllowMethodPolicy(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}
	if err = d.Set("triggered_action", flattenWafAllowMethodPolicy(o["triggered-action"], d, "triggered_action", sv)); err != nil {
		if !fortiAPIPatch(o["triggered-action"]) {
			return fmt.Errorf("Error reading triggered_action: %v", err)
		}
	}
	if err = d.Set("allow_method_exception", flattenWafAllowMethodPolicy(o["allow-method-exception"], d, "allow_method_exception", sv)); err != nil {
		if !fortiAPIPatch(o["allow-method-exception"]) {
			return fmt.Errorf("Error reading allow_method_exception: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafAllowMethodPolicy(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafAllowMethodPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafAllowMethodPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("allow_method"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "allow_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-method"] = t
		}
	}
	if v, ok := d.GetOk("override_header"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "override_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-header"] = t
		}
	}
	if v, ok := d.GetOk("override_parameter"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "override_parameter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-parameter"] = t
		}
	}
	if v, ok := d.GetOk("severity"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	if v, ok := d.GetOk("triggered_action"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "triggered_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["triggered-action"] = t
		}
	}
	if v, ok := d.GetOk("allow_method_exception"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "allow_method_exception", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-method-exception"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafAllowMethodPolicy(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
