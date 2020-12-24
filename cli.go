package cmdnote

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	version bool
	help    bool

	ToRead   string
	ToWrite  string
	IsList   bool
	Keywords string
	ToRemove string
	ToRename string

	IsInteractive bool
)

func init() {
	flag.BoolVar(&version, "version", false, "display version number")
	flag.BoolVar(&version, "v", false, "display version number")

	flag.BoolVar(&help, "help", false, "display help info")
	flag.BoolVar(&help, "h", false, "display help info")

	flag.StringVar(&ToRead, "read", "", "read a note")
	flag.StringVar(&ToRead, "r", "", "read a note")

	flag.StringVar(&ToWrite, "write", "", "write a note")
	flag.StringVar(&ToWrite, "w", "", "write a note")

	flag.BoolVar(&IsList, "list", false, "list a note")
	flag.BoolVar(&IsList, "l", false, "list a note")

	flag.StringVar(&Keywords, "search", "", "search a note")
	flag.StringVar(&Keywords, "s", "", "search a note")

	flag.StringVar(&ToRemove, "remove", "", "remove a note")

	flag.StringVar(&ToRename, "rename", "", "rename a note")

	flag.BoolVar(&IsInteractive, "interactive", false, "use interactive mode")
	flag.BoolVar(&IsInteractive, "i", false, "use interactive mode")

	flag.Usage = usage
}

func parseArges() {
	flag.Parse()

	if version {
		fmt.Println(VersionNumber)
	}

	if help {
		flag.Usage()
	}

	if ToRead != "" {
		readNote(CONF.Browser, FullNotePath(ToRead))
	}

	if ToWrite != "" {
		writeNote(CONF.Editor, FullNotePath(ToWrite))
		IndexNote(ToWrite)
	}

	if IsList {
		if IsInteractive {
			listNotesInteractive()
		} else {
			listNotes()
		}
	}

	if Keywords != "" {
		if IsInteractive {
			searchNotesInteractive(Keywords)
		} else {
			searchNotes(Keywords)
		}
	}

	if ToRemove != "" {
		remoteNote(FullNotePath(ToRemove))
		DeleteIndex(ToRemove)
	}

	if ToRename != "" {
		params := strings.Split(ToRename, ":")
		if len(params) != 2 {
			log.Fatal("This flag takes an arg in <oldname:newname> format")
		}
		renameNote(FullNotePath(params[0]), FullNotePath(params[1]))
		DeleteIndex(params[0])
		IndexNote(params[1])
	}

}

var helpUsage = `
Usage: cmdnote [options]

-r --read   <notename>                read a note
-w --write  <notename>                write a note
-l --list                             list all notes
     -i --interactive                   provide interactive inspection
--remove    <notename>                remove a note
--rename    <oldname:newname>         rename a note
-s --search <keywords>                search for a note
     -i --interactive                   provide interactive inspection
`

func usage() {
	fmt.Println(helpUsage)
}
