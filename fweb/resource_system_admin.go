// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/admin.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAdmin() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemAdminRead,
		Update: resourceSystemAdminUpdate,
		Create: resourceSystemAdminCreate,
		Delete: resourceSystemAdminDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"access_profile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trusthostv4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trusthostv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"email_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"phone_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mobile_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hidden": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_global_menu_favorites": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"gui_vdom_menu_favorites": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"column": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"layout_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"permanent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"widget_table": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"admin_usergrp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"wildcard": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"accprofile_override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fortiai": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sshkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"passwd_set_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password_pos": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password0": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password3": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password4": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password5": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password7": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password8": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"history_password9": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"force_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"feature_info_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"old_password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemAdminCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemAdmin: type error")
	}

	obj, err := getObjectSystemAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource while getting object: %v", err)
	}

	_, err = c.CreateSystemAdmin(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemAdmin resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAdmin(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemAdmin resource: %v", err)
	}

	return resourceSystemAdminRead(d, m)
}

func resourceSystemAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemAdmin(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAdminRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemAdmin(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAdmin resource from API: %v", err)
	}

	return nil
}

func flattenSystemAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("access_profile", flattenSystemAdmin(o["access-profile"], d, "access_profile", sv)); err != nil {
		if !fortiAPIPatch(o["access-profile"]) {
			return fmt.Errorf("Error reading access_profile: %v", err)
		}
	}
	if err = d.Set("trusthostv4", flattenSystemAdmin(o["trusthostv4"], d, "trusthostv4", sv)); err != nil {
		if !fortiAPIPatch(o["trusthostv4"]) {
			return fmt.Errorf("Error reading trusthostv4: %v", err)
		}
	}
	if err = d.Set("trusthostv6", flattenSystemAdmin(o["trusthostv6"], d, "trusthostv6", sv)); err != nil {
		if !fortiAPIPatch(o["trusthostv6"]) {
			return fmt.Errorf("Error reading trusthostv6: %v", err)
		}
	}
	if err = d.Set("last_name", flattenSystemAdmin(o["last-name"], d, "last_name", sv)); err != nil {
		if !fortiAPIPatch(o["last-name"]) {
			return fmt.Errorf("Error reading last_name: %v", err)
		}
	}
	if err = d.Set("first_name", flattenSystemAdmin(o["first-name"], d, "first_name", sv)); err != nil {
		if !fortiAPIPatch(o["first-name"]) {
			return fmt.Errorf("Error reading first_name: %v", err)
		}
	}
	if err = d.Set("email_address", flattenSystemAdmin(o["email-address"], d, "email_address", sv)); err != nil {
		if !fortiAPIPatch(o["email-address"]) {
			return fmt.Errorf("Error reading email_address: %v", err)
		}
	}
	if err = d.Set("phone_number", flattenSystemAdmin(o["phone-number"], d, "phone_number", sv)); err != nil {
		if !fortiAPIPatch(o["phone-number"]) {
			return fmt.Errorf("Error reading phone_number: %v", err)
		}
	}
	if err = d.Set("mobile_number", flattenSystemAdmin(o["mobile-number"], d, "mobile_number", sv)); err != nil {
		if !fortiAPIPatch(o["mobile-number"]) {
			return fmt.Errorf("Error reading mobile_number: %v", err)
		}
	}
	if err = d.Set("hidden", flattenSystemAdmin(o["hidden"], d, "hidden", sv)); err != nil {
		if !fortiAPIPatch(o["hidden"]) {
			return fmt.Errorf("Error reading hidden: %v", err)
		}
	}
	if err = d.Set("domains", flattenSystemAdmin(o["domains"], d, "domains", sv)); err != nil {
		if !fortiAPIPatch(o["domains"]) {
			return fmt.Errorf("Error reading domains: %v", err)
		}
	}
	if err = d.Set("gui_global_menu_favorites", flattenSystemAdmin(o["gui-global-menu-favorites"], d, "gui_global_menu_favorites", sv)); err != nil {
		if !fortiAPIPatch(o["gui-global-menu-favorites"]) {
			return fmt.Errorf("Error reading gui_global_menu_favorites: %v", err)
		}
	}
	if err = d.Set("gui_vdom_menu_favorites", flattenSystemAdmin(o["gui-vdom-menu-favorites"], d, "gui_vdom_menu_favorites", sv)); err != nil {
		if !fortiAPIPatch(o["gui-vdom-menu-favorites"]) {
			return fmt.Errorf("Error reading gui_vdom_menu_favorites: %v", err)
		}
	}
	if err = d.Set("column", flattenSystemAdmin(o["column"], d, "column", sv)); err != nil {
		if !fortiAPIPatch(o["column"]) {
			return fmt.Errorf("Error reading column: %v", err)
		}
	}
	if err = d.Set("status", flattenSystemAdmin(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}
	if err = d.Set("name", flattenSystemAdmin(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("layout_type", flattenSystemAdmin(o["layout-type"], d, "layout_type", sv)); err != nil {
		if !fortiAPIPatch(o["layout-type"]) {
			return fmt.Errorf("Error reading layout_type: %v", err)
		}
	}
	if err = d.Set("permanent", flattenSystemAdmin(o["permanent"], d, "permanent", sv)); err != nil {
		if !fortiAPIPatch(o["permanent"]) {
			return fmt.Errorf("Error reading permanent: %v", err)
		}
	}
	if err = d.Set("widget_table", flattenSystemAdmin(o["widget-table"], d, "widget_table", sv)); err != nil {
		if !fortiAPIPatch(o["widget-table"]) {
			return fmt.Errorf("Error reading widget_table: %v", err)
		}
	}
	if err = d.Set("type", flattenSystemAdmin(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("admin_usergrp", flattenSystemAdmin(o["admin-usergrp"], d, "admin_usergrp", sv)); err != nil {
		if !fortiAPIPatch(o["admin-usergrp"]) {
			return fmt.Errorf("Error reading admin_usergrp: %v", err)
		}
	}
	if err = d.Set("password", flattenSystemAdmin(o["password"], d, "password", sv)); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}
	if err = d.Set("wildcard", flattenSystemAdmin(o["wildcard"], d, "wildcard", sv)); err != nil {
		if !fortiAPIPatch(o["wildcard"]) {
			return fmt.Errorf("Error reading wildcard: %v", err)
		}
	}
	if err = d.Set("accprofile_override", flattenSystemAdmin(o["accprofile-override"], d, "accprofile_override", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile-override"]) {
			return fmt.Errorf("Error reading accprofile_override: %v", err)
		}
	}
	if err = d.Set("fortiai", flattenSystemAdmin(o["fortiai"], d, "fortiai", sv)); err != nil {
		if !fortiAPIPatch(o["fortiai"]) {
			return fmt.Errorf("Error reading fortiai: %v", err)
		}
	}
	if err = d.Set("sshkey", flattenSystemAdmin(o["sshkey"], d, "sshkey", sv)); err != nil {
		if !fortiAPIPatch(o["sshkey"]) {
			return fmt.Errorf("Error reading sshkey: %v", err)
		}
	}
	if err = d.Set("passwd_set_time", flattenSystemAdmin(o["passwd-set-time"], d, "passwd_set_time", sv)); err != nil {
		if !fortiAPIPatch(o["passwd-set-time"]) {
			return fmt.Errorf("Error reading passwd_set_time: %v", err)
		}
	}
	if err = d.Set("history_password_pos", flattenSystemAdmin(o["history-password-pos"], d, "history_password_pos", sv)); err != nil {
		if !fortiAPIPatch(o["history-password-pos"]) {
			return fmt.Errorf("Error reading history_password_pos: %v", err)
		}
	}
	if err = d.Set("history_password0", flattenSystemAdmin(o["history-password0"], d, "history_password0", sv)); err != nil {
		if !fortiAPIPatch(o["history-password0"]) {
			return fmt.Errorf("Error reading history_password0: %v", err)
		}
	}
	if err = d.Set("history_password1", flattenSystemAdmin(o["history-password1"], d, "history_password1", sv)); err != nil {
		if !fortiAPIPatch(o["history-password1"]) {
			return fmt.Errorf("Error reading history_password1: %v", err)
		}
	}
	if err = d.Set("history_password2", flattenSystemAdmin(o["history-password2"], d, "history_password2", sv)); err != nil {
		if !fortiAPIPatch(o["history-password2"]) {
			return fmt.Errorf("Error reading history_password2: %v", err)
		}
	}
	if err = d.Set("history_password3", flattenSystemAdmin(o["history-password3"], d, "history_password3", sv)); err != nil {
		if !fortiAPIPatch(o["history-password3"]) {
			return fmt.Errorf("Error reading history_password3: %v", err)
		}
	}
	if err = d.Set("history_password4", flattenSystemAdmin(o["history-password4"], d, "history_password4", sv)); err != nil {
		if !fortiAPIPatch(o["history-password4"]) {
			return fmt.Errorf("Error reading history_password4: %v", err)
		}
	}
	if err = d.Set("history_password5", flattenSystemAdmin(o["history-password5"], d, "history_password5", sv)); err != nil {
		if !fortiAPIPatch(o["history-password5"]) {
			return fmt.Errorf("Error reading history_password5: %v", err)
		}
	}
	if err = d.Set("history_password6", flattenSystemAdmin(o["history-password6"], d, "history_password6", sv)); err != nil {
		if !fortiAPIPatch(o["history-password6"]) {
			return fmt.Errorf("Error reading history_password6: %v", err)
		}
	}
	if err = d.Set("history_password7", flattenSystemAdmin(o["history-password7"], d, "history_password7", sv)); err != nil {
		if !fortiAPIPatch(o["history-password7"]) {
			return fmt.Errorf("Error reading history_password7: %v", err)
		}
	}
	if err = d.Set("history_password8", flattenSystemAdmin(o["history-password8"], d, "history_password8", sv)); err != nil {
		if !fortiAPIPatch(o["history-password8"]) {
			return fmt.Errorf("Error reading history_password8: %v", err)
		}
	}
	if err = d.Set("history_password9", flattenSystemAdmin(o["history-password9"], d, "history_password9", sv)); err != nil {
		if !fortiAPIPatch(o["history-password9"]) {
			return fmt.Errorf("Error reading history_password9: %v", err)
		}
	}
	if err = d.Set("force_password_change", flattenSystemAdmin(o["force-password-change"], d, "force_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["force-password-change"]) {
			return fmt.Errorf("Error reading force_password_change: %v", err)
		}
	}
	if err = d.Set("feature_info_ver", flattenSystemAdmin(o["feature-info-ver"], d, "feature_info_ver", sv)); err != nil {
		if !fortiAPIPatch(o["feature-info-ver"]) {
			return fmt.Errorf("Error reading feature_info_ver: %v", err)
		}
	}
	if err = d.Set("old_password", flattenSystemAdmin(o["old-password"], d, "old_password", sv)); err != nil {
		if !fortiAPIPatch(o["old-password"]) {
			return fmt.Errorf("Error reading old_password: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemAdmin(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("access_profile"); ok {
		t, err := expandSystemAdmin(d, v, "access_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-profile"] = t
		}
	}
	if v, ok := d.GetOk("trusthostv4"); ok {
		t, err := expandSystemAdmin(d, v, "trusthostv4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthostv4"] = t
		}
	}
	if v, ok := d.GetOk("trusthostv6"); ok {
		t, err := expandSystemAdmin(d, v, "trusthostv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusthostv6"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemAdmin(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("force_password_change"); ok {
		t, err := expandSystemAdmin(d, v, "force_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-password-change"] = t
		}
	}
	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemAdmin(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemAdmin(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
