
set THIS_FOLDER=%~dp0
:: trim trailing backslash
if %THIS_FOLDER:~-1%==\ set THIS_FOLDER=%THIS_FOLDER:~0,-1%

set GOROOT=%THIS_FOLDER%\..\..\Components\go\go-1.1\Win64
set PATH=%PATH%;%GOROOT%\bin

set GOPATH=%THIS_FOLDER%

go build readjson.go

cmd

