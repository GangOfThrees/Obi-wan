BEGIN TRANSACTION;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE OrganisationStatus AS ENUM (
  'inactive',
  'active'
);

CREATE SCHEMA obiwan;

CREATE TABLE obiwan.users (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  firstName TEXT NOT NULL,
  lastName TEXT NOT NULL,
  dob DATE NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP,
  
  CONSTRAINT uniqUsersEmail UNIQUE (email)
);

CREATE TABLE obiwan.organisations (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  name TEXT NOT NULL,
  description TEXT,
  address TEXT,
  city TEXT,
  state TEXT,
  zip TEXT,
  phone TEXT,
  website TEXT,
  email TEXT NOT NULL,
  status OrganisationStatus NOT NULL DEFAULT 'inactive',
  preferredBotService TEXT,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP
);

CREATE TABLE obiwan.organisationProfiles (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  organisationId UUID NOT NULL,
  initialQuestions JSON NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP,
  
  CONSTRAINT fkOrganisationProfilesOrganisationId FOREIGN KEY (organisationId) REFERENCES obiwan.organisations(id) ON DELETE RESTRICT
);

CREATE TABLE obiwan.projects (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  organisationId UUID NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP
);

CREATE TABLE obiwan.projectsOrganisations (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  projectId UUID NOT NULL,
  organisationId UUID NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP,
  
  CONSTRAINT fkProjectsOrganisationsProjectId FOREIGN KEY (projectId) REFERENCES obiwan.projects(id) ON DELETE RESTRICT,
  CONSTRAINT fkProjectsOrganisationsOrganisationId FOREIGN KEY (organisationId) REFERENCES obiwan.organisations(id) ON DELETE RESTRICT,
  CONSTRAINT uniqProjectsOrganisationsProjectOrganisation UNIQUE (projectId, organisationId)
);

CREATE TABLE obiwan.usersOrganisations (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  organisationId UUID NOT NULL,
  userId UUID NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP,
  
  CONSTRAINT fkUsersOrganisationsOrganisationId FOREIGN KEY (organisationId) REFERENCES obiwan.organisations(id) ON DELETE RESTRICT,
  CONSTRAINT fkUsersOrganisationsUserId FOREIGN KEY (userId) REFERENCES obiwan.users(id) ON DELETE RESTRICT,
  CONSTRAINT uniqUsersOrganisationsOrganisationUser UNIQUE (organisationId, userId)
);

CREATE TABLE obiwan.apiKeys (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  projectId UUID NOT NULL,
  key TEXT NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deletedAt TIMESTAMP,

  CONSTRAINT fkApiKeysProjectId FOREIGN KEY (projectId) REFERENCES obiwan.projects(id) ON DELETE RESTRICT
);

END TRANSACTION;
