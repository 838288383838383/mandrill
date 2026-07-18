@echo off
echo 8b    d8    db    88b 88 8888b.  88""Yb 88 88     88
echo 88b  d88   dPYb   88Yb88  8I  Yb 88__dP 88 88     88
echo 88YbdP88  dP__Yb  88 Y88  8I  dY 88"Yb  88 88  .o 88  .o
echo 88 YY 88 dP""""Yb 88  Y8 8888Y"  88  Yb 88 88ood8 88ood8
echo.
echo   Package manager and git TUI — Windows installer
echo.

set INSTALL_DIR=%LOCALAPPDATA%\mandrill\bin
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"

echo Downloading mandrill-windows-amd64.exe...
curl -sL "https://github.com/838288383838383/mandrill/releases/latest/download/mandrill-windows-amd64.exe" -o "%INSTALL_DIR%\mandrill.exe"

echo Installed to %INSTALL_DIR%\mandrill.exe
echo.
echo Add %INSTALL_DIR% to your PATH:
echo   setx PATH "%PATH%;%INSTALL_DIR%"
echo.
echo Or add it manually in System Properties ^> Environment Variables
echo.
echo Done! Run 'mandrill --help' to get started.
