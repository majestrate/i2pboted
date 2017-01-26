package bote

type RPC struct {
	router *Router
}

type RPCArgs struct {
	Graceful bool
}

type RPCResult struct {
	Code int
}

func (rpc *RPC) Shutdown(args *RPCArgs, res *RPCResult) error {
	if args.Graceful {
		rpc.router.GracefulClose()
	} else {
		rpc.router.Close()
	}
	res = &RPCResult{
		Code: 0,
	}
	return nil
}
