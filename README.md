## CLIコマンドライブラリ

### flagパッケージ

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

var (
    commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
)

type TextMarshaler interface {
    MarshalText() (text []byte, err error)
}

type TextUnmarshaler interface {
    UnmarshalText(text []byte) error
}

func run(args []string) int {
    max := commandLine.Int("max", 255, "max value")
    name := commandLine.String("name", "", "my name")
    if err := commandLine.Parse(args); err != nil {
        fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
    }

    if *max > 999 {
        fmt.Fprintf(os.Stderr, "invalid max value: %v\n", *max)
        return 1
    }

    if *name == "" {
        fmt.Fprintf(os.Stderr, "name must be provided")
        return 1
    }

    return 0
}

func main() {
    max := flag.Int("max", 255, "max value")
    name := flag.String("name", "something", "my name")
    flag.Parse()

    println(*name, *max)

    // ポインタではなく値が欲しい場合
    var name string
    var max int

    flag.IntVar(&max, "max", 255, "max value")
    flag.StringVar(&name, "name", "", "my name")
    flag.Parse()

    //引数取得
    for _, arg := range flag.Args() {
        fmt.Println(arg)
    }

    // 引数無しの場合
    if flag.NArg() == 0 {
        flag.Usage()
        os.Exit(1)
    }

    var sleep time.Duration
    flag.DurationVar(&sleep, "sleep", time.Second, "sleep time")
    flag.Parse()
    time.Sleep(sleep)

    var ip net.IP
    flag.TextVar(&ip, "ip", net.IPv4(127, 0, 0, 1), "ip address")
    flag.Parse()

    fmt.Println(ip)

    os.Exit(run(os.Args[1:]))
}
```

### 上記のコードのテストコード

```go
package main

import (
    "flag"
    "os"
    "testing"
)

func TestFlagVar(t *testing.T) {
    tests := []struct{
        name string
        args []string
        want int
    } {
        {name: "test1", args: []string{"-name", "foo"}, want: 0},
        {name: "test2", args: []string{"-name", "foo", "-max", "1000"}, want: 1},
        {name: "test1", args: []string{"-name", "", "-max", "123"}, want: 1},
    }

    for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            commandLine  = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
            if got := run(tt.args); got != tt.want {
                t.Errorf("%v: run() = %v, want %v", tt.name, got, tt.want)
            }
        })
    }
}
```

### urfave/cliパッケージ

```go
package main

import (
    "os"
    "github.com/urfave/cli/v2"
)

func cmdList(c *cli.Context) error {
    if c.Bool("json") {
        return listStudentsAsJSON()
    }
    return listStudents()
}

func main() {
    app := cli.NewApp()
    app.Flags = []cli.Flag{
        &cli.StringFlag{
            Name: "config",
            Aliases: []string{"c"}
            Usage: "Load configuration from `FILE`",
        }
    }
    // サブコマンド対応
    app.Commands = []*cli.Command{
        {
            Name: "list",
            Usage: "list students",
            Flags: []cli.Flag{
                &cli.BoolFlag{
                    Name: "json",
                    Usage: "outpuy as JSON",
                    Value: false,
                },
            },
            Action: cmdList,
        },
    }
    app.Name = "score"
    app.Usage = "Show student's score"
    app.Run(os.Args)
}
```

### alecthomas/kingpinパッケージ

```go
package main

import (
    "os"
    "strings"

    "gopkg.in/alecthomas/kingpin.v2"
)

var (
    app = kingpin.New("score", "Show student's score")
    debug = app.Flag("debug", "Enable debug mode").Bool()
    serverIP = app.Flag("server", "Server address").Default("127.0.0.1").IP()

    register = app.Command("register", "Register a new user")
    registerNick = register.Arg("nick", "Nickname for user").Required().String()
    registerName = register.Arg("name", "Name of user").Required().String()

    post = app.Command("post", "Post a message to a channel")
    postImage = post.Flag("image", "Image to post").File()
    postChannel = post.Arg("channel", "Channel to post to.").Required().String()
    postText = post.Arg("text", "Text to post").Strings()
)

func main() {
    switch kingpin.MustParse(app.Parse(os.Args[1:])) {
        case register.FullCommand():
            println(*registerNick)
        case post.FullCommand():
        if *postImage != nil {}

        text := strings.Join(*postText, " ")
        println("Post:", text)
    }
}
```

### spf13/cobra

```zsh
go install github.com/spf13/cobra-cli@latest

cobra-cli init -l MIT -a "test golang@example.com>"
# .cobra.yamlで設定することもあり

./main
# cmd/root.goに記述されている
# サブコマンド追加
cobra-cli add update
# cmd/update.go
```

```go
#フラグの追加
func init() {
    rootCmd.AddCommand(updateCmd)
    updateCmd.PersistentFlags().String("foo","", "A help for foo")
    updateCmd.Flags().BoolIP("toggle", "t", false, "Help message for toggle")
}

```

### 端末制御用ライブラリ

### olekukonko/tablewriter

```go
package main

import (
    "os"
    "github.com/olekukonko/tablewriter"
)

func main() {
    data := [][]string{
        {"A","The Good", "500"},
        {"B","The Very very Bad Man", "288"},
        {"C","The Ugly", "120"},
        {"D","The Gopher", "800"},
    }

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Name", "Sign", "Rating"})
    table.SetColumnAlignment([]int{
        tablewriter.ALIGN_CENTER,
        tablewriter.ALIGN_DEFAULT,
        tablewriter.ALIGN_DEFAULT,
    })

    table.SetHeaderColor(
        tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
        tablewriter.COlors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
        tablewriter.COlors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
    )

    table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
    table.SetFooter([]string{"","","427.0"})

    for _, v := range data {
        table.Append(v)
    }
    table.Render()
}
```

### mattn/go-runewidth

```go
package main

import (
    "fmt"
    "strings"

    "github.com/mattn/go-runewidth"
)

func main() {

    s := "Go言語でCLIアプリケーション作成"
    fmt.Println(s)
    width := runewidth.StringWidth(s)
    fmt.Println(strings.Repeat("~", width))

    fmt.Println(runewidth.Wrap(s, 11))
}
```

### jroimartin/gocui

```go
package main

import (
    "fmt"
    "log"
    "github.com/jroimartin/gocui"
)

func main() {
    g, err := gocui.NewGui(gocui.OutputNormal)
    if err != nil {
        log.Panicln(err)
    }
    defer g.Close()

    g.ASCII = true
    g.SetManagerFunc(layout)

    if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
        log.Panicln(err)
    }

    if err := g.MainLopp(); err != nil && err != gocui.ErrQuit {
        log.Panicln(err)
    }
}

func layout(g *gocui.Gui) error {
    s := `Go is an open source programming language that makes it simple to build secure, scalable systems.`
    maxX, maxY := g.Size()

    if v, err := g.SetView("hello", maxX/2-21, maxY.2-2, maxX/2+21, maxY/2+2); err != nil {
        if err != gocui.ErrUnknownView {
            return err
        }
        fmt.Fprintln(v, s)
    }
    return nil
}

func quit(g *gogui.Gui, v *gocui.View) error {
    return gocui.ErrQuit
}
```