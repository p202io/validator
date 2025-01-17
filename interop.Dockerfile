FROM gcr.io/prysmaticlabs/build-agent AS builder

WORKDIR /workspace

COPY . /workspace/.

# Build binaries for minimal configuration.
RUN bazel build --define network=minimal --jobs=auto --remote_cache= \
  //cmd/beacon-chain:beacon-chain \
  //cmd/validator:validator \
  //tools/interop/convert-keys \
  //tools/bootnode


FROM gcr.io/whiteblock/base:ubuntu1804

COPY --from=builder /workspace/bazel-bin/beacon-chain/linux_amd64_stripped/beacon-chain .
COPY --from=builder /workspace/bazel-bin/validator/linux_amd64_pure_stripped/validator .
COPY --from=builder /workspace/bazel-bin/tools/interop/convert-keys/linux_amd64_stripped/convert-keys .

RUN mkdir /launch

COPY hack/interop_start.sh /launch/start.sh

ENTRYPOINT ["start.sh"]
