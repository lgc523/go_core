package main

import (
	"testing"
	"time"

	"go.uber.org/zap"
)
var url = "https://liguangchang.cn"

func TestZap(t *testing.T) {
	log, _ := zap.NewProduction()
	defer log.Sync()
	sl := log.Sugar()
	sl.Infow("failed to fetch URL", "url", url, "attemp", 3, "backoff", time.Second)
	sl.Infof("Failed to fetch URl:%s", url)
}

func TestZapSafety(t *testing.T){
	//critical performance
	lg, _ := zap.NewProduction()
	defer lg.Sync()
	lg.Info("failed to fetch URL",
		//structed context as stronly type field values
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
