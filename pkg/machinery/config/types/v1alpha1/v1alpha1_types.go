// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

/*
Package v1alpha1 configuration file contains all the options available for configuring a machine.

To generate a set of basic configuration files, run:
```bash
talosctl gen config --version v1alpha1 <cluster name> <cluster endpoint>
````

This will generate a machine config for each node type, and a talosconfig for the CLI.
*/
package v1alpha1

//go:generate docgen ./v1alpha1_types.go ./v1alpha1_types_doc.go Configuration

import (
	"net/url"
	"os"
	"time"

	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/talos-systems/crypto/x509"

	"github.com/talos-systems/talos/pkg/machinery/config"
)

func init() {
	config.Register("v1alpha1", func(version string) (target interface{}) {
		target = &Config{}

		return target
	})
}

var (

	// Examples section.

	machineConfigRegistriesExample = &RegistriesConfig{
		RegistryMirrors: map[string]*RegistryMirrorConfig{
			"docker.io": {
				MirrorEndpoints: []string{"https://registry.local"},
			},
			"ghcr.io": {
				MirrorEndpoints: []string{"https://registry.insecure", "https://ghcr.io/v2/"},
			},
		},
		RegistryConfig: map[string]*RegistryConfig{
			"registry.local": {
				RegistryTLS: &RegistryTLSConfig{
					TLSClientIdentity: pemEncodedCertificateExample,
				},
				RegistryAuth: &RegistryAuthConfig{
					RegistryUsername: "username",
					RegistryPassword: "password",
				},
			},
			"registry.insecure": {
				RegistryTLS: &RegistryTLSConfig{
					TLSInsecureSkipVerify: true,
				},
			},
		},
	}

	pemEncodedCertificateExample *x509.PEMEncodedCertificateAndKey = &x509.PEMEncodedCertificateAndKey{
		Crt: []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJIekNCMHF..."),
		Key: []byte("LS0tLS1CRUdJTiBFRDI1NTE5IFBSSVZBVEUgS0VZLS0tLS0KTUM..."),
	}

	machineKubeletExample = &KubeletConfig{
		KubeletImage: (&KubeletConfig{}).Image(),
		KubeletExtraArgs: map[string]string{
			"key": "value",
		},
	}

	kubeletImageExample = (&KubeletConfig{}).Image()

	machineNetworkConfigExample = &NetworkConfig{
		NetworkHostname: "worker-1",
		NetworkInterfaces: []*Device{
			{},
		},
		NameServers: []string{"9.8.7.6", "8.7.6.5"},
	}

	machineDisksExample []*MachineDisk = []*MachineDisk{
		{
			DeviceName: "/dev/sdb",
			DiskPartitions: []*DiskPartition{
				{
					DiskMountPoint: "/var/mnt/extra",
					DiskSize:       100000000,
				},
			},
		},
	}

	machineInstallExample = &InstallConfig{
		InstallDisk:            "/dev/sda",
		InstallExtraKernelArgs: []string{"option=value"},
		InstallImage:           "ghcr.io/talos-systems/installer:latest",
		InstallBootloader:      true,
		InstallWipe:            false,
	}

	machineFilesExample = []*MachineFile{
		{
			FileContent:     "...",
			FilePermissions: 0o666,
			FilePath:        "/tmp/file.txt",
			FileOp:          "append",
		},
	}

	machineEnvExamples = []Env{
		{
			"GRPC_GO_LOG_VERBOSITY_LEVEL": "99",
			"GRPC_GO_LOG_SEVERITY_LEVEL":  "info",
			"https_proxy":                 "http://SERVER:PORT/",
		},
		{
			"GRPC_GO_LOG_SEVERITY_LEVEL": "error",
			"https_proxy":                "https://USERNAME:PASSWORD@SERVER:PORT/",
		},
		{
			"https_proxy": "http://DOMAIN\\USERNAME:PASSWORD@SERVER:PORT/",
		},
	}

	machineTimeExample = &TimeConfig{
		TimeServers: []string{"time.cloudflare.com"},
	}

	machineSysctlsExample map[string]string = map[string]string{
		"kernel.domainname":   "talos.dev",
		"net.ipv4.ip_forward": "0",
	}

	clusterControlPlaneExample = &ControlPlaneConfig{
		Endpoint: &Endpoint{
			&url.URL{
				Host:   "1.2.3.4",
				Scheme: "https",
			},
		},
		LocalAPIServerPort: 443,
	}

	clusterNetworkExample = &ClusterNetworkConfig{
		CNI: &CNIConfig{
			CNIName: "flannel",
		},
		DNSDomain:     "cluster.local",
		PodSubnet:     []string{"10.244.0.0/16"},
		ServiceSubnet: []string{"10.96.0.0/12"},
	}

	clusterAPIServerExample = &APIServerConfig{
		ContainerImage: (&APIServerConfig{}).Image(),
		ExtraArgsConfig: map[string]string{
			"key": "value", // TODO: add more real live examples
		},
		CertSANs: []string{
			"1.2.3.4",
			"4.5.6.7",
		},
	}

	clusterControllerManagerExample = &ControllerManagerConfig{
		ContainerImage: (&ControllerManagerConfig{}).Image(),
		ExtraArgsConfig: map[string]string{
			"key": "value", // TODO: add more real live examples
		},
	}

	clusterProxyExample = &ProxyConfig{
		ContainerImage: (&ProxyConfig{}).Image(),
		ExtraArgsConfig: map[string]string{
			"key": "value", // TODO: add more real live examples
		},
		ModeConfig: "ipvs",
	}

	clusterSchedulerConfig = &SchedulerConfig{
		ContainerImage: (&SchedulerConfig{}).Image(),
		ExtraArgsConfig: map[string]string{
			"key": "value", // TODO: add more real live examples
		},
	}

	clusterEtcdConfig = &EtcdConfig{
		ContainerImage: (&EtcdConfig{}).Image(),
		EtcdExtraArgs: map[string]string{
			"key": "value", // TODO: add more real live examples
		},
		RootCA: pemEncodedCertificateExample,
	}

	clusterPodCheckpointerExample = &PodCheckpointer{
		PodCheckpointerImage: "...",
	}

	clusterCoreDNSExample = &CoreDNS{
		CoreDNSImage: (&CoreDNS{}).Image(),
	}

	clusterAdminKubeconfigExample = AdminKubeconfigConfig{
		AdminKubeconfigCertLifetime: time.Hour,
	}

	kubeletExtraMountsExample = []specs.Mount{
		{
			Source:      "/var/lib/example",
			Destination: "/var/lib/example",
			Type:        "bind",
			Options: []string{
				"rshared",
				"ro",
			},
		},
	}

	networkConfigExtraHostsExample = []*ExtraHost{
		{
			HostIP: "192.168.1.100",
			HostAliases: []string{
				"test",
				"test.domain.tld",
			},
		},
	}

	clusterCustomCNIExample = &CNIConfig{
		CNIName: "custom",
		CNIUrls: []string{
			"https://www.mysweethttpserver.com/supersecretcni.yaml",
		},
	}
)

