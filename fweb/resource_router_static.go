// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure router static.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterStaticRead,
		Update: resourceRouterStaticUpdate,
		Create: resourceRouterStaticCreate,
		Delete: resourceRouterStaticDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"device": &schema.Schema{
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

func resourceRouterStaticCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterStatic: type error")
	}

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic resource while getting object: %v", err)
	}

	_, err = c.CreateRouterStatic(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterStaticRead(d, m)
}

func resourceRouterStaticUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterStatic(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic resource: %v", err)
	}

	return resourceRouterStaticRead(d, m)
}

func resourceRouterStaticDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterStatic(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterStatic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStaticRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterStatic(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic resource from API: %v", err)
	}

	return nil
}
func flattenRouterStatic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterStatic(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dst", flattenRouterStatic(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterStatic(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterStatic(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterStatic(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterStatic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandRouterStatic(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterStatic(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandRouterStatic(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterStatic(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
