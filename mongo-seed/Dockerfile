FROM mongo

ARG MONGO_INITDB_ROOT_USERNAME
ARG MONGO_INITDB_ROOT_PASSWORD

COPY init.json /init.json

CMD mongoimport --host mongo --port 27017 --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --authenticationDatabase admin --db go-americanas-teste --collection planets --drop --file /init.json --jsonArray