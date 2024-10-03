package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/opsdata-io/opsdata/backend/pkg/models"
	"github.com/opsdata-io/opsdata/backend/pkg/utils"
)

// CreateUploadLinkRequest defines the request structure for creating an upload link
type CreateUploadLinkRequest struct {
	Customer   string `json:"customer" binding:"required"`
	CaseNumber string `json:"caseNumber" binding:"required"`
	Subject    string `json:"subject" binding:"required"`
	Notes      string `json:"notes" binding:"required"`
}

// CreateUploadLink handles creating an upload link for a specific customer case
// @Summary Create an upload link
// @Description Creates an upload link for a specific customer case
// @Tags Uploads
// @Accept json
// @Produce json
// @Param request body CreateUploadLinkRequest true "Upload link details"
// @Security ApiKeyAuth
// @Success 201 {object} models.UploadLink
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/upload-link [post]
func CreateUploadLink(c *fiber.Ctx) error {
	var req CreateUploadLinkRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Cannot parse file"})
	}

	// Create an UploadLink object
	link := models.UploadLink{
		Customer:   req.Customer,
		CaseNumber: req.CaseNumber,
		Subject:    req.Subject,
		Notes:      req.Notes,
	}

	// Save the upload link to the database
	if err := utils.SaveUploadLink(&link); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]interface{}{"error": "Failed to save upload link"})
	}

	return c.Status(http.StatusCreated).JSON(link)
}
