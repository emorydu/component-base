package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件没有扩展名，则需要指定配置文件的格式，告诉viper使用何种方式解析文件
	} else {
		viper.AddConfigPath(".")          // 将当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.iam") // 配置文件搜索路径，可以配置多个
		viper.SetConfigName("config")     // 配置文件名称 !!! 没有文件扩展名
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件，如果指定了配置文件名，则使用指定的配置文件，否则在注册的路径中搜索
		panic(fmt.Errorf("fatal error config file `%s`: \n", err.Error()))
	}

	//if err := viper.ReadInConfig(); err != nil {
	//	var configFileNotFoundError viper.ConfigFileNotFoundError
	//	if errors.As(err, &configFileNotFoundError) {
	//		// 配置文件未找到错误；如果需要可以忽略
	//	} else {
	//		// 配置文件找到，但产生了其他错误
	//	}
	//}

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())

	// 写入配置文件
	//WriteConfig：保存当前配置到viper当前使用的配置文件中，如果配置文件不存在会报错，如果配置文件存在则覆盖当前配置文件
	//SafeWriteConfig：保存当前配置到viper当前使用的配置文件中，如果配置文件不存在会报错，如果配置文件存在则返回file exists错误
	//WriteConfigAs：保存当前配置到指定文件中，如果文件不存在则新建，如果文件存在则会覆盖文件
	//SafeWriteConfigAs：保存当前配置到指定文件中，如果文件不存在则新建，如果文件存在则返回file exists错误
	//viper.WriteConfig()
	//viper.SafeWriteConfig()
	//viper.WriteConfigAs()
	//viper.SafeWriteConfigAs()

	// 监听和重新读取配置文件
	// viper支持实时读取配置文件（热加载配置），通过WatchConfig函数热加载配置，在调用WatchConfig函数之前，确保已经添加了配置文件的搜索路径
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	// 配置文件发生变更之后会调用的回调函数
	//})

	// 设置配置值
	// 通过viper.Set()函数显示设置配置
	//viper.Set("username", "emory")

	// 注册和使用别名
	// viper别名可以允许多个键引用单个值
	//viper.RegisterAlias("loud", "verbose")
	//viper.Set("verbose", true)
	//viper.Set("loud", true)
	//
	//viper.GetBool("loud")    // true
	//viper.GetBool("verbose") // true

	// 使用环境变量
	//AutomaticEnv()
	//BindEnv(input ...string) error
	//SetEnvPrefix(in string)
	//SetEnvKeyReplacer(r *string.Replacer)
	//AllowEmptyEnv(allowEmptyEnv bool)
	//这里要注意：viper 读取环境变量是区分大小写的。viper 提供了一种机制来确保 ENV 变量是唯一的。通过使用 SetEnvPrefix，可以告诉 Viper 在读取环境变量时使用前缀。
	//BindEnv 和 AutomaticEnv 都将使用此前缀。比如，我们设置了 viper.SetEnvPrefix("VIPER")，当使用 viper.Get("apiversion") 时，实际读取的环境变量是 VIPER_APIVERSION。

	//BindEnv 需要一个或两个参数。第一个参数是键名，第二个是环境变量的名称，环境变量的名称区分大小写。如果未提供 ENV 变量名，则viper将假定ENV变量名为：
	//环境变量前缀_键名全大写 ，例如：前缀为VIPER，key为username，则ENV变量名为： VIPER_USERNAME 。当显式提供 ENV 变量名（第二个参数）时，它不会自动添加前缀。
	//例如，如果第二个参数是 id，Viper 将查找环境变量 ID。

	//在使用 ENV 变量时，需要注意的一件重要事情是，每次访问该值时都将读取它。Viper在调用 BindEnv 时不固定该值。
	//还有一个魔法函数 SetEnvKeyReplacer，SetEnvKeyReplacer 允许你使用 strings.Replacer 对象来重写 Env 键。如果你想在 Get() 调用中使用 - 或者 . ，但希望你的环境变量使用 _ 分隔符，
	//可以通过 SetEnvKeyReplacer 来实现。比如，我们设置了环境变量 USER_SECRET_KEY=bVix2WBv0VPfrDrvlLWrhEdzjLpPCNYb，但我们想用 viper.Get("user.secret-key")，我们调用函数：
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	//os.Setenv("VIPER_USER_SECRET_ID", "QLdywI2MrmDVjSSv6e95weNRvmteRjfKAuNV")
	//os.Setenv("VIPER_USER_SECRET_KEY", "bVix2WBv0VPfrDrvILWrhEdzjLpPCNYb")
	//
	//viper.AutomaticEnv() // 读取环境变量
	//viper.SetEnvPrefix("VIPER")
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	//viper.BindEnv("user.secret-key")
	//viper.BindEnv("user.secret-id", "USER_SECRET_ID")

	// 绑定标志
	//viper.BindPFlag("token", pflag.Lookup("token")) // 绑定单个标志
	//viper.BindPFlags(pflag.CommandLine)             // 绑定标志集

	// viper读取配置
	//Get(key string) interface{}
	//Get<Type>(key string) <Type>
	//AllSettings() map[string]interface{}
	//IsSet(key string) bool

}
