package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/xUnholy/go-proxy/pkg/execute"
	"github.com/xUnholy/go-proxy/pkg/prompt"

	"github.com/xUnholy/go-proxy/internal/cntlm"
	config "github.com/xUnholy/go-proxy/internal/tools"
)

var (
	cntlmFile = "/usr/local/etc/cntlm.conf"
	port      int
)

func SetupSetCli() *cobra.Command {
	// Top-level command
	var setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set CNTLM Proxy Config",
	}

	var setNpmCmd = &cobra.Command{
		Use:   "npm",
		Short: "This command will set the NPM proxy values. Both https-proxy and proxy will be set",
		Run:   setNpmCmd,
	}
	setNpmCmd.Flags().IntVar(&port, "port", 3128, "set custom CNTLM `PORT`")

	var setGitCmd = &cobra.Command{
		Use:   "git",
		Short: "This command will set the GIT global proxy values. Both http.proxy and https.proxy will be set",
		Run:   setGitCmd,
	}
	setGitCmd.Flags().IntVar(&port, "port", 3128, "set custom CNTLM `PORT`")

	var setUsernameCmd = &cobra.Command{
		Use:   "username",
		Short: "This command will update the Username value in your CNTLM.conf file",
		Run:   setUsernameCmd,
	}

	var setPasswordCmd = &cobra.Command{
		Use:   "password",
		Short: "This command will update the Password value in your CNTLM.conf file",
		Run:   setPasswordCmd,
	}

	var setDomainCmd = &cobra.Command{
		Use:   "domain",
		Short: "This command will update the domain value in your CNTLM.conf file",
		Run:   setDomainCmd,
	}

	// add subcommands to set
	setCmd.AddCommand(setNpmCmd, setGitCmd, setUsernameCmd, setPasswordCmd, setDomainCmd)
	return setCmd
}

func setNpmCmd(cmd *cobra.Command, args []string) {
	p := makeProxyURL(port)
	err := config.EnableNPMProxyConfiguration(p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Set npm config successfully")
}

func setGitCmd(cmd *cobra.Command, args []string) {
	p := makeProxyURL(port)
	err := config.EnableGITProxyConfiguration(p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Set git config successfully")
}

func setUsernameCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("Enter Username: ")
	output, err := prompt.GetInput(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	update := fmt.Sprintln("Username\t", output)
	err = cntlm.UpdateFile(cntlmFile, update)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Set CNTLM username successfully")
}

func setPasswordCmd(cmd *cobra.Command, args []string) {
	// TODO: changing the password should restart cntlm to re-auth [go-proxy/#47]
	fmt.Printf("Enter Password: ")
	e := execute.Command{Cmd: "cntlm", Args: []string{"-H"}, Stdin: os.Stdin}
	out, err := execute.RunCommand(e)
	if err != nil {
		log.Fatalln(err)
	}
	err = cntlm.UpdateFile(cntlmFile, string(out))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Set CNTLM password successfully")
}

func setDomainCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("Enter Proxy Domain: ")
	output, err := prompt.GetInput(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}
	update := fmt.Sprintln("Domain\t", output)
	err = cntlm.UpdateFile(cntlmFile, update)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Set CNTLM domain successfully")
}
