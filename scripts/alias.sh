project_name='gin-api-template-development'

alias docker-compose="docker-compose -f deployments/development/docker-compose.yml -p $project_name"

alias build="docker-compose build"
alias up="docker-compose up"
alias stop="docker-compose stop"
alias down="docker-compose down"

alias run="docker-compose run --rm"
alias exec="docker-compose exec"
