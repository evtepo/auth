package user

import (
	userDomain "github.com/evtepo/auth/internal/domain/user"
	"github.com/evtepo/auth/internal/usecase/user"
	"github.com/gin-gonic/gin"
	baseUtils "github.com/evtepo/auth/internal/utils"
)

type Handler struct {
	uc *user.Usecase
}

func (h *Handler) RegisterUserRoutes(r *gin.Engine) {
	group := r.Group("/api/v1/user")
	group.POST("/register", h.Register)
}

func NewHandler(uc *user.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": baseUtils.ValidateError(err), "success": false})
		return
	}

	email := userDomain.Email(req.Email)
	username := userDomain.Username(req.Username)
	password := userDomain.Password(req.Password)

	err := h.uc.Register(c.Request.Context(), email, username, password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "success": false})
	}

	return
}
