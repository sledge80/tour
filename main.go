package main

import (
	"flag"
	"log"
	"tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("util.Execute err: %v", err)
	}
}

func ch2() {
	var name string
	//flag.StringVar(&name, "name", "默认值", "帮助信息")
	//flag.StringVar(&name, "n", "默认值", "帮助信息")

	flag.Parse()
	//log.Printf("name:%s", name)

	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "java":
		javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
		javaCmd.StringVar(&name, "n", "Java 语言", "帮助信息")
		_ = javaCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}

func ch1() {
	var name string
	flag.StringVar(&name, "name", "默认值", "帮助信息")
	flag.StringVar(&name, "n", "默认值", "帮助信息")

	flag.Parse()
	log.Printf("name:%s", name)
}
