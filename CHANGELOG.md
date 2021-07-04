# ChangeLog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

### Project release
## [0.0.9] - 2021-07-04 - Calculate tax v3
> Add: reverse tax calculator  
> Add: Command `show_tax_year_list`  
> Add: Command `show_tax_year_used`  
> Add: Command `select_tax_year`  
> Add: Tax metrics for 2019 and 2020  
> Update: Change `Percentage` field to `Rate` in `Tranche` struct  
> Update: Change type of `Rate` float to string  
## [0.0.8] - 2021-07-03 - Refactoring modules
> Add: `Makefile` with bunch commands  
> Add: index parameter for each commands  
> Update: Simplify entrypoint `main.go`  
> Update: `README` documentation  
> Fix: Modules `core`, `tax`, `user`, `config`, `utils`, `colors`  
## [0.0.7] - 2021-06-29 - Restructure project
> Fix: Rename package `core` to `tax`  
> Fix: Change config structure and add tag fiels in config struct  
> Del: Remove struct.go file and add struct into package file  
## [0.0.6] - 2021-06-28 - Restructure project
> Add: LICENSE.md file (GPL-3.0 License)  
> Update: Change module name  
> Fix: Restructure folders  
> Fix: Update Readme  
## [0.0.5] - 2021-06-26 - table tax tranches
> Add: table to get tax tranches  
## [0.0.4] - 2021-06-25 - Calculate tax v2
> Add: new process to integrate couple, children to process part and including them to the tax process  
## [0.0.3] - 2021-06-23 - Fix
> Add: Testing script for config and tax modules  
> Fix: Config doesn't exist  
## [0.0.2] - 2021-06-22 - Calculate tax v1
> Add: Configuration management
> Add: Process to calculate tax from income 
## [0.0.1] - 2021-06-21 - Init project
