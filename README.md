# korupt-monitor
Monitoring Use(full)less data

## 📋 Prerequisites

- Docker installed
- MacOS (for client setup)
- Git

## 🛠 Installation

### Server Setup

1. Clone the repository
```bash
git clone https://github.com/KoruptTinker/korupt-monitor.git
cd korupt-monitor
```

2. Configure the server
```bash
# Copy the example config file
cp ./config/dummy_server.yaml ./config/prod.yaml

# Edit the config file with your settings
nano config.yaml
```

3. Build and run Docker container
```bash
# Build the Docker image
docker build -t monitor-server .

# Run the container (replace 8080 with your configured port)
docker run -d -p 7070:7070 monitor-server
```

### Client Setup (MacOS Only) ⚠️

1. Clone the repository
```bash
git clone https://github.com/KoruptTinker/korupt-monitor.git
cd korupt-monitor
```

2. Make the setup script executable and run it
```bash
chmod +x setup.sh
./setup.sh
```

3. Follow the on-screen instructions to complete the setup


## ⚠️ Important Notes

- The client setup currently only supports MacOS

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📮 Support

For support, please open an issue in the GitHub repository.

---
