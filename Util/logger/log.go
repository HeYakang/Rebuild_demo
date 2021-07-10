package logger


import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Log *zap.Logger


func InitLogger(project, path string) *zap.Logger {
	// 方法1

	//要将日志写到哪里去
	writeSyncer := getLogWriter(project, path)
	//编解码配置(如何写入日志)
	encoder := getEncoder()
	//打包编码与写入位置与写入等级（下面创建会使用）
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	//Zap.New手动传递所有配置
	//需要配置所有Encoder WriteSyncer LogLevel（上面一条方法创建为core传入）
	//接下来，我们将修改zap logger代码，添加将调用函数信息记录到日志中的功能。为此，我们将在zap.New(..)函数中添加一个Option。
	Log = zap.New(core, zap.AddCaller())
	// sugarLogger = logger.Sugar()

	// 方法2, 比较简洁
	// cfg := zap.NewProductionConfig()
	// cfg.OutputPaths = []string{"stdout", "./test.log"}
	// l, _ := cfg.Build()
	// logger = l
	return Log
}

func getEncoder() zapcore.Encoder {
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	//此处要做的事情是
	//1.修改时间编码器
	//2.在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//此处希望编码器从JSON Encoder更改为普通Encoder。因此将NewJSONEncoder（）改为下面返回
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 所在目录必须预先创建

func getLogWriter(project, path string) zapcore.WriteSyncer {
	logDir := path+"/log"
	if project != "" {
		logDir = "log/" + project
	}
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		panic(err)
	}
	consolSyncers, _, err := zap.Open(fmt.Sprintf("%s/%s.log", logDir, time.Now().Format("2006-01-02_15:04:05")), "stdout") // stdout:打印日志在consol
	if err != nil {
		panic(err)
	}

	// defer close()
	return consolSyncers
}

//参考资料：
//https://www.cnblogs.com/Golanguage/p/12285584.html


/*下面附上网络写法 进行参考
func InitLogger() {
    writeSyncer := getLogWriter()
    encoder := getEncoder()
    core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

    logger := zap.New(core)
    sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
    return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
    file, _ := os.Create("./test.log")
    return zapcore.AddSync(file)
}

*/
