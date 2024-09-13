package queryresult

import "github.com/turbot/pipe-fittings/queryresult"

// Result is a type alias for queryresult.Result[TimingResultStream]
type Result = queryresult.Result[TimingResultStream]

func NewResult(cols []*queryresult.ColumnDef) *Result {
	return queryresult.NewResult[TimingResultStream](cols, make(TimingResultStream, 1))
}

// SyncQueryResult is a type alias for queryresult.SyncQueryResult[TimingResult]
type SyncQueryResult = queryresult.SyncQueryResult[*TimingResult]

// ResultStreamer is a type alias for queryresult.ResultStreamer[TimingResultStream]
type ResultStreamer = queryresult.ResultStreamer[TimingResultStream]

func NewResultStreamer() *ResultStreamer {
	return queryresult.NewResultStreamer[TimingResultStream]()
}
