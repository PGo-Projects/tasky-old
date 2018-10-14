package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/public"
	"github.com/PGo-Projects/tasky/internal/security"
	"github.com/PGo-Projects/tplmgr"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	security.MustSetupSecurity()

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	security.SetupSecurityRouter(mux)

	staticAssetsPath := http.Dir(viper.GetString(config.WebpackAssetsPathKey))
	mux.Method(http.MethodGet, "/public/*", http.StripPrefix("/public/", gzipped.FileServer(staticAssetsPath)))

	public.RegisterRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
