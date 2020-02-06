#! bash
dist=pkg

date
rm -rf $dist
GOOS=linux GOARCH=amd64 go build -o $dist/main main.go
zip -r $dist/main.zip $dist/main public