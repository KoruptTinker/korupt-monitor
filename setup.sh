#!/bin/sh

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to display header
header() {
  clear
  echo "${BLUE}"
  echo "#############################################"
  echo "#          KORUPT Monitor Installer         #"
  echo "#############################################"
  echo "${NC}"
  echo
}

# Function to display error messages
error() {
  echo "${RED}[ERROR] $1${NC}"
  exit 1
}

header

# Build the binary
echo "${YELLOW}Building monitor-bg executable...${NC}"
go build -o monitor-bg ./cmd/monitor-bg || error "Failed to build executable"

# Code signing
echo "${YELLOW}Signing executable...${NC}"
codesign -s - -f -v --timestamp --options runtime monitor-bg || error "Code signing failed"

# Install binary
echo "${YELLOW}Installing to /usr/local/bin...${NC}"
sudo mv monitor-bg /usr/local/bin/monitor-bg || error "Failed to move executable"
sudo chmod +x /usr/local/bin/monitor-bg || error "Failed to set executable permissions"

# Create config directory
echo "${YELLOW}Creating configuration directory...${NC}"
sudo mkdir -p /etc/korupt-monitor || error "Failed to create config directory"

# Get hostname from user
header
echo "${GREEN}Please enter the server hostname (e.g.: http://your-server:7070)${NC}"
read -p "Hostname: " SERVER_HOST

# Generate config file
echo "${YELLOW}Generating configuration file...${NC}"
sudo tee /etc/korupt-monitor/client.yaml >/dev/null <<EOF
external:
  korupt_monitor_server:
    hostname: ${SERVER_HOST}
EOF

# Install launch agent
header
echo "${YELLOW}Installing launch agent...${NC}"
LAUNCH_AGENT_PATH="$HOME/Library/LaunchAgents/com.korupt.monitor-bg.plist"
cp ./launch.plist "$LAUNCH_AGENT_PATH" || error "Failed to copy launch agent"

# Load service
load_service() {
  echo "${YELLOW}Attempting to load service...${NC}"
  launchctl load "$LAUNCH_AGENT_PATH"

  echo "${RED}Failed to load service.${NC}"
  echo "You may need to grant Accessibility permissions in:"
  echo "System Preferences → Security & Privacy → Privacy → Accessibility"
  echo "Click the lock icon and add monitor-bg to the list"

  while true; do
    read -p "${YELLOW}Did you grant the permissions? (Y/n/q): ${NC}" -n 1 -r
    echo # move to new line

    case $REPLY in
    [Yy]*)
      echo "${GREEN}Retrying service load...${NC}"
      launchctl unload "$LAUNCH_AGENT_PATH" 2>/dev/null
      launchctl load "$LAUNCH_AGENT_PATH" && return 0
      echo "${RED}Still failed to load service. Please check permissions.${NC}"
      ;;
    [Qq]*)
      error "User aborted installation"
      ;;
    *)
      echo "${YELLOW}Please grant accessibility permissions or press Q to quit${NC}"
      ;;
    esac
  done
}

header
echo "${YELLOW}Preparing service...${NC}"
# Initial load attempt
load_service || error "Failed to load service after permission grant"

header
echo "${GREEN}Installation completed successfully!${NC}"
echo
echo "Service configuration: /etc/korupt-monitor/client.yaml"
echo "Launch agent: $LAUNCH_AGENT_PATH"
echo
echo "You can manage the service with:"
echo "  launchctl start com.korupt.monitor-bg"
echo "  launchctl stop com.korupt.monitor-bg"
echo "  launchctl unload $LAUNCH_AGENT_PATH"
