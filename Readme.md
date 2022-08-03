 fission fn create --name gomod1 --env go --entrypoint Handler --src go.mod --src go.sum --src fission-functions/echoHandler.go

fission httptrigger create --url /multiply --method GET --function multiply

argo submit -n argo --watch argo-worflows/echowhalesay.wf.yaml 

fission fn create --name multiply --env go --entrypoint MultiplyNumberHandler --src fission-functions/go.mod --src fission-functions/go.sum --src fission-functions/multiplyNumber.go 