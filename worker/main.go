package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/temporalio/samples-go/helloworld"
	"go.temporal.io/sdk/client"
	sdklog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
)

var defaultHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelDebug,
})

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		Namespace: "default",
		HostPort:  os.Getenv("TEMPORAL_CLIENT_HOSTPORT"),
		Logger:    sdklog.NewStructuredLogger(slog.New(defaultHandler)),
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})

	w.RegisterWorkflow(helloworld.Workflow)
	w.RegisterActivity(helloworld.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
