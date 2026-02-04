// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/kerberos_user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserKerberosUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserKerberosUserRead,
		Update: resourceUserKerberosUserUpdate,
		Create: resourceUserKerberosUserCreate,
		Delete: resourceUserKerberosUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"shortname": &schema.Schema{
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

func resourceUserKerberosUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserKerberosUser: type error")
	}

	obj, err := getObjectUserKerberosUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserKerberosUser resource while getting object: %v", err)
	}

	_, err = c.CreateUserKerberosUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserKerberosUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserKerberosUserRead(d, m)
}

func resourceUserKerberosUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserKerberosUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserKerberosUser resource while getting object: %v", err)
	}

	_, err = c.UpdateUserKerberosUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserKerberosUser resource: %v", err)
	}

	return resourceUserKerberosUserRead(d, m)
}

func resourceUserKerberosUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserKerberosUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserKerberosUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserKerberosUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserKerberosUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserKerberosUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserKerberosUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserKerberosUser resource from API: %v", err)
	}

	return nil
}

func flattenUserKerberosUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserKerberosUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("realm", flattenUserKerberosUser(o["realm"], d, "realm", sv)); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}
	if err = d.Set("shortname", flattenUserKerberosUser(o["shortname"], d, "shortname", sv)); err != nil {
		if !fortiAPIPatch(o["shortname"]) {
			return fmt.Errorf("Error reading shortname: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserKerberosUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserKerberosUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserKerberosUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("realm"); ok {
		t, err := expandUserKerberosUser(d, v, "realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}
	if v, ok := d.GetOk("shortname"); ok {
		t, err := expandUserKerberosUser(d, v, "shortname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shortname"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserKerberosUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
