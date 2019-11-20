#!/bin/bash
cd ./integration-tests
bats setup.bats
bats integration.bats
bats teardown.bats
