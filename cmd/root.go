package cmd

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/dnephin/cobra"
	"github.com/iwalz/tdoc/outputs"
	"github.com/iwalz/tdoc/table"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	SvgDir    string
	extension string
	verbose   bool
	debug     bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tdoc file.tdoc",
	Short: "Short description",
	Long:  `Long description`,
	Args:  CheckFile,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			log.SetLevel(log.InfoLevel)
		}
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		fs := afero.NewOsFs()
		file := outputs.NewFile(fs, extension, SvgDir)
		executor := outputs.NewExecutor(fs, extension)
		executor.Exec(file, args)

		/*filename := args[0]
		newFilename := strings.Replace(filename, "."+extension, ".svg", 1)

		content, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}
		cl := elements.NewComponentsList(SvgDir)
		cl.Parse()
		p := &parser.TdocParserImpl{}
		l := parser.NewLexer(string(content), cl)
		p.Parse(l)
		ast := p.AST()

		m := renderer.NewMiddleware(ast, cl)
		file, err := os.Create(newFilename)
		svg := svg.New(file)
		table := m.Scan(ast, cl)
		svg.Start(table.Width(), table.Height())
		m.Render(svg)
		svg.End()
		*/
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.tdoc.yaml)")
	RootCmd.PersistentFlags().StringVarP(&extension, "extension", "e", "tdoc", "file extension to use (default 'tdoc')")
	RootCmd.PersistentFlags().StringVarP(&SvgDir, "svgdir", "s", "/home/ingo/svg", "Source directory for components. foo.svg will make component foo available")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enables verbose mode")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enables debug mode")
	RootCmd.PersistentFlags().BoolVarP(&table.Wireframe, "wireframe", "w", false, "Render wireframes")
	RootCmd.PersistentFlags().IntVar(&table.Dimension, "dimension", 120, "Set the width and height dimension per cell")
	RootCmd.PersistentFlags().IntVar(&table.Border, "border", 40, "Set border dimension per cell")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".tdoc") // name of config file (without extension)
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}