// Config defines the v1alpha1 configuration file.
type Config struct {
	//   description: |
	//     Indicates the schema used to decode the contents.
	//   values:
	//     - "v1alpha1"
	ConfigVersion string `yaml:"version"`
	//   description: |
	//     Enable verbose logging.
	//   values:
	//     - true
	//     - yes
	//     - false
	//     - no
	ConfigDebug bool `yaml:"debug"`
	//   description: |
	//     Indicates whether to pull the machine config upon every boot.
	//   values:
	//     - true
	//     - yes
	//     - false
	//     - no
	ConfigPersist bool `yaml:"persist"`
	//   description: |
	//     Provides machine specific configuration options.
	MachineConfig *MachineConfig `yaml:"machine"`
	//   description: |
	//     Provides cluster specific configuration options.
	ClusterConfig *ClusterConfig `yaml:"cluster"`
}

// MachineConfig reperesents the machine-specific config values.
type MachineConfig struct {
	//   description: |
	//     Defines the role of the machine within the cluster.
	//
	//     #### Init
	//
	//     Init node type designates the first control plane node to come up.
	//     You can think of it like a bootstrap node.
	//     This node will perform the initial steps to bootstrap the cluster -- generation of TLS assets, starting of the control plane, etc.
	//
	//     #### Control Plane
	//
	//     Control Plane node type designates the node as a control plane member.
	//     This means it will host etcd along with the Kubernetes master components such as API Server, Controller Manager, Scheduler.
	//
	//     #### Worker
	//
	//     Worker node type designates the node as a worker node.
	//     This means it will be an available compute node for scheduling workloads.
	//   values:
	//     - "init"
	//     - "controlplane"
	//     - "join"
	MachineType string `yaml:"type"`
	//   description: |
	//     The `token` is used by a machine to join the PKI of the cluster.
	//     Using this token, a machine will create a certificate signing request (CSR), and request a certificate that will be used as its' identity.
	//   examples:
	//     - name: example token
	//       value: "\"328hom.uqjzh6jnn2eie9oi\""
	MachineToken string `yaml:"token"` // Warning: It is important to ensure that this token is correct since a machine's certificate has a short TTL by default
	//   description: |
	//     The root certificate authority of the PKI.
	//     It is composed of a base64 encoded `crt` and `key`.
	//   examples:
	//     - value: pemEncodedCertificateExample
	//       name: machine CA example
	MachineCA *x509.PEMEncodedCertificateAndKey `yaml:"ca,omitempty"`
	//   description: |
	//     Extra certificate subject alternative names for the machine's certificate.
	//     By default, all non-loopback interface IPs are automatically added to the certificate's SANs.
	//   examples:
	//     - name: Uncomment this to enable SANs.
	//       value: '[]string{"10.0.0.10", "172.16.0.10", "192.168.0.10"}'
	MachineCertSANs []string `yaml:"certSANs"`
	//   description: |
	//     Used to provide additional options to the kubelet.
	//   examples:
	//     - name: Kubelet definition example.
	//       value: machineKubeletExample
	MachineKubelet *KubeletConfig `yaml:"kubelet,omitempty"`
	//   description: |
	//     Used to configure the machine's network.
	//   examples:
	//     - name: Network definition example.
	//       value: machineNetworkConfigExample
	MachineNetwork *NetworkConfig `yaml:"network,omitempty"`
	//   description: |
	//     Used to partition, format and mount additional disks.
	//     Since the rootfs is read only with the exception of `/var`, mounts are only valid if they are under `/var`.
	//     Note that the partitioning and formating is done only once, if and only if no existing  partitions are found.
	//     If `size:` is omitted, the partition is sized to occupy full disk.
	//   examples:
	//     - name: MachineDisks list example.
	//       value: machineDisksExample
	MachineDisks []*MachineDisk `yaml:"disks,omitempty"` // Note: `size` is in units of bytes.
	//   description: |
	//     Used to provide instructions for bare-metal installations.
	//   examples:
	//     - name: MachineInstall config usage example.
	//       value: machineInstallExample
	MachineInstall *InstallConfig `yaml:"install,omitempty"`
	//   description: |
	//     Allows the addition of user specified files.
	//     The value of `op` can be `create`, `overwrite`, or `append`.
	//     In the case of `create`, `path` must not exist.
	//     In the case of `overwrite`, and `append`, `path` must be a valid file.
	//     If an `op` value of `append` is used, the existing file will be appended.
	//     Note that the file contents are not required to be base64 encoded.
	//   examples:
	//      - name: MachineFiles usage example.
	//        value: machineFilesExample
	MachineFiles []*MachineFile `yaml:"files,omitempty"` // Note: The specified `path` is relative to `/var`.
	//   description: |
	//     The `env` field allows for the addition of environment variables to a machine.
	//     All environment variables are set on the machine in addition to every service.
	//   values:
	//     - "`GRPC_GO_LOG_VERBOSITY_LEVEL`"
	//     - "`GRPC_GO_LOG_SEVERITY_LEVEL`"
	//     - "`http_proxy`"
	//     - "`https_proxy`"
	//     - "`no_proxy`"
	//   examples:
	//     - name: Environment variables definition examples.
	//       value: machineEnvExamples[0]
	//     - value: machineEnvExamples[1]
	//     - value: machineEnvExamples[2]
	MachineEnv Env `yaml:"env,omitempty"`
	//   description: |
	//     Used to configure the machine's time settings.
	//   examples:
	//     - name: Example configuration for cloudflare ntp server.
	//       value: machineTimeExample
	MachineTime *TimeConfig `yaml:"time,omitempty"`
	//   description: |
	//     Used to configure the machine's sysctls.
	//   examples:
	//     - name: MachineSysctls usage example.
	//       value: machineSysctlsExample
	MachineSysctls map[string]string `yaml:"sysctls,omitempty"`
	//   description: |
	//     Used to configure the machine's container image registry mirrors.
	//
	//     Automatically generates matching CRI configuration for registry mirrors.
	//
	//     Section `mirrors` allows to redirect requests for images to non-default registry,
	//     which might be local registry or caching mirror.
	//
	//     Section `config` provides a way to authenticate to the registry with TLS client
	//     identity, provide registry CA, or authentication information.
	//     Authentication information has same meaning with the corresponding field in `.docker/config.json`.
	//
	//     See also matching configuration for [CRI containerd plugin](https://github.com/containerd/cri/blob/master/docs/registry.md).
	//   examples:
	//     - value: machineConfigRegistriesExample
	MachineRegistries RegistriesConfig `yaml:"registries,omitempty"`
}

