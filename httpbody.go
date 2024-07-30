package downbuff

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	A_IM                                               string
	ACCEPT                                             string
	ACCEPT_CHARSET                                     string
	ACCEPT_DATETIME                                    string
	ACCEPT_ENCODING                                    string
	ACCEPT_LANGUAGE                                    string
	ACCESS_CONTROL_REQUEST_METHOD                      string
	ACCESS_CONTROL_REQUEST_HEADERS                     string
	AUTHORIZATION                                      string
	COOKIE                                             string
	EXPECT                                             string
	FORWARDED                                          string
	FROM                                               string
	HOST                                               string
	HTTP2_SETTINGS                                     string
	IF_MATCH                                           string
	IF_MODIFIED_SINCE                                  time.Time
	IF_NONE_MATCH                                      string
	IF_RANGE                                           string
	IF_UNMODIFIED_SINCE                                time.Time
	MAX_FORWARDS                                       int64
	ORIGIN                                             string
	PRAGMA                                             string
	PREFER                                             string
	PROXY_AUTHORIZATION                                string
	RANGE                                              string
	REFERER                                            string
	TE                                                 string
	TRAILER                                            string
	TRANSFER_ENCODING                                  string
	USER_AGENT                                         string
	UPGRADE                                            string
	VIA                                                string
	WARNING                                            string
	UPGRADE_INSECURE_REQUESTS                          string
	X_REQUESTED_WITH                                   string
	DNT                                                string
	X_FORWARDED_FOR                                    string
	X_FORWARDED_HOST                                   string
	X_FORWARDED_PROTO                                  string
	FRONT_END_HTTPS                                    string
	X_HTTP_METHOD_OVERRIDE                             string
	X_ATT_DEVICEID                                     string
	X_WAP_PROFILE                                      string
	PROXY_CONNECTION                                   string
	X_UIDH                                             string
	X_CSRF_TOKEN                                       string
	X_REQUEST_ID                                       string
	X_CORRELATION_ID                                   string
	SAVE_DATA                                          string
	SEC_GPC                                            string
	ACCEPT_CH                                          string
	ACCESS_CONTROL_ALLOW_ORIGIN                        string
	ACCESS_CONTROL_ALLOW_CREDENTIALS                   string
	ACCESS_CONTROL_EXPOSE_HEADERS                      string
	ACCESS_CONTROL_MAX_AGE                             string
	ACCESS_CONTROL_ALLOW_METHODS                       string
	ACCESS_CONTROL_ALLOW_HEADERS                       string
	ACCEPT_PATCH                                       string
	ACCEPT_RANGES                                      string
	AGE                                                int64
	ALLOW                                              string
	ALT_SVC                                            string
	CACHE_CONTROL                                      string
	CONNECTION                                         string
	CONTENT_DISPOSITION                                string
	CONTENT_ENCODING                                   string
	CONTENT_LANGUAGE                                   string
	CONTENT_LENGTH                                     int64
	CONTENT_LOCATION                                   string
	CONTENT_MD5                                        string
	CONTENT_RANGE                                      string
	CONTENT_TYPE                                       string
	DATE                                               time.Time
	DELTA_BASE                                         string
	ETAG                                               string
	EXPIRES                                            time.Time
	IM                                                 string
	LAST_MODIFIED                                      time.Time
	LINK                                               string
	LOCATION                                           string
	P3P                                                string
	PREFERENCE_APPLIED                                 string
	PROXY_AUTHENTICATE                                 string
	PUBLIC_KEY_PINS                                    string
	RETRY_AFTER                                        string
	SERVER                                             string
	SET_COOKIE                                         string
	STRICT_TRANSPORT_SECURITY                          string
	TK                                                 string
	VARY                                               string
	WWW_AUTHENTICATE                                   string
	X_FRAME_OPTIONS                                    string
	CONTENT_SECURITY_POLICY                            string
	EXPECT_CT                                          string
	NEL                                                string
	PERMISSIONS_POLICY                                 string
	REFRESH                                            string
	REPORT_TO                                          string
	STATUS                                             string
	TIMING_ALLOW_ORIGIN                                string
	X_CONTENT_DURATION                                 int64
	X_CONTENT_TYPE_OPTIONS                             string
	X_POWERED_BY                                       string
	X_REDIRECT_BY                                      string
	X_UA_COMPATIBLE                                    string
	X_TURBO_CHARGED_BY                                 string
	PRIORITY                                           string
	SEC_CH_UA                                          string
	SEC_CH_UA_ARCH                                     string
	SEC_CH_UA_BITNESS                                  string
	SEC_CH_UA_FULL_VERSION                             string
	SEC_CH_UA_FULL_VERSION_LIST                        string
	SEC_CH_UA_MOBILE                                   string
	SEC_CH_UA_MODEL                                    string
	SEC_CH_UA_PLATFORM                                 string
	SEC_CH_UA_PLATFORM_VERSION                         string
	SEC_CH_UA_WOW64                                    string
	SEC_FETCH_DEST                                     string
	SEC_FETCH_MODE                                     string
	SEC_FETCH_SITE                                     string
	SEC_FETCH_USER                                     string
	SEC_PURPOSE                                        string
	SEC_WEBSOCKET_ACCEPT                               string
	SEC_METADATA                                       string
	KEEP_ALIVE                                         string
	ACCEPT_POST                                        string
	ALT_USED                                           string
	ATTRIBUTION_REPORTING_ELIGIBLEEXPERIMENTAL         string
	ATTRIBUTION_REPORTING_REGISTER_SOURCEEXPERIMENTAL  string
	ATTRIBUTION_REPORTING_REGISTER_TRIGGEREXPERIMENTAL string
	CLEAR_SITE_DATA                                    string
	CONTENT_DIGESTEXPERIMENTAL                         string
	CONTENT_DPRNON_STANDARDDEPRECATED                  string
	CONTENT_SECURITY_POLICY_REPORT_ONLY                string
	CRITICAL_CHEXPERIMENTAL                            string
	CROSS_ORIGIN_EMBEDDER_POLICY                       string
	CROSS_ORIGIN_OPENER_POLICY                         string
	CROSS_ORIGIN_RESOURCE_POLICY                       string
	DEVICE_MEMORY                                      string
	DIGESTNON_STANDARDDEPRECATED                       string
	DNTNON_STANDARDDEPRECATED                          string
	DOWNLINKEXPERIMENTAL                               string
	DPRNON_STANDARDDEPRECATED                          string
	EARLY_DATAEXPERIMENTAL                             string
	ECTEXPERIMENTAL                                    string
	NELEXPERIMENTAL                                    string
	NO_VARY_SEARCHEXPERIMENTAL                         string
	OBSERVE_BROWSING_TOPICSEXPERIMENTALNON_STANDARD    string
	ORIGIN_AGENT_CLUSTEREXPERIMENTAL                   string
	PRAGMADEPRECATED                                   string
	REFERRER_POLICY                                    string
	REPORTING_ENDPOINTS                                string
	REPR_DIGESTEXPERIMENTAL                            string
	RTTEXPERIMENTAL                                    string
	SAVE_DATAEXPERIMENTAL                              string
	SEC_BROWSING_TOPICSEXPERIMENTALNON_STANDARD        string
	SEC_CH_PREFERS_COLOR_SCHEMEEXPERIMENTAL            string
	SEC_CH_PREFERS_REDUCED_MOTIONEXPERIMENTAL          string
	SEC_CH_PREFERS_REDUCED_TRANSPARENCYEXPERIMENTAL    string
	SEC_CH_UAEXPERIMENTAL                              string
	SEC_CH_UA_ARCHEXPERIMENTAL                         string
	SEC_CH_UA_BITNESSEXPERIMENTAL                      string
	SEC_CH_UA_FULL_VERSIONDEPRECATED                   string
	SEC_CH_UA_FULL_VERSION_LISTEXPERIMENTAL            string
	SEC_CH_UA_MOBILEEXPERIMENTAL                       string
	SEC_CH_UA_MODELEXPERIMENTAL                        string
	SEC_CH_UA_PLATFORMEXPERIMENTAL                     string
	SEC_CH_UA_PLATFORM_VERSIONEXPERIMENTAL             string
	SEC_GPCEXPERIMENTALNON_STANDARD                    string
	SERVER_TIMING                                      string
	SERVICE_WORKER_NAVIGATION_PRELOAD                  string
	SET_LOGINEXPERIMENTAL                              string
	SOURCEMAP                                          string
	SPECULATION_RULESEXPERIMENTAL                      string
	SUPPORTS_LOADING_MODEEXPERIMENTAL                  string
	TKNON_STANDARDDEPRECATED                           string
	VIEWPORT_WIDTHNON_STANDARDDEPRECATED               string
	WANT_CONTENT_DIGESTEXPERIMENTAL                    string
	WANT_DIGESTNON_STANDARDDEPRECATED                  string
	WANT_REPR_DIGESTEXPERIMENTAL                       string
	WARNINGDEPRECATED                                  string
	WIDTHNON_STANDARDDEPRECATED                        string
	X_DNS_PREFETCH_CONTROLNON_STANDARD                 string
	X_FORWARDED_FORNON_STANDARD                        string
	X_FORWARDED_HOSTNON_STANDARD                       string
	X_FORWARDED_PROTONON_STANDARD                      string
	X_XSS_PROTECTIONNON_STANDARD                       string
	X_CACHE                                            string
	X_SERVED_BY                                        string
	CF_CONNECTING_IP                                   string
	CF_IPCOUNTRY                                       string
	CF_RAY                                             string
	CF_VISITOR                                         string
	CF_WORKER                                          string
	CF_CACHE_STATUS                                    string
	CF_REQUEST_ID                                      string
	CF_REQUEST_CONTROL                                 string
	CF_TRACE_ID                                        string
	CF_PROXY_ID                                        string
	CF_BGJ                                             string
	CF_POP                                             string
	CF_CACHE_TAG                                       string
	CF_CLOUDFLARED                                     string
	CF_ACCESS_AUTHENTICATED_USER_EMAIL                 string
	CF_ACCESS_AUTHENTICATED_USER_IDENTITY_PROVIDER     string
	CF_ACCESS_JWT_ASSERTION                            string
	CF_EDGE_CDN_COUNTRY                                string
	CF_EDGE_CDN_CONTINENT                              string
	CF_EDGE_CDN_CITY                                   string
	CF_EDGE_CDN_REGION                                 string
	CF_EDGE_CDN_TIMEZONE                               string
	CF_EDGE_CDN_LATITUDE                               string
	CF_EDGE_CDN_LONGITUDE                              string
	CF_REQUEST_ACCEPTED_LANGUAGE                       string
	CF_REQUEST_REFERER                                 string
	CF_REQUEST_USER_AGENT                              string
	CF_REQUEST_SCHEME                                  string
}

