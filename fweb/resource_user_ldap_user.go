// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure user/ldap_user.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserLdapUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserLdapUserRead,
		Update: resourceUserLdapUserUpdate,
		Create: resourceUserLdapUserCreate,
		Delete: resourceUserLdapUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"common_name_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distinguished_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"bind_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ssl_connection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"concatenation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"config_change_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"attribute": &schema.Schema{
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

func resourceUserLdapUserCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing UserLdapUser: type error")
	}

	obj, err := getObjectUserLdapUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserLdapUser resource while getting object: %v", err)
	}

	_, err = c.CreateUserLdapUser(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating UserLdapUser resource: %v", err)
	}

	d.SetId(mkey)

	return resourceUserLdapUserRead(d, m)
}

func resourceUserLdapUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectUserLdapUser(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserLdapUser resource while getting object: %v", err)
	}

	_, err = c.UpdateUserLdapUser(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating UserLdapUser resource: %v", err)
	}

	return resourceUserLdapUserRead(d, m)
}

func resourceUserLdapUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteUserLdapUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting UserLdapUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserLdapUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadUserLdapUser(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading UserLdapUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserLdapUser(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserLdapUser resource from API: %v", err)
	}

	return nil
}

func flattenUserLdapUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserLdapUser(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("server", flattenUserLdapUser(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}
	if err = d.Set("port", flattenUserLdapUser(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}
	if err = d.Set("common_name_id", flattenUserLdapUser(o["common-name-id"], d, "common_name_id", sv)); err != nil {
		if !fortiAPIPatch(o["common-name-id"]) {
			return fmt.Errorf("Error reading common_name_id: %v", err)
		}
	}
	if err = d.Set("distinguished_name", flattenUserLdapUser(o["distinguished-name"], d, "distinguished_name", sv)); err != nil {
		if !fortiAPIPatch(o["distinguished-name"]) {
			return fmt.Errorf("Error reading distinguished_name: %v", err)
		}
	}
	if err = d.Set("bind_type", flattenUserLdapUser(o["bind-type"], d, "bind_type", sv)); err != nil {
		if !fortiAPIPatch(o["bind-type"]) {
			return fmt.Errorf("Error reading bind_type: %v", err)
		}
	}
	if err = d.Set("ssl_connection", flattenUserLdapUser(o["ssl-connection"], d, "ssl_connection", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-connection"]) {
			return fmt.Errorf("Error reading ssl_connection: %v", err)
		}
	}
	if err = d.Set("protocol", flattenUserLdapUser(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}
	if err = d.Set("ca_cert", flattenUserLdapUser(o["ca-cert"], d, "ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}
	if err = d.Set("username", flattenUserLdapUser(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}
	if err = d.Set("password", flattenUserLdapUser(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}
	if err = d.Set("filter", flattenUserLdapUser(o["filter"], d, "filter", sv)); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}
	if err = d.Set("group_authentication", flattenUserLdapUser(o["group_authentication"], d, "group_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["group_authentication"]) {
			return fmt.Errorf("Error reading group_authentication: %v", err)
		}
	}
	if err = d.Set("group_type", flattenUserLdapUser(o["group_type"], d, "group_type", sv)); err != nil {
		if !fortiAPIPatch(o["group_type"]) {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}
	if err = d.Set("group_dn", flattenUserLdapUser(o["group_dn"], d, "group_dn", sv)); err != nil {
		if !fortiAPIPatch(o["group_dn"]) {
			return fmt.Errorf("Error reading group_dn: %v", err)
		}
	}
	if err = d.Set("concatenation", flattenUserLdapUser(o["concatenation"], d, "concatenation", sv)); err != nil {
		if !fortiAPIPatch(o["concatenation"]) {
			return fmt.Errorf("Error reading concatenation: %v", err)
		}
	}
	if err = d.Set("config_change_time", flattenUserLdapUser(o["config-change-time"], d, "config_change_time", sv)); err != nil {
		if !fortiAPIPatch(o["config-change-time"]) {
			return fmt.Errorf("Error reading config_change_time: %v", err)
		}
	}
	if err = d.Set("name", flattenUserLdapUser(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("attribute", flattenUserLdapUser(o["attribute"], d, "attribute", sv)); err != nil {
		if !fortiAPIPatch(o["attribute"]) {
			return fmt.Errorf("Error reading attribute: %v", err)
		}
	}
	if err = d.Set("mkey", flattenUserLdapUser(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandUserLdapUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserLdapUser(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserLdapUser(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}
	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserLdapUser(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}
	if v, ok := d.GetOk("common_name_id"); ok {
		t, err := expandUserLdapUser(d, v, "common_name_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["common-name-id"] = t
		}
	}
	if v, ok := d.GetOk("distinguished_name"); ok {
		t, err := expandUserLdapUser(d, v, "distinguished_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distinguished-name"] = t
		}
	}
	if v, ok := d.GetOk("bind_type"); ok {
		t, err := expandUserLdapUser(d, v, "bind_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bind-type"] = t
		}
	}
	if v, ok := d.GetOk("ssl_connection"); ok {
		t, err := expandUserLdapUser(d, v, "ssl_connection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-connection"] = t
		}
	}
	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandUserLdapUser(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}
	if v, ok := d.GetOk("ca_cert"); ok {
		t, err := expandUserLdapUser(d, v, "ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-cert"] = t
		}
	}
	if v, ok := d.GetOk("username"); ok {
		t, err := expandUserLdapUser(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}
	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserLdapUser(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}
	if v, ok := d.GetOk("filter"); ok {
		t, err := expandUserLdapUser(d, v, "filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}
	if v, ok := d.GetOk("group_authentication"); ok {
		t, err := expandUserLdapUser(d, v, "group_authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group_authentication"] = t
		}
	}
	if v, ok := d.GetOk("group_type"); ok {
		t, err := expandUserLdapUser(d, v, "group_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group_type"] = t
		}
	}
	if v, ok := d.GetOk("group_dn"); ok {
		t, err := expandUserLdapUser(d, v, "group_dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group_dn"] = t
		}
	}
	if v, ok := d.GetOk("concatenation"); ok {
		t, err := expandUserLdapUser(d, v, "concatenation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["concatenation"] = t
		}
	}
	if v, ok := d.GetOk("config_change_time"); ok {
		t, err := expandUserLdapUser(d, v, "config_change_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["config-change-time"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserLdapUser(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("attribute"); ok {
		t, err := expandUserLdapUser(d, v, "attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["attribute"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandUserLdapUser(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
