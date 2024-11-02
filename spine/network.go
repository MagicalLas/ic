package spine

import (
	"fmt"
	"net/http"
	"time"
)

// CheckNetworkInit 만약 네트워크가 연결되지 않은 상태에서 서버가
// 먼저 실행되는 경우에 서버가 바로 종료되지 않도록 한다.
// 이러한 상황은 istio를 사용하는 k8s에서 종종 일어나는데, sidecar가
// 준비되지 않은 상태에서 먼저 container가 뜨게된다면 network가 되지
// 않은 상태에서 서버 init에 실패할 수 있다.
func CheckNetworkInit() {
	fmt.Printf("network init check started\n")

	var retryDelay = 500 * time.Millisecond
	const maxDelay = 16 * time.Second
	const url = "https://magical.dev"

	for {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			fmt.Printf("network init done\n")
			resp.Body.Close()
			return
		}

		fmt.Printf("network init failed, retrying in %v...", retryDelay)
		time.Sleep(retryDelay)

		// exponential backoff
		if retryDelay < maxDelay {
			retryDelay *= 2
		}
	}
}
