package forticlient

func (c *FortiSDKClient) CreateRouterStatic(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/router/static"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterStatic(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/router/static"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterStatic(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/router/static"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterStatic(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/router/static"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/dns"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemDns(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/dns"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)

	return
}

func (c *FortiSDKClient) CreateServerPolicyServerPool(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/server-pool"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicyServerPool(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/server-pool"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicyServerPool(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/server-pool"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicyServerPool(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/server-pool"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateServerPolicyPersistencePolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/persistence-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicyPersistencePolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/persistence-policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicyPersistencePolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/persistence-policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicyPersistencePolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/persistence-policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserRecaptchaUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/recaptcha-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserRecaptchaUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/recaptcha-user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserRecaptchaUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/recaptcha-user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserRecaptchaUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/recaptcha-user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserTacacsUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/tacacs+-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserTacacsUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/tacacs+-user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserTacacsUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/tacacs+-user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserTacacsUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/tacacs+-user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserKerberosUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/kerberos-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserKerberosUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/kerberos-user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserKerberosUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/kerberos-user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserKerberosUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/kerberos-user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserKerberosUserServerMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/kerberos-user/server-members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserKerberosUserServerMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/kerberos-user/server-members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserKerberosUserServerMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/kerberos-user/server-members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserKerberosUserServerMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/kerberos-user/server-members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserNtlmUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/ntlm-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserNtlmUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/ntlm-user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserNtlmUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/ntlm-user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserNtlmUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/ntlm-user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserRadiusUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/radius-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserRadiusUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/radius-user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserRadiusUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/radius-user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserRadiusUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/radius-user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserLdapUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/ldap-user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserLdapUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/ldap-user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserLdapUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/ldap-user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserLdapUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/ldap-user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserLdapUserExtractedAttributes(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/ldap-user/extracted-attributes"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserLdapUserExtractedAttributes(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/ldap-user/extracted-attributes"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserLdapUserExtractedAttributes(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/ldap-user/extracted-attributes"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserLdapUserExtractedAttributes(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/ldap-user/extracted-attributes"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserAdminUsergrp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/admin-usergrp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserAdminUsergrp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/admin-usergrp"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserAdminUsergrp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/admin-usergrp"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserAdminUsergrp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/admin-usergrp"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserAdminUsergrpMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/admin-usergrp/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserAdminUsergrpMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/admin-usergrp/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserAdminUsergrpMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/admin-usergrp/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserAdminUsergrpMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/admin-usergrp/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateUserOauthUserServer(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/user/oauth-user.server"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteUserOauthUserServer(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/user/oauth-user.server"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateUserOauthUserServer(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/user/oauth-user.server"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadUserOauthUserServer(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/user/oauth-user.server"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlAccessUrlAccessPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlAccessUrlAccessPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlAccessUrlAccessPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlAccessUrlAccessPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlAccessUrlAccessPolicyRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy/rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlAccessUrlAccessPolicyRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy/rule"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlAccessUrlAccessPolicyRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy/rule"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlAccessUrlAccessPolicyRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-policy/rule"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlAccessUrlAccessRule(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlAccessUrlAccessRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlAccessUrlAccessRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlAccessUrlAccessRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlAccessUrlAccessRuleMatchCondition(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule/match-condition"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlAccessUrlAccessRuleMatchCondition(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule/match-condition"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlAccessUrlAccessRuleMatchCondition(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule/match-condition"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlAccessUrlAccessRuleMatchCondition(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-rule/match-condition"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlAccessUrlAccessParameter(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlAccessUrlAccessParameter(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlAccessUrlAccessParameter(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlAccessUrlAccessParameter(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlAccessUrlAccessParameterUrlAccessParameterList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter/url-access-parameter-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlAccessUrlAccessParameterUrlAccessParameterList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter/url-access-parameter-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlAccessUrlAccessParameterUrlAccessParameterList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter/url-access-parameter-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlAccessUrlAccessParameterUrlAccessParameterList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-access.url-access-parameter/url-access-parameter-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAdmin(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAdmin(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/admin"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAdmin(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/admin"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAdmin(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/admin"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSsoAdmin(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/sso-admin"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemSsoAdmin(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/sso-admin"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemSsoAdmin(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/sso-admin"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSsoAdmin(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/sso-admin"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAccprofile(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/accprofile"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAccprofile(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/accprofile"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAccprofile(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/accprofile"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAccprofile(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/accprofile"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemAdminSettings(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/admin.settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemAdminSettings(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/system/admin.settings"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemAdminSettings(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/system/admin.settings"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemAdminSettings(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/admin.settings"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafAllowMethodExceptions(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafAllowMethodExceptions(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafAllowMethodExceptions(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafAllowMethodExceptions(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafAllowMethodExceptionsAllowMethodExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions/allow-method-exception-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafAllowMethodExceptionsAllowMethodExceptionList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions/allow-method-exception-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafAllowMethodExceptionsAllowMethodExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions/allow-method-exception-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafAllowMethodExceptionsAllowMethodExceptionList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/allow-method-exceptions/allow-method-exception-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafAllowMethodPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/allow-method-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafAllowMethodPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/allow-method-policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafAllowMethodPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/allow-method-policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafAllowMethodPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/allow-method-policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafAllowedOrigins(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/allowed-origins"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafAllowedOrigins(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/allowed-origins"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafAllowedOrigins(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/allowed-origins"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafAllowedOrigins(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/allowed-origins"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafAllowedOriginsOriginList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/allowed-origins/origin-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafAllowedOriginsOriginList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/allowed-origins/origin-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafAllowedOriginsOriginList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/allowed-origins/origin-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafAllowedOriginsOriginList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/allowed-origins/origin-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCorsProtectionRule(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCorsProtectionRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCorsProtectionRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCorsProtectionRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCorsProtectionRuleAllowedHeadersList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-headers-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCorsProtectionRuleAllowedHeadersList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-headers-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCorsProtectionRuleAllowedHeadersList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-headers-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCorsProtectionRuleAllowedHeadersList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-headers-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCorsProtectionRuleAllowedMethodsList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-methods-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCorsProtectionRuleAllowedMethodsList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-methods-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCorsProtectionRuleAllowedMethodsList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-methods-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCorsProtectionRuleAllowedMethodsList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/allowed-methods-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCorsProtectionRuleExposedHeadersList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/exposed-headers-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCorsProtectionRuleExposedHeadersList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/exposed-headers-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCorsProtectionRuleExposedHeadersList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/exposed-headers-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCorsProtectionRuleExposedHeadersList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/cors-protection-rule/exposed-headers-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCa(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.ca"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCa(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.ca"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCa(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCa(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.ca"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateTslCa(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.tslca"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateTslCa(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.tsl-ca"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateTslCa(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateTslCa(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.tslca"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCaGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.ca-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCaGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.ca-group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCaGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.ca-group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCaGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.ca-group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCaGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.ca-group/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCaGroupMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.ca-group/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCaGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.ca-group/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCaGroupMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.ca-group/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateSignCa(params *map[string]interface{}, passwd string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.sign" + passwd
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateSignCa(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.sign-ca"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateSignCa(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.sign-ca"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateSignCa(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.sign"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateIntermediateCertificate(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.intermediateca"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateIntermediateCertificate(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateIntermediateCertificate(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateIntermediateCertificate(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.intermediateca"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateIntermediateCertificateGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateIntermediateCertificateGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateIntermediateCertificateGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateIntermediateCertificateGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateIntermediateCertificateGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateIntermediateCertificateGroupMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateIntermediateCertificateGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateIntermediateCertificateGroupMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.intermediate-certificate-group/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCrl(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.crl"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCrl(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.crl"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCrl(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCrl(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.crl"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCrlGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.crl-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCrlGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.crl-group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCrlGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.crl-group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCrlGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.crl-group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateCrlGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.crl-group/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateCrlGroupMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.crl-group/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateCrlGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.crl-group/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateCrlGroupMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.crl-group/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateOcspStapling(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-stapling"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateOcspStapling(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-stapling"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateOcspStapling(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-stapling"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateOcspStapling(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-stapling"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateOcspSigningCerts(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.ocsp_sign_certs"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateOcspSigningCerts(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-signing-certs"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateOcspSigningCerts(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateOcspSigningCerts(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.ocsp_sign_certs"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateOcspResponder(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-responder"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateOcspResponder(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-responder"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateOcspResponder(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-responder"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateOcspResponder(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.ocsp-responder"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateVerify(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.verify"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateVerify(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.verify"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateVerify(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.verify"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateVerify(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.verify"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateServerCertificateVerify(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.server-certificate-verify"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateServerCertificateVerify(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.server-certificate-verify"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateServerCertificateVerify(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.server-certificate-verify"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateServerCertificateVerify(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.server-certificate-verify"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateSni(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.sni"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateSni(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.sni"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateSni(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.sni"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateSni(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.sni"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateSniMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.sni/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateSniMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.sni/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateSniMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.sni/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateSniMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.sni/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateOfflineSni(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateOfflineSni(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateOfflineSni(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateOfflineSni(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateOfflineSniMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateOfflineSniMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateOfflineSniMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateOfflineSniMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.offline-sni/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateUrlcert(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.urlcert"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateUrlcert(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.urlcert"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateUrlcert(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.urlcert"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateUrlcert(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.urlcert"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateUrlcertList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.urlcert/list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateUrlcertList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.urlcert/list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateUrlcertList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.urlcert/list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateUrlcertList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.urlcert/list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLetsencrypt(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLetsencrypt(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLetsencrypt(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLetsencrypt(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateXmlClientCertificate(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.xmlclientcertificate"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateXmlClientCertificate(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateXmlClientCertificate(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateXmlClientCertificate(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.xmlclientcertificate"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateXmlServerCertificate(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.xmlservercertificate"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateXmlServerCertificate(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.xml-server-certificate"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateXmlServerCertificate(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateXmlServerCertificate(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.xmlservercertificate"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateXmlClientCertificateGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateXmlClientCertificateGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateXmlClientCertificateGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateXmlClientCertificateGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateXmlClientCertificateGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group/members"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateXmlClientCertificateGroupMembers(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group/members"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateXmlClientCertificateGroupMembers(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group/members"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateXmlClientCertificateGroupMembers(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.xml-client-certificate-group/members"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLocalGenerateCsr(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.local.generate_csr"
	output = make(map[string]interface{})

	err = createUpload2(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLocalGenerateCsr(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.local"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLocalGenerateCsr(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLocalGenerateCsr(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.local"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLocalImportCertificate(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/system/certificate.local.import_certificate"
	output = make(map[string]interface{})

	err = createUpload(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLocalImportCertificate(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.local"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLocalImportCertificate(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLocalImportCertificate(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/certificate.local"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateMultiLocal(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.multi-local"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateMultiLocal(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.multi-local"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateMultiLocal(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.multi-local"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateMultiLocal(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.multi-local"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateHpkp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.hpkp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateHpkp(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.hpkp"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateHpkp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.hpkp"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateHpkp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.hpkp"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemCertificateLetsencryptSanList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt/san-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemCertificateLetsencryptSanList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt/san-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemCertificateLetsencryptSanList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt/san-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemCertificateLetsencryptSanList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/certificate.letsencrypt/san-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSnmpSysinfo(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) DeleteSystemSnmpSysinfo(mkey string, vdom string) (err error) {
	return
}

func (c *FortiSDKClient) UpdateSystemSnmpSysinfo(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/snmp.sysinfo"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSnmpSysinfo(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/snmp.sysinfo"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSnmpCommunity(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/snmp.community"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemSnmpCommunity(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/snmp.community"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemSnmpCommunity(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/snmp.community"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSnmpCommunity(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/snmp.community"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSnmpCommunityHosts(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/snmp.community/hosts"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemSnmpCommunityHosts(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/snmp.community/hosts"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemSnmpCommunityHosts(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/snmp.community/hosts"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSnmpCommunityHosts(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/snmp.community/hosts"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSnmpUser(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/snmp.user"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemSnmpUser(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/snmp.user"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemSnmpUser(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/snmp.user"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSnmpUser(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/snmp.user"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemSnmpUserHosts(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/snmp.user/hosts"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemSnmpUserHosts(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/snmp.user/hosts"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemSnmpUserHosts(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/snmp.user/hosts"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemSnmpUserHosts(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/snmp.user/hosts"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemHa(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) DeleteSystemHa(mkey string, vdom string) (err error) {
	return
}

func (c *FortiSDKClient) UpdateSystemHa(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/ha"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemHa(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/ha"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemNtpNtpServer(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/ntp/ntpserver"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemNtpNtpServer(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/ntp/ntpserver"
	path += mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemNtpNtpServer(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/ntp/ntpserver"
	path += mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemNtpNtpServer(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/ntp/ntpserver"
	path += mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemNtp(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) DeleteSystemNtp(mkey string, vdom string) (err error) {
	return
}

func (c *FortiSDKClient) UpdateSystemNtp(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/ntp"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemNtp(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/ntp"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemMaintenanceSystemTime(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) DeleteSystemMaintenanceSystemTime(mkey string, vdom string) (err error) {
	return
}

func (c *FortiSDKClient) UpdateSystemMaintenanceSystemTime(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/system/maintenance.systemtime"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemMaintenanceSystemTime(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/system/maintenance.systemtime"

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafFileUploadRestrictionPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafFileUploadRestrictionPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafFileUploadRestrictionPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafFileUploadRestrictionPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafFileUploadRestrictionPolicyRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy/rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafFileUploadRestrictionPolicyRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy/rule"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafFileUploadRestrictionPolicyRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy/rule"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafFileUploadRestrictionPolicyRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-policy/rule"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafFileUploadRestrictionRule(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafFileUploadRestrictionRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-rule"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafFileUploadRestrictionRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafFileUploadRestrictionRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-rule"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) ReadWafFileSecurityFileTypes(mkey string, vdom string) (resultsArray []interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/waf/filesecurity.filetypes"

	resultsArray, err = read2(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafFileUploadRestrictionRuleFileTypes(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/waf/filesecurity.filetypes"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafFileUploadRestrictionRuleFileTypes(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-rule/file-types"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafFileUploadRestrictionRuleFileTypes(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	return
}

func (c *FortiSDKClient) ReadWafFileUploadRestrictionRuleFileTypes(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/file-upload-restriction-rule/file-types"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCustomProtectionGroup(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/custom-protection-group"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCustomProtectionGroup(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/custom-protection-group"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCustomProtectionGroup(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/custom-protection-group"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCustomProtectionGroup(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/custom-protection-group"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCustomProtectionRule(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCustomProtectionRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCustomProtectionRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCustomProtectionRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCustomProtectionRuleMeetCondition(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule/meet-condition"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCustomProtectionRuleMeetCondition(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule/meet-condition"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCustomProtectionRuleMeetCondition(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule/meet-condition"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCustomProtectionRuleMeetCondition(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/custom-protection-rule/meet-condition"
	path += "?mkey=" + mkey
	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafCustomProtectionGroupTypeList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/custom-protection-group/type-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafCustomProtectionGroupTypeList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/custom-protection-group/type-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafCustomProtectionGroupTypeList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/custom-protection-group/type-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafCustomProtectionGroupTypeList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/custom-protection-group/type-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafSignature(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/waf/signatures"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafSignature(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/signature"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafSignature(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/waf/signatures"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafSignature(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/waf/signatures"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateServerPolicyServerPoolPserverList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/server-pool/pserver-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicyServerPoolPserverList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/server-pool/pserver-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicyServerPoolPserverList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/server-pool/pserver-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicyServerPoolPserverList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/server-pool/pserver-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateServerPolicyServiceCustom(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/service.custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicyServiceCustom(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/service.custom"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicyServiceCustom(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/service.custom"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicyServiceCustom(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/service.custom"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateServerPolicySslCiphersCustom(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/ssl-ciphers.custom"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicySslCiphersCustom(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/ssl-ciphers.custom"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicySslCiphersCustom(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/ssl-ciphers.custom"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicySslCiphersCustom(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/ssl-ciphers.custom"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateServerPolicyVserver(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/vserver"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicyVserver(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/vserver"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicyVserver(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/vserver"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicyVserver(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/vserver"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateServerPolicyVserverVipList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/server-policy/vserver/vip-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteServerPolicyVserverVipList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/server-policy/vserver/vip-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateServerPolicyVserverVipList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/server-policy/vserver/vip-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadServerPolicyVserverVipList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/server-policy/vserver/vip-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateSystemVip(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/system/vip"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteSystemVip(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/system/vip"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateSystemVip(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/system/vip"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadSystemVip(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/system/vip"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateRouterPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/router/policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteRouterPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/router/policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateRouterPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/router/policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadRouterPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/router/policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpHeaderSecurity(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-header-security"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpHeaderSecurity(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-header-security"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpHeaderSecurity(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-header-security"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpHeaderSecurity(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-header-security"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpHeaderSecurityException(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpHeaderSecurityException(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpHeaderSecurityException(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpHeaderSecurityException(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpHeaderSecurityExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception/list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpHeaderSecurityExceptionList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception/list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpHeaderSecurityExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception/list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpHeaderSecurityExceptionList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-header-security-exception/list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpHeaderSecurityHttpHeaderSecurityList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-header-security/http-header-security-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpHeaderSecurityHttpHeaderSecurityList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-header-security/http-header-security-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpHeaderSecurityHttpHeaderSecurityList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-header-security/http-header-security-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpHeaderSecurityHttpHeaderSecurityList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-header-security/http-header-security-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpConstraintsExceptions(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpConstraintsExceptions(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpConstraintsExceptions(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpConstraintsExceptions(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpConstraintsExceptionsHttpConstraintsExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions/http_constraints-exception-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpConstraintsExceptionsHttpConstraintsExceptionList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions/http_constraints-exception-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpConstraintsExceptionsHttpConstraintsExceptionList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions/http_constraints-exception-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpConstraintsExceptionsHttpConstraintsExceptionList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-constraints-exceptions/http_constraints-exception-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafHttpProtocolParameterRestriction(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/http-protocol-parameter-restriction"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafHttpProtocolParameterRestriction(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/http-protocol-parameter-restriction"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafHttpProtocolParameterRestriction(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/http-protocol-parameter-restriction"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafHttpProtocolParameterRestriction(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/http-protocol-parameter-restriction"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlEncryptionUrlEncryptionRule(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlEncryptionUrlEncryptionRule(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlEncryptionUrlEncryptionRule(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlEncryptionUrlEncryptionRule(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlEncryptionUrlEncryptionRuleUrlList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/url-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlEncryptionUrlEncryptionRuleUrlList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/url-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlEncryptionUrlEncryptionRuleUrlList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/url-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlEncryptionUrlEncryptionRuleUrlList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/url-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlEncryptionUrlEncryptionRuleExceptions(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/exceptions"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlEncryptionUrlEncryptionRuleExceptions(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/exceptions"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlEncryptionUrlEncryptionRuleExceptions(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/exceptions"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlEncryptionUrlEncryptionRuleExceptions(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-rule/exceptions"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlEncryptionUrlEncryptionPolicy(params *map[string]interface{}, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy"
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlEncryptionUrlEncryptionPolicy(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy"
	path += "?mkey=" + escapeURLString(mkey)

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlEncryptionUrlEncryptionPolicy(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlEncryptionUrlEncryptionPolicy(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy"
	path += "?mkey=" + escapeURLString(mkey)

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}

func (c *FortiSDKClient) CreateWafUrlEncryptionUrlEncryptionPolicyRuleList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "POST"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy/rule-list"
	path += "?mkey=" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) DeleteWafUrlEncryptionUrlEncryptionPolicyRuleList(mkey string, vdom string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy/rule-list"
	path += "?mkey=" + mkey

	err = delete(c, HTTPMethod, path, vdom)
	return
}

func (c *FortiSDKClient) UpdateWafUrlEncryptionUrlEncryptionPolicyRuleList(params *map[string]interface{}, mkey string, vdom string) (output map[string]interface{}, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy/rule-list"
	path += "?mkey=" + mkey
	output = make(map[string]interface{})

	err = createUpdate(c, HTTPMethod, path, params, output, vdom)
	return
}

func (c *FortiSDKClient) ReadWafUrlEncryptionUrlEncryptionPolicyRuleList(mkey string, vdom string) (mapTmp map[string]interface{}, err error) {
	HTTPMethod := "GET"
	path := "/api/v2.0/cmdb/waf/url-encryption.url-encryption-policy/rule-list"
	path += "?mkey=" + mkey

	mapTmp, err = read(c, HTTPMethod, path, false, vdom)
	return
}
