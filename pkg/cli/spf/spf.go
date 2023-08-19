// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

/*
type Flag struct {
	Name                string              // flag 长选项的名称
	Shorthand           string              // flag 短选项的名称，一个缩写的字符
	Usage               string              // flag 的使用文本
	Value               Value               // flag 的值
	DefValue            string              // flag 的默认值
	Changed             bool                // 记录 flag 的值是否有被设置过
	NoOptDefVal         string              // 当 flag 出现在命令行，但是没有指定选项值时的默认值
	Deprecated          string              // 记录该 flag 是否被放弃
	Hidden              bool                // 如果值为 true，则从 help / usage 输出信息中隐藏该 flag
	ShorthandDeprecated string              // 如果 flag 的短选项被废弃，当使用 flag 的短选项时打印该信息
	Annotations         map[string][]string // 给 flag 设置注解
}

type Value interface {
	String() string
	Set(string) error
	Type() string
}

pflag除了支持单个的Flag之外，还支持FlagSet。FlagSet是一些预先定义好的Flag的集合，几乎所有的pflag操作，都需要借助
FlagSet提供的方法来完成
1. NewFlagSet -> FlagSet
2. pflag 包定义了全局的FlagSet: CommandLine -> (NewFlagSet(os.Args[0], ExitOnError))
*/

func main() {
	//var version bool
	//flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	//flagSet.BoolVar(&version, "version", true, "Print version information and quit.")

	//var name *string = pflag.String("name", "emory", "Input Your Name")
	//
	//name = pflag.StringP("name", "n", "emory", "Input Your Name")
	//_ = name
	//
	//var nm string
	//pflag.StringVar(&nm, "name", "emory", "Input Your Name")
	//
	//pflag.StringVarP(&nm, "name", "n", "emory", "Input Your Name")

	//number := pflag.Int("number", 1234, "help message for number")
	//pflag.Parse()
	//
	//fmt.Printf("argument number is: %v\n", pflag.NArg())    // 非命令行标志参数个数（除了文件名之外都算）
	//fmt.Printf("argument list is: %v\n", pflag.Args())      // 非命令行标志参数（除了文件名之外都是）
	//fmt.Printf("the first argument is: %v\n", pflag.Arg(0)) // 首个非命令行标志参数
	//_ = number

	//var ip = pflag.StringP("ip", "p", "127.0.0.1", "help for ip address")
	//pflag.Lookup("ip").NoOptDefVal = "0.0.0.0"

}
