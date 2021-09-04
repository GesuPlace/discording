
@ECHO OFF
CLS
SETLOCAL ENABLEDELAYEDEXPANSION

SET PATH_ROOT=%~dp0..\
@REM SET GOPATH=C:\Users\User\go
CD /D "%PATH_ROOT%"

ECHO This script will build the project.

go.exe build

ECHO Done working.

PAUSE
ENDLOCAL
EXIT /B %ERRORLEVEL%
