// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_access.url_access_parameter.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlAccessUrlAccessParameter() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlAccessUrlAccessParameterRead,
		Update: resourceWafUrlAccessUrlAccessParameterUpdate,
		Create: resourceWafUrlAccessUrlAccessParameterCreate,
		Delete: resourceWafUrlAccessUrlAccessParameterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"severity": &schema.Schema{
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

func resourceWafUrlAccessUrlAccessParameterCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessParameter: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessParameter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessParameter resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlAccessUrlAccessParameter(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessParameter resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlAccessUrlAccessParameterRead(d, m)
}

func resourceWafUrlAccessUrlAccessParameterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafUrlAccessUrlAccessParameter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessParameter resource while getting object: %v", err)
	}

	_, err = c.UpdateWafUrlAccessUrlAccessParameter(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessParameter resource: %v", err)
	}

	return resourceWafUrlAccessUrlAccessParameterRead(d, m)
}

func resourceWafUrlAccessUrlAccessParameterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafUrlAccessUrlAccessParameter(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlAccessUrlAccessParameter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlAccessUrlAccessParameterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafUrlAccessUrlAccessParameter(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessParameter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlAccessUrlAccessParameter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessParameter resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlAccessUrlAccessParameter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlAccessUrlAccessParameter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("action", flattenWafUrlAccessUrlAccessParameter(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("host_status", flattenWafUrlAccessUrlAccessParameter(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}

	if err = d.Set("severity", flattenWafUrlAccessUrlAccessParameter(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("mkey", flattenWafUrlAccessUrlAccessParameter(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlAccessUrlAccessParameter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlAccessUrlAccessParameter(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandWafUrlAccessUrlAccessParameter(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafUrlAccessUrlAccessParameter(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok {
		t, err := expandWafUrlAccessUrlAccessParameter(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessParameter(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
