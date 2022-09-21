# DAG scenario

This is a simple example explaining how argo workflows can be used to create Workflows in Fission.

![insuranceCalc](https://user-images.githubusercontent.com/23420056/190299151-bb9fcaad-4a6f-4865-aa3e-00ea61fec81e.png)

## Steps

- Run - `./fission-functions/spec.sh`
- Run- argo submit --watch insuranceDag.yaml
For simplicity, we have added some default values you can go ahead and play with the inputs.
