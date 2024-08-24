FROM go:1.22

COPY . .

RUN go build

CMD ./memory_wall