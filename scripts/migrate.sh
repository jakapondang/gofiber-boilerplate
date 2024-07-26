#!/bin/bash

# Configurations
MIGRATION_DIR=./migrations
DATABASE_URL=postgres://admin:admin@localhost:5432/goamartha?sslmode=disable

# Command to execute
COMMAND=$1

# Function to show usage
show_usage() {
    echo "Usage: migrate.sh [COMMAND]"
    echo "Commands:"
    echo "  up         - Apply all up migrations"
    echo "  down       - Apply all down migrations"
    echo "  drop       - Drop everything"
    echo "  force VERSION - Set version"
    echo "  version    - Print current migration version"
    echo "  create NAME - Create a new migration"
    echo "  help       - Show this usage"
}

# Function to run migrations up
migrate_up() {
    migrate -path ${MIGRATION_DIR} -database ${DATABASE_URL} up
}

# Function to run migrations down
migrate_down() {
    migrate -path ${MIGRATION_DIR} -database ${DATABASE_URL} down
}

# Function to drop everything
migrate_drop() {
    migrate -path ${MIGRATION_DIR} -database ${DATABASE_URL} drop
}

# Function to force set version
migrate_force() {
    migrate -path ${MIGRATION_DIR} -database ${DATABASE_URL} force $1
}

# Function to show current version
migrate_version() {
    migrate -path ${MIGRATION_DIR} -database ${DATABASE_URL} version
}

# Function to create a new migration
migrate_create() {
    migrate create -ext sql -dir ${MIGRATION_DIR} -seq $1
}

# Check if migrate is installed
if ! [ -x "$(command -v migrate)" ]; then
    echo 'Error: migrate is not installed.' >&2
    exit 1
fi

# Execute the specified command
case ${COMMAND} in
    up)
        migrate_up
        ;;
    down)
        migrate_down
        ;;
    drop)
        migrate_drop
        ;;
    force)
        if [ -z "$2" ]; then
            echo "Error: VERSION must be specified for force command."
            show_usage
            exit 1
        fi
        migrate_force $2
        ;;
    version)
        migrate_version
        ;;
    create)
        if [ -z "$2" ]; then
            echo "Error: NAME must be specified for create command."
            show_usage
            exit 1
        fi
        migrate_create $2
        ;;
    help|*)
        show_usage
        ;;
esac