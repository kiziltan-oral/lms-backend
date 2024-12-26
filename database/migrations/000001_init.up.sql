CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- BEGIN SYSTEMUSERS
CREATE TABLE "SystemUsers" (
    "Id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "Name" varchar(50) NOT NULL,
    "Surname" varchar(50) NOT NULL,
    "Email" varchar(255) NOT NULL,
    "Password" varchar(64) NOT NULL,
    "PasswordSalt" varchar(15) NOT NULL,
    "IsActive" boolean NOT NULL DEFAULT true
);

-- Create an index on Email and IsActive columns
CREATE INDEX idx_systemusers_email_isactive ON "SystemUsers" ("Email", "IsActive");

-- Create a unique index on Email where IsActive is true
CREATE UNIQUE INDEX uix_systemusers_email_active ON "SystemUsers" ("Email")
WHERE "IsActive" = true;

ALTER TABLE "SystemUsers" OWNER TO postgres;
-- END SYSTEMUSERS

-- BEGIN SYSTEMUSERSETTINGS
CREATE TABLE "SystemUserSettings" (
    "Id" serial PRIMARY KEY,
    "SystemUserId" uuid NOT NULL,
    "Key" varchar(50) NOT NULL,
    "Value" varchar(200) NOT NULL,
    "Description" varchar(200),
    CONSTRAINT fk_systemusersettings_systemuserid FOREIGN KEY ("SystemUserId") REFERENCES "SystemUsers" ("Id") ON DELETE CASCADE
);

CREATE INDEX idx_systemusersettings_systemuserid ON "SystemUserSettings" ("SystemUserId");

ALTER TABLE "SystemUserSettings" OWNER TO postgres;
-- END SYSTEMUSERSETTINGS

-- BEGIN CLIENTS
CREATE TABLE "Clients" (
    "Id" serial PRIMARY KEY,
    "ShortTitle" varchar(50) NOT NULL,
    "Title" varchar(200) NOT NULL,
    "Notes" text,
    "IsActive" boolean NOT NULL DEFAULT true
);

CREATE UNIQUE INDEX uix_clients_shorttitle_active ON "Clients" ("ShortTitle")
WHERE "IsActive" = true;

ALTER TABLE "Clients" OWNER TO postgres;
-- END CLIENTS

-- BEGIN CLIENTPROJECTS
CREATE TABLE "ClientProjects" (
    "Id" serial PRIMARY KEY,
    "ClientId" integer NOT NULL,
    "Name" varchar(100) NOT NULL,
    "IsActive" boolean NOT NULL DEFAULT true,
    CONSTRAINT fk_clientprojects_clientid FOREIGN KEY ("ClientId") REFERENCES "Clients" ("Id") ON DELETE CASCADE
);

CREATE INDEX idx_clientprojects_clientid_active ON "ClientProjects" ("ClientId", "IsActive");

ALTER TABLE "ClientProjects" OWNER TO postgres;
-- END CLIENTPROJECTS

-- BEGIN TIMINGS
CREATE TABLE "Timings" (
    "Id" serial PRIMARY KEY,
    "ClientProjectId" integer NOT NULL,
    "SystemUserId" uuid NOT NULL,
    "Title" varchar(100) NOT NULL,
    "Description" text,
    "StartDateTime" timestamptz NOT NULL,
    "EndDateTime" timestamptz NOT NULL,
    "Status" varchar(50) NOT NULL,
    CONSTRAINT fk_timings_clientprojectid FOREIGN KEY ("ClientProjectId") REFERENCES "ClientProjects" ("Id") ON DELETE CASCADE,
    CONSTRAINT fk_timings_systemuserid FOREIGN KEY ("SystemUserId") REFERENCES "SystemUsers" ("Id") ON DELETE CASCADE
);

CREATE INDEX idx_timings_clientprojectid_status ON "Timings" ("ClientProjectId", "Status");

ALTER TABLE "Timings" OWNER TO postgres;
-- END TIMINGS
