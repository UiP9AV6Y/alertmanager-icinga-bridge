#!/bin/sh
set -e

wget -q -O- \
  "https://raw.githubusercontent.com/Icinga/icingaweb2/refs/tags/v${ICINGAWEB_VERSION}/schema/pgsql.schema.sql" \
| psql -v ON_ERROR_STOP=1 --username "$ICINGAWEB_USER" --dbname "$ICINGAWEB_DB"
