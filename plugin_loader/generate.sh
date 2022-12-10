go build   -buildmode=plugin -gcflags="all=-N -l" -o=plugin.so plugin.go

go build   -gcflags="all=-N -l"