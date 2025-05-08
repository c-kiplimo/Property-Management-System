package routes

import (
	"github.com/gin-gonic/gin"
	"tenant-management/cmd/internal/tenant/handler/http"
	"tenant-management/cmd/internal/tenant/usecase"
)

func RegisterTenantRoutes(r *gin.Engine, tenantUC *usecase.TenantUseCase) {
	tenantHandler := http.NewTenantHandler(tenantUC)
	tenantGroup := r.Group("/tenants")
	{
		tenantGroup.POST("", tenantHandler.Save)
		tenantGroup.GET("/:tenantId", tenantHandler.FindTenantByTenantId)
	}
}
