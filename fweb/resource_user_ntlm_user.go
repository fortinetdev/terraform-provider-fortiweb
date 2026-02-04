// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/ntlm_user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserNtlmUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserNtlmUserRead,
		Update: resourceUserNtlmUserUpdate,
		Create: resourceUserNtlmUserCreate,
		Delete: resourceUserNtlmUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server": &schema.Schema{
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

func resourceUserNtlmUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserNtlmUser: type error")
	}

	obj, err := getObjectUserNtlmUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserNtlmUser resource while getting object: %v", err)
	}

	_, err = c.CreateUserNtlmUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserNtlmUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserNtlmUserRead(d, m)
}

func resourceUserNtlmUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserNtlmUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserNtlmUser resource while getting object: %v", err)
	}

	_, err = c.UpdateUserNtlmUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserNtlmUser resource: %v", err)
	}

	return resourceUserNtlmUserRead(d, m)
}

func resourceUserNtlmUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserNtlmUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserNtlmUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserNtlmUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserNtlmUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserNtlmUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserNtlmUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserNtlmUser resource from API: %v", err)
	}

	return nil
}

func flattenUserNtlmUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserNtlmUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("port", flattenUserNtlmUser(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}
	if err = d.Set("server", flattenUserNtlmUser(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserNtlmUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserNtlmUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserNtlmUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserNtlmUser(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}
	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserNtlmUser(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserNtlmUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
