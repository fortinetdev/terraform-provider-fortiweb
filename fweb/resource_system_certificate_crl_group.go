// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.crl_group.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateCrlGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateCrlGroupRead,
		Update: resourceSystemCertificateCrlGroupUpdate,
		Create: resourceSystemCertificateCrlGroupCreate,
		Delete: resourceSystemCertificateCrlGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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

func resourceSystemCertificateCrlGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateCrlGroup: type error")
	}

	obj, err := getObjectSystemCertificateCrlGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrlGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateCrlGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateCrlGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateCrlGroupRead(d, m)
}

func resourceSystemCertificateCrlGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateCrlGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrlGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateCrlGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateCrlGroup resource: %v", err)
	}

	return resourceSystemCertificateCrlGroupRead(d, m)
}

func resourceSystemCertificateCrlGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateCrlGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateCrlGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateCrlGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateCrlGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrlGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateCrlGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateCrlGroup resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateCrlGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateCrlGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateCrlGroup(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateCrlGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateCrlGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateCrlGroup(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
