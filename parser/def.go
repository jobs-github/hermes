package parser

import "Q/ast"

type parseBlockStmtFn func() *ast.BlockStmt
type parseExpressionFn func(precedence int) ast.Expression
