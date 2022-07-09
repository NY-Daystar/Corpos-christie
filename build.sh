#!/bin/bash

# Script to create archive files for new project release
# First you have to $ make build

Program_Name=corpos-christie

# OS list
declare -a OS_LIST=("linux" "windows" "mac")

# setup for all os
all_build() {
    Version=$1
    echo "Build ${Program_Name} Version: ${Version}"

    cd build

    # Iterate the string array using for loop
    for os in ${OS_LIST[@]}; do
        build_os $os
    done
}

# Build executables for one os
build_os() {
    os=$1
    echo "Build ${Program_Name} for ${os}"

    extension=""
    if [ $os == "windows" ]; then
        extension=".exe"
    fi

    core_name=${os}-${Program_Name}
    program_file=${os}-${Program_Name}${extension}
    exe_name=${Program_Name}${extension}

    # echo DEBUG: cp ${program_file} ${exe_name}
    cp ${program_file} ${exe_name}

    zip_file=${core_name}-${Version}.zip
    tar_file=${core_name}-${Version}.tar
    targz_file=${core_name}-${Version}.tar.gz

    echo "Zip Program ${Program_Name}"
    zip ${zip_file} ${exe_name}

    echo "Tar Program ${Program_Name}"
    tar -cf ${tar_file} ${exe_name}

    echo "Tar.gz Program ${Program_Name}"
    tar -zcf ${targz_file} ${exe_name}
}

# Clean old files for exe
clean() {
    current_directory=${PWD}
    cd build
    rm -rf *-${Program_Name}-*.tar
    rm -rf *-${Program_Name}-*.tar.gz
    rm -rf *-${Program_Name}-*.zip
    rm ${Program_Name}.exe
    rm ${Program_Name}
    cd ${current_directory}
}

main() {

    # Clean old builds
    clean

    # if no version
    if [ -z $1 ]; then
        echo "Set an argument to the version. Ex: 0.0.2"
        exit
    fi

    # Creating new build
    all_build $@
}

main $@
