// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.sign_ca.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateSignCa() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateSignCaRead,
		Update: resourceSystemCertificateSignCaUpdate,
		Create: resourceSystemCertificateSignCaCreate,
		Delete: resourceSystemCertificateSignCaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"upfile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"keyfile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"passwd": &schema.Schema{
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

func resourceSystemCertificateSignCaCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateSignCa: type error")
	}

	obj, err := getObjectSystemCertificateSignCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSignCa resource while getting object: %v", err)
	}

	passwd := ""
	if (*obj)["keyfile"] == nil || (*obj)["keyfile"] == "" {
		passwd = passwd + "?mode=1"
	} else {
		passwd = passwd + "?mode=2"
	}

	if (*obj)["passwd"] != nil && (*obj)["passwd"] != "" {
		passwd = passwd + "&passwd=" + (*obj)["passwd"].(string)
	}

	_, err = c.CreateSystemCertificateSignCa(obj, passwd, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateSignCa resource: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func resourceSystemCertificateSignCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateSignCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSignCa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateSignCa(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateSignCa resource: %v", err)
	}

	return resourceSystemCertificateSignCaRead(d, m)
}

func resourceSystemCertificateSignCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateSignCa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateSignCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateSignCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateSignCa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSignCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateSignCa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateSignCa resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateSignCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateSignCa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("comment", flattenSystemCertificateSignCa(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateSignCa(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateSignCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateSignCa(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemCertificateSignCa(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}
	if v, ok := d.GetOk("upfile"); ok {
		t, err := expandSystemCertificateSignCa(d, v, "upfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upfile"] = t
		}
	}
	if v, ok := d.GetOk("keyfile"); ok {
		t, err := expandSystemCertificateSignCa(d, v, "keyfile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keyfile"] = t
		}
	}
	if v, ok := d.GetOk("passwd"); ok {
		t, err := expandSystemCertificateSignCa(d, v, "passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateSignCa(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
