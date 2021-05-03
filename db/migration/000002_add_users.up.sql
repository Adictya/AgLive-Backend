CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL
);

ALTER TABLE "streams" ADD "streamer" varchar;

ALTER TABLE "streams" ADD FOREIGN KEY ("streamer") REFERENCES "users" ("username");
