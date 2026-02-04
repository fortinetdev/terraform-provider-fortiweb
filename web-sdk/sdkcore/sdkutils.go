package forticlient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	//"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var LOGING_URL string = "/logincheck"
var LOGOUT_URL string = "/logout"
var GET_CSRF_TOKEN_URL string = "/api/v2.0/system/state"

func createHTTPClient() *http.Client {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: tr}
}

func updateClientAuth(c *FortiSDKClient) error { //update token and cookie
	token, cookies, _, csrfToken, err := Login(c.Config.Auth.Hostname, c.Config.Auth.Username, c.Config.Auth.Password, "443")
	if err != nil {
		return fmt.Errorf("failed to Login request: %v", err)
	}
	c.Config.Auth.Token = token
	c.Config.Auth.Cookies = cookies
	c.Config.Auth.CSRFToken = csrfToken
	return nil
}

func configureRequestHeaders(req *http.Request, c *FortiSDKClient, http_method string) { //configure request headers
	for _, cookie := range c.Config.Auth.Cookies {
		req.AddCookie(cookie)
	}

	// Add headers for POST/PUT/DELETE requests
	if http_method == http.MethodPost || http_method == http.MethodPut || http_method == http.MethodDelete {
		req.Header.Set("X-CSRFTOKEN", c.Config.Auth.CSRFToken)
	} //else GET request
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
}

func configureUrl(c *FortiSDKClient, api string, vdomparam string) string {
	url := "https://" + c.Config.FwTarget + api
	if vdomparam != "" {
		url = url + "?vdom=" + vdomparam
	}

	fmt.Println("url: ", url)
	return url
}

func processResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	buffer, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		if len(buffer) != 0 {
			return buffer, nil
		}
		return nil, fmt.Errorf("empty response body")
	}
	return nil, fmt.Errorf("failed to read response body: %v", err)
}

func parseJSONResponse(buffer []byte) (map[string]interface{}, error) {
	var tmp_obj interface{}
	err := json.Unmarshal(buffer, &tmp_obj)

	if err == nil {
		if result, ok := tmp_obj.(map[string]interface{}); ok {
			return result, nil
		}
		return nil, fmt.Errorf("Error converting response to map")
	}
	return nil, fmt.Errorf("Error parsing JSON: %v, response: %s", err, string(buffer))
}

func handleResponse(resp *http.Response) (bool, string, error) {
	fmt.Println("resp.StatusCode = ", resp.StatusCode)

	buffer, err := processResponse(resp)
	if resp.StatusCode == 200 {
		return true, string(buffer), nil
	}

	if err == nil {
		tmp_obj, err := parseJSONResponse(buffer)
		if err == nil {
			output_str, _ := json.MarshalIndent(tmp_obj, "", "\t")
			return false, string(output_str), fmt.Errorf("response status code: %v, response: %s", resp.StatusCode, output_str)
		}
	}
	return false, "", err
}

func make_call(c *FortiSDKClient, api string, http_method string, body interface{}, vdomparam string) (bool, string, error) {
	client := createHTTPClient()
	var body_string []byte

	login_err := updateClientAuth(c)
	if login_err != nil {
		return false, "", login_err
	}

	url := configureUrl(c, api, vdomparam)

	if http_method == http.MethodPost || http_method == http.MethodPut {
		dataMap := map[string]interface{}{"data": body}
		body_string, _ = json.Marshal(dataMap)
	} else {
		body_string, _ = json.Marshal(body)
	}

	if strings.Contains(url, "maintenance.systemtime") { //special case for api SystemMaintenanceSystemTime
		body_string, _ = json.Marshal(body)
	}

	if strings.Contains(url, "filesecurity.filetypes") && http_method == http.MethodPost { //special case for api WafFileUploadRestrictionRuleFileTypes
		dataMap := map[string]interface{}{"data": []interface{}{body}}
		body_string, _ = json.Marshal(dataMap)
	}

	req_body := strings.NewReader(string(body_string))
	req, err := http.NewRequest(http_method, url, req_body)
	if err != nil {
		return false, "", fmt.Errorf("failed to create request: %v", err)
	}

	configureRequestHeaders(req, c, http_method)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err1 := client.Do(req)

	if err1 != nil {
		return false, "", fmt.Errorf("http request failed: %v", err1)
	}

	return handleResponse(resp)
}

