:: Build the executable file
go build -o ./dist/LOGREADER.exe ../main.go

:: Copy the required config files
mkdir ../config/realms
xcopy /s /e /i /y "../config" "./dist/config"
echo f | xcopy /f /y "./install.bat" "./dist/install.bat"

:: Zip files
powershell "Compress-Archive ./dist/* ./LOGREADER.zip -Update"
rd /s /q "./dist"

:: Move zip into dist folder
mkdir dist
move "LOGREADER.zip" "./dist"