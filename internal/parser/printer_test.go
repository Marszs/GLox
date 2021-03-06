package parser

import (
	"GLox/internal/scanner/token"
	"testing"
)

func TestPrinter_Print(t *testing.T) {

	printer := new(Printer)

	expr1 := NewBinary(NewLiteral(1),
		token.NewToken(token.PLUS, "+", nil, 1),
		NewLiteral(2))
	expected := "(+ 1 2)"
	if result := printer.Print(expr1); result != expected {
		t.Fatalf("expected %s, but got %s", expected, result)
	}

	expr2 := NewBinary(
		NewUnary(token.NewToken(token.BANG, "!", nil, 1), NewLiteral(123)),
		token.NewToken(token.STAR, "*", nil, 1),
		NewGrouping(NewLiteral("45.1")))
	expected = "(* (! 123) (group 45.1))"
	if result := printer.Print(expr2); result != expected {
		t.Fatalf("expected %s, but got %s", expected, result)
	}
}
