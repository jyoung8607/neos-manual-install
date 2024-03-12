#!/usr/bin/bash

set -e

DEFAULT_GITHUB_OWNER="commaai"
DEFAULT_GITHUB_BRANCH="release2"

cd /data
rm -rf openpilot
time git clone https://github.com/${1:-$DEFAULT_GITHUB_OWNER}/openpilot.git -b ${2:-$DEFAULT_GITHUB_BRANCH} openpilot --recurse-submodules --depth 1

cd /data/data/com.termux/files
echo $'#!/usr/bin/bash\n\ncd /data/openpilot\n./launch_openpilot.sh\n' > continue.sh
chmod +x continue.sh
echo "Install complete, rebooting..."
reboot
