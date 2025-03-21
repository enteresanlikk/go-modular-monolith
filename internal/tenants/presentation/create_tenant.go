package tenantsPresentation

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	tenantsApplication "github.com/enteresanlikk/go-modular-monolith/internal/tenants/application"
	"github.com/goccy/go-json"

	"github.com/valyala/fasthttp"
)

func (h *TenantsHandler) CreateTenant(c *fasthttp.RequestCtx) {
	var req tenantsApplication.CreateTenantRequest
	if err := json.Unmarshal(c.PostBody(), &req); err != nil {
		commonPresentation.JsonResponseWithStatus(c, fasthttp.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := h.createTenantService.CreateTenant(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(c, fasthttp.StatusInternalServerError, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(c, fasthttp.StatusOK, commonDomain.SuccessDataResult("tenant_created_successfully", response))
}
