// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.local.import_certificate

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLocalImportCertificate() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLocalImportCertificateRead,
		Update: resourceSystemCertificateLocalImportCertificateUpdate,
		Create: resourceSystemCertificateLocalImportCertificateCreate,
		Delete: resourceSystemCertificateLocalImportCertificateDelete,

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
			"certificatewithkeyfile": &schema.Schema{
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

func resourceSystemCertificateLocalImportCertificateCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLocalImportCertificate: type error")
	}

	obj, err := getObjectSystemCertificateLocalImportCertificate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalImportCertificate resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateLocalImportCertificate(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLocalImportCertificate resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateLocalImportCertificateUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceSystemCertificateLocalImportCertificateRead(d, m)
}

func resourceSystemCertificateLocalImportCertificateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateLocalImportCertificate(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLocalImportCertificate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateLocalImportCertificateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateLocalImportCertificate(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalImportCertificate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLocalImportCertificate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLocalImportCertificate resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateLocalImportCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLocalImportCertificate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenSystemCertificateLocalImportCertificate(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLocalImportCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLocalImportCertificate(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("certificatefile"); ok {
		t, err := expandSystemCertificateLocalImportCertificate(d, v, "certificatefile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificateFile"] = t
			obj["type"] = "localCertificate"
		}
	}
	if v, ok := d.GetOk("keyfile"); ok {
		t, err := expandSystemCertificateLocalImportCertificate(d, v, "keyfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyFile"] = t
			obj["type"] = "certificate"
			obj["hsm"] = "undefined"
		}
	}
	if v, ok := d.GetOk("certificatewithkeyfile"); ok {
		t, err := expandSystemCertificateLocalImportCertificate(d, v, "certificatewithkeyfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificateWithKeyFile"] = t
			obj["type"] = "PKCS12Certificate"
		}
	}
	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemCertificateLocalImportCertificate(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLocalImportCertificate(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
