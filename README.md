# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**

# Snap collector plugin - facter
This plugin collects facts about nodes (systems) such as hardware details, network settings, OS type and version, IP addresses, MAC addresses, SSH keys and more from Facter and converts them into Snap metrics.

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
   * [System Requirements](#system-requirements)
   * [Installation](#installation)
   * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
   * [Collected Metrics](#collected-metrics)
   * [Examples](#examples)
   * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [facter 3.x](https://github.com/puppetlabs/facter/blob/master/README.md)
* [golang 1.5+](https://golang.org/dl/) (needed only for building)

NOTE: Facter 2.x may also work, but it's currently not part of our test matrix.

### Operating systems
All OSs currently supported by Snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-facter/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-facter
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-facter.git
```

Build the Snap facter plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap#getting-started)
* Load the plugin and create a task, see example in [Examples](#examples).

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [facter](http://docs.puppetlabs.com/facter/3.1/core_facts.html) 
* [Snap facter integration test](facter/facter_integration_test.go)
* [Snap facter unit test](facter/facter_test.go)
* [Snap facter examples](#Examples)

### Collected Metrics
List of metrics collected by this plugin can be found in [METRICS.md file](METRICS.md).

### Examples
Example of running Snap facter collector and writing data to file.

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `service snap-telemetry start`
* systemd: `systemctl start snap-telemetry`
* command line: `snapteld -l 1 -t 0 &`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-facter/latest/linux/x86_64/snap-plugin-collector-facter
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-facter
$ snaptel plugin load snap-plugin-publisher-file
```

See all available metrics:

```
$ snaptel metric list
```

Download an [example task file](examples/tasks/task-facter.json) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-facter/master/examples/tasks/task-facter.json
$ snaptel task create -t task-facter.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)

```
2015-12-02 08:49:33.679791899 -0500 EST|[intel facter architecture]|amd64|gklab-044-107
2015-12-02 08:49:33.679792606 -0500 EST|[intel facter bios_release_date]|07/14/2011|gklab-044-107
2015-12-02 08:49:33.679797562 -0500 EST|[intel facter bios_vendor]|Hewlett-Packard|gklab-044-107
2015-12-02 08:49:33.679797789 -0500 EST|[intel facter bios_version]|786H1 v01.13|gklab-044-107
2015-12-02 08:49:34.680113961 -0500 EST|[intel facter architecture]|amd64|gklab-044-107
2015-12-02 08:49:34.680114801 -0500 EST|[intel facter bios_release_date]|07/14/2011|gklab-044-107
2015-12-02 08:49:34.680115034 -0500 EST|[intel facter bios_vendor]|Hewlett-Packard|gklab-044-107
2015-12-02 08:49:34.680115224 -0500 EST|[intel facter bios_version]|786H1 v01.13|gklab-044-107
2015-12-02 08:49:35.680925816 -0500 EST|[intel facter architecture]|amd64|gklab-044-107
2015-12-02 08:49:35.680926797 -0500 EST|[intel facter bios_release_date]|07/14/2011|gklab-044-107
```

Stop task:
```
$ snaptel task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-facter/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-facter/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@ppalucki](https://github.com/ppalucki)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
