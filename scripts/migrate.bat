@echo off
setlocal

REM Configurations
set "MIGRATION_DIR=.\migrations"
set "DATABASE_URL=postgres://admin:admin@localhost:5432/goamartha?sslmode=disable"

REM Command to execute
set "COMMAND=%1"

REM Function to show usage
:show_usage
    echo Usage: migrate.bat [COMMAND]
    echo Commands:
    echo   up         - Apply all up migrations
    echo   down       - Apply all down migrations
    echo   drop       - Drop everything
    echo   force VERSION - Set version
    echo   version    - Print current migration version
    echo   create NAME - Create a new migration
    echo   help       - Show this usage
    goto :eof

REM Function to run migrations up
:migrate_up
    migrate -path %MIGRATION_DIR% -database %DATABASE_URL% up
    goto :eof

REM Function to run migrations down
:migrate_down
    migrate -path %MIGRATION_DIR% -database %DATABASE_URL% down
    goto :eof

REM Function to drop everything
:migrate_drop
    migrate -path %MIGRATION_DIR% -database %DATABASE_URL% drop
    goto :eof

REM Function to force set version
:migrate_force
    migrate -path %MIGRATION_DIR% -database %DATABASE_URL% force %2
    goto :eof

REM Function to show current version
:migrate_version
    migrate -path %MIGRATION_DIR% -database %DATABASE_URL% version
    goto :eof

REM Function to create a new migration
:migrate_create
    migrate create -ext sql -dir %MIGRATION_DIR% -seq %2
    goto :eof

REM Check if migrate is installed
where migrate >nul 2>nul
if %errorlevel% neq 0 (
    echo Error: migrate is not installed.
    exit /b 1
)

REM Execute the specified command
if "%COMMAND%"=="up" goto migrate_up
if "%COMMAND%"=="down" goto migrate_down
if "%COMMAND%"=="drop" goto migrate_drop
if "%COMMAND%"=="force" (
    if "%2"=="" (
        echo Error: VERSION must be specified for force command.
        goto show_usage
    )
    goto migrate_force
)
if "%COMMAND%"=="version" goto migrate_version
if "%COMMAND%"=="create" (
    if "%2"=="" (
        echo Error: NAME must be specified for create command.
        goto show_usage
    )
    goto migrate_create
)
goto show_usage