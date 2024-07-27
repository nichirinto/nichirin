package logger

import "go.uber.org/zap"

type Logger = zap.Logger

var Log, _ = zap.NewProduction()
