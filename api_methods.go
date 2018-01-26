package goOAAPI

func (oa *OA) ActivateST(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.activateST", params)
}

func (oa *OA) ActivateSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.activateSubscription", params)
}

func (binding *Binding) AddUserBinding(params interface{}) (map[string]interface{}, error) {
	return binding.SendRequest("pem.ad.binding.addUserBinding", params)
}

func (binding *Binding) RemoveUserBinding(params interface{}) (map[string]interface{}, error) {
	return binding.SendRequest("pem.ad.binding.removeUserBinding", params)
}

func (binding *Binding) GetUserInfo(params interface{}) (map[string]interface{}, error) {
	return binding.SendRequest("pem.ad.binding.getUserInfo", params)
}

func (oa *OA) AddAccount(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addAccount", params)
}

func (oa *OA) AddAccountMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addAccountMember", params)
}

func (oa *OA) AddDNSHosting(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addDNSHosting", params)
}

func (oa *OA) AddDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addDomain", params)
}

func (oa *OA) AddDomainRequest(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addDomainRequest", params)
}

func (oa *OA) AddDomainToAccount(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addDomainToAccount", params)
}

func (aps *APS) GetApplicationInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationInstance", params)
}

func (aps *APS) GetApplicationInstanceSettings(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationInstanceSettings", params)
}

func (aps *APS) GetApplicationInstances(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationInstances", params)
}

func (aps *APS) GetApplicationInstanceToken(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationInstanceToken", params)
}

func (aps *APS) GetAccountToken(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getAccountToken", params)
}

func (aps *APS) GetUserToken(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getUserToken", params)
}

func (aps *APS) GetApplicationSettings(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationSettings", params)
}

func (aps *APS) GetPackage(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getPackage", params)
}

func (aps *APS) GetProvisioningSettings(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getProvisioningSettings", params)
}

func (aps *APS) GetSubscriptionApplicationInstances(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getSubscriptionApplicationInstances", params)
}

func (aps *APS) ImportPackage(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.importPackage", params)
}

func (aps *APS) ProvideApplicationInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.provideApplicationInstance", params)
}

func (aps *APS) RemoveApplication(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.removeApplication", params)
}

func (aps *APS) SetApplicationInstanceSettings(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.setApplicationInstanceSettings", params)
}

func (aps *APS) UnimportPackage(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.unimportPackage", params)
}

func (aps *APS) UnprovideApplicationInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.unprovideApplicationInstance", params)
}

func (aps *APS) UpgradeApplicationInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.upgradeApplicationInstance", params)
}

func (aps *APS) GetApplicationInstanceLicenseActivationData(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationInstanceLicenseActivationData", params)
}

func (aps *APS) GetApplicationLicenseInfo(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationLicenseInfo", params)
}

func (aps *APS) GetServiceInstances(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getServiceInstances", params)
}

func (aps *APS) SetServiceInstanceResourceType(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.setServiceInstanceResourceType", params)
}

func (aps *APS) ProvideServiceInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.provideServiceInstance", params)
}

func (aps *APS) UnprovideServiceInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.unprovideServiceInstance", params)
}

func (aps *APS) SetServiceInstanceSettings(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.setServiceInstanceSettings", params)
}

func (aps *APS) GetApplicationInstanceServiceResourceTypes(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplicationInstanceServiceResourceTypes", params)
}

func (aps *APS) GetApplications(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getApplications", params)
}

func (aps *APS) GetUserServiceInstances(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.getUserServiceInstances", params)
}

func (aps *APS) InstallApplicationInstanceLicense(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.installApplicationInstanceLicense", params)
}

func (aps *APS) RemoveApplicationInstanceLicense(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.removeApplicationInstanceLicense", params)
}

func (aps *APS) RegisterUserInApplicationInstance(params interface{}) (map[string]interface{}, error) {
	return aps.SendRequest("pem.APS.registerUserInApplicationInstance", params)
}

func (oa *OA) BindServicesToDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.bindServicesToDomain", params)
}

func (oa *OA) GetRequiredNameServers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getRequiredNameServers", params)
}

func (oa *OA) SyncNameServers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.syncNameServers", params)
}

