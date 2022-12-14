package mexcsdk

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"mexc-sdk.Base",
		reflect.TypeOf((*Base)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "publicRequest", GoMethod: "PublicRequest"},
			_jsii_.MemberMethod{JsiiMethod: "signRequest", GoMethod: "SignRequest"},
		},
		func() interface{} {
			return &jsiiProxy_Base{}
		},
	)
	_jsii_.RegisterClass(
		"mexc-sdk.Common",
		reflect.TypeOf((*Common)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "aggTrades", GoMethod: "AggTrades"},
			_jsii_.MemberMethod{JsiiMethod: "avgPrice", GoMethod: "AvgPrice"},
			_jsii_.MemberMethod{JsiiMethod: "bookTicker", GoMethod: "BookTicker"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "depth", GoMethod: "Depth"},
			_jsii_.MemberMethod{JsiiMethod: "exchangeInfo", GoMethod: "ExchangeInfo"},
			_jsii_.MemberMethod{JsiiMethod: "historicalTrades", GoMethod: "HistoricalTrades"},
			_jsii_.MemberMethod{JsiiMethod: "klines", GoMethod: "Klines"},
			_jsii_.MemberMethod{JsiiMethod: "ping", GoMethod: "Ping"},
			_jsii_.MemberMethod{JsiiMethod: "publicRequest", GoMethod: "PublicRequest"},
			_jsii_.MemberMethod{JsiiMethod: "signRequest", GoMethod: "SignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "ticker24hr", GoMethod: "Ticker24hr"},
			_jsii_.MemberMethod{JsiiMethod: "tickerPrice", GoMethod: "TickerPrice"},
			_jsii_.MemberMethod{JsiiMethod: "time", GoMethod: "Time"},
			_jsii_.MemberMethod{JsiiMethod: "trades", GoMethod: "Trades"},
		},
		func() interface{} {
			j := jsiiProxy_Common{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Market)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"mexc-sdk.Market",
		reflect.TypeOf((*Market)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "aggTrades", GoMethod: "AggTrades"},
			_jsii_.MemberMethod{JsiiMethod: "avgPrice", GoMethod: "AvgPrice"},
			_jsii_.MemberMethod{JsiiMethod: "bookTicker", GoMethod: "BookTicker"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "depth", GoMethod: "Depth"},
			_jsii_.MemberMethod{JsiiMethod: "exchangeInfo", GoMethod: "ExchangeInfo"},
			_jsii_.MemberMethod{JsiiMethod: "historicalTrades", GoMethod: "HistoricalTrades"},
			_jsii_.MemberMethod{JsiiMethod: "klines", GoMethod: "Klines"},
			_jsii_.MemberMethod{JsiiMethod: "publicRequest", GoMethod: "PublicRequest"},
			_jsii_.MemberMethod{JsiiMethod: "signRequest", GoMethod: "SignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "ticker24hr", GoMethod: "Ticker24hr"},
			_jsii_.MemberMethod{JsiiMethod: "tickerPrice", GoMethod: "TickerPrice"},
			_jsii_.MemberMethod{JsiiMethod: "trades", GoMethod: "Trades"},
		},
		func() interface{} {
			j := jsiiProxy_Market{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Base)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"mexc-sdk.Spot",
		reflect.TypeOf((*Spot)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "accountInfo", GoMethod: "AccountInfo"},
			_jsii_.MemberMethod{JsiiMethod: "accountTradeList", GoMethod: "AccountTradeList"},
			_jsii_.MemberMethod{JsiiMethod: "aggTrades", GoMethod: "AggTrades"},
			_jsii_.MemberMethod{JsiiMethod: "allOrders", GoMethod: "AllOrders"},
			_jsii_.MemberMethod{JsiiMethod: "avgPrice", GoMethod: "AvgPrice"},
			_jsii_.MemberMethod{JsiiMethod: "bookTicker", GoMethod: "BookTicker"},
			_jsii_.MemberMethod{JsiiMethod: "cancelOpenOrders", GoMethod: "CancelOpenOrders"},
			_jsii_.MemberMethod{JsiiMethod: "cancelOrder", GoMethod: "CancelOrder"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "depth", GoMethod: "Depth"},
			_jsii_.MemberMethod{JsiiMethod: "exchangeInfo", GoMethod: "ExchangeInfo"},
			_jsii_.MemberMethod{JsiiMethod: "historicalTrades", GoMethod: "HistoricalTrades"},
			_jsii_.MemberMethod{JsiiMethod: "klines", GoMethod: "Klines"},
			_jsii_.MemberMethod{JsiiMethod: "newOrder", GoMethod: "NewOrder"},
			_jsii_.MemberMethod{JsiiMethod: "newOrderTest", GoMethod: "NewOrderTest"},
			_jsii_.MemberMethod{JsiiMethod: "openOrders", GoMethod: "OpenOrders"},
			_jsii_.MemberMethod{JsiiMethod: "ping", GoMethod: "Ping"},
			_jsii_.MemberMethod{JsiiMethod: "publicRequest", GoMethod: "PublicRequest"},
			_jsii_.MemberMethod{JsiiMethod: "queryOrder", GoMethod: "QueryOrder"},
			_jsii_.MemberMethod{JsiiMethod: "signRequest", GoMethod: "SignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "ticker24hr", GoMethod: "Ticker24hr"},
			_jsii_.MemberMethod{JsiiMethod: "tickerPrice", GoMethod: "TickerPrice"},
			_jsii_.MemberMethod{JsiiMethod: "time", GoMethod: "Time"},
			_jsii_.MemberMethod{JsiiMethod: "trades", GoMethod: "Trades"},
		},
		func() interface{} {
			j := jsiiProxy_Spot{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Trade)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"mexc-sdk.Trade",
		reflect.TypeOf((*Trade)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "accountInfo", GoMethod: "AccountInfo"},
			_jsii_.MemberMethod{JsiiMethod: "accountTradeList", GoMethod: "AccountTradeList"},
			_jsii_.MemberMethod{JsiiMethod: "aggTrades", GoMethod: "AggTrades"},
			_jsii_.MemberMethod{JsiiMethod: "allOrders", GoMethod: "AllOrders"},
			_jsii_.MemberMethod{JsiiMethod: "avgPrice", GoMethod: "AvgPrice"},
			_jsii_.MemberMethod{JsiiMethod: "bookTicker", GoMethod: "BookTicker"},
			_jsii_.MemberMethod{JsiiMethod: "cancelOpenOrders", GoMethod: "CancelOpenOrders"},
			_jsii_.MemberMethod{JsiiMethod: "cancelOrder", GoMethod: "CancelOrder"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "depth", GoMethod: "Depth"},
			_jsii_.MemberMethod{JsiiMethod: "exchangeInfo", GoMethod: "ExchangeInfo"},
			_jsii_.MemberMethod{JsiiMethod: "historicalTrades", GoMethod: "HistoricalTrades"},
			_jsii_.MemberMethod{JsiiMethod: "klines", GoMethod: "Klines"},
			_jsii_.MemberMethod{JsiiMethod: "newOrder", GoMethod: "NewOrder"},
			_jsii_.MemberMethod{JsiiMethod: "newOrderTest", GoMethod: "NewOrderTest"},
			_jsii_.MemberMethod{JsiiMethod: "openOrders", GoMethod: "OpenOrders"},
			_jsii_.MemberMethod{JsiiMethod: "ping", GoMethod: "Ping"},
			_jsii_.MemberMethod{JsiiMethod: "publicRequest", GoMethod: "PublicRequest"},
			_jsii_.MemberMethod{JsiiMethod: "queryOrder", GoMethod: "QueryOrder"},
			_jsii_.MemberMethod{JsiiMethod: "signRequest", GoMethod: "SignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "ticker24hr", GoMethod: "Ticker24hr"},
			_jsii_.MemberMethod{JsiiMethod: "tickerPrice", GoMethod: "TickerPrice"},
			_jsii_.MemberMethod{JsiiMethod: "time", GoMethod: "Time"},
			_jsii_.MemberMethod{JsiiMethod: "trades", GoMethod: "Trades"},
		},
		func() interface{} {
			j := jsiiProxy_Trade{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_UserData)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"mexc-sdk.UserData",
		reflect.TypeOf((*UserData)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "accountInfo", GoMethod: "AccountInfo"},
			_jsii_.MemberMethod{JsiiMethod: "accountTradeList", GoMethod: "AccountTradeList"},
			_jsii_.MemberMethod{JsiiMethod: "aggTrades", GoMethod: "AggTrades"},
			_jsii_.MemberMethod{JsiiMethod: "avgPrice", GoMethod: "AvgPrice"},
			_jsii_.MemberMethod{JsiiMethod: "bookTicker", GoMethod: "BookTicker"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberMethod{JsiiMethod: "depth", GoMethod: "Depth"},
			_jsii_.MemberMethod{JsiiMethod: "exchangeInfo", GoMethod: "ExchangeInfo"},
			_jsii_.MemberMethod{JsiiMethod: "historicalTrades", GoMethod: "HistoricalTrades"},
			_jsii_.MemberMethod{JsiiMethod: "klines", GoMethod: "Klines"},
			_jsii_.MemberMethod{JsiiMethod: "ping", GoMethod: "Ping"},
			_jsii_.MemberMethod{JsiiMethod: "publicRequest", GoMethod: "PublicRequest"},
			_jsii_.MemberMethod{JsiiMethod: "signRequest", GoMethod: "SignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "ticker24hr", GoMethod: "Ticker24hr"},
			_jsii_.MemberMethod{JsiiMethod: "tickerPrice", GoMethod: "TickerPrice"},
			_jsii_.MemberMethod{JsiiMethod: "time", GoMethod: "Time"},
			_jsii_.MemberMethod{JsiiMethod: "trades", GoMethod: "Trades"},
		},
		func() interface{} {
			j := jsiiProxy_UserData{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Common)
			return &j
		},
	)
}
