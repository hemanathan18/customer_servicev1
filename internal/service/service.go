package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCustomerService, NewOffMailServiceService, NewOffPhoneServiceService,
	NewPermAddressServiceService, NewPersMailServiceService, NewPersPhoneServiceService, NewTempAddressServiceService, NewCustomerRelationServiceService)
