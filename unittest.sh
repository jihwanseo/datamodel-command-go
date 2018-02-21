###############################################################################
# Copyright 2018 Samsung Electronics All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
###############################################################################

#!/bin/bash
export GOPATH=$PWD
mkdir src
cp -r formatter/ src/

pkg_list=("formatter")

count=0
for pkg in "${pkg_list[@]}"; do
    go test -c -v -gcflags "-N -l" $pkg
    go test -coverprofile=$count.cover.out $pkg
    if [ $? -ne 0 ]; then
        echo "Unittest is failed."
        rm *.out *.test
        exit 1
    fi
    count=$count.0
done

echo "mode: set" > coverage.out && cat *.cover.out | grep -v mode: | sort -r | \
awk '{if($1 != last) {print $0;last=$1}}' >> coverage.out

go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverall.html

rm *.out *.test
rm -rf src
