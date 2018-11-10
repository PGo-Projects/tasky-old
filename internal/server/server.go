package server

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PGo-Projects/output"
	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/public"
	"github.com/PGo-Projects/tasky/internal/security"
	"github.com/PGo-Projects/tasky/internal/taskyapi"
	"github.com/PGo-Projects/tplmgr"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-webpack/webpack"
	"github.com/lpar/gzipped"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const SOCK = "/tmp/tasky.sock"

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
	security.MustSetup()
	taskyapi.Setup()

	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	security.SetupSecurityRouter(mux)

	staticAssetsPath := http.Dir(viper.GetString(config.WebpackAssetsPathKey))
	mux.Method(http.MethodGet, "/public/*", http.StripPrefix("/public/", gzipped.FileServer(staticAssetsPath)))

	public.RegisterRoutes(mux)
	taskyapi.RegisterRoutes(mux)

	if config.DevRun {
		output.Println("Attempting to run on localhost...", output.BLUE)
		log.Fatal(http.ListenAndServe(":8080", mux))
	} else {
		os.Remove(SOCK)
		unixListener, err := net.Listen("unix", SOCK)
		if err != nil {
			output.Error(err)
		}
		os.Chmod(SOCK, 0666)
		defer unixListener.Close()

		output.Println("Attempting to run on unix socket...", output.BLUE)
		log.Fatal(http.Serve(unixListener, mux))
	}
}
