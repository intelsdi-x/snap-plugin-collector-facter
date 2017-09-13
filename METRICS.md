## Collected metrics

This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-----------------------
/intel/facter/architecture | string | Operating system’s hardware architecture
/intel/facter/bios_release_date | string | Release date of the system BIOS
/intel/facter/bios_vendor | string | Vendor of the system BIOS
/intel/facter/bios_version | string | Version of the system BIOS
/intel/facter/blockdevices | string | Comma-separated list of block devices
/intel/facter/blockdevice_[DEVICE]_model | string | Model name of block devices attached to the system
/intel/facter/blockdevice_[DEVICE]_size | string | Size of a block device in bytes
/intel/facter/blockdevice_[DEVICE]_vendor | string | Vendor name of block devices attached to the system
/intel/facter/boardmanufacturer | string | System board manufacturer
/intel/facter/boardproductname | string | System board product name
/intel/facter/boardserialnumber | string | System board serial number
/intel/facter/dhcp_servers | string | DHCP servers for the system
/intel/facter/domain | string | Network domain of the system
/intel/facter/facterversion | string | Version of facter
/intel/facter/filesystems | string | Usable file systems for block or disk devices
/intel/facter/fqdn | string | Fully-qualified domain name of the system
/intel/facter/gid | string | Group identifier (GID) of the user running Snap
/intel/facter/hardwareisa | string | Hardware instruction set architecture (ISA)
/intel/facter/hardwaremodel | string | Operating system’s hardware model
/intel/facter/hostname | string | Host name of the system
/intel/facter/id | string | User identifier (UID) of the user running Snap
/intel/facter/interfaces | string | Comma-separated list of network interface names
/intel/facter/ipaddress | string | IPv4 address for the default network interface
/intel/facter/ipaddress_[INTERFACE] | string | IPv4 address for a network interface
/intel/facter/is_virtual | string | Whether or not the host is a virtual machine
/intel/facter/kernel | string | Kernel’s name
/intel/facter/kernelmajversion | string | Kernel’s major version
/intel/facter/kernelrelease | string | Kernel’s release
/intel/facter/kernelversion | string | Kernel’s version
/intel/facter/lsbdistcodename | string | Linux Standard Base (LSB) distribution code name
/intel/facter/lsbdistdescription | string | Linux Standard Base (LSB) distribution description
/intel/facter/lsbdistid | string | Linux Standard Base (LSB) distribution identifier
/intel/facter/lsbdistrelease | string | Linux Standard Base (LSB) distribution release
/intel/facter/lsbmajdistrelease | string | Linux Standard Base (LSB) major distribution release
/intel/facter/lsbminordistrelease | string | Linux Standard Base (LSB) minor distribution release
/intel/facter/lsbrelease | string | Linux Standard Base (LSB) release
/intel/facter/macaddress | string | MAC address for the default network interface
/intel/facter/macaddress_[INTERFACE] | string | MAC address for a network interface
/intel/facter/manufacturer | string | System manufacturer
/intel/facter/memoryfree | string | Display size of the free system memory (e.g. “1 GiB”)
/intel/facter/memoryfree_mb | string | Size of the free system memory, in mebibytes
/intel/facter/memorysize | string | display size of the total system memory (e.g. “1 GiB”)
/intel/facter/memorysize_mb | string | Size of the total system memory, in mebibytes
/intel/facter/mtu_[INTERFACE] | string | Maximum Transmission Unit (MTU) for a network interface
/intel/facter/netmask | string | IPv4 netmask for the default network interface
/intel/facter/netmask_[INTERFACE] | string | IPv4 netmask for a network interface
/intel/facter/network | string | IPv4 network for the default network interface
/intel/facter/network_[INTERFACE] | string | IPv4 network for a network interface
/intel/facter/operatingsystem | string | Name of the operating system
/intel/facter/operatingsystemmajrelease | string | Major release of the operating system
/intel/facter/operatingsystemrelease | string | Release of the operating system
/intel/facter/osfamily | string | Family of the operating system
/intel/facter/partitions | string | Disk partitions of the system
/intel/facter/path | string | PATH environment variable
/intel/facter/physicalprocessorcount | string | Count of physical processors
/intel/facter/processor[N] | string | Model string of processor N
/intel/facter/processorcount | string | Count of logical processors
/intel/facter/productname | string | System product name
/intel/facter/rubyplatform | string | Platform Ruby was built for
/intel/facter/rubysitedir | string | Path to Ruby’s site library directory
/intel/facter/rubyversion | string | Version of Ruby
/intel/facter/selinux | string | Whether Security-Enhanced Linux (SELinux) is enabled
/intel/facter/selinux_config_mode | string | Configured Security-Enhanced Linux (SELinux) mode
/intel/facter/selinux_config_policy | string | Configured Security-Enhanced Linux (SELinux) policy
/intel/facter/selinux_current_mode | string | Current Security-Enhanced Linux (SELinux) mode
/intel/facter/selinux_enforced | string | Whether Security-Enhanced Linux (SELinux) is enforced
/intel/facter/selinux_policyversion | string | Security-Enhanced Linux (SELinux) policy version
/intel/facter/serialnumber | string | System product serial number
/intel/facter/ssh[ALGORITHM]key | string | SSH public key for the algorithm
/intel/facter/swapfree | string | Display size of the free swap memory (e.g. “1 GiB”)
/intel/facter/swapfree_mb | string | Size of the free swap memory, in mebibytes
/intel/facter/swapsize | string | Display size of the total swap memory (e.g. “1 GiB”)
/intel/facter/swapsize_mb | string | Size of the total swap memory, in mebibytes
/intel/facter/timezone | string | System timezone
/intel/facter/uptime | string | System uptime
/intel/facter/uptime_days | string | System uptime days
/intel/facter/uptime_hours | string | System uptime hours
/intel/facter/uptime_seconds | string | System uptime seconds
/intel/facter/uuid | string | System product unique identifier
/intel/facter/virtual | string | Hypervisor name for virtual machines or “physical” for physical machines