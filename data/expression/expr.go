package expression

import (
	"github.com/AiRISTAFlowInc/flow-studio-core/data"
	"github.com/AiRISTAFlowInc/flow-studio-core/data/resolve"
)

type Factory interface {
	NewExpr(exprStr string) (Expr, error)
}

type Expr interface {
	Eval(scope data.Scope) (interface{}, error)
}

type FactoryCreatorFunc func(resolve.CompositeResolver) Factory
