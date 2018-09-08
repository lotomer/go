@echo off
rem 作者：Lotomer
rem 时间：2018-05-08
rem 最后更新时间：2018-05-11
set CURR_PATH=%~dp0
set OUTPUT_PATH=output
cd %CURR_PATH%

set app_name=datastore
call :buildGO %app_name% linux amd64 "" ".\application\%app_name%"
call :buildGO %app_name% linux 386 "" ".\application\%app_name%"
call :buildGO %app_name% linux arm64 "" ".\application\%app_name%"
call :buildGO %app_name% windows amd64 ".exe" ".\application\%app_name%"

set app_name=crawler
call :buildGO %app_name% linux amd64 "" ".\application\%app_name%"
call :buildGO %app_name% linux 386 "" ".\application\%app_name%"
call :buildGO %app_name% linux arm64 "" ".\application\%app_name%"
call :buildGO %app_name% windows amd64 ".exe" ".\application\%app_name%"


goto :exit

:buildGO
set APPNAME=%~1
set GOOS=%~2
set GOARCH=%~3
set GOEXE=%~4
set SRCPATH=%~5
echo go build -o %OUTPUT_PATH%\%APPNAME%-%GOOS%-%GOARCH%%GOEXE% %SRCPATH%
go build -o %OUTPUT_PATH%\%APPNAME%-%GOOS%-%GOARCH%%GOEXE% %SRCPATH%
goto :eof

:exit
mkdir %OUTPUT_PATH%\etc\
copy /y etc\* %OUTPUT_PATH%\etc\ 
pause