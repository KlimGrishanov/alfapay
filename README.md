# alfapay

Go client library for Alfa Payments API (Alfa-Bank Payment Gateway).

## Installation

```bash
go get github.com/alfapay
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/alfapay"
)

func main() {
    // Create a new client
    client := alfapay.NewClient("your-username", "your-password")

    ctx := context.Background()

    // Register an order
    resp, err := client.Orders.Register(ctx, &alfapay.RegisterOrderRequest{
        OrderNumber: "ORDER-12345",
        Amount:      100000, // 1000.00 RUB (in kopecks)
        ReturnURL:   "https://your-site.com/success",
        Description: "Payment for order #12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Redirect customer to: %s\n", resp.FormURL)
}
```

## Features

- **Order Management**: Register single-stage and two-stage (pre-auth) payments
- **Payment Operations**: Deposit, refund, reverse payments
- **Card Bindings**: Save and manage customer cards for recurring payments
- **Recurrent Payments**: Process subscription and recurring payments
- **Mobile Payments**: Apple Pay, Google Pay, Samsung Pay, MIR Pay, Yandex Pay
- **SBP (Fast Payment System)**: QR code payments, B2B and B2C transfers
- **3D Secure**: Support for 3DS authentication flow

## Configuration Options

```go
// Custom base URL (production)
client := alfapay.NewClient(
    "username",
    "password",
    alfapay.WithBaseURL("https://pay.alfabank.ru/payment"),
)

// Custom timeout
client := alfapay.NewClient(
    "username",
    "password",
    alfapay.WithTimeout(60 * time.Second),
)

// Custom HTTP client
httpClient := &http.Client{
    Transport: &http.Transport{
        MaxIdleConns: 100,
    },
}
client := alfapay.NewClient(
    "username",
    "password",
    alfapay.WithHTTPClient(httpClient),
)
```

## API Reference

### Orders

```go
// Single-stage payment (immediate charge)
client.Orders.Register(ctx, &alfapay.RegisterOrderRequest{...})

// Two-stage payment (pre-authorization)
client.Orders.RegisterPreAuth(ctx, &alfapay.RegisterOrderRequest{...})

// Cancel unpaid order
client.Orders.Decline(ctx, &alfapay.DeclineRequest{...})

// Add parameters to order
client.Orders.AddParams(ctx, &alfapay.AddParamsRequest{...})
```

### Status

```go
// Get extended order status
client.Status.GetExtended(ctx, &alfapay.GetOrderStatusRequest{...})

// Quick status check by order ID
client.Status.GetByOrderID(ctx, "order-id")

// Quick status check by order number
client.Status.GetByOrderNumber(ctx, "ORDER-123")

// Get orders for date range
client.Status.GetLastOrders(ctx, &alfapay.GetLastOrdersRequest{...})

// Check 3DS enrollment
client.Status.VerifyEnrollment(ctx, &alfapay.VerifyEnrollmentRequest{...})
```

### Payments

```go
// Complete pre-authorized payment
client.Payments.Deposit(ctx, &alfapay.DepositRequest{...})

// Cancel authorized payment (before settlement)
client.Payments.Reverse(ctx, &alfapay.ReverseRequest{...})

// Pay using saved card
client.Payments.PayWithBinding(ctx, &alfapay.PaymentOrderBindingRequest{...})

// Instant payment (register + pay)
client.Payments.Instant(ctx, &alfapay.InstantPaymentRequest{...})

// Recurrent/subscription payment
client.Payments.Recurrent(ctx, &alfapay.RecurrentPaymentRequest{...})

// Complete 3DS authentication
client.Payments.Finish3DS(ctx, &alfapay.Finish3DSPaymentRequest{...})
```

### Refunds

```go
// Refund payment (partial or full)
client.Refunds.Refund(ctx, &alfapay.RefundRequest{...})

// Instant refund
client.Refunds.InstantRefund(ctx, "order-id", 50000)
```

### Bindings (Saved Cards)

```go
// Get customer's saved cards
client.Bindings.GetBindings(ctx, &alfapay.GetBindingsRequest{...})

// Activate binding
client.Bindings.Activate(ctx, &alfapay.BindingRequest{...})

// Deactivate binding
client.Bindings.Deactivate(ctx, &alfapay.UnbindRequest{...})

// Extend expiry date
client.Bindings.Extend(ctx, &alfapay.ExtendBindingRequest{...})
```

### Mobile Payments

```go
// Apple Pay
client.ApplePay.Payment(ctx, &alfapay.ApplePayPaymentRequest{...})

// Google Pay
client.GooglePay.Payment(ctx, &alfapay.GooglePayRequest{...})

// Samsung Pay
client.SamsungPay.Payment(ctx, &alfapay.SamsungPayPaymentRequest{...})

// MIR Pay
client.MirPay.Payment(ctx, &alfapay.MirPayPaymentRequest{...})

// Yandex Pay
client.YandexPay.Payment(ctx, &alfapay.YandexPayRequest{...})
```

### SBP (Fast Payment System)

```go
// Get QR code for payment
client.SBP.GetQR(ctx, &alfapay.SBPGetQRRequest{...})

// Check QR payment status
client.SBP.GetQRStatus(ctx, "order-id")

// Cancel QR payment
client.SBP.RejectQR(ctx, "order-id")

// Create SBP binding
client.SBP.Bind(ctx, "order-id")

// B2B payment
client.SBP.B2BPerform(ctx, &alfapay.SBPB2BPerformRequest{...})

// B2C payout
client.SBP.B2CPerformPayout(ctx, &alfapay.SBPB2CPayoutRequest{...})
```

## Order Status Values

| Value | Description |
|-------|-------------|
| 0 | Order registered, not paid |
| 1 | Pre-authorized amount held |
| 2 | Full authorization completed |
| 3 | Authorization cancelled |
| 4 | Refund completed |
| 5 | ACS authorization initiated |
| 6 | Authorization declined |

## Error Handling

```go
resp, err := client.Orders.Register(ctx, req)
if err != nil {
    // Network or HTTP error
    if apiErr, ok := err.(*alfapay.APIError); ok {
        fmt.Printf("API error (status %d): %s\n", apiErr.StatusCode, apiErr.Message)
    }
    return err
}

// Check business logic error
if !resp.IsSuccess() {
    fmt.Printf("Error code: %s, message: %s\n", resp.ErrorCode, resp.ErrorMessage)
    return errors.New(resp.ErrorMessage)
}

// Success
fmt.Printf("Order ID: %s\n", resp.OrderID)
```

## Amount Format

All amounts are specified in the smallest currency unit (kopecks for RUB):
- 100.00 RUB = 10000 kopecks
- 1,234.56 RUB = 123456 kopecks

## License

MIT License
