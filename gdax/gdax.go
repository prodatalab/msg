package gdax

//go:generate msg

// Channel blah
type Channel struct {
	Name       string   `json:"name" msg:"name"`
	ProductIDs []string `json:"product_ids" msg:"product_ids"`
}

// SubscribeAuth blah
type SubscribeAuth struct {
	Channels   []string `json:"channels" msg:"channels"`
	Key        string   `json:"key" msg:"key"`
	Passphrase string   `json:"passphrase" msg:"passphrase"`
	ProductIds []string `json:"product_ids" msg:"product_ids"`
	Signature  string   `json:"signature" msg:"signature"`
	Timestamp  string   `json:"timestamp" msg:"timestamp"`
	Type       string   `json:"type" msg:"type"`
}

// SubscribeReq blah
type SubscribeReq struct {
	Type       string    `json:"type" msg:"type"`
	ProductIds []string  `json:"product_ids" msg:"product_ids"`
	Channels   []Channel `json:"channels" msg:"channels"`
}

// SubscribeRep blah
type SubscribeRep struct {
	Type     string    `json:"type" msg:"type"`
	Channels []Channel `json:"channels" msg:"channels"`
}

// UnSubscribeReq blah
type UnSubscribeReq struct {
	Type       string   `json:"type" msg:"type"`
	ProductIds []string `json:"product_ids" msg:"product_ids"`
	Channels   []string `json:"channels" msg:"channels"`
}

// UnSubscribeRep blah
type UnSubscribeRep struct {
	Type     string   `json:"type" msg:"type"`
	Channels []string `json:"channels" msg:"channels"`
}

// Ticker blah
type Ticker struct {
	BestAsk   string `json:"best_ask" msg:"best_ask"`
	BestBid   string `json:"best_bid" msg:"best_bid"`
	LastSize  string `json:"last_size" msg:"last_size"`
	Price     string `json:"price" msg:"price"`
	ProductID string `json:"product_id" msg:"product_id"`
	Sequence  int64  `json:"sequence" msg:"sequence"`
	Side      string `json:"side" msg:"side"`
	Time      string `json:"time" msg:"time"`
	TradeID   int64  `json:"trade_id" msg:"trade_id"`
	Type      string `json:"type" msg:"type"`
}

// Level3Recv blah
type Level3Recv struct {
	Funds     string `json:"funds" msg:"funds"`
	OrderID   string `json:"order_id" msg:"order_id"`
	OrderType string `json:"order_type" msg:"order_type"`
	Price     string `json:"price" msg:"price"`
	ProductID string `json:"product_id" msg:"product_id"`
	Sequence  int64  `json:"sequence" msg:"sequence"`
	Side      string `json:"side" msg:"side"`
	Size      string `json:"size" msg:"size"`
	Time      string `json:"time" msg:"time"`
	Type      string `json:"type" msg:"type"`
}

// Level3Open blah
type Level3Open struct {
	OrderID       string `json:"order_id" msg:"order_id"`
	Price         string `json:"price" msg:"price"`
	ProductID     string `json:"product_id" msg:"product_id"`
	RemainingSize string `json:"remaining_size" msg:"remaining_size"`
	Sequence      int64  `json:"sequence" msg:"sequence"`
	Side          string `json:"side" msg:"side"`
	Time          string `json:"time" msg:"time"`
	Type          string `json:"type" msg:"type"`
}

// Level3Match blah
type Level3Match struct {
	MakerOrderID   string `json:"maker_order_id" msg:"maker_order_id"`
	Price          string `json:"price" msg:"price"`
	ProductID      string `json:"product_id" msg:"product_id"`
	ProfileID      string `json:"profile_id" msg:"profile_id"`
	Sequence       int64  `json:"sequence" msg:"sequence"`
	Side           string `json:"side" msg:"side"`
	Size           string `json:"size" msg:"size"`
	TakerOrderID   string `json:"taker_order_id" msg:"taker_order_id"`
	TakerProfileID string `json:"taker_profile_id" msg:"taker_profile_id"`
	TakerUserID    string `json:"taker_user_id" msg:"taker_user_id"`
	Time           string `json:"time" msg:"time"`
	TradeID        int64  `json:"trade_id" msg:"trade_id"`
	Type           string `json:"type" msg:"type"`
	UserID         string `json:"user_id" msg:"user_id"`
}

