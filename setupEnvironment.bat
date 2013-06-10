
set THIS_FOLDER_TAIL=%~dp0
set THIS_FOLDER=%THIS_FOLDER_TAIL:~0,-1%

set SMCP_ROOT=%THIS_FOLDER%

set SMCP_CPU_ARCH=x64
if %PROCESSOR_ARCHITECTURE%==x86 (
    set SMCP_CPU_ARCH=x86
) else (
    set SMCP_CPU_ARCH=x64
)

set SMCP_UNZIP_TOOL=%SMCP_ROOT%\Tools\unzip\unzip.exe

set TARGET_PATH=%SMCP_ROOT%\Components\go\go-1.1
if not exist %TARGET_PATH% md %TARGET_PATH%
if not exist %TARGET_PATH%\Win64  %SMCP_UNZIP_TOOL%  %SMCP_ROOT%\Components\Packages\go-1.1-win64.zip -d %TARGET_PATH%

set TARGET_PATH=%SMCP_ROOT%\Components\GoAppEngineSDK\GoAppEngineSDK-1.80
if not exist %TARGET_PATH% md %TARGET_PATH%
if not exist %TARGET_PATH%\Win64  %SMCP_UNZIP_TOOL%  %SMCP_ROOT%\Components\Packages\goapprnginesdk-1.80-win64 -d %TARGET_PATH%


