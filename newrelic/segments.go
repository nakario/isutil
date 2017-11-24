package newrelic

import "github.com/newrelic/go-agent"

const (
	INSERT = "INSERT"
	UPDATE = "UPDATE"
	SELECT = "SELECT"
	DELETE = "DELETE"
)

func StartMySQLInsertSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreMySQL,
		Operation: INSERT,
		Collection: collection,
	}
}

func StartMySQLUpdateSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreMySQL,
		Operation: UPDATE,
		Collection: collection,
	}
}

func StartMySQLSelectSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreMySQL,
		Operation: SELECT,
		Collection: collection,
	}
}

func StartMySQLDeleteSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreMySQL,
		Operation: DELETE,
		Collection: collection,
	}
}

func StartPostgresInsertSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastorePostgres,
		Operation: INSERT,
		Collection: collection,
	}
}

func StartPostgresUpdateSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastorePostgres,
		Operation: UPDATE,
		Collection: collection,
	}
}

func StartPostgresSelectSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastorePostgres,
		Operation: SELECT,
		Collection: collection,
	}
}

func StartPostgresDeleteSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastorePostgres,
		Operation: DELETE,
		Collection: collection,
	}
}

func StartSQLiteInsertSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreSQLite,
		Operation: INSERT,
		Collection: collection,
	}
}

func StartSQLiteUpdateSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreSQLite,
		Operation: UPDATE,
		Collection: collection,
	}
}

func StartSQLiteSelectSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreSQLite,
		Operation: SELECT,
		Collection: collection,
	}
}

func StartSQLiteDeleteSegment(txn newrelic.Transaction, collection string) newrelic.DatastoreSegment {
	return newrelic.DatastoreSegment{
		StartTime: txn.StartSegmentNow(),
		Product: newrelic.DatastoreSQLite,
		Operation: DELETE,
		Collection: collection,
	}
}
