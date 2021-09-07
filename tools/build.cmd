
CALL "%~dp0_config.cmd"

ECHO This script will build the project.

go.exe build

ECHO Done working.

IF "%PAUSE_IN_END%" == "1" PAUSE
ENDLOCAL
EXIT /B %ERRORLEVEL%
