// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package httpssl

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_debugHttp2ConnStreamT struct {
	PidConn  bpf_debugPidConnectionInfoT
	StreamId uint32
}

type bpf_debugHttp2GrpcRequestT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_debugConnectionInfoT
	Data            [256]uint8
	RetData         [64]uint8
	Type            uint8
	_               [1]byte
	Len             int32
	_               [4]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Ssl uint8
	_   [3]byte
	Tp  struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_debugHttpConnectionMetadataT struct {
	Pid struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Type uint8
}

type bpf_debugHttpInfoT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_debugConnectionInfoT
	_               [2]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [160]uint8
	Len             uint32
	RespLen         uint32
	Status          uint16
	Type            uint8
	Ssl             uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_debugPidConnectionInfoT struct {
	Conn bpf_debugConnectionInfoT
	Pid  uint32
}

type bpf_debugPidKeyT struct {
	Pid uint32
	Ns  uint32
}

type bpf_debugSslArgsT struct {
	Ssl    uint64
	Buf    uint64
	LenPtr uint64
}

type bpf_debugSslPidConnectionInfoT struct {
	P_conn    bpf_debugPidConnectionInfoT
	OrigDport uint16
	C_tid     bpf_debugPidKeyT
	_         [2]byte
}

type bpf_debugTcpReqT struct {
	Flags           uint8
	_               [1]byte
	ConnInfo        bpf_debugConnectionInfoT
	_               [2]byte
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	Rbuf            [128]uint8
	Len             uint32
	RespLen         uint32
	Ssl             uint8
	Direction       uint8
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_  [2]byte
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
}

type bpf_debugTpInfoPidT struct {
	Tp struct {
		TraceId  [16]uint8
		SpanId   [8]uint8
		ParentId [8]uint8
		Ts       uint64
		Flags    uint8
		_        [7]byte
	}
	Pid   uint32
	Valid uint8
	_     [3]byte
}

// loadBpf_debug returns the embedded CollectionSpec for bpf_debug.
func loadBpf_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_debug: %w", err)
	}

	return spec, err
}

// loadBpf_debugObjects loads bpf_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_debugObjects
//	*bpf_debugPrograms
//	*bpf_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugSpecs struct {
	bpf_debugProgramSpecs
	bpf_debugMapSpecs
}

// bpf_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugProgramSpecs struct {
	UprobeSslDoHandshake    *ebpf.ProgramSpec `ebpf:"uprobe_ssl_do_handshake"`
	UprobeSslRead           *ebpf.ProgramSpec `ebpf:"uprobe_ssl_read"`
	UprobeSslReadEx         *ebpf.ProgramSpec `ebpf:"uprobe_ssl_read_ex"`
	UprobeSslShutdown       *ebpf.ProgramSpec `ebpf:"uprobe_ssl_shutdown"`
	UprobeSslWrite          *ebpf.ProgramSpec `ebpf:"uprobe_ssl_write"`
	UprobeSslWriteEx        *ebpf.ProgramSpec `ebpf:"uprobe_ssl_write_ex"`
	UretprobeSslDoHandshake *ebpf.ProgramSpec `ebpf:"uretprobe_ssl_do_handshake"`
	UretprobeSslRead        *ebpf.ProgramSpec `ebpf:"uretprobe_ssl_read"`
	UretprobeSslReadEx      *ebpf.ProgramSpec `ebpf:"uretprobe_ssl_read_ex"`
	UretprobeSslWrite       *ebpf.ProgramSpec `ebpf:"uretprobe_ssl_write"`
	UretprobeSslWriteEx     *ebpf.ProgramSpec `ebpf:"uretprobe_ssl_write_ex"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	ActiveSslConnections    *ebpf.MapSpec `ebpf:"active_ssl_connections"`
	ActiveSslHandshakes     *ebpf.MapSpec `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs       *ebpf.MapSpec `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.MapSpec `ebpf:"active_ssl_write_args"`
	CloneMap                *ebpf.MapSpec `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.MapSpec `ebpf:"connection_meta_mem"`
	DebugEvents             *ebpf.MapSpec `ebpf:"debug_events"`
	Events                  *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections     *ebpf.MapSpec `ebpf:"filtered_connections"`
	Http2InfoMem            *ebpf.MapSpec `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.MapSpec `ebpf:"http_info_mem"`
	OngoingHttp             *ebpf.MapSpec `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.MapSpec `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.MapSpec `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.MapSpec `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.MapSpec `ebpf:"ongoing_tcp_req"`
	PidCache                *ebpf.MapSpec `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.MapSpec `ebpf:"pid_tid_to_conn"`
	ServerTraces            *ebpf.MapSpec `ebpf:"server_traces"`
	SslToConn               *ebpf.MapSpec `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.MapSpec `ebpf:"ssl_to_pid_tid"`
	TcpReqMem               *ebpf.MapSpec `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.MapSpec `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.MapSpec `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids               *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugObjects struct {
	bpf_debugPrograms
	bpf_debugMaps
}

