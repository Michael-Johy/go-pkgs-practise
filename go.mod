module github.com/Michael-Johy/go-pkgs-practise

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/spf13/viper v1.7.1
	github.com/unknwon/com v1.0.1
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

)