// ClusterConfig reperesents the cluster-wide config values.
type ClusterConfig struct {
	//   description: |
	//     Provides control plane specific configuration options.
	//   examples:
	//     - name: Setting controlplain endpoint address to 1.2.3.4 and port to 443 example.
	//       value: clusterControlPlaneExample
	ControlPlane *ControlPlaneConfig `yaml:"controlPlane"`
	//   description: |
	//     Configures the cluster's name.
	ClusterName string `yaml:"clusterName,omitempty"`
	//   description: |
	//     Provides cluster network configuration.
	//   examples:
	//     - name: Configuring with flannel cni and setting up subnets.
	//       value:  clusterNetworkExample
	ClusterNetwork *ClusterNetworkConfig `yaml:"network,omitempty"`
	//   description: |
	//     The [bootstrap token](https://kubernetes.io/docs/reference/access-authn-authz/bootstrap-tokens/).
	//   examples:
	//     - name: Bootstrap token example (do not use in production!).
	//       value: '"wlzjyw.bei2zfylhs2by0wd"'
	BootstrapToken string `yaml:"token,omitempty"`
	//   description: |
	//     The key used for the [encryption of secret data at rest](https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/).
	//   examples:
	//     - name: Decryption secret example (do not use in production!).
	//       value: '"z01mye6j16bspJYtTB/5SFX8j7Ph4JXxM2Xuu4vsBPM="'
	ClusterAESCBCEncryptionSecret string `yaml:"aescbcEncryptionSecret"`
	//   description: |
	//     The base64 encoded root certificate authority used by Kubernetes.
	//   examples:
	//     - name: ClusterCA example.
	//       value: pemEncodedCertificateExample
	ClusterCA *x509.PEMEncodedCertificateAndKey `yaml:"ca,omitempty"`
	//   description: |
	//     API server specific configuration options.
	//   examples:
	//     - value: clusterAPIServerExample
	APIServerConfig *APIServerConfig `yaml:"apiServer,omitempty"`
	//   description: |
	//     Controller manager server specific configuration options.
	//   examples:
	//     - value: clusterControllerManagerExample
	ControllerManagerConfig *ControllerManagerConfig `yaml:"controllerManager,omitempty"`
	//   description: |
	//     Kube-proxy server-specific configuration options
	//   examples:
	//     - value: clusterProxyExample
	ProxyConfig *ProxyConfig `yaml:"proxy,omitempty"`
	//   description: |
	//     Scheduler server specific configuration options.
	//   examples:
	//     - value: clusterSchedulerConfig
	SchedulerConfig *SchedulerConfig `yaml:"scheduler,omitempty"`
	//   description: |
	//     Etcd specific configuration options.
	//   examples:
	//     - value: clusterEtcdConfig
	EtcdConfig *EtcdConfig `yaml:"etcd,omitempty"`
	//   description: |
	//     Pod Checkpointer specific configuration options.
	//   examples:
	//     - value: clusterPodCheckpointerExample
	PodCheckpointerConfig *PodCheckpointer `yaml:"podCheckpointer,omitempty"`
	//   description: |
	//     Core DNS specific configuration options.
	//   examples:
	//     - value: clusterCoreDNSExample
	CoreDNSConfig *CoreDNS `yaml:"coreDNS,omitempty"`
	//   description: |
	//     A list of urls that point to additional manifests.
	//     These will get automatically deployed by bootkube.
	//   examples:
	//     - value: >
	//        []string{
	//         "https://www.mysweethttpserver.com/manifest1.yaml",
	//         "https://www.mysweethttpserver.com/manifest2.yaml",
	//        }
	ExtraManifests []string `yaml:"extraManifests,omitempty"`
	//   description: |
	//     A map of key value pairs that will be added while fetching the ExtraManifests.
	//   examples:
	//     - value: >
	//         map[string]string{
	//           "Token": "1234567",
	//           "X-ExtraInfo": "info",
	//         }
	ExtraManifestHeaders map[string]string `yaml:"extraManifestHeaders,omitempty"`
	//   description: |
	//     Settings for admin kubeconfig generation.
	//     Certificate lifetime can be configured.
	//   examples:
	//     - value: clusterAdminKubeconfigExample
	AdminKubeconfigConfig AdminKubeconfigConfig `yaml:"adminKubeconfig,omitempty"`
	//   description: |
	//     Indicates if master nodes are schedulable.
	//   values:
	//     - true
	//     - yes
	//     - false
	//     - no
	AllowSchedulingOnMasters bool `yaml:"allowSchedulingOnMasters,omitempty"`
}

