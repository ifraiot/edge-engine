#!/bin/bash

# Configuration
SERVICE_NAME="my-service"
SERVICE_EXECUTABLE="/path/to/your/service/executable"
SERVICE_CONFIG="/path/to/your/config/file"

# Function to install and start the service
install_and_start_service() {
    echo "Installing $SERVICE_NAME..."
    
    # Copy service executable and config
    cp "$SERVICE_EXECUTABLE" /usr/local/bin/
    cp "$SERVICE_CONFIG" /etc/$SERVICE_NAME/
    
    # Create service definition file
    echo "[Unit]
    Description=My Custom Service
    After=network.target
    
    [Service]
    ExecStart=/usr/local/bin/$SERVICE_NAME
    Restart=always
    
    [Install]
    WantedBy=multi-user.target" > /etc/systemd/system/$SERVICE_NAME.service
    
    # Reload Systemd to pick up new service definitions
    systemctl daemon-reload
    
    # Enable and start the service
    systemctl enable $SERVICE_NAME
    systemctl start $SERVICE_NAME
    
    echo "$SERVICE_NAME installed and started."
}

# Run the installation and start
install_and_start_service
