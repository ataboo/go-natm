CREATE TYPE task_type_enum as ENUM('Task', 'Problem', 'Admin');
CREATE TYPE associations_enum as ENUM('Owner', 'Writer', 'Reader');

CREATE TABLE projects (
   id UUID PRIMARY KEY,
   name VARCHAR (64) NOT NULL,
   abbreviation VARCHAR (64) NOT NULL,
   description VARCHAR (128) NOT NULL,
   active BOOLEAN NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE task_statuses (
   id UUID PRIMARY KEY,
   project_id UUID NOT NULL REFERENCES projects(id),
   name VARCHAR(64) NOT NULL,
   ordinal INTEGER NOT NULL,
   active BOOLEAN NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks (
   id UUID PRIMARY KEY,
   task_status_id UUID NOT NULL REFERENCES task_statuses(id),
   number INTEGER NOT NULL,
   assignee_id UUID REFERENCES users(id),
   ordinal INTEGER NOT NULL,
   title VARCHAR(64) NOT NULL,
   estimate INTEGER,
   description TEXT NOT NULL,
   task_type task_type_enum NOT NULL,
   active BOOLEAN NOT NULL DEFAULT TRUE,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE project_associations (
   id UUID PRIMARY KEY,
   project_id UUID NOT NULL REFERENCES projects(id),
   user_id UUID NOT NULL REFERENCES users(id),
   association associations_enum NOT NULL
);

CREATE TABLE work_logs (
   id UUID PRIMARY KEY,
   task_id UUID NOT NULL REFERENCES tasks(id),
   user_id UUID NOT NULL REFERENCES users(id),
   start_time TIMESTAMP NOT NULL,
   end_time TIMESTAMP NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);