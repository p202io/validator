FROM gcr.io/prysmaticlabs/build-agent AS builder

WORKDIR /workspace

COPY . /workspace/.

# Build binaries for minimal configuration.
RUN bazel build --define network=minimal --jobs=auto --remote_cache= \
  //cmd/beacon-chain:beacon-chain \
  //cmd/validator:validator \
  //tools/interop/convert-keys


FROM ubuntu:18.04

COPY --from=builder /workspace/bazel-bin/cmd/beacon-chain/beacon-chain_/beacon-chain /usr/local/bin/
COPY --from=builder /workspace/p202-data/ /p202-data/

ENTRYPOINT ["beacon-chain"]
