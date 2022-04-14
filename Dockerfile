FROM golang:1.18-stretch as build

WORKDIR /go/src/github.com/CERIT-SC/nvidia-device-plugin
COPY . .

RUN export CGO_LDFLAGS_ALLOW='-Wl,--unresolved-symbols=ignore-in-object-files' && \
    go mod download && \
    go build -ldflags="-s -w" -o /go/bin/nvidia-device-plugin-v2 cmd/nvidia/main.go

RUN go build -o /go/bin/kubectl-inspect-nvidia-v2 cmd/inspect/*.go

FROM debian:bullseye-slim

ENV NVIDIA_VISIBLE_DEVICES=all
ENV NVIDIA_DRIVER_CAPABILITIES=utility

COPY --from=build /go/bin/nvidia-device-plugin-v2 /usr/bin/nvidia-device-plugin-v2

COPY --from=build /go/bin/kubectl-inspect-nvidia-v2 /usr/bin/kubectl-inspect-nvidia-v2

CMD ["nvidia-device-plugin-v2","-logtostderr"]
