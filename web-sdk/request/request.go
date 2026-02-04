package request

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/fortinetdev/terraform-provider-fortiweb/web-sdk/config"
)

// Request describes the request to FortiOS service
type Request struct {
	Config       config.Config
	HTTPRequest  *http.Request
	HTTPResponse *http.Response
	Path         string
	Params       interface{}
	Data         *bytes.Buffer
}

// New creates reqeust object with http method, path, params and data,
// It will save the http request, path, etc. for the next operations
// such as sending data, getting response, etc.
// It returns the created request object to the gobal plugin client.
func New(c config.Config, method string, path string, params interface{}, data *bytes.Buffer) *Request {
	var h *http.Request

	if data == nil {
		h, _ = http.NewRequest(method, "", nil)
	} else {
		h, _ = http.NewRequest(method, "", data)
	}

	r := &Request{
		Config:      c,
		Path:        path,
		HTTPRequest: h,
		Params:      params,
		Data:        data,
	}
	return r
}

// Build Request header

// Build Request Sign/Login Info

// Send request data to FortiOS.
// If errors are encountered, it returns the error.
func (r *Request) Send() error {
	return r.Send2(15, false)
}

// Send2 request data to FortiOS with bIgnoreVdom.
// If errors are encountered, it returns the error.
func (r *Request) Send2(retries int, ignvdom bool) error {
	//Build FortiOS
	//build Sign/Login INfo

	//httpReq.URL, err = url.Parse(clientInfo.Endpoint + operation.HTTPPath)

	r.HTTPRequest.Header.Set("Content-Type", "application/json")

	// Set authentication based on available credentials
	if r.Config.Auth.Username != "" && r.Config.Auth.Password != "" {
		// Use cookies for session-based authentication
		for _, cookie := range r.Config.Auth.Cookies {
			r.HTTPRequest.AddCookie(cookie)
		}
		// Add CSRF token if available
		if r.Config.Auth.CSRFToken != "" {
			r.HTTPRequest.Header.Set("X-CSRF-Token", r.Config.Auth.CSRFToken)
		}
	} else {
		// Use token-based authentication
		r.HTTPRequest.Header.Set("APITOKEN", r.Config.Auth.Token)
	}

	u := ""
	if ignvdom == true {
		u = buildURLWithoutVdom(r)
	} else {
		u = buildURL(r)
	}

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func buildURL3(r *Request, vdomparam string) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path

	if strings.Contains(u, "?") {
		u += "&"
	} else {
		u += "?"
	}

	if vdomparam != "" {
		u += "vdom="
		u += vdomparam
		u += "&"
	} else {
		if r.Config.Auth.Vdom != "" {
			u += "vdom="
			u += r.Config.Auth.Vdom
			u += "&"
		}
	}

	// Add access_token parameter only when not using username/password authentication
	if r.Config.Auth.Username == "" || r.Config.Auth.Password == "" {
		u += "access_token="
		u += r.Config.Auth.Token
		u += "&"
	}

	// Remove trailing & or ? if no parameters were added
	if strings.HasSuffix(u, "&") || strings.HasSuffix(u, "?") {
		u = strings.TrimSuffix(u, "&")
		u = strings.TrimSuffix(u, "?")
	}

	return u
}

