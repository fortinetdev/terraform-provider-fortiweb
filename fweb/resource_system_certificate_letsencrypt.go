// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/certificate.letsencrypt.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemCertificateLetsencrypt() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemCertificateLetsencryptRead,
		Update: resourceSystemCertificateLetsencryptUpdate,
		Create: resourceSystemCertificateLetsencryptCreate,
		Delete: resourceSystemCertificateLetsencryptDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"validation_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"renew_period": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceSystemCertificateLetsencryptCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemCertificateLetsencrypt: type error")
	}

	obj, err := getObjectSystemCertificateLetsencrypt(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLetsencrypt resource while getting object: %v", err)
	}

	_, err = c.CreateSystemCertificateLetsencrypt(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemCertificateLetsencrypt resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemCertificateLetsencryptRead(d, m)
}

func resourceSystemCertificateLetsencryptUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemCertificateLetsencrypt(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLetsencrypt resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCertificateLetsencrypt(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemCertificateLetsencrypt resource: %v", err)
	}

	return resourceSystemCertificateLetsencryptRead(d, m)
}

func resourceSystemCertificateLetsencryptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemCertificateLetsencrypt(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCertificateLetsencrypt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCertificateLetsencryptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemCertificateLetsencrypt(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLetsencrypt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCertificateLetsencrypt(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCertificateLetsencrypt resource from API: %v", err)
	}

	return nil
}

func flattenSystemCertificateLetsencrypt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCertificateLetsencrypt(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("domain", flattenSystemCertificateLetsencrypt(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}
	if err = d.Set("validation_method", flattenSystemCertificateLetsencrypt(o["validation-method"], d, "validation_method", sv)); err != nil {
		if !fortiAPIPatch(o["validation-method"]) {
			return fmt.Errorf("Error reading validation_method: %v", err)
		}
	}
	if err = d.Set("key_type", flattenSystemCertificateLetsencrypt(o["key-type"], d, "key_type", sv)); err != nil {
		if !fortiAPIPatch(o["key-type"]) {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}
	if err = d.Set("renew_period", flattenSystemCertificateLetsencrypt(o["renew-period"], d, "renew_period", sv)); err != nil {
		if !fortiAPIPatch(o["renew-period"]) {
			return fmt.Errorf("Error reading renew_period: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemCertificateLetsencrypt(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemCertificateLetsencrypt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCertificateLetsencrypt(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemCertificateLetsencrypt(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}
	if v, ok := d.GetOk("validation_method"); ok {
		t, err := expandSystemCertificateLetsencrypt(d, v, "validation_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validation-method"] = t
		}
	}
	if v, ok := d.GetOk("key_type"); ok {
		t, err := expandSystemCertificateLetsencrypt(d, v, "key_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}
	if v, ok := d.GetOk("retry_times"); ok {
		t, err := expandSystemCertificateLetsencrypt(d, v, "retry_times", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry-times"] = t
		}
	}
	if v, ok := d.GetOk("renew_period"); ok {
		t, err := expandSystemCertificateLetsencrypt(d, v, "renew_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renew-period"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemCertificateLetsencrypt(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
