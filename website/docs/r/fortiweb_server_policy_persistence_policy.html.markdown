---
subcategory: "FortiWEB Server Policy"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_server_policy_persistence_policy"
description: |-
  Configure FortiWEB server policy persistence policy configuration.
---

# fortiweb_server_policy_persistence_policy
Configure FortiWEB server policy persistence policy configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_persistence_policy" "test" {
	mkey = "test"
	ipv6_mask_length = 128
	ipv4_netmask = "255.255.255.255"
	timeout = 300
	type = "source-ip"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The Name of the server policy persistence policy entry.
* `type` - Type setting. Valid values: source-ip, persistent-cookie, asp-sessionid, php-sessionid, jsp-sessionid, insert-cookie, HTTP-header, url-parameter, rewrite-cookie, embedded-cookie or ssl-session-id.
* `cookie_name` - The value to match or the name of the cookie that FortiWeb inserts.
* `timeout` - The maximum amount of time between requests that FortiWeb maintains persistence, in seconds. Valid values: 10 - 86400.
* `ipv4_netmask` - The IPv4 subnet used for session persistence.
* `ipv6_mask_length` - The IPv6 network prefix used for session persistence.
* `http_header` - The name of the HTTP header that the persistence feature uses to route requests.
* `url_parameter` - The name of the URL parameter that the persistence feature uses to route requests.
* `cookie_path` - The path attribute for the cookie that FortiWeb inserts.
* `cookie_domain` - The domain attribute for the cookie that FortiWeb inserts.
* `secure_cookie` - The secure cookie to force browsers to return the cookie only for HTTPS traffic.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Persistence Policy can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_persistence_policy.labelname {mkey}
```
