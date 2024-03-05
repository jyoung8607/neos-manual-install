# neos-manual-install
Work around an issue with NEOS Setup no longer seeing Internet available.

This script is made specifically for NEOS devices (EON, comma two). Do not use this script on a comma three or threex device.

# Basic usage

Re-installing stock openpilot 0.8.13.1.

1. Connect to Wi-Fi normally.
2. When connected, go to More Options.
3. Touch the triple-dot icon in the upper right corner, select Advanced.
4. Scroll down and note the IPv4 address, will look like "192.168.202.191".
5. Download the [default/setup SSH key](https://github.com/commaai/openpilot/blob/master/tools/ssh/id_rsa).
6. Connect to that IP via the SSH client of your choice, or [Openpilot Toolkit](https://github.com/spektor56/OpenpilotToolkit). SSH is always enabled while still in NEOS Setup.
7. Paste the following command into your SSH session: `curl -Ls https://tinyurl.com/bdhse3xn | bash -e`

# Advanced usage

Let's be honest, we're here because you were tempted to uninstall openpilot and try some other random fork!
The installer script above will accept two optional parameters, a GitHub repository owner and a branch name. This mirrors the format used by comma's install generator.

Where you would normally use the install URL: `https://installer.comma.ai/commaai/release2`

Use this command instead: `curl -Ls https://tinyurl.com/bdhse3xn | bash -e commaai release2`
