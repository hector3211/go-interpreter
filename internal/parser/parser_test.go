package parser

import (
	"go-interpreter/internal/ast"
	"go-interpreter/internal/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	// incorrectInput := `
	//    let x = 5;
	//    let = 10;
	//
	//    let 8080;
	//    `
	input := `
    let x = 5;
    let y = 10;

    let foobar = 8080;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("prgram.Statements returned less than 3 statements\n got: %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("does not equal 'let'")
		return false
	}

	letstmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got: %T", s)
		return false
	}

	if letstmt.Name.Value != name {
		t.Errorf("expected: %s got: %s", letstmt.Name.Value, name)
		return false
	}

	return true
}
