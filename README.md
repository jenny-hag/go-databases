# Databases
In this project I will explore working with databases in go-lang

## sqlite

### Create database
```
CREATE TABLE `bucket_list` (
      `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
      `place` VARCHAR(64) NULL,
      `country` VARCHAR(64) NULL,
      `photo` VARCHAR(64) NULL,
      `visited` BOOL DEFAULT FALSE,
      `created_at` DATE NULL
  );
```