// Level3Done blah
type Level3Done struct {
	OrderID       string `json:"order_id" msg:"order_id"`
	Price         string `json:"price" msg:"price"`
	ProductID     string `json:"product_id" msg:"product_id"`
	Reason        string `json:"reason" msg:"reason"`
	RemainingSize string `json:"remaining_size" msg:"remaining_size"`
	Sequence      int64  `json:"sequence" msg:"sequence"`
	Side          string `json:"side" msg:"side"`
	Time          string `json:"time" msg:"time"`
	Type          string `json:"type" msg:"type"`
}

// Level3Change blah
type Level3Change struct {
	NewFunds  string `json:"new_funds" msg:"new_funds"`
	NewSize   string `json:"new_size" msg:"new_size"`
	OldFunds  string `json:"old_funds" msg:"old_funds"`
	OldSize   string `json:"old_size" msg:"old_size"`
	OrderID   string `json:"order_id" msg:"order_id"`
	Price     string `json:"price" msg:"price"`
	ProductID string `json:"product_id" msg:"product_id"`
	Sequence  int64  `json:"sequence" msg:"sequence"`
	Side      string `json:"side" msg:"side"`
	Time      string `json:"time" msg:"time"`
	Type      string `json:"type" msg:"type"`
}

// Heartbeat blah
type Heartbeat struct {
	LastTradeID int64  `json:"last_trade_id" msg:"last_trade_id"`
	ProductID   string `json:"product_id" msg:"product_id"`
	Sequence    int64  `json:"sequence" msg:"sequence"`
	Time        string `json:"time" msg:"time"`
	Type        string `json:"type" msg:"type"`
}

// WSError blah
type WSError struct {
	Message string `json:"message" msg:"message"`
	Type    string `json:"type" msg:"type"`
}

// Activate blah
type Activate struct {
	Funds        string `json:"funds" msg:"funds"`
	OrderID      string `json:"order_id" msg:"order_id"`
	Private      bool   `json:"private" msg:"private"`
	ProductID    string `json:"product_id" msg:"product_id"`
	ProfileID    string `json:"profile_id" msg:"profile_id"`
	Side         string `json:"side" msg:"side"`
	Size         string `json:"size" msg:"size"`
	StopPrice    string `json:"stop_price" msg:"stop_price"`
	StopType     string `json:"stop_type" msg:"stop_type"`
	TakerFeeRate string `json:"taker_fee_rate" msg:"taker_fee_rate"`
	Timestamp    string `json:"timestamp" msg:"timestamp"`
	Type         string `json:"type" msg:"type"`
	UserID       string `json:"user_id" msg:"user_id"`
}

// Level2Update blah
type Level2Update struct {
	Changes   [][]string `json:"changes" msg:"changes"`
	ProductID string     `json:"product_id" msg:"product_id"`
	Type      string     `json:"type" msg:"type"`
}

// Level2Snapshot blah
type Level2Snapshot struct {
	Asks      [][]string `json:"asks" msg:"asks"`
	Bids      [][]string `json:"bids" msg:"bids"`
	ProductID string     `json:"product_id" msg:"product_id"`
	Type      string     `json:"type" msg:"type"`
}

// Time blah
type Time struct {
	Epoch float64 `json:"epoch" msg:"epoch"`
	Iso   string  `json:"iso" msg:"iso"`
}

// UserAccount blah
type UserAccount []struct {
	ExchangeVolume string `json:"exchange_volume" msg:"exchange_volume"`
	ProductID      string `json:"product_id" msg:"product_id"`
	RecordedAt     string `json:"recorded_at" msg:"recorded_at"`
	Volume         string `json:"volume" msg:"volume"`
}

