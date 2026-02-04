---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_access_url_access_rule_match_condition"
description: |-
  Configure FortiWEB waf url access url access rule match condition configuration.
---

# fortiweb_waf_url_access_url_access_rule_match_condition
Configure FortiWEB waf url access url access rule match condition configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_access_url_access_rule_match_condition" "test" {
	sub_mkey = "1"
	mkey = "test"
	sip_address_check = "enable"
	sip_address_type = "sdomain"
	sdomain_type = "ipv4"
	sip_address_domain = "test"
	type = "simple-string"
	reg_exp = "/test"
	url_access_parameter = "test"
	only_method_check = "enable"
	only_method = "get post"
	only_protocol_check = "enable"
	only_protocol = "http https"
	reverse_match = "no"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url access url access rule entry.
* `sub_mkey` - The id of the waf url access url access rule match condition entry.
* `sip_address_check` - Enable to add the client’s source IP address as a criteria for matching the URL access rule. Valid values: enable, disable.
* `sip_address_type` - Valid values: sip, sdomain or source-domain.
* `sip_address_domain` - Specifies the domain to match the client source IP after DNS lookup.
* `type` - Select how to use the text in reg-exp to determine whether or not a request URL meets the conditions for this rule. Valid values: simple-string, regular-expression.
* `reg_exp` - Depending on your selection in type and reverse-match, type a regular expression that defines either all matching or all non-matching URLs. 
* `url_access_parameter` - Enter the URL Access Parameter rule you have created by config waf url-access-parameter.
* `only_method_check` - Enable or disable the HTTP method.
* `only_method` - Select the HTTP methods. Valid values: get, post, head, options, trace, connect, delete, put, patch, webdav, rpc, others.
* `only_protocol_check` - Enable or disable the HTTP protocol.
* `only_protocol` - Select the HTTP protocols. Valid values: http, https, ws, wss.
* `reverse_match` - Indicate how to use reg-exp when determining whether or not this rule’s condition has been met. Valid values: no, yes.
* `reverse_dns_timeout` - Reverse-Dns-Timeout setting.
* `sip_address_value` - Sip-Address-Value setting.
* `sdomain_type` -  Specifies the type of IP address FortiWeb retrieves from the DNS lookup of the domain specified by sip-address-domain.
* `source_domain_type` - Available only if sip-address-type is source-domain.
* `source_domain` - Enter a literal domain or a regular expression that is designed to match multiple URLs.
* `reverse_dns_timeout` - To avoid the process hanging for a long time, you can set this option to limit the time when FortiWeb performs the reverse DNS lookup for an IP address. Valid values: 0-600.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Access Url Access Rule Match Condition can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_access_url_access_rule_match_condition.labelname {mkey}
```
