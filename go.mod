module github.com/justcy/uc

go 1.13

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.49.0
	cloud.google.com/go/bigquery => github.com/googleapis/google-cloud-go/bigquery v1.3.0
	cloud.google.com/go/datastore => github.com/googleapis/google-cloud-go/datastore v1.0.0
	cloud.google.com/go/pubsub => github.com/googleapis/google-cloud-go/pubsub v1.1.0
	cloud.google.com/go/storage => github.com/googleapis/google-cloud-go/storage v1.4.0
	github.com/go-xorm/core v0.6.3 => xorm.io/core v0.6.3
	github.com/openzipkin/zipkin-go-opentracing => github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	go.uber.org/atomic => github.com/uber-go/atomic v1.5.1
	go.uber.org/multierr => github.com/uber-go/multierr v1.4.0
	go.uber.org/zap => github.com/uber-go/zap v1.13.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/exp => github.com/golang/exp v0.0.0-20191129062945-2f5052295587
	golang.org/x/image => github.com/golang/image v0.0.0-20191214001246-9130b4cfad52
	golang.org/x/lint => github.com/golang/lint v0.0.0-20191125180803-fdd1cda4f05f
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20191210151939-1a1fef82734d
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191210023423-ac6580df4449
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20191217011448-c39ce2148d8e
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.14.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.5
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20191216205247-b31c10ee225f
	google.golang.org/grpc => github.com/grpc/grpc-go v1.25.1
)

require (
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.0
	github.com/lightstep/lightstep-tracer-go v0.20.0
	github.com/oklog/oklog v0.3.2
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin/zipkin-go-opentracing v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.5.1
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	google.golang.org/grpc v1.26.0
	sourcegraph.com/sourcegraph/appdash v0.0.0-20190731080439-ebfcffb1b5c0
)
