#!/usr/bin/env bats

@test "invoke install command - install latest with --json" {
  cd cmd/cli/
  run go run main.go install --json
  echo "status = ${status}"
  echo "output trace = ${output}"
    [ "$status" -eq 0 ]
}

@test "invoke status -j command - output = '{"status":"stopped","installed-versions":["latest"]}'" {
  cd cmd/cli/
  run go run main.go status -j
  echo "status = ${status}"
  echo "output trace = ${output}"
  [ "$output" = '{"status":"stopped","installed-versions":["latest"]}' ]
  [ "$status" -eq 0 ]
}

@test "invoke start command - Start dockerhub images (latest)" {
  cd cmd/cli/
  run go run main.go start -t latest
  echo "status = ${status}"
  echo "output trace = ${output}"
  [ "$status" -eq 0 ]
}