package internal

import (
	"context"
	"errors"
	"fmt"
	"database/sql"
	pb "cartservice/proto/productpb"
)

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than zero")
	ErrCartExpired     = errors.New("your cart has expired")
)


type Service struct {
	repo    *Repository
	product pb.ProductServiceClient
}

func NewService(repo *Repository,productService pb.ProductServiceClient) *Service {
	return &Service{
		repo:    repo,
		product: productService,
	}
}


func (s *Service) AddItemToCart(ctx context.Context, userID int64, productID int64, qty int) error {
	
	if qty <= 0 {
		return ErrInvalidQuantity
	}
	
	_,err:=s.product.IsExistProduct(ctx,&pb.IsExistProductRequest{
		Id: productID,
	})

	if err!=nil{
		return  err
	}
	
	cartID, err := s.repo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return fmt.Errorf("service failed to resolve cart: %w", err)
	}

	p,err:=s.product.GetProductPriceById(ctx,&pb.GetProductPriceByIdRequest{Id: productID})
	if err!=nil {
		return  err
	}
	err = s.repo.AddItem(ctx, cartID,p.Id , qty,float32(p.Price))
	if err != nil {
		return fmt.Errorf("service failed to add item: %w", err)
	}

	// err = s.repo.UpdateExpiration(ctx, cartID, s.cartTTL)
	// if err != nil {
		
	// 	fmt.Printf("warning: failed to update cart TTL for cart %d: %v\n", cartID, err)
	// }

	return nil
}

func (s *Service) GetUserCart(ctx context.Context, userID int64) (CartResponse, error) {
	// 1. Fetch the cart
	cart, err := s.repo.GetCart(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrCartNotFound) {
			return CartResponse{UserID: userID, Items: make([]CartItemResponse, 0)}, nil
		}
		return CartResponse{}, err
	}

	// expiration, err := s.repo.GetExpiration(ctx, cart.ID)
	// if err != nil {
	// 	// if errors.Is(err, ErrCartNotFound) {
	// 	// 	// Expiration record missing; set a default one
	// 	// 	_ = s.repo.UpdateExpiration(ctx, cart.ID, s.cartTTL)
	// 	// 	return cart, nil
	// 	// }
	// 	return CartResponse{}, err
	// }

	var ids []int64

	for _,c:= range cart.Items{
		ids = append(ids, int64(c.ProductID))
	}


	res,err:=s.product.GetProductsByIds(ctx,&pb.GetProductsByIdsRequest{
		Ids: ids,
	})

	if err!=nil {
		return  CartResponse{},err
	}


	productMap := make(map[int64]*pb.Product)

	for _, p := range res.Products {
		productMap[p.Id] = p
	}

	for i := range cart.Items {
    if p, ok := productMap[cart.Items[i].ProductID]; ok {
        cart.Items[i].ProductName = p.Name
        cart.Items[i].ProductDescription = p.Description
        cart.Items[i].ProducImgUrl = p.ImgUrl
    }
}

	
	// if time.Now().After(expiration.ExpiresAt) {
	// 	_ = s.repo.ClearCart(ctx, cart.ID)
	// 	_ = s.repo.DeleteExpiration(ctx, cart.ID) 
	// 	return CartResponse{}, ErrCartExpired
	// }

	return cart, nil
}


func (s *Service) UpdateItemQuantity(ctx context.Context, userID int64, cartItemID int64, qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}

	_, err := s.repo.GetCart(ctx, userID)
	if err != nil {
		return err
	}

	
	err = s.repo.UpdateQuantity(ctx, cartItemID, qty)
	if err != nil {
		return err
	}


	// _ = extendCartTTL(ctx, s, cart.ID)

	return nil
}


func (s *Service) RemoveItemFromCart(ctx context.Context, userID int64, cartItemID int64) error {
	_, err := s.repo.GetCart(ctx, userID)
	if err != nil {
		return err
	}

	err = s.repo.RemoveItem(ctx, cartItemID)
	if err != nil {
		return err
	}

	//  _ = extendCartTTL(ctx, s, cart.ID)
	return nil
}

func (s *Service) EmptyCartTx(ctx context.Context,tx *sql.Tx, userID int64) error {
	cart, err := s.repo.GetCart(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrCartNotFound) {
			return nil
		}
		return err
	}

	err = s.repo.ClearCartTx(ctx,tx, cart.ID)
	if err != nil {
		return err
	}

	// _ = s.repo.DeleteExpiration(ctx, cart.ID)
	return nil
}


func (s *Service) EmptyCart (ctx context.Context ,userID int64)error{
	cart, err := s.repo.GetCart(ctx ,userID)
	if err != nil {
		if errors.Is(err, ErrCartNotFound) {
			return nil
		}
		return err
	}

	err = s.repo.ClearCart( cart.ID)
	if err != nil {
		return err
	}

	// _ = s.repo.DeleteExpiration(ctx, cart.ID)
	return nil
}

// func extendCartTTL(ctx context.Context, s *Service, cartID int64) error {
// 	return s.repo.UpdateExpiration(ctx, cartID, s.cartTTL)
// }