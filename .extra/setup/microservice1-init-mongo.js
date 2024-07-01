function seed(dbName, user, password) {
  db = db.getSiblingDB(dbName);
  db.createUser({
    user: user,
    pwd: password,
    roles: [{ role: "readWrite", db: dbName }],
  });
}

seed("microservice1-prod-db", "microservice1-prod-db-user", "changeit");
seed("microservice1-dev-db", "microservice1-dev-db-user", "changeit");
seed("microservice1-test-db", "microservice1-test-db-user", "changeit");
