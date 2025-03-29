#! /usr/bin/bash

set -e

OUTPUT_FILE_NAME=sysd

go build -o $OUTPUT_FILE_NAME ./cmd/sysd

export PATH=${PATH}:${PWD}/${OUTPUT_FILE_NAME}

# Adding root as user:group as handling systemd requires root level access.
sudo chown root:root ${OUTPUT_FILE_NAME}

# Adding current logged-in user's UID to the executable so the user have all the root level access for this binary.
sudo chmod u+s ${OUTPUT_FILE_NAME}

# Making binary executable for root user.
sudo chmod +x ${OUTPUT_FILE_NAME}
