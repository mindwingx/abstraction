package abstraction

// drivers abstraction to handle theme by interface across the API services

type AbstractCtx interface {
	Param(key string) string
	Query(key string) string
	GetHeader(key string) string
	BindJSON(obj interface{}) error
	ShouldBindJSON(obj interface{}) error
	JSON(code int, obj interface{})
	AbortWithStatusJSON(status int, data interface{})
	Abort()
	Next()
}

// handler function wrapper

type CustomCtxFunc func(ctx AbstractCtx)

// todo: wrap file upload, request validation, file download

type ApiService interface {
	InitApiService()
	StartHttp()
	RouteGroup
	Routes
}

type RouteGroup interface {
	RouteGroup(prefix string) RouteGroup
	NestedGroup(prefix string) RouteGroup
	Routes
}

type Routes interface {
	Get(path string, handler ...CustomCtxFunc)
	Post(path string, handler ...CustomCtxFunc)
	Put(path string, handler ...CustomCtxFunc)
	Delete(path string, handler ...CustomCtxFunc)
}

type RpcService interface {
	InitRpcService(rpcEntities []interface{})
	StartRpc()
	Caller(destinationPort string, rpcMethod string, args interface{}, reply interface{}) error
}

type CoreService interface {
	Stop()
}
