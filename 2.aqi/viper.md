# viper

环境变量：是储存在操作系统本地的一些键值对？

设置环境变量的前缀，会匹配前缀为设置值的环境变量，（环境变量还没用到过，这一块需要多了解）
viper.SetEnvPrefix("")

绑定环境变量，让viper优先读取环境变量
viper.AutomaticEnv()

设置配置文件名和类型，例如config.yaml。如设置了环境变量，环境变量中没有时会搜索配置文件。
viper.SetConfigName(acf.ConfigName)
viper.SetConfigType(acf.ConfigType)

设置配置文件路径，一般是相对路径？
viper.AddConfigPath(acf.ConfigPath)

读取配置文件后，保存至缓存，增加读取速度
viper.ReadInConfig()

返回viper.ReadInConfig()读取到的缓存值，以键值对的形式存在，保存在appconfig的Configs map[string]any 内部
viper.AllSettings()