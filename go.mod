module github.com/go-gluon/example

go 1.17

require (
	github.com/go-gluon/gluon v0.0.0
	github.com/go-gluon/zerolog v0.0.0
	github.com/go-gluon/fiber v0.0.0
	github.com/gofiber/fiber/v2 v2.19.0
	github.com/rs/zerolog v1.22.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/andybalholm/brotli v1.0.2 // indirect
	github.com/klauspost/compress v1.13.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.29.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
)

replace (
	github.com/go-gluon/fiber => ../fiber
	github.com/go-gluon/gluon => ../gluon
	github.com/go-gluon/zerolog => ../zerolog
)