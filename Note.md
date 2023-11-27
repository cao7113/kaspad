EnableNonNativeSubnetworks=true

2023-11-27 16:32:14.231 [CRT] RPCS: Exiting: Fatal error in goroutine `routerInitializer-handleIncomingMessages 6`: We currently don't support non native subnetworks

/Users/rj/.asdf/installs/golang/1.20.1/go/src/runtime/panic.go:884 +0x1f4
github.com/kaspanet/kaspad/domain/miningmanager/blocktemplatebuilder.(*blockTemplateBuilder).BuildBlockTemplate(0x140001c5da0, 0x20?)
/Users/rj/dev/kaspa-net/kaspad/domain/miningmanager/blocktemplatebuilder/blocktemplatebuilder.go:127 +0x690
github.com/kaspanet/kaspad/domain/miningmanager.(*miningManager).GetBlockTemplate(0x14000181860, 0x14015387520)
/Users/rj/dev/kaspa-net/kaspad/domain/miningmanager/miningmanager.go:70 +0x328
github.com/kaspanet/kaspad/app/rpc/rpchandlers.HandleGetBlockTemplate(0x140080db3b0, 0x140002fa180?, {0x100c29ff0?, 0x14015380fc0})
/Users/rj/dev/kaspa-net/kaspad/app/rpc/rpchandlers/get_block_template.go:32 +0x21c
github.com/kaspanet/kaspad/app/rpc.(\*Manager).handleIncomingMessages(0x140081b9218, 0x140081da5a0, 0x10?)
