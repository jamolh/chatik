FROM golang:1.16.1-alpine as builder

RUN apk add --no-cache libc6-compat
RUN apk add --no-cache git

# RUN export GOPRIVATE=git.alifpay.tj/terminals/*
# RUN git config --global \
#     url."https://alif-terminals:z3bh3_dWsuNUzzXW2unq@git.alifpay.tj".insteadOf "https://git.alifpay.tj" 

ENV GO111MODULE=on

WORKDIR /app

COPY .env .
COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify                                                                                                                                                                                                                               

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chatik .        

# FROM scratch
# COPY --from=builder /app/control-panel-api /app/
# COPY config.json .

EXPOSE 50001

CMD ["sh"]

ENTRYPOINT ["/app/chatik"]