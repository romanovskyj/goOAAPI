package goOAAPI

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"github.com/kolo/xmlrpc"
)

// TODO: Consider glog using

// APS is just a link to oa object. It needs for comfortable api using (the same as in doc), e.g:
// oa.APS.GetapplicationInstance(params)
type APS struct {
	client *xmlrpc.Client
}

type OA struct {
	client *xmlrpc.Client
	APS    *APS
}

type Response struct {
	ErrorMessage string `xmlrpc:"error_message"`
	Result       interface{} `xmlrpc:"result"`
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

	return &oa, nil
}

func call(client *xmlrpc.Client, method string, params interface{}) (interface{}, error) {
	var response Response

	err := client.Call(method, params, &response)

	if err != nil {
		fmt.Printf("Could not send request: %v\n", err)
		return response, err
	}

	if response.ErrorMessage != "" {
		fmt.Printf("Failed to execute request: %v\n", response.ErrorMessage)
		return response.Result, errors.New(response.ErrorMessage)
	}

	return response.Result, nil
}

func (oa *OA) SendRequest(method string, params interface{}) (interface{}, error) {
	return call(oa.client, method, params)
}

func (aps *APS) SendRequest(method string, params interface{}) (interface{}, error) {
	return call(aps.client, method, params)
}

func (oa *OA) activateST(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.activateST", params)
}

func (oa *OA) activateSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.activateSubscription", params)
}

func (oa *OA) addAccount(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addAccount", params)
}

func (oa *OA) addAccountMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addAccountMember", params)
}

func (oa *OA) addDNSHosting(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addDNSHosting", params)
}

func (oa *OA) addDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addDomain", params)
}

func (oa *OA) addDomainRequest(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addDomainRequest", params)
}

func (oa *OA) addDomainToAccount(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addDomainToAccount", params)
}

func (oa *OA) bindServicesToDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.bindServicesToDomain", params)
}

func (oa *OA) getRequiredNameServers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getRequiredNameServers", params)
}

func (oa *OA) syncNameServers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.syncNameServers", params)
}

func (oa *OA) unbindServicesFromDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.unbindServicesFromDomain", params)
}

func (oa *OA) addProvisioningAttributes(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addProvisioningAttributes", params)
}

func (oa *OA) addPTRRecord(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addPTRRecord", params)
}

func (oa *OA) addResourceType(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addResourceType", params)
}

func (oa *OA) addResourceTypeToServiceTemplate(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addResourceTypeToServiceTemplate", params)
}

func (oa *OA) addServiceTemplate(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addServiceTemplate", params)
}

func (oa *OA) addSubdomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addSubdomain", params)
}

func (oa *OA) addSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addSubscription", params)
}

func (oa *OA) addUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.addUser", params)
}

func (oa *OA) assignIPPool(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.assignIPPool", params)
}

func (oa *OA) assignRolesToMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.assignRolesToMember", params)
}

func (oa *OA) attachIPPool(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.attachIPPool", params)
}

func (oa *OA) batchRequest(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.batchRequest", params)
}

func (oa *OA) bindIPPool(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.bindIPPool", params)
}

func (oa *OA) brandDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.brandDomain", params)
}

func (oa *OA) changeAccountOwner(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.changeAccountOwner", params)
}

func (oa *OA) changeUserPassword(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.changeUserPassword", params)
}

func (oa *OA) checkMoveSubscriptions(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.checkMoveSubscriptions", params)
}

func (oa *OA) checkServiceTemplatesForProvisioning(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.checkServiceTemplatesForProvisioning", params)
}

func (oa *OA) checkPassword(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.checkPassword", params)
}

func (oa *OA) cloneServiceTemplate(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.cloneServiceTemplate", params)
}

func (oa *OA) createDatabase(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.createDatabase", params)
}

func (oa *OA) createDatabaseUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.createDatabaseUser", params)
}

func (oa *OA) createDNSRecord(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.createDNSRecord", params)
}

