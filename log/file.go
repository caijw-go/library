package log

import (
    "encoding/json"
    "fmt"
    "github.com/caijw-go/library/consts"
    "github.com/caijw-go/library/server"
    "github.com/gin-gonic/gin"
    "log"
    "os"
    "path"
    "sync"
    "time"
)

var _locker sync.Mutex    //全局锁
var _logger *log.Logger   //当前正在操作日志的指针实例
var _currentFile *os.File //当前正在写入日志的文件指针

type content struct { //FlowID后期再考虑要不要加上
    //Context *gin.Context
    //Flow    string
    //FlowID  string
    Level string
    Msg   string
    Ext   []interface{}
}

func (t *content) ToString() string {
    var now = time.Now()
    bytes, err := json.Marshal(gin.H{
        "app":   server.AppName(),
        "level": t.Level,
        "msg":   t.Msg,
        "ext":   t.Ext,
        "time":  now.Unix(),
        "micro": now.UnixNano() / 1e6,
    })
    if err != nil {
        return ""
    }
    return string(bytes)
}

// Init 初始化文件类日志系统
func Init(logPath string) error {
    newName := fmt.Sprintf("%s/%s.json", path.Dir(logPath), time.Now().Format(consts.TimeFormatDate))

    //log ready
    if _currentFile != nil {
        if _currentFile.Name() != newName {
            _locker.Lock()
            _currentFile.Close()
        } else if _, err := os.Stat(_currentFile.Name()); err != nil {
            _locker.Lock()
            _currentFile.Close()
        } else {
            return nil
        }
    } else {
        _locker.Lock()
    }

    //lock log file to ready
    defer _locker.Unlock()
    var f *os.File
    var err error
    f, err = os.OpenFile(newName, os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        if os.IsNotExist(err) {
            if f, err = os.Create(newName); err == nil {
                _currentFile = f
                goto SUCCESS
            }
        }
    } else {
        _currentFile = f
        goto SUCCESS
    }

    _logger = log.New(os.Stdout, "", 0)
    return err

SUCCESS:
    _logger = log.New(_currentFile, "", 0)
    return err
}

//Info 记录本地文本类Info日志
func Info(msg string, ext ...interface{}) {
    write("info", msg, ext)
}

func Error(msg string, ext ...interface{}) {
    write("error", msg, ext)
}

func write(level, msg string, ext []interface{}) {
    str := (&content{
        Level: level,
        Msg:   msg,
        Ext:   ext,
    }).ToString()
    if !server.IsProd() {
        log.Println(str)
    }
    go _logger.Println(str)
}
