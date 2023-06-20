FROM golang:1.18-bullseye AS build
# FROMPATTERN 1.18

COPY ./ /go-markdown-renderer/
WORKDIR /go-markdown-renderer
RUN go build -buildvcs=false

#-------------------------------------------------------------------------------
FROM stefanfritsch/baseimage_statup:20 AS app
# FROMPATTERN 20

COPY --from=build /go-markdown-renderer/go-markdown-renderer /usr/local/bin/go-markdown-renderer
COPY run.sh /etc/runit/runsvdir/default/go-markdown-renderer/run
COPY assets/ /markdown-renderer/assets/
COPY README.md /markdown-renderer/docs/README.md
COPY templates/ /markdown-renderer/templates/

WORKDIR /markdown-renderer

RUN  echo "markdown-renderer:7518215" | add-users /dev/stdin \
  && usermod -a -G docker_env markdown-renderer \
  && chmod +x /etc/runit/runsvdir/default/go-markdown-renderer/run \
  && chmod +x /usr/local/bin/go-markdown-renderer \
  && chown -R markdown-renderer:markdown-renderer /markdown-renderer

ENV PORT=8090
EXPOSE 8090
