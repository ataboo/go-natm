ALTER TABLE project_associations DROP CONSTRAINT project_associations_project_id_fkey;
ALTER TABLE project_associations DROP CONSTRAINT project_associations_user_id_fkey;

ALTER TABLE tasks DROP CONSTRAINT tasks_assignee_id_fkey;
ALTER TABLE tasks DROP CONSTRAINT tasks_project_id_fkey;
ALTER TABLE tasks DROP CONSTRAINT tasks_task_status_id_fkey;

ALTER TABLE work_logs DROP CONSTRAINT work_logs_task_id_fkey;

DROP TABLE projects;
DROP TABLE task_statuses;
DROP TABLE tasks;
DROP TABLE project_associations;
DROP TABLE work_logs;
DROP TYPE task_type_enum;
DROP TYPE associations_enum;
