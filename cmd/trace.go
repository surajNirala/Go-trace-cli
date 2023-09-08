package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace IP",
	Long:  "Trace IP Address Application",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
				// var url = "https://ipinfo.io/" + ip + "/geo"
				// fmt.Println(url)
			}
		} else {

			fmt.Println("Please provide the IP.")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	Ip       string `json::"ip"`
	Hostname string `json::"hostname"`
	Anycast  string `json::"anycast"`
	City     string `json::"city"`
	Region   string `json::"region"`
	Country  string `json::"country"`
	Loc      string `json::"loc"`
	Org      string `json::"org"`
	Postal   string `json::"postal"`
	Timezone string `json::"timezone"`
	Readme   string `json::"readme"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)
	data := Ip{}
	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}
	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND")
	fmt.Println(data)
	fmt.Printf("IP : %s\nCITY : %s\nREGION : %s\nCOUNTRY : %s\n", data.Ip, data.City, data.Region, data.Country)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to  get the response.")
	}

	responseByte, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("Unable to read the response.")
	}

	return responseByte

}
