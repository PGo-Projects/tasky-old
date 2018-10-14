package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tplmgr"
	"github.com/go-webpack/webpack"
	"github.com/lpar/gzipped"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func setup() {
	webpack.Plugin = "manifest"
	webpack.FsPath = viper.GetString(config.WebpackFsPathKey)
	webpack.WebPath = viper.GetString(config.WebpackWebPathKey)
	webpack.Init(true)

	funcMap := template.FuncMap{
		"asset": webpack.AssetHelper,
	}

	templateLayoutPath := viper.GetString(config.TemplatesLayoutPathKey)
	templatePath := viper.GetString(config.TemplatesPathKey)
	tplmgr.SetConfig(templateLayoutPath, templatePath)
	tplmgr.SetDelimiters("{%", "%}")
	tplmgr.MustLoadWithFuncs(funcMap)
}

func Run(cmd *cobra.Command, args []string) {
	setup()
	initStaticServer()
	http.HandleFunc("/", IndexPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	tplmgr.Render(w, "index.tmpl", nil)
}

func initStaticServer() {
	var staticAssetsPath = http.Dir(viper.GetString(config.WebpackAssetsPathKey))
	http.Handle("/public/", http.StripPrefix("/public/", gzipped.FileServer(staticAssetsPath)))
}
