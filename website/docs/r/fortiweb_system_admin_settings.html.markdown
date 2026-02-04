---
subcategory: "FortiWEB System Settings"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_admin_settings"
description: |-
  Configure FortiWEB system admin settings configuration.
---

# fortiweb_system_admin
Configure FortiWEB system admin settings configuration.

## Example Usage
```hcl
resource "fortiweb_system_admin_settings" "test" {
	hostname = "FortiWeb"
	http = 80
	https = 443
	httpsservercertificate = "defaulthttpscert"
	admin_tls_v12 = "enable"
	admin_tls_v13 = "disable"
	confsync_enable = "enable"
	configsync = 995
	idletimeout = 350
	language = "english"
	passwordpolicy {
			min_length_option = "enable"
			min_length = 8
			single_admin_mode = "enable"
			character_requirements = "enable"
			min_upper_case_letter = 5
			min_lower_case_letter = 1
			min_number = 5
			min_non_alphanumeric = 0
			forbid_password_reuse = "enable"
			history_password_number = 4
			expire_status = "enable"
			expire_day = 80
			status = "enable"
	}
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `hostname` - The host name of this FortiWeb appliance.
* `http` - Enter the port number on which the FortiWeb appliance listens for HTTP access to the web UI. Valid values: 1-65535.
* `https` - Enter the port number on which the FortiWeb appliance listens for HTTPS (SSL-secured) access to the web UI.  Valid values: 1-65535.
* `httpsservercertificate` - Specifies the certificate that FortiWeb uses for the accesses to its Web UI through HTTPS.
* `httpsIntermediateCertificate` - Specifies the intermediate CA group if any.
* `admin_tls_v12` - Enable to specify TSL 1.2 clients can use to connect securely to the FortiWeb appliance. Valid values: enable, disable.
* `admin_tls_v13` - Enable to specify TSL 1.3 clients can use to connect securely to the FortiWeb appliance. Valid values: enable, disable.
* `confsync_enable` - Enable to be able to enter configsync local port via the CLI.
* `configsync` - Enter the port number the local FortiWeb uses to listen for a remote (peer) FortiWeb. Valid values: 1-65535.
* `idletimeout` - Enter the amount of time after which an idle administrative session with the web UI or CLI will be automatically logged out.  Valid values: 1-480.
* `language` - Select which language to use when displaying the web UI. Valid values: english, japanese, simch, trach.
* `passwordpolicy` - Access-Profile setting.
* `status` - Enable to enforce password rules for administrator accounts. Valid values: enable, disable.
* `min_length_option` - Enable to set the minimum length for the password. Valid values: enable, disable.
* `min_length` - Enter the minimum password length. Valid values: 8-128.
* `single_admin_mode` - Enable to activate single admin user login. Valid values: enable, disable.
* `character_requirements` - Enable to set characters. Valid values: enable, disable.
* `min_upper_case_letter` - Enter the number of upper case characters. Valid values: 0-128.
* `min_lower_case_letter` - Enter the number of lower case characters. Valid values: 0-128.
* `min_number` - Enter the number of number characters. Valid values: 0-128.
* `min_non_alphanumeric` - Enter the number of special characters. Valid values: 0-128.
* `forbid_password_reuse` - Enable forbidding password re-use. Valid values: enable, disable.
* `history_password_number` - Enter the number of history passwords that can not be reused. Valid values: 1-10.
* `expire_status` - Enable password expiration. Valid values: enable, disable.
* `expire_day` - Enter the valid period for the password. Valid values: 1-999.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Admin can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_admin_settings.labelname {mkey}
```
