package cmd

import (
	"net/http"
	"sync"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/ory/graceful"

	//"github.com/ory/herodot"
	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
	"github.com/songvi/kratos-selfservice-ui-go/driver"
	"github.com/songvi/kratos-selfservice-ui-go/driver/configuration"
	//"github.com/urfave/negroni"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start http service",
	Long:  "Start http service",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		wg.Add(1)
		go start(&wg, cmd, args)
		wg.Wait()
	},
}

func start(wg *sync.WaitGroup, cmd *cobra.Command, args []string) {
	defer wg.Done()
	l := logrus.New()
	r := httprouter.New()

	c := configuration.NewViperProvider()

	reg := driver.NewDefaultRegistry()

	reg.WithConfig(c)
	reg.WithLogger(l)

	reg.LoginHandler().RegisterRouter(r)
	reg.ProfileHandler().RegisterRouter(r)
	reg.RegistrationHandler().RegisterRouter(r)
	reg.ErrorHandler().RegisterRouter(r)
	reg.LogoutHandler().RegisterRouter(r)
	reg.WhoamiHandler().RegisterRouter(r)

	server := graceful.WithDefaults(&http.Server{
		Addr:    c.ListenOn(),
		Handler: context.ClearHandler(r),
	})

	l.Printf("Strating the http server on: %s", c.ListenOn())
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		l.Errorln(err)
		l.Fatalln("Failed to gracefully shutdown http server")
	}

	l.Println("Admin httpd was shutdown gracefully")
}
