package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTextNextToken(t *testing.T) {
	input := "bar foo"

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{TEXT, "bar"},
		{TEXT, "foo"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestEmptyInput(t *testing.T) {
	input := ``

	l := NewLexer(input)
	lval := &TdocSymType{}
	tok := l.Lex(lval)
	assert.Equal(t, 0, tok)
}

func TestComplexTextNextToken(t *testing.T) {
	input := `foo bar blubb
  baz
    quoo
    la
          le`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{TEXT, "foo"},
		{TEXT, "bar"},
		{TEXT, "blubb"},
		{TEXT, "baz"},
		{TEXT, "quoo"},
		{TEXT, "la"},
		{TEXT, "le"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestSimpleComponentNextToken(t *testing.T) {
	input := `cloud actor node`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "cloud"},
		{ERROR, "A component can't be next to a component - need an identifier first"},
		{COMPONENT, "node"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestSimpleMixNextToken(t *testing.T) {
	input := `cloud foo actor bar node duck`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "cloud"},
		{IDENTIFIER, "foo"},
		{COMPONENT, "actor"},
		{IDENTIFIER, "bar"},
		{COMPONENT, "node"},
		{IDENTIFIER, "duck"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestSingleQuoteIdentifier(t *testing.T) {
	input := `actor 'test foo'`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test foo"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestDoubleQuoteIdentifier(t *testing.T) {
	input := `actor "test foo"`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test foo"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestDoubleQuoteMultilineIdentifier(t *testing.T) {
	input := `actor "test
foo"`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test\nfoo"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestAliasDoubleQuoteMultilineIdentifier(t *testing.T) {
	input := `actor "test
foo" as foo`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test\nfoo"},
		{ALIAS, "as"},
		{TEXT, "foo"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestSimpleAliasDeclaration(t *testing.T) {
	input := `actor test as foo`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test"},
		{ALIAS, "as"},
		{TEXT, "foo"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestDigitContainingAndUnicodeAliasDeclaration(t *testing.T) {
	input := `actor test as fo12☂o`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test"},
		{ALIAS, "as"},
		{TEXT, "fo12☂o"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestAliasAsIdentifierDeclaration(t *testing.T) {
	input := `actor test as "foo 12"`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test"},
		{ALIAS, "as"},
		{ERROR, "Aliases are not allowed to be quoted"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestUnicodeMixNextToken(t *testing.T) {
	input := `cloud ✓ actor ✓ node`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "cloud"},
		{IDENTIFIER, "✓"},
		{COMPONENT, "actor"},
		{IDENTIFIER, "✓"},
		{COMPONENT, "node"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}

func TestDeclarationCombination(t *testing.T) {
	input := `actor "test for multiple words" as f✓o cloud 'and again' as bar☂`

	tests := []struct {
		expectedType    int
		expectedLiteral string
	}{
		{COMPONENT, "actor"},
		{IDENTIFIER, "test for multiple words"},
		{ALIAS, "as"},
		{TEXT, "f✓o"},
		{COMPONENT, "cloud"},
		{IDENTIFIER, "and again"},
		{ALIAS, "as"},
		{TEXT, "bar☂"},
	}

	l := NewLexer(input)
	for _, tt := range tests {
		lval := &TdocSymType{}
		tok := l.Lex(lval)
		assert.Equal(t, tt.expectedType, tok)
		assert.Equal(t, tt.expectedLiteral, lval.val)
	}
}