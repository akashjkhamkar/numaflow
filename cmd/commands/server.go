/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"strings"

	"github.com/spf13/cobra"

	sharedutil "github.com/numaproj/numaflow/pkg/shared/util"
	svrcmd "github.com/numaproj/numaflow/server/cmd/server"
)

func NewServerCommand() *cobra.Command {
	var (
		insecure         bool
		port             int
		namespaced       bool
		managedNamespace string
		baseHref         string
		disableAuth      bool
		dexServerAddr    string
		serverAddr       string
	)

	command := &cobra.Command{
		Use:   "server",
		Short: "Start a Numaflow server",
		Run: func(cmd *cobra.Command, args []string) {
			if !cmd.Flags().Changed("port") && insecure {
				port = 8080
			}
			if !strings.HasSuffix(baseHref, "/") {
				baseHref = baseHref + "/"
			}
			opts := svrcmd.ServerOptions{
				Insecure:         insecure,
				Port:             port,
				Namespaced:       namespaced,
				ManagedNamespace: managedNamespace,
				BaseHref:         baseHref,
				DisableAuth:      disableAuth,
				DexServerAddr:    dexServerAddr,
				ServerAddr:       serverAddr,
			}
			server := svrcmd.NewServer(opts)
			server.Start()
		},
	}
	command.Flags().BoolVar(&insecure, "insecure", sharedutil.LookupEnvBoolOr("NUMAFLOW_SERVER_INSECURE", false), "Whether to disable TLS, defaults to false.")
	command.Flags().IntVarP(&port, "port", "p", sharedutil.LookupEnvIntOr("NUMAFLOW_SERVER_PORT_NUMBER", 8443), "Port to listen on, defaults to 8443 or 8080 if insecure is set")
	command.Flags().BoolVar(&namespaced, "namespaced", sharedutil.LookupEnvBoolOr("NUMAFLOW_SERVER_NAMESPACED", false), "Whether to run in namespaced scope, defaults to false.")
	command.Flags().StringVar(&managedNamespace, "managed-namespace", sharedutil.LookupEnvStringOr("NUMAFLOW_SERVER_MANAGED_NAMESPACE", sharedutil.LookupEnvStringOr("NAMESPACE", "numaflow-system")), "The namespace that the server watches when \"--namespaced\" is \"true\".")
	command.Flags().StringVar(&baseHref, "base-href", sharedutil.LookupEnvStringOr("NUMAFLOW_SERVER_BASE_HREF", "/"), "Base href for Numaflow server, defaults to '/'.")
	command.Flags().BoolVar(&disableAuth, "disable-auth", sharedutil.LookupEnvBoolOr("NUMAFLOW_SERVER_DISABLE_AUTH", false), "Whether to disable authentication and authorization, defaults to false.")
	command.Flags().StringVar(&dexServerAddr, "dex-server-addr", sharedutil.LookupEnvStringOr("NUMAFLOW_SERVER_DEX_SERVER_ADDR", "http://numaflow-dex-server:5556/dex"), "The address of the Dex server.")
	command.Flags().StringVar(&serverAddr, "server-addr", sharedutil.LookupEnvStringOr("NUMAFLOW_SERVER_ADDRESS", "https://localhost:8443"), "The external address of the Numaflow server.")
	return command
}
