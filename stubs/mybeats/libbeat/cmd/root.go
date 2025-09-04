package cmd

import (
	"github.com/spf13/cobra"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"
)

func GenRootCmdWithSettings(creator beat.Creator, settings instance.Settings) *cobra.Command {
	return &cobra.Command{}
}