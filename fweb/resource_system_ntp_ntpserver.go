// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/ntp/netserver.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemNtpNtpServer() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemNtpNtpServerRead,
		Update: resourceSystemNtpNtpServerUpdate,
		Create: resourceSystemNtpNtpServerCreate,
		Delete: resourceSystemNtpNtpServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemNtpNtpServerCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemNtpNtpServer: type error")
	}

	obj, err := getObjectSystemNtpNtpServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemNtpNtpServer resource while getting object: %v", err)
	}

	_, err = c.CreateSystemNtpNtpServer(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemNtpNtpServer resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemNtpNtpServerRead(d, m)
}

func resourceSystemNtpNtpServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing SystemNtpNtpServer: type error")
	}

	obj, err := getObjectSystemNtpNtpServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtpNtpServer resource while getting object: %v", err)
	}

	mkey = "?sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemNtpNtpServer(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtpNtpServer resource: %v", err)
	}

	return resourceSystemNtpNtpServerRead(d, m)
}

func resourceSystemNtpNtpServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "2"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing SystemNtpNtpServer: type error")
	}

	mkey = "?sub_mkey=" + sub_mkey
	err := c.DeleteSystemNtpNtpServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNtpNtpServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpNtpServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "2"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing SystemNtpNtpServer: type error")
	}

	mkey = "?sub_mkey=" + sub_mkey
	o, err := c.ReadSystemNtpNtpServer(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtpNtpServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNtpNtpServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtpNtpServer resource from API: %v", err)
	}

	return nil
}

func flattenSystemNtpNtpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNtpNtpServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemNtpNtpServer(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("key_id", flattenSystemNtpNtpServer(o["key-id"], d, "key_id", sv)); err != nil {
		if !fortiAPIPatch(o["key-id"]) {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}
	if err = d.Set("key_type", flattenSystemNtpNtpServer(o["key-type"], d, "key_type", sv)); err != nil {
		if !fortiAPIPatch(o["key-type"]) {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}
	if err = d.Set("key", flattenSystemNtpNtpServer(o["key"], d, "key", sv)); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}
	if err = d.Set("authentication", flattenSystemNtpNtpServer(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}
	if err = d.Set("ip_type", flattenSystemNtpNtpServer(o["ip-type"], d, "ip_type", sv)); err != nil {
		if !fortiAPIPatch(o["ip-type"]) {
			return fmt.Errorf("Error reading ip_type: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemNtpNtpServer(o["server"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemNtpNtpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtpNtpServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("key_id"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "key_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-id"] = t
		}
	}
	if v, ok := d.GetOk("key_type"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "key_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-type"] = t
		}
	}
	if v, ok := d.GetOk("key"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}
	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}
	if v, ok := d.GetOk("ip_type"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "ip_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-type"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemNtpNtpServer(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	return &obj, nil
}
