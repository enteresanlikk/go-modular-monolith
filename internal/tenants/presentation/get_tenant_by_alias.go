package tenantsPresentation

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	tenantsApplication "github.com/enteresanlikk/go-modular-monolith/internal/tenants/application"
	"github.com/gofiber/fiber/v2"
)

func (h *TenantsHandler) GetTenantByAlias(c *fiber.Ctx) {
	var req tenantsApplication.GetTenantByAliasRequest
	if err := c.BodyParser(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(c, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return
	}

	response, err := h.getTenantByAliasService.GetTenantByAlias(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(c, fiber.StatusInternalServerError, commonDomain.ErrorResult(err.Error()))
		return
	}

	commonPresentation.JsonResponseWithStatus(c, fiber.StatusOK, commonDomain.SuccessDataResult("tenant_fetched_successfully", response))
}
