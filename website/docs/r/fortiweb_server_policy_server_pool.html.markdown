---
subcategory: "FortiWEB Server Policy"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_server_policy_server_pool"
description: |-
  Configure FortiWEB server policy server pool configuration.
---

# fortiweb_server_policy_server_pool
Configure FortiWEB server policy server pool configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_server_pool" "test" {
	mkey = "sp3"
	type = "reverse-proxy"
	server_balance = "disable"
	comment = "test"
	health = "HLTHCK_TCP"
	lb_algo = "round-robin"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The Name of the server policy server pool entry.
* `type` - Type setting. Default is reverse-proxy.
* `protocol` - Specific options for a server pool become available. Default is HTTP.
* `server_balance` - Specifies whether the pool contains a single server or multiple members. Valid values: enable or disable.
* `health` - The name of a server health check FortiWeb uses to determine the responsiveness of server pool members.
* `hlck_sip` - The IP address of Hlck-Sip setting.
* `hlck_sip6` - The IPv6 address of Hlck-Sip6 setting.
* `lb_algo` - Select the load-balancing algorithms that FortiWeb uses when it distributes new connections among server pool members.
* `persistence` - Enter the name of the persistence policy that specifies a session persistence method and timeout to apply to the pool.
* `comment` - Enter a description or other comment.
* `http_reuse` - Configure multiplexing so that FortiWeb uses a single connection to a server for requests from multiple clients. Valid values: aggressive, always, never, safe.
* `reuse_conn_total_time` - Enter the maximum time limit in which a cached server connection may be reused. Valid values: 1-1000.
* `reuse_conn_idle_time` - Enter an idle time limit for a cached server connection. Valid values: 1-1000.
* `reuse_conn_max_request` - Enter the maximum number of HTTP responses that the cached server connection may handle. Valid values: 1-1000.
* `reuse_conn_max_count` - Enter the maximum number of allowed cached server connections. Valid values: 1-1000.
* `reuse_conn_max_count` - Enter the maximum number of allowed cached server connections. Valid values: 1-1000.
* `adfs_server_name` - Enter a name for the AD FS Server.
* `proxy_protocol` - If the back-end server enables proxy protocol, you need to enable the Proxy Protocol option on FortiWeb so that the TCP SSL and HTTP traffic can successfully go through. Valid values: enable, disable.
* `proxy_protocol_version` - Select the proxy protocol version for the back-end server. Valid values: v1, v2.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Server Pool can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_server_pool.labelname {mkey}
```
