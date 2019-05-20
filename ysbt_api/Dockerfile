from golang:latest as builder
workdir /app/

copy go.mod .
copy go.sum .

run go mod download

copy ysbt_api ysbt_api/

run make -C ysbt_api build

from scratch
copy --from=builder /app/ysbt_api/ysbt_api /run/ysbt_api
entrypoint ["/run/ysbt_api"]
expose 8080
