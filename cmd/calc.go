package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/martijnkorbee/michael/cmd/internal/pkg/mortgage"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	sum      int
	duration int
	interest float64
)

var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate mortgage costs over time",
	Long: `Calculate the mortgage costs over time for the total duration of the mortgage. 
Outputs a sqlite database, populated with a rowset per month/year with the compounding costs of the mortgage`,
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"linear"},
}

func init() {
	calcCmd.AddCommand(calcLinearCmd)
	calcCmd.AddCommand(calcAnnuitiesCmd)

	calcLinearCmd.Flags().IntVar(&sum, "sum", 0, "mortgage sum")
	calcLinearCmd.Flags().IntVar(&duration, "duration", 0, "mortgage duration in years")
	calcLinearCmd.Flags().Float64Var(&interest, "interest", 0.0, "interest as decimal")

	calcLinearCmd.MarkFlagRequired("sum")
	calcLinearCmd.MarkFlagRequired("duration")
	calcLinearCmd.MarkFlagRequired("interest")

	calcAnnuitiesCmd.Flags().IntVar(&sum, "sum", 0, "mortgage sum")
	calcAnnuitiesCmd.Flags().IntVar(&duration, "duration", 0, "mortgage duration in years")
	calcAnnuitiesCmd.Flags().Float64Var(&interest, "interest", 0.0, "interest as decimal")

	calcAnnuitiesCmd.MarkFlagRequired("sum")
	calcAnnuitiesCmd.MarkFlagRequired("duration")
	calcAnnuitiesCmd.MarkFlagRequired("interest")
}

var calcLinearCmd = &cobra.Command{
	Use:   "linear",
	Short: "Calculate a linear mortgage",
	Long:  "Calculates a mortgage based on a linear model.",
	Run: func(cmd *cobra.Command, args []string) {
		mustConnectToDB()
		mustPrepareDB("linear")

		defer app.database.Close()

		// create new mortgage
		mg := mortgage.New("linear", sum, duration, interest)

		collection := app.database.Collection("linear_mortgage")

		log.Println("calculating mortgage")
		// calculate and inserts mortgage rows
		for i := 1; i <= mg.Terms; i++ {
			// calculate next month
			if err := mg.CalculateNextMonth(); err != nil {
				log.Fatalln(err.Error())
			}

			_, err := collection.Insert(mg)
			if err != nil {
				log.Println(err.Error())
			}

		}

		log.Println("finished, printing mortgage by yearly increment")

		osCmd := exec.Command("sqlite3", "-header", "-column", fmt.Sprintf("%s/mortgage.db", app.rootpath), "select * from linear_mortgage where month = 12 or (month = 1 and year = 1);")

		if out, err := osCmd.CombinedOutput(); err != nil {
			log.Println("failed to print", err.Error())
		} else {
			log.Printf("\n%v\n", string(out))
		}
	},
}

var calcAnnuitiesCmd = &cobra.Command{
	Use:   "annuities",
	Short: "Calculate a annuities mortgage",
	Long:  "Calculates a mortgage based on a annuities model.",
	Run: func(cmd *cobra.Command, args []string) {
		mustConnectToDB()
		mustPrepareDB("annuities")

		// create new mortgage
		mg := mortgage.New("annuities", sum, duration, interest)

		collection := app.database.Collection("annuities_mortgage")

		log.Println("calculating mortgage")
		// calculate and inserts mortgage rows
		for i := 1; i <= mg.Terms; i++ {
			// calculate next month
			if err := mg.CalculateNextMonth(); err != nil {
				log.Fatalln(err.Error())
			}

			_, err := collection.Insert(mg)
			if err != nil {
				log.Println(err.Error())
			}

		}

		log.Println("finished, printing mortgage by yearly increment")

		osCmd := exec.Command("sqlite3", "-header", "-column", fmt.Sprintf("%s/mortgage.db", app.rootpath), "select * from annuities_mortgage where month = 12 or (month = 1 and year = 1);")

		if out, err := osCmd.CombinedOutput(); err != nil {
			log.Println("failed to print", err.Error())
		} else {
			log.Printf("\n%v\n", string(out))
		}
	},
}
