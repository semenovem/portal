#!/bin/sh

[ -z "$S3_API_URL" ] && echo "error: s3 url not set" && exit 1
[ -z "$S3_ACCESS_KEY" ] && echo "error: s3 access key not set" && exit 1
[ -z "$S3_SECRET_KEY" ] && echo "error: s3 secret key not set" && exit 1

echo "
{
  \"version\": \"10\",
  \"aliases\": {
    \"local\": {
      \"url\": \"https://${S3_API_URL}\",
      \"accessKey\": \"${S3_ACCESS_KEY}\",
      \"secretKey\": \"${S3_SECRET_KEY}\",
      \"api\": \"s3v4\",
      \"path\": \"auto\"
    }
  }
}
" >"${HOME}/.mc/config.json"

while true; do
  mc --insecure mb --ignore-existing "local/${S3_BUCKET_NAME}" && break
  echo "warn: unsuccessful attempt to create a bucket"
  sleep 1
done
