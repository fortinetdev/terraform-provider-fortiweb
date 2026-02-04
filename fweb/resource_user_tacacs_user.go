// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/Tacacs_user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserTacacsUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserTacacsUserRead,
		Update: resourceUserTacacsUserUpdate,
		Create: resourceUserTacacsUserCreate,
		Delete: resourceUserTacacsUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"passwordedited": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
				Optional: true,
			},
			"secret": &schema.Schema{
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

func resourceUserTacacsUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserTacacsUser: type error")
	}

	obj, err := getObjectUserTacacsUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserTacacsUser resource while getting object: %v", err)
	}

	_, err = c.CreateUserTacacsUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserTacacsUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserTacacsUserRead(d, m)
}

func resourceUserTacacsUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserTacacsUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserTacacsUser resource while getting object: %v", err)
	}

	_, err = c.UpdateUserTacacsUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserTacacsUser resource: %v", err)
	}

	return resourceUserTacacsUserRead(d, m)
}

func resourceUserTacacsUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserTacacsUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserTacacsUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserTacacsUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserTacacsUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserTacacsUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserTacacsUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserTacacsUser resource from API: %v", err)
	}

	return nil
}

func flattenUserTacacsUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserTacacsUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("auth_type", flattenUserTacacsUser(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}
	if err = d.Set("type", flattenUserTacacsUser(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("server", flattenUserTacacsUser(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}
	if err = d.Set("passwordedited", flattenUserTacacsUser(o["passwordEdited"], d, "passwordedited", sv)); err != nil {
		if !fortiAPIPatch(o["passwordEdited"]) {
			return fmt.Errorf("Error reading passwordedited: %v", err)
		}
	}
	if err = d.Set("secret", flattenUserTacacsUser(o["secret"], d, "secret", sv)); err != nil {
		if !fortiAPIPatch(o["secret"]) {
			return fmt.Errorf("Error reading secret: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserTacacsUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserTacacsUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserTacacsUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandUserTacacsUser(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserTacacsUser(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserTacacsUser(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}
	if v, ok := d.GetOk("passwordedited"); ok {
		t, err := expandUserTacacsUser(d, v, "passwordedited", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwordEdited"] = t
		}
	}
	if v, ok := d.GetOk("secret"); ok {
		t, err := expandUserTacacsUser(d, v, "secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserTacacsUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
