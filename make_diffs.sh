# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# this script is a helper to track differences from each step

mkdir -p diff
diff        main.go step01/main.go > diff/01.go
diff step01/main.go step02/main.go > diff/12.go
diff step02/main.go step03/main.go > diff/23.go
diff step03/main.go step04/main.go > diff/34.go
diff step04/main.go step05/main.go > diff/45.go
diff step05/main.go step06/main.go > diff/56.go
diff step06/main.go step07/main.go > diff/67.go
# diff step07/main.go step08/main.go > diff/78.go