func (oa *OA) deactivateST(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.deactivateST", params)
}

func (oa *OA) deleteDNSRecord(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.deleteDNSRecord", params)
}

func (oa *OA) disableDNSRecord(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.disableDNSRecord", params)
}

func (oa *OA) detachIPPool(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.detachIPPool", params)
}

func (oa *OA) disableAccount(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.disableAccount", params)
}

func (oa *OA) disableAccountMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.disableAccountMember", params)
}

func (oa *OA) disableDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.disableDomain", params)
}

func (oa *OA) enableDNSRecord(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableDNSRecord", params)
}

func (oa *OA) disableSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.disableSubscription", params)
}

func (oa *OA) disableUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.disableUser", params)
}

func (oa *OA) enableAccount(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableAccount", params)
}

func (oa *OA) enableAccountMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableAccountMember", params)
}

func (oa *OA) enableAccountWideServiceUsers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableAccountWideServiceUsers", params)
}

func (oa *OA) enableDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableDomain", params)
}

func (oa *OA) enableSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableSubscription", params)
}

func (oa *OA) enableUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.enableUser", params)
}

func (oa *OA) findHost(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.findHost", params)
}

func (oa *OA) getAccountDomains(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountDomains", params)
}

func (oa *OA) getAccountRoles(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountRoles", params)
}

func (oa *OA) getAccountMemberRoles(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountMemberRoles", params)
}

func (oa *OA) getAvailableSkins(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAvailableSkins", params)
}

func (oa *OA) getBrandInfo(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getBrandInfo", params)
}

func (oa *OA) getCustomerSubscriptionsResources(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getCustomerSubscriptionsResources", params)
}

func (oa *OA) getDomainsForBrandCreation(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getDomainsForBrandCreation", params)
}

func (oa *OA) importBrandCertificate(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.importBrandCertificate", params)
}

func (oa *OA) getDomainNameServers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getDomainNameServers", params)
}

func (oa *OA) getHost(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getHost", params)
}

func (oa *OA) getResourceTypesByClass(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getResourceTypesByClass", params)
}

func (oa *OA) getServiceTemplateList(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getServiceTemplateList", params)
}

func (oa *OA) getVendorCustomers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getVendorCustomers", params)
}

func (oa *OA) getCustomerSubscriptions(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getCustomerSubscriptions", params)
}

func (oa *OA) getAccountInfo(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountInfo", params)
}

func (oa *OA) getAccountMembers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountMembers", params)
}

func (oa *OA) getAccountMemberByLogin(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountMemberByLogin", params)
}

func (oa *OA) getAccountSubscriptions(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountSubscriptions", params)
}

func (oa *OA) getAccountServiceUsers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getAccountServiceUsers", params)
}

func (oa *OA) getDomainByName(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getDomainByName", params)
}

func (oa *OA) getDomainSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getDomainSubscription", params)
}

func (oa *OA) getDomainList(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getDomainList", params)
}

func (oa *OA) getDNSRecords(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getDNSRecords", params)
}

func (oa *OA) getFTPUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getFTPUser", params)
}

func (oa *OA) getNameServers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getNameServers", params)
}

func (oa *OA) getProvisioningAttributes(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getProvisioningAttributes", params)
}

func (oa *OA) getRequestStatus(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getRequestStatus", params)
}

func (oa *OA) getResourceUsage(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getResourceUsage", params)
}

func (oa *OA) getResourceUsageForPeriod(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getResourceUsageForPeriod", params)
}

func (oa *OA) getServiceTemplate(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getServiceTemplate", params)
}

func (oa *OA) getSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getSubscription", params)
}

func (oa *OA) getSubscriptionWebspaces(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getSubscriptionWebspaces", params)
}

func (oa *OA) getUserByLogin(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getUserByLogin", params)
}

func (oa *OA) getUserInfo(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getUserInfo", params)
}

func (oa *OA) getUsers(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getUsers", params)
}

