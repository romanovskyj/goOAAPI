## Odin Automation API library

### Overview
A golang wrapper for managing [Odin Automation](http://www.odin.com/products/premium/) API

### Installation
To install goOAAPI package run `go get https://github.com/romanovskyj/goOAAPI`. To use it in application add `github.com/kolo/xmlrpc` string to import statement.

### Usage

Below you can find the sample of creating new subscription:

```
type Resource struct{
    ResourceId int `xmlrpc:"resource_id"`
    ResourceLimit int `xmlrpc:"resource_limit"`
}

type RequestParams struct{
    AccountId int `xmlrpc:"account_id"`
    SubscriptionId int `xmlrpc:"subscription_id"`
    ServiceTemplateId int `xmlrpc:"service_template_id"`
    ResourceLimits []Resource `xmlrpc:"resource_limits"`
    ApplicationInstanceId int `xmlrpc:"application_instance_id"`
}

func main() {
    oa, _ := goOAAPI.GetClient("oa.hostname", "admin", "password", false, 8440)

    resources := []Resource {
        {1001210,1},
        {1001209, 1},
        {1001214, 0},
        {1001212,0},
        {1001211, 0},
        {1001213, 1},
    }

    params := RequestParams{
        AccountId: 201,
        ApplicationInstanceId: 222,
        SubscriptionId: 100012,
        ServiceTemplateId: 165,
        ResourceLimits: resources,
    }

    answer, _ := oa.ActivateSubscription(params)

    fmt.Println(answer["subscription_id"])
}

```

As it can be seen at the code, it needs to precreate structures for passing it as request parameters. As for the response - it response string map with answer values

### Implementation details
The [xmlrpc](https://github.com/kolo/xmlrpc) library is used for sending requests.