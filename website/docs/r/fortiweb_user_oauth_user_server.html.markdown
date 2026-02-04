---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_oauth_user_server"
description: |-
  Configure FortiWEB user oauth user server configuration.
---

# fortiweb_user_oauth_user_server
Configure FortiWEB user oauth user server configuration.

## Example Usage
```hcl
resource "fortiweb_user_oauth_user_server" "test" {
	mkey = "test"
	mode = "both"
	scope = "openid"
	oidc = "enable"
	client_id = "172.23.133.1"
	client_secret = "1234"
	redirect_endpoint = "123"
	authz_req = "Google Authorization"
	token_req = "Google Token"
	refresh_req = "Google Refresh"
	jwks_req = "Google JWK Set"
	validate_req = "Google Validate"
	userinfo_req = "Google Userinfo"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user oauth user server entry.
* `mode` - Select whether FortiWeb works as an authorization client or a resource server, or both.
* `scope` - The scope field for OAuth.
* `oidc` - Enable to use OIDC authentication.
* `client_id` - A client credential.
* `client_secret` - A client credential.
* `redirect_endpoint` - Redirection URL back to FortiWeb.
* `authz_req` - The authorization request created in config user oauth-user request.
* `token_req` - The token request created in config user oauth-user request.
* `refresh_req` - The refresh request created in config user oauth-user request.
* `jwks_req` - The JWKS request created in config user oauth-user request.
* `validate_req` - The valid request created in config user oauth-user request.
* `userinfo_req` - The user info request created in config user oauth-user request.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Oauth User Server can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_oauth_user_server.labelname {mkey}
```