func (oa *OA) UnbindServicesFromDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.unbindServicesFromDomain", params)
}

func (oa *OA) AddProvisioningAttributes(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addProvisioningAttributes", params)
}

func (oa *OA) AddPTRRecord(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addPTRRecord", params)
}

func (oa *OA) AddResourceType(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addResourceType", params)
}

func (oa *OA) AddResourceTypeToServiceTemplate(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addResourceTypeToServiceTemplate", params)
}

func (oa *OA) AddServiceTemplate(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addServiceTemplate", params)
}

func (oa *OA) AddSubdomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addSubdomain", params)
}

func (oa *OA) AddSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addSubscription", params)
}

func (oa *OA) AddUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.addUser", params)
}

func (oa *OA) AssignIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.assignIPPool", params)
}

func (oa *OA) AssignRolesToMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.assignRolesToMember", params)
}

func (oa *OA) AttachIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.attachIPPool", params)
}

func (oa *OA) BatchRequest(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.batchRequest", params)
}

func (oa *OA) BindIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.bindIPPool", params)
}

func (oa *OA) BrandDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.brandDomain", params)
}

func (oa *OA) ChangeAccountOwner(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.changeAccountOwner", params)
}

func (oa *OA) ChangeUserPassword(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.changeUserPassword", params)
}

func (oa *OA) CheckMoveSubscriptions(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.checkMoveSubscriptions", params)
}

func (oa *OA) CheckServiceTemplatesForProvisioning(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.checkServiceTemplatesForProvisioning", params)
}

func (oa *OA) CheckPassword(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.checkPassword", params)
}

func (cqmail *Cqmail) AddMailbox(params interface{}) (map[string]interface{}, error) {
	return cqmail.SendRequest("pem.cqmail.addMailbox", params)
}

func (cqmail *Cqmail) DelMailname(params interface{}) (map[string]interface{}, error) {
	return cqmail.SendRequest("pem.cqmail.delMailname", params)
}

func (cqmail *Cqmail) EditMailname(params interface{}) (map[string]interface{}, error) {
	return cqmail.SendRequest("pem.cqmail.editMailname", params)
}

func (cqmail *Cqmail) AddMailForwarding(params interface{}) (map[string]interface{}, error) {
	return cqmail.SendRequest("pem.cqmail.addMailForwarding", params)
}

func (cqmail *Cqmail) EditEmailAddresses(params interface{}) (map[string]interface{}, error) {
	return cqmail.SendRequest("pem.cqmail.editEmailAddresses", params)
}

func (oa *OA) CloneServiceTemplate(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.cloneServiceTemplate", params)
}

func (oa *OA) CreateDatabase(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.createDatabase", params)
}

func (oa *OA) CreateDatabaseUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.createDatabaseUser", params)
}

func (oa *OA) CreateDNSRecord(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.createDNSRecord", params)
}

func (oa *OA) CreateIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.createIPPool", params)
}

func (oa *OA) DeactivateST(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.deactivateST", params)
}

func (oa *OA) DeleteDNSRecord(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.deleteDNSRecord", params)
}

func (oa *OA) DisableDNSRecord(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.disableDNSRecord", params)
}

func (oa *OA) DetachIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.detachIPPool", params)
}

func (oa *OA) DisableAccount(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.disableAccount", params)
}

func (oa *OA) DisableAccountMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.disableAccountMember", params)
}

func (oa *OA) DisableDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.disableDomain", params)
}

func (oa *OA) EnableDNSRecord(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableDNSRecord", params)
}

func (oa *OA) DisableSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.disableSubscription", params)
}

func (oa *OA) DisableUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.disableUser", params)
}

func (oa *OA) EnableAccount(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableAccount", params)
}

func (oa *OA) EnableAccountMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableAccountMember", params)
}

func (oa *OA) EnableAccountWideServiceUsers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableAccountWideServiceUsers", params)
}

func (oa *OA) EnableDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableDomain", params)
}

func (oa *OA) EnableSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableSubscription", params)
}

func (oa *OA) EnableUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.enableUser", params)
}

func (exchange *Exchange) AddEmailAddresses(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addEmailAddresses", params)
}

