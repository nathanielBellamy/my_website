#!/bin/bash
awk '/func \(lc \*LogsController\) StreamLogsHandler/{in_handler=1} /func \(lc \*LogsController\) streamNewLines/{in_handler=0} in_handler {gsub(/log\.Ctx\(ctx\)/, "log.Ctx(r.Context())")} {print}' backend/go/logs/sse.go > tmp.go && mv tmp.go backend/go/logs/sse.go
