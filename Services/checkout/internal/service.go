package internal

import (
	"context"
	"fmt"
	"log"
	// "strconv"
	"time"

	"checkout/utils"

	addressClient "proto/address"
	cartClient "proto/cart"
	invetory "proto/inventory"
	order "proto/order"
)

type Service struct {
	cartService      cartClient.CartServiceClient
	inventoryService invetory.InventoryServiceClient
	orderService     order.OrderServiceClient
	addressService   addressClient.AddressServiceClient
}


func NewService(
	cartService cartClient.CartServiceClient,
	orderService order.OrderServiceClient,
	inventoryClient invetory.InventoryServiceClient,
	addressService  addressClient.AddressServiceClient,
) *Service {

	return &Service{
		cartService:      cartService,
		orderService:     orderService,
		inventoryService:   inventoryClient,

		addressService:   addressService,
	}
}

func (s *Service) loadCart(
	ctx context.Context,
	userID int64,
) (utils.CartResponse, error) {

	cartResponse, err := s.cartService.GetUserCart(
		ctx,
		&cartClient.GetUserCartRequest{
			UserId: userID,
		},
	)
	if err != nil {
		return utils.CartResponse{}, err
	}

	if len(cartResponse.Items) == 0 {
		return utils.CartResponse{}, fmt.Errorf("cart is empty")
	}

	items := make([]utils.CartItemResponse, 0, len(cartResponse.Items))

	for _, item := range cartResponse.Items {
		items = append(items, utils.CartItemResponse{
			ID:                 item.Id,
			ProductID:          item.ProductId,
			Quantity:           int(item.Quantity),
			UnitPrice:          float64(item.UnitPrice),
			TotalPrice:         float64(item.TotalPrice),
			ProductName:        item.ProductName,
			ProductDescription: item.ProductDescription,
			ProductImgUrl:      item.ProductImgUrl, // or ProducImgUrl if you keep the typo
		})
	}

	return utils.CartResponse{
		ID:         cartResponse.Id,
		UserID:     cartResponse.UserId,
		Items:      items,
		Subtotal:   float64(cartResponse.SubTotal),
		TotalItems: int(cartResponse.TotalItems),
	}, nil
}

func (s *Service) createPendingOrder(
	ctx context.Context,
	userID int64,
	addressID int64,
	cart utils.CartResponse,
) (
	orderID int64,
	orderNumber string,
	shipping float64,
	tax float64,
	discount float64,
	total float64,
	err error,
) {

	subtotal := cart.Subtotal

	shipping = calculateShipping(subtotal)
	tax = calculateTax(subtotal)
	discount = calculateDiscount(subtotal)

	total = subtotal + shipping + tax - discount

	orderNumber = generateOrderNumber()
	// parsedOrderNumber, err := strconv.ParseInt(orderNumber, 10, 64)
	// if err != nil {
	// 	log.Println("error to conver the orderNumber in int64",err)
	// 	return 0, "", 0, 0, 0, 0, err
	// }


	resp, err := s.orderService.CreateOrder(
		ctx,
		&order.CreatOrderRequest{
			UserId:userID,
			AddressId: addressID,
			OrderNumber: orderNumber,
			Subtotal: float32(subtotal),
			Shipping: float32(shipping),
			Tax: float32(tax),
			Discount: float32(discount),
			Total: float32(total),
		})

	if err != nil {
		return 0, "", 0, 0, 0, 0, err
	}

	return resp.OrderId, orderNumber, shipping, tax, discount, total, nil
}

func (s *Service) processItems(
	ctx context.Context,
	orderID int64,
	items []utils.CartItemResponse,
) error {

	for _, item := range items {

		resp, err := s.inventoryService.FindWarehouseForProduct(
			ctx,
			&invetory.FindWarehouseForProductRequest{
				ProductId: item.ProductID,
				Quantity:  int32(item.Quantity),
			},
		)
		if err != nil {
			log.Printf("Error to findwarehouse")
			return err
		}

		_, err = s.inventoryService.ReserveStock(
			ctx,
			&invetory.ReserveStockRequest{
				ProductId:   item.ProductID,
				WarehouseId: resp.WarehouseId,
				Quantity:    int64(item.Quantity),
			},
		)
		if err != nil {
			log.Printf("error to reesve Stock %v", err)
			return err
		}

		_,err = s.orderService.CreateOrderItem(
		ctx,
		&order.CreateOrderItemRequest{
			OrderId: orderID,
			ProductId:  item.ProductID,
			WarehouseId:  resp.WarehouseId,
		  	Quantity: int64(item.Quantity),
			UnitPrice: float32(item.UnitPrice),
			TotalPrice: float32(item.TotalPrice),
		},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Checkout(
	ctx context.Context,
	userID int64,
	body CheckoutBody,
) (CheckoutResponse, error) {
	log.Println("User id in checkout",userID)
	err := s.validateAddress(
		ctx,
		userID,
		body.AddressID,
	)
	if err != nil {
		return CheckoutResponse{}, err
	}

	cartResponse, err := s.loadCart(
		ctx,
		userID,
	)
	if err != nil {
		log.Println("error to load cart",err)
		return CheckoutResponse{}, err
	}

	// tx, err := s.orderService.BeginTx(ctx)
	if err != nil {
		return CheckoutResponse{}, err
	}

	// defer func() {
	// 	_ = s.orderService.Rollback(tx)
	// }()

	orderID,
		orderNumber,
		shipping,
		tax,
		discount,
		total,
		err := s.createPendingOrder(
		ctx,
		userID,
		body.AddressID,
		cartResponse,
	)

	if err != nil {
		log.Println("error to create order ",err)
		return CheckoutResponse{}, err
	}

	err = s.processItems(
		ctx,
		orderID,
		cartResponse.Items,
	)

	if err != nil {
		log.Println("error to process item ",err)
		return CheckoutResponse{}, err
	}

	// if err := s.orderService.Commit(tx); err != nil {
	// 	return CheckoutResponse{}, err
	// }

	return CheckoutResponse{
		OrderID:     orderID,
		OrderNumber: orderNumber,
		Status:      "pending",
		Subtotal:    cartResponse.Subtotal,
		Shipping:    shipping,
		Tax:         tax,
		Discount:    discount,
		Total:       total,
	}, nil
}

func generateOrderNumber() string {
	return fmt.Sprintf(
		"ORD-%d",
		time.Now().UnixNano(),
	)
}

func calculateShipping(subtotal float64) float64 {
	return 50
}

func calculateTax(subtotal float64) float64 {
	return subtotal * 0.18
}

func calculateDiscount(subtotal float64) float64 {
	log.Println("rnn the discoutn ", subtotal)
	return 0
}

func (s *Service) validateAddress(
	ctx context.Context,
	userID int64,
	addressID int64,
) error {

	_, err := s.addressService.GetAddress(
		ctx,
		&addressClient.GetAddressRequest{
			UserId: userID,
	 	AddressId: addressID,
		},
	)

	return err
}
