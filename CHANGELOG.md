# CHANGELOG

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

### Project releases

## 2.2.0 - May, 16th 2024 - Long fixes

### Added

-   New unit tests for GUI and settings
-   Optimize setup when launching application
-   Add screen in documentation (README)
-   Add new configuration to change year of tax calculation in GUI
-   Add an autoupdater
-   Add new command in console and GUI for tax history

### Fixed

-   Refreshing menu when changing languages
-   Remove definitely duplicate menu item 'Quit'

### Changed

-   Refactoring GUI code structure with MVC

## 2.1.0 - January, 15th 2024 - Small fixes

### Added

-   Deepsource analyzed and fix issues
-   Add new command `show_tax_tranche`
-   Add taxes scales for `2023` & `2024`

### Changed

-   Update documentation (README)

## 2.0.0 - August, 2nd 2022 - GUI

### Added

-   First version of GUI app
-   Add Dark and light theme
-   Add Languages in GUI for `en` and `fr`
-   Add `about` options in console
-   Add `settings` into setting file for `theme`, `language` and `currency`
-   Add [zap librairy](https://github.com/uber-go/zap) to handle logs with log rotation

### Changed

-   Rename `build.sh` to `package.sh` with all folowing (Readme, Makefile)
-   Reorganize `gui` package

### Removed

-   Remove Docker setup
    -   Files: `Dockerfile` `.dockerignore`
    -   Command in the `Makefile`

## 1.1.0 - July, 9th 2022 - 2022 update

### Added

-   Add tax metrics for 2022
-   Add icon for windows executable

### Changed

-   Update `Makefile` and `build.sh` script
-   Update Golang version from `v1.16.4` to `v1.18.3`

# Fixed

-   Fix Shares counter for `isolated parent`

### Removed

-   Remove `config.json` file

## 1.0.0 - July, 10th 2021 - Public version

### Featured

-   Create Docker images and features
-   Create Go documentation

### Added

-   Add tax details such as income and year of tax metrics

### Changed

-   Split Execution mode `GUI` & `Console`
-   Change `Parts` field into `User` struct to `Shares`

### Fixed

-   Fix method `GetShares` get shares of the user
-   Fix method to read data from console

## 0.0.9 - July, 4th 2021 - Calculate tax v3

### Featured

-   reverse tax calculator

### Added

-   Add command `show_tax_year_list`
-   Add command `show_tax_year_used`
-   Add command `select_tax_year`
-   Add tax metrics for 2019 and 2020

### Changed

-   Change `Percentage` field to `Rate` in `Tranche` struct
-   Change type of `Rate` float to string

## 0.0.8 - July, 3rd 2021 - Refactoring modules

### Added

-   Create `Makefile` with bunch of commands
-   Set an index parameter for each commands

### Changed

-   Simplify entrypoint `main.go`
-   Update `README` documentation

### Fixed

## 0.0.7 - June 29th 2021 - Restructure project

### Changed

-   Rename package `core` to `tax`
-   Change config structure and add tag fiels in config struct

### Removed

-   Remove `struct.go` file and add struct into package file

## 0.0.6 - June, 28th 2021 - Restructure project

### Added

-   `LICENSE.md` file (GPL-3.0 License)

### Changed

-   Reorganize folders
-   Change module's name
-   Update Readme

## 0.0.5 - June, 26th 2021 - table tax tranches

### Added

-   Create table structure to get tax tranches

## 0.0.4 - June, 25th 2021 - Calculate tax v2

### Added

-   New process to integrate couple, children to process part and including them to the tax process

## 0.0.3 - June, 23rd 2021 - Fix

### Added

-   Testing scripts for config and tax modules

### Fixed

-   Set handler if config doesn't exist

## 0.0.2 - June, 22nd 2021 - Calculate tax v1

### Added

-   Configuration management
-   Process to calculate tax from income

## 0.0.1 - June, 21st 2021 - Init project
