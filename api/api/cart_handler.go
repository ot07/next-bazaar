package api

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
	cart_domain "github.com/ot07/next-bazaar/api/domain/cart"
	cart_service "github.com/ot07/next-bazaar/api/service/cart"
	"github.com/ot07/next-bazaar/api/validation"
	db "github.com/ot07/next-bazaar/db/sqlc"
)

type cartProductResponse struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description db.NullString `json:"description" swaggertype:"string"`
	Price       string        `json:"price"`
	Quantity    int32         `json:"quantity"`
	Subtotal    string        `json:"subtotal"`
}

func newCartProductResponse(cartProduct cart_domain.CartProduct) cartProductResponse {
	return cartProductResponse{
		ID:          cartProduct.ID,
		Name:        cartProduct.Name,
		Description: db.NullString{NullString: cartProduct.Description},
		Price:       cartProduct.Price,
		Quantity:    cartProduct.Quantity,
		Subtotal:    cartProduct.Subtotal,
	}
}

type cartResponse []cartProductResponse

func newCartResponse(products []cart_domain.CartProduct) cartResponse {
	rsp := make(cartResponse, 0, len(products))
	for _, product := range products {
		rsp = append(rsp, newCartProductResponse(product))
	}
	return rsp
}

type getProductsRequest struct {
	ID uuid.UUID `params:"user_id"`
}

type cartHandler struct {
	service *cart_service.CartService
}

func newCartHandler(s *cart_service.CartService) *cartHandler {
	return &cartHandler{
		service: s,
	}
}

// @Summary      Get cart
// @Tags         cart
// @Param        userId path string true "User ID"
// @Success      200 {object} productResponse
// @Failure      400 {object} errorResponse
// @Failure      404 {object} errorResponse
// @Failure      500 {object} errorResponse
// @Router       /cart-products/{userId} [get]
func (h *cartHandler) getCart(c *fiber.Ctx) error {
	req := new(getProductsRequest)
	if err := c.ParamsParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse(err))
	}

	cartProducts, err := h.service.GetProductsByUserID(c.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(newErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse(err))
	}

	rsp := newCartResponse(cartProducts)
	return c.Status(fiber.StatusOK).JSON(rsp)
}

type addProductRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int32     `json:"quantity" validate:"required,min=1"`
}

// @Summary      Add product to cart
// @Tags         cart
// @Param        body body addProductRequest true "Cart product object"
// @Success      200 {object} messageResponse
// @Failure      400 {object} errorResponse
// @Failure      403 {object} errorResponse
// @Failure      500 {object} errorResponse
// @Router       /cart-products [post]
func (h *cartHandler) addProduct(c *fiber.Ctx) error {
	session, ok := c.Locals(ctxLocalSessionKey).(db.Session)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(newErrorResponse(
			fmt.Errorf("session token not found"),
		))
	}

	req := new(addProductRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse(err))
	}

	validate := validation.NewValidator()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse(err))
	}

	err := h.service.AddProduct(c.Context(), cart_service.NewAddProductParams(
		session.UserID,
		req.ProductID,
		req.Quantity,
	))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.Status(fiber.StatusForbidden).JSON(newErrorResponse(err))
			}
		}
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse(err))
	}

	rsp := newMessageResponse("Cart product added successfully")
	return c.Status(fiber.StatusOK).JSON(rsp)
}