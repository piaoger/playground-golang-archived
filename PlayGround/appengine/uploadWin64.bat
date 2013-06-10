
set THIS_FOLDER=%~dp0
:: trim trailing backslash
if %THIS_FOLDER:~-1%==\ set THIS_FOLDER=%THIS_FOLDER:~0,-1%

set GOROOT=%THIS_FOLDER%\..\go\go-1.0.3\Win64
set APPENGINESDKROOT=%THIS_FOLDER%\..\GoAppEngineSDK\GoAppEngineSDK-1.80\Win64
set PYTHONROOT=S:\Workspace\Tools\Python\Win-2.7.3.2\App

set PATH=%PATH%;%GOROOT%\bin;%APPENGINESDKROOT%;%PYTHONROOT%

set GOPATH=%THIS_FOLDER%

python %APPENGINESDKROOT%\appcfg.py update solidmcp

cmd

