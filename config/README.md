# config

This directory contains the config files used in the application or other tools that we might use e.g. linters.
Config files can be .yaml, .json, ... files.

We also declare a config package in this directory and implement the code for loading the config data into type-safe structs.

We should respect the order of configuration loading that are listed as below:

* explicit call to Set
* flag
* env
* config
* key/value store
* default

The package <https://github.com/spf13/viper> could be used.
