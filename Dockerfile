FROM ubuntu:latest as builder

WORKDIR /workspace

RUN apt-get update && apt-get install -y curl bash git

RUN echo "echo quack" > duck.sh && chmod +x duck.sh

FROM alpine:latest

COPY --from=builder /workspace/duck.sh .

# RUN apt-get update && apt-get install -y git
RUN apt-get update && apt-get install -y ssh

# RUN apt-get update && apt-get install -y curl bash git

CMD [ "sh", "duck.sh" ]
