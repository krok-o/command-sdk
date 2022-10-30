package pkg

import "github.com/spf13/cobra"

const (
	artifactURL = "artifact-url"
	platform    = "platform"
	payload     = "payload"
	eventType   = "event-type"
)

// Options define a set of default values to bind to.
type Options struct {
	ArtifactURL string
	Platform    string
	Payload     string
	EventType   string
}

// AddRequiredFlagsToCommand adds all required default flags to a command implementation.
// This is done currently only for cobra.Command(s).
func AddRequiredFlagsToCommand(cmd *cobra.Command, opts *Options) {
	cmd.Flags().StringVar(&opts.Platform, platform, "", "Platform defines where the hook came from, github, gitlab, etc.")
	cmd.Flags().StringVar(&opts.ArtifactURL, artifactURL, "", "Artifact URL provides access to a reconciled source.")
	cmd.Flags().StringVar(&opts.EventType, eventType, "", "Event type defines what the event was; push, pull, issue-comment etc.")
	cmd.Flags().StringVar(&opts.Payload, payload, "", "Payload contains the payload that triggered the event.")
}
