#!/bin/bash

# Fail on anything
set -e

./hack/update-generated-completions.sh
./hack/update-generated-docs.sh
./hack/update-generated-openapi.sh
./hack/update-generated-deep-copies.sh
./hack/update-generated-protobuf.sh
./hack/update-generated-conversions.sh
