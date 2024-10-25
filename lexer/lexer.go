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

package lexer

// ----------------------------------------------------------------

type TokenType int
type DataType int

type ConstraintType int
type ColumnConstraintType int

// ----------------------------------------------------------------

const (
	TokenEOF TokenType = iota

	TokenIdentifier
	TokenDataType

	TokenLeftParen
	TokenRightParen
	TokenComma
	TokenSemicolon

	TokenCreate
	TokenTable
	TokenIf
	TokenNot
	TokenExists

	TokenCharset
	TokenCharacter
	TokenSet

	TokenEngine
	TokenCollate
	TokenComment

	TokenPrimary
	TokenUnique
	TokenForeign
	TokenIndex
	TokenKey
	TokenUnsigned
	TokenZerofill
	TokenAutoIncrement
	TokenNull
	TokenNotNull
	TokenDefault
	TokenCurrentTimestamp
	TokenOn
	TokenUpdate
	TokenReferences

	TokenConstraint
	TokenPartition
	TokenBy

	TokenCheck
)

// ----------------------------------------------------------------

const (
	TokenKeywordCreate           = "CREATE"
	TokenKeywordTable            = "TABLE"
	TokenKeywordIf               = "IF"
	TokenKeywordNot              = "NOT"
	TokenKeywordExists           = "EXISTS"
	TokenKeywordPrimary          = "PRIMARY"
	TokenKeywordUnique           = "UNIQUE"
	TokenKeywordForeign          = "FOREIGN"
	TokenKeywordReferences       = "REFERENCES"
	TokenKeywordKey              = "KEY"
	TokenKeywordUnsigned         = "UNSIGNED"
	TokenKeywordZerofill         = "ZEROFILL"
	TokenKeywordNull             = "NULL"
	TokenKeywordAutoIncrement    = "AUTO_INCREMENT"
	TokenKeywordDefault          = "DEFAULT"
	TokenKeywordOn               = "ON"
	TokenKeywordUpdate           = "UPDATE"
	TokenKeywordCurrentTimestamp = "CURRENT_TIMESTAMP"
	TokenKeywordCharset          = "CHARSET"
	TokenKeywordCharacter        = "CHARACTER"
	TokenKeywordCollate          = "COLLATE"
	TokenKeywordComment          = "COMMENT"
	TokenKeywordPartition        = "PARTITION"
	TokenKeywordBy               = "BY"
	TokenKeywordIndex            = "INDEX"
	TokenKeywordConstraint       = "CONSTRAINT"
	TokenKeywordSet              = "SET"
	TokenKeywordCheck            = "CHECK"
	TokenKeywordEngine           = "ENGINE"
	TokenKeywordComma            = ","
	TokenKeywordLeftParen        = "("
	TokenKeywordRightParen       = ")"
)

// ----------------------------------------------------------------

const (
	TokenBit DataType = iota
	TokenTinyInt
	TokenSmallInt
	TokenMediumInt
	TokenInt
	TokenBigint

	// ----------------------------------------------------------------

	TokenFloat
	TokenDouble
	TokenDecimal

	// ----------------------------------------------------------------

	TokenChar
	TokenVarchar

	TokenTinyText
	TokenText
	TokenMediumText
	TokenLongText

	// ----------------------------------------------------------------

	TokenDate
	TokenTime
	TokenDateTime
	TokenTimestamp
	TokenYear

	// ----------------------------------------------------------------

	TokenTinyBlob
	TokenBlob
	TokenMediumBlob
	TokenLongBlob

	// ----------------------------------------------------------------

	TokenEnum
	TokenCollectionSet

	// ----------------------------------------------------------------

	TokenGeometry
	TokenPoint
	TokenLineString
	TokenPolygon
	TokenMultiPoint
	TokenMultiLineString
	TokenMultiPolygon
	TokenGeometryCollection

	// ----------------------------------------------------------------

	TokenJson
)

// ----------------------------------------------------------------

const (
	ColumnConstraintPrimary ConstraintType = iota
	ColumnConstraintForeign
	ColumnConstraintUnique
)

const (
	ColumnConstraintNotNull ConstraintType = iota + 3
	ColumnConstraintDefault
	ColumnConstraintCheck
	ColumnConstraintUnsigned
	ColumnConstraintZerofill
)

var (
	MySQLDataTypes = []string{
		"TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT",
		"FLOAT", "DOUBLE", "DECIMAL",
		"CHAR", "VARCHAR",
		"TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT",
		"DATE", "TIME", "DATETIME", "TIMESTAMP", "YEAR",
		"TINYBLOB", "BLOB", "MEDIUMBLOB", "LONGBLOB",
		"ENUM", "SET",
		"GEOMETRY", "POINT", "LINESTRING", "POLYGON", "MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON", "GEOMETRYCOLLECTION",
	}
)
