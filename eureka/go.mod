module github.com/tietang/go-eureka-client/eureka

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
	golang.org/x/net => github.com/golang/net v0.0.0-20190327025741-74e053c68e29
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
	golang.org/x/text => github.com/golang/text v0.3.0
)

require (
	github.com/go-ini/ini v1.57.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/tietang/go-utils v0.1.3
	github.com/tietang/props/v3 latest
	github.com/valyala/fasttemplate v1.1.0 // indirect
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/ini.v1 v1.42.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
)
