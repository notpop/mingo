package mingo

import (
	"fmt"
	"go/ast"
	"strings"
)

func stringifySelectExpr(expr *ast.SelectorExpr) string {
	switch x := expr.X.(type) {
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", stringifySelectExpr(x), expr.Sel.Name)
	default:
		return fmt.Sprintf("%s.%s", x.(*ast.Ident).Name, expr.Sel.Name)
	}
}

func stringifyExpr(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return stringifyBasicLit(x)
	case *ast.CallExpr:
		return stringifyCallExpr(x)
	case *ast.SelectorExpr:
		return stringifySelectExpr(x)
	case *ast.StarExpr:
		return stringifyStarExpr(x)
	case *ast.ArrayType:
		return stringifyArrayType(x)
	case *ast.Ellipsis:
		return stringifyEllipsis(x)
	case *ast.FuncLit:
		// TODO
	case *ast.BinaryExpr:
		return stringifyBinaryExpr(x)
	case *ast.SliceExpr:
		return stringifySliceExpr(x)
	case *ast.UnaryExpr:
		return stringifyUnaryExpr(x)
	case *ast.CompositeLit:
		return stringifyCompositeLit(x)
	case *ast.ParenExpr:
		return stringifyParenExpr(x)
	case *ast.IndexExpr:
		return stringifyIndexExpr(x)
	case *ast.KeyValueExpr:
		return stringifyKeyValueExpr(x)
	case *ast.TypeAssertExpr:
		return stringifyTypeAssertExpr(x)
	case *ast.ChanType:
		return stringifyChanType(x)
	case *ast.MapType:
		return stringifyMapType(x)
	case *ast.InterfaceType:
		return stringifyInterfaceType(x)
	case *ast.StructType:
		return stringifyStructType(x)
	case *ast.FuncType:
		return stringifyFuncType(x)
	}

	return expr.(*ast.Ident).Name
}

func stringifyBasicLit(lit *ast.BasicLit) string {
	return lit.Value
}

func stringifyCallExpr(expr *ast.CallExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.Fun))
	sb.WriteString("(")
	for i, arg := range expr.Args {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(arg))
	}
	sb.WriteString(")")

	return sb.String()
}

func stringifyStarExpr(expr *ast.StarExpr) string {
	return fmt.Sprintf("*%s", stringifyExpr(expr.X))
}

func stringifyArrayType(expr *ast.ArrayType) string {
	return fmt.Sprintf("[]%s", stringifyExpr(expr.Elt))
}

func stringifyEllipsis(expr *ast.Ellipsis) string {
	return fmt.Sprintf("...%s", stringifyExpr(expr.Elt))
}

func stringifyBinaryExpr(expr *ast.BinaryExpr) string {
	return fmt.Sprintf("%s%s%s", stringifyExpr(expr.X), expr.Op.String(), stringifyExpr(expr.Y))
}

func stringifySliceExpr(expr *ast.SliceExpr) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.X))
	sb.WriteString("[")
	if expr.Low != nil {
		sb.WriteString(stringifyExpr(expr.Low))
	}
	sb.WriteString(":") // FIXME
	if expr.High != nil {
		sb.WriteString(stringifyExpr(expr.High))
	}
	if expr.Max != nil {
		sb.WriteString(":")
		sb.WriteString(stringifyExpr(expr.Max))
	}
	sb.WriteString("]")

	return sb.String()
}

func stringifyUnaryExpr(expr *ast.UnaryExpr) string {
	return fmt.Sprintf("%s%s", expr.Op.String(), stringifyExpr(expr.X))
}

func stringifyCompositeLit(expr *ast.CompositeLit) string {
	sb := new(strings.Builder)

	sb.WriteString(stringifyExpr(expr.Type))
	sb.WriteString("{")
	for i, elt := range expr.Elts {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(stringifyExpr(elt))
	}
	sb.WriteString("}")

	return sb.String()
}

func stringifyParenExpr(expr *ast.ParenExpr) string {
	return fmt.Sprintf("(%s)", stringifyExpr(expr.X))
}

func stringifyIndexExpr(expr *ast.IndexExpr) string {
	return fmt.Sprintf("%s[%s]", stringifyExpr(expr.X), stringifyExpr(expr.Index))
}

func stringifyKeyValueExpr(expr *ast.KeyValueExpr) string {
	return fmt.Sprintf("%s:%s", stringifyExpr(expr.Key), stringifyExpr(expr.Value))
}

func stringifyTypeAssertExpr(expr *ast.TypeAssertExpr) string {
	return fmt.Sprintf("%s.(%s)", stringifyExpr(expr.X), stringifyExpr(expr.Type))
}

func stringifyChanType(expr *ast.ChanType) string {
	sb := new(strings.Builder)

	if expr.Dir == ast.RECV {
		sb.WriteString("<-")
	}
	sb.WriteString("chan")
	sb.WriteString(stringifyExpr(expr.Value))

	return sb.String()
}

func stringifyMapType(expr *ast.MapType) string {
	return fmt.Sprintf("map[%s]%s", stringifyExpr(expr.Key), stringifyExpr(expr.Value))
}

func stringifyInterfaceType(expr *ast.InterfaceType) string {
	sb := new(strings.Builder)

	sb.WriteString("interface{")
	for i, field := range expr.Methods.List {
		if i > 0 {
			sb.WriteString(";")
		}
		for j, name := range field.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}
		sb.WriteString(" ")
		sb.WriteString(stringifyExpr(field.Type))
	}
	sb.WriteString("}")

	return sb.String()
}

func stringifyStructType(expr *ast.StructType) string {
	sb := new(strings.Builder)

	sb.WriteString("struct{")
	for i, field := range expr.Fields.List {
		if i > 0 {
			sb.WriteString(";")
		}
		for j, name := range field.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}
		sb.WriteString(" ")
		sb.WriteString(stringifyExpr(field.Type))
	}
	sb.WriteString("}")

	return sb.String()
}

func stringifyFuncType(expr *ast.FuncType) string {
	sb := new(strings.Builder)

	sb.WriteString("func(")

	// args
	for i, arg := range expr.Params.List {
		if i > 0 {
			sb.WriteString(",")
		}
		for j, name := range arg.Names {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(name.Name)
		}
		sb.WriteString(" ")
		sb.WriteString(stringifyExpr(arg.Type))
	}
	sb.WriteString(")")

	// result
	if expr.Results != nil {
		rb := new(strings.Builder)

		needParens := false
		if len(expr.Results.List) > 1 {
			needParens = true
		}

		for i, rslt := range expr.Results.List {
			if i > 0 {
				rb.WriteString(",")
			}

			for j, name := range rslt.Names {
				needParens = true
				if j > 0 {
					rb.WriteString(",")
				}
				rb.WriteString(name.Name)
				rb.WriteString(" ")
			}

			rb.WriteString(stringifyExpr(rslt.Type))
		}

		if needParens {
			sb.WriteString("(")
		}
		sb.WriteString(rb.String())
		if needParens {
			sb.WriteString(")")
		}
	}

	return sb.String()
}