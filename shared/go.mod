module uptime-monitor/shared

go 1.20

require github.com/lib/pq v1.10.9

replace uptime-monitor/api => ../api

replace uptime-monitor/shared => ../shared

replace uptime-monitor/worker => ../worker
