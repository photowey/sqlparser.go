/*
 * Copyright © 2024 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"fmt"

	"github.com/photowey/sqlparser.go/lexer"
)

type Token struct {
	Type    lexer.TokenType
	Literal string
}

type Tokenizer struct {
	Tokens   []Token
	position int
}

func NewTokenizer(tokens []Token) *Tokenizer {
	return &Tokenizer{Tokens: tokens, position: 0}
}

func (t *Tokenizer) Next() Token {
	if t.position >= len(t.Tokens) {
		return Token{Type: lexer.TokenEOF} // 处理结束标记
	}
	token := t.Tokens[t.position]
	t.position++

	// printToken(token)

	return token
}

func printToken(token Token) {
	fmt.Printf("Next token: Type=%v, Literal='%s'\n", token.Type, token.Literal)
}

func (t *Tokenizer) Peek() Token {
	if t.position >= len(t.Tokens) {
		return Token{Type: lexer.TokenEOF} // 处理结束标记
	}
	return t.Tokens[t.position]
}

func (t *Tokenizer) HasNext() bool {
	return t.position < len(t.Tokens)
}

func tokenize(sql string) []string {
	var tokens []string
	var current []rune

	insideQuotes := false
	quoteChar := rune(0)

	for _, ch := range sql {
		if insideQuotes {
			if ch == quoteChar {
				insideQuotes = false
				current = append(current, ch)
				tokens = append(tokens, string(current))
				current = []rune{}
			} else {
				current = append(current, ch)
			}
		} else {
			switch ch {
			case ' ', '\t', '\n', '\r':
				if len(current) > 0 {
					tokens = append(tokens, string(current))
					current = []rune{}
				}
			case '(', ')', ',', ';', '=':
				if len(current) > 0 {
					tokens = append(tokens, string(current))
					current = []rune{}
				}
				tokens = append(tokens, string(ch))
			case '\'', '"':
				if len(current) > 0 {
					tokens = append(tokens, string(current))
					current = []rune{}
				}
				insideQuotes = true
				quoteChar = ch
				current = append(current, ch)
			default:
				current = append(current, ch)
			}
		}
	}

	if len(current) > 0 {
		tokens = append(tokens, string(current))
	}

	return tokens
}
