aws lambda update-function-configuration \
  --function-name my-go-function \
  --environment Variables="{APP_ENV=aws,REQUEST_TO_SERVER_URL_SCHEMA=http,REQUEST_TO_SERVER_URL_PATH=app.pointservice.internal,REQUEST_TO_SERVER_METHOD=GET}"