// KubeletConfig reperesents the kubelet config values.
type KubeletConfig struct {
	//   description: |
	//     The `image` field is an optional reference to an alternative kubelet image.
	//   examples:
	//     - value: kubeletImageExample
	KubeletImage string `yaml:"image,omitempty"`
	//   description: |
	//     The `extraArgs` field is used to provide additional flags to the kubelet.
	//   examples:
	//     - value: >
	//         map[string]string{
	//           "key": "value",
	//         }
	KubeletExtraArgs map[string]string `yaml:"extraArgs,omitempty"`
	//   description: |
	//     The `extraMounts` field is used to add additional mounts to the kubelet container.
	//   examples:
	//     - value: kubeletExtraMountsExample
	KubeletExtraMounts []specs.Mount `yaml:"extraMounts,omitempty"`
}

// NetworkConfig reperesents the machine's networking config values.
type NetworkConfig struct {
	//   description: |
	//     Used to statically set the hostname for the host.
	NetworkHostname string `yaml:"hostname,omitempty"`
	//   description: |
	//     `interfaces` is used to define the network interface configuration.
	//     By default all network interfaces will attempt a DHCP discovery.
	//     This can be further tuned through this configuration parameter.
	//
	//     #### machine.network.interfaces.interface
	//
	//     This is the interface name that should be configured.
	//
	//     #### machine.network.interfaces.cidr
	//
	//     `cidr` is used to specify a static IP address to the interface.
	//     This should be in proper CIDR notation ( `192.168.2.5/24` ).
	//
	//     > Note: This option is mutually exclusive with DHCP.
	//
	//     #### machine.network.interfaces.dhcp
	//
	//     `dhcp` is used to specify that this device should be configured via DHCP.
	//
	//     The following DHCP options are supported:
	//
	//     - `OptionClasslessStaticRoute`
	//     - `OptionDomainNameServer`
	//     - `OptionDNSDomainSearchList`
	//     - `OptionHostName`
	//
	//     > Note: This option is mutually exclusive with CIDR.
	//     >
	//     > Note: To configure an interface with *only* IPv6 SLAAC addressing, CIDR should be set to "" and DHCP to false
	//     > in order for Talos to skip configuration of addresses.
	//     > All other options will still apply.
	//
	//     #### machine.network.interfaces.ignore
	//
	//     `ignore` is used to exclude a specific interface from configuration.
	//     This parameter is optional.
	//
	//     #### machine.network.interfaces.dummy
	//
	//     `dummy` is used to specify that this interface should be a virtual-only, dummy interface.
	//     This parameter is optional.
	//
	//     #### machine.network.interfaces.routes
	//
	//     `routes` is used to specify static routes that may be necessary.
	//     This parameter is optional.
	//
	//     Routes can be repeated and includes a `Network` and `Gateway` field.
	NetworkInterfaces []*Device `yaml:"interfaces,omitempty"`
	//   description: |
	//     Used to statically set the nameservers for the host.
	//     Defaults to `1.1.1.1` and `8.8.8.8`
	NameServers []string `yaml:"nameservers,omitempty"`
	//   description: |
	//     Allows for extra entries to be added to /etc/hosts file
	//   examples:
	//     - value: networkConfigExtraHostsExample
	ExtraHostEntries []*ExtraHost `yaml:"extraHostEntries,omitempty"`
}

