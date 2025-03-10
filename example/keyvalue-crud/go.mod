module github.com/jamesstocktonj1/wasi-lib/example/keyvalue-crud

go 1.23.0

toolchain go1.23.7

replace github.com/jamesstocktonj1/wasi-lib => ../../

require (
	github.com/jamesstocktonj1/wasi-lib v0.0.0-20250308110010-3e2202d9ad48
	github.com/julienschmidt/httprouter v1.3.0
	go.bytecodealliance.org v0.5.0
	go.bytecodealliance.org/cm v0.1.0
	go.wasmcloud.dev/component v0.0.5
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/regclient/regclient v0.7.2 // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/samber/slog-common v0.17.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	github.com/urfave/cli/v3 v3.0.0-beta1 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)
