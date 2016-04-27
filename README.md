# snap collector plugin - facter
This plugin collects facts about nodes (systems) such as hardware details, network settings, OS type and version, IP addresses, MAC addresses, SSH keys and more from Facter and converts them into snap metrics. 

It's used in the [snap framework](http://github.com:intelsdi-x/snap).

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
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements
* [facter](https://github.com/puppetlabs/facter/blob/master/README.md) 
* [golang 1.5+](https://golang.org/dl/) (needed only for building)

### Operating systems
All OSs currently supported by snap:
* Linux/amd64
* Darwin/amd64

### Installation
#### Download facter plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-facter
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-facter.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [facter](http://docs.puppetlabs.com/facter/3.1/core_facts.html) 
* [snap facter integration test](https://github.com/intelsdi-x/snap-plugin-collector-facter/blob/master/facter/facter_integration_test.go)
* [snap facter unit test](https://github.com/intelsdi-x/snap-plugin-collector-facter/blob/master/facter/facter_test.go)
* [snap facter examples](#Examples)

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Description
----------|-----------------------
/intel/facter/architecture |
/intel/facter/bios_release_date |
/intel/facter/bios_vendor |
/intel/facter/bios_version |
/intel/facter/blockdevices |
/intel/facter/blockdevice_sda_model |
/intel/facter/blockdevice_sda_size |
/intel/facter/blockdevice_sda_vendor |
/intel/facter/blockdevice_sdb_model |
/intel/facter/blockdevice_sdb_size |
/intel/facter/blockdevice_sdb_vendor |
/intel/facter/boardmanufacturer |
/intel/facter/boardproductname |
/intel/facter/boardserialnumber |
/intel/facter/dhcp_servers |
/intel/facter/domain |
/intel/facter/facterversion |
/intel/facter/filesystems |
/intel/facter/fqdn |
/intel/facter/gid |
/intel/facter/hardwareisa |
/intel/facter/hardwaremodel |
/intel/facter/hostname |
/intel/facter/id |
/intel/facter/interfaces |
/intel/facter/ipaddress |
/intel/facter/ipaddress_[INTERFACE] |
/intel/facter/is_virtual |
/intel/facter/kernel |
/intel/facter/kernelmajversion |
/intel/facter/kernelrelease |
/intel/facter/kernelversion |
/intel/facter/lsbdistcodename |
/intel/facter/lsbdistdescription |
/intel/facter/lsbdistid |
/intel/facter/lsbdistrelease |
/intel/facter/lsbmajdistrelease |
/intel/facter/lsbminordistrelease |
/intel/facter/lsbrelease |
/intel/facter/macaddress |
/intel/facter/macaddress_[INTERFACE] |
/intel/facter/manufacturer |
/intel/facter/memoryfree |
/intel/facter/memoryfree_mb |
/intel/facter/memorysize |
/intel/facter/memorysize_mb |
/intel/facter/mtu_[INTERFACE] |
/intel/facter/mtu_lo |
/intel/facter/netmask |
/intel/facter/netmask_[INTERFACE] |
/intel/facter/network_[INTERFACE] |
/intel/facter/operatingsystem |
/intel/facter/operatingsystemmajrelease |
/intel/facter/operatingsystemrelease |
/intel/facter/os |
/intel/facter/osfamily |
/intel/facter/partitions |
/intel/facter/path |
/intel/facter/physicalprocessorcount |
/intel/facter/processor[ID] |
/intel/facter/processorcount |
/intel/facter/processors |
/intel/facter/productname |
/intel/facter/ps |
/intel/facter/rubyplatform |
/intel/facter/rubysitedir |
/intel/facter/rubyversion |
/intel/facter/selinux |
/intel/facter/selinux_config_mode |
/intel/facter/selinux_config_policy |
/intel/facter/selinux_current_mode |
/intel/facter/selinux_enforced |
/intel/facter/selinux_policyversion |
/intel/facter/serialnumber |
/intel/facter/ssh[KEY] |
/intel/facter/swapfree |
/intel/facter/swapfree_mb |
/intel/facter/swapsize |
/intel/facter/swapsize_mb |
/intel/facter/system_uptime |
/intel/facter/timezone |
/intel/facter/type |
/intel/facter/uniqueid |
/intel/facter/uptime |
/intel/facter/uptime_days |
/intel/facter/uptime_hours |
/intel/facter/uptime_seconds |
/intel/facter/uuid |
/intel/facter/virtual |

### Examples
Example running facter, passthru processor, and writing data to a file.

This is done from the snap directory.

In one terminal window, open the snap daemon (in this case with logging set to 1 and trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0
```

In another terminal window:
Load facter plugin
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-collector-facter
```
See available metrics for your system
```
$ $SNAP_PATH/bin/snapctl metric list
```

Create a task manifest file (e.g. `facter-file.json`):    
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "1s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/facter/architecture": {},
                "/intel/facter/bios_release_date": {},
                "/intel/facter/bios_vendor": {},
                "/intel/facter/bios_version": {}
            },
            "config": {
                "/intel/mock": {
                    "password": "secret",
                    "user": "root"
                }
            },
            "process": [
                {
                    "plugin_name": "passthru",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "file",
                            "config": {
                                "file": "/tmp/published_facter"
                            }
                        }
                    ],
                    "config": null
                }
            ],
            "publish": null
        }
    }
}
```

Load passthru plugin for processing:
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-processor-passthru
Plugin loaded
Name: passthru
Version: 1
Type: processor
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:44:03 PST
```

Load file plugin for publishing:
```
$ $SNAP_PATH/bin/snapctl plugin load build/plugin/snap-publisher-file
Plugin loaded
Name: file
Version: 3
Type: publisher
Signed: false
Loaded Time: Fri, 20 Nov 2015 11:45:03 PST
```

Create task:
```
$ $SNAP_PATH/bin/snapctl task create -t examples/tasks/facter-file.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See file output (this is just part of the file):
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
$ $SNAP_PATH/bin/snapctl task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-facter/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-facter/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@ppalucki](https://github.com/ppalucki)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
