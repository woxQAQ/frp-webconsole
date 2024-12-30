#!/bin/bash
set -e

# 1. 生成原始 swagger 文档
swag init -g api/meta.go -o api

# 2. 转换为 OpenAPI 3.0 并处理
swagger2openapi -y api/swagger.yaml > api/openapi_bundle.yaml

schemas=$(yq -o=yaml eval '.components.schemas | keys' api/openapi_bundle.yaml | yq -r '.[0]')

# 3. 获取所有 tags
tags=$(yq -r '.paths[].[].tags[]' api/openapi_bundle.yaml | sort -u)

for tag in $tags; do
    echo "Generating config for tag: $tag"
    mkdir -p "api/${tag}"
    cat > "api/${tag}/oapi-codegen.yaml" << EOF
package: ${tag}
output: pkg/controller/${tag}/gen.go
generate:
  gin-server: true
  models: true
import-mapping:
  ModelsSystem: github.com/woxQAQ/frp-webconsole/pkg/models.System
output-options:
  skip-prune: true
  exclude-schemas:
  $(echo "$schemas" | sed 's/^/  - /')
EOF
    echo "Generating code for tag: $tag"
    mkdir -p "pkg/controller/${tag}"
    go run gen.go \
    -f api/openapi_bundle.yaml \
    -o pkg/controller/${tag}/gen.go \
    -c api/${tag}/oapi-codegen.yaml
    rm -rf api/${tag}
done

rm -rf api/openapi*.yaml