// Account blah
type Account struct {
	Available string `json:"available" msg:"available"`
	Balance   string `json:"balance" msg:"balance"`
	Currency  string `json:"currency" msg:"currency"`
	Hold      string `json:"hold" msg:"hold"`
	ID        string `json:"id" msg:"id"`
	ProfileID string `json:"profile_id" msg:"profile_id"`
}

// AccountDetails blah
type AccountDetails struct {
	OrderID   string `json:"order_id" msg:"order_id"`
	ProductID string `json:"product_id" msg:"product_id"`
	TradeID   string `json:"trade_id" msg:"trade_id"`
}

// AccountHistory blah
type AccountHistory []struct {
	Amount    string         `json:"amount" msg:"amount"`
	Balance   string         `json:"balance" msg:"balance"`
	CreatedAt string         `json:"created_at" msg:"created_at"`
	Details   AccountDetails `json:"details" msg:"details"`
	ID        string         `json:"id" msg:"id"`
	Type      string         `json:"type" msg:"type"`
}

// Accounts blah
type Accounts []struct {
	Available string `json:"available" msg:"available"`
	Balance   string `json:"balance" msg:"balance"`
	Currency  string `json:"currency" msg:"currency"`
	Hold      string `json:"hold" msg:"hold"`
	ID        string `json:"id" msg:"id"`
	ProfileID string `json:"profile_id" msg:"profile_id"`
}

// BankCountry blah
type BankCountry struct {
	Code string `json:"code" msg:"code"`
	Name string `json:"name" msg:"name"`
}

// WireDepositInfo blah
type WireDepositInfo struct {
	AccountAddress string      `json:"account_address" msg:"account_address"`
	AccountName    string      `json:"account_name" msg:"account_name"`
	AccountNumber  string      `json:"account_number" msg:"account_number"`
	BankAddress    string      `json:"bank_address" msg:"bank_address"`
	BnkCountry     BankCountry `json:"bank_country" msg:"bank_country"`
	BankName       string      `json:"bank_name" msg:"bank_name"`
	Reference      string      `json:"reference" msg:"reference"`
	RoutingNumber  string      `json:"routing_number" msg:"routing_number"`
}

// SepaDepositInfo blah
type SepaDepositInfo struct {
	AccountAddress  string `json:"account_address" msg:"account_address"`
	AccountName     string `json:"account_name" msg:"account_name"`
	BankAddress     string `json:"bank_address" msg:"bank_address"`
	BankCountryName string `json:"bank_country_name" msg:"bank_country_name"`
	BankName        string `json:"bank_name" msg:"bank_name"`
	Iban            string `json:"iban" msg:"iban"`
	Reference       string `json:"reference" msg:"reference"`
	Swift           string `json:"swift" msg:"swift"`
}

// AccountsCoinbase blah
type AccountsCoinbase []struct {
	Active                 bool            `json:"active" msg:"active"`
	Balance                string          `json:"balance" msg:"balance"`
	Currency               string          `json:"currency" msg:"currency"`
	ID                     string          `json:"id" msg:"id"`
	Name                   string          `json:"name" msg:"name"`
	Primary                bool            `json:"primary" msg:"primary"`
	SepaDepositInformation SepaDepositInfo `json:"sepa_deposit_information" msg:"sepa_deposit_information"`
	Type                   string          `json:"type" msg:"type"`
	WireDepositInformation WireDepositInfo `json:"wire_deposit_information" msg:"wire_deposit_information"`
}

// Currencies blah
type Currencies []struct {
	ID      string `json:"id" msg:"id"`
	MinSize string `json:"min_size" msg:"min_size"`
	Name    string `json:"name" msg:"name"`
}

// DepositsPaymentMethodReq blah
type DepositsPaymentMethodReq struct {
	Amount          int64  `json:"amount" msg:"amount"`
	Currency        string `json:"currency" msg:"currency"`
	PaymentMethodID string `json:"payment_method_id" msg:"payment_method_id"`
}

