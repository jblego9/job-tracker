@echo off

cd "%~dp0..\frontend"
npx vite build --outDir "%~dp0..\build"
