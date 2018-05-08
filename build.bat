@echo off
rem 作者：Lotomer
rem 时间：2018-05-08
rem 最后更新时间：2018-05-08
set CURR_PATH=%~dp0
set OUTPUT_PATH=output
cd %CURR_PATH%
call :buildGO datastore linux amd64 "" ".\application\datastore"
call :buildGO datastore linux 386 "" ".\application\datastore"
call :buildGO datastore windows amd64 ".exe" ".\application\datastore"

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