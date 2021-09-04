
@ECHO OFF
CLS
SETLOCAL ENABLEDELAYEDEXPANSION

SET PATH_ROOT=%~dp0..\
@REM SET GOPATH=C:\Users\User\go
CD /D "%PATH_ROOT%"

ECHO This script will install the project dependences.

go.exe version
@REM go get -u github.com/kardianos/govendor
go.exe mod init
go.exe mod vendor

ECHO Done working.

PAUSE
ENDLOCAL
EXIT /B %ERRORLEVEL%
