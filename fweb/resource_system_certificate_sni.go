// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.sni.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateSni() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateSniRead,
		Update: resourceSystemCertificateSniUpdate,
		Create: resourceSystemCertificateSniCreate,
		Delete: resourceSystemCertificateSniDelete,

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

func resourceSystemCertificateSniCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateSni: type error")
	}

	obj, err := getObjectSystemCertificateSni(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSni resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateSni(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSni resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateSniRead(d, m)
}

func resourceSystemCertificateSniUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateSni(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSni resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateSni(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSni resource: %v", err)
	}

	return resourceSystemCertificateSniRead(d, m)
}

func resourceSystemCertificateSniDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateSni(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateSni resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateSniRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateSni(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSni resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateSni(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSni resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateSni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateSni(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateSni(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateSni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateSni(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateSni(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
