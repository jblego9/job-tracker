@echo off

set "BUILD_DIR=%~dp0..\build"
set "DIST_DIR=%~dp0..\dist"

echo Starting build...

if exist "%BUILD_DIR%" rd /s /q "%BUILD_DIR%"
mkdir "%BUILD_DIR%"

echo Copying backend source into build dir...
robocopy "%~dp0..\backend" "%BUILD_DIR%" /e >nul
if %errorlevel% leq 7 set errorlevel=0

echo Building frontend...
cd "%~dp0..\frontend"
call npx vite build --outDir "%BUILD_DIR%\internal\frontend-build" --logLevel error

if exist "%DIST_DIR%" rd /s /q "%DIST_DIR%"
mkdir "%DIST_DIR%"

echo Building backend...
cd "%BUILD_DIR%"
go build -o "%DIST_DIR%\server.exe" .\cmd\server\

echo Creating default config files
copy NUL "%DIST_DIR%\data.db"
(
    echo PORT=8080
    echo DB_PATH=./data.db
) > "%DIST_DIR%\.env"

echo Build finished
