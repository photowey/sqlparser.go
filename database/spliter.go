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

package database

import (
	"strings"
)

func SplitSQLStatements(sql string) []string {
	var statements []string
	var buf strings.Builder

	inString := false
	for _, char := range sql {
		if char == '\'' || char == '"' {
			inString = !inString
		}
		if char == ';' && !inString {
			statements = append(statements, buf.String())
			buf.Reset()
		} else {
			buf.WriteRune(char)
		}
	}

	if buf.Len() > 0 {
		statements = append(statements, buf.String())
	}

	return statements
}
