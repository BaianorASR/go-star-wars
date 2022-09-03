FROM mongo

ARG MONGO_PORT
ARG MONGO_INITDB_ROOT_USERNAME
ARG MONGO_INITDB_ROOT_PASSWORD

COPY init.json /init.json

RUN echo "🌱🌱🌱 Seeding database with data from init.json 🌱🌱🌱"

CMD mongoimport --host mongo --port $MONGO_PORT --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --authenticationDatabase admin --db go-americanas-teste --collection planets --drop --file /init.json --jsonArray

RUN echo "🌱🌱🌱 End Seed 🌱🌱🌱"