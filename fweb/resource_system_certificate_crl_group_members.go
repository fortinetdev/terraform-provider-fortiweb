// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.crl_group/members.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCrlGroupMembers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCrlGroupMembersRead,
		Update: resourceSystemCertificateCrlGroupMembersUpdate,
		Create: resourceSystemCertificateCrlGroupMembersCreate,
		Delete: resourceSystemCertificateCrlGroupMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
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

func resourceSystemCertificateCrlGroupMembersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrlGroupMembers: type error")
	}

	obj, err := getObjectSystemCertificateCrlGroupMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrlGroupMembers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCrlGroupMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrlGroupMembers resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateCrlGroupMembersRead(d, m)
}

func resourceSystemCertificateCrlGroupMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrlGroupMembers: type error")
	}

	obj, err := getObjectSystemCertificateCrlGroupMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrlGroupMembers resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemCertificateCrlGroupMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrlGroupMembers resource: %v", err)
	}

	return resourceSystemCertificateCrlGroupMembersRead(d, m)
}

func resourceSystemCertificateCrlGroupMembersDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrlGroupMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemCertificateCrlGroupMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCrlGroupMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCrlGroupMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrlGroupMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemCertificateCrlGroupMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrlGroupMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCrlGroupMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrlGroupMembers resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateCrlGroupMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCrlGroupMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemCertificateCrlGroupMembers(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("name", flattenSystemCertificateCrlGroupMembers(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateCrlGroupMembers(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCrlGroupMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCrlGroupMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemCertificateCrlGroupMembers(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemCertificateCrlGroupMembers(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCrlGroupMembers(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
