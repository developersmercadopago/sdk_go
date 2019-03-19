package entity

//API Preference Doc: https://api.mercadopago.com/checkout/preferences#options

//Items set the information
type Items struct {
	ID          string  `json:"id"`          //String(256)
	Title       string  `json:"title"`       //String(256)
	Description string  `json:"description"` //String(256)
	PictureURL  string  `json:"picture_url"` //String(600)
	CategoryID  string  `json:"category_id"` //String(256) Related: /item_categories/:category_id
	Quantity    int     `json:"quantity"`    //Integer
	CurrencyID  string  `json:"currency_id"` //String(3) Related: /currencies/:currency_id
	UnitPrice   float64 `json:"unit_price"`  //Float
}

//Phone set the buyer phone information
type Phone struct {
	AreaCode string `json:"area_code"` //String(256)
	Number   string `json:"number"`    //String(256)
}

//Identification set de personal buyer identification
type Identification struct {
	IdentificationType string `json:"type"`   //String(256)
	Number             string `json:"number"` //String(256)
}

//Address set the buyer personal address
type Address struct {
	ZipCode      string `json:"zip_code"`      //String(256)
	StreetName   string `json:"street_name"`   //String(256)
	StreetNumber int64  `json:"street_number"` //Integer
}

//Payer set the buyer Information
type Payer struct {
	Name           string          `json:"name"`           //String(256)
	Surname        string          `json:"surname"`        //String(256)
	Email          string          `json:"email"`          //String(256)
	Phone          *Phone          `json:"phone"`          //Object
	Identification *Identification `json:"identification"` //Object
	Address        *Address        `json:"address"`        //Object
	DateCreated    string          `json:"date_created"`   //Date(ISO_8601)
}

//PaymentMethods set up payment methods to be excluded from the payment process
type PaymentMethods struct {
	ExcludedPaymentMethods []string `json:"excluded_payment_methods"`  //Array String(256) | Related: /sites/:site_id/payment_methods/:payment_method_id
	ExcludedPaymentTypes   []string `json:"excluded_payment_types"`    //Array String(256) | Related: /sites/:site_id/payment_methods/:payment_method_id
	DefaultPaymentMethodID string   `json:"default_payment_method_id"` //String(256)
	Installments           int64    `json:"installments"`              //Integer
	DefaultInstallments    int64    `json:"default_installments"`      //Integer
}

//ReceiverAddress set the shipping address
type ReceiverAddress struct {
	ZipCode      string `json:"zip_code"`      //String(256)
	StreetName   string `json:"street_name"`   //String(256)
	StreetNumber int64  `json:"street_number"` //Integer
	Floor        string `json:"floor"`         //String(256)
	Apartment    string `json:"apartment"`     //String(256)
}

//Shipments set the shipments information
type Shipments struct {
	Mode                  string           `json:"mode"`                    //String | Values: custom:Custom shipping, me2: MercadoEnvíos, not_specified: Shipping mode not specified
	LocalPickup           bool             `json:"local_pickup"`            //Boolean
	Dimensions            string           `json:"dimensions"`              //String
	DefaultShippingMethod string           `json:"default_shipping_method"` //Integer
	FreeMethods           []string         `json:"free_methods"`            //Array(Integer) | Related resource/sites/:site/shipping_methods?shipping_mode=me2
	Cost                  float64          `json:"cost"`                    //Float
	FreeShipping          bool             `json:"free_shipping"`           //Boolean
	ReceiverAddress       *ReceiverAddress `json:"receiver_address"`        //Object
}

//BackURLs set the URLs to return to the sellers website
type BackURLs struct {
	Success string `json:"success"` //String(600)
	Pending string `json:"pending"` //String(600)
	Failure string `json:"failure"` //String(600)
}

//Preference allows you to set up, during the payment process, all the item information, any accepted means of payment and many other options
type Preference struct {
	Items               []*Items        `json:"items"`                //Object
	Payer               *Payer          `json:"payer"`                //Object
	PaymentMethods      *PaymentMethods `json:"payment_methods"`      //Object
	Shipments           *Shipments      `json:"shipments"`            //Object
	BackURLs            *BackURLs       `json:"back_urls"`            //Object
	NotificationURL     string          `json:"notification_url"`     //String(500)
	ID                  string          `json:"id"`                   //String
	InitPoint           string          `json:"init_point"`           //String
	SandboxInitPoint    string          `json:"sandbox_init_point"`   //String
	DateCreated         string          `json:"date_created"`         //Date(ISO_8601)
	OperationType       string          `json:"operation_type"`       //String Values | regular_payment: Normal payment, money_transfer: Money request
	AdditionalInfo      string          `json:"additional_info"`      //String(600)
	AutoReturn          string          `json:"auto_return"`          //Values | approved, all
	ExternalReference   string          `json:"external_reference"`   //String(256)
	Expires             bool            `json:"expires"`              //Boolean
	ExpirationDateFrom  string          `json:"expiration_date_from"` //Date(ISO_8601)
	ExpirationDateTo    int64           `json:"expiration_date_to"`   //Date(ISO_8601)
	CollectorID         string          `json:"collector_id"`         //Integer
	ClientID            string          `json:"client_id"`            //String
	Marketplace         string          `json:"marketplace"`          //String(256)
	MarketplaceFee      float64         `json:"marketplace_fee"`      //Float
	DifferentialPricing int64           `json:"differential_pricing"` //Integer Related resource /users/:user_id/differential_pricing/:id
	BinaryMode          bool            `json:"binary_mode"`          //Boolean
}

//LoadPreference make something
func (p *Preference) LoadPreference() {
	items := make([]*Items, 0)

	item := Items{
		Title:       "Dummy Item",
		Description: "Multicolor Item",
		Quantity:    1,
		CurrencyID:  "ARS",
		UnitPrice:   10.0,
	}

	items = append(items, &item)

	p.Items = items

}
