aws lambda update-function-configuration \
  --function-name point-service-batch \
  --environment Variables="{APP_ENV=aws,REQUEST_TO_SERVER_URL_SCHEMA=http,REQUEST_TO_SERVER_URL_PATH=app.pointservice.internal:1323,REQUEST_TO_SERVER_METHOD=GET}" \
  --timeout 30
