// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package nethttp

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

type bpf_debugGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

type bpf_debugHttpClientDataT struct {
	Method        [7]uint8
	Path          [100]uint8
	_             [5]byte
	ContentLength int64
	Pid           struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_ [4]byte
}

type bpf_debugHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              bpf_debugTpInfoT
}

type bpf_debugServerHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              bpf_debugTpInfoT
	Response        uint64
	Method          [7]uint8
	Path            [100]uint8
	_               [5]byte
	ContentLength   uint64
	Http2           uint8
	_               [7]byte
}

type bpf_debugSqlFuncInvocationT struct {
	StartMonotimeNs uint64
	SqlParam        uint64
	QueryLen        uint64
	Tp              bpf_debugTpInfoT
}

type bpf_debugTpInfoPidT struct {
	Tp    bpf_debugTpInfoT
	Pid   uint32
	Valid uint8
	_     [3]byte
}

type bpf_debugTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
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
	UprobeServeHTTP                           *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTP"`
	UprobeServeHTTPReturns                    *ebpf.ProgramSpec `ebpf:"uprobe_ServeHTTPReturns"`
	UprobeConnServe                           *ebpf.ProgramSpec `ebpf:"uprobe_connServe"`
	UprobeConnServeRet                        *ebpf.ProgramSpec `ebpf:"uprobe_connServeRet"`
	UprobeExecDC                              *ebpf.ProgramSpec `ebpf:"uprobe_execDC"`
	UprobeHttp2FramerWriteHeaders             *ebpf.ProgramSpec `ebpf:"uprobe_http2FramerWriteHeaders"`
	UprobeHttp2FramerWriteHeadersReturns      *ebpf.ProgramSpec `ebpf:"uprobe_http2FramerWriteHeaders_returns"`
	UprobeHttp2ResponseWriterStateWriteHeader *ebpf.ProgramSpec `ebpf:"uprobe_http2ResponseWriterStateWriteHeader"`
	UprobeHttp2RoundTrip                      *ebpf.ProgramSpec `ebpf:"uprobe_http2RoundTrip"`
	UprobeHttp2serverConnRunHandler           *ebpf.ProgramSpec `ebpf:"uprobe_http2serverConn_runHandler"`
	UprobeNetFdRead                           *ebpf.ProgramSpec `ebpf:"uprobe_netFdRead"`
	UprobePersistConnRoundTrip                *ebpf.ProgramSpec `ebpf:"uprobe_persistConnRoundTrip"`
	UprobeQueryDC                             *ebpf.ProgramSpec `ebpf:"uprobe_queryDC"`
	UprobeQueryReturn                         *ebpf.ProgramSpec `ebpf:"uprobe_queryReturn"`
	UprobeReadRequestReturns                  *ebpf.ProgramSpec `ebpf:"uprobe_readRequestReturns"`
	UprobeReadRequestStart                    *ebpf.ProgramSpec `ebpf:"uprobe_readRequestStart"`
	UprobeRoundTrip                           *ebpf.ProgramSpec `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn                     *ebpf.ProgramSpec `ebpf:"uprobe_roundTripReturn"`
	UprobeWriteSubset                         *ebpf.ProgramSpec `ebpf:"uprobe_writeSubset"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	DebugEvents                   *ebpf.MapSpec `ebpf:"debug_events"`
	Events                        *ebpf.MapSpec `ebpf:"events"`
	GoTraceMap                    *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap     *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	OngoingClientConnections      *ebpf.MapSpec `ebpf:"ongoing_client_connections"`
	OngoingGoroutines             *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests     *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingHttpClientRequestsData *ebpf.MapSpec `ebpf:"ongoing_http_client_requests_data"`
	OngoingHttpServerRequests     *ebpf.MapSpec `ebpf:"ongoing_http_server_requests"`
	OngoingServerConnections      *ebpf.MapSpec `ebpf:"ongoing_server_connections"`
	OngoingSqlQueries             *ebpf.MapSpec `ebpf:"ongoing_sql_queries"`
	TraceMap                      *ebpf.MapSpec `ebpf:"trace_map"`
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
	DebugEvents                   *ebpf.Map `ebpf:"debug_events"`
	Events                        *ebpf.Map `ebpf:"events"`
	GoTraceMap                    *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap     *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	OngoingClientConnections      *ebpf.Map `ebpf:"ongoing_client_connections"`
	OngoingGoroutines             *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingHttpClientRequests     *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingHttpClientRequestsData *ebpf.Map `ebpf:"ongoing_http_client_requests_data"`
	OngoingHttpServerRequests     *ebpf.Map `ebpf:"ongoing_http_server_requests"`
	OngoingServerConnections      *ebpf.Map `ebpf:"ongoing_server_connections"`
	OngoingSqlQueries             *ebpf.Map `ebpf:"ongoing_sql_queries"`
	TraceMap                      *ebpf.Map `ebpf:"trace_map"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.DebugEvents,
		m.Events,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.OngoingClientConnections,
		m.OngoingGoroutines,
		m.OngoingHttpClientRequests,
		m.OngoingHttpClientRequestsData,
		m.OngoingHttpServerRequests,
		m.OngoingServerConnections,
		m.OngoingSqlQueries,
		m.TraceMap,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	UprobeServeHTTP                           *ebpf.Program `ebpf:"uprobe_ServeHTTP"`
	UprobeServeHTTPReturns                    *ebpf.Program `ebpf:"uprobe_ServeHTTPReturns"`
	UprobeConnServe                           *ebpf.Program `ebpf:"uprobe_connServe"`
	UprobeConnServeRet                        *ebpf.Program `ebpf:"uprobe_connServeRet"`
	UprobeExecDC                              *ebpf.Program `ebpf:"uprobe_execDC"`
	UprobeHttp2FramerWriteHeaders             *ebpf.Program `ebpf:"uprobe_http2FramerWriteHeaders"`
	UprobeHttp2FramerWriteHeadersReturns      *ebpf.Program `ebpf:"uprobe_http2FramerWriteHeaders_returns"`
	UprobeHttp2ResponseWriterStateWriteHeader *ebpf.Program `ebpf:"uprobe_http2ResponseWriterStateWriteHeader"`
	UprobeHttp2RoundTrip                      *ebpf.Program `ebpf:"uprobe_http2RoundTrip"`
	UprobeHttp2serverConnRunHandler           *ebpf.Program `ebpf:"uprobe_http2serverConn_runHandler"`
	UprobeNetFdRead                           *ebpf.Program `ebpf:"uprobe_netFdRead"`
	UprobePersistConnRoundTrip                *ebpf.Program `ebpf:"uprobe_persistConnRoundTrip"`
	UprobeQueryDC                             *ebpf.Program `ebpf:"uprobe_queryDC"`
	UprobeQueryReturn                         *ebpf.Program `ebpf:"uprobe_queryReturn"`
	UprobeReadRequestReturns                  *ebpf.Program `ebpf:"uprobe_readRequestReturns"`
	UprobeReadRequestStart                    *ebpf.Program `ebpf:"uprobe_readRequestStart"`
	UprobeRoundTrip                           *ebpf.Program `ebpf:"uprobe_roundTrip"`
	UprobeRoundTripReturn                     *ebpf.Program `ebpf:"uprobe_roundTripReturn"`
	UprobeWriteSubset                         *ebpf.Program `ebpf:"uprobe_writeSubset"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.UprobeServeHTTP,
		p.UprobeServeHTTPReturns,
		p.UprobeConnServe,
		p.UprobeConnServeRet,
		p.UprobeExecDC,
		p.UprobeHttp2FramerWriteHeaders,
		p.UprobeHttp2FramerWriteHeadersReturns,
		p.UprobeHttp2ResponseWriterStateWriteHeader,
		p.UprobeHttp2RoundTrip,
		p.UprobeHttp2serverConnRunHandler,
		p.UprobeNetFdRead,
		p.UprobePersistConnRoundTrip,
		p.UprobeQueryDC,
		p.UprobeQueryReturn,
		p.UprobeReadRequestReturns,
		p.UprobeReadRequestStart,
		p.UprobeRoundTrip,
		p.UprobeRoundTripReturn,
		p.UprobeWriteSubset,
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
