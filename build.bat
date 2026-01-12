@echo off
echo Building go-android-bruteforce-pin...

set CGO_ENABLED=1
"%ProgramFiles%\Go\bin\go.exe" build -ldflags="-s -w" -o go-android-bruteforce-pin.exe cmd/sony-xperia-z1-compact/main.go

echo Done.
