// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/custom_protection_rule/meet-condition.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCustomProtectionRuleMeetCondition() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCustomProtectionRuleMeetConditionRead,
		Update: resourceWafCustomProtectionRuleMeetConditionUpdate,
		Create: resourceWafCustomProtectionRuleMeetConditionCreate,
		Delete: resourceWafCustomProtectionRuleMeetConditionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"operator": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"expression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"response_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceWafCustomProtectionRuleMeetConditionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionRuleMeetCondition: type error")
	}

	obj, err := getObjectWafCustomProtectionRuleMeetCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionRuleMeetCondition resource while getting object: %v", err)
	}

	_, err = c.CreateWafCustomProtectionRuleMeetCondition(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCustomProtectionRuleMeetCondition resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCustomProtectionRuleMeetConditionRead(d, m)
}

func resourceWafCustomProtectionRuleMeetConditionUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionRuleMeetCondition: type error")
	}

	obj, err := getObjectWafCustomProtectionRuleMeetCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCustomProtectionRuleMeetCondition resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafCustomProtectionRuleMeetCondition(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCustomProtectionRuleMeetCondition resource: %v", err)
	}

	return resourceWafCustomProtectionRuleMeetConditionRead(d, m)
}

func resourceWafCustomProtectionRuleMeetConditionDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionRuleMeetCondition: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafCustomProtectionRuleMeetCondition(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCustomProtectionRuleMeetCondition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCustomProtectionRuleMeetConditionRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCustomProtectionRuleMeetCondition: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafCustomProtectionRuleMeetCondition(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionRuleMeetCondition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCustomProtectionRuleMeetCondition(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCustomProtectionRuleMeetCondition resource from API: %v", err)
	}

	return nil
}

func flattenWafCustomProtectionRuleMeetCondition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCustomProtectionRuleMeetCondition(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafCustomProtectionRuleMeetCondition(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("operator", flattenWafCustomProtectionRuleMeetCondition(o["operator"], d, "operator", sv)); err != nil {
		if !fortiAPIPatch(o["operator"]) {
			return fmt.Errorf("Error reading operator: %v", err)
		}
	}
	if err = d.Set("request_target", flattenWafCustomProtectionRuleMeetCondition(o["request-target"], d, "request_target", sv)); err != nil {
		if !fortiAPIPatch(o["request-target"]) {
			return fmt.Errorf("Error reading request_target: %v", err)
		}
	}
	if err = d.Set("response_target", flattenWafCustomProtectionRuleMeetCondition(o["response-target"], d, "response_target", sv)); err != nil {
		if !fortiAPIPatch(o["response-target"]) {
			return fmt.Errorf("Error reading response_target: %v", err)
		}
	}
	if err = d.Set("threshold", flattenWafCustomProtectionRuleMeetCondition(o["threshold"], d, "threshold", sv)); err != nil {
		if !fortiAPIPatch(o["threshold"]) {
			return fmt.Errorf("Error reading threshold: %v", err)
		}
	}
	if err = d.Set("case_sensitive", flattenWafCustomProtectionRuleMeetCondition(o["case-sensitive"], d, "case_sensitive", sv)); err != nil {
		if !fortiAPIPatch(o["case-sensitive"]) {
			return fmt.Errorf("Error reading case_sensitive: %v", err)
		}
	}
	if err = d.Set("expression", flattenWafCustomProtectionRuleMeetCondition(o["expression"], d, "expression", sv)); err != nil {
		if !fortiAPIPatch(o["expression"]) {
			return fmt.Errorf("Error reading expression: %v", err)
		}
	}

	return nil
}

func expandWafCustomProtectionRuleMeetCondition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCustomProtectionRuleMeetCondition(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("operator"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "operator", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["operator"] = t
		}
	}
	if v, ok := d.GetOk("request_target"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "request_target", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-target"] = t
		}
	}
	if v, ok := d.GetOk("response_target"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "response_target", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-target"] = t
		}
	}
	if v, ok := d.GetOk("threshold"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold"] = t
		}
	}
	if v, ok := d.GetOk("case_sensitive"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "case_sensitive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["case-sensitive"] = t
		}
	}
	if v, ok := d.GetOk("expression"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "expression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expression"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCustomProtectionRuleMeetCondition(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
