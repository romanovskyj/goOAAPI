package goOAAPI

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/kolo/xmlrpc"
)

// TODO: Consider glog using

type OA struct {
	client *xmlrpc.Client
	APS    *APS
	Ad     *Ad
	Cqmail *Cqmail
	IIS    *IIS
	Mssql *Mssql
	Mysql *Mysql
	ProFTPD *ProFTPD
}

// APS is just a link to oa object. It needs for comfortable api using (the same as in doc), e.g:
// oa.APS.GetapplicationInstance(params). The same is for other auxiliary structures such as ad.binding etc
type APS struct {
	client *xmlrpc.Client
}

type Binding struct {
	client *xmlrpc.Client
}

type Ad struct {
	Binding *Binding
}

type Cqmail struct {
	client *xmlrpc.Client
}

type Exchange struct {
	client *xmlrpc.Client
}

type IIS struct {
	client *xmlrpc.Client
}

type Mssql struct {
	client *xmlrpc.Client
}

type Mysql struct {
	client *xmlrpc.Client
}

type ProFTPD struct {
	client *xmlrpc.Client
}

type Response struct {
	ErrorMessage string `xmlrpc:"error_message"`
	Result       map[string]interface{} `xmlrpc:"result"`
}

func GetClient(hostname string, login string, password string, ssl bool, port int) (*OA, error) {
	var urlComponents []string

	if ssl {
		urlComponents = append(urlComponents, "https://")
	} else {
		urlComponents = append(urlComponents, "http://")
	}

	urlComponents = append(urlComponents, login, ":", password, "@", hostname, ":", strconv.Itoa(port))

	url := strings.Join(urlComponents, "")

	client, _ := xmlrpc.NewClient(url, nil)

	oa := OA{client: client}
	oa.APS = &APS{client: client}
	oa.Ad = &Ad{Binding: &Binding{client: client}}
	oa.Cqmail = &Cqmail{client: client}
	oa.IIS = &IIS{client: client}

	return &oa, nil
}

func call(client *xmlrpc.Client, method string, params interface{}) (map[string]interface{}, error) {
	var response Response

	err := client.Call(method, params, &response)

	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		return response.Result, err
	}

	if response.ErrorMessage != "" {
		fmt.Printf("Failed to execute request: %v\n", response.ErrorMessage)
		return response.Result, errors.New(response.ErrorMessage)
	}

	return response.Result, nil
}

func (oa *OA) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(oa.client, method, params)
}

func (aps *APS) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(aps.client, method, params)
}

func (binding *Binding) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(binding.client, method, params)
}

func (cqmail *Cqmail) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(cqmail.client, method, params)
}

func (exchange *Exchange) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(exchange.client, method, params)
}

func (iis *IIS) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(iis.client, method, params)
}

func (mssql *Mssql) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(mssql.client, method, params)
}

func (mysql *Mysql) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(mysql.client, method, params)
}

func (proftpd *ProFTPD) SendRequest(method string, params interface{}) (map[string]interface{}, error) {
	return call(proftpd.client, method, params)
}