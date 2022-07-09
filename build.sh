#!/bin/bash

# ------------------------------------------------------------------
# [Title] : Archange
# [Description] : Script to create archive files for new project release
# [Version] : v1.0.0
# [Author] : Lucas Noga
# [Shell] : Bash v5.1.0
# [Usage] : ./build.sh 1.1.0
#           make build
# ------------------------------------------------------------------

EXECUTABLE=corpos-christie

# OS list
declare -a OS_LIST=("linux" "windows" "mac")

function main {

    # Clean old builds
    clean

    # if no version
    if [ -z $1 ]; then
        echo "Set an argument to the version. Ex: 0.0.2"
        exit
    fi

    # Creating new build
    version=$1

    cd build

    # Iterate the string array using for loop
    for os in ${OS_LIST[@]}; do
        build_os $os $version
    done

    # build_os windows v1
}

# Build executables for one os
# TODO faire un folder par os
function build_os {
    os=$1
    version=$2
    echo "Build ${EXECUTABLE} for ${os} in version: $version"

    # Move into right folder
    mkdir ${os}

    extension=""
    if [ $os == "windows" ]; then
        extension=".exe"
    fi

    # Handle syso file
    syso_file="corpos-christie.syso"
    cp ${syso_file} ${os}/${syso_file}

    # move executable to change path and name
    executable=${os}-${EXECUTABLE}${extension}
    executable_name="${EXECUTABLE}${extension}"
    echo "mv ${executable} ${os}/${executable_name}"
    mv ${executable} ${os}/${executable_name}

    cd ${os}

    zip_file="${EXECUTABLE}-${version}.zip"
    tar_file="${EXECUTABLE}-${version}.tar"
    targz_file="${EXECUTABLE}-${version}.tar.gz"

    echo "Zip Program into ${zip_file}"
    zip ${zip_file} ${executable_name} ${syso_file}

    echo "Tar Program ${tar_file}"
    tar -cf ${tar_file} ${executable_name} ${syso_file}

    echo "Tar.gz Program ${targz_file}"
    tar -zcf ${targz_file} ${executable_name} ${syso_file}

    # Return in build folder
    cd ..
}

# Clean old files for exe
function clean {
    current_directory=${PWD}
    cd build # Go into build directory

    declare -a FILES_EXT=(".tar" ".tar.gz" ".zip")
    for ext in ${FILES_EXT[@]}; do
        rm -f *-${EXECUTABLE}-*${ext}
    done

    cd ${current_directory} # return in previous directory
}

main $@
