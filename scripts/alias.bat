DOSKEY build=docker-compose -f deployments\development\docker-compose.yml -p gin-api-template-development build $*
DOSKEY up=docker-compose -f deployments\development\docker-compose.yml -p gin-api-template-development up $*
DOSKEY stop=docker-compose -f deployments\development\docker-compose.yml -p gin-api-template-development stop $*
DOSKEY down=docker-compose -f deployments\development\docker-compose.yml -p gin-api-template-development down $*
DOSKEY run=docker-compose -f deployments\development\docker-compose.yml -p gin-api-template-development run --rm $*
DOSKEY exec=docker-compose -f deployments\development\docker-compose.yml -p gin-api-template-development exec $*