// InstallConfig represents the installation options for preparing a node.
type InstallConfig struct {
	//   description: |
	//     The disk used to install the bootloader, and ephemeral partitions.
	//   examples:
	//     - value: '"/dev/sda"'
	//     - value: '"/dev/nvme0"'
	InstallDisk string `yaml:"disk,omitempty"`
	//   description: |
	//     Allows for supplying extra kernel args to the bootloader config.
	//   examples:
	//     - value: '[]string{"a=b"}'
	InstallExtraKernelArgs []string `yaml:"extraKernelArgs,omitempty"`
	//   description: |
	//     Allows for supplying the image used to perform the installation.
	//   examples:
	//     - value: '"docker.io/<org>/installer:latest"'
	InstallImage string `yaml:"image,omitempty"`
	//   description: |
	//     Indicates if a bootloader should be installed.
	//   values:
	//     - true
	//     - yes
	//     - false
	//     - no
	InstallBootloader bool `yaml:"bootloader,omitempty"`
	//   description: |
	//     Indicates if zeroes should be written to the `disk` before performing and installation.
	//     Defaults to `true`.
	//   values:
	//     - true
	//     - yes
	//     - false
	//     - no
	InstallWipe bool `yaml:"wipe"`
}

// TimeConfig represents the options for configuring time on a node.
type TimeConfig struct {
	//   description: |
	//     Indicates if time (ntp) is disabled for the machine
	//     Defaults to `false`.
	TimeDisabled bool `yaml:"disabled"`
	//   description: |
	//     Specifies time (ntp) servers to use for setting system time.
	//     Defaults to `pool.ntp.org`
	//
	//     > Note: This parameter only supports a single time server
	TimeServers []string `yaml:"servers,omitempty"`
}

// RegistriesConfig represents the image pull options.
type RegistriesConfig struct {
	//   description: |
	//     Specifies mirror configuration for each registry.
	//     This setting allows to use local pull-through caching registires,
	//     air-gapped installations, etc.
	//
	//     Registry name is the first segment of image identifier, with 'docker.io'
	//     being default one.
	//     Name '*' catches any registry names not specified explicitly.
	RegistryMirrors map[string]*RegistryMirrorConfig `yaml:"mirrors,omitempty"`
	//   description: |
	//     Specifies TLS & auth configuration for HTTPS image registries.
	//     Mutual TLS can be enabled with 'clientIdentity' option.
	//
	//     TLS configuration can be skipped if registry has trusted
	//     server certificate.
	RegistryConfig map[string]*RegistryConfig `yaml:"config,omitempty"`
}

