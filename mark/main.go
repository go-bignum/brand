// Command mark writes the go-bignum mark, svg/color/go-bignum.svg:
//
//	cd mark && go run . > ../svg/color/go-bignum.svg
//
// go-bignum is arbitrary precision, and the mark is a very long number: the
// first digits, then an ellipsis because the number does not stop. The glyph
// is a "1", a "0" and three dots, set on one baseline.
//
// Nothing is placed by hand. The items below list what is drawn, left to
// right, each with its width; the layout puts the same gap between two
// neighbours and centres the whole row on the square. Change the digits and the
// mark redraws itself.
//
// The rest is the family grammar every go-* mark shares: a 256-unit rounded
// square (rx 56) filled with a diagonal two-stop gradient, a white line glyph
// at stroke-width 13 with round caps and joins. logos/brandgen renders the
// format matrix from the file this writes; the kit workflow runs this again
// and fails if a byte of svg/color differs.
package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	light = "#E27AB8" // the gradient's light stop
	dark  = "#A02A74" // its dark stop

	stroke = 13.0 // the family's line width

	height = 88.0 // height of a digit, between stroke centres
	gap    = 26.0 // clear space between two neighbours, between stroke centres
	pitch  = 22.0 // distance between two dots of the ellipsis
	dots   = 3    // an ellipsis has three
)

// item is one thing on the baseline: its width between stroke centres and a
// function that draws it from its left edge x, top y and bottom y.
type item struct {
	width float64
	draw  func(x, top, bottom float64) []string
}

// digitOne is a stem with a flag that leans back from its top.
func digitOne(x, top, bottom float64) []string {
	const w = 22.0
	return []string{fmt.Sprintf(`<path d="M%s %sL%s %sV%s"/>`,
		num(x), num(top+0.25*height), num(x+w), num(top), num(bottom))}
}

// digitZero is a pill as tall as the digit is high.
func digitZero(x, top, bottom float64) []string {
	const w = 46.0
	return []string{fmt.Sprintf(`<rect x="%s" y="%s" width="%s" height="%s" rx="%s"/>`,
		num(x), num(top), num(w), num(bottom-top), num(w/2))}
}

// ellipsis is dots zero-length lines: the round caps are the dots.
func ellipsis(x, top, bottom float64) []string {
	var out []string
	for i := 0; i < dots; i++ {
		out = append(out, fmt.Sprintf(`<path d="M%s %sh0"/>`, num(x+float64(i)*pitch), num(bottom)))
	}
	return out
}

// num prints a coordinate with at most one decimal and no trailing zero.
func num(v float64) string {
	s := fmt.Sprintf("%.1f", math.Round(v*10)/10)
	s = strings.TrimSuffix(s, ".0")
	if s == "-0" {
		s = "0"
	}
	return s
}

func main() {
	row := []item{
		{22, digitOne},
		{46, digitZero},
		{float64(dots-1) * pitch, ellipsis},
	}

	total := float64(len(row)-1) * gap
	for _, it := range row {
		total += it.width
	}

	x := 128 - total/2
	top, bottom := 128-height/2, 128+height/2
	var paths []string
	for _, it := range row {
		paths = append(paths, it.draw(x, top, bottom)...)
		x += it.width + gap
	}

	fmt.Fprintf(os.Stdout, `<svg viewBox="0 0 256 256" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="go-bignum">
  <defs><linearGradient id="bngrad" x1="0" y1="0" x2="1" y2="1"><stop offset="0" stop-color="%s"/><stop offset="1" stop-color="%s"/></linearGradient></defs>
  <rect width="256" height="256" rx="56" fill="url(#bngrad)"/>
  <g fill="none" stroke="#FFFFFF" stroke-width="%s" stroke-linecap="round" stroke-linejoin="round">
    %s
  </g>
</svg>
`, light, dark, num(stroke), strings.Join(paths, "\n    "))
}
