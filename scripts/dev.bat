@echo off

:: Assumes that the backend .env is configured
cd "%~dp0..\backend\cmd"
start cmd /k  "go run .\server"

:: Assumes that the vite.config.js server proxy target is configured
cd "%~dp0..\frontend"
start cmd /k "npm run dev"
