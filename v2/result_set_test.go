package go_ora

import "testing"

// TestResultSetColumnsNilSafe verifies that Columns() does not panic when the
// ResultSet is uninitialized (cols == nil). This guards against the nil-pointer
// panic reported in issue #691, where a connection reset during _query() could
// yield a DataSet with an uninitialized ResultSet.
func TestResultSetColumnsNilSafe(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Columns() panicked on nil cols: %v", r)
		}
	}()
	rs := ResultSet{}
	if got := rs.Columns(); got != nil {
		t.Errorf("expected nil columns for uninitialized ResultSet, got %v", got)
	}
}
