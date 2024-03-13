FROM gcr.io/distroless/static:latest
COPY dist/server /app/server
ENTRYPOINT ["/app/server"]