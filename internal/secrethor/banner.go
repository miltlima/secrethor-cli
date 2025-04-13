package secrethor

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func PrintBanner(version string) {
	fig := figure.NewFigure("Secrethor", "rectangles", true)
	cyan := color.New(color.FgCyan).Add(color.Bold)
	cyan.Println(fig.String())
	fmt.Println(" Secrethor-CLI - Version " + version + "\n")

}
