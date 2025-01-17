FROM alpine:3.20.3
RUN apk --no-cache add ca-certificates>=20220614-r0 bash~=5

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/provider" /usr/local/bin/provider

ENV USER_ID=65532

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
