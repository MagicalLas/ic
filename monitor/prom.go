package monitor

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/examples/middleware/httpmiddleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gomod.usaken.org/ic/config"
	"gomod.usaken.org/ic/spine"
	"net/http"
	"time"
)

func RunPrometheusServer(c *config.Config) error {
	spine.SystemGroup.Add(1)
	defer spine.SystemGroup.Done()

	if !c.EnablePrometheusMetric {
		return nil
	}

	middleware := httpmiddleware.New(prometheus.DefaultRegisterer, nil)
	metricHandler := middleware.WrapHandler(
		"/metrics",
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{},
		),
	)

	server := http.Server{Addr: ":9000", Handler: metricHandler}

	go func() {
		spine.SystemGroup.Add(1)
		defer spine.SystemGroup.Done()

		reason := <-spine.C.Done()
		fmt.Printf("prom server shutdown started due to %s\n", reason)
		// 5분은 휴리스틱하게 정해진 시간이다.
		// prometheus 서버를 내리기전에 이미 충분하게 요청이 들어오지 않은 상태이겠지만,
		// 혹시 1분이상 실행중인 요청이 있다면 실패하도록한다.
		// timeout값보다 크게 하여 최대한 보수적으로 잡는다.
		ctx, _ := context.WithTimeout(context.Background(), time.Minute)
		err := server.Shutdown(ctx)
		if err != nil {
			fmt.Printf("prom server shutdown failed %e\n", err)
		}
		fmt.Printf("prom server successfully shutdown\n")
	}()

	go func() {
		spine.SystemGroup.Add(1)
		defer spine.SystemGroup.Done()

		server.ListenAndServe()
		fmt.Printf("prom server shutdown end\n")
	}()

	fmt.Printf("prom server running... \n")

	return nil
}

func CollectHTTPRequest(uri string) {

}

func CollectHTTPResponse(uri string, code int, d time.Duration) {

}
