# Gin API template

## Project Layout

See https://github.com/golang-standards/project-layout/ for this project layout

- Set routing in `pkg/server/routes.go`
- Create application models, controllers, etc. in `internal/app/`

## Initialize

1. モジュール名を変更する
    - `scripts/alias.sh`
    - `go.mod`
    - import
        - `cmd/main.go`
        - `pkg/database/database.go`
        - `pkg/server/server.go`
        - `pkg/server/routes.go`
2. alias設定する
    - `$ source scripts/alias.sh`
    - Windows
        - `alias.bat` は動作未検証
        - WSL2の使用を推奨
3. ビルドする
    - `$ build`