func (exchange *Exchange) AddMailbox(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addMailbox", params)
}

func (exchange *Exchange) ChangePrimaryEmailAddress(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.changePrimaryEmailAddress", params)
}

func (exchange *Exchange) DisableForwarding(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.disableForwarding", params)
}

func (exchange *Exchange) EnableForwarding(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.enableForwarding", params)
}

func (exchange *Exchange) GetEmailAddresses(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getEmailAddresses", params)
}

func (exchange *Exchange) GetEmailDomains(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getEmailDomains", params)
}

func (exchange *Exchange) GetMailboxByEmailAddress(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxByEmailAddress", params)
}

func (exchange *Exchange) GetMailboxes(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxes", params)
}

func (exchange *Exchange) GetMailboxInfo(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxInfo", params)
}

func (exchange *Exchange) GetMailboxStores(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxStores", params)
}

func (exchange *Exchange) GetUserMailbox(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getUserMailbox", params)
}

func (exchange *Exchange) GrantPublicFolderRoles(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.grantPublicFolderRoles", params)
}

func (exchange *Exchange) RevokePublicFolderRoles(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.revokePublicFolderRoles", params)
}

func (exchange *Exchange) ListPublicFolderRoles(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.listPublicFolderRoles", params)
}

func (exchange *Exchange) EnableSPLAFeature(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.enableSPLAFeature", params)
}

func (exchange *Exchange) DisableSPLAFeature(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.disableSPLAFeature", params)
}

func (exchange *Exchange) ModifyMailbox(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.modifyMailbox", params)
}

func (exchange *Exchange) MoveMailboxes(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.moveMailboxes", params)
}

func (exchange *Exchange) AddMailboxWithTemplate(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addMailboxWithTemplate", params)
}

func (exchange *Exchange) AssignMasterAccount(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.assignMasterAccount", params)
}

func (exchange *Exchange) ChangeMailboxTemplate(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.changeMailboxTemplate", params)
}

func (exchange *Exchange) GetMailboxTemplates(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxTemplates", params)
}

func (exchange *Exchange) ForceMailboxTemplatesUsing(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.forceMailboxTemplatesUsing", params)
}

func (exchange *Exchange) RemoveEmailAddresses(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removeEmailAddresses", params)
}

func (exchange *Exchange) RemoveMailbox(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removeMailbox", params)
}

func (exchange *Exchange) AddBlackBerry(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addBlackBerry", params)
}

func (exchange *Exchange) GetBlackBerryInfo(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getBlackBerryInfo", params)
}

func (exchange *Exchange) GetSubscriptionBlackBerryInfo(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getSubscriptionBlackBerryInfo", params)
}

func (exchange *Exchange) RemoveBlackBerry(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removeBlackBerry", params)
}

func (exchange *Exchange) SetBlackBerryActivationPassword(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.setBlackBerryActivationPassword", params)
}

func (exchange *Exchange) SetBlackBerryITPolicy(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.setBlackBerryITPolicy", params)
}

func (exchange *Exchange) WipeBlackBerry(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.wipeBlackBerry", params)
}

func (exchange *Exchange) AddDistributionList(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addDistributionList", params)
}

func (exchange *Exchange) AddDistributionListMembers(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addDistributionListMembers", params)
}

func (exchange *Exchange) ModifyDistributionList(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.modifyDistributionList", params)
}

func (exchange *Exchange) RemoveDistributionListMembers(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removeDistributionListMembers", params)
}

func (exchange *Exchange) RemoveDistributionList(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removeDistributionList", params)
}

func (exchange *Exchange) AddPublicFolder(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addPublicFolder", params)
}

func (exchange *Exchange) MailEnablePublicFolder(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.mailEnablePublicFolder", params)
}

func (exchange *Exchange) MailDisablePublicFolder(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.mailDisablePublicFolder", params)
}

func (exchange *Exchange) GetRootFolderName(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getRootFolderName", params)
}

func (exchange *Exchange) RemovePublicFolder(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removePublicFolder", params)
}

func (exchange *Exchange) AddContact(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addContact", params)
}

func (exchange *Exchange) GetContact(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getContact", params)
}