// DepositsPaymentMethodRep blah
type DepositsPaymentMethodRep struct {
	Amount   string `json:"amount" msg:"amount"`
	Currency string `json:"currency" msg:"currency"`
	ID       string `json:"id" msg:"id"`
	PayoutAt string `json:"payout_at" msg:"payout_at"`
}

// DepositsCoinbaseReq blah
type DepositsCoinbaseReq struct {
	Amount            int64  `json:"amount" msg:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id" msg:"coinbase_account_id"`
	Currency          string `json:"currency" msg:"currency"`
}

// DepositsCoinbaseRep blah
type DepositsCoinbaseRep struct {
	Amount   string `json:"amount" msg:"amount"`
	Currency string `json:"currency" msg:"currency"`
	ID       string `json:"id" msg:"id"`
}

// ErrorREST blah
type ErrorREST struct {
	Message string `json:"message" msg:"message"`
}

// Fills blah
type Fills []struct {
	CreatedAt string `json:"created_at" msg:"created_at"`
	Fee       string `json:"fee" msg:"fee"`
	Liquidity string `json:"liquidity" msg:"liquidity"`
	OrderID   string `json:"order_id" msg:"order_id"`
	Price     string `json:"price" msg:"price"`
	ProductID string `json:"product_id" msg:"product_id"`
	Settled   bool   `json:"settled" msg:"settled"`
	Side      string `json:"side" msg:"side"`
	Size      string `json:"size" msg:"size"`
	TradeID   int64  `json:"trade_id" msg:"trade_id"`
}

// Holds blah
type Holds []struct {
	AccountID string `json:"account_id" msg:"account_id"`
	Amount    string `json:"amount" msg:"amount"`
	CreatedAt string `json:"created_at" msg:"created_at"`
	ID        string `json:"id" msg:"id"`
	Ref       string `json:"ref" msg:"ref"`
	Type      string `json:"type" msg:"type"`
	UpdatedAt string `json:"updated_at" msg:"updated_at"`
}

// OrderCancel blah
type OrderCancel []string

// OrderGet blah
type OrderGet struct {
	CreatedAt      string `json:"created_at" msg:"created_at"`
	DoneAt         string `json:"done_at" msg:"done_at"`
	DoneReason     string `json:"done_reason" msg:"done_reason"`
	ExecutedValue  string `json:"executed_value" msg:"executed_value"`
	FillFees       string `json:"fill_fees" msg:"fill_fees"`
	FilledSize     string `json:"filled_size" msg:"filled_size"`
	Funds          string `json:"funds" msg:"funds"`
	ID             string `json:"id" msg:"id"`
	PostOnly       bool   `json:"post_only" msg:"post_only"`
	ProductID      string `json:"product_id" msg:"product_id"`
	Settled        bool   `json:"settled" msg:"settled"`
	Side           string `json:"side" msg:"side"`
	Size           string `json:"size" msg:"size"`
	SpecifiedFunds string `json:"specified_funds" msg:"specified_funds"`
	Status         string `json:"status" msg:"status"`
	Stp            string `json:"stp" msg:"stp"`
	Type           string `json:"type" msg:"type"`
}

// OrderList blah
type OrderList []struct {
	CreatedAt     string `json:"created_at" msg:"created_at"`
	ExecutedValue string `json:"executed_value" msg:"executed_value"`
	FillFees      string `json:"fill_fees" msg:"fill_fees"`
	FilledSize    string `json:"filled_size" msg:"filled_size"`
	ID            string `json:"id" msg:"id"`
	PostOnly      bool   `json:"post_only" msg:"post_only"`
	Price         string `json:"price" msg:"price"`
	ProductID     string `json:"product_id" msg:"product_id"`
	Settled       bool   `json:"settled" msg:"settled"`
	Side          string `json:"side" msg:"side"`
	Size          string `json:"size" msg:"size"`
	Status        string `json:"status" msg:"status"`
	Stp           string `json:"stp" msg:"stp"`
	TimeInForce   string `json:"time_in_force" msg:"time_in_force"`
	Type          string `json:"type" msg:"type"`
}

