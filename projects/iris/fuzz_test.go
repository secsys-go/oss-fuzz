// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package iris_test

import (
         "testing"
         "github.com/kataras/iris/v12"
         "github.com/kataras/iris/v12/httptest"
)

func FuzzTestUseRouterParentDisallow(f *testing.F) {
	f.Add("no_userouter_allowed", "always", "_2", "_3", "/index", "/", "/user")
	f.Fuzz(func(t *testing.T, data1 string, data2 string, data3 string, data4 string, data5 string, data6 string, data7 string) {
			app := iris.New()
			app.UseRouter(func(ctx iris.Context) {
					ctx.WriteString(data2)
					ctx.Next()
			})
			app.Get(data5, func(ctx iris.Context) {
					ctx.WriteString(data1)
			})

			app.SetPartyMatcher(func(ctx iris.Context, p iris.Party) bool {
					// modifies the PartyMatcher to not match any UseRouter,
					// tests should receive the handlers response alone.
					return false
			})

			app.PartyFunc(data6, func(p iris.Party) { // it's the same instance of app.
					p.UseRouter(func(ctx iris.Context) {
							ctx.WriteString(data3)
							ctx.Next()
					})
					p.Get(data6, func(ctx iris.Context) {
							ctx.WriteString(data1)
					})
			})

			app.PartyFunc(data7, func(p iris.Party) {
					p.UseRouter(func(ctx iris.Context) {
							ctx.WriteString(data4)
							ctx.Next()
					})

					p.Get(data6, func(ctx iris.Context) {
							ctx.WriteString(data1)
					})
			})

			e := httptest.New(t, app)
			e.GET(data5)
			e.GET(data6)
			e.GET(data7)

	})
}
