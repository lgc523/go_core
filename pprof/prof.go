package main

import (
	"log"
	"math/rand"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
	"go_core/http"
)

const row, col = 1123, 1123

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	c, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	//get cpu info
	if err := pprof.StartCPUProfile(c); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	//mem
	m, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	//runtime.GC()
	if err := pprof.WriteHeapProfile(m); err != nil {
		log.Fatal("could nnot write memeory profile: ", err)
	}
	m.Close()

	//goroutine

	g, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroute profile: ", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil {
		log.Fatal("could not write goroutine profile: ")
	} else {
		gProf.WriteTo(g, 0)
	}
	g.Close()

}