// PodCheckpointer represents the pod-checkpointer config values.
type PodCheckpointer struct {
	//   description: |
	//     The `image` field is an override to the default pod-checkpointer image.
	PodCheckpointerImage string `yaml:"image,omitempty"`
}

// CoreDNS represents the coredns config values.
type CoreDNS struct {
	//   description: |
	//     The `image` field is an override to the default coredns image.
	CoreDNSImage string `yaml:"image,omitempty"`
}

// Endpoint struct holds the endpoint url parsed out of machine config.
type Endpoint struct {
	*url.URL
}

// UnmarshalYAML is a custom unmarshaller for the endpoint struct.
func (e *Endpoint) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var endpoint string

	if err := unmarshal(&endpoint); err != nil {
		return err
	}

	url, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	*e = Endpoint{url}

	return nil
}

// MarshalYAML is a custom unmarshaller for the endpoint struct.
func (e *Endpoint) MarshalYAML() (interface{}, error) {
	return e.URL.String(), nil
}

// ControlPlaneConfig represents control plane config vals.
type ControlPlaneConfig struct {
	//   description: |
	//     Endpoint is the canonical controlplane endpoint, which can be an IP address or a DNS hostname.
	//     It is single-valued, and may optionally include a port number.
	//   examples:
	//     - value: '"https://1.2.3.4:443"'
	Endpoint *Endpoint `yaml:"endpoint"`
	//   description: |
	//     The port that the API server listens on internally.
	//     This may be different than the port portion listed in the endpoint field above.
	//     The default is 6443.
	LocalAPIServerPort int `yaml:"localAPIServerPort,omitempty"`
}

// APIServerConfig represents kube apiserver config vals.
type APIServerConfig struct {
	//   description: |
	//     The container image used in the API server manifest.
	ContainerImage string `yaml:"image,omitempty"`
	//   description: |
	//     Extra arguments to supply to the API server.
	ExtraArgsConfig map[string]string `yaml:"extraArgs,omitempty"`
	//   description: |
	//     Extra certificate subject alternative names for the API server's certificate.
	CertSANs []string `yaml:"certSANs,omitempty"`
}

// ControllerManagerConfig represents kube controller manager config vals.
type ControllerManagerConfig struct {
	//   description: |
	//     The container image used in the controller manager manifest.
	ContainerImage string `yaml:"image,omitempty"`
	//   description: |
	//     Extra arguments to supply to the controller manager.
	ExtraArgsConfig map[string]string `yaml:"extraArgs,omitempty"`
}

// ProxyConfig represents the kube proxy configuration values.
type ProxyConfig struct {
	//   description: |
	//     The container image used in the kube-proxy manifest.
	ContainerImage string `yaml:"image,omitempty"`
	//   description: |
	//     proxy mode of kube-proxy.
	//     By default, this is 'iptables'.
	ModeConfig string `yaml:"mode,omitempty"`
	//   description: |
	//     Extra arguments to supply to kube-proxy.
	ExtraArgsConfig map[string]string `yaml:"extraArgs,omitempty"`
}

// SchedulerConfig represents kube scheduler config vals.
type SchedulerConfig struct {
	//   description: |
	//     The container image used in the scheduler manifest.
	ContainerImage string `yaml:"image,omitempty"`
	//   description: |
	//     Extra arguments to supply to the scheduler.
	ExtraArgsConfig map[string]string `yaml:"extraArgs,omitempty"`
}

// EtcdConfig represents etcd config vals.
type EtcdConfig struct {
	//   description: |
	//     The container image used to create the etcd service.
	ContainerImage string `yaml:"image,omitempty"`
	//   description: |
	//     The `ca` is the root certificate authority of the PKI.
	//     It is composed of a base64 encoded `crt` and `key`.
	//   examples:
	//     - value: pemEncodedCertificateExample
	RootCA *x509.PEMEncodedCertificateAndKey `yaml:"ca"`
	//   description: |
	//     Extra arguments to supply to etcd.
	//     Note that the following args are not allowed:
	//
	//     - `name`
	//     - `data-dir`
	//     - `initial-cluster-state`
	//     - `listen-peer-urls`
	//     - `listen-client-urls`
	//     - `cert-file`
	//     - `key-file`
	//     - `trusted-ca-file`
	//     - `peer-client-cert-auth`
	//     - `peer-cert-file`
	//     - `peer-trusted-ca-file`
	//     - `peer-key-file`
	//   examples:
	//     - values: >
	//         map[string]string{
	//           "initial-cluster": "https://1.2.3.4:2380",
	//           "advertise-client-urls": "https://1.2.3.4:2379",
	//         }
	EtcdExtraArgs map[string]string `yaml:"extraArgs,omitempty"`
}

