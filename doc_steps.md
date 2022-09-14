 fission fn create --name gomod1 --env go --entrypoint Handler --src go.mod --src go.sum --src fission-functions/echoHandler.go

fission httptrigger create --url /add --name add --method GET --function add

argo submit -n argo --watch argo-worflows/echowhalesay.wf.yaml 

fission fn create --name multiply --env go --entrypoint AddNumberHandler --src fission-functions/go.mod --src fission-functions/go.sum --src fission-functions/multiplyNumber.go 
fission httptrigger create --url /multiply --name add --method GET --function multiply

todo: 
implement -
- failwhale
- noop

- To print something we need to run the container and cant call a fission function because we are restricted to http triggers.
- argowf allows us to store output on a location. This feature could be used to view output.Maybe we could use it as place of logging or something
- fissionwf had internal functions. But here I see to perform any actions they mostly write a inline script maybe in python or a shell script.

- Inline functions- argowf has script which is wrapper of container to where we can perform inline functions