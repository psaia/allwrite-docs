# For development...
go install && \
	STORAGE=postgres \
	ACTIVE_DIR=0B4pmjFk2yyz2NUtHRzVUS1RBMVk \
	PG_USER=petesaia \
	PORT=":8000" \
	PG_DB=allwrite \
	PG_HOST=localhost \
	allwrite-docs
