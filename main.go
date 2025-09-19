package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	pterodactyl_running_servers = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_running_servers",
		Help: "Number of running servers",
	}, []string{"node"})
	pterodactyl_total_servers = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_total_servers",
		Help: "Number of total servers",
	}, []string{"node"})
	pterodacytl_server_memory_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodacytl_server_memory_megabytes",
		Help: "Memory used by server in megabytes",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_cpu_absolute = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_cpu_absolute",
		Help: "Absolute cpu usage by server",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_disk_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_disk_megabytes",
		Help: "Disk space used by server in megabytes",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_network_rx_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_network_rx_megabytes",
		Help: "Megabytes received by server via network",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_network_tx_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_network_tx_megabytes",
		Help: "Megabytes transmitted by server via network",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_uptime_milliseconds = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_uptime_milliseconds",
		Help: "Server uptime in milliseconds",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_max_memory_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_max_memory_megabytes",
		Help: "Maximum memory allocated to server in megabytes",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_max_swap_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_max_swap_megabytes",
		Help: "Maximum swap allocated to server in megabytes",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_max_disk_megabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_max_disk_megabytes",
		Help: "Maximum disk space allocated to server in megabytes",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_io = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_io",
		Help: "IO weight of server",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_max_cpu_absolute = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_max_cpu_absolute",
		Help: "Maximum cpu load allowed to server",
	}, []string{"server_name", "id", "node"})
	pterodactyl_server_most_recent_backup_time = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pterodactyl_server_most_recent_backup_time",
		Help: "Timestamp of the most recent backup",
	}, []string{"server_name", "id", "node"})
)

func recordMetrics() {
	go func() {
		for {
			servers, err := getServerList()
			servers_per_node := make(map[string]int)
			online_servers_per_node := make(map[string]int)
			if err == nil {
				for _, server := range servers {
					if !server.Attributes.IsSuspended {
						servers_per_node[server.Attributes.Node]++
						// Get more information
						resources, err := getServerResources(server.Attributes.Identifier)
						if err == nil {
							if resources.Attributes.CurrentState == "running" {
								online_servers_per_node[server.Attributes.Node]++
							}
							pterodacytl_server_memory_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(resources.Attributes.Resources.MemoryBytes) / 1000000)
							pterodactyl_server_cpu_absolute.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(resources.Attributes.Resources.CPUAbsolute)
							pterodactyl_server_disk_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(resources.Attributes.Resources.DiskBytes) / 1000000)
							pterodactyl_server_network_rx_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(resources.Attributes.Resources.NetworkRXBytes) / 1000000)
							pterodactyl_server_network_tx_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(resources.Attributes.Resources.NetworkTXBytes) / 1000000)
							pterodactyl_server_uptime_milliseconds.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(resources.Attributes.Resources.Uptime))
						}
						pterodactyl_server_max_memory_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(server.Attributes.Limits.Memory))
						pterodactyl_server_max_swap_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(server.Attributes.Limits.Swap))
						pterodactyl_server_max_disk_megabytes.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(server.Attributes.Limits.Disk))
						pterodactyl_server_max_cpu_absolute.WithLabelValues(server.Attributes.Name, server.Attributes.Identifier, server.Attributes.Node).Set(float64(server.Attributes.Limits.CPU))
					}
				}
				for node, servers := range servers_per_node {
					pterodactyl_total_servers.WithLabelValues(node).Set(float64(servers))
				}
				for node, servers := range online_servers_per_node {
					pterodactyl_running_servers.WithLabelValues(node).Set(float64(servers))
				}
			} else {
				panic(err)
			}
			fmt.Println("Finished requesting servers")

			time.Sleep(15 * time.Second)
		}
	}()
}

// Get the server list with the type above
func getServerList() ([]ServerListServer, error) {
	// Get the server list
	// https://dashflo.net/docs/api/pterodactyl/v1/#req_26cd9ef4a75540d6be8b4ef683e2b1a2
	page := 1
	var allServers []ServerListServer
	for {
		fmt.Println("Requesting servers from page " + strconv.Itoa(page))
		req, err := http.NewRequest("GET", panelHost+"/api/client?type=admin-all", nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", "Bearer "+apiKey)
		req.Header.Add("Accept", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		bodyBytes, _ := io.ReadAll(resp.Body)
		var servers ServerListResponse
		json.Unmarshal(bodyBytes, &servers)
		allServers = append(allServers, servers.Data...)
		if servers.Meta.Pagination.CurrentPage == servers.Meta.Pagination.TotalPages {
			break
		}
		fmt.Println("Loading next page")
		page++
	}
	fmt.Println("Got " + strconv.Itoa(len(allServers)) + " servers")
	return allServers, nil
}

func getServerResources(identifier string) (Stats, error) {
	req, err := http.NewRequest("GET", panelHost+"/api/client/servers/"+identifier+"/resources", nil)
	if err != nil {
		return Stats{}, err
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return Stats{}, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	var stats Stats
	json.Unmarshal(bodyBytes, &stats)
	return stats, nil
}

var client *http.Client
var panelHost string
var apiKey string

func main() {
	panelHost = os.Getenv("PANEL_URL")
	if panelHost == "" {
		panic("PANEL_URL not set")
	}
	apiKey = os.Getenv("CLIENT_API_TOKEN")
	if apiKey == "" {
		panic("CLIENT_API_TOKEN not set")
	}

	client = &http.Client{}
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9531", nil)
}
