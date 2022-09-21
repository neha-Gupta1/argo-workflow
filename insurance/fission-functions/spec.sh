fission function create --name set-first-name --env go --src fission-function/postFirstName.go --entrypoint PostFirstNameHandler
fission httptrigger create --name postfirstname --url /postfirstname --method GET --function set-first-name

fission function create --name set-last-name --env go --src fission-function/postLastName.go --entrypoint PostLastNameHandler
fission httptrigger create --name postlastname --url /postlastname --method GET --function set-last-name

fission function create --name get-full-name --env go --src fission-functions/getFullName.go --entrypoint GetFullNameHandler
fission httptrigger create --name getfullname --url /getfullnamehandler --method GET --function get-full-name

fission function create --name set-age --env go --src fission-functions/getAge.go --entrypoint PostAgeHandler
fission httptrigger create --name postage --url /setage --method GET --function set-age

fission function create --name set-salary --env go --src fission-functions/getMonthlySalary.go --entrypoint PostSalaryHandler
fission httptrigger create --name postsalary --url /setmonthlysalary --method GET --function set-salary

fission function create --name lowriskinsurance --env go --src fission-functions/calculateEligibililty.go --entrypoint GetLowRiskInsuranceHandler
fission httptrigger create --name getlowriskinsurance --url /getLowRiskEligibility --method GET --function lowriskinsurance

fission function create --name highriskinsurance --env go --src fission-functions/calculateEligibililty.go --entrypoint GetHighRiskInsuranceHandler
fission httptrigger create --name gethighriskinsurance --url /getHighRiskEligibility --method GET --function highriskinsurance
