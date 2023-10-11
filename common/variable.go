package common

import (
	"io/fs"
	"strings"
)

const (
	LumikaVersionNum                = 3
	LumikaVersionString             = "v3.11.0"
	LumikaWorkDirName               = "lumika_data"
	LumikaConfigFileName            = "lumika_config"
	InitStr                         = "Init:"
	WebStr                          = "WebServer:"
	DbStr                           = "Database:"
	EnStr                           = "Encode:"
	DeStr                           = "Decode:"
	AddStr                          = "AddInput:"
	GetStr                          = "GetInput:"
	DlStr                           = "Dl:"
	BDlStr                          = "BDl:"
	BUlStr                          = "BUl:"
	ArStr                           = "AutoRun:"
	ErStr                           = "Error:"
	AddMLevel                       = 90
	AddKLevel                       = 81
	AddMGLevel                      = 200
	AddKGLevel                      = 130
	EncodeVideoSizeLevel            = 32
	EncodeOutputFPSLevel            = 24
	EncodeMaxSecondsLevel           = 35990
	EncodeFFmpegModeLevel           = "medium"
	DefaultHashLength               = 7
	DefaultBlankSeconds             = 3
	DefaultBlankByte                = 85
	DefaultBlankStartByte           = 86
	DefaultBlankEndByte             = 87
	DefaultDeleteFecFiles           = true
	DefaultBiliDownloadReferer      = "https://www.bilibili.com"
	DefaultDlMaxRetries             = 3
	DefaultWebServerDebugMode       = false
	DefaultWebServerHost            = ""
	DefaultWebServerPort            = 7860
	DefaultWebServerRandomPortMin   = 10000
	DefaultWebServerRandomPortMax   = 65535
	DefaultBiliDownloadGoRoutines   = 16
	DefaultBiliDownloadsMaxQueueNum = 5
	DefaultBiliUploadLines          = "ws"
	DefaultBiliUploadThreads        = 10
	DefaultTaskWorkerGoRoutines     = 5
	DefaultDbCrontabSeconds         = 10
)

var (
	EpPath                 string
	LumikaWorkDirPath      string
	LumikaEncodePath       string
	LumikaDecodePath       string
	LumikaEncodeOutputPath string
	LumikaDecodeOutputPath string
	MobileFFmpegPath       = ""
	MobileFFprobePath      = ""
	UISubFiles             fs.FS
)

type CommonError struct {
	Msg string
}

func (e *CommonError) Error() string {
	return e.Msg
}

type Database struct {
	DlTaskList  map[string]*DlTaskListData  `json:"dlTaskList"`
	BDlTaskList map[string]*BDlTaskListData `json:"bDlTaskList"`
	AddTaskList map[string]*AddTaskListData `json:"addTaskList"`
	GetTaskList map[string]*GetTaskListData `json:"getTaskList"`
	BUlTaskList map[string]*BUlTaskListData `json:"bUlTaskList"`
	VarSettings *VarSettings                `json:"VarSettingsVariable"`
}

type VarSettings struct {
	DefaultMaxThreads               int `json:"defaultMaxThreads"`
	DefaultBiliDownloadGoRoutines   int `json:"defaultBiliDownloadGoRoutines"`
	DefaultBiliDownloadsMaxQueueNum int `json:"defaultBiliDownloadsMaxQueueNum"`
	DefaultTaskWorkerGoRoutines     int `json:"defaultTaskWorkerGoRoutines"`
	DefaultDbCrontabSeconds         int `json:"defaultDbCrontabSeconds"`
}

type FecFileConfig struct {
	Version       int      `json:"v"`
	Name          string   `json:"n"`
	Summary       string   `json:"s"`
	Hash          string   `json:"h"`
	M             int      `json:"m"`
	K             int      `json:"k"`
	MG            int      `json:"mg"`
	KG            int      `json:"kg"`
	Length        int64    `json:"l"`
	SegmentLength int64    `json:"sl"`
	FecHashList   []string `json:"fhl"`
}

type ThreadInfo struct {
	ThreadIndex  int
	StartOffset  int64
	EndOffset    int64
	TempFilePath string
}

type FileInfo struct {
	Filename  string `json:"filename"`
	ParentDir string `json:"parentDir"`
	Type      string `json:"type"`
	SizeNum   int64  `json:"sizeNum"`
	SizeStr   string `json:"sizeStr"`
	Timestamp string `json:"timestamp"`
}

type SystemResourceUsage struct {
	OSName                string  `json:"osName"`
	CpuUsagePercent       float64 `json:"cpuUsagePercent"`
	MemUsageTotalAndUsed  string  `json:"memUsageTotalAndUsed"`
	MemUsagePercent       float64 `json:"memUsagePercent"`
	DiskUsageTotalAndUsed string  `json:"diskUsageTotalAndUsed"`
	DiskUsagePercent      float64 `json:"diskUsagePercent"`
	NetworkInterfaceName  string  `json:"networkInterfaceName"`
	UploadSpeed           string  `json:"uploadSpeed"`
	DownloadSpeed         string  `json:"downloadSpeed"`
	UploadTotal           string  `json:"uploadTotal"`
	DownloadTotal         string  `json:"downloadTotal"`
}

type DlTaskInfo struct {
	Url        string `json:"url"`
	FileName   string `json:"fileName"`
	Referer    string `json:"referer"`
	UserAgent  string `json:"userAgent"`
	NumThreads int    `json:"numThreads"`
}

