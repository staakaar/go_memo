package main

/** 上から順番にinit関数が呼ばれる */
/** SQLドライバーなどの初期化に使用するケースがある */
// _ "init/init1"
// _ "init/init2"

var name = "Jiro"

func init() {
	println("hi" + name)
}