func (exchange *Exchange) ModifyContact(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.modifyContact", params)
}

func (exchange *Exchange) RemoveContact(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.removeContact", params)
}

func (exchange *Exchange) AddResourceMailbox(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.addResourceMailbox", params)
}

func (exchange *Exchange) GetResourceMailboxes(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getResourceMailboxes", params)
}

func (exchange *Exchange) ModifyResourceMailbox(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.modifyResourceMailbox", params)
}

func (exchange *Exchange) GetOutlookLicense(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getOutlookLicense", params)
}

func (exchange *Exchange) EnableOutlookLicense(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.enableOutlookLicense", params)
}

func (exchange *Exchange) DisableOutlookLicense(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.disableOutlookLicense", params)
}

func (exchange *Exchange) GetDistributionList(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getDistributionList", params)
}

func (exchange *Exchange) GetDistributionLists(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getDistributionLists", params)
}

func (exchange *Exchange) GetPublicFolders(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getPublicFolders", params)
}

func (exchange *Exchange) GetContacts(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getContacts", params)
}

func (exchange *Exchange) GetDeliveryPermissions(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getDeliveryPermissions", params)
}

func (exchange *Exchange) GetDeliveryPermissionsCandidates(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getDeliveryPermissionsCandidates", params)
}

func (exchange *Exchange) GrantDeliveryPermissions(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.grantDeliveryPermissions", params)
}

func (exchange *Exchange) RevokeDeliveryPermissions(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.revokeDeliveryPermissions", params)
}

func (exchange *Exchange) GetMailboxPermissions(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxPermissions", params)
}

func (exchange *Exchange) GetMailboxPermissionsCandidates(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMailboxPermissionsCandidates", params)
}

func (exchange *Exchange) GetMasterAccountInfo(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.getMasterAccountInfo", params)
}

func (exchange *Exchange) GrantMailboxPermissions(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.grantMailboxPermissions", params)
}

func (exchange *Exchange) RevokeMailboxPermissions(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.revokeMailboxPermissions", params)
}

func (exchange *Exchange) UnassignMasterAccount(params interface{}) (map[string]interface{}, error) {
	return exchange.SendRequest("pem.exchange.unassignMasterAccount", params)
}

func (oa *OA) FindHost(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.findHost", params)
}

func (oa *OA) GetAccountDomains(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountDomains", params)
}

func (oa *OA) GetAccountRoles(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountRoles", params)
}

func (oa *OA) GetAccountMemberRoles(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountMemberRoles", params)
}

func (oa *OA) GetAvailableSkins(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAvailableSkins", params)
}

func (oa *OA) GetBrandInfo(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getBrandInfo", params)
}

func (oa *OA) GetCustomerSubscriptionsResources(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getCustomerSubscriptionsResources", params)
}

func (oa *OA) GetDomainsForBrandCreation(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getDomainsForBrandCreation", params)
}

func (oa *OA) ImportBrandCertificate(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.importBrandCertificate", params)
}

func (oa *OA) GetDomainNameServers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getDomainNameServers", params)
}

func (oa *OA) GetHost(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getHost", params)
}

func (oa *OA) GetMemberSubscriptionRestrictions(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getMemberSubscriptionRestrictions", params)
}

func (oa *OA) GetResourceTypesByClass(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getResourceTypesByClass", params)
}

func (oa *OA) GetServiceTemplateList(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getServiceTemplateList", params)
}

func (oa *OA) GetVendorCustomers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getVendorCustomers", params)
}

func (oa *OA) GetCustomerSubscriptions(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getCustomerSubscriptions", params)
}

func (oa *OA) GetAccountInfo(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountInfo", params)
}

func (oa *OA) GetAccountMembers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountMembers", params)
}

func (oa *OA) GetAccountMemberByLogin(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountMemberByLogin", params)
}

func (oa *OA) GetAccountSubscriptions(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountSubscriptions", params)
}

func (oa *OA) GetAccountServiceUsers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getAccountServiceUsers", params)
}

func (oa *OA) GetDomainByName(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getDomainByName", params)
}

func (oa *OA) GetDomainSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getDomainSubscription", params)
}

