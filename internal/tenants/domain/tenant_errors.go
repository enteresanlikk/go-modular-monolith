package tenantsDomain

import "errors"

var (
	ErrTenantNotFound = errors.New("tenant_not_found")
	ErrTenantExists   = errors.New("tenant_already_exists")
)
