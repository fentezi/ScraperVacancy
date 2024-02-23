FROM ubuntu:latest
LABEL authors="serou"

ENTRYPOINT ["top", "-b"]