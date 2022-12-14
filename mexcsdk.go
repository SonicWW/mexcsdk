// mexc sdk
package mexcsdk

import (
	_init_ "mexc-sdk/mexcsdk/jsii"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type Base interface {
	Config() interface{}
	SetConfig(val interface{})
	PublicRequest(method *string, path *string, paramsObj interface{}) interface{}
	SignRequest(method *string, path *string, paramsObj interface{}) interface{}
}

// The jsii proxy struct for Base
type jsiiProxy_Base struct {
	_ byte // padding
}

func (j *jsiiProxy_Base) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewBase(apiKey *string, apiSecret *string) Base {
	_init_.Initialize()

	j := jsiiProxy_Base{}

	_jsii_.Create(
		"mexc-sdk.Base",
		[]interface{}{apiKey, apiSecret},
		&j,
	)

	return &j
}

func NewBase_Override(b Base, apiKey *string, apiSecret *string) {
	_init_.Initialize()

	_jsii_.Create(
		"mexc-sdk.Base",
		[]interface{}{apiKey, apiSecret},
		b,
	)
}

func (j *jsiiProxy_Base) SetConfig(val interface{}) {
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

func (b *jsiiProxy_Base) PublicRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"publicRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_Base) SignRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"signRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

type Common interface {
	Market
	Config() interface{}
	SetConfig(val interface{})
	AggTrades(symbol *string, options interface{}) interface{}
	AvgPrice(symbol *string) interface{}
	BookTicker(symbol *string) interface{}
	Depth(symbol *string, options interface{}) interface{}
	ExchangeInfo(options interface{}) interface{}
	HistoricalTrades(symbol *string, options interface{}) interface{}
	Klines(symbol *string, interval *string, options interface{}) interface{}
	Ping() interface{}
	PublicRequest(method *string, path *string, paramsObj interface{}) interface{}
	SignRequest(method *string, path *string, paramsObj interface{}) interface{}
	Ticker24hr(symbol *string) interface{}
	TickerPrice(symbol *string) interface{}
	Time() interface{}
	Trades(symbol *string, options interface{}) interface{}
}

// The jsii proxy struct for Common
type jsiiProxy_Common struct {
	jsiiProxy_Market
}

func (j *jsiiProxy_Common) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewCommon(apiKey *string, apiSecret *string) Common {
	_init_.Initialize()

	j := jsiiProxy_Common{}

	_jsii_.Create(
		"mexc-sdk.Common",
		[]interface{}{apiKey, apiSecret},
		&j,
	)

	return &j
}

func NewCommon_Override(c Common, apiKey *string, apiSecret *string) {
	_init_.Initialize()

	_jsii_.Create(
		"mexc-sdk.Common",
		[]interface{}{apiKey, apiSecret},
		c,
	)
}

func (j *jsiiProxy_Common) SetConfig(val interface{}) {
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Compressed/Aggregate Trades List.
//
// Note: If sending startTime and endTime, the interval must be less than one hour
func (c *jsiiProxy_Common) AggTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"aggTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Current Average Price.
func (c *jsiiProxy_Common) AvgPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"avgPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Order Book Ticker.
func (c *jsiiProxy_Common) BookTicker(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"bookTicker",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Order Book.
func (c *jsiiProxy_Common) Depth(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"depth",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Exchange Information.
func (c *jsiiProxy_Common) ExchangeInfo(options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"exchangeInfo",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Old Trade Lookup.
func (c *jsiiProxy_Common) HistoricalTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"historicalTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Kline/Candlestick Data.
func (c *jsiiProxy_Common) Klines(symbol *string, interval *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"klines",
		[]interface{}{symbol, interval, options},
		&returns,
	)

	return returns
}

// Test Connectivity.
func (c *jsiiProxy_Common) Ping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"ping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Common) PublicRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"publicRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Common) SignRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"signRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// 24hr Ticker Price Change Statistics.
func (c *jsiiProxy_Common) Ticker24hr(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"ticker24hr",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Price Ticker.
func (c *jsiiProxy_Common) TickerPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"tickerPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Check Server Time.
func (c *jsiiProxy_Common) Time() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"time",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Recent Trades List.
func (c *jsiiProxy_Common) Trades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"trades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

type Market interface {
	Base
	Config() interface{}
	SetConfig(val interface{})
	AggTrades(symbol *string, options interface{}) interface{}
	AvgPrice(symbol *string) interface{}
	BookTicker(symbol *string) interface{}
	Depth(symbol *string, options interface{}) interface{}
	ExchangeInfo(options interface{}) interface{}
	HistoricalTrades(symbol *string, options interface{}) interface{}
	Klines(symbol *string, interval *string, options interface{}) interface{}
	PublicRequest(method *string, path *string, paramsObj interface{}) interface{}
	SignRequest(method *string, path *string, paramsObj interface{}) interface{}
	Ticker24hr(symbol *string) interface{}
	TickerPrice(symbol *string) interface{}
	Trades(symbol *string, options interface{}) interface{}
}

// The jsii proxy struct for Market
type jsiiProxy_Market struct {
	jsiiProxy_Base
}

func (j *jsiiProxy_Market) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewMarket(apiKey *string, apiSecret *string) Market {
	_init_.Initialize()

	j := jsiiProxy_Market{}

	_jsii_.Create(
		"mexc-sdk.Market",
		[]interface{}{apiKey, apiSecret},
		&j,
	)

	return &j
}

func NewMarket_Override(m Market, apiKey *string, apiSecret *string) {
	_init_.Initialize()

	_jsii_.Create(
		"mexc-sdk.Market",
		[]interface{}{apiKey, apiSecret},
		m,
	)
}

func (j *jsiiProxy_Market) SetConfig(val interface{}) {
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Compressed/Aggregate Trades List.
//
// Note: If sending startTime and endTime, the interval must be less than one hour
func (m *jsiiProxy_Market) AggTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"aggTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Current Average Price.
func (m *jsiiProxy_Market) AvgPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"avgPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Order Book Ticker.
func (m *jsiiProxy_Market) BookTicker(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"bookTicker",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Order Book.
func (m *jsiiProxy_Market) Depth(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"depth",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Exchange Information.
func (m *jsiiProxy_Market) ExchangeInfo(options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"exchangeInfo",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Old Trade Lookup.
func (m *jsiiProxy_Market) HistoricalTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"historicalTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Kline/Candlestick Data.
func (m *jsiiProxy_Market) Klines(symbol *string, interval *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"klines",
		[]interface{}{symbol, interval, options},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Market) PublicRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"publicRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Market) SignRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"signRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// 24hr Ticker Price Change Statistics.
func (m *jsiiProxy_Market) Ticker24hr(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"ticker24hr",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Price Ticker.
func (m *jsiiProxy_Market) TickerPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"tickerPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Recent Trades List.
func (m *jsiiProxy_Market) Trades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"trades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

type Spot interface {
	Trade
	Config() interface{}
	SetConfig(val interface{})
	AccountInfo() interface{}
	AccountTradeList(symbol *string, options interface{}) interface{}
	AggTrades(symbol *string, options interface{}) interface{}
	AllOrders(symbol *string, options interface{}) interface{}
	AvgPrice(symbol *string) interface{}
	BookTicker(symbol *string) interface{}
	CancelOpenOrders(symbol *string) interface{}
	CancelOrder(symbol *string, options interface{}) interface{}
	Depth(symbol *string, options interface{}) interface{}
	ExchangeInfo(options interface{}) interface{}
	HistoricalTrades(symbol *string, options interface{}) interface{}
	Klines(symbol *string, interval *string, options interface{}) interface{}
	NewOrder(symbol *string, side *string, orderType *string, options interface{}) interface{}
	NewOrderTest(symbol *string, side *string, orderType *string, options interface{}) interface{}
	OpenOrders(symbol *string) interface{}
	Ping() interface{}
	PublicRequest(method *string, path *string, paramsObj interface{}) interface{}
	QueryOrder(symbol *string, options interface{}) interface{}
	SignRequest(method *string, path *string, paramsObj interface{}) interface{}
	Ticker24hr(symbol *string) interface{}
	TickerPrice(symbol *string) interface{}
	Time() interface{}
	Trades(symbol *string, options interface{}) interface{}
}

// The jsii proxy struct for Spot
type jsiiProxy_Spot struct {
	jsiiProxy_Trade
}

func (j *jsiiProxy_Spot) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewSpot(apiKey *string, apiSecret *string) Spot {
	_init_.Initialize()

	j := jsiiProxy_Spot{}

	_jsii_.Create(
		"mexc-sdk.Spot",
		[]interface{}{apiKey, apiSecret},
		&j,
	)

	return &j
}

func NewSpot_Override(s Spot, apiKey *string, apiSecret *string) {
	_init_.Initialize()

	_jsii_.Create(
		"mexc-sdk.Spot",
		[]interface{}{apiKey, apiSecret},
		s,
	)
}

func (j *jsiiProxy_Spot) SetConfig(val interface{}) {
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Account Information.
func (s *jsiiProxy_Spot) AccountInfo() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"accountInfo",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Account Trade List.
func (s *jsiiProxy_Spot) AccountTradeList(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"accountTradeList",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Compressed/Aggregate Trades List.
//
// Note: If sending startTime and endTime, the interval must be less than one hour
func (s *jsiiProxy_Spot) AggTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"aggTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// All Orders.
func (s *jsiiProxy_Spot) AllOrders(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"allOrders",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Current Average Price.
func (s *jsiiProxy_Spot) AvgPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"avgPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Order Book Ticker.
func (s *jsiiProxy_Spot) BookTicker(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"bookTicker",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Cancel all Open Orders on a Symbol.
func (s *jsiiProxy_Spot) CancelOpenOrders(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"cancelOpenOrders",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Cancel Order.
func (s *jsiiProxy_Spot) CancelOrder(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"cancelOrder",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Order Book.
func (s *jsiiProxy_Spot) Depth(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"depth",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Exchange Information.
func (s *jsiiProxy_Spot) ExchangeInfo(options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"exchangeInfo",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Old Trade Lookup.
func (s *jsiiProxy_Spot) HistoricalTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"historicalTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Kline/Candlestick Data.
func (s *jsiiProxy_Spot) Klines(symbol *string, interval *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"klines",
		[]interface{}{symbol, interval, options},
		&returns,
	)

	return returns
}

// New Order.
func (s *jsiiProxy_Spot) NewOrder(symbol *string, side *string, orderType *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"newOrder",
		[]interface{}{symbol, side, orderType, options},
		&returns,
	)

	return returns
}

// Test New Order.
func (s *jsiiProxy_Spot) NewOrderTest(symbol *string, side *string, orderType *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"newOrderTest",
		[]interface{}{symbol, side, orderType, options},
		&returns,
	)

	return returns
}

// Current Open Orders.
func (s *jsiiProxy_Spot) OpenOrders(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"openOrders",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Test Connectivity.
func (s *jsiiProxy_Spot) Ping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"ping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Spot) PublicRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"publicRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// Query Order.
func (s *jsiiProxy_Spot) QueryOrder(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"queryOrder",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Spot) SignRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"signRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// 24hr Ticker Price Change Statistics.
func (s *jsiiProxy_Spot) Ticker24hr(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"ticker24hr",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Price Ticker.
func (s *jsiiProxy_Spot) TickerPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"tickerPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Check Server Time.
func (s *jsiiProxy_Spot) Time() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"time",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Recent Trades List.
func (s *jsiiProxy_Spot) Trades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"trades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

type Trade interface {
	UserData
	Config() interface{}
	SetConfig(val interface{})
	AccountInfo() interface{}
	AccountTradeList(symbol *string, options interface{}) interface{}
	AggTrades(symbol *string, options interface{}) interface{}
	AllOrders(symbol *string, options interface{}) interface{}
	AvgPrice(symbol *string) interface{}
	BookTicker(symbol *string) interface{}
	CancelOpenOrders(symbol *string) interface{}
	CancelOrder(symbol *string, options interface{}) interface{}
	Depth(symbol *string, options interface{}) interface{}
	ExchangeInfo(options interface{}) interface{}
	HistoricalTrades(symbol *string, options interface{}) interface{}
	Klines(symbol *string, interval *string, options interface{}) interface{}
	NewOrder(symbol *string, side *string, orderType *string, options interface{}) interface{}
	NewOrderTest(symbol *string, side *string, orderType *string, options interface{}) interface{}
	OpenOrders(symbol *string) interface{}
	Ping() interface{}
	PublicRequest(method *string, path *string, paramsObj interface{}) interface{}
	QueryOrder(symbol *string, options interface{}) interface{}
	SignRequest(method *string, path *string, paramsObj interface{}) interface{}
	Ticker24hr(symbol *string) interface{}
	TickerPrice(symbol *string) interface{}
	Time() interface{}
	Trades(symbol *string, options interface{}) interface{}
}

// The jsii proxy struct for Trade
type jsiiProxy_Trade struct {
	jsiiProxy_UserData
}

func (j *jsiiProxy_Trade) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewTrade(apiKey *string, apiSecret *string) Trade {
	_init_.Initialize()

	j := jsiiProxy_Trade{}

	_jsii_.Create(
		"mexc-sdk.Trade",
		[]interface{}{apiKey, apiSecret},
		&j,
	)

	return &j
}

func NewTrade_Override(t Trade, apiKey *string, apiSecret *string) {
	_init_.Initialize()

	_jsii_.Create(
		"mexc-sdk.Trade",
		[]interface{}{apiKey, apiSecret},
		t,
	)
}

func (j *jsiiProxy_Trade) SetConfig(val interface{}) {
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Account Information.
func (t *jsiiProxy_Trade) AccountInfo() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"accountInfo",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Account Trade List.
func (t *jsiiProxy_Trade) AccountTradeList(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"accountTradeList",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Compressed/Aggregate Trades List.
//
// Note: If sending startTime and endTime, the interval must be less than one hour
func (t *jsiiProxy_Trade) AggTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"aggTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// All Orders.
func (t *jsiiProxy_Trade) AllOrders(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"allOrders",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Current Average Price.
func (t *jsiiProxy_Trade) AvgPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"avgPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Order Book Ticker.
func (t *jsiiProxy_Trade) BookTicker(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"bookTicker",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Cancel all Open Orders on a Symbol.
func (t *jsiiProxy_Trade) CancelOpenOrders(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"cancelOpenOrders",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Cancel Order.
func (t *jsiiProxy_Trade) CancelOrder(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"cancelOrder",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Order Book.
func (t *jsiiProxy_Trade) Depth(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"depth",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Exchange Information.
func (t *jsiiProxy_Trade) ExchangeInfo(options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"exchangeInfo",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Old Trade Lookup.
func (t *jsiiProxy_Trade) HistoricalTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"historicalTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Kline/Candlestick Data.
func (t *jsiiProxy_Trade) Klines(symbol *string, interval *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"klines",
		[]interface{}{symbol, interval, options},
		&returns,
	)

	return returns
}

// New Order.
func (t *jsiiProxy_Trade) NewOrder(symbol *string, side *string, orderType *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"newOrder",
		[]interface{}{symbol, side, orderType, options},
		&returns,
	)

	return returns
}

// Test New Order.
func (t *jsiiProxy_Trade) NewOrderTest(symbol *string, side *string, orderType *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"newOrderTest",
		[]interface{}{symbol, side, orderType, options},
		&returns,
	)

	return returns
}

// Current Open Orders.
func (t *jsiiProxy_Trade) OpenOrders(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"openOrders",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Test Connectivity.
func (t *jsiiProxy_Trade) Ping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"ping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trade) PublicRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"publicRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// Query Order.
func (t *jsiiProxy_Trade) QueryOrder(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"queryOrder",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trade) SignRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"signRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// 24hr Ticker Price Change Statistics.
func (t *jsiiProxy_Trade) Ticker24hr(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"ticker24hr",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Price Ticker.
func (t *jsiiProxy_Trade) TickerPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"tickerPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Check Server Time.
func (t *jsiiProxy_Trade) Time() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"time",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Recent Trades List.
func (t *jsiiProxy_Trade) Trades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"trades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

type UserData interface {
	Common
	Config() interface{}
	SetConfig(val interface{})
	AccountInfo() interface{}
	AccountTradeList(symbol *string, options interface{}) interface{}
	AggTrades(symbol *string, options interface{}) interface{}
	AvgPrice(symbol *string) interface{}
	BookTicker(symbol *string) interface{}
	Depth(symbol *string, options interface{}) interface{}
	ExchangeInfo(options interface{}) interface{}
	HistoricalTrades(symbol *string, options interface{}) interface{}
	Klines(symbol *string, interval *string, options interface{}) interface{}
	Ping() interface{}
	PublicRequest(method *string, path *string, paramsObj interface{}) interface{}
	SignRequest(method *string, path *string, paramsObj interface{}) interface{}
	Ticker24hr(symbol *string) interface{}
	TickerPrice(symbol *string) interface{}
	Time() interface{}
	Trades(symbol *string, options interface{}) interface{}
}

// The jsii proxy struct for UserData
type jsiiProxy_UserData struct {
	jsiiProxy_Common
}

func (j *jsiiProxy_UserData) Config() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"config",
		&returns,
	)
	return returns
}


func NewUserData(apiKey *string, apiSecret *string) UserData {
	_init_.Initialize()

	j := jsiiProxy_UserData{}

	_jsii_.Create(
		"mexc-sdk.UserData",
		[]interface{}{apiKey, apiSecret},
		&j,
	)

	return &j
}

func NewUserData_Override(u UserData, apiKey *string, apiSecret *string) {
	_init_.Initialize()

	_jsii_.Create(
		"mexc-sdk.UserData",
		[]interface{}{apiKey, apiSecret},
		u,
	)
}

func (j *jsiiProxy_UserData) SetConfig(val interface{}) {
	_jsii_.Set(
		j,
		"config",
		val,
	)
}

// Account Information.
func (u *jsiiProxy_UserData) AccountInfo() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"accountInfo",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Account Trade List.
func (u *jsiiProxy_UserData) AccountTradeList(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"accountTradeList",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Compressed/Aggregate Trades List.
//
// Note: If sending startTime and endTime, the interval must be less than one hour
func (u *jsiiProxy_UserData) AggTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"aggTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Current Average Price.
func (u *jsiiProxy_UserData) AvgPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"avgPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Order Book Ticker.
func (u *jsiiProxy_UserData) BookTicker(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"bookTicker",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Order Book.
func (u *jsiiProxy_UserData) Depth(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"depth",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Exchange Information.
func (u *jsiiProxy_UserData) ExchangeInfo(options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"exchangeInfo",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Old Trade Lookup.
func (u *jsiiProxy_UserData) HistoricalTrades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"historicalTrades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

// Kline/Candlestick Data.
func (u *jsiiProxy_UserData) Klines(symbol *string, interval *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"klines",
		[]interface{}{symbol, interval, options},
		&returns,
	)

	return returns
}

// Test Connectivity.
func (u *jsiiProxy_UserData) Ping() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"ping",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserData) PublicRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"publicRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UserData) SignRequest(method *string, path *string, paramsObj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"signRequest",
		[]interface{}{method, path, paramsObj},
		&returns,
	)

	return returns
}

// 24hr Ticker Price Change Statistics.
func (u *jsiiProxy_UserData) Ticker24hr(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"ticker24hr",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Symbol Price Ticker.
func (u *jsiiProxy_UserData) TickerPrice(symbol *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"tickerPrice",
		[]interface{}{symbol},
		&returns,
	)

	return returns
}

// Check Server Time.
func (u *jsiiProxy_UserData) Time() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"time",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Recent Trades List.
func (u *jsiiProxy_UserData) Trades(symbol *string, options interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		u,
		"trades",
		[]interface{}{symbol, options},
		&returns,
	)

	return returns
}

