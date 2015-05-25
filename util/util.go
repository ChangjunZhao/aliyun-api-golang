package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ChangjunZhao/aliyun-api-golang/signer"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// HTTPExecuteError signals that a call to httpExecute failed.
type HTTPExecuteError struct {
	// RequestHeaders provides a stringified listing of request headers.
	RequestHeaders string
	// ResponseBodyBytes is the response read into a byte slice.
	ResponseBodyBytes []byte
	// Status is the status code string response.
	Status string
	// StatusCode is the parsed status code.
	StatusCode int
}

// Error provides a printable string description of an HTTPExecuteError.
func (e HTTPExecuteError) Error() string {
	return "HTTP response is not 200/OK as expected. Actual response: \n" +
		"\tResponse Status: '" + e.Status + "'\n" +
		"\tResponse Code: " + strconv.Itoa(e.StatusCode) + "\n" +
		"\tResponse Body: " + string(e.ResponseBodyBytes) + "\n" +
		"\tRequest Headers: " + e.RequestHeaders
}

//执行http请求
func httpExecute(
	method string, urlStr string, contentType string, body string, oauthParams *OrderedParams) (*http.Response, error) {
	// Create base request.
	req, err := http.NewRequest(method, urlStr, strings.NewReader(body))
	if err != nil {
		return nil, errors.New("NewRequest failed: " + err.Error())
	}
	HttpClient := &http.Client{}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, errors.New("Do: " + err.Error())
	}

	debugHeader := ""
	for k, vals := range req.Header {
		for _, val := range vals {
			debugHeader += "[key: " + k + ", val: " + val + "]"
		}
	}

	// StatusMultipleChoices is 300, any 2xx response should be treated as success
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)

		return resp, HTTPExecuteError{
			RequestHeaders:    debugHeader,
			ResponseBodyBytes: bytes,
			Status:            resp.Status,
			StatusCode:        resp.StatusCode,
		}
	}
	return resp, err
}

//获得HTTP请求的body部分内容
func getBody(method, url string, oauthParams *OrderedParams) (*string, error) {
	resp, err := httpExecute(method, url, "", "", oauthParams)
	if err != nil {
		return nil, errors.New("httpExecute: " + err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, errors.New("ReadAll: " + err.Error())
	}
	bodyStr := string(bodyBytes)
	/*
		if c.debug {
			fmt.Printf("STATUS: %d %s\n", resp.StatusCode, resp.Status)
			fmt.Println("BODY RESPONSE: " + bodyStr)
		}
	*/
	return &bodyStr, nil
}

//构建加上签名的API请求地址
func buildRequestUrl(serverUrl string, base64Signature string, params *OrderedParams) string {
	result := serverUrl + "?"
	for pos, key := range params.Keys() {
		if pos != 0 {
			result += "&"
		}
		result += fmt.Sprintf("%s=%s", key, params.Get(key))
	}
	result += "&Signature=" + url.QueryEscape(base64Signature)
	return result
}

//计算签名的字符串
func requestString(method string, urlPath string, params *OrderedParams) string {
	result := method + "&" + url.QueryEscape(urlPath)
	for pos, key := range params.Keys() {
		if pos == 0 {
			result += "&"
		} else {
			result += url.QueryEscape("&")
		}
		result += url.QueryEscape(fmt.Sprintf("%s=%s", key, params.Get(key)))
	}
	return result
}

//调用API
func CallApiServer(server string, signer *signer.SHA1Signer, params *OrderedParams, i interface{}) error {
	reqString := requestString("GET", "/", params)
	base64Signature, _ := signer.Sign(reqString)
	requestUrl := buildRequestUrl(server, base64Signature, params)
	result, err := getBody("GET", requestUrl, nil)
	if err == nil {
		json.Unmarshal([]byte(*result), &i)
		return nil
	} else {
		return err
	}
}

//排序后的参数列表
type OrderedParams struct {
	allParams   map[string]string
	keyOrdering []string
}

func NewOrderedParams() *OrderedParams {
	return &OrderedParams{
		allParams:   make(map[string]string),
		keyOrdering: make([]string, 0),
	}
}

func (o *OrderedParams) Get(key string) string {
	return o.allParams[key]
}

func (o *OrderedParams) Keys() []string {
	sort.Sort(o)
	return o.keyOrdering
}

func (o *OrderedParams) Add(key, value string) {
	o.AddUnescaped(key, url.QueryEscape(value))
}

func (o *OrderedParams) AddUnescaped(key, value string) {
	o.allParams[key] = value
	o.keyOrdering = append(o.keyOrdering, key)
}

func (o *OrderedParams) Len() int {
	return len(o.keyOrdering)
}

func (o *OrderedParams) Less(i int, j int) bool {
	return o.keyOrdering[i] < o.keyOrdering[j]
}

func (o *OrderedParams) Swap(i int, j int) {
	o.keyOrdering[i], o.keyOrdering[j] = o.keyOrdering[j], o.keyOrdering[i]
}

func (o *OrderedParams) Clone() *OrderedParams {
	clone := NewOrderedParams()
	for _, key := range o.Keys() {
		clone.AddUnescaped(key, o.Get(key))
	}
	return clone
}
