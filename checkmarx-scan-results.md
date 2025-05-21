# Checkmarx Scan Results
Scan ID: d300c97f-6d06-41b4-8514-d6a5b00278f8

## Top IaC Vulnerabilities
1. **Missing User Instruction** (HIGH) - A user should be specified in the Dockerfile, otherwise the image will run as root. File: /Dockerfile, Line: 2
2. **Image Version Using 'latest'** (MEDIUM) - Images should be tagged with specific versions, not 'latest'. File: /Dockerfile, Line: 2
3. **Healthcheck Instruction Missing** (LOW) - The Dockerfile should contain a HEALTHCHECK instruction. File: /Dockerfile, Line: 2
4. **Apt Get Install Pin Version Not Defined** (LOW) - Installed packages should have a pinned version. File: /Dockerfile, Line: 8
5. **Apt Get Install Lists Were Not Deleted** (INFO) - After using apt-get install, the apt-get lists should be deleted. File: /Dockerfile, Line: 8
6. **APT-GET Not Avoiding Additional Packages** (INFO) - The '--no-install-recommends' flag should be used to avoid installing unnecessary packages. File: /Dockerfile, Line: 8

