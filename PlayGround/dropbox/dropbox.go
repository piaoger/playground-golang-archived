/*
// How to get Access Token
// https://github.com/wen866595/godropbox
// http://coderbee.net/index.php/go/20130829/427
package main

import (

    "github.com/scottferg/Dropbox-Go/dropbox"
    "fmt"
    "os"
    "io/ioutil"
)

func main() {

    s := dropbox.Session{
        AppKey:     "my_app_key",
        AppSecret:  "my_app_secret",
        AccessType: "token",
    }
    p := dropbox.Parameters{
        Rev: "01b3f45a",
        FileLimit: "50",
    }
    s.ObtainRequestToken()
    dropbox.GenerateAuthorizeUrl(s.Token.Key, &p)
    s.ObtainAccessToken()

    u := dropbox.Uri{
        Root: "Documents",
        Path: "Documents/setup.md",
    }

    // Download the file
    if file, _, err := dropbox.GetFile(s, u, nil); err == nil {
        ioutil.WriteFile("./test_result.pdf", file, os.ModePerm)
    } else {
        fmt.Println(err.Error())
    }
}

*/


package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/EluctariLLC/dropbox"
   // "github.com/EluctariLLC/dropbox/experiments/config"
    "github.com/davecgh/go-spew/spew"
)

var printf = fmt.Printf

func main() {
    dropbox.Debug = true

        // Available Root:
        // Expected 'root' to be 'dropbox', 'sandbox', or 'auto'
    client, err := dropbox.NewClient(dropbox.ClientOptions{
        Key:    "my_app_key",
        Secret: "my_app_secret",
        Token:  "my_acess_token-wcaqw",
        Uid:    "my_user_id",
        Root:   "dropbox",
    })

    failIf(err)

    { // demo GetAccountInfo
        println("calling client.GetAccountInfo")
        info, err := client.GetAccountInfo()
        failIf(err)
        spew.Dump(info)
        print("\n\n")
    }

    { // demo Mkdir, ReadDir and Delete
        const path = "hello-directory"

        println("calling client.Mkdir")
        stat, err := client.Mkdir(path)
        failIf(err)
        spew.Dump(stat)

        println("calling client.Mkdir (again)")
        stat, err = client.Mkdir(path + "/foobar")
        failIf(err)
        spew.Dump(stat)

        // println("calling client.Copy")
        // stat, err = client.Copy(path+"/foobar", path+"/binbar")
        // failIf(err)
        // spew.Dump(stat)

        // println("calling client.Stat")
        // stat, err = client.Stat(path, nil)
        // failIf(err)
        // spew.Dump(stat)

        // println("calling client.ReadDir")
        // entryNames, dirStat, entryStats, err := client.ReadDir(path, nil)
        // failIf(err)
        // spew.Dump(entryNames, dirStat, entryStats)

        // println("calling client.Delete")
        // stat, err = client.Delete(path)
        // failIf(err)
        // spew.Dump(stat)

        print("\n\n")
    }

    { // demo WriteFile, Stat, ReadFile and Delete
        const path = "hello-file"

        file, err := ioutil.TempFile(os.TempDir(), "dropbox-sdk-go")
        failIf(err)
        defer func(f *os.File) {
            f.Close()
            failIf(os.Remove(f.Name()))
        }(file)

        { // write to local tempfile;
            _, err = file.WriteString("Hello World!\n")
            failIf(err)

            _, err = file.Seek(0, 0)
            failIf(err)
            failIf(file.Sync())
        }

        println("calling client.WriteFile")
        stat, err := client.WriteFile(path, file, nil)
        failIf(err)
        spew.Dump(stat)

        println("calling client.Stat")
        stat, err = client.Stat(path, nil)
        failIf(err)
        spew.Dump(stat)

        println("calling client.ReadFile")
        var buffer bytes.Buffer
        stat, err = client.ReadFile(path, &buffer, nil)
        failIf(err)
        printf("File contents: %s", buffer.Bytes())
        spew.Dump(stat)

        // println("calling client.Delete")
        // stat, err = client.Delete(path)
        // failIf(err)
        // spew.Dump(stat)

        print("\n\n")
    }

    println("Client smoke test finished!")
}

func failIf(err error) {
    if err != nil {
        panic(err)
    }
}
