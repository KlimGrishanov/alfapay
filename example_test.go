package alfapay_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/alfapay"
)

func Example_registerOrder() {
	// Create a new client with your credentials
	client := alfapay.NewClient("your-username", "your-password")

	ctx := context.Background()

	// Register a new order
	resp, err := client.Orders.Register(ctx, &alfapay.RegisterOrderRequest{
		OrderNumber: "ORDER-12345",
		Amount:      100000, // 1000.00 RUB (amount in kopecks)
		ReturnURL:   "https://your-site.com/payment/success",
		FailURL:     "https://your-site.com/payment/fail",
		Description: "Payment for order #12345",
		Language:    "ru",
		Email:       "customer@example.com",
		Phone:       "+79001234567",
	})
	if err != nil {
		log.Fatalf("Failed to register order: %v", err)
	}

	if resp.IsSuccess() {
		fmt.Printf("Order registered successfully!\n")
		fmt.Printf("Order ID: %s\n", resp.OrderID)
		fmt.Printf("Payment form URL: %s\n", resp.FormURL)
	} else {
		fmt.Printf("Error: %s - %s\n", resp.ErrorCode, resp.ErrorMessage)
	}
}

func Example_registerPreAuthOrder() {
	// Create a new client for two-stage payments (pre-authorization)
	client := alfapay.NewClient("your-username", "your-password")

	ctx := context.Background()

	// Register a pre-authorized order (funds will be held, not charged)
	resp, err := client.Orders.RegisterPreAuth(ctx, &alfapay.RegisterOrderRequest{
		OrderNumber: "ORDER-12346",
		Amount:      50000, // 500.00 RUB
		ReturnURL:   "https://your-site.com/payment/success",
		Description: "Pre-authorization for hotel booking",
	})
	if err != nil {
		log.Fatalf("Failed to register pre-auth order: %v", err)
	}

	fmt.Printf("Pre-auth order ID: %s\n", resp.OrderID)
}

func Example_getOrderStatus() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Get order status by order ID
	status, err := client.Status.GetByOrderID(ctx, "your-order-id")
	if err != nil {
		log.Fatalf("Failed to get order status: %v", err)
	}

	fmt.Printf("Order Number: %s\n", status.OrderNumber)
	fmt.Printf("Order Status: %d\n", status.OrderStatus)
	fmt.Printf("Amount: %d kopecks\n", status.Amount)

	if status.CardAuthInfo != nil {
		fmt.Printf("Card: %s\n", status.CardAuthInfo.MaskedPan)
	}
}

func Example_depositPayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Complete a pre-authorized payment
	resp, err := client.Payments.Deposit(ctx, &alfapay.DepositRequest{
		OrderID: "your-order-id",
		Amount:  45000, // Can be less than pre-authorized amount
	})
	if err != nil {
		log.Fatalf("Failed to deposit: %v", err)
	}

	if resp.IsSuccess() {
		fmt.Println("Payment completed successfully!")
	}
}

func Example_refundPayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Refund a payment (partial or full)
	resp, err := client.Refunds.Refund(ctx, &alfapay.RefundRequest{
		OrderID: "your-order-id",
		Amount:  25000, // Partial refund of 250.00 RUB
	})
	if err != nil {
		log.Fatalf("Failed to refund: %v", err)
	}

	if resp.IsSuccess() {
		fmt.Println("Refund processed successfully!")
	}
}

func Example_reversePayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Cancel a pre-authorized payment before settlement
	resp, err := client.Payments.Reverse(ctx, &alfapay.ReverseRequest{
		OrderID: "your-order-id",
	})
	if err != nil {
		log.Fatalf("Failed to reverse: %v", err)
	}

	if resp.IsSuccess() {
		fmt.Println("Payment reversed successfully!")
	}
}

func Example_recurrentPayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Perform a recurrent (subscription) payment using saved card binding
	resp, err := client.Payments.Recurrent(ctx, &alfapay.RecurrentPaymentRequest{
		OrderNumber: "SUB-12345-202311",
		BindingID:   "saved-binding-id",
		Amount:      99900, // 999.00 RUB monthly subscription
		Description: "Monthly subscription payment",
	})
	if err != nil {
		log.Fatalf("Failed to process recurrent payment: %v", err)
	}

	if resp.Success {
		fmt.Printf("Recurrent payment successful! Order ID: %s\n", resp.Data.OrderID)
	} else {
		fmt.Printf("Recurrent payment failed: %s\n", resp.Error.Message)
	}
}