// OrderPlaceRep blah
type OrderPlaceRep struct {
	CreatedAt     string `json:"created_at" msg:"created_at"`
	ExecutedValue string `json:"executed_value" msg:"executed_value"`
	FillFees      string `json:"fill_fees" msg:"fill_fees"`
	FilledSize    string `json:"filled_size" msg:"filled_size"`
	ID            string `json:"id" msg:"id"`
	PostOnly      bool   `json:"post_only" msg:"post_only"`
	Price         string `json:"price" msg:"price"`
	ProductID     string `json:"product_id" msg:"product_id"`
	Settled       bool   `json:"settled" msg:"settled"`
	Side          string `json:"side" msg:"side"`
	Size          string `json:"size" msg:"size"`
	Status        string `json:"status" msg:"status"`
	Stp           string `json:"stp" msg:"stp"`
	TimeInForce   string `json:"time_in_force" msg:"time_in_force"`
	Type          string `json:"type" msg:"type"`
}

// OrderPlaceReq blah
type OrderPlaceReq struct {
	Price     string `json:"price" msg:"price"`
	ProductID string `json:"product_id" msg:"product_id"`
	Side      string `json:"side" msg:"side"`
	Size      string `json:"size" msg:"size"`
}

// Total blah
type Total struct {
	Amount   string `json:"amount" msg:"amount"`
	Currency string `json:"currency" msg:"currency"`
}

// Remaining blah
type Remaining Total

// LimitsDetail blah
type LimitsDetail struct {
	PeriodInDays int16     `json:"period_in_days" msg:"period_in_days"`
	Total        Total     `json:"total" msg:"total"`
	Remaining    Remaining `json:"remaining" msg:"remaining"`
}

// Limits blah
type Limits struct {
	Buy        []LimitsDetail `json:"buy" msg:"buy"`
	Deposit    []LimitsDetail `json:"deposit" msg:"deposit"`
	InstantBuy []LimitsDetail `json:"instant_buy" msg:"instant_buy"`
	Sell       []LimitsDetail `json:"sell" msg:"sell"`
}

// PaymentMethods blah
type PaymentMethods []struct {
	AllowBuy      bool   `json:"allow_buy" msg:"allow_buy"`
	AllowDeposit  bool   `json:"allow_deposit" msg:"allow_deposit"`
	AllowSell     bool   `json:"allow_sell" msg:"allow_sell"`
	AllowWithdraw bool   `json:"allow_withdraw" msg:"allow_withdraw"`
	Currency      string `json:"currency" msg:"currency"`
	ID            string `json:"id" msg:"id"`
	Limits        Limits `json:"limits" msg:"limits"`
	Name          string `json:"name" msg:"name"`
	PrimaryBuy    bool   `json:"primary_buy" msg:"primary_buy"`
	PrimarySell   bool   `json:"primary_sell" msg:"primary_sell"`
	Type          string `json:"type" msg:"type"`
}

// Product24HrStats blah
type Product24HrStats struct {
	High   string `json:"high" msg:"high"`
	Low    string `json:"low" msg:"low"`
	Open   string `json:"open" msg:"open"`
	Volume string `json:"volume" msg:"volume"`
}

// ProductHistoricRates blah
// TODO: THIS MAY NOT GOING TO WORK... FIXME !!!!!!!!
// [
//     [ time, low, high, open, close, volume ],
//     [ 1415398768, 0.32, 4.2, 0.35, 4.2, 12.3 ],
// ]
type ProductHistoricRates [][]interface{}

