package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var launchRocketCmd = &cobra.Command{
	Use:   "rocket <name>",
	Short: "Launches a rocket",
	Long: `Launches a rocket into space for its mission. For example:

	rocketctl launch rocket r1 --countdown=10
	`,

	Run: func(cmd *cobra.Command, args []string) {
		countdown, err := cmd.Flags().GetInt("countdown")
		if err != nil {
			fmt.Println(err)
		}

		rocket := getRocketByName(args[0])

		var b strings.Builder
		b.WriteString(fmt.Sprintf("launching rocket %s: ", rocket.Name))
		for i := countdown; i >= 0; i-- {
			b.WriteString(fmt.Sprintf("%d...", i))
		}
		b.WriteString(" -- blast off!!")
		fmt.Println(b.String())
	},
}

func init() {
	launchRocketCmd.PersistentFlags().IntP("countdown", "c", 10, "countdown length for launch")
	//nolint:errcheck
	launchRocketCmd.MarkPersistentFlagRequired("countdown")
	launchCmd.AddCommand(launchRocketCmd)
}
