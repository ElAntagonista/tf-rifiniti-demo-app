# tf-rifiniti-modules

# About 
A simple golang web app to be used as a deployment artifact for AWS Fargate

# Docker build strategy
A multi-stage build approach is used to reduce the size of the final image drastically.
The runtime image size is just 13MB.
