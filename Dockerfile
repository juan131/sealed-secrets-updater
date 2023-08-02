FROM gcr.io/distroless/static-debian11:latest
LABEL maintainer "Juan Ariza <jariza@vmware.com>"

USER 1001

ARG TARGETARCH
COPY dist/sealed-secrets-updater_linux_${TARGETARCH}*/sealed-secrets-updater /usr/local/bin/

ENTRYPOINT ["sealed-secrets-updater"]
