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

typeName1 = myRpc
TypeName1 = MyRpc
typeIndex = 1
serviceName = myRpcService
clientName = myRpcClient

.PHONY: all install clean
all: $(serviceName) $(clientName)


$(typeName1).go: myRpc.t
	$(ETN2GO) . $(typeName1) main $^


$(serviceName): $(serviceName).go $(typeName1).go
	ethosGo $^

$(clientName): $(clientName).go $(typeName1).go
	ethosGo $^

install:
	sudo rm -rf server/
	(ethosParams server && cd server/ && ethosBuilder && minimaltdBuilder)	
	ethosDirCreate $(ETHOSROOT)/types/spec/$(typeName1)	$(ETHOSROOT)/types/spec/kernelTypes/HashValue all
	install -D $(typeName1)/*				$(ETHOSROOT)/types/all
	install -D $(typeName1)Index/*				$(ETHOSROOT)/types/spec/$(typeName1)
	ethosDirCreate $(ETHOSROOT)/services/$(typeName1)	$(ETHOSROOT)/types/spec/$(typeName1)/$(TypeName1) all
	install -D $(clientName) $(serviceName)			$(ETHOSROOT)/programs
	ethosStringEncode /programs/$(serviceName)>		$(ETHOSROOT)/etc/init/services/$(serviceName)
	ethosStringEncode /programs/$(clientName)>		$(ETHOSROOT)/etc/init/services/$(clientName)
	sudo mkdir $(ETHOSROOT)/user/nobody
	sudo mkdir $(ETHOSROOT)/user/printService
	sudo cp printFile $(ETHOSROOT)/user/nobody/printFile
	
clean:
	rm -rf $(typeName1)/ $(typeIndex)/
	rm -f $(typeName1).go
	rm -f $(serviceName)
	rm -f $(clientName)
	rm -f $(serviceName).goo.ethos
	rm -rf $(clientName).goo.ethos
