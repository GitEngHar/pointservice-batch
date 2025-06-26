aws lambda create-function \
  --function-name point-service-batch \
  --runtime provided.al2 \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --role xxx \
  --region ap-northeast-1 \
  --vpc-config SubnetIds=xxx,xxx,SecurityGroupIds=xxx