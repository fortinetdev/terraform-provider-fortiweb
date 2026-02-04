// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.xml_client_certificate.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateXmlClientCertificate() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateXmlClientCertificateRead,
		Update: resourceSystemCertificateXmlClientCertificateUpdate,
		Create: resourceSystemCertificateXmlClientCertificateCreate,
		Delete: resourceSystemCertificateXmlClientCertificateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"certificatefile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"keyfile": &schema.Schema{
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

func resourceSystemCertificateXmlClientCertificateCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateXmlClientCertificate: type error")
	}

	obj, err := getObjectSystemCertificateXmlClientCertificate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateXmlClientCertificate resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateXmlClientCertificate(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateXmlClientCertificate resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateXmlClientCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateXmlClientCertificateRead(d, m)
}

func resourceSystemCertificateXmlClientCertificateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateXmlClientCertificate(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateXmlClientCertificate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateXmlClientCertificateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateXmlClientCertificate(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateXmlClientCertificate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateXmlClientCertificate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateXmlClientCertificate resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateXmlClientCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateXmlClientCertificate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateXmlClientCertificate(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateXmlClientCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateXmlClientCertificate(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificatefile"); ok {
		t, err := expandSystemCertificateXmlClientCertificate(d, v, "certificatefile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificatefile"] = t
		}
	}
	if v, ok := d.GetOk("keyfile"); ok {
		t, err := expandSystemCertificateXmlClientCertificate(d, v, "keyfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyfile"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateXmlClientCertificate(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
