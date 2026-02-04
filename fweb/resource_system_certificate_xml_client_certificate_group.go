// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.xml_client_certificate_group.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateXmlClientCertificateGroup() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateXmlClientCertificateGroupRead,
		Update: resourceSystemCertificateXmlClientCertificateGroupUpdate,
		Create: resourceSystemCertificateXmlClientCertificateGroupCreate,
		Delete: resourceSystemCertificateXmlClientCertificateGroupDelete,

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

func resourceSystemCertificateXmlClientCertificateGroupCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateXmlClientCertificateGroup: type error")
	}

	obj, err := getObjectSystemCertificateXmlClientCertificateGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateXmlClientCertificateGroup resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateXmlClientCertificateGroup(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateXmlClientCertificateGroup resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateXmlClientCertificateGroupRead(d, m)
}

func resourceSystemCertificateXmlClientCertificateGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateXmlClientCertificateGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateXmlClientCertificateGroup resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateXmlClientCertificateGroup(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateXmlClientCertificateGroup resource: %v", err)
	}

	return resourceSystemCertificateXmlClientCertificateGroupRead(d, m)
}

func resourceSystemCertificateXmlClientCertificateGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateXmlClientCertificateGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateXmlClientCertificateGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateXmlClientCertificateGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateXmlClientCertificateGroup(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateXmlClientCertificateGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateXmlClientCertificateGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateXmlClientCertificateGroup resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateXmlClientCertificateGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateXmlClientCertificateGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateXmlClientCertificateGroup(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateXmlClientCertificateGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateXmlClientCertificateGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateXmlClientCertificateGroup(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
