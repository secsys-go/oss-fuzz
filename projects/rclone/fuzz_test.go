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

package serve

import (
         "github.com/rclone/rclone/fstest/mockobject"
)
import (
         "net/http/httptest"
)
import (
         "testing"
)

func FuzzTestObjectBadRange(f *testing.F) {
	f.Add("aFile", "0123456789")
	f.Fuzz(func(t *testing.T, data1 string, data2 string) {
        w := httptest.NewRecorder()
        r := httptest.NewRequest("GET", "http://example.com/aFile", nil)
        r.Header.Add("Range", "xxxbytes=3-5")
        o := mockobject.New(data1).WithContent([]byte(data2), mockobject.SeekModeNone)
        Object(w, r, o)
        w.Result()
	})
}
