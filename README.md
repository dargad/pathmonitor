# pathmonitor

This is a purely educational attempt to practice Go skills. I needed a daemon
that will monitor some directories in my filesystem and run a command if a new
file is detected in monitored locations.

I am aware that it is also possible to implement this with Python, a bash
script or in a hundred different technologies, but my primary goal was to
practice my newly obtained Go skills to do something practical for my own use.

# Usage

Eventually the daemon will be controlled by a systemd service and a config
file under /etc/pathmonitor.yaml. Example unit file and config file are in
the source code.
