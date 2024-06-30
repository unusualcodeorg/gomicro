function seed(dbName, user, password) {
  db = db.getSiblingDB(dbName);
  db.createUser({
    user: user,
    pwd: password,
    roles: [{ role: "readWrite", db: dbName }],
  });
}

seed("article-service-prod-db", "article-service-prod-db-user", "changeit");
seed("article-service-dev-db", "article-service-dev-db-user", "changeit");
seed("article-service-test-db", "article-service-test-db-user", "changeit");