func Example_getBindings() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Get all saved card bindings for a customer
	resp, err := client.Bindings.GetBindings(ctx, &alfapay.GetBindingsRequest{
		ClientID:    "customer-123",
		ShowExpired: "false",
	})
	if err != nil {
		log.Fatalf("Failed to get bindings: %v", err)
	}

	for _, binding := range resp.Bindings {
		fmt.Printf("Card: %s (expires: %s)\n", binding.MaskedPan, binding.ExpiryDate)
	}
}

func Example_applePayPayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Process Apple Pay payment
	resp, err := client.ApplePay.Payment(ctx, &alfapay.ApplePayPaymentRequest{
		Merchant:     "your-merchant-name",
		OrderNumber:  "ORDER-AP-001",
		PaymentToken: "base64-encoded-apple-pay-token",
		Amount:       150000,
		Email:        "customer@example.com",
	})
	if err != nil {
		log.Fatalf("Apple Pay payment failed: %v", err)
	}

	if resp.Success {
		fmt.Printf("Apple Pay payment successful! Order ID: %s\n", resp.Data.OrderID)
	}
}

func Example_googlePayPayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Process Google Pay payment
	resp, err := client.GooglePay.Payment(ctx, &alfapay.GooglePayRequest{
		Merchant:     "your-merchant-name",
		OrderNumber:  "ORDER-GP-001",
		PaymentToken: "base64-encoded-google-pay-token",
		Amount:       200000,
		IP:           "192.168.1.1",
		ReturnURL:    "https://your-site.com/payment/success",
	})
	if err != nil {
		log.Fatalf("Google Pay payment failed: %v", err)
	}

	if resp.Success {
		fmt.Printf("Google Pay payment successful! Order ID: %s\n", resp.Data.OrderID)
	}
}

func Example_sbpQRPayment() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// First, register an order
	orderResp, err := client.Orders.Register(ctx, &alfapay.RegisterOrderRequest{
		OrderNumber: "SBP-ORDER-001",
		Amount:      50000,
		ReturnURL:   "https://your-site.com/payment/success",
	})
	if err != nil {
		log.Fatalf("Failed to register order: %v", err)
	}

	// Get QR code for SBP payment
	qrResp, err := client.SBP.GetQR(ctx, &alfapay.SBPGetQRRequest{
		MDOrder:  orderResp.OrderID,
		QRHeight: 300,
		QRWidth:  300,
		QRFormat: "image",
	})
	if err != nil {
		log.Fatalf("Failed to get SBP QR: %v", err)
	}

	fmt.Printf("QR URL: %s\n", qrResp.QRURL)
	// Display QR image (base64 encoded in qrResp.QRImage)
}

func Example_customTimeout() {
	// Create client with custom timeout
	client := alfapay.NewClient(
		"your-username",
		"your-password",
		alfapay.WithTimeout(60*time.Second),
	)

	_ = client
}

func Example_customBaseURL() {
	// Create client with production URL
	client := alfapay.NewClient(
		"your-username",
		"your-password",
		alfapay.WithBaseURL("https://pay.alfabank.ru/payment"),
	)

	_ = client
}

func Example_orderWithCart() {
	client := alfapay.NewClient("your-username", "your-password")
	ctx := context.Background()

	// Create order with detailed cart items (for fiscal operations)
	taxType := alfapay.TaxTypeVAT20
	resp, err := client.Orders.Register(ctx, &alfapay.RegisterOrderRequest{
		OrderNumber: "ORDER-CART-001",
		Amount:      250000, // 2500.00 RUB
		ReturnURL:   "https://your-site.com/payment/success",
		Email:       "customer@example.com",
		OrderBundle: &alfapay.OrderBundle{
			CartItems: &alfapay.CartItems{
				Items: []alfapay.Item{
					{
						PositionID: 1,
						Name:       "Product A",
						Quantity: &alfapay.Quantity{
							Value:   2,
							Measure: "pcs",
						},
						ItemAmount: 100000, // 1000.00 RUB total
						ItemPrice:  50000,  // 500.00 RUB per item
						Tax: &alfapay.Tax{
							TaxType: taxType,
						},
					},
					{
						PositionID: 2,
						Name:       "Product B",
						Quantity: &alfapay.Quantity{
							Value:   1,
							Measure: "pcs",
						},
						ItemAmount: 150000, // 1500.00 RUB
						ItemPrice:  150000,
						Tax: &alfapay.Tax{
							TaxType: taxType,
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to register order: %v", err)
	}

	fmt.Printf("Order with cart registered: %s\n", resp.OrderID)
}
