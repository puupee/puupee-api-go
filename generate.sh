java -jar openapi-generator-cli.jar generate -g go -o ./ \
  -c config.json \
  -i swagger.json

git add .
git commit -a -m "update"
git push origin master