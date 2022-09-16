# Retry Scenario

The example show how Argo workflows can be used to handle `failure` while running a Fission Workflow.

### Steps

- Run `spec.sh` - this will create 2 fission functions `set-first-name` and `get-full-name`. And on HTTPTrigger - `postfirstname`
- Submit the Argo workflow - `argo submit --watch retry-backoff.yaml` This will start the workflow execution. And will fail on second task i.e.- `getfullname` because we have not created the trigger for same.
- Open a new terminal and create the trigger using command - `fission httptrigger create --name getfullname --url /getfullnamehandler --method GET --function get-full-name`
- On successful creation of HTTPTrigger the workflow tasks completes successfully.
  
### The above examples proves that -

- We can retry the failed stage for any number of time. Also we can configure the backoff time i.e the time after which retry will be attempted.
- The output from earlier stages stays intact in case of failure.
