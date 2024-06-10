## データーベースパッケージ

## データベース サードパーティ

|データベース|リポジトリ|
|:---|:---|
|MySQL|https://github.com/go-sql-driver/mysql https://github.com/ziutek/mymysql|
|PostgreSQL|https://github.com/lib/pq https://github.com/jackc/pgx|
|SQLite|https://github.com/mattn/go-sqlite3|

## Query
```go
package main

import (
    "database/sql"
    "log"
    _ "github.com/lib/pg"
)

func main() {
    db, err := sql.Open("postgres", "postgres://postgres:postgres@dbserver/database")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    id := 3
    /** 件数１件の場合
     * err = db.QueryRow(`SELECT name, age FROM users WHERE id = $1`, id).Scan(&name, &age) 
     */
    /** 
     * Context クエリの途中中断
     * ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
     * defer cancel()
     * rows, err := db.QueryContext(ctx, `SELECT name, age FROM users WHERE id < $1`, id)
     * if err != nil {
     *  log.Fatal(err)
     * }
     */
    rows, err := db.Query(`SELECT name, age FROM users WHERE id < $1`, id)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var name string
        var age int
        err = rows.Scan(&name, &age)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("name=%v, age=%v\n", name, age)
    }
}
```

## Exec

```go
err := db.Exec(`UPDATE users SET name = $1 WHERE id = $2`, name, id)
if err != nil {
    log.Fatal(err)
}
```

## Prepare
複数のクエリや実行を扱いたい場合、ステートメントをあらかじめ作成、再利用
```go
stmt, err := db.Prepare(`SELECT name, age FROM users WHERE id < $1`)
if err != nil {
    log.Fatal(err)
}
err = stmt.QueryRow(id).Scan(&name, &age)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("name=%v, age%v\n", name, age)
```

## Begin Commit Rollback
```go
tx, err := db.Begin()
if err != nil {
    return err
}
defer tx.Rollback()
result, err := tx.Exec(`INSERT INTO users(name, age) VALUES($1, $2)`, "taro", 41)
if err != nil {
    return err
}
if affected, err := result.RowsAffected(); err != nil {
    return err
} else if affected == 0 {
    return errors.New("no record affected")
}
tx.Commit()
```

## ent/ent

### スキーマ定義の作成
```zsh
$ mkdir todo
$ cd $_
$ go mod init todo

$ go run -mod=mod entgo.io/ent/cmd/ent init Todo
```

```go
package schema

import "entgo.io/ent"

type Todo struct {
    ent.Schema
}

func (Todo) Fields() []ent.Field {
    return []ent.Field{
        field.Text("text").NotEmpty(),
        field.Time("created_at").Default(time.Now).Immutable(),
        field.Enum("status").NamedValues("InProgress", "IN_PROGRESS", "Completed", "COMPLETED",).Default("IN_PROGRESS"),
        field.Int("priority").Default(0),
    }
}

func (Todo) Edges() []ent.Edge {
    return nil
}

/** go generate ./ent　ディレクトリ内にTodoを操作するためのAPI */
```

```go
client, err := ent.Open("postgres", "postgres://postgres:postgres@localhost/test?sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer client.Close()

err := db.Schema.Create(context.Background())
if err != nil {
    log.Fatalf("failed creating schema resources: %v", err)
}

for _, e := range client.Query().AIIX(context.Background()) {
    fmt.Println(e.Text)
}

_, err = client.Todo.Create().SetText("test").Save(context.Background())
if err != nil {
    log.Fatalf("failed creating a todo: %v", err)
}

```

## OpenAPI生成

entが生成したディレクトリ内でent/generate_openapi.go

```go
//go:build ignore
// +build ignore

package main

import (
    "log"
    "entgo.io/contrib/entoas"
    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
)

func main() {
    ex, err := entoas.NewExtension()
    if err != nil {
        log.Fatalf("creating entoas extension: %v", err)
    }

    err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))

    if err != nil {
        log.Fatalf("running ent codegen: %v", err)
    }
}

/** ent/generate.goの編集 */

package ent
//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
//go:generate go run -mod=mod generate_openapi.go

/** 以下のコマンド実行後にopenapi.jsonが生成される */
```

```zsh
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
$ oapi-codegen -package main -generate server -old-config-style ent/openapi.json > oapi.go
```

```go
package main

import (
    "context"
    "databse/sql"
    "encoding/json"
    "errors"
    "log"
    "os"

    "todo/ent"
    "todo/ent/todo"
    "github.com/labstack/echo/v4"
    _ "github.com/lib/pq"
)

type Api struct {
    client *ent.Client
}

type ListTodoParams struct {
    Page *int
    ItemsPerPage *int
}

func (a *Api) ListTodo(ctx echo.Context, params ListTodoParams) error {
    page := 0
    if params.Page != nil {
        page = *params.Page
    }
    itemsPerPage := 5

    if params.ItemsPerPage != nil {
        itemsPerPage = *params.ItemsPerPage
    }

    ees, err := a.client.Todo.Query().Order(ent.Desc(todo.FieldID)).Offset(page * itemsPerPage).Limit(itemsPerPage).All(context.Background())

    if err != nil {
        log.Println(err)
        return echo.ErrBadRequest
    }
    return ctx.JSON(200, ees)
}

func (a *Api) CreateTodo(ctx echo.Context) error {
    var ee ent.Todo
    err := json.NewDecoder(ctx.Request().Body()).Decode(&ee)
    if err != nil {
        log.println(err)
        return echo.ErrBadRequest
    }

    e := a.client.Todo.Create().SetText(ee.Text).SetStatus(ee.Status).SetPriority(ee.Priority)

    if !ee.CreatedAt.IsZero() {
        e.SetCreatedAt(ee.CreatedAt)
    }

    if ee2, err := e.Save(context.Background()); err != nil {
        log.Println(err)
        return echo.ErrBadRequest
    } else {
        ee = *ee2
    }
    return ctx.JSON(200, ee)
}

func (a *Api) DeleteTodo(ctx echo.Context, id int) error {
    e := a.client.DeleteOneID(int(id))
    err := e.Exec(context.Background())
    if err != nil {
        log.Println(err)
        return echo.ErrBadRequest
    }
    return nil
}

func (a *Api) ReadTodo(ctx echo.Context, id int) error {
    e, err := a.client.Todo.Get(context.Bckground(), int(id))
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return echo.ErrNotFound
        }
        log.Println(err)
        return echo.ErrBadRequest
    }
    return ctx.JSON(200, e)
}

func (a *Api) UpdateTodo(ctx echo.Context, id int) error {
    var ee ent.Todo
    err := json.NewDecoder(ctx.Request().Body).Decode(&ee)
    if err != nil {
        log.Println(err)
        return echo.ErrBadRequest
    }
    e := a.client.Todo.UpdateOneID(int(id)).SetText(ee.Text).SetStatus(ee.Status).SetPriority(ee,Priority)
    if ee2, err := e.Save(context.Background()); err != nil {
        log.Println(err)
        return echo.ErrBadRequest
    } else {
        ee = *ee2
    }
    return ctx.JSON(200, ee)
}

func main() {
    client, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("failed opening connection to sqlite: %v", err)
    }

    client.Schema.Create(context.Background())

    e := echo.New()
    myApi := &Api{client: client}
    RegisterHandlers(e, myApi)
    e.Static("/", "static")
    e.Logger.Fatal(e.Start(":8989"))
}
```