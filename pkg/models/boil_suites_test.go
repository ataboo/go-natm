// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Comments", testComments)
	t.Run("ProjectAssociations", testProjectAssociations)
	t.Run("Projects", testProjects)
	t.Run("TaskStatuses", testTaskStatuses)
	t.Run("Tasks", testTasks)
	t.Run("Users", testUsers)
	t.Run("WorkLogs", testWorkLogs)
}

func TestDelete(t *testing.T) {
	t.Run("Comments", testCommentsDelete)
	t.Run("ProjectAssociations", testProjectAssociationsDelete)
	t.Run("Projects", testProjectsDelete)
	t.Run("TaskStatuses", testTaskStatusesDelete)
	t.Run("Tasks", testTasksDelete)
	t.Run("Users", testUsersDelete)
	t.Run("WorkLogs", testWorkLogsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("ProjectAssociations", testProjectAssociationsQueryDeleteAll)
	t.Run("Projects", testProjectsQueryDeleteAll)
	t.Run("TaskStatuses", testTaskStatusesQueryDeleteAll)
	t.Run("Tasks", testTasksQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("WorkLogs", testWorkLogsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("ProjectAssociations", testProjectAssociationsSliceDeleteAll)
	t.Run("Projects", testProjectsSliceDeleteAll)
	t.Run("TaskStatuses", testTaskStatusesSliceDeleteAll)
	t.Run("Tasks", testTasksSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("WorkLogs", testWorkLogsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Comments", testCommentsExists)
	t.Run("ProjectAssociations", testProjectAssociationsExists)
	t.Run("Projects", testProjectsExists)
	t.Run("TaskStatuses", testTaskStatusesExists)
	t.Run("Tasks", testTasksExists)
	t.Run("Users", testUsersExists)
	t.Run("WorkLogs", testWorkLogsExists)
}

func TestFind(t *testing.T) {
	t.Run("Comments", testCommentsFind)
	t.Run("ProjectAssociations", testProjectAssociationsFind)
	t.Run("Projects", testProjectsFind)
	t.Run("TaskStatuses", testTaskStatusesFind)
	t.Run("Tasks", testTasksFind)
	t.Run("Users", testUsersFind)
	t.Run("WorkLogs", testWorkLogsFind)
}

func TestBind(t *testing.T) {
	t.Run("Comments", testCommentsBind)
	t.Run("ProjectAssociations", testProjectAssociationsBind)
	t.Run("Projects", testProjectsBind)
	t.Run("TaskStatuses", testTaskStatusesBind)
	t.Run("Tasks", testTasksBind)
	t.Run("Users", testUsersBind)
	t.Run("WorkLogs", testWorkLogsBind)
}

func TestOne(t *testing.T) {
	t.Run("Comments", testCommentsOne)
	t.Run("ProjectAssociations", testProjectAssociationsOne)
	t.Run("Projects", testProjectsOne)
	t.Run("TaskStatuses", testTaskStatusesOne)
	t.Run("Tasks", testTasksOne)
	t.Run("Users", testUsersOne)
	t.Run("WorkLogs", testWorkLogsOne)
}

func TestAll(t *testing.T) {
	t.Run("Comments", testCommentsAll)
	t.Run("ProjectAssociations", testProjectAssociationsAll)
	t.Run("Projects", testProjectsAll)
	t.Run("TaskStatuses", testTaskStatusesAll)
	t.Run("Tasks", testTasksAll)
	t.Run("Users", testUsersAll)
	t.Run("WorkLogs", testWorkLogsAll)
}

func TestCount(t *testing.T) {
	t.Run("Comments", testCommentsCount)
	t.Run("ProjectAssociations", testProjectAssociationsCount)
	t.Run("Projects", testProjectsCount)
	t.Run("TaskStatuses", testTaskStatusesCount)
	t.Run("Tasks", testTasksCount)
	t.Run("Users", testUsersCount)
	t.Run("WorkLogs", testWorkLogsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Comments", testCommentsHooks)
	t.Run("ProjectAssociations", testProjectAssociationsHooks)
	t.Run("Projects", testProjectsHooks)
	t.Run("TaskStatuses", testTaskStatusesHooks)
	t.Run("Tasks", testTasksHooks)
	t.Run("Users", testUsersHooks)
	t.Run("WorkLogs", testWorkLogsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("ProjectAssociations", testProjectAssociationsInsert)
	t.Run("ProjectAssociations", testProjectAssociationsInsertWhitelist)
	t.Run("Projects", testProjectsInsert)
	t.Run("Projects", testProjectsInsertWhitelist)
	t.Run("TaskStatuses", testTaskStatusesInsert)
	t.Run("TaskStatuses", testTaskStatusesInsertWhitelist)
	t.Run("Tasks", testTasksInsert)
	t.Run("Tasks", testTasksInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("WorkLogs", testWorkLogsInsert)
	t.Run("WorkLogs", testWorkLogsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CommentToTaskUsingTask", testCommentToOneTaskUsingTask)
	t.Run("CommentToUserUsingUser", testCommentToOneUserUsingUser)
	t.Run("ProjectAssociationToProjectUsingProject", testProjectAssociationToOneProjectUsingProject)
	t.Run("ProjectAssociationToUserUsingUser", testProjectAssociationToOneUserUsingUser)
	t.Run("TaskStatusToProjectUsingProject", testTaskStatusToOneProjectUsingProject)
	t.Run("TaskToUserUsingAssignee", testTaskToOneUserUsingAssignee)
	t.Run("TaskToTaskStatusUsingTaskStatus", testTaskToOneTaskStatusUsingTaskStatus)
	t.Run("WorkLogToTaskUsingTask", testWorkLogToOneTaskUsingTask)
	t.Run("WorkLogToUserUsingUser", testWorkLogToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ProjectToProjectAssociations", testProjectToManyProjectAssociations)
	t.Run("ProjectToTaskStatuses", testProjectToManyTaskStatuses)
	t.Run("TaskStatusToTasks", testTaskStatusToManyTasks)
	t.Run("TaskToComments", testTaskToManyComments)
	t.Run("TaskToWorkLogs", testTaskToManyWorkLogs)
	t.Run("UserToComments", testUserToManyComments)
	t.Run("UserToProjectAssociations", testUserToManyProjectAssociations)
	t.Run("UserToAssigneeTasks", testUserToManyAssigneeTasks)
	t.Run("UserToWorkLogs", testUserToManyWorkLogs)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CommentToTaskUsingComments", testCommentToOneSetOpTaskUsingTask)
	t.Run("CommentToUserUsingComments", testCommentToOneSetOpUserUsingUser)
	t.Run("ProjectAssociationToProjectUsingProjectAssociations", testProjectAssociationToOneSetOpProjectUsingProject)
	t.Run("ProjectAssociationToUserUsingProjectAssociations", testProjectAssociationToOneSetOpUserUsingUser)
	t.Run("TaskStatusToProjectUsingTaskStatuses", testTaskStatusToOneSetOpProjectUsingProject)
	t.Run("TaskToUserUsingAssigneeTasks", testTaskToOneSetOpUserUsingAssignee)
	t.Run("TaskToTaskStatusUsingTasks", testTaskToOneSetOpTaskStatusUsingTaskStatus)
	t.Run("WorkLogToTaskUsingWorkLogs", testWorkLogToOneSetOpTaskUsingTask)
	t.Run("WorkLogToUserUsingWorkLogs", testWorkLogToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("TaskToUserUsingAssigneeTasks", testTaskToOneRemoveOpUserUsingAssignee)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ProjectToProjectAssociations", testProjectToManyAddOpProjectAssociations)
	t.Run("ProjectToTaskStatuses", testProjectToManyAddOpTaskStatuses)
	t.Run("TaskStatusToTasks", testTaskStatusToManyAddOpTasks)
	t.Run("TaskToComments", testTaskToManyAddOpComments)
	t.Run("TaskToWorkLogs", testTaskToManyAddOpWorkLogs)
	t.Run("UserToComments", testUserToManyAddOpComments)
	t.Run("UserToProjectAssociations", testUserToManyAddOpProjectAssociations)
	t.Run("UserToAssigneeTasks", testUserToManyAddOpAssigneeTasks)
	t.Run("UserToWorkLogs", testUserToManyAddOpWorkLogs)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("UserToAssigneeTasks", testUserToManySetOpAssigneeTasks)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("UserToAssigneeTasks", testUserToManyRemoveOpAssigneeTasks)
}

func TestReload(t *testing.T) {
	t.Run("Comments", testCommentsReload)
	t.Run("ProjectAssociations", testProjectAssociationsReload)
	t.Run("Projects", testProjectsReload)
	t.Run("TaskStatuses", testTaskStatusesReload)
	t.Run("Tasks", testTasksReload)
	t.Run("Users", testUsersReload)
	t.Run("WorkLogs", testWorkLogsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Comments", testCommentsReloadAll)
	t.Run("ProjectAssociations", testProjectAssociationsReloadAll)
	t.Run("Projects", testProjectsReloadAll)
	t.Run("TaskStatuses", testTaskStatusesReloadAll)
	t.Run("Tasks", testTasksReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("WorkLogs", testWorkLogsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Comments", testCommentsSelect)
	t.Run("ProjectAssociations", testProjectAssociationsSelect)
	t.Run("Projects", testProjectsSelect)
	t.Run("TaskStatuses", testTaskStatusesSelect)
	t.Run("Tasks", testTasksSelect)
	t.Run("Users", testUsersSelect)
	t.Run("WorkLogs", testWorkLogsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Comments", testCommentsUpdate)
	t.Run("ProjectAssociations", testProjectAssociationsUpdate)
	t.Run("Projects", testProjectsUpdate)
	t.Run("TaskStatuses", testTaskStatusesUpdate)
	t.Run("Tasks", testTasksUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("WorkLogs", testWorkLogsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("ProjectAssociations", testProjectAssociationsSliceUpdateAll)
	t.Run("Projects", testProjectsSliceUpdateAll)
	t.Run("TaskStatuses", testTaskStatusesSliceUpdateAll)
	t.Run("Tasks", testTasksSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("WorkLogs", testWorkLogsSliceUpdateAll)
}