type DlTaskListData struct {
	UUID         string      `json:"uuid"`
	TimeStamp    string      `json:"timestamp"`
	FileName     string      `json:"fileName"`
	TaskInfo     *DlTaskInfo `json:"taskInfo"`
	LogCat       string      `json:"logCat"`
	ProgressRate int         `json:"progressRate"`
	ProgressNum  float64     `json:"progressNum"`
	Status       string      `json:"status"`
	StatusMsg    string      `json:"statusMsg"`
	Duration     string      `json:"duration"`
}

type BDlTaskInfo struct {
	ResourceID string `json:"resourceID"`
	ParentDir  string `json:"parentDir"`
	BaseStr    string `json:"baseStr"`
}

type BDlTaskListData struct {
	UUID         string       `json:"uuid"`
	TimeStamp    string       `json:"timestamp"`
	ResourceID   string       `json:"resourceId"`
	TaskInfo     *BDlTaskInfo `json:"taskInfo"`
	BaseStr      string       `json:"baseStr"`
	LogCat       string       `json:"logCat"`
	ProgressRate int          `json:"progressRate"`
	ProgressNum  float64      `json:"progressNum"`
	Status       string       `json:"status"`
	StatusMsg    string       `json:"statusMsg"`
	Duration     string       `json:"duration"`
}

type AddTaskInfo struct {
	FileNameList     []string `json:"fileNameList"`
	DefaultM         int      `json:"defaultM"`
	DefaultK         int      `json:"defaultK"`
	MGValue          int      `json:"mgValue"`
	KGValue          int      `json:"kgValue"`
	VideoSize        int      `json:"videoSize"`
	OutputFPS        int      `json:"outputFPS"`
	EncodeMaxSeconds int      `json:"encodeMaxSeconds"`
	EncodeThread     int      `json:"encodeThread"`
	EncodeFFmpegMode string   `json:"encodeFFmpegMode"`
	DefaultSummary   string   `json:"defaultSummary"`
}

type AddTaskListData struct {
	UUID         string       `json:"uuid"`
	TimeStamp    string       `json:"timestamp"`
	TaskInfo     *AddTaskInfo `json:"taskInfo"`
	LogCat       string       `json:"logCat"`
	BaseStr      string       `json:"baseStr"`
	IsPaused     bool         `json:"isPaused"`
	ProgressRate int          `json:"progressRate"`
	ProgressNum  float64      `json:"progressNum"`
	Status       string       `json:"status"`
	StatusMsg    string       `json:"statusMsg"`
	Duration     string       `json:"duration"`
}

type GetTaskInfo struct {
	DirName      string `json:"dirName"`
	DecodeThread int    `json:"decodeThread"`
	BaseStr      string `json:"baseStr"`
}

type GetTaskListData struct {
	UUID         string       `json:"uuid"`
	TimeStamp    string       `json:"timestamp"`
	TaskInfo     *GetTaskInfo `json:"taskInfo"`
	LogCat       string       `json:"logCat"`
	IsPaused     bool         `json:"isPaused"`
	Filename     string       `json:"filename"`
	ProgressRate int          `json:"progressRate"`
	ProgressNum  float64      `json:"progressNum"`
	Status       string       `json:"status"`
	StatusMsg    string       `json:"statusMsg"`
	Duration     string       `json:"duration"`
}

type biliUpUser struct {
	SESSDATA        string `json:"SESSDATA"`
	BiliJct         string `json:"bili_jct"`
	DedeUserID      string `json:"DedeUserID"`
	DedeuseridCkmd5 string `json:"DedeUserID__ckMd5"`
	AccessToken     string `json:"access_token"`
}

type biliUpVideoInfos struct {
	Tid         int      `json:"tid"`
	Title       string   `json:"title"`
	Aid         string   `json:"aid,omitempty"`
	Tag         []string `json:"tag,omitempty"`
	Source      string   `json:"source,omitempty"`
	Cover       string   `json:"cover,omitempty"`
	CoverPath   string   `json:"cover_path,omitempty"`
	Description string   `json:"description,omitempty"`
	Copyright   int      `json:"copyright,omitempty"`
}

type BUlTaskInfo struct {
	FileName    string           `json:"fileName"`
	Cookie      *biliUpUser      `json:"cookie"`
	UploadLines string           `json:"uploadLines"`
	Threads     int              `json:"threads"`
	VideoInfos  biliUpVideoInfos `json:"videoInfos"`
}

type BUlTaskListData struct {
	UUID         string       `json:"uuid"`
	TimeStamp    string       `json:"timestamp"`
	FileName     string       `json:"fileName"`
	BVID         string       `json:"bvid"`
	TaskInfo     *BUlTaskInfo `json:"taskInfo"`
	LogCat       string       `json:"logCat"`
	ProgressRate int          `json:"progressRate"`
	ProgressNum  float64      `json:"progressNum"`
	Status       string       `json:"status"`
	StatusMsg    string       `json:"statusMsg"`
	Duration     string       `json:"duration"`
}

var DatabaseVariable Database

var VarSettingsVariable VarSettings

var LogVariable strings.Builder

var DlTaskQueue chan *DlTaskListData
var DlTaskList map[string]*DlTaskListData

var BDlTaskQueue chan *BDlTaskListData
var BDlTaskList map[string]*BDlTaskListData

var AddTaskQueue chan *AddTaskListData
var AddTaskList map[string]*AddTaskListData

var GetTaskQueue chan *GetTaskListData
var GetTaskList map[string]*GetTaskListData

var BUlTaskQueue chan *BUlTaskListData
var BUlTaskList map[string]*BUlTaskListData

var UploadTotalStart int64 = -1
var DownloadTotalStart int64 = -1