type BodyRequest struct {
	Path        string
	HttpVersion string // Available 1.0, 1.1, 2, 3
	Header      Header
	UnkHeaders  map[string]string
	Data        []byte
}

func (response BodyRequest) Summary() (summary string) {
	v := reflect.ValueOf(response.Header)
	for i := 0; i < v.NumField(); i++ {
		rawValue := v.Field(i).Interface()
		headerName := strings.ReplaceAll(v.Type().Field(i).Name, "_", "-")
		switch rawValue := rawValue.(type) {
		case int64:
			if rawValue != 0 {
				summary += fmt.Sprintf("%s: %d\n", headerName, rawValue)
			}
		case time.Time:
			if !rawValue.IsZero() {
				summary += fmt.Sprintf("%s: %s\n", headerName, rawValue.Format("Mon, 02 Jan 2006 15:04:05 MST"))
			}
		case string:
			if rawValue != "" {
				summary += fmt.Sprintf("%s: %s\n", headerName, rawValue)
			}
		}
	}
	for headerName, value := range response.UnkHeaders {
		summary += fmt.Sprintf("%s: %s\n", strings.ToUpper(headerName), value)
	}
	if len(summary) < 1 {
		return ""
	}
	return summary[:len(summary)-1]
}

func (response BodyResponse) Summary() (summary string) {
	v := reflect.ValueOf(response.Header)
	for i := 0; i < v.NumField(); i++ {
		rawValue := v.Field(i).Interface()
		headerName := strings.ReplaceAll(v.Type().Field(i).Name, "_", "-")
		switch rawValue := rawValue.(type) {
		case int64:
			if rawValue != 0 {
				summary += fmt.Sprintf("%s: %d\n", headerName, rawValue)
			}
		case time.Time:
			if !rawValue.IsZero() {
				summary += fmt.Sprintf("%s: %s\n", headerName, rawValue.Format("Mon, 02 Jan 2006 15:04:05 MST"))
			}
		case string:
			if rawValue != "" {
				summary += fmt.Sprintf("%s: %s\n", headerName, rawValue)
			}
		}
	}
	for headerName, value := range response.UnkHeaders {
		summary += fmt.Sprintf("%s: %s\n", headerName, value)
	}
	if len(summary) < 1 {
		return ""
	}
	return summary[:len(summary)-1]
}

