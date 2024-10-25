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

type Table struct {
	Database string
	Name     string
	Comment  string
	Columns  []*Column
	Indexes  []*Index
}

type Column struct {
	Name          string
	DataType      string
	Length        *int
	Precision     *int
	Scale         *int
	NotNull       bool
	AutoIncrement bool
	PrimaryKey    bool
	ForeignKey    bool
	Unique        bool
	Unsigned      bool
	Zerofill      bool
	Default       string
	Comment       string
}

type Index struct {
	Name    string
	Columns []string
}
