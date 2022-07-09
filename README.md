# Hello HTTPD

Minimal HTTPD server and a Dockerfile to build a minimal (smallest) Docker
image with that server. Using `upx --brute` allows to get the app binary even
smaller but takes more time to execute.

The server listens on port 8080 or on port defined in `HELLO_HTTPD_PORT`
environment variable.
