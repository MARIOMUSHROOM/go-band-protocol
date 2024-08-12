package http

import (
	"band_protocol_go/internal/ports"
	"band_protocol_go/pkg/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ActionHandler struct {
	service ports.ActionService
	client  *service.APIService
}

func NewActionHandler(service ports.ActionService, client *service.APIService) *ActionHandler {
	return &ActionHandler{service: service, client: client}
}

func (h *ActionHandler) GetBossBaby(c *fiber.Ctx) error {
	text := c.Params("value")

	action := h.service.CheckRevenge(text)
	return c.JSON(action)
}

func (h *ActionHandler) SupermanChicken(c *fiber.Ctx) error {
	var req struct {
		N        int   `json:"n"`
		K        int   `json:"k"`
		Position []int `json:"position"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid request")
	}
	action := h.service.MaxChickensProtected(req.N, req.K, req.Position)
	return c.JSON(action)
}

func (h *ActionHandler) Transaction(c *fiber.Ctx) error {
	var req struct {
		Symbol    string  `json:"symbol"`
		Price     float64 `json:"price"`
		Timestamp float64 `json:"timestamp"`
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(http.StatusBadRequest, "Invalid request")
	}

	postData := make(map[string]interface{})
	postData["symbol"] = req.Symbol
	postData["price"] = req.Price
	postData["timestamp"] = req.Timestamp
	responseStatus, err := h.client.PostTransaction("/broadcast", postData)
	if err != nil {
		return fiber.NewError(http.StatusBadGateway, err.Error())
	}
	if responseStatus == nil {
		return fiber.NewError(http.StatusBadGateway, err.Error())
	}

	endpoint2 := "/check/" + responseStatus.TxHash
	response, err := h.client.GetData(endpoint2)
	if err != nil {
		return fiber.NewError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(response)
}
