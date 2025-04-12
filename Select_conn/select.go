
package Select_conn

func FastestResponse(ch1, ch2 <- chan string) string {
	select {
	case msg := <- ch1:
		return msg
	case msg := <- ch2:
		return msg
	}
}