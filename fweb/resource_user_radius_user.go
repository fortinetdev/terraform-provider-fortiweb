// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/radius_user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserRadiusUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserRadiusUserRead,
		Update: resourceUserRadiusUserUpdate,
		Create: resourceUserRadiusUserCreate,
		Delete: resourceUserRadiusUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secret": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary_server_port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secondary_secret": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password1edited": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password2edited": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fac_push": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"require_msg_auth": &schema.Schema{
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

func resourceUserRadiusUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserRadiusUser: type error")
	}

	obj, err := getObjectUserRadiusUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserRadiusUser resource while getting object: %v", err)
	}

	_, err = c.CreateUserRadiusUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserRadiusUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserRadiusUserRead(d, m)
}

func resourceUserRadiusUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserRadiusUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserRadiusUser resource while getting object: %v", err)
	}

	_, err = c.UpdateUserRadiusUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserRadiusUser resource: %v", err)
	}

	return resourceUserRadiusUserRead(d, m)
}

func resourceUserRadiusUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserRadiusUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserRadiusUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserRadiusUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserRadiusUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserRadiusUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserRadiusUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserRadiusUser resource from API: %v", err)
	}

	return nil
}

func flattenUserRadiusUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserRadiusUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("server", flattenUserRadiusUser(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}
	if err = d.Set("server_port", flattenUserRadiusUser(o["server-port"], d, "server_port", sv)); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}
	if err = d.Set("secret", flattenUserRadiusUser(o["secret"], d, "secret", sv)); err != nil {
		if !fortiAPIPatch(o["secret"]) {
			return fmt.Errorf("Error reading secret: %v", err)
		}
	}
	if err = d.Set("secondary_server", flattenUserRadiusUser(o["secondary-server"], d, "secondary_server", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}
	if err = d.Set("secondary_server_port", flattenUserRadiusUser(o["secondary-server-port"], d, "secondary_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server-port"]) {
			return fmt.Errorf("Error reading secondary_server_port: %v", err)
		}
	}
	if err = d.Set("secondary_secret", flattenUserRadiusUser(o["secondary-secret"], d, "secondary_secret", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-secret"]) {
			return fmt.Errorf("Error reading secondary_secret: %v", err)
		}
	}
	if err = d.Set("password1edited", flattenUserRadiusUser(o["password1Edited"], d, "password1edited", sv)); err != nil {
		if !fortiAPIPatch(o["password1Edited"]) {
			return fmt.Errorf("Error reading password1edited: %v", err)
		}
	}
	if err = d.Set("password2edited", flattenUserRadiusUser(o["password2Edited"], d, "password2edited", sv)); err != nil {
		if !fortiAPIPatch(o["password2Edited"]) {
			return fmt.Errorf("Error reading password2edited: %v", err)
		}
	}
	if err = d.Set("nas_ip", flattenUserRadiusUser(o["nas-ip"], d, "nas_ip", sv)); err != nil {
		if !fortiAPIPatch(o["nas-ip"]) {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}
	if err = d.Set("auth_type", flattenUserRadiusUser(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}
	if err = d.Set("fac_push", flattenUserRadiusUser(o["fac-push"], d, "fac_push", sv)); err != nil {
		if !fortiAPIPatch(o["fac-push"]) {
			return fmt.Errorf("Error reading fac_push: %v", err)
		}
	}
	if err = d.Set("require_msg_auth", flattenUserRadiusUser(o["require-msg-auth"], d, "require_msg_auth", sv)); err != nil {
		if !fortiAPIPatch(o["require-msg-auth"]) {
			return fmt.Errorf("Error reading require_msg_auth: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserRadiusUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserRadiusUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserRadiusUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserRadiusUser(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}
	if v, ok := d.GetOk("server_port"); ok {
		t, err := expandUserRadiusUser(d, v, "server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}
	if v, ok := d.GetOk("secret"); ok {
		t, err := expandUserRadiusUser(d, v, "secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}
	if v, ok := d.GetOk("secondary_server"); ok {
		t, err := expandUserRadiusUser(d, v, "secondary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}
	if v, ok := d.GetOk("secondary_server_port"); ok {
		t, err := expandUserRadiusUser(d, v, "secondary_server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server-port"] = t
		}
	}
	if v, ok := d.GetOk("secondary_secret"); ok {
		t, err := expandUserRadiusUser(d, v, "secondary_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}
	if v, ok := d.GetOk("password1edited"); ok {
		t, err := expandUserRadiusUser(d, v, "password1edited", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password1Edited"] = t
		}
	}
	if v, ok := d.GetOk("password2edited"); ok {
		t, err := expandUserRadiusUser(d, v, "password2edited", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2Edited"] = t
		}
	}
	if v, ok := d.GetOk("nas_ip"); ok {
		t, err := expandUserRadiusUser(d, v, "nas_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}
	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandUserRadiusUser(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}
	if v, ok := d.GetOk("fac_push"); ok {
		t, err := expandUserRadiusUser(d, v, "fac_push", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fac-push"] = t
		}
	}
	if v, ok := d.GetOk("require_msg_auth"); ok {
		t, err := expandUserRadiusUser(d, v, "require_msg_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["require-msg-auth"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserRadiusUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
