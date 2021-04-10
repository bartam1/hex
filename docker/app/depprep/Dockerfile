FROM golang:1.15.11-buster

ENV USER=appuser
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
COPY shorter/ /shorter
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/shorter

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group# Copy our static executable.
COPY --from=builder /go/bin/shorter /go/bin/shorter # Use an unprivileged user.
USER appuser:appuser# Run the hello binary.
ENTRYPOINT ["/go/bin/shorter"]