// ClusterNetworkConfig represents kube networking config vals.
type ClusterNetworkConfig struct {
	//   description: |
	//     The CNI used.
	//     Composed of "name" and "url".
	//     The "name" key only supports upstream bootkube options of "flannel" or "custom".
	//     URLs is only used if name is equal to "custom".
	//     URLs should point to a single yaml file that will get deployed.
	//     Empty struct or any other name will default to bootkube's flannel.
	//   examples:
	//     - value: clusterCustomCNIExample
	CNI *CNIConfig `yaml:"cni,omitempty"`
	//   description: |
	//     The domain used by Kubernetes DNS.
	//     The default is `cluster.local`
	//   examples:
	//     - value: '"cluser.local"'
	DNSDomain string `yaml:"dnsDomain"`
	//   description: |
	//     The pod subnet CIDR.
	//   examples:
	//     -  value: >
	//          []string{"10.244.0.0/16"}
	PodSubnet []string `yaml:"podSubnets"`
	//   description: |
	//     The service subnet CIDR.
	//   examples:
	//   examples:
	//     -  value: >
	//          []string{"10.96.0.0/12"}
	ServiceSubnet []string `yaml:"serviceSubnets"`
}

// CNIConfig contains the info about which CNI we'll deploy.
type CNIConfig struct {
	//   description: |
	//     Name of CNI to use.
	CNIName string `yaml:"name"`
	//   description: |
	//     URLs containing manifests to apply for CNI.
	CNIUrls []string `yaml:"urls,omitempty"`
}

// AdminKubeconfigConfig contains admin kubeconfig settings.
type AdminKubeconfigConfig struct {
	//   description: |
	//     Admin kubeconfig certificate lifetime (default is 1 year).
	//     Field format accepts any Go time.Duration format ('1h' for one hour, '10m' for ten minutes).
	AdminKubeconfigCertLifetime time.Duration `yaml:"certLifetime,omitempty"`
}

// MachineDisk represents the options available for partitioning, formatting, and
// mounting extra disks.
type MachineDisk struct {
	//   description: The name of the disk to use.
	DeviceName string `yaml:"device,omitempty"`
	//   description: A list of partitions to create on the disk.
	DiskPartitions []*DiskPartition `yaml:"partitions,omitempty"`
}

// DiskPartition represents the options for a device partition.
type DiskPartition struct {
	//   description: |
	//     This size of the partition in bytes.
	DiskSize uint64 `yaml:"size,omitempty"`
	//   description:
	//     Where to mount the partition.
	DiskMountPoint string `yaml:"mountpoint,omitempty"`
}

// Env represents a set of environment variables.
type Env = map[string]string

// MachineFile represents a file to write to disk.
type MachineFile struct {
	//   description: The contents of file.
	FileContent string `yaml:"content"`
	//   description: The file's permissions in octal.
	FilePermissions os.FileMode `yaml:"permissions"`
	//   description: The path of the file.
	FilePath string `yaml:"path"`
	//   description: The operation to use
	//   values:
	//     - create
	//     - append
	FileOp string `yaml:"op"`
}

// ExtraHost represents a host entry in /etc/hosts.
type ExtraHost struct {
	//   description: The IP of the host.
	HostIP string `yaml:"ip"`
	//   description: The host alias.
	HostAliases []string `yaml:"aliases"`
}

// Device represents a network interface.
type Device struct {
	//   description: The interface name.
	DeviceInterface string `yaml:"interface"`
	//   description: The CIDR to use.
	DeviceCIDR string `yaml:"cidr"`
	//   description: |
	//     A list of routes associated with the interface.
	//     If used in combination with DHCP, these routes will be appended to routes returned by DHCP server.
	DeviceRoutes []*Route `yaml:"routes"`
	//   description: Bond specific options.
	DeviceBond *Bond `yaml:"bond"`
	//   description: VLAN specific options.
	DeviceVlans []*Vlan `yaml:"vlans"`
	//   description: |
	//     The interface's MTU.
	//     If used in combination with DHCP, this will override any MTU settings returned from DHCP server.
	DeviceMTU int `yaml:"mtu"`
	//   description: Indicates if DHCP should be used.
	DeviceDHCP bool `yaml:"dhcp"`
	//   description: Indicates if the interface should be ignored.
	DeviceIgnore bool `yaml:"ignore"`
	//   description: Indicates if the interface is a dummy interface.
	DeviceDummy bool `yaml:"dummy"`
	//   description: |
	//     DHCP specific options.
	//     DHCP *must* be set to true for these to take effect.
	DeviceDHCPOptions *DHCPOptions `yaml:"dhcpOptions"`
}

// DHCPOptions contains options for configuring the DHCP settings for a given interface.
type DHCPOptions struct {
	//   description: The priority of all routes received via DHCP
	DHCPRouteMetric uint32 `yaml:"routeMetric"`
}

