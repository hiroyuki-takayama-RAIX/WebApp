module github.com/hiroyuki-takayama-RAIX/WebApp

go 1.20

require (
	github.com/hiroyuki-takayama-RAIX/config v0.0.0-00010101000000-000000000000
	golang.org/x/sync v0.3.0
)

require github.com/caarlos0/env/v6 v6.10.1 // indirect

replace github.com/hiroyuki-takayama-RAIX/config => ../config
