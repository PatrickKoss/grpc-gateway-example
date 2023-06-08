#!/bin/bash

# create swagger yaml
protoc -I="." --openapiv2_out json_names_for_fields=false:./api --openapiv2_opt disable_default_responses=true,output_format=yaml,logtostderr=true,allow_merge=true,merge_file_name=merged internal/adapter/api/openapi.proto internal/adapter/api/domain/proto/message.proto internal/adapter/api/domain/proto/error.proto internal/adapter/api/v1/user/user.proto internal/adapter/api/v1/product/product.proto

# convert swagger.yaml to openapi.yaml (version 3)
# from https://github.com/OpenAPITools/openapi-generator
mkdir ./openapi
openapi-generator generate -i ./api/merged.swagger.yaml -g openapi-yaml --minimal-update -o ./openapi
mv ./openapi/openapi/openapi.yaml ./api/openapi.yaml
rm -rf ./openapi

# remove swagger 2.0 stuff
sed -i '' '/url:/d' ./api/openapi.yaml
sed -i '' '/servers:/d' ./api/openapi.yaml
sed -i '' '/.*localhost:3000.*/d' ./api/openapi.yaml
sed -i '' '/.*x-original-swagger.*/d' ./api/openapi.yaml
sed -i '' '/.*ErrorService.*/d' ./api/openapi.yaml

# add openapi stuff
bearerAuth="components:\n  securitySchemes:\n    bearerAuth:\n      type: http\n      scheme: bearer\n      bearerFormat: JWT"
sed -i "" "s|___|]|g" ./api/openapi.yaml
sed -i "" "s|__|[|g" ./api/openapi.yaml
sed -i "" "s|components:|$bearerAuth|g" ./api/openapi.yaml

echo "
servers:
  - url: 'https://api.prd.example.com'
    description: prod
  - url: 'https://api.stg.example.com'
    description: stg
  - url: 'https://api.dev.example.com'
    description: dev
security:
  - bearerAuth: []
" >> ./api/openapi.yaml
