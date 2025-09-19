# Pterodactyl Prometheus Exporter

A Prometheus exporter for Pterodactyl panel metrics.

## Installation

Go to your Pterodactyl Account page and create a new API key, it should start with `ptlc_`.

Then, you can run the exporter using Docker:
```yaml
services:
  pterodactyl-exporter:
    image: cr.bl19.dev/exporters/pterodactyl-exporter-go:latest
    restart: unless-stopped
    ports:
      - "9531:9531"
    environment:
      - PANEL_URL=<panel_url>
      - CLIENT_API_TOKEN=ptlc_xxxxxxxxxxxxxxxxxxxxxxxxxx
```

Replace `<panel_url>` with your Pterodactyl panel URL (e.g., `https://panel.example.com`).
Replace `ptlc_xxxxxxxxxxxxxxxxxxxxxxxxxx` with your actual API key from the Pterodactyl panel (`https://panel.example.com/account/api`).

## Configuration
The exporter can be configured using the following environment variables:
- `PANEL_URL`: The URL of your Pterodactyl panel (e.g., `https://panel.example.com`).
- `CLIENT_API_TOKEN`: Your Pterodactyl API key (e.g., `ptlc_xxxxxxxxxxxxxxxxxxxxxxxxxx`).

## Metrics
The exporter exposes the following metrics:
- `pterodactyl_running_servers`: Number of running servers per node.
- `pterodactyl_total_servers`: Total number of servers per node.
- `pterodactyl_server_cpu_absolute`: CPU usage percentage of each server.
- `pterodactyl_server_max_cpu_absolute`: Maximum CPU allocation percentage of each server.
- `pterodactyl_server_disk_megabytes`: Disk usage in megabytes of each server.
- `pterodactyl_server_max_disk_megabytes`: Maximum disk allocation in megabytes of each server.
- `pterodacytl_server_memory_megabytes`: Memory usage in megabytes of each server.
- `pterodactyl_server_max_memory_megabytes`: Maximum memory allocation in megabytes of each server.
- `pterodactyl_server_max_swap_megabytes`: Maximum swap allocation in megabytes of each server.
- `pterodactyl_server_network_rx_megabytes`: Network received in megabytes of each server.
- `pterodactyl_server_network_tx_megabytes`: Network transmitted in megabytes of each server.
- `pterodactyl_server_uptime_milliseconds`: Uptime in milliseconds of each server.

### Example Metrics Output

Here is an example of the metrics output:

