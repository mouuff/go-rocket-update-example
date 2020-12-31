set format=binaries_{{.OS}}\{{.Dir}}_{{.OS}}_{{.Arch}}
gox -osarch windows/amd64 -output %format%
gox -osarch windows/386 -output %format%
gox -osarch linux/amd64 -output %format%
gox -osarch linux/386 -output %format%
gox -osarch darwin/amd64 -output %format%
