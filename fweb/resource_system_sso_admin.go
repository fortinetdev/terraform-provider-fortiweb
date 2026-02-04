// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/sso_admin.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSsoAdmin() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSsoAdminRead,
		Update: resourceSystemSsoAdminUpdate,
		Create: resourceSystemSsoAdminCreate,
		Delete: resourceSystemSsoAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"access_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemSsoAdminCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSsoAdmin: type error")
	}

	obj, err := getObjectSystemSsoAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoAdmin resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSsoAdmin(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoAdmin resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSsoAdminRead(d, m)
}

func resourceSystemSsoAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemSsoAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoAdmin resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSsoAdmin(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoAdmin resource: %v", err)
	}

	return resourceSystemSsoAdminRead(d, m)
}

func resourceSystemSsoAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemSsoAdmin(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoAdminRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemSsoAdmin(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoAdmin resource from API: %v", err)
	}

	return nil
}

func flattenSystemSsoAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSsoAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("access_profile", flattenSystemSsoAdmin(o["access-profile"], d, "access_profile", sv)); err != nil {
		if !fortiAPIPatch(o["access-profile"]) {
			return fmt.Errorf("Error reading access_profile: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemSsoAdmin(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemSsoAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSsoAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_profile"); ok {
		t, err := expandSystemSsoAdmin(d, v, "access_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-profile"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSsoAdmin(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
