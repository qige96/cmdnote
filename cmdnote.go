package cmdnote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/blevesearch/bleve"
)

const (
	BleveFolder   = ".notes.bleve" // where note indices live
	VersionNumber = "0.1.0"
)

type Configuration struct {
	Editor        string `json:"editor"`
	Browser       string `json:"browser"`
	LocalRepoDir  string `json:"localRepoDir,omitempty"`
	RemoteRepoDir string `json:"remoteRepoDir,mitempty"`
}

var CONF Configuration = Configuration{
	GetDefaultEditorBrowser(),
	GetDefaultEditorBrowser(),
	GetDefaultLocalRepoDir(),
	"",
}

var (
	confPath  = path.Join(path.Dir(CONF.LocalRepoDir), "conf.json")
	BlevePath = path.Join(CONF.LocalRepoDir, BleveFolder)
)

func FullNotePath(noteTitle string) string {
	return path.Join(CONF.LocalRepoDir, noteTitle)
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func init() {
	if !Exist(confPath) {
		DumpConf(confPath)
	} else {
		LoadConf(confPath)
	}
	if !Exist(CONF.LocalRepoDir) {
		os.MkdirAll(CONF.LocalRepoDir, os.ModePerm)
	}
	if !Exist(BlevePath) {
		mapping := bleve.NewIndexMapping()
		_, err := bleve.New(BlevePath, mapping)
		if err != nil {
			log.Fatal(err)
		}
	}
	if !Exist(path.Join(CONF.LocalRepoDir, ".git")) {
		// git init
	}
}

func GetDefaultLocalRepoDir() string {
	return path.Join(getExecuteDir(), "cmd_notes")
}

func GetDefaultEditorBrowser() string {
	sysType := runtime.GOOS
	if sysType == "windows" {
		return "notepad"
	}
	if sysType == "linux" || sysType == "darwin" {
		return "nano"
	}
	return "vim"
}

// load configuration from CONF file
func LoadConf(confFilePath string) Configuration {
	content, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &CONF)
	if err != nil {
		log.Fatal(err)
	}
	return CONF
}

// dump confFilePath to CONF file
func DumpConf(confFilePath string) {
	content, err := json.MarshalIndent(CONF, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(confFilePath, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// get a time-formated filename
func getTimeFileName() string {
	t := time.Now()
	filename := t.Format("2006-01-02_15-04-05") + ".txt"
	return filename
}

// invoke external commands
func invoke(prog string, args []string) {
	cmd := exec.Command(prog, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// read a note
func readNote(prog, notePath string) {
	invoke(prog, []string{notePath})
}

// write a note
func writeNote(prog, notePath string) {
	fileDir := path.Dir(notePath)
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	invoke(prog, []string{notePath})

}

// recursively list all files under a directory
func AllFilePaths(dir string) []string {
	allFiles := []string{}
	err := filepath.Walk(dir,
		func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// fmt.Println(p, info.Size())
			if s, _ := os.Stat(p); !s.IsDir() {
				allFiles = append(allFiles, p)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}

	return allFiles
}

func AllNoteNames() []string {
	allFilePaths := AllFilePaths(CONF.LocalRepoDir)
	allNotes := []string{}
	for _, fpath := range allFilePaths {
		noteTitle, _ := filepath.Rel(CONF.LocalRepoDir, fpath)
		if strings.HasPrefix(noteTitle, BleveFolder) {
			continue
		}
		allNotes = append(allNotes, noteTitle)
	}
	return allNotes
}

// list all notes in local repository
func listNotes() {
	allNotes := AllNoteNames()
	for _, noteTitle := range allNotes {
		fmt.Println(noteTitle)
	}
}

// list all notes in local repository, and provide interactive inspection
func listNotesInteractive() {
	allNotes := AllNoteNames()
	for i, noteTitle := range allNotes {
		fmt.Printf("%5d) %s\n", i, noteTitle)
	}

	interactiveSession(allNotes)
}

// interactive sesion where users can continueously inspect a note
// with the program he wants
func interactiveSession(noteTitles []string) {
	if len(noteTitles) == 0 {
		fmt.Println("Sorry, nothing found!")
		return
	}

	var (
		noteId int
		prog   string
		err    error
	)

	for {
		fmt.Print("> Which note would you like to check? ")
		_, err = fmt.Scanln(&noteId)
		if err != nil {
			if err.Error() == "unexpected newline" {
				break
			} else {
				fmt.Println(err.Error())
			}
		}

		fmt.Print("> Which program would you like to use? ")
		_, err = fmt.Scanln(&prog)
		if err != nil {
			if err.Error() == "unexpected newline" {
				prog = CONF.Browser
			} else {
				fmt.Println(err.Error())
			}
		}

		if noteId < len(noteTitles) && noteId >= 0 {
			notePath := path.Join(CONF.LocalRepoDir, noteTitles[noteId])
			invoke(prog, []string{notePath})
		} else {
			fmt.Printf("noteId %d out of range [%d] - [%d] \n", noteId, 0, len(noteTitles)-1)
		}

	}
}

// remove a note
func remoteNote(notePath string) {
	err := os.Remove(notePath)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// rename a note
func renameNote(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// get the path of the executable file
func getExecuteDir() string {
	dir, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	exPath := filepath.Dir(dir)

	return exPath
}

func main() {
	parseArges()
}
