function seed(dbName, user, password) {
  db = db.getSiblingDB(dbName);
  db.createUser({
    user: user,
    pwd: password,
    roles: [{ role: "readWrite", db: dbName }],
  });
}

seed("blog-prod-db", "blog-prod-db-user", "changeit");
seed("blog-dev-db", "blog-dev-db-user", "changeit");
seed("blog-test-db", "blog-test-db-user", "changeit");
