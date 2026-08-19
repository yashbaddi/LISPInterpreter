package lexer

import (
	"errors"
	"testing"

	"github.com/yashbaddi/golisp/internal/token"
)

func TestNewLexer(t *testing.T) {
	input := "(foo 123 \"bar\")"
	l := NewLexer(input)
	if l == nil {
		t.Fatal("expected non-nil Lexer")
	}
	if string(l.input) != input {
		t.Errorf("expected input %q, got %q", input, string(l.input))
	}
	if l.pos != 0 {
		t.Errorf("expected pos 0, got %d", l.pos)
	}

	// Test with UTF-8 non-ASCII characters
	utf8Input := "(⌘ \"日本語\")"
	lUTF8 := NewLexer(utf8Input)
	if len(lUTF8.input) != len([]rune(utf8Input)) {
		t.Errorf("expected rune slice length %d, got %d", len([]rune(utf8Input)), len(lUTF8.input))
	}
}

func TestNextToken(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
		wantErr  bool
	}{
		{
			name:  "parentheses and whitespace",
			input: " ( ) ",
			expected: []token.Token{
				{Type: token.LPAREN, Literal: "("},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.EOF, Literal: ""},
			},
			wantErr: false,
		},
		{
			name:  "string literal",
			input: `"hello world"`,
			expected: []token.Token{
				{Type: token.STRING, Literal: "hello world"},
				{Type: token.EOF, Literal: ""},
			},
			wantErr: false,
		},
		{
			name:  "number literal",
			input: "12345",
			expected: []token.Token{
				{Type: token.NUMBER, Literal: "12345"},
				{Type: token.EOF, Literal: ""},
			},
			wantErr: false,
		},
		{
			name:  "boolean literals",
			input: "#t #f",
			expected: []token.Token{
				{Type: token.BOOLEAN, Literal: "#t"},
				{Type: token.BOOLEAN, Literal: "#f"},
				{Type: token.EOF, Literal: ""},
			},
			wantErr: false,
		},
		{
			name:  "identifier",
			input: "foo",
			expected: []token.Token{
				{Type: token.IDENTIFIER, Literal: "foo"},
				{Type: token.EOF, Literal: ""},
			},
			wantErr: false,
		},
		{
			name:  "complex lisp expression",
			input: "(define x 42)",
			expected: []token.Token{
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENTIFIER, Literal: "define"},
				{Type: token.IDENTIFIER, Literal: "x"},
				{Type: token.NUMBER, Literal: "42"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.EOF, Literal: ""},
			},
			wantErr: false,
		},
		{
			name:    "invalid token error",
			input:   "@",
			wantErr: true,
		},
		{
			name:    "unclosed string error",
			input:   `"unclosed string`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			for i, exp := range tt.expected {
				tok, err := l.NextToken()
				if (err != nil) != tt.wantErr {
					t.Fatalf("step %d: NextToken() error = %v, wantErr %v", i, err, tt.wantErr)
				}
				if tt.wantErr {
					return
				}
				if tok != exp {
					t.Errorf("step %d: got token %+v, want %+v", i, tok, exp)
				}
			}
			if tt.wantErr {
				_, err := l.NextToken()
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			}
		})
	}
}

func TestSkipWhiteSpace(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedPos int
	}{
		{
			name:        "leading space tab newline",
			input:       "  \t\n\ra",
			expectedPos: 5,
		},
		{
			name:        "only whitespace",
			input:       "   \t  ",
			expectedPos: 6,
		},
		{
			name:        "empty string",
			input:       "",
			expectedPos: 0,
		},
		{
			name:        "no leading whitespace",
			input:       "abc",
			expectedPos: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			l.skipWhiteSpace()
			if l.pos != tt.expectedPos {
				t.Errorf("expected pos %d, got %d", tt.expectedPos, l.pos)
			}
		})
	}
}

func TestReadInteger(t *testing.T) {
	t.Run("valid integer", func(t *testing.T) {
		l := NewLexer("123abc")
		val, err := l.readInteger()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val != "123" {
			t.Errorf("expected %q, got %q", "123", val)
		}
		if l.pos != 3 {
			t.Errorf("expected pos 3, got %d", l.pos)
		}
	})

	t.Run("not a number", func(t *testing.T) {
		l := NewLexer("abc")
		val, err := l.readInteger()
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if val != "" {
			t.Errorf("expected empty string, got %q", val)
		}
		var lexErr LexerError
		if !errors.As(err, &lexErr) {
			t.Errorf("expected LexerError type, got %T", err)
		}
		if lexErr.Pos != 0 || lexErr.Msg != "expected a number" {
			t.Errorf("unexpected LexerError details: %+v", lexErr)
		}
	})

	t.Run("at EOF", func(t *testing.T) {
		l := NewLexer("")
		val, err := l.readInteger()
		if err == nil {
			t.Fatal("expected error at EOF, got nil")
		}
		if val != "" {
			t.Errorf("expected empty string, got %q", val)
		}
	})
}

