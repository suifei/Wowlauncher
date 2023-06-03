package main

import (
	"fmt"
	"os"
	"os/exec"
)

func writeRealmlist(path, content string) error {
	if _, err := os.Stat(path); err == nil {
		fmt.Println(path, err)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.WriteString(content)
		return err
	} else{
		return err
	}
}

func startWow() error {
	cmd := exec.Command("start", wowExecute)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

var realmlists = []string{
	"Data/zhTW/realmlist.wtf",
	"Data/enTW/realmlist.wtf",
	"Data/zhCN/realmlist.wtf",
}
var wowExecute = "Wow.exe"

func main() {

	errCount := 0
	for _, realmlist := range realmlists {
		err := writeRealmlist(realmlist,"SET realmlist \"47.111.254.152\"")
		if err != nil {
			errCount ++
		}
	}

	if errCount == len(realmlists){
		fmt.Println("realmlist not found，请确认是否安装魔兽世界客户端，并将该程序放到和 WoW.exe 同样位置。")
		return
	}

	err := startWow()
	if err != nil {
		fmt.Println("启动魔兽世界 WoW.exe 失败，请确认是否安装魔兽世界客户端，并将该程序放到和 WoW.exe 同样位置。")
	}
}
