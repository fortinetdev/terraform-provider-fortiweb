// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.intermediate_certificate_group.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateIntermediateCertificateGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateIntermediateCertificateGroupRead,
		Update: resourceSystemCertificateIntermediateCertificateGroupUpdate,
		Create: resourceSystemCertificateIntermediateCertificateGroupCreate,
		Delete: resourceSystemCertificateIntermediateCertificateGroupDelete,

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

func resourceSystemCertificateIntermediateCertificateGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateIntermediateCertificateGroup: type error")
	}

	obj, err := getObjectSystemCertificateIntermediateCertificateGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCertificateGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateIntermediateCertificateGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateIntermediateCertificateGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateIntermediateCertificateGroupRead(d, m)
}

func resourceSystemCertificateIntermediateCertificateGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateIntermediateCertificateGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCertificateGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateIntermediateCertificateGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateIntermediateCertificateGroup resource: %v", err)
	}

	return resourceSystemCertificateIntermediateCertificateGroupRead(d, m)
}

func resourceSystemCertificateIntermediateCertificateGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateIntermediateCertificateGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateIntermediateCertificateGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateIntermediateCertificateGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateIntermediateCertificateGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCertificateGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateIntermediateCertificateGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateIntermediateCertificateGroup resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateIntermediateCertificateGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateIntermediateCertificateGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateIntermediateCertificateGroup(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateIntermediateCertificateGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateIntermediateCertificateGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateIntermediateCertificateGroup(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