```
# HELP pterodactyl_running_servers Number of running servers
# TYPE pterodactyl_running_servers gauge
pterodactyl_running_servers{node="Node 1"} 2
pterodactyl_running_servers{node="Node 2"} 2
pterodactyl_running_servers{node="Node 3"} 2
# HELP pterodactyl_server_cpu_absolute Absolute cpu usage by server
# TYPE pterodactyl_server_cpu_absolute gauge
pterodactyl_server_cpu_absolute{id="026d5ba0",node="Node 2",server_name="Server 1"} 0
pterodactyl_server_cpu_absolute{id="02f80e67",node="Node 3",server_name="Server 2"} 9.284
pterodactyl_server_cpu_absolute{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 0
pterodactyl_server_cpu_absolute{id="11ed6cbf",node="Node 2",server_name="Server 4"} 0
pterodactyl_server_cpu_absolute{id="180f05d2",node="Node 2",server_name="Server 5"} 0
pterodactyl_server_cpu_absolute{id="2627cbfa",node="Node 3",server_name="Server 6"} 0
pterodactyl_server_cpu_absolute{id="28bef150",node="Node 1",server_name="Server 7"} 12.681
pterodactyl_server_cpu_absolute{id="2d5d3b17",node="Node 1",server_name="Server 8"} 10.482
pterodactyl_server_cpu_absolute{id="2f149d7a",node="Node 1",server_name="Server 9"} 0
# HELP pterodactyl_server_disk_megabytes Disk space used by server in megabytes
# TYPE pterodactyl_server_disk_megabytes gauge
pterodactyl_server_disk_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 18.167935
pterodactyl_server_disk_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 6448.259044
pterodactyl_server_disk_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 856.052619
pterodactyl_server_disk_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 709.637238
pterodactyl_server_disk_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 224.018085
pterodactyl_server_disk_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 4161.279375
pterodactyl_server_disk_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 5735.567281
pterodactyl_server_disk_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 908.039774
pterodactyl_server_disk_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 307.775894
# HELP pterodactyl_server_max_cpu_absolute Maximum cpu load allowed to server
# TYPE pterodactyl_server_max_cpu_absolute gauge
pterodactyl_server_max_cpu_absolute{id="026d5ba0",node="Node 2",server_name="Server 1"} 50
pterodactyl_server_max_cpu_absolute{id="02f80e67",node="Node 3",server_name="Server 2"} 800
pterodactyl_server_max_cpu_absolute{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 300
pterodactyl_server_max_cpu_absolute{id="11ed6cbf",node="Node 2",server_name="Server 4"} 200
pterodactyl_server_max_cpu_absolute{id="180f05d2",node="Node 2",server_name="Server 5"} 600
pterodactyl_server_max_cpu_absolute{id="2627cbfa",node="Node 3",server_name="Server 6"} 600
pterodactyl_server_max_cpu_absolute{id="28bef150",node="Node 1",server_name="Server 7"} 600
pterodactyl_server_max_cpu_absolute{id="2d5d3b17",node="Node 1",server_name="Server 8"} 600
pterodactyl_server_max_cpu_absolute{id="2f149d7a",node="Node 1",server_name="Server 9"} 400
# HELP pterodactyl_server_max_disk_megabytes Maximum disk space allocated to server in megabytes
# TYPE pterodactyl_server_max_disk_megabytes gauge
pterodactyl_server_max_disk_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 128
pterodactyl_server_max_disk_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 32768
pterodactyl_server_max_disk_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 8192
pterodactyl_server_max_disk_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 8192
pterodactyl_server_max_disk_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 12288
pterodactyl_server_max_disk_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 40960
pterodactyl_server_max_disk_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 10240
pterodactyl_server_max_disk_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 8192
pterodactyl_server_max_disk_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 4096
# HELP pterodactyl_server_max_memory_megabytes Maximum memory allocated to server in megabytes
# TYPE pterodactyl_server_max_memory_megabytes gauge
pterodactyl_server_max_memory_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 512
pterodactyl_server_max_memory_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 16384
pterodactyl_server_max_memory_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 6144
pterodactyl_server_max_memory_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 4096
pterodactyl_server_max_memory_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 8192
pterodactyl_server_max_memory_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 32768
pterodactyl_server_max_memory_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 16384
pterodactyl_server_max_memory_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 4096
pterodactyl_server_max_memory_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 1024
# HELP pterodactyl_server_max_swap_megabytes Maximum swap allocated to server in megabytes
# TYPE pterodactyl_server_max_swap_megabytes gauge
pterodactyl_server_max_swap_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 0
pterodactyl_server_max_swap_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 0
pterodactyl_server_max_swap_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 0
pterodactyl_server_max_swap_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 0
pterodactyl_server_max_swap_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 0
pterodactyl_server_max_swap_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 0
pterodactyl_server_max_swap_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 0
pterodactyl_server_max_swap_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 0
pterodactyl_server_max_swap_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 0
# HELP pterodactyl_server_network_rx_megabytes Megabytes received by server via network
# TYPE pterodactyl_server_network_rx_megabytes gauge
pterodactyl_server_network_rx_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 0
pterodactyl_server_network_rx_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 144.355624
pterodactyl_server_network_rx_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 0
pterodactyl_server_network_rx_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 0
pterodactyl_server_network_rx_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 0
pterodactyl_server_network_rx_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 0
pterodactyl_server_network_rx_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 347.458931
pterodactyl_server_network_rx_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 1573.739269
pterodactyl_server_network_rx_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 0
# HELP pterodactyl_server_network_tx_megabytes Megabytes transmitted by server via network
# TYPE pterodactyl_server_network_tx_megabytes gauge
pterodactyl_server_network_tx_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 0
pterodactyl_server_network_tx_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 112.635329
pterodactyl_server_network_tx_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 0
pterodactyl_server_network_tx_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 0
pterodactyl_server_network_tx_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 0
pterodactyl_server_network_tx_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 0
pterodactyl_server_network_tx_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 3951.565632
pterodactyl_server_network_tx_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 7.215089
pterodactyl_server_network_tx_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 0
# HELP pterodactyl_server_uptime_milliseconds Server uptime in milliseconds
# TYPE pterodactyl_server_uptime_milliseconds gauge
pterodactyl_server_uptime_milliseconds{id="026d5ba0",node="Node 2",server_name="Server 1"} 0
pterodactyl_server_uptime_milliseconds{id="02f80e67",node="Node 3",server_name="Server 2"} 6.92637051e+08
pterodactyl_server_uptime_milliseconds{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 0
pterodactyl_server_uptime_milliseconds{id="11ed6cbf",node="Node 2",server_name="Server 4"} 0
pterodactyl_server_uptime_milliseconds{id="180f05d2",node="Node 2",server_name="Server 5"} 0
pterodactyl_server_uptime_milliseconds{id="2627cbfa",node="Node 3",server_name="Server 6"} 0
pterodactyl_server_uptime_milliseconds{id="28bef150",node="Node 1",server_name="Server 7"} 2.8455646587e+10
pterodactyl_server_uptime_milliseconds{id="2d5d3b17",node="Node 1",server_name="Server 8"} 2.3457115423e+10
pterodactyl_server_uptime_milliseconds{id="2f149d7a",node="Node 1",server_name="Server 9"} 0
# HELP pterodactyl_total_servers Number of total servers
# TYPE pterodactyl_total_servers gauge
pterodactyl_total_servers{node="Node 1"} 9
pterodactyl_total_servers{node="Node 2"} 11
pterodactyl_total_servers{node="Node 3"} 5
pterodactyl_total_servers{node="Node 4"} 9
pterodactyl_total_servers{node="Node 5"} 1
# HELP pterodacytl_server_memory_megabytes Memory used by server in megabytes
# TYPE pterodacytl_server_memory_megabytes gauge
pterodacytl_server_memory_megabytes{id="026d5ba0",node="Node 2",server_name="Server 1"} 0
pterodacytl_server_memory_megabytes{id="02f80e67",node="Node 3",server_name="Server 2"} 7135.232
pterodacytl_server_memory_megabytes{id="0f8b7cb4",node="Node 1",server_name="Server 3"} 0
pterodacytl_server_memory_megabytes{id="11ed6cbf",node="Node 2",server_name="Server 4"} 0
pterodacytl_server_memory_megabytes{id="180f05d2",node="Node 2",server_name="Server 5"} 0
pterodacytl_server_memory_megabytes{id="2627cbfa",node="Node 3",server_name="Server 6"} 0
pterodacytl_server_memory_megabytes{id="28bef150",node="Node 1",server_name="Server 7"} 12685.336576
pterodacytl_server_memory_megabytes{id="2d5d3b17",node="Node 1",server_name="Server 8"} 649.306112
pterodacytl_server_memory_megabytes{id="2f149d7a",node="Node 1",server_name="Server 9"} 0
```