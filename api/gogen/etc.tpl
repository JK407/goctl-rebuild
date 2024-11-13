Name: {{.ServiceName}}
Host: {{.Host}}
Port: {{.Port}}
Mysql:
    - {{ .Mysql.TableName }}:
        User: {{ .Mysql.User }}
        Password: {{ .Mysql.Password }}
        Addr: {{ .Mysql.Addr }}
        Port: {{ .Mysql.Port }}
Log:
    # 服务名称
    ServiceName: {{.Log.ServiceName}}
    # 日志打印模式，console 控制台 console,file,volume
    Mode: {{.Log.Mode}}
    # 日志格式, json 格式 或者 plain 纯文本	json, plain
    Encoding: {{.Log.Encoding}}
    # 日期格式化
    TimeFormat: {{.Log.TimeFormat}}
    # 日志在文件输出模式下，日志输出路径
    Path: {{.Log.Path}}
    # 日志级别
    Level: {{.Log.Level}}
    # 是否压缩日志
    Compress: {{.Log.Compress}}
    # 是否开启 stat 日志
    Stat: {{.Log.Stat}}
    # 日志保留天数，只有在文件模式才会生效
    KeepDays: {{.Log.KeepDays}}
    # 堆栈打印冷却时间
    StackCoolDownMillis: {{.Log.StackCoolDownMillis}}
    MaxSize: {{.Log.MaxSize}}
    Rotation: {{.Log.Rotation}}