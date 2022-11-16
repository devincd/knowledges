package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
Viper 使用以下优先级顺序，上面会覆盖下面的配置
- 明确调用 Set() 函数
- flag
- env
- config
- key/value store
- default
*/

func main() {
	// Establishing Defaults
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	fmt.Println(viper.Get("ContentDir"))

	// Reading Config Files
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("not found", err)
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
		}
	}

	// Writing Config Files
	/*
		WriteConfig - writes the current viper configuration to the predefined path, if exists.
		Errors if no predefined path. Will overwrite the current config file, if it exists.
		SafeWriteConfig - writes the current viper configuration to the predefined path.
		Errors if no predefined path. Will not overwrite the current config file, if it exists.
		WriteConfigAs - writes the current viper configuration ti the given filepath.
		Will overwrite the given file, if it exists.
		SafeWriteConfigAs - writes the current viper configuration to the given filepath.
		Will not overwrite the given file, if it exists.
	*/
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}

	// Watching and re-reading config files
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}
