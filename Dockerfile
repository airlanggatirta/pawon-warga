FROM kitabisa/debian-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/airlanggatirta/pawon-warga"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/pawon-warga

WORKDIR /opt/pawon-warga

COPY ./bin/pawon-warga /opt/pawon-warga/

RUN chmod +x /opt/pawon-warga/pawon-warga

# Create appuser
RUN adduser --disabled-password --gecos '' pawon-warga
USER pawon-warga

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/pawon-warga/bin/pawon-warga"]
