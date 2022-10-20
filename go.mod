module github.com/mxyns/twitchatbeat

go 1.17

replace (
	github.com/Microsoft/go-winio => github.com/bi-zone/go-winio v0.4.15
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20210823122811-11c3ef800752
	github.com/apoydence/eachers => github.com/poy/eachers v0.0.0-20181020210610-23942921fe77 //indirect, see https://github.com/elastic/beats/pull/29780 for details.
	github.com/cucumber/godog => github.com/cucumber/godog v0.8.1
	github.com/docker/docker => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
	github.com/docker/go-plugins-helpers => github.com/elastic/go-plugins-helpers v0.0.0-20200207104224-bdf17607b79f
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/dop251/goja_nodejs => github.com/dop251/goja_nodejs v0.0.0-20171011081505-adff31b136e6
	github.com/fsnotify/fsevents => github.com/elastic/fsevents v0.0.0-20181029231046-e1d381a4d270
	github.com/fsnotify/fsnotify => github.com/adriansr/fsnotify v1.4.8-0.20211018144411-a81f2b630e7c
	github.com/golang/glog => github.com/elastic/glog v1.0.1-0.20210831205241-7d8b5c89dfc4
	github.com/google/gopacket => github.com/adriansr/gopacket v1.1.18-0.20200327165309-dd62abfa8a41
	github.com/insomniacslk/dhcp => github.com/elastic/dhcp v0.0.0-20200227161230-57ec251c7eb3 // indirect
	github.com/tonistiigi/fifo => github.com/containerd/fifo v0.0.0-20190816180239-bda0ff6ed73c
)

require (
	github.com/blakesmith/ar v0.0.0-20190502131153-809d4375e1fb
	github.com/cavaliergopher/rpm v1.2.0
	github.com/elastic/beats/v7 v7.17.6
	github.com/magefile/mage v1.11.0
	github.com/mitchellh/gox v1.0.1
	github.com/mxyns/justlog v1.0.1
	github.com/pierrre/gotestcover v0.0.0-20160517101806-924dca7d15f0
	github.com/tsg/go-daemon v0.0.0-20200207173439-e704b93fd89b
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/tools v0.1.12
	gotest.tools/gotestsum v1.8.2
)

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/Shopify/sarama v0.0.0-00010101000000-000000000000 // indirect
	github.com/StackExchange/wmi v0.0.0-20170221213301-9f32b5905fd6 // indirect
	github.com/akavel/rsrc v0.8.0 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/containerd/containerd v1.5.7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dlclark/regexp2 v1.1.7-0.20171009020623-7632a260cbaf // indirect
	github.com/dnephin/pflag v1.0.7 // indirect
	github.com/docker/distribution v2.8.0+incompatible // indirect
	github.com/docker/docker v1.4.2-0.20190924003213-a8608b5b67c7 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/dop251/goja v0.0.0-20200831102558-9af81ddcf0e1 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20171011081505-adff31b136e6 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/ecs v1.12.0 // indirect
	github.com/elastic/elastic-agent-client/v7 v7.0.0-20210727140539-f0905d9377f6 // indirect
	github.com/elastic/go-concert v0.2.0 // indirect
	github.com/elastic/go-lumber v0.1.0 // indirect
	github.com/elastic/go-seccomp-bpf v1.2.0 // indirect
	github.com/elastic/go-structform v0.0.9 // indirect
	github.com/elastic/go-sysinfo v1.8.1 // indirect
	github.com/elastic/go-txfile v0.0.7 // indirect
	github.com/elastic/go-ucfg v0.8.6 // indirect
	github.com/elastic/go-windows v1.0.1 // indirect
	github.com/elastic/gosigar v0.14.2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gempir/go-twitch-irc/v3 v3.0.0 // indirect
	github.com/go-logr/logr v0.4.0 // indirect
	github.com/go-ole/go-ole v1.2.5-0.20190920104607-14974a1cf647 // indirect
	github.com/go-sourcemap/sourcemap v2.1.2+incompatible // indirect
	github.com/gofrs/flock v0.7.2-0.20190320160742-5135e617513b // indirect
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.8.3 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/h2non/filetype v1.1.1 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.0.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/josephspurrier/goversioninfo v0.0.0-20190209210621-63e6d1acd3dd // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/miekg/dns v1.1.25 // indirect
	github.com/mitchellh/hashstructure v0.0.0-20170116052023-ab25296c0f51 // indirect
	github.com/mitchellh/iochan v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nicklaw5/helix v1.25.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2-0.20190823105129-775207bd45b6 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/santhosh-tekuri/jsonschema v1.2.4 // indirect
	github.com/shirou/gopsutil v3.20.12+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/urso/diag v0.0.0-20200210123136-21b3cc8eb797 // indirect
	github.com/urso/go-bin v0.0.0-20180220135811-781c575c9f0e // indirect
	github.com/urso/magetools v0.0.0-20190919040553-290c89e0c230 // indirect
	github.com/urso/sderr v0.0.0-20210525210834-52b04e8f5c71 // indirect
	github.com/xdg/scram v1.0.3 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	go.elastic.co/apm v1.11.0 // indirect
	go.elastic.co/apm/module/apmelasticsearch v1.7.2 // indirect
	go.elastic.co/apm/module/apmhttp v1.7.2 // indirect
	go.elastic.co/ecszap v0.3.0 // indirect
	go.elastic.co/fastjson v1.1.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.14.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211021150943-2b146023228c // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	howett.net/plist v1.0.0 // indirect
	k8s.io/api v0.21.1 // indirect
	k8s.io/apimachinery v0.21.1 // indirect
	k8s.io/client-go v0.21.1 // indirect
	k8s.io/klog/v2 v2.8.0 // indirect
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