func (oa *OA) GetDomainList(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getDomainList", params)
}

func (oa *OA) GetDNSRecords(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getDNSRecords", params)
}

func (oa *OA) GetFTPUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getFTPUser", params)
}

func (oa *OA) GetNameServers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getNameServers", params)
}

func (oa *OA) GetProvisioningAttributes(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getProvisioningAttributes", params)
}

func (oa *OA) GetRequestStatus(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getRequestStatus", params)
}

func (oa *OA) GetResourceUsage(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getResourceUsage", params)
}

func (oa *OA) GetResourceUsageForPeriod(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getResourceUsageForPeriod", params)
}

func (oa *OA) GetServiceTemplate(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getServiceTemplate", params)
}

func (oa *OA) GetSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getSubscription", params)
}

func (oa *OA) GetSubscriptionWebspaces(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getSubscriptionWebspaces", params)
}

func (oa *OA) GetUserByLogin(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getUserByLogin", params)
}

func (oa *OA) GetUserInfo(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getUserInfo", params)
}

func (oa *OA) GetUsers(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getUsers", params)
}

func (oa *OA) GetWebspacesList(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getWebspacesList", params)
}

func (iis *IIS) AddDomainMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.addDomainMapping", params)
}

func (iis *IIS) AddSystemMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.addSystemMapping", params)
}

func (iis *IIS) ApplySystemMappingsToHosts(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.applySystemMappingsToHosts", params)
}

func (iis *IIS) BulkAddDomainMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.bulkAddDomainMapping", params)
}

func (iis *IIS) DeleteSystemMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.deleteSystemMapping", params)
}

func (iis *IIS) GetDomainEngines(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getDomainEngines", params)
}

func (iis *IIS) GetDomainMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getDomainMapping", params)
}

func (iis *IIS) GetDomainMappings(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getDomainMappings", params)
}

func (iis *IIS) GetSystemEngines(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getSystemEngines", params)
}

func (iis *IIS) GetSystemMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getSystemMapping", params)
}

func (iis *IIS) GetSystemMappings(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getSystemMappings", params)
}

func (iis *IIS) GetWebspaceFeaturesInfo(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getWebspaceFeaturesInfo", params)
}

func (iis *IIS) ResetDomainMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.resetDomainMapping", params)
}

func (iis *IIS) SetDomainMappingEnabled(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.setDomainMappingEnabled", params)
}

func (iis *IIS) GetSubscriptionWebsites(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getSubscriptionWebsites", params)
}

func (iis *IIS) GetSubscriptionSharePointSites(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getSubscriptionSharePointSites", params)
}

func (iis *IIS) GetWebspacesInfo(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getWebspacesInfo", params)
}

func (iis *IIS) SetWebspaceFeaturesInfo(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.setWebspaceFeaturesInfo", params)
}

func (iis *IIS) UpdateDomainMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.updateDomainMapping", params)
}

func (iis *IIS) UpdateLimits(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.updateLimits", params)
}

func (iis *IIS) UpdatePHPExtensionsList(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.updatePHPExtensionsList", params)
}

func (iis *IIS) UpdateSystemEngine(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.updateSystemEngine", params)
}

func (iis *IIS) UpdateSystemMapping(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.updateSystemMapping", params)
}

func (iis *IIS) GrantAuthorizedWebAccess(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.grantAuthorizedWebAccess", params)
}

func (iis *IIS) RevokeAuthorizedWebAccess(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.revokeAuthorizedWebAccess", params)
}

func (iis *IIS) GetAuthorizedWebAccessDomainsForUser(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getAuthorizedWebAccessDomainsForUser", params)
}

func (iis *IIS) GetAuthorizedWebAccessUsersForDomain(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getAuthorizedWebAccessUsersForDomain", params)
}

func (iis *IIS) GetFTPAccessStatusForUser(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getFTPAccessStatusForUser", params)
}

func (iis *IIS) GetFTPAccessStatusesForDomain(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.getFTPAccessStatusesForDomain", params)
}

func (iis *IIS) GrantFTPAccessToWebsite(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.grantFTPAccessToWebsite", params)
}

