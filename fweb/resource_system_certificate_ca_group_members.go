// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.ca_group/members.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCaGroupMembers() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCaGroupMembersRead,
		Update: resourceSystemCertificateCaGroupMembersUpdate,
		Create: resourceSystemCertificateCaGroupMembersCreate,
		Delete: resourceSystemCertificateCaGroupMembersDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tsl": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"publish_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

func resourceSystemCertificateCaGroupMembersCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCaGroupMembers: type error")
	}

	obj, err := getObjectSystemCertificateCaGroupMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCaGroupMembers resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCaGroupMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCaGroupMembers resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateCaGroupMembersRead(d, m)
}

func resourceSystemCertificateCaGroupMembersUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCaGroupMembers: type error")
	}

	obj, err := getObjectSystemCertificateCaGroupMembers(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCaGroupMembers resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemCertificateCaGroupMembers(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCaGroupMembers resource: %v", err)
	}

	return resourceSystemCertificateCaGroupMembersRead(d, m)
}

func resourceSystemCertificateCaGroupMembersDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCaGroupMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemCertificateCaGroupMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCaGroupMembers resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCaGroupMembersRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCaGroupMembers: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemCertificateCaGroupMembers(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCaGroupMembers resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCaGroupMembers(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCaGroupMembers resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateCaGroupMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCaGroupMembers(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemCertificateCaGroupMembers(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("type", flattenSystemCertificateCaGroupMembers(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("name", flattenSystemCertificateCaGroupMembers(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("tsl", flattenSystemCertificateCaGroupMembers(o["tsl"], d, "tsl", sv)); err != nil {
		if !fortiAPIPatch(o["tsl"]) {
			return fmt.Errorf("Error reading tsl: %v", err)
		}
	}
	if err = d.Set("publish_dn", flattenSystemCertificateCaGroupMembers(o["publish-dn"], d, "publish_dn", sv)); err != nil {
		if !fortiAPIPatch(o["publish-dn"]) {
			return fmt.Errorf("Error reading publish_dn: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateCaGroupMembers(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCaGroupMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCaGroupMembers(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemCertificateCaGroupMembers(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemCertificateCaGroupMembers(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemCertificateCaGroupMembers(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("tsl"); ok {
		t, err := expandSystemCertificateCaGroupMembers(d, v, "tsl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tsl"] = t
		}
	}
	if v, ok := d.GetOk("publish_dn"); ok {
		t, err := expandSystemCertificateCaGroupMembers(d, v, "publish_dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["publish-dn"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCaGroupMembers(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
