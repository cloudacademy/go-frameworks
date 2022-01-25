package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createRocketCmd = &cobra.Command{
	Use:   "rocket <name>",
	Short: "Creates a new rocket",
	Long: `Creates a new rocket. For example:

	rocketctl create rocket r1 --type=saturnv --mission=apollo11 --fuel=5000 --maxspeed=25000
	`,

	Run: func(cmd *cobra.Command, args []string) {
		rocketType, err := cmd.Flags().GetString("type")
		if err != nil {
			fmt.Println(err)
		}

		mission, err := cmd.Flags().GetString("mission")
		if err != nil {
			fmt.Println(err)
		}

		fuel, err := cmd.Flags().GetInt("fuel")
		if err != nil {
			fmt.Println(err)
		}

		maxspeed, err := cmd.Flags().GetInt("maxspeed")
		if err != nil {
			fmt.Println(err)
		}

		rocket := rocket{
			Name:     args[0],
			Type:     rocketType,
			Mission:  mission,
			Fuel:     fuel,
			Maxspeed: maxspeed,
		}

		createRocket(rocket)
	},
}

func init() {
	createRocketCmd.PersistentFlags().StringP("type", "t", "", "rocket type, possible values: saturnv, falcon9, electron")
	createRocketCmd.PersistentFlags().StringP("mission", "m", "", "name of the mission")
	createRocketCmd.PersistentFlags().IntP("fuel", "f", 100, "amount of fuel required")
	createRocketCmd.PersistentFlags().IntP("maxspeed", "s", 1000, "maximum speed")
	createCmd.AddCommand(createRocketCmd)
}
