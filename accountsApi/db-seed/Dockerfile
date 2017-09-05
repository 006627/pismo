FROM mongo:3.3

COPY account.json /account.json

CMD mongoimport --host db --db pismo-accounts --collection Accounts --type json --file /account.json --jsonArray --drop