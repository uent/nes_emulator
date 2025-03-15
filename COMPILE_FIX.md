To fix the compilation errors:

1. The code changes removing the unused 'line' variable and fixing the color type are correct
2. However, to fully compile the project, you need to install X11 development libraries:

```bash
# On Debian/Ubuntu:
sudo apt-get install libx11-dev

# On Fedora/RHEL:
sudo dnf install libX11-devel

# On Alpine:
apk add libx11-dev
```

After installing the required dependencies, the compilation should succeed.