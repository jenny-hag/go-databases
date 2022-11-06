# Databases
In this project I will explore working with databases in go-lang

### Create database
```
-- SQLite
CREATE TABLE `bucket_list` (
      `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
      `place` VARCHAR(64) NULL,
      `country` VARCHAR(64) NULL,
      `photo` VARCHAR(64) NULL,
      `visited` BOOL DEFAULT FALSE,
      `created_at` DATE NULL
  );
```

```
-- MySQL
CREATE TABLE `bucket_list` (
      `uid` INTEGER PRIMARY KEY AUTO_INCREMENT,
      `place` VARCHAR(64) NULL,
      `country` VARCHAR(64) NULL,
      `photo` VARCHAR(256) NULL,
      `visited` BOOLEAN DEFAULT FALSE,
      `created_at` DATE NULL
  );
```

```
-- PostgreSQL
CREATE TABLE "bucket_list" (
      "uid" SERIAL PRIMARY KEY,
      "place" VARCHAR(64) NULL,
      "country" VARCHAR(64) NULL,
      "photo" VARCHAR(256) NULL,
      "visited" BOOLEAN DEFAULT FALSE,
      "created_at" DATE NULL
  );
  ``
