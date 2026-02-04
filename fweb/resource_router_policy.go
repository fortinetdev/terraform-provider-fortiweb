// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure router/policy.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterPolicyRead,
		Update: resourceRouterPolicyUpdate,
		Create: resourceRouterPolicyCreate,
		Delete: resourceRouterPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"iif": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"oif": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"fwmark": &schema.Schema{
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

func resourceRouterPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterPolicy: type error")
	}

	obj, err := getObjectRouterPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateRouterPolicy(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterPolicyRead(d, m)
}

func resourceRouterPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterPolicy(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy resource: %v", err)
	}

	return resourceRouterPolicyRead(d, m)
}

func resourceRouterPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy resource from API: %v", err)
	}

	return nil
}

func flattenRouterPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("src", flattenRouterPolicy(o["src"], d, "src", sv)); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}
	if err = d.Set("dst", flattenRouterPolicy(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}
	if err = d.Set("action", flattenRouterPolicy(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}
	if err = d.Set("gateway", flattenRouterPolicy(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}
	if err = d.Set("iif", flattenRouterPolicy(o["iif"], d, "iif", sv)); err != nil {
		if !fortiAPIPatch(o["iif"]) {
			return fmt.Errorf("Error reading iif: %v", err)
		}
	}
	if err = d.Set("oif", flattenRouterPolicy(o["oif"], d, "oif", sv)); err != nil {
		if !fortiAPIPatch(o["oif"]) {
			return fmt.Errorf("Error reading oif: %v", err)
		}
	}
	if err = d.Set("priority", flattenRouterPolicy(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}
	if err = d.Set("fwmark", flattenRouterPolicy(o["fwmark"], d, "fwmark", sv)); err != nil {
		if !fortiAPIPatch(o["fwmark"]) {
			return fmt.Errorf("Error reading fwmark: %v", err)
		}
	}
	if err = d.Set("mkey", flattenRouterPolicy(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("src"); ok {
		t, err := expandRouterPolicy(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}
	if v, ok := d.GetOk("dst"); ok {
		t, err := expandRouterPolicy(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}
	if v, ok := d.GetOk("action"); ok {
		t, err := expandRouterPolicy(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}
	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterPolicy(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}
	if v, ok := d.GetOk("iif"); ok {
		t, err := expandRouterPolicy(d, v, "iif", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iif"] = t
		}
	}
	if v, ok := d.GetOk("oif"); ok {
		t, err := expandRouterPolicy(d, v, "oif", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oif"] = t
		}
	}
	if v, ok := d.GetOk("priority"); ok {
		t, err := expandRouterPolicy(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}
	if v, ok := d.GetOk("fwmark"); ok {
		t, err := expandRouterPolicy(d, v, "fwmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fwmark"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterPolicy(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
