package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/newrelic/nr-integrations-builder/scaffold"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	defaultProtocolVersion = 1
	defaultOS              = "linux"
	defaultInterval        = 15
	descriptionTmpl        = "Reports status and metrics for %s service"
)

var iname, idestinationPath, icompanyName, icompanyPrefix, entityType string

var initCmd = &cobra.Command{
	Use:   "init [integration name]",
	Short: "Initialize an integration",
	Long:  `Initialize an integration generating a scaffold.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("An integration name is mandatory")
		}
		return checkFlags(cmd.Flags())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		iname = args[0]
		s := scaffold.Scaffold{
			DestinationPath: filepath.Join(idestinationPath, iname),
			Date:            time.Now().Format("2006-01-02"),
			Integration: scaffold.Integration{
				Name:            iname,
				Description:     fmt.Sprintf(descriptionTmpl, iname),
				ProtocolVersion: defaultProtocolVersion,
				OS:              defaultOS,
				Prefix:          "config/" + icompanyPrefix + "-" + iname,
				Interval:        defaultInterval,
				CompanyPrefix:   icompanyPrefix,
				CompanyName:     icompanyName,
				BinaryName:      icompanyPrefix + "-" + iname,
				EventType:       strings.Title(strings.ToLower(icompanyPrefix)) + strings.Title(strings.ToLower(iname)) + "Sample",
				EntityType:      entityType,
			},
		}
		err := s.Generate(verbose)
		if err != nil {
			return err
		}
		return s.InitVendoring()
	},
}

func checkFlags(flags *pflag.FlagSet) error {

	eType, err := flags.GetString("entity-type")
	if err != nil {
		return err
	}
	if eType != "" && eType != "local" && eType != "remote" {
		return fmt.Errorf("Wrong entity-type argument, valid values are: [remote, local]")
	}

	return checkRequiredFlags(flags)
}

func checkRequiredFlags(flags *pflag.FlagSet) error {
	missingFlagNames := []string{}

	flags.VisitAll(func(flag *pflag.Flag) {
		requiredAnnotation := flag.Annotations[cobra.BashCompOneRequiredFlag]
		if len(requiredAnnotation) == 0 {
			return
		}

		flagRequired := requiredAnnotation[0] == "true"

		if flagRequired && !flag.Changed {
			missingFlagNames = append(missingFlagNames, flag.Name)
		}
	})

	if len(missingFlagNames) > 0 {
		return fmt.Errorf("Required flag/flags: \"%s\" has/have not been set", strings.Join(missingFlagNames, "\", \""))
	}

	return nil
}

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVarP(&icompanyPrefix, "company-prefix", "c", "", "Company prefix identifier (required)")
	err := initCmd.MarkPersistentFlagRequired("company-prefix")
	if err != nil {
		panic(err)
	}
	initCmd.PersistentFlags().StringVarP(&icompanyName, "company-name", "n", "", "Company name (required)")
	err = initCmd.MarkPersistentFlagRequired("company-name")
	if err != nil {
		panic(err)
	}
	initCmd.PersistentFlags().StringVarP(&idestinationPath, "destination-path", "p", "./", "Destination path for initialized integration")
	initCmd.PersistentFlags().StringVarP(&entityType, "entity-type", "e", "remote", "Type of entity to generate: [remote,local] (required)")
}
