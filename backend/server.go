package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ariefsn/gotik/apps/tiktok/repository"
	"github.com/ariefsn/gotik/apps/tiktok/services"
	"github.com/ariefsn/gotik/graph"
	"github.com/ariefsn/gotik/graph/resolvers"
	"github.com/ariefsn/gotik/middlewares"
	"github.com/ariefsn/gotik/queue"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"github.com/playwright-community/playwright-go"
)

const defaultPort = "8080"

func main() {
	err := playwright.Install()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	corsCfg := cors.Options{
		AllowedOrigins: []string{"http://localhost:5173", "http://localhost:" + port},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}
	router.Use(cors.Handler(corsCfg))
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/video" {
				w.Header().Set("Content-Type", "application/json")
			}
			next.ServeHTTP(w, r)
		})
	})

	router.Use(middlewares.Inject())

	router.Get("/video", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		videoUrlEncoded := strings.ReplaceAll(q.Get("url"), " ", "+")
		videoUrl, _ := base64.StdEncoding.DecodeString(videoUrlEncoded)
		urlParsed, _ := url.Parse(string(videoUrl))
		videoToken := strings.ReplaceAll(urlParsed.Query().Get("token"), " ", "+")

		client := &http.Client{}

		req, err := http.NewRequest("GET", urlParsed.String(), nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Cookie", "tt_chain_token="+videoToken)

		res, _ := client.Do(req)

		for k, v := range res.Header {
			w.Header().Add(k, strings.Join(v, " "))
		}
		io.Copy(w, res.Body)
	})

	natsUrl := os.Getenv("NATS_URL")
	if natsUrl == "" {
		natsUrl = nats.DefaultURL
	}

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	tiktokRepo := repository.NewTiktokRepository()
	tiktokSvc := services.NewTiktokService(tiktokRepo, nc)

	que := queue.NewQueue(queue.QueueOptions{
		Conn:          nc,
		TiktokService: tiktokSvc,
	})

	que.Listen()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		TiktokService: tiktokSvc,
	}}))

	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)

		fmt.Println(fmt.Sprintf("[%s]", strings.ToUpper(string(oc.Operation.Operation))), map[string]interface{}{
			"operationName": oc.OperationName,
			"variables":     oc.Variables,
		})

		return next(ctx)
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	// aktifkan ws
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				for _, v := range corsCfg.AllowedOrigins {
					if v == origin {
						return true
					}
				}
				return false
			},
		},
	})

	srv.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
