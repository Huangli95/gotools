#!/bin/bash
function mod_env() {
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct
}

function x86() {
    mod_env
    x86_release
}

function arm() {
    mod_env
    arm_release
}

function x86_release() {
  # export GOARCH=x86
  # export GOOS=linux
  export CGO_ENABLED=0
  target_dir="./output/x86-linux/release"
  target_bin="${target_dir}/gotools"
  go build -ldflags "$flags" -o ${target_bin}
  if [[ $? != 0 ]]; then
    echo "build error"
    return 1
  else
    echo "build success"
    tar -cvzf ${target_dir}/gotools.tar.gz -C ${target_dir} gotools
    return 0
  fi
}

function x86_debug() {
  export CGO_ENABLED=0
  target_dir="./output/x86-linux/debug"
  target_bin="${target_dir}/gotools"
  flags="-X main.Dbg=true"
  go build -ldflags "$flags" -o ${target_bin}
  if [[ $? != 0 ]]; then
    echo "build error"
    return 1
  else
    echo "build success"
    tar -cvzf ${target_dir}/gotools.tar.gz -C ${target_dir} gotools
    return 0
  fi
}

function arm_release() {
  export GOARCH=arm64
  export GOOS=linux
  export CGO_ENABLED=0
  target_dir="./output/arm-linux/release"
  target_bin="${target_dir}/gotools"

  go build  -ldflags "$flags" -o ${target_bin}
  if [[ $? != 0 ]]; then
    echo "build error"
    return 1
  else
    echo "build success"
    tar -cvzf ${target_dir}/gotools.tar.gz -C ${target_dir} gotools
    return 0
  fi
}

function arm_debug() {
  export GOARCH=arm64
  export GOOS=linux
  export CGO_ENABLED=0
  target_dir="./output/arm-linux/debug"
  target_bin="${target_dir}/gotools"
  flags="-X main.Dbg=true"

  go build  -ldflags "$flags" -o ${target_bin}
  if [[ $? != 0 ]]; then
    echo "build error"
    return 1
  else
    echo "build success"
    tar -cvzf ${target_dir}/gotools.tar.gz -C ${target_dir} gotools
    return 0
  fi
}

function help() {
  echo "options: "
  echo "  x86           Build x86_release component"
  echo "  arm           Build arm_release component"
  echo "  x86_debug     Build x86_debug component"
  echo "  arm_debug     Build arm_debug component"
  echo "  example ./build.sh x86"
}

if [[ $# < 1 ]]; then
  help
fi

$@