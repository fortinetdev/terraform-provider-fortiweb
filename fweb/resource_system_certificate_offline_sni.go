// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.offline_sni.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateOfflineSni() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateOfflineSniRead,
		Update: resourceSystemCertificateOfflineSniUpdate,
		Create: resourceSystemCertificateOfflineSniCreate,
		Delete: resourceSystemCertificateOfflineSniDelete,

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

func resourceSystemCertificateOfflineSniCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateOfflineSni: type error")
	}

	obj, err := getObjectSystemCertificateOfflineSni(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOfflineSni resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateOfflineSni(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateOfflineSni resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateOfflineSniRead(d, m)
}

func resourceSystemCertificateOfflineSniUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateOfflineSni(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOfflineSni resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateOfflineSni(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateOfflineSni resource: %v", err)
	}

	return resourceSystemCertificateOfflineSniRead(d, m)
}

func resourceSystemCertificateOfflineSniDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateOfflineSni(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateOfflineSni resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateOfflineSniRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateOfflineSni(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOfflineSni resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateOfflineSni(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateOfflineSni resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateOfflineSni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateOfflineSni(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateOfflineSni(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateOfflineSni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateOfflineSni(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateOfflineSni(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
