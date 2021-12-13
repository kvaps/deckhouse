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

SYSTEM_PACKAGES="curl wget virt-what inotify-tools bash-completion lvm2 parted apt-transport-https sudo nfs-common"
KUBERNETES_DEPENDENCIES="iptables iproute2 socat util-linux mount ebtables ethtool conntrack"

bb-apt-install ${SYSTEM_PACKAGES} ${KUBERNETES_DEPENDENCIES}

bb-rp-install "jq:7bf9a38af84c5d14c4484d6cb53d5a562e48bd1e30618c8e82e62c32-1638042868105" "curl:39b235aae4e9f50990f1a7213cd2b6a9a9df5f849599de0a12e7d3f1-1638562274869"
