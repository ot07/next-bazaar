package product_domain

import (
	"context"

	"github.com/google/uuid"
	db "github.com/ot07/next-bazaar/db/sqlc"
)

type ProductService struct {
	store db.Store
}

func NewProductService(store db.Store) *ProductService {
	return &ProductService{
		store: store,
	}
}

func (s *ProductService) GetProduct(ctx context.Context, id uuid.UUID) (Product, error) {
	product, err := s.store.GetProduct(ctx, id)
	if err != nil {
		return Product{}, err
	}

	category, err := s.store.GetCategory(ctx, product.CategoryID)
	if err != nil {
		return Product{}, err
	}

	seller, err := s.store.GetUser(ctx, product.SellerID)
	if err != nil {
		return Product{}, err
	}

	return toProductDomain(product, category, seller), nil
}

type GetProductsServiceParams struct {
	PageID   int32
	PageSize int32
}

func (s *ProductService) GetProducts(ctx context.Context, params GetProductsServiceParams) ([]Product, error) {
	pageSize := params.PageSize
	pageID := params.PageID

	arg := db.ListProductsParams{
		Limit:  pageSize,
		Offset: (pageID - 1) * pageSize,
	}

	products, err := s.store.ListProducts(ctx, arg)
	if err != nil {
		return nil, err
	}

	categoryIDs := productsToCategoryIDs(products)
	categories, err := s.store.GetCategoriesByIDs(ctx, categoryIDs)
	if err != nil {
		return nil, err
	}

	categoriesMap := make(map[uuid.UUID]db.Category)
	for _, category := range categories {
		categoriesMap[category.ID] = category
	}

	sellersIDs := productsToSellersIDs(products)
	sellers, err := s.store.GetUsersByIDs(ctx, sellersIDs)
	if err != nil {
		return nil, err
	}

	sellersMap := make(map[uuid.UUID]db.User)
	for _, seller := range sellers {
		sellersMap[seller.ID] = seller
	}

	rsp := make([]Product, len(products))
	for i, product := range products {
		rsp[i] = toProductDomain(product, categoriesMap[product.CategoryID], sellersMap[product.SellerID])
	}

	return rsp, nil
}

func (s *ProductService) CountProducts(ctx context.Context) (int64, error) {
	return s.store.CountProducts(ctx)
}
