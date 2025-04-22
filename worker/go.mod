module uptime-monitor/worker

go 1.20

require uptime-monitor/shared v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.10.9 // indirect

replace uptime-monitor/api => ../api

replace uptime-monitor/shared => ../shared

replace uptime-monitor/worker => ../worker
