export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G = et2g
export EG2GO = eg2go

export GOARCH = amd64
export TARGET_ARG = x86_64
export GOETHOSINCLUDE = /usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE = /usr/lib64/go/pkg/linux/$(GOARCH)

export ETHOSROOT = server/rootfs
export MINIMALTDROOT = server/minimaltdfs

typeName = myRpc
TypeName = MyRpc
typeIndex = 2
serviceName = myRpcService
clientName = myRpcClient

.PHONY: all install clean
all: $(serviceName) $(clientName)


$(typeName).go: myRpc.t
	$(ETN2GO) . $(typeName) main $^

$(serviceName): $(serviceName).go $(typeName).go
	ethosGo $^

$(clientName): $(clientName).go $(typeName).go
	ethosGo $^

install: 
	rm -rf $(ETHOSROOT)/services/$(typeName) $(ETHOSROOT)/types/spec/$(typeName)
	ethosDirCreate $(ETHOSROOT)/types/spec/$(typeName)	$(ETHOSROOT)/types/spec/kernelTypes/HashValue all
	install -D $(typeName)/*				$(ETHOSROOT)/types/all
	install -D $(typeName)Index/*				$(ETHOSROOT)/types/spec/$(typeName)
	ethosDirCreate $(ETHOSROOT)/services/$(typeName)	$(ETHOSROOT)/types/spec/$(typeName)/$(TypeName) all
	install -D $(clientName) $(serviceName)			$(ETHOSROOT)/programs
	ethosStringEncode /programs/$(serviceName)>		$(ETHOSROOT)/etc/init/services/$(serviceName)
	ethosStringEncode /programs/$(clientName)>		$(ETHOSROOT)/etc/init/services/$(clientName)
	
clean:
	rm -rf $(typeName)/ $(typeIndex)/
	rm -f $(typeName).go
	rm -f $(serviceName)
	rm -f $(clientName)
	rm -f $(serviceName).goo.ethos
	rm-rf $(clientName).goo.ethos
