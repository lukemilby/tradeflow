package engine

import "github.com/lukemilby/learn-fp-go/2-design-patterns/ch04-solid/02_maybe/src/maybe"

type Engine interface {
	Run
	Stop
	Backtest
	SetConfig
	SetStrategy
}


// Default engine for tradeflow cli
type TradeFlowEngine struct {

}