func BuildPOST(postdata map[string]string) (data []byte) {
	for key, value := range postdata {
		data = append(data, []byte(fmt.Sprintf("%s=%s&", key, value))...)
	}
	return data[:len(data)-1]
}

func SerializeHeaders(method reqmethod, body BodyRequest) (serializedBody []byte) {
	if err := checkValid(map[string]string{"REQ_METHOD": string(method), "REQ_HTTPVERSION": body.HttpVersion}); err != nil {
		panic(err)
	}
	serializedBody = addElement(serializedBody, []byte(fmt.Sprintf("%s %s HTTP/%s", method, body.Path, body.HttpVersion)))
	v := reflect.ValueOf(body.Header)
	for i := 0; i < v.NumField(); i++ {
		rawValue := v.Field(i).Interface()
		headerName := strings.ReplaceAll(v.Type().Field(i).Name, "_", "-")
		switch rawValue := rawValue.(type) {
		case int64:
			if rawValue != 0 {
				serializedBody = addElement(serializedBody, []byte(fmt.Sprintf("%s: %d", headerName, rawValue)))
			}
		case time.Time:
			if !rawValue.IsZero() {
				serializedBody = addElement(serializedBody, []byte(fmt.Sprintf("%s: %s", headerName, rawValue.Format("Mon, 02 Jan 2006 15:04:05 MST"))))
			}
		case string:
			if rawValue != "" {
				serializedBody = addElement(serializedBody, []byte(fmt.Sprintf("%s: %s", headerName, rawValue)))
			}
		}
	}
	for headerName, value := range body.UnkHeaders {
		serializedBody = addElement(serializedBody, []byte(fmt.Sprintf("%s: %s", headerName, value)))
	}
	return addElement(serializedBody, nil)
}

