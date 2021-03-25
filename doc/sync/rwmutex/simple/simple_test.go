package simple

import (
	"log"
	"testing"
	"time"
)

func Test_Process(t *testing.T) {
	p := &process{}
	p.w(5, func() {})

	go func() {
		ticker := time.NewTimer(time.Second * 11)
		for {
			select {
			case <-ticker.C:
				return
			default:
				val := p.r()
				log.Printf("子routine读数： %d", val)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 2)
		p.w(10, func() {
			log.Println("func1 写10")
			time.Sleep(time.Second * 3)
			log.Println("func1 写完毕")
		})
	}()
	go func() {
		time.Sleep(time.Second * 2)
		p.w(7, func() {
			log.Println("func2 写7")
			time.Sleep(time.Second * 3)
			log.Println("func2 写完毕")
		})
	}()

	ticker := time.NewTimer(time.Second * 11)
	for {
		select {
		case <-ticker.C:
			return
		default:
			val := p.r()
			log.Printf("读数： %d", val)
			time.Sleep(time.Second)
		}
	}
}
