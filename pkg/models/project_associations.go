// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ProjectAssociation is an object representing the database table.
type ProjectAssociation struct {
	ID          string `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProjectID   string `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	Email       string `boil:"email" json:"email" toml:"email" yaml:"email"`
	Association string `boil:"association" json:"association" toml:"association" yaml:"association"`

	R *projectAssociationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L projectAssociationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProjectAssociationColumns = struct {
	ID          string
	ProjectID   string
	Email       string
	Association string
}{
	ID:          "id",
	ProjectID:   "project_id",
	Email:       "email",
	Association: "association",
}

// Generated where

var ProjectAssociationWhere = struct {
	ID          whereHelperstring
	ProjectID   whereHelperstring
	Email       whereHelperstring
	Association whereHelperstring
}{
	ID:          whereHelperstring{field: "\"project_associations\".\"id\""},
	ProjectID:   whereHelperstring{field: "\"project_associations\".\"project_id\""},
	Email:       whereHelperstring{field: "\"project_associations\".\"email\""},
	Association: whereHelperstring{field: "\"project_associations\".\"association\""},
}

// ProjectAssociationRels is where relationship names are stored.
var ProjectAssociationRels = struct {
	Project string
}{
	Project: "Project",
}

// projectAssociationR is where relationships are stored.
type projectAssociationR struct {
	Project *Project `boil:"Project" json:"Project" toml:"Project" yaml:"Project"`
}

// NewStruct creates a new relationship struct
func (*projectAssociationR) NewStruct() *projectAssociationR {
	return &projectAssociationR{}
}

// projectAssociationL is where Load methods for each relationship are stored.
type projectAssociationL struct{}

var (
	projectAssociationAllColumns            = []string{"id", "project_id", "email", "association"}
	projectAssociationColumnsWithoutDefault = []string{"id", "project_id", "email", "association"}
	projectAssociationColumnsWithDefault    = []string{}
	projectAssociationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ProjectAssociationSlice is an alias for a slice of pointers to ProjectAssociation.
	// This should generally be used opposed to []ProjectAssociation.
	ProjectAssociationSlice []*ProjectAssociation
	// ProjectAssociationHook is the signature for custom ProjectAssociation hook methods
	ProjectAssociationHook func(context.Context, boil.ContextExecutor, *ProjectAssociation) error

	projectAssociationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	projectAssociationType                 = reflect.TypeOf(&ProjectAssociation{})
	projectAssociationMapping              = queries.MakeStructMapping(projectAssociationType)
	projectAssociationPrimaryKeyMapping, _ = queries.BindMapping(projectAssociationType, projectAssociationMapping, projectAssociationPrimaryKeyColumns)
	projectAssociationInsertCacheMut       sync.RWMutex
	projectAssociationInsertCache          = make(map[string]insertCache)
	projectAssociationUpdateCacheMut       sync.RWMutex
	projectAssociationUpdateCache          = make(map[string]updateCache)
	projectAssociationUpsertCacheMut       sync.RWMutex
	projectAssociationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var projectAssociationBeforeInsertHooks []ProjectAssociationHook
var projectAssociationBeforeUpdateHooks []ProjectAssociationHook
var projectAssociationBeforeDeleteHooks []ProjectAssociationHook
var projectAssociationBeforeUpsertHooks []ProjectAssociationHook

var projectAssociationAfterInsertHooks []ProjectAssociationHook
var projectAssociationAfterSelectHooks []ProjectAssociationHook
var projectAssociationAfterUpdateHooks []ProjectAssociationHook
var projectAssociationAfterDeleteHooks []ProjectAssociationHook
var projectAssociationAfterUpsertHooks []ProjectAssociationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProjectAssociation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProjectAssociation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProjectAssociation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProjectAssociation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProjectAssociation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProjectAssociation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProjectAssociation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProjectAssociation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProjectAssociation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range projectAssociationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProjectAssociationHook registers your hook function for all future operations.
func AddProjectAssociationHook(hookPoint boil.HookPoint, projectAssociationHook ProjectAssociationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		projectAssociationBeforeInsertHooks = append(projectAssociationBeforeInsertHooks, projectAssociationHook)
	case boil.BeforeUpdateHook:
		projectAssociationBeforeUpdateHooks = append(projectAssociationBeforeUpdateHooks, projectAssociationHook)
	case boil.BeforeDeleteHook:
		projectAssociationBeforeDeleteHooks = append(projectAssociationBeforeDeleteHooks, projectAssociationHook)
	case boil.BeforeUpsertHook:
		projectAssociationBeforeUpsertHooks = append(projectAssociationBeforeUpsertHooks, projectAssociationHook)
	case boil.AfterInsertHook:
		projectAssociationAfterInsertHooks = append(projectAssociationAfterInsertHooks, projectAssociationHook)
	case boil.AfterSelectHook:
		projectAssociationAfterSelectHooks = append(projectAssociationAfterSelectHooks, projectAssociationHook)
	case boil.AfterUpdateHook:
		projectAssociationAfterUpdateHooks = append(projectAssociationAfterUpdateHooks, projectAssociationHook)
	case boil.AfterDeleteHook:
		projectAssociationAfterDeleteHooks = append(projectAssociationAfterDeleteHooks, projectAssociationHook)
	case boil.AfterUpsertHook:
		projectAssociationAfterUpsertHooks = append(projectAssociationAfterUpsertHooks, projectAssociationHook)
	}
}

// One returns a single projectAssociation record from the query.
func (q projectAssociationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProjectAssociation, error) {
	o := &ProjectAssociation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for project_associations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProjectAssociation records from the query.
func (q projectAssociationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProjectAssociationSlice, error) {
	var o []*ProjectAssociation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProjectAssociation slice")
	}

	if len(projectAssociationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProjectAssociation records in the query.
func (q projectAssociationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count project_associations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q projectAssociationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if project_associations exists")
	}

	return count > 0, nil
}

// Project pointed to by the foreign key.
func (o *ProjectAssociation) Project(mods ...qm.QueryMod) projectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	query := Projects(queryMods...)
	queries.SetFrom(query.Query, "\"projects\"")

	return query
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (projectAssociationL) LoadProject(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProjectAssociation interface{}, mods queries.Applicator) error {
	var slice []*ProjectAssociation
	var object *ProjectAssociation

	if singular {
		object = maybeProjectAssociation.(*ProjectAssociation)
	} else {
		slice = *maybeProjectAssociation.(*[]*ProjectAssociation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &projectAssociationR{}
		}
		args = append(args, object.ProjectID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &projectAssociationR{}
			}

			for _, a := range args {
				if a == obj.ProjectID {
					continue Outer
				}
			}

			args = append(args, obj.ProjectID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projects`),
		qm.WhereIn(`projects.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Project")
	}

	var resultSlice []*Project
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Project")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for projects")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projects")
	}

	if len(projectAssociationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Project = foreign
		if foreign.R == nil {
			foreign.R = &projectR{}
		}
		foreign.R.ProjectAssociations = append(foreign.R.ProjectAssociations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProjectID == foreign.ID {
				local.R.Project = foreign
				if foreign.R == nil {
					foreign.R = &projectR{}
				}
				foreign.R.ProjectAssociations = append(foreign.R.ProjectAssociations, local)
				break
			}
		}
	}

	return nil
}

// SetProject of the projectAssociation to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.ProjectAssociations.
func (o *ProjectAssociation) SetProject(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Project) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"project_associations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
		strmangle.WhereClause("\"", "\"", 2, projectAssociationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProjectID = related.ID
	if o.R == nil {
		o.R = &projectAssociationR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &projectR{
			ProjectAssociations: ProjectAssociationSlice{o},
		}
	} else {
		related.R.ProjectAssociations = append(related.R.ProjectAssociations, o)
	}

	return nil
}

// ProjectAssociations retrieves all the records using an executor.
func ProjectAssociations(mods ...qm.QueryMod) projectAssociationQuery {
	mods = append(mods, qm.From("\"project_associations\""))
	return projectAssociationQuery{NewQuery(mods...)}
}

// FindProjectAssociation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProjectAssociation(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*ProjectAssociation, error) {
	projectAssociationObj := &ProjectAssociation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"project_associations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, projectAssociationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from project_associations")
	}

	return projectAssociationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProjectAssociation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project_associations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectAssociationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	projectAssociationInsertCacheMut.RLock()
	cache, cached := projectAssociationInsertCache[key]
	projectAssociationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			projectAssociationAllColumns,
			projectAssociationColumnsWithDefault,
			projectAssociationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(projectAssociationType, projectAssociationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(projectAssociationType, projectAssociationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"project_associations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"project_associations\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into project_associations")
	}

	if !cached {
		projectAssociationInsertCacheMut.Lock()
		projectAssociationInsertCache[key] = cache
		projectAssociationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProjectAssociation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProjectAssociation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	projectAssociationUpdateCacheMut.RLock()
	cache, cached := projectAssociationUpdateCache[key]
	projectAssociationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			projectAssociationAllColumns,
			projectAssociationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update project_associations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"project_associations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, projectAssociationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(projectAssociationType, projectAssociationMapping, append(wl, projectAssociationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update project_associations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for project_associations")
	}

	if !cached {
		projectAssociationUpdateCacheMut.Lock()
		projectAssociationUpdateCache[key] = cache
		projectAssociationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q projectAssociationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for project_associations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for project_associations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProjectAssociationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectAssociationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"project_associations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, projectAssociationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in projectAssociation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all projectAssociation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProjectAssociation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no project_associations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(projectAssociationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	projectAssociationUpsertCacheMut.RLock()
	cache, cached := projectAssociationUpsertCache[key]
	projectAssociationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			projectAssociationAllColumns,
			projectAssociationColumnsWithDefault,
			projectAssociationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			projectAssociationAllColumns,
			projectAssociationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert project_associations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(projectAssociationPrimaryKeyColumns))
			copy(conflict, projectAssociationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"project_associations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(projectAssociationType, projectAssociationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(projectAssociationType, projectAssociationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert project_associations")
	}

	if !cached {
		projectAssociationUpsertCacheMut.Lock()
		projectAssociationUpsertCache[key] = cache
		projectAssociationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProjectAssociation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProjectAssociation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProjectAssociation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), projectAssociationPrimaryKeyMapping)
	sql := "DELETE FROM \"project_associations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from project_associations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for project_associations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q projectAssociationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no projectAssociationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from project_associations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project_associations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProjectAssociationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(projectAssociationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectAssociationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"project_associations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectAssociationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from projectAssociation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for project_associations")
	}

	if len(projectAssociationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ProjectAssociation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProjectAssociation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProjectAssociationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProjectAssociationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), projectAssociationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"project_associations\".* FROM \"project_associations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, projectAssociationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProjectAssociationSlice")
	}

	*o = slice

	return nil
}

// ProjectAssociationExists checks if the ProjectAssociation row exists.
func ProjectAssociationExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"project_associations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if project_associations exists")
	}

	return exists, nil
}