func (oa *OA) getWebspacesList(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getWebspacesList", params)
}

func (oa *OA) importCertificate(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.importCertificate", params)
}

func (oa *OA) installPleskLicense(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.installPleskLicense", params)
}

func (oa *OA) migrateSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.migrateSubscription", params)
}

func (oa *OA) modifyUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.modifyUser", params)
}

func (oa *OA) moveSubscriptions(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.moveSubscriptions", params)
}

func (oa *OA) orderSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.orderSubscription", params)
}

func (oa *OA) promoteToReseller(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.promoteToReseller", params)
}

func (oa *OA) provisionSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.provisionSubscription", params)
}

func (oa *OA) registerExternalSystem(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.registerExternalSystem", params)
}

func (oa *OA) setExternalSystemConfig(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setExternalSystemConfig", params)
}

func (oa *OA) getExternalSystemList(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.getExternalSystemList", params)
}

func (oa *OA) unregisterExternalSystem(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.unregisterExternalSystem", params)
}

func (oa *OA) registerSharedNode(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.registerSharedNode", params)
}

func (oa *OA) registerWindowsNode(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.registerWindowsNode", params)
}

func (oa *OA) removeAccount(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeAccount", params)
}

func (oa *OA) removeAccountMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeAccountMember", params)
}

func (oa *OA) removeDNSHosting(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeDNSHosting", params)
}

func (oa *OA) removeDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeDomain", params)
}

func (oa *OA) removeLicense(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeLicense", params)
}

func (oa *OA) removePTRRecord(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removePTRRecord", params)
}

func (oa *OA) removeSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeSubscription", params)
}

func (oa *OA) removeUser(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.removeUser", params)
}

func (oa *OA) resetResourceUsage(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.resetResourceUsage", params)
}

func (oa *OA) revokeIPPool(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.revokeIPPool", params)
}

func (oa *OA) revokeRolesFromMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.revokeRolesFromMember", params)
}

func (oa *OA) revokePleskLicense(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.revokePleskLicense", params)
}

func (oa *OA) setAccountAuthData(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setAccountAuthData", params)
}

func (oa *OA) setAccountInfo(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setAccountInfo", params)
}

func (oa *OA) setDomainRegistrarStatus(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setDomainRegistrarStatus", params)
}

func (oa *OA) setFTPPassword(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setFTPPassword", params)
}

func (oa *OA) setHostAttributes(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setHostAttributes", params)
}

func (oa *OA) setHostReadyToProvide(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setHostReadyToProvide", params)
}

func (oa *OA) setMemberInfo(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setMemberInfo", params)
}

func (oa *OA) setMemberPassword(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setMemberPassword", params)
}

func (oa *OA) setResourceTypeLimit(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setResourceTypeLimit", params)
}

func (oa *OA) setResourceTypeLimits(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setResourceTypeLimits", params)
}

func (oa *OA) setRTActParams(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setRTActParams", params)
}

func (oa *OA) setRTAttributes(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setRTAttributes", params)
}

func (oa *OA) setSTActivationParams(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setSTActivationParams", params)
}

func (oa *OA) setSTRTLimits(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setSTRTLimits", params)
}

func (oa *OA) setSubscriptionName(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setSubscriptionName", params)
}

func (oa *OA) setSystemProperty(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.setSystemProperty", params)
}

func (oa *OA) uploadLicense(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.uploadLicense", params)
}

func (oa *OA) unbindIPPool(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.unbindIPPool", params)
}

func (oa *OA) unbrandDomain(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.unbrandDomain", params)
}

func (oa *OA) unsetHostAttributes(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.unsetHostAttributes", params)
}

func (oa *OA) updateAccountAndAccountMember(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.updateAccountAndAccountMember", params)
}

func (oa *OA) upgradeSubscription(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.upgradeSubscription", params)
}

func (oa *OA) bindUserToSubscriptio(params interface{}) (interface{}, error) {
	return oa.SendRequest("pem.bindUserToSubscriptio", params)
}