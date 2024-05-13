#!/bin/bash

# ------------------------------------------------------------------
# [Title] : Archange
# [Description] : Script to create archive files for new project release
# [Version] : v1.1.0
# [Author] : Lucas Noga
# [Shell] : Bash v5.1.0
# [Usage] : ./package.sh 1.1.0
#           make build
# ------------------------------------------------------------------

APP=corpos-christie
DEFAULT_VERSION=2.2.0

###
# Main body of script starts here
# $1 : Version of the app
###
function main {
    clean # Clean old builds
    version=$1

    # if no version
    if [ -z $version ]; then
        log_color "Set default version: ${DEFAULT_VERSION}" "yellow"
        version=${DEFAULT_VERSION}
    fi

    cd build

    # Creating new builds
    declare -a OS_LIST=("linux" "windows" "mac")
    for os in ${OS_LIST[@]}; do
        package_app $os $version
    done
}

###
# Package Apps for specific os
# $1 : [string] os built (ex: linux, windows, mac)
# $2 : [string] version of the app
###
function package_app {
    os=$1
    version=$2
    log_color "Build ${APP} for ${os} in version: $version" "yellow"

    # Move into right folder
    mkdir -p ${os}

    extension=""
    if [ $os == "windows" ]; then
        extension=".exe"
    fi

    # Copy resources folder into os folder
    resources="resources"
    cp -r ${resources} ${os}/${resources}
    log_color "Copy resources folder:  cp -r ${resources} ${os}/${resources}" "yellow"

    # Copy app and change its name into os folder
    app_old=${os}-${APP}${extension}
    app="${APP}${extension}"
    cp ${app_old} ${os}/${app}
    log_color "cp ${app_old} ${os}/${app}" "yellow"

    log_color "Go into os folder: $(pwd)/${os}" "yellow"
    cd ${os}

    archive_file=${os}-${APP}-${version}
    declare -a ARCHIVE_LIST=("zip" "tar" "tar.gz")

    for arch in ${ARCHIVE_LIST[@]}; do
        archive $arch $archive_file $resources
    done

    # Return in build folder
    cd ..
    log_color "Return to build folder: $(pwd)" "yellow"
}

###
# Archive program depend of archive extension (zip, tar)
# $1 : [string] archive extension (ex: zip, tar)
# $2 : [string] archive name (ex: linux-corpos-christie-1.0.0)
# $3 : [string] path of the resources folder
###
function archive {
    arch_ext=$1
    arch_name=$2
    resources=$3

    arch_file=${arch_name}.${arch_ext}

    log_color "$(capitalize $arch_ext) Program into ${arch_file}" "green"
    arch_cmd=""
    if [ $arch_ext = "zip" ]; then
        arch_cmd="zip -r ${arch_file} ${app} ${resources}"
        # zip -r ${arch_file} ${app} ${resources}
    elif [ $arch_ext = "tar" ]; then
        arch_cmd="tar -cf ${arch_file} ${app} ${resources}"
        # tar -cf ${arch_file} ${app} ${resources}
    elif [ $arch_ext = "tar.gz" ]; then
        arch_cmd="tar -zcf ${arch_file} ${app} ${resources}"
        # tar -zcf ${arch_file} ${app} ${resources}
    fi

    log_color "Archiving ${arch_cmd}" "green"
    # Execute command
    eval $arch_cmd
}

###
# Clean old files for exe
###
function clean {
    log_color "Cleaning old build..." "blue"

    current_directory=${PWD}
    cd build # Go into build directory

    declare -a FILES_EXT=(".tar" ".tar.gz" ".zip")
    for ext in ${FILES_EXT[@]}; do
        rm -f *-${EXECUTABLE}-*${ext}
    done

    log_color "Old build cleaned" "blue"

    cd ${current_directory} # return in previous directory
}

################################################################### Logging functions ###################################################################

###
# Capitalize string
###
function capitalize {
    v=$1
    echo "${v^}"
}

###
# Simple log with date function to support color
###
function log {
    echo -e $(date '+%Y-%m-%d %H:%M:%S') : $@
}

typeset -A COLORS=(
    [default]='\033[0;39m'
    [black]='\033[0;30m'
    [red]='\033[0;31m'
    [green]='\033[0;32m'
    [yellow]='\033[0;33m'
    [blue]='\033[0;34m'
    [magenta]='\033[0;35m'
    [cyan]='\033[0;36m'
    [light_gray]='\033[0;37m'
    [light_grey]='\033[0;37m'
    [dark_gray]='\033[0;90m'
    [dark_grey]='\033[0;90m'
    [light_red]='\033[0;91m'
    [light_green]='\033[0;92m'
    [light_yellow]='\033[0;93m'
    [light_blue]='\033[0;94m'
    [light_magenta]='\033[0;95m'
    [light_cyan]='\033[0;96m'
    [nc]='\033[0m' # No Color
)

###
# Log the message in specific color
###
function log_color {
    message=$1
    color=$2
    log ${COLORS[$color]}$message${COLORS[nc]}
}

main $@
