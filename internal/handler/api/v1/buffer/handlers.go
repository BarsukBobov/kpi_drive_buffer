package buffer

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"kpi_drive_buffer/internal/handler/api/baseapi"
	"kpi_drive_buffer/internal/repository/kpidrive"
)

// @Summary saveFact
// @Tags buffer
// @Accept multipart/form-data
// @Produce json
// @Param form formData kpidrive.SaveFactForm true "form"
// @Success 200 {object} kpidrive.SaveFactResponse
// @Failure 404 {object} baseapi.ErrorResponse
// @router /buffer/save_fact [post]
func (h *router) saveFact(c *gin.Context) {
	var form kpidrive.SaveFactForm

	err := c.ShouldBindWith(&form, binding.Form)
	if err != nil {
		baseapi.Response404(c, err)
		return
	}
	resp, err := h.service.SaveFact(&form)
	if err != nil {
		baseapi.Response404(c, err)
		return
	}
	baseapi.Response(c, resp.StatusCode, resp.Body)
}
