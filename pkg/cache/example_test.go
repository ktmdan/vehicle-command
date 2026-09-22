package cache_test

import (
	"context"
	"fmt"

	"github.com/teslamotors/vehicle-command/pkg/cache"
	"github.com/teslamotors/vehicle-command/pkg/connector/inet"
	"github.com/teslamotors/vehicle-command/pkg/protocol"
	"github.com/teslamotors/vehicle-command/pkg/vehicle"
)

func Example() {
	const cacheFilename = "my_cache.json"
	const privateKeyFilename = "private_key.pem"

	// A BLE connection (ble.NewConnection(ctx, "myvin123")) works the same way;
	// inet is used here so this example has no third-party dependencies.
	conn := inet.NewConnection("myvin123", "", "https://my-proxy.example.com", "my-agent")
	defer conn.Close()

	// Try to load cache from disk if it doesn't already exist
	myCache, err := cache.ImportFromFile(cacheFilename)
	if err != nil {
		myCache = cache.New(5) // Create a cache that holds sessions for up to five vehicles
	}

	privateKey, err := protocol.LoadPrivateKey(privateKeyFilename)
	if err != nil {
		panic(err)
	}

	car, err := vehicle.NewVehicle(conn, privateKey, myCache)
	if err != nil {
		panic(err)
	}

	if err := car.Connect(context.Background()); err != nil {
		panic(err)
	}
	defer car.Disconnect()

	// StartSession(...) will load from myCache when possible.
	if err := car.StartSession(context.Background(), nil); err != nil {
		panic(err)
	}

	defer func() {
		if err := car.UpdateCachedSessions(myCache); err != nil {
			fmt.Printf("Error updating session cache: %s\n", err)
			return
		}
		if err := myCache.ExportToFile(cacheFilename); err != nil {
			fmt.Printf("Error saving session cache: %s\n", err)
		}
	}()

	// Interact with car
}
