
@REM Config file. For import in others scripts.
@REM CALL "%~dp0_config.cmd"

@ECHO OFF
CLS
ENDLOCAL

CALL "%~dp0heroku_info.cmd"

SET PATH_ROOT=%~dp0..\
@REM SET GOPATH=C:\Users\User\go
SET PAUSE_IN_END=0

CD /D "%PATH_ROOT%"

SETLOCAL ENABLEDELAYEDEXPANSION
ECHO Variables declared.

EXIT /B %ERRORLEVEL%
