package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

func main() {

	// export OS_CLOUD=dk-stage
	// Use cloud from environment (OS_CLOUD) or set it directly
	opts := &clientconfig.ClientOpts{
		Cloud: os.Getenv("OS_CLOUD"),
	}

	// Authenticate
	provider, err := clientconfig.AuthenticatedClient(opts)
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}
	// Authentication successful

	// Create a Compute client (Nova)
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		log.Fatalf("Failed to create compute client: %v", err)
	}

	// List flavors
	allPages, err := flavors.ListDetail(client, nil).AllPages()
	if err != nil {
		log.Fatalf("Failed to list flavors: %v", err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		log.Fatalf("Failed to extract flavors: %v", err)
	}

	// Print flavor info
	// Print header
	fmt.Printf("%-20s %-36s %-6s %-6s %-6s\n", "NAME", "ID", "RAM", "VCPUs", "DISK")
	fmt.Println("--------------------------------------------------------------------------------")

	// Print each flavor row
	for _, f := range allFlavors {
		fmt.Printf("%-20s %-36s %-6d %-6d %-6d\n", f.Name, f.ID, f.RAM, f.VCPUs, f.Disk)
	}

}