func make_call2(c *FortiSDKClient, api string, http_method string, body interface{}, vdomparam string) (bool, string, error) {
	client := createHTTPClient()

	bodyPtr := body.(*map[string]interface{})
	bodyMap := *bodyPtr

	login_err := updateClientAuth(c)
	if login_err != nil {
		return false, "", login_err
	}

	url := configureUrl(c, api, vdomparam)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for key, value := range bodyMap {
		if value != nil && value != "" {
			switch key {
			case "name", "certType", "password", "hsm", "type", "httpUrl", "scepUrl", "url":
				fw, err := bodyWriter.CreateFormField(key)
				_, err = io.Copy(fw, strings.NewReader(value.(string)))
				if err != nil {
					return false, "", fmt.Errorf("Error form created %s name: %v\n", key, err)
				}
			case "srcfile", "upfile", "keyfile", "keyFile", "uploadedFile", "certificatefile", "certificateFile", "certificateWithKeyFile":
				fw, err := bodyWriter.CreateFormFile(key, filepath.Base(value.(string)))
				fh, err := os.Open(value.(string))
				_, err = io.Copy(fw, fh)
				if err != nil {
					fh.Close()
					return false, "", fmt.Errorf("Error form open %s file: %v\n", key, err)
				}
				fh.Close()
			default:
				fmt.Println("Do match in key list (key: %s)", key)
			}
		}
	}

	bodyWriter.Close()
	req, _ := http.NewRequest(http.MethodPost, url, bodyBuf)
	configureRequestHeaders(req, c, http.MethodPost)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	resp, err1 := client.Do(req)
	if err1 != nil {
		return false, "", fmt.Errorf("http request failed: %v", err1)
	}

	return handleResponse(resp)
}

// make_call3 only for certificate local generate csr
func make_call3(c *FortiSDKClient, api string, http_method string, body interface{}, vdomparam string) (bool, string, error) {
	client := createHTTPClient()

	bodyPtr := body.(*map[string]interface{})
	bodyMap := *bodyPtr

	login_err := updateClientAuth(c)
	if login_err != nil {
		return false, "", login_err
	}

	url := configureUrl(c, api, vdomparam)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for key, value := range bodyMap {
		fw, err := bodyWriter.CreateFormField(key)
		_, err = io.Copy(fw, strings.NewReader(value.(string)))
		if err != nil {
			return false, "", fmt.Errorf("Error form created key %s: %v\n", key, err)
		}
	}

	bodyWriter.Close()
	req, _ := http.NewRequest(http.MethodPost, url, bodyBuf)
	configureRequestHeaders(req, c, http.MethodPost)
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

	resp, err1 := client.Do(req)
	if err1 != nil {
		return false, "", fmt.Errorf("http request failed: %v", err1)
	}

	return handleResponse(resp)
}

func createUpload(c *FortiSDKClient, method string, path string, params *map[string]interface{}, output map[string]interface{}, vdomparam string) (err error) {
	success, respBody, err := make_call2(c, path, method, params, vdomparam)
	if !success {
		return err
	}

	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return fmt.Errorf("Error parsing JSON in createUpload function: %v, response: %s", err, respBody)
	}

	if results, ok := result["results"].(map[string]interface{}); ok {
		fmt.Println("Results:", results)
	}

	// Handle response data
	if result["vdom"] != nil {
		output["vdom"] = result["vdom"]
	}

	if result["mkey"] != nil {
		output["mkey"] = result["mkey"]
	}

	if _, ok := result["payload"].(float64); ok {
		ret_code := result["payload"].(float64)
		if ret_code != 0 {
			str_code := strconv.FormatInt(int64(ret_code), 10)
			err = errors.New(fweb_err[str_code])
		}
	}

	return
}

// createUpload2 only for certificate local generate csr
func createUpload2(c *FortiSDKClient, method string, path string, params *map[string]interface{}, output map[string]interface{}, vdomparam string) (err error) {
	success, respBody, err := make_call3(c, path, method, params, vdomparam)
	if !success {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return fmt.Errorf("Error parsing JSON in createUpload function: %v, response: %s", err, respBody)
	}

	if results, ok := result["results"].(map[string]interface{}); ok {
		fmt.Println("Results:", results)
	}

	if result["vdom"] != nil {
		output["vdom"] = result["vdom"]
	}

	if result["mkey"] != nil {
		output["mkey"] = result["mkey"]
	}

	if _, ok := result["payload"].(float64); ok {
		ret_code := result["payload"].(float64)
		if ret_code != 0 {
			str_code := strconv.FormatInt(int64(ret_code), 10)
			err = errors.New(fweb_err[str_code])
		}
	}

	return
}

