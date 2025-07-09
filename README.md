# C2 NEOS Alternate Fix Install

Stuck on the NEOS setup screen? You're not alone.

![Screenshot](neos-installer-stuck.jpg)

This repository provides a simple, all-in-one tool to bypass the NEOS Setup screen and install openpilot on your comma two or EON. This is NOT designed or needed for AGNOS devices like the comma three.

## Usage

This tool can be run on Windows or macOS.

### Prerequisites

1.  **Connect to Wi-Fi**
    *   On your comma device, connect to your Wi-Fi network normally.

2.  **Find Your Device's IP Address**
    *   On the device, go to **More Options**.
    *   Touch the triple-dot icon in the upper right corner and select **Advanced**.
    *   Scroll down and note the **IPv4 address**. It will look like `192.168.x.x`.

### Windows Instructions

1.  **Download the Installer**
    *   [**Click here to download `c2-neos-alt-fix-install.exe`**](https://github.com/ophwug/c2-neos-alt-fix-install/releases/latest/download/c2-neos-alt-fix-install.exe)
    *   Save the file to a convenient location, like your Downloads folder.

2.  **Run the Installer**
    *   Find the `c2-neos-alt-fix-install.exe` file you downloaded.
    *   **Double-click** the file to run it.
    *   The application will open a window and guide you through the rest of the process.
    *   When the process is finished, the window will stay open until you press the Enter key.

### macOS Instructions

1.  **Run the Installer**
    *   Open the **Terminal** application on your Mac.
    *   Copy and paste the following command into the Terminal and press Enter. This will download, make executable, and run the installer in one step.
        ```bash
        curl -L https://github.com/ophwug/c2-neos-alt-fix-install/releases/latest/download/c2-neos-alt-fix-install-darwin -o c2-neos-alt-fix-install-darwin && chmod +x c2-neos-alt-fix-install-darwin && ./c2-neos-alt-fix-install-darwin
        ```
    *   The application will then guide you through the rest of the process.

## For Developers

If you want to build the application yourself:

1.  Clone this repository.
2.  Make sure you have Go installed (version 1.22 or later).
3.  Run `make` to build both the Windows and macOS executables.

The build process is automated via GitHub Actions. Every push to the `main` branch will trigger a new build and update the "Latest Build" release.

## Credits

This project is a Go-based evolution of the original shell script installer created by [jyoung8607](https://github.com/jyoung8607). A big thank you for the original work and inspiration!
