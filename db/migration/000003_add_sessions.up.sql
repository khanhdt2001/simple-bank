CREATE TABLE "sessions" (
                         "id" uuid PRIMARY KEY,
                         "username" varchar not null ,
                         "refresh_token" varchar NOT NULL,
                         "user_agent" varchar NOT NULL,
                         "client_ip" varchar NOT NULL,
                         "is_blocked" boolean NOT NULL DEFAULT FALSE,
                         "created_at" timestamptz NOT NULL DEFAULT (now()),
                         "expires_at" timestamptz NOT NULL
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");