// Bond contains the various options for configuring a
// bonded interface.
type Bond struct {
	//   description: The interfaces that make up the bond.
	BondInterfaces []string `yaml:"interfaces"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondARPIPTarget []string `yaml:"arpIPTarget"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondMode string `yaml:"mode"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondHashPolicy string `yaml:"xmitHashPolicy"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondLACPRate string `yaml:"lacpRate"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondADActorSystem string `yaml:"adActorSystem"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondARPValidate string `yaml:"arpValidate"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondARPAllTargets string `yaml:"arpAllTargets"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondPrimary string `yaml:"primary"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondPrimaryReselect string `yaml:"primaryReselect"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondFailOverMac string `yaml:"failOverMac"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondADSelect string `yaml:"adSelect"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondMIIMon uint32 `yaml:"miimon"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondUpDelay uint32 `yaml:"updelay"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondDownDelay uint32 `yaml:"downdelay"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondARPInterval uint32 `yaml:"arpInterval"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondResendIGMP uint32 `yaml:"resendIgmp"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondMinLinks uint32 `yaml:"minLinks"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondLPInterval uint32 `yaml:"lpInterval"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondPacketsPerSlave uint32 `yaml:"packetsPerSlave"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondNumPeerNotif uint8 `yaml:"numPeerNotif"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondTLBDynamicLB uint8 `yaml:"tlbDynamicLb"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondAllSlavesActive uint8 `yaml:"allSlavesActive"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondUseCarrier bool `yaml:"useCarrier"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondADActorSysPrio uint16 `yaml:"adActorSysPrio"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondADUserPortKey uint16 `yaml:"adUserPortKey"`
	//   description: |
	//     A bond option.
	//     Please see the official kernel documentation.
	BondPeerNotifyDelay uint32 `yaml:"peerNotifyDelay"`
}

// Vlan represents vlan settings for a device.
type Vlan struct {
	//   description: The CIDR to use.
	VlanCIDR string `yaml:"cidr"`
	//   description: A list of routes associated with the VLAN.
	VlanRoutes []*Route `yaml:"routes"`
	//   description: Indicates if DHCP should be used.
	VlanDHCP bool `yaml:"dhcp"`
	//   description: The VLAN's ID.
	VlanID uint16 `yaml:"vlanId"`
}

// Route represents a network route.
type Route struct {
	//   description: The route's network.
	RouteNetwork string `yaml:"network"`
	//   description: The route's gateway.
	RouteGateway string `yaml:"gateway"`
}

// RegistryMirrorConfig represents mirror configuration for a registry.
type RegistryMirrorConfig struct {
	//   description: |
	//     List of endpoints (URLs) for registry mirrors to use.
	//     Endpoint configures HTTP/HTTPS access mode, host name,
	//     port and path (if path is not set, it defaults to `/v2`).
	MirrorEndpoints []string `yaml:"endpoints"`
}

// RegistryConfig specifies auth & TLS config per registry.
type RegistryConfig struct {
	//   description: The TLS configuration for this registry.
	RegistryTLS *RegistryTLSConfig `yaml:"tls,omitempty"`
	//   description: The auth configuration for this registry.
	RegistryAuth *RegistryAuthConfig `yaml:"auth,omitempty"`
}

// RegistryAuthConfig specifies authentication configuration for a registry.
type RegistryAuthConfig struct {
	//   description: |
	//     Optional registry authentication.
	//     The meaning of each field is the same with the corresponding field in .docker/config.json.
	RegistryUsername string `yaml:"username,omitempty"`
	//   description: |
	//     Optional registry authentication.
	//     The meaning of each field is the same with the corresponding field in .docker/config.json.
	RegistryPassword string `yaml:"password,omitempty"`
	//   description: |
	//     Optional registry authentication.
	//     The meaning of each field is the same with the corresponding field in .docker/config.json.
	RegistryAuth string `yaml:"auth,omitempty"`
	//   description: |
	//     Optional registry authentication.
	//     The meaning of each field is the same with the corresponding field in .docker/config.json.
	RegistryIdentityToken string `yaml:"identityToken,omitempty"`
}

// RegistryTLSConfig specifies TLS config for HTTPS registries.
type RegistryTLSConfig struct {
	//   description: |
	//     Enable mutual TLS authentication with the registry.
	//     Client certificate and key should be base64-encoded.
	//   examples:
	//     - value: pemEncodedCertificateExample
	TLSClientIdentity *x509.PEMEncodedCertificateAndKey `yaml:"clientIdentity,omitempty"`
	//   description: |
	//     CA registry certificate to add the list of trusted certificates.
	//     Certificate should be base64-encoded.
	TLSCA Base64Bytes `yaml:"ca,omitempty"`
	//   description: |
	//     Skip TLS server certificate verification (not recommended).
	TLSInsecureSkipVerify bool `yaml:"insecureSkipVerify,omitempty"`
}