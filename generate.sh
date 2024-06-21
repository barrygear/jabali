#!/usr/bin/env bash
openapi-generator generate \
  -i https://raw.githubusercontent.com/barrygear/jabali-api/main/pkg/swagger/swagger.yml \
  -g go \
  --additional-properties packageName=jabalisdkgo,packageVersion=0.0.1,useTags=true \
  --git-user-id barrygear \
  --git-repo-id jabali-sdk-go
