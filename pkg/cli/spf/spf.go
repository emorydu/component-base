// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/spf13/pflag"
)

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
*/

func main() {
	var version bool
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.BoolVar(&version, "version", true, "Print version information and quit.")
}
