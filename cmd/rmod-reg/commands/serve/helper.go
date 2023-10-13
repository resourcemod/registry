// Package serve is an entry point and help tools for serve commands
package serve

func getExample() string {
	return `
	rmod-rmod-reg serve --port=8888 --addr=0.0.0.0

	rmod-rmod-reg serve

	rmod-rmod-reg serve --port=80 --addr=0.0.0.0 --ui=false
	`
}
