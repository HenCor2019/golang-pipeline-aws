package models

type ShipengineShipping struct {
	Currency string
	Amount   float32
}
type Shipengine struct {
	Rate_type               string
	Carrier_id              string
	Shipping_amount         ShipengineShipping
	Insurance_amount        ShipengineShipping
	Confirmation_amount     ShipengineShipping
	Other_amount            ShipengineShipping
	Zone                    int
	Package_type            string
	Delivery_days           int
	Guaranteed_service      bool
	Estimated_delivery_date string
	Carrier_delivery_days   string
	Ship_date               string
	Negotiated_rate         bool
	Service_type            string
	Service_code            string
	Trackable               bool
	Carrier_code            string
	Carrier_nickname        string
	Carrier_friendly_name   string
	Validation_status       string
	Warning_messages        []string
	Error_messages          []string
}