// ProductLevel2 blah
// TODO: THIS MAY NOT GOING TO WORK... FIXME !!!!!!!!
// {
//     "sequence": "3",
//     "bids": [
//         [ "price", "size", "num-orders" ],
//         [ "295.96", "4.39088265", 2 ]
//     ],
//     "asks": [
//         [ "price", "size", "num-orders" ],
//         [ "295.97", "25.23542881", 12 ]
//     ]
// }
type ProductLevel2 struct {
	Asks     [][]interface{} `json:"asks" msg:"asks"`
	Bids     [][]interface{} `json:"bids" msg:"bids"`
	Sequence string          `json:"sequence" msg:"sequence"`
}

// ProductLevel3 blah
// TODO: THIS MAY NOT GOING TO WORK... FIXME !!!!!!!!
// {
//     "sequence": "3",
//     "bids": [
//         [ price, size, order_id ],
//         [ "295.96","0.05088265","3b0f1225-7f84-490b-a29f-0faef9de823a" ],
//         ...
//     ],
//     "asks": [
//         [ price, size, order_id ],
//         [ "295.97","5.72036512","da863862-25f4-4868-ac41-005d11ab0a5f" ],
//         ...
//     ]
// }
type ProductLevel3 {
	Sequence string `json:"sequence" msg:"sequence"`
	Bids [][]interface{} `json:"bids" msg:"bids"`
	Asks [][]interface{} `json:"asks" msg:"asks"`
}

// ProductLevel1 blah
// This is not correct FIXME !!!!!!!!!!!!!!
// 	{
// 		"sequence": "3",
// 		"bids": [
// 			[ price, size, num-orders ],
// 		],
// 		"asks": [
// 			[ price, size, num-orders ],
// 		]
// 	}
// }
type ProductLevel1 struct {
	Sequence int64          `json:"sequence" msg:"sequence"`
	Bids     [][]interface{} `json:"bids" msg:"bids"`
	Asks     [][]interface{} `json:"asks" msg:"asks"`
}

// ProductTicker blah
type ProductTicker struct {
	Ask     string `json:"ask" msg:"ask"`
	Bid     string `json:"bid" msg:"bid"`
	Price   string `json:"price" msg:"price"`
	Size    string `json:"size" msg:"size"`
	Time    string `json:"time" msg:"time"`
	TradeID int64  `json:"trade_id" msg:"trade_id"`
	Volume  string `json:"volume" msg:"volume"`
}

// ProductTrades blah
type ProductTrades []struct {
	Price   string `json:"price" msg:"price"`
	Side    string `json:"side" msg:"side"`
	Size    string `json:"size" msg:"size"`
	Time    string `json:"time" msg:"time"`
	TradeID int64  `json:"trade_id" msg:"trade_id"`
}

// Products blah
type Products []struct {
	BaseCurrency   string `json:"base_currency" msg:"base_currency"`
	BaseMaxSize    string `json:"base_max_size" msg:"base_max_size"`
	BaseMinSize    string `json:"base_min_size" msg:"base_min_size"`
	ID             string `json:"id" msg:"id"`
	QuoteCurrency  string `json:"quote_currency" msg:"quote_currency"`
	QuoteIncrement string `json:"quote_increment" msg:"quote_increment"`
}

// Parameters blah
type Parameters struct {
	EndDate   string `json:"end_date" msg:"end_date"`
	StartDate string `json:"start_date" msg:"start_date"`
}

// ReportsRep blah
type ReportsRep struct {
	CompletedAt string     `json:"completed_at" msg:"completed_at"`
	CreatedAt   string     `json:"created_at" msg:"created_at"`
	ExpiresAt   string     `json:"expires_at" msg:"expires_at"`
	FileURL     string     `json:"file_url" msg:"file_url"`
	ID          string     `json:"id" msg:"id"`
	Params      Parameters `json:"params" msg:"params"`
	Status      string     `json:"status" msg:"status"`
	Type        string     `json:"type" msg:"type"`
}

// ReportsReq blah
type ReportsReq struct {
	EndDate   string `json:"end_date" msg:"end_date"`
	StartDate string `json:"start_date" msg:"start_date"`
	Type      string `json:"type" msg:"type"`
}
