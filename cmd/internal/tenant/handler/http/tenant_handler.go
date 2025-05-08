package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tenant-management/cmd/internal/tenant/domain"
	_ "tenant-management/cmd/internal/tenant/domain"
	"tenant-management/cmd/internal/tenant/usecase"
	"tenant-management/pkg/idgen"
)

type TenantHandler struct {
	usecase *usecase.TenantUseCase
}

func NewTenantHandler(usecase *usecase.TenantUseCase) *TenantHandler {
	return &TenantHandler{usecase}
}

func (handler *TenantHandler) Save(c *gin.Context) {
	type req struct {
		Email        string `json:"email"`
		MobileNumber string `json:"mobile_number"`
		Names        string `json:"names"`
		UnitID       string `json:"unit_id"`
	}
	var t req
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tenant := domain.Tenant{
		ID:           idgen.GenerateID(),
		Email:        t.Email,
		MobileNumber: t.MobileNumber,
		Names:        t.Names,
		UnitID:       t.UnitID,
	}
	err := handler.usecase.CreateTenant(tenant)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"tenant": tenant})
}

// fetch tenant by id
func (handler *TenantHandler) FindTenantByTenantId(c *gin.Context) {
	type req struct {
		ID string `json:"id"`
	}
	var t req
	_ = c.ShouldBindJSON(&t)
	//return
	tenant, err := handler.usecase.FindTenantById(t.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tenant": tenant})
}
