#!/usr/bin/bash

GITHUB_OWNER="commaai"
GITHUB_BRANCH="release2"

cd /data
git clone https://github.com/$GITHUB_OWNER/openpilot.git -b $GITHUB_BRANCH --recurse-submodules --depth 1

cd /data/data/com.termux/files
echo $'#!/usr/bin/bash\n\ncd /data/openpilot\n./launch_openpilot.sh\n" > continue.sh
chmod +x continue.sh
reboot