// Send3 request data to FortiOS with custom vdom.
// If errors are encountered, it returns the error.
func (r *Request) Send4(vdomparam string, content_type string) error {
	retries := 15

	r.HTTPRequest.Header.Set("Content-Type", content_type)

	// Set authentication based on available credentials
	if r.Config.Auth.Username != "" && r.Config.Auth.Password != "" {
		// Use cookies for session-based authentication
		for _, cookie := range r.Config.Auth.Cookies {
			r.HTTPRequest.AddCookie(cookie)
		}
		// Add CSRF token if available
		if r.Config.Auth.CSRFToken != "" {
			r.HTTPRequest.Header.Set("X-CSRF-Token", r.Config.Auth.CSRFToken)
		}
	} else {
		// Use token-based authentication
		r.HTTPRequest.Header.Set("APITOKEN", r.Config.Auth.Token)
	}

	u := buildURL3(r, vdomparam)
	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

// Send3 request data to FortiOS with custom vdom.
// If errors are encountered, it returns the error.
func (r *Request) Send3(vdomparam string) error {
	retries := 15

	r.HTTPRequest.Header.Set("Content-Type", "application/json")

	// Set authentication based on available credentials
	if r.Config.Auth.Username != "" && r.Config.Auth.Password != "" {
		// Use cookies for session-based authentication
		for _, cookie := range r.Config.Auth.Cookies {
			r.HTTPRequest.AddCookie(cookie)
		}
		// Add CSRF token if available
		if r.Config.Auth.CSRFToken != "" {
			r.HTTPRequest.Header.Set("X-CSRF-Token", r.Config.Auth.CSRFToken)
		}
	} else {
		// Use token-based authentication
		r.HTTPRequest.Header.Set("APITOKEN", r.Config.Auth.Token)
	}

	u := buildURL3(r, vdomparam)
	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

// Send_ADC request data to FortADC with custom vdom.
// If errors are encountered, it returns the error.
func (r *Request) Send_ADC(vdomparam string) error {
	retries := 15

	r.HTTPRequest.Header.Set("X-CSRF-Token", r.Config.Auth.CSRFToken)
	r.HTTPRequest.Header.Set("Content-Type", "application/json;charset=utf-8")
	r.HTTPRequest.Header.Set("Cache-Control", "no-cache")
	r.HTTPRequest.Header.Set("Accept", "application/json, text/plain, */*")
	r.HTTPRequest.Header.Set("Accept-Language", "en-US,en;q=0.5")
	r.HTTPRequest.Header.Set("Accept-Encoding", "gzip, deflate, br")

	u := buildURL3(r, vdomparam)
	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > retries {
				err = fmt.Errorf("lost connection to firewall with error: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}

func filterapikey(v string) string {
	re, _ := regexp.Compile("access_token=.*?\"")
	res := re.ReplaceAllString(v, "access_token=***************\"")

	return res
}

func buildURLWithoutVdom(r *Request) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path
	u += "?"

	// Add access_token parameter only when not using username/password authentication
	if r.Config.Auth.Username == "" || r.Config.Auth.Password == "" {
		u += "access_token="
		u += r.Config.Auth.Token
	}

	// Remove trailing & or ? if no parameters were added
	if strings.HasSuffix(u, "&") || strings.HasSuffix(u, "?") {
		u = strings.TrimSuffix(u, "&")
		u = strings.TrimSuffix(u, "?")
	}

	return u
}

func buildURL(r *Request) string {
	u := "https://"
	u += r.Config.FwTarget
	u += r.Path
	u += "?"

	if r.Config.Auth.Vdom != "" {
		u += "vdom="
		u += r.Config.Auth.Vdom
		u += "&"
	}

	// Add access_token parameter only when not using username/password authentication
	if r.Config.Auth.Username == "" || r.Config.Auth.Password == "" {
		u += "access_token="
		u += r.Config.Auth.Token
		u += "&"
	}

	// Remove trailing & or ? if no parameters were added
	if strings.HasSuffix(u, "&") || strings.HasSuffix(u, "?") {
		u = strings.TrimSuffix(u, "&")
		u = strings.TrimSuffix(u, "?")
	}

	return u
}

// SendWithSpecialParams sends request data to FortiOS with special URL paramaters.
// If errors are encountered, it returns the error.
func (r *Request) SendWithSpecialParams(s, vdomparam string) error {
	r.HTTPRequest.Header.Set("Content-Type", "application/json")

	// Set authentication based on available credentials
	if r.Config.Auth.Username != "" && r.Config.Auth.Password != "" {
		// Use cookies for session-based authentication
		for _, cookie := range r.Config.Auth.Cookies {
			r.HTTPRequest.AddCookie(cookie)
		}
		// Add CSRF token if available
		if r.Config.Auth.CSRFToken != "" {
			r.HTTPRequest.Header.Set("X-CSRF-Token", r.Config.Auth.CSRFToken)
		}
	} else {
		// Use token-based authentication
		r.HTTPRequest.Header.Set("APITOKEN", r.Config.Auth.Token)
	}

	u := buildURL3(r, vdomparam)

	if s != "" {
		u += "&"
		u += s
	}

	var err error
	r.HTTPRequest.URL, err = url.Parse(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	retry := 0
	for {
		//Send
		rsp, errdo := r.Config.HTTPCon.Do(r.HTTPRequest)
		r.HTTPResponse = rsp
		if errdo != nil {
			if strings.Contains(errdo.Error(), "x509: ") {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}

			if retry > 15 {
				err = fmt.Errorf("Error found: %v", filterapikey(errdo.Error()))
				break
			}
			time.Sleep(time.Duration(1) * time.Second)
			log.Printf("Error found: %v, will resend again %s, %d", filterapikey(errdo.Error()), u, retry)

			retry++

		} else {
			break
		}
	}

	return err
}
