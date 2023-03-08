package termtablespkg

import (
	"fmt"

	"github.com/scylladb/termtables"
)

func TTP() {
	t := termtables.CreateTable()
	t.AddHeaders("User", "Age")
	t.AddRow("dj", 18)
	t.AddRow("darjun", 30)
	fmt.Println(t.Render())
}

func TTP2() {
	t := termtables.CreateTable()
	t.AddHeaders("User", "Age")
	t.AddRow("dj", 18)
	t.AddRow("darjun", 30)
	fmt.Println("HTML:")
	t.SetModeHTML()
	fmt.Println(t.Render())

	fmt.Println("Markdown:")
	t.SetModeMarkdown()
	fmt.Println(t.Render())
}
