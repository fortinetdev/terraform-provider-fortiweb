---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_ldap_user"
description: |-
  Configure FortiWEB user ldap user configuration.
---

# fortiweb_user_ldap_user
Configure FortiWEB user ldap user configuration.

## Example Usage
```hcl
resource "fortiweb_user_ldap_user" "test" {
	mkey = "test"
	server = "172.23.133.1"
	port = 389
	common_name_id = "commonname"
	distinguished_name = "distinguishedname"
	bind_type = "regular"
	username = "172.23.133.2"
	password = "1234"
	filter = "filter"
	group_authentication = "enable"
	group_type = "windows-ad"
	group_dn = "test_group"
	ssl_connection = "enable"
	protocol = "ldaps"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user ldap user entry.
* `server` - Type the server IP or domain address of the LDAP server.
* `port` - The port number where the LDAP server listens. Valid values: 1-65535.
* `common_name_id` - Enter the identifier, often cn, for the common name (CN) attribute whose value is the user name.
* `distinguished_name` - Enter the distinguished name (DN).
* `bind_type` - Select one of the following LDAP query binding styles. Valid values: anonymous, simple, regular.
* `username` - The bind distinguished-name.
* `password` - The password of the username.
* `filter` - An LDAP query filter string, if any, that will be used to filter out results from the query’s results based upon any attribute in the record set.
* `group_authentication` - Enable to only include users that are members of an LDAP group.
* `group_type` - Select the schema that matches your server’s LDAP directory. Valid values: edirectory, open-ldap, windows-ad.
* `group_dn` - The distinguished name of the LDAP user group.
* `ssl_connection` - Enable to connect to the LDAP servers using an encrypted connection, then select the style of the encryption in protocol.
* `protocol` - Select whether to secure the LDAP query using LDAPS or STARTTLS. Valid values: ldaps, starttls.
* `ca_cert` - Enter the name of the certificate so the FortiWeb will only accept a certificate from the LDAP server that is signed by this CA.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Ldap User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_ldap_user.labelname {mkey}
```
