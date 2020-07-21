# cmdnote

A simple note taking system for command line interface.



#### Installation

1. Install Go：https://golang.org/

2.  Download project source code

    ```shell
    $ git clone https://github.com/qige96/cmdnote.git
    ```

3. Compile (and install)

   ```shell
   $ cd cmdnote
   $ go build cmdnote.go # require third party dependency, may demand network
   $ go install # install to $GOPATH
   ```

#### Usage

You could do basic CURD to the notes. All notes are stored as files in your local repository.

```shell
$ cmdnote -w hello.txt # invoke your preferred editor to write a file
$ cmdnote -r hello.txt # invoke your preferred reader to read a file
$ cmdnote -l # list all available notes
$ cmdnote --rename hello.txt world.txt # rename a note
$ cmdnote --remove world.txt # remove a note
```

**Support full text search by keywords**

```shell
$ git -s "hello"
```

#### Configuration

Configuration file `conf.json` is located under the same directory as the executable file live, config for Windows may look like：

```json
{
        "editor": "notepad",
        "browser": "notepad",
        "localRepoDir": "F:\\cmdnote\\cmd_notes",
        "remoteRepoDir": ""
}
```

config for Linux or MacOS(Darwin) may look like：

```json
{
        "editor": "nano",
        "browser": "nano",
        "localRepoDir": "/home/ubuntu/cmdnote",
        "remoteRepoDir": ""
}
```



#### License

MIT