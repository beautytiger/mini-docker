#!/usr/bin/env bash
# 此脚本用于从kafka引入其他集群的日志
# 运行前需保证当前集群中已经安装好elasticsearch
set -o errexit
set -o nounset
set -o pipefail

DIR=$(cd $(dirname "${BASH_SOURCE[0]}") >/dev/null 2>&1 && pwd)
cd $DIR

find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
