FROM mongo:3.3

COPY transactions.json /transactions.json
COPY operations.json /operations.json

CMD mongoimport --host db --db pismo-transactions --collection Transactions --type json --file /transactions.json --jsonArray --drop && \
    mongoimport --host db --db pismo-transactions --collection OperationsTypes --type json --file /operations.json --jsonArray --drop