func TestReadString(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		wantErr     bool
		expectedMsg string
	}{
		{
			name:     "plain string",
			input:    `hello world"`,
			expected: "hello world",
			wantErr:  false,
		},
		{
			name:     "escaped characters",
			input:    `\" \\ \b \f \n \r \t"`,
			expected: "\" \\ \b \f \n \r \t",
			wantErr:  false,
		},
		{
			name:     "unicode escape valid ASCII hex",
			input:    `\u0041"`,
			expected: "A",
			wantErr:  false,
		},
		{
			name:     "unicode escape non-ASCII snowman",
			input:    `\u2603"`,
			expected: "☃",
			wantErr:  false,
		},
		{
			name:        "unexpected EOF in string",
			input:       `hello world`,
			wantErr:     true,
			expectedMsg: "Unexpected End of file",
		},
		{
			name:        "unexpected EOF after backslash",
			input:       `hello\`,
			wantErr:     true,
			expectedMsg: "Unexpected End of file",
		},
		{
			name:        "short unicode escape sequence",
			input:       `\u00"`,
			wantErr:     true,
			expectedMsg: "Invalid Hex Value",
		},
		{
			name:        "invalid hex in unicode escape sequence",
			input:       `\uGGGG"`,
			wantErr:     true,
			expectedMsg: "Invalid Hex Value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			str, err := l.readString()
			if (err != nil) != tt.wantErr {
				t.Fatalf("readString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				var lexErr LexerError
				if !errors.As(err, &lexErr) {
					t.Fatalf("expected LexerError, got %T (%v)", err, err)
				}
				if tt.expectedMsg != "" && lexErr.Msg != tt.expectedMsg {
					t.Errorf("expected error msg %q, got %q", tt.expectedMsg, lexErr.Msg)
				}
				return
			}
			if str != tt.expected {
				t.Errorf("got %q, want %q", str, tt.expected)
			}
		})
	}
}

func TestReadIdentifier(t *testing.T) {
	t.Run("valid identifier with space", func(t *testing.T) {
		l := NewLexer("foo 123")
		id, err := l.readIdentifier()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != "foo" {
			t.Errorf("expected %q, got %q", "foo", id)
		}
		if l.pos != 3 {
			t.Errorf("expected pos 3, got %d", l.pos)
		}
	})

	t.Run("valid alphanumeric identifier", func(t *testing.T) {
		l := NewLexer("foo123")
		id, err := l.readIdentifier()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != "foo123" {
			t.Errorf("expected %q, got %q", "foo123", id)
		}
		if l.pos != 6 {
			t.Errorf("expected pos 6, got %d", l.pos)
		}
	})

	t.Run("empty identifier if not starting with letter", func(t *testing.T) {
		l := NewLexer("123")
		id, err := l.readIdentifier()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if id != "" {
			t.Errorf("expected empty string, got %q", id)
		}
		if l.pos != 0 {
			t.Errorf("expected pos 0, got %d", l.pos)
		}
	})
}

func TestCurrent(t *testing.T) {
	l := NewLexer("a⌘c")
	if l.current() != 'a' {
		t.Errorf("expected 'a', got %c", l.current())
	}
	l.advance()
	if l.current() != '⌘' {
		t.Errorf("expected '⌘', got %c", l.current())
	}
	l.advance()
	if l.current() != 'c' {
		t.Errorf("expected 'c', got %c", l.current())
	}
}

func TestAdvance(t *testing.T) {
	l := NewLexer("abc")
	if l.pos != 0 {
		t.Errorf("expected pos 0, got %d", l.pos)
	}
	l.advance()
	if l.pos != 1 {
		t.Errorf("expected pos 1, got %d", l.pos)
	}
	l.advance()
	if l.pos != 2 {
		t.Errorf("expected pos 2, got %d", l.pos)
	}
}

func TestAdvanceN(t *testing.T) {
	t.Run("advance within bounds", func(t *testing.T) {
		l := NewLexer("hello world")
		l.advanceN(5)
		if l.pos != 5 {
			t.Errorf("expected pos 5, got %d", l.pos)
		}
	})

	t.Run("advance out of bounds clamped to len", func(t *testing.T) {
		l := NewLexer("hi")
		l.advanceN(10)
		if l.pos != 2 {
			t.Errorf("expected pos clamped to 2, got %d", l.pos)
		}
	})
}

func TestEOF(t *testing.T) {
	l := NewLexer("a")
	if l.eof() {
		t.Errorf("expected eof to be false at start")
	}
	l.advance()
	if !l.eof() {
		t.Errorf("expected eof to be true after advancing past length")
	}
	l.advance()
	if !l.eof() {
		t.Errorf("expected eof to stay true when pos > len")
	}
}

func TestLexerError(t *testing.T) {
	err := LexerError{
		Pos: 42,
		Msg: "syntax error",
	}
	expectedStr := "Lexical Error: 42: syntax error"
	if err.Error() != expectedStr {
		t.Errorf("expected %q, got %q", expectedStr, err.Error())
	}
}
