package goOAAPI

import "github.com/kolo/xmlrpc"

type OA struct{
	Hostname string
	Username string
}

type Result struct{
	ErrorMessage string `xmlrpc:"error_message"`
}

type RequestParams struct{
	subscription_id int
	service_template_id int
}

func (oa *OA) ActivateSubscription() Result {
	result := Result{}

	client, _ := xmlrpc.NewClient(oa.Hostname, nil)
	client.Call("pem.activateSubscription", RequestParams{3,4}, &result)

	return result
}

func GetClient(hostname string, username string, password string, ssl bool, port int) (*OA, error) {
	oa := OA{hostname, username}

	return &oa, nil
}