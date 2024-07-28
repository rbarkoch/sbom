#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )


version="$1"

if [ -z "$version" ]; then
    echo 'No version number was provided'
    exit 1
fi

echo "Publishing version '$version'..."

echo "Building..."
pushd ./src
./build.sh
popd


mkdir ./artifacts
cp LICENSE.md ./artifacts/LICENSE.md
cp README.md ./artifacts/README.md

binaries=("sbom-linux-386" "sbom-linux-amd64" "sbom-windows-386.exe" "sbom-windows-amd64.exe")

for binary in ${binaries[@]}; do
    echo "Copying '$binary'"
    cp ./bin/$binary ./artifacts/$binary
done

branch="$(git rev-parse --abbrev-ref HEAD)"
commit="$(git rev-parse HEAD)"

pushd ./artifacts
./sbom-linux-amd64 init sbom \
    --name "SBOM" \
    --description "A tool for creating basic Software Bill-of-Materials (SBOM) files." \
    --author "Ronnie Bar-Kochba" \
    --version $version \
    --repository "https://github.com/rbarkoch/sbom" \
    --branch "$branch" \
    --commit "$commit" \
    --license "MIT" \
    --license-url "https://github.com/rbarkoch/sbom/blob/main/LICENSE.md" \
    --type "executable"

./sbom-linux-amd64 package add go \
    --name "The Go Programming Language" \
    --description "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software." \
    --version "$(go version)" \
    --license "BSD 4-Clause" \
    --license-url "https://github.com/golang/go/blob/master/LICENSE" \
    --repository "https://github.com/golang/go" \
    --type "build-tool"


zip -r ./sbom-$version.zip .
popd