type BodyResponse struct {
	HttpVersion string
	StatusCode  status
	StatusText  string
	Header      Header
	UnkHeaders   map[string]string
	Data        []byte
}

func packResponse(content []byte) (body BodyResponse) {
	var header []byte
	var ok bool
	if header, body.Data, ok = bytes.Cut(content, []byte{0x0d, 0x0a, 0x0d, 0x0a}); !ok {
		if len(header) == 0 {
			return BodyResponse{}
		}
		body.Data = nil
	}
	body.UnkHeaders = make(map[string]string)
	for index, headerline := range bytes.Split(header, []byte{0x0d, 0x0a}) {
		if index == 0 {
			responseline := bytes.SplitN(headerline, []byte(" "), 3)
			if len(responseline) != 3 {
				return BodyResponse{}
			}
			body.HttpVersion = string(responseline[0])
			body.StatusCode = func() status { i, _ := strconv.Atoi(string(responseline[1])); return status(i) }()
			body.StatusText = string(responseline[2])
		} else {
			headerlineKV := bytes.Split(headerline, []byte{0x3a, 0x20})
			if len(headerlineKV) == 2 {
				fn := strings.ToUpper(string(headerlineKV[0]))
				field := reflect.ValueOf(&body.Header).Elem().FieldByName(strings.ReplaceAll(fn, "-", "_"))
				if field.IsValid() {
					switch field.Kind() {
					case reflect.String:
						field.SetString(string(headerlineKV[1]))
					case reflect.Int64:
						if i, _ := strconv.ParseInt(string(headerlineKV[1]), 10, 64); i != 0 {
							field.SetInt(i)
						}
					case reflect.Struct:
						if convt, _ := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", string(headerlineKV[1])); convt != (time.Time{}) {
							field.Set(reflect.ValueOf(convt))
						}
					}
				} else {
					body.UnkHeaders[fn] = string(headerlineKV[1])
				}
			}
		}
	}
	return body
}

func addElement(body []byte, element []byte) (updatedBody []byte) {
	body = append(body, element...)
	return append(body, []byte{0x0d, 0x0a}...)
}
