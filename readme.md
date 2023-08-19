# Edge-Engine: Container-Based Edge Computing Software for IoT Application


**Description:**
Edge-Engine is a powerful container-based edge computing software designed to support the deployment of IoT applications in edge environments. It facilitates seamless management of edge devices and applications from a cloud platform, enhancing data processing, analysis, and visualization capabilities. The software provides robust support for various communication protocols, including MQTT, OPC UA, Modbus TCP/RTU, and REST API, allowing efficient data collection and streaming. Edge-Engine integrates NATS for data streaming and analysis, enabling the creation of sophisticated data analyzer applications. The processed results are then seamlessly updated to the cloud platform for the construction of interactive dashboards and real-time monitoring.

## Key Features:

1. **Containerization:** Edge-Engine leverages container technology for efficient application deployment, isolation, and scalability across edge devices, ensuring consistent behavior regardless of underlying hardware.

2. **Cloud-Managed Deployment:** The software offers centralized management capabilities from a cloud platform, allowing administrators to remotely deploy, update, and monitor applications across a fleet of edge devices.

3. **Protocol Support:**
   - **MQTT:** Edge-Engine supports MQTT protocol for seamless communication between edge devices and applications, enabling efficient data collection and control.
   - **OPC UA:** The software provides built-in OPC UA support for secure and reliable exchange of data and information between heterogeneous industrial devices.
   - **Modbus TCP/RTU:** Edge-Engine facilitates communication with legacy industrial devices through Modbus TCP/RTU, ensuring compatibility with a wide range of equipment.
   - **REST API:** RESTful API support enables easy integration with web services and allows edge devices to expose their capabilities for external access.

4. **Data Collection:**
   - **MQTT Data Collection:** Edge-Engine efficiently collects data from MQTT-enabled devices, aggregating information from multiple sources.
   - **OPC UA and Modbus:** The software facilitates seamless data retrieval from OPC UA and Modbus-enabled devices, simplifying integration with industrial machinery and processes.
   - **REST API Integration:** Edge-Engine enables data collection from external systems and services through RESTful API interactions.

5. **Data Streaming with NATS:**
   - Edge-Engine integrates NATS for high-performance, real-time data streaming and messaging across edge devices and applications.
   - NATS supports the creation of data analyzer applications for complex event processing, analytics, and pattern recognition at the edge.

6. **Cloud Integration:**
   - Edge-Engine seamlessly synchronizes processed data and analysis results with the cloud platform for further analysis, visualization, and storage.
   - Data-driven insights obtained from edge processing can be used to build interactive dashboards and real-time monitoring interfaces.

7. **Security and Reliability:**
   - The software prioritizes security by offering secure communication protocols and encryption mechanisms to safeguard sensitive data in transit.
   - Edge-Engine ensures high availability and reliability through container orchestration and failover mechanisms.

## Benefits:

- **Efficient Edge Computing:** Edge-Engine optimizes data processing by executing applications directly on edge devices, reducing latency and bandwidth usage.
- **Flexible Deployment:** The container-based architecture provides flexibility in deploying and managing applications across various edge devices and locations.
- **Integration Versatility:** Edge-Engine's protocol support ensures seamless integration with a diverse set of IoT devices, legacy equipment, and web services.
- **Advanced Analytics:** Integration with NATS enables the creation of data analyzer applications, supporting real-time analytics and insights at the edge.
- **Cloud Synergy:** The software bridges the gap between edge and cloud environments, enabling data synchronization and dashboard creation for comprehensive insights.

## Roadmap:
- [ ] Send a secure command via MQTT and HTTP, CLI
- [ ] Data Ingressing with MQTT,OPC UA, Modbus
- [ ] Support Machine Leaning Application
