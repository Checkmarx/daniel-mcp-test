# Vulnerable Dockerfile for IaC scanning demo
FROM ubuntu:latest

# 1. Use of latest tag (not pinned to a specific version)
# 2. Running as root user (default in Ubuntu images)
# 3. Exposing an insecure port

RUN apt-get update && apt-get install -y curl

# Insecure: Exposing a sensitive port
EXPOSE 2375

# Insecure: Adding a file with world-writable permissions
RUN echo 'echo Hello, world!' > /entrypoint.sh && chmod 777 /entrypoint.sh

CMD ["/bin/bash", "/entrypoint.sh"]