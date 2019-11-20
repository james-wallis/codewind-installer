#!/bin/bash
# ############################
# # Removal command tests  # #
# ############################

# @test "invoke stop-all command - Stop dockerhub images (latest)" {
#   cd cmd/cli/
#   run go run main.go stop-all
#   echo "status = ${status}"
#   echo "output trace = ${output}"
#   [ "$status" -eq 0 ]
# }

# @test "invoke remove command - remove all dockerhub images" {
#   cd cmd/cli/
#   run go run main.go remove
#   echo "status = ${status}"
#   echo "output trace = ${output}"
#   [ "$status" -eq 0 ]
# }