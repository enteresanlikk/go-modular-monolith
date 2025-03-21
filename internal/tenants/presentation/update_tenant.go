package tenantsPresentation

import (
	commonDomain "github.com/enteresanlikk/go-modular-monolith/internal/common/domain"
	commonPresentation "github.com/enteresanlikk/go-modular-monolith/internal/common/presentation"
	tenantsApplication "github.com/enteresanlikk/go-modular-monolith/internal/tenants/application"
	"github.com/gofiber/fiber/v2"
)

func (h *TenantsHandler) UpdateTenant(c *fiber.Ctx) error {
	var req tenantsApplication.UpdateTenantRequest
	if err := c.BodyParser(&req); err != nil {
		commonPresentation.JsonResponseWithStatus(c, fiber.StatusBadRequest, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	response, err := h.updateTenantService.UpdateTenant(&req)
	if err != nil {
		commonPresentation.JsonResponseWithStatus(c, fiber.StatusInternalServerError, commonDomain.ErrorResult(err.Error()))
		return nil
	}

	commonPresentation.JsonResponseWithStatus(c, fiber.StatusOK, commonDomain.SuccessDataResult("tenant_updated_successfully", response))
	return nil
}
