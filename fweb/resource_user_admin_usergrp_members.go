// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/admin_usergrp/members.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserAdminUsergrpMembers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserAdminUsergrpMembersRead,
		Update: resourceUserAdminUsergrpMembersUpdate,
		Create: resourceUserAdminUsergrpMembersCreate,
		Delete: resourceUserAdminUsergrpMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ldap_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"radius_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"pki_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tacacs_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_mkey": &schema.Schema{
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

func resourceUserAdminUsergrpMembersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserAdminUsergrpMembers: type error")
	}

	obj, err := getObjectUserAdminUsergrpMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserAdminUsergrpMembers resource while getting object: %v", err)
	}

	_, err = c.CreateUserAdminUsergrpMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserAdminUsergrpMembers resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserAdminUsergrpMembersRead(d, m)
}

func resourceUserAdminUsergrpMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserAdminUsergrpMembers: type error")
	}

	obj, err := getObjectUserAdminUsergrpMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserAdminUsergrpMembers resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateUserAdminUsergrpMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserAdminUsergrpMembers resource: %v", err)
	}

	return resourceUserAdminUsergrpMembersRead(d, m)
}

func resourceUserAdminUsergrpMembersDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserAdminUsergrpMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteUserAdminUsergrpMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserAdminUsergrpMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserAdminUsergrpMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserAdminUsergrpMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadUserAdminUsergrpMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserAdminUsergrpMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserAdminUsergrpMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserAdminUsergrpMembers resource from API: %v", err)
	}

	return nil
}

func flattenUserAdminUsergrpMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserAdminUsergrpMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenUserAdminUsergrpMembers(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}
	if err = d.Set("type", flattenUserAdminUsergrpMembers(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("ldap_name", flattenUserAdminUsergrpMembers(o["ldap-name"], d, "ldap_name", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-name"]) {
			return fmt.Errorf("Error reading ldap_name: %v", err)
		}
	}
	if err = d.Set("radius_name", flattenUserAdminUsergrpMembers(o["radius-name"], d, "radius_name", sv)); err != nil {
		if !fortiAPIPatch(o["radius-name"]) {
			return fmt.Errorf("Error reading radius_name: %v", err)
		}
	}
	if err = d.Set("pki_name", flattenUserAdminUsergrpMembers(o["pki-name"], d, "pki_name", sv)); err != nil {
		if !fortiAPIPatch(o["pki-name"]) {
			return fmt.Errorf("Error reading pki_name: %v", err)
		}
	}
	if err = d.Set("tacacs_name", flattenUserAdminUsergrpMembers(o["tacacs+-name"], d, "tacacs_name", sv)); err != nil {
		if !fortiAPIPatch(o["tacacs+-name"]) {
			return fmt.Errorf("Error reading tacacs+_name: %v", err)
		}
	}
	if err = d.Set("group_name", flattenUserAdminUsergrpMembers(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserAdminUsergrpMembers(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserAdminUsergrpMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserAdminUsergrpMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("ldap_name"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "ldap_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-name"] = t
		}
	}
	if v, ok := d.GetOk("radius_name"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "radius_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-name"] = t
		}
	}
	if v, ok := d.GetOk("pki_name"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "pki_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pki-name"] = t
		}
	}
	if v, ok := d.GetOk("tacacs_name"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "tacacs_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tacacs+-name"] = t
		}
	}
	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "group_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserAdminUsergrpMembers(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