func (iis *IIS) RevokeFTPAccessFromWebsite(params interface{}) (map[string]interface{}, error) {
	return iis.SendRequest("pem.iis.revokeFTPAccessFromWebsite", params)
}

func (oa *OA) ImportCertificate(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.importCertificate", params)
}

func (oa *OA) InstallPleskLicense(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.installPleskLicense", params)
}

func (oa *OA) MigrateSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.migrateSubscription", params)
}

func (oa *OA) ModifyUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.modifyUser", params)
}

func (oa *OA) MoveSubscriptions(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.moveSubscriptions", params)
}

func (mssql *Mssql) SetQuota(params interface{}) (map[string]interface{}, error) {
	return mssql.SendRequest("pem.mssql.setQuota", params)
}

func (mysql *Mysql) ApplyMySQLActivationParams(params interface{}) (map[string]interface{}, error) {
	return mysql.SendRequest("pem.mysql.applyMySQLActivationParams", params)
}

func (oa *OA) OrderSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.orderSubscription", params)
}

func (oa *OA) PromoteToReseller(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.promoteToReseller", params)
}

func (oa *OA) ProvisionSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.provisionSubscription", params)
}

func (proftpd *ProFTPD) CreateCustomFTPUser(params interface{}) (map[string]interface{}, error) {
	return proftpd.SendRequest("pem.ProFTPD.createCustomFTPUser", params)
}

func (proftpd *ProFTPD) DeleteCustomFTPUser(params interface{}) (map[string]interface{}, error) {
	return proftpd.SendRequest("pem.ProFTPD.deleteCustomFTPUser", params)
}

func (proftpd *ProFTPD) GetCustomFTPUsersList(params interface{}) (map[string]interface{}, error) {
	return proftpd.SendRequest("pem.ProFTPD.getCustomFTPUsersList", params)
}

func (oa *OA) RegisterExternalSystem(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.registerExternalSystem", params)
}

func (oa *OA) SetExternalSystemConfig(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setExternalSystemConfig", params)
}

func (oa *OA) SetMemberSubscriptionRestrictions(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setMemberSubscriptionRestrictions", params)
}

func (oa *OA) GetExternalSystemList(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.getExternalSystemList", params)
}

func (oa *OA) UnregisterExternalSystem(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.unregisterExternalSystem", params)
}

func (oa *OA) RegisterSharedNode(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.registerSharedNode", params)
}

func (oa *OA) RegisterWindowsNode(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.registerWindowsNode", params)
}

func (oa *OA) RemoveAccount(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeAccount", params)
}

func (oa *OA) RemoveAccountMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeAccountMember", params)
}

func (oa *OA) RemoveDNSHosting(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeDNSHosting", params)
}

func (oa *OA) RemoveDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeDomain", params)
}

func (oa *OA) RemoveLicense(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeLicense", params)
}

func (oa *OA) RemovePTRRecord(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removePTRRecord", params)
}

func (oa *OA) RemoveSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeSubscription", params)
}

func (oa *OA) RemoveUser(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.removeUser", params)
}

func (oa *OA) ResetResourceUsage(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.resetResourceUsage", params)
}

func (oa *OA) RevokeIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.revokeIPPool", params)
}

func (oa *OA) RevokeRolesFromMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.revokeRolesFromMember", params)
}

func (oa *OA) RevokePleskLicense(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.revokePleskLicense", params)
}

func (oa *OA) SetAccountAuthData(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setAccountAuthData", params)
}

func (oa *OA) SetAccountInfo(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setAccountInfo", params)
}

func (oa *OA) SetDomainRegistrarStatus(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setDomainRegistrarStatus", params)
}

func (oa *OA) SetFTPPassword(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setFTPPassword", params)
}

func (oa *OA) SetHostAttributes(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setHostAttributes", params)
}

func (oa *OA) SetHostReadyToProvide(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setHostReadyToProvide", params)
}

func (oa *OA) SetMemberInfo(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setMemberInfo", params)
}

func (oa *OA) SetMemberPassword(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setMemberPassword", params)
}

func (oa *OA) SetResourceTypeLimit(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setResourceTypeLimit", params)
}

func (oa *OA) SetResourceTypeLimits(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setResourceTypeLimits", params)
}

