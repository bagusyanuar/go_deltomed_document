package builder

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BuildGroup(route *gin.Engine, db *gorm.DB) {
	AdminBuilder(route, db)
	SupportBuilder(route, db)
}
