CREATE TABLE IF NOT EXISTS "people" (
    "id" INTEGER not NULL,
    "first_name" TEXT not NULL,
    "last_name" TEXT not NULL,
    "email" TEXT not NULL,
    "ip_address" TEXT not NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
);