func (oa *OA) SetRTActParams(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setRTActParams", params)
}

func (oa *OA) SetRTAttributes(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setRTAttributes", params)
}

func (oa *OA) SetSTActivationParams(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setSTActivationParams", params)
}

func (oa *OA) SetSTRTLimits(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setSTRTLimits", params)
}

func (oa *OA) SetSubscriptionName(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setSubscriptionName", params)
}

func (oa *OA) SetSystemProperty(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.setSystemProperty", params)
}

func (sharepoint *Sharepoint) CreateSharePointSite(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.createSharePointSite", params)
}

func (spamAssassin *SpamAssassin) AddItems(params interface{}) (interface{}, error) {
	return spamAssassin.SendRequest("pem.spam_assassin.addItems", params)
}

func (spamAssassin *SpamAssassin) DeleteItems(params interface{}) (interface{}, error) {
	return spamAssassin.SendRequest("pem.spam_assassin.deleteItems", params)
}

func (spamAssassin *SpamAssassin) GetItems(params interface{}) (interface{}, error) {
	return spamAssassin.SendRequest("pem.spam_assassin.getItems", params)
}

func (tasks *Tasks) RescheduleTask(params interface{}) (interface{}, error) {
	return tasks.SendRequest("pem.tasks.rescheduleTask", params)
}

func (sharepoint *Sharepoint) DeleteSharePointSite(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.deleteSharePointSite", params)
}

func (sharepoint *Sharepoint) GetAvailableSharePointSiteTemplates(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.getAvailableSharePointSiteTemplates", params)
}

func (sharepoint *Sharepoint) AddSharePointUser(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.addSharePointUser", params)
}

func (sharepoint *Sharepoint) RemoveSharePointUser(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.removeSharePointUser", params)
}

func (sharepoint *Sharepoint) CreateSharePointSiteInSharedApplication(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.createSharePointSiteInSharedApplication", params)
}

func (sharepoint *Sharepoint) ModifySharePointSiteLimits(params interface{}) (interface{}, error) {
	return sharepoint.SendRequest("pem.sharepoint.modifySharePointSiteLimits", params)
}

func (oa *OA) UploadLicense(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.uploadLicense", params)
}

func (oa *OA) UnbindIPPool(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.unbindIPPool", params)
}

func (oa *OA) UnbrandDomain(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.unbrandDomain", params)
}

func (oa *OA) UnsetHostAttributes(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.unsetHostAttributes", params)
}

func (oa *OA) UpdateAccountAndAccountMember(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.updateAccountAndAccountMember", params)
}

func (oa *OA) UpgradeSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.upgradeSubscription", params)
}

func (virtuozzo *Virtuozzo) AttachVPS(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.attachVPS", params)
}

func (virtuozzo *Virtuozzo) ImportTemplate(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.importTemplate", params)
}

func (virtuozzo *Virtuozzo) InstallTemplate(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.installTemplate", params)
}

func (virtuozzo *Virtuozzo) InstallTemplates(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.installTemplates", params)
}

func (virtuozzo *Virtuozzo) RemoveTemplates(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.removeTemplates", params)
}

func (virtuozzo *Virtuozzo) UpdateEZTemplates(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.updateEZTemplates", params)
}

func (virtuozzo *Virtuozzo) RegisterHWNode(params interface{}) (interface{}, error) {
	return virtuozzo.SendRequest("pem.virtuozzo.registerHWNode", params)
}

func (webCluster *WebCluster) ChangeNFS(params interface{}) (interface{}, error) {
	return webCluster.SendRequest("pem.web_cluster.changeNFS", params)
}

func (txn *TXN) Begin(params interface{}) (map[string]interface{}, error) {
	return txn.SendRequest("txn.begin", params)
}

func (txn *TXN) Commit(params interface{}) (map[string]interface{}, error) {
	return txn.SendRequest("txn.begin", params)
}
func (txn *TXN) Rollback(params interface{}) (map[string]interface{}, error) {
	return txn.SendRequest("txn.begin", params)
}

func (oa *OA) BindUserToSubscription(params interface{}) (map[string]interface{}, error) {
	return oa.SendRequest("pem.bindUserToSubscription", params)
}