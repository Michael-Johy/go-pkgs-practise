module github.com/Michael-Johy/go-pkgs-practise

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.5.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20210326220855-61e056675ecf // indirect
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181 // indirect
	google.golang.org/genproto v0.0.0-20210325224202-eed09b1b5210 // indirect
	google.golang.org/grpc v1.36.1 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.0.5
	gorm.io/gorm v1.21.6
)

replace (
	github.com/Michael-Johy/go-pkgs-practise/gin/conf => ../go-pkgs-practise/gin/conf
	github.com/Michael-Johy/go-pkgs-practise/gin/models => ../go-pkgs-practise/gin/models
	github.com/Michael-Johy/go-pkgs-practise/gin/models/blog => ../go-pkgs-practise/gin/models/blog

	github.com/Michael-Johy/go-pkgs-practise/gin/pkg => ../go-pkgs-practise/gin/pkg
	github.com/Michael-Johy/go-pkgs-practise/gin/pkg/e => ../go-pkgs-practise/gin/pkg/e
	github.com/Michael-Johy/go-pkgs-practise/gin/pkg/setting => ../go-pkgs-practise/gin/pkg/setting
	github.com/Michael-Johy/go-pkgs-practise/gin/pkg/util => ../go-pkgs-practise/gin/pkg/util

	github.com/Michael-Johy/go-pkgs-practise/gin/routers => ../go-pkgs-practise/gin/routers
	github.com/Michael-Johy/go-pkgs-practise/gin/routers/api/v1v => ./gin/routers/api/v1v

	github.com/Michael-Johy/go-pkgs-practise/grpc/proto/proto/grpc => ../go-pkgs-practise/grpc/proto/proto/grpc

)
