package parser

import (
	"go-interpreter/internal/ast"
	"go-interpreter/internal/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;

    let foobar = 8080;
    `

	l := lexer.NewLexer(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parseProgram() returned nil")
	}

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
