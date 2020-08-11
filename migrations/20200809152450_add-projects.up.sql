CREATE TABLE projects (
   id UUID PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   active BOOLEAN NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE task_states (
   id UUID PRIMARY KEY,
   project_id UUID NOT NULL REFERENCES projects(id),
   name VARCHAR(50) NOT NULL,
   position INTEGER NOT NULL,
   active BOOLEAN NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks (
   id UUID PRIMARY KEY,
   project_id UUID NOT NULL REFERENCES projects(id),
   task_state_id UUID NOT NULL REFERENCES task_states(id),
   name VARCHAR(50) NOT NULL,
   description TEXT NOT NULL
);

CREATE TABLE task_labels (
   id UUID PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   project_id UUID NOT NULL REFERENCES projects(id)
);

CREATE TABLE project_users (
   id UUID PRIMARY KEY,
   project_id UUID NOT NULL REFERENCES projects(id),
   user_id UUID NOT NULL REFERENCES users(id),
   can_read BOOLEAN NOT NULL,
   can_write BOOLEAN NOT NULL,
   is_owner BOOLEAN NOT NULL
);

CREATE TABLE work_logs (
   id UUID PRIMARY KEY,
   task_id UUID NOT NULL REFERENCES tasks(id),
   user_id UUID NOT NULL REFERENCES users(id),
   start_time TIMESTAMP NOT NULL,
   end_time TIMESTAMP NOT NULL
);