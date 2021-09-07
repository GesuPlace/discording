
CALL "%~dp0_config.cmd"

ECHO This script will install the project dependences (for the first run).

go.exe version
@REM go get -u github.com/kardianos/govendor
@REM govendor list
@REM govendor remove +u
go.exe mod init
go.exe mod vendor

ECHO Done working.

IF "%PAUSE_IN_END%" == "1" PAUSE
ENDLOCAL
EXIT /B %ERRORLEVEL%
