package handler

import (
	"context"
	"fmt"

	paymentStatus "github.com/harryhcs/conservio_web/view/payStatus"
	"github.com/hasura/go-graphql-client"
	"github.com/labstack/echo/v4"
)

type PayStatusHandler struct {
	Client *graphql.Client
}

var m struct {
	DirectBookingPayments []struct {
		Email            string
		Amount           int
		Status           string
		StaySummary      string
		LastName         string
		FirstName        string
		CheckInDate      string
		CheckOutDate     string
		BookingReference string
		PropertyName     string
	} `graphql:"directBookingPayments(input: {id: $payID, propertyID: 0, bookingReference: \"\"})"`
}

func (h PayStatusHandler) HandleShowPayStatus(c echo.Context) error {

	id := c.Param("payID") // use this to get payment from API
	payID := graphql.ID(id)

	variables := map[string]interface{}{
		"payID": payID,
	}

	err := h.Client.Query(context.Background(), &m, variables)
	if err != nil {
		fmt.Printf("GQL Error: %s", err)
	} else {
		fmt.Printf("Found payment: %v\n", m.DirectBookingPayments[0].Email)
	}

	payment := m.DirectBookingPayments[0]
	return render(c, paymentStatus.Show(payment))
}
