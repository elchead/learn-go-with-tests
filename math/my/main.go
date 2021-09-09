package main

import (
	"os"
	"time"
)

func main() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
}

// const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
// <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
// <svg xmlns="http://www.w3.org/2000/svg"
//      width="100%"
//      height="100%"
//      viewBox="0 0 300 300"
//      version="2.0">`

// const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

// const svgEnd = `</svg>`
