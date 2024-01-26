#!/bin/sh
go install github.com/volatiletech/sqlboiler/v4@v4.14.2
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.2
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@v4.14.2
# sqlboiler mysql
./sqlboiler mysql
