#!/bin/bash

# WARNING: DO NOT EDIT, THIS FILE IS PROBABLY A COPY
#
# The original version of this file is located in the https://github.com/khulnasoft/common-files repo.
# If you're looking at this file in a different repo and want to make a change, please go to the
# common-files repo, make the change there and check it in. Then come back to this repo and run
# "make update-common".

# Copyright Khulnasoft Authors
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

if BUILD_GIT_REVISION=$(git rev-parse HEAD 2> /dev/null); then
  if [[ -z "${IGNORE_DIRTY_TREE}" ]] && [[ -n "$(git status --porcelain 2>/dev/null)" ]]; then
    BUILD_GIT_REVISION=${BUILD_GIT_REVISION}"-dirty"
  fi
else
  BUILD_GIT_REVISION=unknown
fi

# Check for local changes
tree_status="Clean"
if [[ -z "${IGNORE_DIRTY_TREE}" ]] && ! git diff-index --quiet HEAD --; then
  tree_status="Modified"
fi

GIT_DESCRIBE_TAG=$(git describe --tags --always)
HUB=${HUB:-"docker.io/khulnasoft"}

# used by common/scripts/gobuild.sh
echo "khulnasoft.com/khulnasoft/pkg/version.buildVersion=${VERSION:-$BUILD_GIT_REVISION}"
echo "khulnasoft.com/khulnasoft/pkg/version.buildGitRevision=${BUILD_GIT_REVISION}"
echo "khulnasoft.com/khulnasoft/pkg/version.buildStatus=${tree_status}"
echo "khulnasoft.com/khulnasoft/pkg/version.buildTag=${GIT_DESCRIBE_TAG}"
echo "khulnasoft.com/khulnasoft/pkg/version.buildHub=${HUB}"
echo "khulnasoft.com/khulnasoft/pkg/version.buildOS=${BUILD_GOOS}"
echo "khulnasoft.com/khulnasoft/pkg/version.buildArch=${BUILD_GOARCH}"
