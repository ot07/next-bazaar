package api

import (
	"database/sql"
	"math"

	"github.com/gofiber/fiber/v2"
	product_domain "github.com/ot07/next-bazaar/api/domain/product"
	product_service "github.com/ot07/next-bazaar/api/service/product"
	"github.com/ot07/next-bazaar/api/validation"
)

type productHandler struct {
	service *product_service.ProductService
}

func newProductHandler(s *product_service.ProductService) *productHandler {
	return &productHandler{
		service: s,
	}
}

// @Summary      Get product
// @Tags         products
// @Param        id path string true "Product ID"
// @Success      200 {object} productResponse
// @Failure      400 {object} errorResponse
// @Failure      404 {object} errorResponse
// @Failure      500 {object} errorResponse
// @Router       /products/{id} [get]
func (h *productHandler) getProduct(c *fiber.Ctx) error {
	req := new(product_domain.GetProductRequest)
	if err := c.ParamsParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse(err))
	}

	product, err := h.service.GetProduct(c.Context(), req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(newErrorResponse(err))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse(err))
	}

	rsp := product_domain.NewProductResponse(*product)
	return c.Status(fiber.StatusOK).JSON(rsp)
}

// @Summary      List products
// @Tags         products
// @Param        query query listProductsRequest true "query"
// @Success      200 {object} listProductsResponse
// @Failure      400 {object} errorResponse
// @Failure      500 {object} errorResponse
// @Router       /products [get]
func (h *productHandler) listProducts(c *fiber.Ctx) error {
	req := new(product_domain.ListProductsRequest)
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse(err))
	}

	validate := validation.NewValidator()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newErrorResponse(err))
	}

	products, err := h.service.GetProducts(c.Context(), req.PageID, req.PageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse(err))
	}

	totalCount, err := h.service.CountAllProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(newErrorResponse(err))
	}

	pageCount := int64(math.Ceil(float64(totalCount) / float64(req.PageSize)))

	rsp := product_domain.ListProductsResponse{
		Meta: product_domain.ListProductsResponseMeta{
			PageID:     req.PageID,
			PageSize:   req.PageSize,
			PageCount:  pageCount,
			TotalCount: totalCount,
		},
		Data: product_domain.NewProductsResponse(products),
	}
	return c.Status(fiber.StatusOK).JSON(rsp)
}