func (o *bpf_debugObjects) Close() error {
	return _Bpf_debugClose(
		&o.bpf_debugPrograms,
		&o.bpf_debugMaps,
	)
}

// bpf_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugMaps struct {
	ActiveSslConnections    *ebpf.Map `ebpf:"active_ssl_connections"`
	ActiveSslHandshakes     *ebpf.Map `ebpf:"active_ssl_handshakes"`
	ActiveSslReadArgs       *ebpf.Map `ebpf:"active_ssl_read_args"`
	ActiveSslWriteArgs      *ebpf.Map `ebpf:"active_ssl_write_args"`
	CloneMap                *ebpf.Map `ebpf:"clone_map"`
	ConnectionMetaMem       *ebpf.Map `ebpf:"connection_meta_mem"`
	DebugEvents             *ebpf.Map `ebpf:"debug_events"`
	Events                  *ebpf.Map `ebpf:"events"`
	FilteredConnections     *ebpf.Map `ebpf:"filtered_connections"`
	Http2InfoMem            *ebpf.Map `ebpf:"http2_info_mem"`
	HttpInfoMem             *ebpf.Map `ebpf:"http_info_mem"`
	OngoingHttp             *ebpf.Map `ebpf:"ongoing_http"`
	OngoingHttp2Connections *ebpf.Map `ebpf:"ongoing_http2_connections"`
	OngoingHttp2Grpc        *ebpf.Map `ebpf:"ongoing_http2_grpc"`
	OngoingHttpFallback     *ebpf.Map `ebpf:"ongoing_http_fallback"`
	OngoingTcpReq           *ebpf.Map `ebpf:"ongoing_tcp_req"`
	PidCache                *ebpf.Map `ebpf:"pid_cache"`
	PidTidToConn            *ebpf.Map `ebpf:"pid_tid_to_conn"`
	ServerTraces            *ebpf.Map `ebpf:"server_traces"`
	SslToConn               *ebpf.Map `ebpf:"ssl_to_conn"`
	SslToPidTid             *ebpf.Map `ebpf:"ssl_to_pid_tid"`
	TcpReqMem               *ebpf.Map `ebpf:"tcp_req_mem"`
	TpCharBufMem            *ebpf.Map `ebpf:"tp_char_buf_mem"`
	TpInfoMem               *ebpf.Map `ebpf:"tp_info_mem"`
	TraceMap                *ebpf.Map `ebpf:"trace_map"`
	ValidPids               *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.ActiveSslConnections,
		m.ActiveSslHandshakes,
		m.ActiveSslReadArgs,
		m.ActiveSslWriteArgs,
		m.CloneMap,
		m.ConnectionMetaMem,
		m.DebugEvents,
		m.Events,
		m.FilteredConnections,
		m.Http2InfoMem,
		m.HttpInfoMem,
		m.OngoingHttp,
		m.OngoingHttp2Connections,
		m.OngoingHttp2Grpc,
		m.OngoingHttpFallback,
		m.OngoingTcpReq,
		m.PidCache,
		m.PidTidToConn,
		m.ServerTraces,
		m.SslToConn,
		m.SslToPidTid,
		m.TcpReqMem,
		m.TpCharBufMem,
		m.TpInfoMem,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	UprobeSslDoHandshake    *ebpf.Program `ebpf:"uprobe_ssl_do_handshake"`
	UprobeSslRead           *ebpf.Program `ebpf:"uprobe_ssl_read"`
	UprobeSslReadEx         *ebpf.Program `ebpf:"uprobe_ssl_read_ex"`
	UprobeSslShutdown       *ebpf.Program `ebpf:"uprobe_ssl_shutdown"`
	UprobeSslWrite          *ebpf.Program `ebpf:"uprobe_ssl_write"`
	UprobeSslWriteEx        *ebpf.Program `ebpf:"uprobe_ssl_write_ex"`
	UretprobeSslDoHandshake *ebpf.Program `ebpf:"uretprobe_ssl_do_handshake"`
	UretprobeSslRead        *ebpf.Program `ebpf:"uretprobe_ssl_read"`
	UretprobeSslReadEx      *ebpf.Program `ebpf:"uretprobe_ssl_read_ex"`
	UretprobeSslWrite       *ebpf.Program `ebpf:"uretprobe_ssl_write"`
	UretprobeSslWriteEx     *ebpf.Program `ebpf:"uretprobe_ssl_write_ex"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.UprobeSslDoHandshake,
		p.UprobeSslRead,
		p.UprobeSslReadEx,
		p.UprobeSslShutdown,
		p.UprobeSslWrite,
		p.UprobeSslWriteEx,
		p.UretprobeSslDoHandshake,
		p.UretprobeSslRead,
		p.UretprobeSslReadEx,
		p.UretprobeSslWrite,
		p.UretprobeSslWriteEx,
	)
}

func _Bpf_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_debug_bpfel_arm64.o
var _Bpf_debugBytes []byte
