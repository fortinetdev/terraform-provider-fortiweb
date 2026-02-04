// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/recaptcha_user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserRecaptchaUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserRecaptchaUserRead,
		Update: resourceUserRecaptchaUserUpdate,
		Create: resourceUserRecaptchaUserCreate,
		Delete: resourceUserRecaptchaUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"site_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secret_key": &schema.Schema{
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

func resourceUserRecaptchaUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserRecaptchaUser: type error")
	}

	obj, err := getObjectUserRecaptchaUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserRecaptchaUser resource while getting object: %v", err)
	}

	_, err = c.CreateUserRecaptchaUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserRecaptchaUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserRecaptchaUserRead(d, m)
}

func resourceUserRecaptchaUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserRecaptchaUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserRecaptchaUser resource while getting object: %v", err)
	}

	_, err = c.UpdateUserRecaptchaUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserRecaptchaUser resource: %v", err)
	}

	return resourceUserRecaptchaUserRead(d, m)
}

func resourceUserRecaptchaUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserRecaptchaUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserRecaptchaUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserRecaptchaUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserRecaptchaUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserRecaptchaUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserRecaptchaUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserRecaptchaUser resource from API: %v", err)
	}

	return nil
}

func flattenUserRecaptchaUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserRecaptchaUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenUserRecaptchaUser(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("site_key", flattenUserRecaptchaUser(o["site-key"], d, "site_key", sv)); err != nil {
		if !fortiAPIPatch(o["site-key"]) {
			return fmt.Errorf("Error reading site_key: %v", err)
		}
	}
	if err = d.Set("secret_key", flattenUserRecaptchaUser(o["secret-key"], d, "secret_key", sv)); err != nil {
		if !fortiAPIPatch(o["secret-key"]) {
			return fmt.Errorf("Error reading secret_key: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserRecaptchaUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserRecaptchaUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserRecaptchaUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserRecaptchaUser(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("site_key"); ok {
		t, err := expandUserRecaptchaUser(d, v, "site_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["site-key"] = t
		}
	}
	if v, ok := d.GetOk("secret_key"); ok {
		t, err := expandUserRecaptchaUser(d, v, "secret_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-key"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserRecaptchaUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
