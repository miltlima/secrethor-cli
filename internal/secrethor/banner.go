package secrethor

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func PrintBanner() {
	fig := figure.NewFigure("Secrethor", "chunky", true)
	cyan := color.New(color.FgCyan).Add(color.Bold)
	cyan.Println(fig.String())
}
