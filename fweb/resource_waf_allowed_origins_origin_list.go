// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/allowed_origins/origin-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafAllowedOriginsOriginList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafAllowedOriginsOriginListRead,
		Update: resourceWafAllowedOriginsOriginListUpdate,
		Create: resourceWafAllowedOriginsOriginListCreate,
		Delete: resourceWafAllowedOriginsOriginListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"origin_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"include_sub_domains": &schema.Schema{
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

func resourceWafAllowedOriginsOriginListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowedOriginsOriginList: type error")
	}

	obj, err := getObjectWafAllowedOriginsOriginList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowedOriginsOriginList resource while getting object: %v", err)
	}

	_, err = c.CreateWafAllowedOriginsOriginList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowedOriginsOriginList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafAllowedOriginsOriginListRead(d, m)
}

func resourceWafAllowedOriginsOriginListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowedOriginsOriginList: type error")
	}

	obj, err := getObjectWafAllowedOriginsOriginList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowedOriginsOriginList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafAllowedOriginsOriginList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowedOriginsOriginList resource: %v", err)
	}

	return resourceWafAllowedOriginsOriginListRead(d, m)
}

func resourceWafAllowedOriginsOriginListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowedOriginsOriginList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafAllowedOriginsOriginList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafAllowedOriginsOriginList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafAllowedOriginsOriginListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowedOriginsOriginList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafAllowedOriginsOriginList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowedOriginsOriginList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafAllowedOriginsOriginList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowedOriginsOriginList resource from API: %v", err)
	}

	return nil
}

func flattenWafAllowedOriginsOriginList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafAllowedOriginsOriginList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafAllowedOriginsOriginList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}
	if err = d.Set("protocol", flattenWafAllowedOriginsOriginList(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}
	if err = d.Set("origin_name", flattenWafAllowedOriginsOriginList(o["origin-name"], d, "origin_name", sv)); err != nil {
		if !fortiAPIPatch(o["origin-name"]) {
			return fmt.Errorf("Error reading origin_name: %v", err)
		}
	}
	if err = d.Set("port", flattenWafAllowedOriginsOriginList(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}
	if err = d.Set("include_sub_domains", flattenWafAllowedOriginsOriginList(o["include-sub-domains"], d, "include_sub_domains", sv)); err != nil {
		if !fortiAPIPatch(o["include-sub-domains"]) {
			return fmt.Errorf("Error reading include_sub_domains: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafAllowedOriginsOriginList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafAllowedOriginsOriginList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafAllowedOriginsOriginList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafAllowedOriginsOriginList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandWafAllowedOriginsOriginList(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}
	if v, ok := d.GetOk("origin_name"); ok {
		t, err := expandWafAllowedOriginsOriginList(d, v, "origin_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["origin-name"] = t
		}
	}
	if v, ok := d.GetOk("port"); ok {
		t, err := expandWafAllowedOriginsOriginList(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}
	if v, ok := d.GetOk("include_sub_domains"); ok {
		t, err := expandWafAllowedOriginsOriginList(d, v, "include_sub_domains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-sub-domains"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafAllowedOriginsOriginList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
