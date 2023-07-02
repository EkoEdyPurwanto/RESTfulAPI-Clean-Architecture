step by step

1. create a project folder
2. initialize git
3. initialize go module
4. create api specification
5. install framework & driver (postgreSQL, Echo Framework, GO playground/validator)
6. create database & table
7. implement models/domain/entity
8. implement repository (interface & impl from interface)
9. implement useCase and struct for requests and responses put in models directory (interface & impl from interface)
10. implement category validation in useCase use https://github.com/go-playground/validator
11. implement delivery (interface & impl from interface)
12. generate constructor in delivery, useCase and repository
13. install spf13/viper for configuration
14. implement router 
15. implement open database and create env
16. implement server
17. implement test.http if there is an error fix it
18. implement error handling 
19. implement Dockerfile
20. implement docker-compose.yaml
21. implement Detecting Security Vulnerabilities In Dependencies (go.mod) by doing 'osv-scanner --lockfile=go.mod' use https://github.com/google/osv-scanner 
22. implement Fixing Bug or Search Code That is Not Common by doing 'golangci-lint run' use https://github.com/golangci/golangci-lint