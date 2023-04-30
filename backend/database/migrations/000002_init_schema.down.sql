DROP TABLE IF EXISTS "User" CASCADE;
DROP TABLE IF EXISTS "QA" CASCADE;
DROP TABLE IF EXISTS "Chat" CASCADE;
DROP TABLE IF EXISTS "Message" CASCADE;

DROP FUNCTION IF EXISTS "message_no_increment" CASCADE;

CREATE TABLE "User" (
  "username" varchar(40) PRIMARY KEY
);

CREATE TABLE "QA" (
  "qa_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "question" varchar(255) NOT NULL,
  "answer" varchar(255) NOT NULL
);

CREATE TABLE "Chat" (
  "chat_id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "username" varchar(40) NOT NULL
);

CREATE TABLE "ChatQA" (
  "chat_id" int,
  "no" int NOT NULL,
  "question" varchar(255) NOT NULL,
  "answer" varchar(255) NOT NULL
);

ALTER TABLE "Chat" ADD FOREIGN KEY ("username") REFERENCES "User" ("username");

ALTER TABLE "ChatQA" ADD FOREIGN KEY ("chat_id") REFERENCES "Chat" ("chat_id");

CREATE OR REPLACE FUNCTION chatqa_no_increment() RETURNS TRIGGER AS 
$$
BEGIN
  SELECT COALESCE(MAX("no"), 0) + 1 INTO NEW."no" FROM "ChatQA" WHERE "chat_id" = NEW."chat_id";
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for inserting Weak Entity ChatQA
CREATE OR REPLACE TRIGGER ChatQA BEFORE INSERT ON "ChatQA"
FOR EACH ROW
EXECUTE FUNCTION chatqa_no_increment();