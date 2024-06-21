#!/usr/bin/env bash
openapi-generator generate \
  -i ./main/pkg/swagger/swagger.yml \
  -g go \
  --additional-properties packageName=jabalisdkgo,packageVersion=0.0.1,useTags=true \
  --git-user-id barrygear \
  --git-repo-id jabali
