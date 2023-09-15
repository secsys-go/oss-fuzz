#!/bin/bash -eu
# Copyright 2020 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################

# required by Go 1.20
export CXX="${CXX} -lresolv"

mv $SRC/fuzzMarshalJSON.go $SRC/tidb/types/
mv $SRC/fuzzNewBitLiteral.go $SRC/tidb/types/
mv $SRC/fuzzNewHexLiteral.go $SRC/tidb/types/
mv $SRC/fuzzBufferIsolation.go $SRC/tidb/br/pkg/membuf/

compile_go_fuzzer github.com/pingcap/tidb/types FuzzUnmarshalJSON fuzzUnmarshalJSON
compile_go_fuzzer github.com/pingcap/tidb/types FuzzNewBitLiteral fuzzNewBitLiteral
compile_go_fuzzer github.com/pingcap/tidb/types FuzzNewHexLiteral fuzzNewHexLiteral

go mod tidy
printf "package membuf\nimport _ \"github.com/AdamKorcz/go-118-fuzz-build/testing\"\n" > $SRC/tidb/br/pkg/membuf/register.go
go mod tidy
compile_native_go_fuzzer github.com/pingcap/tidb/br/pkg/membuf FuzzTestBufferIsolation fuzzTestBufferIsolation