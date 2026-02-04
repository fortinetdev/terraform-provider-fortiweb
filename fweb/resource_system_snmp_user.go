// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/snmp.user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSnmpUserRead,
		Update: resourceSystemSnmpUserUpdate,
		Create: resourceSystemSnmpUserCreate,
		Delete: resourceSystemSnmpUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//noauthnopriv, authnopriv, authpriv
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"auth_pwd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"priv_pwd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"query_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"query_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trap_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trapport_local": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trapport_remote": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"trapevent": &schema.Schema{
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

func resourceSystemSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpUser: type error")
	}

	obj, err := getObjectSystemSnmpUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUser resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSnmpUserRead(d, m)
}

func resourceSystemSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemSnmpUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUser resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSnmpUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUser resource: %v", err)
	}

	return resourceSystemSnmpUserRead(d, m)
}

func resourceSystemSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemSnmpUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemSnmpUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUser resource from API: %v", err)
	}

	return nil
}

func flattenSystemSnmpUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemSnmpUser(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}
	if err = d.Set("security_level", flattenSystemSnmpUser(o["security-level"], d, "security_level", sv)); err != nil {
		if !fortiAPIPatch(o["security-level"]) {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}
	if err = d.Set("auth_proto", flattenSystemSnmpUser(o["auth-proto"], d, "auth_proto", sv)); err != nil {
		if !fortiAPIPatch(o["auth-proto"]) {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}
	if err = d.Set("auth_pwd", flattenSystemSnmpUser(o["auth-pwd"], d, "auth_pwd", sv)); err != nil {
		if !fortiAPIPatch(o["auth-pwd"]) {
			return fmt.Errorf("Error reading auth_pwd: %v", err)
		}
	}
	if err = d.Set("priv_proto", flattenSystemSnmpUser(o["priv-proto"], d, "priv_proto", sv)); err != nil {
		if !fortiAPIPatch(o["priv-proto"]) {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}
	if err = d.Set("priv_pwd", flattenSystemSnmpUser(o["priv-pwd"], d, "priv_pwd", sv)); err != nil {
		if !fortiAPIPatch(o["priv-pwd"]) {
			return fmt.Errorf("Error reading priv_pwd: %v", err)
		}
	}
	if err = d.Set("query_status", flattenSystemSnmpUser(o["query-status"], d, "query_status", sv)); err != nil {
		if !fortiAPIPatch(o["query-status"]) {
			return fmt.Errorf("Error reading query_status: %v", err)
		}
	}
	if err = d.Set("query_port", flattenSystemSnmpUser(o["query-port"], d, "query_port", sv)); err != nil {
		if !fortiAPIPatch(o["query-port"]) {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}
	if err = d.Set("trap_status", flattenSystemSnmpUser(o["trap-status"], d, "trap_status", sv)); err != nil {
		if !fortiAPIPatch(o["trap-status"]) {
			return fmt.Errorf("Error reading trap_status: %v", err)
		}
	}
	if err = d.Set("trapport_local", flattenSystemSnmpUser(o["trapport-local"], d, "trapport_local", sv)); err != nil {
		if !fortiAPIPatch(o["trapport-local"]) {
			return fmt.Errorf("Error reading trapport_local: %v", err)
		}
	}
	if err = d.Set("trapport_remote", flattenSystemSnmpUser(o["trapport-remote"], d, "trapport_remote", sv)); err != nil {
		if !fortiAPIPatch(o["trapport-remote"]) {
			return fmt.Errorf("Error reading trapport_remote: %v", err)
		}
	}
	if err = d.Set("trapevent", flattenSystemSnmpUser(o["trapevent"], d, "trapevent", sv)); err != nil {
		if !fortiAPIPatch(o["trapevent"]) {
			return fmt.Errorf("Error reading trapevent: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemSnmpUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemSnmpUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpUser(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}
	if v, ok := d.GetOk("security_level"); ok {
		t, err := expandSystemSnmpUser(d, v, "security_level", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}
	if v, ok := d.GetOk("auth_proto"); ok {
		t, err := expandSystemSnmpUser(d, v, "auth_proto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}
	if v, ok := d.GetOk("auth_pwd"); ok {
		t, err := expandSystemSnmpUser(d, v, "auth_pwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}
	if v, ok := d.GetOk("priv_proto"); ok {
		t, err := expandSystemSnmpUser(d, v, "priv_proto", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}
	if v, ok := d.GetOk("priv_pwd"); ok {
		t, err := expandSystemSnmpUser(d, v, "priv_pwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}
	if v, ok := d.GetOk("query_status"); ok {
		t, err := expandSystemSnmpUser(d, v, "query_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-status"] = t
		}
	}
	if v, ok := d.GetOk("query_port"); ok {
		t, err := expandSystemSnmpUser(d, v, "query_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-port"] = t
		}
	}
	if v, ok := d.GetOk("trap_status"); ok {
		t, err := expandSystemSnmpUser(d, v, "trap_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-status"] = t
		}
	}
	if v, ok := d.GetOk("trapport_local"); ok {
		t, err := expandSystemSnmpUser(d, v, "trapport_local", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trapport-local"] = t
		}
	}
	if v, ok := d.GetOk("trapport_remote"); ok {
		t, err := expandSystemSnmpUser(d, v, "trapport_remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trapport-remote"] = t
		}
	}
	if v, ok := d.GetOk("trapevent"); ok {
		t, err := expandSystemSnmpUser(d, v, "trapevent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trapevent"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSnmpUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
