#!/bin/sh

go build -o monitor-bg ./cmd/monitor-bg

mv monitor-bg /usr/local/bin/monitor-bg

chmod +x /usr/local/bin/monitor-bg

mkdir /etc/korupt-monitor

cp ./config/dummy_client.yaml /etc/korupt-monitor/client.yaml

cp ./launch.plist ~/Library/LaunchAgents/com.korupt.monitor-bg.plist

launchctl load ~/Library/LaunchAgents/com.korupt.monitor-bg.plist

echo "Successfully setup monitor-bg"
