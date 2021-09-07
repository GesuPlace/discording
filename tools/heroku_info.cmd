
@REM Heroku info pull script. For import in config only.
@REM CALL "%~dp0heroku_info.cmd"

ECHO Printing info about Heroku app.

heroku addons --all
@REM heroku apps:info

EXIT /B %ERRORLEVEL%
