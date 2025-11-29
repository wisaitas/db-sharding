package create

import "github.com/gofiber/fiber/v3"

type Handler struct {
	service Service
}

func NewHandler(
	service Service,
) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateExample(c fiber.Ctx) error {
	var req Request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return h.service.CreateExample(c, req)
}
