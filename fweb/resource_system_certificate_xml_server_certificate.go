// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.xml_server_certificate.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateXmlServerCertificate() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateXmlServerCertificateRead,
		Update: resourceSystemCertificateXmlServerCertificateUpdate,
		Create: resourceSystemCertificateXmlServerCertificateCreate,
		Delete: resourceSystemCertificateXmlServerCertificateDelete,

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
			"password": &schema.Schema{
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

func resourceSystemCertificateXmlServerCertificateCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateXmlServerCertificate: type error")
	}

	obj, err := getObjectSystemCertificateXmlServerCertificate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateXmlServerCertificate resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateXmlServerCertificate(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateXmlServerCertificate resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateXmlServerCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateXmlServerCertificateRead(d, m)
}

func resourceSystemCertificateXmlServerCertificateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateXmlServerCertificate(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateXmlServerCertificate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateXmlServerCertificateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateXmlServerCertificate(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateXmlServerCertificate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateXmlServerCertificate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateXmlServerCertificate resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateXmlServerCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateXmlServerCertificate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateXmlServerCertificate(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateXmlServerCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateXmlServerCertificate(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificatefile"); ok {
		t, err := expandSystemCertificateXmlServerCertificate(d, v, "certificatefile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificatefile"] = t
		}
	}
	if v, ok := d.GetOk("keyfile"); ok {
		t, err := expandSystemCertificateXmlServerCertificate(d, v, "keyfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyfile"] = t
		}
	}
	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemCertificateXmlServerCertificate(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateXmlServerCertificate(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
