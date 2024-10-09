# neos-manual-install

Stuck here? You're not alone.

![Screenshot](neos-installer-stuck.jpg)

This repo contains a downloadable, executable script that bypasses NEOS Setup to install openpilot. This appears to be necessary as of late February 2024; the reason is not yet understood. This script is made specifically for NEOS devices (EON, comma two). This is NOT designed or needed for AGNOS devices such as comma three or threex.

# Basic Usage

Installing stock openpilot 0.8.13.1.

1. Connect to Wi-Fi normally.
2. When connected, go to More Options.
3. Touch the triple-dot icon in the upper right corner, select Advanced.
4. Scroll down and note the IPv4 address, will look like "192.168.202.191".
5. Download and save the [NEOS default/setup SSH key](https://raw.githubusercontent.com/commaai/openpilot/master/system/hardware/tici/id_rsa) to your machine as "id_rsa".
    * This is different from any GitHub SSH key you may have.
6. Connect to your EON/C2 using that key.
    * Command line: `ssh -i [pathname of saved id_rsa] comma@[your-c2-ip-address]`
    * Your browser may have saved the id_rsa file as "id_rsa.txt" or "id_rsa.pem"
    * If you get a "bad permissions" error, `chmod 600 [pathname of saved id_rsa]` and try again
7. Once connected, paste this command into your SSH session: `curl -Ls https://tinyurl.com/bdhse3xn | bash -s`

# Advanced Usage (Installing a Custom Fork)

Let's be honest, we're here because you were tempted to uninstall openpilot and try some other random fork!
The installer script in Step 7 above will accept two optional parameters, a GitHub repository owner and a branch name. This mirrors the format used by comma's install generator.

For a custom install URL such as: `installer.comma.ai/ErichMoraga/813`

Instead use this command: `curl -Ls https://tinyurl.com/bdhse3xn | bash -s ErichMoraga 813`
