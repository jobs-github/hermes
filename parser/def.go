package parser

import "hermes/ast"

type parseBlockStmtFn func() *ast.BlockStmt
type parseExpressionFn func(precedence int) ast.Expression
