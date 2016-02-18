==>Runtime and SDK:

Dockerfile https://github.com/uzmo/docker/blob/master/OpenUI5_RT_SDK

Supported Tags:1.28.28
OpenUI5 Runtime and SDK for version 1.28.28
Version 1.28.28 is the only stable maintenance release by SAP.
SAP Fiori Design Guideline(https://experience.sap.com/fiori-design/) is based on the controls available with SAPUI5 1.28.
This image is build from Alpine which is quite smaller than other Linux.
The content is hosted on a static file server powered by Golang.

Run the container:
docker run -d -p 5000:9090 uzmo/openui5:1.28.28

Access the libraries:
Runtime: http://<docker machine IP>:5000
SDK: http://<docker machine IP>:5000/sdk
