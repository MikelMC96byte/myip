#!/bin/bash
workdir="$(pwd)"
buildir="${workdir}/build"
sourcedir="${workdir}/src"

packagename="myip"

archs=("amd64" "arm7")
oss=("windows" "linux")

rm -rf ${buildir}
mkdir -p ${buildir}

# Compilation for Windows AMD64
env GOOS=windows go build -o ${buildir}/windows/amd64/${packagename}.exe ${sourcedir}
cd ${buildir}/windows/amd64
zip -r ${workdir}/build/${packagename}-windows-amd64.zip *
cd ${workdir}

# Compilation for Linux AMD64
env GOOS=linux go build -o ${buildir}/linux/amd64/${packagename} ${sourcedir}
cd ${buildir}/linux/amd64
zip -r ${buildir}/${packagename}-linux-amd64.zip *
cd ${workdir}

# Compilation for Linux ARMv7
env GOOS=linux GOARCH=arm GOARM=7 go build -o ${buildir}/linux/arm7/${packagename} ${sourcedir}
cd ${buildir}/linux/arm7
zip -r ${workdir}/build/${packagename}-linux-arm7.zip *
cd ${workdir}