func createUpdate(c *FortiSDKClient, method string, path string, params *map[string]interface{}, output map[string]interface{}, vdomparam string) (err error) {
	success, respBody, err := make_call(c, path, method, params, vdomparam)
	if !success {
		return err
	}

	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return fmt.Errorf("Error parsing JSON in createUpdate function: %v, response: %s", err, respBody)
	}

	if results, ok := result["results"].(map[string]interface{}); ok {
		fmt.Println("Results:", results)
	}

	/*
		// Handle response data
		if result["vdom"] != nil {
			output["vdom"] = result["vdom"]
		}

		if result["mkey"] != nil {
			output["mkey"] = result["mkey"]
		}


			if _, ok := result["result"].(float64); ok {
				ret_code := result["payload"].(float64)
				if ret_code != 0 {
					str_code := strconv.FormatInt(int64(ret_code), 10)
					err = errors.New(fweb_err[str_code])
				}
			}
	*/
	return
}

func delete(c *FortiSDKClient, method string, path string, vdomparam string) (err error) {
	success, respBody, err := make_call(c, path, method, make(map[string]interface{}), vdomparam)
	if !success {
		fmt.Println("err: ", err)
		return err
	}

	// Parse response
	var result map[string]interface{}
	json.Unmarshal([]byte(respBody), &result)

	if results, ok := result["results"].(map[string]interface{}); ok {
		if results["status"] == "success" {
			return
		} else {
			fmt.Println("Results:", results)
		}
	}

	err = fortiAPIErrorFormat(result, string(respBody))

	return
}

func read(c *FortiSDKClient, method string, path string, bcomplex bool, vdomparam string) (mapTmp map[string]interface{}, err error) {
	success, respBody, err := make_call(c, path, method, make(map[string]interface{}), vdomparam)
	if !success {
		return nil, err
	}

	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON in read function: %v, response: %s", err, respBody)
	}

	if fortiAPIHttpStatus404Checking(result) == true {
		mapTmp = nil
		return
	}

	// Extract results section if it exists - this is the primary success case
	// First check if results is a map
	if results, ok := result["results"].(map[string]interface{}); ok {
		fmt.Println("Results:", results)
		mapTmp = results
		return mapTmp, nil
	}

	// Then check if results is an array and extract the first element
	if resultsArray, ok := result["results"].([]interface{}); ok && len(resultsArray) > 0 {
		if firstResult, ok := resultsArray[0].(map[string]interface{}); ok {
			fmt.Println("Results array first element:", firstResult)
			mapTmp = firstResult
			return mapTmp, nil
		}
	}

	// Only check for errors if we don't have results
	err = fortiAPIErrorFormat(result, string(respBody))
	if err == nil {
		// Fallback to original logic for other response formats
		switch result["results"].(type) {
		case map[string]interface{}:
			mapTmp = result["results"].(map[string]interface{})
			for k, v := range mapTmp {
				switch v.(type) {
				case string:
					mapTmp[k] = strings.TrimSuffix(v.(string), " ")
				}
			}
		}
	}
	return mapTmp, err
}

func read2(c *FortiSDKClient, method string, path string, bcomplex bool, vdomparam string) (resultsArray []interface{}, err error) {
	success, respBody, err := make_call(c, path, method, make(map[string]interface{}), vdomparam)
	if !success {
		return nil, err
	}

	// Parse response
	var result map[string]interface{}
	err = json.Unmarshal([]byte(respBody), &result)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON in read function: %v, response: %s", err, respBody)
	}

	if resultsArray, ok := result["results"].([]interface{}); ok && len(resultsArray) > 0 {
		return resultsArray, nil
	} else {
		return nil, fmt.Errorf("Error get resultsArray")
	}
}

