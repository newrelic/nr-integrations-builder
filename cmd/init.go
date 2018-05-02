package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/newrelic/nr-integrations-builder.v1/scaffold"
)

const (
	defaultProtocolVersion = 1
	defaultOS              = "linux"
	defaultInterval        = 15
	descriptionTmpl        = "Reports status and metrics for %s service"
)

var iname, idestinationPath, icompanyName, icompanyPrefix string

var initCmd = &cobra.Command{
	Use:   "init [integration name]",
	Short: "Initialize an integration",
	Long:  `Initialize an integration generating a scaffold.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("An integration name is mandatory")
		}
		iname = args[0]
		s := scaffold.Scaffold{
			DestinationPath: filepath.Join(idestinationPath, iname),
			Date:            time.Now().Format("2006-01-02"),
			Integration: scaffold.Integration{
				Name:            iname,
				Description:     fmt.Sprintf(descriptionTmpl, iname),
				ProtocolVersion: defaultProtocolVersion,
				OS:              defaultOS,
				Prefix:          "config/" + iname,
				Interval:        defaultInterval,
				CompanyPrefix:   icompanyPrefix,
				CompanyName:     icompanyName,
				BinaryName:      icompanyPrefix + "-" + iname,
				EventType:       strings.Title(strings.ToLower(iname)) + "Sample",
			},
		}
		err := s.Generate(verbose)
		if err != nil {
			return err
		}
		return s.InitVendoring()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVarP(&icompanyPrefix, "company-prefix", "c", "custom", "Company prefix identifier")
	initCmd.PersistentFlags().StringVarP(&icompanyName, "company-name", "n", "custom", "Company name")
	initCmd.PersistentFlags().StringVarP(&idestinationPath, "destination-path", "p", "./", "Destination path for initialized integration")
}
