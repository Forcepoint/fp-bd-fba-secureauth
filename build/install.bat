:: Change to directory of SecureAuth Log Reader
cd %~dp0%

:: Take in required values
set /P password="Please enter your administrator password: "
for /F "tokens=*" %%g in ('whoami') do (SET user=%%g)
cls

:: Set required variables
set zipfile=%~dp0%resources/nssm.zip
set unzipfile=%~dp0%resources/unzip.exe
set runfile=%~dp0%LOGREADER.exe

:: Download required files
mkdir resources
powershell "Invoke-WebRequest -Outfile %zipfile% 'http://nssm.cc/release/nssm-2.24.zip'"
powershell "Invoke-WebRequest -Outfile %unzipfile% 'http://stahlworks.com/dev/unzip.exe'"

:: Decompress required files
cd resources
%unzipfile% %zipfile%

:: Install service
.\nssm-2.24\win64\nssm.exe install SecureAuthLogReader %runfile%
.\nssm-2.24\win64\nssm.exe set SecureAuthLogReader AppDirectory %~dp0%
.\nssm-2.24\win64\nssm.exe set SecureAuthLogReader ObjectName %user% %password%
.\nssm-2.24\win64\nssm.exe start SecureAuthLogReader

:: Clean up files
cd ..
rmdir /s /y %~dp0%resources

PAUSE