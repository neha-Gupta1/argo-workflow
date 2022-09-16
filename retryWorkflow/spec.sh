fission function create --name set-first-name --env go --src fissionFunction/postFirstName.go --entrypoint PostFirstNameHandler
fission httptrigger create --name postfirstname --url /postfirstname --method GET --function set-first-name

fission function create --name get-full-name --env go --src fissionFunctions/getFullName.go --entrypoint GetFullNameHandler

# command to run after retry scenario recreated
# fission httptrigger create --name getfullname --url /getfullnamehandler --method GET --function get-full-name

