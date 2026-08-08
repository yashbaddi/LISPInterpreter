package token

import "testing"

func TestTokenTypes(t *testing.T) {
	tokens := []struct {
		tokenType TokenType
		expected  TokenType
	}{
		{LPAREN, 0},
		{RPAREN, 1},
		{STRING, 2},
		{NUMBER, 3},
		{IDENTIFIER, 4},
		{ILLEGAL, 5},
		{EOF, 6},
	}

	for _, tt := range tokens {
		if tt.tokenType != tt.expected {
			t.Errorf("expected TokenType value %d, got %d", tt.expected, tt.tokenType)
		}
	}
}

func TestTokenCreation(t *testing.T) {
	tests := []struct {
		name     string
		token    Token
		wantType TokenType
		wantLit  string
	}{
		{
			name:     "LPAREN token",
			token:    Token{Type: LPAREN, Literal: "("},
			wantType: LPAREN,
			wantLit:  "(",
		},
		{
			name:     "RPAREN token",
			token:    Token{Type: RPAREN, Literal: ")"},
			wantType: RPAREN,
			wantLit:  ")",
		},
		{
			name:     "STRING token",
			token:    Token{Type: STRING, Literal: "hello"},
			wantType: STRING,
			wantLit:  "hello",
		},
		{
			name:     "NUMBER token",
			token:    Token{Type: NUMBER, Literal: "123"},
			wantType: NUMBER,
			wantLit:  "123",
		},
		{
			name:     "IDENTIFIER token",
			token:    Token{Type: IDENTIFIER, Literal: "foo"},
			wantType: IDENTIFIER,
			wantLit:  "foo",
		},
		{
			name:     "ILLEGAL token",
			token:    Token{Type: ILLEGAL, Literal: "@"},
			wantType: ILLEGAL,
			wantLit:  "@",
		},
		{
			name:     "EOF token",
			token:    Token{Type: EOF, Literal: ""},
			wantType: EOF,
			wantLit:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.token.Type != tt.wantType {
				t.Errorf("got Type %v, want %v", tt.token.Type, tt.wantType)
			}
			if tt.token.Literal != tt.wantLit {
				t.Errorf("got Literal %q, want %q", tt.token.Literal, tt.wantLit)
			}
		})
	}
}
