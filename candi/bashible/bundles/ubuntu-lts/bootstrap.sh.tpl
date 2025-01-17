{{- /*
# Copyright 2021 Flant JSC
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
*/}}
#!/bin/bash
export LANG=C
if ! type jq 2>/dev/null || ! type curl 2>/dev/null || ! type nc 2>/dev/null; then
  apt update
  export DEBIAN_FRONTEND=noninteractive
  until apt install jq netcat-openbsd curl -y; do
    echo "Error installing packages"
    apt update
    sleep 10
  done
fi

mkdir -p /var/lib/bashible/