func Login(ip string, username string, passwd string, port string) (string, []*http.Cookie, map[string]interface{}, string, error) {
	//check ip
	var respcookie []*http.Cookie
	errorcode := make(map[string]interface{}, 0)

	//ip := c.Hostname
	csrf_result := ""
	if net.ParseIP(ip) == nil {
		return "", respcookie, errorcode, "", fmt.Errorf("FortiWeb IP is required or invalid")
	}

	client := createHTTPClient()

	var response map[string]string
	//username := c.Username, passwd := c.Password, port := c.Port
	url := "https://" + ip + ":" + port + LOGING_URL

	fmt.Println("login url", url)
	body_string := "ajax=1&username=" + username + "&secretkey=" + passwd
	req_body := strings.NewReader(string(body_string))
	req, _ := http.NewRequest("POST", url, req_body)

	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("Login FortiWeb %s failed: %v\n", ip, err1)
		return "", respcookie, errorcode, "", fmt.Errorf("%v", err1)
	}

	defer resp.Body.Close()

	if 200 != resp.StatusCode {
		var tmp_obj interface{}
		buffer, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(buffer, &tmp_obj)
		output_str, _ := json.MarshalIndent(tmp_obj, "", "\t")
		return "", respcookie, errorcode, "", fmt.Errorf("%s", output_str)
	}

	respcookie = resp.Cookies()
	buffer, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buffer, &response)
	token := response["token"]

	for i := 0; i < 3; i++ {
		url = "https://" + ip + ":" + port + GET_CSRF_TOKEN_URL
		fmt.Printf("%s\n", url)
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		for _, v := range respcookie {
			req.AddCookie(v)
		}

		resp_err, err2 := client.Do(req)
		if err2 == nil && resp_err.StatusCode == 200 {
			csrf_buffer, _ := ioutil.ReadAll(resp_err.Body)
			csrf_result = gjson.Get(string(csrf_buffer), "resutls.admin.csrf_token").String()
			fmt.Println("Got CSRF Success\n")
			defer resp_err.Body.Close()
			break
		} else {
			if err2 == nil {
				resp_err.Body.Close()
			}
			fmt.Println("Got CSRF Failed %d \n", resp_err.StatusCode)
		}
		time.Sleep(1)
	}

	for i := 0; i < 3; i++ {
		url = "https://" + ip + ":" + port + GET_CSRF_TOKEN_URL
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		for _, v := range respcookie {
			req.AddCookie(v)
		}
		resp_err, err2 := client.Do(req)
		if err2 == nil && resp_err.StatusCode == 200 {
			csrf_buffer, _ := ioutil.ReadAll(resp.Body)
			var jsonResponse map[string]interface{}
			err := json.Unmarshal(csrf_buffer, &jsonResponse)
			if err != nil {
				fmt.Printf("Error parsing JSON: %v\n", err)
				continue
			}
			if results, ok := jsonResponse["results"].(map[string]interface{}); ok {
				if admin, ok := results["admin"].(map[string]interface{}); ok {
					if csrfToken, ok := admin["csrf_token"].(string); ok {
						csrf_result = csrfToken
						fmt.Println("Got CSRF Success")
						break
					}
				}
			}
		} else {
			if err2 == nil {
				resp_err.Body.Close()
			}
			fmt.Println("Got CSRF Failed %d \n", resp_err.StatusCode)
		}
		time.Sleep(1 * time.Second)
	}

	errorcode["0"] = "succeeded"
	return token, respcookie, errorcode, csrf_result, nil
}

func Logout(c *FortiSDKClient) error {
	//make_call(c, LOGOUT_URL, http.MethodGet, make(map[string]interface{}), "logout")
	client := createHTTPClient()
	var body_string []byte

	body := make(map[string]interface{})
	url := "https://" + c.Config.FwTarget + ":443" + LOGOUT_URL
	body_string, _ = json.Marshal(body)

	req_body := strings.NewReader(string(body_string))
	req, _ := http.NewRequest(http.MethodGet, url, req_body)
	configureRequestHeaders(req, c, http.MethodGet)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("http request Logout failed: %v\n", err1)
		return err1
	}

	defer resp.Body.Close()
	var tmp_obj interface{}

	buffer, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buffer, &tmp_obj)
	output_str, _ := json.MarshalIndent(tmp_obj, "", "\t")
	fmt.Println("http request Logout output_str: %s\n", string(output_str))

	return nil
}
